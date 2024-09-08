// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetDatabaseInstancesResult {
    /**
     * @return See Argument Reference above.
     * 
     */
    private String clusterId;
    /**
     * @return The list of databases ids of the database cluster associated with the project.
     * 
     */
    private List<String> databaseIds;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String engine;
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

    private GetDatabaseInstancesResult() {}
    /**
     * @return See Argument Reference above.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return The list of databases ids of the database cluster associated with the project.
     * 
     */
    public List<String> databaseIds() {
        return this.databaseIds;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String engine() {
        return this.engine;
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

    public static Builder builder(GetDatabaseInstancesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private List<String> databaseIds;
        private String engine;
        private String id;
        private String serviceName;
        public Builder() {}
        public Builder(GetDatabaseInstancesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.databaseIds = defaults.databaseIds;
    	      this.engine = defaults.engine;
    	      this.id = defaults.id;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            if (clusterId == null) {
              throw new MissingRequiredPropertyException("GetDatabaseInstancesResult", "clusterId");
            }
            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder databaseIds(List<String> databaseIds) {
            if (databaseIds == null) {
              throw new MissingRequiredPropertyException("GetDatabaseInstancesResult", "databaseIds");
            }
            this.databaseIds = databaseIds;
            return this;
        }
        public Builder databaseIds(String... databaseIds) {
            return databaseIds(List.of(databaseIds));
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            if (engine == null) {
              throw new MissingRequiredPropertyException("GetDatabaseInstancesResult", "engine");
            }
            this.engine = engine;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDatabaseInstancesResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetDatabaseInstancesResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetDatabaseInstancesResult build() {
            final var _resultValue = new GetDatabaseInstancesResult();
            _resultValue.clusterId = clusterId;
            _resultValue.databaseIds = databaseIds;
            _resultValue.engine = engine;
            _resultValue.id = id;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}