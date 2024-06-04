// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vps.Outputs
{

    [OutputType]
    public sealed class VpsPlanOptionConfiguration
    {
        /// <summary>
        /// Identifier of the resource
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// Path to the resource in api.ovh.com
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private VpsPlanOptionConfiguration(
            string label,

            string value)
        {
            Label = label;
            Value = value;
        }
    }
}
