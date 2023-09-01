// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vrack

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attach a Public Cloud Project to a VRack.
//
// ## Example Usage
//
// ## Import
//
// Attachment of a public cloud project and a VRack can be imported using the `project_id`, the `service_name` (vRackID) and the `attach_id`, separated by "/" E.g., bash <break><break>```sh<break> $ pulumi import ovh:Vrack/cloudProject:CloudProject myattach ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678/vrack_pn-12345678-cloudproject_ookie9mee8Shaeghaeleeju7Xeghohv6e-attach <break>```<break><break>
type CloudProject struct {
	pulumi.CustomResourceState

	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The service name of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewCloudProject registers a new resource with the given unique name, arguments, and options.
func NewCloudProject(ctx *pulumi.Context,
	name string, args *CloudProjectArgs, opts ...pulumi.ResourceOption) (*CloudProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CloudProject
	err := ctx.RegisterResource("ovh:Vrack/cloudProject:CloudProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProject gets an existing CloudProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectState, opts ...pulumi.ResourceOption) (*CloudProject, error) {
	var resource CloudProject
	err := ctx.ReadResource("ovh:Vrack/cloudProject:CloudProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProject resources.
type cloudProjectState struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ProjectId *string `pulumi:"projectId"`
	// The service name of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

type CloudProjectState struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ProjectId pulumi.StringPtrInput
	// The service name of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (CloudProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectState)(nil)).Elem()
}

type cloudProjectArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ProjectId string `pulumi:"projectId"`
	// The service name of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProject resource.
type CloudProjectArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ProjectId pulumi.StringInput
	// The service name of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (CloudProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectArgs)(nil)).Elem()
}

type CloudProjectInput interface {
	pulumi.Input

	ToCloudProjectOutput() CloudProjectOutput
	ToCloudProjectOutputWithContext(ctx context.Context) CloudProjectOutput
}

func (*CloudProject) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProject)(nil)).Elem()
}

func (i *CloudProject) ToCloudProjectOutput() CloudProjectOutput {
	return i.ToCloudProjectOutputWithContext(context.Background())
}

func (i *CloudProject) ToCloudProjectOutputWithContext(ctx context.Context) CloudProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectOutput)
}

// CloudProjectArrayInput is an input type that accepts CloudProjectArray and CloudProjectArrayOutput values.
// You can construct a concrete instance of `CloudProjectArrayInput` via:
//
//	CloudProjectArray{ CloudProjectArgs{...} }
type CloudProjectArrayInput interface {
	pulumi.Input

	ToCloudProjectArrayOutput() CloudProjectArrayOutput
	ToCloudProjectArrayOutputWithContext(context.Context) CloudProjectArrayOutput
}

type CloudProjectArray []CloudProjectInput

func (CloudProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProject)(nil)).Elem()
}

func (i CloudProjectArray) ToCloudProjectArrayOutput() CloudProjectArrayOutput {
	return i.ToCloudProjectArrayOutputWithContext(context.Background())
}

func (i CloudProjectArray) ToCloudProjectArrayOutputWithContext(ctx context.Context) CloudProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectArrayOutput)
}

// CloudProjectMapInput is an input type that accepts CloudProjectMap and CloudProjectMapOutput values.
// You can construct a concrete instance of `CloudProjectMapInput` via:
//
//	CloudProjectMap{ "key": CloudProjectArgs{...} }
type CloudProjectMapInput interface {
	pulumi.Input

	ToCloudProjectMapOutput() CloudProjectMapOutput
	ToCloudProjectMapOutputWithContext(context.Context) CloudProjectMapOutput
}

type CloudProjectMap map[string]CloudProjectInput

func (CloudProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProject)(nil)).Elem()
}

func (i CloudProjectMap) ToCloudProjectMapOutput() CloudProjectMapOutput {
	return i.ToCloudProjectMapOutputWithContext(context.Background())
}

func (i CloudProjectMap) ToCloudProjectMapOutputWithContext(ctx context.Context) CloudProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectMapOutput)
}

type CloudProjectOutput struct{ *pulumi.OutputState }

func (CloudProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProject)(nil)).Elem()
}

func (o CloudProjectOutput) ToCloudProjectOutput() CloudProjectOutput {
	return o
}

func (o CloudProjectOutput) ToCloudProjectOutputWithContext(ctx context.Context) CloudProjectOutput {
	return o
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o CloudProjectOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProject) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The service name of the vrack. If omitted,
// the `OVH_VRACK_SERVICE` environment variable is used.
func (o CloudProjectOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProject) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type CloudProjectArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProject)(nil)).Elem()
}

func (o CloudProjectArrayOutput) ToCloudProjectArrayOutput() CloudProjectArrayOutput {
	return o
}

func (o CloudProjectArrayOutput) ToCloudProjectArrayOutputWithContext(ctx context.Context) CloudProjectArrayOutput {
	return o
}

func (o CloudProjectArrayOutput) Index(i pulumi.IntInput) CloudProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudProject {
		return vs[0].([]*CloudProject)[vs[1].(int)]
	}).(CloudProjectOutput)
}

type CloudProjectMapOutput struct{ *pulumi.OutputState }

func (CloudProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProject)(nil)).Elem()
}

func (o CloudProjectMapOutput) ToCloudProjectMapOutput() CloudProjectMapOutput {
	return o
}

func (o CloudProjectMapOutput) ToCloudProjectMapOutputWithContext(ctx context.Context) CloudProjectMapOutput {
	return o
}

func (o CloudProjectMapOutput) MapIndex(k pulumi.StringInput) CloudProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudProject {
		return vs[0].(map[string]*CloudProject)[vs[1].(string)]
	}).(CloudProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectInput)(nil)).Elem(), &CloudProject{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectArrayInput)(nil)).Elem(), CloudProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectMapInput)(nil)).Elem(), CloudProjectMap{})
	pulumi.RegisterOutputType(CloudProjectOutput{})
	pulumi.RegisterOutputType(CloudProjectArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectMapOutput{})
}
