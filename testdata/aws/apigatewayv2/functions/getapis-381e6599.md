Provides details about multiple Amazon API Gateway Version 2 APIs.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.apigatewayv2.getApis({
    protocolType: "HTTP",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.get_apis(protocol_type="HTTP")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ApiGatewayV2.GetApis.Invoke(new()
    {
        ProtocolType = "HTTP",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigatewayv2.GetApis(ctx, &apigatewayv2.GetApisArgs{
			ProtocolType: pulumi.StringRef("HTTP"),
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
import com.pulumi.aws.apigatewayv2.Apigatewayv2Functions;
import com.pulumi.aws.apigatewayv2.inputs.GetApisArgs;
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
        final var example = Apigatewayv2Functions.getApis(GetApisArgs.builder()
            .protocolType("HTTP")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:apigatewayv2:getApis
      Arguments:
        protocolType: HTTP
```
{{% /example %}}
{{% /examples %}}