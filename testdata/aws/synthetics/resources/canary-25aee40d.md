Provides a Synthetics Canary resource.

> **NOTE:** When you create a canary, AWS creates supporting implicit resources. See the Amazon CloudWatch Synthetics documentation on [DeleteCanary](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_DeleteCanary.html) for a full list. Neither AWS nor this provider deletes these implicit resources automatically when the canary is deleted. Before deleting a canary, ensure you have all the information about the canary that you need to delete the implicit resources using the AWS Console, or AWS CLI.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const some = new aws.synthetics.Canary("some", {
    artifactS3Location: "s3://some-bucket/",
    executionRoleArn: "some-role",
    handler: "exports.handler",
    runtimeVersion: "syn-1.0",
    schedule: {
        expression: "rate(0 minute)",
    },
    zipFile: "test-fixtures/lambdatest.zip",
});
```
```python
import pulumi
import pulumi_aws as aws

some = aws.synthetics.Canary("some",
    artifact_s3_location="s3://some-bucket/",
    execution_role_arn="some-role",
    handler="exports.handler",
    runtime_version="syn-1.0",
    schedule=aws.synthetics.CanaryScheduleArgs(
        expression="rate(0 minute)",
    ),
    zip_file="test-fixtures/lambdatest.zip")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var some = new Aws.Synthetics.Canary("some", new()
    {
        ArtifactS3Location = "s3://some-bucket/",
        ExecutionRoleArn = "some-role",
        Handler = "exports.handler",
        RuntimeVersion = "syn-1.0",
        Schedule = new Aws.Synthetics.Inputs.CanaryScheduleArgs
        {
            Expression = "rate(0 minute)",
        },
        ZipFile = "test-fixtures/lambdatest.zip",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/synthetics"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := synthetics.NewCanary(ctx, "some", &synthetics.CanaryArgs{
			ArtifactS3Location: pulumi.String("s3://some-bucket/"),
			ExecutionRoleArn:   pulumi.String("some-role"),
			Handler:            pulumi.String("exports.handler"),
			RuntimeVersion:     pulumi.String("syn-1.0"),
			Schedule: &synthetics.CanaryScheduleArgs{
				Expression: pulumi.String("rate(0 minute)"),
			},
			ZipFile: pulumi.String("test-fixtures/lambdatest.zip"),
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
import com.pulumi.aws.synthetics.Canary;
import com.pulumi.aws.synthetics.CanaryArgs;
import com.pulumi.aws.synthetics.inputs.CanaryScheduleArgs;
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
        var some = new Canary("some", CanaryArgs.builder()        
            .artifactS3Location("s3://some-bucket/")
            .executionRoleArn("some-role")
            .handler("exports.handler")
            .runtimeVersion("syn-1.0")
            .schedule(CanaryScheduleArgs.builder()
                .expression("rate(0 minute)")
                .build())
            .zipFile("test-fixtures/lambdatest.zip")
            .build());

    }
}
```
```yaml
resources:
  some:
    type: aws:synthetics:Canary
    properties:
      artifactS3Location: s3://some-bucket/
      executionRoleArn: some-role
      handler: exports.handler
      runtimeVersion: syn-1.0
      schedule:
        expression: rate(0 minute)
      zipFile: test-fixtures/lambdatest.zip
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Synthetics Canaries using the `name`. For example:

```sh
 $ pulumi import aws:synthetics/canary:Canary some some-canary
```
 