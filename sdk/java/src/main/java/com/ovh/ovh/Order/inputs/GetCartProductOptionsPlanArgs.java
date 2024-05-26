// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Order.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetCartProductOptionsPlanArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetCartProductOptionsPlanArgs Empty = new GetCartProductOptionsPlanArgs();

    /**
     * Cart identifier
     * 
     */
    @Import(name="cartId", required=true)
    private Output<String> cartId;

    /**
     * @return Cart identifier
     * 
     */
    public Output<String> cartId() {
        return this.cartId;
    }

    /**
     * Catalog name
     * 
     */
    @Import(name="catalogName")
    private @Nullable Output<String> catalogName;

    /**
     * @return Catalog name
     * 
     */
    public Optional<Output<String>> catalogName() {
        return Optional.ofNullable(this.catalogName);
    }

    /**
     * options plan code.
     * 
     */
    @Import(name="optionsPlanCode", required=true)
    private Output<String> optionsPlanCode;

    /**
     * @return options plan code.
     * 
     */
    public Output<String> optionsPlanCode() {
        return this.optionsPlanCode;
    }

    /**
     * Product offer identifier
     * 
     */
    @Import(name="planCode", required=true)
    private Output<String> planCode;

    /**
     * @return Product offer identifier
     * 
     */
    public Output<String> planCode() {
        return this.planCode;
    }

    /**
     * Capacity of the pricing (type of pricing)
     * 
     */
    @Import(name="priceCapacity", required=true)
    private Output<String> priceCapacity;

    /**
     * @return Capacity of the pricing (type of pricing)
     * 
     */
    public Output<String> priceCapacity() {
        return this.priceCapacity;
    }

    /**
     * Product
     * 
     */
    @Import(name="product", required=true)
    private Output<String> product;

    /**
     * @return Product
     * 
     */
    public Output<String> product() {
        return this.product;
    }

    private GetCartProductOptionsPlanArgs() {}

    private GetCartProductOptionsPlanArgs(GetCartProductOptionsPlanArgs $) {
        this.cartId = $.cartId;
        this.catalogName = $.catalogName;
        this.optionsPlanCode = $.optionsPlanCode;
        this.planCode = $.planCode;
        this.priceCapacity = $.priceCapacity;
        this.product = $.product;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetCartProductOptionsPlanArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetCartProductOptionsPlanArgs $;

        public Builder() {
            $ = new GetCartProductOptionsPlanArgs();
        }

        public Builder(GetCartProductOptionsPlanArgs defaults) {
            $ = new GetCartProductOptionsPlanArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cartId Cart identifier
         * 
         * @return builder
         * 
         */
        public Builder cartId(Output<String> cartId) {
            $.cartId = cartId;
            return this;
        }

        /**
         * @param cartId Cart identifier
         * 
         * @return builder
         * 
         */
        public Builder cartId(String cartId) {
            return cartId(Output.of(cartId));
        }

        /**
         * @param catalogName Catalog name
         * 
         * @return builder
         * 
         */
        public Builder catalogName(@Nullable Output<String> catalogName) {
            $.catalogName = catalogName;
            return this;
        }

        /**
         * @param catalogName Catalog name
         * 
         * @return builder
         * 
         */
        public Builder catalogName(String catalogName) {
            return catalogName(Output.of(catalogName));
        }

        /**
         * @param optionsPlanCode options plan code.
         * 
         * @return builder
         * 
         */
        public Builder optionsPlanCode(Output<String> optionsPlanCode) {
            $.optionsPlanCode = optionsPlanCode;
            return this;
        }

        /**
         * @param optionsPlanCode options plan code.
         * 
         * @return builder
         * 
         */
        public Builder optionsPlanCode(String optionsPlanCode) {
            return optionsPlanCode(Output.of(optionsPlanCode));
        }

        /**
         * @param planCode Product offer identifier
         * 
         * @return builder
         * 
         */
        public Builder planCode(Output<String> planCode) {
            $.planCode = planCode;
            return this;
        }

        /**
         * @param planCode Product offer identifier
         * 
         * @return builder
         * 
         */
        public Builder planCode(String planCode) {
            return planCode(Output.of(planCode));
        }

        /**
         * @param priceCapacity Capacity of the pricing (type of pricing)
         * 
         * @return builder
         * 
         */
        public Builder priceCapacity(Output<String> priceCapacity) {
            $.priceCapacity = priceCapacity;
            return this;
        }

        /**
         * @param priceCapacity Capacity of the pricing (type of pricing)
         * 
         * @return builder
         * 
         */
        public Builder priceCapacity(String priceCapacity) {
            return priceCapacity(Output.of(priceCapacity));
        }

        /**
         * @param product Product
         * 
         * @return builder
         * 
         */
        public Builder product(Output<String> product) {
            $.product = product;
            return this;
        }

        /**
         * @param product Product
         * 
         * @return builder
         * 
         */
        public Builder product(String product) {
            return product(Output.of(product));
        }

        public GetCartProductOptionsPlanArgs build() {
            if ($.cartId == null) {
                throw new MissingRequiredPropertyException("GetCartProductOptionsPlanArgs", "cartId");
            }
            if ($.optionsPlanCode == null) {
                throw new MissingRequiredPropertyException("GetCartProductOptionsPlanArgs", "optionsPlanCode");
            }
            if ($.planCode == null) {
                throw new MissingRequiredPropertyException("GetCartProductOptionsPlanArgs", "planCode");
            }
            if ($.priceCapacity == null) {
                throw new MissingRequiredPropertyException("GetCartProductOptionsPlanArgs", "priceCapacity");
            }
            if ($.product == null) {
                throw new MissingRequiredPropertyException("GetCartProductOptionsPlanArgs", "product");
            }
            return $;
        }
    }

}
