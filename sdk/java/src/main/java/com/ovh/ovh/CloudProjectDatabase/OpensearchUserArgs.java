// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProjectDatabase;

import com.ovh.ovh.CloudProjectDatabase.inputs.OpensearchUserAclArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OpensearchUserArgs extends com.pulumi.resources.ResourceArgs {

    public static final OpensearchUserArgs Empty = new OpensearchUserArgs();

    /**
     * Acls of the user.
     * 
     */
    @Import(name="acls")
    private @Nullable Output<List<OpensearchUserAclArgs>> acls;

    /**
     * @return Acls of the user.
     * 
     */
    public Optional<Output<List<OpensearchUserAclArgs>>> acls() {
        return Optional.ofNullable(this.acls);
    }

    /**
     * Cluster ID.
     * 
     */
    @Import(name="clusterId", required=true)
    private Output<String> clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }

    /**
     * Username affected by this acl. A user named &#34;avnadmin&#34; is mapped with already created admin user and reset his password instead of creating a new user.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Username affected by this acl. A user named &#34;avnadmin&#34; is mapped with already created admin user and reset his password instead of creating a new user.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Arbitrary string to change to trigger a password update
     * 
     */
    @Import(name="passwordReset")
    private @Nullable Output<String> passwordReset;

    /**
     * @return Arbitrary string to change to trigger a password update
     * 
     */
    public Optional<Output<String>> passwordReset() {
        return Optional.ofNullable(this.passwordReset);
    }

    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private OpensearchUserArgs() {}

    private OpensearchUserArgs(OpensearchUserArgs $) {
        this.acls = $.acls;
        this.clusterId = $.clusterId;
        this.name = $.name;
        this.passwordReset = $.passwordReset;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OpensearchUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OpensearchUserArgs $;

        public Builder() {
            $ = new OpensearchUserArgs();
        }

        public Builder(OpensearchUserArgs defaults) {
            $ = new OpensearchUserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param acls Acls of the user.
         * 
         * @return builder
         * 
         */
        public Builder acls(@Nullable Output<List<OpensearchUserAclArgs>> acls) {
            $.acls = acls;
            return this;
        }

        /**
         * @param acls Acls of the user.
         * 
         * @return builder
         * 
         */
        public Builder acls(List<OpensearchUserAclArgs> acls) {
            return acls(Output.of(acls));
        }

        /**
         * @param acls Acls of the user.
         * 
         * @return builder
         * 
         */
        public Builder acls(OpensearchUserAclArgs... acls) {
            return acls(List.of(acls));
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param name Username affected by this acl. A user named &#34;avnadmin&#34; is mapped with already created admin user and reset his password instead of creating a new user.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Username affected by this acl. A user named &#34;avnadmin&#34; is mapped with already created admin user and reset his password instead of creating a new user.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param passwordReset Arbitrary string to change to trigger a password update
         * 
         * @return builder
         * 
         */
        public Builder passwordReset(@Nullable Output<String> passwordReset) {
            $.passwordReset = passwordReset;
            return this;
        }

        /**
         * @param passwordReset Arbitrary string to change to trigger a password update
         * 
         * @return builder
         * 
         */
        public Builder passwordReset(String passwordReset) {
            return passwordReset(Output.of(passwordReset));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public OpensearchUserArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("OpensearchUserArgs", "clusterId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("OpensearchUserArgs", "serviceName");
            }
            return $;
        }
    }

}
