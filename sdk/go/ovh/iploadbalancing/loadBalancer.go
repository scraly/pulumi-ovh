// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/iploadbalancing"
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/me"
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/order"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			myaccount, err := me.GetMe(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			mycart, err := order.GetCart(ctx, &order.GetCartArgs{
//				OvhSubsidiary: myaccount.OvhSubsidiary,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			iplb, err := order.GetCartProductPlan(ctx, &order.GetCartProductPlanArgs{
//				CartId:        mycart.Id,
//				PriceCapacity: "renew",
//				Product:       "ipLoadbalancing",
//				PlanCode:      "iplb-lb1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			bhs, err := order.GetCartProductOptionsPlan(ctx, &order.GetCartProductOptionsPlanArgs{
//				CartId:          iplb.CartId,
//				PriceCapacity:   iplb.PriceCapacity,
//				Product:         iplb.Product,
//				PlanCode:        iplb.PlanCode,
//				OptionsPlanCode: "iplb-zone-lb1-rbx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = iploadbalancing.NewLoadBalancer(ctx, "iplb-lb1", &iploadbalancing.LoadBalancerArgs{
//				OvhSubsidiary: pulumi.String(mycart.OvhSubsidiary),
//				DisplayName:   pulumi.String("my ip loadbalancing"),
//				Plan: &iploadbalancing.LoadBalancerPlanArgs{
//					Duration:    pulumi.String(iplb.SelectedPrices[0].Duration),
//					PlanCode:    pulumi.String(iplb.PlanCode),
//					PricingMode: pulumi.String(iplb.SelectedPrices[0].PricingMode),
//				},
//				PlanOptions: iploadbalancing.LoadBalancerPlanOptionArray{
//					&iploadbalancing.LoadBalancerPlanOptionArgs{
//						Duration:    pulumi.String(bhs.SelectedPrices[0].Duration),
//						PlanCode:    pulumi.String(bhs.PlanCode),
//						PricingMode: pulumi.String(bhs.SelectedPrices[0].PricingMode),
//					},
//				},
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
// OVHcloud IP load balancing services can be imported using its `service_name`.
//
// Using the following configuration:
//
// hcl
//
// import {
//
//	to = ovh_iploadbalancing.iplb
//
//	id = "<service name>"
//
// }
//
// You can then run:
//
// bash
//
// $ pulumi preview -generate-config-out=iplb.tf
//
// $ pulumi up
//
// The file `iplb.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above.
//
// See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
type LoadBalancer struct {
	pulumi.CustomResourceState

	// URN of the load balancer, used when writing IAM policies
	LoadBalancerURN pulumi.StringOutput `pulumi:"LoadBalancerURN"`
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Your IP load balancing
	IpLoadbalancing pulumi.StringOutput `pulumi:"ipLoadbalancing"`
	// The IPV4 associated to your IP load balancing
	Ipv4 pulumi.StringOutput `pulumi:"ipv4"`
	// The IPV6 associated to your IP load balancing. DEPRECATED.
	Ipv6 pulumi.StringOutput `pulumi:"ipv6"`
	// The metrics token associated with your IP load balancing
	MetricsToken pulumi.StringOutput `pulumi:"metricsToken"`
	// The offer of your IP load balancing
	Offer pulumi.StringOutput `pulumi:"offer"`
	// Available additional zone for your Load Balancer
	OrderableZones LoadBalancerOrderableZoneArrayOutput `pulumi:"orderableZones"`
	// Details about an Order
	Orders LoadBalancerOrderArrayOutput `pulumi:"orders"`
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary pulumi.StringOutput `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrOutput `pulumi:"paymentMean"`
	// Product Plan to order
	Plan LoadBalancerPlanOutput `pulumi:"plan"`
	// Product Plan to order
	PlanOptions LoadBalancerPlanOptionArrayOutput `pulumi:"planOptions"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
	SslConfiguration pulumi.StringOutput `pulumi:"sslConfiguration"`
	// Current state of your IP
	State pulumi.StringOutput `pulumi:"state"`
	// Vrack eligibility
	VrackEligibility pulumi.BoolOutput `pulumi:"vrackEligibility"`
	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	VrackName pulumi.StringOutput `pulumi:"vrackName"`
	// Location where your service is
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewLoadBalancer registers a new resource with the given unique name, arguments, and options.
func NewLoadBalancer(ctx *pulumi.Context,
	name string, args *LoadBalancerArgs, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	if args == nil {
		args = &LoadBalancerArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"metricsToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LoadBalancer
	err := ctx.RegisterResource("ovh:IpLoadBalancing/loadBalancer:LoadBalancer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoadBalancer gets an existing LoadBalancer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoadBalancer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoadBalancerState, opts ...pulumi.ResourceOption) (*LoadBalancer, error) {
	var resource LoadBalancer
	err := ctx.ReadResource("ovh:IpLoadBalancing/loadBalancer:LoadBalancer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoadBalancer resources.
type loadBalancerState struct {
	// URN of the load balancer, used when writing IAM policies
	LoadBalancerURN *string `pulumi:"LoadBalancerURN"`
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName *string `pulumi:"displayName"`
	// Your IP load balancing
	IpLoadbalancing *string `pulumi:"ipLoadbalancing"`
	// The IPV4 associated to your IP load balancing
	Ipv4 *string `pulumi:"ipv4"`
	// The IPV6 associated to your IP load balancing. DEPRECATED.
	Ipv6 *string `pulumi:"ipv6"`
	// The metrics token associated with your IP load balancing
	MetricsToken *string `pulumi:"metricsToken"`
	// The offer of your IP load balancing
	Offer *string `pulumi:"offer"`
	// Available additional zone for your Load Balancer
	OrderableZones []LoadBalancerOrderableZone `pulumi:"orderableZones"`
	// Details about an Order
	Orders []LoadBalancerOrder `pulumi:"orders"`
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *LoadBalancerPlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []LoadBalancerPlanOption `pulumi:"planOptions"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
	SslConfiguration *string `pulumi:"sslConfiguration"`
	// Current state of your IP
	State *string `pulumi:"state"`
	// Vrack eligibility
	VrackEligibility *bool `pulumi:"vrackEligibility"`
	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	VrackName *string `pulumi:"vrackName"`
	// Location where your service is
	Zones []string `pulumi:"zones"`
}

type LoadBalancerState struct {
	// URN of the load balancer, used when writing IAM policies
	LoadBalancerURN pulumi.StringPtrInput
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName pulumi.StringPtrInput
	// Your IP load balancing
	IpLoadbalancing pulumi.StringPtrInput
	// The IPV4 associated to your IP load balancing
	Ipv4 pulumi.StringPtrInput
	// The IPV6 associated to your IP load balancing. DEPRECATED.
	Ipv6 pulumi.StringPtrInput
	// The metrics token associated with your IP load balancing
	MetricsToken pulumi.StringPtrInput
	// The offer of your IP load balancing
	Offer pulumi.StringPtrInput
	// Available additional zone for your Load Balancer
	OrderableZones LoadBalancerOrderableZoneArrayInput
	// Details about an Order
	Orders LoadBalancerOrderArrayInput
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan LoadBalancerPlanPtrInput
	// Product Plan to order
	PlanOptions LoadBalancerPlanOptionArrayInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
	SslConfiguration pulumi.StringPtrInput
	// Current state of your IP
	State pulumi.StringPtrInput
	// Vrack eligibility
	VrackEligibility pulumi.BoolPtrInput
	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	VrackName pulumi.StringPtrInput
	// Location where your service is
	Zones pulumi.StringArrayInput
}

func (LoadBalancerState) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerState)(nil)).Elem()
}

type loadBalancerArgs struct {
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName *string `pulumi:"displayName"`
	// Details about an Order
	Orders []LoadBalancerOrder `pulumi:"orders"`
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *LoadBalancerPlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []LoadBalancerPlanOption `pulumi:"planOptions"`
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
	SslConfiguration *string `pulumi:"sslConfiguration"`
}

// The set of arguments for constructing a LoadBalancer resource.
type LoadBalancerArgs struct {
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName pulumi.StringPtrInput
	// Details about an Order
	Orders LoadBalancerOrderArrayInput
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan LoadBalancerPlanPtrInput
	// Product Plan to order
	PlanOptions LoadBalancerPlanOptionArrayInput
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
	SslConfiguration pulumi.StringPtrInput
}

func (LoadBalancerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loadBalancerArgs)(nil)).Elem()
}

type LoadBalancerInput interface {
	pulumi.Input

	ToLoadBalancerOutput() LoadBalancerOutput
	ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput
}

func (*LoadBalancer) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (i *LoadBalancer) ToLoadBalancerOutput() LoadBalancerOutput {
	return i.ToLoadBalancerOutputWithContext(context.Background())
}

func (i *LoadBalancer) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerOutput)
}

// LoadBalancerArrayInput is an input type that accepts LoadBalancerArray and LoadBalancerArrayOutput values.
// You can construct a concrete instance of `LoadBalancerArrayInput` via:
//
//	LoadBalancerArray{ LoadBalancerArgs{...} }
type LoadBalancerArrayInput interface {
	pulumi.Input

	ToLoadBalancerArrayOutput() LoadBalancerArrayOutput
	ToLoadBalancerArrayOutputWithContext(context.Context) LoadBalancerArrayOutput
}

type LoadBalancerArray []LoadBalancerInput

func (LoadBalancerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancerArray) ToLoadBalancerArrayOutput() LoadBalancerArrayOutput {
	return i.ToLoadBalancerArrayOutputWithContext(context.Background())
}

func (i LoadBalancerArray) ToLoadBalancerArrayOutputWithContext(ctx context.Context) LoadBalancerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerArrayOutput)
}

// LoadBalancerMapInput is an input type that accepts LoadBalancerMap and LoadBalancerMapOutput values.
// You can construct a concrete instance of `LoadBalancerMapInput` via:
//
//	LoadBalancerMap{ "key": LoadBalancerArgs{...} }
type LoadBalancerMapInput interface {
	pulumi.Input

	ToLoadBalancerMapOutput() LoadBalancerMapOutput
	ToLoadBalancerMapOutputWithContext(context.Context) LoadBalancerMapOutput
}

type LoadBalancerMap map[string]LoadBalancerInput

func (LoadBalancerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancer)(nil)).Elem()
}

func (i LoadBalancerMap) ToLoadBalancerMapOutput() LoadBalancerMapOutput {
	return i.ToLoadBalancerMapOutputWithContext(context.Background())
}

func (i LoadBalancerMap) ToLoadBalancerMapOutputWithContext(ctx context.Context) LoadBalancerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerMapOutput)
}

type LoadBalancerOutput struct{ *pulumi.OutputState }

func (LoadBalancerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerOutput) ToLoadBalancerOutput() LoadBalancerOutput {
	return o
}

func (o LoadBalancerOutput) ToLoadBalancerOutputWithContext(ctx context.Context) LoadBalancerOutput {
	return o
}

// URN of the load balancer, used when writing IAM policies
func (o LoadBalancerOutput) LoadBalancerURN() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.LoadBalancerURN }).(pulumi.StringOutput)
}

// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
func (o LoadBalancerOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Your IP load balancing
func (o LoadBalancerOutput) IpLoadbalancing() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.IpLoadbalancing }).(pulumi.StringOutput)
}

// The IPV4 associated to your IP load balancing
func (o LoadBalancerOutput) Ipv4() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Ipv4 }).(pulumi.StringOutput)
}

// The IPV6 associated to your IP load balancing. DEPRECATED.
func (o LoadBalancerOutput) Ipv6() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Ipv6 }).(pulumi.StringOutput)
}

// The metrics token associated with your IP load balancing
func (o LoadBalancerOutput) MetricsToken() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.MetricsToken }).(pulumi.StringOutput)
}

// The offer of your IP load balancing
func (o LoadBalancerOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.Offer }).(pulumi.StringOutput)
}

// Available additional zone for your Load Balancer
func (o LoadBalancerOutput) OrderableZones() LoadBalancerOrderableZoneArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerOrderableZoneArrayOutput { return v.OrderableZones }).(LoadBalancerOrderableZoneArrayOutput)
}

// Details about an Order
func (o LoadBalancerOutput) Orders() LoadBalancerOrderArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerOrderArrayOutput { return v.Orders }).(LoadBalancerOrderArrayOutput)
}

// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
func (o LoadBalancerOutput) OvhSubsidiary() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.OvhSubsidiary }).(pulumi.StringOutput)
}

// Ovh payment mode
//
// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
func (o LoadBalancerOutput) PaymentMean() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringPtrOutput { return v.PaymentMean }).(pulumi.StringPtrOutput)
}

// Product Plan to order
func (o LoadBalancerOutput) Plan() LoadBalancerPlanOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerPlanOutput { return v.Plan }).(LoadBalancerPlanOutput)
}

// Product Plan to order
func (o LoadBalancerOutput) PlanOptions() LoadBalancerPlanOptionArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) LoadBalancerPlanOptionArrayOutput { return v.PlanOptions }).(LoadBalancerPlanOptionArrayOutput)
}

// The internal name of your IP load balancing
func (o LoadBalancerOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
func (o LoadBalancerOutput) SslConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.SslConfiguration }).(pulumi.StringOutput)
}

// Current state of your IP
func (o LoadBalancerOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Vrack eligibility
func (o LoadBalancerOutput) VrackEligibility() pulumi.BoolOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.BoolOutput { return v.VrackEligibility }).(pulumi.BoolOutput)
}

// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
func (o LoadBalancerOutput) VrackName() pulumi.StringOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringOutput { return v.VrackName }).(pulumi.StringOutput)
}

// Location where your service is
func (o LoadBalancerOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LoadBalancer) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

type LoadBalancerArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerArrayOutput) ToLoadBalancerArrayOutput() LoadBalancerArrayOutput {
	return o
}

func (o LoadBalancerArrayOutput) ToLoadBalancerArrayOutputWithContext(ctx context.Context) LoadBalancerArrayOutput {
	return o
}

func (o LoadBalancerArrayOutput) Index(i pulumi.IntInput) LoadBalancerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoadBalancer {
		return vs[0].([]*LoadBalancer)[vs[1].(int)]
	}).(LoadBalancerOutput)
}

type LoadBalancerMapOutput struct{ *pulumi.OutputState }

func (LoadBalancerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoadBalancer)(nil)).Elem()
}

func (o LoadBalancerMapOutput) ToLoadBalancerMapOutput() LoadBalancerMapOutput {
	return o
}

func (o LoadBalancerMapOutput) ToLoadBalancerMapOutputWithContext(ctx context.Context) LoadBalancerMapOutput {
	return o
}

func (o LoadBalancerMapOutput) MapIndex(k pulumi.StringInput) LoadBalancerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoadBalancer {
		return vs[0].(map[string]*LoadBalancer)[vs[1].(string)]
	}).(LoadBalancerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerInput)(nil)).Elem(), &LoadBalancer{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerArrayInput)(nil)).Elem(), LoadBalancerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoadBalancerMapInput)(nil)).Elem(), LoadBalancerMap{})
	pulumi.RegisterOutputType(LoadBalancerOutput{})
	pulumi.RegisterOutputType(LoadBalancerArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerMapOutput{})
}
