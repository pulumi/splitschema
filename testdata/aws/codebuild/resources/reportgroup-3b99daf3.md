Provides a CodeBuild Report Groups Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const examplePolicyDocument = current.then(current => aws.iam.getPolicyDocument({
    statements: [{
        sid: "Enable IAM User Permissions",
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: [`arn:aws:iam::${current.accountId}:root`],
        }],
        actions: ["kms:*"],
        resources: ["*"],
    }],
}));
const exampleKey = new aws.kms.Key("exampleKey", {
    description: "my test kms key",
    deletionWindowInDays: 7,
    policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleReportGroup = new aws.codebuild.ReportGroup("exampleReportGroup", {
    type: "TEST",
    exportConfig: {
        type: "S3",
        s3Destination: {
            bucket: exampleBucketV2.id,
            encryptionDisabled: false,
            encryptionKey: exampleKey.arn,
            packaging: "NONE",
            path: "/some",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="Enable IAM User Permissions",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=[f"arn:aws:iam::{current.account_id}:root"],
    )],
    actions=["kms:*"],
    resources=["*"],
)])
example_key = aws.kms.Key("exampleKey",
    description="my test kms key",
    deletion_window_in_days=7,
    policy=example_policy_document.json)
example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_report_group = aws.codebuild.ReportGroup("exampleReportGroup",
    type="TEST",
    export_config=aws.codebuild.ReportGroupExportConfigArgs(
        type="S3",
        s3_destination=aws.codebuild.ReportGroupExportConfigS3DestinationArgs(
            bucket=example_bucket_v2.id,
            encryption_disabled=False,
            encryption_key=example_key.arn,
            packaging="NONE",
            path="/some",
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "Enable IAM User Permissions",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            $"arn:aws:iam::{current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:root",
                        },
                    },
                },
                Actions = new[]
                {
                    "kms:*",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "my test kms key",
        DeletionWindowInDays = 7,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleReportGroup = new Aws.CodeBuild.ReportGroup("exampleReportGroup", new()
    {
        Type = "TEST",
        ExportConfig = new Aws.CodeBuild.Inputs.ReportGroupExportConfigArgs
        {
            Type = "S3",
            S3Destination = new Aws.CodeBuild.Inputs.ReportGroupExportConfigS3DestinationArgs
            {
                Bucket = exampleBucketV2.Id,
                EncryptionDisabled = false,
                EncryptionKey = exampleKey.Arn,
                Packaging = "NONE",
                Path = "/some",
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codebuild"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("Enable IAM User Permissions"),
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "AWS",
							Identifiers: []string{
								fmt.Sprintf("arn:aws:iam::%v:root", current.AccountId),
							},
						},
					},
					Actions: []string{
						"kms:*",
					},
					Resources: []string{
						"*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description:          pulumi.String("my test kms key"),
			DeletionWindowInDays: pulumi.Int(7),
			Policy:               *pulumi.String(examplePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		_, err = codebuild.NewReportGroup(ctx, "exampleReportGroup", &codebuild.ReportGroupArgs{
			Type: pulumi.String("TEST"),
			ExportConfig: &codebuild.ReportGroupExportConfigArgs{
				Type: pulumi.String("S3"),
				S3Destination: &codebuild.ReportGroupExportConfigS3DestinationArgs{
					Bucket:             exampleBucketV2.ID(),
					EncryptionDisabled: pulumi.Bool(false),
					EncryptionKey:      exampleKey.Arn,
					Packaging:          pulumi.String("NONE"),
					Path:               pulumi.String("/some"),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.codebuild.ReportGroup;
import com.pulumi.aws.codebuild.ReportGroupArgs;
import com.pulumi.aws.codebuild.inputs.ReportGroupExportConfigArgs;
import com.pulumi.aws.codebuild.inputs.ReportGroupExportConfigS3DestinationArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("Enable IAM User Permissions")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers(String.format("arn:aws:iam::%s:root", current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                    .build())
                .actions("kms:*")
                .resources("*")
                .build())
            .build());

        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("my test kms key")
            .deletionWindowInDays(7)
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleReportGroup = new ReportGroup("exampleReportGroup", ReportGroupArgs.builder()        
            .type("TEST")
            .exportConfig(ReportGroupExportConfigArgs.builder()
                .type("S3")
                .s3Destination(ReportGroupExportConfigS3DestinationArgs.builder()
                    .bucket(exampleBucketV2.id())
                    .encryptionDisabled(false)
                    .encryptionKey(exampleKey.arn())
                    .packaging("NONE")
                    .path("/some")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: my test kms key
      deletionWindowInDays: 7
      policy: ${examplePolicyDocument.json}
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleReportGroup:
    type: aws:codebuild:ReportGroup
    properties:
      type: TEST
      exportConfig:
        type: S3
        s3Destination:
          bucket: ${exampleBucketV2.id}
          encryptionDisabled: false
          encryptionKey: ${exampleKey.arn}
          packaging: NONE
          path: /some
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: Enable IAM User Permissions
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - arn:aws:iam::${current.accountId}:root
            actions:
              - kms:*
            resources:
              - '*'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeBuild Report Group using the CodeBuild Report Group arn. For example:

```sh
 $ pulumi import aws:codebuild/reportGroup:ReportGroup example arn:aws:codebuild:us-west-2:123456789:report-group/report-group-name
```
 