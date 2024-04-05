Resource for managing an AWS QuickSight VPC Connection.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const vpcConnectionRole = new aws.iam.Role("vpcConnectionRole", {
    assumeRolePolicy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Action: "sts:AssumeRole",
            Principal: {
                Service: "quicksight.amazonaws.com",
            },
        }],
    }),
    inlinePolicies: [{
        name: "QuickSightVPCConnectionRolePolicy",
        policy: JSON.stringify({
            Version: "2012-10-17",
            Statement: [{
                Effect: "Allow",
                Action: [
                    "ec2:CreateNetworkInterface",
                    "ec2:ModifyNetworkInterfaceAttribute",
                    "ec2:DeleteNetworkInterface",
                    "ec2:DescribeSubnets",
                    "ec2:DescribeSecurityGroups",
                ],
                Resource: ["*"],
            }],
        }),
    }],
});
const example = new aws.quicksight.VpcConnection("example", {
    vpcConnectionId: "example-connection-id",
    roleArn: vpcConnectionRole.arn,
    securityGroupIds: ["sg-00000000000000000"],
    subnetIds: [
        "subnet-00000000000000000",
        "subnet-00000000000000001",
    ],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

vpc_connection_role = aws.iam.Role("vpcConnectionRole",
    assume_role_policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Action": "sts:AssumeRole",
            "Principal": {
                "Service": "quicksight.amazonaws.com",
            },
        }],
    }),
    inline_policies=[aws.iam.RoleInlinePolicyArgs(
        name="QuickSightVPCConnectionRolePolicy",
        policy=json.dumps({
            "Version": "2012-10-17",
            "Statement": [{
                "Effect": "Allow",
                "Action": [
                    "ec2:CreateNetworkInterface",
                    "ec2:ModifyNetworkInterfaceAttribute",
                    "ec2:DeleteNetworkInterface",
                    "ec2:DescribeSubnets",
                    "ec2:DescribeSecurityGroups",
                ],
                "Resource": ["*"],
            }],
        }),
    )])
example = aws.quicksight.VpcConnection("example",
    vpc_connection_id="example-connection-id",
    role_arn=vpc_connection_role.arn,
    security_group_ids=["sg-00000000000000000"],
    subnet_ids=[
        "subnet-00000000000000000",
        "subnet-00000000000000001",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var vpcConnectionRole = new Aws.Iam.Role("vpcConnectionRole", new()
    {
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Effect"] = "Allow",
                    ["Action"] = "sts:AssumeRole",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "quicksight.amazonaws.com",
                    },
                },
            },
        }),
        InlinePolicies = new[]
        {
            new Aws.Iam.Inputs.RoleInlinePolicyArgs
            {
                Name = "QuickSightVPCConnectionRolePolicy",
                Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
                {
                    ["Version"] = "2012-10-17",
                    ["Statement"] = new[]
                    {
                        new Dictionary<string, object?>
                        {
                            ["Effect"] = "Allow",
                            ["Action"] = new[]
                            {
                                "ec2:CreateNetworkInterface",
                                "ec2:ModifyNetworkInterfaceAttribute",
                                "ec2:DeleteNetworkInterface",
                                "ec2:DescribeSubnets",
                                "ec2:DescribeSecurityGroups",
                            },
                            ["Resource"] = new[]
                            {
                                "*",
                            },
                        },
                    },
                }),
            },
        },
    });

    var example = new Aws.Quicksight.VpcConnection("example", new()
    {
        VpcConnectionId = "example-connection-id",
        RoleArn = vpcConnectionRole.Arn,
        SecurityGroupIds = new[]
        {
            "sg-00000000000000000",
        },
        SubnetIds = new[]
        {
            "subnet-00000000000000000",
            "subnet-00000000000000001",
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Effect": "Allow",
					"Action": "sts:AssumeRole",
					"Principal": map[string]interface{}{
						"Service": "quicksight.amazonaws.com",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		tmpJSON1, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Effect": "Allow",
					"Action": []string{
						"ec2:CreateNetworkInterface",
						"ec2:ModifyNetworkInterfaceAttribute",
						"ec2:DeleteNetworkInterface",
						"ec2:DescribeSubnets",
						"ec2:DescribeSecurityGroups",
					},
					"Resource": []string{
						"*",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json1 := string(tmpJSON1)
		vpcConnectionRole, err := iam.NewRole(ctx, "vpcConnectionRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(json0),
			InlinePolicies: iam.RoleInlinePolicyArray{
				&iam.RoleInlinePolicyArgs{
					Name:   pulumi.String("QuickSightVPCConnectionRolePolicy"),
					Policy: pulumi.String(json1),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = quicksight.NewVpcConnection(ctx, "example", &quicksight.VpcConnectionArgs{
			VpcConnectionId: pulumi.String("example-connection-id"),
			RoleArn:         vpcConnectionRole.Arn,
			SecurityGroupIds: pulumi.StringArray{
				pulumi.String("sg-00000000000000000"),
			},
			SubnetIds: pulumi.StringArray{
				pulumi.String("subnet-00000000000000000"),
				pulumi.String("subnet-00000000000000001"),
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.inputs.RoleInlinePolicyArgs;
import com.pulumi.aws.quicksight.VpcConnection;
import com.pulumi.aws.quicksight.VpcConnectionArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var vpcConnectionRole = new Role("vpcConnectionRole", RoleArgs.builder()        
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Action", "sts:AssumeRole"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "quicksight.amazonaws.com")
                        ))
                    )))
                )))
            .inlinePolicies(RoleInlinePolicyArgs.builder()
                .name("QuickSightVPCConnectionRolePolicy")
                .policy(serializeJson(
                    jsonObject(
                        jsonProperty("Version", "2012-10-17"),
                        jsonProperty("Statement", jsonArray(jsonObject(
                            jsonProperty("Effect", "Allow"),
                            jsonProperty("Action", jsonArray(
                                "ec2:CreateNetworkInterface", 
                                "ec2:ModifyNetworkInterfaceAttribute", 
                                "ec2:DeleteNetworkInterface", 
                                "ec2:DescribeSubnets", 
                                "ec2:DescribeSecurityGroups"
                            )),
                            jsonProperty("Resource", jsonArray("*"))
                        )))
                    )))
                .build())
            .build());

        var example = new VpcConnection("example", VpcConnectionArgs.builder()        
            .vpcConnectionId("example-connection-id")
            .roleArn(vpcConnectionRole.arn())
            .securityGroupIds("sg-00000000000000000")
            .subnetIds(            
                "subnet-00000000000000000",
                "subnet-00000000000000001")
            .build());

    }
}
```
```yaml
resources:
  vpcConnectionRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Effect: Allow
              Action: sts:AssumeRole
              Principal:
                Service: quicksight.amazonaws.com
      inlinePolicies:
        - name: QuickSightVPCConnectionRolePolicy
          policy:
            fn::toJSON:
              Version: 2012-10-17
              Statement:
                - Effect: Allow
                  Action:
                    - ec2:CreateNetworkInterface
                    - ec2:ModifyNetworkInterfaceAttribute
                    - ec2:DeleteNetworkInterface
                    - ec2:DescribeSubnets
                    - ec2:DescribeSecurityGroups
                  Resource:
                    - '*'
  example:
    type: aws:quicksight:VpcConnection
    properties:
      vpcConnectionId: example-connection-id
      roleArn: ${vpcConnectionRole.arn}
      securityGroupIds:
        - sg-00000000000000000
      subnetIds:
        - subnet-00000000000000000
        - subnet-00000000000000001
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import QuickSight VPC connection using the AWS account ID and VPC connection ID separated by commas (`,`). For example:

```sh
 $ pulumi import aws:quicksight/vpcConnection:VpcConnection example 123456789012,example
```
 