// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetDatabaseInstancePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatabaseInstancePlainArgs Empty = new GetDatabaseInstancePlainArgs();

    /**
     * Cluster ID
     * 
     */
    @Import(name="clusterId", required=true)
    private String clusterId;

    /**
     * @return Cluster ID
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }

    /**
     * The engine of the database cluster you want database information. To get a full list of available engine visit:
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * Available engines:
     * 
     */
    @Import(name="engine", required=true)
    private String engine;

    /**
     * @return The engine of the database cluster you want database information. To get a full list of available engine visit:
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * Available engines:
     * 
     */
    public String engine() {
        return this.engine;
    }

    /**
     * Name of the database.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name of the database.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetDatabaseInstancePlainArgs() {}

    private GetDatabaseInstancePlainArgs(GetDatabaseInstancePlainArgs $) {
        this.clusterId = $.clusterId;
        this.engine = $.engine;
        this.name = $.name;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatabaseInstancePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatabaseInstancePlainArgs $;

        public Builder() {
            $ = new GetDatabaseInstancePlainArgs();
        }

        public Builder(GetDatabaseInstancePlainArgs defaults) {
            $ = new GetDatabaseInstancePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param engine The engine of the database cluster you want database information. To get a full list of available engine visit:
         * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * Available engines:
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param name Name of the database.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
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
            $.serviceName = serviceName;
            return this;
        }

        public GetDatabaseInstancePlainArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("GetDatabaseInstancePlainArgs", "clusterId");
            }
            if ($.engine == null) {
                throw new MissingRequiredPropertyException("GetDatabaseInstancePlainArgs", "engine");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetDatabaseInstancePlainArgs", "name");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetDatabaseInstancePlainArgs", "serviceName");
            }
            return $;
        }
    }

}