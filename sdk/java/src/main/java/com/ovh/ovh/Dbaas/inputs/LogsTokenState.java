// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dbaas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LogsTokenState extends com.pulumi.resources.ResourceArgs {

    public static final LogsTokenState Empty = new LogsTokenState();

    /**
     * Cluster ID. If not provided, the default cluster_id is used
     * 
     */
    @Import(name="clusterId")
    private @Nullable Output<String> clusterId;

    /**
     * @return Cluster ID. If not provided, the default cluster_id is used
     * 
     */
    public Optional<Output<String>> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }

    /**
     * Token creation date
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Token creation date
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * Name of the token
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the token
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The LDP service name
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The LDP service name
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * ID of the token
     * 
     */
    @Import(name="tokenId")
    private @Nullable Output<String> tokenId;

    /**
     * @return ID of the token
     * 
     */
    public Optional<Output<String>> tokenId() {
        return Optional.ofNullable(this.tokenId);
    }

    /**
     * Token last update date
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return Token last update date
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    /**
     * Token value
     * 
     */
    @Import(name="value")
    private @Nullable Output<String> value;

    /**
     * @return Token value
     * 
     */
    public Optional<Output<String>> value() {
        return Optional.ofNullable(this.value);
    }

    private LogsTokenState() {}

    private LogsTokenState(LogsTokenState $) {
        this.clusterId = $.clusterId;
        this.createdAt = $.createdAt;
        this.name = $.name;
        this.serviceName = $.serviceName;
        this.tokenId = $.tokenId;
        this.updatedAt = $.updatedAt;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LogsTokenState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LogsTokenState $;

        public Builder() {
            $ = new LogsTokenState();
        }

        public Builder(LogsTokenState defaults) {
            $ = new LogsTokenState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID. If not provided, the default cluster_id is used
         * 
         * @return builder
         * 
         */
        public Builder clusterId(@Nullable Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId Cluster ID. If not provided, the default cluster_id is used
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param createdAt Token creation date
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Token creation date
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param name Name of the token
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the token
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param serviceName The LDP service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The LDP service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param tokenId ID of the token
         * 
         * @return builder
         * 
         */
        public Builder tokenId(@Nullable Output<String> tokenId) {
            $.tokenId = tokenId;
            return this;
        }

        /**
         * @param tokenId ID of the token
         * 
         * @return builder
         * 
         */
        public Builder tokenId(String tokenId) {
            return tokenId(Output.of(tokenId));
        }

        /**
         * @param updatedAt Token last update date
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt Token last update date
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        /**
         * @param value Token value
         * 
         * @return builder
         * 
         */
        public Builder value(@Nullable Output<String> value) {
            $.value = value;
            return this;
        }

        /**
         * @param value Token value
         * 
         * @return builder
         * 
         */
        public Builder value(String value) {
            return value(Output.of(value));
        }

        public LogsTokenState build() {
            return $;
        }
    }

}
