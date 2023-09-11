// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hosting

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
//
// ## Import
//
// OVHcloud Webhosting database can be imported using the `service_name`, E.g.,
type PrivateDatabase struct {
	pulumi.CustomResourceState

	// Number of CPU on your private database
	Cpu pulumi.IntOutput `pulumi:"cpu"`
	// Datacenter where this private database is located
	Datacenter pulumi.StringOutput `pulumi:"datacenter"`
	// Name displayed in customer panel for your private database
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Private database hostname
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Private database FTP hostname
	HostnameFtp pulumi.StringOutput `pulumi:"hostnameFtp"`
	// Infrastructure where service was stored
	Infrastructure pulumi.StringOutput `pulumi:"infrastructure"`
	// Type of the private database offer
	Offer pulumi.StringOutput `pulumi:"offer"`
	// Details about your Order
	Orders PrivateDatabaseOrderArrayOutput `pulumi:"orders"`
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary pulumi.StringOutput `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrOutput `pulumi:"paymentMean"`
	// Product Plan to order
	Plan PrivateDatabasePlanOutput `pulumi:"plan"`
	// Product Plan to order
	PlanOptions PrivateDatabasePlanOptionArrayOutput `pulumi:"planOptions"`
	// Private database service port
	Port pulumi.IntOutput `pulumi:"port"`
	// Private database FTP port
	PortFtp pulumi.IntOutput `pulumi:"portFtp"`
	// Space allowed (in MB) on your private database
	QuotaSize pulumi.IntOutput `pulumi:"quotaSize"`
	// Sapce used (in MB) on your private database
	QuotaUsed pulumi.IntOutput `pulumi:"quotaUsed"`
	// Amount of ram (in MB) on your private database
	Ram pulumi.IntOutput `pulumi:"ram"`
	// Private database server name
	Server pulumi.StringOutput `pulumi:"server"`
	// Service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Private database state
	State pulumi.StringOutput `pulumi:"state"`
	// Private database type
	Type pulumi.StringOutput `pulumi:"type"`
	// URN of the private database, used when writing IAM policies
	Urn pulumi.StringOutput `pulumi:"urn"`
	// Private database available versions
	Version pulumi.StringOutput `pulumi:"version"`
	// Private database version label
	VersionLabel pulumi.StringOutput `pulumi:"versionLabel"`
	// Private database version number
	VersionNumber pulumi.Float64Output `pulumi:"versionNumber"`
}

// NewPrivateDatabase registers a new resource with the given unique name, arguments, and options.
func NewPrivateDatabase(ctx *pulumi.Context,
	name string, args *PrivateDatabaseArgs, opts ...pulumi.ResourceOption) (*PrivateDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OvhSubsidiary == nil {
		return nil, errors.New("invalid value for required argument 'OvhSubsidiary'")
	}
	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PrivateDatabase
	err := ctx.RegisterResource("ovh:Hosting/privateDatabase:PrivateDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateDatabase gets an existing PrivateDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateDatabaseState, opts ...pulumi.ResourceOption) (*PrivateDatabase, error) {
	var resource PrivateDatabase
	err := ctx.ReadResource("ovh:Hosting/privateDatabase:PrivateDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateDatabase resources.
type privateDatabaseState struct {
	// Number of CPU on your private database
	Cpu *int `pulumi:"cpu"`
	// Datacenter where this private database is located
	Datacenter *string `pulumi:"datacenter"`
	// Name displayed in customer panel for your private database
	DisplayName *string `pulumi:"displayName"`
	// Private database hostname
	Hostname *string `pulumi:"hostname"`
	// Private database FTP hostname
	HostnameFtp *string `pulumi:"hostnameFtp"`
	// Infrastructure where service was stored
	Infrastructure *string `pulumi:"infrastructure"`
	// Type of the private database offer
	Offer *string `pulumi:"offer"`
	// Details about your Order
	Orders []PrivateDatabaseOrder `pulumi:"orders"`
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *PrivateDatabasePlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []PrivateDatabasePlanOption `pulumi:"planOptions"`
	// Private database service port
	Port *int `pulumi:"port"`
	// Private database FTP port
	PortFtp *int `pulumi:"portFtp"`
	// Space allowed (in MB) on your private database
	QuotaSize *int `pulumi:"quotaSize"`
	// Sapce used (in MB) on your private database
	QuotaUsed *int `pulumi:"quotaUsed"`
	// Amount of ram (in MB) on your private database
	Ram *int `pulumi:"ram"`
	// Private database server name
	Server *string `pulumi:"server"`
	// Service name
	ServiceName *string `pulumi:"serviceName"`
	// Private database state
	State *string `pulumi:"state"`
	// Private database type
	Type *string `pulumi:"type"`
	// URN of the private database, used when writing IAM policies
	Urn *string `pulumi:"urn"`
	// Private database available versions
	Version *string `pulumi:"version"`
	// Private database version label
	VersionLabel *string `pulumi:"versionLabel"`
	// Private database version number
	VersionNumber *float64 `pulumi:"versionNumber"`
}

type PrivateDatabaseState struct {
	// Number of CPU on your private database
	Cpu pulumi.IntPtrInput
	// Datacenter where this private database is located
	Datacenter pulumi.StringPtrInput
	// Name displayed in customer panel for your private database
	DisplayName pulumi.StringPtrInput
	// Private database hostname
	Hostname pulumi.StringPtrInput
	// Private database FTP hostname
	HostnameFtp pulumi.StringPtrInput
	// Infrastructure where service was stored
	Infrastructure pulumi.StringPtrInput
	// Type of the private database offer
	Offer pulumi.StringPtrInput
	// Details about your Order
	Orders PrivateDatabaseOrderArrayInput
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan PrivateDatabasePlanPtrInput
	// Product Plan to order
	PlanOptions PrivateDatabasePlanOptionArrayInput
	// Private database service port
	Port pulumi.IntPtrInput
	// Private database FTP port
	PortFtp pulumi.IntPtrInput
	// Space allowed (in MB) on your private database
	QuotaSize pulumi.IntPtrInput
	// Sapce used (in MB) on your private database
	QuotaUsed pulumi.IntPtrInput
	// Amount of ram (in MB) on your private database
	Ram pulumi.IntPtrInput
	// Private database server name
	Server pulumi.StringPtrInput
	// Service name
	ServiceName pulumi.StringPtrInput
	// Private database state
	State pulumi.StringPtrInput
	// Private database type
	Type pulumi.StringPtrInput
	// URN of the private database, used when writing IAM policies
	Urn pulumi.StringPtrInput
	// Private database available versions
	Version pulumi.StringPtrInput
	// Private database version label
	VersionLabel pulumi.StringPtrInput
	// Private database version number
	VersionNumber pulumi.Float64PtrInput
}

func (PrivateDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDatabaseState)(nil)).Elem()
}

type privateDatabaseArgs struct {
	// Name displayed in customer panel for your private database
	DisplayName *string `pulumi:"displayName"`
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan PrivateDatabasePlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []PrivateDatabasePlanOption `pulumi:"planOptions"`
	// Service name
	ServiceName *string `pulumi:"serviceName"`
}

// The set of arguments for constructing a PrivateDatabase resource.
type PrivateDatabaseArgs struct {
	// Name displayed in customer panel for your private database
	DisplayName pulumi.StringPtrInput
	// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
	OvhSubsidiary pulumi.StringInput
	// Ovh payment mode
	//
	// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan PrivateDatabasePlanInput
	// Product Plan to order
	PlanOptions PrivateDatabasePlanOptionArrayInput
	// Service name
	ServiceName pulumi.StringPtrInput
}

func (PrivateDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDatabaseArgs)(nil)).Elem()
}

type PrivateDatabaseInput interface {
	pulumi.Input

	ToPrivateDatabaseOutput() PrivateDatabaseOutput
	ToPrivateDatabaseOutputWithContext(ctx context.Context) PrivateDatabaseOutput
}

func (*PrivateDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDatabase)(nil)).Elem()
}

func (i *PrivateDatabase) ToPrivateDatabaseOutput() PrivateDatabaseOutput {
	return i.ToPrivateDatabaseOutputWithContext(context.Background())
}

func (i *PrivateDatabase) ToPrivateDatabaseOutputWithContext(ctx context.Context) PrivateDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseOutput)
}

func (i *PrivateDatabase) ToOutput(ctx context.Context) pulumix.Output[*PrivateDatabase] {
	return pulumix.Output[*PrivateDatabase]{
		OutputState: i.ToPrivateDatabaseOutputWithContext(ctx).OutputState,
	}
}

// PrivateDatabaseArrayInput is an input type that accepts PrivateDatabaseArray and PrivateDatabaseArrayOutput values.
// You can construct a concrete instance of `PrivateDatabaseArrayInput` via:
//
//	PrivateDatabaseArray{ PrivateDatabaseArgs{...} }
type PrivateDatabaseArrayInput interface {
	pulumi.Input

	ToPrivateDatabaseArrayOutput() PrivateDatabaseArrayOutput
	ToPrivateDatabaseArrayOutputWithContext(context.Context) PrivateDatabaseArrayOutput
}

type PrivateDatabaseArray []PrivateDatabaseInput

func (PrivateDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateDatabase)(nil)).Elem()
}

func (i PrivateDatabaseArray) ToPrivateDatabaseArrayOutput() PrivateDatabaseArrayOutput {
	return i.ToPrivateDatabaseArrayOutputWithContext(context.Background())
}

func (i PrivateDatabaseArray) ToPrivateDatabaseArrayOutputWithContext(ctx context.Context) PrivateDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseArrayOutput)
}

func (i PrivateDatabaseArray) ToOutput(ctx context.Context) pulumix.Output[[]*PrivateDatabase] {
	return pulumix.Output[[]*PrivateDatabase]{
		OutputState: i.ToPrivateDatabaseArrayOutputWithContext(ctx).OutputState,
	}
}

// PrivateDatabaseMapInput is an input type that accepts PrivateDatabaseMap and PrivateDatabaseMapOutput values.
// You can construct a concrete instance of `PrivateDatabaseMapInput` via:
//
//	PrivateDatabaseMap{ "key": PrivateDatabaseArgs{...} }
type PrivateDatabaseMapInput interface {
	pulumi.Input

	ToPrivateDatabaseMapOutput() PrivateDatabaseMapOutput
	ToPrivateDatabaseMapOutputWithContext(context.Context) PrivateDatabaseMapOutput
}

type PrivateDatabaseMap map[string]PrivateDatabaseInput

func (PrivateDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateDatabase)(nil)).Elem()
}

func (i PrivateDatabaseMap) ToPrivateDatabaseMapOutput() PrivateDatabaseMapOutput {
	return i.ToPrivateDatabaseMapOutputWithContext(context.Background())
}

func (i PrivateDatabaseMap) ToPrivateDatabaseMapOutputWithContext(ctx context.Context) PrivateDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateDatabaseMapOutput)
}

func (i PrivateDatabaseMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*PrivateDatabase] {
	return pulumix.Output[map[string]*PrivateDatabase]{
		OutputState: i.ToPrivateDatabaseMapOutputWithContext(ctx).OutputState,
	}
}

type PrivateDatabaseOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateDatabase)(nil)).Elem()
}

func (o PrivateDatabaseOutput) ToPrivateDatabaseOutput() PrivateDatabaseOutput {
	return o
}

func (o PrivateDatabaseOutput) ToPrivateDatabaseOutputWithContext(ctx context.Context) PrivateDatabaseOutput {
	return o
}

func (o PrivateDatabaseOutput) ToOutput(ctx context.Context) pulumix.Output[*PrivateDatabase] {
	return pulumix.Output[*PrivateDatabase]{
		OutputState: o.OutputState,
	}
}

// Number of CPU on your private database
func (o PrivateDatabaseOutput) Cpu() pulumi.IntOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.IntOutput { return v.Cpu }).(pulumi.IntOutput)
}

// Datacenter where this private database is located
func (o PrivateDatabaseOutput) Datacenter() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.Datacenter }).(pulumi.StringOutput)
}

// Name displayed in customer panel for your private database
func (o PrivateDatabaseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Private database hostname
func (o PrivateDatabaseOutput) Hostname() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.Hostname }).(pulumi.StringOutput)
}

// Private database FTP hostname
func (o PrivateDatabaseOutput) HostnameFtp() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.HostnameFtp }).(pulumi.StringOutput)
}

// Infrastructure where service was stored
func (o PrivateDatabaseOutput) Infrastructure() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.Infrastructure }).(pulumi.StringOutput)
}

// Type of the private database offer
func (o PrivateDatabaseOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.Offer }).(pulumi.StringOutput)
}

// Details about your Order
func (o PrivateDatabaseOutput) Orders() PrivateDatabaseOrderArrayOutput {
	return o.ApplyT(func(v *PrivateDatabase) PrivateDatabaseOrderArrayOutput { return v.Orders }).(PrivateDatabaseOrderArrayOutput)
}

// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
func (o PrivateDatabaseOutput) OvhSubsidiary() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.OvhSubsidiary }).(pulumi.StringOutput)
}

// Ovh payment mode
//
// Deprecated: This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
func (o PrivateDatabaseOutput) PaymentMean() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringPtrOutput { return v.PaymentMean }).(pulumi.StringPtrOutput)
}

// Product Plan to order
func (o PrivateDatabaseOutput) Plan() PrivateDatabasePlanOutput {
	return o.ApplyT(func(v *PrivateDatabase) PrivateDatabasePlanOutput { return v.Plan }).(PrivateDatabasePlanOutput)
}

// Product Plan to order
func (o PrivateDatabaseOutput) PlanOptions() PrivateDatabasePlanOptionArrayOutput {
	return o.ApplyT(func(v *PrivateDatabase) PrivateDatabasePlanOptionArrayOutput { return v.PlanOptions }).(PrivateDatabasePlanOptionArrayOutput)
}

// Private database service port
func (o PrivateDatabaseOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// Private database FTP port
func (o PrivateDatabaseOutput) PortFtp() pulumi.IntOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.IntOutput { return v.PortFtp }).(pulumi.IntOutput)
}

// Space allowed (in MB) on your private database
func (o PrivateDatabaseOutput) QuotaSize() pulumi.IntOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.IntOutput { return v.QuotaSize }).(pulumi.IntOutput)
}

// Sapce used (in MB) on your private database
func (o PrivateDatabaseOutput) QuotaUsed() pulumi.IntOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.IntOutput { return v.QuotaUsed }).(pulumi.IntOutput)
}

// Amount of ram (in MB) on your private database
func (o PrivateDatabaseOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.IntOutput { return v.Ram }).(pulumi.IntOutput)
}

// Private database server name
func (o PrivateDatabaseOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

// Service name
func (o PrivateDatabaseOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Private database state
func (o PrivateDatabaseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Private database type
func (o PrivateDatabaseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// URN of the private database, used when writing IAM policies
func (o PrivateDatabaseOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.Urn }).(pulumi.StringOutput)
}

// Private database available versions
func (o PrivateDatabaseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// Private database version label
func (o PrivateDatabaseOutput) VersionLabel() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.StringOutput { return v.VersionLabel }).(pulumi.StringOutput)
}

// Private database version number
func (o PrivateDatabaseOutput) VersionNumber() pulumi.Float64Output {
	return o.ApplyT(func(v *PrivateDatabase) pulumi.Float64Output { return v.VersionNumber }).(pulumi.Float64Output)
}

type PrivateDatabaseArrayOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PrivateDatabase)(nil)).Elem()
}

func (o PrivateDatabaseArrayOutput) ToPrivateDatabaseArrayOutput() PrivateDatabaseArrayOutput {
	return o
}

func (o PrivateDatabaseArrayOutput) ToPrivateDatabaseArrayOutputWithContext(ctx context.Context) PrivateDatabaseArrayOutput {
	return o
}

func (o PrivateDatabaseArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*PrivateDatabase] {
	return pulumix.Output[[]*PrivateDatabase]{
		OutputState: o.OutputState,
	}
}

func (o PrivateDatabaseArrayOutput) Index(i pulumi.IntInput) PrivateDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PrivateDatabase {
		return vs[0].([]*PrivateDatabase)[vs[1].(int)]
	}).(PrivateDatabaseOutput)
}

type PrivateDatabaseMapOutput struct{ *pulumi.OutputState }

func (PrivateDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PrivateDatabase)(nil)).Elem()
}

func (o PrivateDatabaseMapOutput) ToPrivateDatabaseMapOutput() PrivateDatabaseMapOutput {
	return o
}

func (o PrivateDatabaseMapOutput) ToPrivateDatabaseMapOutputWithContext(ctx context.Context) PrivateDatabaseMapOutput {
	return o
}

func (o PrivateDatabaseMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*PrivateDatabase] {
	return pulumix.Output[map[string]*PrivateDatabase]{
		OutputState: o.OutputState,
	}
}

func (o PrivateDatabaseMapOutput) MapIndex(k pulumi.StringInput) PrivateDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PrivateDatabase {
		return vs[0].(map[string]*PrivateDatabase)[vs[1].(string)]
	}).(PrivateDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseInput)(nil)).Elem(), &PrivateDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseArrayInput)(nil)).Elem(), PrivateDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateDatabaseMapInput)(nil)).Elem(), PrivateDatabaseMap{})
	pulumi.RegisterOutputType(PrivateDatabaseOutput{})
	pulumi.RegisterOutputType(PrivateDatabaseArrayOutput{})
	pulumi.RegisterOutputType(PrivateDatabaseMapOutput{})
}
