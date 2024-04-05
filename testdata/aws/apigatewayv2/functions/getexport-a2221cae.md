Exports a definition of an API in a particular output format and specification.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.apigatewayv2.getExport({
    apiId: aws_apigatewayv2_route.test.api_id,
    specification: "OAS30",
    outputType: "JSON",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.apigatewayv2.get_export(api_id=aws_apigatewayv2_route["test"]["api_id"],
    specification="OAS30",
    output_type="JSON")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.ApiGatewayV2.GetExport.Invoke(new()
    {
        ApiId = aws_apigatewayv2_route.Test.Api_id,
        Specification = "OAS30",
        OutputType = "JSON",
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
		_, err := apigatewayv2.GetExport(ctx, &apigatewayv2.GetExportArgs{
			ApiId:         aws_apigatewayv2_route.Test.Api_id,
			Specification: "OAS30",
			OutputType:    "JSON",
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
import com.pulumi.aws.apigatewayv2.inputs.GetExportArgs;
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
        final var test = Apigatewayv2Functions.getExport(GetExportArgs.builder()
            .apiId(aws_apigatewayv2_route.test().api_id())
            .specification("OAS30")
            .outputType("JSON")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:apigatewayv2:getExport
      Arguments:
        apiId: ${aws_apigatewayv2_route.test.api_id}
        specification: OAS30
        outputType: JSON
```
{{% /example %}}
{{% /examples %}}