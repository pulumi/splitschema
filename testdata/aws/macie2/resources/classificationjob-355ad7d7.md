Provides a resource to manage an [AWS Macie Classification Job](https://docs.aws.amazon.com/macie/latest/APIReference/jobs.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testAccount = new aws.macie2.Account("testAccount", {});
const testClassificationJob = new aws.macie2.ClassificationJob("testClassificationJob", {
    jobType: "ONE_TIME",
    s3JobDefinition: {
        bucketDefinitions: [{
            accountId: "ACCOUNT ID",
            buckets: ["S3 BUCKET NAME"],
        }],
    },
}, {
    dependsOn: [testAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

test_account = aws.macie2.Account("testAccount")
test_classification_job = aws.macie2.ClassificationJob("testClassificationJob",
    job_type="ONE_TIME",
    s3_job_definition=aws.macie2.ClassificationJobS3JobDefinitionArgs(
        bucket_definitions=[aws.macie2.ClassificationJobS3JobDefinitionBucketDefinitionArgs(
            account_id="ACCOUNT ID",
            buckets=["S3 BUCKET NAME"],
        )],
    ),
    opts=pulumi.ResourceOptions(depends_on=[test_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testAccount = new Aws.Macie2.Account("testAccount");

    var testClassificationJob = new Aws.Macie2.ClassificationJob("testClassificationJob", new()
    {
        JobType = "ONE_TIME",
        S3JobDefinition = new Aws.Macie2.Inputs.ClassificationJobS3JobDefinitionArgs
        {
            BucketDefinitions = new[]
            {
                new Aws.Macie2.Inputs.ClassificationJobS3JobDefinitionBucketDefinitionArgs
                {
                    AccountId = "ACCOUNT ID",
                    Buckets = new[]
                    {
                        "S3 BUCKET NAME",
                    },
                },
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            testAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/macie2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testAccount, err := macie2.NewAccount(ctx, "testAccount", nil)
		if err != nil {
			return err
		}
		_, err = macie2.NewClassificationJob(ctx, "testClassificationJob", &macie2.ClassificationJobArgs{
			JobType: pulumi.String("ONE_TIME"),
			S3JobDefinition: &macie2.ClassificationJobS3JobDefinitionArgs{
				BucketDefinitions: macie2.ClassificationJobS3JobDefinitionBucketDefinitionArray{
					&macie2.ClassificationJobS3JobDefinitionBucketDefinitionArgs{
						AccountId: pulumi.String("ACCOUNT ID"),
						Buckets: pulumi.StringArray{
							pulumi.String("S3 BUCKET NAME"),
						},
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			testAccount,
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
import com.pulumi.aws.macie2.Account;
import com.pulumi.aws.macie2.ClassificationJob;
import com.pulumi.aws.macie2.ClassificationJobArgs;
import com.pulumi.aws.macie2.inputs.ClassificationJobS3JobDefinitionArgs;
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
        var testAccount = new Account("testAccount");

        var testClassificationJob = new ClassificationJob("testClassificationJob", ClassificationJobArgs.builder()        
            .jobType("ONE_TIME")
            .s3JobDefinition(ClassificationJobS3JobDefinitionArgs.builder()
                .bucketDefinitions(ClassificationJobS3JobDefinitionBucketDefinitionArgs.builder()
                    .accountId("ACCOUNT ID")
                    .buckets("S3 BUCKET NAME")
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(testAccount)
                .build());

    }
}
```
```yaml
resources:
  testAccount:
    type: aws:macie2:Account
  testClassificationJob:
    type: aws:macie2:ClassificationJob
    properties:
      jobType: ONE_TIME
      s3JobDefinition:
        bucketDefinitions:
          - accountId: ACCOUNT ID
            buckets:
              - S3 BUCKET NAME
    options:
      dependson:
        - ${testAccount}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_macie2_classification_job` using the id. For example:

```sh
 $ pulumi import aws:macie2/classificationJob:ClassificationJob example abcd1
```
 