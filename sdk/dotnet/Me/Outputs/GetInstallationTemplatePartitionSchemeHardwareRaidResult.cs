// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ediri.Ovh.Me.Outputs
{

    [OutputType]
    public sealed class GetInstallationTemplatePartitionSchemeHardwareRaidResult
    {
        /// <summary>
        /// Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
        /// </summary>
        public readonly ImmutableArray<string> Disks;
        /// <summary>
        /// RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
        /// </summary>
        public readonly string Mode;
        /// <summary>
        /// Hardware RAID name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies the creation order of the hardware RAID.
        /// </summary>
        public readonly int Step;

        [OutputConstructor]
        private GetInstallationTemplatePartitionSchemeHardwareRaidResult(
            ImmutableArray<string> disks,

            string mode,

            string name,

            int step)
        {
            Disks = disks;
            Mode = mode;
            Name = name;
            Step = step;
        }
    }
}