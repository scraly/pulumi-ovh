// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class DatabaseNode {
    /**
     * @return Private network id in which the node should be deployed. It&#39;s the regional openstackId of the private network
     * 
     */
    private @Nullable String networkId;
    /**
     * @return Public cloud region in which the node should be deployed.
     * Ex: &#34;GRA&#39;.
     * 
     */
    private String region;
    /**
     * @return Private subnet ID in which the node is.
     * 
     */
    private @Nullable String subnetId;

    private DatabaseNode() {}
    /**
     * @return Private network id in which the node should be deployed. It&#39;s the regional openstackId of the private network
     * 
     */
    public Optional<String> networkId() {
        return Optional.ofNullable(this.networkId);
    }
    /**
     * @return Public cloud region in which the node should be deployed.
     * Ex: &#34;GRA&#39;.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return Private subnet ID in which the node is.
     * 
     */
    public Optional<String> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(DatabaseNode defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String networkId;
        private String region;
        private @Nullable String subnetId;
        public Builder() {}
        public Builder(DatabaseNode defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.networkId = defaults.networkId;
    	      this.region = defaults.region;
    	      this.subnetId = defaults.subnetId;
        }

        @CustomType.Setter
        public Builder networkId(@Nullable String networkId) {

            this.networkId = networkId;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("DatabaseNode", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder subnetId(@Nullable String subnetId) {

            this.subnetId = subnetId;
            return this;
        }
        public DatabaseNode build() {
            final var _resultValue = new DatabaseNode();
            _resultValue.networkId = networkId;
            _resultValue.region = region;
            _resultValue.subnetId = subnetId;
            return _resultValue;
        }
    }
}