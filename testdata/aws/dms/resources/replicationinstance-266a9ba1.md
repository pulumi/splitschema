Provides a DMS (Data Migration Service) replication instance resource. DMS replication instances can be created, updated, deleted, and imported.

{{% examples %}}
## Example Usage
{{% example %}}

Create required roles and then create a DMS instance, setting the depends_on to the required role policy attachments.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dmsAssumeRole = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            identifiers: ["dms.amazonaws.com"],
            type: "Service",
        }],
    }],
});
const dms_access_for_endpoint = new aws.iam.Role("dms-access-for-endpoint", {assumeRolePolicy: dmsAssumeRole.then(dmsAssumeRole => dmsAssumeRole.json)});
const dms_access_for_endpoint_AmazonDMSRedshiftS3Role = new aws.iam.RolePolicyAttachment("dms-access-for-endpoint-AmazonDMSRedshiftS3Role", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role",
    role: dms_access_for_endpoint.name,
});
const dms_cloudwatch_logs_role = new aws.iam.Role("dms-cloudwatch-logs-role", {assumeRolePolicy: dmsAssumeRole.then(dmsAssumeRole => dmsAssumeRole.json)});
const dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole = new aws.iam.RolePolicyAttachment("dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole",
    role: dms_cloudwatch_logs_role.name,
});
const dms_vpc_role = new aws.iam.Role("dms-vpc-role", {assumeRolePolicy: dmsAssumeRole.then(dmsAssumeRole => dmsAssumeRole.json)});
const dms_vpc_role_AmazonDMSVPCManagementRole = new aws.iam.RolePolicyAttachment("dms-vpc-role-AmazonDMSVPCManagementRole", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole",
    role: dms_vpc_role.name,
});
// Create a new replication instance
const test = new aws.dms.ReplicationInstance("test", {
    allocatedStorage: 20,
    applyImmediately: true,
    autoMinorVersionUpgrade: true,
    availabilityZone: "us-west-2c",
    engineVersion: "3.1.4",
    kmsKeyArn: "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
    multiAz: false,
    preferredMaintenanceWindow: "sun:10:30-sun:14:30",
    publiclyAccessible: true,
    replicationInstanceClass: "dms.t2.micro",
    replicationInstanceId: "test-dms-replication-instance-tf",
    replicationSubnetGroupId: aws_dms_replication_subnet_group["test-dms-replication-subnet-group-tf"].id,
    tags: {
        Name: "test",
    },
    vpcSecurityGroupIds: ["sg-12345678"],
}, {
    dependsOn: [
        dms_access_for_endpoint_AmazonDMSRedshiftS3Role,
        dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole,
        dms_vpc_role_AmazonDMSVPCManagementRole,
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

dms_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        identifiers=["dms.amazonaws.com"],
        type="Service",
    )],
)])
dms_access_for_endpoint = aws.iam.Role("dms-access-for-endpoint", assume_role_policy=dms_assume_role.json)
dms_access_for_endpoint__amazon_dms_redshift_s3_role = aws.iam.RolePolicyAttachment("dms-access-for-endpoint-AmazonDMSRedshiftS3Role",
    policy_arn="arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role",
    role=dms_access_for_endpoint.name)
dms_cloudwatch_logs_role = aws.iam.Role("dms-cloudwatch-logs-role", assume_role_policy=dms_assume_role.json)
dms_cloudwatch_logs_role__amazon_dms_cloud_watch_logs_role = aws.iam.RolePolicyAttachment("dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole",
    policy_arn="arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole",
    role=dms_cloudwatch_logs_role.name)
dms_vpc_role = aws.iam.Role("dms-vpc-role", assume_role_policy=dms_assume_role.json)
dms_vpc_role__amazon_dmsvpc_management_role = aws.iam.RolePolicyAttachment("dms-vpc-role-AmazonDMSVPCManagementRole",
    policy_arn="arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole",
    role=dms_vpc_role.name)
# Create a new replication instance
test = aws.dms.ReplicationInstance("test",
    allocated_storage=20,
    apply_immediately=True,
    auto_minor_version_upgrade=True,
    availability_zone="us-west-2c",
    engine_version="3.1.4",
    kms_key_arn="arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
    multi_az=False,
    preferred_maintenance_window="sun:10:30-sun:14:30",
    publicly_accessible=True,
    replication_instance_class="dms.t2.micro",
    replication_instance_id="test-dms-replication-instance-tf",
    replication_subnet_group_id=aws_dms_replication_subnet_group["test-dms-replication-subnet-group-tf"]["id"],
    tags={
        "Name": "test",
    },
    vpc_security_group_ids=["sg-12345678"],
    opts=pulumi.ResourceOptions(depends_on=[
            dms_access_for_endpoint__amazon_dms_redshift_s3_role,
            dms_cloudwatch_logs_role__amazon_dms_cloud_watch_logs_role,
            dms_vpc_role__amazon_dmsvpc_management_role,
        ]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var dmsAssumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                        Identifiers = new[]
                        {
                            "dms.amazonaws.com",
                        },
                        Type = "Service",
                    },
                },
            },
        },
    });

    var dms_access_for_endpoint = new Aws.Iam.Role("dms-access-for-endpoint", new()
    {
        AssumeRolePolicy = dmsAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var dms_access_for_endpoint_AmazonDMSRedshiftS3Role = new Aws.Iam.RolePolicyAttachment("dms-access-for-endpoint-AmazonDMSRedshiftS3Role", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role",
        Role = dms_access_for_endpoint.Name,
    });

    var dms_cloudwatch_logs_role = new Aws.Iam.Role("dms-cloudwatch-logs-role", new()
    {
        AssumeRolePolicy = dmsAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole = new Aws.Iam.RolePolicyAttachment("dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole",
        Role = dms_cloudwatch_logs_role.Name,
    });

    var dms_vpc_role = new Aws.Iam.Role("dms-vpc-role", new()
    {
        AssumeRolePolicy = dmsAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var dms_vpc_role_AmazonDMSVPCManagementRole = new Aws.Iam.RolePolicyAttachment("dms-vpc-role-AmazonDMSVPCManagementRole", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole",
        Role = dms_vpc_role.Name,
    });

    // Create a new replication instance
    var test = new Aws.Dms.ReplicationInstance("test", new()
    {
        AllocatedStorage = 20,
        ApplyImmediately = true,
        AutoMinorVersionUpgrade = true,
        AvailabilityZone = "us-west-2c",
        EngineVersion = "3.1.4",
        KmsKeyArn = "arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012",
        MultiAz = false,
        PreferredMaintenanceWindow = "sun:10:30-sun:14:30",
        PubliclyAccessible = true,
        ReplicationInstanceClass = "dms.t2.micro",
        ReplicationInstanceId = "test-dms-replication-instance-tf",
        ReplicationSubnetGroupId = aws_dms_replication_subnet_group.Test_dms_replication_subnet_group_tf.Id,
        Tags = 
        {
            { "Name", "test" },
        },
        VpcSecurityGroupIds = new[]
        {
            "sg-12345678",
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            dms_access_for_endpoint_AmazonDMSRedshiftS3Role,
            dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole,
            dms_vpc_role_AmazonDMSVPCManagementRole,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		dmsAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Identifiers: []string{
								"dms.amazonaws.com",
							},
							Type: "Service",
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewRole(ctx, "dms-access-for-endpoint", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(dmsAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "dms-access-for-endpoint-AmazonDMSRedshiftS3Role", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role"),
			Role:      dms_access_for_endpoint.Name,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRole(ctx, "dms-cloudwatch-logs-role", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(dmsAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole"),
			Role:      dms_cloudwatch_logs_role.Name,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRole(ctx, "dms-vpc-role", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(dmsAssumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "dms-vpc-role-AmazonDMSVPCManagementRole", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole"),
			Role:      dms_vpc_role.Name,
		})
		if err != nil {
			return err
		}
		_, err = dms.NewReplicationInstance(ctx, "test", &dms.ReplicationInstanceArgs{
			AllocatedStorage:           pulumi.Int(20),
			ApplyImmediately:           pulumi.Bool(true),
			AutoMinorVersionUpgrade:    pulumi.Bool(true),
			AvailabilityZone:           pulumi.String("us-west-2c"),
			EngineVersion:              pulumi.String("3.1.4"),
			KmsKeyArn:                  pulumi.String("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012"),
			MultiAz:                    pulumi.Bool(false),
			PreferredMaintenanceWindow: pulumi.String("sun:10:30-sun:14:30"),
			PubliclyAccessible:         pulumi.Bool(true),
			ReplicationInstanceClass:   pulumi.String("dms.t2.micro"),
			ReplicationInstanceId:      pulumi.String("test-dms-replication-instance-tf"),
			ReplicationSubnetGroupId:   pulumi.Any(aws_dms_replication_subnet_group.TestDmsReplicationSubnetGroupTf.Id),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("test"),
			},
			VpcSecurityGroupIds: pulumi.StringArray{
				pulumi.String("sg-12345678"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			dms_access_for_endpoint_AmazonDMSRedshiftS3Role,
			dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole,
			dms_vpc_role_AmazonDMSVPCManagementRole,
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
import com.pulumi.aws.dms.ReplicationInstance;
import com.pulumi.aws.dms.ReplicationInstanceArgs;
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
        final var dmsAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .identifiers("dms.amazonaws.com")
                    .type("Service")
                    .build())
                .build())
            .build());

        var dms_access_for_endpoint = new Role("dms-access-for-endpoint", RoleArgs.builder()        
            .assumeRolePolicy(dmsAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var dms_access_for_endpoint_AmazonDMSRedshiftS3Role = new RolePolicyAttachment("dms-access-for-endpoint-AmazonDMSRedshiftS3Role", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role")
            .role(dms_access_for_endpoint.name())
            .build());

        var dms_cloudwatch_logs_role = new Role("dms-cloudwatch-logs-role", RoleArgs.builder()        
            .assumeRolePolicy(dmsAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole = new RolePolicyAttachment("dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole")
            .role(dms_cloudwatch_logs_role.name())
            .build());

        var dms_vpc_role = new Role("dms-vpc-role", RoleArgs.builder()        
            .assumeRolePolicy(dmsAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var dms_vpc_role_AmazonDMSVPCManagementRole = new RolePolicyAttachment("dms-vpc-role-AmazonDMSVPCManagementRole", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole")
            .role(dms_vpc_role.name())
            .build());

        var test = new ReplicationInstance("test", ReplicationInstanceArgs.builder()        
            .allocatedStorage(20)
            .applyImmediately(true)
            .autoMinorVersionUpgrade(true)
            .availabilityZone("us-west-2c")
            .engineVersion("3.1.4")
            .kmsKeyArn("arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012")
            .multiAz(false)
            .preferredMaintenanceWindow("sun:10:30-sun:14:30")
            .publiclyAccessible(true)
            .replicationInstanceClass("dms.t2.micro")
            .replicationInstanceId("test-dms-replication-instance-tf")
            .replicationSubnetGroupId(aws_dms_replication_subnet_group.test-dms-replication-subnet-group-tf().id())
            .tags(Map.of("Name", "test"))
            .vpcSecurityGroupIds("sg-12345678")
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    dms_access_for_endpoint_AmazonDMSRedshiftS3Role,
                    dms_cloudwatch_logs_role_AmazonDMSCloudWatchLogsRole,
                    dms_vpc_role_AmazonDMSVPCManagementRole)
                .build());

    }
}
```
```yaml
resources:
  dms-access-for-endpoint:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${dmsAssumeRole.json}
  dms-access-for-endpoint-AmazonDMSRedshiftS3Role:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/service-role/AmazonDMSRedshiftS3Role
      role: ${["dms-access-for-endpoint"].name}
  dms-cloudwatch-logs-role:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${dmsAssumeRole.json}
  dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/service-role/AmazonDMSCloudWatchLogsRole
      role: ${["dms-cloudwatch-logs-role"].name}
  dms-vpc-role:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${dmsAssumeRole.json}
  dms-vpc-role-AmazonDMSVPCManagementRole:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole
      role: ${["dms-vpc-role"].name}
  # Create a new replication instance
  test:
    type: aws:dms:ReplicationInstance
    properties:
      allocatedStorage: 20
      applyImmediately: true
      autoMinorVersionUpgrade: true
      availabilityZone: us-west-2c
      engineVersion: 3.1.4
      kmsKeyArn: arn:aws:kms:us-east-1:123456789012:key/12345678-1234-1234-1234-123456789012
      multiAz: false
      preferredMaintenanceWindow: sun:10:30-sun:14:30
      publiclyAccessible: true
      replicationInstanceClass: dms.t2.micro
      replicationInstanceId: test-dms-replication-instance-tf
      replicationSubnetGroupId: ${aws_dms_replication_subnet_group"test-dms-replication-subnet-group-tf"[%!s(MISSING)].id}
      tags:
        Name: test
      vpcSecurityGroupIds:
        - sg-12345678
    options:
      dependson:
        - ${["dms-access-for-endpoint-AmazonDMSRedshiftS3Role"]}
        - ${["dms-cloudwatch-logs-role-AmazonDMSCloudWatchLogsRole"]}
        - ${["dms-vpc-role-AmazonDMSVPCManagementRole"]}
variables:
  dmsAssumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            principals:
              - identifiers:
                  - dms.amazonaws.com
                type: Service
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import replication instances using the `replication_instance_id`. For example:

```sh
 $ pulumi import aws:dms/replicationInstance:ReplicationInstance test test-dms-replication-instance-tf
```
 