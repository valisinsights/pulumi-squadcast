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

__all__ = ['ServiceMaintenanceArgs', 'ServiceMaintenance']

@pulumi.input_type
class ServiceMaintenanceArgs:
    def __init__(__self__, *,
                 service_id: pulumi.Input[str],
                 windows: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceMaintenanceWindowArgs']]]] = None):
        """
        The set of arguments for constructing a ServiceMaintenance resource.
        :param pulumi.Input[str] service_id: Service id.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceMaintenanceWindowArgs']]] windows: Date and Time range during which maintenance would be carried out
        """
        pulumi.set(__self__, "service_id", service_id)
        if windows is not None:
            pulumi.set(__self__, "windows", windows)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        """
        Service id.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter
    def windows(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceMaintenanceWindowArgs']]]]:
        """
        Date and Time range during which maintenance would be carried out
        """
        return pulumi.get(self, "windows")

    @windows.setter
    def windows(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceMaintenanceWindowArgs']]]]):
        pulumi.set(self, "windows", value)


@pulumi.input_type
class _ServiceMaintenanceState:
    def __init__(__self__, *,
                 service_id: Optional[pulumi.Input[str]] = None,
                 windows: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceMaintenanceWindowArgs']]]] = None):
        """
        Input properties used for looking up and filtering ServiceMaintenance resources.
        :param pulumi.Input[str] service_id: Service id.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceMaintenanceWindowArgs']]] windows: Date and Time range during which maintenance would be carried out
        """
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)
        if windows is not None:
            pulumi.set(__self__, "windows", windows)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        """
        Service id.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter
    def windows(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceMaintenanceWindowArgs']]]]:
        """
        Date and Time range during which maintenance would be carried out
        """
        return pulumi.get(self, "windows")

    @windows.setter
    def windows(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceMaintenanceWindowArgs']]]]):
        pulumi.set(self, "windows", value)


class ServiceMaintenance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 windows: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceMaintenanceWindowArgs']]]]] = None,
                 __props__=None):
        """
        [Maintenance Mode](https://support.squadcast.com/docs/maintenance-mode) enables you to reduce alert noise during the scheduled maintenance window. Alerts generated during active maintenance windows would be automatically suppressed and hence, no notifications are generated for those suppressed alerts.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_squadcast as squadcast

        example_team = squadcast.get_team(name="example team name")
        example_service = squadcast.get_service(name="example service name",
            team_id=example_team.id)
        example_service_maintenance = squadcast.ServiceMaintenance("exampleServiceMaintenance",
            service_id=example_service.id,
            windows=[
                squadcast.ServiceMaintenanceWindowArgs(
                    from_="2032-06-01T10:30:00.000Z",
                    till="2032-06-01T11:30:00.000Z",
                    repeat_till="2032-06-30T10:30:00.000Z",
                    repeat_frequency="week",
                ),
                squadcast.ServiceMaintenanceWindowArgs(
                    from_="2032-07-01T10:30:00.000Z",
                    till="2032-07-02T10:30:00.000Z",
                ),
            ])
        ```

        ## Import

        teamID:serviceID Use 'Get All Teams' and 'Get All Services' APIs to get the id of the team and service respectively

        ```sh
         $ pulumi import squadcast:index/serviceMaintenance:ServiceMaintenance test 62d2fe23a57381088224d726:62da76c088f407f9ca756ca5
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] service_id: Service id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceMaintenanceWindowArgs']]]] windows: Date and Time range during which maintenance would be carried out
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceMaintenanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        [Maintenance Mode](https://support.squadcast.com/docs/maintenance-mode) enables you to reduce alert noise during the scheduled maintenance window. Alerts generated during active maintenance windows would be automatically suppressed and hence, no notifications are generated for those suppressed alerts.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_squadcast as squadcast

        example_team = squadcast.get_team(name="example team name")
        example_service = squadcast.get_service(name="example service name",
            team_id=example_team.id)
        example_service_maintenance = squadcast.ServiceMaintenance("exampleServiceMaintenance",
            service_id=example_service.id,
            windows=[
                squadcast.ServiceMaintenanceWindowArgs(
                    from_="2032-06-01T10:30:00.000Z",
                    till="2032-06-01T11:30:00.000Z",
                    repeat_till="2032-06-30T10:30:00.000Z",
                    repeat_frequency="week",
                ),
                squadcast.ServiceMaintenanceWindowArgs(
                    from_="2032-07-01T10:30:00.000Z",
                    till="2032-07-02T10:30:00.000Z",
                ),
            ])
        ```

        ## Import

        teamID:serviceID Use 'Get All Teams' and 'Get All Services' APIs to get the id of the team and service respectively

        ```sh
         $ pulumi import squadcast:index/serviceMaintenance:ServiceMaintenance test 62d2fe23a57381088224d726:62da76c088f407f9ca756ca5
        ```

        :param str resource_name: The name of the resource.
        :param ServiceMaintenanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceMaintenanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 windows: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceMaintenanceWindowArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ServiceMaintenanceArgs.__new__(ServiceMaintenanceArgs)

            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["windows"] = windows
        super(ServiceMaintenance, __self__).__init__(
            'squadcast:index/serviceMaintenance:ServiceMaintenance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            service_id: Optional[pulumi.Input[str]] = None,
            windows: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceMaintenanceWindowArgs']]]]] = None) -> 'ServiceMaintenance':
        """
        Get an existing ServiceMaintenance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] service_id: Service id.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceMaintenanceWindowArgs']]]] windows: Date and Time range during which maintenance would be carried out
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceMaintenanceState.__new__(_ServiceMaintenanceState)

        __props__.__dict__["service_id"] = service_id
        __props__.__dict__["windows"] = windows
        return ServiceMaintenance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        """
        Service id.
        """
        return pulumi.get(self, "service_id")

    @property
    @pulumi.getter
    def windows(self) -> pulumi.Output[Optional[Sequence['outputs.ServiceMaintenanceWindow']]]:
        """
        Date and Time range during which maintenance would be carried out
        """
        return pulumi.get(self, "windows")
