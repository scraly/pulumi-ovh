// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Scraly.PulumiPackage.Ovh.Ip.Outputs
{

    [OutputType]
    public sealed class IpServiceRoutedTo
    {
        /// <summary>
        /// service name
        /// </summary>
        public readonly string? ServiceName;

        [OutputConstructor]
        private IpServiceRoutedTo(string? serviceName)
        {
            ServiceName = serviceName;
        }
    }
}