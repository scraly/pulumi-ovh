// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class KubeKubeconfigAttribute {
    /**
     * @return The kubernetes API server client certificate.
     * 
     */
    private @Nullable String clientCertificate;
    /**
     * @return The kubernetes API server client key.
     * 
     */
    private @Nullable String clientKey;
    /**
     * @return The kubernetes API server CA certificate.
     * 
     */
    private @Nullable String clusterCaCertificate;
    /**
     * @return The kubernetes API server URL.
     * 
     */
    private @Nullable String host;

    private KubeKubeconfigAttribute() {}
    /**
     * @return The kubernetes API server client certificate.
     * 
     */
    public Optional<String> clientCertificate() {
        return Optional.ofNullable(this.clientCertificate);
    }
    /**
     * @return The kubernetes API server client key.
     * 
     */
    public Optional<String> clientKey() {
        return Optional.ofNullable(this.clientKey);
    }
    /**
     * @return The kubernetes API server CA certificate.
     * 
     */
    public Optional<String> clusterCaCertificate() {
        return Optional.ofNullable(this.clusterCaCertificate);
    }
    /**
     * @return The kubernetes API server URL.
     * 
     */
    public Optional<String> host() {
        return Optional.ofNullable(this.host);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KubeKubeconfigAttribute defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String clientCertificate;
        private @Nullable String clientKey;
        private @Nullable String clusterCaCertificate;
        private @Nullable String host;
        public Builder() {}
        public Builder(KubeKubeconfigAttribute defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientCertificate = defaults.clientCertificate;
    	      this.clientKey = defaults.clientKey;
    	      this.clusterCaCertificate = defaults.clusterCaCertificate;
    	      this.host = defaults.host;
        }

        @CustomType.Setter
        public Builder clientCertificate(@Nullable String clientCertificate) {

            this.clientCertificate = clientCertificate;
            return this;
        }
        @CustomType.Setter
        public Builder clientKey(@Nullable String clientKey) {

            this.clientKey = clientKey;
            return this;
        }
        @CustomType.Setter
        public Builder clusterCaCertificate(@Nullable String clusterCaCertificate) {

            this.clusterCaCertificate = clusterCaCertificate;
            return this;
        }
        @CustomType.Setter
        public Builder host(@Nullable String host) {

            this.host = host;
            return this;
        }
        public KubeKubeconfigAttribute build() {
            final var _resultValue = new KubeKubeconfigAttribute();
            _resultValue.clientCertificate = clientCertificate;
            _resultValue.clientKey = clientKey;
            _resultValue.clusterCaCertificate = clusterCaCertificate;
            _resultValue.host = host;
            return _resultValue;
        }
    }
}
