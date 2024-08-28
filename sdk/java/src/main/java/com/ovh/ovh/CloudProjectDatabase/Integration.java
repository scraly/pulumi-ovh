// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProjectDatabase;

import com.ovh.ovh.CloudProjectDatabase.IntegrationArgs;
import com.ovh.ovh.CloudProjectDatabase.inputs.IntegrationState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates an integration for a database cluster associated with a public cloud project.
 * 
 * With this resource you can create an integration for all engine exept `mongodb`.
 * 
 * Please take a look at the list of available `types` in the `Argument references` section in order to know the list of avaulable integrations. For example, thanks to the integration feature you can have your PostgreSQL logs in your OpenSearch Database.
 * 
 * ## Example Usage
 * 
 * Push PostgreSQL logs in an OpenSearch DB:
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProjectDatabase.CloudProjectDatabaseFunctions;
 * import com.pulumi.ovh.CloudProjectDatabase.inputs.GetDatabaseArgs;
 * import com.pulumi.ovh.CloudProjectDatabase.Integration;
 * import com.pulumi.ovh.CloudProjectDatabase.IntegrationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var dbpostgresql = CloudProjectDatabaseFunctions.getDatabase(GetDatabaseArgs.builder()
 *             .serviceName("XXXX")
 *             .engine("postgresql")
 *             .id("ZZZZ")
 *             .build());
 * 
 *         final var dbopensearch = CloudProjectDatabaseFunctions.getDatabase(GetDatabaseArgs.builder()
 *             .serviceName("XXXX")
 *             .engine("opensearch")
 *             .id("ZZZZ")
 *             .build());
 * 
 *         var integration = new Integration("integration", IntegrationArgs.builder()
 *             .serviceName(dbpostgresql.applyValue(getDatabaseResult -> getDatabaseResult.serviceName()))
 *             .engine(dbpostgresql.applyValue(getDatabaseResult -> getDatabaseResult.engine()))
 *             .clusterId(dbpostgresql.applyValue(getDatabaseResult -> getDatabaseResult.id()))
 *             .sourceServiceId(dbpostgresql.applyValue(getDatabaseResult -> getDatabaseResult.id()))
 *             .destinationServiceId(dbopensearch.applyValue(getDatabaseResult -> getDatabaseResult.id()))
 *             .type("opensearchLogs")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * OVHcloud Managed database clusters users can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the user, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:CloudProjectDatabase/integration:Integration my_user service_name/engine/cluster_id/id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProjectDatabase/integration:Integration")
public class Integration extends com.pulumi.resources.CustomResource {
    /**
     * Cluster ID.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * ID of the destination service.
     * 
     */
    @Export(name="destinationServiceId", refs={String.class}, tree="[0]")
    private Output<String> destinationServiceId;

    /**
     * @return ID of the destination service.
     * 
     */
    public Output<String> destinationServiceId() {
        return this.destinationServiceId;
    }
    /**
     * The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * All engines available exept `mongodb`.
     * 
     */
    @Export(name="engine", refs={String.class}, tree="[0]")
    private Output<String> engine;

    /**
     * @return The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * All engines available exept `mongodb`.
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }
    /**
     * Parameters for the integration.
     * 
     */
    @Export(name="parameters", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> parameters;

    /**
     * @return Parameters for the integration.
     * 
     */
    public Output<Optional<Map<String,String>>> parameters() {
        return Codegen.optional(this.parameters);
    }
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * ID of the source service.
     * 
     */
    @Export(name="sourceServiceId", refs={String.class}, tree="[0]")
    private Output<String> sourceServiceId;

    /**
     * @return ID of the source service.
     * 
     */
    public Output<String> sourceServiceId() {
        return this.sourceServiceId;
    }
    /**
     * Current status of the integration.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Current status of the integration.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Type of the integration.
     * Available types:
     * * `grafanaDashboard`
     * * `grafanaDatasource`
     * * `kafkaConnect`
     * * `kafkaLogs`
     * * `kafkaMirrorMaker`
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of the integration.
     * Available types:
     * * `grafanaDashboard`
     * * `grafanaDatasource`
     * * `kafkaConnect`
     * * `kafkaLogs`
     * * `kafkaMirrorMaker`
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Integration(String name) {
        this(name, IntegrationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Integration(String name, IntegrationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Integration(String name, IntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/integration:Integration", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private Integration(String name, Output<String> id, @Nullable IntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/integration:Integration", name, state, makeResourceOptions(options, id));
    }

    private static IntegrationArgs makeArgs(IntegrationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? IntegrationArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Integration get(String name, Output<String> id, @Nullable IntegrationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Integration(name, id, state, options);
    }
}
