Provides a S3 bucket resource.

> This resource provides functionality for managing S3 general purpose buckets in an AWS Partition. To manage Amazon S3 Express directory buckets, use the `aws_directory_bucket` resource. To manage [S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html), use the `aws.s3control.Bucket` resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Private Bucket With Tags

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3.BucketV2("example", {tags: {
    Environment: "Dev",
    Name: "My bucket",
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3.BucketV2("example", tags={
    "Environment": "Dev",
    "Name": "My bucket",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3.BucketV2("example", new()
    {
        Tags = 
        {
            { "Environment", "Dev" },
            { "Name", "My bucket" },
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
		_, err := s3.NewBucketV2(ctx, "example", &s3.BucketV2Args{
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Dev"),
				"Name":        pulumi.String("My bucket"),
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
import com.pulumi.aws.s3.BucketV2Args;
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
        var example = new BucketV2("example", BucketV2Args.builder()        
            .tags(Map.ofEntries(
                Map.entry("Environment", "Dev"),
                Map.entry("Name", "My bucket")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3:BucketV2
    properties:
      tags:
        Environment: Dev
        Name: My bucket
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 bucket using the `bucket`. For example:

```sh
 $ pulumi import aws:s3/bucketV2:BucketV2 bucket bucket-name
```
 