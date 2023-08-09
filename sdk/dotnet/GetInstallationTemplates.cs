// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Ovh
{
    public static class GetInstallationTemplates
    {
        /// <summary>
        /// Use this data source to get the list of installation templates available for dedicated servers.
        /// 
        /// ## Example Usage
        /// 
        /// ```hcl
        /// data "ovh_dedicated_installation_templates" "templates" {}
        /// ```
        /// </summary>
        public static Task<GetInstallationTemplatesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstallationTemplatesResult>("ovh:index/getInstallationTemplates:getInstallationTemplates", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetInstallationTemplatesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of installation templates IDs available for dedicated servers.
        /// </summary>
        public readonly ImmutableArray<string> Results;

        [OutputConstructor]
        private GetInstallationTemplatesResult(
            string id,

            ImmutableArray<string> results)
        {
            Id = id;
            Results = results;
        }
    }
}