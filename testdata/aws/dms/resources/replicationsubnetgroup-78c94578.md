Provides a DMS (Data Migration Service) replication subnet group resource. DMS replication subnet groups can be created, updated, deleted, and imported.

> **Note:** AWS requires a special IAM role called `dms-vpc-role` when using this resource. See the example below to create it as part of your configuration.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new replication subnet group
const example = new aws.dms.ReplicationSubnetGroup("example", {
    replicationSubnetGroupDescription: "Example replication subnet group",
    replicationSubnetGroupId: "example-dms-replication-subnet-group-tf",
    subnetIds: [
        "subnet-12345678",
        "subnet-12345679",
    ],
    tags: {
        Name: "example",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

# Create a new replication subnet group
example = aws.dms.ReplicationSubnetGroup("example",
    replication_subnet_group_description="Example replication subnet group",
    replication_subnet_group_id="example-dms-replication-subnet-group-tf",
    subnet_ids=[
        "subnet-12345678",
        "subnet-12345679",
    ],
    tags={
        "Name": "example",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Create a new replication subnet group
    var example = new Aws.Dms.ReplicationSubnetGroup("example", new()
    {
        ReplicationSubnetGroupDescription = "Example replication subnet group",
        ReplicationSubnetGroupId = "example-dms-replication-subnet-group-tf",
        SubnetIds = new[]
        {
            "subnet-12345678",
            "subnet-12345679",
        },
        Tags = 
        {
            { "Name", "example" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.NewReplicationSubnetGroup(ctx, "example", &dms.ReplicationSubnetGroupArgs{
			ReplicationSubnetGroupDescription: pulumi.String("Example replication subnet group"),
			ReplicationSubnetGroupId:          pulumi.String("example-dms-replication-subnet-group-tf"),
			SubnetIds: pulumi.StringArray{
				pulumi.String("subnet-12345678"),
				pulumi.String("subnet-12345679"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example"),
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
import com.pulumi.aws.dms.ReplicationSubnetGroup;
import com.pulumi.aws.dms.ReplicationSubnetGroupArgs;
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
        var example = new ReplicationSubnetGroup("example", ReplicationSubnetGroupArgs.builder()        
            .replicationSubnetGroupDescription("Example replication subnet group")
            .replicationSubnetGroupId("example-dms-replication-subnet-group-tf")
            .subnetIds(            
                "subnet-12345678",
                "subnet-12345679")
            .tags(Map.of("Name", "example"))
            .build());

    }
}
```
```yaml
resources:
  # Create a new replication subnet group
  example:
    type: aws:dms:ReplicationSubnetGroup
    properties:
      replicationSubnetGroupDescription: Example replication subnet group
      replicationSubnetGroupId: example-dms-replication-subnet-group-tf
      subnetIds:
        - subnet-12345678
        - subnet-12345679
      tags:
        Name: example
```
{{% /example %}}
{{% example %}}
### Creating special IAM role

If your account does not already include the `dms-vpc-role` IAM role, you will need to create it to allow DMS to manage subnets in the VPC.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dms_vpc_role = new aws.iam.Role("dms-vpc-role", {
    description: "Allows DMS to manage VPC",
    assumeRolePolicy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Principal: {
                Service: "dms.amazonaws.com",
            },
            Action: "sts:AssumeRole",
        }],
    }),
});
const exampleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment", {
    role: dms_vpc_role.name,
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole",
});
const exampleReplicationSubnetGroup = new aws.dms.ReplicationSubnetGroup("exampleReplicationSubnetGroup", {
    replicationSubnetGroupDescription: "Example",
    replicationSubnetGroupId: "example-id",
    subnetIds: [
        "subnet-12345678",
        "subnet-12345679",
    ],
    tags: {
        Name: "example-id",
    },
}, {
    dependsOn: [exampleRolePolicyAttachment],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

dms_vpc_role = aws.iam.Role("dms-vpc-role",
    description="Allows DMS to manage VPC",
    assume_role_policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Principal": {
                "Service": "dms.amazonaws.com",
            },
            "Action": "sts:AssumeRole",
        }],
    }))
example_role_policy_attachment = aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment",
    role=dms_vpc_role.name,
    policy_arn="arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole")
example_replication_subnet_group = aws.dms.ReplicationSubnetGroup("exampleReplicationSubnetGroup",
    replication_subnet_group_description="Example",
    replication_subnet_group_id="example-id",
    subnet_ids=[
        "subnet-12345678",
        "subnet-12345679",
    ],
    tags={
        "Name": "example-id",
    },
    opts=pulumi.ResourceOptions(depends_on=[example_role_policy_attachment]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var dms_vpc_role = new Aws.Iam.Role("dms-vpc-role", new()
    {
        Description = "Allows DMS to manage VPC",
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Effect"] = "Allow",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "dms.amazonaws.com",
                    },
                    ["Action"] = "sts:AssumeRole",
                },
            },
        }),
    });

    var exampleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("exampleRolePolicyAttachment", new()
    {
        Role = dms_vpc_role.Name,
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole",
    });

    var exampleReplicationSubnetGroup = new Aws.Dms.ReplicationSubnetGroup("exampleReplicationSubnetGroup", new()
    {
        ReplicationSubnetGroupDescription = "Example",
        ReplicationSubnetGroupId = "example-id",
        SubnetIds = new[]
        {
            "subnet-12345678",
            "subnet-12345679",
        },
        Tags = 
        {
            { "Name", "example-id" },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleRolePolicyAttachment,
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "dms.amazonaws.com",
					},
					"Action": "sts:AssumeRole",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = iam.NewRole(ctx, "dms-vpc-role", &iam.RoleArgs{
			Description:      pulumi.String("Allows DMS to manage VPC"),
			AssumeRolePolicy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		exampleRolePolicyAttachment, err := iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      dms_vpc_role.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole"),
		})
		if err != nil {
			return err
		}
		_, err = dms.NewReplicationSubnetGroup(ctx, "exampleReplicationSubnetGroup", &dms.ReplicationSubnetGroupArgs{
			ReplicationSubnetGroupDescription: pulumi.String("Example"),
			ReplicationSubnetGroupId:          pulumi.String("example-id"),
			SubnetIds: pulumi.StringArray{
				pulumi.String("subnet-12345678"),
				pulumi.String("subnet-12345679"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("example-id"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleRolePolicyAttachment,
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.dms.ReplicationSubnetGroup;
import com.pulumi.aws.dms.ReplicationSubnetGroupArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var dms_vpc_role = new Role("dms-vpc-role", RoleArgs.builder()        
            .description("Allows DMS to manage VPC")
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "dms.amazonaws.com")
                        )),
                        jsonProperty("Action", "sts:AssumeRole")
                    )))
                )))
            .build());

        var exampleRolePolicyAttachment = new RolePolicyAttachment("exampleRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .role(dms_vpc_role.name())
            .policyArn("arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole")
            .build());

        var exampleReplicationSubnetGroup = new ReplicationSubnetGroup("exampleReplicationSubnetGroup", ReplicationSubnetGroupArgs.builder()        
            .replicationSubnetGroupDescription("Example")
            .replicationSubnetGroupId("example-id")
            .subnetIds(            
                "subnet-12345678",
                "subnet-12345679")
            .tags(Map.of("Name", "example-id"))
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleRolePolicyAttachment)
                .build());

    }
}
```
```yaml
resources:
  dms-vpc-role:
    type: aws:iam:Role
    properties:
      description: Allows DMS to manage VPC
      assumeRolePolicy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Effect: Allow
              Principal:
                Service: dms.amazonaws.com
              Action: sts:AssumeRole
  exampleRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${["dms-vpc-role"].name}
      policyArn: arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole
  exampleReplicationSubnetGroup:
    type: aws:dms:ReplicationSubnetGroup
    properties:
      replicationSubnetGroupDescription: Example
      replicationSubnetGroupId: example-id
      subnetIds:
        - subnet-12345678
        - subnet-12345679
      tags:
        Name: example-id
    options:
      dependson:
        - ${exampleRolePolicyAttachment}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import replication subnet groups using the `replication_subnet_group_id`. For example:

```sh
 $ pulumi import aws:dms/replicationSubnetGroup:ReplicationSubnetGroup test test-dms-replication-subnet-group-tf
```
 