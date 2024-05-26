// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dbaas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetLogsClusterResult {
    /**
     * @return is the URN of the DBaas logs instance
     * 
     */
    private String DBaasURN;
    /**
     * @return is allowed networks for ARCHIVE flow type
     * 
     */
    private List<String> archiveAllowedNetworks;
    private @Nullable String clusterId;
    /**
     * @return is type of cluster (DEDICATED, PRO or TRIAL)
     * 
     */
    private String clusterType;
    /**
     * @return is PEM for dedicated inputs
     * 
     */
    private String dedicatedInputPem;
    /**
     * @return is allowed networks for DIRECT_INPUT flow type
     * 
     */
    private List<String> directInputAllowedNetworks;
    /**
     * @return is PEM for direct inputs
     * 
     */
    private String directInputPem;
    /**
     * @return is cluster hostname hosting the tenant
     * 
     */
    private String hostname;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return is true if all content generated by given service will be placed on this cluster
     * 
     */
    private Boolean isDefault;
    /**
     * @return is true if given service can perform advanced operations on cluster
     * 
     */
    private Boolean isUnlocked;
    /**
     * @return is allowed networks for QUERY flow type
     * 
     */
    private List<String> queryAllowedNetworks;
    /**
     * @return is datacenter localization
     * 
     */
    private String region;
    private String serviceName;

    private GetLogsClusterResult() {}
    /**
     * @return is the URN of the DBaas logs instance
     * 
     */
    public String DBaasURN() {
        return this.DBaasURN;
    }
    /**
     * @return is allowed networks for ARCHIVE flow type
     * 
     */
    public List<String> archiveAllowedNetworks() {
        return this.archiveAllowedNetworks;
    }
    public Optional<String> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }
    /**
     * @return is type of cluster (DEDICATED, PRO or TRIAL)
     * 
     */
    public String clusterType() {
        return this.clusterType;
    }
    /**
     * @return is PEM for dedicated inputs
     * 
     */
    public String dedicatedInputPem() {
        return this.dedicatedInputPem;
    }
    /**
     * @return is allowed networks for DIRECT_INPUT flow type
     * 
     */
    public List<String> directInputAllowedNetworks() {
        return this.directInputAllowedNetworks;
    }
    /**
     * @return is PEM for direct inputs
     * 
     */
    public String directInputPem() {
        return this.directInputPem;
    }
    /**
     * @return is cluster hostname hosting the tenant
     * 
     */
    public String hostname() {
        return this.hostname;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return is true if all content generated by given service will be placed on this cluster
     * 
     */
    public Boolean isDefault() {
        return this.isDefault;
    }
    /**
     * @return is true if given service can perform advanced operations on cluster
     * 
     */
    public Boolean isUnlocked() {
        return this.isUnlocked;
    }
    /**
     * @return is allowed networks for QUERY flow type
     * 
     */
    public List<String> queryAllowedNetworks() {
        return this.queryAllowedNetworks;
    }
    /**
     * @return is datacenter localization
     * 
     */
    public String region() {
        return this.region;
    }
    public String serviceName() {
        return this.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLogsClusterResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String DBaasURN;
        private List<String> archiveAllowedNetworks;
        private @Nullable String clusterId;
        private String clusterType;
        private String dedicatedInputPem;
        private List<String> directInputAllowedNetworks;
        private String directInputPem;
        private String hostname;
        private String id;
        private Boolean isDefault;
        private Boolean isUnlocked;
        private List<String> queryAllowedNetworks;
        private String region;
        private String serviceName;
        public Builder() {}
        public Builder(GetLogsClusterResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.DBaasURN = defaults.DBaasURN;
    	      this.archiveAllowedNetworks = defaults.archiveAllowedNetworks;
    	      this.clusterId = defaults.clusterId;
    	      this.clusterType = defaults.clusterType;
    	      this.dedicatedInputPem = defaults.dedicatedInputPem;
    	      this.directInputAllowedNetworks = defaults.directInputAllowedNetworks;
    	      this.directInputPem = defaults.directInputPem;
    	      this.hostname = defaults.hostname;
    	      this.id = defaults.id;
    	      this.isDefault = defaults.isDefault;
    	      this.isUnlocked = defaults.isUnlocked;
    	      this.queryAllowedNetworks = defaults.queryAllowedNetworks;
    	      this.region = defaults.region;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder DBaasURN(String DBaasURN) {
            if (DBaasURN == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "DBaasURN");
            }
            this.DBaasURN = DBaasURN;
            return this;
        }
        @CustomType.Setter
        public Builder archiveAllowedNetworks(List<String> archiveAllowedNetworks) {
            if (archiveAllowedNetworks == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "archiveAllowedNetworks");
            }
            this.archiveAllowedNetworks = archiveAllowedNetworks;
            return this;
        }
        public Builder archiveAllowedNetworks(String... archiveAllowedNetworks) {
            return archiveAllowedNetworks(List.of(archiveAllowedNetworks));
        }
        @CustomType.Setter
        public Builder clusterId(@Nullable String clusterId) {

            this.clusterId = clusterId;
            return this;
        }
        @CustomType.Setter
        public Builder clusterType(String clusterType) {
            if (clusterType == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "clusterType");
            }
            this.clusterType = clusterType;
            return this;
        }
        @CustomType.Setter
        public Builder dedicatedInputPem(String dedicatedInputPem) {
            if (dedicatedInputPem == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "dedicatedInputPem");
            }
            this.dedicatedInputPem = dedicatedInputPem;
            return this;
        }
        @CustomType.Setter
        public Builder directInputAllowedNetworks(List<String> directInputAllowedNetworks) {
            if (directInputAllowedNetworks == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "directInputAllowedNetworks");
            }
            this.directInputAllowedNetworks = directInputAllowedNetworks;
            return this;
        }
        public Builder directInputAllowedNetworks(String... directInputAllowedNetworks) {
            return directInputAllowedNetworks(List.of(directInputAllowedNetworks));
        }
        @CustomType.Setter
        public Builder directInputPem(String directInputPem) {
            if (directInputPem == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "directInputPem");
            }
            this.directInputPem = directInputPem;
            return this;
        }
        @CustomType.Setter
        public Builder hostname(String hostname) {
            if (hostname == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "hostname");
            }
            this.hostname = hostname;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder isDefault(Boolean isDefault) {
            if (isDefault == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "isDefault");
            }
            this.isDefault = isDefault;
            return this;
        }
        @CustomType.Setter
        public Builder isUnlocked(Boolean isUnlocked) {
            if (isUnlocked == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "isUnlocked");
            }
            this.isUnlocked = isUnlocked;
            return this;
        }
        @CustomType.Setter
        public Builder queryAllowedNetworks(List<String> queryAllowedNetworks) {
            if (queryAllowedNetworks == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "queryAllowedNetworks");
            }
            this.queryAllowedNetworks = queryAllowedNetworks;
            return this;
        }
        public Builder queryAllowedNetworks(String... queryAllowedNetworks) {
            return queryAllowedNetworks(List.of(queryAllowedNetworks));
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetLogsClusterResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetLogsClusterResult build() {
            final var _resultValue = new GetLogsClusterResult();
            _resultValue.DBaasURN = DBaasURN;
            _resultValue.archiveAllowedNetworks = archiveAllowedNetworks;
            _resultValue.clusterId = clusterId;
            _resultValue.clusterType = clusterType;
            _resultValue.dedicatedInputPem = dedicatedInputPem;
            _resultValue.directInputAllowedNetworks = directInputAllowedNetworks;
            _resultValue.directInputPem = directInputPem;
            _resultValue.hostname = hostname;
            _resultValue.id = id;
            _resultValue.isDefault = isDefault;
            _resultValue.isUnlocked = isUnlocked;
            _resultValue.queryAllowedNetworks = queryAllowedNetworks;
            _resultValue.region = region;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
