// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dedicated.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ServerPlanConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerPlanConfigurationArgs Empty = new ServerPlanConfigurationArgs();

    /**
     * Label for your configuration item
     * 
     */
    @Import(name="label", required=true)
    private Output<String> label;

    /**
     * @return Label for your configuration item
     * 
     */
    public Output<String> label() {
        return this.label;
    }

    /**
     * Value or resource URL on API.OVH.COM of your configuration item
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return Value or resource URL on API.OVH.COM of your configuration item
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    private ServerPlanConfigurationArgs() {}

    private ServerPlanConfigurationArgs(ServerPlanConfigurationArgs $) {
        this.label = $.label;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerPlanConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerPlanConfigurationArgs $;

        public Builder() {
            $ = new ServerPlanConfigurationArgs();
        }

        public Builder(ServerPlanConfigurationArgs defaults) {
            $ = new ServerPlanConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param label Label for your configuration item
         * 
         * @return builder
         * 
         */
        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label Label for your configuration item
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param value Value or resource URL on API.OVH.COM of your configuration item
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Value or resource URL on API.OVH.COM of your configuration item
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public ServerPlanConfigurationArgs build() {
            if ($.label == null) {
                throw new MissingRequiredPropertyException("ServerPlanConfigurationArgs", "label");
            }
            if ($.value == null) {
                throw new MissingRequiredPropertyException("ServerPlanConfigurationArgs", "value");
            }
            return $;
        }
    }

}
