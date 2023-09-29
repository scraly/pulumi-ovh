// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to retrieve information about a dedicated server associated with your OVHcloud Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ovh.GetServer(ctx, &ovh.GetServerArgs{
//				ServiceName: "XXXXXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetServer(ctx *pulumi.Context, args *GetServerArgs, opts ...pulumi.InvokeOption) (*GetServerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServerResult
	err := ctx.Invoke("ovh:index/getServer:getServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServer.
type GetServerArgs struct {
	// The serviceName of your dedicated server.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getServer.
type GetServerResult struct {
	// URN of the dedicated server instance
	ServerURN string `pulumi:"ServerURN"`
	// boot id of the server
	BootId int `pulumi:"bootId"`
	// dedicated server commercial range
	CommercialRange string `pulumi:"commercialRange"`
	// dedicated datacenter localisation (bhs1,bhs2,...)
	Datacenter string `pulumi:"datacenter"`
	// List of enabled public VNI uuids
	EnabledPublicVnis []string `pulumi:"enabledPublicVnis"`
	// List of enabled vrackAggregation VNI uuids
	EnabledVrackAggregationVnis []string `pulumi:"enabledVrackAggregationVnis"`
	// List of enabled vrack VNI uuids
	EnabledVrackVnis []string `pulumi:"enabledVrackVnis"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// dedicated server ip (IPv4)
	Ip string `pulumi:"ip"`
	// dedicated server ip blocks
	Ips []string `pulumi:"ips"`
	// link speed of the server
	LinkSpeed int `pulumi:"linkSpeed"`
	// Icmp monitoring state
	Monitoring bool `pulumi:"monitoring"`
	// User defined VirtualNetworkInterface name
	Name string `pulumi:"name"`
	// Operating system
	Os string `pulumi:"os"`
	// Does this server have professional use option
	ProfessionalUse bool `pulumi:"professionalUse"`
	// rack id of the server
	Rack string `pulumi:"rack"`
	// rescue mail of the server
	RescueMail string `pulumi:"rescueMail"`
	// dedicated server reverse
	Reverse string `pulumi:"reverse"`
	// root device of the server
	RootDevice string `pulumi:"rootDevice"`
	// your server id
	ServerId    int    `pulumi:"serverId"`
	ServiceName string `pulumi:"serviceName"`
	// error, hacked, hackedBlocked, ok
	State string `pulumi:"state"`
	// Dedicated server support level (critical, fastpath, gs, pro)
	SupportLevel string `pulumi:"supportLevel"`
	// the list of Virtualnetworkinterface assiociated with this server
	Vnis []GetServerVni `pulumi:"vnis"`
}

func GetServerOutput(ctx *pulumi.Context, args GetServerOutputArgs, opts ...pulumi.InvokeOption) GetServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServerResult, error) {
			args := v.(GetServerArgs)
			r, err := GetServer(ctx, &args, opts...)
			var s GetServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServerResultOutput)
}

// A collection of arguments for invoking getServer.
type GetServerOutputArgs struct {
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerArgs)(nil)).Elem()
}

// A collection of values returned by getServer.
type GetServerResultOutput struct{ *pulumi.OutputState }

func (GetServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerResult)(nil)).Elem()
}

func (o GetServerResultOutput) ToGetServerResultOutput() GetServerResultOutput {
	return o
}

func (o GetServerResultOutput) ToGetServerResultOutputWithContext(ctx context.Context) GetServerResultOutput {
	return o
}

func (o GetServerResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetServerResult] {
	return pulumix.Output[GetServerResult]{
		OutputState: o.OutputState,
	}
}

// URN of the dedicated server instance
func (o GetServerResultOutput) ServerURN() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.ServerURN }).(pulumi.StringOutput)
}

// boot id of the server
func (o GetServerResultOutput) BootId() pulumi.IntOutput {
	return o.ApplyT(func(v GetServerResult) int { return v.BootId }).(pulumi.IntOutput)
}

// dedicated server commercial range
func (o GetServerResultOutput) CommercialRange() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.CommercialRange }).(pulumi.StringOutput)
}

// dedicated datacenter localisation (bhs1,bhs2,...)
func (o GetServerResultOutput) Datacenter() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.Datacenter }).(pulumi.StringOutput)
}

// List of enabled public VNI uuids
func (o GetServerResultOutput) EnabledPublicVnis() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerResult) []string { return v.EnabledPublicVnis }).(pulumi.StringArrayOutput)
}

// List of enabled vrackAggregation VNI uuids
func (o GetServerResultOutput) EnabledVrackAggregationVnis() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerResult) []string { return v.EnabledVrackAggregationVnis }).(pulumi.StringArrayOutput)
}

// List of enabled vrack VNI uuids
func (o GetServerResultOutput) EnabledVrackVnis() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerResult) []string { return v.EnabledVrackVnis }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.Id }).(pulumi.StringOutput)
}

// dedicated server ip (IPv4)
func (o GetServerResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.Ip }).(pulumi.StringOutput)
}

// dedicated server ip blocks
func (o GetServerResultOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetServerResult) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

// link speed of the server
func (o GetServerResultOutput) LinkSpeed() pulumi.IntOutput {
	return o.ApplyT(func(v GetServerResult) int { return v.LinkSpeed }).(pulumi.IntOutput)
}

// Icmp monitoring state
func (o GetServerResultOutput) Monitoring() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServerResult) bool { return v.Monitoring }).(pulumi.BoolOutput)
}

// User defined VirtualNetworkInterface name
func (o GetServerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.Name }).(pulumi.StringOutput)
}

// Operating system
func (o GetServerResultOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.Os }).(pulumi.StringOutput)
}

// Does this server have professional use option
func (o GetServerResultOutput) ProfessionalUse() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServerResult) bool { return v.ProfessionalUse }).(pulumi.BoolOutput)
}

// rack id of the server
func (o GetServerResultOutput) Rack() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.Rack }).(pulumi.StringOutput)
}

// rescue mail of the server
func (o GetServerResultOutput) RescueMail() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.RescueMail }).(pulumi.StringOutput)
}

// dedicated server reverse
func (o GetServerResultOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.Reverse }).(pulumi.StringOutput)
}

// root device of the server
func (o GetServerResultOutput) RootDevice() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.RootDevice }).(pulumi.StringOutput)
}

// your server id
func (o GetServerResultOutput) ServerId() pulumi.IntOutput {
	return o.ApplyT(func(v GetServerResult) int { return v.ServerId }).(pulumi.IntOutput)
}

func (o GetServerResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// error, hacked, hackedBlocked, ok
func (o GetServerResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.State }).(pulumi.StringOutput)
}

// Dedicated server support level (critical, fastpath, gs, pro)
func (o GetServerResultOutput) SupportLevel() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerResult) string { return v.SupportLevel }).(pulumi.StringOutput)
}

// the list of Virtualnetworkinterface assiociated with this server
func (o GetServerResultOutput) Vnis() GetServerVniArrayOutput {
	return o.ApplyT(func(v GetServerResult) []GetServerVni { return v.Vnis }).(GetServerVniArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerResultOutput{})
}
