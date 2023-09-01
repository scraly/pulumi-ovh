// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this resource to create a hardware raid group in the partition scheme of a custom installation template available for dedicated servers.
//
// ## Example Usage
//
// ## Import
//
// The resource can be imported using the `template_name`, `scheme_name`, `name` of the cluster, separated by "/" E.g., bash <break><break>```sh<break> $ pulumi import ovh:Me/installationTemplatePartitionSchemeHardwareRaid:InstallationTemplatePartitionSchemeHardwareRaid group1 template_name/scheme_name/name <break>```<break><break>
type InstallationTemplatePartitionSchemeHardwareRaid struct {
	pulumi.CustomResourceState

	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
	Disks pulumi.StringArrayOutput `pulumi:"disks"`
	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
	Mode pulumi.StringOutput `pulumi:"mode"`
	// Hardware RAID name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The partition scheme name.
	SchemeName pulumi.StringOutput `pulumi:"schemeName"`
	// Specifies the creation order of the hardware RAID.
	Step pulumi.IntOutput `pulumi:"step"`
	// The template name of the partition scheme.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
}

// NewInstallationTemplatePartitionSchemeHardwareRaid registers a new resource with the given unique name, arguments, and options.
func NewInstallationTemplatePartitionSchemeHardwareRaid(ctx *pulumi.Context,
	name string, args *InstallationTemplatePartitionSchemeHardwareRaidArgs, opts ...pulumi.ResourceOption) (*InstallationTemplatePartitionSchemeHardwareRaid, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Disks == nil {
		return nil, errors.New("invalid value for required argument 'Disks'")
	}
	if args.Mode == nil {
		return nil, errors.New("invalid value for required argument 'Mode'")
	}
	if args.SchemeName == nil {
		return nil, errors.New("invalid value for required argument 'SchemeName'")
	}
	if args.Step == nil {
		return nil, errors.New("invalid value for required argument 'Step'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstallationTemplatePartitionSchemeHardwareRaid
	err := ctx.RegisterResource("ovh:Me/installationTemplatePartitionSchemeHardwareRaid:InstallationTemplatePartitionSchemeHardwareRaid", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstallationTemplatePartitionSchemeHardwareRaid gets an existing InstallationTemplatePartitionSchemeHardwareRaid resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstallationTemplatePartitionSchemeHardwareRaid(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstallationTemplatePartitionSchemeHardwareRaidState, opts ...pulumi.ResourceOption) (*InstallationTemplatePartitionSchemeHardwareRaid, error) {
	var resource InstallationTemplatePartitionSchemeHardwareRaid
	err := ctx.ReadResource("ovh:Me/installationTemplatePartitionSchemeHardwareRaid:InstallationTemplatePartitionSchemeHardwareRaid", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstallationTemplatePartitionSchemeHardwareRaid resources.
type installationTemplatePartitionSchemeHardwareRaidState struct {
	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
	Disks []string `pulumi:"disks"`
	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
	Mode *string `pulumi:"mode"`
	// Hardware RAID name.
	Name *string `pulumi:"name"`
	// The partition scheme name.
	SchemeName *string `pulumi:"schemeName"`
	// Specifies the creation order of the hardware RAID.
	Step *int `pulumi:"step"`
	// The template name of the partition scheme.
	TemplateName *string `pulumi:"templateName"`
}

type InstallationTemplatePartitionSchemeHardwareRaidState struct {
	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
	Disks pulumi.StringArrayInput
	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
	Mode pulumi.StringPtrInput
	// Hardware RAID name.
	Name pulumi.StringPtrInput
	// The partition scheme name.
	SchemeName pulumi.StringPtrInput
	// Specifies the creation order of the hardware RAID.
	Step pulumi.IntPtrInput
	// The template name of the partition scheme.
	TemplateName pulumi.StringPtrInput
}

func (InstallationTemplatePartitionSchemeHardwareRaidState) ElementType() reflect.Type {
	return reflect.TypeOf((*installationTemplatePartitionSchemeHardwareRaidState)(nil)).Elem()
}

type installationTemplatePartitionSchemeHardwareRaidArgs struct {
	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
	Disks []string `pulumi:"disks"`
	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
	Mode string `pulumi:"mode"`
	// Hardware RAID name.
	Name *string `pulumi:"name"`
	// The partition scheme name.
	SchemeName string `pulumi:"schemeName"`
	// Specifies the creation order of the hardware RAID.
	Step int `pulumi:"step"`
	// The template name of the partition scheme.
	TemplateName string `pulumi:"templateName"`
}

// The set of arguments for constructing a InstallationTemplatePartitionSchemeHardwareRaid resource.
type InstallationTemplatePartitionSchemeHardwareRaidArgs struct {
	// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
	Disks pulumi.StringArrayInput
	// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
	Mode pulumi.StringInput
	// Hardware RAID name.
	Name pulumi.StringPtrInput
	// The partition scheme name.
	SchemeName pulumi.StringInput
	// Specifies the creation order of the hardware RAID.
	Step pulumi.IntInput
	// The template name of the partition scheme.
	TemplateName pulumi.StringInput
}

func (InstallationTemplatePartitionSchemeHardwareRaidArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*installationTemplatePartitionSchemeHardwareRaidArgs)(nil)).Elem()
}

type InstallationTemplatePartitionSchemeHardwareRaidInput interface {
	pulumi.Input

	ToInstallationTemplatePartitionSchemeHardwareRaidOutput() InstallationTemplatePartitionSchemeHardwareRaidOutput
	ToInstallationTemplatePartitionSchemeHardwareRaidOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemeHardwareRaidOutput
}

func (*InstallationTemplatePartitionSchemeHardwareRaid) ElementType() reflect.Type {
	return reflect.TypeOf((**InstallationTemplatePartitionSchemeHardwareRaid)(nil)).Elem()
}

func (i *InstallationTemplatePartitionSchemeHardwareRaid) ToInstallationTemplatePartitionSchemeHardwareRaidOutput() InstallationTemplatePartitionSchemeHardwareRaidOutput {
	return i.ToInstallationTemplatePartitionSchemeHardwareRaidOutputWithContext(context.Background())
}

func (i *InstallationTemplatePartitionSchemeHardwareRaid) ToInstallationTemplatePartitionSchemeHardwareRaidOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemeHardwareRaidOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstallationTemplatePartitionSchemeHardwareRaidOutput)
}

// InstallationTemplatePartitionSchemeHardwareRaidArrayInput is an input type that accepts InstallationTemplatePartitionSchemeHardwareRaidArray and InstallationTemplatePartitionSchemeHardwareRaidArrayOutput values.
// You can construct a concrete instance of `InstallationTemplatePartitionSchemeHardwareRaidArrayInput` via:
//
//	InstallationTemplatePartitionSchemeHardwareRaidArray{ InstallationTemplatePartitionSchemeHardwareRaidArgs{...} }
type InstallationTemplatePartitionSchemeHardwareRaidArrayInput interface {
	pulumi.Input

	ToInstallationTemplatePartitionSchemeHardwareRaidArrayOutput() InstallationTemplatePartitionSchemeHardwareRaidArrayOutput
	ToInstallationTemplatePartitionSchemeHardwareRaidArrayOutputWithContext(context.Context) InstallationTemplatePartitionSchemeHardwareRaidArrayOutput
}

type InstallationTemplatePartitionSchemeHardwareRaidArray []InstallationTemplatePartitionSchemeHardwareRaidInput

func (InstallationTemplatePartitionSchemeHardwareRaidArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstallationTemplatePartitionSchemeHardwareRaid)(nil)).Elem()
}

func (i InstallationTemplatePartitionSchemeHardwareRaidArray) ToInstallationTemplatePartitionSchemeHardwareRaidArrayOutput() InstallationTemplatePartitionSchemeHardwareRaidArrayOutput {
	return i.ToInstallationTemplatePartitionSchemeHardwareRaidArrayOutputWithContext(context.Background())
}

func (i InstallationTemplatePartitionSchemeHardwareRaidArray) ToInstallationTemplatePartitionSchemeHardwareRaidArrayOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemeHardwareRaidArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstallationTemplatePartitionSchemeHardwareRaidArrayOutput)
}

// InstallationTemplatePartitionSchemeHardwareRaidMapInput is an input type that accepts InstallationTemplatePartitionSchemeHardwareRaidMap and InstallationTemplatePartitionSchemeHardwareRaidMapOutput values.
// You can construct a concrete instance of `InstallationTemplatePartitionSchemeHardwareRaidMapInput` via:
//
//	InstallationTemplatePartitionSchemeHardwareRaidMap{ "key": InstallationTemplatePartitionSchemeHardwareRaidArgs{...} }
type InstallationTemplatePartitionSchemeHardwareRaidMapInput interface {
	pulumi.Input

	ToInstallationTemplatePartitionSchemeHardwareRaidMapOutput() InstallationTemplatePartitionSchemeHardwareRaidMapOutput
	ToInstallationTemplatePartitionSchemeHardwareRaidMapOutputWithContext(context.Context) InstallationTemplatePartitionSchemeHardwareRaidMapOutput
}

type InstallationTemplatePartitionSchemeHardwareRaidMap map[string]InstallationTemplatePartitionSchemeHardwareRaidInput

func (InstallationTemplatePartitionSchemeHardwareRaidMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstallationTemplatePartitionSchemeHardwareRaid)(nil)).Elem()
}

func (i InstallationTemplatePartitionSchemeHardwareRaidMap) ToInstallationTemplatePartitionSchemeHardwareRaidMapOutput() InstallationTemplatePartitionSchemeHardwareRaidMapOutput {
	return i.ToInstallationTemplatePartitionSchemeHardwareRaidMapOutputWithContext(context.Background())
}

func (i InstallationTemplatePartitionSchemeHardwareRaidMap) ToInstallationTemplatePartitionSchemeHardwareRaidMapOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemeHardwareRaidMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstallationTemplatePartitionSchemeHardwareRaidMapOutput)
}

type InstallationTemplatePartitionSchemeHardwareRaidOutput struct{ *pulumi.OutputState }

func (InstallationTemplatePartitionSchemeHardwareRaidOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstallationTemplatePartitionSchemeHardwareRaid)(nil)).Elem()
}

func (o InstallationTemplatePartitionSchemeHardwareRaidOutput) ToInstallationTemplatePartitionSchemeHardwareRaidOutput() InstallationTemplatePartitionSchemeHardwareRaidOutput {
	return o
}

func (o InstallationTemplatePartitionSchemeHardwareRaidOutput) ToInstallationTemplatePartitionSchemeHardwareRaidOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemeHardwareRaidOutput {
	return o
}

// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
func (o InstallationTemplatePartitionSchemeHardwareRaidOutput) Disks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemeHardwareRaid) pulumi.StringArrayOutput { return v.Disks }).(pulumi.StringArrayOutput)
}

// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
func (o InstallationTemplatePartitionSchemeHardwareRaidOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemeHardwareRaid) pulumi.StringOutput { return v.Mode }).(pulumi.StringOutput)
}

// Hardware RAID name.
func (o InstallationTemplatePartitionSchemeHardwareRaidOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemeHardwareRaid) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The partition scheme name.
func (o InstallationTemplatePartitionSchemeHardwareRaidOutput) SchemeName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemeHardwareRaid) pulumi.StringOutput { return v.SchemeName }).(pulumi.StringOutput)
}

// Specifies the creation order of the hardware RAID.
func (o InstallationTemplatePartitionSchemeHardwareRaidOutput) Step() pulumi.IntOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemeHardwareRaid) pulumi.IntOutput { return v.Step }).(pulumi.IntOutput)
}

// The template name of the partition scheme.
func (o InstallationTemplatePartitionSchemeHardwareRaidOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemeHardwareRaid) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

type InstallationTemplatePartitionSchemeHardwareRaidArrayOutput struct{ *pulumi.OutputState }

func (InstallationTemplatePartitionSchemeHardwareRaidArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstallationTemplatePartitionSchemeHardwareRaid)(nil)).Elem()
}

func (o InstallationTemplatePartitionSchemeHardwareRaidArrayOutput) ToInstallationTemplatePartitionSchemeHardwareRaidArrayOutput() InstallationTemplatePartitionSchemeHardwareRaidArrayOutput {
	return o
}

func (o InstallationTemplatePartitionSchemeHardwareRaidArrayOutput) ToInstallationTemplatePartitionSchemeHardwareRaidArrayOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemeHardwareRaidArrayOutput {
	return o
}

func (o InstallationTemplatePartitionSchemeHardwareRaidArrayOutput) Index(i pulumi.IntInput) InstallationTemplatePartitionSchemeHardwareRaidOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstallationTemplatePartitionSchemeHardwareRaid {
		return vs[0].([]*InstallationTemplatePartitionSchemeHardwareRaid)[vs[1].(int)]
	}).(InstallationTemplatePartitionSchemeHardwareRaidOutput)
}

type InstallationTemplatePartitionSchemeHardwareRaidMapOutput struct{ *pulumi.OutputState }

func (InstallationTemplatePartitionSchemeHardwareRaidMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstallationTemplatePartitionSchemeHardwareRaid)(nil)).Elem()
}

func (o InstallationTemplatePartitionSchemeHardwareRaidMapOutput) ToInstallationTemplatePartitionSchemeHardwareRaidMapOutput() InstallationTemplatePartitionSchemeHardwareRaidMapOutput {
	return o
}

func (o InstallationTemplatePartitionSchemeHardwareRaidMapOutput) ToInstallationTemplatePartitionSchemeHardwareRaidMapOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemeHardwareRaidMapOutput {
	return o
}

func (o InstallationTemplatePartitionSchemeHardwareRaidMapOutput) MapIndex(k pulumi.StringInput) InstallationTemplatePartitionSchemeHardwareRaidOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstallationTemplatePartitionSchemeHardwareRaid {
		return vs[0].(map[string]*InstallationTemplatePartitionSchemeHardwareRaid)[vs[1].(string)]
	}).(InstallationTemplatePartitionSchemeHardwareRaidOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstallationTemplatePartitionSchemeHardwareRaidInput)(nil)).Elem(), &InstallationTemplatePartitionSchemeHardwareRaid{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstallationTemplatePartitionSchemeHardwareRaidArrayInput)(nil)).Elem(), InstallationTemplatePartitionSchemeHardwareRaidArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstallationTemplatePartitionSchemeHardwareRaidMapInput)(nil)).Elem(), InstallationTemplatePartitionSchemeHardwareRaidMap{})
	pulumi.RegisterOutputType(InstallationTemplatePartitionSchemeHardwareRaidOutput{})
	pulumi.RegisterOutputType(InstallationTemplatePartitionSchemeHardwareRaidArrayOutput{})
	pulumi.RegisterOutputType(InstallationTemplatePartitionSchemeHardwareRaidMapOutput{})
}
