// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Iam
{
    /// <summary>
    /// Creates an IAM policy.
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
    ///     var account = Ovh.Me.GetMe.Invoke();
    /// 
    ///     var myGroup = new Ovh.Me.IdentityGroup("myGroup", new()
    ///     {
    ///         Description = "my_group created in Terraform",
    ///     });
    /// 
    ///     var manager = new Ovh.Iam.Policy("manager", new()
    ///     {
    ///         Description = "Users are allowed to use the OVH manager",
    ///         Identities = new[]
    ///         {
    ///             myGroup.GroupURN,
    ///         },
    ///         Resources = new[]
    ///         {
    ///             account.Apply(getMeResult =&gt; getMeResult.AccountURN),
    ///         },
    ///         Allows = new[]
    ///         {
    ///             "account:apiovh:me/get",
    ///             "account:apiovh:me/supportLevel/get",
    ///             "account:apiovh:me/certificates/get",
    ///             "account:apiovh:me/tag/get",
    ///             "account:apiovh:services/get",
    ///             "account:apiovh:*",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Iam/policy:Policy")]
    public partial class Policy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of actions allowed on resources by identities
        /// </summary>
        [Output("allows")]
        public Output<ImmutableArray<string>> Allows { get; private set; } = null!;

        /// <summary>
        /// Creation date of this group.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Description of the policy
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of overrides of action that must not be allowed even if they are caught by allow. Only makes sens if allow contains wildcards.
        /// </summary>
        [Output("excepts")]
        public Output<ImmutableArray<string>> Excepts { get; private set; } = null!;

        /// <summary>
        /// List of identities affected by the policy
        /// </summary>
        [Output("identities")]
        public Output<ImmutableArray<string>> Identities { get; private set; } = null!;

        /// <summary>
        /// Name of the policy, must be unique
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Owner of the policy.
        /// </summary>
        [Output("owner")]
        public Output<string> Owner { get; private set; } = null!;

        /// <summary>
        /// Indicates that the policy is a default one.
        /// </summary>
        [Output("readOnly")]
        public Output<bool> ReadOnly { get; private set; } = null!;

        /// <summary>
        /// List of resources affected by the policy
        /// </summary>
        [Output("resources")]
        public Output<ImmutableArray<string>> Resources { get; private set; } = null!;

        /// <summary>
        /// Date of the last update of this group.
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a Policy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Policy(string name, PolicyArgs args, CustomResourceOptions? options = null)
            : base("ovh:Iam/policy:Policy", name, args ?? new PolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Policy(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Iam/policy:Policy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Policy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Policy Get(string name, Input<string> id, PolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Policy(name, id, state, options);
        }
    }

    public sealed class PolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("allows")]
        private InputList<string>? _allows;

        /// <summary>
        /// List of actions allowed on resources by identities
        /// </summary>
        public InputList<string> Allows
        {
            get => _allows ?? (_allows = new InputList<string>());
            set => _allows = value;
        }

        /// <summary>
        /// Description of the policy
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("excepts")]
        private InputList<string>? _excepts;

        /// <summary>
        /// List of overrides of action that must not be allowed even if they are caught by allow. Only makes sens if allow contains wildcards.
        /// </summary>
        public InputList<string> Excepts
        {
            get => _excepts ?? (_excepts = new InputList<string>());
            set => _excepts = value;
        }

        [Input("identities", required: true)]
        private InputList<string>? _identities;

        /// <summary>
        /// List of identities affected by the policy
        /// </summary>
        public InputList<string> Identities
        {
            get => _identities ?? (_identities = new InputList<string>());
            set => _identities = value;
        }

        /// <summary>
        /// Name of the policy, must be unique
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("resources", required: true)]
        private InputList<string>? _resources;

        /// <summary>
        /// List of resources affected by the policy
        /// </summary>
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        public PolicyArgs()
        {
        }
        public static new PolicyArgs Empty => new PolicyArgs();
    }

    public sealed class PolicyState : global::Pulumi.ResourceArgs
    {
        [Input("allows")]
        private InputList<string>? _allows;

        /// <summary>
        /// List of actions allowed on resources by identities
        /// </summary>
        public InputList<string> Allows
        {
            get => _allows ?? (_allows = new InputList<string>());
            set => _allows = value;
        }

        /// <summary>
        /// Creation date of this group.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Description of the policy
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("excepts")]
        private InputList<string>? _excepts;

        /// <summary>
        /// List of overrides of action that must not be allowed even if they are caught by allow. Only makes sens if allow contains wildcards.
        /// </summary>
        public InputList<string> Excepts
        {
            get => _excepts ?? (_excepts = new InputList<string>());
            set => _excepts = value;
        }

        [Input("identities")]
        private InputList<string>? _identities;

        /// <summary>
        /// List of identities affected by the policy
        /// </summary>
        public InputList<string> Identities
        {
            get => _identities ?? (_identities = new InputList<string>());
            set => _identities = value;
        }

        /// <summary>
        /// Name of the policy, must be unique
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Owner of the policy.
        /// </summary>
        [Input("owner")]
        public Input<string>? Owner { get; set; }

        /// <summary>
        /// Indicates that the policy is a default one.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        [Input("resources")]
        private InputList<string>? _resources;

        /// <summary>
        /// List of resources affected by the policy
        /// </summary>
        public InputList<string> Resources
        {
            get => _resources ?? (_resources = new InputList<string>());
            set => _resources = value;
        }

        /// <summary>
        /// Date of the last update of this group.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public PolicyState()
        {
        }
        public static new PolicyState Empty => new PolicyState();
    }
}
