// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * A Runbook is a compilation of routine procedures and operations that are documented for reference while working on a critical incident. Sometimes, it can also be referred to as a Playbook.Use this data source to get information about a specific Runbook that you can use for other Squadcast resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as squadcast from "@pulumi/squadcast";
 *
 * const test = squadcast.getRunbook({
 *     name: squadcast_runbook.test.name,
 *     teamId: "team id",
 * });
 * ```
 */
export function getRunbook(args: GetRunbookArgs, opts?: pulumi.InvokeOptions): Promise<GetRunbookResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("squadcast:index/getRunbook:getRunbook", {
        "name": args.name,
        "teamId": args.teamId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRunbook.
 */
export interface GetRunbookArgs {
    /**
     * Name of the Runbook
     */
    name: string;
    /**
     * Team id.
     */
    teamId: string;
}

/**
 * A collection of values returned by getRunbook.
 */
export interface GetRunbookResult {
    /**
     * Runbooks owner
     */
    readonly entityOwners: outputs.GetRunbookEntityOwner[];
    /**
     * Runbook id.
     */
    readonly id: string;
    /**
     * Name of the Runbook
     */
    readonly name: string;
    /**
     * Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
     */
    readonly steps: outputs.GetRunbookStep[];
    /**
     * Team id.
     */
    readonly teamId: string;
}
/**
 * A Runbook is a compilation of routine procedures and operations that are documented for reference while working on a critical incident. Sometimes, it can also be referred to as a Playbook.Use this data source to get information about a specific Runbook that you can use for other Squadcast resources.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as squadcast from "@pulumi/squadcast";
 *
 * const test = squadcast.getRunbook({
 *     name: squadcast_runbook.test.name,
 *     teamId: "team id",
 * });
 * ```
 */
export function getRunbookOutput(args: GetRunbookOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRunbookResult> {
    return pulumi.output(args).apply((a: any) => getRunbook(a, opts))
}

/**
 * A collection of arguments for invoking getRunbook.
 */
export interface GetRunbookOutputArgs {
    /**
     * Name of the Runbook
     */
    name: pulumi.Input<string>;
    /**
     * Team id.
     */
    teamId: pulumi.Input<string>;
}
