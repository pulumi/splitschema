Manages a Glue Security Configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.SecurityConfiguration("example", {encryptionConfiguration: {
    cloudwatchEncryption: {
        cloudwatchEncryptionMode: "DISABLED",
    },
    jobBookmarksEncryption: {
        jobBookmarksEncryptionMode: "DISABLED",
    },
    s3Encryption: {
        kmsKeyArn: data.aws_kms_key.example.arn,
        s3EncryptionMode: "SSE-KMS",
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.SecurityConfiguration("example", encryption_configuration=aws.glue.SecurityConfigurationEncryptionConfigurationArgs(
    cloudwatch_encryption=aws.glue.SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionArgs(
        cloudwatch_encryption_mode="DISABLED",
    ),
    job_bookmarks_encryption=aws.glue.SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionArgs(
        job_bookmarks_encryption_mode="DISABLED",
    ),
    s3_encryption=aws.glue.SecurityConfigurationEncryptionConfigurationS3EncryptionArgs(
        kms_key_arn=data["aws_kms_key"]["example"]["arn"],
        s3_encryption_mode="SSE-KMS",
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
    var example = new Aws.Glue.SecurityConfiguration("example", new()
    {
        EncryptionConfiguration = new Aws.Glue.Inputs.SecurityConfigurationEncryptionConfigurationArgs
        {
            CloudwatchEncryption = new Aws.Glue.Inputs.SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionArgs
            {
                CloudwatchEncryptionMode = "DISABLED",
            },
            JobBookmarksEncryption = new Aws.Glue.Inputs.SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionArgs
            {
                JobBookmarksEncryptionMode = "DISABLED",
            },
            S3Encryption = new Aws.Glue.Inputs.SecurityConfigurationEncryptionConfigurationS3EncryptionArgs
            {
                KmsKeyArn = data.Aws_kms_key.Example.Arn,
                S3EncryptionMode = "SSE-KMS",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewSecurityConfiguration(ctx, "example", &glue.SecurityConfigurationArgs{
			EncryptionConfiguration: &glue.SecurityConfigurationEncryptionConfigurationArgs{
				CloudwatchEncryption: &glue.SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionArgs{
					CloudwatchEncryptionMode: pulumi.String("DISABLED"),
				},
				JobBookmarksEncryption: &glue.SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionArgs{
					JobBookmarksEncryptionMode: pulumi.String("DISABLED"),
				},
				S3Encryption: &glue.SecurityConfigurationEncryptionConfigurationS3EncryptionArgs{
					KmsKeyArn:        pulumi.Any(data.Aws_kms_key.Example.Arn),
					S3EncryptionMode: pulumi.String("SSE-KMS"),
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
import com.pulumi.aws.glue.SecurityConfiguration;
import com.pulumi.aws.glue.SecurityConfigurationArgs;
import com.pulumi.aws.glue.inputs.SecurityConfigurationEncryptionConfigurationArgs;
import com.pulumi.aws.glue.inputs.SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionArgs;
import com.pulumi.aws.glue.inputs.SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionArgs;
import com.pulumi.aws.glue.inputs.SecurityConfigurationEncryptionConfigurationS3EncryptionArgs;
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
        var example = new SecurityConfiguration("example", SecurityConfigurationArgs.builder()        
            .encryptionConfiguration(SecurityConfigurationEncryptionConfigurationArgs.builder()
                .cloudwatchEncryption(SecurityConfigurationEncryptionConfigurationCloudwatchEncryptionArgs.builder()
                    .cloudwatchEncryptionMode("DISABLED")
                    .build())
                .jobBookmarksEncryption(SecurityConfigurationEncryptionConfigurationJobBookmarksEncryptionArgs.builder()
                    .jobBookmarksEncryptionMode("DISABLED")
                    .build())
                .s3Encryption(SecurityConfigurationEncryptionConfigurationS3EncryptionArgs.builder()
                    .kmsKeyArn(data.aws_kms_key().example().arn())
                    .s3EncryptionMode("SSE-KMS")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:SecurityConfiguration
    properties:
      encryptionConfiguration:
        cloudwatchEncryption:
          cloudwatchEncryptionMode: DISABLED
        jobBookmarksEncryption:
          jobBookmarksEncryptionMode: DISABLED
        s3Encryption:
          kmsKeyArn: ${data.aws_kms_key.example.arn}
          s3EncryptionMode: SSE-KMS
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Security Configurations using `name`. For example:

```sh
 $ pulumi import aws:glue/securityConfiguration:SecurityConfiguration example example
```
 