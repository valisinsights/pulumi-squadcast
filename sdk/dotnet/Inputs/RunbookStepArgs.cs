// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast.Inputs
{

    public sealed class RunbookStepArgs : global::Pulumi.ResourceArgs
    {
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        public RunbookStepArgs()
        {
        }
        public static new RunbookStepArgs Empty => new RunbookStepArgs();
    }
}
