// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dbaas
{
    public static class GetLogsClusters
    {
        /// <summary>
        /// Use this data source to retrieve UUIDs of DBaas logs clusters.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var logstash = Ovh.Dbaas.GetLogsClusters.Invoke(new()
        ///     {
        ///         ServiceName = "ldp-xx-xxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetLogsClustersResult> InvokeAsync(GetLogsClustersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLogsClustersResult>("ovh:Dbaas/getLogsClusters:getLogsClusters", args ?? new GetLogsClustersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve UUIDs of DBaas logs clusters.
        /// 
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var logstash = Ovh.Dbaas.GetLogsClusters.Invoke(new()
        ///     {
        ///         ServiceName = "ldp-xx-xxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetLogsClustersResult> Invoke(GetLogsClustersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLogsClustersResult>("ovh:Dbaas/getLogsClusters:getLogsClusters", args ?? new GetLogsClustersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLogsClustersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The service name. It's the ID of your Logs Data Platform instance.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetLogsClustersArgs()
        {
        }
        public static new GetLogsClustersArgs Empty => new GetLogsClustersArgs();
    }

    public sealed class GetLogsClustersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The service name. It's the ID of your Logs Data Platform instance.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetLogsClustersInvokeArgs()
        {
        }
        public static new GetLogsClustersInvokeArgs Empty => new GetLogsClustersInvokeArgs();
    }


    [OutputType]
    public sealed class GetLogsClustersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ServiceName;
        public readonly string Urn;
        /// <summary>
        /// is the cluster id
        /// </summary>
        public readonly ImmutableArray<string> Uuids;

        [OutputConstructor]
        private GetLogsClustersResult(
            string id,

            string serviceName,

            string urn,

            ImmutableArray<string> uuids)
        {
            Id = id;
            ServiceName = serviceName;
            Urn = urn;
            Uuids = uuids;
        }
    }
}
