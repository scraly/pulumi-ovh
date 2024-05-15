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
    /// Creates an identity user.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myUser = new Ovh.Me.IdentityUser("myUser", new()
    ///     {
    ///         Description = "Some custom description",
    ///         Email = "my_login@example.com",
    ///         Group = "DEFAULT",
    ///         Login = "my_login",
    ///         Password = "super-s3cr3t!password",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:Me/identityUser:IdentityUser")]
    public partial class IdentityUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// URN of the user, used when writing IAM policies
        /// </summary>
        [Output("UserURN")]
        public Output<string> UserURN { get; private set; } = null!;

        /// <summary>
        /// Creation date of this user.
        /// </summary>
        [Output("creation")]
        public Output<string> Creation { get; private set; } = null!;

        /// <summary>
        /// User description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// User's email.
        /// </summary>
        [Output("email")]
        public Output<string> Email { get; private set; } = null!;

        /// <summary>
        /// User's group.
        /// </summary>
        [Output("group")]
        public Output<string?> Group { get; private set; } = null!;

        /// <summary>
        /// Last update of this user.
        /// </summary>
        [Output("lastUpdate")]
        public Output<string> LastUpdate { get; private set; } = null!;

        /// <summary>
        /// User's login suffix.
        /// </summary>
        [Output("login")]
        public Output<string> Login { get; private set; } = null!;

        /// <summary>
        /// User's password.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// When the user changed his password for the last time.
        /// </summary>
        [Output("passwordLastUpdate")]
        public Output<string> PasswordLastUpdate { get; private set; } = null!;

        /// <summary>
        /// Current user's status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityUser(string name, IdentityUserArgs args, CustomResourceOptions? options = null)
            : base("ovh:Me/identityUser:IdentityUser", name, args ?? new IdentityUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityUser(string name, Input<string> id, IdentityUserState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Me/identityUser:IdentityUser", name, state, MakeResourceOptions(options, id))
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
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IdentityUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityUser Get(string name, Input<string> id, IdentityUserState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityUser(name, id, state, options);
        }
    }

    public sealed class IdentityUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// User description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User's email.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// User's group.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// User's login suffix.
        /// </summary>
        [Input("login", required: true)]
        public Input<string> Login { get; set; } = null!;

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// User's password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public IdentityUserArgs()
        {
        }
        public static new IdentityUserArgs Empty => new IdentityUserArgs();
    }

    public sealed class IdentityUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URN of the user, used when writing IAM policies
        /// </summary>
        [Input("UserURN")]
        public Input<string>? UserURN { get; set; }

        /// <summary>
        /// Creation date of this user.
        /// </summary>
        [Input("creation")]
        public Input<string>? Creation { get; set; }

        /// <summary>
        /// User description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User's email.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// User's group.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

        /// <summary>
        /// Last update of this user.
        /// </summary>
        [Input("lastUpdate")]
        public Input<string>? LastUpdate { get; set; }

        /// <summary>
        /// User's login suffix.
        /// </summary>
        [Input("login")]
        public Input<string>? Login { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// User's password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// When the user changed his password for the last time.
        /// </summary>
        [Input("passwordLastUpdate")]
        public Input<string>? PasswordLastUpdate { get; set; }

        /// <summary>
        /// Current user's status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public IdentityUserState()
        {
        }
        public static new IdentityUserState Empty => new IdentityUserState();
    }
}
