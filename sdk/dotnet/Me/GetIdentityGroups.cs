// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me
{
    public static class GetIdentityGroups
    {
        /// <summary>
        /// Use this data source to retrieve the list of the account's identity groups
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
        ///     var groups = Ovh.Me.GetIdentityGroups.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIdentityGroupsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdentityGroupsResult>("ovh:Me/getIdentityGroups:getIdentityGroups", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve the list of the account's identity groups
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
        ///     var groups = Ovh.Me.GetIdentityGroups.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIdentityGroupsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityGroupsResult>("ovh:Me/getIdentityGroups:getIdentityGroups", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetIdentityGroupsResult
    {
        /// <summary>
        /// The list of the group names of all the identity groups.
        /// </summary>
        public readonly ImmutableArray<string> Groups;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetIdentityGroupsResult(
            ImmutableArray<string> groups,

            string id)
        {
            Groups = groups;
            Id = id;
        }
    }
}
