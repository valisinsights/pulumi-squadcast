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
    public sealed class GetWebformSeverityResult
    {
        /// <summary>
        /// Severity description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Severity type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetWebformSeverityResult(
            string description,

            string type)
        {
            Description = description;
            Type = type;
        }
    }
}
