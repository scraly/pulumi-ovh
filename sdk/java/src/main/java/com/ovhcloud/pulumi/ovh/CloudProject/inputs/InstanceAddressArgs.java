// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceAddressArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstanceAddressArgs Empty = new InstanceAddressArgs();

    /**
     * IP address
     * 
     */
    @Import(name="ip")
    private @Nullable Output<String> ip;

    /**
     * @return IP address
     * 
     */
    public Optional<Output<String>> ip() {
        return Optional.ofNullable(this.ip);
    }

    /**
     * IP version
     * 
     */
    @Import(name="version")
    private @Nullable Output<Integer> version;

    /**
     * @return IP version
     * 
     */
    public Optional<Output<Integer>> version() {
        return Optional.ofNullable(this.version);
    }

    private InstanceAddressArgs() {}

    private InstanceAddressArgs(InstanceAddressArgs $) {
        this.ip = $.ip;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceAddressArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceAddressArgs $;

        public Builder() {
            $ = new InstanceAddressArgs();
        }

        public Builder(InstanceAddressArgs defaults) {
            $ = new InstanceAddressArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ip IP address
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip IP address
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param version IP version
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<Integer> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version IP version
         * 
         * @return builder
         * 
         */
        public Builder version(Integer version) {
            return version(Output.of(version));
        }

        public InstanceAddressArgs build() {
            return $;
        }
    }

}
