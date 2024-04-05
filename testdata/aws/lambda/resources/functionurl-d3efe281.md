Provides a Lambda function URL resource. A function URL is a dedicated HTTP(S) endpoint for a Lambda function.

See the [AWS Lambda documentation](https://docs.aws.amazon.com/lambda/latest/dg/lambda-urls.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testLatest = new aws.lambda.FunctionUrl("testLatest", {
    functionName: aws_lambda_function.test.function_name,
    authorizationType: "NONE",
});
const testLive = new aws.lambda.FunctionUrl("testLive", {
    functionName: aws_lambda_function.test.function_name,
    qualifier: "my_alias",
    authorizationType: "AWS_IAM",
    cors: {
        allowCredentials: true,
        allowOrigins: ["*"],
        allowMethods: ["*"],
        allowHeaders: [
            "date",
            "keep-alive",
        ],
        exposeHeaders: [
            "keep-alive",
            "date",
        ],
        maxAge: 86400,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_latest = aws.lambda_.FunctionUrl("testLatest",
    function_name=aws_lambda_function["test"]["function_name"],
    authorization_type="NONE")
test_live = aws.lambda_.FunctionUrl("testLive",
    function_name=aws_lambda_function["test"]["function_name"],
    qualifier="my_alias",
    authorization_type="AWS_IAM",
    cors=aws.lambda_.FunctionUrlCorsArgs(
        allow_credentials=True,
        allow_origins=["*"],
        allow_methods=["*"],
        allow_headers=[
            "date",
            "keep-alive",
        ],
        expose_headers=[
            "keep-alive",
            "date",
        ],
        max_age=86400,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testLatest = new Aws.Lambda.FunctionUrl("testLatest", new()
    {
        FunctionName = aws_lambda_function.Test.Function_name,
        AuthorizationType = "NONE",
    });

    var testLive = new Aws.Lambda.FunctionUrl("testLive", new()
    {
        FunctionName = aws_lambda_function.Test.Function_name,
        Qualifier = "my_alias",
        AuthorizationType = "AWS_IAM",
        Cors = new Aws.Lambda.Inputs.FunctionUrlCorsArgs
        {
            AllowCredentials = true,
            AllowOrigins = new[]
            {
                "*",
            },
            AllowMethods = new[]
            {
                "*",
            },
            AllowHeaders = new[]
            {
                "date",
                "keep-alive",
            },
            ExposeHeaders = new[]
            {
                "keep-alive",
                "date",
            },
            MaxAge = 86400,
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
		_, err := lambda.NewFunctionUrl(ctx, "testLatest", &lambda.FunctionUrlArgs{
			FunctionName:      pulumi.Any(aws_lambda_function.Test.Function_name),
			AuthorizationType: pulumi.String("NONE"),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewFunctionUrl(ctx, "testLive", &lambda.FunctionUrlArgs{
			FunctionName:      pulumi.Any(aws_lambda_function.Test.Function_name),
			Qualifier:         pulumi.String("my_alias"),
			AuthorizationType: pulumi.String("AWS_IAM"),
			Cors: &lambda.FunctionUrlCorsArgs{
				AllowCredentials: pulumi.Bool(true),
				AllowOrigins: pulumi.StringArray{
					pulumi.String("*"),
				},
				AllowMethods: pulumi.StringArray{
					pulumi.String("*"),
				},
				AllowHeaders: pulumi.StringArray{
					pulumi.String("date"),
					pulumi.String("keep-alive"),
				},
				ExposeHeaders: pulumi.StringArray{
					pulumi.String("keep-alive"),
					pulumi.String("date"),
				},
				MaxAge: pulumi.Int(86400),
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
import com.pulumi.aws.lambda.FunctionUrl;
import com.pulumi.aws.lambda.FunctionUrlArgs;
import com.pulumi.aws.lambda.inputs.FunctionUrlCorsArgs;
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
        var testLatest = new FunctionUrl("testLatest", FunctionUrlArgs.builder()        
            .functionName(aws_lambda_function.test().function_name())
            .authorizationType("NONE")
            .build());

        var testLive = new FunctionUrl("testLive", FunctionUrlArgs.builder()        
            .functionName(aws_lambda_function.test().function_name())
            .qualifier("my_alias")
            .authorizationType("AWS_IAM")
            .cors(FunctionUrlCorsArgs.builder()
                .allowCredentials(true)
                .allowOrigins("*")
                .allowMethods("*")
                .allowHeaders(                
                    "date",
                    "keep-alive")
                .exposeHeaders(                
                    "keep-alive",
                    "date")
                .maxAge(86400)
                .build())
            .build());

    }
}
```
```yaml
resources:
  testLatest:
    type: aws:lambda:FunctionUrl
    properties:
      functionName: ${aws_lambda_function.test.function_name}
      authorizationType: NONE
  testLive:
    type: aws:lambda:FunctionUrl
    properties:
      functionName: ${aws_lambda_function.test.function_name}
      qualifier: my_alias
      authorizationType: AWS_IAM
      cors:
        allowCredentials: true
        allowOrigins:
          - '*'
        allowMethods:
          - '*'
        allowHeaders:
          - date
          - keep-alive
        exposeHeaders:
          - keep-alive
          - date
        maxAge: 86400
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lambda function URLs using the `function_name` or `function_name/qualifier`. For example:

```sh
 $ pulumi import aws:lambda/functionUrl:FunctionUrl test_lambda_url my_test_lambda_function
```
 