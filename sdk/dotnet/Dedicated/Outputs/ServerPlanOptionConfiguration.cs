// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Outputs
{

    [OutputType]
    public sealed class ServerPlanOptionConfiguration
    {
        /// <summary>
        /// Label for your configuration item
        /// </summary>
        public readonly string Label;
        /// <summary>
        /// Value or resource URL on API.OVH.COM of your configuration item
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private ServerPlanOptionConfiguration(
            string label,

            string value)
        {
            Label = label;
            Value = value;
        }
    }
}
