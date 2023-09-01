// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

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
	case "ovh:IpLoadBalancing/httpFarm:HttpFarm":
		r = &HttpFarm{}
	case "ovh:IpLoadBalancing/httpFarmServer:HttpFarmServer":
		r = &HttpFarmServer{}
	case "ovh:IpLoadBalancing/httpFrontend:HttpFrontend":
		r = &HttpFrontend{}
	case "ovh:IpLoadBalancing/httpRoute:HttpRoute":
		r = &HttpRoute{}
	case "ovh:IpLoadBalancing/httpRouteRule:HttpRouteRule":
		r = &HttpRouteRule{}
	case "ovh:IpLoadBalancing/loadBalancer:LoadBalancer":
		r = &LoadBalancer{}
	case "ovh:IpLoadBalancing/refresh:Refresh":
		r = &Refresh{}
	case "ovh:IpLoadBalancing/tcpFarm:TcpFarm":
		r = &TcpFarm{}
	case "ovh:IpLoadBalancing/tcpFarmServer:TcpFarmServer":
		r = &TcpFarmServer{}
	case "ovh:IpLoadBalancing/tcpFrontend:TcpFrontend":
		r = &TcpFrontend{}
	case "ovh:IpLoadBalancing/tcpRoute:TcpRoute":
		r = &TcpRoute{}
	case "ovh:IpLoadBalancing/tcpRouteRule:TcpRouteRule":
		r = &TcpRouteRule{}
	case "ovh:IpLoadBalancing/vrackNetwork:VrackNetwork":
		r = &VrackNetwork{}
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
		"IpLoadBalancing/httpFarm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/httpFarmServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/httpFrontend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/httpRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/httpRouteRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/loadBalancer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/refresh",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/tcpFarm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/tcpFarmServer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/tcpFrontend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/tcpRoute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/tcpRouteRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"ovh",
		"IpLoadBalancing/vrackNetwork",
		&module{version},
	)
}
