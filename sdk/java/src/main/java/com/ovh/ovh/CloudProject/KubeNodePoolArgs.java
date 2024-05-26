// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject;

import com.ovh.ovh.CloudProject.inputs.KubeNodePoolTemplateArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubeNodePoolArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubeNodePoolArgs Empty = new KubeNodePoolArgs();

    /**
     * should the pool use the anti-affinity feature. Default to `false`. **Changing this value recreates the resource.**
     * 
     */
    @Import(name="antiAffinity")
    private @Nullable Output<Boolean> antiAffinity;

    /**
     * @return should the pool use the anti-affinity feature. Default to `false`. **Changing this value recreates the resource.**
     * 
     */
    public Optional<Output<Boolean>> antiAffinity() {
        return Optional.ofNullable(this.antiAffinity);
    }

    /**
     * Enable auto-scaling for the pool. Default to `false`.
     * * ` template  ` - (Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:
     * 
     */
    @Import(name="autoscale")
    private @Nullable Output<Boolean> autoscale;

    /**
     * @return Enable auto-scaling for the pool. Default to `false`.
     * * ` template  ` - (Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:
     * 
     */
    public Optional<Output<Boolean>> autoscale() {
        return Optional.ofNullable(this.autoscale);
    }

    /**
     * number of nodes to start.
     * 
     */
    @Import(name="desiredNodes")
    private @Nullable Output<Integer> desiredNodes;

    /**
     * @return number of nodes to start.
     * 
     */
    public Optional<Output<Integer>> desiredNodes() {
        return Optional.ofNullable(this.desiredNodes);
    }

    /**
     * a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: &#34;b2-7&#34;. You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
     * **Changing this value recreates the resource.**
     * 
     */
    @Import(name="flavorName", required=true)
    private Output<String> flavorName;

    /**
     * @return a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: &#34;b2-7&#34;. You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
     * **Changing this value recreates the resource.**
     * 
     */
    public Output<String> flavorName() {
        return this.flavorName;
    }

    /**
     * The id of the managed kubernetes cluster. **Changing this value recreates the resource.**
     * 
     */
    @Import(name="kubeId", required=true)
    private Output<String> kubeId;

    /**
     * @return The id of the managed kubernetes cluster. **Changing this value recreates the resource.**
     * 
     */
    public Output<String> kubeId() {
        return this.kubeId;
    }

    /**
     * maximum number of nodes allowed in the pool. Setting `desired_nodes` over this value will raise an error.
     * 
     */
    @Import(name="maxNodes")
    private @Nullable Output<Integer> maxNodes;

    /**
     * @return maximum number of nodes allowed in the pool. Setting `desired_nodes` over this value will raise an error.
     * 
     */
    public Optional<Output<Integer>> maxNodes() {
        return Optional.ofNullable(this.maxNodes);
    }

    /**
     * minimum number of nodes allowed in the pool. Setting `desired_nodes` under this value will raise an error.
     * 
     */
    @Import(name="minNodes")
    private @Nullable Output<Integer> minNodes;

    /**
     * @return minimum number of nodes allowed in the pool. Setting `desired_nodes` under this value will raise an error.
     * 
     */
    public Optional<Output<Integer>> minNodes() {
        return Optional.ofNullable(this.minNodes);
    }

    /**
     * should the nodes be billed on a monthly basis. Default to `false`. **Changing this value recreates the resource.**
     * 
     */
    @Import(name="monthlyBilled")
    private @Nullable Output<Boolean> monthlyBilled;

    /**
     * @return should the nodes be billed on a monthly basis. Default to `false`. **Changing this value recreates the resource.**
     * 
     */
    public Optional<Output<Boolean>> monthlyBilled() {
        return Optional.ofNullable(this.monthlyBilled);
    }

    /**
     * The name of the nodepool. Warning: `_` char is not allowed! **Changing this value recreates the resource.**
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the nodepool. Warning: `_` char is not allowed! **Changing this value recreates the resource.**
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     * Node pool template
     * 
     */
    @Import(name="template")
    private @Nullable Output<KubeNodePoolTemplateArgs> template;

    /**
     * @return Node pool template
     * 
     */
    public Optional<Output<KubeNodePoolTemplateArgs>> template() {
        return Optional.ofNullable(this.template);
    }

    private KubeNodePoolArgs() {}

    private KubeNodePoolArgs(KubeNodePoolArgs $) {
        this.antiAffinity = $.antiAffinity;
        this.autoscale = $.autoscale;
        this.desiredNodes = $.desiredNodes;
        this.flavorName = $.flavorName;
        this.kubeId = $.kubeId;
        this.maxNodes = $.maxNodes;
        this.minNodes = $.minNodes;
        this.monthlyBilled = $.monthlyBilled;
        this.name = $.name;
        this.serviceName = $.serviceName;
        this.template = $.template;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeNodePoolArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeNodePoolArgs $;

        public Builder() {
            $ = new KubeNodePoolArgs();
        }

        public Builder(KubeNodePoolArgs defaults) {
            $ = new KubeNodePoolArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param antiAffinity should the pool use the anti-affinity feature. Default to `false`. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder antiAffinity(@Nullable Output<Boolean> antiAffinity) {
            $.antiAffinity = antiAffinity;
            return this;
        }

        /**
         * @param antiAffinity should the pool use the anti-affinity feature. Default to `false`. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder antiAffinity(Boolean antiAffinity) {
            return antiAffinity(Output.of(antiAffinity));
        }

        /**
         * @param autoscale Enable auto-scaling for the pool. Default to `false`.
         * * ` template  ` - (Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:
         * 
         * @return builder
         * 
         */
        public Builder autoscale(@Nullable Output<Boolean> autoscale) {
            $.autoscale = autoscale;
            return this;
        }

        /**
         * @param autoscale Enable auto-scaling for the pool. Default to `false`.
         * * ` template  ` - (Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:
         * 
         * @return builder
         * 
         */
        public Builder autoscale(Boolean autoscale) {
            return autoscale(Output.of(autoscale));
        }

        /**
         * @param desiredNodes number of nodes to start.
         * 
         * @return builder
         * 
         */
        public Builder desiredNodes(@Nullable Output<Integer> desiredNodes) {
            $.desiredNodes = desiredNodes;
            return this;
        }

        /**
         * @param desiredNodes number of nodes to start.
         * 
         * @return builder
         * 
         */
        public Builder desiredNodes(Integer desiredNodes) {
            return desiredNodes(Output.of(desiredNodes));
        }

        /**
         * @param flavorName a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: &#34;b2-7&#34;. You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
         * **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder flavorName(Output<String> flavorName) {
            $.flavorName = flavorName;
            return this;
        }

        /**
         * @param flavorName a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: &#34;b2-7&#34;. You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.
         * **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder flavorName(String flavorName) {
            return flavorName(Output.of(flavorName));
        }

        /**
         * @param kubeId The id of the managed kubernetes cluster. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder kubeId(Output<String> kubeId) {
            $.kubeId = kubeId;
            return this;
        }

        /**
         * @param kubeId The id of the managed kubernetes cluster. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder kubeId(String kubeId) {
            return kubeId(Output.of(kubeId));
        }

        /**
         * @param maxNodes maximum number of nodes allowed in the pool. Setting `desired_nodes` over this value will raise an error.
         * 
         * @return builder
         * 
         */
        public Builder maxNodes(@Nullable Output<Integer> maxNodes) {
            $.maxNodes = maxNodes;
            return this;
        }

        /**
         * @param maxNodes maximum number of nodes allowed in the pool. Setting `desired_nodes` over this value will raise an error.
         * 
         * @return builder
         * 
         */
        public Builder maxNodes(Integer maxNodes) {
            return maxNodes(Output.of(maxNodes));
        }

        /**
         * @param minNodes minimum number of nodes allowed in the pool. Setting `desired_nodes` under this value will raise an error.
         * 
         * @return builder
         * 
         */
        public Builder minNodes(@Nullable Output<Integer> minNodes) {
            $.minNodes = minNodes;
            return this;
        }

        /**
         * @param minNodes minimum number of nodes allowed in the pool. Setting `desired_nodes` under this value will raise an error.
         * 
         * @return builder
         * 
         */
        public Builder minNodes(Integer minNodes) {
            return minNodes(Output.of(minNodes));
        }

        /**
         * @param monthlyBilled should the nodes be billed on a monthly basis. Default to `false`. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder monthlyBilled(@Nullable Output<Boolean> monthlyBilled) {
            $.monthlyBilled = monthlyBilled;
            return this;
        }

        /**
         * @param monthlyBilled should the nodes be billed on a monthly basis. Default to `false`. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder monthlyBilled(Boolean monthlyBilled) {
            return monthlyBilled(Output.of(monthlyBilled));
        }

        /**
         * @param name The name of the nodepool. Warning: `_` char is not allowed! **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the nodepool. Warning: `_` char is not allowed! **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param template Node pool template
         * 
         * @return builder
         * 
         */
        public Builder template(@Nullable Output<KubeNodePoolTemplateArgs> template) {
            $.template = template;
            return this;
        }

        /**
         * @param template Node pool template
         * 
         * @return builder
         * 
         */
        public Builder template(KubeNodePoolTemplateArgs template) {
            return template(Output.of(template));
        }

        public KubeNodePoolArgs build() {
            if ($.flavorName == null) {
                throw new MissingRequiredPropertyException("KubeNodePoolArgs", "flavorName");
            }
            if ($.kubeId == null) {
                throw new MissingRequiredPropertyException("KubeNodePoolArgs", "kubeId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("KubeNodePoolArgs", "serviceName");
            }
            return $;
        }
    }

}
