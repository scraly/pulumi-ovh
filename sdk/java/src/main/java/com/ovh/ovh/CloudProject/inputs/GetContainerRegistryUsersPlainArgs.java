// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetContainerRegistryUsersPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetContainerRegistryUsersPlainArgs Empty = new GetContainerRegistryUsersPlainArgs();

    /**
     * Registry ID
     * 
     */
    @Import(name="registryId", required=true)
    private String registryId;

    /**
     * @return Registry ID
     * 
     */
    public String registryId() {
        return this.registryId;
    }

    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetContainerRegistryUsersPlainArgs() {}

    private GetContainerRegistryUsersPlainArgs(GetContainerRegistryUsersPlainArgs $) {
        this.registryId = $.registryId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetContainerRegistryUsersPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetContainerRegistryUsersPlainArgs $;

        public Builder() {
            $ = new GetContainerRegistryUsersPlainArgs();
        }

        public Builder(GetContainerRegistryUsersPlainArgs defaults) {
            $ = new GetContainerRegistryUsersPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param registryId Registry ID
         * 
         * @return builder
         * 
         */
        public Builder registryId(String registryId) {
            $.registryId = registryId;
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
            $.serviceName = serviceName;
            return this;
        }

        public GetContainerRegistryUsersPlainArgs build() {
            if ($.registryId == null) {
                throw new MissingRequiredPropertyException("GetContainerRegistryUsersPlainArgs", "registryId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetContainerRegistryUsersPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
