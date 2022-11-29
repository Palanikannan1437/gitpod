// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package server

import (
	"context"
	"testing"
	"time"

	"github.com/gitpod-io/gitpod/common-go/baseserver"
	api "github.com/gitpod-io/gitpod/ide-service-api"
	"github.com/gitpod-io/gitpod/ide-service-api/config"
)

func TestResolveStartWorkspaceSpec(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cfg := &config.ServiceConfiguration{
		Server:        &baseserver.Configuration{},
		IDEConfigPath: "",
	}
	server := New(cfg)
	server.ResolveStartWorkspaceSpec(ctx, &api.ResolveStartWorkspaceSpecRequest{})
}
