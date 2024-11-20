// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a backend server group (farm) to be used by loadbalancing frontend(s)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/IpLoadBalancing"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			lb, err := IpLoadBalancing.GetIpLoadBalancing(ctx, &iploadbalancing.GetIpLoadBalancingArgs{
//				ServiceName: pulumi.StringRef("ip-1.2.3.4"),
//				State:       pulumi.StringRef("ok"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = IpLoadBalancing.NewUdpFarm(ctx, "farmName", &IpLoadBalancing.UdpFarmArgs{
//				DisplayName: pulumi.String("ingress-8080-gra"),
//				Port:        pulumi.Float64(80),
//				ServiceName: pulumi.String(lb.ServiceName),
//				Zone:        pulumi.String("gra"),
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
// UDP Farm can be imported using the following format `service_name` and the `id` of the farm, separated by "/" e.g.
//
// bash
//
// ```sh
// $ pulumi import ovh:IpLoadBalancing/udpFarm:UdpFarm farmname service_name/farm_id
// ```
type UdpFarm struct {
	pulumi.CustomResourceState

	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Id of your farm.
	FarmId pulumi.Float64Output `pulumi:"farmId"`
	// Port attached to your farm ([1..49151]). Inherited from frontend if null
	Port pulumi.Float64Output `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.Float64PtrOutput `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewUdpFarm registers a new resource with the given unique name, arguments, and options.
func NewUdpFarm(ctx *pulumi.Context,
	name string, args *UdpFarmArgs, opts ...pulumi.ResourceOption) (*UdpFarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UdpFarm
	err := ctx.RegisterResource("ovh:IpLoadBalancing/udpFarm:UdpFarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUdpFarm gets an existing UdpFarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUdpFarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UdpFarmState, opts ...pulumi.ResourceOption) (*UdpFarm, error) {
	var resource UdpFarm
	err := ctx.ReadResource("ovh:IpLoadBalancing/udpFarm:UdpFarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UdpFarm resources.
type udpFarmState struct {
	// Readable label for loadbalancer farm
	DisplayName *string `pulumi:"displayName"`
	// Id of your farm.
	FarmId *float64 `pulumi:"farmId"`
	// Port attached to your farm ([1..49151]). Inherited from frontend if null
	Port *float64 `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId *float64 `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone *string `pulumi:"zone"`
}

type UdpFarmState struct {
	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrInput
	// Id of your farm.
	FarmId pulumi.Float64PtrInput
	// Port attached to your farm ([1..49151]). Inherited from frontend if null
	Port pulumi.Float64PtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.Float64PtrInput
	// Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringPtrInput
}

func (UdpFarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*udpFarmState)(nil)).Elem()
}

type udpFarmArgs struct {
	// Readable label for loadbalancer farm
	DisplayName *string `pulumi:"displayName"`
	// Port attached to your farm ([1..49151]). Inherited from frontend if null
	Port float64 `pulumi:"port"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId *float64 `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a UdpFarm resource.
type UdpFarmArgs struct {
	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrInput
	// Port attached to your farm ([1..49151]). Inherited from frontend if null
	Port pulumi.Float64Input
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.Float64PtrInput
	// Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
	Zone pulumi.StringInput
}

func (UdpFarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*udpFarmArgs)(nil)).Elem()
}

type UdpFarmInput interface {
	pulumi.Input

	ToUdpFarmOutput() UdpFarmOutput
	ToUdpFarmOutputWithContext(ctx context.Context) UdpFarmOutput
}

func (*UdpFarm) ElementType() reflect.Type {
	return reflect.TypeOf((**UdpFarm)(nil)).Elem()
}

func (i *UdpFarm) ToUdpFarmOutput() UdpFarmOutput {
	return i.ToUdpFarmOutputWithContext(context.Background())
}

func (i *UdpFarm) ToUdpFarmOutputWithContext(ctx context.Context) UdpFarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UdpFarmOutput)
}

// UdpFarmArrayInput is an input type that accepts UdpFarmArray and UdpFarmArrayOutput values.
// You can construct a concrete instance of `UdpFarmArrayInput` via:
//
//	UdpFarmArray{ UdpFarmArgs{...} }
type UdpFarmArrayInput interface {
	pulumi.Input

	ToUdpFarmArrayOutput() UdpFarmArrayOutput
	ToUdpFarmArrayOutputWithContext(context.Context) UdpFarmArrayOutput
}

type UdpFarmArray []UdpFarmInput

func (UdpFarmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UdpFarm)(nil)).Elem()
}

func (i UdpFarmArray) ToUdpFarmArrayOutput() UdpFarmArrayOutput {
	return i.ToUdpFarmArrayOutputWithContext(context.Background())
}

func (i UdpFarmArray) ToUdpFarmArrayOutputWithContext(ctx context.Context) UdpFarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UdpFarmArrayOutput)
}

// UdpFarmMapInput is an input type that accepts UdpFarmMap and UdpFarmMapOutput values.
// You can construct a concrete instance of `UdpFarmMapInput` via:
//
//	UdpFarmMap{ "key": UdpFarmArgs{...} }
type UdpFarmMapInput interface {
	pulumi.Input

	ToUdpFarmMapOutput() UdpFarmMapOutput
	ToUdpFarmMapOutputWithContext(context.Context) UdpFarmMapOutput
}

type UdpFarmMap map[string]UdpFarmInput

func (UdpFarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UdpFarm)(nil)).Elem()
}

func (i UdpFarmMap) ToUdpFarmMapOutput() UdpFarmMapOutput {
	return i.ToUdpFarmMapOutputWithContext(context.Background())
}

func (i UdpFarmMap) ToUdpFarmMapOutputWithContext(ctx context.Context) UdpFarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UdpFarmMapOutput)
}

type UdpFarmOutput struct{ *pulumi.OutputState }

func (UdpFarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UdpFarm)(nil)).Elem()
}

func (o UdpFarmOutput) ToUdpFarmOutput() UdpFarmOutput {
	return o
}

func (o UdpFarmOutput) ToUdpFarmOutputWithContext(ctx context.Context) UdpFarmOutput {
	return o
}

// Readable label for loadbalancer farm
func (o UdpFarmOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UdpFarm) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Id of your farm.
func (o UdpFarmOutput) FarmId() pulumi.Float64Output {
	return o.ApplyT(func(v *UdpFarm) pulumi.Float64Output { return v.FarmId }).(pulumi.Float64Output)
}

// Port attached to your farm ([1..49151]). Inherited from frontend if null
func (o UdpFarmOutput) Port() pulumi.Float64Output {
	return o.ApplyT(func(v *UdpFarm) pulumi.Float64Output { return v.Port }).(pulumi.Float64Output)
}

// The internal name of your IP load balancing
func (o UdpFarmOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *UdpFarm) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
func (o UdpFarmOutput) VrackNetworkId() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *UdpFarm) pulumi.Float64PtrOutput { return v.VrackNetworkId }).(pulumi.Float64PtrOutput)
}

// Zone where the farm will be defined (ie. `gra`, `bhs` also supports `all`)
func (o UdpFarmOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *UdpFarm) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type UdpFarmArrayOutput struct{ *pulumi.OutputState }

func (UdpFarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UdpFarm)(nil)).Elem()
}

func (o UdpFarmArrayOutput) ToUdpFarmArrayOutput() UdpFarmArrayOutput {
	return o
}

func (o UdpFarmArrayOutput) ToUdpFarmArrayOutputWithContext(ctx context.Context) UdpFarmArrayOutput {
	return o
}

func (o UdpFarmArrayOutput) Index(i pulumi.IntInput) UdpFarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UdpFarm {
		return vs[0].([]*UdpFarm)[vs[1].(int)]
	}).(UdpFarmOutput)
}

type UdpFarmMapOutput struct{ *pulumi.OutputState }

func (UdpFarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UdpFarm)(nil)).Elem()
}

func (o UdpFarmMapOutput) ToUdpFarmMapOutput() UdpFarmMapOutput {
	return o
}

func (o UdpFarmMapOutput) ToUdpFarmMapOutputWithContext(ctx context.Context) UdpFarmMapOutput {
	return o
}

func (o UdpFarmMapOutput) MapIndex(k pulumi.StringInput) UdpFarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UdpFarm {
		return vs[0].(map[string]*UdpFarm)[vs[1].(string)]
	}).(UdpFarmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UdpFarmInput)(nil)).Elem(), &UdpFarm{})
	pulumi.RegisterInputType(reflect.TypeOf((*UdpFarmArrayInput)(nil)).Elem(), UdpFarmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UdpFarmMapInput)(nil)).Elem(), UdpFarmMap{})
	pulumi.RegisterOutputType(UdpFarmOutput{})
	pulumi.RegisterOutputType(UdpFarmArrayOutput{})
	pulumi.RegisterOutputType(UdpFarmMapOutput{})
}
