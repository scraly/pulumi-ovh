// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package order

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to create a temporary order cart to retrieve information order cart products.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Me"
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Order"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myAccount, err := Me.GetMe(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = Order.GetCart(ctx, &order.GetCartArgs{
//				OvhSubsidiary: myAccount.OvhSubsidiary,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCart(ctx *pulumi.Context, args *GetCartArgs, opts ...pulumi.InvokeOption) (*GetCartResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCartResult
	err := ctx.Invoke("ovh:Order/getCart:getCart", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCart.
type GetCartArgs struct {
	// Assign a shopping cart to a logged in client. Values can be `true` or `false`.
	Assign *bool `pulumi:"assign"`
	// Description of your cart
	Description *string `pulumi:"description"`
	// Expiration time (format: 2006-01-02T15:04:05+00:00)
	Expire *string `pulumi:"expire"`
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
}

// A collection of values returned by getCart.
type GetCartResult struct {
	Assign *bool `pulumi:"assign"`
	// Cart identifier
	CartId      string  `pulumi:"cartId"`
	Description *string `pulumi:"description"`
	Expire      string  `pulumi:"expire"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Items of your cart
	Items         []int  `pulumi:"items"`
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
	// Indicates if the cart has already been validated
	ReadOnly bool `pulumi:"readOnly"`
}

func GetCartOutput(ctx *pulumi.Context, args GetCartOutputArgs, opts ...pulumi.InvokeOption) GetCartResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCartResultOutput, error) {
			args := v.(GetCartArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Order/getCart:getCart", args, GetCartResultOutput{}, options).(GetCartResultOutput), nil
		}).(GetCartResultOutput)
}

// A collection of arguments for invoking getCart.
type GetCartOutputArgs struct {
	// Assign a shopping cart to a logged in client. Values can be `true` or `false`.
	Assign pulumi.BoolPtrInput `pulumi:"assign"`
	// Description of your cart
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Expiration time (format: 2006-01-02T15:04:05+00:00)
	Expire pulumi.StringPtrInput `pulumi:"expire"`
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary pulumi.StringInput `pulumi:"ovhSubsidiary"`
}

func (GetCartOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCartArgs)(nil)).Elem()
}

// A collection of values returned by getCart.
type GetCartResultOutput struct{ *pulumi.OutputState }

func (GetCartResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCartResult)(nil)).Elem()
}

func (o GetCartResultOutput) ToGetCartResultOutput() GetCartResultOutput {
	return o
}

func (o GetCartResultOutput) ToGetCartResultOutputWithContext(ctx context.Context) GetCartResultOutput {
	return o
}

func (o GetCartResultOutput) Assign() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetCartResult) *bool { return v.Assign }).(pulumi.BoolPtrOutput)
}

// Cart identifier
func (o GetCartResultOutput) CartId() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartResult) string { return v.CartId }).(pulumi.StringOutput)
}

func (o GetCartResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetCartResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetCartResultOutput) Expire() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartResult) string { return v.Expire }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetCartResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartResult) string { return v.Id }).(pulumi.StringOutput)
}

// Items of your cart
func (o GetCartResultOutput) Items() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetCartResult) []int { return v.Items }).(pulumi.IntArrayOutput)
}

func (o GetCartResultOutput) OvhSubsidiary() pulumi.StringOutput {
	return o.ApplyT(func(v GetCartResult) string { return v.OvhSubsidiary }).(pulumi.StringOutput)
}

// Indicates if the cart has already been validated
func (o GetCartResultOutput) ReadOnly() pulumi.BoolOutput {
	return o.ApplyT(func(v GetCartResult) bool { return v.ReadOnly }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCartResultOutput{})
}
