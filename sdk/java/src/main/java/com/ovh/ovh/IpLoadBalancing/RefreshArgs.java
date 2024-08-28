// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.IpLoadBalancing;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class RefreshArgs extends com.pulumi.resources.ResourceArgs {

    public static final RefreshArgs Empty = new RefreshArgs();

    /**
     * List of values tracked to trigger refresh, used also to form implicit dependencies
     * 
     */
    @Import(name="keepers", required=true)
    private Output<List<String>> keepers;

    /**
     * @return List of values tracked to trigger refresh, used also to form implicit dependencies
     * 
     */
    public Output<List<String>> keepers() {
        return this.keepers;
    }

    /**
     * The internal name of your IP load balancing
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The internal name of your IP load balancing
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private RefreshArgs() {}

    private RefreshArgs(RefreshArgs $) {
        this.keepers = $.keepers;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RefreshArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RefreshArgs $;

        public Builder() {
            $ = new RefreshArgs();
        }

        public Builder(RefreshArgs defaults) {
            $ = new RefreshArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param keepers List of values tracked to trigger refresh, used also to form implicit dependencies
         * 
         * @return builder
         * 
         */
        public Builder keepers(Output<List<String>> keepers) {
            $.keepers = keepers;
            return this;
        }

        /**
         * @param keepers List of values tracked to trigger refresh, used also to form implicit dependencies
         * 
         * @return builder
         * 
         */
        public Builder keepers(List<String> keepers) {
            return keepers(Output.of(keepers));
        }

        /**
         * @param keepers List of values tracked to trigger refresh, used also to form implicit dependencies
         * 
         * @return builder
         * 
         */
        public Builder keepers(String... keepers) {
            return keepers(List.of(keepers));
        }

        /**
         * @param serviceName The internal name of your IP load balancing
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your IP load balancing
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public RefreshArgs build() {
            if ($.keepers == null) {
                throw new MissingRequiredPropertyException("RefreshArgs", "keepers");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("RefreshArgs", "serviceName");
            }
            return $;
        }
    }

}
