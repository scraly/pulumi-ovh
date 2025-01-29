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

// Creates a subnet in a private network of a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/cloudproject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudproject.NewNetworkPrivateSubnet(ctx, "subnet", &cloudproject.NetworkPrivateSubnetArgs{
//				Dhcp:        pulumi.Bool(true),
//				End:         pulumi.String("192.168.168.200"),
//				Network:     pulumi.String("192.168.168.0/24"),
//				NetworkId:   pulumi.String("0234543"),
//				NoGateway:   pulumi.Bool(false),
//				Region:      pulumi.String("GRA1"),
//				ServiceName: pulumi.String("xxxxx"),
//				Start:       pulumi.String("192.168.168.100"),
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
// Subnet in a private network of a public cloud project can be imported using the `service_name` , the `network_id` and the `subnet_id`, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet mysubnet ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678/0f0b73a4-403b-45e4-86d0-b438f1291909
// ```
type NetworkPrivateSubnet struct {
	pulumi.CustomResourceState

	// Ip Block representing the subnet cidr.
	Cidr pulumi.StringOutput `pulumi:"cidr"`
	// Enable DHCP.
	// Changing this forces a new resource to be created. Defaults to false.
	Dhcp pulumi.BoolPtrOutput `pulumi:"dhcp"`
	// Last ip for this region.
	// Changing this value recreates the subnet.
	End pulumi.StringOutput `pulumi:"end"`
	// The IP of the gateway
	GatewayIp pulumi.StringOutput `pulumi:"gatewayIp"`
	// List of ip pools allocated in the subnet.
	// * `ip_pools/network` - Global network with cidr.
	// * `ip_pools/region` - Region where this subnet is created.
	// * `ip_pools/dhcp` - DHCP enabled.
	// * `ip_pools/end` - Last ip for this region.
	// * `ip_pools/start` - First ip for this region.
	IpPools NetworkPrivateSubnetIpPoolArrayOutput `pulumi:"ipPools"`
	// Global network in CIDR format.
	// Changing this value recreates the subnet
	Network pulumi.StringOutput `pulumi:"network"`
	// The id of the network.
	// Changing this forces a new resource to be created.
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// Set to true if you don't want to set a default gateway IP.
	// Changing this value recreates the resource. Defaults to false.
	NoGateway pulumi.BoolPtrOutput `pulumi:"noGateway"`
	// The region in which the network subnet will be created.
	// Ex.: "GRA1". Changing this value recreates the resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// First ip for this region.
	// Changing this value recreates the subnet.
	Start pulumi.StringOutput `pulumi:"start"`
}

// NewNetworkPrivateSubnet registers a new resource with the given unique name, arguments, and options.
func NewNetworkPrivateSubnet(ctx *pulumi.Context,
	name string, args *NetworkPrivateSubnetArgs, opts ...pulumi.ResourceOption) (*NetworkPrivateSubnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.End == nil {
		return nil, errors.New("invalid value for required argument 'End'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Start == nil {
		return nil, errors.New("invalid value for required argument 'Start'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkPrivateSubnet
	err := ctx.RegisterResource("ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPrivateSubnet gets an existing NetworkPrivateSubnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPrivateSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkPrivateSubnetState, opts ...pulumi.ResourceOption) (*NetworkPrivateSubnet, error) {
	var resource NetworkPrivateSubnet
	err := ctx.ReadResource("ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkPrivateSubnet resources.
type networkPrivateSubnetState struct {
	// Ip Block representing the subnet cidr.
	Cidr *string `pulumi:"cidr"`
	// Enable DHCP.
	// Changing this forces a new resource to be created. Defaults to false.
	Dhcp *bool `pulumi:"dhcp"`
	// Last ip for this region.
	// Changing this value recreates the subnet.
	End *string `pulumi:"end"`
	// The IP of the gateway
	GatewayIp *string `pulumi:"gatewayIp"`
	// List of ip pools allocated in the subnet.
	// * `ip_pools/network` - Global network with cidr.
	// * `ip_pools/region` - Region where this subnet is created.
	// * `ip_pools/dhcp` - DHCP enabled.
	// * `ip_pools/end` - Last ip for this region.
	// * `ip_pools/start` - First ip for this region.
	IpPools []NetworkPrivateSubnetIpPool `pulumi:"ipPools"`
	// Global network in CIDR format.
	// Changing this value recreates the subnet
	Network *string `pulumi:"network"`
	// The id of the network.
	// Changing this forces a new resource to be created.
	NetworkId *string `pulumi:"networkId"`
	// Set to true if you don't want to set a default gateway IP.
	// Changing this value recreates the resource. Defaults to false.
	NoGateway *bool `pulumi:"noGateway"`
	// The region in which the network subnet will be created.
	// Ex.: "GRA1". Changing this value recreates the resource.
	Region *string `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// First ip for this region.
	// Changing this value recreates the subnet.
	Start *string `pulumi:"start"`
}

type NetworkPrivateSubnetState struct {
	// Ip Block representing the subnet cidr.
	Cidr pulumi.StringPtrInput
	// Enable DHCP.
	// Changing this forces a new resource to be created. Defaults to false.
	Dhcp pulumi.BoolPtrInput
	// Last ip for this region.
	// Changing this value recreates the subnet.
	End pulumi.StringPtrInput
	// The IP of the gateway
	GatewayIp pulumi.StringPtrInput
	// List of ip pools allocated in the subnet.
	// * `ip_pools/network` - Global network with cidr.
	// * `ip_pools/region` - Region where this subnet is created.
	// * `ip_pools/dhcp` - DHCP enabled.
	// * `ip_pools/end` - Last ip for this region.
	// * `ip_pools/start` - First ip for this region.
	IpPools NetworkPrivateSubnetIpPoolArrayInput
	// Global network in CIDR format.
	// Changing this value recreates the subnet
	Network pulumi.StringPtrInput
	// The id of the network.
	// Changing this forces a new resource to be created.
	NetworkId pulumi.StringPtrInput
	// Set to true if you don't want to set a default gateway IP.
	// Changing this value recreates the resource. Defaults to false.
	NoGateway pulumi.BoolPtrInput
	// The region in which the network subnet will be created.
	// Ex.: "GRA1". Changing this value recreates the resource.
	Region pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// First ip for this region.
	// Changing this value recreates the subnet.
	Start pulumi.StringPtrInput
}

func (NetworkPrivateSubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPrivateSubnetState)(nil)).Elem()
}

type networkPrivateSubnetArgs struct {
	// Enable DHCP.
	// Changing this forces a new resource to be created. Defaults to false.
	Dhcp *bool `pulumi:"dhcp"`
	// Last ip for this region.
	// Changing this value recreates the subnet.
	End string `pulumi:"end"`
	// Global network in CIDR format.
	// Changing this value recreates the subnet
	Network string `pulumi:"network"`
	// The id of the network.
	// Changing this forces a new resource to be created.
	NetworkId string `pulumi:"networkId"`
	// Set to true if you don't want to set a default gateway IP.
	// Changing this value recreates the resource. Defaults to false.
	NoGateway *bool `pulumi:"noGateway"`
	// The region in which the network subnet will be created.
	// Ex.: "GRA1". Changing this value recreates the resource.
	Region string `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// First ip for this region.
	// Changing this value recreates the subnet.
	Start string `pulumi:"start"`
}

// The set of arguments for constructing a NetworkPrivateSubnet resource.
type NetworkPrivateSubnetArgs struct {
	// Enable DHCP.
	// Changing this forces a new resource to be created. Defaults to false.
	Dhcp pulumi.BoolPtrInput
	// Last ip for this region.
	// Changing this value recreates the subnet.
	End pulumi.StringInput
	// Global network in CIDR format.
	// Changing this value recreates the subnet
	Network pulumi.StringInput
	// The id of the network.
	// Changing this forces a new resource to be created.
	NetworkId pulumi.StringInput
	// Set to true if you don't want to set a default gateway IP.
	// Changing this value recreates the resource. Defaults to false.
	NoGateway pulumi.BoolPtrInput
	// The region in which the network subnet will be created.
	// Ex.: "GRA1". Changing this value recreates the resource.
	Region pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
	// First ip for this region.
	// Changing this value recreates the subnet.
	Start pulumi.StringInput
}

func (NetworkPrivateSubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPrivateSubnetArgs)(nil)).Elem()
}

type NetworkPrivateSubnetInput interface {
	pulumi.Input

	ToNetworkPrivateSubnetOutput() NetworkPrivateSubnetOutput
	ToNetworkPrivateSubnetOutputWithContext(ctx context.Context) NetworkPrivateSubnetOutput
}

func (*NetworkPrivateSubnet) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPrivateSubnet)(nil)).Elem()
}

func (i *NetworkPrivateSubnet) ToNetworkPrivateSubnetOutput() NetworkPrivateSubnetOutput {
	return i.ToNetworkPrivateSubnetOutputWithContext(context.Background())
}

func (i *NetworkPrivateSubnet) ToNetworkPrivateSubnetOutputWithContext(ctx context.Context) NetworkPrivateSubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPrivateSubnetOutput)
}

// NetworkPrivateSubnetArrayInput is an input type that accepts NetworkPrivateSubnetArray and NetworkPrivateSubnetArrayOutput values.
// You can construct a concrete instance of `NetworkPrivateSubnetArrayInput` via:
//
//	NetworkPrivateSubnetArray{ NetworkPrivateSubnetArgs{...} }
type NetworkPrivateSubnetArrayInput interface {
	pulumi.Input

	ToNetworkPrivateSubnetArrayOutput() NetworkPrivateSubnetArrayOutput
	ToNetworkPrivateSubnetArrayOutputWithContext(context.Context) NetworkPrivateSubnetArrayOutput
}

type NetworkPrivateSubnetArray []NetworkPrivateSubnetInput

func (NetworkPrivateSubnetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPrivateSubnet)(nil)).Elem()
}

func (i NetworkPrivateSubnetArray) ToNetworkPrivateSubnetArrayOutput() NetworkPrivateSubnetArrayOutput {
	return i.ToNetworkPrivateSubnetArrayOutputWithContext(context.Background())
}

func (i NetworkPrivateSubnetArray) ToNetworkPrivateSubnetArrayOutputWithContext(ctx context.Context) NetworkPrivateSubnetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPrivateSubnetArrayOutput)
}

// NetworkPrivateSubnetMapInput is an input type that accepts NetworkPrivateSubnetMap and NetworkPrivateSubnetMapOutput values.
// You can construct a concrete instance of `NetworkPrivateSubnetMapInput` via:
//
//	NetworkPrivateSubnetMap{ "key": NetworkPrivateSubnetArgs{...} }
type NetworkPrivateSubnetMapInput interface {
	pulumi.Input

	ToNetworkPrivateSubnetMapOutput() NetworkPrivateSubnetMapOutput
	ToNetworkPrivateSubnetMapOutputWithContext(context.Context) NetworkPrivateSubnetMapOutput
}

type NetworkPrivateSubnetMap map[string]NetworkPrivateSubnetInput

func (NetworkPrivateSubnetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPrivateSubnet)(nil)).Elem()
}

func (i NetworkPrivateSubnetMap) ToNetworkPrivateSubnetMapOutput() NetworkPrivateSubnetMapOutput {
	return i.ToNetworkPrivateSubnetMapOutputWithContext(context.Background())
}

func (i NetworkPrivateSubnetMap) ToNetworkPrivateSubnetMapOutputWithContext(ctx context.Context) NetworkPrivateSubnetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPrivateSubnetMapOutput)
}

type NetworkPrivateSubnetOutput struct{ *pulumi.OutputState }

func (NetworkPrivateSubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPrivateSubnet)(nil)).Elem()
}

func (o NetworkPrivateSubnetOutput) ToNetworkPrivateSubnetOutput() NetworkPrivateSubnetOutput {
	return o
}

func (o NetworkPrivateSubnetOutput) ToNetworkPrivateSubnetOutputWithContext(ctx context.Context) NetworkPrivateSubnetOutput {
	return o
}

// Ip Block representing the subnet cidr.
func (o NetworkPrivateSubnetOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPrivateSubnet) pulumi.StringOutput { return v.Cidr }).(pulumi.StringOutput)
}

// Enable DHCP.
// Changing this forces a new resource to be created. Defaults to false.
func (o NetworkPrivateSubnetOutput) Dhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkPrivateSubnet) pulumi.BoolPtrOutput { return v.Dhcp }).(pulumi.BoolPtrOutput)
}

// Last ip for this region.
// Changing this value recreates the subnet.
func (o NetworkPrivateSubnetOutput) End() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPrivateSubnet) pulumi.StringOutput { return v.End }).(pulumi.StringOutput)
}

// The IP of the gateway
func (o NetworkPrivateSubnetOutput) GatewayIp() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPrivateSubnet) pulumi.StringOutput { return v.GatewayIp }).(pulumi.StringOutput)
}

// List of ip pools allocated in the subnet.
// * `ip_pools/network` - Global network with cidr.
// * `ip_pools/region` - Region where this subnet is created.
// * `ip_pools/dhcp` - DHCP enabled.
// * `ip_pools/end` - Last ip for this region.
// * `ip_pools/start` - First ip for this region.
func (o NetworkPrivateSubnetOutput) IpPools() NetworkPrivateSubnetIpPoolArrayOutput {
	return o.ApplyT(func(v *NetworkPrivateSubnet) NetworkPrivateSubnetIpPoolArrayOutput { return v.IpPools }).(NetworkPrivateSubnetIpPoolArrayOutput)
}

// Global network in CIDR format.
// Changing this value recreates the subnet
func (o NetworkPrivateSubnetOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPrivateSubnet) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// The id of the network.
// Changing this forces a new resource to be created.
func (o NetworkPrivateSubnetOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPrivateSubnet) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// Set to true if you don't want to set a default gateway IP.
// Changing this value recreates the resource. Defaults to false.
func (o NetworkPrivateSubnetOutput) NoGateway() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkPrivateSubnet) pulumi.BoolPtrOutput { return v.NoGateway }).(pulumi.BoolPtrOutput)
}

// The region in which the network subnet will be created.
// Ex.: "GRA1". Changing this value recreates the resource.
func (o NetworkPrivateSubnetOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPrivateSubnet) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o NetworkPrivateSubnetOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPrivateSubnet) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// First ip for this region.
// Changing this value recreates the subnet.
func (o NetworkPrivateSubnetOutput) Start() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPrivateSubnet) pulumi.StringOutput { return v.Start }).(pulumi.StringOutput)
}

type NetworkPrivateSubnetArrayOutput struct{ *pulumi.OutputState }

func (NetworkPrivateSubnetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPrivateSubnet)(nil)).Elem()
}

func (o NetworkPrivateSubnetArrayOutput) ToNetworkPrivateSubnetArrayOutput() NetworkPrivateSubnetArrayOutput {
	return o
}

func (o NetworkPrivateSubnetArrayOutput) ToNetworkPrivateSubnetArrayOutputWithContext(ctx context.Context) NetworkPrivateSubnetArrayOutput {
	return o
}

func (o NetworkPrivateSubnetArrayOutput) Index(i pulumi.IntInput) NetworkPrivateSubnetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkPrivateSubnet {
		return vs[0].([]*NetworkPrivateSubnet)[vs[1].(int)]
	}).(NetworkPrivateSubnetOutput)
}

type NetworkPrivateSubnetMapOutput struct{ *pulumi.OutputState }

func (NetworkPrivateSubnetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPrivateSubnet)(nil)).Elem()
}

func (o NetworkPrivateSubnetMapOutput) ToNetworkPrivateSubnetMapOutput() NetworkPrivateSubnetMapOutput {
	return o
}

func (o NetworkPrivateSubnetMapOutput) ToNetworkPrivateSubnetMapOutputWithContext(ctx context.Context) NetworkPrivateSubnetMapOutput {
	return o
}

func (o NetworkPrivateSubnetMapOutput) MapIndex(k pulumi.StringInput) NetworkPrivateSubnetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkPrivateSubnet {
		return vs[0].(map[string]*NetworkPrivateSubnet)[vs[1].(string)]
	}).(NetworkPrivateSubnetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPrivateSubnetInput)(nil)).Elem(), &NetworkPrivateSubnet{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPrivateSubnetArrayInput)(nil)).Elem(), NetworkPrivateSubnetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPrivateSubnetMapInput)(nil)).Elem(), NetworkPrivateSubnetMap{})
	pulumi.RegisterOutputType(NetworkPrivateSubnetOutput{})
	pulumi.RegisterOutputType(NetworkPrivateSubnetArrayOutput{})
	pulumi.RegisterOutputType(NetworkPrivateSubnetMapOutput{})
}
