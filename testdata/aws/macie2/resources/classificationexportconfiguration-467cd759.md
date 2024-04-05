Provides a resource to manage an [Amazon Macie Classification Export Configuration](https://docs.aws.amazon.com/macie/latest/APIReference/classification-export-configuration.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.macie2.Account("exampleAccount", {});
const exampleClassificationExportConfiguration = new aws.macie2.ClassificationExportConfiguration("exampleClassificationExportConfiguration", {s3Destination: {
    bucketName: aws_s3_bucket.example.bucket,
    keyPrefix: "exampleprefix/",
    kmsKeyArn: aws_kms_key.example.arn,
}}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.macie2.Account("exampleAccount")
example_classification_export_configuration = aws.macie2.ClassificationExportConfiguration("exampleClassificationExportConfiguration", s3_destination=aws.macie2.ClassificationExportConfigurationS3DestinationArgs(
    bucket_name=aws_s3_bucket["example"]["bucket"],
    key_prefix="exampleprefix/",
    kms_key_arn=aws_kms_key["example"]["arn"],
),
opts=pulumi.ResourceOptions(depends_on=[example_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.Macie2.Account("exampleAccount");

    var exampleClassificationExportConfiguration = new Aws.Macie2.ClassificationExportConfiguration("exampleClassificationExportConfiguration", new()
    {
        S3Destination = new Aws.Macie2.Inputs.ClassificationExportConfigurationS3DestinationArgs
        {
            BucketName = aws_s3_bucket.Example.Bucket,
            KeyPrefix = "exampleprefix/",
            KmsKeyArn = aws_kms_key.Example.Arn,
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccount,
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
		exampleAccount, err := macie2.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = macie2.NewClassificationExportConfiguration(ctx, "exampleClassificationExportConfiguration", &macie2.ClassificationExportConfigurationArgs{
			S3Destination: &macie2.ClassificationExportConfigurationS3DestinationArgs{
				BucketName: pulumi.Any(aws_s3_bucket.Example.Bucket),
				KeyPrefix:  pulumi.String("exampleprefix/"),
				KmsKeyArn:  pulumi.Any(aws_kms_key.Example.Arn),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccount,
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
import com.pulumi.aws.macie2.ClassificationExportConfiguration;
import com.pulumi.aws.macie2.ClassificationExportConfigurationArgs;
import com.pulumi.aws.macie2.inputs.ClassificationExportConfigurationS3DestinationArgs;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleClassificationExportConfiguration = new ClassificationExportConfiguration("exampleClassificationExportConfiguration", ClassificationExportConfigurationArgs.builder()        
            .s3Destination(ClassificationExportConfigurationS3DestinationArgs.builder()
                .bucketName(aws_s3_bucket.example().bucket())
                .keyPrefix("exampleprefix/")
                .kmsKeyArn(aws_kms_key.example().arn())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:macie2:Account
  exampleClassificationExportConfiguration:
    type: aws:macie2:ClassificationExportConfiguration
    properties:
      s3Destination:
        bucketName: ${aws_s3_bucket.example.bucket}
        keyPrefix: exampleprefix/
        kmsKeyArn: ${aws_kms_key.example.arn}
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_macie2_classification_export_configuration` using the account ID and region. For example:

```sh
 $ pulumi import aws:macie2/classificationExportConfiguration:ClassificationExportConfiguration example 123456789012:us-west-2
```
 