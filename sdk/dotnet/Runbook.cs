// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast
{
    /// <summary>
    /// [Squadcast Runbook](https://support.squadcast.com/docs/runbooks) is a compilation of routine procedures and operations that are documented for reference while working on a critical incident. Sometimes, it can also be referred to as a Playbook.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Squadcast = Pulumi.Squadcast;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleTeam = Squadcast.GetTeam.Invoke(new()
    ///     {
    ///         Name = "example team name",
    ///     });
    /// 
    ///     var exampleUser = Squadcast.GetUser.Invoke(new()
    ///     {
    ///         Email = "test@example.com",
    ///     });
    /// 
    ///     var exampleRunbook = new Squadcast.Runbook("exampleRunbook", new()
    ///     {
    ///         TeamId = exampleTeam.Apply(getTeamResult =&gt; getTeamResult.Id),
    ///         Steps = new[]
    ///         {
    ///             new Squadcast.Inputs.RunbookStepArgs
    ///             {
    ///                 Content = "some text here",
    ///             },
    ///             new Squadcast.Inputs.RunbookStepArgs
    ///             {
    ///                 Content = "some text here 2",
    ///             },
    ///         },
    ///         EntityOwner = new Squadcast.Inputs.RunbookEntityOwnerArgs
    ///         {
    ///             Id = exampleUser.Apply(getUserResult =&gt; getUserResult.Id),
    ///             Type = "user",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// teamID:runbookID Use 'Get All Teams' and 'Get All Runbooks' APIs to get the id of the team and runbook respectively
    /// 
    /// ```sh
    ///  $ pulumi import squadcast:index/runbook:Runbook test 62d2fe23a57381088224d726:62da76c088f407f9ca756ca5
    /// ```
    /// </summary>
    [SquadcastResourceType("squadcast:index/runbook:Runbook")]
    public partial class Runbook : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Runbooks owner.
        /// </summary>
        [Output("entityOwner")]
        public Output<Outputs.RunbookEntityOwner> EntityOwner { get; private set; } = null!;

        /// <summary>
        /// Name of the Runbook.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
        /// </summary>
        [Output("steps")]
        public Output<ImmutableArray<Outputs.RunbookStep>> Steps { get; private set; } = null!;

        /// <summary>
        /// Team id.
        /// </summary>
        [Output("teamId")]
        public Output<string> TeamId { get; private set; } = null!;


        /// <summary>
        /// Create a Runbook resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Runbook(string name, RunbookArgs args, CustomResourceOptions? options = null)
            : base("squadcast:index/runbook:Runbook", name, args ?? new RunbookArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Runbook(string name, Input<string> id, RunbookState? state = null, CustomResourceOptions? options = null)
            : base("squadcast:index/runbook:Runbook", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Runbook resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Runbook Get(string name, Input<string> id, RunbookState? state = null, CustomResourceOptions? options = null)
        {
            return new Runbook(name, id, state, options);
        }
    }

    public sealed class RunbookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Runbooks owner.
        /// </summary>
        [Input("entityOwner")]
        public Input<Inputs.RunbookEntityOwnerArgs>? EntityOwner { get; set; }

        /// <summary>
        /// Name of the Runbook.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("steps", required: true)]
        private InputList<Inputs.RunbookStepArgs>? _steps;

        /// <summary>
        /// Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
        /// </summary>
        public InputList<Inputs.RunbookStepArgs> Steps
        {
            get => _steps ?? (_steps = new InputList<Inputs.RunbookStepArgs>());
            set => _steps = value;
        }

        /// <summary>
        /// Team id.
        /// </summary>
        [Input("teamId", required: true)]
        public Input<string> TeamId { get; set; } = null!;

        public RunbookArgs()
        {
        }
        public static new RunbookArgs Empty => new RunbookArgs();
    }

    public sealed class RunbookState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Runbooks owner.
        /// </summary>
        [Input("entityOwner")]
        public Input<Inputs.RunbookEntityOwnerGetArgs>? EntityOwner { get; set; }

        /// <summary>
        /// Name of the Runbook.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("steps")]
        private InputList<Inputs.RunbookStepGetArgs>? _steps;

        /// <summary>
        /// Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
        /// </summary>
        public InputList<Inputs.RunbookStepGetArgs> Steps
        {
            get => _steps ?? (_steps = new InputList<Inputs.RunbookStepGetArgs>());
            set => _steps = value;
        }

        /// <summary>
        /// Team id.
        /// </summary>
        [Input("teamId")]
        public Input<string>? TeamId { get; set; }

        public RunbookState()
        {
        }
        public static new RunbookState Empty => new RunbookState();
    }
}