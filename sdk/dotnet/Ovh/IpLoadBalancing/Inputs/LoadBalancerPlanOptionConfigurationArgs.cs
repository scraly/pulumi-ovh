// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Scraly.PulumiPackage.Ovh.IpLoadBalancing.Inputs
{

    public sealed class LoadBalancerPlanOptionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier of the resource
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// Path to the resource in API.OVH.COM
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public LoadBalancerPlanOptionConfigurationArgs()
        {
        }
        public static new LoadBalancerPlanOptionConfigurationArgs Empty => new LoadBalancerPlanOptionConfigurationArgs();
    }
}