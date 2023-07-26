// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package squadcast

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/valisinsights/pulumi-squadcast/sdk/go/squadcast/internal"
)

// A Runbook is a compilation of routine procedures and operations that are documented for reference while working on a critical incident. Sometimes, it can also be referred to as a Playbook.Use this data source to get information about a specific Runbook that you can use for other Squadcast resources.
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
//			_, err := squadcast.LookupRunbook(ctx, &squadcast.LookupRunbookArgs{
//				Name:   squadcast_runbook.Test.Name,
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
func LookupRunbook(ctx *pulumi.Context, args *LookupRunbookArgs, opts ...pulumi.InvokeOption) (*LookupRunbookResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRunbookResult
	err := ctx.Invoke("squadcast:index/getRunbook:getRunbook", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRunbook.
type LookupRunbookArgs struct {
	// Name of the Runbook
	Name string `pulumi:"name"`
	// Team id.
	TeamId string `pulumi:"teamId"`
}

// A collection of values returned by getRunbook.
type LookupRunbookResult struct {
	// Runbooks owner
	EntityOwners []GetRunbookEntityOwner `pulumi:"entityOwners"`
	// Runbook id.
	Id string `pulumi:"id"`
	// Name of the Runbook
	Name string `pulumi:"name"`
	// Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
	Steps []GetRunbookStep `pulumi:"steps"`
	// Team id.
	TeamId string `pulumi:"teamId"`
}

func LookupRunbookOutput(ctx *pulumi.Context, args LookupRunbookOutputArgs, opts ...pulumi.InvokeOption) LookupRunbookResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRunbookResult, error) {
			args := v.(LookupRunbookArgs)
			r, err := LookupRunbook(ctx, &args, opts...)
			var s LookupRunbookResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRunbookResultOutput)
}

// A collection of arguments for invoking getRunbook.
type LookupRunbookOutputArgs struct {
	// Name of the Runbook
	Name pulumi.StringInput `pulumi:"name"`
	// Team id.
	TeamId pulumi.StringInput `pulumi:"teamId"`
}

func (LookupRunbookOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRunbookArgs)(nil)).Elem()
}

// A collection of values returned by getRunbook.
type LookupRunbookResultOutput struct{ *pulumi.OutputState }

func (LookupRunbookResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRunbookResult)(nil)).Elem()
}

func (o LookupRunbookResultOutput) ToLookupRunbookResultOutput() LookupRunbookResultOutput {
	return o
}

func (o LookupRunbookResultOutput) ToLookupRunbookResultOutputWithContext(ctx context.Context) LookupRunbookResultOutput {
	return o
}

// Runbooks owner
func (o LookupRunbookResultOutput) EntityOwners() GetRunbookEntityOwnerArrayOutput {
	return o.ApplyT(func(v LookupRunbookResult) []GetRunbookEntityOwner { return v.EntityOwners }).(GetRunbookEntityOwnerArrayOutput)
}

// Runbook id.
func (o LookupRunbookResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRunbookResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the Runbook
func (o LookupRunbookResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRunbookResult) string { return v.Name }).(pulumi.StringOutput)
}

// Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
func (o LookupRunbookResultOutput) Steps() GetRunbookStepArrayOutput {
	return o.ApplyT(func(v LookupRunbookResult) []GetRunbookStep { return v.Steps }).(GetRunbookStepArrayOutput)
}

// Team id.
func (o LookupRunbookResultOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRunbookResult) string { return v.TeamId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRunbookResultOutput{})
}
