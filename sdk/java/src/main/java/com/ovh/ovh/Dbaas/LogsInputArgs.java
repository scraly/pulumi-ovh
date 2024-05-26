// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dbaas;

import com.ovh.ovh.Dbaas.inputs.LogsInputConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LogsInputArgs extends com.pulumi.resources.ResourceArgs {

    public static final LogsInputArgs Empty = new LogsInputArgs();

    /**
     * List of IP blocks
     * 
     */
    @Import(name="allowedNetworks")
    private @Nullable Output<List<String>> allowedNetworks;

    /**
     * @return List of IP blocks
     * 
     */
    public Optional<Output<List<String>>> allowedNetworks() {
        return Optional.ofNullable(this.allowedNetworks);
    }

    /**
     * Input configuration
     * 
     */
    @Import(name="configuration", required=true)
    private Output<LogsInputConfigurationArgs> configuration;

    /**
     * @return Input configuration
     * 
     */
    public Output<LogsInputConfigurationArgs> configuration() {
        return this.configuration;
    }

    /**
     * Input description
     * 
     */
    @Import(name="description", required=true)
    private Output<String> description;

    /**
     * @return Input description
     * 
     */
    public Output<String> description() {
        return this.description;
    }

    /**
     * Input engine ID
     * 
     */
    @Import(name="engineId", required=true)
    private Output<String> engineId;

    /**
     * @return Input engine ID
     * 
     */
    public Output<String> engineId() {
        return this.engineId;
    }

    /**
     * Port
     * 
     */
    @Import(name="exposedPort")
    private @Nullable Output<String> exposedPort;

    /**
     * @return Port
     * 
     */
    public Optional<Output<String>> exposedPort() {
        return Optional.ofNullable(this.exposedPort);
    }

    /**
     * Number of instance running
     * 
     */
    @Import(name="nbInstance")
    private @Nullable Output<Integer> nbInstance;

    /**
     * @return Number of instance running
     * 
     */
    public Optional<Output<Integer>> nbInstance() {
        return Optional.ofNullable(this.nbInstance);
    }

    /**
     * service name
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     * Associated Graylog stream
     * 
     */
    @Import(name="streamId", required=true)
    private Output<String> streamId;

    /**
     * @return Associated Graylog stream
     * 
     */
    public Output<String> streamId() {
        return this.streamId;
    }

    /**
     * Input title
     * 
     */
    @Import(name="title", required=true)
    private Output<String> title;

    /**
     * @return Input title
     * 
     */
    public Output<String> title() {
        return this.title;
    }

    private LogsInputArgs() {}

    private LogsInputArgs(LogsInputArgs $) {
        this.allowedNetworks = $.allowedNetworks;
        this.configuration = $.configuration;
        this.description = $.description;
        this.engineId = $.engineId;
        this.exposedPort = $.exposedPort;
        this.nbInstance = $.nbInstance;
        this.serviceName = $.serviceName;
        this.streamId = $.streamId;
        this.title = $.title;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LogsInputArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LogsInputArgs $;

        public Builder() {
            $ = new LogsInputArgs();
        }

        public Builder(LogsInputArgs defaults) {
            $ = new LogsInputArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedNetworks List of IP blocks
         * 
         * @return builder
         * 
         */
        public Builder allowedNetworks(@Nullable Output<List<String>> allowedNetworks) {
            $.allowedNetworks = allowedNetworks;
            return this;
        }

        /**
         * @param allowedNetworks List of IP blocks
         * 
         * @return builder
         * 
         */
        public Builder allowedNetworks(List<String> allowedNetworks) {
            return allowedNetworks(Output.of(allowedNetworks));
        }

        /**
         * @param allowedNetworks List of IP blocks
         * 
         * @return builder
         * 
         */
        public Builder allowedNetworks(String... allowedNetworks) {
            return allowedNetworks(List.of(allowedNetworks));
        }

        /**
         * @param configuration Input configuration
         * 
         * @return builder
         * 
         */
        public Builder configuration(Output<LogsInputConfigurationArgs> configuration) {
            $.configuration = configuration;
            return this;
        }

        /**
         * @param configuration Input configuration
         * 
         * @return builder
         * 
         */
        public Builder configuration(LogsInputConfigurationArgs configuration) {
            return configuration(Output.of(configuration));
        }

        /**
         * @param description Input description
         * 
         * @return builder
         * 
         */
        public Builder description(Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Input description
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param engineId Input engine ID
         * 
         * @return builder
         * 
         */
        public Builder engineId(Output<String> engineId) {
            $.engineId = engineId;
            return this;
        }

        /**
         * @param engineId Input engine ID
         * 
         * @return builder
         * 
         */
        public Builder engineId(String engineId) {
            return engineId(Output.of(engineId));
        }

        /**
         * @param exposedPort Port
         * 
         * @return builder
         * 
         */
        public Builder exposedPort(@Nullable Output<String> exposedPort) {
            $.exposedPort = exposedPort;
            return this;
        }

        /**
         * @param exposedPort Port
         * 
         * @return builder
         * 
         */
        public Builder exposedPort(String exposedPort) {
            return exposedPort(Output.of(exposedPort));
        }

        /**
         * @param nbInstance Number of instance running
         * 
         * @return builder
         * 
         */
        public Builder nbInstance(@Nullable Output<Integer> nbInstance) {
            $.nbInstance = nbInstance;
            return this;
        }

        /**
         * @param nbInstance Number of instance running
         * 
         * @return builder
         * 
         */
        public Builder nbInstance(Integer nbInstance) {
            return nbInstance(Output.of(nbInstance));
        }

        /**
         * @param serviceName service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param streamId Associated Graylog stream
         * 
         * @return builder
         * 
         */
        public Builder streamId(Output<String> streamId) {
            $.streamId = streamId;
            return this;
        }

        /**
         * @param streamId Associated Graylog stream
         * 
         * @return builder
         * 
         */
        public Builder streamId(String streamId) {
            return streamId(Output.of(streamId));
        }

        /**
         * @param title Input title
         * 
         * @return builder
         * 
         */
        public Builder title(Output<String> title) {
            $.title = title;
            return this;
        }

        /**
         * @param title Input title
         * 
         * @return builder
         * 
         */
        public Builder title(String title) {
            return title(Output.of(title));
        }

        public LogsInputArgs build() {
            if ($.configuration == null) {
                throw new MissingRequiredPropertyException("LogsInputArgs", "configuration");
            }
            if ($.description == null) {
                throw new MissingRequiredPropertyException("LogsInputArgs", "description");
            }
            if ($.engineId == null) {
                throw new MissingRequiredPropertyException("LogsInputArgs", "engineId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("LogsInputArgs", "serviceName");
            }
            if ($.streamId == null) {
                throw new MissingRequiredPropertyException("LogsInputArgs", "streamId");
            }
            if ($.title == null) {
                throw new MissingRequiredPropertyException("LogsInputArgs", "title");
            }
            return $;
        }
    }

}
