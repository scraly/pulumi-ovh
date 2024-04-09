// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a backend server group (frontend) to be used by loadbalancing frontend(s)
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const lb = ovh.IpLoadBalancing.getIpLoadBalancing({
 *     serviceName: "ip-1.2.3.4",
 *     state: "ok",
 * });
 * const testfrontend = new ovh.iploadbalancing.UdpFrontend("testfrontend", {
 *     serviceName: lb.then(lb => lb.serviceName),
 *     displayName: "ingress-8080-gra",
 *     zone: "all",
 *     port: "10,11",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class UdpFrontend extends pulumi.CustomResource {
    /**
     * Get an existing UdpFrontend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UdpFrontendState, opts?: pulumi.CustomResourceOptions): UdpFrontend {
        return new UdpFrontend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:IpLoadBalancing/udpFrontend:UdpFrontend';

    /**
     * Returns true if the given object is an instance of UdpFrontend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UdpFrontend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UdpFrontend.__pulumiType;
    }

    /**
     * Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     */
    public readonly dedicatedIpfos!: pulumi.Output<string[]>;
    /**
     * Default UDP Farm of your frontend
     */
    public readonly defaultFarmId!: pulumi.Output<number>;
    /**
     * Disable your frontend. Default: 'false'
     */
    public readonly disabled!: pulumi.Output<boolean>;
    /**
     * Human readable name for your frontend
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Id of your frontend
     */
    public /*out*/ readonly frontendId!: pulumi.Output<number>;
    /**
     * Port(s) attached to your frontend. Supports single port (numerical value), 
     * range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
     * and/or 'range'. Each port must be in the [1;49151] range
     */
    public readonly port!: pulumi.Output<string>;
    /**
     * The internal name of your IP load balancing
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a UdpFrontend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UdpFrontendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UdpFrontendArgs | UdpFrontendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UdpFrontendState | undefined;
            resourceInputs["dedicatedIpfos"] = state ? state.dedicatedIpfos : undefined;
            resourceInputs["defaultFarmId"] = state ? state.defaultFarmId : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["frontendId"] = state ? state.frontendId : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as UdpFrontendArgs | undefined;
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["dedicatedIpfos"] = args ? args.dedicatedIpfos : undefined;
            resourceInputs["defaultFarmId"] = args ? args.defaultFarmId : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["frontendId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UdpFrontend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UdpFrontend resources.
 */
export interface UdpFrontendState {
    /**
     * Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     */
    dedicatedIpfos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default UDP Farm of your frontend
     */
    defaultFarmId?: pulumi.Input<number>;
    /**
     * Disable your frontend. Default: 'false'
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Human readable name for your frontend
     */
    displayName?: pulumi.Input<string>;
    /**
     * Id of your frontend
     */
    frontendId?: pulumi.Input<number>;
    /**
     * Port(s) attached to your frontend. Supports single port (numerical value), 
     * range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
     * and/or 'range'. Each port must be in the [1;49151] range
     */
    port?: pulumi.Input<string>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UdpFrontend resource.
 */
export interface UdpFrontendArgs {
    /**
     * Only attach frontend on these ip. No restriction if null. List of Ip blocks.
     */
    dedicatedIpfos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Default UDP Farm of your frontend
     */
    defaultFarmId?: pulumi.Input<number>;
    /**
     * Disable your frontend. Default: 'false'
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * Human readable name for your frontend
     */
    displayName?: pulumi.Input<string>;
    /**
     * Port(s) attached to your frontend. Supports single port (numerical value), 
     * range (2 dash-delimited increasing ports) and comma-separated list of 'single port'
     * and/or 'range'. Each port must be in the [1;49151] range
     */
    port: pulumi.Input<string>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName: pulumi.Input<string>;
    /**
     * Zone where the frontend will be defined (ie. `gra`, `bhs` also supports `all`)
     */
    zone: pulumi.Input<string>;
}
