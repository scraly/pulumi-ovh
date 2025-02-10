// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

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
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/ip"
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/me"
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/order"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := me.GetMe(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			mycart, err := order.GetCart(ctx, &order.GetCartArgs{
//				OvhSubsidiary: "fr",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ipblockCartProductPlan, err := order.GetCartProductPlan(ctx, &order.GetCartProductPlanArgs{
//				CartId:        mycart.Id,
//				PriceCapacity: "renew",
//				Product:       "ip",
//				PlanCode:      "ip-v4-s30-ripe",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ip.NewIpService(ctx, "ipblockIpService", &ip.IpServiceArgs{
//				OvhSubsidiary: pulumi.String(mycart.OvhSubsidiary),
//				Description:   pulumi.String("my ip block"),
//				Plan: &ip.IpServicePlanArgs{
//					Duration:    pulumi.String(ipblockCartProductPlan.SelectedPrices[0].Duration),
//					PlanCode:    pulumi.String(ipblockCartProductPlan.PlanCode),
//					PricingMode: pulumi.String(ipblockCartProductPlan.SelectedPrices[0].PricingMode),
//					Configurations: ip.IpServicePlanConfigurationArray{
//						&ip.IpServicePlanConfigurationArgs{
//							Label: pulumi.String("country"),
//							Value: pulumi.String("FR"),
//						},
//						&ip.IpServicePlanConfigurationArgs{
//							Label: pulumi.String("region"),
//							Value: pulumi.String("europe"),
//						},
//						&ip.IpServicePlanConfigurationArgs{
//							Label: pulumi.String("destination"),
//							Value: pulumi.String("parking"),
//						},
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
// The resource can be imported using its `service_name`, E.g.,
//
// hcl
//
// import {
//
//	to = ovh_ip_service.ipblock
//
//	id = "ip-xx.xx.xx.xx"
//
// }
//
// bash
//
// $ pulumi preview -generate-config-out=ipblock.tf
//
// $ pulumi up
//
// The file `ipblock.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above.
//
// See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
type IpService struct {
	pulumi.CustomResourceState

	// can be terminated
	CanBeTerminated pulumi.BoolOutput `pulumi:"canBeTerminated"`
	// country
	Country pulumi.StringOutput `pulumi:"country"`
	// Custom description on your ip.
	Description pulumi.StringOutput `pulumi:"description"`
	// ip block
	Ip pulumi.StringOutput `pulumi:"ip"`
	// Details about an Order
	Orders IpServiceOrderArrayOutput `pulumi:"orders"`
	// IP block organisation Id
	OrganisationId pulumi.StringOutput `pulumi:"organisationId"`
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary pulumi.StringOutput `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrOutput `pulumi:"paymentMean"`
	// Product Plan to order
	Plan IpServicePlanOutput `pulumi:"plan"`
	// Product Plan to order
	PlanOptions IpServicePlanOptionArrayOutput `pulumi:"planOptions"`
	// Routage information
	RoutedTos IpServiceRoutedToArrayOutput `pulumi:"routedTos"`
	// service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Possible values for ip type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIpService registers a new resource with the given unique name, arguments, and options.
func NewIpService(ctx *pulumi.Context,
	name string, args *IpServiceArgs, opts ...pulumi.ResourceOption) (*IpService, error) {
	if args == nil {
		args = &IpServiceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IpService
	err := ctx.RegisterResource("ovh:Ip/ipService:IpService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpService gets an existing IpService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpServiceState, opts ...pulumi.ResourceOption) (*IpService, error) {
	var resource IpService
	err := ctx.ReadResource("ovh:Ip/ipService:IpService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpService resources.
type ipServiceState struct {
	// can be terminated
	CanBeTerminated *bool `pulumi:"canBeTerminated"`
	// country
	Country *string `pulumi:"country"`
	// Custom description on your ip.
	Description *string `pulumi:"description"`
	// ip block
	Ip *string `pulumi:"ip"`
	// Details about an Order
	Orders []IpServiceOrder `pulumi:"orders"`
	// IP block organisation Id
	OrganisationId *string `pulumi:"organisationId"`
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *IpServicePlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []IpServicePlanOption `pulumi:"planOptions"`
	// Routage information
	RoutedTos []IpServiceRoutedTo `pulumi:"routedTos"`
	// service name
	ServiceName *string `pulumi:"serviceName"`
	// Possible values for ip type
	Type *string `pulumi:"type"`
}

type IpServiceState struct {
	// can be terminated
	CanBeTerminated pulumi.BoolPtrInput
	// country
	Country pulumi.StringPtrInput
	// Custom description on your ip.
	Description pulumi.StringPtrInput
	// ip block
	Ip pulumi.StringPtrInput
	// Details about an Order
	Orders IpServiceOrderArrayInput
	// IP block organisation Id
	OrganisationId pulumi.StringPtrInput
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan IpServicePlanPtrInput
	// Product Plan to order
	PlanOptions IpServicePlanOptionArrayInput
	// Routage information
	RoutedTos IpServiceRoutedToArrayInput
	// service name
	ServiceName pulumi.StringPtrInput
	// Possible values for ip type
	Type pulumi.StringPtrInput
}

func (IpServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipServiceState)(nil)).Elem()
}

type ipServiceArgs struct {
	// Custom description on your ip.
	Description *string `pulumi:"description"`
	// Details about an Order
	Orders []IpServiceOrder `pulumi:"orders"`
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *IpServicePlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []IpServicePlanOption `pulumi:"planOptions"`
}

// The set of arguments for constructing a IpService resource.
type IpServiceArgs struct {
	// Custom description on your ip.
	Description pulumi.StringPtrInput
	// Details about an Order
	Orders IpServiceOrderArrayInput
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan IpServicePlanPtrInput
	// Product Plan to order
	PlanOptions IpServicePlanOptionArrayInput
}

func (IpServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipServiceArgs)(nil)).Elem()
}

type IpServiceInput interface {
	pulumi.Input

	ToIpServiceOutput() IpServiceOutput
	ToIpServiceOutputWithContext(ctx context.Context) IpServiceOutput
}

func (*IpService) ElementType() reflect.Type {
	return reflect.TypeOf((**IpService)(nil)).Elem()
}

func (i *IpService) ToIpServiceOutput() IpServiceOutput {
	return i.ToIpServiceOutputWithContext(context.Background())
}

func (i *IpService) ToIpServiceOutputWithContext(ctx context.Context) IpServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpServiceOutput)
}

// IpServiceArrayInput is an input type that accepts IpServiceArray and IpServiceArrayOutput values.
// You can construct a concrete instance of `IpServiceArrayInput` via:
//
//	IpServiceArray{ IpServiceArgs{...} }
type IpServiceArrayInput interface {
	pulumi.Input

	ToIpServiceArrayOutput() IpServiceArrayOutput
	ToIpServiceArrayOutputWithContext(context.Context) IpServiceArrayOutput
}

type IpServiceArray []IpServiceInput

func (IpServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpService)(nil)).Elem()
}

func (i IpServiceArray) ToIpServiceArrayOutput() IpServiceArrayOutput {
	return i.ToIpServiceArrayOutputWithContext(context.Background())
}

func (i IpServiceArray) ToIpServiceArrayOutputWithContext(ctx context.Context) IpServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpServiceArrayOutput)
}

// IpServiceMapInput is an input type that accepts IpServiceMap and IpServiceMapOutput values.
// You can construct a concrete instance of `IpServiceMapInput` via:
//
//	IpServiceMap{ "key": IpServiceArgs{...} }
type IpServiceMapInput interface {
	pulumi.Input

	ToIpServiceMapOutput() IpServiceMapOutput
	ToIpServiceMapOutputWithContext(context.Context) IpServiceMapOutput
}

type IpServiceMap map[string]IpServiceInput

func (IpServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpService)(nil)).Elem()
}

func (i IpServiceMap) ToIpServiceMapOutput() IpServiceMapOutput {
	return i.ToIpServiceMapOutputWithContext(context.Background())
}

func (i IpServiceMap) ToIpServiceMapOutputWithContext(ctx context.Context) IpServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpServiceMapOutput)
}

type IpServiceOutput struct{ *pulumi.OutputState }

func (IpServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpService)(nil)).Elem()
}

func (o IpServiceOutput) ToIpServiceOutput() IpServiceOutput {
	return o
}

func (o IpServiceOutput) ToIpServiceOutputWithContext(ctx context.Context) IpServiceOutput {
	return o
}

// can be terminated
func (o IpServiceOutput) CanBeTerminated() pulumi.BoolOutput {
	return o.ApplyT(func(v *IpService) pulumi.BoolOutput { return v.CanBeTerminated }).(pulumi.BoolOutput)
}

// country
func (o IpServiceOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v *IpService) pulumi.StringOutput { return v.Country }).(pulumi.StringOutput)
}

// Custom description on your ip.
func (o IpServiceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *IpService) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// ip block
func (o IpServiceOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *IpService) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// Details about an Order
func (o IpServiceOutput) Orders() IpServiceOrderArrayOutput {
	return o.ApplyT(func(v *IpService) IpServiceOrderArrayOutput { return v.Orders }).(IpServiceOrderArrayOutput)
}

// IP block organisation Id
func (o IpServiceOutput) OrganisationId() pulumi.StringOutput {
	return o.ApplyT(func(v *IpService) pulumi.StringOutput { return v.OrganisationId }).(pulumi.StringOutput)
}

// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
func (o IpServiceOutput) OvhSubsidiary() pulumi.StringOutput {
	return o.ApplyT(func(v *IpService) pulumi.StringOutput { return v.OvhSubsidiary }).(pulumi.StringOutput)
}

// Ovh payment mode
//
// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
func (o IpServiceOutput) PaymentMean() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpService) pulumi.StringPtrOutput { return v.PaymentMean }).(pulumi.StringPtrOutput)
}

// Product Plan to order
func (o IpServiceOutput) Plan() IpServicePlanOutput {
	return o.ApplyT(func(v *IpService) IpServicePlanOutput { return v.Plan }).(IpServicePlanOutput)
}

// Product Plan to order
func (o IpServiceOutput) PlanOptions() IpServicePlanOptionArrayOutput {
	return o.ApplyT(func(v *IpService) IpServicePlanOptionArrayOutput { return v.PlanOptions }).(IpServicePlanOptionArrayOutput)
}

// Routage information
func (o IpServiceOutput) RoutedTos() IpServiceRoutedToArrayOutput {
	return o.ApplyT(func(v *IpService) IpServiceRoutedToArrayOutput { return v.RoutedTos }).(IpServiceRoutedToArrayOutput)
}

// service name
func (o IpServiceOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpService) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Possible values for ip type
func (o IpServiceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *IpService) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type IpServiceArrayOutput struct{ *pulumi.OutputState }

func (IpServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpService)(nil)).Elem()
}

func (o IpServiceArrayOutput) ToIpServiceArrayOutput() IpServiceArrayOutput {
	return o
}

func (o IpServiceArrayOutput) ToIpServiceArrayOutputWithContext(ctx context.Context) IpServiceArrayOutput {
	return o
}

func (o IpServiceArrayOutput) Index(i pulumi.IntInput) IpServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpService {
		return vs[0].([]*IpService)[vs[1].(int)]
	}).(IpServiceOutput)
}

type IpServiceMapOutput struct{ *pulumi.OutputState }

func (IpServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpService)(nil)).Elem()
}

func (o IpServiceMapOutput) ToIpServiceMapOutput() IpServiceMapOutput {
	return o
}

func (o IpServiceMapOutput) ToIpServiceMapOutputWithContext(ctx context.Context) IpServiceMapOutput {
	return o
}

func (o IpServiceMapOutput) MapIndex(k pulumi.StringInput) IpServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpService {
		return vs[0].(map[string]*IpService)[vs[1].(string)]
	}).(IpServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpServiceInput)(nil)).Elem(), &IpService{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpServiceArrayInput)(nil)).Elem(), IpServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpServiceMapInput)(nil)).Elem(), IpServiceMap{})
	pulumi.RegisterOutputType(IpServiceOutput{})
	pulumi.RegisterOutputType(IpServiceArrayOutput{})
	pulumi.RegisterOutputType(IpServiceMapOutput{})
}
