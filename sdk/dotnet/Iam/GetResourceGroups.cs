// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Iam
{
    public static class GetResourceGroups
    {
        /// <summary>
        /// Use this data source to list the existing IAM policies of an account.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myGroups = Ovh.Iam.GetResourceGroups.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetResourceGroupsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetResourceGroupsResult>("ovh:Iam/getResourceGroups:getResourceGroups", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to list the existing IAM policies of an account.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myGroups = Ovh.Iam.GetResourceGroups.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetResourceGroupsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetResourceGroupsResult>("ovh:Iam/getResourceGroups:getResourceGroups", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetResourceGroupsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of the resource groups IDs.
        /// </summary>
        public readonly ImmutableArray<string> ResourceGroups;

        [OutputConstructor]
        private GetResourceGroupsResult(
            string id,

            ImmutableArray<string> resourceGroups)
        {
            Id = id;
            ResourceGroups = resourceGroups;
        }
    }
}
