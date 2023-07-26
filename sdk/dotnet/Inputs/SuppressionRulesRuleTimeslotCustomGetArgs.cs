// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast.Inputs
{

    public sealed class SuppressionRulesRuleTimeslotCustomGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("repeats", required: true)]
        public Input<string> Repeats { get; set; } = null!;

        [Input("repeatsCount")]
        public Input<int>? RepeatsCount { get; set; }

        [Input("repeatsOnMonth")]
        public Input<string>? RepeatsOnMonth { get; set; }

        [Input("repeatsOnWeekdays")]
        private InputList<int>? _repeatsOnWeekdays;
        public InputList<int> RepeatsOnWeekdays
        {
            get => _repeatsOnWeekdays ?? (_repeatsOnWeekdays = new InputList<int>());
            set => _repeatsOnWeekdays = value;
        }

        public SuppressionRulesRuleTimeslotCustomGetArgs()
        {
        }
        public static new SuppressionRulesRuleTimeslotCustomGetArgs Empty => new SuppressionRulesRuleTimeslotCustomGetArgs();
    }
}