// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dbaas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LogsOutputOpenSearchIndexState extends com.pulumi.resources.ResourceArgs {

    public static final LogsOutputOpenSearchIndexState Empty = new LogsOutputOpenSearchIndexState();

    /**
     * If set, notify when size is near 80, 90 or 100 % of its maximum capacity
     * 
     */
    @Import(name="alertNotifyEnabled")
    private @Nullable Output<Boolean> alertNotifyEnabled;

    /**
     * @return If set, notify when size is near 80, 90 or 100 % of its maximum capacity
     * 
     */
    public Optional<Output<Boolean>> alertNotifyEnabled() {
        return Optional.ofNullable(this.alertNotifyEnabled);
    }

    /**
     * Index creation
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Index creation
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * Current index size (in bytes)
     * 
     */
    @Import(name="currentSize")
    private @Nullable Output<Integer> currentSize;

    /**
     * @return Current index size (in bytes)
     * 
     */
    public Optional<Output<Integer>> currentSize() {
        return Optional.ofNullable(this.currentSize);
    }

    /**
     * Index description
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Index description
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Index ID
     * 
     */
    @Import(name="indexId")
    private @Nullable Output<String> indexId;

    /**
     * @return Index ID
     * 
     */
    public Optional<Output<String>> indexId() {
        return Optional.ofNullable(this.indexId);
    }

    /**
     * Indicates if you are allowed to edit entry
     * 
     */
    @Import(name="isEditable")
    private @Nullable Output<Boolean> isEditable;

    /**
     * @return Indicates if you are allowed to edit entry
     * 
     */
    public Optional<Output<Boolean>> isEditable() {
        return Optional.ofNullable(this.isEditable);
    }

    /**
     * Maximum index size (in bytes)
     * 
     */
    @Import(name="maxSize")
    private @Nullable Output<Integer> maxSize;

    /**
     * @return Maximum index size (in bytes)
     * 
     */
    public Optional<Output<Integer>> maxSize() {
        return Optional.ofNullable(this.maxSize);
    }

    /**
     * Index name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Index name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Number of shards
     * 
     */
    @Import(name="nbShard")
    private @Nullable Output<Integer> nbShard;

    /**
     * @return Number of shards
     * 
     */
    public Optional<Output<Integer>> nbShard() {
        return Optional.ofNullable(this.nbShard);
    }

    /**
     * The service name
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The service name
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Index suffix
     * 
     */
    @Import(name="suffix")
    private @Nullable Output<String> suffix;

    /**
     * @return Index suffix
     * 
     */
    public Optional<Output<String>> suffix() {
        return Optional.ofNullable(this.suffix);
    }

    /**
     * Index last update
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return Index last update
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    private LogsOutputOpenSearchIndexState() {}

    private LogsOutputOpenSearchIndexState(LogsOutputOpenSearchIndexState $) {
        this.alertNotifyEnabled = $.alertNotifyEnabled;
        this.createdAt = $.createdAt;
        this.currentSize = $.currentSize;
        this.description = $.description;
        this.indexId = $.indexId;
        this.isEditable = $.isEditable;
        this.maxSize = $.maxSize;
        this.name = $.name;
        this.nbShard = $.nbShard;
        this.serviceName = $.serviceName;
        this.suffix = $.suffix;
        this.updatedAt = $.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LogsOutputOpenSearchIndexState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LogsOutputOpenSearchIndexState $;

        public Builder() {
            $ = new LogsOutputOpenSearchIndexState();
        }

        public Builder(LogsOutputOpenSearchIndexState defaults) {
            $ = new LogsOutputOpenSearchIndexState(Objects.requireNonNull(defaults));
        }

        /**
         * @param alertNotifyEnabled If set, notify when size is near 80, 90 or 100 % of its maximum capacity
         * 
         * @return builder
         * 
         */
        public Builder alertNotifyEnabled(@Nullable Output<Boolean> alertNotifyEnabled) {
            $.alertNotifyEnabled = alertNotifyEnabled;
            return this;
        }

        /**
         * @param alertNotifyEnabled If set, notify when size is near 80, 90 or 100 % of its maximum capacity
         * 
         * @return builder
         * 
         */
        public Builder alertNotifyEnabled(Boolean alertNotifyEnabled) {
            return alertNotifyEnabled(Output.of(alertNotifyEnabled));
        }

        /**
         * @param createdAt Index creation
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Index creation
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param currentSize Current index size (in bytes)
         * 
         * @return builder
         * 
         */
        public Builder currentSize(@Nullable Output<Integer> currentSize) {
            $.currentSize = currentSize;
            return this;
        }

        /**
         * @param currentSize Current index size (in bytes)
         * 
         * @return builder
         * 
         */
        public Builder currentSize(Integer currentSize) {
            return currentSize(Output.of(currentSize));
        }

        /**
         * @param description Index description
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Index description
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param indexId Index ID
         * 
         * @return builder
         * 
         */
        public Builder indexId(@Nullable Output<String> indexId) {
            $.indexId = indexId;
            return this;
        }

        /**
         * @param indexId Index ID
         * 
         * @return builder
         * 
         */
        public Builder indexId(String indexId) {
            return indexId(Output.of(indexId));
        }

        /**
         * @param isEditable Indicates if you are allowed to edit entry
         * 
         * @return builder
         * 
         */
        public Builder isEditable(@Nullable Output<Boolean> isEditable) {
            $.isEditable = isEditable;
            return this;
        }

        /**
         * @param isEditable Indicates if you are allowed to edit entry
         * 
         * @return builder
         * 
         */
        public Builder isEditable(Boolean isEditable) {
            return isEditable(Output.of(isEditable));
        }

        /**
         * @param maxSize Maximum index size (in bytes)
         * 
         * @return builder
         * 
         */
        public Builder maxSize(@Nullable Output<Integer> maxSize) {
            $.maxSize = maxSize;
            return this;
        }

        /**
         * @param maxSize Maximum index size (in bytes)
         * 
         * @return builder
         * 
         */
        public Builder maxSize(Integer maxSize) {
            return maxSize(Output.of(maxSize));
        }

        /**
         * @param name Index name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Index name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param nbShard Number of shards
         * 
         * @return builder
         * 
         */
        public Builder nbShard(@Nullable Output<Integer> nbShard) {
            $.nbShard = nbShard;
            return this;
        }

        /**
         * @param nbShard Number of shards
         * 
         * @return builder
         * 
         */
        public Builder nbShard(Integer nbShard) {
            return nbShard(Output.of(nbShard));
        }

        /**
         * @param serviceName The service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param suffix Index suffix
         * 
         * @return builder
         * 
         */
        public Builder suffix(@Nullable Output<String> suffix) {
            $.suffix = suffix;
            return this;
        }

        /**
         * @param suffix Index suffix
         * 
         * @return builder
         * 
         */
        public Builder suffix(String suffix) {
            return suffix(Output.of(suffix));
        }

        /**
         * @param updatedAt Index last update
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt Index last update
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        public LogsOutputOpenSearchIndexState build() {
            return $;
        }
    }

}
