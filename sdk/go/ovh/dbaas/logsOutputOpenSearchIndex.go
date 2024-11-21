// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a DBaaS Logs Opensearch output index.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/Dbaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dbaas.NewLogsOutputOpenSearchIndex(ctx, "index", &Dbaas.LogsOutputOpenSearchIndexArgs{
//				Description: pulumi.String("my opensearch index"),
//				ServiceName: pulumi.String("...."),
//				Suffix:      pulumi.String("index"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type LogsOutputOpenSearchIndex struct {
	pulumi.CustomResourceState

	// If set, notify when size is near 80, 90 or 100 % of its maximum capacity
	AlertNotifyEnabled pulumi.BoolOutput `pulumi:"alertNotifyEnabled"`
	// Index creation
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Current index size (in bytes)
	CurrentSize pulumi.IntOutput `pulumi:"currentSize"`
	// Index description
	Description pulumi.StringOutput `pulumi:"description"`
	// Index ID
	IndexId pulumi.StringOutput `pulumi:"indexId"`
	// Indicates if you are allowed to edit entry
	IsEditable pulumi.BoolOutput `pulumi:"isEditable"`
	// Maximum index size (in bytes)
	MaxSize pulumi.IntOutput `pulumi:"maxSize"`
	// Index name
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of shards
	NbShard pulumi.IntOutput `pulumi:"nbShard"`
	// The service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Index suffix
	Suffix pulumi.StringOutput `pulumi:"suffix"`
	// Index last update
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewLogsOutputOpenSearchIndex registers a new resource with the given unique name, arguments, and options.
func NewLogsOutputOpenSearchIndex(ctx *pulumi.Context,
	name string, args *LogsOutputOpenSearchIndexArgs, opts ...pulumi.ResourceOption) (*LogsOutputOpenSearchIndex, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.NbShard == nil {
		return nil, errors.New("invalid value for required argument 'NbShard'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Suffix == nil {
		return nil, errors.New("invalid value for required argument 'Suffix'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogsOutputOpenSearchIndex
	err := ctx.RegisterResource("ovh:Dbaas/logsOutputOpenSearchIndex:LogsOutputOpenSearchIndex", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogsOutputOpenSearchIndex gets an existing LogsOutputOpenSearchIndex resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogsOutputOpenSearchIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogsOutputOpenSearchIndexState, opts ...pulumi.ResourceOption) (*LogsOutputOpenSearchIndex, error) {
	var resource LogsOutputOpenSearchIndex
	err := ctx.ReadResource("ovh:Dbaas/logsOutputOpenSearchIndex:LogsOutputOpenSearchIndex", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogsOutputOpenSearchIndex resources.
type logsOutputOpenSearchIndexState struct {
	// If set, notify when size is near 80, 90 or 100 % of its maximum capacity
	AlertNotifyEnabled *bool `pulumi:"alertNotifyEnabled"`
	// Index creation
	CreatedAt *string `pulumi:"createdAt"`
	// Current index size (in bytes)
	CurrentSize *int `pulumi:"currentSize"`
	// Index description
	Description *string `pulumi:"description"`
	// Index ID
	IndexId *string `pulumi:"indexId"`
	// Indicates if you are allowed to edit entry
	IsEditable *bool `pulumi:"isEditable"`
	// Maximum index size (in bytes)
	MaxSize *int `pulumi:"maxSize"`
	// Index name
	Name *string `pulumi:"name"`
	// Number of shards
	NbShard *int `pulumi:"nbShard"`
	// The service name
	ServiceName *string `pulumi:"serviceName"`
	// Index suffix
	Suffix *string `pulumi:"suffix"`
	// Index last update
	UpdatedAt *string `pulumi:"updatedAt"`
}

type LogsOutputOpenSearchIndexState struct {
	// If set, notify when size is near 80, 90 or 100 % of its maximum capacity
	AlertNotifyEnabled pulumi.BoolPtrInput
	// Index creation
	CreatedAt pulumi.StringPtrInput
	// Current index size (in bytes)
	CurrentSize pulumi.IntPtrInput
	// Index description
	Description pulumi.StringPtrInput
	// Index ID
	IndexId pulumi.StringPtrInput
	// Indicates if you are allowed to edit entry
	IsEditable pulumi.BoolPtrInput
	// Maximum index size (in bytes)
	MaxSize pulumi.IntPtrInput
	// Index name
	Name pulumi.StringPtrInput
	// Number of shards
	NbShard pulumi.IntPtrInput
	// The service name
	ServiceName pulumi.StringPtrInput
	// Index suffix
	Suffix pulumi.StringPtrInput
	// Index last update
	UpdatedAt pulumi.StringPtrInput
}

func (LogsOutputOpenSearchIndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*logsOutputOpenSearchIndexState)(nil)).Elem()
}

type logsOutputOpenSearchIndexArgs struct {
	// Index description
	Description string `pulumi:"description"`
	// Number of shards
	NbShard int `pulumi:"nbShard"`
	// The service name
	ServiceName string `pulumi:"serviceName"`
	// Index suffix
	Suffix string `pulumi:"suffix"`
}

// The set of arguments for constructing a LogsOutputOpenSearchIndex resource.
type LogsOutputOpenSearchIndexArgs struct {
	// Index description
	Description pulumi.StringInput
	// Number of shards
	NbShard pulumi.IntInput
	// The service name
	ServiceName pulumi.StringInput
	// Index suffix
	Suffix pulumi.StringInput
}

func (LogsOutputOpenSearchIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logsOutputOpenSearchIndexArgs)(nil)).Elem()
}

type LogsOutputOpenSearchIndexInput interface {
	pulumi.Input

	ToLogsOutputOpenSearchIndexOutput() LogsOutputOpenSearchIndexOutput
	ToLogsOutputOpenSearchIndexOutputWithContext(ctx context.Context) LogsOutputOpenSearchIndexOutput
}

func (*LogsOutputOpenSearchIndex) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsOutputOpenSearchIndex)(nil)).Elem()
}

func (i *LogsOutputOpenSearchIndex) ToLogsOutputOpenSearchIndexOutput() LogsOutputOpenSearchIndexOutput {
	return i.ToLogsOutputOpenSearchIndexOutputWithContext(context.Background())
}

func (i *LogsOutputOpenSearchIndex) ToLogsOutputOpenSearchIndexOutputWithContext(ctx context.Context) LogsOutputOpenSearchIndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsOutputOpenSearchIndexOutput)
}

// LogsOutputOpenSearchIndexArrayInput is an input type that accepts LogsOutputOpenSearchIndexArray and LogsOutputOpenSearchIndexArrayOutput values.
// You can construct a concrete instance of `LogsOutputOpenSearchIndexArrayInput` via:
//
//	LogsOutputOpenSearchIndexArray{ LogsOutputOpenSearchIndexArgs{...} }
type LogsOutputOpenSearchIndexArrayInput interface {
	pulumi.Input

	ToLogsOutputOpenSearchIndexArrayOutput() LogsOutputOpenSearchIndexArrayOutput
	ToLogsOutputOpenSearchIndexArrayOutputWithContext(context.Context) LogsOutputOpenSearchIndexArrayOutput
}

type LogsOutputOpenSearchIndexArray []LogsOutputOpenSearchIndexInput

func (LogsOutputOpenSearchIndexArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogsOutputOpenSearchIndex)(nil)).Elem()
}

func (i LogsOutputOpenSearchIndexArray) ToLogsOutputOpenSearchIndexArrayOutput() LogsOutputOpenSearchIndexArrayOutput {
	return i.ToLogsOutputOpenSearchIndexArrayOutputWithContext(context.Background())
}

func (i LogsOutputOpenSearchIndexArray) ToLogsOutputOpenSearchIndexArrayOutputWithContext(ctx context.Context) LogsOutputOpenSearchIndexArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsOutputOpenSearchIndexArrayOutput)
}

// LogsOutputOpenSearchIndexMapInput is an input type that accepts LogsOutputOpenSearchIndexMap and LogsOutputOpenSearchIndexMapOutput values.
// You can construct a concrete instance of `LogsOutputOpenSearchIndexMapInput` via:
//
//	LogsOutputOpenSearchIndexMap{ "key": LogsOutputOpenSearchIndexArgs{...} }
type LogsOutputOpenSearchIndexMapInput interface {
	pulumi.Input

	ToLogsOutputOpenSearchIndexMapOutput() LogsOutputOpenSearchIndexMapOutput
	ToLogsOutputOpenSearchIndexMapOutputWithContext(context.Context) LogsOutputOpenSearchIndexMapOutput
}

type LogsOutputOpenSearchIndexMap map[string]LogsOutputOpenSearchIndexInput

func (LogsOutputOpenSearchIndexMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogsOutputOpenSearchIndex)(nil)).Elem()
}

func (i LogsOutputOpenSearchIndexMap) ToLogsOutputOpenSearchIndexMapOutput() LogsOutputOpenSearchIndexMapOutput {
	return i.ToLogsOutputOpenSearchIndexMapOutputWithContext(context.Background())
}

func (i LogsOutputOpenSearchIndexMap) ToLogsOutputOpenSearchIndexMapOutputWithContext(ctx context.Context) LogsOutputOpenSearchIndexMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsOutputOpenSearchIndexMapOutput)
}

type LogsOutputOpenSearchIndexOutput struct{ *pulumi.OutputState }

func (LogsOutputOpenSearchIndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsOutputOpenSearchIndex)(nil)).Elem()
}

func (o LogsOutputOpenSearchIndexOutput) ToLogsOutputOpenSearchIndexOutput() LogsOutputOpenSearchIndexOutput {
	return o
}

func (o LogsOutputOpenSearchIndexOutput) ToLogsOutputOpenSearchIndexOutputWithContext(ctx context.Context) LogsOutputOpenSearchIndexOutput {
	return o
}

// If set, notify when size is near 80, 90 or 100 % of its maximum capacity
func (o LogsOutputOpenSearchIndexOutput) AlertNotifyEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.BoolOutput { return v.AlertNotifyEnabled }).(pulumi.BoolOutput)
}

// Index creation
func (o LogsOutputOpenSearchIndexOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Current index size (in bytes)
func (o LogsOutputOpenSearchIndexOutput) CurrentSize() pulumi.IntOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.IntOutput { return v.CurrentSize }).(pulumi.IntOutput)
}

// Index description
func (o LogsOutputOpenSearchIndexOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Index ID
func (o LogsOutputOpenSearchIndexOutput) IndexId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.StringOutput { return v.IndexId }).(pulumi.StringOutput)
}

// Indicates if you are allowed to edit entry
func (o LogsOutputOpenSearchIndexOutput) IsEditable() pulumi.BoolOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.BoolOutput { return v.IsEditable }).(pulumi.BoolOutput)
}

// Maximum index size (in bytes)
func (o LogsOutputOpenSearchIndexOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.IntOutput { return v.MaxSize }).(pulumi.IntOutput)
}

// Index name
func (o LogsOutputOpenSearchIndexOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of shards
func (o LogsOutputOpenSearchIndexOutput) NbShard() pulumi.IntOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.IntOutput { return v.NbShard }).(pulumi.IntOutput)
}

// The service name
func (o LogsOutputOpenSearchIndexOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Index suffix
func (o LogsOutputOpenSearchIndexOutput) Suffix() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.StringOutput { return v.Suffix }).(pulumi.StringOutput)
}

// Index last update
func (o LogsOutputOpenSearchIndexOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsOutputOpenSearchIndex) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type LogsOutputOpenSearchIndexArrayOutput struct{ *pulumi.OutputState }

func (LogsOutputOpenSearchIndexArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogsOutputOpenSearchIndex)(nil)).Elem()
}

func (o LogsOutputOpenSearchIndexArrayOutput) ToLogsOutputOpenSearchIndexArrayOutput() LogsOutputOpenSearchIndexArrayOutput {
	return o
}

func (o LogsOutputOpenSearchIndexArrayOutput) ToLogsOutputOpenSearchIndexArrayOutputWithContext(ctx context.Context) LogsOutputOpenSearchIndexArrayOutput {
	return o
}

func (o LogsOutputOpenSearchIndexArrayOutput) Index(i pulumi.IntInput) LogsOutputOpenSearchIndexOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogsOutputOpenSearchIndex {
		return vs[0].([]*LogsOutputOpenSearchIndex)[vs[1].(int)]
	}).(LogsOutputOpenSearchIndexOutput)
}

type LogsOutputOpenSearchIndexMapOutput struct{ *pulumi.OutputState }

func (LogsOutputOpenSearchIndexMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogsOutputOpenSearchIndex)(nil)).Elem()
}

func (o LogsOutputOpenSearchIndexMapOutput) ToLogsOutputOpenSearchIndexMapOutput() LogsOutputOpenSearchIndexMapOutput {
	return o
}

func (o LogsOutputOpenSearchIndexMapOutput) ToLogsOutputOpenSearchIndexMapOutputWithContext(ctx context.Context) LogsOutputOpenSearchIndexMapOutput {
	return o
}

func (o LogsOutputOpenSearchIndexMapOutput) MapIndex(k pulumi.StringInput) LogsOutputOpenSearchIndexOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogsOutputOpenSearchIndex {
		return vs[0].(map[string]*LogsOutputOpenSearchIndex)[vs[1].(string)]
	}).(LogsOutputOpenSearchIndexOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogsOutputOpenSearchIndexInput)(nil)).Elem(), &LogsOutputOpenSearchIndex{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogsOutputOpenSearchIndexArrayInput)(nil)).Elem(), LogsOutputOpenSearchIndexArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogsOutputOpenSearchIndexMapInput)(nil)).Elem(), LogsOutputOpenSearchIndexMap{})
	pulumi.RegisterOutputType(LogsOutputOpenSearchIndexOutput{})
	pulumi.RegisterOutputType(LogsOutputOpenSearchIndexArrayOutput{})
	pulumi.RegisterOutputType(LogsOutputOpenSearchIndexMapOutput{})
}