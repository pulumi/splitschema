Creates a AWS Batch compute environment. Compute environments contain the Amazon ECS container instances that are used to run containerized batch jobs.

For information about AWS Batch, see [What is AWS Batch?](http://docs.aws.amazon.com/batch/latest/userguide/what-is-batch.html) .
For information about compute environment, see [Compute Environments](http://docs.aws.amazon.com/batch/latest/userguide/compute_environments.html) .

> **Note:** To prevent a race condition during environment deletion, make sure to set `depends_on` to the related `aws.iam.RolePolicyAttachment`;
otherwise, the policy may be destroyed too soon and the compute environment will then get stuck in the `DELETING` state, see [Troubleshooting AWS Batch](http://docs.aws.amazon.com/batch/latest/userguide/troubleshooting.html) .

{{% examples %}}
## Example Usage
{{% example %}}
### EC2 Type

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ec2AssumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["ec2.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const ecsInstanceRoleRole = new aws.iam.Role("ecsInstanceRoleRole", {assumeRolePolicy: ec2AssumeRole.then(ec2AssumeRole => ec2AssumeRole.json)});
const ecsInstanceRoleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("ecsInstanceRoleRolePolicyAttachment", {
    role: ecsInstanceRoleRole.name,
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role",
});
const ecsInstanceRoleInstanceProfile = new aws.iam.InstanceProfile("ecsInstanceRoleInstanceProfile", {role: ecsInstanceRoleRole.name});
const batchAssumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["batch.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const awsBatchServiceRoleRole = new aws.iam.Role("awsBatchServiceRoleRole", {assumeRolePolicy: batchAssumeRole.then(batchAssumeRole => batchAssumeRole.json)});
const awsBatchServiceRoleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("awsBatchServiceRoleRolePolicyAttachment", {
    role: awsBatchServiceRoleRole.name,
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole",
});
const sampleSecurityGroup = new aws.ec2.SecurityGroup("sampleSecurityGroup", {egress: [{
    fromPort: 0,
    toPort: 0,
    protocol: "-1",
    cidrBlocks: ["0.0.0.0/0"],
}]});
const sampleVpc = new aws.ec2.Vpc("sampleVpc", {cidrBlock: "10.1.0.0/16"});
const sampleSubnet = new aws.ec2.Subnet("sampleSubnet", {
    vpcId: sampleVpc.id,
    cidrBlock: "10.1.1.0/24",
});
const samplePlacementGroup = new aws.ec2.PlacementGroup("samplePlacementGroup", {strategy: "cluster"});
const sampleComputeEnvironment = new aws.batch.ComputeEnvironment("sampleComputeEnvironment", {
    computeEnvironmentName: "sample",
    computeResources: {
        instanceRole: ecsInstanceRoleInstanceProfile.arn,
        instanceTypes: ["c4.large"],
        maxVcpus: 16,
        minVcpus: 0,
        placementGroup: samplePlacementGroup.name,
        securityGroupIds: [sampleSecurityGroup.id],
        subnets: [sampleSubnet.id],
        type: "EC2",
    },
    serviceRole: awsBatchServiceRoleRole.arn,
    type: "MANAGED",
}, {
    dependsOn: [awsBatchServiceRoleRolePolicyAttachment],
});
```
```python
import pulumi
import pulumi_aws as aws

ec2_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["ec2.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
ecs_instance_role_role = aws.iam.Role("ecsInstanceRoleRole", assume_role_policy=ec2_assume_role.json)
ecs_instance_role_role_policy_attachment = aws.iam.RolePolicyAttachment("ecsInstanceRoleRolePolicyAttachment",
    role=ecs_instance_role_role.name,
    policy_arn="arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role")
ecs_instance_role_instance_profile = aws.iam.InstanceProfile("ecsInstanceRoleInstanceProfile", role=ecs_instance_role_role.name)
batch_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["batch.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
aws_batch_service_role_role = aws.iam.Role("awsBatchServiceRoleRole", assume_role_policy=batch_assume_role.json)
aws_batch_service_role_role_policy_attachment = aws.iam.RolePolicyAttachment("awsBatchServiceRoleRolePolicyAttachment",
    role=aws_batch_service_role_role.name,
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole")
sample_security_group = aws.ec2.SecurityGroup("sampleSecurityGroup", egress=[aws.ec2.SecurityGroupEgressArgs(
    from_port=0,
    to_port=0,
    protocol="-1",
    cidr_blocks=["0.0.0.0/0"],
)])
sample_vpc = aws.ec2.Vpc("sampleVpc", cidr_block="10.1.0.0/16")
sample_subnet = aws.ec2.Subnet("sampleSubnet",
    vpc_id=sample_vpc.id,
    cidr_block="10.1.1.0/24")
sample_placement_group = aws.ec2.PlacementGroup("samplePlacementGroup", strategy="cluster")
sample_compute_environment = aws.batch.ComputeEnvironment("sampleComputeEnvironment",
    compute_environment_name="sample",
    compute_resources=aws.batch.ComputeEnvironmentComputeResourcesArgs(
        instance_role=ecs_instance_role_instance_profile.arn,
        instance_types=["c4.large"],
        max_vcpus=16,
        min_vcpus=0,
        placement_group=sample_placement_group.name,
        security_group_ids=[sample_security_group.id],
        subnets=[sample_subnet.id],
        type="EC2",
    ),
    service_role=aws_batch_service_role_role.arn,
    type="MANAGED",
    opts=pulumi.ResourceOptions(depends_on=[aws_batch_service_role_role_policy_attachment]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ec2AssumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "ec2.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var ecsInstanceRoleRole = new Aws.Iam.Role("ecsInstanceRoleRole", new()
    {
        AssumeRolePolicy = ec2AssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var ecsInstanceRoleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("ecsInstanceRoleRolePolicyAttachment", new()
    {
        Role = ecsInstanceRoleRole.Name,
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role",
    });

    var ecsInstanceRoleInstanceProfile = new Aws.Iam.InstanceProfile("ecsInstanceRoleInstanceProfile", new()
    {
        Role = ecsInstanceRoleRole.Name,
    });

    var batchAssumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "batch.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var awsBatchServiceRoleRole = new Aws.Iam.Role("awsBatchServiceRoleRole", new()
    {
        AssumeRolePolicy = batchAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var awsBatchServiceRoleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("awsBatchServiceRoleRolePolicyAttachment", new()
    {
        Role = awsBatchServiceRoleRole.Name,
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole",
    });

    var sampleSecurityGroup = new Aws.Ec2.SecurityGroup("sampleSecurityGroup", new()
    {
        Egress = new[]
        {
            new Aws.Ec2.Inputs.SecurityGroupEgressArgs
            {
                FromPort = 0,
                ToPort = 0,
                Protocol = "-1",
                CidrBlocks = new[]
                {
                    "0.0.0.0/0",
                },
            },
        },
    });

    var sampleVpc = new Aws.Ec2.Vpc("sampleVpc", new()
    {
        CidrBlock = "10.1.0.0/16",
    });

    var sampleSubnet = new Aws.Ec2.Subnet("sampleSubnet", new()
    {
        VpcId = sampleVpc.Id,
        CidrBlock = "10.1.1.0/24",
    });

    var samplePlacementGroup = new Aws.Ec2.PlacementGroup("samplePlacementGroup", new()
    {
        Strategy = "cluster",
    });

    var sampleComputeEnvironment = new Aws.Batch.ComputeEnvironment("sampleComputeEnvironment", new()
    {
        ComputeEnvironmentName = "sample",
        ComputeResources = new Aws.Batch.Inputs.ComputeEnvironmentComputeResourcesArgs
        {
            InstanceRole = ecsInstanceRoleInstanceProfile.Arn,
            InstanceTypes = new[]
            {
                "c4.large",
            },
            MaxVcpus = 16,
            MinVcpus = 0,
            PlacementGroup = samplePlacementGroup.Name,
            SecurityGroupIds = new[]
            {
                sampleSecurityGroup.Id,
            },
            Subnets = new[]
            {
                sampleSubnet.Id,
            },
            Type = "EC2",
        },
        ServiceRole = awsBatchServiceRoleRole.Arn,
        Type = "MANAGED",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            awsBatchServiceRoleRolePolicyAttachment,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/batch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ec2AssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"ec2.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		ecsInstanceRoleRole, err := iam.NewRole(ctx, "ecsInstanceRoleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(ec2AssumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "ecsInstanceRoleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      ecsInstanceRoleRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"),
		})
		if err != nil {
			return err
		}
		ecsInstanceRoleInstanceProfile, err := iam.NewInstanceProfile(ctx, "ecsInstanceRoleInstanceProfile", &iam.InstanceProfileArgs{
			Role: ecsInstanceRoleRole.Name,
		})
		if err != nil {
			return err
		}
		batchAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"batch.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		awsBatchServiceRoleRole, err := iam.NewRole(ctx, "awsBatchServiceRoleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(batchAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		awsBatchServiceRoleRolePolicyAttachment, err := iam.NewRolePolicyAttachment(ctx, "awsBatchServiceRoleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      awsBatchServiceRoleRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole"),
		})
		if err != nil {
			return err
		}
		sampleSecurityGroup, err := ec2.NewSecurityGroup(ctx, "sampleSecurityGroup", &ec2.SecurityGroupArgs{
			Egress: ec2.SecurityGroupEgressArray{
				&ec2.SecurityGroupEgressArgs{
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
					Protocol: pulumi.String("-1"),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		sampleVpc, err := ec2.NewVpc(ctx, "sampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.1.0.0/16"),
		})
		if err != nil {
			return err
		}
		sampleSubnet, err := ec2.NewSubnet(ctx, "sampleSubnet", &ec2.SubnetArgs{
			VpcId:     sampleVpc.ID(),
			CidrBlock: pulumi.String("10.1.1.0/24"),
		})
		if err != nil {
			return err
		}
		samplePlacementGroup, err := ec2.NewPlacementGroup(ctx, "samplePlacementGroup", &ec2.PlacementGroupArgs{
			Strategy: pulumi.String("cluster"),
		})
		if err != nil {
			return err
		}
		_, err = batch.NewComputeEnvironment(ctx, "sampleComputeEnvironment", &batch.ComputeEnvironmentArgs{
			ComputeEnvironmentName: pulumi.String("sample"),
			ComputeResources: &batch.ComputeEnvironmentComputeResourcesArgs{
				InstanceRole: ecsInstanceRoleInstanceProfile.Arn,
				InstanceTypes: pulumi.StringArray{
					pulumi.String("c4.large"),
				},
				MaxVcpus:       pulumi.Int(16),
				MinVcpus:       pulumi.Int(0),
				PlacementGroup: samplePlacementGroup.Name,
				SecurityGroupIds: pulumi.StringArray{
					sampleSecurityGroup.ID(),
				},
				Subnets: pulumi.StringArray{
					sampleSubnet.ID(),
				},
				Type: pulumi.String("EC2"),
			},
			ServiceRole: awsBatchServiceRoleRole.Arn,
			Type:        pulumi.String("MANAGED"),
		}, pulumi.DependsOn([]pulumi.Resource{
			awsBatchServiceRoleRolePolicyAttachment,
		}))
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
import com.pulumi.aws.iam.InstanceProfile;
import com.pulumi.aws.iam.InstanceProfileArgs;
import com.pulumi.aws.ec2.SecurityGroup;
import com.pulumi.aws.ec2.SecurityGroupArgs;
import com.pulumi.aws.ec2.inputs.SecurityGroupEgressArgs;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.ec2.PlacementGroup;
import com.pulumi.aws.ec2.PlacementGroupArgs;
import com.pulumi.aws.batch.ComputeEnvironment;
import com.pulumi.aws.batch.ComputeEnvironmentArgs;
import com.pulumi.aws.batch.inputs.ComputeEnvironmentComputeResourcesArgs;
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
        final var ec2AssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("ec2.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var ecsInstanceRoleRole = new Role("ecsInstanceRoleRole", RoleArgs.builder()        
            .assumeRolePolicy(ec2AssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var ecsInstanceRoleRolePolicyAttachment = new RolePolicyAttachment("ecsInstanceRoleRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .role(ecsInstanceRoleRole.name())
            .policyArn("arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role")
            .build());

        var ecsInstanceRoleInstanceProfile = new InstanceProfile("ecsInstanceRoleInstanceProfile", InstanceProfileArgs.builder()        
            .role(ecsInstanceRoleRole.name())
            .build());

        final var batchAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("batch.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var awsBatchServiceRoleRole = new Role("awsBatchServiceRoleRole", RoleArgs.builder()        
            .assumeRolePolicy(batchAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var awsBatchServiceRoleRolePolicyAttachment = new RolePolicyAttachment("awsBatchServiceRoleRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .role(awsBatchServiceRoleRole.name())
            .policyArn("arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole")
            .build());

        var sampleSecurityGroup = new SecurityGroup("sampleSecurityGroup", SecurityGroupArgs.builder()        
            .egress(SecurityGroupEgressArgs.builder()
                .fromPort(0)
                .toPort(0)
                .protocol("-1")
                .cidrBlocks("0.0.0.0/0")
                .build())
            .build());

        var sampleVpc = new Vpc("sampleVpc", VpcArgs.builder()        
            .cidrBlock("10.1.0.0/16")
            .build());

        var sampleSubnet = new Subnet("sampleSubnet", SubnetArgs.builder()        
            .vpcId(sampleVpc.id())
            .cidrBlock("10.1.1.0/24")
            .build());

        var samplePlacementGroup = new PlacementGroup("samplePlacementGroup", PlacementGroupArgs.builder()        
            .strategy("cluster")
            .build());

        var sampleComputeEnvironment = new ComputeEnvironment("sampleComputeEnvironment", ComputeEnvironmentArgs.builder()        
            .computeEnvironmentName("sample")
            .computeResources(ComputeEnvironmentComputeResourcesArgs.builder()
                .instanceRole(ecsInstanceRoleInstanceProfile.arn())
                .instanceTypes("c4.large")
                .maxVcpus(16)
                .minVcpus(0)
                .placementGroup(samplePlacementGroup.name())
                .securityGroupIds(sampleSecurityGroup.id())
                .subnets(sampleSubnet.id())
                .type("EC2")
                .build())
            .serviceRole(awsBatchServiceRoleRole.arn())
            .type("MANAGED")
            .build(), CustomResourceOptions.builder()
                .dependsOn(awsBatchServiceRoleRolePolicyAttachment)
                .build());

    }
}
```
```yaml
resources:
  ecsInstanceRoleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${ec2AssumeRole.json}
  ecsInstanceRoleRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${ecsInstanceRoleRole.name}
      policyArn: arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role
  ecsInstanceRoleInstanceProfile:
    type: aws:iam:InstanceProfile
    properties:
      role: ${ecsInstanceRoleRole.name}
  awsBatchServiceRoleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${batchAssumeRole.json}
  awsBatchServiceRoleRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${awsBatchServiceRoleRole.name}
      policyArn: arn:aws:iam::aws:policy/service-role/AWSBatchServiceRole
  sampleSecurityGroup:
    type: aws:ec2:SecurityGroup
    properties:
      egress:
        - fromPort: 0
          toPort: 0
          protocol: '-1'
          cidrBlocks:
            - 0.0.0.0/0
  sampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.1.0.0/16
  sampleSubnet:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${sampleVpc.id}
      cidrBlock: 10.1.1.0/24
  samplePlacementGroup:
    type: aws:ec2:PlacementGroup
    properties:
      strategy: cluster
  sampleComputeEnvironment:
    type: aws:batch:ComputeEnvironment
    properties:
      computeEnvironmentName: sample
      computeResources:
        instanceRole: ${ecsInstanceRoleInstanceProfile.arn}
        instanceTypes:
          - c4.large
        maxVcpus: 16
        minVcpus: 0
        placementGroup: ${samplePlacementGroup.name}
        securityGroupIds:
          - ${sampleSecurityGroup.id}
        subnets:
          - ${sampleSubnet.id}
        type: EC2
      serviceRole: ${awsBatchServiceRoleRole.arn}
      type: MANAGED
    options:
      dependson:
        - ${awsBatchServiceRoleRolePolicyAttachment}
variables:
  ec2AssumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - ec2.amazonaws.com
            actions:
              - sts:AssumeRole
  batchAssumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - batch.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% example %}}
### Fargate Type

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sample = new aws.batch.ComputeEnvironment("sample", {
    computeEnvironmentName: "sample",
    computeResources: {
        maxVcpus: 16,
        securityGroupIds: [aws_security_group.sample.id],
        subnets: [aws_subnet.sample.id],
        type: "FARGATE",
    },
    serviceRole: aws_iam_role.aws_batch_service_role.arn,
    type: "MANAGED",
}, {
    dependsOn: [aws_iam_role_policy_attachment.aws_batch_service_role],
});
```
```python
import pulumi
import pulumi_aws as aws

sample = aws.batch.ComputeEnvironment("sample",
    compute_environment_name="sample",
    compute_resources=aws.batch.ComputeEnvironmentComputeResourcesArgs(
        max_vcpus=16,
        security_group_ids=[aws_security_group["sample"]["id"]],
        subnets=[aws_subnet["sample"]["id"]],
        type="FARGATE",
    ),
    service_role=aws_iam_role["aws_batch_service_role"]["arn"],
    type="MANAGED",
    opts=pulumi.ResourceOptions(depends_on=[aws_iam_role_policy_attachment["aws_batch_service_role"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var sample = new Aws.Batch.ComputeEnvironment("sample", new()
    {
        ComputeEnvironmentName = "sample",
        ComputeResources = new Aws.Batch.Inputs.ComputeEnvironmentComputeResourcesArgs
        {
            MaxVcpus = 16,
            SecurityGroupIds = new[]
            {
                aws_security_group.Sample.Id,
            },
            Subnets = new[]
            {
                aws_subnet.Sample.Id,
            },
            Type = "FARGATE",
        },
        ServiceRole = aws_iam_role.Aws_batch_service_role.Arn,
        Type = "MANAGED",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_iam_role_policy_attachment.Aws_batch_service_role,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/batch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := batch.NewComputeEnvironment(ctx, "sample", &batch.ComputeEnvironmentArgs{
			ComputeEnvironmentName: pulumi.String("sample"),
			ComputeResources: &batch.ComputeEnvironmentComputeResourcesArgs{
				MaxVcpus: pulumi.Int(16),
				SecurityGroupIds: pulumi.StringArray{
					aws_security_group.Sample.Id,
				},
				Subnets: pulumi.StringArray{
					aws_subnet.Sample.Id,
				},
				Type: pulumi.String("FARGATE"),
			},
			ServiceRole: pulumi.Any(aws_iam_role.Aws_batch_service_role.Arn),
			Type:        pulumi.String("MANAGED"),
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_iam_role_policy_attachment.Aws_batch_service_role,
		}))
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
import com.pulumi.aws.batch.ComputeEnvironment;
import com.pulumi.aws.batch.ComputeEnvironmentArgs;
import com.pulumi.aws.batch.inputs.ComputeEnvironmentComputeResourcesArgs;
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
        var sample = new ComputeEnvironment("sample", ComputeEnvironmentArgs.builder()        
            .computeEnvironmentName("sample")
            .computeResources(ComputeEnvironmentComputeResourcesArgs.builder()
                .maxVcpus(16)
                .securityGroupIds(aws_security_group.sample().id())
                .subnets(aws_subnet.sample().id())
                .type("FARGATE")
                .build())
            .serviceRole(aws_iam_role.aws_batch_service_role().arn())
            .type("MANAGED")
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_iam_role_policy_attachment.aws_batch_service_role())
                .build());

    }
}
```
```yaml
resources:
  sample:
    type: aws:batch:ComputeEnvironment
    properties:
      computeEnvironmentName: sample
      computeResources:
        maxVcpus: 16
        securityGroupIds:
          - ${aws_security_group.sample.id}
        subnets:
          - ${aws_subnet.sample.id}
        type: FARGATE
      serviceRole: ${aws_iam_role.aws_batch_service_role.arn}
      type: MANAGED
    options:
      dependson:
        - ${aws_iam_role_policy_attachment.aws_batch_service_role}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS Batch compute using the `compute_environment_name`. For example:

```sh
 $ pulumi import aws:batch/computeEnvironment:ComputeEnvironment sample sample
```
 