// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase
{
    /// <summary>
    /// ## Import
    /// 
    /// OVHcloud Managed Redis clusters users can be imported using the `service_name`, `cluster_id` and `id` of the user, separated by "/" E.g., bash
    /// 
    /// ```sh
    ///  $ pulumi import ovh:CloudProjectDatabase/redisUser:RedisUser my_user service_name/cluster_id/id
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProjectDatabase/redisUser:RedisUser")]
    public partial class RedisUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Categories of the user.
        /// </summary>
        [Output("categories")]
        public Output<ImmutableArray<string>> Categories { get; private set; } = null!;

        /// <summary>
        /// Channels of the user.
        /// </summary>
        [Output("channels")]
        public Output<ImmutableArray<string>> Channels { get; private set; } = null!;

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Commands of the user.
        /// </summary>
        [Output("commands")]
        public Output<ImmutableArray<string>> Commands { get; private set; } = null!;

        /// <summary>
        /// Date of the creation of the user.
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Keys of the user.
        /// </summary>
        [Output("keys")]
        public Output<ImmutableArray<string>> Keys { get; private set; } = null!;

        /// <summary>
        /// Name of the user. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// (Sensitive) Password of the user.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Arbitrary string to change to trigger a password update
        /// </summary>
        [Output("passwordReset")]
        public Output<string?> PasswordReset { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Current status of the user.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a RedisUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RedisUser(string name, RedisUserArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/redisUser:RedisUser", name, args ?? new RedisUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RedisUser(string name, Input<string> id, RedisUserState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/redisUser:RedisUser", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RedisUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RedisUser Get(string name, Input<string> id, RedisUserState? state = null, CustomResourceOptions? options = null)
        {
            return new RedisUser(name, id, state, options);
        }
    }

    public sealed class RedisUserArgs : global::Pulumi.ResourceArgs
    {
        [Input("categories")]
        private InputList<string>? _categories;

        /// <summary>
        /// Categories of the user.
        /// </summary>
        public InputList<string> Categories
        {
            get => _categories ?? (_categories = new InputList<string>());
            set => _categories = value;
        }

        [Input("channels")]
        private InputList<string>? _channels;

        /// <summary>
        /// Channels of the user.
        /// </summary>
        public InputList<string> Channels
        {
            get => _channels ?? (_channels = new InputList<string>());
            set => _channels = value;
        }

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// Commands of the user.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        [Input("keys")]
        private InputList<string>? _keys;

        /// <summary>
        /// Keys of the user.
        /// </summary>
        public InputList<string> Keys
        {
            get => _keys ?? (_keys = new InputList<string>());
            set => _keys = value;
        }

        /// <summary>
        /// Name of the user. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Arbitrary string to change to trigger a password update
        /// </summary>
        [Input("passwordReset")]
        public Input<string>? PasswordReset { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public RedisUserArgs()
        {
        }
        public static new RedisUserArgs Empty => new RedisUserArgs();
    }

    public sealed class RedisUserState : global::Pulumi.ResourceArgs
    {
        [Input("categories")]
        private InputList<string>? _categories;

        /// <summary>
        /// Categories of the user.
        /// </summary>
        public InputList<string> Categories
        {
            get => _categories ?? (_categories = new InputList<string>());
            set => _categories = value;
        }

        [Input("channels")]
        private InputList<string>? _channels;

        /// <summary>
        /// Channels of the user.
        /// </summary>
        public InputList<string> Channels
        {
            get => _channels ?? (_channels = new InputList<string>());
            set => _channels = value;
        }

        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// Commands of the user.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        /// <summary>
        /// Date of the creation of the user.
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("keys")]
        private InputList<string>? _keys;

        /// <summary>
        /// Keys of the user.
        /// </summary>
        public InputList<string> Keys
        {
            get => _keys ?? (_keys = new InputList<string>());
            set => _keys = value;
        }

        /// <summary>
        /// Name of the user. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// (Sensitive) Password of the user.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Arbitrary string to change to trigger a password update
        /// </summary>
        [Input("passwordReset")]
        public Input<string>? PasswordReset { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Current status of the user.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public RedisUserState()
        {
        }
        public static new RedisUserState Empty => new RedisUserState();
    }
}
