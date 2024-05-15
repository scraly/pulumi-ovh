// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage rules for TCP route.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const reject = new ovh.iploadbalancing.TcpRoute("reject", {
 *     serviceName: "loadbalancer-xxxxxxxxxxxxxxxxxx",
 *     weight: 1,
 *     frontendId: 11111,
 *     action: {
 *         type: "reject",
 *     },
 * });
 * const examplerule = new ovh.iploadbalancing.TcpRouteRule("examplerule", {
 *     serviceName: "loadbalancer-xxxxxxxxxxxxxxxxxx",
 *     routeId: reject.id,
 *     displayName: "Match example.com host",
 *     field: "sni",
 *     match: "is",
 *     negate: false,
 *     pattern: "example.com",
 * });
 * ```
 *
 * ## Import
 *
 * TCP route rule can be imported using the following format `serviceName`, the `id` of the route and the `id` of the rule separated by "/" e.g.
 */
export class TcpRouteRule extends pulumi.CustomResource {
    /**
     * Get an existing TcpRouteRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TcpRouteRuleState, opts?: pulumi.CustomResourceOptions): TcpRouteRule {
        return new TcpRouteRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:IpLoadBalancing/tcpRouteRule:TcpRouteRule';

    /**
     * Returns true if the given object is an instance of TcpRouteRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TcpRouteRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TcpRouteRule.__pulumiType;
    }

    /**
     * Human readable name for your rule, this field is for you
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
     */
    public readonly field!: pulumi.Output<string>;
    /**
     * Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
     */
    public readonly match!: pulumi.Output<string>;
    /**
     * Invert the matching operator effect
     */
    public readonly negate!: pulumi.Output<boolean>;
    /**
     * Value to match against this match. Interpretation if this field depends on the match and field
     */
    public readonly pattern!: pulumi.Output<string | undefined>;
    /**
     * The route to apply this rule
     */
    public readonly routeId!: pulumi.Output<string>;
    /**
     * The internal name of your IP load balancing
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Name of sub-field, if applicable. This may be a Cookie or Header name for instance
     */
    public readonly subField!: pulumi.Output<string | undefined>;

    /**
     * Create a TcpRouteRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TcpRouteRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TcpRouteRuleArgs | TcpRouteRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TcpRouteRuleState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["field"] = state ? state.field : undefined;
            resourceInputs["match"] = state ? state.match : undefined;
            resourceInputs["negate"] = state ? state.negate : undefined;
            resourceInputs["pattern"] = state ? state.pattern : undefined;
            resourceInputs["routeId"] = state ? state.routeId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["subField"] = state ? state.subField : undefined;
        } else {
            const args = argsOrState as TcpRouteRuleArgs | undefined;
            if ((!args || args.field === undefined) && !opts.urn) {
                throw new Error("Missing required property 'field'");
            }
            if ((!args || args.match === undefined) && !opts.urn) {
                throw new Error("Missing required property 'match'");
            }
            if ((!args || args.routeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeId'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["field"] = args ? args.field : undefined;
            resourceInputs["match"] = args ? args.match : undefined;
            resourceInputs["negate"] = args ? args.negate : undefined;
            resourceInputs["pattern"] = args ? args.pattern : undefined;
            resourceInputs["routeId"] = args ? args.routeId : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["subField"] = args ? args.subField : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TcpRouteRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TcpRouteRule resources.
 */
export interface TcpRouteRuleState {
    /**
     * Human readable name for your rule, this field is for you
     */
    displayName?: pulumi.Input<string>;
    /**
     * Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
     */
    field?: pulumi.Input<string>;
    /**
     * Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
     */
    match?: pulumi.Input<string>;
    /**
     * Invert the matching operator effect
     */
    negate?: pulumi.Input<boolean>;
    /**
     * Value to match against this match. Interpretation if this field depends on the match and field
     */
    pattern?: pulumi.Input<string>;
    /**
     * The route to apply this rule
     */
    routeId?: pulumi.Input<string>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Name of sub-field, if applicable. This may be a Cookie or Header name for instance
     */
    subField?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TcpRouteRule resource.
 */
export interface TcpRouteRuleArgs {
    /**
     * Human readable name for your rule, this field is for you
     */
    displayName?: pulumi.Input<string>;
    /**
     * Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules
     */
    field: pulumi.Input<string>;
    /**
     * Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"
     */
    match: pulumi.Input<string>;
    /**
     * Invert the matching operator effect
     */
    negate?: pulumi.Input<boolean>;
    /**
     * Value to match against this match. Interpretation if this field depends on the match and field
     */
    pattern?: pulumi.Input<string>;
    /**
     * The route to apply this rule
     */
    routeId: pulumi.Input<string>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName: pulumi.Input<string>;
    /**
     * Name of sub-field, if applicable. This may be a Cookie or Header name for instance
     */
    subField?: pulumi.Input<string>;
}
