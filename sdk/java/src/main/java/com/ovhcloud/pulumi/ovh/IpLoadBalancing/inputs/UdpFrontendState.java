// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.IpLoadBalancing.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UdpFrontendState extends com.pulumi.resources.ResourceArgs {

    public static final UdpFrontendState Empty = new UdpFrontendState();

    /**
     * Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     * 
     */
    @Import(name="dedicatedIpfos")
    private @Nullable Output<List<String>> dedicatedIpfos;

    /**
     * @return Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     * 
     */
    public Optional<Output<List<String>>> dedicatedIpfos() {
        return Optional.ofNullable(this.dedicatedIpfos);
    }

    /**
     * Default UDP Farm of your frontend
     * 
     */
    @Import(name="defaultFarmId")
    private @Nullable Output<Double> defaultFarmId;

    /**
     * @return Default UDP Farm of your frontend
     * 
     */
    public Optional<Output<Double>> defaultFarmId() {
        return Optional.ofNullable(this.defaultFarmId);
    }

    /**
     * Disable your frontend. Default: &#39;false&#39;
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return Disable your frontend. Default: &#39;false&#39;
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * Human readable name for your frontend
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Human readable name for your frontend
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Id of your frontend
     * 
     */
    @Import(name="frontendId")
    private @Nullable Output<Double> frontendId;

    /**
     * @return Id of your frontend
     * 
     */
    public Optional<Output<Double>> frontendId() {
        return Optional.ofNullable(this.frontendId);
    }

    /**
     * Port(s) attached to your frontend. Supports single port (numerical value),
     * range (2 dash-delimited increasing ports) and comma-separated list of &#39;single port&#39;
     * and/or &#39;range&#39;. Each port must be in the [1;49151] range
     * 
     */
    @Import(name="port")
    private @Nullable Output<String> port;

    /**
     * @return Port(s) attached to your frontend. Supports single port (numerical value),
     * range (2 dash-delimited increasing ports) and comma-separated list of &#39;single port&#39;
     * and/or &#39;range&#39;. Each port must be in the [1;49151] range
     * 
     */
    public Optional<Output<String>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * The internal name of your IP load balancing
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The internal name of your IP load balancing
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     * 
     */
    @Import(name="zone")
    private @Nullable Output<String> zone;

    /**
     * @return Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     * 
     */
    public Optional<Output<String>> zone() {
        return Optional.ofNullable(this.zone);
    }

    private UdpFrontendState() {}

    private UdpFrontendState(UdpFrontendState $) {
        this.dedicatedIpfos = $.dedicatedIpfos;
        this.defaultFarmId = $.defaultFarmId;
        this.disabled = $.disabled;
        this.displayName = $.displayName;
        this.frontendId = $.frontendId;
        this.port = $.port;
        this.serviceName = $.serviceName;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UdpFrontendState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UdpFrontendState $;

        public Builder() {
            $ = new UdpFrontendState();
        }

        public Builder(UdpFrontendState defaults) {
            $ = new UdpFrontendState(Objects.requireNonNull(defaults));
        }

        /**
         * @param dedicatedIpfos Only attach frontend on these ip. No restriction if null. List of Ip blocks.
         * 
         * @return builder
         * 
         */
        public Builder dedicatedIpfos(@Nullable Output<List<String>> dedicatedIpfos) {
            $.dedicatedIpfos = dedicatedIpfos;
            return this;
        }

        /**
         * @param dedicatedIpfos Only attach frontend on these ip. No restriction if null. List of Ip blocks.
         * 
         * @return builder
         * 
         */
        public Builder dedicatedIpfos(List<String> dedicatedIpfos) {
            return dedicatedIpfos(Output.of(dedicatedIpfos));
        }

        /**
         * @param dedicatedIpfos Only attach frontend on these ip. No restriction if null. List of Ip blocks.
         * 
         * @return builder
         * 
         */
        public Builder dedicatedIpfos(String... dedicatedIpfos) {
            return dedicatedIpfos(List.of(dedicatedIpfos));
        }

        /**
         * @param defaultFarmId Default UDP Farm of your frontend
         * 
         * @return builder
         * 
         */
        public Builder defaultFarmId(@Nullable Output<Double> defaultFarmId) {
            $.defaultFarmId = defaultFarmId;
            return this;
        }

        /**
         * @param defaultFarmId Default UDP Farm of your frontend
         * 
         * @return builder
         * 
         */
        public Builder defaultFarmId(Double defaultFarmId) {
            return defaultFarmId(Output.of(defaultFarmId));
        }

        /**
         * @param disabled Disable your frontend. Default: &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled Disable your frontend. Default: &#39;false&#39;
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param displayName Human readable name for your frontend
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Human readable name for your frontend
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param frontendId Id of your frontend
         * 
         * @return builder
         * 
         */
        public Builder frontendId(@Nullable Output<Double> frontendId) {
            $.frontendId = frontendId;
            return this;
        }

        /**
         * @param frontendId Id of your frontend
         * 
         * @return builder
         * 
         */
        public Builder frontendId(Double frontendId) {
            return frontendId(Output.of(frontendId));
        }

        /**
         * @param port Port(s) attached to your frontend. Supports single port (numerical value),
         * range (2 dash-delimited increasing ports) and comma-separated list of &#39;single port&#39;
         * and/or &#39;range&#39;. Each port must be in the [1;49151] range
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<String> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port(s) attached to your frontend. Supports single port (numerical value),
         * range (2 dash-delimited increasing ports) and comma-separated list of &#39;single port&#39;
         * and/or &#39;range&#39;. Each port must be in the [1;49151] range
         * 
         * @return builder
         * 
         */
        public Builder port(String port) {
            return port(Output.of(port));
        }

        /**
         * @param serviceName The internal name of your IP load balancing
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
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

        /**
         * @param zone Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
         * 
         * @return builder
         * 
         */
        public Builder zone(@Nullable Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public UdpFrontendState build() {
            return $;
        }
    }

}