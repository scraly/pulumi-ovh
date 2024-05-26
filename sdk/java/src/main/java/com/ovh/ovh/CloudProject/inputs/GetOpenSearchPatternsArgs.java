// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetOpenSearchPatternsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOpenSearchPatternsArgs Empty = new GetOpenSearchPatternsArgs();

    /**
     * Cluster ID
     * 
     */
    @Import(name="clusterId", required=true)
    private Output<String> clusterId;

    /**
     * @return Cluster ID
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
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

    private GetOpenSearchPatternsArgs() {}

    private GetOpenSearchPatternsArgs(GetOpenSearchPatternsArgs $) {
        this.clusterId = $.clusterId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOpenSearchPatternsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOpenSearchPatternsArgs $;

        public Builder() {
            $ = new GetOpenSearchPatternsArgs();
        }

        public Builder(GetOpenSearchPatternsArgs defaults) {
            $ = new GetOpenSearchPatternsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID
         * 
         * @return builder
         * 
         */
        public Builder clusterId(Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId Cluster ID
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
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

        public GetOpenSearchPatternsArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("GetOpenSearchPatternsArgs", "clusterId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetOpenSearchPatternsArgs", "serviceName");
            }
            return $;
        }
    }

}
