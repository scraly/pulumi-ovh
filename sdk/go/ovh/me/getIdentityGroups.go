// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the list of the account's identity groups
//
// ## Example Usage
func GetIdentityGroups(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetIdentityGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetIdentityGroupsResult
	err := ctx.Invoke("ovh:Me/getIdentityGroups:getIdentityGroups", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getIdentityGroups.
type GetIdentityGroupsResult struct {
	// The list of the group names of all the identity groups.
	Groups []string `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}
