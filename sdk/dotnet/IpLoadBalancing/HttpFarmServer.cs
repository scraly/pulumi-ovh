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
    /// Creates a backend server entry linked to HTTP loadbalancing group (farm)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var lb = Ovh.IpLoadBalancing.GetIpLoadBalancing.Invoke(new()
    ///     {
    ///         ServiceName = "ip-1.2.3.4",
    ///         State = "ok",
    ///     });
    /// 
    ///     var farmname = new Ovh.IpLoadBalancing.HttpFarm("farmname", new()
    ///     {
    ///         Port = 8080,
    ///         ServiceName = lb.Apply(getIpLoadBalancingResult =&gt; getIpLoadBalancingResult.ServiceName),
    ///         Zone = "all",
    ///     });
    /// 
    ///     var backend = new Ovh.IpLoadBalancing.HttpFarmServer("backend", new()
    ///     {
    ///         Address = "4.5.6.7",
    ///         Backup = true,
    ///         DisplayName = "mybackend",
    ///         FarmId = farmname.Id,
    ///         Port = 80,
    ///         Probe = true,
    ///         ProxyProtocolVersion = "v2",
    ///         ServiceName = lb.Apply(getIpLoadBalancingResult =&gt; getIpLoadBalancingResult.ServiceName),
    ///         Ssl = false,
    ///         Status = "active",
    ///         Weight = 2,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// HTTP farm server can be imported using the following format `service_name`, the `id` of the farm and the `id` of the server separated by "/" e.g.
    /// </summary>
    [OvhResourceType("ovh:IpLoadBalancing/httpFarmServer:HttpFarmServer")]
    public partial class HttpFarmServer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Address of the backend server (IP from either internal or OVHcloud network)
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// is it a backup server used in case of failure of all the non-backup backends
        /// </summary>
        [Output("backup")]
        public Output<bool?> Backup { get; private set; } = null!;

        [Output("chain")]
        public Output<string?> Chain { get; private set; } = null!;

        /// <summary>
        /// Value of the stickiness cookie used for this backend.
        /// </summary>
        [Output("cookie")]
        public Output<string> Cookie { get; private set; } = null!;

        /// <summary>
        /// Label for the server
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// ID of the farm this server is attached to
        /// </summary>
        [Output("farmId")]
        public Output<int> FarmId { get; private set; } = null!;

        /// <summary>
        /// enable action when backend marked down. (`shutdown-sessions`)
        /// </summary>
        [Output("onMarkedDown")]
        public Output<string?> OnMarkedDown { get; private set; } = null!;

        /// <summary>
        /// Port that backend will respond on
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// defines if backend will be probed to determine health and keep as active in farm if healthy
        /// </summary>
        [Output("probe")]
        public Output<bool?> Probe { get; private set; } = null!;

        /// <summary>
        /// version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
        /// </summary>
        [Output("proxyProtocolVersion")]
        public Output<string?> ProxyProtocolVersion { get; private set; } = null!;

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// is the connection ciphered with SSL (TLS)
        /// </summary>
        [Output("ssl")]
        public Output<bool?> Ssl { get; private set; } = null!;

        /// <summary>
        /// backend status - `active` or `inactive`
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// used in loadbalancing algorithm
        /// </summary>
        [Output("weight")]
        public Output<int?> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a HttpFarmServer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HttpFarmServer(string name, HttpFarmServerArgs args, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/httpFarmServer:HttpFarmServer", name, args ?? new HttpFarmServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HttpFarmServer(string name, Input<string> id, HttpFarmServerState? state = null, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/httpFarmServer:HttpFarmServer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HttpFarmServer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HttpFarmServer Get(string name, Input<string> id, HttpFarmServerState? state = null, CustomResourceOptions? options = null)
        {
            return new HttpFarmServer(name, id, state, options);
        }
    }

    public sealed class HttpFarmServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address of the backend server (IP from either internal or OVHcloud network)
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// is it a backup server used in case of failure of all the non-backup backends
        /// </summary>
        [Input("backup")]
        public Input<bool>? Backup { get; set; }

        [Input("chain")]
        public Input<string>? Chain { get; set; }

        /// <summary>
        /// Label for the server
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// ID of the farm this server is attached to
        /// </summary>
        [Input("farmId", required: true)]
        public Input<int> FarmId { get; set; } = null!;

        /// <summary>
        /// enable action when backend marked down. (`shutdown-sessions`)
        /// </summary>
        [Input("onMarkedDown")]
        public Input<string>? OnMarkedDown { get; set; }

        /// <summary>
        /// Port that backend will respond on
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// defines if backend will be probed to determine health and keep as active in farm if healthy
        /// </summary>
        [Input("probe")]
        public Input<bool>? Probe { get; set; }

        /// <summary>
        /// version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
        /// </summary>
        [Input("proxyProtocolVersion")]
        public Input<string>? ProxyProtocolVersion { get; set; }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// is the connection ciphered with SSL (TLS)
        /// </summary>
        [Input("ssl")]
        public Input<bool>? Ssl { get; set; }

        /// <summary>
        /// backend status - `active` or `inactive`
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        /// <summary>
        /// used in loadbalancing algorithm
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public HttpFarmServerArgs()
        {
        }
        public static new HttpFarmServerArgs Empty => new HttpFarmServerArgs();
    }

    public sealed class HttpFarmServerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address of the backend server (IP from either internal or OVHcloud network)
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// is it a backup server used in case of failure of all the non-backup backends
        /// </summary>
        [Input("backup")]
        public Input<bool>? Backup { get; set; }

        [Input("chain")]
        public Input<string>? Chain { get; set; }

        /// <summary>
        /// Value of the stickiness cookie used for this backend.
        /// </summary>
        [Input("cookie")]
        public Input<string>? Cookie { get; set; }

        /// <summary>
        /// Label for the server
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// ID of the farm this server is attached to
        /// </summary>
        [Input("farmId")]
        public Input<int>? FarmId { get; set; }

        /// <summary>
        /// enable action when backend marked down. (`shutdown-sessions`)
        /// </summary>
        [Input("onMarkedDown")]
        public Input<string>? OnMarkedDown { get; set; }

        /// <summary>
        /// Port that backend will respond on
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// defines if backend will be probed to determine health and keep as active in farm if healthy
        /// </summary>
        [Input("probe")]
        public Input<bool>? Probe { get; set; }

        /// <summary>
        /// version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (`v1`, `v2`, `v2-ssl`, `v2-ssl-cn`)
        /// </summary>
        [Input("proxyProtocolVersion")]
        public Input<string>? ProxyProtocolVersion { get; set; }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// is the connection ciphered with SSL (TLS)
        /// </summary>
        [Input("ssl")]
        public Input<bool>? Ssl { get; set; }

        /// <summary>
        /// backend status - `active` or `inactive`
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// used in loadbalancing algorithm
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public HttpFarmServerState()
        {
        }
        public static new HttpFarmServerState Empty => new HttpFarmServerState();
    }
}
