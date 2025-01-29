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

// Use this resource to create a partition in the partition scheme of a custom installation template available for dedicated servers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mytemplate, err := me.NewInstallationTemplate(ctx, "mytemplate", &me.InstallationTemplateArgs{
//				BaseTemplateName: pulumi.String("debian12_64"),
//				TemplateName:     pulumi.String("mytemplate"),
//			})
//			if err != nil {
//				return err
//			}
//			scheme, err := me.NewInstallationTemplatePartitionScheme(ctx, "scheme", &me.InstallationTemplatePartitionSchemeArgs{
//				TemplateName: mytemplate.TemplateName,
//				Priority:     pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = me.NewInstallationTemplatePartitionSchemePartition(ctx, "root", &me.InstallationTemplatePartitionSchemePartitionArgs{
//				TemplateName: scheme.TemplateName,
//				SchemeName:   scheme.Name,
//				Mountpoint:   pulumi.String("/"),
//				Filesystem:   pulumi.String("ext4"),
//				Size:         pulumi.Int(400),
//				Order:        pulumi.Int(1),
//				Type:         pulumi.String("primary"),
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
// The resource can be imported using the `template_name`, `scheme_name`, `mountpoint` of the cluster, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:Me/installationTemplatePartitionSchemePartition:InstallationTemplatePartitionSchemePartition root template_name/scheme_name/mountpoint
// ```
type InstallationTemplatePartitionSchemePartition struct {
	pulumi.CustomResourceState

	// Partition filesystem. Enum with possibles values:
	// - btrfs
	// - ext3
	// - ext4
	// - ntfs
	// - reiserfs
	// - swap
	// - ufs
	// - xfs
	// - zfs
	Filesystem pulumi.StringOutput `pulumi:"filesystem"`
	// partition mount point.
	Mountpoint pulumi.StringOutput `pulumi:"mountpoint"`
	// step or order. specifies the creation order of the partition on the disk
	Order pulumi.IntOutput `pulumi:"order"`
	// raid partition type. Enum with possible values:
	// - raid0
	// - raid1
	// - raid10
	// - raid5
	// - raid6
	Raid pulumi.StringOutput `pulumi:"raid"`
	// The partition scheme name.
	SchemeName pulumi.StringOutput `pulumi:"schemeName"`
	// size of partition in MB, 0 => rest of the space.
	Size pulumi.IntOutput `pulumi:"size"`
	// The template name of the partition scheme.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
	// partition type. Enum with possible values:
	// - lv
	// - primary
	// - logical
	Type pulumi.StringOutput `pulumi:"type"`
	// The volume name needed for proxmox distribution
	VolumeName pulumi.StringOutput `pulumi:"volumeName"`
}

// NewInstallationTemplatePartitionSchemePartition registers a new resource with the given unique name, arguments, and options.
func NewInstallationTemplatePartitionSchemePartition(ctx *pulumi.Context,
	name string, args *InstallationTemplatePartitionSchemePartitionArgs, opts ...pulumi.ResourceOption) (*InstallationTemplatePartitionSchemePartition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Filesystem == nil {
		return nil, errors.New("invalid value for required argument 'Filesystem'")
	}
	if args.Mountpoint == nil {
		return nil, errors.New("invalid value for required argument 'Mountpoint'")
	}
	if args.Order == nil {
		return nil, errors.New("invalid value for required argument 'Order'")
	}
	if args.SchemeName == nil {
		return nil, errors.New("invalid value for required argument 'SchemeName'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource InstallationTemplatePartitionSchemePartition
	err := ctx.RegisterResource("ovh:Me/installationTemplatePartitionSchemePartition:InstallationTemplatePartitionSchemePartition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstallationTemplatePartitionSchemePartition gets an existing InstallationTemplatePartitionSchemePartition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstallationTemplatePartitionSchemePartition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstallationTemplatePartitionSchemePartitionState, opts ...pulumi.ResourceOption) (*InstallationTemplatePartitionSchemePartition, error) {
	var resource InstallationTemplatePartitionSchemePartition
	err := ctx.ReadResource("ovh:Me/installationTemplatePartitionSchemePartition:InstallationTemplatePartitionSchemePartition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstallationTemplatePartitionSchemePartition resources.
type installationTemplatePartitionSchemePartitionState struct {
	// Partition filesystem. Enum with possibles values:
	// - btrfs
	// - ext3
	// - ext4
	// - ntfs
	// - reiserfs
	// - swap
	// - ufs
	// - xfs
	// - zfs
	Filesystem *string `pulumi:"filesystem"`
	// partition mount point.
	Mountpoint *string `pulumi:"mountpoint"`
	// step or order. specifies the creation order of the partition on the disk
	Order *int `pulumi:"order"`
	// raid partition type. Enum with possible values:
	// - raid0
	// - raid1
	// - raid10
	// - raid5
	// - raid6
	Raid *string `pulumi:"raid"`
	// The partition scheme name.
	SchemeName *string `pulumi:"schemeName"`
	// size of partition in MB, 0 => rest of the space.
	Size *int `pulumi:"size"`
	// The template name of the partition scheme.
	TemplateName *string `pulumi:"templateName"`
	// partition type. Enum with possible values:
	// - lv
	// - primary
	// - logical
	Type *string `pulumi:"type"`
	// The volume name needed for proxmox distribution
	VolumeName *string `pulumi:"volumeName"`
}

type InstallationTemplatePartitionSchemePartitionState struct {
	// Partition filesystem. Enum with possibles values:
	// - btrfs
	// - ext3
	// - ext4
	// - ntfs
	// - reiserfs
	// - swap
	// - ufs
	// - xfs
	// - zfs
	Filesystem pulumi.StringPtrInput
	// partition mount point.
	Mountpoint pulumi.StringPtrInput
	// step or order. specifies the creation order of the partition on the disk
	Order pulumi.IntPtrInput
	// raid partition type. Enum with possible values:
	// - raid0
	// - raid1
	// - raid10
	// - raid5
	// - raid6
	Raid pulumi.StringPtrInput
	// The partition scheme name.
	SchemeName pulumi.StringPtrInput
	// size of partition in MB, 0 => rest of the space.
	Size pulumi.IntPtrInput
	// The template name of the partition scheme.
	TemplateName pulumi.StringPtrInput
	// partition type. Enum with possible values:
	// - lv
	// - primary
	// - logical
	Type pulumi.StringPtrInput
	// The volume name needed for proxmox distribution
	VolumeName pulumi.StringPtrInput
}

func (InstallationTemplatePartitionSchemePartitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*installationTemplatePartitionSchemePartitionState)(nil)).Elem()
}

type installationTemplatePartitionSchemePartitionArgs struct {
	// Partition filesystem. Enum with possibles values:
	// - btrfs
	// - ext3
	// - ext4
	// - ntfs
	// - reiserfs
	// - swap
	// - ufs
	// - xfs
	// - zfs
	Filesystem string `pulumi:"filesystem"`
	// partition mount point.
	Mountpoint string `pulumi:"mountpoint"`
	// step or order. specifies the creation order of the partition on the disk
	Order int `pulumi:"order"`
	// raid partition type. Enum with possible values:
	// - raid0
	// - raid1
	// - raid10
	// - raid5
	// - raid6
	Raid *string `pulumi:"raid"`
	// The partition scheme name.
	SchemeName string `pulumi:"schemeName"`
	// size of partition in MB, 0 => rest of the space.
	Size int `pulumi:"size"`
	// The template name of the partition scheme.
	TemplateName string `pulumi:"templateName"`
	// partition type. Enum with possible values:
	// - lv
	// - primary
	// - logical
	Type string `pulumi:"type"`
	// The volume name needed for proxmox distribution
	VolumeName *string `pulumi:"volumeName"`
}

// The set of arguments for constructing a InstallationTemplatePartitionSchemePartition resource.
type InstallationTemplatePartitionSchemePartitionArgs struct {
	// Partition filesystem. Enum with possibles values:
	// - btrfs
	// - ext3
	// - ext4
	// - ntfs
	// - reiserfs
	// - swap
	// - ufs
	// - xfs
	// - zfs
	Filesystem pulumi.StringInput
	// partition mount point.
	Mountpoint pulumi.StringInput
	// step or order. specifies the creation order of the partition on the disk
	Order pulumi.IntInput
	// raid partition type. Enum with possible values:
	// - raid0
	// - raid1
	// - raid10
	// - raid5
	// - raid6
	Raid pulumi.StringPtrInput
	// The partition scheme name.
	SchemeName pulumi.StringInput
	// size of partition in MB, 0 => rest of the space.
	Size pulumi.IntInput
	// The template name of the partition scheme.
	TemplateName pulumi.StringInput
	// partition type. Enum with possible values:
	// - lv
	// - primary
	// - logical
	Type pulumi.StringInput
	// The volume name needed for proxmox distribution
	VolumeName pulumi.StringPtrInput
}

func (InstallationTemplatePartitionSchemePartitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*installationTemplatePartitionSchemePartitionArgs)(nil)).Elem()
}

type InstallationTemplatePartitionSchemePartitionInput interface {
	pulumi.Input

	ToInstallationTemplatePartitionSchemePartitionOutput() InstallationTemplatePartitionSchemePartitionOutput
	ToInstallationTemplatePartitionSchemePartitionOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemePartitionOutput
}

func (*InstallationTemplatePartitionSchemePartition) ElementType() reflect.Type {
	return reflect.TypeOf((**InstallationTemplatePartitionSchemePartition)(nil)).Elem()
}

func (i *InstallationTemplatePartitionSchemePartition) ToInstallationTemplatePartitionSchemePartitionOutput() InstallationTemplatePartitionSchemePartitionOutput {
	return i.ToInstallationTemplatePartitionSchemePartitionOutputWithContext(context.Background())
}

func (i *InstallationTemplatePartitionSchemePartition) ToInstallationTemplatePartitionSchemePartitionOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemePartitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstallationTemplatePartitionSchemePartitionOutput)
}

// InstallationTemplatePartitionSchemePartitionArrayInput is an input type that accepts InstallationTemplatePartitionSchemePartitionArray and InstallationTemplatePartitionSchemePartitionArrayOutput values.
// You can construct a concrete instance of `InstallationTemplatePartitionSchemePartitionArrayInput` via:
//
//	InstallationTemplatePartitionSchemePartitionArray{ InstallationTemplatePartitionSchemePartitionArgs{...} }
type InstallationTemplatePartitionSchemePartitionArrayInput interface {
	pulumi.Input

	ToInstallationTemplatePartitionSchemePartitionArrayOutput() InstallationTemplatePartitionSchemePartitionArrayOutput
	ToInstallationTemplatePartitionSchemePartitionArrayOutputWithContext(context.Context) InstallationTemplatePartitionSchemePartitionArrayOutput
}

type InstallationTemplatePartitionSchemePartitionArray []InstallationTemplatePartitionSchemePartitionInput

func (InstallationTemplatePartitionSchemePartitionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstallationTemplatePartitionSchemePartition)(nil)).Elem()
}

func (i InstallationTemplatePartitionSchemePartitionArray) ToInstallationTemplatePartitionSchemePartitionArrayOutput() InstallationTemplatePartitionSchemePartitionArrayOutput {
	return i.ToInstallationTemplatePartitionSchemePartitionArrayOutputWithContext(context.Background())
}

func (i InstallationTemplatePartitionSchemePartitionArray) ToInstallationTemplatePartitionSchemePartitionArrayOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemePartitionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstallationTemplatePartitionSchemePartitionArrayOutput)
}

// InstallationTemplatePartitionSchemePartitionMapInput is an input type that accepts InstallationTemplatePartitionSchemePartitionMap and InstallationTemplatePartitionSchemePartitionMapOutput values.
// You can construct a concrete instance of `InstallationTemplatePartitionSchemePartitionMapInput` via:
//
//	InstallationTemplatePartitionSchemePartitionMap{ "key": InstallationTemplatePartitionSchemePartitionArgs{...} }
type InstallationTemplatePartitionSchemePartitionMapInput interface {
	pulumi.Input

	ToInstallationTemplatePartitionSchemePartitionMapOutput() InstallationTemplatePartitionSchemePartitionMapOutput
	ToInstallationTemplatePartitionSchemePartitionMapOutputWithContext(context.Context) InstallationTemplatePartitionSchemePartitionMapOutput
}

type InstallationTemplatePartitionSchemePartitionMap map[string]InstallationTemplatePartitionSchemePartitionInput

func (InstallationTemplatePartitionSchemePartitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstallationTemplatePartitionSchemePartition)(nil)).Elem()
}

func (i InstallationTemplatePartitionSchemePartitionMap) ToInstallationTemplatePartitionSchemePartitionMapOutput() InstallationTemplatePartitionSchemePartitionMapOutput {
	return i.ToInstallationTemplatePartitionSchemePartitionMapOutputWithContext(context.Background())
}

func (i InstallationTemplatePartitionSchemePartitionMap) ToInstallationTemplatePartitionSchemePartitionMapOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemePartitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstallationTemplatePartitionSchemePartitionMapOutput)
}

type InstallationTemplatePartitionSchemePartitionOutput struct{ *pulumi.OutputState }

func (InstallationTemplatePartitionSchemePartitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstallationTemplatePartitionSchemePartition)(nil)).Elem()
}

func (o InstallationTemplatePartitionSchemePartitionOutput) ToInstallationTemplatePartitionSchemePartitionOutput() InstallationTemplatePartitionSchemePartitionOutput {
	return o
}

func (o InstallationTemplatePartitionSchemePartitionOutput) ToInstallationTemplatePartitionSchemePartitionOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemePartitionOutput {
	return o
}

// Partition filesystem. Enum with possibles values:
// - btrfs
// - ext3
// - ext4
// - ntfs
// - reiserfs
// - swap
// - ufs
// - xfs
// - zfs
func (o InstallationTemplatePartitionSchemePartitionOutput) Filesystem() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemePartition) pulumi.StringOutput { return v.Filesystem }).(pulumi.StringOutput)
}

// partition mount point.
func (o InstallationTemplatePartitionSchemePartitionOutput) Mountpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemePartition) pulumi.StringOutput { return v.Mountpoint }).(pulumi.StringOutput)
}

// step or order. specifies the creation order of the partition on the disk
func (o InstallationTemplatePartitionSchemePartitionOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemePartition) pulumi.IntOutput { return v.Order }).(pulumi.IntOutput)
}

// raid partition type. Enum with possible values:
// - raid0
// - raid1
// - raid10
// - raid5
// - raid6
func (o InstallationTemplatePartitionSchemePartitionOutput) Raid() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemePartition) pulumi.StringOutput { return v.Raid }).(pulumi.StringOutput)
}

// The partition scheme name.
func (o InstallationTemplatePartitionSchemePartitionOutput) SchemeName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemePartition) pulumi.StringOutput { return v.SchemeName }).(pulumi.StringOutput)
}

// size of partition in MB, 0 => rest of the space.
func (o InstallationTemplatePartitionSchemePartitionOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemePartition) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// The template name of the partition scheme.
func (o InstallationTemplatePartitionSchemePartitionOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemePartition) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

// partition type. Enum with possible values:
// - lv
// - primary
// - logical
func (o InstallationTemplatePartitionSchemePartitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemePartition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The volume name needed for proxmox distribution
func (o InstallationTemplatePartitionSchemePartitionOutput) VolumeName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstallationTemplatePartitionSchemePartition) pulumi.StringOutput { return v.VolumeName }).(pulumi.StringOutput)
}

type InstallationTemplatePartitionSchemePartitionArrayOutput struct{ *pulumi.OutputState }

func (InstallationTemplatePartitionSchemePartitionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstallationTemplatePartitionSchemePartition)(nil)).Elem()
}

func (o InstallationTemplatePartitionSchemePartitionArrayOutput) ToInstallationTemplatePartitionSchemePartitionArrayOutput() InstallationTemplatePartitionSchemePartitionArrayOutput {
	return o
}

func (o InstallationTemplatePartitionSchemePartitionArrayOutput) ToInstallationTemplatePartitionSchemePartitionArrayOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemePartitionArrayOutput {
	return o
}

func (o InstallationTemplatePartitionSchemePartitionArrayOutput) Index(i pulumi.IntInput) InstallationTemplatePartitionSchemePartitionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstallationTemplatePartitionSchemePartition {
		return vs[0].([]*InstallationTemplatePartitionSchemePartition)[vs[1].(int)]
	}).(InstallationTemplatePartitionSchemePartitionOutput)
}

type InstallationTemplatePartitionSchemePartitionMapOutput struct{ *pulumi.OutputState }

func (InstallationTemplatePartitionSchemePartitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstallationTemplatePartitionSchemePartition)(nil)).Elem()
}

func (o InstallationTemplatePartitionSchemePartitionMapOutput) ToInstallationTemplatePartitionSchemePartitionMapOutput() InstallationTemplatePartitionSchemePartitionMapOutput {
	return o
}

func (o InstallationTemplatePartitionSchemePartitionMapOutput) ToInstallationTemplatePartitionSchemePartitionMapOutputWithContext(ctx context.Context) InstallationTemplatePartitionSchemePartitionMapOutput {
	return o
}

func (o InstallationTemplatePartitionSchemePartitionMapOutput) MapIndex(k pulumi.StringInput) InstallationTemplatePartitionSchemePartitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstallationTemplatePartitionSchemePartition {
		return vs[0].(map[string]*InstallationTemplatePartitionSchemePartition)[vs[1].(string)]
	}).(InstallationTemplatePartitionSchemePartitionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstallationTemplatePartitionSchemePartitionInput)(nil)).Elem(), &InstallationTemplatePartitionSchemePartition{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstallationTemplatePartitionSchemePartitionArrayInput)(nil)).Elem(), InstallationTemplatePartitionSchemePartitionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstallationTemplatePartitionSchemePartitionMapInput)(nil)).Elem(), InstallationTemplatePartitionSchemePartitionMap{})
	pulumi.RegisterOutputType(InstallationTemplatePartitionSchemePartitionOutput{})
	pulumi.RegisterOutputType(InstallationTemplatePartitionSchemePartitionArrayOutput{})
	pulumi.RegisterOutputType(InstallationTemplatePartitionSchemePartitionMapOutput{})
}
