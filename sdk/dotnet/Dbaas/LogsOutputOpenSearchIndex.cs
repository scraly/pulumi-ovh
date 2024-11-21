// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dbaas
{
    /// <summary>
    /// Creates a DBaaS Logs Opensearch output index.
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
    ///     var index = new Ovh.Dbaas.LogsOutputOpenSearchIndex("index", new()
    ///     {
    ///         Description = "my opensearch index",
    ///         ServiceName = "....",
    ///         Suffix = "index",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Dbaas/logsOutputOpenSearchIndex:LogsOutputOpenSearchIndex")]
    public partial class LogsOutputOpenSearchIndex : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If set, notify when size is near 80, 90 or 100 % of its maximum capacity
        /// </summary>
        [Output("alertNotifyEnabled")]
        public Output<bool> AlertNotifyEnabled { get; private set; } = null!;

        /// <summary>
        /// Index creation
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Current index size (in bytes)
        /// </summary>
        [Output("currentSize")]
        public Output<int> CurrentSize { get; private set; } = null!;

        /// <summary>
        /// Index description
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Index ID
        /// </summary>
        [Output("indexId")]
        public Output<string> IndexId { get; private set; } = null!;

        /// <summary>
        /// Indicates if you are allowed to edit entry
        /// </summary>
        [Output("isEditable")]
        public Output<bool> IsEditable { get; private set; } = null!;

        /// <summary>
        /// Maximum index size (in bytes)
        /// </summary>
        [Output("maxSize")]
        public Output<int> MaxSize { get; private set; } = null!;

        /// <summary>
        /// Index name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Number of shards
        /// </summary>
        [Output("nbShard")]
        public Output<int> NbShard { get; private set; } = null!;

        /// <summary>
        /// The service name
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Index suffix
        /// </summary>
        [Output("suffix")]
        public Output<string> Suffix { get; private set; } = null!;

        /// <summary>
        /// Index last update
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a LogsOutputOpenSearchIndex resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LogsOutputOpenSearchIndex(string name, LogsOutputOpenSearchIndexArgs args, CustomResourceOptions? options = null)
            : base("ovh:Dbaas/logsOutputOpenSearchIndex:LogsOutputOpenSearchIndex", name, args ?? new LogsOutputOpenSearchIndexArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LogsOutputOpenSearchIndex(string name, Input<string> id, LogsOutputOpenSearchIndexState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Dbaas/logsOutputOpenSearchIndex:LogsOutputOpenSearchIndex", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LogsOutputOpenSearchIndex resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LogsOutputOpenSearchIndex Get(string name, Input<string> id, LogsOutputOpenSearchIndexState? state = null, CustomResourceOptions? options = null)
        {
            return new LogsOutputOpenSearchIndex(name, id, state, options);
        }
    }

    public sealed class LogsOutputOpenSearchIndexArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Index description
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// Number of shards
        /// </summary>
        [Input("nbShard", required: true)]
        public Input<int> NbShard { get; set; } = null!;

        /// <summary>
        /// The service name
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Index suffix
        /// </summary>
        [Input("suffix", required: true)]
        public Input<string> Suffix { get; set; } = null!;

        public LogsOutputOpenSearchIndexArgs()
        {
        }
        public static new LogsOutputOpenSearchIndexArgs Empty => new LogsOutputOpenSearchIndexArgs();
    }

    public sealed class LogsOutputOpenSearchIndexState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If set, notify when size is near 80, 90 or 100 % of its maximum capacity
        /// </summary>
        [Input("alertNotifyEnabled")]
        public Input<bool>? AlertNotifyEnabled { get; set; }

        /// <summary>
        /// Index creation
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Current index size (in bytes)
        /// </summary>
        [Input("currentSize")]
        public Input<int>? CurrentSize { get; set; }

        /// <summary>
        /// Index description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Index ID
        /// </summary>
        [Input("indexId")]
        public Input<string>? IndexId { get; set; }

        /// <summary>
        /// Indicates if you are allowed to edit entry
        /// </summary>
        [Input("isEditable")]
        public Input<bool>? IsEditable { get; set; }

        /// <summary>
        /// Maximum index size (in bytes)
        /// </summary>
        [Input("maxSize")]
        public Input<int>? MaxSize { get; set; }

        /// <summary>
        /// Index name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Number of shards
        /// </summary>
        [Input("nbShard")]
        public Input<int>? NbShard { get; set; }

        /// <summary>
        /// The service name
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Index suffix
        /// </summary>
        [Input("suffix")]
        public Input<string>? Suffix { get; set; }

        /// <summary>
        /// Index last update
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public LogsOutputOpenSearchIndexState()
        {
        }
        public static new LogsOutputOpenSearchIndexState Empty => new LogsOutputOpenSearchIndexState();
    }
}