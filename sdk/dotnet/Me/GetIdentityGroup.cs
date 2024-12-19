// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me
{
    public static class GetIdentityGroup
    {
        /// <summary>
        /// Use this data source to retrieve information about an identity group.
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
        ///     var myGroup = Ovh.Me.GetIdentityGroup.Invoke(new()
        ///     {
        ///         Name = "my_group_name",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIdentityGroupResult> InvokeAsync(GetIdentityGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIdentityGroupResult>("ovh:Me/getIdentityGroup:getIdentityGroup", args ?? new GetIdentityGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an identity group.
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
        ///     var myGroup = Ovh.Me.GetIdentityGroup.Invoke(new()
        ///     {
        ///         Name = "my_group_name",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIdentityGroupResult> Invoke(GetIdentityGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityGroupResult>("ovh:Me/getIdentityGroup:getIdentityGroup", args ?? new GetIdentityGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an identity group.
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
        ///     var myGroup = Ovh.Me.GetIdentityGroup.Invoke(new()
        ///     {
        ///         Name = "my_group_name",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIdentityGroupResult> Invoke(GetIdentityGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetIdentityGroupResult>("ovh:Me/getIdentityGroup:getIdentityGroup", args ?? new GetIdentityGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIdentityGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Group name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetIdentityGroupArgs()
        {
        }
        public static new GetIdentityGroupArgs Empty => new GetIdentityGroupArgs();
    }

    public sealed class GetIdentityGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Group name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetIdentityGroupInvokeArgs()
        {
        }
        public static new GetIdentityGroupInvokeArgs Empty => new GetIdentityGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetIdentityGroupResult
    {
        /// <summary>
        /// Identity URN of the group.
        /// </summary>
        public readonly string GroupURN;
        /// <summary>
        /// Creation date of this group.
        /// </summary>
        public readonly string Creation;
        /// <summary>
        /// Is the group a default and immutable one.
        /// </summary>
        public readonly bool DefaultGroup;
        /// <summary>
        /// Group description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Date of the last update of this group.
        /// </summary>
        public readonly string LastUpdate;
        public readonly string Name;
        /// <summary>
        /// Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
        /// </summary>
        public readonly string Role;

        [OutputConstructor]
        private GetIdentityGroupResult(
            string GroupURN,

            string creation,

            bool defaultGroup,

            string description,

            string id,

            string lastUpdate,

            string name,

            string role)
        {
            this.GroupURN = GroupURN;
            Creation = creation;
            DefaultGroup = defaultGroup;
            Description = description;
            Id = id;
            LastUpdate = lastUpdate;
            Name = name;
            Role = role;
        }
    }
}
