// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package server

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gitpod-io/gitpod/common-go/baseserver"
	"github.com/gitpod-io/gitpod/common-go/log"
	"github.com/gitpod-io/gitpod/common-go/watch"
	api "github.com/gitpod-io/gitpod/ide-service-api"
	"github.com/gitpod-io/gitpod/ide-service-api/config"
	"github.com/heptiolabs/healthcheck"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type IDEServiceServer struct {
	config                 *config.ServiceConfiguration
	originIDEConfig        []byte
	parsedIDEConfigContent string
	ideConfig              *config.IDEConfig
	ideConfigFileName      string

	api.UnimplementedIDEServiceServer
}

func Start(logger *logrus.Entry, cfg *config.ServiceConfiguration) error {

	ctx := context.Background()
	logger.WithField("config", cfg).Info("Starting ide-service.")

	registry := prometheus.NewRegistry()
	health := healthcheck.NewHandler()

	srv, err := baseserver.New("ide-service",
		baseserver.WithLogger(logger),
		baseserver.WithConfig(cfg.Server),
		baseserver.WithMetricsRegistry(registry),
		baseserver.WithHealthHandler(health),
	)
	if err != nil {
		return fmt.Errorf("failed to initialize ide-service: %w", err)
	}

	s := New(cfg)
	go s.watchIDEConfig(ctx)
	go s.scheduleUpdate(ctx)
	s.register(srv.GRPC())

	health.AddReadinessCheck("ide-service", func() error {
		if s.ideConfig == nil {
			return fmt.Errorf("ide config is not ready")
		}
		return nil
	})
	health.AddReadinessCheck("grpc-server", grpcProbe(*cfg.Server.Services.GRPC))

	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to serve ide-service: %w", err)
	}

	return nil
}

func New(cfg *config.ServiceConfiguration) *IDEServiceServer {
	fn, err := filepath.Abs(cfg.IDEConfigPath)
	if err != nil {
		log.WithField("path", cfg.IDEConfigPath).WithError(err).Fatal("cannot convert ide config path to abs path")
	}
	s := &IDEServiceServer{
		config:            cfg,
		ideConfigFileName: fn,
	}
	return s
}

func (s *IDEServiceServer) register(grpcServer *grpc.Server) {
	api.RegisterIDEServiceServer(grpcServer, s)
}

func (s *IDEServiceServer) GetConfig(ctx context.Context, req *api.GetConfigRequest) (*api.GetConfigResponse, error) {
	return &api.GetConfigResponse{
		Content: s.parsedIDEConfigContent,
	}, nil
}

func (s *IDEServiceServer) readIDEConfig(ctx context.Context, isInit bool) {
	b, err := os.ReadFile(s.ideConfigFileName)
	if err != nil {
		log.WithError(err).Warn("cannot read ide config file")
		return
	}
	if config, err := ParseConfig(ctx, b); err != nil {
		if !isInit {
			log.WithError(err).Fatal("cannot parse ide config")
		}
		log.WithError(err).Error("cannot parse ide config")
		return
	} else {
		parsedConfig, err := json.Marshal(config)
		if err != nil {
			log.WithError(err).Error("cannot marshal ide config")
			return
		}
		s.parsedIDEConfigContent = string(parsedConfig)
		s.ideConfig = config
		s.originIDEConfig = b

		log.Info("ide config updated")
	}
}

func (s *IDEServiceServer) watchIDEConfig(ctx context.Context) {
	go s.readIDEConfig(ctx, true)

	// `watch.File` only watch for create and remove event
	// so with locally debugging, we cannot watch example ide config file change
	// but in k8s, configmap change will create/remove file to replace it
	if err := watch.File(ctx, s.ideConfigFileName, func() {
		s.readIDEConfig(ctx, false)
	}); err != nil {
		log.WithError(err).Fatal("cannot start watch of ide config file")
	}
}

func (s *IDEServiceServer) scheduleUpdate(ctx context.Context) {
	t := time.NewTicker(time.Hour * 1)
	for {
		select {
		case <-t.C:
			log.Info("schedule update config")
			s.readIDEConfig(ctx, false)
		case <-ctx.Done():
			t.Stop()
			return
		}
	}
}

func grpcProbe(cfg baseserver.ServerConfiguration) func() error {
	return func() error {
		creds := insecure.NewCredentials()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(ctx, cfg.Address, grpc.WithTransportCredentials(creds))
		if err != nil {
			return err
		}
		defer conn.Close()

		client := grpc_health_v1.NewHealthClient(conn)
		check, err := client.Check(ctx, &grpc_health_v1.HealthCheckRequest{})
		if err != nil {
			return err
		}

		if check.Status == grpc_health_v1.HealthCheckResponse_SERVING {
			return nil
		}
		return fmt.Errorf("grpc service not ready")
	}
}

type IDESettings struct {
	DefaultIde       string `json:"defaultIde,omitempty"`
	UseLatestVersion bool   `json:"useLatestVersion,omitempty"`
}

type WorkspaceContext struct {
	Referrer    string `json:"referrer,omitempty"`
	ReferrerIde string `json:"referrerIde,omitempty"`
}

type JetBrainsPrebuilds struct {
	Version string `json:"version,omitempty"`
}

type JetBrainsProductConfig struct {
	Prebuilds JetBrainsPrebuilds `json:"prebuilds,omitempty"`
	Vmoptions string             `json:"vmoptions,omitempty"`
}

type JetBrainsConfig struct {
	Intellij *JetBrainsProductConfig `json:"intellij,omitempty"`
	Goland   *JetBrainsProductConfig `json:"goland,omitempty"`
	Pycharm  *JetBrainsProductConfig `json:"pycharm,omitempty"`
	Phpstorm *JetBrainsProductConfig `json:"phpstorm,omitempty"`
	Rubymine *JetBrainsProductConfig `json:"rubymine,omitempty"`
	Webstorm *JetBrainsProductConfig `json:"webstorm,omitempty"`
	Rider    *JetBrainsProductConfig `json:"rider,omitempty"`
	Clion    *JetBrainsProductConfig `json:"clion,omitempty"`
}

type WorkspaceConfig struct {
	Jetbrains *JetBrainsConfig `json:"jetbrains,omitempty"`
}

func (s *IDEServiceServer) resolveReferrerIDE(wsCtx WorkspaceContext, ideSettings IDESettings) (config.IDEOption, bool) {
	if wsCtx.Referrer == "" || wsCtx.ReferrerIde == "" {
		return config.IDEOption{}, false
	}

	client, ok := s.ideConfig.IdeOptions.Clients[wsCtx.Referrer]
	if !ok {
		return config.IDEOption{}, false
	}

	providedIde := wsCtx.ReferrerIde
	providedOption, ok := s.ideConfig.IdeOptions.Options[providedIde]
	if !ok {
		return config.IDEOption{}, false
	}

	hasDesktopIde := false
	for _, desktopIde := range client.DesktopIDEs {
		if desktopIde == providedIde {
			hasDesktopIde = true
		}
	}

	if hasDesktopIde {
		return providedOption, true
	}

	defaultDesktopIdeOption, ok := s.ideConfig.IdeOptions.Options[client.DefaultDesktopIDE]
	if !ok {
		return config.IDEOption{}, false
	}

	return defaultDesktopIdeOption, true
}

func (s *IDEServiceServer) ResolveStartWorkspaceSpec(ctx context.Context, req *api.ResolveStartWorkspaceSpecRequest) (resp *api.ResolveStartWorkspaceSpecResponse, err error) {
	resp = &api.ResolveStartWorkspaceSpecResponse{
		SupervisorImage: s.ideConfig.SupervisorImage,
	}

	if req.Type == api.WorkspaceType_IMAGEBUILD {
		return resp, nil
	}

	var wsConfig WorkspaceConfig
	var wsContext WorkspaceContext
	var ideSettings IDESettings

	err = json.Unmarshal([]byte(req.WorkspaceConfig), &wsConfig)
	if err != nil {
		// todo(af): define default
		log.WithError(err).Error("failed to parse workspace config")
		return resp, nil
	}

	err = json.Unmarshal([]byte(req.Context), &wsContext)
	if err != nil {
		// todo(af): define default
		log.WithError(err).Error("failed to parse context")
		return resp, nil
	}

	err = json.Unmarshal([]byte(req.WorkspaceConfig), &ideSettings)
	if err != nil {
		// todo(af): define default
		log.WithError(err).Error("failed to parse ide settings")
		return resp, nil
	}
	if req.Type == api.WorkspaceType_PREBUILD && wsConfig.Jetbrains != nil {
		// do some logic
		return resp, nil
	}

	defaultIde := s.ideConfig.IdeOptions.Options[s.ideConfig.IdeOptions.DefaultIde]

	chosenIDEName := ideSettings.DefaultIde
	chosenIDEVersion := ideSettings.UseLatestVersion

	chosenIDE := defaultIde

	getIDEImage := func(ideOption config.IDEOption, useLatest bool) string {
		if useLatest && ideOption.LatestImage != "" {
			return ideOption.LatestImage
		}

		return ideOption.Image
	}

	getPluginImage := func(ideOption config.IDEOption, useLatest bool) string {
		if useLatest && ideOption.PluginLatestImage != "" {
			return ideOption.PluginLatestImage
		}

		return ideOption.PluginImage
	}

	if chosenIDEName != "" {
		if ide, ok := s.ideConfig.IdeOptions.Options[chosenIDEName]; ok {
			chosenIDE = ide
		}
	}

	// we always need WebImage for when the user chooses a desktop ide
	resp.WebImage = getIDEImage(defaultIde, chosenIDEVersion)

	var desktopImageLayer string
	var desktopPluginImageLayer string
	if chosenIDE.Type == config.IDETypeDesktop {
		desktopImageLayer = getIDEImage(chosenIDE, chosenIDEVersion)
		desktopPluginImageLayer = getPluginImage(chosenIDE, chosenIDEVersion)

	} else {
		resp.WebImage = getIDEImage(chosenIDE, chosenIDEVersion)
	}

	referrer, ok := s.resolveReferrerIDE(wsContext, ideSettings)
	if ok {
		desktopImageLayer = getIDEImage(referrer, chosenIDEVersion)
		desktopPluginImageLayer = getIDEImage(referrer, chosenIDEVersion)
	}

	resp.IdeImageLayers = append(resp.IdeImageLayers, desktopImageLayer)

	if desktopPluginImageLayer != "" {
		resp.IdeImageLayers = append(resp.IdeImageLayers, desktopPluginImageLayer)
	}

	return resp, nil
}
