// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hosting

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to retrieve information about an hosting database.
//
// ## Example Usage
func LookupPrivateDatabase(ctx *pulumi.Context, args *LookupPrivateDatabaseArgs, opts ...pulumi.InvokeOption) (*LookupPrivateDatabaseResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateDatabaseResult
	err := ctx.Invoke("ovh:Hosting/getPrivateDatabase:getPrivateDatabase", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivateDatabase.
type LookupPrivateDatabaseArgs struct {
	// The internal name of your private database
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getPrivateDatabase.
type LookupPrivateDatabaseResult struct {
	// Number of CPU on your private database
	Cpu int `pulumi:"cpu"`
	// Datacenter where this private database is located
	Datacenter string `pulumi:"datacenter"`
	// Name displayed in customer panel for your private database
	DisplayName string `pulumi:"displayName"`
	// Private database hostname
	Hostname string `pulumi:"hostname"`
	// Private database FTP hostname
	HostnameFtp string `pulumi:"hostnameFtp"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Infrastructure where service was stored
	Infrastructure string `pulumi:"infrastructure"`
	// Type of the private database offer
	Offer string `pulumi:"offer"`
	// Private database service port
	Port int `pulumi:"port"`
	// Private database FTP port
	PortFtp int `pulumi:"portFtp"`
	// Space allowed (in MB) on your private database
	QuotaSize int `pulumi:"quotaSize"`
	// Sapce used (in MB) on your private database
	QuotaUsed int `pulumi:"quotaUsed"`
	// Amount of ram (in MB) on your private database
	Ram int `pulumi:"ram"`
	// Private database server name
	Server      string `pulumi:"server"`
	ServiceName string `pulumi:"serviceName"`
	// Private database state
	State string `pulumi:"state"`
	Type  string `pulumi:"type"`
	// URN of the private database
	Urn string `pulumi:"urn"`
	// Private database available versions
	Version string `pulumi:"version"`
	// Private database version label
	VersionLabel string `pulumi:"versionLabel"`
	// Private database version number
	VersionNumber float64 `pulumi:"versionNumber"`
}

func LookupPrivateDatabaseOutput(ctx *pulumi.Context, args LookupPrivateDatabaseOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateDatabaseResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateDatabaseResult, error) {
			args := v.(LookupPrivateDatabaseArgs)
			r, err := LookupPrivateDatabase(ctx, &args, opts...)
			var s LookupPrivateDatabaseResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateDatabaseResultOutput)
}

// A collection of arguments for invoking getPrivateDatabase.
type LookupPrivateDatabaseOutputArgs struct {
	// The internal name of your private database
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupPrivateDatabaseOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDatabaseArgs)(nil)).Elem()
}

// A collection of values returned by getPrivateDatabase.
type LookupPrivateDatabaseResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateDatabaseResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDatabaseResult)(nil)).Elem()
}

func (o LookupPrivateDatabaseResultOutput) ToLookupPrivateDatabaseResultOutput() LookupPrivateDatabaseResultOutput {
	return o
}

func (o LookupPrivateDatabaseResultOutput) ToLookupPrivateDatabaseResultOutputWithContext(ctx context.Context) LookupPrivateDatabaseResultOutput {
	return o
}

func (o LookupPrivateDatabaseResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrivateDatabaseResult] {
	return pulumix.Output[LookupPrivateDatabaseResult]{
		OutputState: o.OutputState,
	}
}

// Number of CPU on your private database
func (o LookupPrivateDatabaseResultOutput) Cpu() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) int { return v.Cpu }).(pulumi.IntOutput)
}

// Datacenter where this private database is located
func (o LookupPrivateDatabaseResultOutput) Datacenter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.Datacenter }).(pulumi.StringOutput)
}

// Name displayed in customer panel for your private database
func (o LookupPrivateDatabaseResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Private database hostname
func (o LookupPrivateDatabaseResultOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.Hostname }).(pulumi.StringOutput)
}

// Private database FTP hostname
func (o LookupPrivateDatabaseResultOutput) HostnameFtp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.HostnameFtp }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPrivateDatabaseResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.Id }).(pulumi.StringOutput)
}

// Infrastructure where service was stored
func (o LookupPrivateDatabaseResultOutput) Infrastructure() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.Infrastructure }).(pulumi.StringOutput)
}

// Type of the private database offer
func (o LookupPrivateDatabaseResultOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.Offer }).(pulumi.StringOutput)
}

// Private database service port
func (o LookupPrivateDatabaseResultOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) int { return v.Port }).(pulumi.IntOutput)
}

// Private database FTP port
func (o LookupPrivateDatabaseResultOutput) PortFtp() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) int { return v.PortFtp }).(pulumi.IntOutput)
}

// Space allowed (in MB) on your private database
func (o LookupPrivateDatabaseResultOutput) QuotaSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) int { return v.QuotaSize }).(pulumi.IntOutput)
}

// Sapce used (in MB) on your private database
func (o LookupPrivateDatabaseResultOutput) QuotaUsed() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) int { return v.QuotaUsed }).(pulumi.IntOutput)
}

// Amount of ram (in MB) on your private database
func (o LookupPrivateDatabaseResultOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) int { return v.Ram }).(pulumi.IntOutput)
}

// Private database server name
func (o LookupPrivateDatabaseResultOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.Server }).(pulumi.StringOutput)
}

func (o LookupPrivateDatabaseResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Private database state
func (o LookupPrivateDatabaseResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupPrivateDatabaseResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.Type }).(pulumi.StringOutput)
}

// URN of the private database
func (o LookupPrivateDatabaseResultOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.Urn }).(pulumi.StringOutput)
}

// Private database available versions
func (o LookupPrivateDatabaseResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.Version }).(pulumi.StringOutput)
}

// Private database version label
func (o LookupPrivateDatabaseResultOutput) VersionLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) string { return v.VersionLabel }).(pulumi.StringOutput)
}

// Private database version number
func (o LookupPrivateDatabaseResultOutput) VersionNumber() pulumi.Float64Output {
	return o.ApplyT(func(v LookupPrivateDatabaseResult) float64 { return v.VersionNumber }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateDatabaseResultOutput{})
}
