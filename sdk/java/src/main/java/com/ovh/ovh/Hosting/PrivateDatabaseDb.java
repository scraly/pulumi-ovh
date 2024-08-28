// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Hosting;

import com.ovh.ovh.Hosting.PrivateDatabaseDbArgs;
import com.ovh.ovh.Hosting.inputs.PrivateDatabaseDbState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Create a new database on your private cloud database service.
 * 
 * ## Example Usage
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Hosting.PrivateDatabaseDb;
 * import com.pulumi.ovh.Hosting.PrivateDatabaseDbArgs;
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
 *         var database = new PrivateDatabaseDb("database", PrivateDatabaseDbArgs.builder()        
 *             .databaseName("XXXXXX")
 *             .serviceName("XXXXXX")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * 
 * ## Import
 * 
 * OVHcloud Webhosting database can be imported using the `service_name` and the `database_name`, separated by &#34;/&#34; E.g.,
 * 
 * ```sh
 *  $ pulumi import ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb database service_name/database_name
 * ```
 * 
 */
@ResourceType(type="ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb")
public class PrivateDatabaseDb extends com.pulumi.resources.CustomResource {
    /**
     * Name of your new database
     * 
     */
    @Export(name="databaseName", refs={String.class}, tree="[0]")
    private Output<String> databaseName;

    /**
     * @return Name of your new database
     * 
     */
    public Output<String> databaseName() {
        return this.databaseName;
    }
    /**
     * The internal name of your private database.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The internal name of your private database.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrivateDatabaseDb(String name) {
        this(name, PrivateDatabaseDbArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrivateDatabaseDb(String name, PrivateDatabaseDbArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrivateDatabaseDb(String name, PrivateDatabaseDbArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb", name, args == null ? PrivateDatabaseDbArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PrivateDatabaseDb(String name, Output<String> id, @Nullable PrivateDatabaseDbState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Hosting/privateDatabaseDb:PrivateDatabaseDb", name, state, makeResourceOptions(options, id));
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
    public static PrivateDatabaseDb get(String name, Output<String> id, @Nullable PrivateDatabaseDbState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrivateDatabaseDb(name, id, state, options);
    }
}
