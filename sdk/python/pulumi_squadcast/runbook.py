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
from ._inputs import *

__all__ = ['RunbookArgs', 'Runbook']

@pulumi.input_type
class RunbookArgs:
    def __init__(__self__, *,
                 steps: pulumi.Input[Sequence[pulumi.Input['RunbookStepArgs']]],
                 team_id: pulumi.Input[str],
                 entity_owner: Optional[pulumi.Input['RunbookEntityOwnerArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Runbook resource.
        :param pulumi.Input[Sequence[pulumi.Input['RunbookStepArgs']]] steps: Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
        :param pulumi.Input[str] team_id: Team id.
        :param pulumi.Input['RunbookEntityOwnerArgs'] entity_owner: Runbooks owner.
        :param pulumi.Input[str] name: Name of the Runbook.
        """
        pulumi.set(__self__, "steps", steps)
        pulumi.set(__self__, "team_id", team_id)
        if entity_owner is not None:
            pulumi.set(__self__, "entity_owner", entity_owner)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def steps(self) -> pulumi.Input[Sequence[pulumi.Input['RunbookStepArgs']]]:
        """
        Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
        """
        return pulumi.get(self, "steps")

    @steps.setter
    def steps(self, value: pulumi.Input[Sequence[pulumi.Input['RunbookStepArgs']]]):
        pulumi.set(self, "steps", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Input[str]:
        """
        Team id.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter(name="entityOwner")
    def entity_owner(self) -> Optional[pulumi.Input['RunbookEntityOwnerArgs']]:
        """
        Runbooks owner.
        """
        return pulumi.get(self, "entity_owner")

    @entity_owner.setter
    def entity_owner(self, value: Optional[pulumi.Input['RunbookEntityOwnerArgs']]):
        pulumi.set(self, "entity_owner", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Runbook.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RunbookState:
    def __init__(__self__, *,
                 entity_owner: Optional[pulumi.Input['RunbookEntityOwnerArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 steps: Optional[pulumi.Input[Sequence[pulumi.Input['RunbookStepArgs']]]] = None,
                 team_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Runbook resources.
        :param pulumi.Input['RunbookEntityOwnerArgs'] entity_owner: Runbooks owner.
        :param pulumi.Input[str] name: Name of the Runbook.
        :param pulumi.Input[Sequence[pulumi.Input['RunbookStepArgs']]] steps: Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
        :param pulumi.Input[str] team_id: Team id.
        """
        if entity_owner is not None:
            pulumi.set(__self__, "entity_owner", entity_owner)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if steps is not None:
            pulumi.set(__self__, "steps", steps)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)

    @property
    @pulumi.getter(name="entityOwner")
    def entity_owner(self) -> Optional[pulumi.Input['RunbookEntityOwnerArgs']]:
        """
        Runbooks owner.
        """
        return pulumi.get(self, "entity_owner")

    @entity_owner.setter
    def entity_owner(self, value: Optional[pulumi.Input['RunbookEntityOwnerArgs']]):
        pulumi.set(self, "entity_owner", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Runbook.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def steps(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RunbookStepArgs']]]]:
        """
        Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
        """
        return pulumi.get(self, "steps")

    @steps.setter
    def steps(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RunbookStepArgs']]]]):
        pulumi.set(self, "steps", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        Team id.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)


class Runbook(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entity_owner: Optional[pulumi.Input[pulumi.InputType['RunbookEntityOwnerArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 steps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunbookStepArgs']]]]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        [Squadcast Runbook](https://support.squadcast.com/docs/runbooks) is a compilation of routine procedures and operations that are documented for reference while working on a critical incident. Sometimes, it can also be referred to as a Playbook.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_squadcast as squadcast

        example_team = squadcast.get_team(name="example team name")
        example_user = squadcast.get_user(email="test@example.com")
        example_runbook = squadcast.Runbook("exampleRunbook",
            team_id=example_team.id,
            steps=[
                squadcast.RunbookStepArgs(
                    content="some text here",
                ),
                squadcast.RunbookStepArgs(
                    content="some text here 2",
                ),
            ],
            entity_owner=squadcast.RunbookEntityOwnerArgs(
                id=example_user.id,
                type="user",
            ))
        ```

        ## Import

        teamID:runbookID Use 'Get All Teams' and 'Get All Runbooks' APIs to get the id of the team and runbook respectively

        ```sh
         $ pulumi import squadcast:index/runbook:Runbook test 62d2fe23a57381088224d726:62da76c088f407f9ca756ca5
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RunbookEntityOwnerArgs']] entity_owner: Runbooks owner.
        :param pulumi.Input[str] name: Name of the Runbook.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunbookStepArgs']]]] steps: Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
        :param pulumi.Input[str] team_id: Team id.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RunbookArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        [Squadcast Runbook](https://support.squadcast.com/docs/runbooks) is a compilation of routine procedures and operations that are documented for reference while working on a critical incident. Sometimes, it can also be referred to as a Playbook.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_squadcast as squadcast

        example_team = squadcast.get_team(name="example team name")
        example_user = squadcast.get_user(email="test@example.com")
        example_runbook = squadcast.Runbook("exampleRunbook",
            team_id=example_team.id,
            steps=[
                squadcast.RunbookStepArgs(
                    content="some text here",
                ),
                squadcast.RunbookStepArgs(
                    content="some text here 2",
                ),
            ],
            entity_owner=squadcast.RunbookEntityOwnerArgs(
                id=example_user.id,
                type="user",
            ))
        ```

        ## Import

        teamID:runbookID Use 'Get All Teams' and 'Get All Runbooks' APIs to get the id of the team and runbook respectively

        ```sh
         $ pulumi import squadcast:index/runbook:Runbook test 62d2fe23a57381088224d726:62da76c088f407f9ca756ca5
        ```

        :param str resource_name: The name of the resource.
        :param RunbookArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RunbookArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entity_owner: Optional[pulumi.Input[pulumi.InputType['RunbookEntityOwnerArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 steps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunbookStepArgs']]]]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RunbookArgs.__new__(RunbookArgs)

            __props__.__dict__["entity_owner"] = entity_owner
            __props__.__dict__["name"] = name
            if steps is None and not opts.urn:
                raise TypeError("Missing required property 'steps'")
            __props__.__dict__["steps"] = steps
            if team_id is None and not opts.urn:
                raise TypeError("Missing required property 'team_id'")
            __props__.__dict__["team_id"] = team_id
        super(Runbook, __self__).__init__(
            'squadcast:index/runbook:Runbook',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            entity_owner: Optional[pulumi.Input[pulumi.InputType['RunbookEntityOwnerArgs']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            steps: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunbookStepArgs']]]]] = None,
            team_id: Optional[pulumi.Input[str]] = None) -> 'Runbook':
        """
        Get an existing Runbook resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['RunbookEntityOwnerArgs']] entity_owner: Runbooks owner.
        :param pulumi.Input[str] name: Name of the Runbook.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RunbookStepArgs']]]] steps: Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
        :param pulumi.Input[str] team_id: Team id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RunbookState.__new__(_RunbookState)

        __props__.__dict__["entity_owner"] = entity_owner
        __props__.__dict__["name"] = name
        __props__.__dict__["steps"] = steps
        __props__.__dict__["team_id"] = team_id
        return Runbook(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="entityOwner")
    def entity_owner(self) -> pulumi.Output['outputs.RunbookEntityOwner']:
        """
        Runbooks owner.
        """
        return pulumi.get(self, "entity_owner")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the Runbook.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def steps(self) -> pulumi.Output[Sequence['outputs.RunbookStep']]:
        """
        Step by Step instructions, you can add as many steps as you want, supports markdown formatting.
        """
        return pulumi.get(self, "steps")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[str]:
        """
        Team id.
        """
        return pulumi.get(self, "team_id")
