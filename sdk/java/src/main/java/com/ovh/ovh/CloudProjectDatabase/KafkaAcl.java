// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProjectDatabase;

import com.ovh.ovh.CloudProjectDatabase.KafkaAclArgs;
import com.ovh.ovh.CloudProjectDatabase.inputs.KafkaAclState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Creates an ACL for a kafka cluster associated with a public cloud project.
 * 
 * ## Example Usage
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProjectDatabase.CloudProjectDatabaseFunctions;
 * import com.pulumi.ovh.CloudProjectDatabase.inputs.GetDatabaseArgs;
 * import com.pulumi.ovh.CloudProjectDatabase.KafkaAcl;
 * import com.pulumi.ovh.CloudProjectDatabase.KafkaAclArgs;
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
 *         final var kafka = CloudProjectDatabaseFunctions.getDatabase(GetDatabaseArgs.builder()
 *             .serviceName("XXX")
 *             .engine("kafka")
 *             .id("ZZZ")
 *             .build());
 * 
 *         var acl = new KafkaAcl("acl", KafkaAclArgs.builder()        
 *             .serviceName(kafka.applyValue(getDatabaseResult -> getDatabaseResult.serviceName()))
 *             .clusterId(kafka.applyValue(getDatabaseResult -> getDatabaseResult.id()))
 *             .permission("read")
 *             .topic("mytopic")
 *             .username("johndoe")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * 
 * ## Import
 * 
 * OVHcloud Managed kafka clusters ACLs can be imported using the `service_name`, `cluster_id` and `id` of the acl, separated by &#34;/&#34; E.g., bash
 * 
 * ```sh
 *  $ pulumi import ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl my_acl service_name/cluster_id/id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl")
public class KafkaAcl extends com.pulumi.resources.CustomResource {
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
     * Permission to give to this username on this topic.
     * Available permissions:
     * 
     */
    @Export(name="permission", refs={String.class}, tree="[0]")
    private Output<String> permission;

    /**
     * @return Permission to give to this username on this topic.
     * Available permissions:
     * 
     */
    public Output<String> permission() {
        return this.permission;
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
     * Topic affected by this ACL.
     * 
     */
    @Export(name="topic", refs={String.class}, tree="[0]")
    private Output<String> topic;

    /**
     * @return Topic affected by this ACL.
     * 
     */
    public Output<String> topic() {
        return this.topic;
    }
    /**
     * Username affected by this ACL.
     * 
     */
    @Export(name="username", refs={String.class}, tree="[0]")
    private Output<String> username;

    /**
     * @return Username affected by this ACL.
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public KafkaAcl(String name) {
        this(name, KafkaAclArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public KafkaAcl(String name, KafkaAclArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public KafkaAcl(String name, KafkaAclArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl", name, args == null ? KafkaAclArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private KafkaAcl(String name, Output<String> id, @Nullable KafkaAclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/kafkaAcl:KafkaAcl", name, state, makeResourceOptions(options, id));
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
    public static KafkaAcl get(String name, Output<String> id, @Nullable KafkaAclState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new KafkaAcl(name, id, state, options);
    }
}