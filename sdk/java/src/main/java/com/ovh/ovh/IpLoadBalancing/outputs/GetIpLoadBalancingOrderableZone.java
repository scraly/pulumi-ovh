// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.IpLoadBalancing.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIpLoadBalancingOrderableZone {
    /**
     * @return The zone three letter code
     * 
     */
    private String name;
    /**
     * @return The billing planCode for this zone
     * 
     */
    private String planCode;

    private GetIpLoadBalancingOrderableZone() {}
    /**
     * @return The zone three letter code
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The billing planCode for this zone
     * 
     */
    public String planCode() {
        return this.planCode;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpLoadBalancingOrderableZone defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String name;
        private String planCode;
        public Builder() {}
        public Builder(GetIpLoadBalancingOrderableZone defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.planCode = defaults.planCode;
        }

        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetIpLoadBalancingOrderableZone", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder planCode(String planCode) {
            if (planCode == null) {
              throw new MissingRequiredPropertyException("GetIpLoadBalancingOrderableZone", "planCode");
            }
            this.planCode = planCode;
            return this;
        }
        public GetIpLoadBalancingOrderableZone build() {
            final var _resultValue = new GetIpLoadBalancingOrderableZone();
            _resultValue.name = name;
            _resultValue.planCode = planCode;
            return _resultValue;
        }
    }
}