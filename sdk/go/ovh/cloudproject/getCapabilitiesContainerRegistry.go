// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the container registry capabilities of a public cloud project.
//
// ## Example Usage
func LookupCapabilitiesContainerRegistry(ctx *pulumi.Context, args *LookupCapabilitiesContainerRegistryArgs, opts ...pulumi.InvokeOption) (*LookupCapabilitiesContainerRegistryResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCapabilitiesContainerRegistryResult
	err := ctx.Invoke("ovh:CloudProject/getCapabilitiesContainerRegistry:getCapabilitiesContainerRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCapabilitiesContainerRegistry.
type LookupCapabilitiesContainerRegistryArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCapabilitiesContainerRegistry.
type LookupCapabilitiesContainerRegistryResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of container registry capability for a single region
	Results     []GetCapabilitiesContainerRegistryResult `pulumi:"results"`
	ServiceName string                                   `pulumi:"serviceName"`
}

func LookupCapabilitiesContainerRegistryOutput(ctx *pulumi.Context, args LookupCapabilitiesContainerRegistryOutputArgs, opts ...pulumi.InvokeOption) LookupCapabilitiesContainerRegistryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCapabilitiesContainerRegistryResult, error) {
			args := v.(LookupCapabilitiesContainerRegistryArgs)
			r, err := LookupCapabilitiesContainerRegistry(ctx, &args, opts...)
			var s LookupCapabilitiesContainerRegistryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCapabilitiesContainerRegistryResultOutput)
}

// A collection of arguments for invoking getCapabilitiesContainerRegistry.
type LookupCapabilitiesContainerRegistryOutputArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCapabilitiesContainerRegistryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapabilitiesContainerRegistryArgs)(nil)).Elem()
}

// A collection of values returned by getCapabilitiesContainerRegistry.
type LookupCapabilitiesContainerRegistryResultOutput struct{ *pulumi.OutputState }

func (LookupCapabilitiesContainerRegistryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCapabilitiesContainerRegistryResult)(nil)).Elem()
}

func (o LookupCapabilitiesContainerRegistryResultOutput) ToLookupCapabilitiesContainerRegistryResultOutput() LookupCapabilitiesContainerRegistryResultOutput {
	return o
}

func (o LookupCapabilitiesContainerRegistryResultOutput) ToLookupCapabilitiesContainerRegistryResultOutputWithContext(ctx context.Context) LookupCapabilitiesContainerRegistryResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCapabilitiesContainerRegistryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapabilitiesContainerRegistryResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of container registry capability for a single region
func (o LookupCapabilitiesContainerRegistryResultOutput) Results() GetCapabilitiesContainerRegistryResultArrayOutput {
	return o.ApplyT(func(v LookupCapabilitiesContainerRegistryResult) []GetCapabilitiesContainerRegistryResult {
		return v.Results
	}).(GetCapabilitiesContainerRegistryResultArrayOutput)
}

func (o LookupCapabilitiesContainerRegistryResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCapabilitiesContainerRegistryResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCapabilitiesContainerRegistryResultOutput{})
}