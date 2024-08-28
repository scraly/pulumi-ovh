// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProjectDatabase;

import com.ovh.ovh.CloudProjectDatabase.UserArgs;
import com.ovh.ovh.CloudProjectDatabase.inputs.UserState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * OVHcloud Managed database clusters users can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the user, separated by &#34;/&#34; E.g., bash
 * 
 * ```sh
 *  $ pulumi import ovh:CloudProjectDatabase/user:User my_user service_name/engine/cluster_id/id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProjectDatabase/user:User")
public class User extends com.pulumi.resources.CustomResource {
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
     * Date of the creation of the user.
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Date of the creation of the user.
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * Available engines:
     * 
     */
    @Export(name="engine", refs={String.class}, tree="[0]")
    private Output<String> engine;

    /**
     * @return The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * Available engines:
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }
    /**
     * Name of the user. A user named &#34;avnadmin&#34; is map with already created admin user and reset his password instead of create a new user. The &#34;Grafana&#34; engine only allows the &#34;avnadmin&#34; mapping.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the user. A user named &#34;avnadmin&#34; is map with already created admin user and reset his password instead of create a new user. The &#34;Grafana&#34; engine only allows the &#34;avnadmin&#34; mapping.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * (Sensitive) Password of the user.
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return (Sensitive) Password of the user.
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * Arbitrary string to change to trigger a password update
     * 
     */
    @Export(name="passwordReset", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> passwordReset;

    /**
     * @return Arbitrary string to change to trigger a password update
     * 
     */
    public Output<Optional<String>> passwordReset() {
        return Codegen.optional(this.passwordReset);
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
     * Current status of the user.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Current status of the user.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public User(String name) {
        this(name, UserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public User(String name, UserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public User(String name, UserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/user:User", name, args == null ? UserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private User(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/user:User", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static User get(String name, Output<String> id, @Nullable UserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new User(name, id, state, options);
    }
}
