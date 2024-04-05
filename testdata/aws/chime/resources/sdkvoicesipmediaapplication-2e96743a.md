A ChimeSDKVoice SIP Media Application is a managed object that passes values from a SIP rule to a target AWS Lambda function.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.chime.SdkvoiceSipMediaApplication("example", {
    awsRegion: "us-east-1",
    endpoints: {
        lambdaArn: aws_lambda_function.test.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.chime.SdkvoiceSipMediaApplication("example",
    aws_region="us-east-1",
    endpoints=aws.chime.SdkvoiceSipMediaApplicationEndpointsArgs(
        lambda_arn=aws_lambda_function["test"]["arn"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Chime.SdkvoiceSipMediaApplication("example", new()
    {
        AwsRegion = "us-east-1",
        Endpoints = new Aws.Chime.Inputs.SdkvoiceSipMediaApplicationEndpointsArgs
        {
            LambdaArn = aws_lambda_function.Test.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/chime"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := chime.NewSdkvoiceSipMediaApplication(ctx, "example", &chime.SdkvoiceSipMediaApplicationArgs{
			AwsRegion: pulumi.String("us-east-1"),
			Endpoints: &chime.SdkvoiceSipMediaApplicationEndpointsArgs{
				LambdaArn: pulumi.Any(aws_lambda_function.Test.Arn),
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
import com.pulumi.aws.chime.SdkvoiceSipMediaApplication;
import com.pulumi.aws.chime.SdkvoiceSipMediaApplicationArgs;
import com.pulumi.aws.chime.inputs.SdkvoiceSipMediaApplicationEndpointsArgs;
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
        var example = new SdkvoiceSipMediaApplication("example", SdkvoiceSipMediaApplicationArgs.builder()        
            .awsRegion("us-east-1")
            .endpoints(SdkvoiceSipMediaApplicationEndpointsArgs.builder()
                .lambdaArn(aws_lambda_function.test().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:chime:SdkvoiceSipMediaApplication
    properties:
      awsRegion: us-east-1
      endpoints:
        lambdaArn: ${aws_lambda_function.test.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a ChimeSDKVoice SIP Media Application using the `id`. For example:

```sh
 $ pulumi import aws:chime/sdkvoiceSipMediaApplication:SdkvoiceSipMediaApplication example abcdef123456
```
 