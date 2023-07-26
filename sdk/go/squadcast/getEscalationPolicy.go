// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package squadcast

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/valisinsights/pulumi-squadcast/sdk/go/squadcast/internal"
)

// [Escalation Policies](https://support.squadcast.com/docs/escalation-policies) defines rules indicating when and how alerts will escalate to various Users, Squads and (or) Schedules within your Organization.Use this data source to get information about a specific escalation policy that you can use for other Squadcast resources.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/valisinsights/pulumi-squadcast/sdk/go/squadcast"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := squadcast.LookupEscalationPolicy(ctx, &squadcast.LookupEscalationPolicyArgs{
//				Name:   squadcast_escalation_policy.Test.Name,
//				TeamId: "team id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupEscalationPolicy(ctx *pulumi.Context, args *LookupEscalationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupEscalationPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEscalationPolicyResult
	err := ctx.Invoke("squadcast:index/getEscalationPolicy:getEscalationPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEscalationPolicy.
type LookupEscalationPolicyArgs struct {
	// Name of the Escalation Policy
	Name string `pulumi:"name"`
	// Team id.
	TeamId string `pulumi:"teamId"`
}

// A collection of values returned by getEscalationPolicy.
type LookupEscalationPolicyResult struct {
	// Detailed description about the nature & purpose Escalation policy
	Description string `pulumi:"description"`
	// Escalation policy owner
	EntityOwners []GetEscalationPolicyEntityOwner `pulumi:"entityOwners"`
	// EscalationPolicy id.
	Id string `pulumi:"id"`
	// Name of the Escalation Policy
	Name string `pulumi:"name"`
	// You can choose to repeate the entire policy, if no one acknowledges the incident even after the Escalation Policy has been executed fully once
	Repeats []GetEscalationPolicyRepeat `pulumi:"repeats"`
	// Rules will have the details of who to notify and when to notify and how to notify them.
	Rules []GetEscalationPolicyRule `pulumi:"rules"`
	// Team id.
	TeamId string `pulumi:"teamId"`
}

func LookupEscalationPolicyOutput(ctx *pulumi.Context, args LookupEscalationPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupEscalationPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEscalationPolicyResult, error) {
			args := v.(LookupEscalationPolicyArgs)
			r, err := LookupEscalationPolicy(ctx, &args, opts...)
			var s LookupEscalationPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEscalationPolicyResultOutput)
}

// A collection of arguments for invoking getEscalationPolicy.
type LookupEscalationPolicyOutputArgs struct {
	// Name of the Escalation Policy
	Name pulumi.StringInput `pulumi:"name"`
	// Team id.
	TeamId pulumi.StringInput `pulumi:"teamId"`
}

func (LookupEscalationPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEscalationPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getEscalationPolicy.
type LookupEscalationPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupEscalationPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEscalationPolicyResult)(nil)).Elem()
}

func (o LookupEscalationPolicyResultOutput) ToLookupEscalationPolicyResultOutput() LookupEscalationPolicyResultOutput {
	return o
}

func (o LookupEscalationPolicyResultOutput) ToLookupEscalationPolicyResultOutputWithContext(ctx context.Context) LookupEscalationPolicyResultOutput {
	return o
}

// Detailed description about the nature & purpose Escalation policy
func (o LookupEscalationPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEscalationPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Escalation policy owner
func (o LookupEscalationPolicyResultOutput) EntityOwners() GetEscalationPolicyEntityOwnerArrayOutput {
	return o.ApplyT(func(v LookupEscalationPolicyResult) []GetEscalationPolicyEntityOwner { return v.EntityOwners }).(GetEscalationPolicyEntityOwnerArrayOutput)
}

// EscalationPolicy id.
func (o LookupEscalationPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEscalationPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the Escalation Policy
func (o LookupEscalationPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEscalationPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// You can choose to repeate the entire policy, if no one acknowledges the incident even after the Escalation Policy has been executed fully once
func (o LookupEscalationPolicyResultOutput) Repeats() GetEscalationPolicyRepeatArrayOutput {
	return o.ApplyT(func(v LookupEscalationPolicyResult) []GetEscalationPolicyRepeat { return v.Repeats }).(GetEscalationPolicyRepeatArrayOutput)
}

// Rules will have the details of who to notify and when to notify and how to notify them.
func (o LookupEscalationPolicyResultOutput) Rules() GetEscalationPolicyRuleArrayOutput {
	return o.ApplyT(func(v LookupEscalationPolicyResult) []GetEscalationPolicyRule { return v.Rules }).(GetEscalationPolicyRuleArrayOutput)
}

// Team id.
func (o LookupEscalationPolicyResultOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEscalationPolicyResult) string { return v.TeamId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEscalationPolicyResultOutput{})
}
