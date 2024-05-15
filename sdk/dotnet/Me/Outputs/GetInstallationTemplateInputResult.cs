// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me.Outputs
{

    [OutputType]
    public sealed class GetInstallationTemplateInputResult
    {
        public readonly string Default;
        /// <summary>
        /// Information about this template.
        /// </summary>
        public readonly string Description;
        public readonly ImmutableArray<string> Enums;
        public readonly bool Mandatory;
        /// <summary>
        /// Hardware RAID name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Partition type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetInstallationTemplateInputResult(
            string @default,

            string description,

            ImmutableArray<string> enums,

            bool mandatory,

            string name,

            string type)
        {
            Default = @default;
            Description = description;
            Enums = enums;
            Mandatory = mandatory;
            Name = name;
            Type = type;
        }
    }
}
