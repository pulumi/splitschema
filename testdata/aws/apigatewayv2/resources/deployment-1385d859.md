Manages an Amazon API Gateway Version 2 deployment.
More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

> **Note:** Creating a deployment for an API requires at least one `aws.apigatewayv2.Route` resource associated with that API. To avoid race conditions when all resources are being created together, you need to add implicit resource references via the `triggers` argument or explicit resource references using the [resource `dependsOn` meta-argument](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson).


{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Deployment("example", {
    apiId: aws_apigatewayv2_api.example.id,
    description: "Example deployment",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.Deployment("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    description="Example deployment")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.Deployment("example", new()
    {
        ApiId = aws_apigatewayv2_api.Example.Id,
        Description = "Example deployment",
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
		_, err := apigatewayv2.NewDeployment(ctx, "example", &apigatewayv2.DeploymentArgs{
			ApiId:       pulumi.Any(aws_apigatewayv2_api.Example.Id),
			Description: pulumi.String("Example deployment"),
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
import com.pulumi.aws.apigatewayv2.Deployment;
import com.pulumi.aws.apigatewayv2.DeploymentArgs;
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
        var example = new Deployment("example", DeploymentArgs.builder()        
            .apiId(aws_apigatewayv2_api.example().id())
            .description("Example deployment")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:Deployment
    properties:
      apiId: ${aws_apigatewayv2_api.example.id}
      description: Example deployment
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_apigatewayv2_deployment` using the API identifier and deployment identifier. For example:

```sh
 $ pulumi import aws:apigatewayv2/deployment:Deployment example aabbccddee/1122334
```
 The `triggers` argument cannot be imported.

