// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    /// <summary>
    /// Creates a subnet in a private network of a public cloud project.
    /// 
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// Subnet in a private network of a public cloud project can be imported using the `service_name` , the `network_id` and the `subnet_id`, separated by "/" E.g., bash &lt;break&gt;&lt;break&gt;```sh&lt;break&gt; $ pulumi import ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet mysubnet ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678/0f0b73a4-403b-45e4-86d0-b438f1291909 &lt;break&gt;```&lt;break&gt;&lt;break&gt;
    /// </summary>
    [OvhResourceType("ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet")]
    public partial class NetworkPrivateSubnet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Ip Block representing the subnet cidr.
        /// </summary>
        [Output("cidr")]
        public Output<string> Cidr { get; private set; } = null!;

        /// <summary>
        /// Enable DHCP.
        /// Changing this forces a new resource to be created. Defaults to false.
        /// _
        /// </summary>
        [Output("dhcp")]
        public Output<bool?> Dhcp { get; private set; } = null!;

        /// <summary>
        /// Last ip for this region.
        /// Changing this value recreates the subnet.
        /// </summary>
        [Output("end")]
        public Output<string> End { get; private set; } = null!;

        /// <summary>
        /// The IP of the gateway
        /// </summary>
        [Output("gatewayIp")]
        public Output<string> GatewayIp { get; private set; } = null!;

        /// <summary>
        /// List of ip pools allocated in the subnet.
        /// * `ip_pools/network` - Global network with cidr.
        /// * `ip_pools/region` - Region where this subnet is created.
        /// * `ip_pools/dhcp` - DHCP enabled.
        /// * `ip_pools/end` - Last ip for this region.
        /// * `ip_pools/start` - First ip for this region.
        /// </summary>
        [Output("ipPools")]
        public Output<ImmutableArray<Outputs.NetworkPrivateSubnetIpPool>> IpPools { get; private set; } = null!;

        /// <summary>
        /// Global network in CIDR format.
        /// Changing this value recreates the subnet
        /// </summary>
        [Output("network")]
        public Output<string> Network { get; private set; } = null!;

        /// <summary>
        /// The id of the network.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// Set to true if you don't want to set a default gateway IP.
        /// Changing this value recreates the resource. Defaults to false.
        /// </summary>
        [Output("noGateway")]
        public Output<bool?> NoGateway { get; private set; } = null!;

        /// <summary>
        /// The region in which the network subnet will be created.
        /// Ex.: "GRA1". Changing this value recreates the resource.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// First ip for this region.
        /// Changing this value recreates the subnet.
        /// </summary>
        [Output("start")]
        public Output<string> Start { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkPrivateSubnet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkPrivateSubnet(string name, NetworkPrivateSubnetArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet", name, args ?? new NetworkPrivateSubnetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkPrivateSubnet(string name, Input<string> id, NetworkPrivateSubnetState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkPrivateSubnet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkPrivateSubnet Get(string name, Input<string> id, NetworkPrivateSubnetState? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkPrivateSubnet(name, id, state, options);
        }
    }

    public sealed class NetworkPrivateSubnetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable DHCP.
        /// Changing this forces a new resource to be created. Defaults to false.
        /// _
        /// </summary>
        [Input("dhcp")]
        public Input<bool>? Dhcp { get; set; }

        /// <summary>
        /// Last ip for this region.
        /// Changing this value recreates the subnet.
        /// </summary>
        [Input("end", required: true)]
        public Input<string> End { get; set; } = null!;

        /// <summary>
        /// Global network in CIDR format.
        /// Changing this value recreates the subnet
        /// </summary>
        [Input("network", required: true)]
        public Input<string> Network { get; set; } = null!;

        /// <summary>
        /// The id of the network.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// Set to true if you don't want to set a default gateway IP.
        /// Changing this value recreates the resource. Defaults to false.
        /// </summary>
        [Input("noGateway")]
        public Input<bool>? NoGateway { get; set; }

        /// <summary>
        /// The region in which the network subnet will be created.
        /// Ex.: "GRA1". Changing this value recreates the resource.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// First ip for this region.
        /// Changing this value recreates the subnet.
        /// </summary>
        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        public NetworkPrivateSubnetArgs()
        {
        }
        public static new NetworkPrivateSubnetArgs Empty => new NetworkPrivateSubnetArgs();
    }

    public sealed class NetworkPrivateSubnetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Ip Block representing the subnet cidr.
        /// </summary>
        [Input("cidr")]
        public Input<string>? Cidr { get; set; }

        /// <summary>
        /// Enable DHCP.
        /// Changing this forces a new resource to be created. Defaults to false.
        /// _
        /// </summary>
        [Input("dhcp")]
        public Input<bool>? Dhcp { get; set; }

        /// <summary>
        /// Last ip for this region.
        /// Changing this value recreates the subnet.
        /// </summary>
        [Input("end")]
        public Input<string>? End { get; set; }

        /// <summary>
        /// The IP of the gateway
        /// </summary>
        [Input("gatewayIp")]
        public Input<string>? GatewayIp { get; set; }

        [Input("ipPools")]
        private InputList<Inputs.NetworkPrivateSubnetIpPoolGetArgs>? _ipPools;

        /// <summary>
        /// List of ip pools allocated in the subnet.
        /// * `ip_pools/network` - Global network with cidr.
        /// * `ip_pools/region` - Region where this subnet is created.
        /// * `ip_pools/dhcp` - DHCP enabled.
        /// * `ip_pools/end` - Last ip for this region.
        /// * `ip_pools/start` - First ip for this region.
        /// </summary>
        public InputList<Inputs.NetworkPrivateSubnetIpPoolGetArgs> IpPools
        {
            get => _ipPools ?? (_ipPools = new InputList<Inputs.NetworkPrivateSubnetIpPoolGetArgs>());
            set => _ipPools = value;
        }

        /// <summary>
        /// Global network in CIDR format.
        /// Changing this value recreates the subnet
        /// </summary>
        [Input("network")]
        public Input<string>? Network { get; set; }

        /// <summary>
        /// The id of the network.
        /// Changing this forces a new resource to be created.
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// Set to true if you don't want to set a default gateway IP.
        /// Changing this value recreates the resource. Defaults to false.
        /// </summary>
        [Input("noGateway")]
        public Input<bool>? NoGateway { get; set; }

        /// <summary>
        /// The region in which the network subnet will be created.
        /// Ex.: "GRA1". Changing this value recreates the resource.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// First ip for this region.
        /// Changing this value recreates the subnet.
        /// </summary>
        [Input("start")]
        public Input<string>? Start { get; set; }

        public NetworkPrivateSubnetState()
        {
        }
        public static new NetworkPrivateSubnetState Empty => new NetworkPrivateSubnetState();
    }
}
