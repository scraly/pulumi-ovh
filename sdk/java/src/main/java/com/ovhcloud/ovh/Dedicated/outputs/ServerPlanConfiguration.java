// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Dedicated.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ServerPlanConfiguration {
    /**
     * @return Label for your configuration item
     * 
     */
    private String label;
    /**
     * @return Value or resource URL on API.OVH.COM of your configuration item
     * 
     */
    private String value;

    private ServerPlanConfiguration() {}
    /**
     * @return Label for your configuration item
     * 
     */
    public String label() {
        return this.label;
    }
    /**
     * @return Value or resource URL on API.OVH.COM of your configuration item
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServerPlanConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String label;
        private String value;
        public Builder() {}
        public Builder(ServerPlanConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.label = defaults.label;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder label(String label) {
            if (label == null) {
              throw new MissingRequiredPropertyException("ServerPlanConfiguration", "label");
            }
            this.label = label;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("ServerPlanConfiguration", "value");
            }
            this.value = value;
            return this;
        }
        public ServerPlanConfiguration build() {
            final var _resultValue = new ServerPlanConfiguration();
            _resultValue.label = label;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}
