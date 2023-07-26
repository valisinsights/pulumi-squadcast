# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, abilities=None, email=None, first_name=None, id=None, is_email_verified=None, is_override_dnd_enabled=None, is_phone_verified=None, last_name=None, name=None, notification_rules=None, oncall_reminder_rules=None, phone=None, role=None, time_zone=None):
        if abilities and not isinstance(abilities, list):
            raise TypeError("Expected argument 'abilities' to be a list")
        pulumi.set(__self__, "abilities", abilities)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if first_name and not isinstance(first_name, str):
            raise TypeError("Expected argument 'first_name' to be a str")
        pulumi.set(__self__, "first_name", first_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_email_verified and not isinstance(is_email_verified, bool):
            raise TypeError("Expected argument 'is_email_verified' to be a bool")
        pulumi.set(__self__, "is_email_verified", is_email_verified)
        if is_override_dnd_enabled and not isinstance(is_override_dnd_enabled, bool):
            raise TypeError("Expected argument 'is_override_dnd_enabled' to be a bool")
        pulumi.set(__self__, "is_override_dnd_enabled", is_override_dnd_enabled)
        if is_phone_verified and not isinstance(is_phone_verified, bool):
            raise TypeError("Expected argument 'is_phone_verified' to be a bool")
        pulumi.set(__self__, "is_phone_verified", is_phone_verified)
        if last_name and not isinstance(last_name, str):
            raise TypeError("Expected argument 'last_name' to be a str")
        pulumi.set(__self__, "last_name", last_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notification_rules and not isinstance(notification_rules, list):
            raise TypeError("Expected argument 'notification_rules' to be a list")
        pulumi.set(__self__, "notification_rules", notification_rules)
        if oncall_reminder_rules and not isinstance(oncall_reminder_rules, list):
            raise TypeError("Expected argument 'oncall_reminder_rules' to be a list")
        pulumi.set(__self__, "oncall_reminder_rules", oncall_reminder_rules)
        if phone and not isinstance(phone, str):
            raise TypeError("Expected argument 'phone' to be a str")
        pulumi.set(__self__, "phone", phone)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if time_zone and not isinstance(time_zone, str):
            raise TypeError("Expected argument 'time_zone' to be a str")
        pulumi.set(__self__, "time_zone", time_zone)

    @property
    @pulumi.getter
    def abilities(self) -> Sequence[str]:
        """
        Denotes the Permissions / abilities of the user.
        """
        return pulumi.get(self, "abilities")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        User email.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> str:
        """
        User first name.
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        User id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isEmailVerified")
    def is_email_verified(self) -> bool:
        """
        Denotes if the user has verified their email or not.
        """
        return pulumi.get(self, "is_email_verified")

    @property
    @pulumi.getter(name="isOverrideDndEnabled")
    def is_override_dnd_enabled(self) -> bool:
        """
        Deprecated, this can be ignored.
        """
        return pulumi.get(self, "is_override_dnd_enabled")

    @property
    @pulumi.getter(name="isPhoneVerified")
    def is_phone_verified(self) -> bool:
        """
        Denotes if the user has verified their phone number or not.
        """
        return pulumi.get(self, "is_phone_verified")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> str:
        """
        User last name.
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        User name, automatically computed from first name and last name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationRules")
    def notification_rules(self) -> Sequence['outputs.GetUserNotificationRuleResult']:
        """
        User Personal Notification Rules.
        """
        return pulumi.get(self, "notification_rules")

    @property
    @pulumi.getter(name="oncallReminderRules")
    def oncall_reminder_rules(self) -> Sequence['outputs.GetUserOncallReminderRuleResult']:
        """
        User's personal on-call reminder notification rules.
        """
        return pulumi.get(self, "oncall_reminder_rules")

    @property
    @pulumi.getter
    def phone(self) -> str:
        """
        User phone number.
        """
        return pulumi.get(self, "phone")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        User role.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="timeZone")
    def time_zone(self) -> str:
        """
        User time_zone.
        """
        return pulumi.get(self, "time_zone")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            abilities=self.abilities,
            email=self.email,
            first_name=self.first_name,
            id=self.id,
            is_email_verified=self.is_email_verified,
            is_override_dnd_enabled=self.is_override_dnd_enabled,
            is_phone_verified=self.is_phone_verified,
            last_name=self.last_name,
            name=self.name,
            notification_rules=self.notification_rules,
            oncall_reminder_rules=self.oncall_reminder_rules,
            phone=self.phone,
            role=self.role,
            time_zone=self.time_zone)


def get_user(email: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    Use this data source to get information about a specific user that you can use for other Squadcast resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_squadcast as squadcast

    test = squadcast.get_user(email=squadcast_user["test"]["email"])
    ```


    :param str email: User email.
    """
    __args__ = dict()
    __args__['email'] = email
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('squadcast:index/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        abilities=pulumi.get(__ret__, 'abilities'),
        email=pulumi.get(__ret__, 'email'),
        first_name=pulumi.get(__ret__, 'first_name'),
        id=pulumi.get(__ret__, 'id'),
        is_email_verified=pulumi.get(__ret__, 'is_email_verified'),
        is_override_dnd_enabled=pulumi.get(__ret__, 'is_override_dnd_enabled'),
        is_phone_verified=pulumi.get(__ret__, 'is_phone_verified'),
        last_name=pulumi.get(__ret__, 'last_name'),
        name=pulumi.get(__ret__, 'name'),
        notification_rules=pulumi.get(__ret__, 'notification_rules'),
        oncall_reminder_rules=pulumi.get(__ret__, 'oncall_reminder_rules'),
        phone=pulumi.get(__ret__, 'phone'),
        role=pulumi.get(__ret__, 'role'),
        time_zone=pulumi.get(__ret__, 'time_zone'))


@_utilities.lift_output_func(get_user)
def get_user_output(email: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    Use this data source to get information about a specific user that you can use for other Squadcast resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_squadcast as squadcast

    test = squadcast.get_user(email=squadcast_user["test"]["email"])
    ```


    :param str email: User email.
    """
    ...