Provides details about a specific API Gateway Authorizer.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.apigateway.getAuthorizer({
    restApiId: aws_api_gateway_rest_api.example.id,
    authorizerId: data.aws_api_gateway_authorizers.example.ids[0],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apigateway.get_authorizer(rest_api_id=aws_api_gateway_rest_api["example"]["id"],
    authorizer_id=data["aws_api_gateway_authorizers"]["example"]["ids"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ApiGateway.GetAuthorizer.Invoke(new()
    {
        RestApiId = aws_api_gateway_rest_api.Example.Id,
        AuthorizerId = data.Aws_api_gateway_authorizers.Example.Ids[0],
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
		_, err := apigateway.LookupAuthorizer(ctx, &apigateway.LookupAuthorizerArgs{
			RestApiId:    aws_api_gateway_rest_api.Example.Id,
			AuthorizerId: data.Aws_api_gateway_authorizers.Example.Ids[0],
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
import com.pulumi.aws.apigateway.inputs.GetAuthorizerArgs;
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
        final var example = ApigatewayFunctions.getAuthorizer(GetAuthorizerArgs.builder()
            .restApiId(aws_api_gateway_rest_api.example().id())
            .authorizerId(data.aws_api_gateway_authorizers().example().ids()[0])
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:apigateway:getAuthorizer
      Arguments:
        restApiId: ${aws_api_gateway_rest_api.example.id}
        authorizerId: ${data.aws_api_gateway_authorizers.example.ids[0]}
```
{{% /example %}}
{{% /examples %}}