// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ContainerRegistryOIDCState extends com.pulumi.resources.ResourceArgs {

    public static final ContainerRegistryOIDCState Empty = new ContainerRegistryOIDCState();

    /**
     * Delete existing users from Harbor. OIDC can&#39;t be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
     * 
     */
    @Import(name="deleteUsers")
    private @Nullable Output<Boolean> deleteUsers;

    /**
     * @return Delete existing users from Harbor. OIDC can&#39;t be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
     * 
     */
    public Optional<Output<Boolean>> deleteUsers() {
        return Optional.ofNullable(this.deleteUsers);
    }

    /**
     * Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
     * 
     */
    @Import(name="oidcAdminGroup")
    private @Nullable Output<String> oidcAdminGroup;

    /**
     * @return Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
     * 
     */
    public Optional<Output<String>> oidcAdminGroup() {
        return Optional.ofNullable(this.oidcAdminGroup);
    }

    /**
     * Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
     * 
     */
    @Import(name="oidcAutoOnboard")
    private @Nullable Output<Boolean> oidcAutoOnboard;

    /**
     * @return Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
     * 
     */
    public Optional<Output<Boolean>> oidcAutoOnboard() {
        return Optional.ofNullable(this.oidcAutoOnboard);
    }

    /**
     * The client ID with which Harbor is registered as client application with the OIDC provider.
     * 
     */
    @Import(name="oidcClientId")
    private @Nullable Output<String> oidcClientId;

    /**
     * @return The client ID with which Harbor is registered as client application with the OIDC provider.
     * 
     */
    public Optional<Output<String>> oidcClientId() {
        return Optional.ofNullable(this.oidcClientId);
    }

    /**
     * The secret for the Harbor client application.
     * 
     */
    @Import(name="oidcClientSecret")
    private @Nullable Output<String> oidcClientSecret;

    /**
     * @return The secret for the Harbor client application.
     * 
     */
    public Optional<Output<String>> oidcClientSecret() {
        return Optional.ofNullable(this.oidcClientSecret);
    }

    /**
     * The URL of an OIDC-compliant server.
     * 
     */
    @Import(name="oidcEndpoint")
    private @Nullable Output<String> oidcEndpoint;

    /**
     * @return The URL of an OIDC-compliant server.
     * 
     */
    public Optional<Output<String>> oidcEndpoint() {
        return Optional.ofNullable(this.oidcEndpoint);
    }

    /**
     * The name of Claim in the ID token whose value is the list of group names.
     * 
     */
    @Import(name="oidcGroupsClaim")
    private @Nullable Output<String> oidcGroupsClaim;

    /**
     * @return The name of Claim in the ID token whose value is the list of group names.
     * 
     */
    public Optional<Output<String>> oidcGroupsClaim() {
        return Optional.ofNullable(this.oidcGroupsClaim);
    }

    /**
     * The name of the OIDC provider.
     * 
     */
    @Import(name="oidcName")
    private @Nullable Output<String> oidcName;

    /**
     * @return The name of the OIDC provider.
     * 
     */
    public Optional<Output<String>> oidcName() {
        return Optional.ofNullable(this.oidcName);
    }

    /**
     * The scope sent to OIDC server during authentication. It&#39;s a comma-separated string that must contain &#39;openid&#39; and usually also contains &#39;profile&#39; and &#39;email&#39;. To obtain refresh tokens it should also contain &#39;offline_access&#39;.
     * 
     */
    @Import(name="oidcScope")
    private @Nullable Output<String> oidcScope;

    /**
     * @return The scope sent to OIDC server during authentication. It&#39;s a comma-separated string that must contain &#39;openid&#39; and usually also contains &#39;profile&#39; and &#39;email&#39;. To obtain refresh tokens it should also contain &#39;offline_access&#39;.
     * 
     */
    public Optional<Output<String>> oidcScope() {
        return Optional.ofNullable(this.oidcScope);
    }

    /**
     * The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to &#39;name&#39; (only useful when automatic Onboarding is enabled).
     * 
     */
    @Import(name="oidcUserClaim")
    private @Nullable Output<String> oidcUserClaim;

    /**
     * @return The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to &#39;name&#39; (only useful when automatic Onboarding is enabled).
     * 
     */
    public Optional<Output<String>> oidcUserClaim() {
        return Optional.ofNullable(this.oidcUserClaim);
    }

    /**
     * Set it to `false` if your OIDC server is hosted via self-signed certificate.
     * 
     */
    @Import(name="oidcVerifyCert")
    private @Nullable Output<Boolean> oidcVerifyCert;

    /**
     * @return Set it to `false` if your OIDC server is hosted via self-signed certificate.
     * 
     */
    public Optional<Output<Boolean>> oidcVerifyCert() {
        return Optional.ofNullable(this.oidcVerifyCert);
    }

    /**
     * The ID of the Managed Private Registry. **Changing this value recreates the resource.**
     * 
     */
    @Import(name="registryId")
    private @Nullable Output<String> registryId;

    /**
     * @return The ID of the Managed Private Registry. **Changing this value recreates the resource.**
     * 
     */
    public Optional<Output<String>> registryId() {
        return Optional.ofNullable(this.registryId);
    }

    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    private ContainerRegistryOIDCState() {}

    private ContainerRegistryOIDCState(ContainerRegistryOIDCState $) {
        this.deleteUsers = $.deleteUsers;
        this.oidcAdminGroup = $.oidcAdminGroup;
        this.oidcAutoOnboard = $.oidcAutoOnboard;
        this.oidcClientId = $.oidcClientId;
        this.oidcClientSecret = $.oidcClientSecret;
        this.oidcEndpoint = $.oidcEndpoint;
        this.oidcGroupsClaim = $.oidcGroupsClaim;
        this.oidcName = $.oidcName;
        this.oidcScope = $.oidcScope;
        this.oidcUserClaim = $.oidcUserClaim;
        this.oidcVerifyCert = $.oidcVerifyCert;
        this.registryId = $.registryId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ContainerRegistryOIDCState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ContainerRegistryOIDCState $;

        public Builder() {
            $ = new ContainerRegistryOIDCState();
        }

        public Builder(ContainerRegistryOIDCState defaults) {
            $ = new ContainerRegistryOIDCState(Objects.requireNonNull(defaults));
        }

        /**
         * @param deleteUsers Delete existing users from Harbor. OIDC can&#39;t be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder deleteUsers(@Nullable Output<Boolean> deleteUsers) {
            $.deleteUsers = deleteUsers;
            return this;
        }

        /**
         * @param deleteUsers Delete existing users from Harbor. OIDC can&#39;t be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder deleteUsers(Boolean deleteUsers) {
            return deleteUsers(Output.of(deleteUsers));
        }

        /**
         * @param oidcAdminGroup Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
         * 
         * @return builder
         * 
         */
        public Builder oidcAdminGroup(@Nullable Output<String> oidcAdminGroup) {
            $.oidcAdminGroup = oidcAdminGroup;
            return this;
        }

        /**
         * @param oidcAdminGroup Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
         * 
         * @return builder
         * 
         */
        public Builder oidcAdminGroup(String oidcAdminGroup) {
            return oidcAdminGroup(Output.of(oidcAdminGroup));
        }

        /**
         * @param oidcAutoOnboard Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
         * 
         * @return builder
         * 
         */
        public Builder oidcAutoOnboard(@Nullable Output<Boolean> oidcAutoOnboard) {
            $.oidcAutoOnboard = oidcAutoOnboard;
            return this;
        }

        /**
         * @param oidcAutoOnboard Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
         * 
         * @return builder
         * 
         */
        public Builder oidcAutoOnboard(Boolean oidcAutoOnboard) {
            return oidcAutoOnboard(Output.of(oidcAutoOnboard));
        }

        /**
         * @param oidcClientId The client ID with which Harbor is registered as client application with the OIDC provider.
         * 
         * @return builder
         * 
         */
        public Builder oidcClientId(@Nullable Output<String> oidcClientId) {
            $.oidcClientId = oidcClientId;
            return this;
        }

        /**
         * @param oidcClientId The client ID with which Harbor is registered as client application with the OIDC provider.
         * 
         * @return builder
         * 
         */
        public Builder oidcClientId(String oidcClientId) {
            return oidcClientId(Output.of(oidcClientId));
        }

        /**
         * @param oidcClientSecret The secret for the Harbor client application.
         * 
         * @return builder
         * 
         */
        public Builder oidcClientSecret(@Nullable Output<String> oidcClientSecret) {
            $.oidcClientSecret = oidcClientSecret;
            return this;
        }

        /**
         * @param oidcClientSecret The secret for the Harbor client application.
         * 
         * @return builder
         * 
         */
        public Builder oidcClientSecret(String oidcClientSecret) {
            return oidcClientSecret(Output.of(oidcClientSecret));
        }

        /**
         * @param oidcEndpoint The URL of an OIDC-compliant server.
         * 
         * @return builder
         * 
         */
        public Builder oidcEndpoint(@Nullable Output<String> oidcEndpoint) {
            $.oidcEndpoint = oidcEndpoint;
            return this;
        }

        /**
         * @param oidcEndpoint The URL of an OIDC-compliant server.
         * 
         * @return builder
         * 
         */
        public Builder oidcEndpoint(String oidcEndpoint) {
            return oidcEndpoint(Output.of(oidcEndpoint));
        }

        /**
         * @param oidcGroupsClaim The name of Claim in the ID token whose value is the list of group names.
         * 
         * @return builder
         * 
         */
        public Builder oidcGroupsClaim(@Nullable Output<String> oidcGroupsClaim) {
            $.oidcGroupsClaim = oidcGroupsClaim;
            return this;
        }

        /**
         * @param oidcGroupsClaim The name of Claim in the ID token whose value is the list of group names.
         * 
         * @return builder
         * 
         */
        public Builder oidcGroupsClaim(String oidcGroupsClaim) {
            return oidcGroupsClaim(Output.of(oidcGroupsClaim));
        }

        /**
         * @param oidcName The name of the OIDC provider.
         * 
         * @return builder
         * 
         */
        public Builder oidcName(@Nullable Output<String> oidcName) {
            $.oidcName = oidcName;
            return this;
        }

        /**
         * @param oidcName The name of the OIDC provider.
         * 
         * @return builder
         * 
         */
        public Builder oidcName(String oidcName) {
            return oidcName(Output.of(oidcName));
        }

        /**
         * @param oidcScope The scope sent to OIDC server during authentication. It&#39;s a comma-separated string that must contain &#39;openid&#39; and usually also contains &#39;profile&#39; and &#39;email&#39;. To obtain refresh tokens it should also contain &#39;offline_access&#39;.
         * 
         * @return builder
         * 
         */
        public Builder oidcScope(@Nullable Output<String> oidcScope) {
            $.oidcScope = oidcScope;
            return this;
        }

        /**
         * @param oidcScope The scope sent to OIDC server during authentication. It&#39;s a comma-separated string that must contain &#39;openid&#39; and usually also contains &#39;profile&#39; and &#39;email&#39;. To obtain refresh tokens it should also contain &#39;offline_access&#39;.
         * 
         * @return builder
         * 
         */
        public Builder oidcScope(String oidcScope) {
            return oidcScope(Output.of(oidcScope));
        }

        /**
         * @param oidcUserClaim The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to &#39;name&#39; (only useful when automatic Onboarding is enabled).
         * 
         * @return builder
         * 
         */
        public Builder oidcUserClaim(@Nullable Output<String> oidcUserClaim) {
            $.oidcUserClaim = oidcUserClaim;
            return this;
        }

        /**
         * @param oidcUserClaim The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to &#39;name&#39; (only useful when automatic Onboarding is enabled).
         * 
         * @return builder
         * 
         */
        public Builder oidcUserClaim(String oidcUserClaim) {
            return oidcUserClaim(Output.of(oidcUserClaim));
        }

        /**
         * @param oidcVerifyCert Set it to `false` if your OIDC server is hosted via self-signed certificate.
         * 
         * @return builder
         * 
         */
        public Builder oidcVerifyCert(@Nullable Output<Boolean> oidcVerifyCert) {
            $.oidcVerifyCert = oidcVerifyCert;
            return this;
        }

        /**
         * @param oidcVerifyCert Set it to `false` if your OIDC server is hosted via self-signed certificate.
         * 
         * @return builder
         * 
         */
        public Builder oidcVerifyCert(Boolean oidcVerifyCert) {
            return oidcVerifyCert(Output.of(oidcVerifyCert));
        }

        /**
         * @param registryId The ID of the Managed Private Registry. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder registryId(@Nullable Output<String> registryId) {
            $.registryId = registryId;
            return this;
        }

        /**
         * @param registryId The ID of the Managed Private Registry. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder registryId(String registryId) {
            return registryId(Output.of(registryId));
        }

        /**
         * @param serviceName The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public ContainerRegistryOIDCState build() {
            return $;
        }
    }

}