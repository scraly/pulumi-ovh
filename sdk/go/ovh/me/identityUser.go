// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates an identity user.
//
// ## Example Usage
type IdentityUser struct {
	pulumi.CustomResourceState

	// Creation date of this user.
	Creation pulumi.StringOutput `pulumi:"creation"`
	// User description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User's email.
	Email pulumi.StringOutput `pulumi:"email"`
	// User's group.
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// Last update of this user.
	LastUpdate pulumi.StringOutput `pulumi:"lastUpdate"`
	// User's login suffix.
	Login pulumi.StringOutput `pulumi:"login"`
	// User's password.
	Password pulumi.StringOutput `pulumi:"password"`
	// When the user changed his password for the last time.
	PasswordLastUpdate pulumi.StringOutput `pulumi:"passwordLastUpdate"`
	// Current user's status.
	Status pulumi.StringOutput `pulumi:"status"`
	// URN of the user, used when writing IAM policies
	Urn pulumi.StringOutput `pulumi:"urn"`
}

// NewIdentityUser registers a new resource with the given unique name, arguments, and options.
func NewIdentityUser(ctx *pulumi.Context,
	name string, args *IdentityUserArgs, opts ...pulumi.ResourceOption) (*IdentityUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IdentityUser
	err := ctx.RegisterResource("ovh:Me/identityUser:IdentityUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityUser gets an existing IdentityUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityUserState, opts ...pulumi.ResourceOption) (*IdentityUser, error) {
	var resource IdentityUser
	err := ctx.ReadResource("ovh:Me/identityUser:IdentityUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityUser resources.
type identityUserState struct {
	// Creation date of this user.
	Creation *string `pulumi:"creation"`
	// User description.
	Description *string `pulumi:"description"`
	// User's email.
	Email *string `pulumi:"email"`
	// User's group.
	Group *string `pulumi:"group"`
	// Last update of this user.
	LastUpdate *string `pulumi:"lastUpdate"`
	// User's login suffix.
	Login *string `pulumi:"login"`
	// User's password.
	Password *string `pulumi:"password"`
	// When the user changed his password for the last time.
	PasswordLastUpdate *string `pulumi:"passwordLastUpdate"`
	// Current user's status.
	Status *string `pulumi:"status"`
	// URN of the user, used when writing IAM policies
	Urn *string `pulumi:"urn"`
}

type IdentityUserState struct {
	// Creation date of this user.
	Creation pulumi.StringPtrInput
	// User description.
	Description pulumi.StringPtrInput
	// User's email.
	Email pulumi.StringPtrInput
	// User's group.
	Group pulumi.StringPtrInput
	// Last update of this user.
	LastUpdate pulumi.StringPtrInput
	// User's login suffix.
	Login pulumi.StringPtrInput
	// User's password.
	Password pulumi.StringPtrInput
	// When the user changed his password for the last time.
	PasswordLastUpdate pulumi.StringPtrInput
	// Current user's status.
	Status pulumi.StringPtrInput
	// URN of the user, used when writing IAM policies
	Urn pulumi.StringPtrInput
}

func (IdentityUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityUserState)(nil)).Elem()
}

type identityUserArgs struct {
	// User description.
	Description *string `pulumi:"description"`
	// User's email.
	Email string `pulumi:"email"`
	// User's group.
	Group *string `pulumi:"group"`
	// User's login suffix.
	Login string `pulumi:"login"`
	// User's password.
	Password string `pulumi:"password"`
}

// The set of arguments for constructing a IdentityUser resource.
type IdentityUserArgs struct {
	// User description.
	Description pulumi.StringPtrInput
	// User's email.
	Email pulumi.StringInput
	// User's group.
	Group pulumi.StringPtrInput
	// User's login suffix.
	Login pulumi.StringInput
	// User's password.
	Password pulumi.StringInput
}

func (IdentityUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityUserArgs)(nil)).Elem()
}

type IdentityUserInput interface {
	pulumi.Input

	ToIdentityUserOutput() IdentityUserOutput
	ToIdentityUserOutputWithContext(ctx context.Context) IdentityUserOutput
}

func (*IdentityUser) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityUser)(nil)).Elem()
}

func (i *IdentityUser) ToIdentityUserOutput() IdentityUserOutput {
	return i.ToIdentityUserOutputWithContext(context.Background())
}

func (i *IdentityUser) ToIdentityUserOutputWithContext(ctx context.Context) IdentityUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityUserOutput)
}

func (i *IdentityUser) ToOutput(ctx context.Context) pulumix.Output[*IdentityUser] {
	return pulumix.Output[*IdentityUser]{
		OutputState: i.ToIdentityUserOutputWithContext(ctx).OutputState,
	}
}

// IdentityUserArrayInput is an input type that accepts IdentityUserArray and IdentityUserArrayOutput values.
// You can construct a concrete instance of `IdentityUserArrayInput` via:
//
//	IdentityUserArray{ IdentityUserArgs{...} }
type IdentityUserArrayInput interface {
	pulumi.Input

	ToIdentityUserArrayOutput() IdentityUserArrayOutput
	ToIdentityUserArrayOutputWithContext(context.Context) IdentityUserArrayOutput
}

type IdentityUserArray []IdentityUserInput

func (IdentityUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityUser)(nil)).Elem()
}

func (i IdentityUserArray) ToIdentityUserArrayOutput() IdentityUserArrayOutput {
	return i.ToIdentityUserArrayOutputWithContext(context.Background())
}

func (i IdentityUserArray) ToIdentityUserArrayOutputWithContext(ctx context.Context) IdentityUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityUserArrayOutput)
}

func (i IdentityUserArray) ToOutput(ctx context.Context) pulumix.Output[[]*IdentityUser] {
	return pulumix.Output[[]*IdentityUser]{
		OutputState: i.ToIdentityUserArrayOutputWithContext(ctx).OutputState,
	}
}

// IdentityUserMapInput is an input type that accepts IdentityUserMap and IdentityUserMapOutput values.
// You can construct a concrete instance of `IdentityUserMapInput` via:
//
//	IdentityUserMap{ "key": IdentityUserArgs{...} }
type IdentityUserMapInput interface {
	pulumi.Input

	ToIdentityUserMapOutput() IdentityUserMapOutput
	ToIdentityUserMapOutputWithContext(context.Context) IdentityUserMapOutput
}

type IdentityUserMap map[string]IdentityUserInput

func (IdentityUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityUser)(nil)).Elem()
}

func (i IdentityUserMap) ToIdentityUserMapOutput() IdentityUserMapOutput {
	return i.ToIdentityUserMapOutputWithContext(context.Background())
}

func (i IdentityUserMap) ToIdentityUserMapOutputWithContext(ctx context.Context) IdentityUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityUserMapOutput)
}

func (i IdentityUserMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*IdentityUser] {
	return pulumix.Output[map[string]*IdentityUser]{
		OutputState: i.ToIdentityUserMapOutputWithContext(ctx).OutputState,
	}
}

type IdentityUserOutput struct{ *pulumi.OutputState }

func (IdentityUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityUser)(nil)).Elem()
}

func (o IdentityUserOutput) ToIdentityUserOutput() IdentityUserOutput {
	return o
}

func (o IdentityUserOutput) ToIdentityUserOutputWithContext(ctx context.Context) IdentityUserOutput {
	return o
}

func (o IdentityUserOutput) ToOutput(ctx context.Context) pulumix.Output[*IdentityUser] {
	return pulumix.Output[*IdentityUser]{
		OutputState: o.OutputState,
	}
}

// Creation date of this user.
func (o IdentityUserOutput) Creation() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityUser) pulumi.StringOutput { return v.Creation }).(pulumi.StringOutput)
}

// User description.
func (o IdentityUserOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityUser) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User's email.
func (o IdentityUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityUser) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// User's group.
func (o IdentityUserOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityUser) pulumi.StringPtrOutput { return v.Group }).(pulumi.StringPtrOutput)
}

// Last update of this user.
func (o IdentityUserOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityUser) pulumi.StringOutput { return v.LastUpdate }).(pulumi.StringOutput)
}

// User's login suffix.
func (o IdentityUserOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityUser) pulumi.StringOutput { return v.Login }).(pulumi.StringOutput)
}

// User's password.
func (o IdentityUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// When the user changed his password for the last time.
func (o IdentityUserOutput) PasswordLastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityUser) pulumi.StringOutput { return v.PasswordLastUpdate }).(pulumi.StringOutput)
}

// Current user's status.
func (o IdentityUserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityUser) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// URN of the user, used when writing IAM policies
func (o IdentityUserOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v *IdentityUser) pulumi.StringOutput { return v.Urn }).(pulumi.StringOutput)
}

type IdentityUserArrayOutput struct{ *pulumi.OutputState }

func (IdentityUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IdentityUser)(nil)).Elem()
}

func (o IdentityUserArrayOutput) ToIdentityUserArrayOutput() IdentityUserArrayOutput {
	return o
}

func (o IdentityUserArrayOutput) ToIdentityUserArrayOutputWithContext(ctx context.Context) IdentityUserArrayOutput {
	return o
}

func (o IdentityUserArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*IdentityUser] {
	return pulumix.Output[[]*IdentityUser]{
		OutputState: o.OutputState,
	}
}

func (o IdentityUserArrayOutput) Index(i pulumi.IntInput) IdentityUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IdentityUser {
		return vs[0].([]*IdentityUser)[vs[1].(int)]
	}).(IdentityUserOutput)
}

type IdentityUserMapOutput struct{ *pulumi.OutputState }

func (IdentityUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IdentityUser)(nil)).Elem()
}

func (o IdentityUserMapOutput) ToIdentityUserMapOutput() IdentityUserMapOutput {
	return o
}

func (o IdentityUserMapOutput) ToIdentityUserMapOutputWithContext(ctx context.Context) IdentityUserMapOutput {
	return o
}

func (o IdentityUserMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*IdentityUser] {
	return pulumix.Output[map[string]*IdentityUser]{
		OutputState: o.OutputState,
	}
}

func (o IdentityUserMapOutput) MapIndex(k pulumi.StringInput) IdentityUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IdentityUser {
		return vs[0].(map[string]*IdentityUser)[vs[1].(string)]
	}).(IdentityUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityUserInput)(nil)).Elem(), &IdentityUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityUserArrayInput)(nil)).Elem(), IdentityUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityUserMapInput)(nil)).Elem(), IdentityUserMap{})
	pulumi.RegisterOutputType(IdentityUserOutput{})
	pulumi.RegisterOutputType(IdentityUserArrayOutput{})
	pulumi.RegisterOutputType(IdentityUserMapOutput{})
}
