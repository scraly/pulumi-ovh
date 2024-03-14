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

// Apply IP restrictions container registry associated with a public cloud project on Harbor UI and API.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			_, err := CloudProject.GetContainerRegistry(ctx, &cloudproject.GetContainerRegistryArgs{
//				ServiceName: "XXXXXX",
//				RegistryId:  "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = CloudProject.NewContainerRegistryIPRestrictionsManagement(ctx, "my-mgt-iprestrictions", &CloudProject.ContainerRegistryIPRestrictionsManagementArgs{
//				ServiceName: pulumi.Any(ovh_cloud_project_containerregistry.Registry.Service_name),
//				RegistryId:  pulumi.Any(ovh_cloud_project_containerregistry.Registry.Id),
//				IpRestrictions: pulumi.MapArray{
//					pulumi.Map{
//						"ip_block":    pulumi.Any("xxx.xxx.xxx.xxx/xx"),
//						"description": pulumi.Any("xxxxxxx"),
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
// <!--End PulumiCodeChooser -->
type ContainerRegistryIPRestrictionsManagement struct {
	pulumi.CustomResourceState

	// IP restrictions applied on Harbor UI and API.
	IpRestrictions pulumi.MapArrayOutput `pulumi:"ipRestrictions"`
	// The id of the Managed Private Registry.
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewContainerRegistryIPRestrictionsManagement registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistryIPRestrictionsManagement(ctx *pulumi.Context,
	name string, args *ContainerRegistryIPRestrictionsManagementArgs, opts ...pulumi.ResourceOption) (*ContainerRegistryIPRestrictionsManagement, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpRestrictions == nil {
		return nil, errors.New("invalid value for required argument 'IpRestrictions'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerRegistryIPRestrictionsManagement
	err := ctx.RegisterResource("ovh:CloudProject/containerRegistryIPRestrictionsManagement:ContainerRegistryIPRestrictionsManagement", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistryIPRestrictionsManagement gets an existing ContainerRegistryIPRestrictionsManagement resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistryIPRestrictionsManagement(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryIPRestrictionsManagementState, opts ...pulumi.ResourceOption) (*ContainerRegistryIPRestrictionsManagement, error) {
	var resource ContainerRegistryIPRestrictionsManagement
	err := ctx.ReadResource("ovh:CloudProject/containerRegistryIPRestrictionsManagement:ContainerRegistryIPRestrictionsManagement", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistryIPRestrictionsManagement resources.
type containerRegistryIPRestrictionsManagementState struct {
	// IP restrictions applied on Harbor UI and API.
	IpRestrictions []map[string]interface{} `pulumi:"ipRestrictions"`
	// The id of the Managed Private Registry.
	RegistryId *string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

type ContainerRegistryIPRestrictionsManagementState struct {
	// IP restrictions applied on Harbor UI and API.
	IpRestrictions pulumi.MapArrayInput
	// The id of the Managed Private Registry.
	RegistryId pulumi.StringPtrInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (ContainerRegistryIPRestrictionsManagementState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryIPRestrictionsManagementState)(nil)).Elem()
}

type containerRegistryIPRestrictionsManagementArgs struct {
	// IP restrictions applied on Harbor UI and API.
	IpRestrictions []map[string]interface{} `pulumi:"ipRestrictions"`
	// The id of the Managed Private Registry.
	RegistryId string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ContainerRegistryIPRestrictionsManagement resource.
type ContainerRegistryIPRestrictionsManagementArgs struct {
	// IP restrictions applied on Harbor UI and API.
	IpRestrictions pulumi.MapArrayInput
	// The id of the Managed Private Registry.
	RegistryId pulumi.StringInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (ContainerRegistryIPRestrictionsManagementArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryIPRestrictionsManagementArgs)(nil)).Elem()
}

type ContainerRegistryIPRestrictionsManagementInput interface {
	pulumi.Input

	ToContainerRegistryIPRestrictionsManagementOutput() ContainerRegistryIPRestrictionsManagementOutput
	ToContainerRegistryIPRestrictionsManagementOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsManagementOutput
}

func (*ContainerRegistryIPRestrictionsManagement) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryIPRestrictionsManagement)(nil)).Elem()
}

func (i *ContainerRegistryIPRestrictionsManagement) ToContainerRegistryIPRestrictionsManagementOutput() ContainerRegistryIPRestrictionsManagementOutput {
	return i.ToContainerRegistryIPRestrictionsManagementOutputWithContext(context.Background())
}

func (i *ContainerRegistryIPRestrictionsManagement) ToContainerRegistryIPRestrictionsManagementOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsManagementOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIPRestrictionsManagementOutput)
}

// ContainerRegistryIPRestrictionsManagementArrayInput is an input type that accepts ContainerRegistryIPRestrictionsManagementArray and ContainerRegistryIPRestrictionsManagementArrayOutput values.
// You can construct a concrete instance of `ContainerRegistryIPRestrictionsManagementArrayInput` via:
//
//	ContainerRegistryIPRestrictionsManagementArray{ ContainerRegistryIPRestrictionsManagementArgs{...} }
type ContainerRegistryIPRestrictionsManagementArrayInput interface {
	pulumi.Input

	ToContainerRegistryIPRestrictionsManagementArrayOutput() ContainerRegistryIPRestrictionsManagementArrayOutput
	ToContainerRegistryIPRestrictionsManagementArrayOutputWithContext(context.Context) ContainerRegistryIPRestrictionsManagementArrayOutput
}

type ContainerRegistryIPRestrictionsManagementArray []ContainerRegistryIPRestrictionsManagementInput

func (ContainerRegistryIPRestrictionsManagementArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistryIPRestrictionsManagement)(nil)).Elem()
}

func (i ContainerRegistryIPRestrictionsManagementArray) ToContainerRegistryIPRestrictionsManagementArrayOutput() ContainerRegistryIPRestrictionsManagementArrayOutput {
	return i.ToContainerRegistryIPRestrictionsManagementArrayOutputWithContext(context.Background())
}

func (i ContainerRegistryIPRestrictionsManagementArray) ToContainerRegistryIPRestrictionsManagementArrayOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsManagementArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIPRestrictionsManagementArrayOutput)
}

// ContainerRegistryIPRestrictionsManagementMapInput is an input type that accepts ContainerRegistryIPRestrictionsManagementMap and ContainerRegistryIPRestrictionsManagementMapOutput values.
// You can construct a concrete instance of `ContainerRegistryIPRestrictionsManagementMapInput` via:
//
//	ContainerRegistryIPRestrictionsManagementMap{ "key": ContainerRegistryIPRestrictionsManagementArgs{...} }
type ContainerRegistryIPRestrictionsManagementMapInput interface {
	pulumi.Input

	ToContainerRegistryIPRestrictionsManagementMapOutput() ContainerRegistryIPRestrictionsManagementMapOutput
	ToContainerRegistryIPRestrictionsManagementMapOutputWithContext(context.Context) ContainerRegistryIPRestrictionsManagementMapOutput
}

type ContainerRegistryIPRestrictionsManagementMap map[string]ContainerRegistryIPRestrictionsManagementInput

func (ContainerRegistryIPRestrictionsManagementMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistryIPRestrictionsManagement)(nil)).Elem()
}

func (i ContainerRegistryIPRestrictionsManagementMap) ToContainerRegistryIPRestrictionsManagementMapOutput() ContainerRegistryIPRestrictionsManagementMapOutput {
	return i.ToContainerRegistryIPRestrictionsManagementMapOutputWithContext(context.Background())
}

func (i ContainerRegistryIPRestrictionsManagementMap) ToContainerRegistryIPRestrictionsManagementMapOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsManagementMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryIPRestrictionsManagementMapOutput)
}

type ContainerRegistryIPRestrictionsManagementOutput struct{ *pulumi.OutputState }

func (ContainerRegistryIPRestrictionsManagementOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryIPRestrictionsManagement)(nil)).Elem()
}

func (o ContainerRegistryIPRestrictionsManagementOutput) ToContainerRegistryIPRestrictionsManagementOutput() ContainerRegistryIPRestrictionsManagementOutput {
	return o
}

func (o ContainerRegistryIPRestrictionsManagementOutput) ToContainerRegistryIPRestrictionsManagementOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsManagementOutput {
	return o
}

// IP restrictions applied on Harbor UI and API.
func (o ContainerRegistryIPRestrictionsManagementOutput) IpRestrictions() pulumi.MapArrayOutput {
	return o.ApplyT(func(v *ContainerRegistryIPRestrictionsManagement) pulumi.MapArrayOutput { return v.IpRestrictions }).(pulumi.MapArrayOutput)
}

// The id of the Managed Private Registry.
func (o ContainerRegistryIPRestrictionsManagementOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryIPRestrictionsManagement) pulumi.StringOutput { return v.RegistryId }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o ContainerRegistryIPRestrictionsManagementOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryIPRestrictionsManagement) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type ContainerRegistryIPRestrictionsManagementArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryIPRestrictionsManagementArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistryIPRestrictionsManagement)(nil)).Elem()
}

func (o ContainerRegistryIPRestrictionsManagementArrayOutput) ToContainerRegistryIPRestrictionsManagementArrayOutput() ContainerRegistryIPRestrictionsManagementArrayOutput {
	return o
}

func (o ContainerRegistryIPRestrictionsManagementArrayOutput) ToContainerRegistryIPRestrictionsManagementArrayOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsManagementArrayOutput {
	return o
}

func (o ContainerRegistryIPRestrictionsManagementArrayOutput) Index(i pulumi.IntInput) ContainerRegistryIPRestrictionsManagementOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerRegistryIPRestrictionsManagement {
		return vs[0].([]*ContainerRegistryIPRestrictionsManagement)[vs[1].(int)]
	}).(ContainerRegistryIPRestrictionsManagementOutput)
}

type ContainerRegistryIPRestrictionsManagementMapOutput struct{ *pulumi.OutputState }

func (ContainerRegistryIPRestrictionsManagementMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistryIPRestrictionsManagement)(nil)).Elem()
}

func (o ContainerRegistryIPRestrictionsManagementMapOutput) ToContainerRegistryIPRestrictionsManagementMapOutput() ContainerRegistryIPRestrictionsManagementMapOutput {
	return o
}

func (o ContainerRegistryIPRestrictionsManagementMapOutput) ToContainerRegistryIPRestrictionsManagementMapOutputWithContext(ctx context.Context) ContainerRegistryIPRestrictionsManagementMapOutput {
	return o
}

func (o ContainerRegistryIPRestrictionsManagementMapOutput) MapIndex(k pulumi.StringInput) ContainerRegistryIPRestrictionsManagementOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerRegistryIPRestrictionsManagement {
		return vs[0].(map[string]*ContainerRegistryIPRestrictionsManagement)[vs[1].(string)]
	}).(ContainerRegistryIPRestrictionsManagementOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryIPRestrictionsManagementInput)(nil)).Elem(), &ContainerRegistryIPRestrictionsManagement{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryIPRestrictionsManagementArrayInput)(nil)).Elem(), ContainerRegistryIPRestrictionsManagementArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryIPRestrictionsManagementMapInput)(nil)).Elem(), ContainerRegistryIPRestrictionsManagementMap{})
	pulumi.RegisterOutputType(ContainerRegistryIPRestrictionsManagementOutput{})
	pulumi.RegisterOutputType(ContainerRegistryIPRestrictionsManagementArrayOutput{})
	pulumi.RegisterOutputType(ContainerRegistryIPRestrictionsManagementMapOutput{})
}
