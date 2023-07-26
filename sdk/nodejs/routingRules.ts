// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * [Routing rules](https://support.squadcast.com/docs/alert-routing) allows you to ensure that alerts are routed to the right responder with the help of `event tags` attached to them.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as squadcast from "@pulumi/squadcast";
 *
 * const exampleTeam = squadcast.getTeam({
 *     name: "example team name",
 * });
 * const exampleUser = squadcast.getUser({
 *     email: "test@example.com",
 * });
 * const exampleService = exampleTeam.then(exampleTeam => squadcast.getService({
 *     name: "example service name",
 *     teamId: exampleTeam.id,
 * }));
 * const exampleEscalaionPolicy = exampleTeam.then(exampleTeam => squadcast.getEscalationPolicy({
 *     name: "example escalation policy name",
 *     teamId: exampleTeam.id,
 * }));
 * const exampleSquad = exampleTeam.then(exampleTeam => squadcast.getSquad({
 *     name: "example squad name",
 *     teamId: exampleTeam.id,
 * }));
 * const exampleRoutingRules = new squadcast.RoutingRules("exampleRoutingRules", {
 *     teamId: exampleTeam.then(exampleTeam => exampleTeam.id),
 *     serviceId: exampleService.then(exampleService => exampleService.id),
 *     rules: [
 *         {
 *             isBasic: false,
 *             expression: "payload[\"event_id\"] == 40",
 *             routeToId: Promise.all([exampleUser, exampleSquad, exampleEscalaionPolicy]).then(([exampleUser, exampleSquad, exampleEscalaionPolicy]) => exampleUser.id / exampleSquad.id / exampleEscalaionPolicy.id),
 *             routeToType: "user/squad/escalation_policy",
 *         },
 *         {
 *             isBasic: true,
 *             basicExpressions: [{
 *                 lhs: "payload[\"foo\"]",
 *                 rhs: "bar",
 *             }],
 *             routeToId: Promise.all([exampleUser, exampleSquad, exampleEscalaionPolicy]).then(([exampleUser, exampleSquad, exampleEscalaionPolicy]) => exampleUser.id / exampleSquad.id / exampleEscalaionPolicy.id),
 *             routeToType: "user/squad/escalation_policy",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * teamID:serviceID Use 'Get All Teams' and 'Get All Services' APIs to get the id of the team and service respectively
 *
 * ```sh
 *  $ pulumi import squadcast:index/routingRules:RoutingRules test 62d2fe23a57381088224d726:62da76c088f407f9ca756ca5
 * ```
 */
export class RoutingRules extends pulumi.CustomResource {
    /**
     * Get an existing RoutingRules resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RoutingRulesState, opts?: pulumi.CustomResourceOptions): RoutingRules {
        return new RoutingRules(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'squadcast:index/routingRules:RoutingRules';

    /**
     * Returns true if the given object is an instance of RoutingRules.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoutingRules {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoutingRules.__pulumiType;
    }

    public readonly rules!: pulumi.Output<outputs.RoutingRulesRule[]>;
    /**
     * Service id.
     */
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * Team id.
     */
    public readonly teamId!: pulumi.Output<string>;

    /**
     * Create a RoutingRules resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoutingRulesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RoutingRulesArgs | RoutingRulesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RoutingRulesState | undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["serviceId"] = state ? state.serviceId : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
        } else {
            const args = argsOrState as RoutingRulesArgs | undefined;
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            if ((!args || args.teamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamId'");
            }
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RoutingRules.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RoutingRules resources.
 */
export interface RoutingRulesState {
    rules?: pulumi.Input<pulumi.Input<inputs.RoutingRulesRule>[]>;
    /**
     * Service id.
     */
    serviceId?: pulumi.Input<string>;
    /**
     * Team id.
     */
    teamId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RoutingRules resource.
 */
export interface RoutingRulesArgs {
    rules: pulumi.Input<pulumi.Input<inputs.RoutingRulesRule>[]>;
    /**
     * Service id.
     */
    serviceId: pulumi.Input<string>;
    /**
     * Team id.
     */
    teamId: pulumi.Input<string>;
}