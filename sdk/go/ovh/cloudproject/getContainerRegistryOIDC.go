// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a OVHcloud Managed Private Registry OIDC.
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
//			my_oidc, err := CloudProject.GetContainerRegistryOIDC(ctx, &cloudproject.GetContainerRegistryOIDCArgs{
//				ServiceName: "XXXXXX",
//				RegistryId:  "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("oidc-client-id", my_oidc.OidcClientId)
//			return nil
//		})
//	}
//
// ```
func LookupContainerRegistryOIDC(ctx *pulumi.Context, args *LookupContainerRegistryOIDCArgs, opts ...pulumi.InvokeOption) (*LookupContainerRegistryOIDCResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupContainerRegistryOIDCResult
	err := ctx.Invoke("ovh:CloudProject/getContainerRegistryOIDC:getContainerRegistryOIDC", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerRegistryOIDC.
type LookupContainerRegistryOIDCArgs struct {
	// Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
	OidcAdminGroup *string `pulumi:"oidcAdminGroup"`
	// Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
	OidcAutoOnboard *bool `pulumi:"oidcAutoOnboard"`
	// The client ID with which Harbor is registered as client application with the OIDC provider.
	OidcClientId *string `pulumi:"oidcClientId"`
	// The URL of an OIDC-compliant server.
	OidcEndpoint *string `pulumi:"oidcEndpoint"`
	// The name of Claim in the ID token whose value is the list of group names.
	OidcGroupsClaim *string `pulumi:"oidcGroupsClaim"`
	// The name of the OIDC provider.
	OidcName *string `pulumi:"oidcName"`
	// The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
	OidcScope *string `pulumi:"oidcScope"`
	// The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
	OidcUserClaim *string `pulumi:"oidcUserClaim"`
	// Set it to `false` if your OIDC server is hosted via self-signed certificate.
	OidcVerifyCert *bool `pulumi:"oidcVerifyCert"`
	// The id of the Managed Private Registry.
	RegistryId string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getContainerRegistryOIDC.
type LookupContainerRegistryOIDCResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
	OidcAdminGroup *string `pulumi:"oidcAdminGroup"`
	// Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
	OidcAutoOnboard *bool `pulumi:"oidcAutoOnboard"`
	// The client ID with which Harbor is registered as client application with the OIDC provider.
	OidcClientId *string `pulumi:"oidcClientId"`
	// The URL of an OIDC-compliant server.
	OidcEndpoint *string `pulumi:"oidcEndpoint"`
	// The name of Claim in the ID token whose value is the list of group names.
	OidcGroupsClaim *string `pulumi:"oidcGroupsClaim"`
	// The name of the OIDC provider.
	OidcName *string `pulumi:"oidcName"`
	// The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
	OidcScope *string `pulumi:"oidcScope"`
	// The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
	OidcUserClaim *string `pulumi:"oidcUserClaim"`
	// Set it to `false` if your OIDC server is hosted via self-signed certificate.
	OidcVerifyCert *bool `pulumi:"oidcVerifyCert"`
	// The ID of the Managed Private Registry.
	RegistryId string `pulumi:"registryId"`
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

func LookupContainerRegistryOIDCOutput(ctx *pulumi.Context, args LookupContainerRegistryOIDCOutputArgs, opts ...pulumi.InvokeOption) LookupContainerRegistryOIDCResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupContainerRegistryOIDCResult, error) {
			args := v.(LookupContainerRegistryOIDCArgs)
			r, err := LookupContainerRegistryOIDC(ctx, &args, opts...)
			var s LookupContainerRegistryOIDCResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupContainerRegistryOIDCResultOutput)
}

// A collection of arguments for invoking getContainerRegistryOIDC.
type LookupContainerRegistryOIDCOutputArgs struct {
	// Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
	OidcAdminGroup pulumi.StringPtrInput `pulumi:"oidcAdminGroup"`
	// Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
	OidcAutoOnboard pulumi.BoolPtrInput `pulumi:"oidcAutoOnboard"`
	// The client ID with which Harbor is registered as client application with the OIDC provider.
	OidcClientId pulumi.StringPtrInput `pulumi:"oidcClientId"`
	// The URL of an OIDC-compliant server.
	OidcEndpoint pulumi.StringPtrInput `pulumi:"oidcEndpoint"`
	// The name of Claim in the ID token whose value is the list of group names.
	OidcGroupsClaim pulumi.StringPtrInput `pulumi:"oidcGroupsClaim"`
	// The name of the OIDC provider.
	OidcName pulumi.StringPtrInput `pulumi:"oidcName"`
	// The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
	OidcScope pulumi.StringPtrInput `pulumi:"oidcScope"`
	// The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
	OidcUserClaim pulumi.StringPtrInput `pulumi:"oidcUserClaim"`
	// Set it to `false` if your OIDC server is hosted via self-signed certificate.
	OidcVerifyCert pulumi.BoolPtrInput `pulumi:"oidcVerifyCert"`
	// The id of the Managed Private Registry.
	RegistryId pulumi.StringInput `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupContainerRegistryOIDCOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerRegistryOIDCArgs)(nil)).Elem()
}

// A collection of values returned by getContainerRegistryOIDC.
type LookupContainerRegistryOIDCResultOutput struct{ *pulumi.OutputState }

func (LookupContainerRegistryOIDCResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupContainerRegistryOIDCResult)(nil)).Elem()
}

func (o LookupContainerRegistryOIDCResultOutput) ToLookupContainerRegistryOIDCResultOutput() LookupContainerRegistryOIDCResultOutput {
	return o
}

func (o LookupContainerRegistryOIDCResultOutput) ToLookupContainerRegistryOIDCResultOutputWithContext(ctx context.Context) LookupContainerRegistryOIDCResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupContainerRegistryOIDCResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) string { return v.Id }).(pulumi.StringOutput)
}

// Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
func (o LookupContainerRegistryOIDCResultOutput) OidcAdminGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) *string { return v.OidcAdminGroup }).(pulumi.StringPtrOutput)
}

// Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
func (o LookupContainerRegistryOIDCResultOutput) OidcAutoOnboard() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) *bool { return v.OidcAutoOnboard }).(pulumi.BoolPtrOutput)
}

// The client ID with which Harbor is registered as client application with the OIDC provider.
func (o LookupContainerRegistryOIDCResultOutput) OidcClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) *string { return v.OidcClientId }).(pulumi.StringPtrOutput)
}

// The URL of an OIDC-compliant server.
func (o LookupContainerRegistryOIDCResultOutput) OidcEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) *string { return v.OidcEndpoint }).(pulumi.StringPtrOutput)
}

// The name of Claim in the ID token whose value is the list of group names.
func (o LookupContainerRegistryOIDCResultOutput) OidcGroupsClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) *string { return v.OidcGroupsClaim }).(pulumi.StringPtrOutput)
}

// The name of the OIDC provider.
func (o LookupContainerRegistryOIDCResultOutput) OidcName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) *string { return v.OidcName }).(pulumi.StringPtrOutput)
}

// The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
func (o LookupContainerRegistryOIDCResultOutput) OidcScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) *string { return v.OidcScope }).(pulumi.StringPtrOutput)
}

// The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
func (o LookupContainerRegistryOIDCResultOutput) OidcUserClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) *string { return v.OidcUserClaim }).(pulumi.StringPtrOutput)
}

// Set it to `false` if your OIDC server is hosted via self-signed certificate.
func (o LookupContainerRegistryOIDCResultOutput) OidcVerifyCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) *bool { return v.OidcVerifyCert }).(pulumi.BoolPtrOutput)
}

// The ID of the Managed Private Registry.
func (o LookupContainerRegistryOIDCResultOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) string { return v.RegistryId }).(pulumi.StringOutput)
}

// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o LookupContainerRegistryOIDCResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupContainerRegistryOIDCResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupContainerRegistryOIDCResultOutput{})
}
