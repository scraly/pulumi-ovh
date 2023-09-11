// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manage rules for TCP route.
//
// ## Example Usage
//
// ## Import
//
// TCP route rule can be imported using the following format `serviceName`, the `id` of the route and the `id` of the rule separated by "/" e.g.
type TcpRouteRule struct {
	pulumi.CustomResourceState

	// Human readable name for your rule, this field is for you
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
	Field pulumi.StringOutput `pulumi:"field"`
	// Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
	Match pulumi.StringOutput `pulumi:"match"`
	// Invert the matching operator effect
	Negate pulumi.BoolOutput `pulumi:"negate"`
	// Value to match against this match. Interpretation if this field depends on the match and field
	Pattern pulumi.StringPtrOutput `pulumi:"pattern"`
	// The route to apply this rule
	RouteId pulumi.StringOutput `pulumi:"routeId"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
	SubField pulumi.StringPtrOutput `pulumi:"subField"`
}

// NewTcpRouteRule registers a new resource with the given unique name, arguments, and options.
func NewTcpRouteRule(ctx *pulumi.Context,
	name string, args *TcpRouteRuleArgs, opts ...pulumi.ResourceOption) (*TcpRouteRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Field == nil {
		return nil, errors.New("invalid value for required argument 'Field'")
	}
	if args.Match == nil {
		return nil, errors.New("invalid value for required argument 'Match'")
	}
	if args.RouteId == nil {
		return nil, errors.New("invalid value for required argument 'RouteId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TcpRouteRule
	err := ctx.RegisterResource("ovh:IpLoadBalancing/tcpRouteRule:TcpRouteRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTcpRouteRule gets an existing TcpRouteRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTcpRouteRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TcpRouteRuleState, opts ...pulumi.ResourceOption) (*TcpRouteRule, error) {
	var resource TcpRouteRule
	err := ctx.ReadResource("ovh:IpLoadBalancing/tcpRouteRule:TcpRouteRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TcpRouteRule resources.
type tcpRouteRuleState struct {
	// Human readable name for your rule, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
	Field *string `pulumi:"field"`
	// Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
	Match *string `pulumi:"match"`
	// Invert the matching operator effect
	Negate *bool `pulumi:"negate"`
	// Value to match against this match. Interpretation if this field depends on the match and field
	Pattern *string `pulumi:"pattern"`
	// The route to apply this rule
	RouteId *string `pulumi:"routeId"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
	SubField *string `pulumi:"subField"`
}

type TcpRouteRuleState struct {
	// Human readable name for your rule, this field is for you
	DisplayName pulumi.StringPtrInput
	// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
	Field pulumi.StringPtrInput
	// Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
	Match pulumi.StringPtrInput
	// Invert the matching operator effect
	Negate pulumi.BoolPtrInput
	// Value to match against this match. Interpretation if this field depends on the match and field
	Pattern pulumi.StringPtrInput
	// The route to apply this rule
	RouteId pulumi.StringPtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
	SubField pulumi.StringPtrInput
}

func (TcpRouteRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*tcpRouteRuleState)(nil)).Elem()
}

type tcpRouteRuleArgs struct {
	// Human readable name for your rule, this field is for you
	DisplayName *string `pulumi:"displayName"`
	// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
	Field string `pulumi:"field"`
	// Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
	Match string `pulumi:"match"`
	// Invert the matching operator effect
	Negate *bool `pulumi:"negate"`
	// Value to match against this match. Interpretation if this field depends on the match and field
	Pattern *string `pulumi:"pattern"`
	// The route to apply this rule
	RouteId string `pulumi:"routeId"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
	SubField *string `pulumi:"subField"`
}

// The set of arguments for constructing a TcpRouteRule resource.
type TcpRouteRuleArgs struct {
	// Human readable name for your rule, this field is for you
	DisplayName pulumi.StringPtrInput
	// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
	Field pulumi.StringInput
	// Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
	Match pulumi.StringInput
	// Invert the matching operator effect
	Negate pulumi.BoolPtrInput
	// Value to match against this match. Interpretation if this field depends on the match and field
	Pattern pulumi.StringPtrInput
	// The route to apply this rule
	RouteId pulumi.StringInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
	SubField pulumi.StringPtrInput
}

func (TcpRouteRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tcpRouteRuleArgs)(nil)).Elem()
}

type TcpRouteRuleInput interface {
	pulumi.Input

	ToTcpRouteRuleOutput() TcpRouteRuleOutput
	ToTcpRouteRuleOutputWithContext(ctx context.Context) TcpRouteRuleOutput
}

func (*TcpRouteRule) ElementType() reflect.Type {
	return reflect.TypeOf((**TcpRouteRule)(nil)).Elem()
}

func (i *TcpRouteRule) ToTcpRouteRuleOutput() TcpRouteRuleOutput {
	return i.ToTcpRouteRuleOutputWithContext(context.Background())
}

func (i *TcpRouteRule) ToTcpRouteRuleOutputWithContext(ctx context.Context) TcpRouteRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpRouteRuleOutput)
}

func (i *TcpRouteRule) ToOutput(ctx context.Context) pulumix.Output[*TcpRouteRule] {
	return pulumix.Output[*TcpRouteRule]{
		OutputState: i.ToTcpRouteRuleOutputWithContext(ctx).OutputState,
	}
}

// TcpRouteRuleArrayInput is an input type that accepts TcpRouteRuleArray and TcpRouteRuleArrayOutput values.
// You can construct a concrete instance of `TcpRouteRuleArrayInput` via:
//
//	TcpRouteRuleArray{ TcpRouteRuleArgs{...} }
type TcpRouteRuleArrayInput interface {
	pulumi.Input

	ToTcpRouteRuleArrayOutput() TcpRouteRuleArrayOutput
	ToTcpRouteRuleArrayOutputWithContext(context.Context) TcpRouteRuleArrayOutput
}

type TcpRouteRuleArray []TcpRouteRuleInput

func (TcpRouteRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TcpRouteRule)(nil)).Elem()
}

func (i TcpRouteRuleArray) ToTcpRouteRuleArrayOutput() TcpRouteRuleArrayOutput {
	return i.ToTcpRouteRuleArrayOutputWithContext(context.Background())
}

func (i TcpRouteRuleArray) ToTcpRouteRuleArrayOutputWithContext(ctx context.Context) TcpRouteRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpRouteRuleArrayOutput)
}

func (i TcpRouteRuleArray) ToOutput(ctx context.Context) pulumix.Output[[]*TcpRouteRule] {
	return pulumix.Output[[]*TcpRouteRule]{
		OutputState: i.ToTcpRouteRuleArrayOutputWithContext(ctx).OutputState,
	}
}

// TcpRouteRuleMapInput is an input type that accepts TcpRouteRuleMap and TcpRouteRuleMapOutput values.
// You can construct a concrete instance of `TcpRouteRuleMapInput` via:
//
//	TcpRouteRuleMap{ "key": TcpRouteRuleArgs{...} }
type TcpRouteRuleMapInput interface {
	pulumi.Input

	ToTcpRouteRuleMapOutput() TcpRouteRuleMapOutput
	ToTcpRouteRuleMapOutputWithContext(context.Context) TcpRouteRuleMapOutput
}

type TcpRouteRuleMap map[string]TcpRouteRuleInput

func (TcpRouteRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TcpRouteRule)(nil)).Elem()
}

func (i TcpRouteRuleMap) ToTcpRouteRuleMapOutput() TcpRouteRuleMapOutput {
	return i.ToTcpRouteRuleMapOutputWithContext(context.Background())
}

func (i TcpRouteRuleMap) ToTcpRouteRuleMapOutputWithContext(ctx context.Context) TcpRouteRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpRouteRuleMapOutput)
}

func (i TcpRouteRuleMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*TcpRouteRule] {
	return pulumix.Output[map[string]*TcpRouteRule]{
		OutputState: i.ToTcpRouteRuleMapOutputWithContext(ctx).OutputState,
	}
}

type TcpRouteRuleOutput struct{ *pulumi.OutputState }

func (TcpRouteRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TcpRouteRule)(nil)).Elem()
}

func (o TcpRouteRuleOutput) ToTcpRouteRuleOutput() TcpRouteRuleOutput {
	return o
}

func (o TcpRouteRuleOutput) ToTcpRouteRuleOutputWithContext(ctx context.Context) TcpRouteRuleOutput {
	return o
}

func (o TcpRouteRuleOutput) ToOutput(ctx context.Context) pulumix.Output[*TcpRouteRule] {
	return pulumix.Output[*TcpRouteRule]{
		OutputState: o.OutputState,
	}
}

// Human readable name for your rule, this field is for you
func (o TcpRouteRuleOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TcpRouteRule) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
func (o TcpRouteRuleOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRouteRule) pulumi.StringOutput { return v.Field }).(pulumi.StringOutput)
}

// Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
func (o TcpRouteRuleOutput) Match() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRouteRule) pulumi.StringOutput { return v.Match }).(pulumi.StringOutput)
}

// Invert the matching operator effect
func (o TcpRouteRuleOutput) Negate() pulumi.BoolOutput {
	return o.ApplyT(func(v *TcpRouteRule) pulumi.BoolOutput { return v.Negate }).(pulumi.BoolOutput)
}

// Value to match against this match. Interpretation if this field depends on the match and field
func (o TcpRouteRuleOutput) Pattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TcpRouteRule) pulumi.StringPtrOutput { return v.Pattern }).(pulumi.StringPtrOutput)
}

// The route to apply this rule
func (o TcpRouteRuleOutput) RouteId() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRouteRule) pulumi.StringOutput { return v.RouteId }).(pulumi.StringOutput)
}

// The internal name of your IP load balancing
func (o TcpRouteRuleOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpRouteRule) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
func (o TcpRouteRuleOutput) SubField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TcpRouteRule) pulumi.StringPtrOutput { return v.SubField }).(pulumi.StringPtrOutput)
}

type TcpRouteRuleArrayOutput struct{ *pulumi.OutputState }

func (TcpRouteRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TcpRouteRule)(nil)).Elem()
}

func (o TcpRouteRuleArrayOutput) ToTcpRouteRuleArrayOutput() TcpRouteRuleArrayOutput {
	return o
}

func (o TcpRouteRuleArrayOutput) ToTcpRouteRuleArrayOutputWithContext(ctx context.Context) TcpRouteRuleArrayOutput {
	return o
}

func (o TcpRouteRuleArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*TcpRouteRule] {
	return pulumix.Output[[]*TcpRouteRule]{
		OutputState: o.OutputState,
	}
}

func (o TcpRouteRuleArrayOutput) Index(i pulumi.IntInput) TcpRouteRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TcpRouteRule {
		return vs[0].([]*TcpRouteRule)[vs[1].(int)]
	}).(TcpRouteRuleOutput)
}

type TcpRouteRuleMapOutput struct{ *pulumi.OutputState }

func (TcpRouteRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TcpRouteRule)(nil)).Elem()
}

func (o TcpRouteRuleMapOutput) ToTcpRouteRuleMapOutput() TcpRouteRuleMapOutput {
	return o
}

func (o TcpRouteRuleMapOutput) ToTcpRouteRuleMapOutputWithContext(ctx context.Context) TcpRouteRuleMapOutput {
	return o
}

func (o TcpRouteRuleMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*TcpRouteRule] {
	return pulumix.Output[map[string]*TcpRouteRule]{
		OutputState: o.OutputState,
	}
}

func (o TcpRouteRuleMapOutput) MapIndex(k pulumi.StringInput) TcpRouteRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TcpRouteRule {
		return vs[0].(map[string]*TcpRouteRule)[vs[1].(string)]
	}).(TcpRouteRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TcpRouteRuleInput)(nil)).Elem(), &TcpRouteRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*TcpRouteRuleArrayInput)(nil)).Elem(), TcpRouteRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TcpRouteRuleMapInput)(nil)).Elem(), TcpRouteRuleMap{})
	pulumi.RegisterOutputType(TcpRouteRuleOutput{})
	pulumi.RegisterOutputType(TcpRouteRuleArrayOutput{})
	pulumi.RegisterOutputType(TcpRouteRuleMapOutput{})
}
