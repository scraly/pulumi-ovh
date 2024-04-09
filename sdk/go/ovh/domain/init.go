// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domain

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
	case "ovh:Domain/zone:Zone":
		r = &Zone{}
	case "ovh:Domain/zoneDNSSec:ZoneDNSSec":
		r = &ZoneDNSSec{}
	case "ovh:Domain/zoneRecord:ZoneRecord":
		r = &ZoneRecord{}
	case "ovh:Domain/zoneRedirection:ZoneRedirection":
		r = &ZoneRedirection{}
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
		"Domain/zone",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/zoneDNSSec",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/zoneRecord",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"Domain/zoneRedirection",
		&module{version},
	)
}
