// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetContainerRegistries
    {
        /// <summary>
        /// Use this data source to get the container registries of a public cloud project.
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
        ///     var registries = Ovh.CloudProject.GetContainerRegistries.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetContainerRegistriesResult> InvokeAsync(GetContainerRegistriesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetContainerRegistriesResult>("ovh:CloudProject/getContainerRegistries:getContainerRegistries", args ?? new GetContainerRegistriesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the container registries of a public cloud project.
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
        ///     var registries = Ovh.CloudProject.GetContainerRegistries.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetContainerRegistriesResult> Invoke(GetContainerRegistriesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerRegistriesResult>("ovh:CloudProject/getContainerRegistries:getContainerRegistries", args ?? new GetContainerRegistriesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the container registries of a public cloud project.
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
        ///     var registries = Ovh.CloudProject.GetContainerRegistries.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetContainerRegistriesResult> Invoke(GetContainerRegistriesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetContainerRegistriesResult>("ovh:CloudProject/getContainerRegistries:getContainerRegistries", args ?? new GetContainerRegistriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetContainerRegistriesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetContainerRegistriesArgs()
        {
        }
        public static new GetContainerRegistriesArgs Empty => new GetContainerRegistriesArgs();
    }

    public sealed class GetContainerRegistriesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetContainerRegistriesInvokeArgs()
        {
        }
        public static new GetContainerRegistriesInvokeArgs Empty => new GetContainerRegistriesInvokeArgs();
    }


    [OutputType]
    public sealed class GetContainerRegistriesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of container registries associated with the project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetContainerRegistriesResultResult> Results;
        public readonly string ServiceName;

        [OutputConstructor]
        private GetContainerRegistriesResult(
            string id,

            ImmutableArray<Outputs.GetContainerRegistriesResultResult> results,

            string serviceName)
        {
            Id = id;
            Results = results;
            ServiceName = serviceName;
        }
    }
}
