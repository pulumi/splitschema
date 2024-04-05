The bucket policy data source returns IAM policy of an S3 bucket.

{{% examples %}}
## Example Usage
{{% example %}}

The following example retrieves IAM policy of a specified S3 bucket.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.s3.getBucketPolicy({
    bucket: "example-bucket-name",
});
export const foo = example.then(example => example.policy);
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.get_bucket_policy(bucket="example-bucket-name")
pulumi.export("foo", example.policy)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.S3.GetBucketPolicy.Invoke(new()
    {
        Bucket = "example-bucket-name",
    });

    return new Dictionary<string, object?>
    {
        ["foo"] = example.Apply(getBucketPolicyResult => getBucketPolicyResult.Policy),
    };
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
		example, err := s3.LookupBucketPolicy(ctx, &s3.LookupBucketPolicyArgs{
			Bucket: "example-bucket-name",
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("foo", example.Policy)
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.inputs.GetBucketPolicyArgs;
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
        final var example = S3Functions.getBucketPolicy(GetBucketPolicyArgs.builder()
            .bucket("example-bucket-name")
            .build());

        ctx.export("foo", example.applyValue(getBucketPolicyResult -> getBucketPolicyResult.policy()));
    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:s3:getBucketPolicy
      Arguments:
        bucket: example-bucket-name
outputs:
  foo: ${example.policy}
```
{{% /example %}}
{{% /examples %}}