// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package squadcast

import (
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	squadcast "github.com/squadcast/terraform-provider-squadcast/shim"
	"github.com/valisinsights/pulumi-squadcast/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "squadcast"
	// modules:
	mainMod = "index" // the squadcast module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	sp := squadcast.New("1.4.4")
	p := shimv2.NewProvider(sp())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "squadcast",
		DisplayName: "Squadcast",
		Publisher:   "valisinsights",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing squadcast cloud resources.",
		Keywords:          []string{"pulumi", "squadcast", "category/cloud"},

		License:  "Apache-2.0",
		Homepage: "https://www.pulumi.com",

		Repository: "https://github.com/valisinsights/pulumi-squadcast",
		GitHubOrg:  "squadcast", // should match the TF provider module's require directive, not any replace directives.

		Config:               map[string]*tfbridge.SchemaInfo{},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"squadcast_deduplication_rules":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "DeduplicationRules")},
			"squadcast_escalation_policy":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "EscalationPolicy")},
			"squadcast_routing_rules":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "RoutingRules")},
			"squadcast_runbook":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Runbook")},
			"squadcast_schedule":             {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Schedule")},
			"squadcast_schedule_v2":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ScheduleV2")},
			"squadcast_schedule_rotation_v2": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ScheduleRotationV2")},
			"squadcast_service_maintenance":  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "ServiceMaintenance")},
			"squadcast_service":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Service")},
			"squadcast_squad":                {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Squad")},
			"squadcast_suppression_rules":    {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SuppressionRules")},
			"squadcast_tagging_rules":        {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TaggingRules")},
			"squadcast_team_member":          {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TeamMember")},
			"squadcast_team_role":            {Tok: tfbridge.MakeResource(mainPkg, mainMod, "TeamRole")},
			"squadcast_team":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Team")},
			"squadcast_user":                 {Tok: tfbridge.MakeResource(mainPkg, mainMod, "User")},
			"squadcast_slo":                  {Tok: tfbridge.MakeResource(mainPkg, mainMod, "SLO")},
			"squadcast_webform":              {Tok: tfbridge.MakeResource(mainPkg, mainMod, "Sebform")},
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: tfbridge.MakeResource(mainPkg, mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: tfbridge.MakeResource(mainPkg, mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: tfbridge.MakeType(mainPkg, "Tags")},
			// 	},
			// },
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getAmi")},
			"squadcast_squad":             {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSquad")},
			"squadcast_service":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getService")},
			"squadcast_escalation_policy": {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getEscalationPolicy")},
			"squadcast_team":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTeam")},
			"squadcast_team_role":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getTeamRole")},
			"squadcast_user":              {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getUser")},
			"squadcast_schedule":          {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getSchedule")},
			"squadcast_schedule_v2":       {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getScheduleV2")},
			"squadcast_runbook":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getRunbook")},
			"squadcast_webform":           {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getWebform")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/valisinsights/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// These are new API's that you may opt to use to automatically compute resource tokens,
	// and apply auto aliasing for full backwards compatibility.
	// See https://pkg.go.dev/github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge#ProviderInfo.ComputeTokens
	prov.MustComputeTokens(tokens.SingleModule("squadcast_", mainMod,
		tokens.MakeStandard(mainPkg)))
	// TODO: Why is this not available?
	// prov.MustApplyAutoAliasing()
	prov.SetAutonaming(255, "-")

	return prov
}
