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

// Set the S3 Policy of a public cloud project user.
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			user, err := CloudProject.NewUser(ctx, "user", &CloudProject.UserArgs{
//				ServiceName: pulumi.String("XXX"),
//				Description: pulumi.String("my user"),
//				RoleNames: pulumi.StringArray{
//					pulumi.String("objectstore_operator"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = CloudProject.NewS3Credential(ctx, "myS3Credentials", &CloudProject.S3CredentialArgs{
//				ServiceName: user.ServiceName,
//				UserId:      user.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"Statement": []map[string]interface{}{
//					map[string]interface{}{
//						"Sid":    "RWContainer",
//						"Effect": "Allow",
//						"Action": []string{
//							"s3:GetObject",
//							"s3:PutObject",
//							"s3:DeleteObject",
//							"s3:ListBucket",
//							"s3:ListMultipartUploadParts",
//							"s3:ListBucketMultipartUploads",
//							"s3:AbortMultipartUpload",
//							"s3:GetBucketLocation",
//						},
//						"Resource": []string{
//							"arn:aws:s3:::hp-bucket",
//							"arn:aws:s3:::hp-bucket/*",
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			_, err = CloudProject.NewS3Policy(ctx, "policy", &CloudProject.S3PolicyArgs{
//				ServiceName: user.ServiceName,
//				UserId:      user.ID(),
//				Policy:      pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// OVHcloud User S3 Policy can be imported using the `service_name`, `user_id` of the policy, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:CloudProject/s3Policy:S3Policy policy service_name/user_id
// ```
type S3Policy struct {
	pulumi.CustomResourceState

	// The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
	Policy pulumi.StringOutput `pulumi:"policy"`
	// The ID of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The ID of a public cloud project's user.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewS3Policy registers a new resource with the given unique name, arguments, and options.
func NewS3Policy(ctx *pulumi.Context,
	name string, args *S3PolicyArgs, opts ...pulumi.ResourceOption) (*S3Policy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Policy == nil {
		return nil, errors.New("invalid value for required argument 'Policy'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.UserId == nil {
		return nil, errors.New("invalid value for required argument 'UserId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource S3Policy
	err := ctx.RegisterResource("ovh:CloudProject/s3Policy:S3Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetS3Policy gets an existing S3Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetS3Policy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *S3PolicyState, opts ...pulumi.ResourceOption) (*S3Policy, error) {
	var resource S3Policy
	err := ctx.ReadResource("ovh:CloudProject/s3Policy:S3Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering S3Policy resources.
type s3policyState struct {
	// The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
	Policy *string `pulumi:"policy"`
	// The ID of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// The ID of a public cloud project's user.
	UserId *string `pulumi:"userId"`
}

type S3PolicyState struct {
	// The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
	Policy pulumi.StringPtrInput
	// The ID of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// The ID of a public cloud project's user.
	UserId pulumi.StringPtrInput
}

func (S3PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*s3policyState)(nil)).Elem()
}

type s3policyArgs struct {
	// The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
	Policy string `pulumi:"policy"`
	// The ID of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// The ID of a public cloud project's user.
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a S3Policy resource.
type S3PolicyArgs struct {
	// The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
	Policy pulumi.StringInput
	// The ID of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
	// The ID of a public cloud project's user.
	UserId pulumi.StringInput
}

func (S3PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*s3policyArgs)(nil)).Elem()
}

type S3PolicyInput interface {
	pulumi.Input

	ToS3PolicyOutput() S3PolicyOutput
	ToS3PolicyOutputWithContext(ctx context.Context) S3PolicyOutput
}

func (*S3Policy) ElementType() reflect.Type {
	return reflect.TypeOf((**S3Policy)(nil)).Elem()
}

func (i *S3Policy) ToS3PolicyOutput() S3PolicyOutput {
	return i.ToS3PolicyOutputWithContext(context.Background())
}

func (i *S3Policy) ToS3PolicyOutputWithContext(ctx context.Context) S3PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3PolicyOutput)
}

// S3PolicyArrayInput is an input type that accepts S3PolicyArray and S3PolicyArrayOutput values.
// You can construct a concrete instance of `S3PolicyArrayInput` via:
//
//	S3PolicyArray{ S3PolicyArgs{...} }
type S3PolicyArrayInput interface {
	pulumi.Input

	ToS3PolicyArrayOutput() S3PolicyArrayOutput
	ToS3PolicyArrayOutputWithContext(context.Context) S3PolicyArrayOutput
}

type S3PolicyArray []S3PolicyInput

func (S3PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*S3Policy)(nil)).Elem()
}

func (i S3PolicyArray) ToS3PolicyArrayOutput() S3PolicyArrayOutput {
	return i.ToS3PolicyArrayOutputWithContext(context.Background())
}

func (i S3PolicyArray) ToS3PolicyArrayOutputWithContext(ctx context.Context) S3PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3PolicyArrayOutput)
}

// S3PolicyMapInput is an input type that accepts S3PolicyMap and S3PolicyMapOutput values.
// You can construct a concrete instance of `S3PolicyMapInput` via:
//
//	S3PolicyMap{ "key": S3PolicyArgs{...} }
type S3PolicyMapInput interface {
	pulumi.Input

	ToS3PolicyMapOutput() S3PolicyMapOutput
	ToS3PolicyMapOutputWithContext(context.Context) S3PolicyMapOutput
}

type S3PolicyMap map[string]S3PolicyInput

func (S3PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*S3Policy)(nil)).Elem()
}

func (i S3PolicyMap) ToS3PolicyMapOutput() S3PolicyMapOutput {
	return i.ToS3PolicyMapOutputWithContext(context.Background())
}

func (i S3PolicyMap) ToS3PolicyMapOutputWithContext(ctx context.Context) S3PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(S3PolicyMapOutput)
}

type S3PolicyOutput struct{ *pulumi.OutputState }

func (S3PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**S3Policy)(nil)).Elem()
}

func (o S3PolicyOutput) ToS3PolicyOutput() S3PolicyOutput {
	return o
}

func (o S3PolicyOutput) ToS3PolicyOutputWithContext(ctx context.Context) S3PolicyOutput {
	return o
}

// The policy document. This is a JSON formatted string. See examples of policies on [public documentation](https://docs.ovh.com/gb/en/storage/s3/identity-and-access-management/).
func (o S3PolicyOutput) Policy() pulumi.StringOutput {
	return o.ApplyT(func(v *S3Policy) pulumi.StringOutput { return v.Policy }).(pulumi.StringOutput)
}

// The ID of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o S3PolicyOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *S3Policy) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// The ID of a public cloud project's user.
func (o S3PolicyOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v *S3Policy) pulumi.StringOutput { return v.UserId }).(pulumi.StringOutput)
}

type S3PolicyArrayOutput struct{ *pulumi.OutputState }

func (S3PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*S3Policy)(nil)).Elem()
}

func (o S3PolicyArrayOutput) ToS3PolicyArrayOutput() S3PolicyArrayOutput {
	return o
}

func (o S3PolicyArrayOutput) ToS3PolicyArrayOutputWithContext(ctx context.Context) S3PolicyArrayOutput {
	return o
}

func (o S3PolicyArrayOutput) Index(i pulumi.IntInput) S3PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *S3Policy {
		return vs[0].([]*S3Policy)[vs[1].(int)]
	}).(S3PolicyOutput)
}

type S3PolicyMapOutput struct{ *pulumi.OutputState }

func (S3PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*S3Policy)(nil)).Elem()
}

func (o S3PolicyMapOutput) ToS3PolicyMapOutput() S3PolicyMapOutput {
	return o
}

func (o S3PolicyMapOutput) ToS3PolicyMapOutputWithContext(ctx context.Context) S3PolicyMapOutput {
	return o
}

func (o S3PolicyMapOutput) MapIndex(k pulumi.StringInput) S3PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *S3Policy {
		return vs[0].(map[string]*S3Policy)[vs[1].(string)]
	}).(S3PolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*S3PolicyInput)(nil)).Elem(), &S3Policy{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3PolicyArrayInput)(nil)).Elem(), S3PolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*S3PolicyMapInput)(nil)).Elem(), S3PolicyMap{})
	pulumi.RegisterOutputType(S3PolicyOutput{})
	pulumi.RegisterOutputType(S3PolicyArrayOutput{})
	pulumi.RegisterOutputType(S3PolicyMapOutput{})
}
