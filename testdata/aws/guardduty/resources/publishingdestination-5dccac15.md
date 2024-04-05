Provides a resource to manage a GuardDuty PublishingDestination. Requires an existing GuardDuty Detector.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentCallerIdentity = aws.getCallerIdentity({});
const currentRegion = aws.getRegion({});
const gdBucket = new aws.s3.BucketV2("gdBucket", {forceDestroy: true});
const bucketPol = aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            sid: "Allow PutObject",
            actions: ["s3:PutObject"],
            resources: [pulumi.interpolate`${gdBucket.arn}/*`],
            principals: [{
                type: "Service",
                identifiers: ["guardduty.amazonaws.com"],
            }],
        },
        {
            sid: "Allow GetBucketLocation",
            actions: ["s3:GetBucketLocation"],
            resources: [gdBucket.arn],
            principals: [{
                type: "Service",
                identifiers: ["guardduty.amazonaws.com"],
            }],
        },
    ],
});
const kmsPol = Promise.all([currentRegion, currentCallerIdentity, currentRegion, currentCallerIdentity, currentCallerIdentity]).then(([currentRegion, currentCallerIdentity, currentRegion1, currentCallerIdentity1, currentCallerIdentity2]) => aws.iam.getPolicyDocument({
    statements: [
        {
            sid: "Allow GuardDuty to encrypt findings",
            actions: ["kms:GenerateDataKey"],
            resources: [`arn:aws:kms:${currentRegion.name}:${currentCallerIdentity.accountId}:key/*`],
            principals: [{
                type: "Service",
                identifiers: ["guardduty.amazonaws.com"],
            }],
        },
        {
            sid: "Allow all users to modify/delete key (test only)",
            actions: ["kms:*"],
            resources: [`arn:aws:kms:${currentRegion1.name}:${currentCallerIdentity1.accountId}:key/*`],
            principals: [{
                type: "AWS",
                identifiers: [`arn:aws:iam::${currentCallerIdentity2.accountId}:root`],
            }],
        },
    ],
}));
const testGd = new aws.guardduty.Detector("testGd", {enable: true});
const gdBucketAcl = new aws.s3.BucketAclV2("gdBucketAcl", {
    bucket: gdBucket.id,
    acl: "private",
});
const gdBucketPolicy = new aws.s3.BucketPolicy("gdBucketPolicy", {
    bucket: gdBucket.id,
    policy: bucketPol.apply(bucketPol => bucketPol.json),
});
const gdKey = new aws.kms.Key("gdKey", {
    description: "Temporary key for AccTest of TF",
    deletionWindowInDays: 7,
    policy: kmsPol.then(kmsPol => kmsPol.json),
});
const test = new aws.guardduty.PublishingDestination("test", {
    detectorId: testGd.id,
    destinationArn: gdBucket.arn,
    kmsKeyArn: gdKey.arn,
}, {
    dependsOn: [gdBucketPolicy],
});
```
```python
import pulumi
import pulumi_aws as aws

current_caller_identity = aws.get_caller_identity()
current_region = aws.get_region()
gd_bucket = aws.s3.BucketV2("gdBucket", force_destroy=True)
bucket_pol = aws.iam.get_policy_document_output(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="Allow PutObject",
        actions=["s3:PutObject"],
        resources=[gd_bucket.arn.apply(lambda arn: f"{arn}/*")],
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=["guardduty.amazonaws.com"],
        )],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="Allow GetBucketLocation",
        actions=["s3:GetBucketLocation"],
        resources=[gd_bucket.arn],
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=["guardduty.amazonaws.com"],
        )],
    ),
])
kms_pol = aws.iam.get_policy_document(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="Allow GuardDuty to encrypt findings",
        actions=["kms:GenerateDataKey"],
        resources=[f"arn:aws:kms:{current_region.name}:{current_caller_identity.account_id}:key/*"],
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="Service",
            identifiers=["guardduty.amazonaws.com"],
        )],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        sid="Allow all users to modify/delete key (test only)",
        actions=["kms:*"],
        resources=[f"arn:aws:kms:{current_region.name}:{current_caller_identity.account_id}:key/*"],
        principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
            type="AWS",
            identifiers=[f"arn:aws:iam::{current_caller_identity.account_id}:root"],
        )],
    ),
])
test_gd = aws.guardduty.Detector("testGd", enable=True)
gd_bucket_acl = aws.s3.BucketAclV2("gdBucketAcl",
    bucket=gd_bucket.id,
    acl="private")
gd_bucket_policy = aws.s3.BucketPolicy("gdBucketPolicy",
    bucket=gd_bucket.id,
    policy=bucket_pol.json)
gd_key = aws.kms.Key("gdKey",
    description="Temporary key for AccTest of TF",
    deletion_window_in_days=7,
    policy=kms_pol.json)
test = aws.guardduty.PublishingDestination("test",
    detector_id=test_gd.id,
    destination_arn=gd_bucket.arn,
    kms_key_arn=gd_key.arn,
    opts=pulumi.ResourceOptions(depends_on=[gd_bucket_policy]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var currentRegion = Aws.GetRegion.Invoke();

    var gdBucket = new Aws.S3.BucketV2("gdBucket", new()
    {
        ForceDestroy = true,
    });

    var bucketPol = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "Allow PutObject",
                Actions = new[]
                {
                    "s3:PutObject",
                },
                Resources = new[]
                {
                    $"{gdBucket.Arn}/*",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "guardduty.amazonaws.com",
                        },
                    },
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "Allow GetBucketLocation",
                Actions = new[]
                {
                    "s3:GetBucketLocation",
                },
                Resources = new[]
                {
                    gdBucket.Arn,
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "guardduty.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var kmsPol = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "Allow GuardDuty to encrypt findings",
                Actions = new[]
                {
                    "kms:GenerateDataKey",
                },
                Resources = new[]
                {
                    $"arn:aws:kms:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:key/*",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "guardduty.amazonaws.com",
                        },
                    },
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "Allow all users to modify/delete key (test only)",
                Actions = new[]
                {
                    "kms:*",
                },
                Resources = new[]
                {
                    $"arn:aws:kms:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:key/*",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            $"arn:aws:iam::{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:root",
                        },
                    },
                },
            },
        },
    });

    var testGd = new Aws.GuardDuty.Detector("testGd", new()
    {
        Enable = true,
    });

    var gdBucketAcl = new Aws.S3.BucketAclV2("gdBucketAcl", new()
    {
        Bucket = gdBucket.Id,
        Acl = "private",
    });

    var gdBucketPolicy = new Aws.S3.BucketPolicy("gdBucketPolicy", new()
    {
        Bucket = gdBucket.Id,
        Policy = bucketPol.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var gdKey = new Aws.Kms.Key("gdKey", new()
    {
        Description = "Temporary key for AccTest of TF",
        DeletionWindowInDays = 7,
        Policy = kmsPol.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var test = new Aws.GuardDuty.PublishingDestination("test", new()
    {
        DetectorId = testGd.Id,
        DestinationArn = gdBucket.Arn,
        KmsKeyArn = gdKey.Arn,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            gdBucketPolicy,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/guardduty"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		gdBucket, err := s3.NewBucketV2(ctx, "gdBucket", &s3.BucketV2Args{
			ForceDestroy: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		bucketPol := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Sid: pulumi.String("Allow PutObject"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:PutObject"),
					},
					Resources: pulumi.StringArray{
						gdBucket.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("Service"),
							Identifiers: pulumi.StringArray{
								pulumi.String("guardduty.amazonaws.com"),
							},
						},
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Sid: pulumi.String("Allow GetBucketLocation"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:GetBucketLocation"),
					},
					Resources: pulumi.StringArray{
						gdBucket.Arn,
					},
					Principals: iam.GetPolicyDocumentStatementPrincipalArray{
						&iam.GetPolicyDocumentStatementPrincipalArgs{
							Type: pulumi.String("Service"),
							Identifiers: pulumi.StringArray{
								pulumi.String("guardduty.amazonaws.com"),
							},
						},
					},
				},
			},
		}, nil)
		kmsPol, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid: pulumi.StringRef("Allow GuardDuty to encrypt findings"),
					Actions: []string{
						"kms:GenerateDataKey",
					},
					Resources: []string{
						fmt.Sprintf("arn:aws:kms:%v:%v:key/*", currentRegion.Name, currentCallerIdentity.AccountId),
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"guardduty.amazonaws.com",
							},
						},
					},
				},
				{
					Sid: pulumi.StringRef("Allow all users to modify/delete key (test only)"),
					Actions: []string{
						"kms:*",
					},
					Resources: []string{
						fmt.Sprintf("arn:aws:kms:%v:%v:key/*", currentRegion.Name, currentCallerIdentity.AccountId),
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "AWS",
							Identifiers: []string{
								fmt.Sprintf("arn:aws:iam::%v:root", currentCallerIdentity.AccountId),
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		testGd, err := guardduty.NewDetector(ctx, "testGd", &guardduty.DetectorArgs{
			Enable: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "gdBucketAcl", &s3.BucketAclV2Args{
			Bucket: gdBucket.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		gdBucketPolicy, err := s3.NewBucketPolicy(ctx, "gdBucketPolicy", &s3.BucketPolicyArgs{
			Bucket: gdBucket.ID(),
			Policy: bucketPol.ApplyT(func(bucketPol iam.GetPolicyDocumentResult) (*string, error) {
				return &bucketPol.Json, nil
			}).(pulumi.StringPtrOutput),
		})
		if err != nil {
			return err
		}
		gdKey, err := kms.NewKey(ctx, "gdKey", &kms.KeyArgs{
			Description:          pulumi.String("Temporary key for AccTest of TF"),
			DeletionWindowInDays: pulumi.Int(7),
			Policy:               *pulumi.String(kmsPol.Json),
		})
		if err != nil {
			return err
		}
		_, err = guardduty.NewPublishingDestination(ctx, "test", &guardduty.PublishingDestinationArgs{
			DetectorId:     testGd.ID(),
			DestinationArn: gdBucket.Arn,
			KmsKeyArn:      gdKey.Arn,
		}, pulumi.DependsOn([]pulumi.Resource{
			gdBucketPolicy,
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.guardduty.Detector;
import com.pulumi.aws.guardduty.DetectorArgs;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.guardduty.PublishingDestination;
import com.pulumi.aws.guardduty.PublishingDestinationArgs;
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
        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        final var currentRegion = AwsFunctions.getRegion();

        var gdBucket = new BucketV2("gdBucket", BucketV2Args.builder()        
            .forceDestroy(true)
            .build());

        final var bucketPol = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .sid("Allow PutObject")
                    .actions("s3:PutObject")
                    .resources(gdBucket.arn().applyValue(arn -> String.format("%s/*", arn)))
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("Service")
                        .identifiers("guardduty.amazonaws.com")
                        .build())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .sid("Allow GetBucketLocation")
                    .actions("s3:GetBucketLocation")
                    .resources(gdBucket.arn())
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("Service")
                        .identifiers("guardduty.amazonaws.com")
                        .build())
                    .build())
            .build());

        final var kmsPol = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .sid("Allow GuardDuty to encrypt findings")
                    .actions("kms:GenerateDataKey")
                    .resources(String.format("arn:aws:kms:%s:%s:key/*", currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("Service")
                        .identifiers("guardduty.amazonaws.com")
                        .build())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .sid("Allow all users to modify/delete key (test only)")
                    .actions("kms:*")
                    .resources(String.format("arn:aws:kms:%s:%s:key/*", currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                    .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                        .type("AWS")
                        .identifiers(String.format("arn:aws:iam::%s:root", currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                        .build())
                    .build())
            .build());

        var testGd = new Detector("testGd", DetectorArgs.builder()        
            .enable(true)
            .build());

        var gdBucketAcl = new BucketAclV2("gdBucketAcl", BucketAclV2Args.builder()        
            .bucket(gdBucket.id())
            .acl("private")
            .build());

        var gdBucketPolicy = new BucketPolicy("gdBucketPolicy", BucketPolicyArgs.builder()        
            .bucket(gdBucket.id())
            .policy(bucketPol.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(bucketPol -> bucketPol.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var gdKey = new Key("gdKey", KeyArgs.builder()        
            .description("Temporary key for AccTest of TF")
            .deletionWindowInDays(7)
            .policy(kmsPol.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var test = new PublishingDestination("test", PublishingDestinationArgs.builder()        
            .detectorId(testGd.id())
            .destinationArn(gdBucket.arn())
            .kmsKeyArn(gdKey.arn())
            .build(), CustomResourceOptions.builder()
                .dependsOn(gdBucketPolicy)
                .build());

    }
}
```
```yaml
resources:
  testGd:
    type: aws:guardduty:Detector
    properties:
      enable: true
  gdBucket:
    type: aws:s3:BucketV2
    properties:
      forceDestroy: true
  gdBucketAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${gdBucket.id}
      acl: private
  gdBucketPolicy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${gdBucket.id}
      policy: ${bucketPol.json}
  gdKey:
    type: aws:kms:Key
    properties:
      description: Temporary key for AccTest of TF
      deletionWindowInDays: 7
      policy: ${kmsPol.json}
  test:
    type: aws:guardduty:PublishingDestination
    properties:
      detectorId: ${testGd.id}
      destinationArn: ${gdBucket.arn}
      kmsKeyArn: ${gdKey.arn}
    options:
      dependson:
        - ${gdBucketPolicy}
variables:
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  bucketPol:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: Allow PutObject
            actions:
              - s3:PutObject
            resources:
              - ${gdBucket.arn}/*
            principals:
              - type: Service
                identifiers:
                  - guardduty.amazonaws.com
          - sid: Allow GetBucketLocation
            actions:
              - s3:GetBucketLocation
            resources:
              - ${gdBucket.arn}
            principals:
              - type: Service
                identifiers:
                  - guardduty.amazonaws.com
  kmsPol:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: Allow GuardDuty to encrypt findings
            actions:
              - kms:GenerateDataKey
            resources:
              - arn:aws:kms:${currentRegion.name}:${currentCallerIdentity.accountId}:key/*
            principals:
              - type: Service
                identifiers:
                  - guardduty.amazonaws.com
          - sid: Allow all users to modify/delete key (test only)
            actions:
              - kms:*
            resources:
              - arn:aws:kms:${currentRegion.name}:${currentCallerIdentity.accountId}:key/*
            principals:
              - type: AWS
                identifiers:
                  - arn:aws:iam::${currentCallerIdentity.accountId}:root
```

> **Note:** Please do not use this simple example for Bucket-Policy and KMS Key Policy in a production environment. It is much too open for such a use-case. Refer to the AWS documentation here: https://docs.aws.amazon.com/guardduty/latest/ug/guardduty_exportfindings.html
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import GuardDuty PublishingDestination using the master GuardDuty detector ID and PublishingDestinationID. For example:

```sh
 $ pulumi import aws:guardduty/publishingDestination:PublishingDestination test a4b86f26fa42e7e7cf0d1c333ea77777:a4b86f27a0e464e4a7e0516d242f1234
```
 