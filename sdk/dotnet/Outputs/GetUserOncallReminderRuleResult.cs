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
    public sealed class GetUserOncallReminderRuleResult
    {
        public readonly int DelayMinutes;
        public readonly string Type;

        [OutputConstructor]
        private GetUserOncallReminderRuleResult(
            int delayMinutes,

            string type)
        {
            DelayMinutes = delayMinutes;
            Type = type;
        }
    }
}
