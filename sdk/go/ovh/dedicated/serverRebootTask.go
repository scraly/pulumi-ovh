// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicated

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
type ServerRebootTask struct {
	pulumi.CustomResourceState

	// Details of this task. (should be `Reboot asked`)
	Comment pulumi.StringOutput `pulumi:"comment"`
	// Completion date in RFC3339 format.
	DoneDate pulumi.StringOutput `pulumi:"doneDate"`
	// Function name (should be `hardReboot`).
	Function pulumi.StringOutput `pulumi:"function"`
	// List of values tracked to trigger reboot, used also to form implicit dependencies.
	Keepers pulumi.StringArrayOutput `pulumi:"keepers"`
	// Last update in RFC3339 format.
	LastUpdate pulumi.StringOutput `pulumi:"lastUpdate"`
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Task creation date in RFC3339 format.
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// Task status (should be `done`)
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewServerRebootTask registers a new resource with the given unique name, arguments, and options.
func NewServerRebootTask(ctx *pulumi.Context,
	name string, args *ServerRebootTaskArgs, opts ...pulumi.ResourceOption) (*ServerRebootTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Keepers == nil {
		return nil, errors.New("invalid value for required argument 'Keepers'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServerRebootTask
	err := ctx.RegisterResource("ovh:Dedicated/serverRebootTask:ServerRebootTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerRebootTask gets an existing ServerRebootTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerRebootTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerRebootTaskState, opts ...pulumi.ResourceOption) (*ServerRebootTask, error) {
	var resource ServerRebootTask
	err := ctx.ReadResource("ovh:Dedicated/serverRebootTask:ServerRebootTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerRebootTask resources.
type serverRebootTaskState struct {
	// Details of this task. (should be `Reboot asked`)
	Comment *string `pulumi:"comment"`
	// Completion date in RFC3339 format.
	DoneDate *string `pulumi:"doneDate"`
	// Function name (should be `hardReboot`).
	Function *string `pulumi:"function"`
	// List of values tracked to trigger reboot, used also to form implicit dependencies.
	Keepers []string `pulumi:"keepers"`
	// Last update in RFC3339 format.
	LastUpdate *string `pulumi:"lastUpdate"`
	// The serviceName of your dedicated server.
	ServiceName *string `pulumi:"serviceName"`
	// Task creation date in RFC3339 format.
	StartDate *string `pulumi:"startDate"`
	// Task status (should be `done`)
	Status *string `pulumi:"status"`
}

type ServerRebootTaskState struct {
	// Details of this task. (should be `Reboot asked`)
	Comment pulumi.StringPtrInput
	// Completion date in RFC3339 format.
	DoneDate pulumi.StringPtrInput
	// Function name (should be `hardReboot`).
	Function pulumi.StringPtrInput
	// List of values tracked to trigger reboot, used also to form implicit dependencies.
	Keepers pulumi.StringArrayInput
	// Last update in RFC3339 format.
	LastUpdate pulumi.StringPtrInput
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringPtrInput
	// Task creation date in RFC3339 format.
	StartDate pulumi.StringPtrInput
	// Task status (should be `done`)
	Status pulumi.StringPtrInput
}

func (ServerRebootTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverRebootTaskState)(nil)).Elem()
}

type serverRebootTaskArgs struct {
	// List of values tracked to trigger reboot, used also to form implicit dependencies.
	Keepers []string `pulumi:"keepers"`
	// The serviceName of your dedicated server.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ServerRebootTask resource.
type ServerRebootTaskArgs struct {
	// List of values tracked to trigger reboot, used also to form implicit dependencies.
	Keepers pulumi.StringArrayInput
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringInput
}

func (ServerRebootTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverRebootTaskArgs)(nil)).Elem()
}

type ServerRebootTaskInput interface {
	pulumi.Input

	ToServerRebootTaskOutput() ServerRebootTaskOutput
	ToServerRebootTaskOutputWithContext(ctx context.Context) ServerRebootTaskOutput
}

func (*ServerRebootTask) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerRebootTask)(nil)).Elem()
}

func (i *ServerRebootTask) ToServerRebootTaskOutput() ServerRebootTaskOutput {
	return i.ToServerRebootTaskOutputWithContext(context.Background())
}

func (i *ServerRebootTask) ToServerRebootTaskOutputWithContext(ctx context.Context) ServerRebootTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerRebootTaskOutput)
}

func (i *ServerRebootTask) ToOutput(ctx context.Context) pulumix.Output[*ServerRebootTask] {
	return pulumix.Output[*ServerRebootTask]{
		OutputState: i.ToServerRebootTaskOutputWithContext(ctx).OutputState,
	}
}

// ServerRebootTaskArrayInput is an input type that accepts ServerRebootTaskArray and ServerRebootTaskArrayOutput values.
// You can construct a concrete instance of `ServerRebootTaskArrayInput` via:
//
//	ServerRebootTaskArray{ ServerRebootTaskArgs{...} }
type ServerRebootTaskArrayInput interface {
	pulumi.Input

	ToServerRebootTaskArrayOutput() ServerRebootTaskArrayOutput
	ToServerRebootTaskArrayOutputWithContext(context.Context) ServerRebootTaskArrayOutput
}

type ServerRebootTaskArray []ServerRebootTaskInput

func (ServerRebootTaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerRebootTask)(nil)).Elem()
}

func (i ServerRebootTaskArray) ToServerRebootTaskArrayOutput() ServerRebootTaskArrayOutput {
	return i.ToServerRebootTaskArrayOutputWithContext(context.Background())
}

func (i ServerRebootTaskArray) ToServerRebootTaskArrayOutputWithContext(ctx context.Context) ServerRebootTaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerRebootTaskArrayOutput)
}

func (i ServerRebootTaskArray) ToOutput(ctx context.Context) pulumix.Output[[]*ServerRebootTask] {
	return pulumix.Output[[]*ServerRebootTask]{
		OutputState: i.ToServerRebootTaskArrayOutputWithContext(ctx).OutputState,
	}
}

// ServerRebootTaskMapInput is an input type that accepts ServerRebootTaskMap and ServerRebootTaskMapOutput values.
// You can construct a concrete instance of `ServerRebootTaskMapInput` via:
//
//	ServerRebootTaskMap{ "key": ServerRebootTaskArgs{...} }
type ServerRebootTaskMapInput interface {
	pulumi.Input

	ToServerRebootTaskMapOutput() ServerRebootTaskMapOutput
	ToServerRebootTaskMapOutputWithContext(context.Context) ServerRebootTaskMapOutput
}

type ServerRebootTaskMap map[string]ServerRebootTaskInput

func (ServerRebootTaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerRebootTask)(nil)).Elem()
}

func (i ServerRebootTaskMap) ToServerRebootTaskMapOutput() ServerRebootTaskMapOutput {
	return i.ToServerRebootTaskMapOutputWithContext(context.Background())
}

func (i ServerRebootTaskMap) ToServerRebootTaskMapOutputWithContext(ctx context.Context) ServerRebootTaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerRebootTaskMapOutput)
}

func (i ServerRebootTaskMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServerRebootTask] {
	return pulumix.Output[map[string]*ServerRebootTask]{
		OutputState: i.ToServerRebootTaskMapOutputWithContext(ctx).OutputState,
	}
}

type ServerRebootTaskOutput struct{ *pulumi.OutputState }

func (ServerRebootTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerRebootTask)(nil)).Elem()
}

func (o ServerRebootTaskOutput) ToServerRebootTaskOutput() ServerRebootTaskOutput {
	return o
}

func (o ServerRebootTaskOutput) ToServerRebootTaskOutputWithContext(ctx context.Context) ServerRebootTaskOutput {
	return o
}

func (o ServerRebootTaskOutput) ToOutput(ctx context.Context) pulumix.Output[*ServerRebootTask] {
	return pulumix.Output[*ServerRebootTask]{
		OutputState: o.OutputState,
	}
}

// Details of this task. (should be `Reboot asked`)
func (o ServerRebootTaskOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerRebootTask) pulumi.StringOutput { return v.Comment }).(pulumi.StringOutput)
}

// Completion date in RFC3339 format.
func (o ServerRebootTaskOutput) DoneDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerRebootTask) pulumi.StringOutput { return v.DoneDate }).(pulumi.StringOutput)
}

// Function name (should be `hardReboot`).
func (o ServerRebootTaskOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerRebootTask) pulumi.StringOutput { return v.Function }).(pulumi.StringOutput)
}

// List of values tracked to trigger reboot, used also to form implicit dependencies.
func (o ServerRebootTaskOutput) Keepers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServerRebootTask) pulumi.StringArrayOutput { return v.Keepers }).(pulumi.StringArrayOutput)
}

// Last update in RFC3339 format.
func (o ServerRebootTaskOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerRebootTask) pulumi.StringOutput { return v.LastUpdate }).(pulumi.StringOutput)
}

// The serviceName of your dedicated server.
func (o ServerRebootTaskOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerRebootTask) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Task creation date in RFC3339 format.
func (o ServerRebootTaskOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerRebootTask) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// Task status (should be `done`)
func (o ServerRebootTaskOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerRebootTask) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type ServerRebootTaskArrayOutput struct{ *pulumi.OutputState }

func (ServerRebootTaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerRebootTask)(nil)).Elem()
}

func (o ServerRebootTaskArrayOutput) ToServerRebootTaskArrayOutput() ServerRebootTaskArrayOutput {
	return o
}

func (o ServerRebootTaskArrayOutput) ToServerRebootTaskArrayOutputWithContext(ctx context.Context) ServerRebootTaskArrayOutput {
	return o
}

func (o ServerRebootTaskArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ServerRebootTask] {
	return pulumix.Output[[]*ServerRebootTask]{
		OutputState: o.OutputState,
	}
}

func (o ServerRebootTaskArrayOutput) Index(i pulumi.IntInput) ServerRebootTaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerRebootTask {
		return vs[0].([]*ServerRebootTask)[vs[1].(int)]
	}).(ServerRebootTaskOutput)
}

type ServerRebootTaskMapOutput struct{ *pulumi.OutputState }

func (ServerRebootTaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerRebootTask)(nil)).Elem()
}

func (o ServerRebootTaskMapOutput) ToServerRebootTaskMapOutput() ServerRebootTaskMapOutput {
	return o
}

func (o ServerRebootTaskMapOutput) ToServerRebootTaskMapOutputWithContext(ctx context.Context) ServerRebootTaskMapOutput {
	return o
}

func (o ServerRebootTaskMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ServerRebootTask] {
	return pulumix.Output[map[string]*ServerRebootTask]{
		OutputState: o.OutputState,
	}
}

func (o ServerRebootTaskMapOutput) MapIndex(k pulumi.StringInput) ServerRebootTaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerRebootTask {
		return vs[0].(map[string]*ServerRebootTask)[vs[1].(string)]
	}).(ServerRebootTaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerRebootTaskInput)(nil)).Elem(), &ServerRebootTask{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerRebootTaskArrayInput)(nil)).Elem(), ServerRebootTaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerRebootTaskMapInput)(nil)).Elem(), ServerRebootTaskMap{})
	pulumi.RegisterOutputType(ServerRebootTaskOutput{})
	pulumi.RegisterOutputType(ServerRebootTaskArrayOutput{})
	pulumi.RegisterOutputType(ServerRebootTaskMapOutput{})
}
