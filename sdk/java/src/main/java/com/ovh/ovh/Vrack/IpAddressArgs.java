// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Vrack;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class IpAddressArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpAddressArgs Empty = new IpAddressArgs();

    /**
     * Your IP block.
     * 
     */
    @Import(name="block", required=true)
    private Output<String> block;

    /**
     * @return Your IP block.
     * 
     */
    public Output<String> block() {
        return this.block;
    }

    /**
     * The internal name of your vrack
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The internal name of your vrack
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private IpAddressArgs() {}

    private IpAddressArgs(IpAddressArgs $) {
        this.block = $.block;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpAddressArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpAddressArgs $;

        public Builder() {
            $ = new IpAddressArgs();
        }

        public Builder(IpAddressArgs defaults) {
            $ = new IpAddressArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param block Your IP block.
         * 
         * @return builder
         * 
         */
        public Builder block(Output<String> block) {
            $.block = block;
            return this;
        }

        /**
         * @param block Your IP block.
         * 
         * @return builder
         * 
         */
        public Builder block(String block) {
            return block(Output.of(block));
        }

        /**
         * @param serviceName The internal name of your vrack
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your vrack
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public IpAddressArgs build() {
            if ($.block == null) {
                throw new MissingRequiredPropertyException("IpAddressArgs", "block");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("IpAddressArgs", "serviceName");
            }
            return $;
        }
    }

}
