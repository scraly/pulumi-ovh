// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vps
{
    public static class GetVps
    {
        /// <summary>
        /// Use this data source to retrieve information about a vps associated with your OVHcloud Account.
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
        ///     var server = Ovh.Vps.GetVps.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetVpsResult> InvokeAsync(GetVpsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVpsResult>("ovh:Vps/getVps:getVps", args ?? new GetVpsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a vps associated with your OVHcloud Account.
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
        ///     var server = Ovh.Vps.GetVps.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetVpsResult> Invoke(GetVpsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVpsResult>("ovh:Vps/getVps:getVps", args ?? new GetVpsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVpsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetVpsArgs()
        {
        }
        public static new GetVpsArgs Empty => new GetVpsArgs();
    }

    public sealed class GetVpsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetVpsInvokeArgs()
        {
        }
        public static new GetVpsInvokeArgs Empty => new GetVpsInvokeArgs();
    }


    [OutputType]
    public sealed class GetVpsResult
    {
        /// <summary>
        /// The URN of the vps
        /// </summary>
        public readonly string VpsURN;
        /// <summary>
        /// The OVHcloud cluster the vps is in
        /// </summary>
        public readonly string Cluster;
        /// <summary>
        /// The datacenter in which the vps is located
        /// </summary>
        public readonly ImmutableDictionary<string, string> Datacenter;
        /// <summary>
        /// The displayed name in the OVHcloud web admin
        /// </summary>
        public readonly string Displayname;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of IPs addresses attached to the vps
        /// </summary>
        public readonly ImmutableArray<string> Ips;
        /// <summary>
        /// The keymap for the ip kvm, valid values "", "fr", "us"
        /// </summary>
        public readonly string Keymap;
        /// <summary>
        /// The amount of memory in MB of the vps.
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// A dict describing the type of vps.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Model;
        public readonly string Name;
        /// <summary>
        /// The source of the boot kernel
        /// </summary>
        public readonly string Netbootmode;
        /// <summary>
        /// The type of offer (ssd, cloud, classic)
        /// </summary>
        public readonly string Offertype;
        public readonly string ServiceName;
        /// <summary>
        /// A boolean to indicate if OVHcloud SLA monitoring is active.
        /// </summary>
        public readonly bool Slamonitoring;
        /// <summary>
        /// The state of the vps
        /// </summary>
        public readonly string State;
        /// <summary>
        /// The type of server
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The number of vcore of the vps
        /// </summary>
        public readonly int Vcore;
        /// <summary>
        /// The OVHcloud zone where the vps is
        /// </summary>
        public readonly string Zone;

        [OutputConstructor]
        private GetVpsResult(
            string VpsURN,

            string cluster,

            ImmutableDictionary<string, string> datacenter,

            string displayname,

            string id,

            ImmutableArray<string> ips,

            string keymap,

            int memory,

            ImmutableDictionary<string, string> model,

            string name,

            string netbootmode,

            string offertype,

            string serviceName,

            bool slamonitoring,

            string state,

            string type,

            int vcore,

            string zone)
        {
            this.VpsURN = VpsURN;
            Cluster = cluster;
            Datacenter = datacenter;
            Displayname = displayname;
            Id = id;
            Ips = ips;
            Keymap = keymap;
            Memory = memory;
            Model = model;
            Name = name;
            Netbootmode = netbootmode;
            Offertype = offertype;
            ServiceName = serviceName;
            Slamonitoring = slamonitoring;
            State = state;
            Type = type;
            Vcore = vcore;
            Zone = zone;
        }
    }
}
