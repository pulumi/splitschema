{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.apigateway.getExport({
    restApiId: aws_api_gateway_stage.example.rest_api_id,
    stageName: aws_api_gateway_stage.example.stage_name,
    exportType: "oas30",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigateway.get_export(rest_api_id=aws_api_gateway_stage["example"]["rest_api_id"],
    stage_name=aws_api_gateway_stage["example"]["stage_name"],
    export_type="oas30")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ApiGateway.GetExport.Invoke(new()
    {
        RestApiId = aws_api_gateway_stage.Example.Rest_api_id,
        StageName = aws_api_gateway_stage.Example.Stage_name,
        ExportType = "oas30",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigateway.GetExport(ctx, &apigateway.GetExportArgs{
			RestApiId:  aws_api_gateway_stage.Example.Rest_api_id,
			StageName:  aws_api_gateway_stage.Example.Stage_name,
			ExportType: "oas30",
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
import com.pulumi.aws.apigateway.ApigatewayFunctions;
import com.pulumi.aws.apigateway.inputs.GetExportArgs;
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
        final var example = ApigatewayFunctions.getExport(GetExportArgs.builder()
            .restApiId(aws_api_gateway_stage.example().rest_api_id())
            .stageName(aws_api_gateway_stage.example().stage_name())
            .exportType("oas30")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:apigateway:getExport
      Arguments:
        restApiId: ${aws_api_gateway_stage.example.rest_api_id}
        stageName: ${aws_api_gateway_stage.example.stage_name}
        exportType: oas30
```
{{% /example %}}
{{% /examples %}}