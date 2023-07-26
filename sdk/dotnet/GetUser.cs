// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Squadcast
{
    public static class GetUser
    {
        /// <summary>
        /// Use this data source to get information about a specific user that you can use for other Squadcast resources.
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
        ///     var test = Squadcast.GetUser.Invoke(new()
        ///     {
        ///         Email = squadcast_user.Test.Email,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUserResult> InvokeAsync(GetUserArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUserResult>("squadcast:index/getUser:getUser", args ?? new GetUserArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a specific user that you can use for other Squadcast resources.
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
        ///     var test = Squadcast.GetUser.Invoke(new()
        ///     {
        ///         Email = squadcast_user.Test.Email,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUserResult> Invoke(GetUserInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUserResult>("squadcast:index/getUser:getUser", args ?? new GetUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// User email.
        /// </summary>
        [Input("email", required: true)]
        public string Email { get; set; } = null!;

        public GetUserArgs()
        {
        }
        public static new GetUserArgs Empty => new GetUserArgs();
    }

    public sealed class GetUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// User email.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        public GetUserInvokeArgs()
        {
        }
        public static new GetUserInvokeArgs Empty => new GetUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetUserResult
    {
        /// <summary>
        /// Denotes the Permissions / abilities of the user.
        /// </summary>
        public readonly ImmutableArray<string> Abilities;
        /// <summary>
        /// User email.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// User first name.
        /// </summary>
        public readonly string FirstName;
        /// <summary>
        /// User id.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Denotes if the user has verified their email or not.
        /// </summary>
        public readonly bool IsEmailVerified;
        /// <summary>
        /// Deprecated, this can be ignored.
        /// </summary>
        public readonly bool IsOverrideDndEnabled;
        /// <summary>
        /// Denotes if the user has verified their phone number or not.
        /// </summary>
        public readonly bool IsPhoneVerified;
        /// <summary>
        /// User last name.
        /// </summary>
        public readonly string LastName;
        /// <summary>
        /// User name, automatically computed from first name and last name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// User Personal Notification Rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserNotificationRuleResult> NotificationRules;
        /// <summary>
        /// User's personal on-call reminder notification rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUserOncallReminderRuleResult> OncallReminderRules;
        /// <summary>
        /// User phone number.
        /// </summary>
        public readonly string Phone;
        /// <summary>
        /// User role.
        /// </summary>
        public readonly string Role;
        /// <summary>
        /// User time_zone.
        /// </summary>
        public readonly string TimeZone;

        [OutputConstructor]
        private GetUserResult(
            ImmutableArray<string> abilities,

            string email,

            string firstName,

            string id,

            bool isEmailVerified,

            bool isOverrideDndEnabled,

            bool isPhoneVerified,

            string lastName,

            string name,

            ImmutableArray<Outputs.GetUserNotificationRuleResult> notificationRules,

            ImmutableArray<Outputs.GetUserOncallReminderRuleResult> oncallReminderRules,

            string phone,

            string role,

            string timeZone)
        {
            Abilities = abilities;
            Email = email;
            FirstName = firstName;
            Id = id;
            IsEmailVerified = isEmailVerified;
            IsOverrideDndEnabled = isOverrideDndEnabled;
            IsPhoneVerified = isPhoneVerified;
            LastName = lastName;
            Name = name;
            NotificationRules = notificationRules;
            OncallReminderRules = oncallReminderRules;
            Phone = phone;
            Role = role;
            TimeZone = timeZone;
        }
    }
}
