// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Ip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ReverseArgs extends com.pulumi.resources.ResourceArgs {

    public static final ReverseArgs Empty = new ReverseArgs();

    /**
     * The IP to set the reverse of
     * 
     */
    @Import(name="ReverseIp", required=true)
    private Output<String> ReverseIp;

    /**
     * @return The IP to set the reverse of
     * 
     */
    public Output<String> ReverseIp() {
        return this.ReverseIp;
    }

    /**
     * The value of the reverse
     * 
     */
    @Import(name="ReverseValue", required=true)
    private Output<String> ReverseValue;

    /**
     * @return The value of the reverse
     * 
     */
    public Output<String> ReverseValue() {
        return this.ReverseValue;
    }

    /**
     * The IP block to which the IP belongs
     * 
     */
    @Import(name="ip", required=true)
    private Output<String> ip;

    /**
     * @return The IP block to which the IP belongs
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }

    private ReverseArgs() {}

    private ReverseArgs(ReverseArgs $) {
        this.ReverseIp = $.ReverseIp;
        this.ReverseValue = $.ReverseValue;
        this.ip = $.ip;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReverseArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReverseArgs $;

        public Builder() {
            $ = new ReverseArgs();
        }

        public Builder(ReverseArgs defaults) {
            $ = new ReverseArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ReverseIp The IP to set the reverse of
         * 
         * @return builder
         * 
         */
        public Builder ReverseIp(Output<String> ReverseIp) {
            $.ReverseIp = ReverseIp;
            return this;
        }

        /**
         * @param ReverseIp The IP to set the reverse of
         * 
         * @return builder
         * 
         */
        public Builder ReverseIp(String ReverseIp) {
            return ReverseIp(Output.of(ReverseIp));
        }

        /**
         * @param ReverseValue The value of the reverse
         * 
         * @return builder
         * 
         */
        public Builder ReverseValue(Output<String> ReverseValue) {
            $.ReverseValue = ReverseValue;
            return this;
        }

        /**
         * @param ReverseValue The value of the reverse
         * 
         * @return builder
         * 
         */
        public Builder ReverseValue(String ReverseValue) {
            return ReverseValue(Output.of(ReverseValue));
        }

        /**
         * @param ip The IP block to which the IP belongs
         * 
         * @return builder
         * 
         */
        public Builder ip(Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip The IP block to which the IP belongs
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        public ReverseArgs build() {
            if ($.ReverseIp == null) {
                throw new MissingRequiredPropertyException("ReverseArgs", "ReverseIp");
            }
            if ($.ReverseValue == null) {
                throw new MissingRequiredPropertyException("ReverseArgs", "ReverseValue");
            }
            if ($.ip == null) {
                throw new MissingRequiredPropertyException("ReverseArgs", "ip");
            }
            return $;
        }
    }

}
