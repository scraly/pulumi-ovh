// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class KubeCustomizationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("apiservers")]
        private InputList<Inputs.KubeCustomizationApiserverGetArgs>? _apiservers;

        /// <summary>
        /// Kubernetes API server customization
        /// </summary>
        [Obsolete(@"Use customization_apiserver instead")]
        public InputList<Inputs.KubeCustomizationApiserverGetArgs> Apiservers
        {
            get => _apiservers ?? (_apiservers = new InputList<Inputs.KubeCustomizationApiserverGetArgs>());
            set => _apiservers = value;
        }

        public KubeCustomizationGetArgs()
        {
        }
        public static new KubeCustomizationGetArgs Empty => new KubeCustomizationGetArgs();
    }
}