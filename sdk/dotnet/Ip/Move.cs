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
    /// Moves a given IP to a different service, or inversely, parks it if empty service is given
    /// 
    /// ## Move IP `1.2.3.4` to service loadbalancer-XXXXX
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
    ///     var moveIpToLoadBalancerXxxxx = new Ovh.Ip.Move("moveIpToLoadBalancerXxxxx", new()
    ///     {
    ///         Ip = "1.2.3.4",
    ///         RoutedTo = new Ovh.Ip.Inputs.MoveRoutedToArgs
    ///         {
    ///             ServiceName = "loadbalancer-XXXXX",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Park IP/Detach IP `1.2.3.4` from any service
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
    ///     var parkIp = new Ovh.Ip.Move("parkIp", new()
    ///     {
    ///         Ip = "1.2.3.4",
    ///         RoutedTo = new Ovh.Ip.Inputs.MoveRoutedToArgs
    ///         {
    ///             ServiceName = "",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [OvhResourceType("ovh:Ip/move:Move")]
    public partial class Move : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether IP service can be terminated
        /// </summary>
        [Output("canBeTerminated")]
        public Output<bool> CanBeTerminated { get; private set; } = null!;

        /// <summary>
        /// Country
        /// </summary>
        [Output("country")]
        public Output<string> Country { get; private set; } = null!;

        /// <summary>
        /// Description attached to the IP
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// IP block that we want to attach to a different service
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// IP block organisation Id
        /// </summary>
        [Output("organisationId")]
        public Output<string> OrganisationId { get; private set; } = null!;

        /// <summary>
        /// Service to route the IP to. If null, the IP will be [parked](https://api.ovh.com/console/#/ip/%7Bip%7D/park~POST)
        /// instead of [moved](https://api.ovh.com/console/#/ip/%7Bip%7D/move~POST)
        /// </summary>
        [Output("routedTo")]
        public Output<Outputs.MoveRoutedTo> RoutedTo { get; private set; } = null!;

        /// <summary>
        /// Name of the service to route the IP to. IP will be parked if this value is an empty string
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Starting date and time field of the current IP task that is in charge of changing the service the IP is attached to
        /// </summary>
        [Output("taskStartDate")]
        public Output<string> TaskStartDate { get; private set; } = null!;

        /// <summary>
        /// Status field of the current IP task that is in charge of changing the service the IP is attached to
        /// </summary>
        [Output("taskStatus")]
        public Output<string> TaskStatus { get; private set; } = null!;

        /// <summary>
        /// Possible values for ip type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Move resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Move(string name, MoveArgs args, CustomResourceOptions? options = null)
            : base("ovh:Ip/move:Move", name, args ?? new MoveArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Move(string name, Input<string> id, MoveState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Ip/move:Move", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Move resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Move Get(string name, Input<string> id, MoveState? state = null, CustomResourceOptions? options = null)
        {
            return new Move(name, id, state, options);
        }
    }

    public sealed class MoveArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description attached to the IP
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// IP block that we want to attach to a different service
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// Service to route the IP to. If null, the IP will be [parked](https://api.ovh.com/console/#/ip/%7Bip%7D/park~POST)
        /// instead of [moved](https://api.ovh.com/console/#/ip/%7Bip%7D/move~POST)
        /// </summary>
        [Input("routedTo", required: true)]
        public Input<Inputs.MoveRoutedToArgs> RoutedTo { get; set; } = null!;

        public MoveArgs()
        {
        }
        public static new MoveArgs Empty => new MoveArgs();
    }

    public sealed class MoveState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether IP service can be terminated
        /// </summary>
        [Input("canBeTerminated")]
        public Input<bool>? CanBeTerminated { get; set; }

        /// <summary>
        /// Country
        /// </summary>
        [Input("country")]
        public Input<string>? Country { get; set; }

        /// <summary>
        /// Description attached to the IP
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// IP block that we want to attach to a different service
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// IP block organisation Id
        /// </summary>
        [Input("organisationId")]
        public Input<string>? OrganisationId { get; set; }

        /// <summary>
        /// Service to route the IP to. If null, the IP will be [parked](https://api.ovh.com/console/#/ip/%7Bip%7D/park~POST)
        /// instead of [moved](https://api.ovh.com/console/#/ip/%7Bip%7D/move~POST)
        /// </summary>
        [Input("routedTo")]
        public Input<Inputs.MoveRoutedToGetArgs>? RoutedTo { get; set; }

        /// <summary>
        /// Name of the service to route the IP to. IP will be parked if this value is an empty string
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Starting date and time field of the current IP task that is in charge of changing the service the IP is attached to
        /// </summary>
        [Input("taskStartDate")]
        public Input<string>? TaskStartDate { get; set; }

        /// <summary>
        /// Status field of the current IP task that is in charge of changing the service the IP is attached to
        /// </summary>
        [Input("taskStatus")]
        public Input<string>? TaskStatus { get; set; }

        /// <summary>
        /// Possible values for ip type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public MoveState()
        {
        }
        public static new MoveState Empty => new MoveState();
    }
}
