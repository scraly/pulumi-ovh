// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.IpLoadBalancing;

import com.ovh.ovh.IpLoadBalancing.RefreshArgs;
import com.ovh.ovh.IpLoadBalancing.inputs.RefreshState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Applies changes from other `ovh_iploadbalancing_*` resources to the production configuration of loadbalancers.
 * 
 * ## Example Usage
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.IpLoadBalancing.IpLoadBalancingFunctions;
 * import com.pulumi.ovh.IpLoadBalancing.inputs.GetIpLoadBalancingArgs;
 * import com.pulumi.ovh.IpLoadBalancing.TcpFarm;
 * import com.pulumi.ovh.IpLoadBalancing.TcpFarmArgs;
 * import com.pulumi.ovh.IpLoadBalancing.TcpFarmServer;
 * import com.pulumi.ovh.IpLoadBalancing.TcpFarmServerArgs;
 * import com.pulumi.ovh.IpLoadBalancing.Refresh;
 * import com.pulumi.ovh.IpLoadBalancing.RefreshArgs;
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
 *         final var lb = IpLoadBalancingFunctions.getIpLoadBalancing(GetIpLoadBalancingArgs.builder()
 *             .serviceName("ip-1.2.3.4")
 *             .state("ok")
 *             .build());
 * 
 *         var farmname = new TcpFarm("farmname", TcpFarmArgs.builder()        
 *             .port(8080)
 *             .serviceName(lb.applyValue(getIpLoadBalancingResult -> getIpLoadBalancingResult.serviceName()))
 *             .zone("all")
 *             .build());
 * 
 *         var backend = new TcpFarmServer("backend", TcpFarmServerArgs.builder()        
 *             .address("4.5.6.7")
 *             .backup(true)
 *             .displayName("mybackend")
 *             .farmId(farmname.id())
 *             .port(80)
 *             .probe(true)
 *             .proxyProtocolVersion("v2")
 *             .serviceName(lb.applyValue(getIpLoadBalancingResult -> getIpLoadBalancingResult.serviceName()))
 *             .ssl(false)
 *             .status("active")
 *             .weight(2)
 *             .build());
 * 
 *         var mylb = new Refresh("mylb", RefreshArgs.builder()        
 *             .keepers(backend.stream().map(element -> element.address()).collect(toList()))
 *             .serviceName(lb.applyValue(getIpLoadBalancingResult -> getIpLoadBalancingResult.serviceName()))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * 
 */
@ResourceType(type="ovh:IpLoadBalancing/refresh:Refresh")
public class Refresh extends com.pulumi.resources.CustomResource {
    /**
     * List of values tracked to trigger refresh, used also to form implicit dependencies
     * 
     */
    @Export(name="keepers", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> keepers;

    /**
     * @return List of values tracked to trigger refresh, used also to form implicit dependencies
     * 
     */
    public Output<List<String>> keepers() {
        return this.keepers;
    }
    /**
     * The internal name of your IP load balancing
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The internal name of your IP load balancing
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Refresh(String name) {
        this(name, RefreshArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Refresh(String name, RefreshArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Refresh(String name, RefreshArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/refresh:Refresh", name, args == null ? RefreshArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Refresh(String name, Output<String> id, @Nullable RefreshState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/refresh:Refresh", name, state, makeResourceOptions(options, id));
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
    public static Refresh get(String name, Output<String> id, @Nullable RefreshState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Refresh(name, id, state, options);
    }
}
