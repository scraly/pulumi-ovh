// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetMongoDbUserResult {
    /**
     * @return See Argument Reference above.
     * 
     */
    private String clusterId;
    /**
     * @return Date of the creation of the user.
     * 
     */
    private String createdAt;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String name;
    /**
     * @return Roles the user belongs to
     * 
     */
    private List<String> roles;
    /**
     * @return Current status of the user.
     * 
     */
    private String serviceName;
    /**
     * @return Current status of the user.
     * 
     */
    private String status;

    private GetMongoDbUserResult() {}
    /**
     * @return See Argument Reference above.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return Date of the creation of the user.
     * 
     */
    public String createdAt() {
        return this.createdAt;
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
    public String name() {
        return this.name;
    }
    /**
     * @return Roles the user belongs to
     * 
     */
    public List<String> roles() {
        return this.roles;
    }
    /**
     * @return Current status of the user.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Current status of the user.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMongoDbUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String createdAt;
        private String id;
        private String name;
        private List<String> roles;
        private String serviceName;
        private String status;
        public Builder() {}
        public Builder(GetMongoDbUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.roles = defaults.roles;
    	      this.serviceName = defaults.serviceName;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetMongoDbUserResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetMongoDbUserResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetMongoDbUserResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetMongoDbUserResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder roles(List<String> roles) {
            if (roles == null) {
              throw new MissingRequiredPropertyException("GetMongoDbUserResult", "roles");
            }
            this.roles = roles;
            return this;
        }
        public Builder roles(String... roles) {
            return roles(List.of(roles));
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetMongoDbUserResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetMongoDbUserResult", "status");
            }
            this.status = status;
            return this;
        }
        public GetMongoDbUserResult build() {
            final var _resultValue = new GetMongoDbUserResult();
            _resultValue.clusterId = clusterId;
            _resultValue.createdAt = createdAt;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.roles = roles;
            _resultValue.serviceName = serviceName;
            _resultValue.status = status;
            return _resultValue;
        }
    }
}
