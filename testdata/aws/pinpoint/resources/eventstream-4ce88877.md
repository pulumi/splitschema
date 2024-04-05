Provides a Pinpoint Event Stream resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const app = new aws.pinpoint.App("app", {});
const testStream = new aws.kinesis.Stream("testStream", {shardCount: 1});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["pinpoint.us-east-1.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const testRole = new aws.iam.Role("testRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const stream = new aws.pinpoint.EventStream("stream", {
    applicationId: app.applicationId,
    destinationStreamArn: testStream.arn,
    roleArn: testRole.arn,
});
const testRolePolicyPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: [
            "kinesis:PutRecords",
            "kinesis:DescribeStream",
        ],
        resources: ["arn:aws:kinesis:us-east-1:*:*/*"],
    }],
});
const testRolePolicyRolePolicy = new aws.iam.RolePolicy("testRolePolicyRolePolicy", {
    role: testRole.id,
    policy: testRolePolicyPolicyDocument.then(testRolePolicyPolicyDocument => testRolePolicyPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

app = aws.pinpoint.App("app")
test_stream = aws.kinesis.Stream("testStream", shard_count=1)
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["pinpoint.us-east-1.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
test_role = aws.iam.Role("testRole", assume_role_policy=assume_role.json)
stream = aws.pinpoint.EventStream("stream",
    application_id=app.application_id,
    destination_stream_arn=test_stream.arn,
    role_arn=test_role.arn)
test_role_policy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=[
        "kinesis:PutRecords",
        "kinesis:DescribeStream",
    ],
    resources=["arn:aws:kinesis:us-east-1:*:*/*"],
)])
test_role_policy_role_policy = aws.iam.RolePolicy("testRolePolicyRolePolicy",
    role=test_role.id,
    policy=test_role_policy_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var app = new Aws.Pinpoint.App("app");

    var testStream = new Aws.Kinesis.Stream("testStream", new()
    {
        ShardCount = 1,
    });

    var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                            "pinpoint.us-east-1.amazonaws.com",
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

    var testRole = new Aws.Iam.Role("testRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var stream = new Aws.Pinpoint.EventStream("stream", new()
    {
        ApplicationId = app.ApplicationId,
        DestinationStreamArn = testStream.Arn,
        RoleArn = testRole.Arn,
    });

    var testRolePolicyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "kinesis:PutRecords",
                    "kinesis:DescribeStream",
                },
                Resources = new[]
                {
                    "arn:aws:kinesis:us-east-1:*:*/*",
                },
            },
        },
    });

    var testRolePolicyRolePolicy = new Aws.Iam.RolePolicy("testRolePolicyRolePolicy", new()
    {
        Role = testRole.Id,
        Policy = testRolePolicyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kinesis"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/pinpoint"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		app, err := pinpoint.NewApp(ctx, "app", nil)
		if err != nil {
			return err
		}
		testStream, err := kinesis.NewStream(ctx, "testStream", &kinesis.StreamArgs{
			ShardCount: pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"pinpoint.us-east-1.amazonaws.com",
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
		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = pinpoint.NewEventStream(ctx, "stream", &pinpoint.EventStreamArgs{
			ApplicationId:        app.ApplicationId,
			DestinationStreamArn: testStream.Arn,
			RoleArn:              testRole.Arn,
		})
		if err != nil {
			return err
		}
		testRolePolicyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"kinesis:PutRecords",
						"kinesis:DescribeStream",
					},
					Resources: []string{
						"arn:aws:kinesis:us-east-1:*:*/*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "testRolePolicyRolePolicy", &iam.RolePolicyArgs{
			Role:   testRole.ID(),
			Policy: *pulumi.String(testRolePolicyPolicyDocument.Json),
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
import com.pulumi.aws.pinpoint.App;
import com.pulumi.aws.kinesis.Stream;
import com.pulumi.aws.kinesis.StreamArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.pinpoint.EventStream;
import com.pulumi.aws.pinpoint.EventStreamArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
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
        var app = new App("app");

        var testStream = new Stream("testStream", StreamArgs.builder()        
            .shardCount(1)
            .build());

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("pinpoint.us-east-1.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var testRole = new Role("testRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var stream = new EventStream("stream", EventStreamArgs.builder()        
            .applicationId(app.applicationId())
            .destinationStreamArn(testStream.arn())
            .roleArn(testRole.arn())
            .build());

        final var testRolePolicyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions(                
                    "kinesis:PutRecords",
                    "kinesis:DescribeStream")
                .resources("arn:aws:kinesis:us-east-1:*:*/*")
                .build())
            .build());

        var testRolePolicyRolePolicy = new RolePolicy("testRolePolicyRolePolicy", RolePolicyArgs.builder()        
            .role(testRole.id())
            .policy(testRolePolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  stream:
    type: aws:pinpoint:EventStream
    properties:
      applicationId: ${app.applicationId}
      destinationStreamArn: ${testStream.arn}
      roleArn: ${testRole.arn}
  app:
    type: aws:pinpoint:App
  testStream:
    type: aws:kinesis:Stream
    properties:
      shardCount: 1
  testRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  testRolePolicyRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${testRole.id}
      policy: ${testRolePolicyPolicyDocument.json}
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - pinpoint.us-east-1.amazonaws.com
            actions:
              - sts:AssumeRole
  testRolePolicyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - kinesis:PutRecords
              - kinesis:DescribeStream
            resources:
              - arn:aws:kinesis:us-east-1:*:*/*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Pinpoint Event Stream using the `application-id`. For example:

```sh
 $ pulumi import aws:pinpoint/eventStream:EventStream stream application-id
```
 