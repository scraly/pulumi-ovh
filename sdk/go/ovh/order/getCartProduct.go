// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package order

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information of order cart product products.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/me"
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/order"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myAccount, err := me.GetMe(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			myCart, err := order.GetCart(ctx, &order.GetCartArgs{
//				OvhSubsidiary: myAccount.OvhSubsidiary,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = order.GetCartProduct(ctx, &order.GetCartProductArgs{
//				CartId:  myCart.Id,
//				Product: "...",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCartProduct(ctx *pulumi.Context, args *LookupCartProductArgs, opts ...pulumi.InvokeOption) (*LookupCartProductResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCartProductResult
	err := ctx.Invoke("ovh:Order/getCartProduct:getCartProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCartProduct.
type LookupCartProductArgs struct {
	// Cart identifier
	CartId string `pulumi:"cartId"`
	// product
	Product string `pulumi:"product"`
}

// A collection of values returned by getCartProduct.
type LookupCartProductResult struct {
	CartId string `pulumi:"cartId"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Product string `pulumi:"product"`
	// products results
	Results []GetCartProductResult `pulumi:"results"`
}

func LookupCartProductOutput(ctx *pulumi.Context, args LookupCartProductOutputArgs, opts ...pulumi.InvokeOption) LookupCartProductResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCartProductResultOutput, error) {
			args := v.(LookupCartProductArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Order/getCartProduct:getCartProduct", args, LookupCartProductResultOutput{}, options).(LookupCartProductResultOutput), nil
		}).(LookupCartProductResultOutput)
}

// A collection of arguments for invoking getCartProduct.
type LookupCartProductOutputArgs struct {
	// Cart identifier
	CartId pulumi.StringInput `pulumi:"cartId"`
	// product
	Product pulumi.StringInput `pulumi:"product"`
}

func (LookupCartProductOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCartProductArgs)(nil)).Elem()
}

// A collection of values returned by getCartProduct.
type LookupCartProductResultOutput struct{ *pulumi.OutputState }

func (LookupCartProductResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCartProductResult)(nil)).Elem()
}

func (o LookupCartProductResultOutput) ToLookupCartProductResultOutput() LookupCartProductResultOutput {
	return o
}

func (o LookupCartProductResultOutput) ToLookupCartProductResultOutputWithContext(ctx context.Context) LookupCartProductResultOutput {
	return o
}

func (o LookupCartProductResultOutput) CartId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCartProductResult) string { return v.CartId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCartProductResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCartProductResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCartProductResultOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCartProductResult) string { return v.Product }).(pulumi.StringOutput)
}

// products results
func (o LookupCartProductResultOutput) Results() GetCartProductResultArrayOutput {
	return o.ApplyT(func(v LookupCartProductResult) []GetCartProductResult { return v.Results }).(GetCartProductResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCartProductResultOutput{})
}
