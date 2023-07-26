// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast.Inputs
{

    public sealed class SuppressionRulesRuleGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("basicExpressions")]
        private InputList<Inputs.SuppressionRulesRuleBasicExpressionGetArgs>? _basicExpressions;

        /// <summary>
        /// The basic expression which needs to be evaluated to be true for this rule to apply.
        /// </summary>
        public InputList<Inputs.SuppressionRulesRuleBasicExpressionGetArgs> BasicExpressions
        {
            get => _basicExpressions ?? (_basicExpressions = new InputList<Inputs.SuppressionRulesRuleBasicExpressionGetArgs>());
            set => _basicExpressions = value;
        }

        /// <summary>
        /// description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The expression which needs to be evaluated to be true for this rule to apply.
        /// </summary>
        [Input("expression")]
        public Input<string>? Expression { get; set; }

        /// <summary>
        /// is_basic will be true when users use the drop down selectors which will have lhs, op &amp; rhs value, whereas it will be false when they use the advanced mode and it would have the expression for it's value
        /// </summary>
        [Input("isBasic", required: true)]
        public Input<bool> IsBasic { get; set; } = null!;

        /// <summary>
        /// is_timebased will be true when users use the time based suppression rule
        /// </summary>
        [Input("isTimebased")]
        public Input<bool>? IsTimebased { get; set; }

        [Input("timeslots")]
        private InputList<Inputs.SuppressionRulesRuleTimeslotGetArgs>? _timeslots;

        /// <summary>
        /// The timeslots for which this rule should be applied.
        /// </summary>
        public InputList<Inputs.SuppressionRulesRuleTimeslotGetArgs> Timeslots
        {
            get => _timeslots ?? (_timeslots = new InputList<Inputs.SuppressionRulesRuleTimeslotGetArgs>());
            set => _timeslots = value;
        }

        public SuppressionRulesRuleGetArgs()
        {
        }
        public static new SuppressionRulesRuleGetArgs Empty => new SuppressionRulesRuleGetArgs();
    }
}
