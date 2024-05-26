// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.IpLoadBalancing.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class LoadBalancerPlanOptionConfiguration {
    /**
     * @return Identifier of the resource
     * 
     */
    private String label;
    /**
     * @return Path to the resource in API.OVH.COM
     * 
     */
    private String value;

    private LoadBalancerPlanOptionConfiguration() {}
    /**
     * @return Identifier of the resource
     * 
     */
    public String label() {
        return this.label;
    }
    /**
     * @return Path to the resource in API.OVH.COM
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LoadBalancerPlanOptionConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String label;
        private String value;
        public Builder() {}
        public Builder(LoadBalancerPlanOptionConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.label = defaults.label;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder label(String label) {
            if (label == null) {
              throw new MissingRequiredPropertyException("LoadBalancerPlanOptionConfiguration", "label");
            }
            this.label = label;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("LoadBalancerPlanOptionConfiguration", "value");
            }
            this.value = value;
            return this;
        }
        public LoadBalancerPlanOptionConfiguration build() {
            final var _resultValue = new LoadBalancerPlanOptionConfiguration();
            _resultValue.label = label;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}
