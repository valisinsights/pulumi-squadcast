// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast.Inputs
{

    public sealed class EscalationPolicyRepeatArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of minutes to wait before repeating the escalation policy
        /// </summary>
        [Input("delayMinutes", required: true)]
        public Input<int> DelayMinutes { get; set; } = null!;

        /// <summary>
        /// The number of times you want this escalation policy to be repeated, maximum allowed to repeat 3 times
        /// </summary>
        [Input("times", required: true)]
        public Input<int> Times { get; set; } = null!;

        public EscalationPolicyRepeatArgs()
        {
        }
        public static new EscalationPolicyRepeatArgs Empty => new EscalationPolicyRepeatArgs();
    }
}