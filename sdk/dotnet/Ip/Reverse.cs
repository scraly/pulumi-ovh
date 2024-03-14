// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Ip
{
    /// <summary>
    /// Provides a OVHcloud IP reverse.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // Set the reverse of an IP
    ///     var test = new Ovh.Ip.Reverse("test", new()
    ///     {
    ///         Ip = "192.0.2.0/24",
    ///         ReverseIp = "192.0.2.1",
    ///         ReverseValue = "example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// The resource can be imported using the `ip`, `ip_reverse` of the address, separated by "|" E.g.,
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import ovh:Ip/reverse:Reverse my_reverse '2001:0db8:c0ff:ee::/64|2001:0db8:c0ff:ee::42'
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Ip/reverse:Reverse")]
    public partial class Reverse : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The IP to set the reverse of
        /// </summary>
        [Output("ReverseIp")]
        public Output<string> ReverseIp { get; private set; } = null!;

        /// <summary>
        /// The value of the reverse
        /// </summary>
        [Output("ReverseValue")]
        public Output<string> ReverseValue { get; private set; } = null!;

        /// <summary>
        /// The IP block to which the IP belongs
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;


        /// <summary>
        /// Create a Reverse resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Reverse(string name, ReverseArgs args, CustomResourceOptions? options = null)
            : base("ovh:Ip/reverse:Reverse", name, args ?? new ReverseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Reverse(string name, Input<string> id, ReverseState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Ip/reverse:Reverse", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Reverse resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Reverse Get(string name, Input<string> id, ReverseState? state = null, CustomResourceOptions? options = null)
        {
            return new Reverse(name, id, state, options);
        }
    }

    public sealed class ReverseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP to set the reverse of
        /// </summary>
        [Input("ReverseIp", required: true)]
        public Input<string> ReverseIp { get; set; } = null!;

        /// <summary>
        /// The value of the reverse
        /// </summary>
        [Input("ReverseValue", required: true)]
        public Input<string> ReverseValue { get; set; } = null!;

        /// <summary>
        /// The IP block to which the IP belongs
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        public ReverseArgs()
        {
        }
        public static new ReverseArgs Empty => new ReverseArgs();
    }

    public sealed class ReverseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IP to set the reverse of
        /// </summary>
        [Input("ReverseIp")]
        public Input<string>? ReverseIp { get; set; }

        /// <summary>
        /// The value of the reverse
        /// </summary>
        [Input("ReverseValue")]
        public Input<string>? ReverseValue { get; set; }

        /// <summary>
        /// The IP block to which the IP belongs
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public ReverseState()
        {
        }
        public static new ReverseState Empty => new ReverseState();
    }
}
