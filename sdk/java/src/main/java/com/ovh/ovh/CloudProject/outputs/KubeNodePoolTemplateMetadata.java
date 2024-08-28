// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class KubeNodePoolTemplateMetadata {
    /**
     * @return annotations
     * 
     */
    private Map<String,String> annotations;
    /**
     * @return finalizers
     * 
     */
    private List<String> finalizers;
    /**
     * @return labels
     * 
     */
    private Map<String,String> labels;

    private KubeNodePoolTemplateMetadata() {}
    /**
     * @return annotations
     * 
     */
    public Map<String,String> annotations() {
        return this.annotations;
    }
    /**
     * @return finalizers
     * 
     */
    public List<String> finalizers() {
        return this.finalizers;
    }
    /**
     * @return labels
     * 
     */
    public Map<String,String> labels() {
        return this.labels;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KubeNodePoolTemplateMetadata defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> annotations;
        private List<String> finalizers;
        private Map<String,String> labels;
        public Builder() {}
        public Builder(KubeNodePoolTemplateMetadata defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.annotations = defaults.annotations;
    	      this.finalizers = defaults.finalizers;
    	      this.labels = defaults.labels;
        }

        @CustomType.Setter
        public Builder annotations(Map<String,String> annotations) {
            if (annotations == null) {
              throw new MissingRequiredPropertyException("KubeNodePoolTemplateMetadata", "annotations");
            }
            this.annotations = annotations;
            return this;
        }
        @CustomType.Setter
        public Builder finalizers(List<String> finalizers) {
            if (finalizers == null) {
              throw new MissingRequiredPropertyException("KubeNodePoolTemplateMetadata", "finalizers");
            }
            this.finalizers = finalizers;
            return this;
        }
        public Builder finalizers(String... finalizers) {
            return finalizers(List.of(finalizers));
        }
        @CustomType.Setter
        public Builder labels(Map<String,String> labels) {
            if (labels == null) {
              throw new MissingRequiredPropertyException("KubeNodePoolTemplateMetadata", "labels");
            }
            this.labels = labels;
            return this;
        }
        public KubeNodePoolTemplateMetadata build() {
            final var _resultValue = new KubeNodePoolTemplateMetadata();
            _resultValue.annotations = annotations;
            _resultValue.finalizers = finalizers;
            _resultValue.labels = labels;
            return _resultValue;
        }
    }
}
