// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.util.Objects;

@CustomType
public final class GetCapabilitiesContainerFilterFeature {
    /**
     * @return Vulnerability scanning
     * 
     */
    private Boolean vulnerability;

    private GetCapabilitiesContainerFilterFeature() {}
    /**
     * @return Vulnerability scanning
     * 
     */
    public Boolean vulnerability() {
        return this.vulnerability;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCapabilitiesContainerFilterFeature defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean vulnerability;
        public Builder() {}
        public Builder(GetCapabilitiesContainerFilterFeature defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.vulnerability = defaults.vulnerability;
        }

        @CustomType.Setter
        public Builder vulnerability(Boolean vulnerability) {
            if (vulnerability == null) {
              throw new MissingRequiredPropertyException("GetCapabilitiesContainerFilterFeature", "vulnerability");
            }
            this.vulnerability = vulnerability;
            return this;
        }
        public GetCapabilitiesContainerFilterFeature build() {
            final var _resultValue = new GetCapabilitiesContainerFilterFeature();
            _resultValue.vulnerability = vulnerability;
            return _resultValue;
        }
    }
}