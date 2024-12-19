// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Okms.Outputs
{

    [OutputType]
    public sealed class GetOkmsServiceKeyPemKeysPemResult
    {
        /// <summary>
        /// The key in base64 encoded PEM format
        /// </summary>
        public readonly string Pem;

        [OutputConstructor]
        private GetOkmsServiceKeyPemKeysPemResult(string pem)
        {
            Pem = pem;
        }
    }
}
