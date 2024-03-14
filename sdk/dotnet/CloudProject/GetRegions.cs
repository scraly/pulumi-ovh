// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetRegions
    {
        /// <summary>
        /// Use this data source to get the regions of a public cloud project.
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
        ///     var regions = Ovh.CloudProject.GetRegions.Invoke(new()
        ///     {
        ///         HasServicesUps = new[]
        ///         {
        ///             "network",
        ///         },
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetRegionsResult> InvokeAsync(GetRegionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRegionsResult>("ovh:CloudProject/getRegions:getRegions", args ?? new GetRegionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the regions of a public cloud project.
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
        ///     var regions = Ovh.CloudProject.GetRegions.Invoke(new()
        ///     {
        ///         HasServicesUps = new[]
        ///         {
        ///             "network",
        ///         },
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetRegionsResult> Invoke(GetRegionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRegionsResult>("ovh:CloudProject/getRegions:getRegions", args ?? new GetRegionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRegionsArgs : global::Pulumi.InvokeArgs
    {
        [Input("hasServicesUps")]
        private List<string>? _hasServicesUps;

        /// <summary>
        /// List of services which has to be UP in regions.
        /// Example: "image", "instance", "network", "storage", "volume", "workflow", ...
        /// If left blank, returns all regions associated with the service_name.
        /// </summary>
        public List<string> HasServicesUps
        {
            get => _hasServicesUps ?? (_hasServicesUps = new List<string>());
            set => _hasServicesUps = value;
        }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetRegionsArgs()
        {
        }
        public static new GetRegionsArgs Empty => new GetRegionsArgs();
    }

    public sealed class GetRegionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("hasServicesUps")]
        private InputList<string>? _hasServicesUps;

        /// <summary>
        /// List of services which has to be UP in regions.
        /// Example: "image", "instance", "network", "storage", "volume", "workflow", ...
        /// If left blank, returns all regions associated with the service_name.
        /// </summary>
        public InputList<string> HasServicesUps
        {
            get => _hasServicesUps ?? (_hasServicesUps = new InputList<string>());
            set => _hasServicesUps = value;
        }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetRegionsInvokeArgs()
        {
        }
        public static new GetRegionsInvokeArgs Empty => new GetRegionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetRegionsResult
    {
        public readonly ImmutableArray<string> HasServicesUps;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of regions associated with the project, filtered by services UP.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string ServiceName;

        [OutputConstructor]
        private GetRegionsResult(
            ImmutableArray<string> hasServicesUps,

            string id,

            ImmutableArray<string> names,

            string serviceName)
        {
            HasServicesUps = hasServicesUps;
            Id = id;
            Names = names;
            ServiceName = serviceName;
        }
    }
}
