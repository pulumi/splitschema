Creates a Lambda function alias. Creates an alias that points to the specified Lambda function version.

For information about Lambda and how to use it, see [What is AWS Lambda?](http://docs.aws.amazon.com/lambda/latest/dg/welcome.html)
For information about function aliases, see [CreateAlias](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateAlias.html) and [AliasRoutingConfiguration](https://docs.aws.amazon.com/lambda/latest/dg/API_AliasRoutingConfiguration.html) in the API docs.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testLambdaAlias = new aws.lambda.Alias("testLambdaAlias", {
    description: "a sample description",
    functionName: aws_lambda_function.lambda_function_test.arn,
    functionVersion: "1",
    routingConfig: {
        additionalVersionWeights: {
            "2": 0.5,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_lambda_alias = aws.lambda_.Alias("testLambdaAlias",
    description="a sample description",
    function_name=aws_lambda_function["lambda_function_test"]["arn"],
    function_version="1",
    routing_config=aws.lambda_.AliasRoutingConfigArgs(
        additional_version_weights={
            "2": 0.5,
        },
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testLambdaAlias = new Aws.Lambda.Alias("testLambdaAlias", new()
    {
        Description = "a sample description",
        FunctionName = aws_lambda_function.Lambda_function_test.Arn,
        FunctionVersion = "1",
        RoutingConfig = new Aws.Lambda.Inputs.AliasRoutingConfigArgs
        {
            AdditionalVersionWeights = 
            {
                { "2", 0.5 },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewAlias(ctx, "testLambdaAlias", &lambda.AliasArgs{
			Description:     pulumi.String("a sample description"),
			FunctionName:    pulumi.Any(aws_lambda_function.Lambda_function_test.Arn),
			FunctionVersion: pulumi.String("1"),
			RoutingConfig: &lambda.AliasRoutingConfigArgs{
				AdditionalVersionWeights: pulumi.Float64Map{
					"2": pulumi.Float64(0.5),
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
import com.pulumi.aws.lambda.Alias;
import com.pulumi.aws.lambda.AliasArgs;
import com.pulumi.aws.lambda.inputs.AliasRoutingConfigArgs;
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
        var testLambdaAlias = new Alias("testLambdaAlias", AliasArgs.builder()        
            .description("a sample description")
            .functionName(aws_lambda_function.lambda_function_test().arn())
            .functionVersion("1")
            .routingConfig(AliasRoutingConfigArgs.builder()
                .additionalVersionWeights(Map.of("2", 0.5))
                .build())
            .build());

    }
}
```
```yaml
resources:
  testLambdaAlias:
    type: aws:lambda:Alias
    properties:
      description: a sample description
      functionName: ${aws_lambda_function.lambda_function_test.arn}
      functionVersion: '1'
      routingConfig:
        additionalVersionWeights:
          '2': 0.5
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lambda Function Aliases using the `function_name/alias`. For example:

```sh
 $ pulumi import aws:lambda/alias:Alias test_lambda_alias my_test_lambda_function/my_alias
```
 