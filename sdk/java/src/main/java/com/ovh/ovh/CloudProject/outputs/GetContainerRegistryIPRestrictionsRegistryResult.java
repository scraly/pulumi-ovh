// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetContainerRegistryIPRestrictionsRegistryResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return IP restrictions applied on artifact manager component.
     * 
     */
    private List<Map<String,Object>> ipRestrictions;
    /**
     * @return The ID of the Managed Private Registry.
     * 
     */
    private String registryId;
    /**
     * @return The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    private String serviceName;

    private GetContainerRegistryIPRestrictionsRegistryResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return IP restrictions applied on artifact manager component.
     * 
     */
    public List<Map<String,Object>> ipRestrictions() {
        return this.ipRestrictions;
    }
    /**
     * @return The ID of the Managed Private Registry.
     * 
     */
    public String registryId() {
        return this.registryId;
    }
    /**
     * @return The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetContainerRegistryIPRestrictionsRegistryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<Map<String,Object>> ipRestrictions;
        private String registryId;
        private String serviceName;
        public Builder() {}
        public Builder(GetContainerRegistryIPRestrictionsRegistryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ipRestrictions = defaults.ipRestrictions;
    	      this.registryId = defaults.registryId;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetContainerRegistryIPRestrictionsRegistryResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ipRestrictions(List<Map<String,Object>> ipRestrictions) {
            if (ipRestrictions == null) {
              throw new MissingRequiredPropertyException("GetContainerRegistryIPRestrictionsRegistryResult", "ipRestrictions");
            }
            this.ipRestrictions = ipRestrictions;
            return this;
        }
        @CustomType.Setter
        public Builder registryId(String registryId) {
            if (registryId == null) {
              throw new MissingRequiredPropertyException("GetContainerRegistryIPRestrictionsRegistryResult", "registryId");
            }
            this.registryId = registryId;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetContainerRegistryIPRestrictionsRegistryResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetContainerRegistryIPRestrictionsRegistryResult build() {
            final var _resultValue = new GetContainerRegistryIPRestrictionsRegistryResult();
            _resultValue.id = id;
            _resultValue.ipRestrictions = ipRestrictions;
            _resultValue.registryId = registryId;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
