// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vps.Inputs
{

    public sealed class VpsPlanConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the resource
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// Path to the resource in api.ovh.com
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public VpsPlanConfigurationArgs()
        {
        }
        public static new VpsPlanConfigurationArgs Empty => new VpsPlanConfigurationArgs();
    }
}