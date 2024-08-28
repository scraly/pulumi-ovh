// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Domain;

import com.ovh.ovh.Domain.ZoneArgs;
import com.ovh.ovh.Domain.inputs.ZoneState;
import com.ovh.ovh.Domain.outputs.ZoneOrder;
import com.ovh.ovh.Domain.outputs.ZonePlan;
import com.ovh.ovh.Domain.outputs.ZonePlanOption;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Me.MeFunctions;
 * import com.pulumi.ovh.Order.OrderFunctions;
 * import com.pulumi.ovh.Order.inputs.GetCartArgs;
 * import com.pulumi.ovh.Order.inputs.GetCartProductPlanArgs;
 * import com.pulumi.ovh.Domain.Zone;
 * import com.pulumi.ovh.Domain.ZoneArgs;
 * import com.pulumi.ovh.Domain.inputs.ZonePlanArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var myaccount = MeFunctions.getMe();
 * 
 *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
 *             .ovhSubsidiary(myaccount.applyValue(getMeResult -> getMeResult.ovhSubsidiary()))
 *             .build());
 * 
 *         final var zoneCartProductPlan = OrderFunctions.getCartProductPlan(GetCartProductPlanArgs.builder()
 *             .cartId(mycart.applyValue(getCartResult -> getCartResult.id()))
 *             .priceCapacity("renew")
 *             .product("dns")
 *             .planCode("zone")
 *             .build());
 * 
 *         var zoneZone = new Zone("zoneZone", ZoneArgs.builder()        
 *             .ovhSubsidiary(mycart.applyValue(getCartResult -> getCartResult.ovhSubsidiary()))
 *             .plan(ZonePlanArgs.builder()
 *                 .duration(zoneCartProductPlan.applyValue(getCartProductPlanResult -> getCartProductPlanResult.selectedPrices()[0].duration()))
 *                 .planCode(zoneCartProductPlan.applyValue(getCartProductPlanResult -> getCartProductPlanResult.planCode()))
 *                 .pricingMode(zoneCartProductPlan.applyValue(getCartProductPlanResult -> getCartProductPlanResult.selectedPrices()[0].pricingMode()))
 *                 .configurations(                
 *                     ZonePlanConfigurationArgs.builder()
 *                         .label("zone")
 *                         .value("myzone.mydomain.com")
 *                         .build(),
 *                     ZonePlanConfigurationArgs.builder()
 *                         .label("template")
 *                         .value("minimized")
 *                         .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * 
 * ## Import
 * 
 * Zone can be imported using the `order_id` that can be retrieved in the [order page](https://www.ovh.com/manager/#/dedicated/billing/orders/orders) at the creation time of the zone.
 * 
 * bash
 * 
 * ```sh
 *  $ pulumi import ovh:Domain/zone:Zone zone order_id
 * ```
 * 
 */
@ResourceType(type="ovh:Domain/zone:Zone")
public class Zone extends com.pulumi.resources.CustomResource {
    @Export(name="ZoneURN", refs={String.class}, tree="[0]")
    private Output<String> ZoneURN;

    public Output<String> ZoneURN() {
        return this.ZoneURN;
    }
    /**
     * Is DNSSEC supported by this zone
     * 
     */
    @Export(name="dnssecSupported", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> dnssecSupported;

    /**
     * @return Is DNSSEC supported by this zone
     * 
     */
    public Output<Boolean> dnssecSupported() {
        return this.dnssecSupported;
    }
    /**
     * hasDnsAnycast flag of the DNS zone
     * 
     */
    @Export(name="hasDnsAnycast", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> hasDnsAnycast;

    /**
     * @return hasDnsAnycast flag of the DNS zone
     * 
     */
    public Output<Boolean> hasDnsAnycast() {
        return this.hasDnsAnycast;
    }
    /**
     * Last update date of the DNS zone
     * 
     */
    @Export(name="lastUpdate", refs={String.class}, tree="[0]")
    private Output<String> lastUpdate;

    /**
     * @return Last update date of the DNS zone
     * 
     */
    public Output<String> lastUpdate() {
        return this.lastUpdate;
    }
    /**
     * Zone name
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Zone name
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Name servers that host the DNS zone
     * 
     */
    @Export(name="nameServers", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> nameServers;

    /**
     * @return Name servers that host the DNS zone
     * 
     */
    public Output<List<String>> nameServers() {
        return this.nameServers;
    }
    /**
     * Details about an Order
     * 
     */
    @Export(name="orders", refs={List.class,ZoneOrder.class}, tree="[0,1]")
    private Output<List<ZoneOrder>> orders;

    /**
     * @return Details about an Order
     * 
     */
    public Output<List<ZoneOrder>> orders() {
        return this.orders;
    }
    /**
     * OVHcloud Subsidiary. Country of OVHcloud legal entity you&#39;ll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
     * 
     */
    @Export(name="ovhSubsidiary", refs={String.class}, tree="[0]")
    private Output<String> ovhSubsidiary;

    /**
     * @return OVHcloud Subsidiary. Country of OVHcloud legal entity you&#39;ll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
     * 
     */
    public Output<String> ovhSubsidiary() {
        return this.ovhSubsidiary;
    }
    /**
     * Ovh payment mode
     * 
     * @deprecated
     * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     * 
     */
    @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
    @Export(name="paymentMean", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> paymentMean;

    /**
     * @return Ovh payment mode
     * 
     */
    public Output<Optional<String>> paymentMean() {
        return Codegen.optional(this.paymentMean);
    }
    /**
     * Product Plan to order
     * 
     */
    @Export(name="plan", refs={ZonePlan.class}, tree="[0]")
    private Output<ZonePlan> plan;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<ZonePlan> plan() {
        return this.plan;
    }
    /**
     * Product Plan to order
     * 
     */
    @Export(name="planOptions", refs={List.class,ZonePlanOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ZonePlanOption>> planOptions;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<Optional<List<ZonePlanOption>>> planOptions() {
        return Codegen.optional(this.planOptions);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Zone(String name) {
        this(name, ZoneArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Zone(String name, ZoneArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Zone(String name, ZoneArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Domain/zone:Zone", name, args == null ? ZoneArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Zone(String name, Output<String> id, @Nullable ZoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Domain/zone:Zone", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static Zone get(String name, Output<String> id, @Nullable ZoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Zone(name, id, state, options);
    }
}
