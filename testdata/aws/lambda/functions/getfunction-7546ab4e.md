Provides information about a Lambda Function.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const functionName = config.require("functionName");
const existing = aws.lambda.getFunction({
    functionName: functionName,
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
function_name = config.require("functionName")
existing = aws.lambda.get_function(function_name=function_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var functionName = config.Require("functionName");
    var existing = Aws.Lambda.GetFunction.Invoke(new()
    {
        FunctionName = functionName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		functionName := cfg.Require("functionName")
		_, err := lambda.LookupFunction(ctx, &lambda.LookupFunctionArgs{
			FunctionName: functionName,
		}, nil)
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
import com.pulumi.aws.lambda.LambdaFunctions;
import com.pulumi.aws.lambda.inputs.GetFunctionArgs;
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
        final var config = ctx.config();
        final var functionName = config.get("functionName");
        final var existing = LambdaFunctions.getFunction(GetFunctionArgs.builder()
            .functionName(functionName)
            .build());

    }
}
```
```yaml
configuration:
  functionName:
    type: string
variables:
  existing:
    fn::invoke:
      Function: aws:lambda:getFunction
      Arguments:
        functionName: ${functionName}
```
{{% /example %}}
{{% /examples %}}