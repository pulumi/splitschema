Provides an API Gateway Gateway Response for a REST API Gateway.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.apigateway.RestApi("main", {});
const test = new aws.apigateway.Response("test", {
    restApiId: main.id,
    statusCode: "401",
    responseType: "UNAUTHORIZED",
    responseTemplates: {
        "application/json": "{\"message\":$context.error.messageString}",
    },
    responseParameters: {
        "gatewayresponse.header.Authorization": "'Basic'",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.apigateway.RestApi("main")
test = aws.apigateway.Response("test",
    rest_api_id=main.id,
    status_code="401",
    response_type="UNAUTHORIZED",
    response_templates={
        "application/json": "{\"message\":$context.error.messageString}",
    },
    response_parameters={
        "gatewayresponse.header.Authorization": "'Basic'",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.ApiGateway.RestApi("main");

    var test = new Aws.ApiGateway.Response("test", new()
    {
        RestApiId = main.Id,
        StatusCode = "401",
        ResponseType = "UNAUTHORIZED",
        ResponseTemplates = 
        {
            { "application/json", "{\"message\":$context.error.messageString}" },
        },
        ResponseParameters = 
        {
            { "gatewayresponse.header.Authorization", "'Basic'" },
        },
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
		main, err := apigateway.NewRestApi(ctx, "main", nil)
		if err != nil {
			return err
		}
		_, err = apigateway.NewResponse(ctx, "test", &apigateway.ResponseArgs{
			RestApiId:    main.ID(),
			StatusCode:   pulumi.String("401"),
			ResponseType: pulumi.String("UNAUTHORIZED"),
			ResponseTemplates: pulumi.StringMap{
				"application/json": pulumi.String("{\"message\":$context.error.messageString}"),
			},
			ResponseParameters: pulumi.StringMap{
				"gatewayresponse.header.Authorization": pulumi.String("'Basic'"),
			},
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
import com.pulumi.aws.apigateway.RestApi;
import com.pulumi.aws.apigateway.Response;
import com.pulumi.aws.apigateway.ResponseArgs;
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
        var main = new RestApi("main");

        var test = new Response("test", ResponseArgs.builder()        
            .restApiId(main.id())
            .statusCode("401")
            .responseType("UNAUTHORIZED")
            .responseTemplates(Map.of("application/json", "{\"message\":$context.error.messageString}"))
            .responseParameters(Map.of("gatewayresponse.header.Authorization", "'Basic'"))
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:apigateway:RestApi
  test:
    type: aws:apigateway:Response
    properties:
      restApiId: ${main.id}
      statusCode: '401'
      responseType: UNAUTHORIZED
      responseTemplates:
        application/json: '{"message":$context.error.messageString}'
      responseParameters:
        gatewayresponse.header.Authorization: '''Basic'''
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_api_gateway_gateway_response` using `REST-API-ID/RESPONSE-TYPE`. For example:

```sh
 $ pulumi import aws:apigateway/response:Response example 12345abcde/UNAUTHORIZED
```
 