// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRegionsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRegionsPlainArgs Empty = new GetRegionsPlainArgs();

    /**
     * List of services which has to be UP in regions.
     * Example: &#34;image&#34;, &#34;instance&#34;, &#34;network&#34;, &#34;storage&#34;, &#34;volume&#34;, &#34;workflow&#34;, ...
     * If left blank, returns all regions associated with the service_name.
     * 
     */
    @Import(name="hasServicesUps")
    private @Nullable List<String> hasServicesUps;

    /**
     * @return List of services which has to be UP in regions.
     * Example: &#34;image&#34;, &#34;instance&#34;, &#34;network&#34;, &#34;storage&#34;, &#34;volume&#34;, &#34;workflow&#34;, ...
     * If left blank, returns all regions associated with the service_name.
     * 
     */
    public Optional<List<String>> hasServicesUps() {
        return Optional.ofNullable(this.hasServicesUps);
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

    private GetRegionsPlainArgs() {}

    private GetRegionsPlainArgs(GetRegionsPlainArgs $) {
        this.hasServicesUps = $.hasServicesUps;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRegionsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRegionsPlainArgs $;

        public Builder() {
            $ = new GetRegionsPlainArgs();
        }

        public Builder(GetRegionsPlainArgs defaults) {
            $ = new GetRegionsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param hasServicesUps List of services which has to be UP in regions.
         * Example: &#34;image&#34;, &#34;instance&#34;, &#34;network&#34;, &#34;storage&#34;, &#34;volume&#34;, &#34;workflow&#34;, ...
         * If left blank, returns all regions associated with the service_name.
         * 
         * @return builder
         * 
         */
        public Builder hasServicesUps(@Nullable List<String> hasServicesUps) {
            $.hasServicesUps = hasServicesUps;
            return this;
        }

        /**
         * @param hasServicesUps List of services which has to be UP in regions.
         * Example: &#34;image&#34;, &#34;instance&#34;, &#34;network&#34;, &#34;storage&#34;, &#34;volume&#34;, &#34;workflow&#34;, ...
         * If left blank, returns all regions associated with the service_name.
         * 
         * @return builder
         * 
         */
        public Builder hasServicesUps(String... hasServicesUps) {
            return hasServicesUps(List.of(hasServicesUps));
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

        public GetRegionsPlainArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetRegionsPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
