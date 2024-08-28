// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.inputs;

import com.ovh.ovh.CloudProject.inputs.KubeNodePoolTemplateMetadataArgs;
import com.ovh.ovh.CloudProject.inputs.KubeNodePoolTemplateSpecArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.util.Objects;


public final class KubeNodePoolTemplateArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubeNodePoolTemplateArgs Empty = new KubeNodePoolTemplateArgs();

    /**
     * Metadata of each node in the pool
     * 
     */
    @Import(name="metadata", required=true)
    private Output<KubeNodePoolTemplateMetadataArgs> metadata;

    /**
     * @return Metadata of each node in the pool
     * 
     */
    public Output<KubeNodePoolTemplateMetadataArgs> metadata() {
        return this.metadata;
    }

    /**
     * Spec of each node in the pool
     * 
     */
    @Import(name="spec", required=true)
    private Output<KubeNodePoolTemplateSpecArgs> spec;

    /**
     * @return Spec of each node in the pool
     * 
     */
    public Output<KubeNodePoolTemplateSpecArgs> spec() {
        return this.spec;
    }

    private KubeNodePoolTemplateArgs() {}

    private KubeNodePoolTemplateArgs(KubeNodePoolTemplateArgs $) {
        this.metadata = $.metadata;
        this.spec = $.spec;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeNodePoolTemplateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeNodePoolTemplateArgs $;

        public Builder() {
            $ = new KubeNodePoolTemplateArgs();
        }

        public Builder(KubeNodePoolTemplateArgs defaults) {
            $ = new KubeNodePoolTemplateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param metadata Metadata of each node in the pool
         * 
         * @return builder
         * 
         */
        public Builder metadata(Output<KubeNodePoolTemplateMetadataArgs> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata Metadata of each node in the pool
         * 
         * @return builder
         * 
         */
        public Builder metadata(KubeNodePoolTemplateMetadataArgs metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param spec Spec of each node in the pool
         * 
         * @return builder
         * 
         */
        public Builder spec(Output<KubeNodePoolTemplateSpecArgs> spec) {
            $.spec = spec;
            return this;
        }

        /**
         * @param spec Spec of each node in the pool
         * 
         * @return builder
         * 
         */
        public Builder spec(KubeNodePoolTemplateSpecArgs spec) {
            return spec(Output.of(spec));
        }

        public KubeNodePoolTemplateArgs build() {
            if ($.metadata == null) {
                throw new MissingRequiredPropertyException("KubeNodePoolTemplateArgs", "metadata");
            }
            if ($.spec == null) {
                throw new MissingRequiredPropertyException("KubeNodePoolTemplateArgs", "spec");
            }
            return $;
        }
    }

}