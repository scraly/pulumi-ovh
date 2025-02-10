// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated
{
    /// <summary>
    /// ## Import
    /// 
    /// Dedicated servers can be imported using the `service_name`.
    /// 
    /// Using the following configuration:
    /// 
    /// hcl
    /// 
    /// import {
    /// 
    ///   to = ovh_dedicated_server.server
    /// 
    ///   id = "&lt;service name&gt;"
    /// 
    /// }
    /// 
    /// You can then run:
    /// 
    /// bash
    /// 
    /// $ pulumi preview -generate-config-out=dedicated.tf
    /// 
    /// $ pulumi up
    /// 
    /// The file `dedicated.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above.
    /// 
    /// See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
    /// </summary>
    [OvhResourceType("ovh:Dedicated/server:Server")]
    public partial class Server : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Dedicated AZ localisation
        /// </summary>
        [Output("availabilityZone")]
        public Output<string> AvailabilityZone { get; private set; } = null!;

        /// <summary>
        /// Boot id of the server
        /// </summary>
        [Output("bootId")]
        public Output<double> BootId { get; private set; } = null!;

        /// <summary>
        /// Boot script of the server
        /// </summary>
        [Output("bootScript")]
        public Output<string> BootScript { get; private set; } = null!;

        /// <summary>
        /// Dedicated server commercial range
        /// </summary>
        [Output("commercialRange")]
        public Output<string> CommercialRange { get; private set; } = null!;

        /// <summary>
        /// Dedicated datacenter localisation (bhs1,bhs2,...)
        /// </summary>
        [Output("datacenter")]
        public Output<string> Datacenter { get; private set; } = null!;

        /// <summary>
        /// A structure describing informations about installation custom
        /// </summary>
        [Output("details")]
        public Output<Outputs.ServerDetails?> Details { get; private set; } = null!;

        /// <summary>
        /// Resource display name
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Path of the EFI bootloader
        /// </summary>
        [Output("efiBootloaderPath")]
        public Output<string> EfiBootloaderPath { get; private set; } = null!;

        /// <summary>
        /// IAM resource information
        /// </summary>
        [Output("iam")]
        public Output<Outputs.ServerIam> Iam { get; private set; } = null!;

        /// <summary>
        /// Dedicated server ip (IPv4)
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// Link speed of the server
        /// </summary>
        [Output("linkSpeed")]
        public Output<double> LinkSpeed { get; private set; } = null!;

        /// <summary>
        /// Icmp monitoring state
        /// </summary>
        [Output("monitoring")]
        public Output<bool> Monitoring { get; private set; } = null!;

        /// <summary>
        /// Dedicated server name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("newUpgradeSystem")]
        public Output<bool> NewUpgradeSystem { get; private set; } = null!;

        /// <summary>
        /// Prevent datacenter intervention
        /// </summary>
        [Output("noIntervention")]
        public Output<bool> NoIntervention { get; private set; } = null!;

        /// <summary>
        /// Details about an Order
        /// </summary>
        [Output("order")]
        public Output<Outputs.ServerOrder> Order { get; private set; } = null!;

        /// <summary>
        /// Operating system
        /// </summary>
        [Output("os")]
        public Output<string> Os { get; private set; } = null!;

        /// <summary>
        /// OVH subsidiaries
        /// </summary>
        [Output("ovhSubsidiary")]
        public Output<string?> OvhSubsidiary { get; private set; } = null!;

        /// <summary>
        /// Partition scheme name
        /// </summary>
        [Output("partitionSchemeName")]
        public Output<string?> PartitionSchemeName { get; private set; } = null!;

        [Output("planOptions")]
        public Output<ImmutableArray<Outputs.ServerPlanOption>> PlanOptions { get; private set; } = null!;

        [Output("plans")]
        public Output<ImmutableArray<Outputs.ServerPlan>> Plans { get; private set; } = null!;

        /// <summary>
        /// Power state of the server (poweron, poweroff)
        /// </summary>
        [Output("powerState")]
        public Output<string> PowerState { get; private set; } = null!;

        /// <summary>
        /// Does this server have professional use option
        /// </summary>
        [Output("professionalUse")]
        public Output<bool> ProfessionalUse { get; private set; } = null!;

        /// <summary>
        /// Rack id of the server
        /// </summary>
        [Output("rack")]
        public Output<string> Rack { get; private set; } = null!;

        /// <summary>
        /// Dedicated region localisation
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Rescue mail of the server
        /// </summary>
        [Output("rescueMail")]
        public Output<string> RescueMail { get; private set; } = null!;

        /// <summary>
        /// Public SSH Key used in the rescue mode
        /// </summary>
        [Output("rescueSshKey")]
        public Output<string> RescueSshKey { get; private set; } = null!;

        /// <summary>
        /// Dedicated server reverse
        /// </summary>
        [Output("reverse")]
        public Output<string> Reverse { get; private set; } = null!;

        /// <summary>
        /// Root device of the server
        /// </summary>
        [Output("rootDevice")]
        public Output<string> RootDevice { get; private set; } = null!;

        /// <summary>
        /// Server id
        /// </summary>
        [Output("serverId")]
        public Output<double> ServerId { get; private set; } = null!;

        /// <summary>
        /// The service_name of your dedicated server
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// All states a Dedicated can be in (error, hacked, hackedBlocked, ok)
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Dedicated server support level (critical, fastpath, gs, pro)
        /// </summary>
        [Output("supportLevel")]
        public Output<string> SupportLevel { get; private set; } = null!;

        /// <summary>
        /// Template name
        /// </summary>
        [Output("templateName")]
        public Output<string?> TemplateName { get; private set; } = null!;

        /// <summary>
        /// Metadata
        /// </summary>
        [Output("userMetadatas")]
        public Output<ImmutableArray<Outputs.ServerUserMetadata>> UserMetadatas { get; private set; } = null!;


        /// <summary>
        /// Create a Server resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Server(string name, ServerArgs? args = null, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/server:Server", name, args ?? new ServerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Server(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Dedicated/server:Server", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Server resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Server Get(string name, Input<string> id, ServerState? state = null, CustomResourceOptions? options = null)
        {
            return new Server(name, id, state, options);
        }
    }

    public sealed class ServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Boot id of the server
        /// </summary>
        [Input("bootId")]
        public Input<double>? BootId { get; set; }

        /// <summary>
        /// Boot script of the server
        /// </summary>
        [Input("bootScript")]
        public Input<string>? BootScript { get; set; }

        /// <summary>
        /// A structure describing informations about installation custom
        /// </summary>
        [Input("details")]
        public Input<Inputs.ServerDetailsArgs>? Details { get; set; }

        /// <summary>
        /// Resource display name
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Path of the EFI bootloader
        /// </summary>
        [Input("efiBootloaderPath")]
        public Input<string>? EfiBootloaderPath { get; set; }

        /// <summary>
        /// Icmp monitoring state
        /// </summary>
        [Input("monitoring")]
        public Input<bool>? Monitoring { get; set; }

        /// <summary>
        /// Prevent datacenter intervention
        /// </summary>
        [Input("noIntervention")]
        public Input<bool>? NoIntervention { get; set; }

        /// <summary>
        /// OVH subsidiaries
        /// </summary>
        [Input("ovhSubsidiary")]
        public Input<string>? OvhSubsidiary { get; set; }

        /// <summary>
        /// Partition scheme name
        /// </summary>
        [Input("partitionSchemeName")]
        public Input<string>? PartitionSchemeName { get; set; }

        [Input("planOptions")]
        private InputList<Inputs.ServerPlanOptionArgs>? _planOptions;
        public InputList<Inputs.ServerPlanOptionArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.ServerPlanOptionArgs>());
            set => _planOptions = value;
        }

        [Input("plans")]
        private InputList<Inputs.ServerPlanArgs>? _plans;
        public InputList<Inputs.ServerPlanArgs> Plans
        {
            get => _plans ?? (_plans = new InputList<Inputs.ServerPlanArgs>());
            set => _plans = value;
        }

        /// <summary>
        /// Rescue mail of the server
        /// </summary>
        [Input("rescueMail")]
        public Input<string>? RescueMail { get; set; }

        /// <summary>
        /// Public SSH Key used in the rescue mode
        /// </summary>
        [Input("rescueSshKey")]
        public Input<string>? RescueSshKey { get; set; }

        /// <summary>
        /// Root device of the server
        /// </summary>
        [Input("rootDevice")]
        public Input<string>? RootDevice { get; set; }

        /// <summary>
        /// All states a Dedicated can be in (error, hacked, hackedBlocked, ok)
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Template name
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        [Input("userMetadatas")]
        private InputList<Inputs.ServerUserMetadataArgs>? _userMetadatas;

        /// <summary>
        /// Metadata
        /// </summary>
        public InputList<Inputs.ServerUserMetadataArgs> UserMetadatas
        {
            get => _userMetadatas ?? (_userMetadatas = new InputList<Inputs.ServerUserMetadataArgs>());
            set => _userMetadatas = value;
        }

        public ServerArgs()
        {
        }
        public static new ServerArgs Empty => new ServerArgs();
    }

    public sealed class ServerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dedicated AZ localisation
        /// </summary>
        [Input("availabilityZone")]
        public Input<string>? AvailabilityZone { get; set; }

        /// <summary>
        /// Boot id of the server
        /// </summary>
        [Input("bootId")]
        public Input<double>? BootId { get; set; }

        /// <summary>
        /// Boot script of the server
        /// </summary>
        [Input("bootScript")]
        public Input<string>? BootScript { get; set; }

        /// <summary>
        /// Dedicated server commercial range
        /// </summary>
        [Input("commercialRange")]
        public Input<string>? CommercialRange { get; set; }

        /// <summary>
        /// Dedicated datacenter localisation (bhs1,bhs2,...)
        /// </summary>
        [Input("datacenter")]
        public Input<string>? Datacenter { get; set; }

        /// <summary>
        /// A structure describing informations about installation custom
        /// </summary>
        [Input("details")]
        public Input<Inputs.ServerDetailsGetArgs>? Details { get; set; }

        /// <summary>
        /// Resource display name
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Path of the EFI bootloader
        /// </summary>
        [Input("efiBootloaderPath")]
        public Input<string>? EfiBootloaderPath { get; set; }

        /// <summary>
        /// IAM resource information
        /// </summary>
        [Input("iam")]
        public Input<Inputs.ServerIamGetArgs>? Iam { get; set; }

        /// <summary>
        /// Dedicated server ip (IPv4)
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Link speed of the server
        /// </summary>
        [Input("linkSpeed")]
        public Input<double>? LinkSpeed { get; set; }

        /// <summary>
        /// Icmp monitoring state
        /// </summary>
        [Input("monitoring")]
        public Input<bool>? Monitoring { get; set; }

        /// <summary>
        /// Dedicated server name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("newUpgradeSystem")]
        public Input<bool>? NewUpgradeSystem { get; set; }

        /// <summary>
        /// Prevent datacenter intervention
        /// </summary>
        [Input("noIntervention")]
        public Input<bool>? NoIntervention { get; set; }

        /// <summary>
        /// Details about an Order
        /// </summary>
        [Input("order")]
        public Input<Inputs.ServerOrderGetArgs>? Order { get; set; }

        /// <summary>
        /// Operating system
        /// </summary>
        [Input("os")]
        public Input<string>? Os { get; set; }

        /// <summary>
        /// OVH subsidiaries
        /// </summary>
        [Input("ovhSubsidiary")]
        public Input<string>? OvhSubsidiary { get; set; }

        /// <summary>
        /// Partition scheme name
        /// </summary>
        [Input("partitionSchemeName")]
        public Input<string>? PartitionSchemeName { get; set; }

        [Input("planOptions")]
        private InputList<Inputs.ServerPlanOptionGetArgs>? _planOptions;
        public InputList<Inputs.ServerPlanOptionGetArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.ServerPlanOptionGetArgs>());
            set => _planOptions = value;
        }

        [Input("plans")]
        private InputList<Inputs.ServerPlanGetArgs>? _plans;
        public InputList<Inputs.ServerPlanGetArgs> Plans
        {
            get => _plans ?? (_plans = new InputList<Inputs.ServerPlanGetArgs>());
            set => _plans = value;
        }

        /// <summary>
        /// Power state of the server (poweron, poweroff)
        /// </summary>
        [Input("powerState")]
        public Input<string>? PowerState { get; set; }

        /// <summary>
        /// Does this server have professional use option
        /// </summary>
        [Input("professionalUse")]
        public Input<bool>? ProfessionalUse { get; set; }

        /// <summary>
        /// Rack id of the server
        /// </summary>
        [Input("rack")]
        public Input<string>? Rack { get; set; }

        /// <summary>
        /// Dedicated region localisation
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Rescue mail of the server
        /// </summary>
        [Input("rescueMail")]
        public Input<string>? RescueMail { get; set; }

        /// <summary>
        /// Public SSH Key used in the rescue mode
        /// </summary>
        [Input("rescueSshKey")]
        public Input<string>? RescueSshKey { get; set; }

        /// <summary>
        /// Dedicated server reverse
        /// </summary>
        [Input("reverse")]
        public Input<string>? Reverse { get; set; }

        /// <summary>
        /// Root device of the server
        /// </summary>
        [Input("rootDevice")]
        public Input<string>? RootDevice { get; set; }

        /// <summary>
        /// Server id
        /// </summary>
        [Input("serverId")]
        public Input<double>? ServerId { get; set; }

        /// <summary>
        /// The service_name of your dedicated server
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// All states a Dedicated can be in (error, hacked, hackedBlocked, ok)
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Dedicated server support level (critical, fastpath, gs, pro)
        /// </summary>
        [Input("supportLevel")]
        public Input<string>? SupportLevel { get; set; }

        /// <summary>
        /// Template name
        /// </summary>
        [Input("templateName")]
        public Input<string>? TemplateName { get; set; }

        [Input("userMetadatas")]
        private InputList<Inputs.ServerUserMetadataGetArgs>? _userMetadatas;

        /// <summary>
        /// Metadata
        /// </summary>
        public InputList<Inputs.ServerUserMetadataGetArgs> UserMetadatas
        {
            get => _userMetadatas ?? (_userMetadatas = new InputList<Inputs.ServerUserMetadataGetArgs>());
            set => _userMetadatas = value;
        }

        public ServerState()
        {
        }
        public static new ServerState Empty => new ServerState();
    }
}
