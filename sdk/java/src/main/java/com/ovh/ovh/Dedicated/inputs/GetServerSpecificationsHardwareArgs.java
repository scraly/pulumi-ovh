// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dedicated.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetServerSpecificationsHardwareArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServerSpecificationsHardwareArgs Empty = new GetServerSpecificationsHardwareArgs();

    /**
     * The internal name of your dedicated server.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The internal name of your dedicated server.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private GetServerSpecificationsHardwareArgs() {}

    private GetServerSpecificationsHardwareArgs(GetServerSpecificationsHardwareArgs $) {
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServerSpecificationsHardwareArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServerSpecificationsHardwareArgs $;

        public Builder() {
            $ = new GetServerSpecificationsHardwareArgs();
        }

        public Builder(GetServerSpecificationsHardwareArgs defaults) {
            $ = new GetServerSpecificationsHardwareArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The internal name of your dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public GetServerSpecificationsHardwareArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetServerSpecificationsHardwareArgs", "serviceName");
            }
            return $;
        }
    }

}