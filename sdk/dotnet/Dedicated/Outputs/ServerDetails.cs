// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Outputs
{

    [OutputType]
    public sealed class ServerDetails
    {
        /// <summary>
        /// Personnal hostname to use in server reinstallation
        /// </summary>
        public readonly string? CustomHostname;
        /// <summary>
        /// Disk group id to process install on (only available for some templates)
        /// </summary>
        public readonly double? DiskGroupId;
        /// <summary>
        /// true if you want to install only on the first disk
        /// </summary>
        public readonly bool? NoRaid;
        /// <summary>
        /// Number of devices to use for system's software RAID
        /// </summary>
        public readonly double? SoftRaidDevices;

        [OutputConstructor]
        private ServerDetails(
            string? customHostname,

            double? diskGroupId,

            bool? noRaid,

            double? softRaidDevices)
        {
            CustomHostname = customHostname;
            DiskGroupId = diskGroupId;
            NoRaid = noRaid;
            SoftRaidDevices = softRaidDevices;
        }
    }
}