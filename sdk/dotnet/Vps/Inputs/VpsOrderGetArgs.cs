// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vps.Inputs
{

    public sealed class VpsOrderGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("date")]
        public Input<string>? Date { get; set; }

        [Input("details")]
        private InputList<Inputs.VpsOrderDetailGetArgs>? _details;
        public InputList<Inputs.VpsOrderDetailGetArgs> Details
        {
            get => _details ?? (_details = new InputList<Inputs.VpsOrderDetailGetArgs>());
            set => _details = value;
        }

        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        [Input("orderId")]
        public Input<double>? OrderId { get; set; }

        public VpsOrderGetArgs()
        {
        }
        public static new VpsOrderGetArgs Empty => new VpsOrderGetArgs();
    }
}