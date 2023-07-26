// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast.Inputs
{

    public sealed class SLORuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the monitoring rule
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Is checked?
        /// </summary>
        [Input("isChecked")]
        public Input<bool>? IsChecked { get; set; }

        /// <summary>
        /// The name of monitoring check."Supported values are "breached*error*budget", "unhealthy*slo","increased*false*positives", "remaining*error_budget"
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the SLO
        /// </summary>
        [Input("sloId")]
        public Input<int>? SloId { get; set; }

        /// <summary>
        /// Threshold for the monitoring checkOnly supported for rules name "increased*false*positives" and "remaining*error*budget"
        /// </summary>
        [Input("threshold")]
        public Input<int>? Threshold { get; set; }

        public SLORuleArgs()
        {
        }
        public static new SLORuleArgs Empty => new SLORuleArgs();
    }
}
