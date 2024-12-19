// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetOpenSearchPatterns
    {
        /// <summary>
        /// Use this data source to get the list of pattern of a opensearch cluster associated with a public cloud project.
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
        ///     var patterns = Ovh.CloudProject.GetOpenSearchPatterns.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["patternIds"] = patterns.Apply(getOpenSearchPatternsResult =&gt; getOpenSearchPatternsResult.PatternIds),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetOpenSearchPatternsResult> InvokeAsync(GetOpenSearchPatternsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOpenSearchPatternsResult>("ovh:CloudProject/getOpenSearchPatterns:getOpenSearchPatterns", args ?? new GetOpenSearchPatternsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the list of pattern of a opensearch cluster associated with a public cloud project.
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
        ///     var patterns = Ovh.CloudProject.GetOpenSearchPatterns.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["patternIds"] = patterns.Apply(getOpenSearchPatternsResult =&gt; getOpenSearchPatternsResult.PatternIds),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetOpenSearchPatternsResult> Invoke(GetOpenSearchPatternsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOpenSearchPatternsResult>("ovh:CloudProject/getOpenSearchPatterns:getOpenSearchPatterns", args ?? new GetOpenSearchPatternsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the list of pattern of a opensearch cluster associated with a public cloud project.
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
        ///     var patterns = Ovh.CloudProject.GetOpenSearchPatterns.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["patternIds"] = patterns.Apply(getOpenSearchPatternsResult =&gt; getOpenSearchPatternsResult.PatternIds),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetOpenSearchPatternsResult> Invoke(GetOpenSearchPatternsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOpenSearchPatternsResult>("ovh:CloudProject/getOpenSearchPatterns:getOpenSearchPatterns", args ?? new GetOpenSearchPatternsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOpenSearchPatternsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetOpenSearchPatternsArgs()
        {
        }
        public static new GetOpenSearchPatternsArgs Empty => new GetOpenSearchPatternsArgs();
    }

    public sealed class GetOpenSearchPatternsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetOpenSearchPatternsInvokeArgs()
        {
        }
        public static new GetOpenSearchPatternsInvokeArgs Empty => new GetOpenSearchPatternsInvokeArgs();
    }


    [OutputType]
    public sealed class GetOpenSearchPatternsResult
    {
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of patterns ids of the opensearch cluster associated with the project.
        /// </summary>
        public readonly ImmutableArray<string> PatternIds;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private GetOpenSearchPatternsResult(
            string clusterId,

            string id,

            ImmutableArray<string> patternIds,

            string serviceName)
        {
            ClusterId = clusterId;
            Id = id;
            PatternIds = patternIds;
            ServiceName = serviceName;
        }
    }
}
