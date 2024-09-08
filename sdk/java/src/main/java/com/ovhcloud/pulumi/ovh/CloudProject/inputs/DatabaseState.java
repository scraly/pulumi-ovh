// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.ovhcloud.pulumi.ovh.CloudProject.inputs.DatabaseEndpointArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.DatabaseIpRestrictionArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.DatabaseNodeArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DatabaseState extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseState Empty = new DatabaseState();

    /**
     * Advanced configuration key / value.
     * 
     */
    @Import(name="advancedConfiguration")
    private @Nullable Output<Map<String,String>> advancedConfiguration;

    /**
     * @return Advanced configuration key / value.
     * 
     */
    public Optional<Output<Map<String,String>>> advancedConfiguration() {
        return Optional.ofNullable(this.advancedConfiguration);
    }

    /**
     * List of region where backups are pushed. Not more than 1 regions for MongoDB. Not more than 2 regions for the other engines with one being the same as the nodes[].region field
     * 
     */
    @Import(name="backupRegions")
    private @Nullable Output<List<String>> backupRegions;

    /**
     * @return List of region where backups are pushed. Not more than 1 regions for MongoDB. Not more than 2 regions for the other engines with one being the same as the nodes[].region field
     * 
     */
    public Optional<Output<List<String>>> backupRegions() {
        return Optional.ofNullable(this.backupRegions);
    }

    /**
     * Time on which backups start every day.
     * 
     */
    @Import(name="backupTime")
    private @Nullable Output<String> backupTime;

    /**
     * @return Time on which backups start every day.
     * 
     */
    public Optional<Output<String>> backupTime() {
        return Optional.ofNullable(this.backupTime);
    }

    /**
     * Date of the creation of the cluster.
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Date of the creation of the cluster.
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * Small description of the database service.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Small description of the database service.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * The disk size (in GB) of the database service.
     * 
     */
    @Import(name="diskSize")
    private @Nullable Output<Integer> diskSize;

    /**
     * @return The disk size (in GB) of the database service.
     * 
     */
    public Optional<Output<Integer>> diskSize() {
        return Optional.ofNullable(this.diskSize);
    }

    /**
     * Defines the disk type of the database service.
     * 
     */
    @Import(name="diskType")
    private @Nullable Output<String> diskType;

    /**
     * @return Defines the disk type of the database service.
     * 
     */
    public Optional<Output<String>> diskType() {
        return Optional.ofNullable(this.diskType);
    }

    /**
     * List of all endpoints objects of the service.
     * 
     */
    @Import(name="endpoints")
    private @Nullable Output<List<DatabaseEndpointArgs>> endpoints;

    /**
     * @return List of all endpoints objects of the service.
     * 
     */
    public Optional<Output<List<DatabaseEndpointArgs>>> endpoints() {
        return Optional.ofNullable(this.endpoints);
    }

    /**
     * The database engine you want to deploy. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    @Import(name="engine")
    private @Nullable Output<String> engine;

    /**
     * @return The database engine you want to deploy. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    public Optional<Output<String>> engine() {
        return Optional.ofNullable(this.engine);
    }

    /**
     * A valid OVHcloud public cloud database flavor name in which the nodes will be started.
     * Ex: &#34;db1-7&#34;. Changing this value upgrade the nodes with the new flavor.
     * You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
     * 
     */
    @Import(name="flavor")
    private @Nullable Output<String> flavor;

    /**
     * @return A valid OVHcloud public cloud database flavor name in which the nodes will be started.
     * Ex: &#34;db1-7&#34;. Changing this value upgrade the nodes with the new flavor.
     * You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
     * 
     */
    public Optional<Output<String>> flavor() {
        return Optional.ofNullable(this.flavor);
    }

    /**
     * IP Blocks authorized to access to the cluster.
     * 
     */
    @Import(name="ipRestrictions")
    private @Nullable Output<List<DatabaseIpRestrictionArgs>> ipRestrictions;

    /**
     * @return IP Blocks authorized to access to the cluster.
     * 
     */
    public Optional<Output<List<DatabaseIpRestrictionArgs>>> ipRestrictions() {
        return Optional.ofNullable(this.ipRestrictions);
    }

    /**
     * Defines whether the REST API is enabled on a kafka cluster
     * 
     */
    @Import(name="kafkaRestApi")
    private @Nullable Output<Boolean> kafkaRestApi;

    /**
     * @return Defines whether the REST API is enabled on a kafka cluster
     * 
     */
    public Optional<Output<Boolean>> kafkaRestApi() {
        return Optional.ofNullable(this.kafkaRestApi);
    }

    /**
     * Defines whether the schema registry is enabled on a Kafka cluster
     * 
     */
    @Import(name="kafkaSchemaRegistry")
    private @Nullable Output<Boolean> kafkaSchemaRegistry;

    /**
     * @return Defines whether the schema registry is enabled on a Kafka cluster
     * 
     */
    public Optional<Output<Boolean>> kafkaSchemaRegistry() {
        return Optional.ofNullable(this.kafkaSchemaRegistry);
    }

    /**
     * Time on which maintenances can start every day.
     * 
     */
    @Import(name="maintenanceTime")
    private @Nullable Output<String> maintenanceTime;

    /**
     * @return Time on which maintenances can start every day.
     * 
     */
    public Optional<Output<String>> maintenanceTime() {
        return Optional.ofNullable(this.maintenanceTime);
    }

    /**
     * Type of network of the cluster.
     * 
     */
    @Import(name="networkType")
    private @Nullable Output<String> networkType;

    /**
     * @return Type of network of the cluster.
     * 
     */
    public Optional<Output<String>> networkType() {
        return Optional.ofNullable(this.networkType);
    }

    /**
     * List of nodes object.
     * Multi region cluster are not yet available, all node should be identical.
     * 
     */
    @Import(name="nodes")
    private @Nullable Output<List<DatabaseNodeArgs>> nodes;

    /**
     * @return List of nodes object.
     * Multi region cluster are not yet available, all node should be identical.
     * 
     */
    public Optional<Output<List<DatabaseNodeArgs>>> nodes() {
        return Optional.ofNullable(this.nodes);
    }

    /**
     * Defines whether the ACLs are enabled on an OpenSearch cluster
     * 
     */
    @Import(name="opensearchAclsEnabled")
    private @Nullable Output<Boolean> opensearchAclsEnabled;

    /**
     * @return Defines whether the ACLs are enabled on an OpenSearch cluster
     * 
     */
    public Optional<Output<Boolean>> opensearchAclsEnabled() {
        return Optional.ofNullable(this.opensearchAclsEnabled);
    }

    /**
     * Plan of the cluster.
     * * MongoDB: Enum: &#34;discovery&#34;, &#34;production&#34;, &#34;advanced&#34;.
     * * Mysql, PosgreSQL, Cassandra, M3DB, : Enum: &#34;essential&#34;, &#34;business&#34;, &#34;enterprise&#34;.
     * * M3 Aggregator: &#34;business&#34;, &#34;enterprise&#34;.
     * * Redis: &#34;essential&#34;, &#34;business&#34;
     * 
     */
    @Import(name="plan")
    private @Nullable Output<String> plan;

    /**
     * @return Plan of the cluster.
     * * MongoDB: Enum: &#34;discovery&#34;, &#34;production&#34;, &#34;advanced&#34;.
     * * Mysql, PosgreSQL, Cassandra, M3DB, : Enum: &#34;essential&#34;, &#34;business&#34;, &#34;enterprise&#34;.
     * * M3 Aggregator: &#34;business&#34;, &#34;enterprise&#34;.
     * * Redis: &#34;essential&#34;, &#34;business&#34;
     * 
     */
    public Optional<Output<String>> plan() {
        return Optional.ofNullable(this.plan);
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

    /**
     * Current status of the cluster.
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Current status of the cluster.
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * The version of the engine in which the service should be deployed
     * 
     */
    @Import(name="version")
    private @Nullable Output<String> version;

    /**
     * @return The version of the engine in which the service should be deployed
     * 
     */
    public Optional<Output<String>> version() {
        return Optional.ofNullable(this.version);
    }

    private DatabaseState() {}

    private DatabaseState(DatabaseState $) {
        this.advancedConfiguration = $.advancedConfiguration;
        this.backupRegions = $.backupRegions;
        this.backupTime = $.backupTime;
        this.createdAt = $.createdAt;
        this.description = $.description;
        this.diskSize = $.diskSize;
        this.diskType = $.diskType;
        this.endpoints = $.endpoints;
        this.engine = $.engine;
        this.flavor = $.flavor;
        this.ipRestrictions = $.ipRestrictions;
        this.kafkaRestApi = $.kafkaRestApi;
        this.kafkaSchemaRegistry = $.kafkaSchemaRegistry;
        this.maintenanceTime = $.maintenanceTime;
        this.networkType = $.networkType;
        this.nodes = $.nodes;
        this.opensearchAclsEnabled = $.opensearchAclsEnabled;
        this.plan = $.plan;
        this.serviceName = $.serviceName;
        this.status = $.status;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseState $;

        public Builder() {
            $ = new DatabaseState();
        }

        public Builder(DatabaseState defaults) {
            $ = new DatabaseState(Objects.requireNonNull(defaults));
        }

        /**
         * @param advancedConfiguration Advanced configuration key / value.
         * 
         * @return builder
         * 
         */
        public Builder advancedConfiguration(@Nullable Output<Map<String,String>> advancedConfiguration) {
            $.advancedConfiguration = advancedConfiguration;
            return this;
        }

        /**
         * @param advancedConfiguration Advanced configuration key / value.
         * 
         * @return builder
         * 
         */
        public Builder advancedConfiguration(Map<String,String> advancedConfiguration) {
            return advancedConfiguration(Output.of(advancedConfiguration));
        }

        /**
         * @param backupRegions List of region where backups are pushed. Not more than 1 regions for MongoDB. Not more than 2 regions for the other engines with one being the same as the nodes[].region field
         * 
         * @return builder
         * 
         */
        public Builder backupRegions(@Nullable Output<List<String>> backupRegions) {
            $.backupRegions = backupRegions;
            return this;
        }

        /**
         * @param backupRegions List of region where backups are pushed. Not more than 1 regions for MongoDB. Not more than 2 regions for the other engines with one being the same as the nodes[].region field
         * 
         * @return builder
         * 
         */
        public Builder backupRegions(List<String> backupRegions) {
            return backupRegions(Output.of(backupRegions));
        }

        /**
         * @param backupRegions List of region where backups are pushed. Not more than 1 regions for MongoDB. Not more than 2 regions for the other engines with one being the same as the nodes[].region field
         * 
         * @return builder
         * 
         */
        public Builder backupRegions(String... backupRegions) {
            return backupRegions(List.of(backupRegions));
        }

        /**
         * @param backupTime Time on which backups start every day.
         * 
         * @return builder
         * 
         */
        public Builder backupTime(@Nullable Output<String> backupTime) {
            $.backupTime = backupTime;
            return this;
        }

        /**
         * @param backupTime Time on which backups start every day.
         * 
         * @return builder
         * 
         */
        public Builder backupTime(String backupTime) {
            return backupTime(Output.of(backupTime));
        }

        /**
         * @param createdAt Date of the creation of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Date of the creation of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param description Small description of the database service.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Small description of the database service.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param diskSize The disk size (in GB) of the database service.
         * 
         * @return builder
         * 
         */
        public Builder diskSize(@Nullable Output<Integer> diskSize) {
            $.diskSize = diskSize;
            return this;
        }

        /**
         * @param diskSize The disk size (in GB) of the database service.
         * 
         * @return builder
         * 
         */
        public Builder diskSize(Integer diskSize) {
            return diskSize(Output.of(diskSize));
        }

        /**
         * @param diskType Defines the disk type of the database service.
         * 
         * @return builder
         * 
         */
        public Builder diskType(@Nullable Output<String> diskType) {
            $.diskType = diskType;
            return this;
        }

        /**
         * @param diskType Defines the disk type of the database service.
         * 
         * @return builder
         * 
         */
        public Builder diskType(String diskType) {
            return diskType(Output.of(diskType));
        }

        /**
         * @param endpoints List of all endpoints objects of the service.
         * 
         * @return builder
         * 
         */
        public Builder endpoints(@Nullable Output<List<DatabaseEndpointArgs>> endpoints) {
            $.endpoints = endpoints;
            return this;
        }

        /**
         * @param endpoints List of all endpoints objects of the service.
         * 
         * @return builder
         * 
         */
        public Builder endpoints(List<DatabaseEndpointArgs> endpoints) {
            return endpoints(Output.of(endpoints));
        }

        /**
         * @param endpoints List of all endpoints objects of the service.
         * 
         * @return builder
         * 
         */
        public Builder endpoints(DatabaseEndpointArgs... endpoints) {
            return endpoints(List.of(endpoints));
        }

        /**
         * @param engine The database engine you want to deploy. To get a full list of available engine visit.
         * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * 
         * @return builder
         * 
         */
        public Builder engine(@Nullable Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine The database engine you want to deploy. To get a full list of available engine visit.
         * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param flavor A valid OVHcloud public cloud database flavor name in which the nodes will be started.
         * Ex: &#34;db1-7&#34;. Changing this value upgrade the nodes with the new flavor.
         * You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
         * 
         * @return builder
         * 
         */
        public Builder flavor(@Nullable Output<String> flavor) {
            $.flavor = flavor;
            return this;
        }

        /**
         * @param flavor A valid OVHcloud public cloud database flavor name in which the nodes will be started.
         * Ex: &#34;db1-7&#34;. Changing this value upgrade the nodes with the new flavor.
         * You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
         * 
         * @return builder
         * 
         */
        public Builder flavor(String flavor) {
            return flavor(Output.of(flavor));
        }

        /**
         * @param ipRestrictions IP Blocks authorized to access to the cluster.
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictions(@Nullable Output<List<DatabaseIpRestrictionArgs>> ipRestrictions) {
            $.ipRestrictions = ipRestrictions;
            return this;
        }

        /**
         * @param ipRestrictions IP Blocks authorized to access to the cluster.
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictions(List<DatabaseIpRestrictionArgs> ipRestrictions) {
            return ipRestrictions(Output.of(ipRestrictions));
        }

        /**
         * @param ipRestrictions IP Blocks authorized to access to the cluster.
         * 
         * @return builder
         * 
         */
        public Builder ipRestrictions(DatabaseIpRestrictionArgs... ipRestrictions) {
            return ipRestrictions(List.of(ipRestrictions));
        }

        /**
         * @param kafkaRestApi Defines whether the REST API is enabled on a kafka cluster
         * 
         * @return builder
         * 
         */
        public Builder kafkaRestApi(@Nullable Output<Boolean> kafkaRestApi) {
            $.kafkaRestApi = kafkaRestApi;
            return this;
        }

        /**
         * @param kafkaRestApi Defines whether the REST API is enabled on a kafka cluster
         * 
         * @return builder
         * 
         */
        public Builder kafkaRestApi(Boolean kafkaRestApi) {
            return kafkaRestApi(Output.of(kafkaRestApi));
        }

        /**
         * @param kafkaSchemaRegistry Defines whether the schema registry is enabled on a Kafka cluster
         * 
         * @return builder
         * 
         */
        public Builder kafkaSchemaRegistry(@Nullable Output<Boolean> kafkaSchemaRegistry) {
            $.kafkaSchemaRegistry = kafkaSchemaRegistry;
            return this;
        }

        /**
         * @param kafkaSchemaRegistry Defines whether the schema registry is enabled on a Kafka cluster
         * 
         * @return builder
         * 
         */
        public Builder kafkaSchemaRegistry(Boolean kafkaSchemaRegistry) {
            return kafkaSchemaRegistry(Output.of(kafkaSchemaRegistry));
        }

        /**
         * @param maintenanceTime Time on which maintenances can start every day.
         * 
         * @return builder
         * 
         */
        public Builder maintenanceTime(@Nullable Output<String> maintenanceTime) {
            $.maintenanceTime = maintenanceTime;
            return this;
        }

        /**
         * @param maintenanceTime Time on which maintenances can start every day.
         * 
         * @return builder
         * 
         */
        public Builder maintenanceTime(String maintenanceTime) {
            return maintenanceTime(Output.of(maintenanceTime));
        }

        /**
         * @param networkType Type of network of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder networkType(@Nullable Output<String> networkType) {
            $.networkType = networkType;
            return this;
        }

        /**
         * @param networkType Type of network of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder networkType(String networkType) {
            return networkType(Output.of(networkType));
        }

        /**
         * @param nodes List of nodes object.
         * Multi region cluster are not yet available, all node should be identical.
         * 
         * @return builder
         * 
         */
        public Builder nodes(@Nullable Output<List<DatabaseNodeArgs>> nodes) {
            $.nodes = nodes;
            return this;
        }

        /**
         * @param nodes List of nodes object.
         * Multi region cluster are not yet available, all node should be identical.
         * 
         * @return builder
         * 
         */
        public Builder nodes(List<DatabaseNodeArgs> nodes) {
            return nodes(Output.of(nodes));
        }

        /**
         * @param nodes List of nodes object.
         * Multi region cluster are not yet available, all node should be identical.
         * 
         * @return builder
         * 
         */
        public Builder nodes(DatabaseNodeArgs... nodes) {
            return nodes(List.of(nodes));
        }

        /**
         * @param opensearchAclsEnabled Defines whether the ACLs are enabled on an OpenSearch cluster
         * 
         * @return builder
         * 
         */
        public Builder opensearchAclsEnabled(@Nullable Output<Boolean> opensearchAclsEnabled) {
            $.opensearchAclsEnabled = opensearchAclsEnabled;
            return this;
        }

        /**
         * @param opensearchAclsEnabled Defines whether the ACLs are enabled on an OpenSearch cluster
         * 
         * @return builder
         * 
         */
        public Builder opensearchAclsEnabled(Boolean opensearchAclsEnabled) {
            return opensearchAclsEnabled(Output.of(opensearchAclsEnabled));
        }

        /**
         * @param plan Plan of the cluster.
         * * MongoDB: Enum: &#34;discovery&#34;, &#34;production&#34;, &#34;advanced&#34;.
         * * Mysql, PosgreSQL, Cassandra, M3DB, : Enum: &#34;essential&#34;, &#34;business&#34;, &#34;enterprise&#34;.
         * * M3 Aggregator: &#34;business&#34;, &#34;enterprise&#34;.
         * * Redis: &#34;essential&#34;, &#34;business&#34;
         * 
         * @return builder
         * 
         */
        public Builder plan(@Nullable Output<String> plan) {
            $.plan = plan;
            return this;
        }

        /**
         * @param plan Plan of the cluster.
         * * MongoDB: Enum: &#34;discovery&#34;, &#34;production&#34;, &#34;advanced&#34;.
         * * Mysql, PosgreSQL, Cassandra, M3DB, : Enum: &#34;essential&#34;, &#34;business&#34;, &#34;enterprise&#34;.
         * * M3 Aggregator: &#34;business&#34;, &#34;enterprise&#34;.
         * * Redis: &#34;essential&#34;, &#34;business&#34;
         * 
         * @return builder
         * 
         */
        public Builder plan(String plan) {
            return plan(Output.of(plan));
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

        /**
         * @param status Current status of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Current status of the cluster.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param version The version of the engine in which the service should be deployed
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Output<String> version) {
            $.version = version;
            return this;
        }

        /**
         * @param version The version of the engine in which the service should be deployed
         * 
         * @return builder
         * 
         */
        public Builder version(String version) {
            return version(Output.of(version));
        }

        public DatabaseState build() {
            return $;
        }
    }

}