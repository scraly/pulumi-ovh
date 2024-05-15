// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a list of OVHcloud Managed Kubernetes nodes.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			nodesKubeNodes, err := CloudProject.GetKubeNodes(ctx, &cloudproject.GetKubeNodesArgs{
//				ServiceName: "XXXXXX",
//				KubeId:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("nodes", nodesKubeNodes)
//			return nil
//		})
//	}
//
// ```
func GetKubeNodes(ctx *pulumi.Context, args *GetKubeNodesArgs, opts ...pulumi.InvokeOption) (*GetKubeNodesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetKubeNodesResult
	err := ctx.Invoke("ovh:CloudProject/getKubeNodes:getKubeNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubeNodes.
type GetKubeNodesArgs struct {
	// The ID of the managed kubernetes cluster.
	KubeId string `pulumi:"kubeId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getKubeNodes.
type GetKubeNodesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	KubeId string `pulumi:"kubeId"`
	// List of all nodes composing the kubernetes cluster
	Nodes []GetKubeNodesNode `pulumi:"nodes"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func GetKubeNodesOutput(ctx *pulumi.Context, args GetKubeNodesOutputArgs, opts ...pulumi.InvokeOption) GetKubeNodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubeNodesResult, error) {
			args := v.(GetKubeNodesArgs)
			r, err := GetKubeNodes(ctx, &args, opts...)
			var s GetKubeNodesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKubeNodesResultOutput)
}

// A collection of arguments for invoking getKubeNodes.
type GetKubeNodesOutputArgs struct {
	// The ID of the managed kubernetes cluster.
	KubeId pulumi.StringInput `pulumi:"kubeId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetKubeNodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubeNodesArgs)(nil)).Elem()
}

// A collection of values returned by getKubeNodes.
type GetKubeNodesResultOutput struct{ *pulumi.OutputState }

func (GetKubeNodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubeNodesResult)(nil)).Elem()
}

func (o GetKubeNodesResultOutput) ToGetKubeNodesResultOutput() GetKubeNodesResultOutput {
	return o
}

func (o GetKubeNodesResultOutput) ToGetKubeNodesResultOutputWithContext(ctx context.Context) GetKubeNodesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetKubeNodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeNodesResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetKubeNodesResultOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeNodesResult) string { return v.KubeId }).(pulumi.StringOutput)
}

// List of all nodes composing the kubernetes cluster
func (o GetKubeNodesResultOutput) Nodes() GetKubeNodesNodeArrayOutput {
	return o.ApplyT(func(v GetKubeNodesResult) []GetKubeNodesNode { return v.Nodes }).(GetKubeNodesNodeArrayOutput)
}

// See Argument Reference above.
func (o GetKubeNodesResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeNodesResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubeNodesResultOutput{})
}
