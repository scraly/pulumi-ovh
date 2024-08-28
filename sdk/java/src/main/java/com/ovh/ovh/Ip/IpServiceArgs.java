// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Ip;

import com.ovh.ovh.Ip.inputs.IpServiceOrderArgs;
import com.ovh.ovh.Ip.inputs.IpServicePlanArgs;
import com.ovh.ovh.Ip.inputs.IpServicePlanOptionArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IpServiceArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpServiceArgs Empty = new IpServiceArgs();

    /**
     * Custom description on your ip.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Custom description on your ip.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Details about an Order
     * 
     */
    @Import(name="orders")
    private @Nullable Output<List<IpServiceOrderArgs>> orders;

    /**
     * @return Details about an Order
     * 
     */
    public Optional<Output<List<IpServiceOrderArgs>>> orders() {
        return Optional.ofNullable(this.orders);
    }

    /**
     * OVHcloud Subsidiary. Country of OVHcloud legal entity you&#39;ll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
     * 
     */
    @Import(name="ovhSubsidiary", required=true)
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
    @Import(name="plan", required=true)
    private Output<IpServicePlanArgs> plan;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<IpServicePlanArgs> plan() {
        return this.plan;
    }

    /**
     * Product Plan to order
     * 
     */
    @Import(name="planOptions")
    private @Nullable Output<List<IpServicePlanOptionArgs>> planOptions;

    /**
     * @return Product Plan to order
     * 
     */
    public Optional<Output<List<IpServicePlanOptionArgs>>> planOptions() {
        return Optional.ofNullable(this.planOptions);
    }

    private IpServiceArgs() {}

    private IpServiceArgs(IpServiceArgs $) {
        this.description = $.description;
        this.orders = $.orders;
        this.ovhSubsidiary = $.ovhSubsidiary;
        this.paymentMean = $.paymentMean;
        this.plan = $.plan;
        this.planOptions = $.planOptions;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpServiceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpServiceArgs $;

        public Builder() {
            $ = new IpServiceArgs();
        }

        public Builder(IpServiceArgs defaults) {
            $ = new IpServiceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Custom description on your ip.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Custom description on your ip.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param orders Details about an Order
         * 
         * @return builder
         * 
         */
        public Builder orders(@Nullable Output<List<IpServiceOrderArgs>> orders) {
            $.orders = orders;
            return this;
        }

        /**
         * @param orders Details about an Order
         * 
         * @return builder
         * 
         */
        public Builder orders(List<IpServiceOrderArgs> orders) {
            return orders(Output.of(orders));
        }

        /**
         * @param orders Details about an Order
         * 
         * @return builder
         * 
         */
        public Builder orders(IpServiceOrderArgs... orders) {
            return orders(List.of(orders));
        }

        /**
         * @param ovhSubsidiary OVHcloud Subsidiary. Country of OVHcloud legal entity you&#39;ll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
         * 
         * @return builder
         * 
         */
        public Builder ovhSubsidiary(Output<String> ovhSubsidiary) {
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
        public Builder plan(Output<IpServicePlanArgs> plan) {
            $.plan = plan;
            return this;
        }

        /**
         * @param plan Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder plan(IpServicePlanArgs plan) {
            return plan(Output.of(plan));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(@Nullable Output<List<IpServicePlanOptionArgs>> planOptions) {
            $.planOptions = planOptions;
            return this;
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(List<IpServicePlanOptionArgs> planOptions) {
            return planOptions(Output.of(planOptions));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(IpServicePlanOptionArgs... planOptions) {
            return planOptions(List.of(planOptions));
        }

        public IpServiceArgs build() {
            if ($.ovhSubsidiary == null) {
                throw new MissingRequiredPropertyException("IpServiceArgs", "ovhSubsidiary");
            }
            if ($.plan == null) {
                throw new MissingRequiredPropertyException("IpServiceArgs", "plan");
            }
            return $;
        }
    }

}
