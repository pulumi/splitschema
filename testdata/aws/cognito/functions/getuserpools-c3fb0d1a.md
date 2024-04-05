Use this data source to get a list of cognito user pools.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const selectedRestApi = aws.apigateway.getRestApi({
    name: _var.api_gateway_name,
});
const selectedUserPools = aws.cognito.getUserPools({
    name: _var.cognito_user_pool_name,
});
const cognito = new aws.apigateway.Authorizer("cognito", {
    type: "COGNITO_USER_POOLS",
    restApi: selectedRestApi.then(selectedRestApi => selectedRestApi.id),
    providerArns: selectedUserPools.then(selectedUserPools => selectedUserPools.arns),
});
```
```python
import pulumi
import pulumi_aws as aws

selected_rest_api = aws.apigateway.get_rest_api(name=var["api_gateway_name"])
selected_user_pools = aws.cognito.get_user_pools(name=var["cognito_user_pool_name"])
cognito = aws.apigateway.Authorizer("cognito",
    type="COGNITO_USER_POOLS",
    rest_api=selected_rest_api.id,
    provider_arns=selected_user_pools.arns)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var selectedRestApi = Aws.ApiGateway.GetRestApi.Invoke(new()
    {
        Name = @var.Api_gateway_name,
    });

    var selectedUserPools = Aws.Cognito.GetUserPools.Invoke(new()
    {
        Name = @var.Cognito_user_pool_name,
    });

    var cognito = new Aws.ApiGateway.Authorizer("cognito", new()
    {
        Type = "COGNITO_USER_POOLS",
        RestApi = selectedRestApi.Apply(getRestApiResult => getRestApiResult.Id),
        ProviderArns = selectedUserPools.Apply(getUserPoolsResult => getUserPoolsResult.Arns),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		selectedRestApi, err := apigateway.LookupRestApi(ctx, &apigateway.LookupRestApiArgs{
			Name: _var.Api_gateway_name,
		}, nil)
		if err != nil {
			return err
		}
		selectedUserPools, err := cognito.GetUserPools(ctx, &cognito.GetUserPoolsArgs{
			Name: _var.Cognito_user_pool_name,
		}, nil)
		if err != nil {
			return err
		}
		_, err = apigateway.NewAuthorizer(ctx, "cognito", &apigateway.AuthorizerArgs{
			Type:         pulumi.String("COGNITO_USER_POOLS"),
			RestApi:      *pulumi.String(selectedRestApi.Id),
			ProviderArns: interface{}(selectedUserPools.Arns),
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
import com.pulumi.aws.apigateway.ApigatewayFunctions;
import com.pulumi.aws.apigateway.inputs.GetRestApiArgs;
import com.pulumi.aws.cognito.CognitoFunctions;
import com.pulumi.aws.cognito.inputs.GetUserPoolsArgs;
import com.pulumi.aws.apigateway.Authorizer;
import com.pulumi.aws.apigateway.AuthorizerArgs;
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
        final var selectedRestApi = ApigatewayFunctions.getRestApi(GetRestApiArgs.builder()
            .name(var_.api_gateway_name())
            .build());

        final var selectedUserPools = CognitoFunctions.getUserPools(GetUserPoolsArgs.builder()
            .name(var_.cognito_user_pool_name())
            .build());

        var cognito = new Authorizer("cognito", AuthorizerArgs.builder()        
            .type("COGNITO_USER_POOLS")
            .restApi(selectedRestApi.applyValue(getRestApiResult -> getRestApiResult.id()))
            .providerArns(selectedUserPools.applyValue(getUserPoolsResult -> getUserPoolsResult.arns()))
            .build());

    }
}
```
```yaml
resources:
  cognito:
    type: aws:apigateway:Authorizer
    properties:
      type: COGNITO_USER_POOLS
      restApi: ${selectedRestApi.id}
      providerArns: ${selectedUserPools.arns}
variables:
  selectedRestApi:
    fn::invoke:
      Function: aws:apigateway:getRestApi
      Arguments:
        name: ${var.api_gateway_name}
  selectedUserPools:
    fn::invoke:
      Function: aws:cognito:getUserPools
      Arguments:
        name: ${var.cognito_user_pool_name}
```
{{% /example %}}
{{% /examples %}}