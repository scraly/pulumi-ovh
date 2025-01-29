// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to manage an IP permanent mitigation.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/ip"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ip.NewMitigation(ctx, "mitigation", &ip.MitigationArgs{
//				Ip:             pulumi.String("XXXXXX"),
//				IpOnMitigation: pulumi.String("XXXXXX"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Mitigation struct {
	pulumi.CustomResourceState

	// Set on true if the IP is on auto-mitigation
	Auto pulumi.BoolOutput `pulumi:"auto"`
	// The IP or the CIDR
	Ip pulumi.StringOutput `pulumi:"ip"`
	// IPv4 address
	// * ` permanent  ` - Set on true if the IP is on permanent mitigation
	IpOnMitigation pulumi.StringOutput `pulumi:"ipOnMitigation"`
	// Set on true if your ip is on permanent mitigation
	Permanent pulumi.BoolOutput `pulumi:"permanent"`
	// Current state of the IP on mitigation
	State pulumi.StringOutput `pulumi:"state"`
}

// NewMitigation registers a new resource with the given unique name, arguments, and options.
func NewMitigation(ctx *pulumi.Context,
	name string, args *MitigationArgs, opts ...pulumi.ResourceOption) (*Mitigation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.IpOnMitigation == nil {
		return nil, errors.New("invalid value for required argument 'IpOnMitigation'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Mitigation
	err := ctx.RegisterResource("ovh:Ip/mitigation:Mitigation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMitigation gets an existing Mitigation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMitigation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MitigationState, opts ...pulumi.ResourceOption) (*Mitigation, error) {
	var resource Mitigation
	err := ctx.ReadResource("ovh:Ip/mitigation:Mitigation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mitigation resources.
type mitigationState struct {
	// Set on true if the IP is on auto-mitigation
	Auto *bool `pulumi:"auto"`
	// The IP or the CIDR
	Ip *string `pulumi:"ip"`
	// IPv4 address
	// * ` permanent  ` - Set on true if the IP is on permanent mitigation
	IpOnMitigation *string `pulumi:"ipOnMitigation"`
	// Set on true if your ip is on permanent mitigation
	Permanent *bool `pulumi:"permanent"`
	// Current state of the IP on mitigation
	State *string `pulumi:"state"`
}

type MitigationState struct {
	// Set on true if the IP is on auto-mitigation
	Auto pulumi.BoolPtrInput
	// The IP or the CIDR
	Ip pulumi.StringPtrInput
	// IPv4 address
	// * ` permanent  ` - Set on true if the IP is on permanent mitigation
	IpOnMitigation pulumi.StringPtrInput
	// Set on true if your ip is on permanent mitigation
	Permanent pulumi.BoolPtrInput
	// Current state of the IP on mitigation
	State pulumi.StringPtrInput
}

func (MitigationState) ElementType() reflect.Type {
	return reflect.TypeOf((*mitigationState)(nil)).Elem()
}

type mitigationArgs struct {
	// The IP or the CIDR
	Ip string `pulumi:"ip"`
	// IPv4 address
	// * ` permanent  ` - Set on true if the IP is on permanent mitigation
	IpOnMitigation string `pulumi:"ipOnMitigation"`
	// Set on true if your ip is on permanent mitigation
	Permanent *bool `pulumi:"permanent"`
}

// The set of arguments for constructing a Mitigation resource.
type MitigationArgs struct {
	// The IP or the CIDR
	Ip pulumi.StringInput
	// IPv4 address
	// * ` permanent  ` - Set on true if the IP is on permanent mitigation
	IpOnMitigation pulumi.StringInput
	// Set on true if your ip is on permanent mitigation
	Permanent pulumi.BoolPtrInput
}

func (MitigationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mitigationArgs)(nil)).Elem()
}

type MitigationInput interface {
	pulumi.Input

	ToMitigationOutput() MitigationOutput
	ToMitigationOutputWithContext(ctx context.Context) MitigationOutput
}

func (*Mitigation) ElementType() reflect.Type {
	return reflect.TypeOf((**Mitigation)(nil)).Elem()
}

func (i *Mitigation) ToMitigationOutput() MitigationOutput {
	return i.ToMitigationOutputWithContext(context.Background())
}

func (i *Mitigation) ToMitigationOutputWithContext(ctx context.Context) MitigationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MitigationOutput)
}

// MitigationArrayInput is an input type that accepts MitigationArray and MitigationArrayOutput values.
// You can construct a concrete instance of `MitigationArrayInput` via:
//
//	MitigationArray{ MitigationArgs{...} }
type MitigationArrayInput interface {
	pulumi.Input

	ToMitigationArrayOutput() MitigationArrayOutput
	ToMitigationArrayOutputWithContext(context.Context) MitigationArrayOutput
}

type MitigationArray []MitigationInput

func (MitigationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mitigation)(nil)).Elem()
}

func (i MitigationArray) ToMitigationArrayOutput() MitigationArrayOutput {
	return i.ToMitigationArrayOutputWithContext(context.Background())
}

func (i MitigationArray) ToMitigationArrayOutputWithContext(ctx context.Context) MitigationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MitigationArrayOutput)
}

// MitigationMapInput is an input type that accepts MitigationMap and MitigationMapOutput values.
// You can construct a concrete instance of `MitigationMapInput` via:
//
//	MitigationMap{ "key": MitigationArgs{...} }
type MitigationMapInput interface {
	pulumi.Input

	ToMitigationMapOutput() MitigationMapOutput
	ToMitigationMapOutputWithContext(context.Context) MitigationMapOutput
}

type MitigationMap map[string]MitigationInput

func (MitigationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mitigation)(nil)).Elem()
}

func (i MitigationMap) ToMitigationMapOutput() MitigationMapOutput {
	return i.ToMitigationMapOutputWithContext(context.Background())
}

func (i MitigationMap) ToMitigationMapOutputWithContext(ctx context.Context) MitigationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MitigationMapOutput)
}

type MitigationOutput struct{ *pulumi.OutputState }

func (MitigationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mitigation)(nil)).Elem()
}

func (o MitigationOutput) ToMitigationOutput() MitigationOutput {
	return o
}

func (o MitigationOutput) ToMitigationOutputWithContext(ctx context.Context) MitigationOutput {
	return o
}

// Set on true if the IP is on auto-mitigation
func (o MitigationOutput) Auto() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mitigation) pulumi.BoolOutput { return v.Auto }).(pulumi.BoolOutput)
}

// The IP or the CIDR
func (o MitigationOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Mitigation) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// IPv4 address
// * ` permanent  ` - Set on true if the IP is on permanent mitigation
func (o MitigationOutput) IpOnMitigation() pulumi.StringOutput {
	return o.ApplyT(func(v *Mitigation) pulumi.StringOutput { return v.IpOnMitigation }).(pulumi.StringOutput)
}

// Set on true if your ip is on permanent mitigation
func (o MitigationOutput) Permanent() pulumi.BoolOutput {
	return o.ApplyT(func(v *Mitigation) pulumi.BoolOutput { return v.Permanent }).(pulumi.BoolOutput)
}

// Current state of the IP on mitigation
func (o MitigationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Mitigation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

type MitigationArrayOutput struct{ *pulumi.OutputState }

func (MitigationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mitigation)(nil)).Elem()
}

func (o MitigationArrayOutput) ToMitigationArrayOutput() MitigationArrayOutput {
	return o
}

func (o MitigationArrayOutput) ToMitigationArrayOutputWithContext(ctx context.Context) MitigationArrayOutput {
	return o
}

func (o MitigationArrayOutput) Index(i pulumi.IntInput) MitigationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Mitigation {
		return vs[0].([]*Mitigation)[vs[1].(int)]
	}).(MitigationOutput)
}

type MitigationMapOutput struct{ *pulumi.OutputState }

func (MitigationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mitigation)(nil)).Elem()
}

func (o MitigationMapOutput) ToMitigationMapOutput() MitigationMapOutput {
	return o
}

func (o MitigationMapOutput) ToMitigationMapOutputWithContext(ctx context.Context) MitigationMapOutput {
	return o
}

func (o MitigationMapOutput) MapIndex(k pulumi.StringInput) MitigationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Mitigation {
		return vs[0].(map[string]*Mitigation)[vs[1].(string)]
	}).(MitigationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MitigationInput)(nil)).Elem(), &Mitigation{})
	pulumi.RegisterInputType(reflect.TypeOf((*MitigationArrayInput)(nil)).Elem(), MitigationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MitigationMapInput)(nil)).Elem(), MitigationMap{})
	pulumi.RegisterOutputType(MitigationOutput{})
	pulumi.RegisterOutputType(MitigationArrayOutput{})
	pulumi.RegisterOutputType(MitigationMapOutput{})
}
