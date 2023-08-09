// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Ovh.Hosting.Inputs
{

    public sealed class PrivateDatabasePlanConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the resource
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// Path to the resource in API.OVH.COM
        /// 
        /// Plan order valid values can be found on OVHcloud [APIv6](https://api.ovh.com/console/#/hosting/privateDatabase/availableOrderCapacities~GET)
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public PrivateDatabasePlanConfigurationGetArgs()
        {
        }
        public static new PrivateDatabasePlanConfigurationGetArgs Empty => new PrivateDatabasePlanConfigurationGetArgs();
    }
}