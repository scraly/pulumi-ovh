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

// Creates an integration for a database cluster associated with a public cloud project.
//
// With this resource you can create an integration for all engine exept `mongodb`.
//
// Please take a look at the list of available `types` in the `Argument references` section in order to know the list of avaulable integrations. For example, thanks to the integration feature you can have your PostgreSQL logs in your OpenSearch Database.
//
// ## Example Usage
//
// Push PostgreSQL logs in an OpenSearch DB:
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
//			dbPostgresql, err := cloudprojectdatabase.GetDatabase(ctx, &cloudprojectdatabase.GetDatabaseArgs{
//				ServiceName: "XXXX",
//				Engine:      "postgresql",
//				Id:          "ZZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			dbOpensearch, err := cloudprojectdatabase.GetDatabase(ctx, &cloudprojectdatabase.GetDatabaseArgs{
//				ServiceName: "XXXX",
//				Engine:      "opensearch",
//				Id:          "ZZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = cloudprojectdatabase.NewIntegration(ctx, "integration", &cloudprojectdatabase.IntegrationArgs{
//				ServiceName:          pulumi.String(dbPostgresql.ServiceName),
//				Engine:               pulumi.String(dbPostgresql.Engine),
//				ClusterId:            pulumi.String(dbPostgresql.Id),
//				SourceServiceId:      pulumi.String(dbPostgresql.Id),
//				DestinationServiceId: pulumi.String(dbOpensearch.Id),
//				Type:                 pulumi.String("opensearchLogs"),
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
// OVHcloud Managed database clusters users can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the user, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:CloudProjectDatabase/integration:Integration my_user service_name/engine/cluster_id/id
// ```
type Integration struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// ID of the destination service.
	DestinationServiceId pulumi.StringOutput `pulumi:"destinationServiceId"`
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// All engines available exept `mongodb`.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Parameters for the integration.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// ID of the source service.
	SourceServiceId pulumi.StringOutput `pulumi:"sourceServiceId"`
	// Current status of the integration.
	Status pulumi.StringOutput `pulumi:"status"`
	// Type of the integration.
	// Available types:
	// * `grafanaDashboard`
	// * `grafanaDatasource`
	// * `kafkaConnect`
	// * `kafkaLogs`
	// * `kafkaMirrorMaker`
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegration registers a new resource with the given unique name, arguments, and options.
func NewIntegration(ctx *pulumi.Context,
	name string, args *IntegrationArgs, opts ...pulumi.ResourceOption) (*Integration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.DestinationServiceId == nil {
		return nil, errors.New("invalid value for required argument 'DestinationServiceId'")
	}
	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.SourceServiceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceServiceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Integration
	err := ctx.RegisterResource("ovh:CloudProjectDatabase/integration:Integration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegration gets an existing Integration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationState, opts ...pulumi.ResourceOption) (*Integration, error) {
	var resource Integration
	err := ctx.ReadResource("ovh:CloudProjectDatabase/integration:Integration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Integration resources.
type integrationState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// ID of the destination service.
	DestinationServiceId *string `pulumi:"destinationServiceId"`
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// All engines available exept `mongodb`.
	Engine *string `pulumi:"engine"`
	// Parameters for the integration.
	Parameters map[string]string `pulumi:"parameters"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// ID of the source service.
	SourceServiceId *string `pulumi:"sourceServiceId"`
	// Current status of the integration.
	Status *string `pulumi:"status"`
	// Type of the integration.
	// Available types:
	// * `grafanaDashboard`
	// * `grafanaDatasource`
	// * `kafkaConnect`
	// * `kafkaLogs`
	// * `kafkaMirrorMaker`
	Type *string `pulumi:"type"`
}

type IntegrationState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// ID of the destination service.
	DestinationServiceId pulumi.StringPtrInput
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// All engines available exept `mongodb`.
	Engine pulumi.StringPtrInput
	// Parameters for the integration.
	Parameters pulumi.StringMapInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// ID of the source service.
	SourceServiceId pulumi.StringPtrInput
	// Current status of the integration.
	Status pulumi.StringPtrInput
	// Type of the integration.
	// Available types:
	// * `grafanaDashboard`
	// * `grafanaDatasource`
	// * `kafkaConnect`
	// * `kafkaLogs`
	// * `kafkaMirrorMaker`
	Type pulumi.StringPtrInput
}

func (IntegrationState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationState)(nil)).Elem()
}

type integrationArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// ID of the destination service.
	DestinationServiceId string `pulumi:"destinationServiceId"`
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// All engines available exept `mongodb`.
	Engine string `pulumi:"engine"`
	// Parameters for the integration.
	Parameters map[string]string `pulumi:"parameters"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// ID of the source service.
	SourceServiceId string `pulumi:"sourceServiceId"`
	// Type of the integration.
	// Available types:
	// * `grafanaDashboard`
	// * `grafanaDatasource`
	// * `kafkaConnect`
	// * `kafkaLogs`
	// * `kafkaMirrorMaker`
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Integration resource.
type IntegrationArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// ID of the destination service.
	DestinationServiceId pulumi.StringInput
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// All engines available exept `mongodb`.
	Engine pulumi.StringInput
	// Parameters for the integration.
	Parameters pulumi.StringMapInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
	// ID of the source service.
	SourceServiceId pulumi.StringInput
	// Type of the integration.
	// Available types:
	// * `grafanaDashboard`
	// * `grafanaDatasource`
	// * `kafkaConnect`
	// * `kafkaLogs`
	// * `kafkaMirrorMaker`
	Type pulumi.StringPtrInput
}

func (IntegrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationArgs)(nil)).Elem()
}

type IntegrationInput interface {
	pulumi.Input

	ToIntegrationOutput() IntegrationOutput
	ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput
}

func (*Integration) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil)).Elem()
}

func (i *Integration) ToIntegrationOutput() IntegrationOutput {
	return i.ToIntegrationOutputWithContext(context.Background())
}

func (i *Integration) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationOutput)
}

// IntegrationArrayInput is an input type that accepts IntegrationArray and IntegrationArrayOutput values.
// You can construct a concrete instance of `IntegrationArrayInput` via:
//
//	IntegrationArray{ IntegrationArgs{...} }
type IntegrationArrayInput interface {
	pulumi.Input

	ToIntegrationArrayOutput() IntegrationArrayOutput
	ToIntegrationArrayOutputWithContext(context.Context) IntegrationArrayOutput
}

type IntegrationArray []IntegrationInput

func (IntegrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Integration)(nil)).Elem()
}

func (i IntegrationArray) ToIntegrationArrayOutput() IntegrationArrayOutput {
	return i.ToIntegrationArrayOutputWithContext(context.Background())
}

func (i IntegrationArray) ToIntegrationArrayOutputWithContext(ctx context.Context) IntegrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationArrayOutput)
}

// IntegrationMapInput is an input type that accepts IntegrationMap and IntegrationMapOutput values.
// You can construct a concrete instance of `IntegrationMapInput` via:
//
//	IntegrationMap{ "key": IntegrationArgs{...} }
type IntegrationMapInput interface {
	pulumi.Input

	ToIntegrationMapOutput() IntegrationMapOutput
	ToIntegrationMapOutputWithContext(context.Context) IntegrationMapOutput
}

type IntegrationMap map[string]IntegrationInput

func (IntegrationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Integration)(nil)).Elem()
}

func (i IntegrationMap) ToIntegrationMapOutput() IntegrationMapOutput {
	return i.ToIntegrationMapOutputWithContext(context.Background())
}

func (i IntegrationMap) ToIntegrationMapOutputWithContext(ctx context.Context) IntegrationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationMapOutput)
}

type IntegrationOutput struct{ *pulumi.OutputState }

func (IntegrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Integration)(nil)).Elem()
}

func (o IntegrationOutput) ToIntegrationOutput() IntegrationOutput {
	return o
}

func (o IntegrationOutput) ToIntegrationOutputWithContext(ctx context.Context) IntegrationOutput {
	return o
}

// Cluster ID.
func (o IntegrationOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// ID of the destination service.
func (o IntegrationOutput) DestinationServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.DestinationServiceId }).(pulumi.StringOutput)
}

// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
// All engines available exept `mongodb`.
func (o IntegrationOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Parameters for the integration.
func (o IntegrationOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o IntegrationOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// ID of the source service.
func (o IntegrationOutput) SourceServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.SourceServiceId }).(pulumi.StringOutput)
}

// Current status of the integration.
func (o IntegrationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Type of the integration.
// Available types:
// * `grafanaDashboard`
// * `grafanaDatasource`
// * `kafkaConnect`
// * `kafkaLogs`
// * `kafkaMirrorMaker`
func (o IntegrationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Integration) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type IntegrationArrayOutput struct{ *pulumi.OutputState }

func (IntegrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Integration)(nil)).Elem()
}

func (o IntegrationArrayOutput) ToIntegrationArrayOutput() IntegrationArrayOutput {
	return o
}

func (o IntegrationArrayOutput) ToIntegrationArrayOutputWithContext(ctx context.Context) IntegrationArrayOutput {
	return o
}

func (o IntegrationArrayOutput) Index(i pulumi.IntInput) IntegrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Integration {
		return vs[0].([]*Integration)[vs[1].(int)]
	}).(IntegrationOutput)
}

type IntegrationMapOutput struct{ *pulumi.OutputState }

func (IntegrationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Integration)(nil)).Elem()
}

func (o IntegrationMapOutput) ToIntegrationMapOutput() IntegrationMapOutput {
	return o
}

func (o IntegrationMapOutput) ToIntegrationMapOutputWithContext(ctx context.Context) IntegrationMapOutput {
	return o
}

func (o IntegrationMapOutput) MapIndex(k pulumi.StringInput) IntegrationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Integration {
		return vs[0].(map[string]*Integration)[vs[1].(string)]
	}).(IntegrationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationInput)(nil)).Elem(), &Integration{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationArrayInput)(nil)).Elem(), IntegrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationMapInput)(nil)).Elem(), IntegrationMap{})
	pulumi.RegisterOutputType(IntegrationOutput{})
	pulumi.RegisterOutputType(IntegrationArrayOutput{})
	pulumi.RegisterOutputType(IntegrationMapOutput{})
}
