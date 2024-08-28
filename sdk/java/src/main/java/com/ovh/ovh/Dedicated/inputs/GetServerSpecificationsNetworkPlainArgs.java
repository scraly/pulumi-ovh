// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dedicated.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetServerSpecificationsNetworkPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetServerSpecificationsNetworkPlainArgs Empty = new GetServerSpecificationsNetworkPlainArgs();

    /**
     * The internal name of your dedicated server.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The internal name of your dedicated server.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetServerSpecificationsNetworkPlainArgs() {}

    private GetServerSpecificationsNetworkPlainArgs(GetServerSpecificationsNetworkPlainArgs $) {
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetServerSpecificationsNetworkPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetServerSpecificationsNetworkPlainArgs $;

        public Builder() {
            $ = new GetServerSpecificationsNetworkPlainArgs();
        }

        public Builder(GetServerSpecificationsNetworkPlainArgs defaults) {
            $ = new GetServerSpecificationsNetworkPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The internal name of your dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetServerSpecificationsNetworkPlainArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetServerSpecificationsNetworkPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
