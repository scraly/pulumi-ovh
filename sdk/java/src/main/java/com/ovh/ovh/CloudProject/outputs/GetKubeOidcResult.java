// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetKubeOidcResult {
    /**
     * @return The OIDC client ID.
     * 
     */
    private @Nullable String clientId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The OIDC issuer url.
     * 
     */
    private @Nullable String issuerUrl;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String kubeId;
    private @Nullable String oidcCaContent;
    private @Nullable List<String> oidcGroupsClaims;
    private @Nullable String oidcGroupsPrefix;
    private @Nullable List<String> oidcRequiredClaims;
    private @Nullable List<String> oidcSigningAlgs;
    private @Nullable String oidcUsernameClaim;
    private @Nullable String oidcUsernamePrefix;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;

    private GetKubeOidcResult() {}
    /**
     * @return The OIDC client ID.
     * 
     */
    public Optional<String> clientId() {
        return Optional.ofNullable(this.clientId);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The OIDC issuer url.
     * 
     */
    public Optional<String> issuerUrl() {
        return Optional.ofNullable(this.issuerUrl);
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String kubeId() {
        return this.kubeId;
    }
    public Optional<String> oidcCaContent() {
        return Optional.ofNullable(this.oidcCaContent);
    }
    public List<String> oidcGroupsClaims() {
        return this.oidcGroupsClaims == null ? List.of() : this.oidcGroupsClaims;
    }
    public Optional<String> oidcGroupsPrefix() {
        return Optional.ofNullable(this.oidcGroupsPrefix);
    }
    public List<String> oidcRequiredClaims() {
        return this.oidcRequiredClaims == null ? List.of() : this.oidcRequiredClaims;
    }
    public List<String> oidcSigningAlgs() {
        return this.oidcSigningAlgs == null ? List.of() : this.oidcSigningAlgs;
    }
    public Optional<String> oidcUsernameClaim() {
        return Optional.ofNullable(this.oidcUsernameClaim);
    }
    public Optional<String> oidcUsernamePrefix() {
        return Optional.ofNullable(this.oidcUsernamePrefix);
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

    public static Builder builder(GetKubeOidcResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String clientId;
        private String id;
        private @Nullable String issuerUrl;
        private String kubeId;
        private @Nullable String oidcCaContent;
        private @Nullable List<String> oidcGroupsClaims;
        private @Nullable String oidcGroupsPrefix;
        private @Nullable List<String> oidcRequiredClaims;
        private @Nullable List<String> oidcSigningAlgs;
        private @Nullable String oidcUsernameClaim;
        private @Nullable String oidcUsernamePrefix;
        private String serviceName;
        public Builder() {}
        public Builder(GetKubeOidcResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clientId = defaults.clientId;
    	      this.id = defaults.id;
    	      this.issuerUrl = defaults.issuerUrl;
    	      this.kubeId = defaults.kubeId;
    	      this.oidcCaContent = defaults.oidcCaContent;
    	      this.oidcGroupsClaims = defaults.oidcGroupsClaims;
    	      this.oidcGroupsPrefix = defaults.oidcGroupsPrefix;
    	      this.oidcRequiredClaims = defaults.oidcRequiredClaims;
    	      this.oidcSigningAlgs = defaults.oidcSigningAlgs;
    	      this.oidcUsernameClaim = defaults.oidcUsernameClaim;
    	      this.oidcUsernamePrefix = defaults.oidcUsernamePrefix;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder clientId(@Nullable String clientId) {

            this.clientId = clientId;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetKubeOidcResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder issuerUrl(@Nullable String issuerUrl) {

            this.issuerUrl = issuerUrl;
            return this;
        }
        @CustomType.Setter
        public Builder kubeId(String kubeId) {
            if (kubeId == null) {
              throw new MissingRequiredPropertyException("GetKubeOidcResult", "kubeId");
            }
            this.kubeId = kubeId;
            return this;
        }
        @CustomType.Setter
        public Builder oidcCaContent(@Nullable String oidcCaContent) {

            this.oidcCaContent = oidcCaContent;
            return this;
        }
        @CustomType.Setter
        public Builder oidcGroupsClaims(@Nullable List<String> oidcGroupsClaims) {

            this.oidcGroupsClaims = oidcGroupsClaims;
            return this;
        }
        public Builder oidcGroupsClaims(String... oidcGroupsClaims) {
            return oidcGroupsClaims(List.of(oidcGroupsClaims));
        }
        @CustomType.Setter
        public Builder oidcGroupsPrefix(@Nullable String oidcGroupsPrefix) {

            this.oidcGroupsPrefix = oidcGroupsPrefix;
            return this;
        }
        @CustomType.Setter
        public Builder oidcRequiredClaims(@Nullable List<String> oidcRequiredClaims) {

            this.oidcRequiredClaims = oidcRequiredClaims;
            return this;
        }
        public Builder oidcRequiredClaims(String... oidcRequiredClaims) {
            return oidcRequiredClaims(List.of(oidcRequiredClaims));
        }
        @CustomType.Setter
        public Builder oidcSigningAlgs(@Nullable List<String> oidcSigningAlgs) {

            this.oidcSigningAlgs = oidcSigningAlgs;
            return this;
        }
        public Builder oidcSigningAlgs(String... oidcSigningAlgs) {
            return oidcSigningAlgs(List.of(oidcSigningAlgs));
        }
        @CustomType.Setter
        public Builder oidcUsernameClaim(@Nullable String oidcUsernameClaim) {

            this.oidcUsernameClaim = oidcUsernameClaim;
            return this;
        }
        @CustomType.Setter
        public Builder oidcUsernamePrefix(@Nullable String oidcUsernamePrefix) {

            this.oidcUsernamePrefix = oidcUsernamePrefix;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetKubeOidcResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        public GetKubeOidcResult build() {
            final var _resultValue = new GetKubeOidcResult();
            _resultValue.clientId = clientId;
            _resultValue.id = id;
            _resultValue.issuerUrl = issuerUrl;
            _resultValue.kubeId = kubeId;
            _resultValue.oidcCaContent = oidcCaContent;
            _resultValue.oidcGroupsClaims = oidcGroupsClaims;
            _resultValue.oidcGroupsPrefix = oidcGroupsPrefix;
            _resultValue.oidcRequiredClaims = oidcRequiredClaims;
            _resultValue.oidcSigningAlgs = oidcSigningAlgs;
            _resultValue.oidcUsernameClaim = oidcUsernameClaim;
            _resultValue.oidcUsernamePrefix = oidcUsernamePrefix;
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}
