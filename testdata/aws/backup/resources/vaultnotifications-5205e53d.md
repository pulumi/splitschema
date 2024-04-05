Provides an AWS Backup vault notifications resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testTopic = new aws.sns.Topic("testTopic", {});
const testPolicyDocument = testTopic.arn.apply(arn => aws.iam.getPolicyDocumentOutput({
    policyId: "__default_policy_ID",
    statements: [{
        actions: ["SNS:Publish"],
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["backup.amazonaws.com"],
        }],
        resources: [arn],
        sid: "__default_statement_ID",
    }],
}));
const testTopicPolicy = new aws.sns.TopicPolicy("testTopicPolicy", {
    arn: testTopic.arn,
    policy: testPolicyDocument.apply(testPolicyDocument => testPolicyDocument.json),
});
const testVaultNotifications = new aws.backup.VaultNotifications("testVaultNotifications", {
    backupVaultName: "example_backup_vault",
    snsTopicArn: testTopic.arn,
    backupVaultEvents: [
        "BACKUP_JOB_STARTED",
        "RESTORE_JOB_COMPLETED",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

test_topic = aws.sns.Topic("testTopic")
test_policy_document = test_topic.arn.apply(lambda arn: aws.iam.get_policy_document_output(policy_id="__default_policy_ID",
    statements=[aws.iam.GetPolicyDocumentStatementArgs(
        actions=["SNS:Publish"],
        effect="Allow",
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=["backup.amazonaws.com"],
        )],
        resources=[arn],
        sid="__default_statement_ID",
    )]))
test_topic_policy = aws.sns.TopicPolicy("testTopicPolicy",
    arn=test_topic.arn,
    policy=test_policy_document.json)
test_vault_notifications = aws.backup.VaultNotifications("testVaultNotifications",
    backup_vault_name="example_backup_vault",
    sns_topic_arn=test_topic.arn,
    backup_vault_events=[
        "BACKUP_JOB_STARTED",
        "RESTORE_JOB_COMPLETED",
    ])
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
        PolicyId = "__default_policy_ID",
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "SNS:Publish",
                },
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "backup.amazonaws.com",
                        },
                    },
                },
                Resources = new[]
                {
                    testTopic.Arn,
                },
                Sid = "__default_statement_ID",
            },
        },
    });

    var testTopicPolicy = new Aws.Sns.TopicPolicy("testTopicPolicy", new()
    {
        Arn = testTopic.Arn,
        Policy = testPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var testVaultNotifications = new Aws.Backup.VaultNotifications("testVaultNotifications", new()
    {
        BackupVaultName = "example_backup_vault",
        SnsTopicArn = testTopic.Arn,
        BackupVaultEvents = new[]
        {
            "BACKUP_JOB_STARTED",
            "RESTORE_JOB_COMPLETED",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
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
testPolicyDocument := testTopic.Arn.ApplyT(func(arn string) (iam.GetPolicyDocumentResult, error) {
return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
PolicyId: "__default_policy_ID",
Statements: []iam.GetPolicyDocumentStatement{
{
Actions: []string{
"SNS:Publish",
},
Effect: "Allow",
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "Service",
Identifiers: []string{
"backup.amazonaws.com",
},
},
},
Resources: interface{}{
arn,
},
Sid: "__default_statement_ID",
},
},
}, nil), nil
}).(iam.GetPolicyDocumentResultOutput)
_, err = sns.NewTopicPolicy(ctx, "testTopicPolicy", &sns.TopicPolicyArgs{
Arn: testTopic.Arn,
Policy: testPolicyDocument.ApplyT(func(testPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
return &testPolicyDocument.Json, nil
}).(pulumi.StringPtrOutput),
})
if err != nil {
return err
}
_, err = backup.NewVaultNotifications(ctx, "testVaultNotifications", &backup.VaultNotificationsArgs{
BackupVaultName: pulumi.String("example_backup_vault"),
SnsTopicArn: testTopic.Arn,
BackupVaultEvents: pulumi.StringArray{
pulumi.String("BACKUP_JOB_STARTED"),
pulumi.String("RESTORE_JOB_COMPLETED"),
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
import com.pulumi.aws.sns.TopicPolicy;
import com.pulumi.aws.sns.TopicPolicyArgs;
import com.pulumi.aws.backup.VaultNotifications;
import com.pulumi.aws.backup.VaultNotificationsArgs;
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
            .policyId("__default_policy_ID")
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("SNS:Publish")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("backup.amazonaws.com")
                    .build())
                .resources(testTopic.arn())
                .sid("__default_statement_ID")
                .build())
            .build());

        var testTopicPolicy = new TopicPolicy("testTopicPolicy", TopicPolicyArgs.builder()        
            .arn(testTopic.arn())
            .policy(testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(testPolicyDocument -> testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var testVaultNotifications = new VaultNotifications("testVaultNotifications", VaultNotificationsArgs.builder()        
            .backupVaultName("example_backup_vault")
            .snsTopicArn(testTopic.arn())
            .backupVaultEvents(            
                "BACKUP_JOB_STARTED",
                "RESTORE_JOB_COMPLETED")
            .build());

    }
}
```
```yaml
resources:
  testTopic:
    type: aws:sns:Topic
  testTopicPolicy:
    type: aws:sns:TopicPolicy
    properties:
      arn: ${testTopic.arn}
      policy: ${testPolicyDocument.json}
  testVaultNotifications:
    type: aws:backup:VaultNotifications
    properties:
      backupVaultName: example_backup_vault
      snsTopicArn: ${testTopic.arn}
      backupVaultEvents:
        - BACKUP_JOB_STARTED
        - RESTORE_JOB_COMPLETED
variables:
  testPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        policyId: __default_policy_ID
        statements:
          - actions:
              - SNS:Publish
            effect: Allow
            principals:
              - type: Service
                identifiers:
                  - backup.amazonaws.com
            resources:
              - ${testTopic.arn}
            sid: __default_statement_ID
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Backup vault notifications using the `name`. For example:

```sh
 $ pulumi import aws:backup/vaultNotifications:VaultNotifications test TestVault
```
 