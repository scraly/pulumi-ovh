// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Outputs
{

    [OutputType]
    public sealed class GetInstallationTemplateLicenseResult
    {
        public readonly ImmutableArray<Outputs.GetInstallationTemplateLicenseOResult> Os;
        public readonly ImmutableArray<Outputs.GetInstallationTemplateLicenseUsageResult> Usages;

        [OutputConstructor]
        private GetInstallationTemplateLicenseResult(
            ImmutableArray<Outputs.GetInstallationTemplateLicenseOResult> os,

            ImmutableArray<Outputs.GetInstallationTemplateLicenseUsageResult> usages)
        {
            Os = os;
            Usages = usages;
        }
    }
}