Provides an S3 bucket (server access) logging resource. For more information, see [Logging requests using server access logging](https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerLogs.html)
in the AWS S3 User Guide.

> **Note:** Amazon S3 supports server access logging, AWS CloudTrail, or a combination of both. Refer to the [Logging options for Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/userguide/logging-with-S3.html)
to decide which method meets your requirements.

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketAclV2 = new aws.s3.BucketAclV2("exampleBucketAclV2", {
    bucket: exampleBucketV2.id,
    acl: "private",
});
const logBucket = new aws.s3.BucketV2("logBucket", {});
const logBucketAcl = new aws.s3.BucketAclV2("logBucketAcl", {
    bucket: logBucket.id,
    acl: "log-delivery-write",
});
const exampleBucketLoggingV2 = new aws.s3.BucketLoggingV2("exampleBucketLoggingV2", {
    bucket: exampleBucketV2.id,
    targetBucket: logBucket.id,
    targetPrefix: "log/",
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_acl_v2 = aws.s3.BucketAclV2("exampleBucketAclV2",
    bucket=example_bucket_v2.id,
    acl="private")
log_bucket = aws.s3.BucketV2("logBucket")
log_bucket_acl = aws.s3.BucketAclV2("logBucketAcl",
    bucket=log_bucket.id,
    acl="log-delivery-write")
example_bucket_logging_v2 = aws.s3.BucketLoggingV2("exampleBucketLoggingV2",
    bucket=example_bucket_v2.id,
    target_bucket=log_bucket.id,
    target_prefix="log/")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketAclV2 = new Aws.S3.BucketAclV2("exampleBucketAclV2", new()
    {
        Bucket = exampleBucketV2.Id,
        Acl = "private",
    });

    var logBucket = new Aws.S3.BucketV2("logBucket");

    var logBucketAcl = new Aws.S3.BucketAclV2("logBucketAcl", new()
    {
        Bucket = logBucket.Id,
        Acl = "log-delivery-write",
    });

    var exampleBucketLoggingV2 = new Aws.S3.BucketLoggingV2("exampleBucketLoggingV2", new()
    {
        Bucket = exampleBucketV2.Id,
        TargetBucket = logBucket.Id,
        TargetPrefix = "log/",
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
		_, err = s3.NewBucketAclV2(ctx, "exampleBucketAclV2", &s3.BucketAclV2Args{
			Bucket: exampleBucketV2.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		logBucket, err := s3.NewBucketV2(ctx, "logBucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "logBucketAcl", &s3.BucketAclV2Args{
			Bucket: logBucket.ID(),
			Acl:    pulumi.String("log-delivery-write"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketLoggingV2(ctx, "exampleBucketLoggingV2", &s3.BucketLoggingV2Args{
			Bucket:       exampleBucketV2.ID(),
			TargetBucket: logBucket.ID(),
			TargetPrefix: pulumi.String("log/"),
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
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.s3.BucketLoggingV2;
import com.pulumi.aws.s3.BucketLoggingV2Args;
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

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .acl("private")
            .build());

        var logBucket = new BucketV2("logBucket");

        var logBucketAcl = new BucketAclV2("logBucketAcl", BucketAclV2Args.builder()        
            .bucket(logBucket.id())
            .acl("log-delivery-write")
            .build());

        var exampleBucketLoggingV2 = new BucketLoggingV2("exampleBucketLoggingV2", BucketLoggingV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .targetBucket(logBucket.id())
            .targetPrefix("log/")
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${exampleBucketV2.id}
      acl: private
  logBucket:
    type: aws:s3:BucketV2
  logBucketAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${logBucket.id}
      acl: log-delivery-write
  exampleBucketLoggingV2:
    type: aws:s3:BucketLoggingV2
    properties:
      bucket: ${exampleBucketV2.id}
      targetBucket: ${logBucket.id}
      targetPrefix: log/
```
{{% /example %}}
{{% /examples %}}

## Import

If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

__Using `pulumi import` to import__ S3 bucket logging using the `bucket` or using the `bucket` and `expected_bucket_owner` separated by a comma (`,`). For example:

If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, import using the `bucket`:

```sh
 $ pulumi import aws:s3/bucketLoggingV2:BucketLoggingV2 example bucket-name
```
 If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

```sh
 $ pulumi import aws:s3/bucketLoggingV2:BucketLoggingV2 example bucket-name,123456789012
```
 