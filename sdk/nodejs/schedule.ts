// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * [Squadcast schedules](https://support.squadcast.com/docs/schedules) are used to manage on-call scheduling & determine who will be notified when an incident is triggered.
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
 * const exampleSchedule = new squadcast.Schedule("exampleSchedule", {
 *     teamId: exampleTeam.then(exampleTeam => exampleTeam.id),
 *     color: "#9900ef",
 * });
 * ```
 *
 * ## Import
 *
 * teamID:scheduleName Use 'Get All Teams' API to get the id of the team
 *
 * ```sh
 *  $ pulumi import squadcast:index/schedule:Schedule test "62d2fe23a57381088224d726:Example Schedule"
 * ```
 */
export class Schedule extends pulumi.CustomResource {
    /**
     * Get an existing Schedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScheduleState, opts?: pulumi.CustomResourceOptions): Schedule {
        return new Schedule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'squadcast:index/schedule:Schedule';

    /**
     * Returns true if the given object is an instance of Schedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Schedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Schedule.__pulumiType;
    }

    /**
     * Calendar color scheme for this schedule, hex values.
     */
    public readonly color!: pulumi.Output<string>;
    /**
     * Detailed description about the Schedule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the Schedule.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Team id.
     */
    public readonly teamId!: pulumi.Output<string>;

    /**
     * Create a Schedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScheduleArgs | ScheduleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScheduleState | undefined;
            resourceInputs["color"] = state ? state.color : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
        } else {
            const args = argsOrState as ScheduleArgs | undefined;
            if ((!args || args.color === undefined) && !opts.urn) {
                throw new Error("Missing required property 'color'");
            }
            if ((!args || args.teamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamId'");
            }
            resourceInputs["color"] = args ? args.color : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Schedule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Schedule resources.
 */
export interface ScheduleState {
    /**
     * Calendar color scheme for this schedule, hex values.
     */
    color?: pulumi.Input<string>;
    /**
     * Detailed description about the Schedule.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the Schedule.
     */
    name?: pulumi.Input<string>;
    /**
     * Team id.
     */
    teamId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Schedule resource.
 */
export interface ScheduleArgs {
    /**
     * Calendar color scheme for this schedule, hex values.
     */
    color: pulumi.Input<string>;
    /**
     * Detailed description about the Schedule.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the Schedule.
     */
    name?: pulumi.Input<string>;
    /**
     * Team id.
     */
    teamId: pulumi.Input<string>;
}
