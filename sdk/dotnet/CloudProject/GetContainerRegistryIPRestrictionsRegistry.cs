// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetContainerRegistryIPRestrictionsRegistry
    {
        /// <summary>
        /// Use this data source to get the list of Registry IP Restrictions of a container registry associated with a public cloud project.
        /// </summary>
        public static Task<GetContainerRegistryIPRestrictionsRegistryResult> InvokeAsync(GetContainerRegistryIPRestrictionsRegistryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContainerRegistryIPRestrictionsRegistryResult>("ovh:CloudProject/getContainerRegistryIPRestrictionsRegistry:getContainerRegistryIPRestrictionsRegistry", args ?? new GetContainerRegistryIPRestrictionsRegistryArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the list of Registry IP Restrictions of a container registry associated with a public cloud project.
        /// </summary>
        public static Output<GetContainerRegistryIPRestrictionsRegistryResult> Invoke(GetContainerRegistryIPRestrictionsRegistryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerRegistryIPRestrictionsRegistryResult>("ovh:CloudProject/getContainerRegistryIPRestrictionsRegistry:getContainerRegistryIPRestrictionsRegistry", args ?? new GetContainerRegistryIPRestrictionsRegistryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContainerRegistryIPRestrictionsRegistryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the Managed Private Registry.
        /// </summary>
        [Input("registryId", required: true)]
        public string RegistryId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetContainerRegistryIPRestrictionsRegistryArgs()
        {
        }
        public static new GetContainerRegistryIPRestrictionsRegistryArgs Empty => new GetContainerRegistryIPRestrictionsRegistryArgs();
    }

    public sealed class GetContainerRegistryIPRestrictionsRegistryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the Managed Private Registry.
        /// </summary>
        [Input("registryId", required: true)]
        public Input<string> RegistryId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetContainerRegistryIPRestrictionsRegistryInvokeArgs()
        {
        }
        public static new GetContainerRegistryIPRestrictionsRegistryInvokeArgs Empty => new GetContainerRegistryIPRestrictionsRegistryInvokeArgs();
    }


    [OutputType]
    public sealed class GetContainerRegistryIPRestrictionsRegistryResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP restrictions applied on artifact manager component.
        /// </summary>
        public readonly ImmutableArray<ImmutableDictionary<string, string>> IpRestrictions;
        /// <summary>
        /// The ID of the Managed Private Registry.
        /// </summary>
        public readonly string RegistryId;
        /// <summary>
        /// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private GetContainerRegistryIPRestrictionsRegistryResult(
            string id,

            ImmutableArray<ImmutableDictionary<string, string>> ipRestrictions,

            string registryId,

            string serviceName)
        {
            Id = id;
            IpRestrictions = ipRestrictions;
            RegistryId = registryId;
            ServiceName = serviceName;
        }
    }
}
