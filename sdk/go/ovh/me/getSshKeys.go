// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Use this data source to retrieve list of names of the account's SSH keys.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Me.GetSshKeys(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetSshKeys(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetSshKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSshKeysResult
	err := ctx.Invoke("ovh:Me/getSshKeys:getSshKeys", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getSshKeys.
type GetSshKeysResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of the names of all the SSH keys.
	Names []string `pulumi:"names"`
}

func GetSshKeysOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetSshKeysResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetSshKeysResult, error) {
		r, err := GetSshKeys(ctx, opts...)
		var s GetSshKeysResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetSshKeysResultOutput)
}

// A collection of values returned by getSshKeys.
type GetSshKeysResultOutput struct{ *pulumi.OutputState }

func (GetSshKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSshKeysResult)(nil)).Elem()
}

func (o GetSshKeysResultOutput) ToGetSshKeysResultOutput() GetSshKeysResultOutput {
	return o
}

func (o GetSshKeysResultOutput) ToGetSshKeysResultOutputWithContext(ctx context.Context) GetSshKeysResultOutput {
	return o
}

func (o GetSshKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSshKeysResult] {
	return pulumix.Output[GetSshKeysResult]{
		OutputState: o.OutputState,
	}
}

// The provider-assigned unique ID for this managed resource.
func (o GetSshKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetSshKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of the names of all the SSH keys.
func (o GetSshKeysResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetSshKeysResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSshKeysResultOutput{})
}
