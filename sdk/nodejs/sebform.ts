// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * [Squadcast Webforms](https://support.squadcast.com/webforms/webforms) allows organizations to expand their customer support by hosting public Webforms, so their customers can quickly create an alert from outside the Squadcast ecosystem. Not only this, but internal stakeholders can also leverage Webforms for easy alert creation.
 *
 * ## Import
 *
 * teamID:webformName Use 'Get All Teams' API to get the id of the team
 *
 * ```sh
 *  $ pulumi import squadcast:index/sebform:Sebform example_webform "63065e992a5f9a1d5792b6c5:Webform Name"
 * ```
 */
export class Sebform extends pulumi.CustomResource {
    /**
     * Get an existing Sebform resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SebformState, opts?: pulumi.CustomResourceOptions): Sebform {
        return new Sebform(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'squadcast:index/sebform:Sebform';

    /**
     * Returns true if the given object is an instance of Sebform.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Sebform {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Sebform.__pulumiType;
    }

    /**
     * Custom domain name (URL).
     */
    public readonly customDomainName!: pulumi.Output<string | undefined>;
    /**
     * Description of the Webform.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Defines when to send email to the reporter (triggered, acknowledged, resolved).
     */
    public readonly emailOns!: pulumi.Output<string[] | undefined>;
    /**
     * Footer link.
     */
    public readonly footerLink!: pulumi.Output<string | undefined>;
    /**
     * Footer text.
     */
    public readonly footerText!: pulumi.Output<string | undefined>;
    /**
     * Webform header.
     */
    public readonly header!: pulumi.Output<string>;
    /**
     * Input Fields added to Webforms. Added as tags to incident based on selection.
     */
    public readonly inputFields!: pulumi.Output<outputs.SebformInputField[] | undefined>;
    /**
     * Name of the Webform.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Form owner.
     */
    public readonly owner!: pulumi.Output<outputs.SebformOwner>;
    /**
     * Public URL of the Webform.
     */
    public /*out*/ readonly publicUrl!: pulumi.Output<string>;
    /**
     * Services added to Webform.
     */
    public readonly services!: pulumi.Output<outputs.SebformService[]>;
    /**
     * Severity of the incident.
     *
     * @deprecated Use `input_field` instead of `severity`.
     */
    public readonly severities!: pulumi.Output<outputs.SebformSeverity[] | undefined>;
    /**
     * Webform Tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Team id.
     */
    public readonly teamId!: pulumi.Output<string>;
    /**
     * Webform title (public).
     */
    public readonly title!: pulumi.Output<string>;

    /**
     * Create a Sebform resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SebformArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SebformArgs | SebformState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SebformState | undefined;
            resourceInputs["customDomainName"] = state ? state.customDomainName : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["emailOns"] = state ? state.emailOns : undefined;
            resourceInputs["footerLink"] = state ? state.footerLink : undefined;
            resourceInputs["footerText"] = state ? state.footerText : undefined;
            resourceInputs["header"] = state ? state.header : undefined;
            resourceInputs["inputFields"] = state ? state.inputFields : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["publicUrl"] = state ? state.publicUrl : undefined;
            resourceInputs["services"] = state ? state.services : undefined;
            resourceInputs["severities"] = state ? state.severities : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["teamId"] = state ? state.teamId : undefined;
            resourceInputs["title"] = state ? state.title : undefined;
        } else {
            const args = argsOrState as SebformArgs | undefined;
            if ((!args || args.header === undefined) && !opts.urn) {
                throw new Error("Missing required property 'header'");
            }
            if ((!args || args.owner === undefined) && !opts.urn) {
                throw new Error("Missing required property 'owner'");
            }
            if ((!args || args.services === undefined) && !opts.urn) {
                throw new Error("Missing required property 'services'");
            }
            if ((!args || args.teamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'teamId'");
            }
            if ((!args || args.title === undefined) && !opts.urn) {
                throw new Error("Missing required property 'title'");
            }
            resourceInputs["customDomainName"] = args ? args.customDomainName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["emailOns"] = args ? args.emailOns : undefined;
            resourceInputs["footerLink"] = args ? args.footerLink : undefined;
            resourceInputs["footerText"] = args ? args.footerText : undefined;
            resourceInputs["header"] = args ? args.header : undefined;
            resourceInputs["inputFields"] = args ? args.inputFields : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["owner"] = args ? args.owner : undefined;
            resourceInputs["services"] = args ? args.services : undefined;
            resourceInputs["severities"] = args ? args.severities : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["teamId"] = args ? args.teamId : undefined;
            resourceInputs["title"] = args ? args.title : undefined;
            resourceInputs["publicUrl"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Sebform.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Sebform resources.
 */
export interface SebformState {
    /**
     * Custom domain name (URL).
     */
    customDomainName?: pulumi.Input<string>;
    /**
     * Description of the Webform.
     */
    description?: pulumi.Input<string>;
    /**
     * Defines when to send email to the reporter (triggered, acknowledged, resolved).
     */
    emailOns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Footer link.
     */
    footerLink?: pulumi.Input<string>;
    /**
     * Footer text.
     */
    footerText?: pulumi.Input<string>;
    /**
     * Webform header.
     */
    header?: pulumi.Input<string>;
    /**
     * Input Fields added to Webforms. Added as tags to incident based on selection.
     */
    inputFields?: pulumi.Input<pulumi.Input<inputs.SebformInputField>[]>;
    /**
     * Name of the Webform.
     */
    name?: pulumi.Input<string>;
    /**
     * Form owner.
     */
    owner?: pulumi.Input<inputs.SebformOwner>;
    /**
     * Public URL of the Webform.
     */
    publicUrl?: pulumi.Input<string>;
    /**
     * Services added to Webform.
     */
    services?: pulumi.Input<pulumi.Input<inputs.SebformService>[]>;
    /**
     * Severity of the incident.
     *
     * @deprecated Use `input_field` instead of `severity`.
     */
    severities?: pulumi.Input<pulumi.Input<inputs.SebformSeverity>[]>;
    /**
     * Webform Tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Team id.
     */
    teamId?: pulumi.Input<string>;
    /**
     * Webform title (public).
     */
    title?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Sebform resource.
 */
export interface SebformArgs {
    /**
     * Custom domain name (URL).
     */
    customDomainName?: pulumi.Input<string>;
    /**
     * Description of the Webform.
     */
    description?: pulumi.Input<string>;
    /**
     * Defines when to send email to the reporter (triggered, acknowledged, resolved).
     */
    emailOns?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Footer link.
     */
    footerLink?: pulumi.Input<string>;
    /**
     * Footer text.
     */
    footerText?: pulumi.Input<string>;
    /**
     * Webform header.
     */
    header: pulumi.Input<string>;
    /**
     * Input Fields added to Webforms. Added as tags to incident based on selection.
     */
    inputFields?: pulumi.Input<pulumi.Input<inputs.SebformInputField>[]>;
    /**
     * Name of the Webform.
     */
    name?: pulumi.Input<string>;
    /**
     * Form owner.
     */
    owner: pulumi.Input<inputs.SebformOwner>;
    /**
     * Services added to Webform.
     */
    services: pulumi.Input<pulumi.Input<inputs.SebformService>[]>;
    /**
     * Severity of the incident.
     *
     * @deprecated Use `input_field` instead of `severity`.
     */
    severities?: pulumi.Input<pulumi.Input<inputs.SebformSeverity>[]>;
    /**
     * Webform Tags.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Team id.
     */
    teamId: pulumi.Input<string>;
    /**
     * Webform title (public).
     */
    title: pulumi.Input<string>;
}