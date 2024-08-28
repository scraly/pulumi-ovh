// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Ip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class IpServicePlanOptionConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final IpServicePlanOptionConfigurationArgs Empty = new IpServicePlanOptionConfigurationArgs();

    /**
     * Identifier of the resource
     * 
     */
    @Import(name="label", required=true)
    private Output<String> label;

    /**
     * @return Identifier of the resource
     * 
     */
    public Output<String> label() {
        return this.label;
    }

    /**
     * Path to the resource in API.OVH.COM
     * 
     */
    @Import(name="value", required=true)
    private Output<String> value;

    /**
     * @return Path to the resource in API.OVH.COM
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    private IpServicePlanOptionConfigurationArgs() {}

    private IpServicePlanOptionConfigurationArgs(IpServicePlanOptionConfigurationArgs $) {
        this.label = $.label;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IpServicePlanOptionConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IpServicePlanOptionConfigurationArgs $;

        public Builder() {
            $ = new IpServicePlanOptionConfigurationArgs();
        }

        public Builder(IpServicePlanOptionConfigurationArgs defaults) {
            $ = new IpServicePlanOptionConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param label Identifier of the resource
         * 
         * @return builder
         * 
         */
        public Builder label(Output<String> label) {
            $.label = label;
            return this;
        }

        /**
         * @param label Identifier of the resource
         * 
         * @return builder
         * 
         */
        public Builder label(String label) {
            return label(Output.of(label));
        }

        /**
         * @param value Path to the resource in API.OVH.COM
         * 
         * @return builder
         * 
         */
        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Path to the resource in API.OVH.COM
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public IpServicePlanOptionConfigurationArgs build() {
            if ($.label == null) {
                throw new MissingRequiredPropertyException("IpServicePlanOptionConfigurationArgs", "label");
            }
            if ($.value == null) {
                throw new MissingRequiredPropertyException("IpServicePlanOptionConfigurationArgs", "value");
            }
            return $;
        }
    }

}
