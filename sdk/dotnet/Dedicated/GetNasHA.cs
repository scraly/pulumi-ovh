// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated
{
    public static class GetNasHA
    {
        /// <summary>
        /// Use this data source to retrieve information about a dedicated HA-NAS.
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
        ///     var my_nas_ha = Ovh.Dedicated.GetNasHA.Invoke(new()
        ///     {
        ///         ServiceName = "zpool-12345",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetNasHAResult> InvokeAsync(GetNasHAArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNasHAResult>("ovh:Dedicated/getNasHA:getNasHA", args ?? new GetNasHAArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about a dedicated HA-NAS.
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
        ///     var my_nas_ha = Ovh.Dedicated.GetNasHA.Invoke(new()
        ///     {
        ///         ServiceName = "zpool-12345",
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetNasHAResult> Invoke(GetNasHAInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNasHAResult>("ovh:Dedicated/getNasHA:getNasHA", args ?? new GetNasHAInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNasHAArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The service_name of your dedicated HA-NAS.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetNasHAArgs()
        {
        }
        public static new GetNasHAArgs Empty => new GetNasHAArgs();
    }

    public sealed class GetNasHAInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The service_name of your dedicated HA-NAS.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetNasHAInvokeArgs()
        {
        }
        public static new GetNasHAInvokeArgs Empty => new GetNasHAInvokeArgs();
    }


    [OutputType]
    public sealed class GetNasHAResult
    {
        /// <summary>
        /// the URN of the HA-NAS instance
        /// </summary>
        public readonly string NasHAURN;
        /// <summary>
        /// True, if partition creation is allowed on this HA-NAS
        /// </summary>
        public readonly bool CanCreatePartition;
        /// <summary>
        /// The name you give to the HA-NAS
        /// </summary>
        public readonly string CustomName;
        /// <summary>
        /// area of HA-NAS
        /// </summary>
        public readonly string Datacenter;
        /// <summary>
        /// the disk type of the HA-NAS. Possible values are: `hdd`, `ssd`, `nvme`
        /// </summary>
        public readonly string DiskType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Access IP of HA-NAS
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// Send an email to customer if any issue is detected
        /// </summary>
        public readonly bool Monitored;
        /// <summary>
        /// The storage service name
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// percentage of HA-NAS space used in %!
        /// (MISSING)
        /// </summary>
        public readonly double ZpoolCapacity;
        /// <summary>
        /// the size of the HA-NAS in GB
        /// </summary>
        public readonly double ZpoolSize;

        [OutputConstructor]
        private GetNasHAResult(
            string NasHAURN,

            bool canCreatePartition,

            string customName,

            string datacenter,

            string diskType,

            string id,

            string ip,

            bool monitored,

            string serviceName,

            double zpoolCapacity,

            double zpoolSize)
        {
            this.NasHAURN = NasHAURN;
            CanCreatePartition = canCreatePartition;
            CustomName = customName;
            Datacenter = datacenter;
            DiskType = diskType;
            Id = id;
            Ip = ip;
            Monitored = monitored;
            ServiceName = serviceName;
            ZpoolCapacity = zpoolCapacity;
            ZpoolSize = zpoolSize;
        }
    }
}
