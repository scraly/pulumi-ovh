// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.IpLoadBalancing;

import com.ovhcloud.pulumi.ovh.IpLoadBalancing.HttpRouteRuleArgs;
import com.ovhcloud.pulumi.ovh.IpLoadBalancing.inputs.HttpRouteRuleState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manage rules for HTTP route.
 * 
 * ## Example Usage
 * 
 * Route which redirect all URL to HTTPs for example.com (Vhost).
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.IpLoadBalancing.HttpRoute;
 * import com.pulumi.ovh.IpLoadBalancing.HttpRouteArgs;
 * import com.pulumi.ovh.IpLoadBalancing.inputs.HttpRouteActionArgs;
 * import com.pulumi.ovh.IpLoadBalancing.HttpRouteRule;
 * import com.pulumi.ovh.IpLoadBalancing.HttpRouteRuleArgs;
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
 *         var httpsredirect = new HttpRoute("httpsredirect", HttpRouteArgs.builder()
 *             .action(HttpRouteActionArgs.builder()
 *                 .status(302)
 *                 .target("https://${host}${path}${arguments}")
 *                 .type("redirect")
 *                 .build())
 *             .displayName("Redirect to HTTPS")
 *             .frontendId(11111)
 *             .serviceName("loadbalancer-xxxxxxxxxxxxxxxxxx")
 *             .weight(1)
 *             .build());
 * 
 *         var examplerule = new HttpRouteRule("examplerule", HttpRouteRuleArgs.builder()
 *             .displayName("Match example.com host")
 *             .field("host")
 *             .match("is")
 *             .negate(false)
 *             .pattern("example.com")
 *             .routeId(httpsredirect.id())
 *             .serviceName("loadbalancer-xxxxxxxxxxxxxxxxxx")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Rule which match a specific header (same effect as the host match above).
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.IpLoadBalancing.HttpRouteRule;
 * import com.pulumi.ovh.IpLoadBalancing.HttpRouteRuleArgs;
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
 *         var examplerule = new HttpRouteRule("examplerule", HttpRouteRuleArgs.builder()
 *             .displayName("Match example.com Host header")
 *             .field("headers")
 *             .match("is")
 *             .negate(false)
 *             .pattern("example.com")
 *             .routeId(ovh_iploadbalancing_http_route.httpsredirect().id())
 *             .serviceName("loadbalancer-xxxxxxxxxxxxxxxxxx")
 *             .subField("Host")
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
 * HTTP route rule can be imported using the following format `service_name`, the `id` of the route and the `id` of the rule separated by &#34;/&#34; e.g.
 * 
 */
@ResourceType(type="ovh:IpLoadBalancing/httpRouteRule:HttpRouteRule")
public class HttpRouteRule extends com.pulumi.resources.CustomResource {
    /**
     * Human readable name for your rule, this field is for you
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Human readable name for your rule, this field is for you
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Name of the field to match like &#34;protocol&#34; or &#34;host&#34;. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34; for a list of available rules
     * 
     */
    @Export(name="field", refs={String.class}, tree="[0]")
    private Output<String> field;

    /**
     * @return Name of the field to match like &#34;protocol&#34; or &#34;host&#34;. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34; for a list of available rules
     * 
     */
    public Output<String> field() {
        return this.field;
    }
    /**
     * Matching operator. Not all operators are available for all fields. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34;
     * 
     */
    @Export(name="match", refs={String.class}, tree="[0]")
    private Output<String> match;

    /**
     * @return Matching operator. Not all operators are available for all fields. See &#34;/ipLoadbalancing/{serviceName}/availableRouteRules&#34;
     * 
     */
    public Output<String> match() {
        return this.match;
    }
    /**
     * Invert the matching operator effect
     * 
     */
    @Export(name="negate", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> negate;

    /**
     * @return Invert the matching operator effect
     * 
     */
    public Output<Boolean> negate() {
        return this.negate;
    }
    /**
     * Value to match against this match. Interpretation if this field depends on the match and field
     * 
     */
    @Export(name="pattern", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pattern;

    /**
     * @return Value to match against this match. Interpretation if this field depends on the match and field
     * 
     */
    public Output<Optional<String>> pattern() {
        return Codegen.optional(this.pattern);
    }
    /**
     * The route to apply this rule
     * 
     */
    @Export(name="routeId", refs={String.class}, tree="[0]")
    private Output<String> routeId;

    /**
     * @return The route to apply this rule
     * 
     */
    public Output<String> routeId() {
        return this.routeId;
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
     * Name of sub-field, if applicable. This may be a Cookie or Header name for instance
     * 
     */
    @Export(name="subField", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> subField;

    /**
     * @return Name of sub-field, if applicable. This may be a Cookie or Header name for instance
     * 
     */
    public Output<Optional<String>> subField() {
        return Codegen.optional(this.subField);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HttpRouteRule(java.lang.String name) {
        this(name, HttpRouteRuleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HttpRouteRule(java.lang.String name, HttpRouteRuleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HttpRouteRule(java.lang.String name, HttpRouteRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/httpRouteRule:HttpRouteRule", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private HttpRouteRule(java.lang.String name, Output<java.lang.String> id, @Nullable HttpRouteRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/httpRouteRule:HttpRouteRule", name, state, makeResourceOptions(options, id), false);
    }

    private static HttpRouteRuleArgs makeArgs(HttpRouteRuleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? HttpRouteRuleArgs.Empty : args;
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
    public static HttpRouteRule get(java.lang.String name, Output<java.lang.String> id, @Nullable HttpRouteRuleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HttpRouteRule(name, id, state, options);
    }
}