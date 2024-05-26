// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkPrivateRegionsStatusArgs extends com.pulumi.resources.ResourceArgs {

    public static final NetworkPrivateRegionsStatusArgs Empty = new NetworkPrivateRegionsStatusArgs();

    @Import(name="region")
    private @Nullable Output<String> region;

    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * the status of the network. should be normally set to &#39;ACTIVE&#39;.
     * 
     */
    @Import(name="status", required=true)
    private Output<String> status;

    /**
     * @return the status of the network. should be normally set to &#39;ACTIVE&#39;.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    private NetworkPrivateRegionsStatusArgs() {}

    private NetworkPrivateRegionsStatusArgs(NetworkPrivateRegionsStatusArgs $) {
        this.region = $.region;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkPrivateRegionsStatusArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkPrivateRegionsStatusArgs $;

        public Builder() {
            $ = new NetworkPrivateRegionsStatusArgs();
        }

        public Builder(NetworkPrivateRegionsStatusArgs defaults) {
            $ = new NetworkPrivateRegionsStatusArgs(Objects.requireNonNull(defaults));
        }

        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param status the status of the network. should be normally set to &#39;ACTIVE&#39;.
         * 
         * @return builder
         * 
         */
        public Builder status(Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status the status of the network. should be normally set to &#39;ACTIVE&#39;.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public NetworkPrivateRegionsStatusArgs build() {
            if ($.status == null) {
                throw new MissingRequiredPropertyException("NetworkPrivateRegionsStatusArgs", "status");
            }
            return $;
        }
    }

}
