// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.IpLoadBalancing;

import com.ovh.ovh.IpLoadBalancing.inputs.TcpFarmProbeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TcpFarmArgs extends com.pulumi.resources.ResourceArgs {

    public static final TcpFarmArgs Empty = new TcpFarmArgs();

    /**
     * Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
     * 
     */
    @Import(name="balance")
    private @Nullable Output<String> balance;

    /**
     * @return Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
     * 
     */
    public Optional<Output<String>> balance() {
        return Optional.ofNullable(this.balance);
    }

    /**
     * Readable label for loadbalancer farm
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Readable label for loadbalancer farm
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Port for backends to receive traffic on.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return Port for backends to receive traffic on.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
    }

    /**
     * define a backend healthcheck probe
     * 
     */
    @Import(name="probe")
    private @Nullable Output<TcpFarmProbeArgs> probe;

    /**
     * @return define a backend healthcheck probe
     * 
     */
    public Optional<Output<TcpFarmProbeArgs>> probe() {
        return Optional.ofNullable(this.probe);
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

    /**
     * Stickiness type. No stickiness if null (`sourceIp`)
     * 
     */
    @Import(name="stickiness")
    private @Nullable Output<String> stickiness;

    /**
     * @return Stickiness type. No stickiness if null (`sourceIp`)
     * 
     */
    public Optional<Output<String>> stickiness() {
        return Optional.ofNullable(this.stickiness);
    }

    /**
     * Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
     * 
     */
    @Import(name="vrackNetworkId")
    private @Nullable Output<Integer> vrackNetworkId;

    /**
     * @return Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
     * 
     */
    public Optional<Output<Integer>> vrackNetworkId() {
        return Optional.ofNullable(this.vrackNetworkId);
    }

    /**
     * Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
     * 
     */
    @Import(name="zone", required=true)
    private Output<String> zone;

    /**
     * @return Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    private TcpFarmArgs() {}

    private TcpFarmArgs(TcpFarmArgs $) {
        this.balance = $.balance;
        this.displayName = $.displayName;
        this.port = $.port;
        this.probe = $.probe;
        this.serviceName = $.serviceName;
        this.stickiness = $.stickiness;
        this.vrackNetworkId = $.vrackNetworkId;
        this.zone = $.zone;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TcpFarmArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TcpFarmArgs $;

        public Builder() {
            $ = new TcpFarmArgs();
        }

        public Builder(TcpFarmArgs defaults) {
            $ = new TcpFarmArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param balance Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
         * 
         * @return builder
         * 
         */
        public Builder balance(@Nullable Output<String> balance) {
            $.balance = balance;
            return this;
        }

        /**
         * @param balance Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
         * 
         * @return builder
         * 
         */
        public Builder balance(String balance) {
            return balance(Output.of(balance));
        }

        /**
         * @param displayName Readable label for loadbalancer farm
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Readable label for loadbalancer farm
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param port Port for backends to receive traffic on.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port Port for backends to receive traffic on.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
        }

        /**
         * @param probe define a backend healthcheck probe
         * 
         * @return builder
         * 
         */
        public Builder probe(@Nullable Output<TcpFarmProbeArgs> probe) {
            $.probe = probe;
            return this;
        }

        /**
         * @param probe define a backend healthcheck probe
         * 
         * @return builder
         * 
         */
        public Builder probe(TcpFarmProbeArgs probe) {
            return probe(Output.of(probe));
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

        /**
         * @param stickiness Stickiness type. No stickiness if null (`sourceIp`)
         * 
         * @return builder
         * 
         */
        public Builder stickiness(@Nullable Output<String> stickiness) {
            $.stickiness = stickiness;
            return this;
        }

        /**
         * @param stickiness Stickiness type. No stickiness if null (`sourceIp`)
         * 
         * @return builder
         * 
         */
        public Builder stickiness(String stickiness) {
            return stickiness(Output.of(stickiness));
        }

        /**
         * @param vrackNetworkId Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
         * 
         * @return builder
         * 
         */
        public Builder vrackNetworkId(@Nullable Output<Integer> vrackNetworkId) {
            $.vrackNetworkId = vrackNetworkId;
            return this;
        }

        /**
         * @param vrackNetworkId Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
         * 
         * @return builder
         * 
         */
        public Builder vrackNetworkId(Integer vrackNetworkId) {
            return vrackNetworkId(Output.of(vrackNetworkId));
        }

        /**
         * @param zone Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
         * 
         * @return builder
         * 
         */
        public Builder zone(Output<String> zone) {
            $.zone = zone;
            return this;
        }

        /**
         * @param zone Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
         * 
         * @return builder
         * 
         */
        public Builder zone(String zone) {
            return zone(Output.of(zone));
        }

        public TcpFarmArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("TcpFarmArgs", "serviceName");
            }
            if ($.zone == null) {
                throw new MissingRequiredPropertyException("TcpFarmArgs", "zone");
            }
            return $;
        }
    }

}
