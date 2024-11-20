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

// Generates a temporary presigned S3 URLs to download or upload an object.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			presignedUrlRegionStoragePresign, err := CloudProject.NewRegionStoragePresign(ctx, "presignedUrlRegionStoragePresign", &CloudProject.RegionStoragePresignArgs{
//				ServiceName: pulumi.String("xxxxxxxxxxxxxxxxx"),
//				RegionName:  pulumi.String("GRA"),
//				Expire:      pulumi.Int(3600),
//				Method:      pulumi.String("GET"),
//				Object:      pulumi.String("an-object-in-the-bucket"),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("presignedUrl", presignedUrlRegionStoragePresign.Url)
//			return nil
//		})
//	}
//
// ```
type RegionStoragePresign struct {
	pulumi.CustomResourceState

	// Define, in seconds, for how long your URL will be
	// valid.
	Expire pulumi.IntOutput `pulumi:"expire"`
	// The method you want to use to interact with your
	// object. Can be either 'GET' or 'PUT'.
	Method pulumi.StringOutput `pulumi:"method"`
	// The name of your S3 storage container/bucket.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the object in your S3 bucket.
	Object pulumi.StringOutput `pulumi:"object"`
	// The region in which your storage is located. Must
	// be in **uppercase**. Ex.: "GRA".
	RegionName pulumi.StringOutput `pulumi:"regionName"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Computed URL result.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewRegionStoragePresign registers a new resource with the given unique name, arguments, and options.
func NewRegionStoragePresign(ctx *pulumi.Context,
	name string, args *RegionStoragePresignArgs, opts ...pulumi.ResourceOption) (*RegionStoragePresign, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Expire == nil {
		return nil, errors.New("invalid value for required argument 'Expire'")
	}
	if args.Method == nil {
		return nil, errors.New("invalid value for required argument 'Method'")
	}
	if args.Object == nil {
		return nil, errors.New("invalid value for required argument 'Object'")
	}
	if args.RegionName == nil {
		return nil, errors.New("invalid value for required argument 'RegionName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RegionStoragePresign
	err := ctx.RegisterResource("ovh:CloudProject/regionStoragePresign:RegionStoragePresign", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionStoragePresign gets an existing RegionStoragePresign resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionStoragePresign(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionStoragePresignState, opts ...pulumi.ResourceOption) (*RegionStoragePresign, error) {
	var resource RegionStoragePresign
	err := ctx.ReadResource("ovh:CloudProject/regionStoragePresign:RegionStoragePresign", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionStoragePresign resources.
type regionStoragePresignState struct {
	// Define, in seconds, for how long your URL will be
	// valid.
	Expire *int `pulumi:"expire"`
	// The method you want to use to interact with your
	// object. Can be either 'GET' or 'PUT'.
	Method *string `pulumi:"method"`
	// The name of your S3 storage container/bucket.
	Name *string `pulumi:"name"`
	// The name of the object in your S3 bucket.
	Object *string `pulumi:"object"`
	// The region in which your storage is located. Must
	// be in **uppercase**. Ex.: "GRA".
	RegionName *string `pulumi:"regionName"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Computed URL result.
	Url *string `pulumi:"url"`
}

type RegionStoragePresignState struct {
	// Define, in seconds, for how long your URL will be
	// valid.
	Expire pulumi.IntPtrInput
	// The method you want to use to interact with your
	// object. Can be either 'GET' or 'PUT'.
	Method pulumi.StringPtrInput
	// The name of your S3 storage container/bucket.
	Name pulumi.StringPtrInput
	// The name of the object in your S3 bucket.
	Object pulumi.StringPtrInput
	// The region in which your storage is located. Must
	// be in **uppercase**. Ex.: "GRA".
	RegionName pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Computed URL result.
	Url pulumi.StringPtrInput
}

func (RegionStoragePresignState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionStoragePresignState)(nil)).Elem()
}

type regionStoragePresignArgs struct {
	// Define, in seconds, for how long your URL will be
	// valid.
	Expire int `pulumi:"expire"`
	// The method you want to use to interact with your
	// object. Can be either 'GET' or 'PUT'.
	Method string `pulumi:"method"`
	// The name of your S3 storage container/bucket.
	Name *string `pulumi:"name"`
	// The name of the object in your S3 bucket.
	Object string `pulumi:"object"`
	// The region in which your storage is located. Must
	// be in **uppercase**. Ex.: "GRA".
	RegionName string `pulumi:"regionName"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a RegionStoragePresign resource.
type RegionStoragePresignArgs struct {
	// Define, in seconds, for how long your URL will be
	// valid.
	Expire pulumi.IntInput
	// The method you want to use to interact with your
	// object. Can be either 'GET' or 'PUT'.
	Method pulumi.StringInput
	// The name of your S3 storage container/bucket.
	Name pulumi.StringPtrInput
	// The name of the object in your S3 bucket.
	Object pulumi.StringInput
	// The region in which your storage is located. Must
	// be in **uppercase**. Ex.: "GRA".
	RegionName pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (RegionStoragePresignArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionStoragePresignArgs)(nil)).Elem()
}

type RegionStoragePresignInput interface {
	pulumi.Input

	ToRegionStoragePresignOutput() RegionStoragePresignOutput
	ToRegionStoragePresignOutputWithContext(ctx context.Context) RegionStoragePresignOutput
}

func (*RegionStoragePresign) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionStoragePresign)(nil)).Elem()
}

func (i *RegionStoragePresign) ToRegionStoragePresignOutput() RegionStoragePresignOutput {
	return i.ToRegionStoragePresignOutputWithContext(context.Background())
}

func (i *RegionStoragePresign) ToRegionStoragePresignOutputWithContext(ctx context.Context) RegionStoragePresignOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionStoragePresignOutput)
}

// RegionStoragePresignArrayInput is an input type that accepts RegionStoragePresignArray and RegionStoragePresignArrayOutput values.
// You can construct a concrete instance of `RegionStoragePresignArrayInput` via:
//
//	RegionStoragePresignArray{ RegionStoragePresignArgs{...} }
type RegionStoragePresignArrayInput interface {
	pulumi.Input

	ToRegionStoragePresignArrayOutput() RegionStoragePresignArrayOutput
	ToRegionStoragePresignArrayOutputWithContext(context.Context) RegionStoragePresignArrayOutput
}

type RegionStoragePresignArray []RegionStoragePresignInput

func (RegionStoragePresignArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionStoragePresign)(nil)).Elem()
}

func (i RegionStoragePresignArray) ToRegionStoragePresignArrayOutput() RegionStoragePresignArrayOutput {
	return i.ToRegionStoragePresignArrayOutputWithContext(context.Background())
}

func (i RegionStoragePresignArray) ToRegionStoragePresignArrayOutputWithContext(ctx context.Context) RegionStoragePresignArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionStoragePresignArrayOutput)
}

// RegionStoragePresignMapInput is an input type that accepts RegionStoragePresignMap and RegionStoragePresignMapOutput values.
// You can construct a concrete instance of `RegionStoragePresignMapInput` via:
//
//	RegionStoragePresignMap{ "key": RegionStoragePresignArgs{...} }
type RegionStoragePresignMapInput interface {
	pulumi.Input

	ToRegionStoragePresignMapOutput() RegionStoragePresignMapOutput
	ToRegionStoragePresignMapOutputWithContext(context.Context) RegionStoragePresignMapOutput
}

type RegionStoragePresignMap map[string]RegionStoragePresignInput

func (RegionStoragePresignMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionStoragePresign)(nil)).Elem()
}

func (i RegionStoragePresignMap) ToRegionStoragePresignMapOutput() RegionStoragePresignMapOutput {
	return i.ToRegionStoragePresignMapOutputWithContext(context.Background())
}

func (i RegionStoragePresignMap) ToRegionStoragePresignMapOutputWithContext(ctx context.Context) RegionStoragePresignMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionStoragePresignMapOutput)
}

type RegionStoragePresignOutput struct{ *pulumi.OutputState }

func (RegionStoragePresignOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionStoragePresign)(nil)).Elem()
}

func (o RegionStoragePresignOutput) ToRegionStoragePresignOutput() RegionStoragePresignOutput {
	return o
}

func (o RegionStoragePresignOutput) ToRegionStoragePresignOutputWithContext(ctx context.Context) RegionStoragePresignOutput {
	return o
}

// Define, in seconds, for how long your URL will be
// valid.
func (o RegionStoragePresignOutput) Expire() pulumi.IntOutput {
	return o.ApplyT(func(v *RegionStoragePresign) pulumi.IntOutput { return v.Expire }).(pulumi.IntOutput)
}

// The method you want to use to interact with your
// object. Can be either 'GET' or 'PUT'.
func (o RegionStoragePresignOutput) Method() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionStoragePresign) pulumi.StringOutput { return v.Method }).(pulumi.StringOutput)
}

// The name of your S3 storage container/bucket.
func (o RegionStoragePresignOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionStoragePresign) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the object in your S3 bucket.
func (o RegionStoragePresignOutput) Object() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionStoragePresign) pulumi.StringOutput { return v.Object }).(pulumi.StringOutput)
}

// The region in which your storage is located. Must
// be in **uppercase**. Ex.: "GRA".
func (o RegionStoragePresignOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionStoragePresign) pulumi.StringOutput { return v.RegionName }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o RegionStoragePresignOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionStoragePresign) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Computed URL result.
func (o RegionStoragePresignOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionStoragePresign) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type RegionStoragePresignArrayOutput struct{ *pulumi.OutputState }

func (RegionStoragePresignArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionStoragePresign)(nil)).Elem()
}

func (o RegionStoragePresignArrayOutput) ToRegionStoragePresignArrayOutput() RegionStoragePresignArrayOutput {
	return o
}

func (o RegionStoragePresignArrayOutput) ToRegionStoragePresignArrayOutputWithContext(ctx context.Context) RegionStoragePresignArrayOutput {
	return o
}

func (o RegionStoragePresignArrayOutput) Index(i pulumi.IntInput) RegionStoragePresignOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegionStoragePresign {
		return vs[0].([]*RegionStoragePresign)[vs[1].(int)]
	}).(RegionStoragePresignOutput)
}

type RegionStoragePresignMapOutput struct{ *pulumi.OutputState }

func (RegionStoragePresignMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionStoragePresign)(nil)).Elem()
}

func (o RegionStoragePresignMapOutput) ToRegionStoragePresignMapOutput() RegionStoragePresignMapOutput {
	return o
}

func (o RegionStoragePresignMapOutput) ToRegionStoragePresignMapOutputWithContext(ctx context.Context) RegionStoragePresignMapOutput {
	return o
}

func (o RegionStoragePresignMapOutput) MapIndex(k pulumi.StringInput) RegionStoragePresignOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegionStoragePresign {
		return vs[0].(map[string]*RegionStoragePresign)[vs[1].(string)]
	}).(RegionStoragePresignOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionStoragePresignInput)(nil)).Elem(), &RegionStoragePresign{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionStoragePresignArrayInput)(nil)).Elem(), RegionStoragePresignArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionStoragePresignMapInput)(nil)).Elem(), RegionStoragePresignMap{})
	pulumi.RegisterOutputType(RegionStoragePresignOutput{})
	pulumi.RegisterOutputType(RegionStoragePresignArrayOutput{})
	pulumi.RegisterOutputType(RegionStoragePresignMapOutput{})
}
