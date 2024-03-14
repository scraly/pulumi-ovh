// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to filter the list of container registry capabilities associated with a public cloud project to match one and only one capability.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := CloudProject.GetCapabilitiesContainerFilter(ctx, &cloudproject.GetCapabilitiesContainerFilterArgs{
//				PlanName:    "SMALL",
//				Region:      "GRA",
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
// <!--End PulumiCodeChooser -->
func GetCapabilitiesContainerFilter(ctx *pulumi.Context, args *GetCapabilitiesContainerFilterArgs, opts ...pulumi.InvokeOption) (*GetCapabilitiesContainerFilterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCapabilitiesContainerFilterResult
	err := ctx.Invoke("ovh:CloudProject/getCapabilitiesContainerFilter:getCapabilitiesContainerFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCapabilitiesContainerFilter.
type GetCapabilitiesContainerFilterArgs struct {
	// The plan name. It can be 'SMALL', 'MEDIUM' or 'LARGE'.
	PlanName string `pulumi:"planName"`
	// The region name
	Region string `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCapabilitiesContainerFilter.
type GetCapabilitiesContainerFilterResult struct {
	// Plan code from the catalog
	Code string `pulumi:"code"`
	// Plan creation date
	CreatedAt string `pulumi:"createdAt"`
	// Features of the plan
	Features []GetCapabilitiesContainerFilterFeature `pulumi:"features"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Plan name
	Name     string `pulumi:"name"`
	PlanName string `pulumi:"planName"`
	Region   string `pulumi:"region"`
	// Container registry limits
	RegistryLimits []GetCapabilitiesContainerFilterRegistryLimit `pulumi:"registryLimits"`
	ServiceName    string                                        `pulumi:"serviceName"`
	// Plan last update date
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetCapabilitiesContainerFilterOutput(ctx *pulumi.Context, args GetCapabilitiesContainerFilterOutputArgs, opts ...pulumi.InvokeOption) GetCapabilitiesContainerFilterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetCapabilitiesContainerFilterResult, error) {
			args := v.(GetCapabilitiesContainerFilterArgs)
			r, err := GetCapabilitiesContainerFilter(ctx, &args, opts...)
			var s GetCapabilitiesContainerFilterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetCapabilitiesContainerFilterResultOutput)
}

// A collection of arguments for invoking getCapabilitiesContainerFilter.
type GetCapabilitiesContainerFilterOutputArgs struct {
	// The plan name. It can be 'SMALL', 'MEDIUM' or 'LARGE'.
	PlanName pulumi.StringInput `pulumi:"planName"`
	// The region name
	Region pulumi.StringInput `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetCapabilitiesContainerFilterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesContainerFilterArgs)(nil)).Elem()
}

// A collection of values returned by getCapabilitiesContainerFilter.
type GetCapabilitiesContainerFilterResultOutput struct{ *pulumi.OutputState }

func (GetCapabilitiesContainerFilterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCapabilitiesContainerFilterResult)(nil)).Elem()
}

func (o GetCapabilitiesContainerFilterResultOutput) ToGetCapabilitiesContainerFilterResultOutput() GetCapabilitiesContainerFilterResultOutput {
	return o
}

func (o GetCapabilitiesContainerFilterResultOutput) ToGetCapabilitiesContainerFilterResultOutputWithContext(ctx context.Context) GetCapabilitiesContainerFilterResultOutput {
	return o
}

// Plan code from the catalog
func (o GetCapabilitiesContainerFilterResultOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesContainerFilterResult) string { return v.Code }).(pulumi.StringOutput)
}

// Plan creation date
func (o GetCapabilitiesContainerFilterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesContainerFilterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Features of the plan
func (o GetCapabilitiesContainerFilterResultOutput) Features() GetCapabilitiesContainerFilterFeatureArrayOutput {
	return o.ApplyT(func(v GetCapabilitiesContainerFilterResult) []GetCapabilitiesContainerFilterFeature {
		return v.Features
	}).(GetCapabilitiesContainerFilterFeatureArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCapabilitiesContainerFilterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesContainerFilterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Plan name
func (o GetCapabilitiesContainerFilterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesContainerFilterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetCapabilitiesContainerFilterResultOutput) PlanName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesContainerFilterResult) string { return v.PlanName }).(pulumi.StringOutput)
}

func (o GetCapabilitiesContainerFilterResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesContainerFilterResult) string { return v.Region }).(pulumi.StringOutput)
}

// Container registry limits
func (o GetCapabilitiesContainerFilterResultOutput) RegistryLimits() GetCapabilitiesContainerFilterRegistryLimitArrayOutput {
	return o.ApplyT(func(v GetCapabilitiesContainerFilterResult) []GetCapabilitiesContainerFilterRegistryLimit {
		return v.RegistryLimits
	}).(GetCapabilitiesContainerFilterRegistryLimitArrayOutput)
}

func (o GetCapabilitiesContainerFilterResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesContainerFilterResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Plan last update date
func (o GetCapabilitiesContainerFilterResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetCapabilitiesContainerFilterResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCapabilitiesContainerFilterResultOutput{})
}
