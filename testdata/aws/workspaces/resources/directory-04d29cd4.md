Provides a WorkSpaces directory in AWS WorkSpaces Service.

> **NOTE:** AWS WorkSpaces service requires [`workspaces_DefaultRole`](https://docs.aws.amazon.com/workspaces/latest/adminguide/workspaces-access-control.html#create-default-role) IAM role to operate normally.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const workspaces = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["workspaces.amazonaws.com"],
        }],
    }],
});
const workspacesDefault = new aws.iam.Role("workspacesDefault", {assumeRolePolicy: workspaces.then(workspaces => workspaces.json)});
const workspacesDefaultServiceAccess = new aws.iam.RolePolicyAttachment("workspacesDefaultServiceAccess", {
    role: workspacesDefault.name,
    policyArn: "arn:aws:iam::aws:policy/AmazonWorkSpacesServiceAccess",
});
const workspacesDefaultSelfServiceAccess = new aws.iam.RolePolicyAttachment("workspacesDefaultSelfServiceAccess", {
    role: workspacesDefault.name,
    policyArn: "arn:aws:iam::aws:policy/AmazonWorkSpacesSelfServiceAccess",
});
const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.0.0.0/16"});
const exampleC = new aws.ec2.Subnet("exampleC", {
    vpcId: exampleVpc.id,
    availabilityZone: "us-east-1c",
    cidrBlock: "10.0.2.0/24",
});
const exampleD = new aws.ec2.Subnet("exampleD", {
    vpcId: exampleVpc.id,
    availabilityZone: "us-east-1d",
    cidrBlock: "10.0.3.0/24",
});
const exampleDirectory = new aws.workspaces.Directory("exampleDirectory", {
    directoryId: exampleDirectoryservice / directoryDirectory.id,
    subnetIds: [
        exampleC.id,
        exampleD.id,
    ],
    tags: {
        Example: "true",
    },
    selfServicePermissions: {
        changeComputeType: true,
        increaseVolumeSize: true,
        rebuildWorkspace: true,
        restartWorkspace: true,
        switchRunningMode: true,
    },
    workspaceAccessProperties: {
        deviceTypeAndroid: "ALLOW",
        deviceTypeChromeos: "ALLOW",
        deviceTypeIos: "ALLOW",
        deviceTypeLinux: "DENY",
        deviceTypeOsx: "ALLOW",
        deviceTypeWeb: "DENY",
        deviceTypeWindows: "DENY",
        deviceTypeZeroclient: "DENY",
    },
    workspaceCreationProperties: {
        customSecurityGroupId: aws_security_group.example.id,
        defaultOu: "OU=AWS,DC=Workgroup,DC=Example,DC=com",
        enableInternetAccess: true,
        enableMaintenanceMode: true,
        userEnabledAsLocalAdministrator: true,
    },
}, {
    dependsOn: [
        workspacesDefaultServiceAccess,
        workspacesDefaultSelfServiceAccess,
    ],
});
const exampleA = new aws.ec2.Subnet("exampleA", {
    vpcId: exampleVpc.id,
    availabilityZone: "us-east-1a",
    cidrBlock: "10.0.0.0/24",
});
const exampleB = new aws.ec2.Subnet("exampleB", {
    vpcId: exampleVpc.id,
    availabilityZone: "us-east-1b",
    cidrBlock: "10.0.1.0/24",
});
const exampleDirectoryservice_directoryDirectory = new aws.directoryservice.Directory("exampleDirectoryservice/directoryDirectory", {
    name: "corp.example.com",
    password: "#S1ncerely",
    size: "Small",
    vpcSettings: {
        vpcId: exampleVpc.id,
        subnetIds: [
            exampleA.id,
            exampleB.id,
        ],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

workspaces = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["workspaces.amazonaws.com"],
    )],
)])
workspaces_default = aws.iam.Role("workspacesDefault", assume_role_policy=workspaces.json)
workspaces_default_service_access = aws.iam.RolePolicyAttachment("workspacesDefaultServiceAccess",
    role=workspaces_default.name,
    policy_arn="arn:aws:iam::aws:policy/AmazonWorkSpacesServiceAccess")
workspaces_default_self_service_access = aws.iam.RolePolicyAttachment("workspacesDefaultSelfServiceAccess",
    role=workspaces_default.name,
    policy_arn="arn:aws:iam::aws:policy/AmazonWorkSpacesSelfServiceAccess")
example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
example_c = aws.ec2.Subnet("exampleC",
    vpc_id=example_vpc.id,
    availability_zone="us-east-1c",
    cidr_block="10.0.2.0/24")
example_d = aws.ec2.Subnet("exampleD",
    vpc_id=example_vpc.id,
    availability_zone="us-east-1d",
    cidr_block="10.0.3.0/24")
example_directory = aws.workspaces.Directory("exampleDirectory",
    directory_id=example_directoryservice / directory_directory["id"],
    subnet_ids=[
        example_c.id,
        example_d.id,
    ],
    tags={
        "Example": "true",
    },
    self_service_permissions=aws.workspaces.DirectorySelfServicePermissionsArgs(
        change_compute_type=True,
        increase_volume_size=True,
        rebuild_workspace=True,
        restart_workspace=True,
        switch_running_mode=True,
    ),
    workspace_access_properties=aws.workspaces.DirectoryWorkspaceAccessPropertiesArgs(
        device_type_android="ALLOW",
        device_type_chromeos="ALLOW",
        device_type_ios="ALLOW",
        device_type_linux="DENY",
        device_type_osx="ALLOW",
        device_type_web="DENY",
        device_type_windows="DENY",
        device_type_zeroclient="DENY",
    ),
    workspace_creation_properties=aws.workspaces.DirectoryWorkspaceCreationPropertiesArgs(
        custom_security_group_id=aws_security_group["example"]["id"],
        default_ou="OU=AWS,DC=Workgroup,DC=Example,DC=com",
        enable_internet_access=True,
        enable_maintenance_mode=True,
        user_enabled_as_local_administrator=True,
    ),
    opts=pulumi.ResourceOptions(depends_on=[
            workspaces_default_service_access,
            workspaces_default_self_service_access,
        ]))
example_a = aws.ec2.Subnet("exampleA",
    vpc_id=example_vpc.id,
    availability_zone="us-east-1a",
    cidr_block="10.0.0.0/24")
example_b = aws.ec2.Subnet("exampleB",
    vpc_id=example_vpc.id,
    availability_zone="us-east-1b",
    cidr_block="10.0.1.0/24")
example_directoryservice_directory_directory = aws.directoryservice.Directory("exampleDirectoryservice/directoryDirectory",
    name="corp.example.com",
    password="#S1ncerely",
    size="Small",
    vpc_settings=aws.directoryservice.DirectoryVpcSettingsArgs(
        vpc_id=example_vpc.id,
        subnet_ids=[
            example_a.id,
            example_b.id,
        ],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var workspaces = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "workspaces.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var workspacesDefault = new Aws.Iam.Role("workspacesDefault", new()
    {
        AssumeRolePolicy = workspaces.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var workspacesDefaultServiceAccess = new Aws.Iam.RolePolicyAttachment("workspacesDefaultServiceAccess", new()
    {
        Role = workspacesDefault.Name,
        PolicyArn = "arn:aws:iam::aws:policy/AmazonWorkSpacesServiceAccess",
    });

    var workspacesDefaultSelfServiceAccess = new Aws.Iam.RolePolicyAttachment("workspacesDefaultSelfServiceAccess", new()
    {
        Role = workspacesDefault.Name,
        PolicyArn = "arn:aws:iam::aws:policy/AmazonWorkSpacesSelfServiceAccess",
    });

    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var exampleC = new Aws.Ec2.Subnet("exampleC", new()
    {
        VpcId = exampleVpc.Id,
        AvailabilityZone = "us-east-1c",
        CidrBlock = "10.0.2.0/24",
    });

    var exampleD = new Aws.Ec2.Subnet("exampleD", new()
    {
        VpcId = exampleVpc.Id,
        AvailabilityZone = "us-east-1d",
        CidrBlock = "10.0.3.0/24",
    });

    var exampleDirectory = new Aws.Workspaces.Directory("exampleDirectory", new()
    {
        DirectoryId = exampleDirectoryservice / directoryDirectory.Id,
        SubnetIds = new[]
        {
            exampleC.Id,
            exampleD.Id,
        },
        Tags = 
        {
            { "Example", "true" },
        },
        SelfServicePermissions = new Aws.Workspaces.Inputs.DirectorySelfServicePermissionsArgs
        {
            ChangeComputeType = true,
            IncreaseVolumeSize = true,
            RebuildWorkspace = true,
            RestartWorkspace = true,
            SwitchRunningMode = true,
        },
        WorkspaceAccessProperties = new Aws.Workspaces.Inputs.DirectoryWorkspaceAccessPropertiesArgs
        {
            DeviceTypeAndroid = "ALLOW",
            DeviceTypeChromeos = "ALLOW",
            DeviceTypeIos = "ALLOW",
            DeviceTypeLinux = "DENY",
            DeviceTypeOsx = "ALLOW",
            DeviceTypeWeb = "DENY",
            DeviceTypeWindows = "DENY",
            DeviceTypeZeroclient = "DENY",
        },
        WorkspaceCreationProperties = new Aws.Workspaces.Inputs.DirectoryWorkspaceCreationPropertiesArgs
        {
            CustomSecurityGroupId = aws_security_group.Example.Id,
            DefaultOu = "OU=AWS,DC=Workgroup,DC=Example,DC=com",
            EnableInternetAccess = true,
            EnableMaintenanceMode = true,
            UserEnabledAsLocalAdministrator = true,
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            workspacesDefaultServiceAccess,
            workspacesDefaultSelfServiceAccess,
        },
    });

    var exampleA = new Aws.Ec2.Subnet("exampleA", new()
    {
        VpcId = exampleVpc.Id,
        AvailabilityZone = "us-east-1a",
        CidrBlock = "10.0.0.0/24",
    });

    var exampleB = new Aws.Ec2.Subnet("exampleB", new()
    {
        VpcId = exampleVpc.Id,
        AvailabilityZone = "us-east-1b",
        CidrBlock = "10.0.1.0/24",
    });

    var exampleDirectoryservice_directoryDirectory = new Aws.DirectoryService.Directory("exampleDirectoryservice/directoryDirectory", new()
    {
        Name = "corp.example.com",
        Password = "#S1ncerely",
        Size = "Small",
        VpcSettings = new Aws.DirectoryService.Inputs.DirectoryVpcSettingsArgs
        {
            VpcId = exampleVpc.Id,
            SubnetIds = new[]
            {
                exampleA.Id,
                exampleB.Id,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/workspaces"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		workspaces, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"workspaces.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		workspacesDefault, err := iam.NewRole(ctx, "workspacesDefault", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(workspaces.Json),
		})
		if err != nil {
			return err
		}
		workspacesDefaultServiceAccess, err := iam.NewRolePolicyAttachment(ctx, "workspacesDefaultServiceAccess", &iam.RolePolicyAttachmentArgs{
			Role:      workspacesDefault.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonWorkSpacesServiceAccess"),
		})
		if err != nil {
			return err
		}
		workspacesDefaultSelfServiceAccess, err := iam.NewRolePolicyAttachment(ctx, "workspacesDefaultSelfServiceAccess", &iam.RolePolicyAttachmentArgs{
			Role:      workspacesDefault.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonWorkSpacesSelfServiceAccess"),
		})
		if err != nil {
			return err
		}
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		exampleC, err := ec2.NewSubnet(ctx, "exampleC", &ec2.SubnetArgs{
			VpcId:            exampleVpc.ID(),
			AvailabilityZone: pulumi.String("us-east-1c"),
			CidrBlock:        pulumi.String("10.0.2.0/24"),
		})
		if err != nil {
			return err
		}
		exampleD, err := ec2.NewSubnet(ctx, "exampleD", &ec2.SubnetArgs{
			VpcId:            exampleVpc.ID(),
			AvailabilityZone: pulumi.String("us-east-1d"),
			CidrBlock:        pulumi.String("10.0.3.0/24"),
		})
		if err != nil {
			return err
		}
		_, err = workspaces.NewDirectory(ctx, "exampleDirectory", &workspaces.DirectoryArgs{
			DirectoryId: exampleDirectoryservice / directoryDirectory.Id,
			SubnetIds: pulumi.StringArray{
				exampleC.ID(),
				exampleD.ID(),
			},
			Tags: pulumi.StringMap{
				"Example": pulumi.String("true"),
			},
			SelfServicePermissions: &workspaces.DirectorySelfServicePermissionsArgs{
				ChangeComputeType:  pulumi.Bool(true),
				IncreaseVolumeSize: pulumi.Bool(true),
				RebuildWorkspace:   pulumi.Bool(true),
				RestartWorkspace:   pulumi.Bool(true),
				SwitchRunningMode:  pulumi.Bool(true),
			},
			WorkspaceAccessProperties: &workspaces.DirectoryWorkspaceAccessPropertiesArgs{
				DeviceTypeAndroid:    pulumi.String("ALLOW"),
				DeviceTypeChromeos:   pulumi.String("ALLOW"),
				DeviceTypeIos:        pulumi.String("ALLOW"),
				DeviceTypeLinux:      pulumi.String("DENY"),
				DeviceTypeOsx:        pulumi.String("ALLOW"),
				DeviceTypeWeb:        pulumi.String("DENY"),
				DeviceTypeWindows:    pulumi.String("DENY"),
				DeviceTypeZeroclient: pulumi.String("DENY"),
			},
			WorkspaceCreationProperties: &workspaces.DirectoryWorkspaceCreationPropertiesArgs{
				CustomSecurityGroupId:           pulumi.Any(aws_security_group.Example.Id),
				DefaultOu:                       pulumi.String("OU=AWS,DC=Workgroup,DC=Example,DC=com"),
				EnableInternetAccess:            pulumi.Bool(true),
				EnableMaintenanceMode:           pulumi.Bool(true),
				UserEnabledAsLocalAdministrator: pulumi.Bool(true),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			workspacesDefaultServiceAccess,
			workspacesDefaultSelfServiceAccess,
		}))
		if err != nil {
			return err
		}
		exampleA, err := ec2.NewSubnet(ctx, "exampleA", &ec2.SubnetArgs{
			VpcId:            exampleVpc.ID(),
			AvailabilityZone: pulumi.String("us-east-1a"),
			CidrBlock:        pulumi.String("10.0.0.0/24"),
		})
		if err != nil {
			return err
		}
		exampleB, err := ec2.NewSubnet(ctx, "exampleB", &ec2.SubnetArgs{
			VpcId:            exampleVpc.ID(),
			AvailabilityZone: pulumi.String("us-east-1b"),
			CidrBlock:        pulumi.String("10.0.1.0/24"),
		})
		if err != nil {
			return err
		}
		_, err = directoryservice.NewDirectory(ctx, "exampleDirectoryservice/directoryDirectory", &directoryservice.DirectoryArgs{
			Name:     pulumi.String("corp.example.com"),
			Password: pulumi.String("#S1ncerely"),
			Size:     pulumi.String("Small"),
			VpcSettings: &directoryservice.DirectoryVpcSettingsArgs{
				VpcId: exampleVpc.ID(),
				SubnetIds: pulumi.StringArray{
					exampleA.ID(),
					exampleB.ID(),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.workspaces.Directory;
import com.pulumi.aws.workspaces.DirectoryArgs;
import com.pulumi.aws.workspaces.inputs.DirectorySelfServicePermissionsArgs;
import com.pulumi.aws.workspaces.inputs.DirectoryWorkspaceAccessPropertiesArgs;
import com.pulumi.aws.workspaces.inputs.DirectoryWorkspaceCreationPropertiesArgs;
import com.pulumi.aws.directoryservice.Directory;
import com.pulumi.aws.directoryservice.DirectoryArgs;
import com.pulumi.aws.directoryservice.inputs.DirectoryVpcSettingsArgs;
import com.pulumi.resources.CustomResourceOptions;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        final var workspaces = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("workspaces.amazonaws.com")
                    .build())
                .build())
            .build());

        var workspacesDefault = new Role("workspacesDefault", RoleArgs.builder()        
            .assumeRolePolicy(workspaces.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var workspacesDefaultServiceAccess = new RolePolicyAttachment("workspacesDefaultServiceAccess", RolePolicyAttachmentArgs.builder()        
            .role(workspacesDefault.name())
            .policyArn("arn:aws:iam::aws:policy/AmazonWorkSpacesServiceAccess")
            .build());

        var workspacesDefaultSelfServiceAccess = new RolePolicyAttachment("workspacesDefaultSelfServiceAccess", RolePolicyAttachmentArgs.builder()        
            .role(workspacesDefault.name())
            .policyArn("arn:aws:iam::aws:policy/AmazonWorkSpacesSelfServiceAccess")
            .build());

        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var exampleC = new Subnet("exampleC", SubnetArgs.builder()        
            .vpcId(exampleVpc.id())
            .availabilityZone("us-east-1c")
            .cidrBlock("10.0.2.0/24")
            .build());

        var exampleD = new Subnet("exampleD", SubnetArgs.builder()        
            .vpcId(exampleVpc.id())
            .availabilityZone("us-east-1d")
            .cidrBlock("10.0.3.0/24")
            .build());

        var exampleDirectory = new Directory("exampleDirectory", DirectoryArgs.builder()        
            .directoryId(exampleDirectoryservice / directoryDirectory.id())
            .subnetIds(            
                exampleC.id(),
                exampleD.id())
            .tags(Map.of("Example", true))
            .selfServicePermissions(DirectorySelfServicePermissionsArgs.builder()
                .changeComputeType(true)
                .increaseVolumeSize(true)
                .rebuildWorkspace(true)
                .restartWorkspace(true)
                .switchRunningMode(true)
                .build())
            .workspaceAccessProperties(DirectoryWorkspaceAccessPropertiesArgs.builder()
                .deviceTypeAndroid("ALLOW")
                .deviceTypeChromeos("ALLOW")
                .deviceTypeIos("ALLOW")
                .deviceTypeLinux("DENY")
                .deviceTypeOsx("ALLOW")
                .deviceTypeWeb("DENY")
                .deviceTypeWindows("DENY")
                .deviceTypeZeroclient("DENY")
                .build())
            .workspaceCreationProperties(DirectoryWorkspaceCreationPropertiesArgs.builder()
                .customSecurityGroupId(aws_security_group.example().id())
                .defaultOu("OU=AWS,DC=Workgroup,DC=Example,DC=com")
                .enableInternetAccess(true)
                .enableMaintenanceMode(true)
                .userEnabledAsLocalAdministrator(true)
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    workspacesDefaultServiceAccess,
                    workspacesDefaultSelfServiceAccess)
                .build());

        var exampleA = new Subnet("exampleA", SubnetArgs.builder()        
            .vpcId(exampleVpc.id())
            .availabilityZone("us-east-1a")
            .cidrBlock("10.0.0.0/24")
            .build());

        var exampleB = new Subnet("exampleB", SubnetArgs.builder()        
            .vpcId(exampleVpc.id())
            .availabilityZone("us-east-1b")
            .cidrBlock("10.0.1.0/24")
            .build());

        var exampleDirectoryservice_directoryDirectory = new Directory("exampleDirectoryservice/directoryDirectory", DirectoryArgs.builder()        
            .name("corp.example.com")
            .password("#S1ncerely")
            .size("Small")
            .vpcSettings(DirectoryVpcSettingsArgs.builder()
                .vpcId(exampleVpc.id())
                .subnetIds(                
                    exampleA.id(),
                    exampleB.id())
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### IP Groups

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleIpGroup = new aws.workspaces.IpGroup("exampleIpGroup", {});
const exampleDirectory = new aws.workspaces.Directory("exampleDirectory", {
    directoryId: aws_directory_service_directory.example.id,
    ipGroupIds: [exampleIpGroup.id],
});
```
```python
import pulumi
import pulumi_aws as aws

example_ip_group = aws.workspaces.IpGroup("exampleIpGroup")
example_directory = aws.workspaces.Directory("exampleDirectory",
    directory_id=aws_directory_service_directory["example"]["id"],
    ip_group_ids=[example_ip_group.id])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleIpGroup = new Aws.Workspaces.IpGroup("exampleIpGroup");

    var exampleDirectory = new Aws.Workspaces.Directory("exampleDirectory", new()
    {
        DirectoryId = aws_directory_service_directory.Example.Id,
        IpGroupIds = new[]
        {
            exampleIpGroup.Id,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/workspaces"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleIpGroup, err := workspaces.NewIpGroup(ctx, "exampleIpGroup", nil)
		if err != nil {
			return err
		}
		_, err = workspaces.NewDirectory(ctx, "exampleDirectory", &workspaces.DirectoryArgs{
			DirectoryId: pulumi.Any(aws_directory_service_directory.Example.Id),
			IpGroupIds: pulumi.StringArray{
				exampleIpGroup.ID(),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.workspaces.IpGroup;
import com.pulumi.aws.workspaces.Directory;
import com.pulumi.aws.workspaces.DirectoryArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var exampleIpGroup = new IpGroup("exampleIpGroup");

        var exampleDirectory = new Directory("exampleDirectory", DirectoryArgs.builder()        
            .directoryId(aws_directory_service_directory.example().id())
            .ipGroupIds(exampleIpGroup.id())
            .build());

    }
}
```
```yaml
resources:
  exampleDirectory:
    type: aws:workspaces:Directory
    properties:
      directoryId: ${aws_directory_service_directory.example.id}
      ipGroupIds:
        - ${exampleIpGroup.id}
  exampleIpGroup:
    type: aws:workspaces:IpGroup
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Workspaces directory using the directory ID. For example:

```sh
 $ pulumi import aws:workspaces/directory:Directory main d-4444444444
```
 