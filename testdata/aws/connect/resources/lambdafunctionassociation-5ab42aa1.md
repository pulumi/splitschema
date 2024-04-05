Provides an Amazon Connect Lambda Function Association. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html) and [Invoke AWS Lambda functions](https://docs.aws.amazon.com/connect/latest/adminguide/connect-lambda-functions.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.LambdaFunctionAssociation("example", {
    functionArn: aws_lambda_function.example.arn,
    instanceId: aws_connect_instance.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.LambdaFunctionAssociation("example",
    function_arn=aws_lambda_function["example"]["arn"],
    instance_id=aws_connect_instance["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.LambdaFunctionAssociation("example", new()
    {
        FunctionArn = aws_lambda_function.Example.Arn,
        InstanceId = aws_connect_instance.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewLambdaFunctionAssociation(ctx, "example", &connect.LambdaFunctionAssociationArgs{
			FunctionArn: pulumi.Any(aws_lambda_function.Example.Arn),
			InstanceId:  pulumi.Any(aws_connect_instance.Example.Id),
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
import com.pulumi.aws.connect.LambdaFunctionAssociation;
import com.pulumi.aws.connect.LambdaFunctionAssociationArgs;
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
        var example = new LambdaFunctionAssociation("example", LambdaFunctionAssociationArgs.builder()        
            .functionArn(aws_lambda_function.example().arn())
            .instanceId(aws_connect_instance.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:LambdaFunctionAssociation
    properties:
      functionArn: ${aws_lambda_function.example.arn}
      instanceId: ${aws_connect_instance.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_connect_lambda_function_association` using the `instance_id` and `function_arn` separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:connect/lambdaFunctionAssociation:LambdaFunctionAssociation example aaaaaaaa-bbbb-cccc-dddd-111111111111,arn:aws:lambda:us-west-2:123456789123:function:example
```
 