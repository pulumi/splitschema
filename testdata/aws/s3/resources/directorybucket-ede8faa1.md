Provides an Amazon S3 Express directory bucket resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.DirectoryBucket("example", {
    bucket: "example--usw2-az1--x-s3",
    location: {
        name: "usw2-az1",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.DirectoryBucket("example",
    bucket="example--usw2-az1--x-s3",
    location=aws.s3.DirectoryBucketLocationArgs(
        name="usw2-az1",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.DirectoryBucket("example", new()
    {
        Bucket = "example--usw2-az1--x-s3",
        Location = new Aws.S3.Inputs.DirectoryBucketLocationArgs
        {
            Name = "usw2-az1",
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
		_, err := s3.NewDirectoryBucket(ctx, "example", &s3.DirectoryBucketArgs{
			Bucket: pulumi.String("example--usw2-az1--x-s3"),
			Location: &s3.DirectoryBucketLocationArgs{
				Name: pulumi.String("usw2-az1"),
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
import com.pulumi.aws.s3.DirectoryBucket;
import com.pulumi.aws.s3.DirectoryBucketArgs;
import com.pulumi.aws.s3.inputs.DirectoryBucketLocationArgs;
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
        var example = new DirectoryBucket("example", DirectoryBucketArgs.builder()        
            .bucket("example--usw2-az1--x-s3")
            .location(DirectoryBucketLocationArgs.builder()
                .name("usw2-az1")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:DirectoryBucket
    properties:
      bucket: example--usw2-az1--x-s3
      location:
        name: usw2-az1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 bucket using `bucket`. For example:

```sh
 $ pulumi import aws:s3/directoryBucket:DirectoryBucket example example--usw2-az1--x-s3
```
 