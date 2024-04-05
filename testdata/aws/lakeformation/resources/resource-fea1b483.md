Registers a Lake Formation resource (e.g., S3 bucket) as managed by the Data Catalog. In other words, the S3 path is added to the data lake.

Choose a role that has read/write access to the chosen Amazon S3 path or use the service-linked role. When you register the S3 path, the service-linked role and a new inline policy are created on your behalf. Lake Formation adds the first path to the inline policy and attaches it to the service-linked role. When you register subsequent paths, Lake Formation adds the path to the existing policy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucket = aws.s3.getBucket({
    bucket: "an-example-bucket",
});
const exampleResource = new aws.lakeformation.Resource("exampleResource", {arn: exampleBucket.then(exampleBucket => exampleBucket.arn)});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket = aws.s3.get_bucket(bucket="an-example-bucket")
example_resource = aws.lakeformation.Resource("exampleResource", arn=example_bucket.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucket = Aws.S3.GetBucket.Invoke(new()
    {
        Bucket = "an-example-bucket",
    });

    var exampleResource = new Aws.LakeFormation.Resource("exampleResource", new()
    {
        Arn = exampleBucket.Apply(getBucketResult => getBucketResult.Arn),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucket, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
			Bucket: "an-example-bucket",
		}, nil)
		if err != nil {
			return err
		}
		_, err = lakeformation.NewResource(ctx, "exampleResource", &lakeformation.ResourceArgs{
			Arn: *pulumi.String(exampleBucket.Arn),
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
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.inputs.GetBucketArgs;
import com.pulumi.aws.lakeformation.Resource;
import com.pulumi.aws.lakeformation.ResourceArgs;
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
        final var exampleBucket = S3Functions.getBucket(GetBucketArgs.builder()
            .bucket("an-example-bucket")
            .build());

        var exampleResource = new Resource("exampleResource", ResourceArgs.builder()        
            .arn(exampleBucket.applyValue(getBucketResult -> getBucketResult.arn()))
            .build());

    }
}
```
```yaml
resources:
  exampleResource:
    type: aws:lakeformation:Resource
    properties:
      arn: ${exampleBucket.arn}
variables:
  exampleBucket:
    fn::invoke:
      Function: aws:s3:getBucket
      Arguments:
        bucket: an-example-bucket
```
{{% /example %}}
{{% /examples %}}