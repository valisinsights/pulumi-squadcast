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
    public sealed class GetEscalationPolicyRuleRoundRobinRotationResult
    {
        public readonly int DelayMinutes;
        public readonly bool Enabled;

        [OutputConstructor]
        private GetEscalationPolicyRuleRoundRobinRotationResult(
            int delayMinutes,

            bool enabled)
        {
            DelayMinutes = delayMinutes;
            Enabled = enabled;
        }
    }
}