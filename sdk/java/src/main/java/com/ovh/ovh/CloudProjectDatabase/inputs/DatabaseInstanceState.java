// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatabaseInstanceState extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseInstanceState Empty = new DatabaseInstanceState();

    /**
     * Cluster ID.
     * 
     */
    @Import(name="clusterId")
    private @Nullable Output<String> clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public Optional<Output<String>> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }

    /**
     * Defines if the database has been created by default.
     * 
     */
    @Import(name="default")
    private @Nullable Output<Boolean> default_;

    /**
     * @return Defines if the database has been created by default.
     * 
     */
    public Optional<Output<Boolean>> default_() {
        return Optional.ofNullable(this.default_);
    }

    /**
     * The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * Available engines:
     * 
     */
    @Import(name="engine")
    private @Nullable Output<String> engine;

    /**
     * @return The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * Available engines:
     * 
     */
    public Optional<Output<String>> engine() {
        return Optional.ofNullable(this.engine);
    }

    /**
     * Name of the database.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the database.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    private DatabaseInstanceState() {}

    private DatabaseInstanceState(DatabaseInstanceState $) {
        this.clusterId = $.clusterId;
        this.default_ = $.default_;
        this.engine = $.engine;
        this.name = $.name;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseInstanceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseInstanceState $;

        public Builder() {
            $ = new DatabaseInstanceState();
        }

        public Builder(DatabaseInstanceState defaults) {
            $ = new DatabaseInstanceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(@Nullable Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param default_ Defines if the database has been created by default.
         * 
         * @return builder
         * 
         */
        public Builder default_(@Nullable Output<Boolean> default_) {
            $.default_ = default_;
            return this;
        }

        /**
         * @param default_ Defines if the database has been created by default.
         * 
         * @return builder
         * 
         */
        public Builder default_(Boolean default_) {
            return default_(Output.of(default_));
        }

        /**
         * @param engine The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * Available engines:
         * 
         * @return builder
         * 
         */
        public Builder engine(@Nullable Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * Available engines:
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param name Name of the database.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the database.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public DatabaseInstanceState build() {
            return $;
        }
    }

}
