// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast
{
    public static class GetSquad
    {
        /// <summary>
        /// [Squads](https://support.squadcast.com/docs/squads) are smaller groups of members within Teams. Squads could correspond to groups of people that are responsible for specific projects within a Team.Use this data source to get information about a specific Squad.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Squadcast = Pulumi.Squadcast;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = Squadcast.GetSquad.Invoke(new()
        ///     {
        ///         Name = "test",
        ///         TeamId = "team id",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSquadResult> InvokeAsync(GetSquadArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSquadResult>("squadcast:index/getSquad:getSquad", args ?? new GetSquadArgs(), options.WithDefaults());

        /// <summary>
        /// [Squads](https://support.squadcast.com/docs/squads) are smaller groups of members within Teams. Squads could correspond to groups of people that are responsible for specific projects within a Team.Use this data source to get information about a specific Squad.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Squadcast = Pulumi.Squadcast;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var test = Squadcast.GetSquad.Invoke(new()
        ///     {
        ///         Name = "test",
        ///         TeamId = "team id",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSquadResult> Invoke(GetSquadInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSquadResult>("squadcast:index/getSquad:getSquad", args ?? new GetSquadInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSquadArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Squad.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Team id.
        /// </summary>
        [Input("teamId", required: true)]
        public string TeamId { get; set; } = null!;

        public GetSquadArgs()
        {
        }
        public static new GetSquadArgs Empty => new GetSquadArgs();
    }

    public sealed class GetSquadInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the Squad.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Team id.
        /// </summary>
        [Input("teamId", required: true)]
        public Input<string> TeamId { get; set; } = null!;

        public GetSquadInvokeArgs()
        {
        }
        public static new GetSquadInvokeArgs Empty => new GetSquadInvokeArgs();
    }


    [OutputType]
    public sealed class GetSquadResult
    {
        /// <summary>
        /// Squad id.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> MemberIds;
        /// <summary>
        /// Name of the Squad.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Team id.
        /// </summary>
        public readonly string TeamId;

        [OutputConstructor]
        private GetSquadResult(
            string id,

            ImmutableArray<string> memberIds,

            string name,

            string teamId)
        {
            Id = id;
            MemberIds = memberIds;
            Name = name;
            TeamId = teamId;
        }
    }
}