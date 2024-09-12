// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Dedicated;

import com.ovhcloud.ovh.Dedicated.ServerArgs;
import com.ovhcloud.ovh.Dedicated.inputs.ServerState;
import com.ovhcloud.ovh.Dedicated.outputs.ServerDetails;
import com.ovhcloud.ovh.Dedicated.outputs.ServerIam;
import com.ovhcloud.ovh.Dedicated.outputs.ServerOrder;
import com.ovhcloud.ovh.Dedicated.outputs.ServerPlan;
import com.ovhcloud.ovh.Dedicated.outputs.ServerPlanOption;
import com.ovhcloud.ovh.Dedicated.outputs.ServerUserMetadata;
import com.ovhcloud.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Dedicated servers can be imported using the `service_name`, e.g.:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:Dedicated/server:Server server service_name
 * ```
 * 
 */
@ResourceType(type="ovh:Dedicated/server:Server")
public class Server extends com.pulumi.resources.CustomResource {
    /**
     * Dedicated AZ localisation
     * 
     */
    @Export(name="availabilityZone", refs={String.class}, tree="[0]")
    private Output<String> availabilityZone;

    /**
     * @return Dedicated AZ localisation
     * 
     */
    public Output<String> availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * Boot id of the server
     * 
     */
    @Export(name="bootId", refs={Double.class}, tree="[0]")
    private Output<Double> bootId;

    /**
     * @return Boot id of the server
     * 
     */
    public Output<Double> bootId() {
        return this.bootId;
    }
    /**
     * Boot script of the server
     * 
     */
    @Export(name="bootScript", refs={String.class}, tree="[0]")
    private Output<String> bootScript;

    /**
     * @return Boot script of the server
     * 
     */
    public Output<String> bootScript() {
        return this.bootScript;
    }
    /**
     * Dedicated server commercial range
     * 
     */
    @Export(name="commercialRange", refs={String.class}, tree="[0]")
    private Output<String> commercialRange;

    /**
     * @return Dedicated server commercial range
     * 
     */
    public Output<String> commercialRange() {
        return this.commercialRange;
    }
    /**
     * Dedicated datacenter localisation (bhs1,bhs2,...)
     * 
     */
    @Export(name="datacenter", refs={String.class}, tree="[0]")
    private Output<String> datacenter;

    /**
     * @return Dedicated datacenter localisation (bhs1,bhs2,...)
     * 
     */
    public Output<String> datacenter() {
        return this.datacenter;
    }
    /**
     * A structure describing informations about installation custom
     * 
     */
    @Export(name="details", refs={ServerDetails.class}, tree="[0]")
    private Output</* @Nullable */ ServerDetails> details;

    /**
     * @return A structure describing informations about installation custom
     * 
     */
    public Output<Optional<ServerDetails>> details() {
        return Codegen.optional(this.details);
    }
    /**
     * Resource display name
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Resource display name
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * IAM resource information
     * 
     */
    @Export(name="iam", refs={ServerIam.class}, tree="[0]")
    private Output<ServerIam> iam;

    /**
     * @return IAM resource information
     * 
     */
    public Output<ServerIam> iam() {
        return this.iam;
    }
    /**
     * Dedicated server ip (IPv4)
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return Dedicated server ip (IPv4)
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * Link speed of the server
     * 
     */
    @Export(name="linkSpeed", refs={Double.class}, tree="[0]")
    private Output<Double> linkSpeed;

    /**
     * @return Link speed of the server
     * 
     */
    public Output<Double> linkSpeed() {
        return this.linkSpeed;
    }
    /**
     * Icmp monitoring state
     * 
     */
    @Export(name="monitoring", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> monitoring;

    /**
     * @return Icmp monitoring state
     * 
     */
    public Output<Boolean> monitoring() {
        return this.monitoring;
    }
    /**
     * Dedicated server name
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Dedicated server name
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="newUpgradeSystem", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> newUpgradeSystem;

    public Output<Boolean> newUpgradeSystem() {
        return this.newUpgradeSystem;
    }
    /**
     * Prevent datacenter intervention
     * 
     */
    @Export(name="noIntervention", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> noIntervention;

    /**
     * @return Prevent datacenter intervention
     * 
     */
    public Output<Boolean> noIntervention() {
        return this.noIntervention;
    }
    /**
     * Details about an Order
     * 
     */
    @Export(name="order", refs={ServerOrder.class}, tree="[0]")
    private Output<ServerOrder> order;

    /**
     * @return Details about an Order
     * 
     */
    public Output<ServerOrder> order() {
        return this.order;
    }
    /**
     * Operating system
     * 
     */
    @Export(name="os", refs={String.class}, tree="[0]")
    private Output<String> os;

    /**
     * @return Operating system
     * 
     */
    public Output<String> os() {
        return this.os;
    }
    /**
     * OVH subsidiaries
     * 
     */
    @Export(name="ovhSubsidiary", refs={String.class}, tree="[0]")
    private Output<String> ovhSubsidiary;

    /**
     * @return OVH subsidiaries
     * 
     */
    public Output<String> ovhSubsidiary() {
        return this.ovhSubsidiary;
    }
    /**
     * Partition scheme name
     * 
     */
    @Export(name="partitionSchemeName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> partitionSchemeName;

    /**
     * @return Partition scheme name
     * 
     */
    public Output<Optional<String>> partitionSchemeName() {
        return Codegen.optional(this.partitionSchemeName);
    }
    @Export(name="planOptions", refs={List.class,ServerPlanOption.class}, tree="[0,1]")
    private Output<List<ServerPlanOption>> planOptions;

    public Output<List<ServerPlanOption>> planOptions() {
        return this.planOptions;
    }
    @Export(name="plans", refs={List.class,ServerPlan.class}, tree="[0,1]")
    private Output<List<ServerPlan>> plans;

    public Output<List<ServerPlan>> plans() {
        return this.plans;
    }
    /**
     * Power state of the server (poweron, poweroff)
     * 
     */
    @Export(name="powerState", refs={String.class}, tree="[0]")
    private Output<String> powerState;

    /**
     * @return Power state of the server (poweron, poweroff)
     * 
     */
    public Output<String> powerState() {
        return this.powerState;
    }
    /**
     * Does this server have professional use option
     * 
     */
    @Export(name="professionalUse", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> professionalUse;

    /**
     * @return Does this server have professional use option
     * 
     */
    public Output<Boolean> professionalUse() {
        return this.professionalUse;
    }
    /**
     * Rack id of the server
     * 
     */
    @Export(name="rack", refs={String.class}, tree="[0]")
    private Output<String> rack;

    /**
     * @return Rack id of the server
     * 
     */
    public Output<String> rack() {
        return this.rack;
    }
    /**
     * Dedicated region localisation
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Dedicated region localisation
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Rescue mail of the server
     * 
     */
    @Export(name="rescueMail", refs={String.class}, tree="[0]")
    private Output<String> rescueMail;

    /**
     * @return Rescue mail of the server
     * 
     */
    public Output<String> rescueMail() {
        return this.rescueMail;
    }
    /**
     * Public SSH Key used in the rescue mode
     * 
     */
    @Export(name="rescueSshKey", refs={String.class}, tree="[0]")
    private Output<String> rescueSshKey;

    /**
     * @return Public SSH Key used in the rescue mode
     * 
     */
    public Output<String> rescueSshKey() {
        return this.rescueSshKey;
    }
    /**
     * Dedicated server reverse
     * 
     */
    @Export(name="reverse", refs={String.class}, tree="[0]")
    private Output<String> reverse;

    /**
     * @return Dedicated server reverse
     * 
     */
    public Output<String> reverse() {
        return this.reverse;
    }
    /**
     * Root device of the server
     * 
     */
    @Export(name="rootDevice", refs={String.class}, tree="[0]")
    private Output<String> rootDevice;

    /**
     * @return Root device of the server
     * 
     */
    public Output<String> rootDevice() {
        return this.rootDevice;
    }
    /**
     * Server id
     * 
     */
    @Export(name="serverId", refs={Double.class}, tree="[0]")
    private Output<Double> serverId;

    /**
     * @return Server id
     * 
     */
    public Output<Double> serverId() {
        return this.serverId;
    }
    /**
     * The service_name of your dedicated server
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The service_name of your dedicated server
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * All states a Dedicated can be in (error, hacked, hackedBlocked, ok)
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return All states a Dedicated can be in (error, hacked, hackedBlocked, ok)
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Dedicated server support level (critical, fastpath, gs, pro)
     * 
     */
    @Export(name="supportLevel", refs={String.class}, tree="[0]")
    private Output<String> supportLevel;

    /**
     * @return Dedicated server support level (critical, fastpath, gs, pro)
     * 
     */
    public Output<String> supportLevel() {
        return this.supportLevel;
    }
    /**
     * Template name
     * 
     */
    @Export(name="templateName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> templateName;

    /**
     * @return Template name
     * 
     */
    public Output<Optional<String>> templateName() {
        return Codegen.optional(this.templateName);
    }
    /**
     * Metadata
     * 
     */
    @Export(name="userMetadatas", refs={List.class,ServerUserMetadata.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ServerUserMetadata>> userMetadatas;

    /**
     * @return Metadata
     * 
     */
    public Output<Optional<List<ServerUserMetadata>>> userMetadatas() {
        return Codegen.optional(this.userMetadatas);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Server(java.lang.String name) {
        this(name, ServerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Server(java.lang.String name, ServerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Server(java.lang.String name, ServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/server:Server", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Server(java.lang.String name, Output<java.lang.String> id, @Nullable ServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/server:Server", name, state, makeResourceOptions(options, id), false);
    }

    private static ServerArgs makeArgs(ServerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServerArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Server get(java.lang.String name, Output<java.lang.String> id, @Nullable ServerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Server(name, id, state, options);
    }
}
