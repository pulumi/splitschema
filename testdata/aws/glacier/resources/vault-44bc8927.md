Provides a Glacier Vault Resource. You can refer to the [Glacier Developer Guide](https://docs.aws.amazon.com/amazonglacier/latest/dev/working-with-vaults.html) for a full explanation of the Glacier Vault functionality

> **NOTE:** When removing a Glacier Vault, the Vault must be empty.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsSnsTopic = new aws.sns.Topic("awsSnsTopic", {});
const myArchivePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        sid: "add-read-only-perm",
        effect: "Allow",
        principals: [{
            type: "*",
            identifiers: ["*"],
        }],
        actions: [
            "glacier:InitiateJob",
            "glacier:GetJobOutput",
        ],
        resources: ["arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive"],
    }],
});
const myArchiveVault = new aws.glacier.Vault("myArchiveVault", {
    notification: {
        snsTopic: awsSnsTopic.arn,
        events: [
            "ArchiveRetrievalCompleted",
            "InventoryRetrievalCompleted",
        ],
    },
    accessPolicy: myArchivePolicyDocument.then(myArchivePolicyDocument => myArchivePolicyDocument.json),
    tags: {
        Test: "MyArchive",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

aws_sns_topic = aws.sns.Topic("awsSnsTopic")
my_archive_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="add-read-only-perm",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="*",
        identifiers=["*"],
    )],
    actions=[
        "glacier:InitiateJob",
        "glacier:GetJobOutput",
    ],
    resources=["arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive"],
)])
my_archive_vault = aws.glacier.Vault("myArchiveVault",
    notification=aws.glacier.VaultNotificationArgs(
        sns_topic=aws_sns_topic.arn,
        events=[
            "ArchiveRetrievalCompleted",
            "InventoryRetrievalCompleted",
        ],
    ),
    access_policy=my_archive_policy_document.json,
    tags={
        "Test": "MyArchive",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var awsSnsTopic = new Aws.Sns.Topic("awsSnsTopic");

    var myArchivePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "add-read-only-perm",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "*",
                        Identifiers = new[]
                        {
                            "*",
                        },
                    },
                },
                Actions = new[]
                {
                    "glacier:InitiateJob",
                    "glacier:GetJobOutput",
                },
                Resources = new[]
                {
                    "arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive",
                },
            },
        },
    });

    var myArchiveVault = new Aws.Glacier.Vault("myArchiveVault", new()
    {
        Notification = new Aws.Glacier.Inputs.VaultNotificationArgs
        {
            SnsTopic = awsSnsTopic.Arn,
            Events = new[]
            {
                "ArchiveRetrievalCompleted",
                "InventoryRetrievalCompleted",
            },
        },
        AccessPolicy = myArchivePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
        Tags = 
        {
            { "Test", "MyArchive" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glacier"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		awsSnsTopic, err := sns.NewTopic(ctx, "awsSnsTopic", nil)
		if err != nil {
			return err
		}
		myArchivePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("add-read-only-perm"),
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "*",
							Identifiers: []string{
								"*",
							},
						},
					},
					Actions: []string{
						"glacier:InitiateJob",
						"glacier:GetJobOutput",
					},
					Resources: []string{
						"arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = glacier.NewVault(ctx, "myArchiveVault", &glacier.VaultArgs{
			Notification: &glacier.VaultNotificationArgs{
				SnsTopic: awsSnsTopic.Arn,
				Events: pulumi.StringArray{
					pulumi.String("ArchiveRetrievalCompleted"),
					pulumi.String("InventoryRetrievalCompleted"),
				},
			},
			AccessPolicy: *pulumi.String(myArchivePolicyDocument.Json),
			Tags: pulumi.StringMap{
				"Test": pulumi.String("MyArchive"),
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
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.glacier.Vault;
import com.pulumi.aws.glacier.VaultArgs;
import com.pulumi.aws.glacier.inputs.VaultNotificationArgs;
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
        var awsSnsTopic = new Topic("awsSnsTopic");

        final var myArchivePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("add-read-only-perm")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("*")
                    .identifiers("*")
                    .build())
                .actions(                
                    "glacier:InitiateJob",
                    "glacier:GetJobOutput")
                .resources("arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive")
                .build())
            .build());

        var myArchiveVault = new Vault("myArchiveVault", VaultArgs.builder()        
            .notification(VaultNotificationArgs.builder()
                .snsTopic(awsSnsTopic.arn())
                .events(                
                    "ArchiveRetrievalCompleted",
                    "InventoryRetrievalCompleted")
                .build())
            .accessPolicy(myArchivePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .tags(Map.of("Test", "MyArchive"))
            .build());

    }
}
```
```yaml
resources:
  awsSnsTopic:
    type: aws:sns:Topic
  myArchiveVault:
    type: aws:glacier:Vault
    properties:
      notification:
        snsTopic: ${awsSnsTopic.arn}
        events:
          - ArchiveRetrievalCompleted
          - InventoryRetrievalCompleted
      accessPolicy: ${myArchivePolicyDocument.json}
      tags:
        Test: MyArchive
variables:
  myArchivePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: add-read-only-perm
            effect: Allow
            principals:
              - type: '*'
                identifiers:
                  - '*'
            actions:
              - glacier:InitiateJob
              - glacier:GetJobOutput
            resources:
              - arn:aws:glacier:eu-west-1:432981146916:vaults/MyArchive
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glacier Vaults using the `name`. For example:

```sh
 $ pulumi import aws:glacier/vault:Vault archive my_archive
```
 