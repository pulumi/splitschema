Associates an AppConfig Extension with a Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testTopic = new aws.sns.Topic("testTopic", {});
const testPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["appconfig.amazonaws.com"],
        }],
    }],
});
const testRole = new aws.iam.Role("testRole", {assumeRolePolicy: testPolicyDocument.then(testPolicyDocument => testPolicyDocument.json)});
const testExtension = new aws.appconfig.Extension("testExtension", {
    description: "test description",
    actionPoints: [{
        point: "ON_DEPLOYMENT_COMPLETE",
        actions: [{
            name: "test",
            roleArn: testRole.arn,
            uri: testTopic.arn,
        }],
    }],
    tags: {
        Type: "AppConfig Extension",
    },
});
const testApplication = new aws.appconfig.Application("testApplication", {});
const testExtensionAssociation = new aws.appconfig.ExtensionAssociation("testExtensionAssociation", {
    extensionArn: testExtension.arn,
    resourceArn: testApplication.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

test_topic = aws.sns.Topic("testTopic")
test_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["appconfig.amazonaws.com"],
    )],
)])
test_role = aws.iam.Role("testRole", assume_role_policy=test_policy_document.json)
test_extension = aws.appconfig.Extension("testExtension",
    description="test description",
    action_points=[aws.appconfig.ExtensionActionPointArgs(
        point="ON_DEPLOYMENT_COMPLETE",
        actions=[aws.appconfig.ExtensionActionPointActionArgs(
            name="test",
            role_arn=test_role.arn,
            uri=test_topic.arn,
        )],
    )],
    tags={
        "Type": "AppConfig Extension",
    })
test_application = aws.appconfig.Application("testApplication")
test_extension_association = aws.appconfig.ExtensionAssociation("testExtensionAssociation",
    extension_arn=test_extension.arn,
    resource_arn=test_application.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testTopic = new Aws.Sns.Topic("testTopic");

    var testPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                            "appconfig.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var testRole = new Aws.Iam.Role("testRole", new()
    {
        AssumeRolePolicy = testPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var testExtension = new Aws.AppConfig.Extension("testExtension", new()
    {
        Description = "test description",
        ActionPoints = new[]
        {
            new Aws.AppConfig.Inputs.ExtensionActionPointArgs
            {
                Point = "ON_DEPLOYMENT_COMPLETE",
                Actions = new[]
                {
                    new Aws.AppConfig.Inputs.ExtensionActionPointActionArgs
                    {
                        Name = "test",
                        RoleArn = testRole.Arn,
                        Uri = testTopic.Arn,
                    },
                },
            },
        },
        Tags = 
        {
            { "Type", "AppConfig Extension" },
        },
    });

    var testApplication = new Aws.AppConfig.Application("testApplication");

    var testExtensionAssociation = new Aws.AppConfig.ExtensionAssociation("testExtensionAssociation", new()
    {
        ExtensionArn = testExtension.Arn,
        ResourceArn = testApplication.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appconfig"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testTopic, err := sns.NewTopic(ctx, "testTopic", nil)
		if err != nil {
			return err
		}
		testPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"appconfig.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(testPolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		testExtension, err := appconfig.NewExtension(ctx, "testExtension", &appconfig.ExtensionArgs{
			Description: pulumi.String("test description"),
			ActionPoints: appconfig.ExtensionActionPointArray{
				&appconfig.ExtensionActionPointArgs{
					Point: pulumi.String("ON_DEPLOYMENT_COMPLETE"),
					Actions: appconfig.ExtensionActionPointActionArray{
						&appconfig.ExtensionActionPointActionArgs{
							Name:    pulumi.String("test"),
							RoleArn: testRole.Arn,
							Uri:     testTopic.Arn,
						},
					},
				},
			},
			Tags: pulumi.StringMap{
				"Type": pulumi.String("AppConfig Extension"),
			},
		})
		if err != nil {
			return err
		}
		testApplication, err := appconfig.NewApplication(ctx, "testApplication", nil)
		if err != nil {
			return err
		}
		_, err = appconfig.NewExtensionAssociation(ctx, "testExtensionAssociation", &appconfig.ExtensionAssociationArgs{
			ExtensionArn: testExtension.Arn,
			ResourceArn:  testApplication.Arn,
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
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.appconfig.Extension;
import com.pulumi.aws.appconfig.ExtensionArgs;
import com.pulumi.aws.appconfig.inputs.ExtensionActionPointArgs;
import com.pulumi.aws.appconfig.Application;
import com.pulumi.aws.appconfig.ExtensionAssociation;
import com.pulumi.aws.appconfig.ExtensionAssociationArgs;
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
        var testTopic = new Topic("testTopic");

        final var testPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("appconfig.amazonaws.com")
                    .build())
                .build())
            .build());

        var testRole = new Role("testRole", RoleArgs.builder()        
            .assumeRolePolicy(testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var testExtension = new Extension("testExtension", ExtensionArgs.builder()        
            .description("test description")
            .actionPoints(ExtensionActionPointArgs.builder()
                .point("ON_DEPLOYMENT_COMPLETE")
                .actions(ExtensionActionPointActionArgs.builder()
                    .name("test")
                    .roleArn(testRole.arn())
                    .uri(testTopic.arn())
                    .build())
                .build())
            .tags(Map.of("Type", "AppConfig Extension"))
            .build());

        var testApplication = new Application("testApplication");

        var testExtensionAssociation = new ExtensionAssociation("testExtensionAssociation", ExtensionAssociationArgs.builder()        
            .extensionArn(testExtension.arn())
            .resourceArn(testApplication.arn())
            .build());

    }
}
```
```yaml
resources:
  testTopic:
    type: aws:sns:Topic
  testRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${testPolicyDocument.json}
  testExtension:
    type: aws:appconfig:Extension
    properties:
      description: test description
      actionPoints:
        - point: ON_DEPLOYMENT_COMPLETE
          actions:
            - name: test
              roleArn: ${testRole.arn}
              uri: ${testTopic.arn}
      tags:
        Type: AppConfig Extension
  testApplication:
    type: aws:appconfig:Application
  testExtensionAssociation:
    type: aws:appconfig:ExtensionAssociation
    properties:
      extensionArn: ${testExtension.arn}
      resourceArn: ${testApplication.arn}
variables:
  testPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            principals:
              - type: Service
                identifiers:
                  - appconfig.amazonaws.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AppConfig Extension Associations using their extension association ID. For example:

```sh
 $ pulumi import aws:appconfig/extensionAssociation:ExtensionAssociation example 71rxuzt
```
 