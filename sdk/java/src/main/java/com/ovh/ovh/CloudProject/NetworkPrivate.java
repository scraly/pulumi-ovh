// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.CloudProject;

import com.ovh.ovh.CloudProject.NetworkPrivateArgs;
import com.ovh.ovh.CloudProject.inputs.NetworkPrivateState;
import com.ovh.ovh.CloudProject.outputs.NetworkPrivateRegionsAttribute;
import com.ovh.ovh.CloudProject.outputs.NetworkPrivateRegionsStatus;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a private network in a public cloud project.
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
 * import com.pulumi.ovh.CloudProject.NetworkPrivate;
 * import com.pulumi.ovh.CloudProject.NetworkPrivateArgs;
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
 *         var net = new NetworkPrivate("net", NetworkPrivateArgs.builder()
 *             .regions(            
 *                 "GRA1",
 *                 "BHS1")
 *             .serviceName("XXXXXX")
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
 * Private network in a public cloud project can be imported using the `service_name` and the `network_id`, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:CloudProject/networkPrivate:NetworkPrivate mynet ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProject/networkPrivate:NetworkPrivate")
public class NetworkPrivate extends com.pulumi.resources.CustomResource {
    /**
     * The name of the network.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the network.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * an array of valid OVHcloud public cloud region ID in which the network
     * will be available. Ex.: &#34;GRA1&#34;. Defaults to all public cloud regions.
     * 
     */
    @Export(name="regions", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> regions;

    /**
     * @return an array of valid OVHcloud public cloud region ID in which the network
     * will be available. Ex.: &#34;GRA1&#34;. Defaults to all public cloud regions.
     * 
     */
    public Output<List<String>> regions() {
        return this.regions;
    }
    /**
     * A map representing information about the region.
     * * `regions_attributes/region` - The id of the region.
     * * `regions_attributes/status` - The status of the network in the region.
     * * `regions_attributes/openstackid` - The private network id in the region.
     * 
     */
    @Export(name="regionsAttributes", refs={List.class,NetworkPrivateRegionsAttribute.class}, tree="[0,1]")
    private Output<List<NetworkPrivateRegionsAttribute>> regionsAttributes;

    /**
     * @return A map representing information about the region.
     * * `regions_attributes/region` - The id of the region.
     * * `regions_attributes/status` - The status of the network in the region.
     * * `regions_attributes/openstackid` - The private network id in the region.
     * 
     */
    public Output<List<NetworkPrivateRegionsAttribute>> regionsAttributes() {
        return this.regionsAttributes;
    }
    /**
     * (Deprecated) A map representing the status of the network per region.
     * * `regions_status/region` - (Deprecated) The id of the region.
     * * `regions_status/status` - (Deprecated) The status of the network in the region.
     * 
     * @deprecated
     * use the regions_attributes field instead
     * 
     */
    @Deprecated /* use the regions_attributes field instead */
    @Export(name="regionsStatuses", refs={List.class,NetworkPrivateRegionsStatus.class}, tree="[0,1]")
    private Output<List<NetworkPrivateRegionsStatus>> regionsStatuses;

    /**
     * @return (Deprecated) A map representing the status of the network per region.
     * * `regions_status/region` - (Deprecated) The id of the region.
     * * `regions_status/status` - (Deprecated) The status of the network in the region.
     * 
     */
    public Output<List<NetworkPrivateRegionsStatus>> regionsStatuses() {
        return this.regionsStatuses;
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
     * the status of the network. should be normally set to &#39;ACTIVE&#39;.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return the status of the network. should be normally set to &#39;ACTIVE&#39;.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * the type of the network. Either &#39;private&#39; or &#39;public&#39;.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return the type of the network. Either &#39;private&#39; or &#39;public&#39;.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * a vlan id to associate with the network.
     * Changing this value recreates the resource. Defaults to 0.
     * 
     */
    @Export(name="vlanId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> vlanId;

    /**
     * @return a vlan id to associate with the network.
     * Changing this value recreates the resource. Defaults to 0.
     * 
     */
    public Output<Optional<Integer>> vlanId() {
        return Codegen.optional(this.vlanId);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkPrivate(String name) {
        this(name, NetworkPrivateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkPrivate(String name, NetworkPrivateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkPrivate(String name, NetworkPrivateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/networkPrivate:NetworkPrivate", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()));
    }

    private NetworkPrivate(String name, Output<String> id, @Nullable NetworkPrivateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/networkPrivate:NetworkPrivate", name, state, makeResourceOptions(options, id));
    }

    private static NetworkPrivateArgs makeArgs(NetworkPrivateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NetworkPrivateArgs.Empty : args;
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
    public static NetworkPrivate get(String name, Output<String> id, @Nullable NetworkPrivateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkPrivate(name, id, state, options);
    }
}
