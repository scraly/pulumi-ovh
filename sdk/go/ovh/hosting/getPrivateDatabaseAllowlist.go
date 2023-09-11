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

// Use this data source to retrieve information about an hosting privatedatabase whitelist.
//
// ## Example Usage
func LookupPrivateDatabaseAllowlist(ctx *pulumi.Context, args *LookupPrivateDatabaseAllowlistArgs, opts ...pulumi.InvokeOption) (*LookupPrivateDatabaseAllowlistResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateDatabaseAllowlistResult
	err := ctx.Invoke("ovh:Hosting/getPrivateDatabaseAllowlist:getPrivateDatabaseAllowlist", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivateDatabaseAllowlist.
type LookupPrivateDatabaseAllowlistArgs struct {
	// The whitelisted IP in your instance
	Ip *string `pulumi:"ip"`
	// The internal name of your private database
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getPrivateDatabaseAllowlist.
type LookupPrivateDatabaseAllowlistResult struct {
	// Creation date of the database
	CreationDate string `pulumi:"creationDate"`
	// The provider-assigned unique ID for this managed resource.
	Id string  `pulumi:"id"`
	Ip *string `pulumi:"ip"`
	// The last update date of this whitelist
	LastUpdate string `pulumi:"lastUpdate"`
	// Custom name for your Whitelisted IP
	Name string `pulumi:"name"`
	// Authorize this IP to access service port
	Service     bool   `pulumi:"service"`
	ServiceName string `pulumi:"serviceName"`
	// Authorize this IP to access SFTP port
	Sftp bool `pulumi:"sftp"`
	// Whitelist status
	Status string `pulumi:"status"`
}

func LookupPrivateDatabaseAllowlistOutput(ctx *pulumi.Context, args LookupPrivateDatabaseAllowlistOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateDatabaseAllowlistResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateDatabaseAllowlistResult, error) {
			args := v.(LookupPrivateDatabaseAllowlistArgs)
			r, err := LookupPrivateDatabaseAllowlist(ctx, &args, opts...)
			var s LookupPrivateDatabaseAllowlistResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateDatabaseAllowlistResultOutput)
}

// A collection of arguments for invoking getPrivateDatabaseAllowlist.
type LookupPrivateDatabaseAllowlistOutputArgs struct {
	// The whitelisted IP in your instance
	Ip pulumi.StringPtrInput `pulumi:"ip"`
	// The internal name of your private database
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupPrivateDatabaseAllowlistOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDatabaseAllowlistArgs)(nil)).Elem()
}

// A collection of values returned by getPrivateDatabaseAllowlist.
type LookupPrivateDatabaseAllowlistResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateDatabaseAllowlistResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDatabaseAllowlistResult)(nil)).Elem()
}

func (o LookupPrivateDatabaseAllowlistResultOutput) ToLookupPrivateDatabaseAllowlistResultOutput() LookupPrivateDatabaseAllowlistResultOutput {
	return o
}

func (o LookupPrivateDatabaseAllowlistResultOutput) ToLookupPrivateDatabaseAllowlistResultOutputWithContext(ctx context.Context) LookupPrivateDatabaseAllowlistResultOutput {
	return o
}

func (o LookupPrivateDatabaseAllowlistResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrivateDatabaseAllowlistResult] {
	return pulumix.Output[LookupPrivateDatabaseAllowlistResult]{
		OutputState: o.OutputState,
	}
}

// Creation date of the database
func (o LookupPrivateDatabaseAllowlistResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseAllowlistResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPrivateDatabaseAllowlistResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseAllowlistResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateDatabaseAllowlistResultOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseAllowlistResult) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

// The last update date of this whitelist
func (o LookupPrivateDatabaseAllowlistResultOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseAllowlistResult) string { return v.LastUpdate }).(pulumi.StringOutput)
}

// Custom name for your Whitelisted IP
func (o LookupPrivateDatabaseAllowlistResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseAllowlistResult) string { return v.Name }).(pulumi.StringOutput)
}

// Authorize this IP to access service port
func (o LookupPrivateDatabaseAllowlistResultOutput) Service() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseAllowlistResult) bool { return v.Service }).(pulumi.BoolOutput)
}

func (o LookupPrivateDatabaseAllowlistResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseAllowlistResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Authorize this IP to access SFTP port
func (o LookupPrivateDatabaseAllowlistResultOutput) Sftp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseAllowlistResult) bool { return v.Sftp }).(pulumi.BoolOutput)
}

// Whitelist status
func (o LookupPrivateDatabaseAllowlistResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseAllowlistResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateDatabaseAllowlistResultOutput{})
}
