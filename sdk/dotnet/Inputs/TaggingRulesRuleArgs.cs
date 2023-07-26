// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast.Inputs
{

    public sealed class TaggingRulesRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("basicExpressions")]
        private InputList<Inputs.TaggingRulesRuleBasicExpressionArgs>? _basicExpressions;

        /// <summary>
        /// The basic expression which needs to be evaluated to be true for this rule to apply.
        /// </summary>
        public InputList<Inputs.TaggingRulesRuleBasicExpressionArgs> BasicExpressions
        {
            get => _basicExpressions ?? (_basicExpressions = new InputList<Inputs.TaggingRulesRuleBasicExpressionArgs>());
            set => _basicExpressions = value;
        }

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

        [Input("tags")]
        private InputList<Inputs.TaggingRulesRuleTagArgs>? _tags;

        /// <summary>
        /// The tags supposed to be set for a given payload(incident), Expression must be set when tags are empty and must contain addTags parameters.
        /// </summary>
        public InputList<Inputs.TaggingRulesRuleTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TaggingRulesRuleTagArgs>());
            set => _tags = value;
        }

        public TaggingRulesRuleArgs()
        {
        }
        public static new TaggingRulesRuleArgs Empty => new TaggingRulesRuleArgs();
    }
}
