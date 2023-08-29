// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vrack
{
    /// <summary>
    /// Attach an IP block to a VRack.
    /// 
    /// ## Example Usage
    /// </summary>
    [OvhResourceType("ovh:Vrack/ipAddress:IpAddress")]
    public partial class IpAddress : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Your IP block.
        /// </summary>
        [Output("block")]
        public Output<string> Block { get; private set; } = null!;

        /// <summary>
        /// Your gateway
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// Your IP block
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The internal name of your vrack
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Where you want your block announced on the network
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a IpAddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpAddress(string name, IpAddressArgs args, CustomResourceOptions? options = null)
            : base("ovh:Vrack/ipAddress:IpAddress", name, args ?? new IpAddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpAddress(string name, Input<string> id, IpAddressState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Vrack/ipAddress:IpAddress", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/scraly/pulumi-ovh",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IpAddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpAddress Get(string name, Input<string> id, IpAddressState? state = null, CustomResourceOptions? options = null)
        {
            return new IpAddress(name, id, state, options);
        }
    }

    public sealed class IpAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Your IP block.
        /// </summary>
        [Input("block", required: true)]
        public Input<string> Block { get; set; } = null!;

        /// <summary>
        /// The internal name of your vrack
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public IpAddressArgs()
        {
        }
        public static new IpAddressArgs Empty => new IpAddressArgs();
    }

    public sealed class IpAddressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Your IP block.
        /// </summary>
        [Input("block")]
        public Input<string>? Block { get; set; }

        /// <summary>
        /// Your gateway
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Your IP block
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The internal name of your vrack
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Where you want your block announced on the network
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public IpAddressState()
        {
        }
        public static new IpAddressState Empty => new IpAddressState();
    }
}