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
    public sealed class TaggingRulesRuleTag
    {
        public readonly string Color;
        public readonly string Key;
        public readonly string Value;

        [OutputConstructor]
        private TaggingRulesRuleTag(
            string color,

            string key,

            string value)
        {
            Color = color;
            Key = key;
            Value = value;
        }
    }
}
