Manages S3 bucket-level Public Access Block configuration. For more information about these settings, see the [AWS S3 Block Public Access documentation](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html).

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketPublicAccessBlock = new aws.s3.BucketPublicAccessBlock("exampleBucketPublicAccessBlock", {
    bucket: exampleBucketV2.id,
    blockPublicAcls: true,
    blockPublicPolicy: true,
    ignorePublicAcls: true,
    restrictPublicBuckets: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_public_access_block = aws.s3.BucketPublicAccessBlock("exampleBucketPublicAccessBlock",
    bucket=example_bucket_v2.id,
    block_public_acls=True,
    block_public_policy=True,
    ignore_public_acls=True,
    restrict_public_buckets=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketPublicAccessBlock = new Aws.S3.BucketPublicAccessBlock("exampleBucketPublicAccessBlock", new()
    {
        Bucket = exampleBucketV2.Id,
        BlockPublicAcls = true,
        BlockPublicPolicy = true,
        IgnorePublicAcls = true,
        RestrictPublicBuckets = true,
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
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketPublicAccessBlock(ctx, "exampleBucketPublicAccessBlock", &s3.BucketPublicAccessBlockArgs{
			Bucket:                exampleBucketV2.ID(),
			BlockPublicAcls:       pulumi.Bool(true),
			BlockPublicPolicy:     pulumi.Bool(true),
			IgnorePublicAcls:      pulumi.Bool(true),
			RestrictPublicBuckets: pulumi.Bool(true),
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
import com.pulumi.aws.s3.BucketPublicAccessBlock;
import com.pulumi.aws.s3.BucketPublicAccessBlockArgs;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleBucketPublicAccessBlock = new BucketPublicAccessBlock("exampleBucketPublicAccessBlock", BucketPublicAccessBlockArgs.builder()        
            .bucket(exampleBucketV2.id())
            .blockPublicAcls(true)
            .blockPublicPolicy(true)
            .ignorePublicAcls(true)
            .restrictPublicBuckets(true)
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketPublicAccessBlock:
    type: aws:s3:BucketPublicAccessBlock
    properties:
      bucket: ${exampleBucketV2.id}
      blockPublicAcls: true
      blockPublicPolicy: true
      ignorePublicAcls: true
      restrictPublicBuckets: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_s3_bucket_public_access_block` using the bucket name. For example:

```sh
 $ pulumi import aws:s3/bucketPublicAccessBlock:BucketPublicAccessBlock example my-bucket
```
 