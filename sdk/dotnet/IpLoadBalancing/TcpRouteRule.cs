// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.IpLoadBalancing
{
    /// <summary>
    /// Manage rules for TCP route.
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
    ///     var reject = new Ovh.IpLoadBalancing.TcpRoute("reject", new()
    ///     {
    ///         ServiceName = "loadbalancer-xxxxxxxxxxxxxxxxxx",
    ///         Weight = 1,
    ///         FrontendId = 11111,
    ///         Action = new Ovh.IpLoadBalancing.Inputs.TcpRouteActionArgs
    ///         {
    ///             Type = "reject",
    ///         },
    ///     });
    /// 
    ///     var examplerule = new Ovh.IpLoadBalancing.TcpRouteRule("examplerule", new()
    ///     {
    ///         ServiceName = "loadbalancer-xxxxxxxxxxxxxxxxxx",
    ///         RouteId = reject.Id,
    ///         DisplayName = "Match example.com host",
    ///         Field = "sni",
    ///         Match = "is",
    ///         Negate = false,
    ///         Pattern = "example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// TCP route rule can be imported using the following format `service_name`, the `id` of the route and the `id` of the rule separated by "/" e.g.
    /// </summary>
    [OvhResourceType("ovh:IpLoadBalancing/tcpRouteRule:TcpRouteRule")]
    public partial class TcpRouteRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Human readable name for your rule, this field is for you
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
        /// </summary>
        [Output("field")]
        public Output<string> Field { get; private set; } = null!;

        /// <summary>
        /// Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
        /// </summary>
        [Output("match")]
        public Output<string> Match { get; private set; } = null!;

        /// <summary>
        /// Invert the matching operator effect
        /// </summary>
        [Output("negate")]
        public Output<bool> Negate { get; private set; } = null!;

        /// <summary>
        /// Value to match against this match. Interpretation if this field depends on the match and field
        /// </summary>
        [Output("pattern")]
        public Output<string?> Pattern { get; private set; } = null!;

        /// <summary>
        /// The route to apply this rule
        /// </summary>
        [Output("routeId")]
        public Output<string> RouteId { get; private set; } = null!;

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
        /// </summary>
        [Output("subField")]
        public Output<string?> SubField { get; private set; } = null!;


        /// <summary>
        /// Create a TcpRouteRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TcpRouteRule(string name, TcpRouteRuleArgs args, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/tcpRouteRule:TcpRouteRule", name, args ?? new TcpRouteRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TcpRouteRule(string name, Input<string> id, TcpRouteRuleState? state = null, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/tcpRouteRule:TcpRouteRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TcpRouteRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TcpRouteRule Get(string name, Input<string> id, TcpRouteRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new TcpRouteRule(name, id, state, options);
        }
    }

    public sealed class TcpRouteRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human readable name for your rule, this field is for you
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
        /// </summary>
        [Input("field", required: true)]
        public Input<string> Field { get; set; } = null!;

        /// <summary>
        /// Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
        /// </summary>
        [Input("match", required: true)]
        public Input<string> Match { get; set; } = null!;

        /// <summary>
        /// Invert the matching operator effect
        /// </summary>
        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        /// <summary>
        /// Value to match against this match. Interpretation if this field depends on the match and field
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// The route to apply this rule
        /// </summary>
        [Input("routeId", required: true)]
        public Input<string> RouteId { get; set; } = null!;

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
        /// </summary>
        [Input("subField")]
        public Input<string>? SubField { get; set; }

        public TcpRouteRuleArgs()
        {
        }
        public static new TcpRouteRuleArgs Empty => new TcpRouteRuleArgs();
    }

    public sealed class TcpRouteRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human readable name for your rule, this field is for you
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
        /// </summary>
        [Input("field")]
        public Input<string>? Field { get; set; }

        /// <summary>
        /// Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
        /// </summary>
        [Input("match")]
        public Input<string>? Match { get; set; }

        /// <summary>
        /// Invert the matching operator effect
        /// </summary>
        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        /// <summary>
        /// Value to match against this match. Interpretation if this field depends on the match and field
        /// </summary>
        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        /// <summary>
        /// The route to apply this rule
        /// </summary>
        [Input("routeId")]
        public Input<string>? RouteId { get; set; }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Name of sub-field, if applicable. This may be a Cookie or Header name for instance
        /// </summary>
        [Input("subField")]
        public Input<string>? SubField { get; set; }

        public TcpRouteRuleState()
        {
        }
        public static new TcpRouteRuleState Empty => new TcpRouteRuleState();
    }
}
