// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Ovh.CloudProjectDatabase
{
    public static class GetDatabases
    {
        /// <summary>
        /// Use this data source to get the list of managed databases of a public cloud project.
        /// 
        /// ## Example Usage
        /// 
        /// To get the list of database clusters service for a given engine:
        /// 
        /// ```hcl
        /// data "ovh_cloud_project_databases" "dbs" {
        ///   service_name  = "XXXXXX"
        ///   engine        = "YYYY"
        /// }
        /// 
        /// output "cluster_ids" {
        ///   value = data.ovh_cloud_project_databases.dbs.cluster_ids
        /// }
        /// ```
        /// </summary>
        public static Task<GetDatabasesResult> InvokeAsync(GetDatabasesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDatabasesResult>("ovh:CloudProjectDatabase/getDatabases:getDatabases", args ?? new GetDatabasesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the list of managed databases of a public cloud project.
        /// 
        /// ## Example Usage
        /// 
        /// To get the list of database clusters service for a given engine:
        /// 
        /// ```hcl
        /// data "ovh_cloud_project_databases" "dbs" {
        ///   service_name  = "XXXXXX"
        ///   engine        = "YYYY"
        /// }
        /// 
        /// output "cluster_ids" {
        ///   value = data.ovh_cloud_project_databases.dbs.cluster_ids
        /// }
        /// ```
        /// </summary>
        public static Output<GetDatabasesResult> Invoke(GetDatabasesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDatabasesResult>("ovh:CloudProjectDatabase/getDatabases:getDatabases", args ?? new GetDatabasesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabasesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database engine you want to list. To get a full list of available engine visit:
        /// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine", required: true)]
        public string Engine { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetDatabasesArgs()
        {
        }
        public static new GetDatabasesArgs Empty => new GetDatabasesArgs();
    }

    public sealed class GetDatabasesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The database engine you want to list. To get a full list of available engine visit:
        /// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetDatabasesInvokeArgs()
        {
        }
        public static new GetDatabasesInvokeArgs Empty => new GetDatabasesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabasesResult
    {
        /// <summary>
        /// The list of managed databases ids of the project.
        /// </summary>
        public readonly ImmutableArray<string> ClusterIds;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Engine;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private GetDatabasesResult(
            ImmutableArray<string> clusterIds,

            string engine,

            string id,

            string serviceName)
        {
            ClusterIds = clusterIds;
            Engine = engine;
            Id = id;
            ServiceName = serviceName;
        }
    }
}