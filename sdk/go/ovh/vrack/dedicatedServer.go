// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vrack

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// > **NOTE:** The resource `Vrack.DedicatedServer` is DEPRECATED and will be removed in a future version.
// Use the resource `Vrack.DedicatedServerInterface` instead.
//
// Attach a dedicated server to a VRack.
//
// ## Example Usage
type DedicatedServer struct {
	pulumi.CustomResourceState

	// The id of the dedicated server.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// The service name of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewDedicatedServer registers a new resource with the given unique name, arguments, and options.
func NewDedicatedServer(ctx *pulumi.Context,
	name string, args *DedicatedServerArgs, opts ...pulumi.ResourceOption) (*DedicatedServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DedicatedServer
	err := ctx.RegisterResource("ovh:Vrack/dedicatedServer:DedicatedServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDedicatedServer gets an existing DedicatedServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDedicatedServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DedicatedServerState, opts ...pulumi.ResourceOption) (*DedicatedServer, error) {
	var resource DedicatedServer
	err := ctx.ReadResource("ovh:Vrack/dedicatedServer:DedicatedServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DedicatedServer resources.
type dedicatedServerState struct {
	// The id of the dedicated server.
	ServerId *string `pulumi:"serverId"`
	// The service name of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

type DedicatedServerState struct {
	// The id of the dedicated server.
	ServerId pulumi.StringPtrInput
	// The service name of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (DedicatedServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedServerState)(nil)).Elem()
}

type dedicatedServerArgs struct {
	// The id of the dedicated server.
	ServerId string `pulumi:"serverId"`
	// The service name of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a DedicatedServer resource.
type DedicatedServerArgs struct {
	// The id of the dedicated server.
	ServerId pulumi.StringInput
	// The service name of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (DedicatedServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dedicatedServerArgs)(nil)).Elem()
}

type DedicatedServerInput interface {
	pulumi.Input

	ToDedicatedServerOutput() DedicatedServerOutput
	ToDedicatedServerOutputWithContext(ctx context.Context) DedicatedServerOutput
}

func (*DedicatedServer) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedServer)(nil)).Elem()
}

func (i *DedicatedServer) ToDedicatedServerOutput() DedicatedServerOutput {
	return i.ToDedicatedServerOutputWithContext(context.Background())
}

func (i *DedicatedServer) ToDedicatedServerOutputWithContext(ctx context.Context) DedicatedServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedServerOutput)
}

func (i *DedicatedServer) ToOutput(ctx context.Context) pulumix.Output[*DedicatedServer] {
	return pulumix.Output[*DedicatedServer]{
		OutputState: i.ToDedicatedServerOutputWithContext(ctx).OutputState,
	}
}

// DedicatedServerArrayInput is an input type that accepts DedicatedServerArray and DedicatedServerArrayOutput values.
// You can construct a concrete instance of `DedicatedServerArrayInput` via:
//
//	DedicatedServerArray{ DedicatedServerArgs{...} }
type DedicatedServerArrayInput interface {
	pulumi.Input

	ToDedicatedServerArrayOutput() DedicatedServerArrayOutput
	ToDedicatedServerArrayOutputWithContext(context.Context) DedicatedServerArrayOutput
}

type DedicatedServerArray []DedicatedServerInput

func (DedicatedServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedServer)(nil)).Elem()
}

func (i DedicatedServerArray) ToDedicatedServerArrayOutput() DedicatedServerArrayOutput {
	return i.ToDedicatedServerArrayOutputWithContext(context.Background())
}

func (i DedicatedServerArray) ToDedicatedServerArrayOutputWithContext(ctx context.Context) DedicatedServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedServerArrayOutput)
}

func (i DedicatedServerArray) ToOutput(ctx context.Context) pulumix.Output[[]*DedicatedServer] {
	return pulumix.Output[[]*DedicatedServer]{
		OutputState: i.ToDedicatedServerArrayOutputWithContext(ctx).OutputState,
	}
}

// DedicatedServerMapInput is an input type that accepts DedicatedServerMap and DedicatedServerMapOutput values.
// You can construct a concrete instance of `DedicatedServerMapInput` via:
//
//	DedicatedServerMap{ "key": DedicatedServerArgs{...} }
type DedicatedServerMapInput interface {
	pulumi.Input

	ToDedicatedServerMapOutput() DedicatedServerMapOutput
	ToDedicatedServerMapOutputWithContext(context.Context) DedicatedServerMapOutput
}

type DedicatedServerMap map[string]DedicatedServerInput

func (DedicatedServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedServer)(nil)).Elem()
}

func (i DedicatedServerMap) ToDedicatedServerMapOutput() DedicatedServerMapOutput {
	return i.ToDedicatedServerMapOutputWithContext(context.Background())
}

func (i DedicatedServerMap) ToDedicatedServerMapOutputWithContext(ctx context.Context) DedicatedServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DedicatedServerMapOutput)
}

func (i DedicatedServerMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*DedicatedServer] {
	return pulumix.Output[map[string]*DedicatedServer]{
		OutputState: i.ToDedicatedServerMapOutputWithContext(ctx).OutputState,
	}
}

type DedicatedServerOutput struct{ *pulumi.OutputState }

func (DedicatedServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DedicatedServer)(nil)).Elem()
}

func (o DedicatedServerOutput) ToDedicatedServerOutput() DedicatedServerOutput {
	return o
}

func (o DedicatedServerOutput) ToDedicatedServerOutputWithContext(ctx context.Context) DedicatedServerOutput {
	return o
}

func (o DedicatedServerOutput) ToOutput(ctx context.Context) pulumix.Output[*DedicatedServer] {
	return pulumix.Output[*DedicatedServer]{
		OutputState: o.OutputState,
	}
}

// The id of the dedicated server.
func (o DedicatedServerOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServer) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

// The service name of the vrack. If omitted,
// the `OVH_VRACK_SERVICE` environment variable is used.
func (o DedicatedServerOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *DedicatedServer) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type DedicatedServerArrayOutput struct{ *pulumi.OutputState }

func (DedicatedServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DedicatedServer)(nil)).Elem()
}

func (o DedicatedServerArrayOutput) ToDedicatedServerArrayOutput() DedicatedServerArrayOutput {
	return o
}

func (o DedicatedServerArrayOutput) ToDedicatedServerArrayOutputWithContext(ctx context.Context) DedicatedServerArrayOutput {
	return o
}

func (o DedicatedServerArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*DedicatedServer] {
	return pulumix.Output[[]*DedicatedServer]{
		OutputState: o.OutputState,
	}
}

func (o DedicatedServerArrayOutput) Index(i pulumi.IntInput) DedicatedServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DedicatedServer {
		return vs[0].([]*DedicatedServer)[vs[1].(int)]
	}).(DedicatedServerOutput)
}

type DedicatedServerMapOutput struct{ *pulumi.OutputState }

func (DedicatedServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DedicatedServer)(nil)).Elem()
}

func (o DedicatedServerMapOutput) ToDedicatedServerMapOutput() DedicatedServerMapOutput {
	return o
}

func (o DedicatedServerMapOutput) ToDedicatedServerMapOutputWithContext(ctx context.Context) DedicatedServerMapOutput {
	return o
}

func (o DedicatedServerMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*DedicatedServer] {
	return pulumix.Output[map[string]*DedicatedServer]{
		OutputState: o.OutputState,
	}
}

func (o DedicatedServerMapOutput) MapIndex(k pulumi.StringInput) DedicatedServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DedicatedServer {
		return vs[0].(map[string]*DedicatedServer)[vs[1].(string)]
	}).(DedicatedServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedServerInput)(nil)).Elem(), &DedicatedServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedServerArrayInput)(nil)).Elem(), DedicatedServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DedicatedServerMapInput)(nil)).Elem(), DedicatedServerMap{})
	pulumi.RegisterOutputType(DedicatedServerOutput{})
	pulumi.RegisterOutputType(DedicatedServerArrayOutput{})
	pulumi.RegisterOutputType(DedicatedServerMapOutput{})
}
