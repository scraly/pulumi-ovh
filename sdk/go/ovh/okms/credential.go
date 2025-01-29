// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package okms

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a credential for an OVHcloud KMS.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//	"os"
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/me"
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/okms"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func readFileOrPanic(path string) pulumi.StringPtrInput {
//		data, err := os.ReadFile(path)
//		if err != nil {
//			panic(err.Error())
//		}
//		return pulumi.String(string(data))
//	}
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := me.GetMe(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = okms.NewCredential(ctx, "credNoCsr", &okms.CredentialArgs{
//				OkmsId: pulumi.String("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
//				IdentityUrns: pulumi.StringArray{
//					pulumi.Sprintf("urn:v1:eu:identity:account:%v", data.Ovh_me.Current_account.Nichandle),
//				},
//				Description: pulumi.String("Credential without CSR"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = okms.NewCredential(ctx, "credFromCsr", &okms.CredentialArgs{
//				OkmsId: pulumi.String("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
//				IdentityUrns: pulumi.StringArray{
//					pulumi.Sprintf("urn:v1:eu:identity:account:%v", data.Ovh_me.Current_account.Nichandle),
//				},
//				Csr:         pulumi.String(readFileOrPanic("cred.csr")),
//				Description: pulumi.String("Credential from CSR"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Credential struct {
	pulumi.CustomResourceState

	// (String) Certificate PEM of the credential.
	CertificatePem pulumi.StringOutput `pulumi:"certificatePem"`
	// (String) Creation time of the credential
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Valid Certificate Signing Request
	Csr pulumi.StringOutput `pulumi:"csr"`
	// Description of the credential (max 200)
	Description pulumi.StringOutput `pulumi:"description"`
	// (String) Expiration time of the credential
	ExpiredAt pulumi.StringOutput `pulumi:"expiredAt"`
	// (Boolean) Whether the credential was generated from a CSR
	FromCsr pulumi.BoolOutput `pulumi:"fromCsr"`
	// List of identity URNs associated with the credential (max 25)
	IdentityUrns pulumi.StringArrayOutput `pulumi:"identityUrns"`
	// Name of the credential (max 50)
	Name pulumi.StringOutput `pulumi:"name"`
	// Okms ID
	OkmsId pulumi.StringOutput `pulumi:"okmsId"`
	// (String, Sensitive) Private Key PEM of the credential if no CSR is provided
	PrivateKeyPem pulumi.StringOutput `pulumi:"privateKeyPem"`
	// (String) Status of the credential
	Status pulumi.StringOutput `pulumi:"status"`
	// Validity in days (default 365, max 365)
	Validity pulumi.Float64Output `pulumi:"validity"`
}

// NewCredential registers a new resource with the given unique name, arguments, and options.
func NewCredential(ctx *pulumi.Context,
	name string, args *CredentialArgs, opts ...pulumi.ResourceOption) (*Credential, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IdentityUrns == nil {
		return nil, errors.New("invalid value for required argument 'IdentityUrns'")
	}
	if args.OkmsId == nil {
		return nil, errors.New("invalid value for required argument 'OkmsId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"privateKeyPem",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Credential
	err := ctx.RegisterResource("ovh:Okms/credential:Credential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCredential gets an existing Credential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CredentialState, opts ...pulumi.ResourceOption) (*Credential, error) {
	var resource Credential
	err := ctx.ReadResource("ovh:Okms/credential:Credential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Credential resources.
type credentialState struct {
	// (String) Certificate PEM of the credential.
	CertificatePem *string `pulumi:"certificatePem"`
	// (String) Creation time of the credential
	CreatedAt *string `pulumi:"createdAt"`
	// Valid Certificate Signing Request
	Csr *string `pulumi:"csr"`
	// Description of the credential (max 200)
	Description *string `pulumi:"description"`
	// (String) Expiration time of the credential
	ExpiredAt *string `pulumi:"expiredAt"`
	// (Boolean) Whether the credential was generated from a CSR
	FromCsr *bool `pulumi:"fromCsr"`
	// List of identity URNs associated with the credential (max 25)
	IdentityUrns []string `pulumi:"identityUrns"`
	// Name of the credential (max 50)
	Name *string `pulumi:"name"`
	// Okms ID
	OkmsId *string `pulumi:"okmsId"`
	// (String, Sensitive) Private Key PEM of the credential if no CSR is provided
	PrivateKeyPem *string `pulumi:"privateKeyPem"`
	// (String) Status of the credential
	Status *string `pulumi:"status"`
	// Validity in days (default 365, max 365)
	Validity *float64 `pulumi:"validity"`
}

type CredentialState struct {
	// (String) Certificate PEM of the credential.
	CertificatePem pulumi.StringPtrInput
	// (String) Creation time of the credential
	CreatedAt pulumi.StringPtrInput
	// Valid Certificate Signing Request
	Csr pulumi.StringPtrInput
	// Description of the credential (max 200)
	Description pulumi.StringPtrInput
	// (String) Expiration time of the credential
	ExpiredAt pulumi.StringPtrInput
	// (Boolean) Whether the credential was generated from a CSR
	FromCsr pulumi.BoolPtrInput
	// List of identity URNs associated with the credential (max 25)
	IdentityUrns pulumi.StringArrayInput
	// Name of the credential (max 50)
	Name pulumi.StringPtrInput
	// Okms ID
	OkmsId pulumi.StringPtrInput
	// (String, Sensitive) Private Key PEM of the credential if no CSR is provided
	PrivateKeyPem pulumi.StringPtrInput
	// (String) Status of the credential
	Status pulumi.StringPtrInput
	// Validity in days (default 365, max 365)
	Validity pulumi.Float64PtrInput
}

func (CredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialState)(nil)).Elem()
}

type credentialArgs struct {
	// Valid Certificate Signing Request
	Csr *string `pulumi:"csr"`
	// Description of the credential (max 200)
	Description *string `pulumi:"description"`
	// List of identity URNs associated with the credential (max 25)
	IdentityUrns []string `pulumi:"identityUrns"`
	// Name of the credential (max 50)
	Name *string `pulumi:"name"`
	// Okms ID
	OkmsId string `pulumi:"okmsId"`
	// Validity in days (default 365, max 365)
	Validity *float64 `pulumi:"validity"`
}

// The set of arguments for constructing a Credential resource.
type CredentialArgs struct {
	// Valid Certificate Signing Request
	Csr pulumi.StringPtrInput
	// Description of the credential (max 200)
	Description pulumi.StringPtrInput
	// List of identity URNs associated with the credential (max 25)
	IdentityUrns pulumi.StringArrayInput
	// Name of the credential (max 50)
	Name pulumi.StringPtrInput
	// Okms ID
	OkmsId pulumi.StringInput
	// Validity in days (default 365, max 365)
	Validity pulumi.Float64PtrInput
}

func (CredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*credentialArgs)(nil)).Elem()
}

type CredentialInput interface {
	pulumi.Input

	ToCredentialOutput() CredentialOutput
	ToCredentialOutputWithContext(ctx context.Context) CredentialOutput
}

func (*Credential) ElementType() reflect.Type {
	return reflect.TypeOf((**Credential)(nil)).Elem()
}

func (i *Credential) ToCredentialOutput() CredentialOutput {
	return i.ToCredentialOutputWithContext(context.Background())
}

func (i *Credential) ToCredentialOutputWithContext(ctx context.Context) CredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialOutput)
}

// CredentialArrayInput is an input type that accepts CredentialArray and CredentialArrayOutput values.
// You can construct a concrete instance of `CredentialArrayInput` via:
//
//	CredentialArray{ CredentialArgs{...} }
type CredentialArrayInput interface {
	pulumi.Input

	ToCredentialArrayOutput() CredentialArrayOutput
	ToCredentialArrayOutputWithContext(context.Context) CredentialArrayOutput
}

type CredentialArray []CredentialInput

func (CredentialArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Credential)(nil)).Elem()
}

func (i CredentialArray) ToCredentialArrayOutput() CredentialArrayOutput {
	return i.ToCredentialArrayOutputWithContext(context.Background())
}

func (i CredentialArray) ToCredentialArrayOutputWithContext(ctx context.Context) CredentialArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialArrayOutput)
}

// CredentialMapInput is an input type that accepts CredentialMap and CredentialMapOutput values.
// You can construct a concrete instance of `CredentialMapInput` via:
//
//	CredentialMap{ "key": CredentialArgs{...} }
type CredentialMapInput interface {
	pulumi.Input

	ToCredentialMapOutput() CredentialMapOutput
	ToCredentialMapOutputWithContext(context.Context) CredentialMapOutput
}

type CredentialMap map[string]CredentialInput

func (CredentialMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Credential)(nil)).Elem()
}

func (i CredentialMap) ToCredentialMapOutput() CredentialMapOutput {
	return i.ToCredentialMapOutputWithContext(context.Background())
}

func (i CredentialMap) ToCredentialMapOutputWithContext(ctx context.Context) CredentialMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CredentialMapOutput)
}

type CredentialOutput struct{ *pulumi.OutputState }

func (CredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Credential)(nil)).Elem()
}

func (o CredentialOutput) ToCredentialOutput() CredentialOutput {
	return o
}

func (o CredentialOutput) ToCredentialOutputWithContext(ctx context.Context) CredentialOutput {
	return o
}

// (String) Certificate PEM of the credential.
func (o CredentialOutput) CertificatePem() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.CertificatePem }).(pulumi.StringOutput)
}

// (String) Creation time of the credential
func (o CredentialOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Valid Certificate Signing Request
func (o CredentialOutput) Csr() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.Csr }).(pulumi.StringOutput)
}

// Description of the credential (max 200)
func (o CredentialOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// (String) Expiration time of the credential
func (o CredentialOutput) ExpiredAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.ExpiredAt }).(pulumi.StringOutput)
}

// (Boolean) Whether the credential was generated from a CSR
func (o CredentialOutput) FromCsr() pulumi.BoolOutput {
	return o.ApplyT(func(v *Credential) pulumi.BoolOutput { return v.FromCsr }).(pulumi.BoolOutput)
}

// List of identity URNs associated with the credential (max 25)
func (o CredentialOutput) IdentityUrns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringArrayOutput { return v.IdentityUrns }).(pulumi.StringArrayOutput)
}

// Name of the credential (max 50)
func (o CredentialOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Okms ID
func (o CredentialOutput) OkmsId() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.OkmsId }).(pulumi.StringOutput)
}

// (String, Sensitive) Private Key PEM of the credential if no CSR is provided
func (o CredentialOutput) PrivateKeyPem() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.PrivateKeyPem }).(pulumi.StringOutput)
}

// (String) Status of the credential
func (o CredentialOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Credential) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Validity in days (default 365, max 365)
func (o CredentialOutput) Validity() pulumi.Float64Output {
	return o.ApplyT(func(v *Credential) pulumi.Float64Output { return v.Validity }).(pulumi.Float64Output)
}

type CredentialArrayOutput struct{ *pulumi.OutputState }

func (CredentialArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Credential)(nil)).Elem()
}

func (o CredentialArrayOutput) ToCredentialArrayOutput() CredentialArrayOutput {
	return o
}

func (o CredentialArrayOutput) ToCredentialArrayOutputWithContext(ctx context.Context) CredentialArrayOutput {
	return o
}

func (o CredentialArrayOutput) Index(i pulumi.IntInput) CredentialOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Credential {
		return vs[0].([]*Credential)[vs[1].(int)]
	}).(CredentialOutput)
}

type CredentialMapOutput struct{ *pulumi.OutputState }

func (CredentialMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Credential)(nil)).Elem()
}

func (o CredentialMapOutput) ToCredentialMapOutput() CredentialMapOutput {
	return o
}

func (o CredentialMapOutput) ToCredentialMapOutputWithContext(ctx context.Context) CredentialMapOutput {
	return o
}

func (o CredentialMapOutput) MapIndex(k pulumi.StringInput) CredentialOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Credential {
		return vs[0].(map[string]*Credential)[vs[1].(string)]
	}).(CredentialOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialInput)(nil)).Elem(), &Credential{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialArrayInput)(nil)).Elem(), CredentialArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CredentialMapInput)(nil)).Elem(), CredentialMap{})
	pulumi.RegisterOutputType(CredentialOutput{})
	pulumi.RegisterOutputType(CredentialArrayOutput{})
	pulumi.RegisterOutputType(CredentialMapOutput{})
}
