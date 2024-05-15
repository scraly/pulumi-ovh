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

// Creates a database for a database cluster associated with a public cloud project.
//
// With this resource you can create a database for the following database engine:
//
//   - `mysql`
//   - `postgresql`
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProjectDatabase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			db, err := CloudProjectDatabase.GetDatabase(ctx, &cloudprojectdatabase.GetDatabaseArgs{
//				ServiceName: "XXXX",
//				Engine:      "YYYY",
//				Id:          "ZZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = CloudProjectDatabase.NewDatabaseInstance(ctx, "database", &CloudProjectDatabase.DatabaseInstanceArgs{
//				ServiceName: pulumi.String(db.ServiceName),
//				Engine:      pulumi.String(db.Engine),
//				ClusterId:   pulumi.String(db.Id),
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
// OVHcloud Managed database clusters databases can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the database, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:CloudProjectDatabase/databaseInstance:DatabaseInstance my_database service_name/engine/cluster_id/id
// ```
type DatabaseInstance struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Defines if the database has been created by default.
	Default pulumi.BoolOutput `pulumi:"default"`
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Name of the database.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewDatabaseInstance registers a new resource with the given unique name, arguments, and options.
func NewDatabaseInstance(ctx *pulumi.Context,
	name string, args *DatabaseInstanceArgs, opts ...pulumi.ResourceOption) (*DatabaseInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DatabaseInstance
	err := ctx.RegisterResource("ovh:CloudProjectDatabase/databaseInstance:DatabaseInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseInstance gets an existing DatabaseInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseInstanceState, opts ...pulumi.ResourceOption) (*DatabaseInstance, error) {
	var resource DatabaseInstance
	err := ctx.ReadResource("ovh:CloudProjectDatabase/databaseInstance:DatabaseInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseInstance resources.
type databaseInstanceState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Defines if the database has been created by default.
	Default *bool `pulumi:"default"`
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine *string `pulumi:"engine"`
	// Name of the database.
	Name *string `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

type DatabaseInstanceState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Defines if the database has been created by default.
	Default pulumi.BoolPtrInput
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine pulumi.StringPtrInput
	// Name of the database.
	Name pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (DatabaseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceState)(nil)).Elem()
}

type databaseInstanceArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine string `pulumi:"engine"`
	// Name of the database.
	Name *string `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a DatabaseInstance resource.
type DatabaseInstanceArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine pulumi.StringInput
	// Name of the database.
	Name pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (DatabaseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseInstanceArgs)(nil)).Elem()
}

type DatabaseInstanceInput interface {
	pulumi.Input

	ToDatabaseInstanceOutput() DatabaseInstanceOutput
	ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput
}

func (*DatabaseInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseInstance)(nil)).Elem()
}

func (i *DatabaseInstance) ToDatabaseInstanceOutput() DatabaseInstanceOutput {
	return i.ToDatabaseInstanceOutputWithContext(context.Background())
}

func (i *DatabaseInstance) ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceOutput)
}

// DatabaseInstanceArrayInput is an input type that accepts DatabaseInstanceArray and DatabaseInstanceArrayOutput values.
// You can construct a concrete instance of `DatabaseInstanceArrayInput` via:
//
//	DatabaseInstanceArray{ DatabaseInstanceArgs{...} }
type DatabaseInstanceArrayInput interface {
	pulumi.Input

	ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput
	ToDatabaseInstanceArrayOutputWithContext(context.Context) DatabaseInstanceArrayOutput
}

type DatabaseInstanceArray []DatabaseInstanceInput

func (DatabaseInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseInstance)(nil)).Elem()
}

func (i DatabaseInstanceArray) ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput {
	return i.ToDatabaseInstanceArrayOutputWithContext(context.Background())
}

func (i DatabaseInstanceArray) ToDatabaseInstanceArrayOutputWithContext(ctx context.Context) DatabaseInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceArrayOutput)
}

// DatabaseInstanceMapInput is an input type that accepts DatabaseInstanceMap and DatabaseInstanceMapOutput values.
// You can construct a concrete instance of `DatabaseInstanceMapInput` via:
//
//	DatabaseInstanceMap{ "key": DatabaseInstanceArgs{...} }
type DatabaseInstanceMapInput interface {
	pulumi.Input

	ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput
	ToDatabaseInstanceMapOutputWithContext(context.Context) DatabaseInstanceMapOutput
}

type DatabaseInstanceMap map[string]DatabaseInstanceInput

func (DatabaseInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseInstance)(nil)).Elem()
}

func (i DatabaseInstanceMap) ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput {
	return i.ToDatabaseInstanceMapOutputWithContext(context.Background())
}

func (i DatabaseInstanceMap) ToDatabaseInstanceMapOutputWithContext(ctx context.Context) DatabaseInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseInstanceMapOutput)
}

type DatabaseInstanceOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseInstance)(nil)).Elem()
}

func (o DatabaseInstanceOutput) ToDatabaseInstanceOutput() DatabaseInstanceOutput {
	return o
}

func (o DatabaseInstanceOutput) ToDatabaseInstanceOutputWithContext(ctx context.Context) DatabaseInstanceOutput {
	return o
}

// Cluster ID.
func (o DatabaseInstanceOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Defines if the database has been created by default.
func (o DatabaseInstanceOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.BoolOutput { return v.Default }).(pulumi.BoolOutput)
}

// The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
// Available engines:
func (o DatabaseInstanceOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// Name of the database.
func (o DatabaseInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o DatabaseInstanceOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *DatabaseInstance) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type DatabaseInstanceArrayOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DatabaseInstance)(nil)).Elem()
}

func (o DatabaseInstanceArrayOutput) ToDatabaseInstanceArrayOutput() DatabaseInstanceArrayOutput {
	return o
}

func (o DatabaseInstanceArrayOutput) ToDatabaseInstanceArrayOutputWithContext(ctx context.Context) DatabaseInstanceArrayOutput {
	return o
}

func (o DatabaseInstanceArrayOutput) Index(i pulumi.IntInput) DatabaseInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DatabaseInstance {
		return vs[0].([]*DatabaseInstance)[vs[1].(int)]
	}).(DatabaseInstanceOutput)
}

type DatabaseInstanceMapOutput struct{ *pulumi.OutputState }

func (DatabaseInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DatabaseInstance)(nil)).Elem()
}

func (o DatabaseInstanceMapOutput) ToDatabaseInstanceMapOutput() DatabaseInstanceMapOutput {
	return o
}

func (o DatabaseInstanceMapOutput) ToDatabaseInstanceMapOutputWithContext(ctx context.Context) DatabaseInstanceMapOutput {
	return o
}

func (o DatabaseInstanceMapOutput) MapIndex(k pulumi.StringInput) DatabaseInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DatabaseInstance {
		return vs[0].(map[string]*DatabaseInstance)[vs[1].(string)]
	}).(DatabaseInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInstanceInput)(nil)).Elem(), &DatabaseInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInstanceArrayInput)(nil)).Elem(), DatabaseInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInstanceMapInput)(nil)).Elem(), DatabaseInstanceMap{})
	pulumi.RegisterOutputType(DatabaseInstanceOutput{})
	pulumi.RegisterOutputType(DatabaseInstanceArrayOutput{})
	pulumi.RegisterOutputType(DatabaseInstanceMapOutput{})
}
