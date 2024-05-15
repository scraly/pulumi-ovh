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
    /// Creates an ACL for a kafka cluster associated with a public cloud project.
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
    ///     var kafka = Ovh.CloudProjectDatabase.GetDatabase.Invoke(new()
    ///     {
    ///         ServiceName = "XXX",
    ///         Engine = "kafka",
    ///         Id = "ZZZ",
    ///     });
    /// 
    ///     var acl = new Ovh.CloudProjectDatabase.KafkaAcl("acl", new()
    ///     {
    ///         ServiceName = kafka.Apply(getDatabaseResult =&gt; getDatabaseResult.ServiceName),
    ///         ClusterId = kafka.Apply(getDatabaseResult =&gt; getDatabaseResult.Id),
    ///         Permission = "read",
    ///         Topic = "mytopic",
    ///         Username = "johndoe",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OVHcloud Managed kafka clusters ACLs can be imported using the `service_name`, `cluster_id` and `id` of the acl, separated by "/" E.g.,
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl my_acl service_name/cluster_id/id
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl")]
    public partial class KafkaAcl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Permission to give to this username on this topic.
        /// Available permissions:
        /// </summary>
        [Output("permission")]
        public Output<string> Permission { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Topic affected by this ACL.
        /// </summary>
        [Output("topic")]
        public Output<string> Topic { get; private set; } = null!;

        /// <summary>
        /// Username affected by this ACL.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a KafkaAcl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KafkaAcl(string name, KafkaAclArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl", name, args ?? new KafkaAclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KafkaAcl(string name, Input<string> id, KafkaAclState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KafkaAcl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KafkaAcl Get(string name, Input<string> id, KafkaAclState? state = null, CustomResourceOptions? options = null)
        {
            return new KafkaAcl(name, id, state, options);
        }
    }

    public sealed class KafkaAclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Permission to give to this username on this topic.
        /// Available permissions:
        /// </summary>
        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Topic affected by this ACL.
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        /// <summary>
        /// Username affected by this ACL.
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public KafkaAclArgs()
        {
        }
        public static new KafkaAclArgs Empty => new KafkaAclArgs();
    }

    public sealed class KafkaAclState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Permission to give to this username on this topic.
        /// Available permissions:
        /// </summary>
        [Input("permission")]
        public Input<string>? Permission { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Topic affected by this ACL.
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        /// <summary>
        /// Username affected by this ACL.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public KafkaAclState()
        {
        }
        public static new KafkaAclState Empty => new KafkaAclState();
    }
}
