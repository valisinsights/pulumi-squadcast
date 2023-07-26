// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast.Inputs
{

    public sealed class SLONotifyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the notification rule
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// The ID of the service in which the user want to create an incident
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        /// <summary>
        /// The ID of the SLO.
        /// </summary>
        [Input("sloId")]
        public Input<int>? SloId { get; set; }

        [Input("squadIds")]
        private InputList<string>? _squadIds;

        /// <summary>
        /// List of Squad ID's who should be alerted via email.
        /// </summary>
        public InputList<string> SquadIds
        {
            get => _squadIds ?? (_squadIds = new InputList<string>());
            set => _squadIds = value;
        }

        [Input("userIds")]
        private InputList<string>? _userIds;

        /// <summary>
        /// List of user ID's who should be alerted via email.
        /// </summary>
        public InputList<string> UserIds
        {
            get => _userIds ?? (_userIds = new InputList<string>());
            set => _userIds = value;
        }

        public SLONotifyArgs()
        {
        }
        public static new SLONotifyArgs Empty => new SLONotifyArgs();
    }
}
