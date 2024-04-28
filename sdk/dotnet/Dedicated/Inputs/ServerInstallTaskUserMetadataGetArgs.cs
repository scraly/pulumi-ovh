// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Inputs
{

    public sealed class ServerInstallTaskUserMetadataGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key for the user_metadata
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The value for the user_metadata
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ServerInstallTaskUserMetadataGetArgs()
        {
        }
        public static new ServerInstallTaskUserMetadataGetArgs Empty => new ServerInstallTaskUserMetadataGetArgs();
    }
}