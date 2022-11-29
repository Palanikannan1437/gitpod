// Copyright (c) 2022 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

package oci_tool

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/containerd/containerd/remotes"
	"github.com/containerd/containerd/remotes/docker"
	"github.com/docker/distribution/reference"
	"github.com/opencontainers/go-digest"
	ociv1 "github.com/opencontainers/image-spec/specs-go/v1"
)

func Resolve(ctx context.Context, ref string) (string, error) {
	newCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	res := docker.NewResolver(docker.ResolverOptions{})

	name, desc, err := res.Resolve(newCtx, ref)
	if err != nil {
		return "", err
	}

	pref, err := reference.ParseNamed(name)
	if err != nil {
		return "", err
	}
	cref, err := reference.WithDigest(pref, desc.Digest)
	if err != nil {
		return "", err
	}
	return cref.String(), nil
}

func fetchManifest(ctx context.Context, res remotes.Resolver, ref string, dgst digest.Digest) (resolvedName string, mf *ociv1.Manifest, err error) {
	name, desc, err := res.Resolve(ctx, ref)
	if err != nil {
		return "", nil, fmt.Errorf("cannot resolve %v: %w", ref, err)
	}
	if dgst != "" {
		desc.Digest = dgst
	}
	fetcher, err := res.Fetcher(ctx, name)
	if err != nil {
		return "", nil, err
	}
	mfin, err := fetcher.Fetch(ctx, desc)
	if err != nil {
		return "", nil, err
	}
	defer mfin.Close()

	var mfo ociv1.Manifest
	err = json.NewDecoder(mfin).Decode(&mfo)
	if err != nil {
		return "", nil, fmt.Errorf("cannot decode manifest: %w", err)
	}

	return name, &mfo, nil
}

func interactiveFetchManifestOrIndex(ctx context.Context, res remotes.Resolver, ref, plt string, dgst digest.Digest) (name string, result *ociv1.Manifest, err error) {
	resolved, desc, err := res.Resolve(ctx, ref)
	if err != nil {
		return "", nil, fmt.Errorf("cannot resolve %v: %w", ref, err)
	}

	if dgst != "" {
		desc.Digest = dgst
	}

	fetcher, err := res.Fetcher(ctx, resolved)
	if err != nil {
		return "", nil, err
	}

	in, err := fetcher.Fetch(ctx, desc)
	if err != nil {
		return "", nil, err
	}
	defer in.Close()
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		return "", nil, err
	}

	var mf ociv1.Manifest
	err = json.Unmarshal(buf, &mf)
	if err != nil {
		return "", nil, fmt.Errorf("cannot unmarshal manifest: %w", err)
	}

	if mf.Config.Size != 0 {
		return resolved, &mf, nil
	}

	var mfl ociv1.Index
	err = json.Unmarshal(buf, &mfl)
	if err != nil {
		return "", nil, err
	}

	if plt != "" {
		var dgst digest.Digest
		for _, mf := range mfl.Manifests {
			if fmt.Sprintf("%s-%s", mf.Platform.OS, mf.Platform.Architecture) == plt {
				dgst = mf.Digest
			}
		}
		if dgst == "" {
			return "", nil, fmt.Errorf("no manifest for platform %s found", plt)
		}

		fmt.Fprintf(os.Stderr, "found manifest for %s: %s\n", plt, dgst)

		var mf *ociv1.Manifest
		_, mf, err = fetchManifest(ctx, res, resolved, dgst)
		if err != nil {
			return "", nil, err
		}

		return resolved, mf, nil
	}

	fmt.Fprintf(os.Stderr, "%s points to an index rather than a manifest.\n", ref)
	fmt.Fprintf(os.Stderr, "Use --platform to select a manifest. Possible choices are:\n")
	for _, mf := range mfl.Manifests {
		fmt.Fprintf(os.Stderr, "\t%s-%s\n", mf.Platform.OS, mf.Platform.Architecture)
	}

	os.Exit(2)
	return "", nil, nil
}

func ResolveIDEVersion(ctx context.Context, ref string, dgst digest.Digest) (string, error) {
	newCtx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	res := docker.NewResolver(docker.ResolverOptions{})

	name, mf, err := interactiveFetchManifestOrIndex(newCtx, res, ref, "", dgst)
	if err != nil {
		return "", err
	}

	fetcher, err := res.Fetcher(ctx, name)
	if err != nil {
		return "", err
	}

	cfgin, err := fetcher.Fetch(ctx, mf.Config)
	if err != nil {
		return "", err
	}
	defer cfgin.Close()

	var ctn map[string]interface{}

	err = json.NewDecoder(cfgin).Decode(&ctn)
	if err != nil {
		return "", err
	}

	version := ctn["config"].(map[string]interface{})["Labels"].(map[string]interface{})["io.gitpod.ide.version"].(string)

	return version, nil
}
