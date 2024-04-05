Provides a S3 bucket [metrics configuration](http://docs.aws.amazon.com/AmazonS3/latest/dev/metrics-configurations.html) resource.

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}
### Add metrics configuration for entire S3 bucket

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketV2("example", {});
const example_entire_bucket = new aws.s3.BucketMetric("example-entire-bucket", {bucket: example.id});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketV2("example")
example_entire_bucket = aws.s3.BucketMetric("example-entire-bucket", bucket=example.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketV2("example");

    var example_entire_bucket = new Aws.S3.BucketMetric("example-entire-bucket", new()
    {
        Bucket = example.Id,
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
		_, err = s3.NewBucketMetric(ctx, "example-entire-bucket", &s3.BucketMetricArgs{
			Bucket: example.ID(),
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
import com.pulumi.aws.s3.BucketMetric;
import com.pulumi.aws.s3.BucketMetricArgs;
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

        var example_entire_bucket = new BucketMetric("example-entire-bucket", BucketMetricArgs.builder()        
            .bucket(example.id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketV2
  example-entire-bucket:
    type: aws:s3:BucketMetric
    properties:
      bucket: ${example.id}
```
{{% /example %}}
{{% example %}}
### Add metrics configuration with S3 object filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketV2("example", {});
const example_filtered = new aws.s3.BucketMetric("example-filtered", {
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
example_filtered = aws.s3.BucketMetric("example-filtered",
    bucket=example.id,
    filter=aws.s3.BucketMetricFilterArgs(
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

    var example_filtered = new Aws.S3.BucketMetric("example-filtered", new()
    {
        Bucket = example.Id,
        Filter = new Aws.S3.Inputs.BucketMetricFilterArgs
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
		_, err = s3.NewBucketMetric(ctx, "example-filtered", &s3.BucketMetricArgs{
			Bucket: example.ID(),
			Filter: &s3.BucketMetricFilterArgs{
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
import com.pulumi.aws.s3.BucketMetric;
import com.pulumi.aws.s3.BucketMetricArgs;
import com.pulumi.aws.s3.inputs.BucketMetricFilterArgs;
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

        var example_filtered = new BucketMetric("example-filtered", BucketMetricArgs.builder()        
            .bucket(example.id())
            .filter(BucketMetricFilterArgs.builder()
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
  example:
    type: aws:s3:BucketV2
  example-filtered:
    type: aws:s3:BucketMetric
    properties:
      bucket: ${example.id}
      filter:
        prefix: documents/
        tags:
          priority: high
          class: blue
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 bucket metric configurations using `bucket:metric`. For example:

```sh
 $ pulumi import aws:s3/bucketMetric:BucketMetric my-bucket-entire-bucket my-bucket:EntireBucket
```
 