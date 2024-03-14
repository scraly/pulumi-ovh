// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a user of a database cluster associated with a public cloud project.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
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
//			user, err := CloudProjectDatabase.GetUser(ctx, &cloudprojectdatabase.GetUserArgs{
//				ServiceName: "XXX",
//				Engine:      "YYY",
//				ClusterId:   "ZZZ",
//				Name:        "UUU",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("userName", user.Name)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupUser(ctx *pulumi.Context, args *LookupUserArgs, opts ...pulumi.InvokeOption) (*LookupUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupUserResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type LookupUserArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// The engine of the database cluster you want user information. To get a full list of available engine visit :
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine string `pulumi:"engine"`
	// Name of the user.
	Name string `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getUser.
type LookupUserResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// Date of the creation of the user.
	CreatedAt string `pulumi:"createdAt"`
	// See Argument Reference above.
	Engine string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the user.
	Name string `pulumi:"name"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
	// Current status of the user.
	Status string `pulumi:"status"`
}

func LookupUserOutput(ctx *pulumi.Context, args LookupUserOutputArgs, opts ...pulumi.InvokeOption) LookupUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupUserResult, error) {
			args := v.(LookupUserArgs)
			r, err := LookupUser(ctx, &args, opts...)
			var s LookupUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupUserResultOutput)
}

// A collection of arguments for invoking getUser.
type LookupUserOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The engine of the database cluster you want user information. To get a full list of available engine visit :
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	// Available engines:
	Engine pulumi.StringInput `pulumi:"engine"`
	// Name of the user.
	Name pulumi.StringInput `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type LookupUserResultOutput struct{ *pulumi.OutputState }

func (LookupUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupUserResult)(nil)).Elem()
}

func (o LookupUserResultOutput) ToLookupUserResultOutput() LookupUserResultOutput {
	return o
}

func (o LookupUserResultOutput) ToLookupUserResultOutputWithContext(ctx context.Context) LookupUserResultOutput {
	return o
}

// See Argument Reference above.
func (o LookupUserResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Date of the creation of the user.
func (o LookupUserResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupUserResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the user.
func (o LookupUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o LookupUserResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the user.
func (o LookupUserResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupUserResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupUserResultOutput{})
}
