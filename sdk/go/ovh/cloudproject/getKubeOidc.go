// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a OVHcloud Managed Kubernetes Service cluster OIDC.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/cloudproject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			oidc, err := cloudproject.GetKubeOidc(ctx, &cloudproject.GetKubeOidcArgs{
//				ServiceName: "XXXXXX",
//				KubeId:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("oidc-val", oidc.ClientId)
//			return nil
//		})
//	}
//
// ```
func LookupKubeOidc(ctx *pulumi.Context, args *LookupKubeOidcArgs, opts ...pulumi.InvokeOption) (*LookupKubeOidcResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKubeOidcResult
	err := ctx.Invoke("ovh:CloudProject/getKubeOidc:getKubeOidc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubeOidc.
type LookupKubeOidcArgs struct {
	// The OIDC client ID.
	ClientId *string `pulumi:"clientId"`
	// The OIDC issuer url.
	IssuerUrl *string `pulumi:"issuerUrl"`
	// The id of the managed kubernetes cluster.
	KubeId string `pulumi:"kubeId"`
	// Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.
	OidcCaContent *string `pulumi:"oidcCaContent"`
	// Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.
	OidcGroupsClaims []string `pulumi:"oidcGroupsClaims"`
	// Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
	OidcGroupsPrefix *string `pulumi:"oidcGroupsPrefix"`
	// Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."
	OidcRequiredClaims []string `pulumi:"oidcRequiredClaims"`
	// Array of signing algorithms accepted. Default is \"RS256\".
	OidcSigningAlgs []string `pulumi:"oidcSigningAlgs"`
	// JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
	OidcUsernameClaim *string `pulumi:"oidcUsernameClaim"`
	// Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and `oidcUsernameClaim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
	OidcUsernamePrefix *string `pulumi:"oidcUsernamePrefix"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getKubeOidc.
type LookupKubeOidcResult struct {
	// The OIDC client ID.
	ClientId *string `pulumi:"clientId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The OIDC issuer url.
	IssuerUrl *string `pulumi:"issuerUrl"`
	// See Argument Reference above.
	KubeId string `pulumi:"kubeId"`
	// Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.
	OidcCaContent *string `pulumi:"oidcCaContent"`
	// Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.
	OidcGroupsClaims []string `pulumi:"oidcGroupsClaims"`
	// Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
	OidcGroupsPrefix *string `pulumi:"oidcGroupsPrefix"`
	// Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."
	OidcRequiredClaims []string `pulumi:"oidcRequiredClaims"`
	// Array of signing algorithms accepted. Default is \"RS256\".
	OidcSigningAlgs []string `pulumi:"oidcSigningAlgs"`
	// JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
	OidcUsernameClaim *string `pulumi:"oidcUsernameClaim"`
	// Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and `oidcUsernameClaim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
	OidcUsernamePrefix *string `pulumi:"oidcUsernamePrefix"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func LookupKubeOidcOutput(ctx *pulumi.Context, args LookupKubeOidcOutputArgs, opts ...pulumi.InvokeOption) LookupKubeOidcResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKubeOidcResultOutput, error) {
			args := v.(LookupKubeOidcArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getKubeOidc:getKubeOidc", args, LookupKubeOidcResultOutput{}, options).(LookupKubeOidcResultOutput), nil
		}).(LookupKubeOidcResultOutput)
}

// A collection of arguments for invoking getKubeOidc.
type LookupKubeOidcOutputArgs struct {
	// The OIDC client ID.
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
	// The OIDC issuer url.
	IssuerUrl pulumi.StringPtrInput `pulumi:"issuerUrl"`
	// The id of the managed kubernetes cluster.
	KubeId pulumi.StringInput `pulumi:"kubeId"`
	// Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.
	OidcCaContent pulumi.StringPtrInput `pulumi:"oidcCaContent"`
	// Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.
	OidcGroupsClaims pulumi.StringArrayInput `pulumi:"oidcGroupsClaims"`
	// Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
	OidcGroupsPrefix pulumi.StringPtrInput `pulumi:"oidcGroupsPrefix"`
	// Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."
	OidcRequiredClaims pulumi.StringArrayInput `pulumi:"oidcRequiredClaims"`
	// Array of signing algorithms accepted. Default is \"RS256\".
	OidcSigningAlgs pulumi.StringArrayInput `pulumi:"oidcSigningAlgs"`
	// JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
	OidcUsernameClaim pulumi.StringPtrInput `pulumi:"oidcUsernameClaim"`
	// Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and `oidcUsernameClaim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
	OidcUsernamePrefix pulumi.StringPtrInput `pulumi:"oidcUsernamePrefix"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupKubeOidcOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubeOidcArgs)(nil)).Elem()
}

// A collection of values returned by getKubeOidc.
type LookupKubeOidcResultOutput struct{ *pulumi.OutputState }

func (LookupKubeOidcResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubeOidcResult)(nil)).Elem()
}

func (o LookupKubeOidcResultOutput) ToLookupKubeOidcResultOutput() LookupKubeOidcResultOutput {
	return o
}

func (o LookupKubeOidcResultOutput) ToLookupKubeOidcResultOutputWithContext(ctx context.Context) LookupKubeOidcResultOutput {
	return o
}

// The OIDC client ID.
func (o LookupKubeOidcResultOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupKubeOidcResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) string { return v.Id }).(pulumi.StringOutput)
}

// The OIDC issuer url.
func (o LookupKubeOidcResultOutput) IssuerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) *string { return v.IssuerUrl }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupKubeOidcResultOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) string { return v.KubeId }).(pulumi.StringOutput)
}

// Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.
func (o LookupKubeOidcResultOutput) OidcCaContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) *string { return v.OidcCaContent }).(pulumi.StringPtrOutput)
}

// Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.
func (o LookupKubeOidcResultOutput) OidcGroupsClaims() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) []string { return v.OidcGroupsClaims }).(pulumi.StringArrayOutput)
}

// Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.
func (o LookupKubeOidcResultOutput) OidcGroupsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) *string { return v.OidcGroupsPrefix }).(pulumi.StringPtrOutput)
}

// Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."
func (o LookupKubeOidcResultOutput) OidcRequiredClaims() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) []string { return v.OidcRequiredClaims }).(pulumi.StringArrayOutput)
}

// Array of signing algorithms accepted. Default is \"RS256\".
func (o LookupKubeOidcResultOutput) OidcSigningAlgs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) []string { return v.OidcSigningAlgs }).(pulumi.StringArrayOutput)
}

// JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.
func (o LookupKubeOidcResultOutput) OidcUsernameClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) *string { return v.OidcUsernameClaim }).(pulumi.StringPtrOutput)
}

// Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and `oidcUsernameClaim` is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.
func (o LookupKubeOidcResultOutput) OidcUsernamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) *string { return v.OidcUsernamePrefix }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupKubeOidcResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeOidcResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubeOidcResultOutput{})
}
