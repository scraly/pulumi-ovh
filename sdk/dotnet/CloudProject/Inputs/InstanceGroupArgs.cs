// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class InstanceGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Group id
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        public InstanceGroupArgs()
        {
        }
        public static new InstanceGroupArgs Empty => new InstanceGroupArgs();
    }
}
