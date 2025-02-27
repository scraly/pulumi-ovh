// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "ovh:Dbaas/logsCluster:LogsCluster":
		r = &LogsCluster{}
	case "ovh:Dbaas/logsInput:LogsInput":
		r = &LogsInput{}
	case "ovh:Dbaas/logsOutputGraylogStream:LogsOutputGraylogStream":
		r = &LogsOutputGraylogStream{}
	case "ovh:Dbaas/logsOutputOpenSearchAlias:LogsOutputOpenSearchAlias":
		r = &LogsOutputOpenSearchAlias{}
	case "ovh:Dbaas/logsOutputOpenSearchIndex:LogsOutputOpenSearchIndex":
		r = &LogsOutputOpenSearchIndex{}
	case "ovh:Dbaas/logsToken:LogsToken":
		r = &LogsToken{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"ovh",
		"Dbaas/logsCluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Dbaas/logsInput",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Dbaas/logsOutputGraylogStream",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Dbaas/logsOutputOpenSearchAlias",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Dbaas/logsOutputOpenSearchIndex",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Dbaas/logsToken",
		&module{version},
	)
}
