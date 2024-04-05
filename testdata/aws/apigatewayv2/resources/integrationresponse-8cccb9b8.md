Manages an Amazon API Gateway Version 2 integration response.
More information can be found in the [Amazon API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-websocket-api.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.IntegrationResponse("example", {
    apiId: aws_apigatewayv2_api.example.id,
    integrationId: aws_apigatewayv2_integration.example.id,
    integrationResponseKey: "/200/",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigatewayv2.IntegrationResponse("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    integration_id=aws_apigatewayv2_integration["example"]["id"],
    integration_response_key="/200/")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.IntegrationResponse("example", new()
    {
        ApiId = aws_apigatewayv2_api.Example.Id,
        IntegrationId = aws_apigatewayv2_integration.Example.Id,
        IntegrationResponseKey = "/200/",
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
		_, err := apigatewayv2.NewIntegrationResponse(ctx, "example", &apigatewayv2.IntegrationResponseArgs{
			ApiId:                  pulumi.Any(aws_apigatewayv2_api.Example.Id),
			IntegrationId:          pulumi.Any(aws_apigatewayv2_integration.Example.Id),
			IntegrationResponseKey: pulumi.String("/200/"),
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
import com.pulumi.aws.apigatewayv2.IntegrationResponse;
import com.pulumi.aws.apigatewayv2.IntegrationResponseArgs;
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
        var example = new IntegrationResponse("example", IntegrationResponseArgs.builder()        
            .apiId(aws_apigatewayv2_api.example().id())
            .integrationId(aws_apigatewayv2_integration.example().id())
            .integrationResponseKey("/200/")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:IntegrationResponse
    properties:
      apiId: ${aws_apigatewayv2_api.example.id}
      integrationId: ${aws_apigatewayv2_integration.example.id}
      integrationResponseKey: /200/
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_apigatewayv2_integration_response` using the API identifier, integration identifier and integration response identifier. For example:

```sh
 $ pulumi import aws:apigatewayv2/integrationResponse:IntegrationResponse example aabbccddee/1122334/998877
```
 