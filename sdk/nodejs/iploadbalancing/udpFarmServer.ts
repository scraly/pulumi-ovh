// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a backend server entry linked to loadbalancing group (farm)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const lb = ovh.IpLoadBalancing.getIpLoadBalancing({
 *     serviceName: "ip-1.2.3.4",
 *     state: "ok",
 * });
 * const farmname = new ovh.iploadbalancing.UdpFarm("farmname", {
 *     displayName: "ingress-8080-gra",
 *     port: 80,
 *     serviceName: lb.then(lb => lb.serviceName),
 *     zone: "gra",
 * });
 * const backend = new ovh.iploadbalancing.UdpFarmServer("backend", {
 *     address: "4.5.6.7",
 *     displayName: "mybackend",
 *     farmId: farmname.farmId,
 *     port: 80,
 *     serviceName: lb.then(lb => lb.serviceName),
 *     status: "active",
 * });
 * ```
 *
 * ## Import
 *
 * UDP farm server can be imported using the following format `serviceName`, the `id` of the farm and the `id` of the server separated by "/" e.g.
 */
export class UdpFarmServer extends pulumi.CustomResource {
    /**
     * Get an existing UdpFarmServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UdpFarmServerState, opts?: pulumi.CustomResourceOptions): UdpFarmServer {
        return new UdpFarmServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:IpLoadBalancing/udpFarmServer:UdpFarmServer';

    /**
     * Returns true if the given object is an instance of UdpFarmServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UdpFarmServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UdpFarmServer.__pulumiType;
    }

    /**
     * Address of the backend server (IP from either internal or OVHcloud network)
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * Synonym for `farmId`.
     */
    public /*out*/ readonly backendId!: pulumi.Output<number>;
    /**
     * Label for the server
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * ID of the farm this server is attached to
     */
    public readonly farmId!: pulumi.Output<number>;
    /**
     * Port that backend will respond on
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Id of your server.
     */
    public /*out*/ readonly serverId!: pulumi.Output<number>;
    /**
     * The internal name of your IP load balancing
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * backend status - `active` or `inactive`
     */
    public readonly status!: pulumi.Output<string>;

    /**
     * Create a UdpFarmServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UdpFarmServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UdpFarmServerArgs | UdpFarmServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UdpFarmServerState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["backendId"] = state ? state.backendId : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["farmId"] = state ? state.farmId : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["serverId"] = state ? state.serverId : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as UdpFarmServerArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.farmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'farmId'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.status === undefined) && !opts.urn) {
                throw new Error("Missing required property 'status'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["farmId"] = args ? args.farmId : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["status"] = args ? args.status : undefined;
            resourceInputs["backendId"] = undefined /*out*/;
            resourceInputs["serverId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UdpFarmServer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UdpFarmServer resources.
 */
export interface UdpFarmServerState {
    /**
     * Address of the backend server (IP from either internal or OVHcloud network)
     */
    address?: pulumi.Input<string>;
    /**
     * Synonym for `farmId`.
     */
    backendId?: pulumi.Input<number>;
    /**
     * Label for the server
     */
    displayName?: pulumi.Input<string>;
    /**
     * ID of the farm this server is attached to
     */
    farmId?: pulumi.Input<number>;
    /**
     * Port that backend will respond on
     */
    port?: pulumi.Input<number>;
    /**
     * Id of your server.
     */
    serverId?: pulumi.Input<number>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName?: pulumi.Input<string>;
    /**
     * backend status - `active` or `inactive`
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UdpFarmServer resource.
 */
export interface UdpFarmServerArgs {
    /**
     * Address of the backend server (IP from either internal or OVHcloud network)
     */
    address: pulumi.Input<string>;
    /**
     * Label for the server
     */
    displayName?: pulumi.Input<string>;
    /**
     * ID of the farm this server is attached to
     */
    farmId: pulumi.Input<number>;
    /**
     * Port that backend will respond on
     */
    port?: pulumi.Input<number>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName: pulumi.Input<string>;
    /**
     * backend status - `active` or `inactive`
     */
    status: pulumi.Input<string>;
}