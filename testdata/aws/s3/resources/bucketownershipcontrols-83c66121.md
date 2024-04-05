Provides a resource to manage S3 Bucket Ownership Controls. For more information, see the [S3 Developer Guide](https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html).

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketOwnershipControls = new aws.s3.BucketOwnershipControls("exampleBucketOwnershipControls", {
    bucket: exampleBucketV2.id,
    rule: {
        objectOwnership: "BucketOwnerPreferred",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_ownership_controls = aws.s3.BucketOwnershipControls("exampleBucketOwnershipControls",
    bucket=example_bucket_v2.id,
    rule=aws.s3.BucketOwnershipControlsRuleArgs(
        object_ownership="BucketOwnerPreferred",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketOwnershipControls = new Aws.S3.BucketOwnershipControls("exampleBucketOwnershipControls", new()
    {
        Bucket = exampleBucketV2.Id,
        Rule = new Aws.S3.Inputs.BucketOwnershipControlsRuleArgs
        {
            ObjectOwnership = "BucketOwnerPreferred",
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
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketOwnershipControls(ctx, "exampleBucketOwnershipControls", &s3.BucketOwnershipControlsArgs{
			Bucket: exampleBucketV2.ID(),
			Rule: &s3.BucketOwnershipControlsRuleArgs{
				ObjectOwnership: pulumi.String("BucketOwnerPreferred"),
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
import com.pulumi.aws.s3.BucketOwnershipControls;
import com.pulumi.aws.s3.BucketOwnershipControlsArgs;
import com.pulumi.aws.s3.inputs.BucketOwnershipControlsRuleArgs;
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

        var exampleBucketOwnershipControls = new BucketOwnershipControls("exampleBucketOwnershipControls", BucketOwnershipControlsArgs.builder()        
            .bucket(exampleBucketV2.id())
            .rule(BucketOwnershipControlsRuleArgs.builder()
                .objectOwnership("BucketOwnerPreferred")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketOwnershipControls:
    type: aws:s3:BucketOwnershipControls
    properties:
      bucket: ${exampleBucketV2.id}
      rule:
        objectOwnership: BucketOwnerPreferred
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 Bucket Ownership Controls using S3 Bucket name. For example:

```sh
 $ pulumi import aws:s3/bucketOwnershipControls:BucketOwnershipControls example my-bucket
```
 