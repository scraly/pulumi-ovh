// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about certificates of a cluster associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProjectDatabase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			certificates, err := CloudProjectDatabase.GetCertificates(ctx, &cloudprojectdatabase.GetCertificatesArgs{
//				ServiceName: "XXX",
//				Engine:      "YYY",
//				ClusterId:   "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("certificatesCa", certificates.Ca)
//			return nil
//		})
//	}
//
// ```
func GetCertificates(ctx *pulumi.Context, args *GetCertificatesArgs, opts ...pulumi.InvokeOption) (*GetCertificatesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCertificatesResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getCertificates:getCertificates", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCertificates.
type GetCertificatesArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// The engine of the database cluster you want database information. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine string `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCertificates.
type GetCertificatesResult struct {
	// CA certificate used for the service.
	Ca string `pulumi:"ca"`
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// See Argument Reference above.
	Engine string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func GetCertificatesOutput(ctx *pulumi.Context, args GetCertificatesOutputArgs, opts ...pulumi.InvokeOption) GetCertificatesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCertificatesResultOutput, error) {
			args := v.(GetCertificatesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProjectDatabase/getCertificates:getCertificates", args, GetCertificatesResultOutput{}, options).(GetCertificatesResultOutput), nil
		}).(GetCertificatesResultOutput)
}

// A collection of arguments for invoking getCertificates.
type GetCertificatesOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The engine of the database cluster you want database information. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine pulumi.StringInput `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetCertificatesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesArgs)(nil)).Elem()
}

// A collection of values returned by getCertificates.
type GetCertificatesResultOutput struct{ *pulumi.OutputState }

func (GetCertificatesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCertificatesResult)(nil)).Elem()
}

func (o GetCertificatesResultOutput) ToGetCertificatesResultOutput() GetCertificatesResultOutput {
	return o
}

func (o GetCertificatesResultOutput) ToGetCertificatesResultOutputWithContext(ctx context.Context) GetCertificatesResultOutput {
	return o
}

// CA certificate used for the service.
func (o GetCertificatesResultOutput) Ca() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesResult) string { return v.Ca }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetCertificatesResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetCertificatesResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCertificatesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetCertificatesResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetCertificatesResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCertificatesResultOutput{})
}
