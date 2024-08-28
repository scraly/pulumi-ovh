// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject;

import com.ovh.ovh.CloudProject.ContainerRegistryIPRestrictionsRegistryArgs;
import com.ovh.ovh.CloudProject.inputs.ContainerRegistryIPRestrictionsRegistryState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import javax.annotation.Nullable;

/**
 * Apply IP restrictions container registry associated with a public cloud project on artifact manager component.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProject.CloudProjectFunctions;
 * import com.pulumi.ovh.CloudProject.inputs.GetContainerRegistryArgs;
 * import com.pulumi.ovh.CloudProject.ContainerRegistryIPRestrictionsRegistry;
 * import com.pulumi.ovh.CloudProject.ContainerRegistryIPRestrictionsRegistryArgs;
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
 *         final var registry = CloudProjectFunctions.getContainerRegistry(GetContainerRegistryArgs.builder()
 *             .serviceName("XXXXXX")
 *             .registryId("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx")
 *             .build());
 * 
 *         var my_registry_iprestrictions = new ContainerRegistryIPRestrictionsRegistry("my-registry-iprestrictions", ContainerRegistryIPRestrictionsRegistryArgs.builder()
 *             .serviceName(ovh_cloud_project_containerregistry.registry().service_name())
 *             .registryId(ovh_cloud_project_containerregistry.registry().id())
 *             .ipRestrictions(Map.ofEntries(
 *                 Map.entry("ip_block", "xxx.xxx.xxx.xxx/xx"),
 *                 Map.entry("description", "xxxxxxx")
 *             ))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:CloudProject/containerRegistryIPRestrictionsRegistry:ContainerRegistryIPRestrictionsRegistry")
public class ContainerRegistryIPRestrictionsRegistry extends com.pulumi.resources.CustomResource {
    /**
     * IP restrictions applied on artifact manager component.
     * 
     */
    @Export(name="ipRestrictions", refs={List.class,Map.class,String.class,Object.class}, tree="[0,[1,2,3]]")
    private Output<List<Map<String,Object>>> ipRestrictions;

    /**
     * @return IP restrictions applied on artifact manager component.
     * 
     */
    public Output<List<Map<String,Object>>> ipRestrictions() {
        return this.ipRestrictions;
    }
    /**
     * The id of the Managed Private Registry.
     * 
     */
    @Export(name="registryId", refs={String.class}, tree="[0]")
    private Output<String> registryId;

    /**
     * @return The id of the Managed Private Registry.
     * 
     */
    public Output<String> registryId() {
        return this.registryId;
    }
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContainerRegistryIPRestrictionsRegistry(String name) {
        this(name, ContainerRegistryIPRestrictionsRegistryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContainerRegistryIPRestrictionsRegistry(String name, ContainerRegistryIPRestrictionsRegistryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContainerRegistryIPRestrictionsRegistry(String name, ContainerRegistryIPRestrictionsRegistryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/containerRegistryIPRestrictionsRegistry:ContainerRegistryIPRestrictionsRegistry", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private ContainerRegistryIPRestrictionsRegistry(String name, Output<String> id, @Nullable ContainerRegistryIPRestrictionsRegistryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/containerRegistryIPRestrictionsRegistry:ContainerRegistryIPRestrictionsRegistry", name, state, makeResourceOptions(options, id));
    }

    private static ContainerRegistryIPRestrictionsRegistryArgs makeArgs(ContainerRegistryIPRestrictionsRegistryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ContainerRegistryIPRestrictionsRegistryArgs.Empty : args;
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
    public static ContainerRegistryIPRestrictionsRegistry get(String name, Output<String> id, @Nullable ContainerRegistryIPRestrictionsRegistryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContainerRegistryIPRestrictionsRegistry(name, id, state, options);
    }
}
