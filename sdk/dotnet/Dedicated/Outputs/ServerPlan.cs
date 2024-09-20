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
    public sealed class ServerPlan
    {
        public readonly ImmutableArray<Outputs.ServerPlanConfiguration> Configurations;
        /// <summary>
        /// Duration selected for the purchase of the product
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// Cart item to be linked
        /// </summary>
        public readonly double? ItemId;
        /// <summary>
        /// Identifier of the option offer
        /// </summary>
        public readonly string PlanCode;
        /// <summary>
        /// Pricing mode selected for the purchase of the product
        /// </summary>
        public readonly string PricingMode;
        /// <summary>
        /// Quantity of product desired
        /// </summary>
        public readonly double? Quantity;

        [OutputConstructor]
        private ServerPlan(
            ImmutableArray<Outputs.ServerPlanConfiguration> configurations,

            string duration,

            double? itemId,

            string planCode,

            string pricingMode,

            double? quantity)
        {
            Configurations = configurations;
            Duration = duration;
            ItemId = itemId;
            PlanCode = planCode;
            PricingMode = pricingMode;
            Quantity = quantity;
        }
    }
}