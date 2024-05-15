// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage rules for HTTP route.
//
// ## Example Usage
//
// Route which redirect all URL to HTTPs for example.com (Vhost).
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/IpLoadBalancing"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			httpsredirect, err := IpLoadBalancing.NewHttpRoute(ctx, "httpsredirect", &IpLoadBalancing.HttpRouteArgs{
//				Action: &iploadbalancing.HttpRouteActionArgs{
//					Status: pulumi.Int(302),
//					Target: pulumi.String("https://${host}${path}${arguments}"),
//					Type:   pulumi.String("redirect"),
//				},
//				DisplayName: pulumi.String("Redirect to HTTPS"),
//				FrontendId:  pulumi.Int(11111),
//				ServiceName: pulumi.String("loadbalancer-xxxxxxxxxxxxxxxxxx"),
//				Weight:      pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = IpLoadBalancing.NewHttpRouteRule(ctx, "examplerule", &IpLoadBalancing.HttpRouteRuleArgs{
//				DisplayName: pulumi.String("Match example.com host"),
//				Field:       pulumi.String("host"),
//				Match:       pulumi.String("is"),
//				Negate:      pulumi.Bool(false),
//				Pattern:     pulumi.String("example.com"),
//				RouteId:     httpsredirect.ID(),
//				ServiceName: pulumi.String("loadbalancer-xxxxxxxxxxxxxxxxxx"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// Rule which match a specific header (same effect as the host match above).
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/IpLoadBalancing"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := IpLoadBalancing.NewHttpRouteRule(ctx, "examplerule", &IpLoadBalancing.HttpRouteRuleArgs{
//				DisplayName: pulumi.String("Match example.com Host header"),
//				Field:       pulumi.String("headers"),
//				Match:       pulumi.String("is"),
//				Negate:      pulumi.Bool(false),
//				Pattern:     pulumi.String("example.com"),
//				RouteId:     pulumi.Any(ovh_iploadbalancing_http_route.Httpsredirect.Id),
//				ServiceName: pulumi.String("loadbalancer-xxxxxxxxxxxxxxxxxx"),
//				SubField:    pulumi.String("Host"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// HTTP route rule can be imported using the following format `serviceName`, the `id` of the route and the `id` of the rule separated by "/" e.g.
type HttpRouteRule struct {
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

// NewHttpRouteRule registers a new resource with the given unique name, arguments, and options.
func NewHttpRouteRule(ctx *pulumi.Context,
	name string, args *HttpRouteRuleArgs, opts ...pulumi.ResourceOption) (*HttpRouteRule, error) {
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
	var resource HttpRouteRule
	err := ctx.RegisterResource("ovh:IpLoadBalancing/httpRouteRule:HttpRouteRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHttpRouteRule gets an existing HttpRouteRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpRouteRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HttpRouteRuleState, opts ...pulumi.ResourceOption) (*HttpRouteRule, error) {
	var resource HttpRouteRule
	err := ctx.ReadResource("ovh:IpLoadBalancing/httpRouteRule:HttpRouteRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HttpRouteRule resources.
type httpRouteRuleState struct {
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

type HttpRouteRuleState struct {
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

func (HttpRouteRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpRouteRuleState)(nil)).Elem()
}

type httpRouteRuleArgs struct {
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

// The set of arguments for constructing a HttpRouteRule resource.
type HttpRouteRuleArgs struct {
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

func (HttpRouteRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*httpRouteRuleArgs)(nil)).Elem()
}

type HttpRouteRuleInput interface {
	pulumi.Input

	ToHttpRouteRuleOutput() HttpRouteRuleOutput
	ToHttpRouteRuleOutputWithContext(ctx context.Context) HttpRouteRuleOutput
}

func (*HttpRouteRule) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpRouteRule)(nil)).Elem()
}

func (i *HttpRouteRule) ToHttpRouteRuleOutput() HttpRouteRuleOutput {
	return i.ToHttpRouteRuleOutputWithContext(context.Background())
}

func (i *HttpRouteRule) ToHttpRouteRuleOutputWithContext(ctx context.Context) HttpRouteRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRouteRuleOutput)
}

// HttpRouteRuleArrayInput is an input type that accepts HttpRouteRuleArray and HttpRouteRuleArrayOutput values.
// You can construct a concrete instance of `HttpRouteRuleArrayInput` via:
//
//	HttpRouteRuleArray{ HttpRouteRuleArgs{...} }
type HttpRouteRuleArrayInput interface {
	pulumi.Input

	ToHttpRouteRuleArrayOutput() HttpRouteRuleArrayOutput
	ToHttpRouteRuleArrayOutputWithContext(context.Context) HttpRouteRuleArrayOutput
}

type HttpRouteRuleArray []HttpRouteRuleInput

func (HttpRouteRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpRouteRule)(nil)).Elem()
}

func (i HttpRouteRuleArray) ToHttpRouteRuleArrayOutput() HttpRouteRuleArrayOutput {
	return i.ToHttpRouteRuleArrayOutputWithContext(context.Background())
}

func (i HttpRouteRuleArray) ToHttpRouteRuleArrayOutputWithContext(ctx context.Context) HttpRouteRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRouteRuleArrayOutput)
}

// HttpRouteRuleMapInput is an input type that accepts HttpRouteRuleMap and HttpRouteRuleMapOutput values.
// You can construct a concrete instance of `HttpRouteRuleMapInput` via:
//
//	HttpRouteRuleMap{ "key": HttpRouteRuleArgs{...} }
type HttpRouteRuleMapInput interface {
	pulumi.Input

	ToHttpRouteRuleMapOutput() HttpRouteRuleMapOutput
	ToHttpRouteRuleMapOutputWithContext(context.Context) HttpRouteRuleMapOutput
}

type HttpRouteRuleMap map[string]HttpRouteRuleInput

func (HttpRouteRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpRouteRule)(nil)).Elem()
}

func (i HttpRouteRuleMap) ToHttpRouteRuleMapOutput() HttpRouteRuleMapOutput {
	return i.ToHttpRouteRuleMapOutputWithContext(context.Background())
}

func (i HttpRouteRuleMap) ToHttpRouteRuleMapOutputWithContext(ctx context.Context) HttpRouteRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpRouteRuleMapOutput)
}

type HttpRouteRuleOutput struct{ *pulumi.OutputState }

func (HttpRouteRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpRouteRule)(nil)).Elem()
}

func (o HttpRouteRuleOutput) ToHttpRouteRuleOutput() HttpRouteRuleOutput {
	return o
}

func (o HttpRouteRuleOutput) ToHttpRouteRuleOutputWithContext(ctx context.Context) HttpRouteRuleOutput {
	return o
}

// Human readable name for your rule, this field is for you
func (o HttpRouteRuleOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRouteRule) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
func (o HttpRouteRuleOutput) Field() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpRouteRule) pulumi.StringOutput { return v.Field }).(pulumi.StringOutput)
}

// Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
func (o HttpRouteRuleOutput) Match() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpRouteRule) pulumi.StringOutput { return v.Match }).(pulumi.StringOutput)
}

// Invert the matching operator effect
func (o HttpRouteRuleOutput) Negate() pulumi.BoolOutput {
	return o.ApplyT(func(v *HttpRouteRule) pulumi.BoolOutput { return v.Negate }).(pulumi.BoolOutput)
}

// Value to match against this match. Interpretation if this field depends on the match and field
func (o HttpRouteRuleOutput) Pattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRouteRule) pulumi.StringPtrOutput { return v.Pattern }).(pulumi.StringPtrOutput)
}

// The route to apply this rule
func (o HttpRouteRuleOutput) RouteId() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpRouteRule) pulumi.StringOutput { return v.RouteId }).(pulumi.StringOutput)
}

// The internal name of your IP load balancing
func (o HttpRouteRuleOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpRouteRule) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
func (o HttpRouteRuleOutput) SubField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpRouteRule) pulumi.StringPtrOutput { return v.SubField }).(pulumi.StringPtrOutput)
}

type HttpRouteRuleArrayOutput struct{ *pulumi.OutputState }

func (HttpRouteRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpRouteRule)(nil)).Elem()
}

func (o HttpRouteRuleArrayOutput) ToHttpRouteRuleArrayOutput() HttpRouteRuleArrayOutput {
	return o
}

func (o HttpRouteRuleArrayOutput) ToHttpRouteRuleArrayOutputWithContext(ctx context.Context) HttpRouteRuleArrayOutput {
	return o
}

func (o HttpRouteRuleArrayOutput) Index(i pulumi.IntInput) HttpRouteRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HttpRouteRule {
		return vs[0].([]*HttpRouteRule)[vs[1].(int)]
	}).(HttpRouteRuleOutput)
}

type HttpRouteRuleMapOutput struct{ *pulumi.OutputState }

func (HttpRouteRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpRouteRule)(nil)).Elem()
}

func (o HttpRouteRuleMapOutput) ToHttpRouteRuleMapOutput() HttpRouteRuleMapOutput {
	return o
}

func (o HttpRouteRuleMapOutput) ToHttpRouteRuleMapOutputWithContext(ctx context.Context) HttpRouteRuleMapOutput {
	return o
}

func (o HttpRouteRuleMapOutput) MapIndex(k pulumi.StringInput) HttpRouteRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HttpRouteRule {
		return vs[0].(map[string]*HttpRouteRule)[vs[1].(string)]
	}).(HttpRouteRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HttpRouteRuleInput)(nil)).Elem(), &HttpRouteRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpRouteRuleArrayInput)(nil)).Elem(), HttpRouteRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpRouteRuleMapInput)(nil)).Elem(), HttpRouteRuleMap{})
	pulumi.RegisterOutputType(HttpRouteRuleOutput{})
	pulumi.RegisterOutputType(HttpRouteRuleArrayOutput{})
	pulumi.RegisterOutputType(HttpRouteRuleMapOutput{})
}
