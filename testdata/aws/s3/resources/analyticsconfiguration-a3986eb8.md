Provides a S3 bucket [analytics configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html) resource.

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}
### Add analytics configuration for entire S3 bucket and export results to a second S3 bucket

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketV2("example", {});
const analytics = new aws.s3.BucketV2("analytics", {});
const example_entire_bucket = new aws.s3.AnalyticsConfiguration("example-entire-bucket", {
    bucket: example.id,
    storageClassAnalysis: {
        dataExport: {
            destination: {
                s3BucketDestination: {
                    bucketArn: analytics.arn,
                },
            },
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketV2("example")
analytics = aws.s3.BucketV2("analytics")
example_entire_bucket = aws.s3.AnalyticsConfiguration("example-entire-bucket",
    bucket=example.id,
    storage_class_analysis=aws.s3.AnalyticsConfigurationStorageClassAnalysisArgs(
        data_export=aws.s3.AnalyticsConfigurationStorageClassAnalysisDataExportArgs(
            destination=aws.s3.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationArgs(
                s3_bucket_destination=aws.s3.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestinationArgs(
                    bucket_arn=analytics.arn,
                ),
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
    var example = new Aws.S3.BucketV2("example");

    var analytics = new Aws.S3.BucketV2("analytics");

    var example_entire_bucket = new Aws.S3.AnalyticsConfiguration("example-entire-bucket", new()
    {
        Bucket = example.Id,
        StorageClassAnalysis = new Aws.S3.Inputs.AnalyticsConfigurationStorageClassAnalysisArgs
        {
            DataExport = new Aws.S3.Inputs.AnalyticsConfigurationStorageClassAnalysisDataExportArgs
            {
                Destination = new Aws.S3.Inputs.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationArgs
                {
                    S3BucketDestination = new Aws.S3.Inputs.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestinationArgs
                    {
                        BucketArn = analytics.Arn,
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := s3.NewBucketV2(ctx, "example", nil)
		if err != nil {
			return err
		}
		analytics, err := s3.NewBucketV2(ctx, "analytics", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewAnalyticsConfiguration(ctx, "example-entire-bucket", &s3.AnalyticsConfigurationArgs{
			Bucket: example.ID(),
			StorageClassAnalysis: &s3.AnalyticsConfigurationStorageClassAnalysisArgs{
				DataExport: &s3.AnalyticsConfigurationStorageClassAnalysisDataExportArgs{
					Destination: &s3.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationArgs{
						S3BucketDestination: &s3.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestinationArgs{
							BucketArn: analytics.Arn,
						},
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.AnalyticsConfiguration;
import com.pulumi.aws.s3.AnalyticsConfigurationArgs;
import com.pulumi.aws.s3.inputs.AnalyticsConfigurationStorageClassAnalysisArgs;
import com.pulumi.aws.s3.inputs.AnalyticsConfigurationStorageClassAnalysisDataExportArgs;
import com.pulumi.aws.s3.inputs.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationArgs;
import com.pulumi.aws.s3.inputs.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestinationArgs;
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
        var example = new BucketV2("example");

        var analytics = new BucketV2("analytics");

        var example_entire_bucket = new AnalyticsConfiguration("example-entire-bucket", AnalyticsConfigurationArgs.builder()        
            .bucket(example.id())
            .storageClassAnalysis(AnalyticsConfigurationStorageClassAnalysisArgs.builder()
                .dataExport(AnalyticsConfigurationStorageClassAnalysisDataExportArgs.builder()
                    .destination(AnalyticsConfigurationStorageClassAnalysisDataExportDestinationArgs.builder()
                        .s3BucketDestination(AnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestinationArgs.builder()
                            .bucketArn(analytics.arn())
                            .build())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example-entire-bucket:
    type: aws:s3:AnalyticsConfiguration
    properties:
      bucket: ${example.id}
      storageClassAnalysis:
        dataExport:
          destination:
            s3BucketDestination:
              bucketArn: ${analytics.arn}
  example:
    type: aws:s3:BucketV2
  analytics:
    type: aws:s3:BucketV2
```
{{% /example %}}
{{% example %}}
### Add analytics configuration with S3 object filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketV2("example", {});
const example_filtered = new aws.s3.AnalyticsConfiguration("example-filtered", {
    bucket: example.id,
    filter: {
        prefix: "documents/",
        tags: {
            priority: "high",
            "class": "blue",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketV2("example")
example_filtered = aws.s3.AnalyticsConfiguration("example-filtered",
    bucket=example.id,
    filter=aws.s3.AnalyticsConfigurationFilterArgs(
        prefix="documents/",
        tags={
            "priority": "high",
            "class": "blue",
        },
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketV2("example");

    var example_filtered = new Aws.S3.AnalyticsConfiguration("example-filtered", new()
    {
        Bucket = example.Id,
        Filter = new Aws.S3.Inputs.AnalyticsConfigurationFilterArgs
        {
            Prefix = "documents/",
            Tags = 
            {
                { "priority", "high" },
                { "class", "blue" },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := s3.NewBucketV2(ctx, "example", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewAnalyticsConfiguration(ctx, "example-filtered", &s3.AnalyticsConfigurationArgs{
			Bucket: example.ID(),
			Filter: &s3.AnalyticsConfigurationFilterArgs{
				Prefix: pulumi.String("documents/"),
				Tags: pulumi.StringMap{
					"priority": pulumi.String("high"),
					"class":    pulumi.String("blue"),
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.AnalyticsConfiguration;
import com.pulumi.aws.s3.AnalyticsConfigurationArgs;
import com.pulumi.aws.s3.inputs.AnalyticsConfigurationFilterArgs;
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
        var example = new BucketV2("example");

        var example_filtered = new AnalyticsConfiguration("example-filtered", AnalyticsConfigurationArgs.builder()        
            .bucket(example.id())
            .filter(AnalyticsConfigurationFilterArgs.builder()
                .prefix("documents/")
                .tags(Map.ofEntries(
                    Map.entry("priority", "high"),
                    Map.entry("class", "blue")
                ))
                .build())
            .build());

    }
}
```
```yaml
resources:
  example-filtered:
    type: aws:s3:AnalyticsConfiguration
    properties:
      bucket: ${example.id}
      filter:
        prefix: documents/
        tags:
          priority: high
          class: blue
  example:
    type: aws:s3:BucketV2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 bucket analytics configurations using `bucket:analytics`. For example:

```sh
 $ pulumi import aws:s3/analyticsConfiguration:AnalyticsConfiguration my-bucket-entire-bucket my-bucket:EntireBucket
```
 