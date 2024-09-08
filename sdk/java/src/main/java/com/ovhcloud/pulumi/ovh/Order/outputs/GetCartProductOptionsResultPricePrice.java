// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Order.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCartProductOptionsResultPricePrice {
    /**
     * @return Currency code
     * 
     */
    private String currencyCode;
    /**
     * @return Textual representation
     * 
     */
    private String text;
    /**
     * @return The effective price
     * 
     */
    private Double value;

    private GetCartProductOptionsResultPricePrice() {}
    /**
     * @return Currency code
     * 
     */
    public String currencyCode() {
        return this.currencyCode;
    }
    /**
     * @return Textual representation
     * 
     */
    public String text() {
        return this.text;
    }
    /**
     * @return The effective price
     * 
     */
    public Double value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCartProductOptionsResultPricePrice defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String currencyCode;
        private String text;
        private Double value;
        public Builder() {}
        public Builder(GetCartProductOptionsResultPricePrice defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.currencyCode = defaults.currencyCode;
    	      this.text = defaults.text;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder currencyCode(String currencyCode) {
            if (currencyCode == null) {
              throw new MissingRequiredPropertyException("GetCartProductOptionsResultPricePrice", "currencyCode");
            }
            this.currencyCode = currencyCode;
            return this;
        }
        @CustomType.Setter
        public Builder text(String text) {
            if (text == null) {
              throw new MissingRequiredPropertyException("GetCartProductOptionsResultPricePrice", "text");
            }
            this.text = text;
            return this;
        }
        @CustomType.Setter
        public Builder value(Double value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("GetCartProductOptionsResultPricePrice", "value");
            }
            this.value = value;
            return this;
        }
        public GetCartProductOptionsResultPricePrice build() {
            final var _resultValue = new GetCartProductOptionsResultPricePrice();
            _resultValue.currencyCode = currencyCode;
            _resultValue.text = text;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}