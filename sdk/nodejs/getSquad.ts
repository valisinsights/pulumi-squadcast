// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * [Squads](https://support.squadcast.com/docs/squads) are smaller groups of members within Teams. Squads could correspond to groups of people that are responsible for specific projects within a Team.Use this data source to get information about a specific Squad.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as squadcast from "@pulumi/squadcast";
 *
 * const test = squadcast.getSquad({
 *     name: "test",
 *     teamId: "team id",
 * });
 * ```
 */
export function getSquad(args: GetSquadArgs, opts?: pulumi.InvokeOptions): Promise<GetSquadResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("squadcast:index/getSquad:getSquad", {
        "name": args.name,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSquad.
 */
export interface GetSquadArgs {
    /**
     * Name of the Squad.
     */
    name: string;
    /**
     * Team id.
     */
    teamId: string;
}

/**
 * A collection of values returned by getSquad.
 */
export interface GetSquadResult {
    /**
     * Squad id.
     */
    readonly id: string;
    readonly memberIds: string[];
    /**
     * Name of the Squad.
     */
    readonly name: string;
    /**
     * Team id.
     */
    readonly teamId: string;
}
/**
 * [Squads](https://support.squadcast.com/docs/squads) are smaller groups of members within Teams. Squads could correspond to groups of people that are responsible for specific projects within a Team.Use this data source to get information about a specific Squad.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as squadcast from "@pulumi/squadcast";
 *
 * const test = squadcast.getSquad({
 *     name: "test",
 *     teamId: "team id",
 * });
 * ```
 */
export function getSquadOutput(args: GetSquadOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSquadResult> {
    return pulumi.output(args).apply((a: any) => getSquad(a, opts))
}

/**
 * A collection of arguments for invoking getSquad.
 */
export interface GetSquadOutputArgs {
    /**
     * Name of the Squad.
     */
    name: pulumi.Input<string>;
    /**
     * Team id.
     */
    teamId: pulumi.Input<string>;
}
