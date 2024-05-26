// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dbaas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetLogsClustersResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String serviceName;
    private String urn;
    /**
     * @return is the cluster id
     * 
     */
    private List<String> uuids;

    private GetLogsClustersResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String serviceName() {
        return this.serviceName;
    }
    public String urn() {
        return this.urn;
    }
    /**
     * @return is the cluster id
     * 
     */
    public List<String> uuids() {
        return this.uuids;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLogsClustersResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String serviceName;
        private String urn;
        private List<String> uuids;
        public Builder() {}
        public Builder(GetLogsClustersResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.serviceName = defaults.serviceName;
    	      this.urn = defaults.urn;
    	      this.uuids = defaults.uuids;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetLogsClustersResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetLogsClustersResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder urn(String urn) {
            if (urn == null) {
              throw new MissingRequiredPropertyException("GetLogsClustersResult", "urn");
            }
            this.urn = urn;
            return this;
        }
        @CustomType.Setter
        public Builder uuids(List<String> uuids) {
            if (uuids == null) {
              throw new MissingRequiredPropertyException("GetLogsClustersResult", "uuids");
            }
            this.uuids = uuids;
            return this;
        }
        public Builder uuids(String... uuids) {
            return uuids(List.of(uuids));
        }
        public GetLogsClustersResult build() {
            final var _resultValue = new GetLogsClustersResult();
            _resultValue.id = id;
            _resultValue.serviceName = serviceName;
            _resultValue.urn = urn;
            _resultValue.uuids = uuids;
            return _resultValue;
        }
    }
}
