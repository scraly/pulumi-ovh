// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.ovhcloud.pulumi.ovh.CloudProject.RegionNetworkArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.RegionNetworkState;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.RegionNetworkSubnet;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Double;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Creates a network in a public cloud project.
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
 * import com.pulumi.ovh.CloudProject.RegionNetwork;
 * import com.pulumi.ovh.CloudProject.RegionNetworkArgs;
 * import com.pulumi.ovh.CloudProject.inputs.RegionNetworkSubnetArgs;
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
 *         var net = new RegionNetwork("net", RegionNetworkArgs.builder()
 *             .regionName("EU-SOUTH-LZ-MAD-A")
 *             .serviceName("XXXXXX")
 *             .subnet(RegionNetworkSubnetArgs.builder()
 *                 .cidr("10.0.0.0/24")
 *                 .enable_dhcp(true)
 *                 .enable_gateway_ip(false)
 *                 .ip_version(4)
 *                 .build())
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
 * A network in a public cloud project can be imported using the `service_name`, `region_name` and `id` attributes.
 * 
 * Using the following configuration:
 * 
 * hcl
 * 
 * import {
 * 
 *   id = &#34;&lt;service_name&gt;/&lt;region_name&gt;/&lt;id&gt;&#34;
 * 
 *   to = ovh_cloud_project_region_network.test
 * 
 * }
 * 
 * You can then run:
 * 
 * bash
 * 
 * $ pulumi preview -generate-config-out=network.tf
 * 
 * $ pulumi up
 * 
 * The file `network.tf` will then contain the imported resource&#39;s configuration, that can be copied next to the `import` block above.
 * 
 * See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
 * 
 */
@ResourceType(type="ovh:CloudProject/regionNetwork:RegionNetwork")
public class RegionNetwork extends com.pulumi.resources.CustomResource {
    /**
     * Name of the network
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the network
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Network region returned by the API
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Network region returned by the API
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Network region
     * 
     */
    @Export(name="regionName", refs={String.class}, tree="[0]")
    private Output<String> regionName;

    /**
     * @return Network region
     * 
     */
    public Output<String> regionName() {
        return this.regionName;
    }
    /**
     * The id of the public cloud project
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Parameters to create a subnet
     * 
     */
    @Export(name="subnet", refs={RegionNetworkSubnet.class}, tree="[0]")
    private Output<RegionNetworkSubnet> subnet;

    /**
     * @return Parameters to create a subnet
     * 
     */
    public Output<RegionNetworkSubnet> subnet() {
        return this.subnet;
    }
    /**
     * Network visibility
     * 
     */
    @Export(name="visibility", refs={String.class}, tree="[0]")
    private Output<String> visibility;

    /**
     * @return Network visibility
     * 
     */
    public Output<String> visibility() {
        return this.visibility;
    }
    /**
     * VLAN ID, between 1 and 4000
     * 
     */
    @Export(name="vlanId", refs={Double.class}, tree="[0]")
    private Output<Double> vlanId;

    /**
     * @return VLAN ID, between 1 and 4000
     * 
     */
    public Output<Double> vlanId() {
        return this.vlanId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RegionNetwork(java.lang.String name) {
        this(name, RegionNetworkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegionNetwork(java.lang.String name, RegionNetworkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegionNetwork(java.lang.String name, RegionNetworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/regionNetwork:RegionNetwork", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RegionNetwork(java.lang.String name, Output<java.lang.String> id, @Nullable RegionNetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/regionNetwork:RegionNetwork", name, state, makeResourceOptions(options, id), false);
    }

    private static RegionNetworkArgs makeArgs(RegionNetworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RegionNetworkArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static RegionNetwork get(java.lang.String name, Output<java.lang.String> id, @Nullable RegionNetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegionNetwork(name, id, state, options);
    }
}
