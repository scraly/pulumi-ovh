// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me
{
    /// <summary>
    /// Creates an OAuth2 service account.
    /// 
    /// ## Example Usage
    /// 
    /// An OAuth2 client for an app hosted at `my-app.com`, that uses the authorization code flow to authenticate.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myOauth2ClientAuthCode = new Ovh.Me.APIOAuth2Client("myOauth2ClientAuthCode", new()
    ///     {
    ///         CallbackUrls = new[]
    ///         {
    ///             "https://my-app.com/callback",
    ///         },
    ///         Description = "An OAuth2 client using the authorization code flow for my-app.com",
    ///         Flow = "AUTHORIZATION_CODE",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// An OAuth2 client for an app hosted at `my-app.com`, that uses the client credentials flow to authenticate.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myOauth2ClientClientCreds = new Ovh.Me.APIOAuth2Client("myOauth2ClientClientCreds", new()
    ///     {
    ///         Description = "An OAuth2 client using the client credentials flow for my app",
    ///         Flow = "CLIENT_CREDENTIALS",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OAuth2 clients can be imported using their `client_id`bash
    /// 
    /// ```sh
    ///  $ pulumi import ovh:Me/aPIOAuth2Client:APIOAuth2Client my_oauth2_client client_id
    /// ```
    /// 
    ///  Because the client_secret is only available for resources created using terraform, OAuth2 clients can also be imported using a `client_id` and a `client_secret` with a pipe separatorbash
    /// 
    /// ```sh
    ///  $ pulumi import ovh:Me/aPIOAuth2Client:APIOAuth2Client my_oauth2_client 'client_id|client_secret'
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Me/aPIOAuth2Client:APIOAuth2Client")]
    public partial class APIOAuth2Client : global::Pulumi.CustomResource
    {
        /// <summary>
        /// List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
        /// </summary>
        [Output("callbackUrls")]
        public Output<ImmutableArray<string>> CallbackUrls { get; private set; } = null!;

        /// <summary>
        /// Client ID of the created service account.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// Client secret of the created service account.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// OAuth2 client description.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
        /// </summary>
        [Output("flow")]
        public Output<string> Flow { get; private set; } = null!;

        /// <summary>
        /// URN that will allow you to associate this oauth2 client with an access policy
        /// </summary>
        [Output("identity")]
        public Output<string> Identity { get; private set; } = null!;

        /// <summary>
        /// OAuth2 client name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a APIOAuth2Client resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public APIOAuth2Client(string name, APIOAuth2ClientArgs args, CustomResourceOptions? options = null)
            : base("ovh:Me/aPIOAuth2Client:APIOAuth2Client", name, args ?? new APIOAuth2ClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private APIOAuth2Client(string name, Input<string> id, APIOAuth2ClientState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Me/aPIOAuth2Client:APIOAuth2Client", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
                AdditionalSecretOutputs =
                {
                    "clientSecret",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing APIOAuth2Client resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static APIOAuth2Client Get(string name, Input<string> id, APIOAuth2ClientState? state = null, CustomResourceOptions? options = null)
        {
            return new APIOAuth2Client(name, id, state, options);
        }
    }

    public sealed class APIOAuth2ClientArgs : global::Pulumi.ResourceArgs
    {
        [Input("callbackUrls")]
        private InputList<string>? _callbackUrls;

        /// <summary>
        /// List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
        /// </summary>
        public InputList<string> CallbackUrls
        {
            get => _callbackUrls ?? (_callbackUrls = new InputList<string>());
            set => _callbackUrls = value;
        }

        /// <summary>
        /// OAuth2 client description.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
        /// </summary>
        [Input("flow", required: true)]
        public Input<string> Flow { get; set; } = null!;

        /// <summary>
        /// OAuth2 client name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public APIOAuth2ClientArgs()
        {
        }
        public static new APIOAuth2ClientArgs Empty => new APIOAuth2ClientArgs();
    }

    public sealed class APIOAuth2ClientState : global::Pulumi.ResourceArgs
    {
        [Input("callbackUrls")]
        private InputList<string>? _callbackUrls;

        /// <summary>
        /// List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
        /// </summary>
        public InputList<string> CallbackUrls
        {
            get => _callbackUrls ?? (_callbackUrls = new InputList<string>());
            set => _callbackUrls = value;
        }

        /// <summary>
        /// Client ID of the created service account.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("clientSecret")]
        private Input<string>? _clientSecret;

        /// <summary>
        /// Client secret of the created service account.
        /// </summary>
        public Input<string>? ClientSecret
        {
            get => _clientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _clientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// OAuth2 client description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
        /// </summary>
        [Input("flow")]
        public Input<string>? Flow { get; set; }

        /// <summary>
        /// URN that will allow you to associate this oauth2 client with an access policy
        /// </summary>
        [Input("identity")]
        public Input<string>? Identity { get; set; }

        /// <summary>
        /// OAuth2 client name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public APIOAuth2ClientState()
        {
        }
        public static new APIOAuth2ClientState Empty => new APIOAuth2ClientState();
    }
}