// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetKubeNodesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetKubeNodesArgs Empty = new GetKubeNodesArgs();

    /**
     * The ID of the managed kubernetes cluster.
     * 
     */
    @Import(name="kubeId", required=true)
    private Output<String> kubeId;

    /**
     * @return The ID of the managed kubernetes cluster.
     * 
     */
    public Output<String> kubeId() {
        return this.kubeId;
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

    private GetKubeNodesArgs() {}

    private GetKubeNodesArgs(GetKubeNodesArgs $) {
        this.kubeId = $.kubeId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKubeNodesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKubeNodesArgs $;

        public Builder() {
            $ = new GetKubeNodesArgs();
        }

        public Builder(GetKubeNodesArgs defaults) {
            $ = new GetKubeNodesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param kubeId The ID of the managed kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder kubeId(Output<String> kubeId) {
            $.kubeId = kubeId;
            return this;
        }

        /**
         * @param kubeId The ID of the managed kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder kubeId(String kubeId) {
            return kubeId(Output.of(kubeId));
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

        public GetKubeNodesArgs build() {
            if ($.kubeId == null) {
                throw new MissingRequiredPropertyException("GetKubeNodesArgs", "kubeId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetKubeNodesArgs", "serviceName");
            }
            return $;
        }
    }

}
