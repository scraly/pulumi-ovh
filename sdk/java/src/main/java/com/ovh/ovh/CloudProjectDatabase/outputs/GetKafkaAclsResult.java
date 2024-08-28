// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetKafkaAclsResult {
    /**
     * @return The list of ACLs ids of the kafka cluster associated with the project.
     * 
     */
    private List<String> aclIds;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String clusterId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;

    private GetKafkaAclsResult() {}
    /**
     * @return The list of ACLs ids of the kafka cluster associated with the project.
     * 
     */
    public List<String> aclIds() {
        return this.aclIds;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKafkaAclsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> aclIds;
        private String clusterId;
        private String id;
        private String serviceName;
        public Builder() {}
        public Builder(GetKafkaAclsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aclIds = defaults.aclIds;
    	      this.clusterId = defaults.clusterId;
    	      this.id = defaults.id;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder aclIds(List<String> aclIds) {
            if (aclIds == null) {
              throw new MissingRequiredPropertyException("GetKafkaAclsResult", "aclIds");
            }
            this.aclIds = aclIds;
            return this;
        }
        public Builder aclIds(String... aclIds) {
            return aclIds(List.of(aclIds));
        }
        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetKafkaAclsResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetKafkaAclsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetKafkaAclsResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetKafkaAclsResult build() {
            final var _resultValue = new GetKafkaAclsResult();
            _resultValue.aclIds = aclIds;
            _resultValue.clusterId = clusterId;
            _resultValue.id = id;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
