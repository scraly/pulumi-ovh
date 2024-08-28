// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.inputs;

import com.ovh.ovh.CloudProject.inputs.GetKubeCustomizationApiserverAdmissionpluginArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.util.List;
import java.util.Objects;


public final class GetKubeCustomizationApiserverArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetKubeCustomizationApiserverArgs Empty = new GetKubeCustomizationApiserverArgs();

    /**
     * Kubernetes API server admission plugins customization
     * 
     */
    @Import(name="admissionplugins", required=true)
    private Output<List<GetKubeCustomizationApiserverAdmissionpluginArgs>> admissionplugins;

    /**
     * @return Kubernetes API server admission plugins customization
     * 
     */
    public Output<List<GetKubeCustomizationApiserverAdmissionpluginArgs>> admissionplugins() {
        return this.admissionplugins;
    }

    private GetKubeCustomizationApiserverArgs() {}

    private GetKubeCustomizationApiserverArgs(GetKubeCustomizationApiserverArgs $) {
        this.admissionplugins = $.admissionplugins;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKubeCustomizationApiserverArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKubeCustomizationApiserverArgs $;

        public Builder() {
            $ = new GetKubeCustomizationApiserverArgs();
        }

        public Builder(GetKubeCustomizationApiserverArgs defaults) {
            $ = new GetKubeCustomizationApiserverArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param admissionplugins Kubernetes API server admission plugins customization
         * 
         * @return builder
         * 
         */
        public Builder admissionplugins(Output<List<GetKubeCustomizationApiserverAdmissionpluginArgs>> admissionplugins) {
            $.admissionplugins = admissionplugins;
            return this;
        }

        /**
         * @param admissionplugins Kubernetes API server admission plugins customization
         * 
         * @return builder
         * 
         */
        public Builder admissionplugins(List<GetKubeCustomizationApiserverAdmissionpluginArgs> admissionplugins) {
            return admissionplugins(Output.of(admissionplugins));
        }

        /**
         * @param admissionplugins Kubernetes API server admission plugins customization
         * 
         * @return builder
         * 
         */
        public Builder admissionplugins(GetKubeCustomizationApiserverAdmissionpluginArgs... admissionplugins) {
            return admissionplugins(List.of(admissionplugins));
        }

        public GetKubeCustomizationApiserverArgs build() {
            if ($.admissionplugins == null) {
                throw new MissingRequiredPropertyException("GetKubeCustomizationApiserverArgs", "admissionplugins");
            }
            return $;
        }
    }

}
