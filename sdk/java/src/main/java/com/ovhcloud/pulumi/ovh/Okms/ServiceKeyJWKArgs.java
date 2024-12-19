// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms;

import com.ovhcloud.pulumi.ovh.Okms.inputs.ServiceKeyJWKKeyArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServiceKeyJWKArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceKeyJWKArgs Empty = new ServiceKeyJWKArgs();

    /**
     * Context of the key
     * 
     */
    @Import(name="context")
    private @Nullable Output<String> context;

    /**
     * @return Context of the key
     * 
     */
    public Optional<Output<String>> context() {
        return Optional.ofNullable(this.context);
    }

    /**
     * Set of JSON Web Keys to import
     * 
     */
    @Import(name="keys", required=true)
    private Output<List<ServiceKeyJWKKeyArgs>> keys;

    /**
     * @return Set of JSON Web Keys to import
     * 
     */
    public Output<List<ServiceKeyJWKKeyArgs>> keys() {
        return this.keys;
    }

    /**
     * Key name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Key name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Okms ID
     * 
     */
    @Import(name="okmsId", required=true)
    private Output<String> okmsId;

    /**
     * @return Okms ID
     * 
     */
    public Output<String> okmsId() {
        return this.okmsId;
    }

    private ServiceKeyJWKArgs() {}

    private ServiceKeyJWKArgs(ServiceKeyJWKArgs $) {
        this.context = $.context;
        this.keys = $.keys;
        this.name = $.name;
        this.okmsId = $.okmsId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceKeyJWKArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceKeyJWKArgs $;

        public Builder() {
            $ = new ServiceKeyJWKArgs();
        }

        public Builder(ServiceKeyJWKArgs defaults) {
            $ = new ServiceKeyJWKArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param context Context of the key
         * 
         * @return builder
         * 
         */
        public Builder context(@Nullable Output<String> context) {
            $.context = context;
            return this;
        }

        /**
         * @param context Context of the key
         * 
         * @return builder
         * 
         */
        public Builder context(String context) {
            return context(Output.of(context));
        }

        /**
         * @param keys Set of JSON Web Keys to import
         * 
         * @return builder
         * 
         */
        public Builder keys(Output<List<ServiceKeyJWKKeyArgs>> keys) {
            $.keys = keys;
            return this;
        }

        /**
         * @param keys Set of JSON Web Keys to import
         * 
         * @return builder
         * 
         */
        public Builder keys(List<ServiceKeyJWKKeyArgs> keys) {
            return keys(Output.of(keys));
        }

        /**
         * @param keys Set of JSON Web Keys to import
         * 
         * @return builder
         * 
         */
        public Builder keys(ServiceKeyJWKKeyArgs... keys) {
            return keys(List.of(keys));
        }

        /**
         * @param name Key name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Key name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param okmsId Okms ID
         * 
         * @return builder
         * 
         */
        public Builder okmsId(Output<String> okmsId) {
            $.okmsId = okmsId;
            return this;
        }

        /**
         * @param okmsId Okms ID
         * 
         * @return builder
         * 
         */
        public Builder okmsId(String okmsId) {
            return okmsId(Output.of(okmsId));
        }

        public ServiceKeyJWKArgs build() {
            if ($.keys == null) {
                throw new MissingRequiredPropertyException("ServiceKeyJWKArgs", "keys");
            }
            if ($.okmsId == null) {
                throw new MissingRequiredPropertyException("ServiceKeyJWKArgs", "okmsId");
            }
            return $;
        }
    }

}
