// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package squadcast

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/valisinsights/pulumi-squadcast/sdk/go/squadcast/internal"
)

// Use this resource to manage the Team meta details like Name, descripton etc.
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
//			_, err := squadcast.NewTeam(ctx, "exampleTeam", nil)
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
// #teamName
//
// ```sh
//
//	$ pulumi import squadcast:index/team:Team test "My Team Name"
//
// ```
type Team struct {
	pulumi.CustomResourceState

	// Team is default?.
	Default        pulumi.BoolOutput      `pulumi:"default"`
	DefaultRoleIds pulumi.StringMapOutput `pulumi:"defaultRoleIds"`
	// Team description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Team name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewTeam registers a new resource with the given unique name, arguments, and options.
func NewTeam(ctx *pulumi.Context,
	name string, args *TeamArgs, opts ...pulumi.ResourceOption) (*Team, error) {
	if args == nil {
		args = &TeamArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Team
	err := ctx.RegisterResource("squadcast:index/team:Team", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeam gets an existing Team resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamState, opts ...pulumi.ResourceOption) (*Team, error) {
	var resource Team
	err := ctx.ReadResource("squadcast:index/team:Team", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Team resources.
type teamState struct {
	// Team is default?.
	Default        *bool             `pulumi:"default"`
	DefaultRoleIds map[string]string `pulumi:"defaultRoleIds"`
	// Team description.
	Description *string `pulumi:"description"`
	// Team name.
	Name *string `pulumi:"name"`
}

type TeamState struct {
	// Team is default?.
	Default        pulumi.BoolPtrInput
	DefaultRoleIds pulumi.StringMapInput
	// Team description.
	Description pulumi.StringPtrInput
	// Team name.
	Name pulumi.StringPtrInput
}

func (TeamState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamState)(nil)).Elem()
}

type teamArgs struct {
	// Team description.
	Description *string `pulumi:"description"`
	// Team name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Team resource.
type TeamArgs struct {
	// Team description.
	Description pulumi.StringPtrInput
	// Team name.
	Name pulumi.StringPtrInput
}

func (TeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamArgs)(nil)).Elem()
}

type TeamInput interface {
	pulumi.Input

	ToTeamOutput() TeamOutput
	ToTeamOutputWithContext(ctx context.Context) TeamOutput
}

func (*Team) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (i *Team) ToTeamOutput() TeamOutput {
	return i.ToTeamOutputWithContext(context.Background())
}

func (i *Team) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamOutput)
}

// TeamArrayInput is an input type that accepts TeamArray and TeamArrayOutput values.
// You can construct a concrete instance of `TeamArrayInput` via:
//
//	TeamArray{ TeamArgs{...} }
type TeamArrayInput interface {
	pulumi.Input

	ToTeamArrayOutput() TeamArrayOutput
	ToTeamArrayOutputWithContext(context.Context) TeamArrayOutput
}

type TeamArray []TeamInput

func (TeamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (i TeamArray) ToTeamArrayOutput() TeamArrayOutput {
	return i.ToTeamArrayOutputWithContext(context.Background())
}

func (i TeamArray) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamArrayOutput)
}

// TeamMapInput is an input type that accepts TeamMap and TeamMapOutput values.
// You can construct a concrete instance of `TeamMapInput` via:
//
//	TeamMap{ "key": TeamArgs{...} }
type TeamMapInput interface {
	pulumi.Input

	ToTeamMapOutput() TeamMapOutput
	ToTeamMapOutputWithContext(context.Context) TeamMapOutput
}

type TeamMap map[string]TeamInput

func (TeamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (i TeamMap) ToTeamMapOutput() TeamMapOutput {
	return i.ToTeamMapOutputWithContext(context.Background())
}

func (i TeamMap) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMapOutput)
}

type TeamOutput struct{ *pulumi.OutputState }

func (TeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (o TeamOutput) ToTeamOutput() TeamOutput {
	return o
}

func (o TeamOutput) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return o
}

// Team is default?.
func (o TeamOutput) Default() pulumi.BoolOutput {
	return o.ApplyT(func(v *Team) pulumi.BoolOutput { return v.Default }).(pulumi.BoolOutput)
}

func (o TeamOutput) DefaultRoleIds() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Team) pulumi.StringMapOutput { return v.DefaultRoleIds }).(pulumi.StringMapOutput)
}

// Team description.
func (o TeamOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Team name.
func (o TeamOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type TeamArrayOutput struct{ *pulumi.OutputState }

func (TeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (o TeamArrayOutput) ToTeamArrayOutput() TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) Index(i pulumi.IntInput) TeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Team {
		return vs[0].([]*Team)[vs[1].(int)]
	}).(TeamOutput)
}

type TeamMapOutput struct{ *pulumi.OutputState }

func (TeamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (o TeamMapOutput) ToTeamMapOutput() TeamMapOutput {
	return o
}

func (o TeamMapOutput) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return o
}

func (o TeamMapOutput) MapIndex(k pulumi.StringInput) TeamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Team {
		return vs[0].(map[string]*Team)[vs[1].(string)]
	}).(TeamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamInput)(nil)).Elem(), &Team{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamArrayInput)(nil)).Elem(), TeamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMapInput)(nil)).Elem(), TeamMap{})
	pulumi.RegisterOutputType(TeamOutput{})
	pulumi.RegisterOutputType(TeamArrayOutput{})
	pulumi.RegisterOutputType(TeamMapOutput{})
}
