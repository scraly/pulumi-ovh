// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dbaas
{
    public static class GetLogsCluster
    {
        /// <summary>
        /// Use this data source to retrieve informations about a DBaas logs cluster tenant.
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
        ///     var logstash = Ovh.Dbaas.GetLogsCluster.Invoke(new()
        ///     {
        ///         ClusterId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        ///         ServiceName = "ldp-xx-xxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetLogsClusterResult> InvokeAsync(GetLogsClusterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLogsClusterResult>("ovh:Dbaas/getLogsCluster:getLogsCluster", args ?? new GetLogsClusterArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve informations about a DBaas logs cluster tenant.
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
        ///     var logstash = Ovh.Dbaas.GetLogsCluster.Invoke(new()
        ///     {
        ///         ClusterId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
        ///         ServiceName = "ldp-xx-xxxxx",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetLogsClusterResult> Invoke(GetLogsClusterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLogsClusterResult>("ovh:Dbaas/getLogsCluster:getLogsCluster", args ?? new GetLogsClusterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLogsClusterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID. If not provided, the default cluster_id is returned
        /// </summary>
        [Input("clusterId")]
        public string? ClusterId { get; set; }

        /// <summary>
        /// The service name. It's the ID of your Logs Data Platform instance.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetLogsClusterArgs()
        {
        }
        public static new GetLogsClusterArgs Empty => new GetLogsClusterArgs();
    }

    public sealed class GetLogsClusterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID. If not provided, the default cluster_id is returned
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The service name. It's the ID of your Logs Data Platform instance.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetLogsClusterInvokeArgs()
        {
        }
        public static new GetLogsClusterInvokeArgs Empty => new GetLogsClusterInvokeArgs();
    }


    [OutputType]
    public sealed class GetLogsClusterResult
    {
        /// <summary>
        /// is allowed networks for ARCHIVE flow type
        /// </summary>
        public readonly ImmutableArray<string> ArchiveAllowedNetworks;
        public readonly string? ClusterId;
        /// <summary>
        /// is type of cluster (DEDICATED, PRO or TRIAL)
        /// </summary>
        public readonly string ClusterType;
        /// <summary>
        /// is PEM for dedicated inputs
        /// </summary>
        public readonly string DedicatedInputPem;
        /// <summary>
        /// is allowed networks for DIRECT_INPUT flow type
        /// </summary>
        public readonly ImmutableArray<string> DirectInputAllowedNetworks;
        /// <summary>
        /// is PEM for direct inputs
        /// </summary>
        public readonly string DirectInputPem;
        /// <summary>
        /// is cluster hostname hosting the tenant
        /// </summary>
        public readonly string Hostname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// is true if all content generated by given service will be placed on this cluster
        /// </summary>
        public readonly bool IsDefault;
        /// <summary>
        /// is true if given service can perform advanced operations on cluster
        /// </summary>
        public readonly bool IsUnlocked;
        /// <summary>
        /// is allowed networks for QUERY flow type
        /// </summary>
        public readonly ImmutableArray<string> QueryAllowedNetworks;
        /// <summary>
        /// is datacenter localization
        /// </summary>
        public readonly string Region;
        public readonly string ServiceName;
        /// <summary>
        /// is the URN of the DBaas logs instance
        /// </summary>
        public readonly string Urn;

        [OutputConstructor]
        private GetLogsClusterResult(
            ImmutableArray<string> archiveAllowedNetworks,

            string? clusterId,

            string clusterType,

            string dedicatedInputPem,

            ImmutableArray<string> directInputAllowedNetworks,

            string directInputPem,

            string hostname,

            string id,

            bool isDefault,

            bool isUnlocked,

            ImmutableArray<string> queryAllowedNetworks,

            string region,

            string serviceName,

            string urn)
        {
            ArchiveAllowedNetworks = archiveAllowedNetworks;
            ClusterId = clusterId;
            ClusterType = clusterType;
            DedicatedInputPem = dedicatedInputPem;
            DirectInputAllowedNetworks = directInputAllowedNetworks;
            DirectInputPem = directInputPem;
            Hostname = hostname;
            Id = id;
            IsDefault = isDefault;
            IsUnlocked = isUnlocked;
            QueryAllowedNetworks = queryAllowedNetworks;
            Region = region;
            ServiceName = serviceName;
            Urn = urn;
        }
    }
}
