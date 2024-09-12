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

// Creates a HTTP backend server group (farm) to be used by loadbalancing frontend(s)
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
//			_, err = IpLoadBalancing.NewHttpFarm(ctx, "farmname", &IpLoadBalancing.HttpFarmArgs{
//				DisplayName: pulumi.String("ingress-8080-gra"),
//				ServiceName: pulumi.String(lb.ServiceName),
//				Zone:        pulumi.String("GRA"),
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
// HTTP farm can be imported using the following format `serviceName` and the `id` of the farm, separated by "/" e.g.
type HttpFarm struct {
	pulumi.CustomResourceState

	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`, `uri`)
	Balance pulumi.StringPtrOutput `pulumi:"balance"`
	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Port attached to your farm ([1..49151]). Inherited from frontend if null
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// define a backend healthcheck probe
	Probe HttpFarmProbePtrOutput `pulumi:"probe"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
	Stickiness pulumi.StringPtrOutput `pulumi:"stickiness"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.IntPtrOutput `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewHttpFarm registers a new resource with the given unique name, arguments, and options.
func NewHttpFarm(ctx *pulumi.Context,
	name string, args *HttpFarmArgs, opts ...pulumi.ResourceOption) (*HttpFarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource HttpFarm
	err := ctx.RegisterResource("ovh:IpLoadBalancing/httpFarm:HttpFarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHttpFarm gets an existing HttpFarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHttpFarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HttpFarmState, opts ...pulumi.ResourceOption) (*HttpFarm, error) {
	var resource HttpFarm
	err := ctx.ReadResource("ovh:IpLoadBalancing/httpFarm:HttpFarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HttpFarm resources.
type httpFarmState struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`, `uri`)
	Balance *string `pulumi:"balance"`
	// Readable label for loadbalancer farm
	DisplayName *string `pulumi:"displayName"`
	// Port attached to your farm ([1..49151]). Inherited from frontend if null
	Port *int `pulumi:"port"`
	// define a backend healthcheck probe
	Probe *HttpFarmProbe `pulumi:"probe"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
	Stickiness *string `pulumi:"stickiness"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId *int `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone *string `pulumi:"zone"`
}

type HttpFarmState struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`, `uri`)
	Balance pulumi.StringPtrInput
	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrInput
	// Port attached to your farm ([1..49151]). Inherited from frontend if null
	Port pulumi.IntPtrInput
	// define a backend healthcheck probe
	Probe HttpFarmProbePtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
	Stickiness pulumi.StringPtrInput
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.IntPtrInput
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone pulumi.StringPtrInput
}

func (HttpFarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*httpFarmState)(nil)).Elem()
}

type httpFarmArgs struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`, `uri`)
	Balance *string `pulumi:"balance"`
	// Readable label for loadbalancer farm
	DisplayName *string `pulumi:"displayName"`
	// Port attached to your farm ([1..49151]). Inherited from frontend if null
	Port *int `pulumi:"port"`
	// define a backend healthcheck probe
	Probe *HttpFarmProbe `pulumi:"probe"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
	Stickiness *string `pulumi:"stickiness"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId *int `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a HttpFarm resource.
type HttpFarmArgs struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`, `uri`)
	Balance pulumi.StringPtrInput
	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrInput
	// Port attached to your farm ([1..49151]). Inherited from frontend if null
	Port pulumi.IntPtrInput
	// define a backend healthcheck probe
	Probe HttpFarmProbePtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
	Stickiness pulumi.StringPtrInput
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.IntPtrInput
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone pulumi.StringInput
}

func (HttpFarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*httpFarmArgs)(nil)).Elem()
}

type HttpFarmInput interface {
	pulumi.Input

	ToHttpFarmOutput() HttpFarmOutput
	ToHttpFarmOutputWithContext(ctx context.Context) HttpFarmOutput
}

func (*HttpFarm) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpFarm)(nil)).Elem()
}

func (i *HttpFarm) ToHttpFarmOutput() HttpFarmOutput {
	return i.ToHttpFarmOutputWithContext(context.Background())
}

func (i *HttpFarm) ToHttpFarmOutputWithContext(ctx context.Context) HttpFarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpFarmOutput)
}

// HttpFarmArrayInput is an input type that accepts HttpFarmArray and HttpFarmArrayOutput values.
// You can construct a concrete instance of `HttpFarmArrayInput` via:
//
//	HttpFarmArray{ HttpFarmArgs{...} }
type HttpFarmArrayInput interface {
	pulumi.Input

	ToHttpFarmArrayOutput() HttpFarmArrayOutput
	ToHttpFarmArrayOutputWithContext(context.Context) HttpFarmArrayOutput
}

type HttpFarmArray []HttpFarmInput

func (HttpFarmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpFarm)(nil)).Elem()
}

func (i HttpFarmArray) ToHttpFarmArrayOutput() HttpFarmArrayOutput {
	return i.ToHttpFarmArrayOutputWithContext(context.Background())
}

func (i HttpFarmArray) ToHttpFarmArrayOutputWithContext(ctx context.Context) HttpFarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpFarmArrayOutput)
}

// HttpFarmMapInput is an input type that accepts HttpFarmMap and HttpFarmMapOutput values.
// You can construct a concrete instance of `HttpFarmMapInput` via:
//
//	HttpFarmMap{ "key": HttpFarmArgs{...} }
type HttpFarmMapInput interface {
	pulumi.Input

	ToHttpFarmMapOutput() HttpFarmMapOutput
	ToHttpFarmMapOutputWithContext(context.Context) HttpFarmMapOutput
}

type HttpFarmMap map[string]HttpFarmInput

func (HttpFarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpFarm)(nil)).Elem()
}

func (i HttpFarmMap) ToHttpFarmMapOutput() HttpFarmMapOutput {
	return i.ToHttpFarmMapOutputWithContext(context.Background())
}

func (i HttpFarmMap) ToHttpFarmMapOutputWithContext(ctx context.Context) HttpFarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpFarmMapOutput)
}

type HttpFarmOutput struct{ *pulumi.OutputState }

func (HttpFarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpFarm)(nil)).Elem()
}

func (o HttpFarmOutput) ToHttpFarmOutput() HttpFarmOutput {
	return o
}

func (o HttpFarmOutput) ToHttpFarmOutputWithContext(ctx context.Context) HttpFarmOutput {
	return o
}

// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`, `uri`)
func (o HttpFarmOutput) Balance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFarm) pulumi.StringPtrOutput { return v.Balance }).(pulumi.StringPtrOutput)
}

// Readable label for loadbalancer farm
func (o HttpFarmOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFarm) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Port attached to your farm ([1..49151]). Inherited from frontend if null
func (o HttpFarmOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HttpFarm) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// define a backend healthcheck probe
func (o HttpFarmOutput) Probe() HttpFarmProbePtrOutput {
	return o.ApplyT(func(v *HttpFarm) HttpFarmProbePtrOutput { return v.Probe }).(HttpFarmProbePtrOutput)
}

// The internal name of your IP load balancing
func (o HttpFarmOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFarm) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Stickiness type. No stickiness if null (`sourceIp`, `cookie`)
func (o HttpFarmOutput) Stickiness() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HttpFarm) pulumi.StringPtrOutput { return v.Stickiness }).(pulumi.StringPtrOutput)
}

// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
func (o HttpFarmOutput) VrackNetworkId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HttpFarm) pulumi.IntPtrOutput { return v.VrackNetworkId }).(pulumi.IntPtrOutput)
}

// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
func (o HttpFarmOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *HttpFarm) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type HttpFarmArrayOutput struct{ *pulumi.OutputState }

func (HttpFarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HttpFarm)(nil)).Elem()
}

func (o HttpFarmArrayOutput) ToHttpFarmArrayOutput() HttpFarmArrayOutput {
	return o
}

func (o HttpFarmArrayOutput) ToHttpFarmArrayOutputWithContext(ctx context.Context) HttpFarmArrayOutput {
	return o
}

func (o HttpFarmArrayOutput) Index(i pulumi.IntInput) HttpFarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HttpFarm {
		return vs[0].([]*HttpFarm)[vs[1].(int)]
	}).(HttpFarmOutput)
}

type HttpFarmMapOutput struct{ *pulumi.OutputState }

func (HttpFarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HttpFarm)(nil)).Elem()
}

func (o HttpFarmMapOutput) ToHttpFarmMapOutput() HttpFarmMapOutput {
	return o
}

func (o HttpFarmMapOutput) ToHttpFarmMapOutputWithContext(ctx context.Context) HttpFarmMapOutput {
	return o
}

func (o HttpFarmMapOutput) MapIndex(k pulumi.StringInput) HttpFarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HttpFarm {
		return vs[0].(map[string]*HttpFarm)[vs[1].(string)]
	}).(HttpFarmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HttpFarmInput)(nil)).Elem(), &HttpFarm{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpFarmArrayInput)(nil)).Elem(), HttpFarmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HttpFarmMapInput)(nil)).Elem(), HttpFarmMap{})
	pulumi.RegisterOutputType(HttpFarmOutput{})
	pulumi.RegisterOutputType(HttpFarmArrayOutput{})
	pulumi.RegisterOutputType(HttpFarmMapOutput{})
}
