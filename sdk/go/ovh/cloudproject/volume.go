// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Volume struct {
	pulumi.CustomResourceState

	// The action of the operation
	Action pulumi.StringOutput `pulumi:"action"`
	// The completed date of the operation
	CompletedAt pulumi.StringOutput `pulumi:"completedAt"`
	// The creation date of the operation
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Volume description
	Description pulumi.StringOutput `pulumi:"description"`
	// Image ID
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// Instance ID
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Volume name
	Name pulumi.StringOutput `pulumi:"name"`
	// Volume status
	Progress pulumi.Float64Output `pulumi:"progress"`
	// Region name
	RegionName pulumi.StringOutput `pulumi:"regionName"`
	// List of regions
	Regions pulumi.StringArrayOutput `pulumi:"regions"`
	// Id of the resource
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// Service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Volume size
	Size pulumi.Float64Output `pulumi:"size"`
	// Snapshot ID
	SnapshotId pulumi.StringOutput `pulumi:"snapshotId"`
	// Datetime of the operation creation
	StartedAt pulumi.StringOutput `pulumi:"startedAt"`
	// Volume status
	Status pulumi.StringOutput `pulumi:"status"`
	// Sub-operations of the operation
	SubOperations VolumeSubOperationArrayOutput `pulumi:"subOperations"`
	// Type of the volume
	Type pulumi.StringOutput `pulumi:"type"`
	// Volume ID
	VolumeId pulumi.StringOutput `pulumi:"volumeId"`
}

// NewVolume registers a new resource with the given unique name, arguments, and options.
func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegionName == nil {
		return nil, errors.New("invalid value for required argument 'RegionName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Volume
	err := ctx.RegisterResource("ovh:CloudProject/volume:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVolume gets an existing Volume resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("ovh:CloudProject/volume:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Volume resources.
type volumeState struct {
	// The action of the operation
	Action *string `pulumi:"action"`
	// The completed date of the operation
	CompletedAt *string `pulumi:"completedAt"`
	// The creation date of the operation
	CreatedAt *string `pulumi:"createdAt"`
	// Volume description
	Description *string `pulumi:"description"`
	// Image ID
	ImageId *string `pulumi:"imageId"`
	// Instance ID
	InstanceId *string `pulumi:"instanceId"`
	// Volume name
	Name *string `pulumi:"name"`
	// Volume status
	Progress *float64 `pulumi:"progress"`
	// Region name
	RegionName *string `pulumi:"regionName"`
	// List of regions
	Regions []string `pulumi:"regions"`
	// Id of the resource
	ResourceId *string `pulumi:"resourceId"`
	// Service name
	ServiceName *string `pulumi:"serviceName"`
	// Volume size
	Size *float64 `pulumi:"size"`
	// Snapshot ID
	SnapshotId *string `pulumi:"snapshotId"`
	// Datetime of the operation creation
	StartedAt *string `pulumi:"startedAt"`
	// Volume status
	Status *string `pulumi:"status"`
	// Sub-operations of the operation
	SubOperations []VolumeSubOperation `pulumi:"subOperations"`
	// Type of the volume
	Type *string `pulumi:"type"`
	// Volume ID
	VolumeId *string `pulumi:"volumeId"`
}

type VolumeState struct {
	// The action of the operation
	Action pulumi.StringPtrInput
	// The completed date of the operation
	CompletedAt pulumi.StringPtrInput
	// The creation date of the operation
	CreatedAt pulumi.StringPtrInput
	// Volume description
	Description pulumi.StringPtrInput
	// Image ID
	ImageId pulumi.StringPtrInput
	// Instance ID
	InstanceId pulumi.StringPtrInput
	// Volume name
	Name pulumi.StringPtrInput
	// Volume status
	Progress pulumi.Float64PtrInput
	// Region name
	RegionName pulumi.StringPtrInput
	// List of regions
	Regions pulumi.StringArrayInput
	// Id of the resource
	ResourceId pulumi.StringPtrInput
	// Service name
	ServiceName pulumi.StringPtrInput
	// Volume size
	Size pulumi.Float64PtrInput
	// Snapshot ID
	SnapshotId pulumi.StringPtrInput
	// Datetime of the operation creation
	StartedAt pulumi.StringPtrInput
	// Volume status
	Status pulumi.StringPtrInput
	// Sub-operations of the operation
	SubOperations VolumeSubOperationArrayInput
	// Type of the volume
	Type pulumi.StringPtrInput
	// Volume ID
	VolumeId pulumi.StringPtrInput
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	// Volume description
	Description *string `pulumi:"description"`
	// Image ID
	ImageId *string `pulumi:"imageId"`
	// Instance ID
	InstanceId *string `pulumi:"instanceId"`
	// Volume name
	Name *string `pulumi:"name"`
	// Region name
	RegionName string `pulumi:"regionName"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
	// Volume size
	Size *float64 `pulumi:"size"`
	// Snapshot ID
	SnapshotId *string `pulumi:"snapshotId"`
	// Type of the volume
	Type *string `pulumi:"type"`
	// Volume ID
	VolumeId *string `pulumi:"volumeId"`
}

// The set of arguments for constructing a Volume resource.
type VolumeArgs struct {
	// Volume description
	Description pulumi.StringPtrInput
	// Image ID
	ImageId pulumi.StringPtrInput
	// Instance ID
	InstanceId pulumi.StringPtrInput
	// Volume name
	Name pulumi.StringPtrInput
	// Region name
	RegionName pulumi.StringInput
	// Service name
	ServiceName pulumi.StringInput
	// Volume size
	Size pulumi.Float64PtrInput
	// Snapshot ID
	SnapshotId pulumi.StringPtrInput
	// Type of the volume
	Type pulumi.StringPtrInput
	// Volume ID
	VolumeId pulumi.StringPtrInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

// VolumeArrayInput is an input type that accepts VolumeArray and VolumeArrayOutput values.
// You can construct a concrete instance of `VolumeArrayInput` via:
//
//	VolumeArray{ VolumeArgs{...} }
type VolumeArrayInput interface {
	pulumi.Input

	ToVolumeArrayOutput() VolumeArrayOutput
	ToVolumeArrayOutputWithContext(context.Context) VolumeArrayOutput
}

type VolumeArray []VolumeInput

func (VolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Volume)(nil)).Elem()
}

func (i VolumeArray) ToVolumeArrayOutput() VolumeArrayOutput {
	return i.ToVolumeArrayOutputWithContext(context.Background())
}

func (i VolumeArray) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeArrayOutput)
}

// VolumeMapInput is an input type that accepts VolumeMap and VolumeMapOutput values.
// You can construct a concrete instance of `VolumeMapInput` via:
//
//	VolumeMap{ "key": VolumeArgs{...} }
type VolumeMapInput interface {
	pulumi.Input

	ToVolumeMapOutput() VolumeMapOutput
	ToVolumeMapOutputWithContext(context.Context) VolumeMapOutput
}

type VolumeMap map[string]VolumeInput

func (VolumeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Volume)(nil)).Elem()
}

func (i VolumeMap) ToVolumeMapOutput() VolumeMapOutput {
	return i.ToVolumeMapOutputWithContext(context.Background())
}

func (i VolumeMap) ToVolumeMapOutputWithContext(ctx context.Context) VolumeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeMapOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

// The action of the operation
func (o VolumeOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// The completed date of the operation
func (o VolumeOutput) CompletedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.CompletedAt }).(pulumi.StringOutput)
}

// The creation date of the operation
func (o VolumeOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Volume description
func (o VolumeOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Image ID
func (o VolumeOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// Instance ID
func (o VolumeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Volume name
func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Volume status
func (o VolumeOutput) Progress() pulumi.Float64Output {
	return o.ApplyT(func(v *Volume) pulumi.Float64Output { return v.Progress }).(pulumi.Float64Output)
}

// Region name
func (o VolumeOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.RegionName }).(pulumi.StringOutput)
}

// List of regions
func (o VolumeOutput) Regions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringArrayOutput { return v.Regions }).(pulumi.StringArrayOutput)
}

// Id of the resource
func (o VolumeOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ResourceId }).(pulumi.StringOutput)
}

// Service name
func (o VolumeOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Volume size
func (o VolumeOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v *Volume) pulumi.Float64Output { return v.Size }).(pulumi.Float64Output)
}

// Snapshot ID
func (o VolumeOutput) SnapshotId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.SnapshotId }).(pulumi.StringOutput)
}

// Datetime of the operation creation
func (o VolumeOutput) StartedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.StartedAt }).(pulumi.StringOutput)
}

// Volume status
func (o VolumeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Sub-operations of the operation
func (o VolumeOutput) SubOperations() VolumeSubOperationArrayOutput {
	return o.ApplyT(func(v *Volume) VolumeSubOperationArrayOutput { return v.SubOperations }).(VolumeSubOperationArrayOutput)
}

// Type of the volume
func (o VolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Volume ID
func (o VolumeOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
}

type VolumeArrayOutput struct{ *pulumi.OutputState }

func (VolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Volume)(nil)).Elem()
}

func (o VolumeArrayOutput) ToVolumeArrayOutput() VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) ToVolumeArrayOutputWithContext(ctx context.Context) VolumeArrayOutput {
	return o
}

func (o VolumeArrayOutput) Index(i pulumi.IntInput) VolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Volume {
		return vs[0].([]*Volume)[vs[1].(int)]
	}).(VolumeOutput)
}

type VolumeMapOutput struct{ *pulumi.OutputState }

func (VolumeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Volume)(nil)).Elem()
}

func (o VolumeMapOutput) ToVolumeMapOutput() VolumeMapOutput {
	return o
}

func (o VolumeMapOutput) ToVolumeMapOutputWithContext(ctx context.Context) VolumeMapOutput {
	return o
}

func (o VolumeMapOutput) MapIndex(k pulumi.StringInput) VolumeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Volume {
		return vs[0].(map[string]*Volume)[vs[1].(string)]
	}).(VolumeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeInput)(nil)).Elem(), &Volume{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeArrayInput)(nil)).Elem(), VolumeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumeMapInput)(nil)).Elem(), VolumeMap{})
	pulumi.RegisterOutputType(VolumeOutput{})
	pulumi.RegisterOutputType(VolumeArrayOutput{})
	pulumi.RegisterOutputType(VolumeMapOutput{})
}