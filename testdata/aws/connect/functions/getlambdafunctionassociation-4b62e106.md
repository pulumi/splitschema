Provides details about a specific Connect Lambda Function Association.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.connect.getLambdaFunctionAssociation({
    functionArn: "arn:aws:lambda:us-west-2:123456789123:function:abcdefg",
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.get_lambda_function_association(function_arn="arn:aws:lambda:us-west-2:123456789123:function:abcdefg",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Connect.GetLambdaFunctionAssociation.Invoke(new()
    {
        FunctionArn = "arn:aws:lambda:us-west-2:123456789123:function:abcdefg",
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
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
		_, err := connect.LookupLambdaFunctionAssociation(ctx, &connect.LookupLambdaFunctionAssociationArgs{
			FunctionArn: "arn:aws:lambda:us-west-2:123456789123:function:abcdefg",
			InstanceId:  "aaaaaaaa-bbbb-cccc-dddd-111111111111",
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
import com.pulumi.aws.connect.ConnectFunctions;
import com.pulumi.aws.connect.inputs.GetLambdaFunctionAssociationArgs;
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
        final var example = ConnectFunctions.getLambdaFunctionAssociation(GetLambdaFunctionAssociationArgs.builder()
            .functionArn("arn:aws:lambda:us-west-2:123456789123:function:abcdefg")
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:connect:getLambdaFunctionAssociation
      Arguments:
        functionArn: arn:aws:lambda:us-west-2:123456789123:function:abcdefg
        instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
```
{{% /example %}}
{{% /examples %}}