// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast.Inputs
{

    public sealed class EscalationPolicyRuleTargetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// EscalationPolicy id.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public EscalationPolicyRuleTargetGetArgs()
        {
        }
        public static new EscalationPolicyRuleTargetGetArgs Empty => new EscalationPolicyRuleTargetGetArgs();
    }
}