Provides an S3 bucket accelerate configuration resource. See the [Requirements for using Transfer Acceleration](https://docs.aws.amazon.com/AmazonS3/latest/userguide/transfer-acceleration.html#transfer-acceleration-requirements) for more details.

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mybucket = new aws.s3.BucketV2("mybucket", {});
const example = new aws.s3.BucketAccelerateConfigurationV2("example", {
    bucket: mybucket.id,
    status: "Enabled",
});
```
```python
import pulumi
import pulumi_aws as aws

mybucket = aws.s3.BucketV2("mybucket")
example = aws.s3.BucketAccelerateConfigurationV2("example",
    bucket=mybucket.id,
    status="Enabled")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var mybucket = new Aws.S3.BucketV2("mybucket");

    var example = new Aws.S3.BucketAccelerateConfigurationV2("example", new()
    {
        Bucket = mybucket.Id,
        Status = "Enabled",
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
		mybucket, err := s3.NewBucketV2(ctx, "mybucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAccelerateConfigurationV2(ctx, "example", &s3.BucketAccelerateConfigurationV2Args{
			Bucket: mybucket.ID(),
			Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.s3.BucketAccelerateConfigurationV2;
import com.pulumi.aws.s3.BucketAccelerateConfigurationV2Args;
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
        var mybucket = new BucketV2("mybucket");

        var example = new BucketAccelerateConfigurationV2("example", BucketAccelerateConfigurationV2Args.builder()        
            .bucket(mybucket.id())
            .status("Enabled")
            .build());

    }
}
```
```yaml
resources:
  mybucket:
    type: aws:s3:BucketV2
  example:
    type: aws:s3:BucketAccelerateConfigurationV2
    properties:
      bucket: ${mybucket.id}
      status: Enabled
```
{{% /example %}}
{{% /examples %}}

## Import

If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

__Using `pulumi import` to import.__ For example:

If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, import using the `bucket`:

```sh
 $ pulumi import aws:s3/bucketAccelerateConfigurationV2:BucketAccelerateConfigurationV2 example bucket-name
```
 If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

```sh
 $ pulumi import aws:s3/bucketAccelerateConfigurationV2:BucketAccelerateConfigurationV2 example bucket-name,123456789012
```
 