// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dbaas.inputs;

import com.ovhcloud.pulumi.ovh.Dbaas.inputs.LogsInputConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LogsInputState extends com.pulumi.resources.ResourceArgs {

    public static final LogsInputState Empty = new LogsInputState();

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
     * Whether the workload is auto-scaled (mutually exclusive with parameter `nb_instance`)
     * 
     */
    @Import(name="autoscale")
    private @Nullable Output<Boolean> autoscale;

    /**
     * @return Whether the workload is auto-scaled (mutually exclusive with parameter `nb_instance`)
     * 
     */
    public Optional<Output<Boolean>> autoscale() {
        return Optional.ofNullable(this.autoscale);
    }

    /**
     * Input configuration
     * 
     */
    @Import(name="configuration")
    private @Nullable Output<LogsInputConfigurationArgs> configuration;

    /**
     * @return Input configuration
     * 
     */
    public Optional<Output<LogsInputConfigurationArgs>> configuration() {
        return Optional.ofNullable(this.configuration);
    }

    /**
     * Input creation
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Input creation
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * Number of instance running (returned by the API)
     * 
     */
    @Import(name="currentNbInstance")
    private @Nullable Output<Integer> currentNbInstance;

    /**
     * @return Number of instance running (returned by the API)
     * 
     */
    public Optional<Output<Integer>> currentNbInstance() {
        return Optional.ofNullable(this.currentNbInstance);
    }

    /**
     * Input description
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Input description
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Input engine ID
     * 
     */
    @Import(name="engineId")
    private @Nullable Output<String> engineId;

    /**
     * @return Input engine ID
     * 
     */
    public Optional<Output<String>> engineId() {
        return Optional.ofNullable(this.engineId);
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
     * Hostname
     * 
     */
    @Import(name="hostname")
    private @Nullable Output<String> hostname;

    /**
     * @return Hostname
     * 
     */
    public Optional<Output<String>> hostname() {
        return Optional.ofNullable(this.hostname);
    }

    /**
     * Input ID
     * 
     */
    @Import(name="inputId")
    private @Nullable Output<String> inputId;

    /**
     * @return Input ID
     * 
     */
    public Optional<Output<String>> inputId() {
        return Optional.ofNullable(this.inputId);
    }

    /**
     * Indicate if input need to be restarted
     * 
     */
    @Import(name="isRestartRequired")
    private @Nullable Output<Boolean> isRestartRequired;

    /**
     * @return Indicate if input need to be restarted
     * 
     */
    public Optional<Output<Boolean>> isRestartRequired() {
        return Optional.ofNullable(this.isRestartRequired);
    }

    /**
     * Maximum number of instances in auto-scaled mode
     * 
     */
    @Import(name="maxScaleInstance")
    private @Nullable Output<Integer> maxScaleInstance;

    /**
     * @return Maximum number of instances in auto-scaled mode
     * 
     */
    public Optional<Output<Integer>> maxScaleInstance() {
        return Optional.ofNullable(this.maxScaleInstance);
    }

    /**
     * Minimum number of instances in auto-scaled mode
     * 
     */
    @Import(name="minScaleInstance")
    private @Nullable Output<Integer> minScaleInstance;

    /**
     * @return Minimum number of instances in auto-scaled mode
     * 
     */
    public Optional<Output<Integer>> minScaleInstance() {
        return Optional.ofNullable(this.minScaleInstance);
    }

    /**
     * Number of instance running (input, mutually exclusive with parameter `autoscale`)
     * 
     */
    @Import(name="nbInstance")
    private @Nullable Output<Integer> nbInstance;

    /**
     * @return Number of instance running (input, mutually exclusive with parameter `autoscale`)
     * 
     */
    public Optional<Output<Integer>> nbInstance() {
        return Optional.ofNullable(this.nbInstance);
    }

    /**
     * Input IP address
     * 
     */
    @Import(name="publicAddress")
    private @Nullable Output<String> publicAddress;

    /**
     * @return Input IP address
     * 
     */
    public Optional<Output<String>> publicAddress() {
        return Optional.ofNullable(this.publicAddress);
    }

    /**
     * service name
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return service name
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Input SSL certificate
     * 
     */
    @Import(name="sslCertificate")
    private @Nullable Output<String> sslCertificate;

    /**
     * @return Input SSL certificate
     * 
     */
    public Optional<Output<String>> sslCertificate() {
        return Optional.ofNullable(this.sslCertificate);
    }

    /**
     * init: configuration required, pending: ready to start, running: available
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return init: configuration required, pending: ready to start, running: available
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Associated Graylog stream
     * 
     */
    @Import(name="streamId")
    private @Nullable Output<String> streamId;

    /**
     * @return Associated Graylog stream
     * 
     */
    public Optional<Output<String>> streamId() {
        return Optional.ofNullable(this.streamId);
    }

    /**
     * Input title
     * 
     */
    @Import(name="title")
    private @Nullable Output<String> title;

    /**
     * @return Input title
     * 
     */
    public Optional<Output<String>> title() {
        return Optional.ofNullable(this.title);
    }

    /**
     * Input last update
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return Input last update
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    private LogsInputState() {}

    private LogsInputState(LogsInputState $) {
        this.allowedNetworks = $.allowedNetworks;
        this.autoscale = $.autoscale;
        this.configuration = $.configuration;
        this.createdAt = $.createdAt;
        this.currentNbInstance = $.currentNbInstance;
        this.description = $.description;
        this.engineId = $.engineId;
        this.exposedPort = $.exposedPort;
        this.hostname = $.hostname;
        this.inputId = $.inputId;
        this.isRestartRequired = $.isRestartRequired;
        this.maxScaleInstance = $.maxScaleInstance;
        this.minScaleInstance = $.minScaleInstance;
        this.nbInstance = $.nbInstance;
        this.publicAddress = $.publicAddress;
        this.serviceName = $.serviceName;
        this.sslCertificate = $.sslCertificate;
        this.status = $.status;
        this.streamId = $.streamId;
        this.title = $.title;
        this.updatedAt = $.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LogsInputState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LogsInputState $;

        public Builder() {
            $ = new LogsInputState();
        }

        public Builder(LogsInputState defaults) {
            $ = new LogsInputState(Objects.requireNonNull(defaults));
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
         * @param autoscale Whether the workload is auto-scaled (mutually exclusive with parameter `nb_instance`)
         * 
         * @return builder
         * 
         */
        public Builder autoscale(@Nullable Output<Boolean> autoscale) {
            $.autoscale = autoscale;
            return this;
        }

        /**
         * @param autoscale Whether the workload is auto-scaled (mutually exclusive with parameter `nb_instance`)
         * 
         * @return builder
         * 
         */
        public Builder autoscale(Boolean autoscale) {
            return autoscale(Output.of(autoscale));
        }

        /**
         * @param configuration Input configuration
         * 
         * @return builder
         * 
         */
        public Builder configuration(@Nullable Output<LogsInputConfigurationArgs> configuration) {
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
         * @param createdAt Input creation
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Input creation
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param currentNbInstance Number of instance running (returned by the API)
         * 
         * @return builder
         * 
         */
        public Builder currentNbInstance(@Nullable Output<Integer> currentNbInstance) {
            $.currentNbInstance = currentNbInstance;
            return this;
        }

        /**
         * @param currentNbInstance Number of instance running (returned by the API)
         * 
         * @return builder
         * 
         */
        public Builder currentNbInstance(Integer currentNbInstance) {
            return currentNbInstance(Output.of(currentNbInstance));
        }

        /**
         * @param description Input description
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
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
        public Builder engineId(@Nullable Output<String> engineId) {
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
         * @param hostname Hostname
         * 
         * @return builder
         * 
         */
        public Builder hostname(@Nullable Output<String> hostname) {
            $.hostname = hostname;
            return this;
        }

        /**
         * @param hostname Hostname
         * 
         * @return builder
         * 
         */
        public Builder hostname(String hostname) {
            return hostname(Output.of(hostname));
        }

        /**
         * @param inputId Input ID
         * 
         * @return builder
         * 
         */
        public Builder inputId(@Nullable Output<String> inputId) {
            $.inputId = inputId;
            return this;
        }

        /**
         * @param inputId Input ID
         * 
         * @return builder
         * 
         */
        public Builder inputId(String inputId) {
            return inputId(Output.of(inputId));
        }

        /**
         * @param isRestartRequired Indicate if input need to be restarted
         * 
         * @return builder
         * 
         */
        public Builder isRestartRequired(@Nullable Output<Boolean> isRestartRequired) {
            $.isRestartRequired = isRestartRequired;
            return this;
        }

        /**
         * @param isRestartRequired Indicate if input need to be restarted
         * 
         * @return builder
         * 
         */
        public Builder isRestartRequired(Boolean isRestartRequired) {
            return isRestartRequired(Output.of(isRestartRequired));
        }

        /**
         * @param maxScaleInstance Maximum number of instances in auto-scaled mode
         * 
         * @return builder
         * 
         */
        public Builder maxScaleInstance(@Nullable Output<Integer> maxScaleInstance) {
            $.maxScaleInstance = maxScaleInstance;
            return this;
        }

        /**
         * @param maxScaleInstance Maximum number of instances in auto-scaled mode
         * 
         * @return builder
         * 
         */
        public Builder maxScaleInstance(Integer maxScaleInstance) {
            return maxScaleInstance(Output.of(maxScaleInstance));
        }

        /**
         * @param minScaleInstance Minimum number of instances in auto-scaled mode
         * 
         * @return builder
         * 
         */
        public Builder minScaleInstance(@Nullable Output<Integer> minScaleInstance) {
            $.minScaleInstance = minScaleInstance;
            return this;
        }

        /**
         * @param minScaleInstance Minimum number of instances in auto-scaled mode
         * 
         * @return builder
         * 
         */
        public Builder minScaleInstance(Integer minScaleInstance) {
            return minScaleInstance(Output.of(minScaleInstance));
        }

        /**
         * @param nbInstance Number of instance running (input, mutually exclusive with parameter `autoscale`)
         * 
         * @return builder
         * 
         */
        public Builder nbInstance(@Nullable Output<Integer> nbInstance) {
            $.nbInstance = nbInstance;
            return this;
        }

        /**
         * @param nbInstance Number of instance running (input, mutually exclusive with parameter `autoscale`)
         * 
         * @return builder
         * 
         */
        public Builder nbInstance(Integer nbInstance) {
            return nbInstance(Output.of(nbInstance));
        }

        /**
         * @param publicAddress Input IP address
         * 
         * @return builder
         * 
         */
        public Builder publicAddress(@Nullable Output<String> publicAddress) {
            $.publicAddress = publicAddress;
            return this;
        }

        /**
         * @param publicAddress Input IP address
         * 
         * @return builder
         * 
         */
        public Builder publicAddress(String publicAddress) {
            return publicAddress(Output.of(publicAddress));
        }

        /**
         * @param serviceName service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
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
         * @param sslCertificate Input SSL certificate
         * 
         * @return builder
         * 
         */
        public Builder sslCertificate(@Nullable Output<String> sslCertificate) {
            $.sslCertificate = sslCertificate;
            return this;
        }

        /**
         * @param sslCertificate Input SSL certificate
         * 
         * @return builder
         * 
         */
        public Builder sslCertificate(String sslCertificate) {
            return sslCertificate(Output.of(sslCertificate));
        }

        /**
         * @param status init: configuration required, pending: ready to start, running: available
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status init: configuration required, pending: ready to start, running: available
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param streamId Associated Graylog stream
         * 
         * @return builder
         * 
         */
        public Builder streamId(@Nullable Output<String> streamId) {
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
        public Builder title(@Nullable Output<String> title) {
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

        /**
         * @param updatedAt Input last update
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt Input last update
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        public LogsInputState build() {
            return $;
        }
    }

}
