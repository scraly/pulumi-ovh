// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Scraly.PulumiPackage.Ovh.CloudProject.Outputs
{

    [OutputType]
    public sealed class GetKubeCustomizationKubeProxyIpvsResult
    {
        /// <summary>
        /// Minimum period that IPVS rules are refreshed in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
        /// </summary>
        public readonly string? MinSyncPeriod;
        /// <summary>
        /// IPVS scheduler.
        /// </summary>
        public readonly string? Scheduler;
        /// <summary>
        /// Minimum period that IPVS rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format.
        /// </summary>
        public readonly string? SyncPeriod;
        /// <summary>
        /// Timeout value used for IPVS TCP sessions after receiving a FIN in RFC3339 duration.
        /// </summary>
        public readonly string? TcpFinTimeout;
        /// <summary>
        /// Timeout value used for idle IPVS TCP sessions in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
        /// </summary>
        public readonly string? TcpTimeout;
        /// <summary>
        /// timeout value used for IPVS UDP packets in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
        /// </summary>
        public readonly string? UdpTimeout;

        [OutputConstructor]
        private GetKubeCustomizationKubeProxyIpvsResult(
            string? minSyncPeriod,

            string? scheduler,

            string? syncPeriod,

            string? tcpFinTimeout,

            string? tcpTimeout,

            string? udpTimeout)
        {
            MinSyncPeriod = minSyncPeriod;
            Scheduler = scheduler;
            SyncPeriod = syncPeriod;
            TcpFinTimeout = tcpFinTimeout;
            TcpTimeout = tcpTimeout;
            UdpTimeout = udpTimeout;
        }
    }
}