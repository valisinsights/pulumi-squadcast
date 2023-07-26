// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package squadcast

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/valisinsights/pulumi-squadcast/sdk/go/squadcast/internal"
)

// [Squadcast schedules](https://support.squadcast.com/docs/schedules) are used to manage on-call scheduling & determine who will be notified when an incident is triggered. Use this data source to get information about a specific schedule that you can use for other Squadcast resources.
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
//			_, err := squadcast.LookupSchedule(ctx, &squadcast.LookupScheduleArgs{
//				Name:   squadcast_schedule.Test.Name,
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
func LookupSchedule(ctx *pulumi.Context, args *LookupScheduleArgs, opts ...pulumi.InvokeOption) (*LookupScheduleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupScheduleResult
	err := ctx.Invoke("squadcast:index/getSchedule:getSchedule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSchedule.
type LookupScheduleArgs struct {
	// Name of the Schedule.
	Name string `pulumi:"name"`
	// Team id.
	TeamId string `pulumi:"teamId"`
}

// A collection of values returned by getSchedule.
type LookupScheduleResult struct {
	// Calendar color scheme for this schedule, hex values.
	Color string `pulumi:"color"`
	// Detailed description about the schedule.
	Description string `pulumi:"description"`
	// Schedule id.
	Id string `pulumi:"id"`
	// Name of the Schedule.
	Name string `pulumi:"name"`
	// Team id.
	TeamId string `pulumi:"teamId"`
}

func LookupScheduleOutput(ctx *pulumi.Context, args LookupScheduleOutputArgs, opts ...pulumi.InvokeOption) LookupScheduleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduleResult, error) {
			args := v.(LookupScheduleArgs)
			r, err := LookupSchedule(ctx, &args, opts...)
			var s LookupScheduleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduleResultOutput)
}

// A collection of arguments for invoking getSchedule.
type LookupScheduleOutputArgs struct {
	// Name of the Schedule.
	Name pulumi.StringInput `pulumi:"name"`
	// Team id.
	TeamId pulumi.StringInput `pulumi:"teamId"`
}

func (LookupScheduleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleArgs)(nil)).Elem()
}

// A collection of values returned by getSchedule.
type LookupScheduleResultOutput struct{ *pulumi.OutputState }

func (LookupScheduleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduleResult)(nil)).Elem()
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutput() LookupScheduleResultOutput {
	return o
}

func (o LookupScheduleResultOutput) ToLookupScheduleResultOutputWithContext(ctx context.Context) LookupScheduleResultOutput {
	return o
}

// Calendar color scheme for this schedule, hex values.
func (o LookupScheduleResultOutput) Color() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Color }).(pulumi.StringOutput)
}

// Detailed description about the schedule.
func (o LookupScheduleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Description }).(pulumi.StringOutput)
}

// Schedule id.
func (o LookupScheduleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the Schedule.
func (o LookupScheduleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.Name }).(pulumi.StringOutput)
}

// Team id.
func (o LookupScheduleResultOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduleResult) string { return v.TeamId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduleResultOutput{})
}