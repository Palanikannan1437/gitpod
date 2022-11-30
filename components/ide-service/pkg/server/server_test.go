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

	ctesting "github.com/gitpod-io/gitpod/common-go/testing"
)

func TestResolveWorkspaceConfig(t *testing.T) {
	type fixture struct {
		api.ResolveWorkspaceConfigRequest
	}
	type gold struct {
		Resp *api.ResolveWorkspaceConfigResponse
		Err  string
	}

	cfg := &config.ServiceConfiguration{
		Server:        &baseserver.Configuration{},
		IDEConfigPath: "../../example-ide-config.json",
	}
	server := New(cfg)
	server.readIDEConfig(context.Background(), true)

	test := ctesting.FixtureTest{
		T:    t,
		Path: "testdata/resolve_ws_config_*.json",
		Test: func(t *testing.T, input any) interface{} {
			fixture := input.(*fixture)

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := server.ResolveWorkspaceConfig(ctx, &fixture.ResolveWorkspaceConfigRequest)
			if err != nil {
				return &gold{Err: err.Error()}
			}
			return &gold{Resp: resp, Err: ""}
		},
		Fixture: func() interface{} { return &fixture{} },
		Gold:    func() interface{} { return &gold{} },
	}
	test.Run()
}

// TODO: migrate to fixture testing
// describe("chooseIDE", async function () {
// 	const baseOpt: IDEOption = {
// 		title: "title",
// 		type: "desktop",
// 		logo: "",
// 		image: "image",
// 		latestImage: "latestImage",
// 	};
// 	const ideOptions: IDEOptions = {
// 		options: {
// 			code: Object.assign({}, baseOpt, { type: "browser" }),
// 			goland: Object.assign({}, baseOpt),
// 			"code-desktop": Object.assign({}, baseOpt),
// 			"no-latest": Object.assign({}, baseOpt),
// 		},
// 		defaultIde: "code",
// 		defaultDesktopIde: "code-desktop",
// 	};
// 	delete ideOptions.options["no-latest"].latestImage;

// 	it("code with latest", function () {
// 		const useLatest = true;
// 		const hasPerm = false;
// 		const result = chooseIDE("code", ideOptions, useLatest, hasPerm);
// 		expect(result.ideImage).to.equal(ideOptions.options["code"].latestImage);
// 	});

// 	it("code without latest", function () {
// 		const useLatest = false;
// 		const hasPerm = false;
// 		const result = chooseIDE("code", ideOptions, useLatest, hasPerm);
// 		expect(result.ideImage).to.equal(ideOptions.options["code"].image);
// 	});

// 	it("desktop ide with latest", function () {
// 		const useLatest = true;
// 		const hasPerm = false;
// 		const result = chooseIDE("code-desktop", ideOptions, useLatest, hasPerm);
// 		expect(result.ideImage).to.equal(ideOptions.options["code"].latestImage);
// 		expect(result.desktopIdeImage).to.equal(ideOptions.options["code-desktop"].latestImage);
// 	});

// 	it("desktop ide (JetBrains) without latest", function () {
// 		const useLatest = false;
// 		const hasPerm = false;
// 		const result = chooseIDE("goland", ideOptions, useLatest, hasPerm);
// 		expect(result.ideImage).to.equal(ideOptions.options["code"].image);
// 		expect(result.desktopIdeImage).to.equal(ideOptions.options["goland"].image);
// 	});

// 	it("desktop ide with no latest image", function () {
// 		const useLatest = true;
// 		const hasPerm = false;
// 		const result = chooseIDE("no-latest", ideOptions, useLatest, hasPerm);
// 		expect(result.ideImage).to.equal(ideOptions.options["code"].latestImage);
// 		expect(result.desktopIdeImage).to.equal(ideOptions.options["no-latest"].image);
// 	});

// 	it("unknown ide with custom permission should be unknown", function () {
// 		const customOptions = Object.assign({}, ideOptions);
// 		customOptions.options["unknown-custom"] = {
// 			title: "unknown title",
// 			type: "browser",
// 			logo: "",
// 			image: "",
// 		};
// 		const useLatest = true;
// 		const hasPerm = true;
// 		const result = chooseIDE("unknown-custom", customOptions, useLatest, hasPerm);
// 		expect(result.ideImage).to.equal("unknown-custom");
// 	});

// 	it("unknown desktop ide with custom permission desktop should be unknown", function () {
// 		const customOptions = Object.assign({}, ideOptions);
// 		customOptions.options["unknown-custom"] = {
// 			title: "unknown title",
// 			type: "desktop",
// 			logo: "",
// 			image: "",
// 		};
// 		const useLatest = true;
// 		const hasPerm = true;
// 		const result = chooseIDE("unknown-custom", customOptions, useLatest, hasPerm);
// 		expect(result.desktopIdeImage).to.equal("unknown-custom");
// 	});

// 	it("unknown browser ide without custom permission should fallback to code", function () {
// 		const customOptions = Object.assign({}, ideOptions);
// 		customOptions.options["unknown-custom"] = {
// 			title: "unknown title",
// 			type: "browser",
// 			logo: "",
// 			image: "",
// 		};
// 		const useLatest = true;
// 		const hasPerm = false;
// 		const result = chooseIDE("unknown-custom", customOptions, useLatest, hasPerm);
// 		expect(result.ideImage).to.equal(ideOptions.options["code"].latestImage);
// 	});

// 	it("not exists ide with custom permission", function () {
// 		const useLatest = true;
// 		const hasPerm = true;
// 		const result = chooseIDE("not-exists", ideOptions, useLatest, hasPerm);
// 		expect(result.ideImage).to.equal(ideOptions.options["code"].latestImage);
// 	});

// 	it("not exists ide with custom permission", function () {
// 		const useLatest = true;
// 		const hasPerm = false;
// 		const result = chooseIDE("not-exists", ideOptions, useLatest, hasPerm);
// 		expect(result.ideImage).to.equal(ideOptions.options["code"].latestImage);
// 	});
// });
