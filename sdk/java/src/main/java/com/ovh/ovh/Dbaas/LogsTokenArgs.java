// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Dbaas;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class LogsTokenArgs extends com.pulumi.resources.ResourceArgs {

    public static final LogsTokenArgs Empty = new LogsTokenArgs();

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
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The LDP service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private LogsTokenArgs() {}

    private LogsTokenArgs(LogsTokenArgs $) {
        this.clusterId = $.clusterId;
        this.name = $.name;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(LogsTokenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private LogsTokenArgs $;

        public Builder() {
            $ = new LogsTokenArgs();
        }

        public Builder(LogsTokenArgs defaults) {
            $ = new LogsTokenArgs(Objects.requireNonNull(defaults));
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
        public Builder serviceName(Output<String> serviceName) {
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

        public LogsTokenArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("LogsTokenArgs", "serviceName");
            }
            return $;
        }
    }

}
