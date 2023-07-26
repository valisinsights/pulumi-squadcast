// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * [Squadcast Services](https://support.squadcast.com/docs/adding-a-service-1) are the core components of your infrastructure/application for which alerts are generated. Services in Squadcast represent specific systems, applications, components, products, or teams for which an incident is created. To check out some of the best practices on creating Services in Squadcast, refer to the guide [here](https://www.squadcast.com/blog/how-to-configure-services-in-squadcast-best-practices-to-reduce-mttr).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as squadcast from "@pulumi/squadcast";
 *
 * const exampleUser = squadcast.getUser({
 *     email: "test@example.com",
 * });
 * const exampleTeam = squadcast.getTeam({
 *     name: "example team name",
 * });
 * const exampleEscalaionPolicy = exampleTeam.then(exampleTeam => squadcast.getEscalationPolicy({
 *     name: "example escalation policy name",
 *     teamId: exampleTeam.id,
 * }));
 * const exampleService = new squadcast.Service("exampleService", {
 *     teamId: exampleTeam.then(exampleTeam => exampleTeam.id),
 *     escalationPolicyId: exampleEscalaionPolicy.then(exampleEscalaionPolicy => exampleEscalaionPolicy.id),
 *     emailPrefix: "example-service-email",
 *     maintainer: {
 *         id: exampleUser.then(exampleUser => exampleUser.id),
 *         type: "user",
 *     },
 *     tags: [
 *         {
 *             key: "testkey",
 *             value: "testval",
 *         },
 *         {
 *             key: "testkey2",
 *             value: "testval2",
 *         },
 *     ],
 *     alertSources: ["example-alert-source"],
 *     slackChannelId: "D0KAQDEPSH",
 * });
 * ```
 *
 * ## Import
 *
 * teamID:serviceID Use 'Get All Teams' and 'Get All Services' APIs to get the id of the team and service respectively
 *
 * ```sh
 *  $ pulumi import squadcast:index/service:Service test_parent 62d2fe23a57381088224d726:62da76c088f407f9ca756ca5
 * ```
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'squadcast:index/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * Active alert source endpoints.
     */
    public /*out*/ readonly activeAlertSourceEndpoints!: pulumi.Output<{[key: string]: string}>;
    /**
     * All available alert source endpoints.
     */
    public /*out*/ readonly alertSourceEndpoints!: pulumi.Output<{[key: string]: string}>;
    /**
     * List of active alert source names.
     */
    public readonly alertSources!: pulumi.Output<string[] | undefined>;
    /**
     * Unique API key of this service.
     */
    public /*out*/ readonly apiKey!: pulumi.Output<string>;
    /**
     * Dependencies (serviceIds)
     */
    public readonly dependencies!: pulumi.Output<string[] | undefined>;
    /**
     * Detailed description about this service.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Email.
     */
    public /*out*/ readonly email!: pulumi.Output<string>;
    /**
     * Email prefix.
     */
    public readonly emailPrefix!: pulumi.Output<string>;
    /**
     * Escalation policy id.
     */
    public readonly escalationPolicyId!: pulumi.Output<string>;
    /**
     * Service owner.
     */
    public readonly maintainer!: pulumi.Output<outputs.ServiceMaintainer>;
    /**
     * Name of the Service.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Slack extension for the service. If set, specifies the ID of the Slack channel associated with the service. If this ID is set, it cannot be removed, but it can be changed to a different slack*channel*id.
     */
    public readonly slackChannelId!: pulumi.Output<string>;
    /**
     * Service tags.
     */
    public readonly tags!: pulumi.Output<outputs.ServiceTag[] | undefined>;
    /**
     * Team id.
     */
    public readonly teamId!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            resourceInputs["activeAlertSourceEndpoints"] = state ? state.activeAlertSourceEndpoints : undefined;
            resourceInputs["alertSourceEndpoints"] = state ? state.alertSourceEndpoints : undefined;
            resourceInputs["alertSources"] = state ? state.alertSources : undefined;
            resourceInputs["apiKey"] = state ? state.apiKey : undefined;
            resourceInputs["dependencies"] = state ? state.dependencies : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["emailPrefix"] = state ? state.emailPrefix : undefined;
            resourceInputs["escalationPolicyId"] = state ? state.escalationPolicyId : undefined;
            resourceInputs["maintainer"] = state ? state.maintainer : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["slackChannelId"] = state ? state.slackChannelId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if ((!args || args.emailPrefix === undefined) && !opts.urn) {
                throw new Error("Missing required property 'emailPrefix'");
            }
            if ((!args || args.escalationPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'escalationPolicyId'");
            }
            if ((!args || args.teamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamId'");
            }
            resourceInputs["alertSources"] = args ? args.alertSources : undefined;
            resourceInputs["dependencies"] = args ? args.dependencies : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["emailPrefix"] = args ? args.emailPrefix : undefined;
            resourceInputs["escalationPolicyId"] = args ? args.escalationPolicyId : undefined;
            resourceInputs["maintainer"] = args ? args.maintainer : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["slackChannelId"] = args ? args.slackChannelId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["activeAlertSourceEndpoints"] = undefined /*out*/;
            resourceInputs["alertSourceEndpoints"] = undefined /*out*/;
            resourceInputs["apiKey"] = undefined /*out*/;
            resourceInputs["email"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * Active alert source endpoints.
     */
    activeAlertSourceEndpoints?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * All available alert source endpoints.
     */
    alertSourceEndpoints?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of active alert source names.
     */
    alertSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Unique API key of this service.
     */
    apiKey?: pulumi.Input<string>;
    /**
     * Dependencies (serviceIds)
     */
    dependencies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Detailed description about this service.
     */
    description?: pulumi.Input<string>;
    /**
     * Email.
     */
    email?: pulumi.Input<string>;
    /**
     * Email prefix.
     */
    emailPrefix?: pulumi.Input<string>;
    /**
     * Escalation policy id.
     */
    escalationPolicyId?: pulumi.Input<string>;
    /**
     * Service owner.
     */
    maintainer?: pulumi.Input<inputs.ServiceMaintainer>;
    /**
     * Name of the Service.
     */
    name?: pulumi.Input<string>;
    /**
     * Slack extension for the service. If set, specifies the ID of the Slack channel associated with the service. If this ID is set, it cannot be removed, but it can be changed to a different slack*channel*id.
     */
    slackChannelId?: pulumi.Input<string>;
    /**
     * Service tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ServiceTag>[]>;
    /**
     * Team id.
     */
    teamId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * List of active alert source names.
     */
    alertSources?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Dependencies (serviceIds)
     */
    dependencies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Detailed description about this service.
     */
    description?: pulumi.Input<string>;
    /**
     * Email prefix.
     */
    emailPrefix: pulumi.Input<string>;
    /**
     * Escalation policy id.
     */
    escalationPolicyId: pulumi.Input<string>;
    /**
     * Service owner.
     */
    maintainer?: pulumi.Input<inputs.ServiceMaintainer>;
    /**
     * Name of the Service.
     */
    name?: pulumi.Input<string>;
    /**
     * Slack extension for the service. If set, specifies the ID of the Slack channel associated with the service. If this ID is set, it cannot be removed, but it can be changed to a different slack*channel*id.
     */
    slackChannelId?: pulumi.Input<string>;
    /**
     * Service tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ServiceTag>[]>;
    /**
     * Team id.
     */
    teamId: pulumi.Input<string>;
}
