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
    'GetServiceResult',
    'AwaitableGetServiceResult',
    'get_service',
    'get_service_output',
]

@pulumi.output_type
class GetServiceResult:
    """
    A collection of values returned by getService.
    """
    def __init__(__self__, active_alert_source_endpoints=None, alert_source_endpoints=None, api_key=None, dependencies=None, description=None, email=None, email_prefix=None, escalation_policy_id=None, id=None, maintainers=None, name=None, slack_channel_id=None, tags=None, team_id=None):
        if active_alert_source_endpoints and not isinstance(active_alert_source_endpoints, dict):
            raise TypeError("Expected argument 'active_alert_source_endpoints' to be a dict")
        pulumi.set(__self__, "active_alert_source_endpoints", active_alert_source_endpoints)
        if alert_source_endpoints and not isinstance(alert_source_endpoints, dict):
            raise TypeError("Expected argument 'alert_source_endpoints' to be a dict")
        pulumi.set(__self__, "alert_source_endpoints", alert_source_endpoints)
        if api_key and not isinstance(api_key, str):
            raise TypeError("Expected argument 'api_key' to be a str")
        pulumi.set(__self__, "api_key", api_key)
        if dependencies and not isinstance(dependencies, list):
            raise TypeError("Expected argument 'dependencies' to be a list")
        pulumi.set(__self__, "dependencies", dependencies)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if email_prefix and not isinstance(email_prefix, str):
            raise TypeError("Expected argument 'email_prefix' to be a str")
        pulumi.set(__self__, "email_prefix", email_prefix)
        if escalation_policy_id and not isinstance(escalation_policy_id, str):
            raise TypeError("Expected argument 'escalation_policy_id' to be a str")
        pulumi.set(__self__, "escalation_policy_id", escalation_policy_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if maintainers and not isinstance(maintainers, list):
            raise TypeError("Expected argument 'maintainers' to be a list")
        pulumi.set(__self__, "maintainers", maintainers)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if slack_channel_id and not isinstance(slack_channel_id, str):
            raise TypeError("Expected argument 'slack_channel_id' to be a str")
        pulumi.set(__self__, "slack_channel_id", slack_channel_id)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if team_id and not isinstance(team_id, str):
            raise TypeError("Expected argument 'team_id' to be a str")
        pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter(name="activeAlertSourceEndpoints")
    def active_alert_source_endpoints(self) -> Mapping[str, str]:
        """
        Active alert source endpoints.
        """
        return pulumi.get(self, "active_alert_source_endpoints")

    @property
    @pulumi.getter(name="alertSourceEndpoints")
    def alert_source_endpoints(self) -> Mapping[str, str]:
        """
        All available alert source endpoints.
        """
        return pulumi.get(self, "alert_source_endpoints")

    @property
    @pulumi.getter(name="apiKey")
    def api_key(self) -> str:
        """
        Unique API key of the service
        """
        return pulumi.get(self, "api_key")

    @property
    @pulumi.getter
    def dependencies(self) -> Sequence[str]:
        """
        dependencies.
        """
        return pulumi.get(self, "dependencies")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Detailed description about the service.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        Email.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="emailPrefix")
    def email_prefix(self) -> str:
        """
        Email prefix.
        """
        return pulumi.get(self, "email_prefix")

    @property
    @pulumi.getter(name="escalationPolicyId")
    def escalation_policy_id(self) -> str:
        """
        Escalation policy id.
        """
        return pulumi.get(self, "escalation_policy_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        Service id.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def maintainers(self) -> Sequence['outputs.GetServiceMaintainerResult']:
        """
        Service owner
        """
        return pulumi.get(self, "maintainers")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the Service.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="slackChannelId")
    def slack_channel_id(self) -> str:
        """
        Slack extension for the service. If set, specifies the ID of the Slack channel associated with the service. If this ID is set, it cannot be removed, but it can be changed to a different slack*channel*id.
        """
        return pulumi.get(self, "slack_channel_id")

    @property
    @pulumi.getter
    def tags(self) -> Sequence['outputs.GetServiceTagResult']:
        """
        Service tags
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> str:
        """
        Team id.
        """
        return pulumi.get(self, "team_id")


class AwaitableGetServiceResult(GetServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceResult(
            active_alert_source_endpoints=self.active_alert_source_endpoints,
            alert_source_endpoints=self.alert_source_endpoints,
            api_key=self.api_key,
            dependencies=self.dependencies,
            description=self.description,
            email=self.email,
            email_prefix=self.email_prefix,
            escalation_policy_id=self.escalation_policy_id,
            id=self.id,
            maintainers=self.maintainers,
            name=self.name,
            slack_channel_id=self.slack_channel_id,
            tags=self.tags,
            team_id=self.team_id)


def get_service(name: Optional[str] = None,
                team_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceResult:
    """
    [Squadcast Services](https://support.squadcast.com/docs/adding-a-service-1) are the core components of your infrastructure/application for which alerts are generated. Services in Squadcast represent specific systems, applications, components, products, or teams for which an incident is created. To check out some of the best practices on creating Services in Squadcast, refer to the guide [here](https://www.squadcast.com/blog/how-to-configure-services-in-squadcast-best-practices-to-reduce-mttr).Use this data source to get information about a specific service.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_squadcast as squadcast

    test = squadcast.get_service(name=squadcast_service["test"]["name"],
        team_id="team id")
    ```


    :param str name: Name of the Service.
    :param str team_id: Team id.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['teamId'] = team_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('squadcast:index/getService:getService', __args__, opts=opts, typ=GetServiceResult).value

    return AwaitableGetServiceResult(
        active_alert_source_endpoints=pulumi.get(__ret__, 'active_alert_source_endpoints'),
        alert_source_endpoints=pulumi.get(__ret__, 'alert_source_endpoints'),
        api_key=pulumi.get(__ret__, 'api_key'),
        dependencies=pulumi.get(__ret__, 'dependencies'),
        description=pulumi.get(__ret__, 'description'),
        email=pulumi.get(__ret__, 'email'),
        email_prefix=pulumi.get(__ret__, 'email_prefix'),
        escalation_policy_id=pulumi.get(__ret__, 'escalation_policy_id'),
        id=pulumi.get(__ret__, 'id'),
        maintainers=pulumi.get(__ret__, 'maintainers'),
        name=pulumi.get(__ret__, 'name'),
        slack_channel_id=pulumi.get(__ret__, 'slack_channel_id'),
        tags=pulumi.get(__ret__, 'tags'),
        team_id=pulumi.get(__ret__, 'team_id'))


@_utilities.lift_output_func(get_service)
def get_service_output(name: Optional[pulumi.Input[str]] = None,
                       team_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceResult]:
    """
    [Squadcast Services](https://support.squadcast.com/docs/adding-a-service-1) are the core components of your infrastructure/application for which alerts are generated. Services in Squadcast represent specific systems, applications, components, products, or teams for which an incident is created. To check out some of the best practices on creating Services in Squadcast, refer to the guide [here](https://www.squadcast.com/blog/how-to-configure-services-in-squadcast-best-practices-to-reduce-mttr).Use this data source to get information about a specific service.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_squadcast as squadcast

    test = squadcast.get_service(name=squadcast_service["test"]["name"],
        team_id="team id")
    ```


    :param str name: Name of the Service.
    :param str team_id: Team id.
    """
    ...