// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast.Inputs
{

    public sealed class DeduplicationRulesRuleBasicExpressionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("lhs", required: true)]
        public Input<string> Lhs { get; set; } = null!;

        [Input("op", required: true)]
        public Input<string> Op { get; set; } = null!;

        [Input("rhs", required: true)]
        public Input<string> Rhs { get; set; } = null!;

        public DeduplicationRulesRuleBasicExpressionGetArgs()
        {
        }
        public static new DeduplicationRulesRuleBasicExpressionGetArgs Empty => new DeduplicationRulesRuleBasicExpressionGetArgs();
    }
}
