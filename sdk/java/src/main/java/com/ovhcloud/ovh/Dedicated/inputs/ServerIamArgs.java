// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Dedicated.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerIamArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerIamArgs Empty = new ServerIamArgs();

    /**
     * Resource display name
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Resource display name
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Unique identifier of the resource in the IAM
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return Unique identifier of the resource in the IAM
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * Resource tags. Tags that were internally computed are prefixed with `ovh:`
     * 
     */
    @Import(name="tags")
    private @Nullable Output<Map<String,String>> tags;

    /**
     * @return Resource tags. Tags that were internally computed are prefixed with `ovh:`
     * 
     */
    public Optional<Output<Map<String,String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * URN of the private database, used when writing IAM policies
     * 
     */
    @Import(name="urn")
    private @Nullable Output<String> urn;

    /**
     * @return URN of the private database, used when writing IAM policies
     * 
     */
    public Optional<Output<String>> urn() {
        return Optional.ofNullable(this.urn);
    }

    private ServerIamArgs() {}

    private ServerIamArgs(ServerIamArgs $) {
        this.displayName = $.displayName;
        this.id = $.id;
        this.tags = $.tags;
        this.urn = $.urn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerIamArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerIamArgs $;

        public Builder() {
            $ = new ServerIamArgs();
        }

        public Builder(ServerIamArgs defaults) {
            $ = new ServerIamArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param displayName Resource display name
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Resource display name
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param id Unique identifier of the resource in the IAM
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id Unique identifier of the resource in the IAM
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param tags Resource tags. Tags that were internally computed are prefixed with `ovh:`
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<Map<String,String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Resource tags. Tags that were internally computed are prefixed with `ovh:`
         * 
         * @return builder
         * 
         */
        public Builder tags(Map<String,String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param urn URN of the private database, used when writing IAM policies
         * 
         * @return builder
         * 
         */
        public Builder urn(@Nullable Output<String> urn) {
            $.urn = urn;
            return this;
        }

        /**
         * @param urn URN of the private database, used when writing IAM policies
         * 
         * @return builder
         * 
         */
        public Builder urn(String urn) {
            return urn(Output.of(urn));
        }

        public ServerIamArgs build() {
            return $;
        }
    }

}
