// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package okms

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a KMS service key, in the JWK format.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Okms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Okms.GetOkmsServiceKey(ctx, &okms.GetOkmsServiceKeyArgs{
//				Id:     "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
//				OkmsId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOkmsServiceKeyJwk(ctx *pulumi.Context, args *GetOkmsServiceKeyJwkArgs, opts ...pulumi.InvokeOption) (*GetOkmsServiceKeyJwkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetOkmsServiceKeyJwkResult
	err := ctx.Invoke("ovh:Okms/getOkmsServiceKeyJwk:getOkmsServiceKeyJwk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOkmsServiceKeyJwk.
type GetOkmsServiceKeyJwkArgs struct {
	// ID of the service key
	Id string `pulumi:"id"`
	// ID of the KMS
	OkmsId string `pulumi:"okmsId"`
}

// A collection of values returned by getOkmsServiceKeyJwk.
type GetOkmsServiceKeyJwkResult struct {
	CreatedAt string                    `pulumi:"createdAt"`
	Iam       GetOkmsServiceKeyJwkIam   `pulumi:"iam"`
	Id        string                    `pulumi:"id"`
	Keys      []GetOkmsServiceKeyJwkKey `pulumi:"keys"`
	Name      string                    `pulumi:"name"`
	OkmsId    string                    `pulumi:"okmsId"`
	Size      float64                   `pulumi:"size"`
	State     string                    `pulumi:"state"`
	Type      string                    `pulumi:"type"`
}

func GetOkmsServiceKeyJwkOutput(ctx *pulumi.Context, args GetOkmsServiceKeyJwkOutputArgs, opts ...pulumi.InvokeOption) GetOkmsServiceKeyJwkResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOkmsServiceKeyJwkResultOutput, error) {
			args := v.(GetOkmsServiceKeyJwkArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Okms/getOkmsServiceKeyJwk:getOkmsServiceKeyJwk", args, GetOkmsServiceKeyJwkResultOutput{}, options).(GetOkmsServiceKeyJwkResultOutput), nil
		}).(GetOkmsServiceKeyJwkResultOutput)
}

// A collection of arguments for invoking getOkmsServiceKeyJwk.
type GetOkmsServiceKeyJwkOutputArgs struct {
	// ID of the service key
	Id pulumi.StringInput `pulumi:"id"`
	// ID of the KMS
	OkmsId pulumi.StringInput `pulumi:"okmsId"`
}

func (GetOkmsServiceKeyJwkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOkmsServiceKeyJwkArgs)(nil)).Elem()
}

// A collection of values returned by getOkmsServiceKeyJwk.
type GetOkmsServiceKeyJwkResultOutput struct{ *pulumi.OutputState }

func (GetOkmsServiceKeyJwkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOkmsServiceKeyJwkResult)(nil)).Elem()
}

func (o GetOkmsServiceKeyJwkResultOutput) ToGetOkmsServiceKeyJwkResultOutput() GetOkmsServiceKeyJwkResultOutput {
	return o
}

func (o GetOkmsServiceKeyJwkResultOutput) ToGetOkmsServiceKeyJwkResultOutputWithContext(ctx context.Context) GetOkmsServiceKeyJwkResultOutput {
	return o
}

func (o GetOkmsServiceKeyJwkResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetOkmsServiceKeyJwkResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o GetOkmsServiceKeyJwkResultOutput) Iam() GetOkmsServiceKeyJwkIamOutput {
	return o.ApplyT(func(v GetOkmsServiceKeyJwkResult) GetOkmsServiceKeyJwkIam { return v.Iam }).(GetOkmsServiceKeyJwkIamOutput)
}

func (o GetOkmsServiceKeyJwkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOkmsServiceKeyJwkResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOkmsServiceKeyJwkResultOutput) Keys() GetOkmsServiceKeyJwkKeyArrayOutput {
	return o.ApplyT(func(v GetOkmsServiceKeyJwkResult) []GetOkmsServiceKeyJwkKey { return v.Keys }).(GetOkmsServiceKeyJwkKeyArrayOutput)
}

func (o GetOkmsServiceKeyJwkResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetOkmsServiceKeyJwkResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetOkmsServiceKeyJwkResultOutput) OkmsId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOkmsServiceKeyJwkResult) string { return v.OkmsId }).(pulumi.StringOutput)
}

func (o GetOkmsServiceKeyJwkResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v GetOkmsServiceKeyJwkResult) float64 { return v.Size }).(pulumi.Float64Output)
}

func (o GetOkmsServiceKeyJwkResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetOkmsServiceKeyJwkResult) string { return v.State }).(pulumi.StringOutput)
}

func (o GetOkmsServiceKeyJwkResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetOkmsServiceKeyJwkResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOkmsServiceKeyJwkResultOutput{})
}
