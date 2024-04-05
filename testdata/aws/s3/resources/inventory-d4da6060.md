Provides a S3 bucket [inventory configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html) resource.

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}
### Add inventory configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testBucketV2 = new aws.s3.BucketV2("testBucketV2", {});
const inventory = new aws.s3.BucketV2("inventory", {});
const testInventory = new aws.s3.Inventory("testInventory", {
    bucket: testBucketV2.id,
    includedObjectVersions: "All",
    schedule: {
        frequency: "Daily",
    },
    destination: {
        bucket: {
            format: "ORC",
            bucketArn: inventory.arn,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_bucket_v2 = aws.s3.BucketV2("testBucketV2")
inventory = aws.s3.BucketV2("inventory")
test_inventory = aws.s3.Inventory("testInventory",
    bucket=test_bucket_v2.id,
    included_object_versions="All",
    schedule=aws.s3.InventoryScheduleArgs(
        frequency="Daily",
    ),
    destination=aws.s3.InventoryDestinationArgs(
        bucket=aws.s3.InventoryDestinationBucketArgs(
            format="ORC",
            bucket_arn=inventory.arn,
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
    var testBucketV2 = new Aws.S3.BucketV2("testBucketV2");

    var inventory = new Aws.S3.BucketV2("inventory");

    var testInventory = new Aws.S3.Inventory("testInventory", new()
    {
        Bucket = testBucketV2.Id,
        IncludedObjectVersions = "All",
        Schedule = new Aws.S3.Inputs.InventoryScheduleArgs
        {
            Frequency = "Daily",
        },
        Destination = new Aws.S3.Inputs.InventoryDestinationArgs
        {
            Bucket = new Aws.S3.Inputs.InventoryDestinationBucketArgs
            {
                Format = "ORC",
                BucketArn = inventory.Arn,
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
		testBucketV2, err := s3.NewBucketV2(ctx, "testBucketV2", nil)
		if err != nil {
			return err
		}
		inventory, err := s3.NewBucketV2(ctx, "inventory", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewInventory(ctx, "testInventory", &s3.InventoryArgs{
			Bucket:                 testBucketV2.ID(),
			IncludedObjectVersions: pulumi.String("All"),
			Schedule: &s3.InventoryScheduleArgs{
				Frequency: pulumi.String("Daily"),
			},
			Destination: &s3.InventoryDestinationArgs{
				Bucket: &s3.InventoryDestinationBucketArgs{
					Format:    pulumi.String("ORC"),
					BucketArn: inventory.Arn,
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
import com.pulumi.aws.s3.Inventory;
import com.pulumi.aws.s3.InventoryArgs;
import com.pulumi.aws.s3.inputs.InventoryScheduleArgs;
import com.pulumi.aws.s3.inputs.InventoryDestinationArgs;
import com.pulumi.aws.s3.inputs.InventoryDestinationBucketArgs;
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
        var testBucketV2 = new BucketV2("testBucketV2");

        var inventory = new BucketV2("inventory");

        var testInventory = new Inventory("testInventory", InventoryArgs.builder()        
            .bucket(testBucketV2.id())
            .includedObjectVersions("All")
            .schedule(InventoryScheduleArgs.builder()
                .frequency("Daily")
                .build())
            .destination(InventoryDestinationArgs.builder()
                .bucket(InventoryDestinationBucketArgs.builder()
                    .format("ORC")
                    .bucketArn(inventory.arn())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testBucketV2:
    type: aws:s3:BucketV2
  inventory:
    type: aws:s3:BucketV2
  testInventory:
    type: aws:s3:Inventory
    properties:
      bucket: ${testBucketV2.id}
      includedObjectVersions: All
      schedule:
        frequency: Daily
      destination:
        bucket:
          format: ORC
          bucketArn: ${inventory.arn}
```
{{% /example %}}
{{% example %}}
### Add inventory configuration with S3 object prefix

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.s3.BucketV2("test", {});
const inventory = new aws.s3.BucketV2("inventory", {});
const test_prefix = new aws.s3.Inventory("test-prefix", {
    bucket: test.id,
    includedObjectVersions: "All",
    schedule: {
        frequency: "Daily",
    },
    filter: {
        prefix: "documents/",
    },
    destination: {
        bucket: {
            format: "ORC",
            bucketArn: inventory.arn,
            prefix: "inventory",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.s3.BucketV2("test")
inventory = aws.s3.BucketV2("inventory")
test_prefix = aws.s3.Inventory("test-prefix",
    bucket=test.id,
    included_object_versions="All",
    schedule=aws.s3.InventoryScheduleArgs(
        frequency="Daily",
    ),
    filter=aws.s3.InventoryFilterArgs(
        prefix="documents/",
    ),
    destination=aws.s3.InventoryDestinationArgs(
        bucket=aws.s3.InventoryDestinationBucketArgs(
            format="ORC",
            bucket_arn=inventory.arn,
            prefix="inventory",
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
    var test = new Aws.S3.BucketV2("test");

    var inventory = new Aws.S3.BucketV2("inventory");

    var test_prefix = new Aws.S3.Inventory("test-prefix", new()
    {
        Bucket = test.Id,
        IncludedObjectVersions = "All",
        Schedule = new Aws.S3.Inputs.InventoryScheduleArgs
        {
            Frequency = "Daily",
        },
        Filter = new Aws.S3.Inputs.InventoryFilterArgs
        {
            Prefix = "documents/",
        },
        Destination = new Aws.S3.Inputs.InventoryDestinationArgs
        {
            Bucket = new Aws.S3.Inputs.InventoryDestinationBucketArgs
            {
                Format = "ORC",
                BucketArn = inventory.Arn,
                Prefix = "inventory",
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
		test, err := s3.NewBucketV2(ctx, "test", nil)
		if err != nil {
			return err
		}
		inventory, err := s3.NewBucketV2(ctx, "inventory", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewInventory(ctx, "test-prefix", &s3.InventoryArgs{
			Bucket:                 test.ID(),
			IncludedObjectVersions: pulumi.String("All"),
			Schedule: &s3.InventoryScheduleArgs{
				Frequency: pulumi.String("Daily"),
			},
			Filter: &s3.InventoryFilterArgs{
				Prefix: pulumi.String("documents/"),
			},
			Destination: &s3.InventoryDestinationArgs{
				Bucket: &s3.InventoryDestinationBucketArgs{
					Format:    pulumi.String("ORC"),
					BucketArn: inventory.Arn,
					Prefix:    pulumi.String("inventory"),
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
import com.pulumi.aws.s3.Inventory;
import com.pulumi.aws.s3.InventoryArgs;
import com.pulumi.aws.s3.inputs.InventoryScheduleArgs;
import com.pulumi.aws.s3.inputs.InventoryFilterArgs;
import com.pulumi.aws.s3.inputs.InventoryDestinationArgs;
import com.pulumi.aws.s3.inputs.InventoryDestinationBucketArgs;
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
        var test = new BucketV2("test");

        var inventory = new BucketV2("inventory");

        var test_prefix = new Inventory("test-prefix", InventoryArgs.builder()        
            .bucket(test.id())
            .includedObjectVersions("All")
            .schedule(InventoryScheduleArgs.builder()
                .frequency("Daily")
                .build())
            .filter(InventoryFilterArgs.builder()
                .prefix("documents/")
                .build())
            .destination(InventoryDestinationArgs.builder()
                .bucket(InventoryDestinationBucketArgs.builder()
                    .format("ORC")
                    .bucketArn(inventory.arn())
                    .prefix("inventory")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:s3:BucketV2
  inventory:
    type: aws:s3:BucketV2
  test-prefix:
    type: aws:s3:Inventory
    properties:
      bucket: ${test.id}
      includedObjectVersions: All
      schedule:
        frequency: Daily
      filter:
        prefix: documents/
      destination:
        bucket:
          format: ORC
          bucketArn: ${inventory.arn}
          prefix: inventory
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 bucket inventory configurations using `bucket:inventory`. For example:

```sh
 $ pulumi import aws:s3/inventory:Inventory my-bucket-entire-bucket my-bucket:EntireBucket
```
 