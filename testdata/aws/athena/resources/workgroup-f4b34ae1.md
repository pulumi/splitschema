Provides an Athena Workgroup.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.athena.Workgroup("example", {configuration: {
    enforceWorkgroupConfiguration: true,
    publishCloudwatchMetricsEnabled: true,
    resultConfiguration: {
        outputLocation: `s3://${aws_s3_bucket.example.bucket}/output/`,
        encryptionConfiguration: {
            encryptionOption: "SSE_KMS",
            kmsKeyArn: aws_kms_key.example.arn,
        },
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.athena.Workgroup("example", configuration=aws.athena.WorkgroupConfigurationArgs(
    enforce_workgroup_configuration=True,
    publish_cloudwatch_metrics_enabled=True,
    result_configuration=aws.athena.WorkgroupConfigurationResultConfigurationArgs(
        output_location=f"s3://{aws_s3_bucket['example']['bucket']}/output/",
        encryption_configuration=aws.athena.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs(
            encryption_option="SSE_KMS",
            kms_key_arn=aws_kms_key["example"]["arn"],
        ),
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
    var example = new Aws.Athena.Workgroup("example", new()
    {
        Configuration = new Aws.Athena.Inputs.WorkgroupConfigurationArgs
        {
            EnforceWorkgroupConfiguration = true,
            PublishCloudwatchMetricsEnabled = true,
            ResultConfiguration = new Aws.Athena.Inputs.WorkgroupConfigurationResultConfigurationArgs
            {
                OutputLocation = $"s3://{aws_s3_bucket.Example.Bucket}/output/",
                EncryptionConfiguration = new Aws.Athena.Inputs.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs
                {
                    EncryptionOption = "SSE_KMS",
                    KmsKeyArn = aws_kms_key.Example.Arn,
                },
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/athena"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := athena.NewWorkgroup(ctx, "example", &athena.WorkgroupArgs{
			Configuration: &athena.WorkgroupConfigurationArgs{
				EnforceWorkgroupConfiguration:   pulumi.Bool(true),
				PublishCloudwatchMetricsEnabled: pulumi.Bool(true),
				ResultConfiguration: &athena.WorkgroupConfigurationResultConfigurationArgs{
					OutputLocation: pulumi.String(fmt.Sprintf("s3://%v/output/", aws_s3_bucket.Example.Bucket)),
					EncryptionConfiguration: &athena.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs{
						EncryptionOption: pulumi.String("SSE_KMS"),
						KmsKeyArn:        pulumi.Any(aws_kms_key.Example.Arn),
					},
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
import com.pulumi.aws.athena.Workgroup;
import com.pulumi.aws.athena.WorkgroupArgs;
import com.pulumi.aws.athena.inputs.WorkgroupConfigurationArgs;
import com.pulumi.aws.athena.inputs.WorkgroupConfigurationResultConfigurationArgs;
import com.pulumi.aws.athena.inputs.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs;
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
        var example = new Workgroup("example", WorkgroupArgs.builder()        
            .configuration(WorkgroupConfigurationArgs.builder()
                .enforceWorkgroupConfiguration(true)
                .publishCloudwatchMetricsEnabled(true)
                .resultConfiguration(WorkgroupConfigurationResultConfigurationArgs.builder()
                    .outputLocation(String.format("s3://%s/output/", aws_s3_bucket.example().bucket()))
                    .encryptionConfiguration(WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs.builder()
                        .encryptionOption("SSE_KMS")
                        .kmsKeyArn(aws_kms_key.example().arn())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:athena:Workgroup
    properties:
      configuration:
        enforceWorkgroupConfiguration: true
        publishCloudwatchMetricsEnabled: true
        resultConfiguration:
          outputLocation: s3://${aws_s3_bucket.example.bucket}/output/
          encryptionConfiguration:
            encryptionOption: SSE_KMS
            kmsKeyArn: ${aws_kms_key.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Athena Workgroups using their name. For example:

```sh
 $ pulumi import aws:athena/workgroup:Workgroup example example
```
 