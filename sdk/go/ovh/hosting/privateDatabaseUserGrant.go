// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hosting

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Add grant on a database in your private cloud database instance.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Hosting"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Hosting.NewPrivateDatabaseUserGrant(ctx, "userGrant", &Hosting.PrivateDatabaseUserGrantArgs{
//				DatabaseName: pulumi.String("ovhcloud"),
//				Grant:        pulumi.String("admin"),
//				ServiceName:  pulumi.String("XXXXXX"),
//				UserName:     pulumi.String("terraform"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// OVHcloud database user's grant can be imported using the `service_name`, the `user_name`, the `database_name` and the `grant`, separated by "/" E.g.,
//
// ```sh
// $ pulumi import ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant user service_name/user_name/database_name/grant
// ```
type PrivateDatabaseUserGrant struct {
	pulumi.CustomResourceState

	// Database name where add grant.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Database name where add grant. Values can be:
	// - admin
	// - none
	// - ro
	// - rw
	Grant pulumi.StringOutput `pulumi:"grant"`
	// The internal name of your private database.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// User name used to connect on your databases.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewPrivateDatabaseUserGrant registers a new resource with the given unique name, arguments, and options.
func NewPrivateDatabaseUserGrant(ctx *pulumi.Context,
	name string, args *PrivateDatabaseUserGrantArgs, opts ...pulumi.ResourceOption) (*PrivateDatabaseUserGrant, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Grant == nil {
		return nil, errors.New("invalid value for required argument 'Grant'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivateDatabaseUserGrant
	err := ctx.RegisterResource("ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateDatabaseUserGrant gets an existing PrivateDatabaseUserGrant resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateDatabaseUserGrant(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateDatabaseUserGrantState, opts ...pulumi.ResourceOption) (*PrivateDatabaseUserGrant, error) {
	var resource PrivateDatabaseUserGrant
	err := ctx.ReadResource("ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateDatabaseUserGrant resources.
type privateDatabaseUserGrantState struct {
	// Database name where add grant.
	DatabaseName *string `pulumi:"databaseName"`
	// Database name where add grant. Values can be:
	// - admin
	// - none
	// - ro
	// - rw
	Grant *string `pulumi:"grant"`
	// The internal name of your private database.
	ServiceName *string `pulumi:"serviceName"`
	// User name used to connect on your databases.
	UserName *string `pulumi:"userName"`
}

type PrivateDatabaseUserGrantState struct {
	// Database name where add grant.
	DatabaseName pulumi.StringPtrInput
	// Database name where add grant. Values can be:
	// - admin
	// - none
	// - ro
	// - rw
	Grant pulumi.StringPtrInput
	// The internal name of your private database.
	ServiceName pulumi.StringPtrInput
	// User name used to connect on your databases.
	UserName pulumi.StringPtrInput
}

func (PrivateDatabaseUserGrantState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDatabaseUserGrantState)(nil)).Elem()
}

type privateDatabaseUserGrantArgs struct {
	// Database name where add grant.
	DatabaseName string `pulumi:"databaseName"`
	// Database name where add grant. Values can be:
	// - admin
	// - none
	// - ro
	// - rw
	Grant string `pulumi:"grant"`
	// The internal name of your private database.
	ServiceName string `pulumi:"serviceName"`
	// User name used to connect on your databases.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a PrivateDatabaseUserGrant resource.
type PrivateDatabaseUserGrantArgs struct {
	// Database name where add grant.
	DatabaseName pulumi.StringInput
	// Database name where add grant. Values can be:
	// - admin
	// - none
	// - ro
	// - rw
	Grant pulumi.StringInput
	// The internal name of your private database.
	ServiceName pulumi.StringInput
	// User name used to connect on your databases.
	UserName pulumi.StringInput
}

func (PrivateDatabaseUserGrantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDatabaseUserGrantArgs)(nil)).Elem()
}

type PrivateDatabaseUserGrantInput interface {
	pulumi.Input

	ToPrivateDatabaseUserGrantOutput() PrivateDatabaseUserGrantOutput
	ToPrivateDatabaseUserGrantOutputWithContext(ctx context.Context) PrivateDatabaseUserGrantOutput
}

func (*PrivateDatabaseUserGrant) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDatabaseUserGrant)(nil)).Elem()
}

func (i *PrivateDatabaseUserGrant) ToPrivateDatabaseUserGrantOutput() PrivateDatabaseUserGrantOutput {
	return i.ToPrivateDatabaseUserGrantOutputWithContext(context.Background())
}

func (i *PrivateDatabaseUserGrant) ToPrivateDatabaseUserGrantOutputWithContext(ctx context.Context) PrivateDatabaseUserGrantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseUserGrantOutput)
}

// PrivateDatabaseUserGrantArrayInput is an input type that accepts PrivateDatabaseUserGrantArray and PrivateDatabaseUserGrantArrayOutput values.
// You can construct a concrete instance of `PrivateDatabaseUserGrantArrayInput` via:
//
//	PrivateDatabaseUserGrantArray{ PrivateDatabaseUserGrantArgs{...} }
type PrivateDatabaseUserGrantArrayInput interface {
	pulumi.Input

	ToPrivateDatabaseUserGrantArrayOutput() PrivateDatabaseUserGrantArrayOutput
	ToPrivateDatabaseUserGrantArrayOutputWithContext(context.Context) PrivateDatabaseUserGrantArrayOutput
}

type PrivateDatabaseUserGrantArray []PrivateDatabaseUserGrantInput

func (PrivateDatabaseUserGrantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateDatabaseUserGrant)(nil)).Elem()
}

func (i PrivateDatabaseUserGrantArray) ToPrivateDatabaseUserGrantArrayOutput() PrivateDatabaseUserGrantArrayOutput {
	return i.ToPrivateDatabaseUserGrantArrayOutputWithContext(context.Background())
}

func (i PrivateDatabaseUserGrantArray) ToPrivateDatabaseUserGrantArrayOutputWithContext(ctx context.Context) PrivateDatabaseUserGrantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseUserGrantArrayOutput)
}

// PrivateDatabaseUserGrantMapInput is an input type that accepts PrivateDatabaseUserGrantMap and PrivateDatabaseUserGrantMapOutput values.
// You can construct a concrete instance of `PrivateDatabaseUserGrantMapInput` via:
//
//	PrivateDatabaseUserGrantMap{ "key": PrivateDatabaseUserGrantArgs{...} }
type PrivateDatabaseUserGrantMapInput interface {
	pulumi.Input

	ToPrivateDatabaseUserGrantMapOutput() PrivateDatabaseUserGrantMapOutput
	ToPrivateDatabaseUserGrantMapOutputWithContext(context.Context) PrivateDatabaseUserGrantMapOutput
}

type PrivateDatabaseUserGrantMap map[string]PrivateDatabaseUserGrantInput

func (PrivateDatabaseUserGrantMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateDatabaseUserGrant)(nil)).Elem()
}

func (i PrivateDatabaseUserGrantMap) ToPrivateDatabaseUserGrantMapOutput() PrivateDatabaseUserGrantMapOutput {
	return i.ToPrivateDatabaseUserGrantMapOutputWithContext(context.Background())
}

func (i PrivateDatabaseUserGrantMap) ToPrivateDatabaseUserGrantMapOutputWithContext(ctx context.Context) PrivateDatabaseUserGrantMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseUserGrantMapOutput)
}

type PrivateDatabaseUserGrantOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseUserGrantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDatabaseUserGrant)(nil)).Elem()
}

func (o PrivateDatabaseUserGrantOutput) ToPrivateDatabaseUserGrantOutput() PrivateDatabaseUserGrantOutput {
	return o
}

func (o PrivateDatabaseUserGrantOutput) ToPrivateDatabaseUserGrantOutputWithContext(ctx context.Context) PrivateDatabaseUserGrantOutput {
	return o
}

// Database name where add grant.
func (o PrivateDatabaseUserGrantOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabaseUserGrant) pulumi.StringOutput { return v.DatabaseName }).(pulumi.StringOutput)
}

// Database name where add grant. Values can be:
// - admin
// - none
// - ro
// - rw
func (o PrivateDatabaseUserGrantOutput) Grant() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabaseUserGrant) pulumi.StringOutput { return v.Grant }).(pulumi.StringOutput)
}

// The internal name of your private database.
func (o PrivateDatabaseUserGrantOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabaseUserGrant) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// User name used to connect on your databases.
func (o PrivateDatabaseUserGrantOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabaseUserGrant) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type PrivateDatabaseUserGrantArrayOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseUserGrantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateDatabaseUserGrant)(nil)).Elem()
}

func (o PrivateDatabaseUserGrantArrayOutput) ToPrivateDatabaseUserGrantArrayOutput() PrivateDatabaseUserGrantArrayOutput {
	return o
}

func (o PrivateDatabaseUserGrantArrayOutput) ToPrivateDatabaseUserGrantArrayOutputWithContext(ctx context.Context) PrivateDatabaseUserGrantArrayOutput {
	return o
}

func (o PrivateDatabaseUserGrantArrayOutput) Index(i pulumi.IntInput) PrivateDatabaseUserGrantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivateDatabaseUserGrant {
		return vs[0].([]*PrivateDatabaseUserGrant)[vs[1].(int)]
	}).(PrivateDatabaseUserGrantOutput)
}

type PrivateDatabaseUserGrantMapOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseUserGrantMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateDatabaseUserGrant)(nil)).Elem()
}

func (o PrivateDatabaseUserGrantMapOutput) ToPrivateDatabaseUserGrantMapOutput() PrivateDatabaseUserGrantMapOutput {
	return o
}

func (o PrivateDatabaseUserGrantMapOutput) ToPrivateDatabaseUserGrantMapOutputWithContext(ctx context.Context) PrivateDatabaseUserGrantMapOutput {
	return o
}

func (o PrivateDatabaseUserGrantMapOutput) MapIndex(k pulumi.StringInput) PrivateDatabaseUserGrantOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivateDatabaseUserGrant {
		return vs[0].(map[string]*PrivateDatabaseUserGrant)[vs[1].(string)]
	}).(PrivateDatabaseUserGrantOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseUserGrantInput)(nil)).Elem(), &PrivateDatabaseUserGrant{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseUserGrantArrayInput)(nil)).Elem(), PrivateDatabaseUserGrantArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseUserGrantMapInput)(nil)).Elem(), PrivateDatabaseUserGrantMap{})
	pulumi.RegisterOutputType(PrivateDatabaseUserGrantOutput{})
	pulumi.RegisterOutputType(PrivateDatabaseUserGrantArrayOutput{})
	pulumi.RegisterOutputType(PrivateDatabaseUserGrantMapOutput{})
}
