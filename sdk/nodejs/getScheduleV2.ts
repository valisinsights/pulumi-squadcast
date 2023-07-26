// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * [Squadcast schedules](https://support.squadcast.com/docs/schedules) are used to manage on-call scheduling & determine who will be notified when an incident is triggered. Use this data source to get information about a specific schedule that you can use for other Squadcast resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as squadcast from "@pulumi/squadcast";
 *
 * const test = squadcast.getScheduleV2({
 *     name: squadcast_schedule_v2.test.name,
 *     teamId: "team_id",
 * });
 * ```
 */
export function getScheduleV2(args: GetScheduleV2Args, opts?: pulumi.InvokeOptions): Promise<GetScheduleV2Result> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("squadcast:index/getScheduleV2:getScheduleV2", {
        "name": args.name,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getScheduleV2.
 */
export interface GetScheduleV2Args {
    /**
     * Name of the Schedule.
     */
    name: string;
    /**
     * Team id.
     */
    teamId: string;
}

/**
 * A collection of values returned by getScheduleV2.
 */
export interface GetScheduleV2Result {
    /**
     * Detailed description about the schedule.
     */
    readonly description: string;
    /**
     * Schedule owner.
     */
    readonly entityOwners: outputs.GetScheduleV2EntityOwner[];
    /**
     * Schedule id.
     */
    readonly id: string;
    /**
     * Name of the Schedule.
     */
    readonly name: string;
    /**
     * Schedule tags.
     */
    readonly tags: outputs.GetScheduleV2Tag[];
    /**
     * Team id.
     */
    readonly teamId: string;
    /**
     * Timezone for the schedule.
     */
    readonly timezone: string;
}
/**
 * [Squadcast schedules](https://support.squadcast.com/docs/schedules) are used to manage on-call scheduling & determine who will be notified when an incident is triggered. Use this data source to get information about a specific schedule that you can use for other Squadcast resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as squadcast from "@pulumi/squadcast";
 *
 * const test = squadcast.getScheduleV2({
 *     name: squadcast_schedule_v2.test.name,
 *     teamId: "team_id",
 * });
 * ```
 */
export function getScheduleV2Output(args: GetScheduleV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScheduleV2Result> {
    return pulumi.output(args).apply((a: any) => getScheduleV2(a, opts))
}

/**
 * A collection of arguments for invoking getScheduleV2.
 */
export interface GetScheduleV2OutputArgs {
    /**
     * Name of the Schedule.
     */
    name: pulumi.Input<string>;
    /**
     * Team id.
     */
    teamId: pulumi.Input<string>;
}