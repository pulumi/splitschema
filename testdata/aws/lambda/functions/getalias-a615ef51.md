Provides information about a Lambda Alias.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const production = aws.lambda.getAlias({
    functionName: "my-lambda-func",
    name: "production",
});
```
```python
import pulumi
import pulumi_aws as aws

production = aws.lambda.get_alias(function_name="my-lambda-func",
    name="production")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var production = Aws.Lambda.GetAlias.Invoke(new()
    {
        FunctionName = "my-lambda-func",
        Name = "production",
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
		_, err := lambda.LookupAlias(ctx, &lambda.LookupAliasArgs{
			FunctionName: "my-lambda-func",
			Name:         "production",
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
import com.pulumi.aws.lambda.inputs.GetAliasArgs;
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
        final var production = LambdaFunctions.getAlias(GetAliasArgs.builder()
            .functionName("my-lambda-func")
            .name("production")
            .build());

    }
}
```
```yaml
variables:
  production:
    fn::invoke:
      Function: aws:lambda:getAlias
      Arguments:
        functionName: my-lambda-func
        name: production
```
{{% /example %}}
{{% /examples %}}