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

// Use this resource to manage the Team roles and their permissions
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
//			_, err = squadcast.NewTeamRole(ctx, "exampleTeamRole", &squadcast.TeamRoleArgs{
//				TeamId: *pulumi.String(exampleTeam.Id),
//				Abilities: pulumi.StringArray{
//					pulumi.String("create-escalation-policies"),
//					pulumi.String("read-escalation-policies"),
//					pulumi.String("update-escalation-policies"),
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
// teamID:teamRole(exAdmin, User, Observer) Use 'Get All Teams' API to get the id of the team
//
// ```sh
//
//	$ pulumi import squadcast:index/teamRole:TeamRole example_resource_name "62d2fe23a57381088224d726:Admin"
//
// ```
type TeamRole struct {
	pulumi.CustomResourceState

	// abilities.
	// Current available abilities are :
	// create-escalation-policies, create-postmortems, create-runbooks, create-schedules, create-services, create-slos, create-squads, create-status-pages, delete-escalation-policies, delete-postmortems, delete-runbooks, delete-schedules, delete-services, delete-slos, delete-squads, delete-status-pages, read-escalation-policies, read-postmortems, read-runbooks, read-schedules, read-services, read-slos, read-squads, read-status-pages, read-team-analytics, update-escalation-policies, update-postmortems, update-runbooks, update-schedules, update-services, update-slos, update-squads, update-status-pages
	Abilities pulumi.StringArrayOutput `pulumi:"abilities"`
	// Team role default.
	Default pulumi.BoolOutput `pulumi:"default"`
	// Team role name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Team id.
	TeamId pulumi.StringOutput `pulumi:"teamId"`
}

// NewTeamRole registers a new resource with the given unique name, arguments, and options.
func NewTeamRole(ctx *pulumi.Context,
	name string, args *TeamRoleArgs, opts ...pulumi.ResourceOption) (*TeamRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Abilities == nil {
		return nil, errors.New("invalid value for required argument 'Abilities'")
	}
	if args.TeamId == nil {
		return nil, errors.New("invalid value for required argument 'TeamId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TeamRole
	err := ctx.RegisterResource("squadcast:index/teamRole:TeamRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeamRole gets an existing TeamRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeamRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamRoleState, opts ...pulumi.ResourceOption) (*TeamRole, error) {
	var resource TeamRole
	err := ctx.ReadResource("squadcast:index/teamRole:TeamRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TeamRole resources.
type teamRoleState struct {
	// abilities.
	// Current available abilities are :
	// create-escalation-policies, create-postmortems, create-runbooks, create-schedules, create-services, create-slos, create-squads, create-status-pages, delete-escalation-policies, delete-postmortems, delete-runbooks, delete-schedules, delete-services, delete-slos, delete-squads, delete-status-pages, read-escalation-policies, read-postmortems, read-runbooks, read-schedules, read-services, read-slos, read-squads, read-status-pages, read-team-analytics, update-escalation-policies, update-postmortems, update-runbooks, update-schedules, update-services, update-slos, update-squads, update-status-pages
	Abilities []string `pulumi:"abilities"`
	// Team role default.
	Default *bool `pulumi:"default"`
	// Team role name.
	Name *string `pulumi:"name"`
	// Team id.
	TeamId *string `pulumi:"teamId"`
}

type TeamRoleState struct {
	// abilities.
	// Current available abilities are :
	// create-escalation-policies, create-postmortems, create-runbooks, create-schedules, create-services, create-slos, create-squads, create-status-pages, delete-escalation-policies, delete-postmortems, delete-runbooks, delete-schedules, delete-services, delete-slos, delete-squads, delete-status-pages, read-escalation-policies, read-postmortems, read-runbooks, read-schedules, read-services, read-slos, read-squads, read-status-pages, read-team-analytics, update-escalation-policies, update-postmortems, update-runbooks, update-schedules, update-services, update-slos, update-squads, update-status-pages
	Abilities pulumi.StringArrayInput
	// Team role default.
	Default pulumi.BoolPtrInput
	// Team role name.
	Name pulumi.StringPtrInput
	// Team id.
	TeamId pulumi.StringPtrInput
}

func (TeamRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamRoleState)(nil)).Elem()
}

type teamRoleArgs struct {
	// abilities.
	// Current available abilities are :
	// create-escalation-policies, create-postmortems, create-runbooks, create-schedules, create-services, create-slos, create-squads, create-status-pages, delete-escalation-policies, delete-postmortems, delete-runbooks, delete-schedules, delete-services, delete-slos, delete-squads, delete-status-pages, read-escalation-policies, read-postmortems, read-runbooks, read-schedules, read-services, read-slos, read-squads, read-status-pages, read-team-analytics, update-escalation-policies, update-postmortems, update-runbooks, update-schedules, update-services, update-slos, update-squads, update-status-pages
	Abilities []string `pulumi:"abilities"`
	// Team role name.
	Name *string `pulumi:"name"`
	// Team id.
	TeamId string `pulumi:"teamId"`
}

// The set of arguments for constructing a TeamRole resource.
type TeamRoleArgs struct {
	// abilities.
	// Current available abilities are :
	// create-escalation-policies, create-postmortems, create-runbooks, create-schedules, create-services, create-slos, create-squads, create-status-pages, delete-escalation-policies, delete-postmortems, delete-runbooks, delete-schedules, delete-services, delete-slos, delete-squads, delete-status-pages, read-escalation-policies, read-postmortems, read-runbooks, read-schedules, read-services, read-slos, read-squads, read-status-pages, read-team-analytics, update-escalation-policies, update-postmortems, update-runbooks, update-schedules, update-services, update-slos, update-squads, update-status-pages
	Abilities pulumi.StringArrayInput
	// Team role name.
	Name pulumi.StringPtrInput
	// Team id.
	TeamId pulumi.StringInput
}

func (TeamRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamRoleArgs)(nil)).Elem()
}

type TeamRoleInput interface {
	pulumi.Input

	ToTeamRoleOutput() TeamRoleOutput
	ToTeamRoleOutputWithContext(ctx context.Context) TeamRoleOutput
}

func (*TeamRole) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamRole)(nil)).Elem()
}

func (i *TeamRole) ToTeamRoleOutput() TeamRoleOutput {
	return i.ToTeamRoleOutputWithContext(context.Background())
}

func (i *TeamRole) ToTeamRoleOutputWithContext(ctx context.Context) TeamRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamRoleOutput)
}

// TeamRoleArrayInput is an input type that accepts TeamRoleArray and TeamRoleArrayOutput values.
// You can construct a concrete instance of `TeamRoleArrayInput` via:
//
//	TeamRoleArray{ TeamRoleArgs{...} }
type TeamRoleArrayInput interface {
	pulumi.Input

	ToTeamRoleArrayOutput() TeamRoleArrayOutput
	ToTeamRoleArrayOutputWithContext(context.Context) TeamRoleArrayOutput
}

type TeamRoleArray []TeamRoleInput

func (TeamRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamRole)(nil)).Elem()
}

func (i TeamRoleArray) ToTeamRoleArrayOutput() TeamRoleArrayOutput {
	return i.ToTeamRoleArrayOutputWithContext(context.Background())
}

func (i TeamRoleArray) ToTeamRoleArrayOutputWithContext(ctx context.Context) TeamRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamRoleArrayOutput)
}

// TeamRoleMapInput is an input type that accepts TeamRoleMap and TeamRoleMapOutput values.
// You can construct a concrete instance of `TeamRoleMapInput` via:
//
//	TeamRoleMap{ "key": TeamRoleArgs{...} }
type TeamRoleMapInput interface {
	pulumi.Input

	ToTeamRoleMapOutput() TeamRoleMapOutput
	ToTeamRoleMapOutputWithContext(context.Context) TeamRoleMapOutput
}

type TeamRoleMap map[string]TeamRoleInput

func (TeamRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamRole)(nil)).Elem()
}

func (i TeamRoleMap) ToTeamRoleMapOutput() TeamRoleMapOutput {
	return i.ToTeamRoleMapOutputWithContext(context.Background())
}

func (i TeamRoleMap) ToTeamRoleMapOutputWithContext(ctx context.Context) TeamRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamRoleMapOutput)
}

type TeamRoleOutput struct{ *pulumi.OutputState }

func (TeamRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamRole)(nil)).Elem()
}

func (o TeamRoleOutput) ToTeamRoleOutput() TeamRoleOutput {
	return o
}

func (o TeamRoleOutput) ToTeamRoleOutputWithContext(ctx context.Context) TeamRoleOutput {
	return o
}

// abilities.
// Current available abilities are :
// create-escalation-policies, create-postmortems, create-runbooks, create-schedules, create-services, create-slos, create-squads, create-status-pages, delete-escalation-policies, delete-postmortems, delete-runbooks, delete-schedules, delete-services, delete-slos, delete-squads, delete-status-pages, read-escalation-policies, read-postmortems, read-runbooks, read-schedules, read-services, read-slos, read-squads, read-status-pages, read-team-analytics, update-escalation-policies, update-postmortems, update-runbooks, update-schedules, update-services, update-slos, update-squads, update-status-pages
func (o TeamRoleOutput) Abilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TeamRole) pulumi.StringArrayOutput { return v.Abilities }).(pulumi.StringArrayOutput)
}

// Team role default.
func (o TeamRoleOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v *TeamRole) pulumi.BoolOutput { return v.Default }).(pulumi.BoolOutput)
}

// Team role name.
func (o TeamRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Team id.
func (o TeamRoleOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v *TeamRole) pulumi.StringOutput { return v.TeamId }).(pulumi.StringOutput)
}

type TeamRoleArrayOutput struct{ *pulumi.OutputState }

func (TeamRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TeamRole)(nil)).Elem()
}

func (o TeamRoleArrayOutput) ToTeamRoleArrayOutput() TeamRoleArrayOutput {
	return o
}

func (o TeamRoleArrayOutput) ToTeamRoleArrayOutputWithContext(ctx context.Context) TeamRoleArrayOutput {
	return o
}

func (o TeamRoleArrayOutput) Index(i pulumi.IntInput) TeamRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TeamRole {
		return vs[0].([]*TeamRole)[vs[1].(int)]
	}).(TeamRoleOutput)
}

type TeamRoleMapOutput struct{ *pulumi.OutputState }

func (TeamRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TeamRole)(nil)).Elem()
}

func (o TeamRoleMapOutput) ToTeamRoleMapOutput() TeamRoleMapOutput {
	return o
}

func (o TeamRoleMapOutput) ToTeamRoleMapOutputWithContext(ctx context.Context) TeamRoleMapOutput {
	return o
}

func (o TeamRoleMapOutput) MapIndex(k pulumi.StringInput) TeamRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TeamRole {
		return vs[0].(map[string]*TeamRole)[vs[1].(string)]
	}).(TeamRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamRoleInput)(nil)).Elem(), &TeamRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamRoleArrayInput)(nil)).Elem(), TeamRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamRoleMapInput)(nil)).Elem(), TeamRoleMap{})
	pulumi.RegisterOutputType(TeamRoleOutput{})
	pulumi.RegisterOutputType(TeamRoleArrayOutput{})
	pulumi.RegisterOutputType(TeamRoleMapOutput{})
}