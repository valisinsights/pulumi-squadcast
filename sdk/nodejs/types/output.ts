// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface DeduplicationRulesRule {
    /**
     * The basic expression which needs to be evaluated to be true for this rule to apply.
     */
    basicExpressions?: outputs.DeduplicationRulesRuleBasicExpression[];
    /**
     * Denotes if dependent services should also be deduplicated
     */
    dependencyDeduplication?: boolean;
    /**
     * description.
     */
    description?: string;
    /**
     * The expression which needs to be evaluated to be true for this rule to apply.
     */
    expression?: string;
    /**
     * is_basic will be true when users use the drop down selectors which will have lhs, op & rhs value, whereas it will be false when they use the advanced mode and it would have the expression for it's value
     */
    isBasic: boolean;
    /**
     * time unit (mins or hours)
     */
    timeUnit?: string;
    /**
     * integer for time_unit
     */
    timeWindow?: number;
}

export interface DeduplicationRulesRuleBasicExpression {
    lhs: string;
    op: string;
    rhs: string;
}

export interface EscalationPolicyEntityOwner {
    /**
     * Escalation policy owner id.
     */
    id: string;
    /**
     * Escalation policy owner type. (user or squad or team)
     */
    type: string;
}

export interface EscalationPolicyRepeat {
    /**
     * The number of minutes to wait before repeating the escalation policy
     */
    delayMinutes: number;
    /**
     * The number of times you want this escalation policy to be repeated, maximum allowed to repeat 3 times
     */
    times: number;
}

export interface EscalationPolicyRule {
    delayMinutes: number;
    notificationChannels?: string[];
    /**
     * repeat this rule
     */
    repeat?: outputs.EscalationPolicyRuleRepeat;
    roundRobin?: outputs.EscalationPolicyRuleRoundRobin;
    targets: outputs.EscalationPolicyRuleTarget[];
}

export interface EscalationPolicyRuleRepeat {
    /**
     * The number of minutes to wait before repeating the escalation policy
     */
    delayMinutes: number;
    /**
     * The number of times you want this escalation policy to be repeated, maximum allowed to repeat 3 times
     */
    times: number;
}

export interface EscalationPolicyRuleRoundRobin {
    enabled: boolean;
    rotation?: outputs.EscalationPolicyRuleRoundRobinRotation;
}

export interface EscalationPolicyRuleRoundRobinRotation {
    delayMinutes?: number;
    enabled?: boolean;
}

export interface EscalationPolicyRuleTarget {
    /**
     * EscalationPolicy id.
     */
    id: string;
    type: string;
}

export interface GetEscalationPolicyEntityOwner {
    /**
     * Escalation policy owner id.
     */
    id: string;
    /**
     * Escalation policy owner type. (user or squad or team)
     */
    type: string;
}

export interface GetEscalationPolicyRepeat {
    delayMinutes: number;
    times: number;
}

export interface GetEscalationPolicyRule {
    delayMinutes: number;
    notificationChannels: string[];
    /**
     * You can choose to repeate the entire policy, if no one acknowledges the incident even after the Escalation Policy has been executed fully once
     */
    repeats: outputs.GetEscalationPolicyRuleRepeat[];
    roundRobins: outputs.GetEscalationPolicyRuleRoundRobin[];
    targets: outputs.GetEscalationPolicyRuleTarget[];
}

export interface GetEscalationPolicyRuleRepeat {
    delayMinutes: number;
    times: number;
}

export interface GetEscalationPolicyRuleRoundRobin {
    enabled: boolean;
    rotations: outputs.GetEscalationPolicyRuleRoundRobinRotation[];
}

export interface GetEscalationPolicyRuleRoundRobinRotation {
    delayMinutes: number;
    enabled: boolean;
}

export interface GetEscalationPolicyRuleTarget {
    /**
     * EscalationPolicy id.
     */
    id: string;
    type: string;
}

export interface GetRunbookEntityOwner {
    /**
     * Runbook owner id.
     */
    id: string;
    /**
     * Runbook owner type. (user or squad or team)
     */
    type: string;
}

export interface GetRunbookStep {
    content: string;
}

export interface GetScheduleV2EntityOwner {
    /**
     * Schedule id.
     */
    id: string;
    type: string;
}

export interface GetScheduleV2Tag {
    color: string;
    key: string;
    value: string;
}

export interface GetServiceMaintainer {
    /**
     * The id of the maintainer.
     */
    id: string;
    /**
     * The type of the maintainer. (user, team or squad)
     */
    type: string;
}

export interface GetServiceTag {
    key: string;
    value: string;
}

export interface GetTeamMember {
    roleIds: string[];
    userId: string;
}

export interface GetTeamRole {
    abilities: string[];
    /**
     * Squadcast has one default team and this field let's us know if this is the default team.
     */
    default: boolean;
    /**
     * Team id.
     */
    id: string;
    /**
     * Name of the Team.
     */
    name: string;
}

export interface GetUserNotificationRule {
    delayMinutes: number;
    type: string;
}

export interface GetUserOncallReminderRule {
    delayMinutes: number;
    type: string;
}

export interface GetWebformInputField {
    /**
     * Input field Label.
     */
    label: string;
    /**
     * Input field options.
     */
    options: string[];
}

export interface GetWebformOwner {
    /**
     * Form owner id.
     */
    id: string;
    /**
     * Form owner name.
     */
    name: string;
    /**
     * Form owner type (user, team, squad).
     */
    type: string;
}

export interface GetWebformService {
    /**
     * Service alias.
     */
    alias: string;
    /**
     * Service name.
     */
    name: string;
    /**
     * Service id.
     */
    serviceId: string;
}

export interface GetWebformSeverity {
    /**
     * Severity description.
     */
    description: string;
    /**
     * Severity type.
     */
    type: string;
}

export interface RoutingRulesRule {
    /**
     * The basic expression which needs to be evaluated to be true for this rule to apply.
     */
    basicExpressions?: outputs.RoutingRulesRuleBasicExpression[];
    /**
     * The expression which needs to be evaluated to be true for this rule to apply.
     */
    expression?: string;
    /**
     * is_basic will be true when users use the drop down selectors which will have lhs, op & rhs value, whereas it will be false when they use the advanced mode and it would have the expression for it's value
     */
    isBasic: boolean;
    /**
     * The id of the entity (user, escalation policy, squad) for which we are routing this incident.
     */
    routeToId: string;
    /**
     * Type of the entity for which we are routing this incident - User, Escalation Policy or Squad
     */
    routeToType: string;
}

export interface RoutingRulesRuleBasicExpression {
    lhs: string;
    rhs: string;
}

export interface RunbookEntityOwner {
    /**
     * Runbook owner id.
     */
    id: string;
    /**
     * Runbook owner type. (user or squad or team)
     */
    type: string;
}

export interface RunbookStep {
    content: string;
}

export interface SLOEntityOwner {
    /**
     * SLO owner id.
     */
    id: string;
    /**
     * SLO owner type (user, team, squad).
     */
    type: string;
}

export interface SLONotify {
    /**
     * The ID of the notification rule
     */
    id: number;
    /**
     * The ID of the service in which the user want to create an incident
     */
    serviceId?: string;
    /**
     * The ID of the SLO.
     */
    sloId: number;
    /**
     * List of Squad ID's who should be alerted via email.
     */
    squadIds?: string[];
    /**
     * List of user ID's who should be alerted via email.
     */
    userIds?: string[];
}

export interface SLORule {
    /**
     * The ID of the monitoring rule
     */
    id: number;
    /**
     * Is checked?
     */
    isChecked: boolean;
    /**
     * The name of monitoring check."Supported values are "breached*error*budget", "unhealthy*slo","increased*false*positives", "remaining*error_budget"
     */
    name: string;
    /**
     * The ID of the SLO
     */
    sloId: number;
    /**
     * Threshold for the monitoring checkOnly supported for rules name "increased*false*positives" and "remaining*error*budget"
     */
    threshold?: number;
}

export interface ScheduleRotationV2ParticipantGroup {
    /**
     * Group participants.
     */
    participants?: outputs.ScheduleRotationV2ParticipantGroupParticipant[];
}

export interface ScheduleRotationV2ParticipantGroupParticipant {
    /**
     * Rotation id.
     */
    id: string;
    type: string;
}

export interface ScheduleRotationV2ShiftTimeslot {
    /**
     * Defines the day of the week for the shift. If not specified, the timeslot is active on all days of the week.
     */
    dayOfWeek?: string;
    /**
     * Defines the duration of each shift. (in minutes)
     */
    duration: number;
    /**
     * Defines the start hour of the each shift in the schedule timezone.
     */
    startHour: number;
    /**
     * Defines the start minute of the each shift in the schedule timezone.
     */
    startMinute: number;
}

export interface ScheduleV2EntityOwner {
    /**
     * Schedule owner id.
     */
    id: string;
    /**
     * Schedule owner type (user, team, squad).
     */
    type: string;
}

export interface ScheduleV2Tag {
    /**
     * Schedule tag color.
     */
    color: string;
    /**
     * Schedule tag key.
     */
    key: string;
    /**
     * Schedule tag value.
     */
    value: string;
}

export interface SebformInputField {
    /**
     * Input field Label.
     */
    label?: string;
    /**
     * Input field options.
     */
    options?: string[];
}

export interface SebformOwner {
    /**
     * Form owner id.
     */
    id: string;
    /**
     * Form owner name.
     */
    name: string;
    /**
     * Form owner type (user, team, squad).
     */
    type: string;
}

export interface SebformService {
    /**
     * Service alias.
     */
    alias?: string;
    /**
     * Service name.
     */
    name: string;
    /**
     * Service ID.
     */
    serviceId: string;
}

export interface SebformSeverity {
    /**
     * Severity description.
     */
    description?: string;
    /**
     * Severity type.
     */
    type: string;
}

export interface ServiceMaintainer {
    /**
     * The id of the maintainer.
     */
    id: string;
    /**
     * The type of the maintainer. (user, team or squad)
     */
    type: string;
}

export interface ServiceMaintenanceWindow {
    /**
     * Starting Time
     */
    from: string;
    /**
     * repeat frequency. ('day', 'week', '2 weeks', '3 weeks', 'month')
     */
    repeatFrequency?: string;
    /**
     * Till when you want to repeat this Maintenance mode
     */
    repeatTill?: string;
    /**
     * End Time.
     */
    till: string;
}

export interface ServiceTag {
    /**
     * key
     */
    key: string;
    /**
     * value
     */
    value: string;
}

export interface SuppressionRulesRule {
    /**
     * The basic expression which needs to be evaluated to be true for this rule to apply.
     */
    basicExpressions?: outputs.SuppressionRulesRuleBasicExpression[];
    /**
     * description.
     */
    description?: string;
    /**
     * The expression which needs to be evaluated to be true for this rule to apply.
     */
    expression?: string;
    /**
     * is_basic will be true when users use the drop down selectors which will have lhs, op & rhs value, whereas it will be false when they use the advanced mode and it would have the expression for it's value
     */
    isBasic: boolean;
    /**
     * is_timebased will be true when users use the time based suppression rule
     */
    isTimebased: boolean;
    /**
     * The timeslots for which this rule should be applied.
     */
    timeslots?: outputs.SuppressionRulesRuleTimeslot[];
}

export interface SuppressionRulesRuleBasicExpression {
    lhs: string;
    op: string;
    rhs: string;
}

export interface SuppressionRulesRuleTimeslot {
    customs?: outputs.SuppressionRulesRuleTimeslotCustom[];
    endTime: string;
    endsNever?: boolean;
    endsOn: string;
    isAllday?: boolean;
    isCustom: boolean;
    repetition: string;
    startTime: string;
    timeZone: string;
}

export interface SuppressionRulesRuleTimeslotCustom {
    repeats: string;
    repeatsCount?: number;
    repeatsOnMonth: string;
    repeatsOnWeekdays?: number[];
}

export interface TaggingRulesRule {
    /**
     * The basic expression which needs to be evaluated to be true for this rule to apply.
     */
    basicExpressions?: outputs.TaggingRulesRuleBasicExpression[];
    /**
     * The expression which needs to be evaluated to be true for this rule to apply.
     */
    expression?: string;
    /**
     * is_basic will be true when users use the drop down selectors which will have lhs, op & rhs value, whereas it will be false when they use the advanced mode and it would have the expression for it's value
     */
    isBasic: boolean;
    /**
     * The tags supposed to be set for a given payload(incident), Expression must be set when tags are empty and must contain addTags parameters.
     */
    tags?: outputs.TaggingRulesRuleTag[];
}

export interface TaggingRulesRuleBasicExpression {
    lhs: string;
    op: string;
    rhs: string;
}

export interface TaggingRulesRuleTag {
    color: string;
    key: string;
    value: string;
}

