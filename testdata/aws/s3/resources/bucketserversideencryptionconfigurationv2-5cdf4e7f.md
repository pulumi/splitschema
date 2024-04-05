Provides a S3 bucket server-side encryption configuration resource.

> **NOTE:** Destroying an `aws.s3.BucketServerSideEncryptionConfigurationV2` resource resets the bucket to [Amazon S3 bucket default encryption](https://docs.aws.amazon.com/AmazonS3/latest/userguide/default-encryption-faq.html).

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const mykey = new aws.kms.Key("mykey", {
    description: "This key is used to encrypt bucket objects",
    deletionWindowInDays: 10,
});
const mybucket = new aws.s3.BucketV2("mybucket", {});
const example = new aws.s3.BucketServerSideEncryptionConfigurationV2("example", {
    bucket: mybucket.id,
    rules: [{
        applyServerSideEncryptionByDefault: {
            kmsMasterKeyId: mykey.arn,
            sseAlgorithm: "aws:kms",
        },
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

mykey = aws.kms.Key("mykey",
    description="This key is used to encrypt bucket objects",
    deletion_window_in_days=10)
mybucket = aws.s3.BucketV2("mybucket")
example = aws.s3.BucketServerSideEncryptionConfigurationV2("example",
    bucket=mybucket.id,
    rules=[aws.s3.BucketServerSideEncryptionConfigurationV2RuleArgs(
        apply_server_side_encryption_by_default=aws.s3.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs(
            kms_master_key_id=mykey.arn,
            sse_algorithm="aws:kms",
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var mykey = new Aws.Kms.Key("mykey", new()
    {
        Description = "This key is used to encrypt bucket objects",
        DeletionWindowInDays = 10,
    });

    var mybucket = new Aws.S3.BucketV2("mybucket");

    var example = new Aws.S3.BucketServerSideEncryptionConfigurationV2("example", new()
    {
        Bucket = mybucket.Id,
        Rules = new[]
        {
            new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationV2RuleArgs
            {
                ApplyServerSideEncryptionByDefault = new Aws.S3.Inputs.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs
                {
                    KmsMasterKeyId = mykey.Arn,
                    SseAlgorithm = "aws:kms",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		mykey, err := kms.NewKey(ctx, "mykey", &kms.KeyArgs{
			Description:          pulumi.String("This key is used to encrypt bucket objects"),
			DeletionWindowInDays: pulumi.Int(10),
		})
		if err != nil {
			return err
		}
		mybucket, err := s3.NewBucketV2(ctx, "mybucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketServerSideEncryptionConfigurationV2(ctx, "example", &s3.BucketServerSideEncryptionConfigurationV2Args{
			Bucket: mybucket.ID(),
			Rules: s3.BucketServerSideEncryptionConfigurationV2RuleArray{
				&s3.BucketServerSideEncryptionConfigurationV2RuleArgs{
					ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs{
						KmsMasterKeyId: mykey.Arn,
						SseAlgorithm:   pulumi.String("aws:kms"),
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketServerSideEncryptionConfigurationV2;
import com.pulumi.aws.s3.BucketServerSideEncryptionConfigurationV2Args;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationV2RuleArgs;
import com.pulumi.aws.s3.inputs.BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs;
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
        var mykey = new Key("mykey", KeyArgs.builder()        
            .description("This key is used to encrypt bucket objects")
            .deletionWindowInDays(10)
            .build());

        var mybucket = new BucketV2("mybucket");

        var example = new BucketServerSideEncryptionConfigurationV2("example", BucketServerSideEncryptionConfigurationV2Args.builder()        
            .bucket(mybucket.id())
            .rules(BucketServerSideEncryptionConfigurationV2RuleArgs.builder()
                .applyServerSideEncryptionByDefault(BucketServerSideEncryptionConfigurationV2RuleApplyServerSideEncryptionByDefaultArgs.builder()
                    .kmsMasterKeyId(mykey.arn())
                    .sseAlgorithm("aws:kms")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  mykey:
    type: aws:kms:Key
    properties:
      description: This key is used to encrypt bucket objects
      deletionWindowInDays: 10
  mybucket:
    type: aws:s3:BucketV2
  example:
    type: aws:s3:BucketServerSideEncryptionConfigurationV2
    properties:
      bucket: ${mybucket.id}
      rules:
        - applyServerSideEncryptionByDefault:
            kmsMasterKeyId: ${mykey.arn}
            sseAlgorithm: aws:kms
```
{{% /example %}}
{{% /examples %}}

## Import

If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

__Using `pulumi import` to import__ S3 bucket server-side encryption configuration using the `bucket` or using the `bucket` and `expected_bucket_owner` separated by a comma (`,`). For example:

If the owner (account ID) of the source bucket is the same account used to configure the AWS Provider, import using the `bucket`:

```sh
 $ pulumi import aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2 example bucket-name
```
 If the owner (account ID) of the source bucket differs from the account used to configure the AWS Provider, import using the `bucket` and `expected_bucket_owner` separated by a comma (`,`):

```sh
 $ pulumi import aws:s3/bucketServerSideEncryptionConfigurationV2:BucketServerSideEncryptionConfigurationV2 example bucket-name,123456789012
```
 