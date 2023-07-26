// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package squadcast

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/valisinsights/pulumi-squadcast/sdk/go/squadcast/internal"
)

// [Escalation Policies](https://support.squadcast.com/docs/escalation-policies) defines rules indicating when and how alerts will escalate to various Users, Squads and (or) Schedules within your Organization.
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
//			exampleTeam, err := squadcast.LookupTeam(ctx, &squadcast.LookupTeamArgs{
//				Name: "example team name",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleUser, err := squadcast.LookupUser(ctx, &squadcast.LookupUserArgs{
//				Email: "test@example.com",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleSquad, err := squadcast.LookupSquad(ctx, &squadcast.LookupSquadArgs{
//				Name:   "example squad name",
//				TeamId: exampleTeam.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			exampleScheduleV2, err := squadcast.LookupScheduleV2(ctx, &squadcast.LookupScheduleV2Args{
//				Name:   "example schedule name",
//				TeamId: exampleTeam.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = squadcast.NewEscalationPolicy(ctx, "exampleEscalaionPolicy", &squadcast.EscalationPolicyArgs{
//				Description: pulumi.String("It's an amazing policy"),
//				TeamId:      *pulumi.String(exampleTeam.Id),
//				Rules: squadcast.EscalationPolicyRuleArray{
//					&squadcast.EscalationPolicyRuleArgs{
//						DelayMinutes: pulumi.Int(0),
//						Targets: squadcast.EscalationPolicyRuleTargetArray{
//							&squadcast.EscalationPolicyRuleTargetArgs{
//								Id:   *pulumi.String(exampleUser.Id),
//								Type: pulumi.String("user"),
//							},
//							&squadcast.EscalationPolicyRuleTargetArgs{
//								Id:   *pulumi.String(exampleScheduleV2.Id),
//								Type: pulumi.String("schedulev2"),
//							},
//						},
//					},
//					&squadcast.EscalationPolicyRuleArgs{
//						DelayMinutes: pulumi.Int(5),
//						Targets: squadcast.EscalationPolicyRuleTargetArray{
//							&squadcast.EscalationPolicyRuleTargetArgs{
//								Id:   *pulumi.String(exampleUser.Id),
//								Type: pulumi.String("user"),
//							},
//							&squadcast.EscalationPolicyRuleTargetArgs{
//								Id:   *pulumi.String(exampleSquad.Id),
//								Type: pulumi.String("squad"),
//							},
//						},
//						NotificationChannels: pulumi.StringArray{
//							pulumi.String("Phone"),
//						},
//						Repeat: &squadcast.EscalationPolicyRuleRepeatArgs{
//							Times:        pulumi.Int(1),
//							DelayMinutes: pulumi.Int(5),
//						},
//					},
//					&squadcast.EscalationPolicyRuleArgs{
//						DelayMinutes: pulumi.Int(10),
//						Targets: squadcast.EscalationPolicyRuleTargetArray{
//							&squadcast.EscalationPolicyRuleTargetArgs{
//								Id:   *pulumi.String(exampleSquad.Id),
//								Type: pulumi.String("squad"),
//							},
//							&squadcast.EscalationPolicyRuleTargetArgs{
//								Id:   *pulumi.String(exampleScheduleV2.Id),
//								Type: pulumi.String("schedulev2"),
//							},
//						},
//						RoundRobin: &squadcast.EscalationPolicyRuleRoundRobinArgs{
//							Enabled: pulumi.Bool(true),
//							Rotation: &squadcast.EscalationPolicyRuleRoundRobinRotationArgs{
//								Enabled:      pulumi.Bool(true),
//								DelayMinutes: pulumi.Int(1),
//							},
//						},
//					},
//				},
//				Repeat: &squadcast.EscalationPolicyRepeatArgs{
//					Times:        pulumi.Int(2),
//					DelayMinutes: pulumi.Int(10),
//				},
//				EntityOwner: &squadcast.EscalationPolicyEntityOwnerArgs{
//					Id:   *pulumi.String(exampleUser.Id),
//					Type: pulumi.String("user"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// teamID:escalationPolicyID Use 'Get All Teams' and 'Get All Escalation Policies' APIs to get the id of the team and escalation policy name respectively
//
// ```sh
//
//	$ pulumi import squadcast:index/escalationPolicy:EscalationPolicy test "62d2fe23a57381088224d726:Example Escalation Policy"
//
// ```
type EscalationPolicy struct {
	pulumi.CustomResourceState

	// Detailed description about the Escalation Policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Escalation policy owner.
	EntityOwner EscalationPolicyEntityOwnerOutput `pulumi:"entityOwner"`
	// Name of the Escalation Policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// You can choose to repeate the entire policy, if no one acknowledges the incident even after the Escalation Policy has been executed fully once
	Repeat EscalationPolicyRepeatPtrOutput `pulumi:"repeat"`
	// Rules will have the details of who to notify and when to notify and how to notify them.
	Rules EscalationPolicyRuleArrayOutput `pulumi:"rules"`
	// Team id.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewEscalationPolicy registers a new resource with the given unique name, arguments, and options.
func NewEscalationPolicy(ctx *pulumi.Context,
	name string, args *EscalationPolicyArgs, opts ...pulumi.ResourceOption) (*EscalationPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EscalationPolicy
	err := ctx.RegisterResource("squadcast:index/escalationPolicy:EscalationPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEscalationPolicy gets an existing EscalationPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEscalationPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EscalationPolicyState, opts ...pulumi.ResourceOption) (*EscalationPolicy, error) {
	var resource EscalationPolicy
	err := ctx.ReadResource("squadcast:index/escalationPolicy:EscalationPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EscalationPolicy resources.
type escalationPolicyState struct {
	// Detailed description about the Escalation Policy.
	Description *string `pulumi:"description"`
	// Escalation policy owner.
	EntityOwner *EscalationPolicyEntityOwner `pulumi:"entityOwner"`
	// Name of the Escalation Policy.
	Name *string `pulumi:"name"`
	// You can choose to repeate the entire policy, if no one acknowledges the incident even after the Escalation Policy has been executed fully once
	Repeat *EscalationPolicyRepeat `pulumi:"repeat"`
	// Rules will have the details of who to notify and when to notify and how to notify them.
	Rules []EscalationPolicyRule `pulumi:"rules"`
	// Team id.
	TeamId *string `pulumi:"teamId"`
}

type EscalationPolicyState struct {
	// Detailed description about the Escalation Policy.
	Description pulumi.StringPtrInput
	// Escalation policy owner.
	EntityOwner EscalationPolicyEntityOwnerPtrInput
	// Name of the Escalation Policy.
	Name pulumi.StringPtrInput
	// You can choose to repeate the entire policy, if no one acknowledges the incident even after the Escalation Policy has been executed fully once
	Repeat EscalationPolicyRepeatPtrInput
	// Rules will have the details of who to notify and when to notify and how to notify them.
	Rules EscalationPolicyRuleArrayInput
	// Team id.
	TeamId pulumi.StringPtrInput
}

func (EscalationPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*escalationPolicyState)(nil)).Elem()
}

type escalationPolicyArgs struct {
	// Detailed description about the Escalation Policy.
	Description *string `pulumi:"description"`
	// Escalation policy owner.
	EntityOwner *EscalationPolicyEntityOwner `pulumi:"entityOwner"`
	// Name of the Escalation Policy.
	Name *string `pulumi:"name"`
	// You can choose to repeate the entire policy, if no one acknowledges the incident even after the Escalation Policy has been executed fully once
	Repeat *EscalationPolicyRepeat `pulumi:"repeat"`
	// Rules will have the details of who to notify and when to notify and how to notify them.
	Rules []EscalationPolicyRule `pulumi:"rules"`
	// Team id.
	TeamId string `pulumi:"teamId"`
}

// The set of arguments for constructing a EscalationPolicy resource.
type EscalationPolicyArgs struct {
	// Detailed description about the Escalation Policy.
	Description pulumi.StringPtrInput
	// Escalation policy owner.
	EntityOwner EscalationPolicyEntityOwnerPtrInput
	// Name of the Escalation Policy.
	Name pulumi.StringPtrInput
	// You can choose to repeate the entire policy, if no one acknowledges the incident even after the Escalation Policy has been executed fully once
	Repeat EscalationPolicyRepeatPtrInput
	// Rules will have the details of who to notify and when to notify and how to notify them.
	Rules EscalationPolicyRuleArrayInput
	// Team id.
	TeamId pulumi.StringInput
}

func (EscalationPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*escalationPolicyArgs)(nil)).Elem()
}

type EscalationPolicyInput interface {
	pulumi.Input

	ToEscalationPolicyOutput() EscalationPolicyOutput
	ToEscalationPolicyOutputWithContext(ctx context.Context) EscalationPolicyOutput
}

func (*EscalationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**EscalationPolicy)(nil)).Elem()
}

func (i *EscalationPolicy) ToEscalationPolicyOutput() EscalationPolicyOutput {
	return i.ToEscalationPolicyOutputWithContext(context.Background())
}

func (i *EscalationPolicy) ToEscalationPolicyOutputWithContext(ctx context.Context) EscalationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EscalationPolicyOutput)
}

// EscalationPolicyArrayInput is an input type that accepts EscalationPolicyArray and EscalationPolicyArrayOutput values.
// You can construct a concrete instance of `EscalationPolicyArrayInput` via:
//
//	EscalationPolicyArray{ EscalationPolicyArgs{...} }
type EscalationPolicyArrayInput interface {
	pulumi.Input

	ToEscalationPolicyArrayOutput() EscalationPolicyArrayOutput
	ToEscalationPolicyArrayOutputWithContext(context.Context) EscalationPolicyArrayOutput
}

type EscalationPolicyArray []EscalationPolicyInput

func (EscalationPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EscalationPolicy)(nil)).Elem()
}

func (i EscalationPolicyArray) ToEscalationPolicyArrayOutput() EscalationPolicyArrayOutput {
	return i.ToEscalationPolicyArrayOutputWithContext(context.Background())
}

func (i EscalationPolicyArray) ToEscalationPolicyArrayOutputWithContext(ctx context.Context) EscalationPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EscalationPolicyArrayOutput)
}

// EscalationPolicyMapInput is an input type that accepts EscalationPolicyMap and EscalationPolicyMapOutput values.
// You can construct a concrete instance of `EscalationPolicyMapInput` via:
//
//	EscalationPolicyMap{ "key": EscalationPolicyArgs{...} }
type EscalationPolicyMapInput interface {
	pulumi.Input

	ToEscalationPolicyMapOutput() EscalationPolicyMapOutput
	ToEscalationPolicyMapOutputWithContext(context.Context) EscalationPolicyMapOutput
}

type EscalationPolicyMap map[string]EscalationPolicyInput

func (EscalationPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EscalationPolicy)(nil)).Elem()
}

func (i EscalationPolicyMap) ToEscalationPolicyMapOutput() EscalationPolicyMapOutput {
	return i.ToEscalationPolicyMapOutputWithContext(context.Background())
}

func (i EscalationPolicyMap) ToEscalationPolicyMapOutputWithContext(ctx context.Context) EscalationPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EscalationPolicyMapOutput)
}

type EscalationPolicyOutput struct{ *pulumi.OutputState }

func (EscalationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EscalationPolicy)(nil)).Elem()
}

func (o EscalationPolicyOutput) ToEscalationPolicyOutput() EscalationPolicyOutput {
	return o
}

func (o EscalationPolicyOutput) ToEscalationPolicyOutputWithContext(ctx context.Context) EscalationPolicyOutput {
	return o
}

// Detailed description about the Escalation Policy.
func (o EscalationPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EscalationPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Escalation policy owner.
func (o EscalationPolicyOutput) EntityOwner() EscalationPolicyEntityOwnerOutput {
	return o.ApplyT(func(v *EscalationPolicy) EscalationPolicyEntityOwnerOutput { return v.EntityOwner }).(EscalationPolicyEntityOwnerOutput)
}

// Name of the Escalation Policy.
func (o EscalationPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EscalationPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// You can choose to repeate the entire policy, if no one acknowledges the incident even after the Escalation Policy has been executed fully once
func (o EscalationPolicyOutput) Repeat() EscalationPolicyRepeatPtrOutput {
	return o.ApplyT(func(v *EscalationPolicy) EscalationPolicyRepeatPtrOutput { return v.Repeat }).(EscalationPolicyRepeatPtrOutput)
}

// Rules will have the details of who to notify and when to notify and how to notify them.
func (o EscalationPolicyOutput) Rules() EscalationPolicyRuleArrayOutput {
	return o.ApplyT(func(v *EscalationPolicy) EscalationPolicyRuleArrayOutput { return v.Rules }).(EscalationPolicyRuleArrayOutput)
}

// Team id.
func (o EscalationPolicyOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *EscalationPolicy) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type EscalationPolicyArrayOutput struct{ *pulumi.OutputState }

func (EscalationPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EscalationPolicy)(nil)).Elem()
}

func (o EscalationPolicyArrayOutput) ToEscalationPolicyArrayOutput() EscalationPolicyArrayOutput {
	return o
}

func (o EscalationPolicyArrayOutput) ToEscalationPolicyArrayOutputWithContext(ctx context.Context) EscalationPolicyArrayOutput {
	return o
}

func (o EscalationPolicyArrayOutput) Index(i pulumi.IntInput) EscalationPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EscalationPolicy {
		return vs[0].([]*EscalationPolicy)[vs[1].(int)]
	}).(EscalationPolicyOutput)
}

type EscalationPolicyMapOutput struct{ *pulumi.OutputState }

func (EscalationPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EscalationPolicy)(nil)).Elem()
}

func (o EscalationPolicyMapOutput) ToEscalationPolicyMapOutput() EscalationPolicyMapOutput {
	return o
}

func (o EscalationPolicyMapOutput) ToEscalationPolicyMapOutputWithContext(ctx context.Context) EscalationPolicyMapOutput {
	return o
}

func (o EscalationPolicyMapOutput) MapIndex(k pulumi.StringInput) EscalationPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EscalationPolicy {
		return vs[0].(map[string]*EscalationPolicy)[vs[1].(string)]
	}).(EscalationPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EscalationPolicyInput)(nil)).Elem(), &EscalationPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*EscalationPolicyArrayInput)(nil)).Elem(), EscalationPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EscalationPolicyMapInput)(nil)).Elem(), EscalationPolicyMap{})
	pulumi.RegisterOutputType(EscalationPolicyOutput{})
	pulumi.RegisterOutputType(EscalationPolicyArrayOutput{})
	pulumi.RegisterOutputType(EscalationPolicyMapOutput{})
}
