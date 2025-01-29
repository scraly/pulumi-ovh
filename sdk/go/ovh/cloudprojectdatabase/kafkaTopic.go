// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a topic for a kafka cluster associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/cloudprojectdatabase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			kafka, err := cloudprojectdatabase.GetDatabase(ctx, &cloudprojectdatabase.GetDatabaseArgs{
//				ServiceName: "XXX",
//				Engine:      "kafka",
//				Id:          "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudprojectdatabase.NewKafkaTopic(ctx, "topic", &cloudprojectdatabase.KafkaTopicArgs{
//				ServiceName:       pulumi.String(kafka.ServiceName),
//				ClusterId:         pulumi.String(kafka.Id),
//				MinInsyncReplicas: pulumi.Int(1),
//				Partitions:        pulumi.Int(3),
//				Replication:       pulumi.Int(2),
//				RetentionBytes:    pulumi.Int(4),
//				RetentionHours:    pulumi.Int(5),
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
// OVHcloud Managed kafka clusters topics can be imported using the `service_name`, `cluster_id` and `id` of the topic, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic my_topic service_name/cluster_id/id
// ```
type KafkaTopic struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Minimum insync replica accepted for this topic. Should be superior to 0
	MinInsyncReplicas pulumi.IntOutput `pulumi:"minInsyncReplicas"`
	// Name of the topic. No spaces allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of partitions for this topic. Should be superior to 0
	Partitions pulumi.IntOutput `pulumi:"partitions"`
	// Number of replication for this topic. Should be superior to 1
	Replication pulumi.IntOutput `pulumi:"replication"`
	// Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
	RetentionBytes pulumi.IntOutput `pulumi:"retentionBytes"`
	// Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
	RetentionHours pulumi.IntOutput `pulumi:"retentionHours"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewKafkaTopic registers a new resource with the given unique name, arguments, and options.
func NewKafkaTopic(ctx *pulumi.Context,
	name string, args *KafkaTopicArgs, opts ...pulumi.ResourceOption) (*KafkaTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource KafkaTopic
	err := ctx.RegisterResource("ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKafkaTopic gets an existing KafkaTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKafkaTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KafkaTopicState, opts ...pulumi.ResourceOption) (*KafkaTopic, error) {
	var resource KafkaTopic
	err := ctx.ReadResource("ovh:CloudProjectDatabase/kafkaTopic:KafkaTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KafkaTopic resources.
type kafkaTopicState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Minimum insync replica accepted for this topic. Should be superior to 0
	MinInsyncReplicas *int `pulumi:"minInsyncReplicas"`
	// Name of the topic. No spaces allowed.
	Name *string `pulumi:"name"`
	// Number of partitions for this topic. Should be superior to 0
	Partitions *int `pulumi:"partitions"`
	// Number of replication for this topic. Should be superior to 1
	Replication *int `pulumi:"replication"`
	// Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
	RetentionBytes *int `pulumi:"retentionBytes"`
	// Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
	RetentionHours *int `pulumi:"retentionHours"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

type KafkaTopicState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Minimum insync replica accepted for this topic. Should be superior to 0
	MinInsyncReplicas pulumi.IntPtrInput
	// Name of the topic. No spaces allowed.
	Name pulumi.StringPtrInput
	// Number of partitions for this topic. Should be superior to 0
	Partitions pulumi.IntPtrInput
	// Number of replication for this topic. Should be superior to 1
	Replication pulumi.IntPtrInput
	// Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
	RetentionBytes pulumi.IntPtrInput
	// Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
	RetentionHours pulumi.IntPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (KafkaTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaTopicState)(nil)).Elem()
}

type kafkaTopicArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Minimum insync replica accepted for this topic. Should be superior to 0
	MinInsyncReplicas *int `pulumi:"minInsyncReplicas"`
	// Name of the topic. No spaces allowed.
	Name *string `pulumi:"name"`
	// Number of partitions for this topic. Should be superior to 0
	Partitions *int `pulumi:"partitions"`
	// Number of replication for this topic. Should be superior to 1
	Replication *int `pulumi:"replication"`
	// Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
	RetentionBytes *int `pulumi:"retentionBytes"`
	// Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
	RetentionHours *int `pulumi:"retentionHours"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a KafkaTopic resource.
type KafkaTopicArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Minimum insync replica accepted for this topic. Should be superior to 0
	MinInsyncReplicas pulumi.IntPtrInput
	// Name of the topic. No spaces allowed.
	Name pulumi.StringPtrInput
	// Number of partitions for this topic. Should be superior to 0
	Partitions pulumi.IntPtrInput
	// Number of replication for this topic. Should be superior to 1
	Replication pulumi.IntPtrInput
	// Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
	RetentionBytes pulumi.IntPtrInput
	// Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
	RetentionHours pulumi.IntPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (KafkaTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kafkaTopicArgs)(nil)).Elem()
}

type KafkaTopicInput interface {
	pulumi.Input

	ToKafkaTopicOutput() KafkaTopicOutput
	ToKafkaTopicOutputWithContext(ctx context.Context) KafkaTopicOutput
}

func (*KafkaTopic) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaTopic)(nil)).Elem()
}

func (i *KafkaTopic) ToKafkaTopicOutput() KafkaTopicOutput {
	return i.ToKafkaTopicOutputWithContext(context.Background())
}

func (i *KafkaTopic) ToKafkaTopicOutputWithContext(ctx context.Context) KafkaTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaTopicOutput)
}

// KafkaTopicArrayInput is an input type that accepts KafkaTopicArray and KafkaTopicArrayOutput values.
// You can construct a concrete instance of `KafkaTopicArrayInput` via:
//
//	KafkaTopicArray{ KafkaTopicArgs{...} }
type KafkaTopicArrayInput interface {
	pulumi.Input

	ToKafkaTopicArrayOutput() KafkaTopicArrayOutput
	ToKafkaTopicArrayOutputWithContext(context.Context) KafkaTopicArrayOutput
}

type KafkaTopicArray []KafkaTopicInput

func (KafkaTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaTopic)(nil)).Elem()
}

func (i KafkaTopicArray) ToKafkaTopicArrayOutput() KafkaTopicArrayOutput {
	return i.ToKafkaTopicArrayOutputWithContext(context.Background())
}

func (i KafkaTopicArray) ToKafkaTopicArrayOutputWithContext(ctx context.Context) KafkaTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaTopicArrayOutput)
}

// KafkaTopicMapInput is an input type that accepts KafkaTopicMap and KafkaTopicMapOutput values.
// You can construct a concrete instance of `KafkaTopicMapInput` via:
//
//	KafkaTopicMap{ "key": KafkaTopicArgs{...} }
type KafkaTopicMapInput interface {
	pulumi.Input

	ToKafkaTopicMapOutput() KafkaTopicMapOutput
	ToKafkaTopicMapOutputWithContext(context.Context) KafkaTopicMapOutput
}

type KafkaTopicMap map[string]KafkaTopicInput

func (KafkaTopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaTopic)(nil)).Elem()
}

func (i KafkaTopicMap) ToKafkaTopicMapOutput() KafkaTopicMapOutput {
	return i.ToKafkaTopicMapOutputWithContext(context.Background())
}

func (i KafkaTopicMap) ToKafkaTopicMapOutputWithContext(ctx context.Context) KafkaTopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KafkaTopicMapOutput)
}

type KafkaTopicOutput struct{ *pulumi.OutputState }

func (KafkaTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KafkaTopic)(nil)).Elem()
}

func (o KafkaTopicOutput) ToKafkaTopicOutput() KafkaTopicOutput {
	return o
}

func (o KafkaTopicOutput) ToKafkaTopicOutputWithContext(ctx context.Context) KafkaTopicOutput {
	return o
}

// Cluster ID.
func (o KafkaTopicOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaTopic) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Minimum insync replica accepted for this topic. Should be superior to 0
func (o KafkaTopicOutput) MinInsyncReplicas() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaTopic) pulumi.IntOutput { return v.MinInsyncReplicas }).(pulumi.IntOutput)
}

// Name of the topic. No spaces allowed.
func (o KafkaTopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaTopic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of partitions for this topic. Should be superior to 0
func (o KafkaTopicOutput) Partitions() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaTopic) pulumi.IntOutput { return v.Partitions }).(pulumi.IntOutput)
}

// Number of replication for this topic. Should be superior to 1
func (o KafkaTopicOutput) Replication() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaTopic) pulumi.IntOutput { return v.Replication }).(pulumi.IntOutput)
}

// Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited
func (o KafkaTopicOutput) RetentionBytes() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaTopic) pulumi.IntOutput { return v.RetentionBytes }).(pulumi.IntOutput)
}

// Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited
func (o KafkaTopicOutput) RetentionHours() pulumi.IntOutput {
	return o.ApplyT(func(v *KafkaTopic) pulumi.IntOutput { return v.RetentionHours }).(pulumi.IntOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o KafkaTopicOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *KafkaTopic) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type KafkaTopicArrayOutput struct{ *pulumi.OutputState }

func (KafkaTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KafkaTopic)(nil)).Elem()
}

func (o KafkaTopicArrayOutput) ToKafkaTopicArrayOutput() KafkaTopicArrayOutput {
	return o
}

func (o KafkaTopicArrayOutput) ToKafkaTopicArrayOutputWithContext(ctx context.Context) KafkaTopicArrayOutput {
	return o
}

func (o KafkaTopicArrayOutput) Index(i pulumi.IntInput) KafkaTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KafkaTopic {
		return vs[0].([]*KafkaTopic)[vs[1].(int)]
	}).(KafkaTopicOutput)
}

type KafkaTopicMapOutput struct{ *pulumi.OutputState }

func (KafkaTopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KafkaTopic)(nil)).Elem()
}

func (o KafkaTopicMapOutput) ToKafkaTopicMapOutput() KafkaTopicMapOutput {
	return o
}

func (o KafkaTopicMapOutput) ToKafkaTopicMapOutputWithContext(ctx context.Context) KafkaTopicMapOutput {
	return o
}

func (o KafkaTopicMapOutput) MapIndex(k pulumi.StringInput) KafkaTopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KafkaTopic {
		return vs[0].(map[string]*KafkaTopic)[vs[1].(string)]
	}).(KafkaTopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaTopicInput)(nil)).Elem(), &KafkaTopic{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaTopicArrayInput)(nil)).Elem(), KafkaTopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KafkaTopicMapInput)(nil)).Elem(), KafkaTopicMap{})
	pulumi.RegisterOutputType(KafkaTopicOutput{})
	pulumi.RegisterOutputType(KafkaTopicArrayOutput{})
	pulumi.RegisterOutputType(KafkaTopicMapOutput{})
}
