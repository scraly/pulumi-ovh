// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Domain.inputs;

import com.ovh.ovh.Domain.inputs.ZoneOrderArgs;
import com.ovh.ovh.Domain.inputs.ZonePlanArgs;
import com.ovh.ovh.Domain.inputs.ZonePlanOptionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ZoneState extends com.pulumi.resources.ResourceArgs {

    public static final ZoneState Empty = new ZoneState();

    @Import(name="ZoneURN")
    private @Nullable Output<String> ZoneURN;

    public Optional<Output<String>> ZoneURN() {
        return Optional.ofNullable(this.ZoneURN);
    }

    /**
     * Is DNSSEC supported by this zone
     * 
     */
    @Import(name="dnssecSupported")
    private @Nullable Output<Boolean> dnssecSupported;

    /**
     * @return Is DNSSEC supported by this zone
     * 
     */
    public Optional<Output<Boolean>> dnssecSupported() {
        return Optional.ofNullable(this.dnssecSupported);
    }

    /**
     * hasDnsAnycast flag of the DNS zone
     * 
     */
    @Import(name="hasDnsAnycast")
    private @Nullable Output<Boolean> hasDnsAnycast;

    /**
     * @return hasDnsAnycast flag of the DNS zone
     * 
     */
    public Optional<Output<Boolean>> hasDnsAnycast() {
        return Optional.ofNullable(this.hasDnsAnycast);
    }

    /**
     * Last update date of the DNS zone
     * 
     */
    @Import(name="lastUpdate")
    private @Nullable Output<String> lastUpdate;

    /**
     * @return Last update date of the DNS zone
     * 
     */
    public Optional<Output<String>> lastUpdate() {
        return Optional.ofNullable(this.lastUpdate);
    }

    /**
     * Zone name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Zone name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Name servers that host the DNS zone
     * 
     */
    @Import(name="nameServers")
    private @Nullable Output<List<String>> nameServers;

    /**
     * @return Name servers that host the DNS zone
     * 
     */
    public Optional<Output<List<String>>> nameServers() {
        return Optional.ofNullable(this.nameServers);
    }

    /**
     * Details about an Order
     * 
     */
    @Import(name="orders")
    private @Nullable Output<List<ZoneOrderArgs>> orders;

    /**
     * @return Details about an Order
     * 
     */
    public Optional<Output<List<ZoneOrderArgs>>> orders() {
        return Optional.ofNullable(this.orders);
    }

    /**
     * OVHcloud Subsidiary. Country of OVHcloud legal entity you&#39;ll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
     * 
     */
    @Import(name="ovhSubsidiary")
    private @Nullable Output<String> ovhSubsidiary;

    /**
     * @return OVHcloud Subsidiary. Country of OVHcloud legal entity you&#39;ll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
     * 
     */
    public Optional<Output<String>> ovhSubsidiary() {
        return Optional.ofNullable(this.ovhSubsidiary);
    }

    /**
     * Ovh payment mode
     * 
     * @deprecated
     * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     * 
     */
    @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
    @Import(name="paymentMean")
    private @Nullable Output<String> paymentMean;

    /**
     * @return Ovh payment mode
     * 
     * @deprecated
     * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     * 
     */
    @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
    public Optional<Output<String>> paymentMean() {
        return Optional.ofNullable(this.paymentMean);
    }

    /**
     * Product Plan to order
     * 
     */
    @Import(name="plan")
    private @Nullable Output<ZonePlanArgs> plan;

    /**
     * @return Product Plan to order
     * 
     */
    public Optional<Output<ZonePlanArgs>> plan() {
        return Optional.ofNullable(this.plan);
    }

    /**
     * Product Plan to order
     * 
     */
    @Import(name="planOptions")
    private @Nullable Output<List<ZonePlanOptionArgs>> planOptions;

    /**
     * @return Product Plan to order
     * 
     */
    public Optional<Output<List<ZonePlanOptionArgs>>> planOptions() {
        return Optional.ofNullable(this.planOptions);
    }

    private ZoneState() {}

    private ZoneState(ZoneState $) {
        this.ZoneURN = $.ZoneURN;
        this.dnssecSupported = $.dnssecSupported;
        this.hasDnsAnycast = $.hasDnsAnycast;
        this.lastUpdate = $.lastUpdate;
        this.name = $.name;
        this.nameServers = $.nameServers;
        this.orders = $.orders;
        this.ovhSubsidiary = $.ovhSubsidiary;
        this.paymentMean = $.paymentMean;
        this.plan = $.plan;
        this.planOptions = $.planOptions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ZoneState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ZoneState $;

        public Builder() {
            $ = new ZoneState();
        }

        public Builder(ZoneState defaults) {
            $ = new ZoneState(Objects.requireNonNull(defaults));
        }

        public Builder ZoneURN(@Nullable Output<String> ZoneURN) {
            $.ZoneURN = ZoneURN;
            return this;
        }

        public Builder ZoneURN(String ZoneURN) {
            return ZoneURN(Output.of(ZoneURN));
        }

        /**
         * @param dnssecSupported Is DNSSEC supported by this zone
         * 
         * @return builder
         * 
         */
        public Builder dnssecSupported(@Nullable Output<Boolean> dnssecSupported) {
            $.dnssecSupported = dnssecSupported;
            return this;
        }

        /**
         * @param dnssecSupported Is DNSSEC supported by this zone
         * 
         * @return builder
         * 
         */
        public Builder dnssecSupported(Boolean dnssecSupported) {
            return dnssecSupported(Output.of(dnssecSupported));
        }

        /**
         * @param hasDnsAnycast hasDnsAnycast flag of the DNS zone
         * 
         * @return builder
         * 
         */
        public Builder hasDnsAnycast(@Nullable Output<Boolean> hasDnsAnycast) {
            $.hasDnsAnycast = hasDnsAnycast;
            return this;
        }

        /**
         * @param hasDnsAnycast hasDnsAnycast flag of the DNS zone
         * 
         * @return builder
         * 
         */
        public Builder hasDnsAnycast(Boolean hasDnsAnycast) {
            return hasDnsAnycast(Output.of(hasDnsAnycast));
        }

        /**
         * @param lastUpdate Last update date of the DNS zone
         * 
         * @return builder
         * 
         */
        public Builder lastUpdate(@Nullable Output<String> lastUpdate) {
            $.lastUpdate = lastUpdate;
            return this;
        }

        /**
         * @param lastUpdate Last update date of the DNS zone
         * 
         * @return builder
         * 
         */
        public Builder lastUpdate(String lastUpdate) {
            return lastUpdate(Output.of(lastUpdate));
        }

        /**
         * @param name Zone name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Zone name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nameServers Name servers that host the DNS zone
         * 
         * @return builder
         * 
         */
        public Builder nameServers(@Nullable Output<List<String>> nameServers) {
            $.nameServers = nameServers;
            return this;
        }

        /**
         * @param nameServers Name servers that host the DNS zone
         * 
         * @return builder
         * 
         */
        public Builder nameServers(List<String> nameServers) {
            return nameServers(Output.of(nameServers));
        }

        /**
         * @param nameServers Name servers that host the DNS zone
         * 
         * @return builder
         * 
         */
        public Builder nameServers(String... nameServers) {
            return nameServers(List.of(nameServers));
        }

        /**
         * @param orders Details about an Order
         * 
         * @return builder
         * 
         */
        public Builder orders(@Nullable Output<List<ZoneOrderArgs>> orders) {
            $.orders = orders;
            return this;
        }

        /**
         * @param orders Details about an Order
         * 
         * @return builder
         * 
         */
        public Builder orders(List<ZoneOrderArgs> orders) {
            return orders(Output.of(orders));
        }

        /**
         * @param orders Details about an Order
         * 
         * @return builder
         * 
         */
        public Builder orders(ZoneOrderArgs... orders) {
            return orders(List.of(orders));
        }

        /**
         * @param ovhSubsidiary OVHcloud Subsidiary. Country of OVHcloud legal entity you&#39;ll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
         * 
         * @return builder
         * 
         */
        public Builder ovhSubsidiary(@Nullable Output<String> ovhSubsidiary) {
            $.ovhSubsidiary = ovhSubsidiary;
            return this;
        }

        /**
         * @param ovhSubsidiary OVHcloud Subsidiary. Country of OVHcloud legal entity you&#39;ll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
         * 
         * @return builder
         * 
         */
        public Builder ovhSubsidiary(String ovhSubsidiary) {
            return ovhSubsidiary(Output.of(ovhSubsidiary));
        }

        /**
         * @param paymentMean Ovh payment mode
         * 
         * @return builder
         * 
         * @deprecated
         * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
         * 
         */
        @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
        public Builder paymentMean(@Nullable Output<String> paymentMean) {
            $.paymentMean = paymentMean;
            return this;
        }

        /**
         * @param paymentMean Ovh payment mode
         * 
         * @return builder
         * 
         * @deprecated
         * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
         * 
         */
        @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
        public Builder paymentMean(String paymentMean) {
            return paymentMean(Output.of(paymentMean));
        }

        /**
         * @param plan Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder plan(@Nullable Output<ZonePlanArgs> plan) {
            $.plan = plan;
            return this;
        }

        /**
         * @param plan Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder plan(ZonePlanArgs plan) {
            return plan(Output.of(plan));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(@Nullable Output<List<ZonePlanOptionArgs>> planOptions) {
            $.planOptions = planOptions;
            return this;
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(List<ZonePlanOptionArgs> planOptions) {
            return planOptions(Output.of(planOptions));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(ZonePlanOptionArgs... planOptions) {
            return planOptions(List.of(planOptions));
        }

        public ZoneState build() {
            return $;
        }
    }

}
