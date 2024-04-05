Manages an API Gateway Request Validator.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigateway.RequestValidator("example", {
    restApi: aws_api_gateway_rest_api.example.id,
    validateRequestBody: true,
    validateRequestParameters: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigateway.RequestValidator("example",
    rest_api=aws_api_gateway_rest_api["example"]["id"],
    validate_request_body=True,
    validate_request_parameters=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGateway.RequestValidator("example", new()
    {
        RestApi = aws_api_gateway_rest_api.Example.Id,
        ValidateRequestBody = true,
        ValidateRequestParameters = true,
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
		_, err := apigateway.NewRequestValidator(ctx, "example", &apigateway.RequestValidatorArgs{
			RestApi:                   pulumi.Any(aws_api_gateway_rest_api.Example.Id),
			ValidateRequestBody:       pulumi.Bool(true),
			ValidateRequestParameters: pulumi.Bool(true),
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
import com.pulumi.aws.apigateway.RequestValidator;
import com.pulumi.aws.apigateway.RequestValidatorArgs;
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
        var example = new RequestValidator("example", RequestValidatorArgs.builder()        
            .restApi(aws_api_gateway_rest_api.example().id())
            .validateRequestBody(true)
            .validateRequestParameters(true)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigateway:RequestValidator
    properties:
      restApi: ${aws_api_gateway_rest_api.example.id}
      validateRequestBody: true
      validateRequestParameters: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_api_gateway_request_validator` using `REST-API-ID/REQUEST-VALIDATOR-ID`. For example:

```sh
 $ pulumi import aws:apigateway/requestValidator:RequestValidator example 12345abcde/67890fghij
```
 