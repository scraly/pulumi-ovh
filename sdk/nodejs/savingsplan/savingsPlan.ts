// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Create and manage an OVHcloud Savings Plan
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const plan = new ovh.savingsplan.SavingsPlan("plan", {
 *     autoRenewal: true,
 *     displayName: "one_month_rancher_savings_plan",
 *     flavor: "Rancher",
 *     period: "P1M",
 *     serviceName: "<public cloud project ID>",
 *     size: 2,
 * });
 * ```
 *
 * ## Import
 *
 * A savings plan can be imported using the following format: `serviceName` and `id` of the savings plan, separated by "/" e.g.
 */
export class SavingsPlan extends pulumi.CustomResource {
    /**
     * Get an existing SavingsPlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SavingsPlanState, opts?: pulumi.CustomResourceOptions): SavingsPlan {
        return new SavingsPlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:SavingsPlan/savingsPlan:SavingsPlan';

    /**
     * Returns true if the given object is an instance of SavingsPlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SavingsPlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SavingsPlan.__pulumiType;
    }

    /**
     * Whether Savings Plan should be renewed at the end of the period (defaults to false)
     */
    public readonly autoRenewal!: pulumi.Output<boolean>;
    /**
     * Custom display name, used in invoices
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * End date of the Savings Plan
     */
    public /*out*/ readonly endDate!: pulumi.Output<string>;
    /**
     * Savings Plan flavor (e.g. Rancher, C3-4, any instance flavor, ...)
     */
    public readonly flavor!: pulumi.Output<string>;
    /**
     * Periodicity of the Savings Plan
     */
    public readonly period!: pulumi.Output<string>;
    /**
     * Action performed when reaching the end of the period (controles by the `autoRenewal` parameter)
     */
    public /*out*/ readonly periodEndAction!: pulumi.Output<string>;
    /**
     * End date of the current period
     */
    public /*out*/ readonly periodEndDate!: pulumi.Output<string>;
    /**
     * Start date of the current period
     */
    public /*out*/ readonly periodStartDate!: pulumi.Output<string>;
    /**
     * Billing ID of the service
     */
    public /*out*/ readonly serviceId!: pulumi.Output<number>;
    /**
     * ID of the public cloud project
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Size of the Savings Plan
     */
    public readonly size!: pulumi.Output<number>;
    /**
     * Start date of the Savings Plan
     */
    public /*out*/ readonly startDate!: pulumi.Output<string>;
    /**
     * Status of the Savings Plan
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a SavingsPlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SavingsPlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SavingsPlanArgs | SavingsPlanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SavingsPlanState | undefined;
            resourceInputs["autoRenewal"] = state ? state.autoRenewal : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["endDate"] = state ? state.endDate : undefined;
            resourceInputs["flavor"] = state ? state.flavor : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["periodEndAction"] = state ? state.periodEndAction : undefined;
            resourceInputs["periodEndDate"] = state ? state.periodEndDate : undefined;
            resourceInputs["periodStartDate"] = state ? state.periodStartDate : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["startDate"] = state ? state.startDate : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as SavingsPlanArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.flavor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'flavor'");
            }
            if ((!args || args.period === undefined) && !opts.urn) {
                throw new Error("Missing required property 'period'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            resourceInputs["autoRenewal"] = args ? args.autoRenewal : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["flavor"] = args ? args.flavor : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["endDate"] = undefined /*out*/;
            resourceInputs["periodEndAction"] = undefined /*out*/;
            resourceInputs["periodEndDate"] = undefined /*out*/;
            resourceInputs["periodStartDate"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["startDate"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SavingsPlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SavingsPlan resources.
 */
export interface SavingsPlanState {
    /**
     * Whether Savings Plan should be renewed at the end of the period (defaults to false)
     */
    autoRenewal?: pulumi.Input<boolean>;
    /**
     * Custom display name, used in invoices
     */
    displayName?: pulumi.Input<string>;
    /**
     * End date of the Savings Plan
     */
    endDate?: pulumi.Input<string>;
    /**
     * Savings Plan flavor (e.g. Rancher, C3-4, any instance flavor, ...)
     */
    flavor?: pulumi.Input<string>;
    /**
     * Periodicity of the Savings Plan
     */
    period?: pulumi.Input<string>;
    /**
     * Action performed when reaching the end of the period (controles by the `autoRenewal` parameter)
     */
    periodEndAction?: pulumi.Input<string>;
    /**
     * End date of the current period
     */
    periodEndDate?: pulumi.Input<string>;
    /**
     * Start date of the current period
     */
    periodStartDate?: pulumi.Input<string>;
    /**
     * Billing ID of the service
     */
    serviceId?: pulumi.Input<number>;
    /**
     * ID of the public cloud project
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Size of the Savings Plan
     */
    size?: pulumi.Input<number>;
    /**
     * Start date of the Savings Plan
     */
    startDate?: pulumi.Input<string>;
    /**
     * Status of the Savings Plan
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SavingsPlan resource.
 */
export interface SavingsPlanArgs {
    /**
     * Whether Savings Plan should be renewed at the end of the period (defaults to false)
     */
    autoRenewal?: pulumi.Input<boolean>;
    /**
     * Custom display name, used in invoices
     */
    displayName: pulumi.Input<string>;
    /**
     * Savings Plan flavor (e.g. Rancher, C3-4, any instance flavor, ...)
     */
    flavor: pulumi.Input<string>;
    /**
     * Periodicity of the Savings Plan
     */
    period: pulumi.Input<string>;
    /**
     * ID of the public cloud project
     */
    serviceName: pulumi.Input<string>;
    /**
     * Size of the Savings Plan
     */
    size: pulumi.Input<number>;
}