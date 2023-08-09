// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to list all the IAM resource types.
//
// ## Important
//
// > Using this resource requires that the account is enrolled in the OVHcloud [IAM beta](https://labs.ovhcloud.com/en/iam/)
//
// ## Example Usage
func GetReferenceResourceType(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetReferenceResourceTypeResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetReferenceResourceTypeResult
	err := ctx.Invoke("ovh:Iam/getReferenceResourceType:getReferenceResourceType", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getReferenceResourceType.
type GetReferenceResourceTypeResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of resource types
	Types []string `pulumi:"types"`
}