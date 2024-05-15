// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me
{
    public static class GetInstallationTemplate
    {
        /// <summary>
        /// Use this data source to get a custom installation template available for dedicated servers.
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
        ///     var mytemplate = Ovh.Me.GetInstallationTemplate.Invoke(new()
        ///     {
        ///         TemplateName = "mytemplate",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetInstallationTemplateResult> InvokeAsync(GetInstallationTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstallationTemplateResult>("ovh:Me/getInstallationTemplate:getInstallationTemplate", args ?? new GetInstallationTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a custom installation template available for dedicated servers.
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
        ///     var mytemplate = Ovh.Me.GetInstallationTemplate.Invoke(new()
        ///     {
        ///         TemplateName = "mytemplate",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetInstallationTemplateResult> Invoke(GetInstallationTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetInstallationTemplateResult>("ovh:Me/getInstallationTemplate:getInstallationTemplate", args ?? new GetInstallationTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetInstallationTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Template name.
        /// </summary>
        [Input("templateName", required: true)]
        public string TemplateName { get; set; } = null!;

        public GetInstallationTemplateArgs()
        {
        }
        public static new GetInstallationTemplateArgs Empty => new GetInstallationTemplateArgs();
    }

    public sealed class GetInstallationTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Template name.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        public GetInstallationTemplateInvokeArgs()
        {
        }
        public static new GetInstallationTemplateInvokeArgs Empty => new GetInstallationTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetInstallationTemplateResult
    {
        /// <summary>
        /// Template bit format (32 or 64).
        /// </summary>
        public readonly int BitFormat;
        /// <summary>
        /// Category of this template (informative only).
        /// </summary>
        public readonly string Category;
        public readonly ImmutableArray<Outputs.GetInstallationTemplateCustomizationResult> Customizations;
        /// <summary>
        /// Information about this template.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Distribution this template is based on.
        /// </summary>
        public readonly string Distribution;
        /// <summary>
        /// End of install date of the template.
        /// </summary>
        public readonly string EndOfInstall;
        /// <summary>
        /// Template family type (bsd,linux,solaris,windows).
        /// </summary>
        public readonly string Family;
        /// <summary>
        /// Filesystems available.
        /// </summary>
        public readonly ImmutableArray<string> Filesystems;
        /// <summary>
        /// Distribution supports hardware raid configuration through the OVHcloud API.
        /// </summary>
        public readonly bool HardRaidConfiguration;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Represents the questions of the expected answers in the userMetadata field.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetInstallationTemplateInputResult> Inputs;
        /// <summary>
        /// Whether this template supports LVM.
        /// </summary>
        public readonly bool LvmReady;
        /// <summary>
        /// Partitioning customization is not available for this OS template.
        /// </summary>
        public readonly bool NoPartitioning;
        public readonly ImmutableArray<Outputs.GetInstallationTemplatePartitionSchemeResult> PartitionSchemes;
        /// <summary>
        /// Template supports RAID0 and RAID1 on 2 disks.
        /// </summary>
        public readonly bool SoftRaidOnlyMirroring;
        /// <summary>
        /// Subfamily of the template.
        /// </summary>
        public readonly string Subfamily;
        public readonly string TemplateName;

        [OutputConstructor]
        private GetInstallationTemplateResult(
            int bitFormat,

            string category,

            ImmutableArray<Outputs.GetInstallationTemplateCustomizationResult> customizations,

            string description,

            string distribution,

            string endOfInstall,

            string family,

            ImmutableArray<string> filesystems,

            bool hardRaidConfiguration,

            string id,

            ImmutableArray<Outputs.GetInstallationTemplateInputResult> inputs,

            bool lvmReady,

            bool noPartitioning,

            ImmutableArray<Outputs.GetInstallationTemplatePartitionSchemeResult> partitionSchemes,

            bool softRaidOnlyMirroring,

            string subfamily,

            string templateName)
        {
            BitFormat = bitFormat;
            Category = category;
            Customizations = customizations;
            Description = description;
            Distribution = distribution;
            EndOfInstall = endOfInstall;
            Family = family;
            Filesystems = filesystems;
            HardRaidConfiguration = hardRaidConfiguration;
            Id = id;
            Inputs = inputs;
            LvmReady = lvmReady;
            NoPartitioning = noPartitioning;
            PartitionSchemes = partitionSchemes;
            SoftRaidOnlyMirroring = softRaidOnlyMirroring;
            Subfamily = subfamily;
            TemplateName = templateName;
        }
    }
}
