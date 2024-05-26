// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetKubeNodePoolNodesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetKubeNodePoolNodesPlainArgs Empty = new GetKubeNodePoolNodesPlainArgs();

    /**
     * The ID of the managed kubernetes cluster.
     * 
     */
    @Import(name="kubeId", required=true)
    private String kubeId;

    /**
     * @return The ID of the managed kubernetes cluster.
     * 
     */
    public String kubeId() {
        return this.kubeId;
    }

    /**
     * Name of the node pool from which we want the nodes.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name of the node pool from which we want the nodes.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetKubeNodePoolNodesPlainArgs() {}

    private GetKubeNodePoolNodesPlainArgs(GetKubeNodePoolNodesPlainArgs $) {
        this.kubeId = $.kubeId;
        this.name = $.name;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKubeNodePoolNodesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKubeNodePoolNodesPlainArgs $;

        public Builder() {
            $ = new GetKubeNodePoolNodesPlainArgs();
        }

        public Builder(GetKubeNodePoolNodesPlainArgs defaults) {
            $ = new GetKubeNodePoolNodesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param kubeId The ID of the managed kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder kubeId(String kubeId) {
            $.kubeId = kubeId;
            return this;
        }

        /**
         * @param name Name of the node pool from which we want the nodes.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param serviceName The ID of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetKubeNodePoolNodesPlainArgs build() {
            if ($.kubeId == null) {
                throw new MissingRequiredPropertyException("GetKubeNodePoolNodesPlainArgs", "kubeId");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetKubeNodePoolNodesPlainArgs", "name");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetKubeNodePoolNodesPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
