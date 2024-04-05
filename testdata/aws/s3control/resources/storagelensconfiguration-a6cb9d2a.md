Provides a resource to manage an S3 Storage Lens configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const example = new aws.s3control.StorageLensConfiguration("example", {
    configId: "example-1",
    storageLensConfiguration: {
        enabled: true,
        accountLevel: {
            activityMetrics: {
                enabled: true,
            },
            bucketLevel: {
                activityMetrics: {
                    enabled: true,
                },
            },
        },
        dataExport: {
            cloudWatchMetrics: {
                enabled: true,
            },
            s3BucketDestination: {
                accountId: current.then(current => current.accountId),
                arn: aws_s3_bucket.target.arn,
                format: "CSV",
                outputSchemaVersion: "V_1",
                encryption: {
                    sseS3s: [{}],
                },
            },
        },
        exclude: {
            buckets: [
                aws_s3_bucket.b1.arn,
                aws_s3_bucket.b2.arn,
            ],
            regions: ["us-east-2"],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
example = aws.s3control.StorageLensConfiguration("example",
    config_id="example-1",
    storage_lens_configuration=aws.s3control.StorageLensConfigurationStorageLensConfigurationArgs(
        enabled=True,
        account_level=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelArgs(
            activity_metrics=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetricsArgs(
                enabled=True,
            ),
            bucket_level=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelArgs(
                activity_metrics=aws.s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetricsArgs(
                    enabled=True,
                ),
            ),
        ),
        data_export=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportArgs(
            cloud_watch_metrics=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportCloudWatchMetricsArgs(
                enabled=True,
            ),
            s3_bucket_destination=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationArgs(
                account_id=current.account_id,
                arn=aws_s3_bucket["target"]["arn"],
                format="CSV",
                output_schema_version="V_1",
                encryption=aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionArgs(
                    sse_s3s=[aws.s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionSseS3Args()],
                ),
            ),
        ),
        exclude=aws.s3control.StorageLensConfigurationStorageLensConfigurationExcludeArgs(
            buckets=[
                aws_s3_bucket["b1"]["arn"],
                aws_s3_bucket["b2"]["arn"],
            ],
            regions=["us-east-2"],
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

    var example = new Aws.S3Control.StorageLensConfiguration("example", new()
    {
        ConfigId = "example-1",
        StorageLensConfigurationDetail = new Aws.S3Control.Inputs.StorageLensConfigurationStorageLensConfigurationArgs
        {
            Enabled = true,
            AccountLevel = new Aws.S3Control.Inputs.StorageLensConfigurationStorageLensConfigurationAccountLevelArgs
            {
                ActivityMetrics = new Aws.S3Control.Inputs.StorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetricsArgs
                {
                    Enabled = true,
                },
                BucketLevel = new Aws.S3Control.Inputs.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelArgs
                {
                    ActivityMetrics = new Aws.S3Control.Inputs.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetricsArgs
                    {
                        Enabled = true,
                    },
                },
            },
            DataExport = new Aws.S3Control.Inputs.StorageLensConfigurationStorageLensConfigurationDataExportArgs
            {
                CloudWatchMetrics = new Aws.S3Control.Inputs.StorageLensConfigurationStorageLensConfigurationDataExportCloudWatchMetricsArgs
                {
                    Enabled = true,
                },
                S3BucketDestination = new Aws.S3Control.Inputs.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationArgs
                {
                    AccountId = current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
                    Arn = aws_s3_bucket.Target.Arn,
                    Format = "CSV",
                    OutputSchemaVersion = "V_1",
                    Encryption = new Aws.S3Control.Inputs.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionArgs
                    {
                        SseS3s = new[]
                        {
                            null,
                        },
                    },
                },
            },
            Exclude = new Aws.S3Control.Inputs.StorageLensConfigurationStorageLensConfigurationExcludeArgs
            {
                Buckets = new[]
                {
                    aws_s3_bucket.B1.Arn,
                    aws_s3_bucket.B2.Arn,
                },
                Regions = new[]
                {
                    "us-east-2",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = s3control.NewStorageLensConfiguration(ctx, "example", &s3control.StorageLensConfigurationArgs{
			ConfigId: pulumi.String("example-1"),
			StorageLensConfiguration: &s3control.StorageLensConfigurationStorageLensConfigurationArgs{
				Enabled: pulumi.Bool(true),
				AccountLevel: &s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelArgs{
					ActivityMetrics: &s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetricsArgs{
						Enabled: pulumi.Bool(true),
					},
					BucketLevel: &s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelArgs{
						ActivityMetrics: &s3control.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetricsArgs{
							Enabled: pulumi.Bool(true),
						},
					},
				},
				DataExport: &s3control.StorageLensConfigurationStorageLensConfigurationDataExportArgs{
					CloudWatchMetrics: &s3control.StorageLensConfigurationStorageLensConfigurationDataExportCloudWatchMetricsArgs{
						Enabled: pulumi.Bool(true),
					},
					S3BucketDestination: &s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationArgs{
						AccountId:           *pulumi.String(current.AccountId),
						Arn:                 pulumi.Any(aws_s3_bucket.Target.Arn),
						Format:              pulumi.String("CSV"),
						OutputSchemaVersion: pulumi.String("V_1"),
						Encryption: &s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionArgs{
							SseS3s: s3control.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionSseS3Array{
								nil,
							},
						},
					},
				},
				Exclude: &s3control.StorageLensConfigurationStorageLensConfigurationExcludeArgs{
					Buckets: pulumi.StringArray{
						aws_s3_bucket.B1.Arn,
						aws_s3_bucket.B2.Arn,
					},
					Regions: pulumi.StringArray{
						pulumi.String("us-east-2"),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.s3control.StorageLensConfiguration;
import com.pulumi.aws.s3control.StorageLensConfigurationArgs;
import com.pulumi.aws.s3control.inputs.StorageLensConfigurationStorageLensConfigurationArgs;
import com.pulumi.aws.s3control.inputs.StorageLensConfigurationStorageLensConfigurationAccountLevelArgs;
import com.pulumi.aws.s3control.inputs.StorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetricsArgs;
import com.pulumi.aws.s3control.inputs.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelArgs;
import com.pulumi.aws.s3control.inputs.StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetricsArgs;
import com.pulumi.aws.s3control.inputs.StorageLensConfigurationStorageLensConfigurationDataExportArgs;
import com.pulumi.aws.s3control.inputs.StorageLensConfigurationStorageLensConfigurationDataExportCloudWatchMetricsArgs;
import com.pulumi.aws.s3control.inputs.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationArgs;
import com.pulumi.aws.s3control.inputs.StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionArgs;
import com.pulumi.aws.s3control.inputs.StorageLensConfigurationStorageLensConfigurationExcludeArgs;
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

        var example = new StorageLensConfiguration("example", StorageLensConfigurationArgs.builder()        
            .configId("example-1")
            .storageLensConfiguration(StorageLensConfigurationStorageLensConfigurationArgs.builder()
                .enabled(true)
                .accountLevel(StorageLensConfigurationStorageLensConfigurationAccountLevelArgs.builder()
                    .activityMetrics(StorageLensConfigurationStorageLensConfigurationAccountLevelActivityMetricsArgs.builder()
                        .enabled(true)
                        .build())
                    .bucketLevel(StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelArgs.builder()
                        .activityMetrics(StorageLensConfigurationStorageLensConfigurationAccountLevelBucketLevelActivityMetricsArgs.builder()
                            .enabled(true)
                            .build())
                        .build())
                    .build())
                .dataExport(StorageLensConfigurationStorageLensConfigurationDataExportArgs.builder()
                    .cloudWatchMetrics(StorageLensConfigurationStorageLensConfigurationDataExportCloudWatchMetricsArgs.builder()
                        .enabled(true)
                        .build())
                    .s3BucketDestination(StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationArgs.builder()
                        .accountId(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
                        .arn(aws_s3_bucket.target().arn())
                        .format("CSV")
                        .outputSchemaVersion("V_1")
                        .encryption(StorageLensConfigurationStorageLensConfigurationDataExportS3BucketDestinationEncryptionArgs.builder()
                            .sseS3s()
                            .build())
                        .build())
                    .build())
                .exclude(StorageLensConfigurationStorageLensConfigurationExcludeArgs.builder()
                    .buckets(                    
                        aws_s3_bucket.b1().arn(),
                        aws_s3_bucket.b2().arn())
                    .regions("us-east-2")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3control:StorageLensConfiguration
    properties:
      configId: example-1
      storageLensConfiguration:
        enabled: true
        accountLevel:
          activityMetrics:
            enabled: true
          bucketLevel:
            activityMetrics:
              enabled: true
        dataExport:
          cloudWatchMetrics:
            enabled: true
          s3BucketDestination:
            accountId: ${current.accountId}
            arn: ${aws_s3_bucket.target.arn}
            format: CSV
            outputSchemaVersion: V_1
            encryption:
              sseS3s:
                - {}
        exclude:
          buckets:
            - ${aws_s3_bucket.b1.arn}
            - ${aws_s3_bucket.b2.arn}
          regions:
            - us-east-2
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 Storage Lens configurations using the `account_id` and `config_id`, separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:s3control/storageLensConfiguration:StorageLensConfiguration example 123456789012:example-1
```
 