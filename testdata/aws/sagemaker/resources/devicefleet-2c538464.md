Provides a SageMaker Device Fleet resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.DeviceFleet("example", {
    deviceFleetName: "example",
    roleArn: aws_iam_role.test.arn,
    outputConfig: {
        s3OutputLocation: `s3://${aws_s3_bucket.example.bucket}/prefix/`,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.DeviceFleet("example",
    device_fleet_name="example",
    role_arn=aws_iam_role["test"]["arn"],
    output_config=aws.sagemaker.DeviceFleetOutputConfigArgs(
        s3_output_location=f"s3://{aws_s3_bucket['example']['bucket']}/prefix/",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.DeviceFleet("example", new()
    {
        DeviceFleetName = "example",
        RoleArn = aws_iam_role.Test.Arn,
        OutputConfig = new Aws.Sagemaker.Inputs.DeviceFleetOutputConfigArgs
        {
            S3OutputLocation = $"s3://{aws_s3_bucket.Example.Bucket}/prefix/",
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewDeviceFleet(ctx, "example", &sagemaker.DeviceFleetArgs{
			DeviceFleetName: pulumi.String("example"),
			RoleArn:         pulumi.Any(aws_iam_role.Test.Arn),
			OutputConfig: &sagemaker.DeviceFleetOutputConfigArgs{
				S3OutputLocation: pulumi.String(fmt.Sprintf("s3://%v/prefix/", aws_s3_bucket.Example.Bucket)),
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
import com.pulumi.aws.sagemaker.DeviceFleet;
import com.pulumi.aws.sagemaker.DeviceFleetArgs;
import com.pulumi.aws.sagemaker.inputs.DeviceFleetOutputConfigArgs;
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
        var example = new DeviceFleet("example", DeviceFleetArgs.builder()        
            .deviceFleetName("example")
            .roleArn(aws_iam_role.test().arn())
            .outputConfig(DeviceFleetOutputConfigArgs.builder()
                .s3OutputLocation(String.format("s3://%s/prefix/", aws_s3_bucket.example().bucket()))
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:DeviceFleet
    properties:
      deviceFleetName: example
      roleArn: ${aws_iam_role.test.arn}
      outputConfig:
        s3OutputLocation: s3://${aws_s3_bucket.example.bucket}/prefix/
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Device Fleets using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/deviceFleet:DeviceFleet example my-fleet
```
 