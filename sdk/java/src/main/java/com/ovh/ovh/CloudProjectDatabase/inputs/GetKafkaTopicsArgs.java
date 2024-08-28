// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetKafkaTopicsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetKafkaTopicsArgs Empty = new GetKafkaTopicsArgs();

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

    private GetKafkaTopicsArgs() {}

    private GetKafkaTopicsArgs(GetKafkaTopicsArgs $) {
        this.clusterId = $.clusterId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKafkaTopicsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKafkaTopicsArgs $;

        public Builder() {
            $ = new GetKafkaTopicsArgs();
        }

        public Builder(GetKafkaTopicsArgs defaults) {
            $ = new GetKafkaTopicsArgs(Objects.requireNonNull(defaults));
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

        public GetKafkaTopicsArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("GetKafkaTopicsArgs", "clusterId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetKafkaTopicsArgs", "serviceName");
            }
            return $;
        }
    }

}
