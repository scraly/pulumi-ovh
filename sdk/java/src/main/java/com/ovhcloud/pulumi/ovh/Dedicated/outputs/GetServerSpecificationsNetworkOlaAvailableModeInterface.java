// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dedicated.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetServerSpecificationsNetworkOlaAvailableModeInterface {
    /**
     * @return Interface aggregation status
     * 
     */
    private Boolean aggregation;
    /**
     * @return Interface count
     * 
     */
    private Double count;
    /**
     * @return Bandwidth offer type (included┃standard)
     * 
     */
    private String type;

    private GetServerSpecificationsNetworkOlaAvailableModeInterface() {}
    /**
     * @return Interface aggregation status
     * 
     */
    public Boolean aggregation() {
        return this.aggregation;
    }
    /**
     * @return Interface count
     * 
     */
    public Double count() {
        return this.count;
    }
    /**
     * @return Bandwidth offer type (included┃standard)
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetServerSpecificationsNetworkOlaAvailableModeInterface defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean aggregation;
        private Double count;
        private String type;
        public Builder() {}
        public Builder(GetServerSpecificationsNetworkOlaAvailableModeInterface defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aggregation = defaults.aggregation;
    	      this.count = defaults.count;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder aggregation(Boolean aggregation) {
            if (aggregation == null) {
              throw new MissingRequiredPropertyException("GetServerSpecificationsNetworkOlaAvailableModeInterface", "aggregation");
            }
            this.aggregation = aggregation;
            return this;
        }
        @CustomType.Setter
        public Builder count(Double count) {
            if (count == null) {
              throw new MissingRequiredPropertyException("GetServerSpecificationsNetworkOlaAvailableModeInterface", "count");
            }
            this.count = count;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetServerSpecificationsNetworkOlaAvailableModeInterface", "type");
            }
            this.type = type;
            return this;
        }
        public GetServerSpecificationsNetworkOlaAvailableModeInterface build() {
            final var _resultValue = new GetServerSpecificationsNetworkOlaAvailableModeInterface();
            _resultValue.aggregation = aggregation;
            _resultValue.count = count;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}