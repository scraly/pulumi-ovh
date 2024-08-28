// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dbaas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LogsOutputGraylogStreamArgs extends com.pulumi.resources.ResourceArgs {

    public static final LogsOutputGraylogStreamArgs Empty = new LogsOutputGraylogStreamArgs();

    /**
     * Cold storage compression method
     * 
     */
    @Import(name="coldStorageCompression")
    private @Nullable Output<String> coldStorageCompression;

    /**
     * @return Cold storage compression method
     * 
     */
    public Optional<Output<String>> coldStorageCompression() {
        return Optional.ofNullable(this.coldStorageCompression);
    }

    /**
     * ColdStorage content
     * 
     */
    @Import(name="coldStorageContent")
    private @Nullable Output<String> coldStorageContent;

    /**
     * @return ColdStorage content
     * 
     */
    public Optional<Output<String>> coldStorageContent() {
        return Optional.ofNullable(this.coldStorageContent);
    }

    /**
     * Is Cold storage enabled?
     * 
     */
    @Import(name="coldStorageEnabled")
    private @Nullable Output<Boolean> coldStorageEnabled;

    /**
     * @return Is Cold storage enabled?
     * 
     */
    public Optional<Output<Boolean>> coldStorageEnabled() {
        return Optional.ofNullable(this.coldStorageEnabled);
    }

    /**
     * Notify on new Cold storage archive
     * 
     */
    @Import(name="coldStorageNotifyEnabled")
    private @Nullable Output<Boolean> coldStorageNotifyEnabled;

    /**
     * @return Notify on new Cold storage archive
     * 
     */
    public Optional<Output<Boolean>> coldStorageNotifyEnabled() {
        return Optional.ofNullable(this.coldStorageNotifyEnabled);
    }

    /**
     * Cold storage retention in year
     * 
     */
    @Import(name="coldStorageRetention")
    private @Nullable Output<Integer> coldStorageRetention;

    /**
     * @return Cold storage retention in year
     * 
     */
    public Optional<Output<Integer>> coldStorageRetention() {
        return Optional.ofNullable(this.coldStorageRetention);
    }

    /**
     * ColdStorage destination
     * 
     */
    @Import(name="coldStorageTarget")
    private @Nullable Output<String> coldStorageTarget;

    /**
     * @return ColdStorage destination
     * 
     */
    public Optional<Output<String>> coldStorageTarget() {
        return Optional.ofNullable(this.coldStorageTarget);
    }

    /**
     * Stream description
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return Stream description
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * Enable ES indexing
     * 
     */
    @Import(name="indexingEnabled")
    private @Nullable Output<Boolean> indexingEnabled;

    /**
     * @return Enable ES indexing
     * 
     */
    public Optional<Output<Boolean>> indexingEnabled() {
        return Optional.ofNullable(this.indexingEnabled);
    }

    /**
     * Maximum indexing size (in GB)
     * 
     */
    @Import(name="indexingMaxSize")
    private @Nullable Output<Integer> indexingMaxSize;

    /**
     * @return Maximum indexing size (in GB)
     * 
     */
    public Optional<Output<Integer>> indexingMaxSize() {
        return Optional.ofNullable(this.indexingMaxSize);
    }

    /**
     * If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
     * 
     */
    @Import(name="indexingNotifyEnabled")
    private @Nullable Output<Boolean> indexingNotifyEnabled;

    /**
     * @return If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
     * 
     */
    public Optional<Output<Boolean>> indexingNotifyEnabled() {
        return Optional.ofNullable(this.indexingNotifyEnabled);
    }

    /**
     * Parent stream ID
     * 
     */
    @Import(name="parentStreamId")
    private @Nullable Output<String> parentStreamId;

    /**
     * @return Parent stream ID
     * 
     */
    public Optional<Output<String>> parentStreamId() {
        return Optional.ofNullable(this.parentStreamId);
    }

    /**
     * If set, pause indexing when maximum size is reach
     * 
     */
    @Import(name="pauseIndexingOnMaxSize")
    private @Nullable Output<Boolean> pauseIndexingOnMaxSize;

    /**
     * @return If set, pause indexing when maximum size is reach
     * 
     */
    public Optional<Output<Boolean>> pauseIndexingOnMaxSize() {
        return Optional.ofNullable(this.pauseIndexingOnMaxSize);
    }

    /**
     * Retention ID
     * 
     */
    @Import(name="retentionId")
    private @Nullable Output<String> retentionId;

    /**
     * @return Retention ID
     * 
     */
    public Optional<Output<String>> retentionId() {
        return Optional.ofNullable(this.retentionId);
    }

    /**
     * The service name
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     * Stream description
     * 
     */
    @Import(name="title", required=true)
    private Output<String> title;

    /**
     * @return Stream description
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    /**
     * Enable Websocket
     * 
     */
    @Import(name="webSocketEnabled")
    private @Nullable Output<Boolean> webSocketEnabled;

    /**
     * @return Enable Websocket
     * 
     */
    public Optional<Output<Boolean>> webSocketEnabled() {
        return Optional.ofNullable(this.webSocketEnabled);
    }

    private LogsOutputGraylogStreamArgs() {}

    private LogsOutputGraylogStreamArgs(LogsOutputGraylogStreamArgs $) {
        this.coldStorageCompression = $.coldStorageCompression;
        this.coldStorageContent = $.coldStorageContent;
        this.coldStorageEnabled = $.coldStorageEnabled;
        this.coldStorageNotifyEnabled = $.coldStorageNotifyEnabled;
        this.coldStorageRetention = $.coldStorageRetention;
        this.coldStorageTarget = $.coldStorageTarget;
        this.description = $.description;
        this.indexingEnabled = $.indexingEnabled;
        this.indexingMaxSize = $.indexingMaxSize;
        this.indexingNotifyEnabled = $.indexingNotifyEnabled;
        this.parentStreamId = $.parentStreamId;
        this.pauseIndexingOnMaxSize = $.pauseIndexingOnMaxSize;
        this.retentionId = $.retentionId;
        this.serviceName = $.serviceName;
        this.title = $.title;
        this.webSocketEnabled = $.webSocketEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LogsOutputGraylogStreamArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LogsOutputGraylogStreamArgs $;

        public Builder() {
            $ = new LogsOutputGraylogStreamArgs();
        }

        public Builder(LogsOutputGraylogStreamArgs defaults) {
            $ = new LogsOutputGraylogStreamArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param coldStorageCompression Cold storage compression method
         * 
         * @return builder
         * 
         */
        public Builder coldStorageCompression(@Nullable Output<String> coldStorageCompression) {
            $.coldStorageCompression = coldStorageCompression;
            return this;
        }

        /**
         * @param coldStorageCompression Cold storage compression method
         * 
         * @return builder
         * 
         */
        public Builder coldStorageCompression(String coldStorageCompression) {
            return coldStorageCompression(Output.of(coldStorageCompression));
        }

        /**
         * @param coldStorageContent ColdStorage content
         * 
         * @return builder
         * 
         */
        public Builder coldStorageContent(@Nullable Output<String> coldStorageContent) {
            $.coldStorageContent = coldStorageContent;
            return this;
        }

        /**
         * @param coldStorageContent ColdStorage content
         * 
         * @return builder
         * 
         */
        public Builder coldStorageContent(String coldStorageContent) {
            return coldStorageContent(Output.of(coldStorageContent));
        }

        /**
         * @param coldStorageEnabled Is Cold storage enabled?
         * 
         * @return builder
         * 
         */
        public Builder coldStorageEnabled(@Nullable Output<Boolean> coldStorageEnabled) {
            $.coldStorageEnabled = coldStorageEnabled;
            return this;
        }

        /**
         * @param coldStorageEnabled Is Cold storage enabled?
         * 
         * @return builder
         * 
         */
        public Builder coldStorageEnabled(Boolean coldStorageEnabled) {
            return coldStorageEnabled(Output.of(coldStorageEnabled));
        }

        /**
         * @param coldStorageNotifyEnabled Notify on new Cold storage archive
         * 
         * @return builder
         * 
         */
        public Builder coldStorageNotifyEnabled(@Nullable Output<Boolean> coldStorageNotifyEnabled) {
            $.coldStorageNotifyEnabled = coldStorageNotifyEnabled;
            return this;
        }

        /**
         * @param coldStorageNotifyEnabled Notify on new Cold storage archive
         * 
         * @return builder
         * 
         */
        public Builder coldStorageNotifyEnabled(Boolean coldStorageNotifyEnabled) {
            return coldStorageNotifyEnabled(Output.of(coldStorageNotifyEnabled));
        }

        /**
         * @param coldStorageRetention Cold storage retention in year
         * 
         * @return builder
         * 
         */
        public Builder coldStorageRetention(@Nullable Output<Integer> coldStorageRetention) {
            $.coldStorageRetention = coldStorageRetention;
            return this;
        }

        /**
         * @param coldStorageRetention Cold storage retention in year
         * 
         * @return builder
         * 
         */
        public Builder coldStorageRetention(Integer coldStorageRetention) {
            return coldStorageRetention(Output.of(coldStorageRetention));
        }

        /**
         * @param coldStorageTarget ColdStorage destination
         * 
         * @return builder
         * 
         */
        public Builder coldStorageTarget(@Nullable Output<String> coldStorageTarget) {
            $.coldStorageTarget = coldStorageTarget;
            return this;
        }

        /**
         * @param coldStorageTarget ColdStorage destination
         * 
         * @return builder
         * 
         */
        public Builder coldStorageTarget(String coldStorageTarget) {
            return coldStorageTarget(Output.of(coldStorageTarget));
        }

        /**
         * @param description Stream description
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Stream description
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param indexingEnabled Enable ES indexing
         * 
         * @return builder
         * 
         */
        public Builder indexingEnabled(@Nullable Output<Boolean> indexingEnabled) {
            $.indexingEnabled = indexingEnabled;
            return this;
        }

        /**
         * @param indexingEnabled Enable ES indexing
         * 
         * @return builder
         * 
         */
        public Builder indexingEnabled(Boolean indexingEnabled) {
            return indexingEnabled(Output.of(indexingEnabled));
        }

        /**
         * @param indexingMaxSize Maximum indexing size (in GB)
         * 
         * @return builder
         * 
         */
        public Builder indexingMaxSize(@Nullable Output<Integer> indexingMaxSize) {
            $.indexingMaxSize = indexingMaxSize;
            return this;
        }

        /**
         * @param indexingMaxSize Maximum indexing size (in GB)
         * 
         * @return builder
         * 
         */
        public Builder indexingMaxSize(Integer indexingMaxSize) {
            return indexingMaxSize(Output.of(indexingMaxSize));
        }

        /**
         * @param indexingNotifyEnabled If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
         * 
         * @return builder
         * 
         */
        public Builder indexingNotifyEnabled(@Nullable Output<Boolean> indexingNotifyEnabled) {
            $.indexingNotifyEnabled = indexingNotifyEnabled;
            return this;
        }

        /**
         * @param indexingNotifyEnabled If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
         * 
         * @return builder
         * 
         */
        public Builder indexingNotifyEnabled(Boolean indexingNotifyEnabled) {
            return indexingNotifyEnabled(Output.of(indexingNotifyEnabled));
        }

        /**
         * @param parentStreamId Parent stream ID
         * 
         * @return builder
         * 
         */
        public Builder parentStreamId(@Nullable Output<String> parentStreamId) {
            $.parentStreamId = parentStreamId;
            return this;
        }

        /**
         * @param parentStreamId Parent stream ID
         * 
         * @return builder
         * 
         */
        public Builder parentStreamId(String parentStreamId) {
            return parentStreamId(Output.of(parentStreamId));
        }

        /**
         * @param pauseIndexingOnMaxSize If set, pause indexing when maximum size is reach
         * 
         * @return builder
         * 
         */
        public Builder pauseIndexingOnMaxSize(@Nullable Output<Boolean> pauseIndexingOnMaxSize) {
            $.pauseIndexingOnMaxSize = pauseIndexingOnMaxSize;
            return this;
        }

        /**
         * @param pauseIndexingOnMaxSize If set, pause indexing when maximum size is reach
         * 
         * @return builder
         * 
         */
        public Builder pauseIndexingOnMaxSize(Boolean pauseIndexingOnMaxSize) {
            return pauseIndexingOnMaxSize(Output.of(pauseIndexingOnMaxSize));
        }

        /**
         * @param retentionId Retention ID
         * 
         * @return builder
         * 
         */
        public Builder retentionId(@Nullable Output<String> retentionId) {
            $.retentionId = retentionId;
            return this;
        }

        /**
         * @param retentionId Retention ID
         * 
         * @return builder
         * 
         */
        public Builder retentionId(String retentionId) {
            return retentionId(Output.of(retentionId));
        }

        /**
         * @param serviceName The service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
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
         * @param title Stream description
         * 
         * @return builder
         * 
         */
        public Builder title(Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title Stream description
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        /**
         * @param webSocketEnabled Enable Websocket
         * 
         * @return builder
         * 
         */
        public Builder webSocketEnabled(@Nullable Output<Boolean> webSocketEnabled) {
            $.webSocketEnabled = webSocketEnabled;
            return this;
        }

        /**
         * @param webSocketEnabled Enable Websocket
         * 
         * @return builder
         * 
         */
        public Builder webSocketEnabled(Boolean webSocketEnabled) {
            return webSocketEnabled(Output.of(webSocketEnabled));
        }

        public LogsOutputGraylogStreamArgs build() {
            if ($.description == null) {
                throw new MissingRequiredPropertyException("LogsOutputGraylogStreamArgs", "description");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("LogsOutputGraylogStreamArgs", "serviceName");
            }
            if ($.title == null) {
                throw new MissingRequiredPropertyException("LogsOutputGraylogStreamArgs", "title");
            }
            return $;
        }
    }

}
