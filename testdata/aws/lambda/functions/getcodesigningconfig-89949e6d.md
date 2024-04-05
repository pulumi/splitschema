Provides information about a Lambda Code Signing Config. A code signing configuration defines a list of allowed signing profiles and defines the code-signing validation policy (action to be taken if deployment validation checks fail).

For information about Lambda code signing configurations and how to use them, see [configuring code signing for Lambda functions](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const existingCsc = aws.lambda.getCodeSigningConfig({
    arn: `arn:aws:lambda:${_var.aws_region}:${_var.aws_account}:code-signing-config:csc-0f6c334abcdea4d8b`,
});
```
```python
import pulumi
import pulumi_aws as aws

existing_csc = aws.lambda.get_code_signing_config(arn=f"arn:aws:lambda:{var['aws_region']}:{var['aws_account']}:code-signing-config:csc-0f6c334abcdea4d8b")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var existingCsc = Aws.Lambda.GetCodeSigningConfig.Invoke(new()
    {
        Arn = $"arn:aws:lambda:{@var.Aws_region}:{@var.Aws_account}:code-signing-config:csc-0f6c334abcdea4d8b",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.LookupCodeSigningConfig(ctx, &lambda.LookupCodeSigningConfigArgs{
			Arn: fmt.Sprintf("arn:aws:lambda:%v:%v:code-signing-config:csc-0f6c334abcdea4d8b", _var.Aws_region, _var.Aws_account),
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
import com.pulumi.aws.lambda.inputs.GetCodeSigningConfigArgs;
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
        final var existingCsc = LambdaFunctions.getCodeSigningConfig(GetCodeSigningConfigArgs.builder()
            .arn(String.format("arn:aws:lambda:%s:%s:code-signing-config:csc-0f6c334abcdea4d8b", var_.aws_region(),var_.aws_account()))
            .build());

    }
}
```
```yaml
variables:
  existingCsc:
    fn::invoke:
      Function: aws:lambda:getCodeSigningConfig
      Arguments:
        arn: arn:aws:lambda:${var.aws_region}:${var.aws_account}:code-signing-config:csc-0f6c334abcdea4d8b
```
{{% /example %}}
{{% /examples %}}