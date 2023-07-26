// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast.Outputs
{

    [OutputType]
    public sealed class DeduplicationRulesRule
    {
        /// <summary>
        /// The basic expression which needs to be evaluated to be true for this rule to apply.
        /// </summary>
        public readonly ImmutableArray<Outputs.DeduplicationRulesRuleBasicExpression> BasicExpressions;
        /// <summary>
        /// Denotes if dependent services should also be deduplicated
        /// </summary>
        public readonly bool? DependencyDeduplication;
        /// <summary>
        /// description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The expression which needs to be evaluated to be true for this rule to apply.
        /// </summary>
        public readonly string? Expression;
        /// <summary>
        /// is_basic will be true when users use the drop down selectors which will have lhs, op &amp; rhs value, whereas it will be false when they use the advanced mode and it would have the expression for it's value
        /// </summary>
        public readonly bool IsBasic;
        /// <summary>
        /// time unit (mins or hours)
        /// </summary>
        public readonly string? TimeUnit;
        /// <summary>
        /// integer for time_unit
        /// </summary>
        public readonly int? TimeWindow;

        [OutputConstructor]
        private DeduplicationRulesRule(
            ImmutableArray<Outputs.DeduplicationRulesRuleBasicExpression> basicExpressions,

            bool? dependencyDeduplication,

            string? description,

            string? expression,

            bool isBasic,

            string? timeUnit,

            int? timeWindow)
        {
            BasicExpressions = basicExpressions;
            DependencyDeduplication = dependencyDeduplication;
            Description = description;
            Expression = expression;
            IsBasic = isBasic;
            TimeUnit = timeUnit;
            TimeWindow = timeWindow;
        }
    }
}
