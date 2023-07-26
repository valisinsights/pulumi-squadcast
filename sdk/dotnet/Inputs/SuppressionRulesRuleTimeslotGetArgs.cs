// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast.Inputs
{

    public sealed class SuppressionRulesRuleTimeslotGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("customs")]
        private InputList<Inputs.SuppressionRulesRuleTimeslotCustomGetArgs>? _customs;
        public InputList<Inputs.SuppressionRulesRuleTimeslotCustomGetArgs> Customs
        {
            get => _customs ?? (_customs = new InputList<Inputs.SuppressionRulesRuleTimeslotCustomGetArgs>());
            set => _customs = value;
        }

        [Input("endTime", required: true)]
        public Input<string> EndTime { get; set; } = null!;

        [Input("endsNever")]
        public Input<bool>? EndsNever { get; set; }

        [Input("endsOn", required: true)]
        public Input<string> EndsOn { get; set; } = null!;

        [Input("isAllday")]
        public Input<bool>? IsAllday { get; set; }

        [Input("isCustom")]
        public Input<bool>? IsCustom { get; set; }

        [Input("repetition", required: true)]
        public Input<string> Repetition { get; set; } = null!;

        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        [Input("timeZone", required: true)]
        public Input<string> TimeZone { get; set; } = null!;

        public SuppressionRulesRuleTimeslotGetArgs()
        {
        }
        public static new SuppressionRulesRuleTimeslotGetArgs Empty => new SuppressionRulesRuleTimeslotGetArgs();
    }
}
