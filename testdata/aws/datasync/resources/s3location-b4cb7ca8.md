Manages an S3 Location within AWS DataSync.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.datasync.S3Location("example", {
    s3BucketArn: aws_s3_bucket.example.arn,
    subdirectory: "/example/prefix",
    s3Config: {
        bucketAccessRoleArn: aws_iam_role.example.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.datasync.S3Location("example",
    s3_bucket_arn=aws_s3_bucket["example"]["arn"],
    subdirectory="/example/prefix",
    s3_config=aws.datasync.S3LocationS3ConfigArgs(
        bucket_access_role_arn=aws_iam_role["example"]["arn"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataSync.S3Location("example", new()
    {
        S3BucketArn = aws_s3_bucket.Example.Arn,
        Subdirectory = "/example/prefix",
        S3Config = new Aws.DataSync.Inputs.S3LocationS3ConfigArgs
        {
            BucketAccessRoleArn = aws_iam_role.Example.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/datasync"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datasync.NewS3Location(ctx, "example", &datasync.S3LocationArgs{
			S3BucketArn:  pulumi.Any(aws_s3_bucket.Example.Arn),
			Subdirectory: pulumi.String("/example/prefix"),
			S3Config: &datasync.S3LocationS3ConfigArgs{
				BucketAccessRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
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
import com.pulumi.aws.datasync.S3Location;
import com.pulumi.aws.datasync.S3LocationArgs;
import com.pulumi.aws.datasync.inputs.S3LocationS3ConfigArgs;
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
        var example = new S3Location("example", S3LocationArgs.builder()        
            .s3BucketArn(aws_s3_bucket.example().arn())
            .subdirectory("/example/prefix")
            .s3Config(S3LocationS3ConfigArgs.builder()
                .bucketAccessRoleArn(aws_iam_role.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:datasync:S3Location
    properties:
      s3BucketArn: ${aws_s3_bucket.example.arn}
      subdirectory: /example/prefix
      s3Config:
        bucketAccessRoleArn: ${aws_iam_role.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_datasync_location_s3` using the DataSync Task Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:datasync/s3Location:S3Location example arn:aws:datasync:us-east-1:123456789012:location/loc-12345678901234567
```
 