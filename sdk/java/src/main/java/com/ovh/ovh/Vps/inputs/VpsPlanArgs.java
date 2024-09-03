// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Vps.inputs;

import com.ovh.ovh.Vps.inputs.VpsPlanConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpsPlanArgs extends com.pulumi.resources.ResourceArgs {

    public static final VpsPlanArgs Empty = new VpsPlanArgs();

    /**
     * Representation of a configuration item for personalizing product
     * 
     */
    @Import(name="configurations")
    private @Nullable Output<List<VpsPlanConfigurationArgs>> configurations;

    /**
     * @return Representation of a configuration item for personalizing product
     * 
     */
    public Optional<Output<List<VpsPlanConfigurationArgs>>> configurations() {
        return Optional.ofNullable(this.configurations);
    }

    /**
     * duration
     * 
     */
    @Import(name="duration", required=true)
    private Output<String> duration;

    /**
     * @return duration
     * 
     */
    public Output<String> duration() {
        return this.duration;
    }

    /**
     * Cart item to be linked
     * 
     */
    @Import(name="itemId")
    private @Nullable Output<Double> itemId;

    /**
     * @return Cart item to be linked
     * 
     */
    public Optional<Output<Double>> itemId() {
        return Optional.ofNullable(this.itemId);
    }

    /**
     * Plan code
     * 
     */
    @Import(name="planCode", required=true)
    private Output<String> planCode;

    /**
     * @return Plan code
     * 
     */
    public Output<String> planCode() {
        return this.planCode;
    }

    /**
     * Pricing model identifier
     * 
     */
    @Import(name="pricingMode", required=true)
    private Output<String> pricingMode;

    /**
     * @return Pricing model identifier
     * 
     */
    public Output<String> pricingMode() {
        return this.pricingMode;
    }

    /**
     * Quantity of product desired
     * 
     */
    @Import(name="quantity")
    private @Nullable Output<Double> quantity;

    /**
     * @return Quantity of product desired
     * 
     */
    public Optional<Output<Double>> quantity() {
        return Optional.ofNullable(this.quantity);
    }

    private VpsPlanArgs() {}

    private VpsPlanArgs(VpsPlanArgs $) {
        this.configurations = $.configurations;
        this.duration = $.duration;
        this.itemId = $.itemId;
        this.planCode = $.planCode;
        this.pricingMode = $.pricingMode;
        this.quantity = $.quantity;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpsPlanArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpsPlanArgs $;

        public Builder() {
            $ = new VpsPlanArgs();
        }

        public Builder(VpsPlanArgs defaults) {
            $ = new VpsPlanArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configurations Representation of a configuration item for personalizing product
         * 
         * @return builder
         * 
         */
        public Builder configurations(@Nullable Output<List<VpsPlanConfigurationArgs>> configurations) {
            $.configurations = configurations;
            return this;
        }

        /**
         * @param configurations Representation of a configuration item for personalizing product
         * 
         * @return builder
         * 
         */
        public Builder configurations(List<VpsPlanConfigurationArgs> configurations) {
            return configurations(Output.of(configurations));
        }

        /**
         * @param configurations Representation of a configuration item for personalizing product
         * 
         * @return builder
         * 
         */
        public Builder configurations(VpsPlanConfigurationArgs... configurations) {
            return configurations(List.of(configurations));
        }

        /**
         * @param duration duration
         * 
         * @return builder
         * 
         */
        public Builder duration(Output<String> duration) {
            $.duration = duration;
            return this;
        }

        /**
         * @param duration duration
         * 
         * @return builder
         * 
         */
        public Builder duration(String duration) {
            return duration(Output.of(duration));
        }

        /**
         * @param itemId Cart item to be linked
         * 
         * @return builder
         * 
         */
        public Builder itemId(@Nullable Output<Double> itemId) {
            $.itemId = itemId;
            return this;
        }

        /**
         * @param itemId Cart item to be linked
         * 
         * @return builder
         * 
         */
        public Builder itemId(Double itemId) {
            return itemId(Output.of(itemId));
        }

        /**
         * @param planCode Plan code
         * 
         * @return builder
         * 
         */
        public Builder planCode(Output<String> planCode) {
            $.planCode = planCode;
            return this;
        }

        /**
         * @param planCode Plan code
         * 
         * @return builder
         * 
         */
        public Builder planCode(String planCode) {
            return planCode(Output.of(planCode));
        }

        /**
         * @param pricingMode Pricing model identifier
         * 
         * @return builder
         * 
         */
        public Builder pricingMode(Output<String> pricingMode) {
            $.pricingMode = pricingMode;
            return this;
        }

        /**
         * @param pricingMode Pricing model identifier
         * 
         * @return builder
         * 
         */
        public Builder pricingMode(String pricingMode) {
            return pricingMode(Output.of(pricingMode));
        }

        /**
         * @param quantity Quantity of product desired
         * 
         * @return builder
         * 
         */
        public Builder quantity(@Nullable Output<Double> quantity) {
            $.quantity = quantity;
            return this;
        }

        /**
         * @param quantity Quantity of product desired
         * 
         * @return builder
         * 
         */
        public Builder quantity(Double quantity) {
            return quantity(Output.of(quantity));
        }

        public VpsPlanArgs build() {
            if ($.duration == null) {
                throw new MissingRequiredPropertyException("VpsPlanArgs", "duration");
            }
            if ($.planCode == null) {
                throw new MissingRequiredPropertyException("VpsPlanArgs", "planCode");
            }
            if ($.pricingMode == null) {
                throw new MissingRequiredPropertyException("VpsPlanArgs", "pricingMode");
            }
            return $;
        }
    }

}