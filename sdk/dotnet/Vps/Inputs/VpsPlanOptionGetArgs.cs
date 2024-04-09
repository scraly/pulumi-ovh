// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vps.Inputs
{

    public sealed class VpsPlanOptionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("configurations")]
        private InputList<Inputs.VpsPlanOptionConfigurationGetArgs>? _configurations;

        /// <summary>
        /// Representation of a configuration item for personalizing product
        /// </summary>
        public InputList<Inputs.VpsPlanOptionConfigurationGetArgs> Configurations
        {
            get => _configurations ?? (_configurations = new InputList<Inputs.VpsPlanOptionConfigurationGetArgs>());
            set => _configurations = value;
        }

        /// <summary>
        /// duration
        /// </summary>
        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        /// <summary>
        /// Cart item to be linked
        /// </summary>
        [Input("itemId", required: true)]
        public Input<double> ItemId { get; set; } = null!;

        /// <summary>
        /// Plan code
        /// </summary>
        [Input("planCode", required: true)]
        public Input<string> PlanCode { get; set; } = null!;

        /// <summary>
        /// Pricing model identifier
        /// </summary>
        [Input("pricingMode", required: true)]
        public Input<string> PricingMode { get; set; } = null!;

        /// <summary>
        /// Quantity of product desired
        /// </summary>
        [Input("quantity", required: true)]
        public Input<double> Quantity { get; set; } = null!;

        public VpsPlanOptionGetArgs()
        {
        }
        public static new VpsPlanOptionGetArgs Empty => new VpsPlanOptionGetArgs();
    }
}
