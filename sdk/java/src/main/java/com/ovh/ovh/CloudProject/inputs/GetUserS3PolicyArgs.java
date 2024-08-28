// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetUserS3PolicyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserS3PolicyArgs Empty = new GetUserS3PolicyArgs();

    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     * The ID of a public cloud project&#39;s user.
     * 
     */
    @Import(name="userId", required=true)
    private Output<String> userId;

    /**
     * @return The ID of a public cloud project&#39;s user.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }

    private GetUserS3PolicyArgs() {}

    private GetUserS3PolicyArgs(GetUserS3PolicyArgs $) {
        this.serviceName = $.serviceName;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserS3PolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserS3PolicyArgs $;

        public Builder() {
            $ = new GetUserS3PolicyArgs();
        }

        public Builder(GetUserS3PolicyArgs defaults) {
            $ = new GetUserS3PolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The ID of the public cloud project. If omitted,
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
         * @param serviceName The ID of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param userId The ID of a public cloud project&#39;s user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Output<String> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of a public cloud project&#39;s user.
         * 
         * @return builder
         * 
         */
        public Builder userId(String userId) {
            return userId(Output.of(userId));
        }

        public GetUserS3PolicyArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetUserS3PolicyArgs", "serviceName");
            }
            if ($.userId == null) {
                throw new MissingRequiredPropertyException("GetUserS3PolicyArgs", "userId");
            }
            return $;
        }
    }

}
