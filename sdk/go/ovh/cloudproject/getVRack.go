// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the linked vrack on your public cloud project.
//
// ## Example Usage
//
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
//			vrackVRack, err := CloudProject.GetVRack(ctx, &cloudproject.GetVRackArgs{
//				ServiceName: "XXXXXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vrack", vrackVRack)
//			return nil
//		})
//	}
//
// ```
func GetVRack(ctx *pulumi.Context, args *GetVRackArgs, opts ...pulumi.InvokeOption) (*GetVRackResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVRackResult
	err := ctx.Invoke("ovh:CloudProject/getVRack:getVRack", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVRack.
type GetVRackArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getVRack.
type GetVRackResult struct {
	Description string `pulumi:"description"`
	// The id of the vrack
	Id string `pulumi:"id"`
	// The name of the vrack
	Name string `pulumi:"name"`
	// The id of the public cloud project
	ServiceName string `pulumi:"serviceName"`
}

func GetVRackOutput(ctx *pulumi.Context, args GetVRackOutputArgs, opts ...pulumi.InvokeOption) GetVRackResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetVRackResultOutput, error) {
			args := v.(GetVRackArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getVRack:getVRack", args, GetVRackResultOutput{}, options).(GetVRackResultOutput), nil
		}).(GetVRackResultOutput)
}

// A collection of arguments for invoking getVRack.
type GetVRackOutputArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetVRackOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVRackArgs)(nil)).Elem()
}

// A collection of values returned by getVRack.
type GetVRackResultOutput struct{ *pulumi.OutputState }

func (GetVRackResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVRackResult)(nil)).Elem()
}

func (o GetVRackResultOutput) ToGetVRackResultOutput() GetVRackResultOutput {
	return o
}

func (o GetVRackResultOutput) ToGetVRackResultOutputWithContext(ctx context.Context) GetVRackResultOutput {
	return o
}

func (o GetVRackResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetVRackResult) string { return v.Description }).(pulumi.StringOutput)
}

// The id of the vrack
func (o GetVRackResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVRackResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the vrack
func (o GetVRackResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetVRackResult) string { return v.Name }).(pulumi.StringOutput)
}

// The id of the public cloud project
func (o GetVRackResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetVRackResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVRackResultOutput{})
}
