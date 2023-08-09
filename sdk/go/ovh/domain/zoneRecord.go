// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package domain

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ## Import
//
// OVHcloud domain zone record can be imported using the `id`, which can be retrieved by using [OVH API portal](https://api.ovh.com/console/#/domain/zone/%7BzoneName%7D/record~GET), and the `zone`, separated by "." E.g., bash <break><break>```sh<break> $ pulumi import ovh:Domain/zoneRecord:ZoneRecord test id.zone <break>```<break><break>
type ZoneRecord struct {
	pulumi.CustomResourceState

	// The type of the record
	Fieldtype pulumi.StringOutput `pulumi:"fieldtype"`
	// The name of the record
	Subdomain pulumi.StringPtrOutput `pulumi:"subdomain"`
	// The value of the record
	Target pulumi.StringOutput `pulumi:"target"`
	// The TTL of the record, it shall be >= to 60.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// The domain to add the record to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewZoneRecord registers a new resource with the given unique name, arguments, and options.
func NewZoneRecord(ctx *pulumi.Context,
	name string, args *ZoneRecordArgs, opts ...pulumi.ResourceOption) (*ZoneRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fieldtype == nil {
		return nil, errors.New("invalid value for required argument 'Fieldtype'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ZoneRecord
	err := ctx.RegisterResource("ovh:Domain/zoneRecord:ZoneRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZoneRecord gets an existing ZoneRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZoneRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZoneRecordState, opts ...pulumi.ResourceOption) (*ZoneRecord, error) {
	var resource ZoneRecord
	err := ctx.ReadResource("ovh:Domain/zoneRecord:ZoneRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZoneRecord resources.
type zoneRecordState struct {
	// The type of the record
	Fieldtype *string `pulumi:"fieldtype"`
	// The name of the record
	Subdomain *string `pulumi:"subdomain"`
	// The value of the record
	Target *string `pulumi:"target"`
	// The TTL of the record, it shall be >= to 60.
	Ttl *int `pulumi:"ttl"`
	// The domain to add the record to
	Zone *string `pulumi:"zone"`
}

type ZoneRecordState struct {
	// The type of the record
	Fieldtype pulumi.StringPtrInput
	// The name of the record
	Subdomain pulumi.StringPtrInput
	// The value of the record
	Target pulumi.StringPtrInput
	// The TTL of the record, it shall be >= to 60.
	Ttl pulumi.IntPtrInput
	// The domain to add the record to
	Zone pulumi.StringPtrInput
}

func (ZoneRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneRecordState)(nil)).Elem()
}

type zoneRecordArgs struct {
	// The type of the record
	Fieldtype string `pulumi:"fieldtype"`
	// The name of the record
	Subdomain *string `pulumi:"subdomain"`
	// The value of the record
	Target string `pulumi:"target"`
	// The TTL of the record, it shall be >= to 60.
	Ttl *int `pulumi:"ttl"`
	// The domain to add the record to
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a ZoneRecord resource.
type ZoneRecordArgs struct {
	// The type of the record
	Fieldtype pulumi.StringInput
	// The name of the record
	Subdomain pulumi.StringPtrInput
	// The value of the record
	Target pulumi.StringInput
	// The TTL of the record, it shall be >= to 60.
	Ttl pulumi.IntPtrInput
	// The domain to add the record to
	Zone pulumi.StringInput
}

func (ZoneRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zoneRecordArgs)(nil)).Elem()
}

type ZoneRecordInput interface {
	pulumi.Input

	ToZoneRecordOutput() ZoneRecordOutput
	ToZoneRecordOutputWithContext(ctx context.Context) ZoneRecordOutput
}

func (*ZoneRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneRecord)(nil)).Elem()
}

func (i *ZoneRecord) ToZoneRecordOutput() ZoneRecordOutput {
	return i.ToZoneRecordOutputWithContext(context.Background())
}

func (i *ZoneRecord) ToZoneRecordOutputWithContext(ctx context.Context) ZoneRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneRecordOutput)
}

// ZoneRecordArrayInput is an input type that accepts ZoneRecordArray and ZoneRecordArrayOutput values.
// You can construct a concrete instance of `ZoneRecordArrayInput` via:
//
//	ZoneRecordArray{ ZoneRecordArgs{...} }
type ZoneRecordArrayInput interface {
	pulumi.Input

	ToZoneRecordArrayOutput() ZoneRecordArrayOutput
	ToZoneRecordArrayOutputWithContext(context.Context) ZoneRecordArrayOutput
}

type ZoneRecordArray []ZoneRecordInput

func (ZoneRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneRecord)(nil)).Elem()
}

func (i ZoneRecordArray) ToZoneRecordArrayOutput() ZoneRecordArrayOutput {
	return i.ToZoneRecordArrayOutputWithContext(context.Background())
}

func (i ZoneRecordArray) ToZoneRecordArrayOutputWithContext(ctx context.Context) ZoneRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneRecordArrayOutput)
}

// ZoneRecordMapInput is an input type that accepts ZoneRecordMap and ZoneRecordMapOutput values.
// You can construct a concrete instance of `ZoneRecordMapInput` via:
//
//	ZoneRecordMap{ "key": ZoneRecordArgs{...} }
type ZoneRecordMapInput interface {
	pulumi.Input

	ToZoneRecordMapOutput() ZoneRecordMapOutput
	ToZoneRecordMapOutputWithContext(context.Context) ZoneRecordMapOutput
}

type ZoneRecordMap map[string]ZoneRecordInput

func (ZoneRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneRecord)(nil)).Elem()
}

func (i ZoneRecordMap) ToZoneRecordMapOutput() ZoneRecordMapOutput {
	return i.ToZoneRecordMapOutputWithContext(context.Background())
}

func (i ZoneRecordMap) ToZoneRecordMapOutputWithContext(ctx context.Context) ZoneRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZoneRecordMapOutput)
}

type ZoneRecordOutput struct{ *pulumi.OutputState }

func (ZoneRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZoneRecord)(nil)).Elem()
}

func (o ZoneRecordOutput) ToZoneRecordOutput() ZoneRecordOutput {
	return o
}

func (o ZoneRecordOutput) ToZoneRecordOutputWithContext(ctx context.Context) ZoneRecordOutput {
	return o
}

// The type of the record
func (o ZoneRecordOutput) Fieldtype() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringOutput { return v.Fieldtype }).(pulumi.StringOutput)
}

// The name of the record
func (o ZoneRecordOutput) Subdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringPtrOutput { return v.Subdomain }).(pulumi.StringPtrOutput)
}

// The value of the record
func (o ZoneRecordOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

// The TTL of the record, it shall be >= to 60.
func (o ZoneRecordOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The domain to add the record to
func (o ZoneRecordOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ZoneRecord) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type ZoneRecordArrayOutput struct{ *pulumi.OutputState }

func (ZoneRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ZoneRecord)(nil)).Elem()
}

func (o ZoneRecordArrayOutput) ToZoneRecordArrayOutput() ZoneRecordArrayOutput {
	return o
}

func (o ZoneRecordArrayOutput) ToZoneRecordArrayOutputWithContext(ctx context.Context) ZoneRecordArrayOutput {
	return o
}

func (o ZoneRecordArrayOutput) Index(i pulumi.IntInput) ZoneRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ZoneRecord {
		return vs[0].([]*ZoneRecord)[vs[1].(int)]
	}).(ZoneRecordOutput)
}

type ZoneRecordMapOutput struct{ *pulumi.OutputState }

func (ZoneRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ZoneRecord)(nil)).Elem()
}

func (o ZoneRecordMapOutput) ToZoneRecordMapOutput() ZoneRecordMapOutput {
	return o
}

func (o ZoneRecordMapOutput) ToZoneRecordMapOutputWithContext(ctx context.Context) ZoneRecordMapOutput {
	return o
}

func (o ZoneRecordMapOutput) MapIndex(k pulumi.StringInput) ZoneRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ZoneRecord {
		return vs[0].(map[string]*ZoneRecord)[vs[1].(string)]
	}).(ZoneRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneRecordInput)(nil)).Elem(), &ZoneRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneRecordArrayInput)(nil)).Elem(), ZoneRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ZoneRecordMapInput)(nil)).Elem(), ZoneRecordMap{})
	pulumi.RegisterOutputType(ZoneRecordOutput{})
	pulumi.RegisterOutputType(ZoneRecordArrayOutput{})
	pulumi.RegisterOutputType(ZoneRecordMapOutput{})
}