Provides a HTTP Method for an API Gateway Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDemoAPI = new aws.apigateway.RestApi("myDemoAPI", {description: "This is my API for demonstration purposes"});
const myDemoResource = new aws.apigateway.Resource("myDemoResource", {
    restApi: myDemoAPI.id,
    parentId: myDemoAPI.rootResourceId,
    pathPart: "mydemoresource",
});
const myDemoMethod = new aws.apigateway.Method("myDemoMethod", {
    restApi: myDemoAPI.id,
    resourceId: myDemoResource.id,
    httpMethod: "GET",
    authorization: "NONE",
});
```
```python
import pulumi
import pulumi_aws as aws

my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
my_demo_resource = aws.apigateway.Resource("myDemoResource",
    rest_api=my_demo_api.id,
    parent_id=my_demo_api.root_resource_id,
    path_part="mydemoresource")
my_demo_method = aws.apigateway.Method("myDemoMethod",
    rest_api=my_demo_api.id,
    resource_id=my_demo_resource.id,
    http_method="GET",
    authorization="NONE")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myDemoAPI = new Aws.ApiGateway.RestApi("myDemoAPI", new()
    {
        Description = "This is my API for demonstration purposes",
    });

    var myDemoResource = new Aws.ApiGateway.Resource("myDemoResource", new()
    {
        RestApi = myDemoAPI.Id,
        ParentId = myDemoAPI.RootResourceId,
        PathPart = "mydemoresource",
    });

    var myDemoMethod = new Aws.ApiGateway.Method("myDemoMethod", new()
    {
        RestApi = myDemoAPI.Id,
        ResourceId = myDemoResource.Id,
        HttpMethod = "GET",
        Authorization = "NONE",
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
		myDemoAPI, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
			Description: pulumi.String("This is my API for demonstration purposes"),
		})
		if err != nil {
			return err
		}
		myDemoResource, err := apigateway.NewResource(ctx, "myDemoResource", &apigateway.ResourceArgs{
			RestApi:  myDemoAPI.ID(),
			ParentId: myDemoAPI.RootResourceId,
			PathPart: pulumi.String("mydemoresource"),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewMethod(ctx, "myDemoMethod", &apigateway.MethodArgs{
			RestApi:       myDemoAPI.ID(),
			ResourceId:    myDemoResource.ID(),
			HttpMethod:    pulumi.String("GET"),
			Authorization: pulumi.String("NONE"),
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
import com.pulumi.aws.apigateway.RestApiArgs;
import com.pulumi.aws.apigateway.Resource;
import com.pulumi.aws.apigateway.ResourceArgs;
import com.pulumi.aws.apigateway.Method;
import com.pulumi.aws.apigateway.MethodArgs;
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
        var myDemoAPI = new RestApi("myDemoAPI", RestApiArgs.builder()        
            .description("This is my API for demonstration purposes")
            .build());

        var myDemoResource = new Resource("myDemoResource", ResourceArgs.builder()        
            .restApi(myDemoAPI.id())
            .parentId(myDemoAPI.rootResourceId())
            .pathPart("mydemoresource")
            .build());

        var myDemoMethod = new Method("myDemoMethod", MethodArgs.builder()        
            .restApi(myDemoAPI.id())
            .resourceId(myDemoResource.id())
            .httpMethod("GET")
            .authorization("NONE")
            .build());

    }
}
```
```yaml
resources:
  myDemoAPI:
    type: aws:apigateway:RestApi
    properties:
      description: This is my API for demonstration purposes
  myDemoResource:
    type: aws:apigateway:Resource
    properties:
      restApi: ${myDemoAPI.id}
      parentId: ${myDemoAPI.rootResourceId}
      pathPart: mydemoresource
  myDemoMethod:
    type: aws:apigateway:Method
    properties:
      restApi: ${myDemoAPI.id}
      resourceId: ${myDemoResource.id}
      httpMethod: GET
      authorization: NONE
```
{{% /example %}}
{{% /examples %}}
## Usage with Cognito User Pool Authorizer

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const cognitoUserPoolName = config.requireObject("cognitoUserPoolName");
const thisUserPools = aws.cognito.getUserPools({
    name: cognitoUserPoolName,
});
const thisRestApi = new aws.apigateway.RestApi("thisRestApi", {});
const thisResource = new aws.apigateway.Resource("thisResource", {
    restApi: thisRestApi.id,
    parentId: thisRestApi.rootResourceId,
    pathPart: "{proxy+}",
});
const thisAuthorizer = new aws.apigateway.Authorizer("thisAuthorizer", {
    type: "COGNITO_USER_POOLS",
    restApi: thisRestApi.id,
    providerArns: thisUserPools.then(thisUserPools => thisUserPools.arns),
});
const any = new aws.apigateway.Method("any", {
    restApi: thisRestApi.id,
    resourceId: thisResource.id,
    httpMethod: "ANY",
    authorization: "COGNITO_USER_POOLS",
    authorizerId: thisAuthorizer.id,
    requestParameters: {
        "method.request.path.proxy": true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
cognito_user_pool_name = config.require_object("cognitoUserPoolName")
this_user_pools = aws.cognito.get_user_pools(name=cognito_user_pool_name)
this_rest_api = aws.apigateway.RestApi("thisRestApi")
this_resource = aws.apigateway.Resource("thisResource",
    rest_api=this_rest_api.id,
    parent_id=this_rest_api.root_resource_id,
    path_part="{proxy+}")
this_authorizer = aws.apigateway.Authorizer("thisAuthorizer",
    type="COGNITO_USER_POOLS",
    rest_api=this_rest_api.id,
    provider_arns=this_user_pools.arns)
any = aws.apigateway.Method("any",
    rest_api=this_rest_api.id,
    resource_id=this_resource.id,
    http_method="ANY",
    authorization="COGNITO_USER_POOLS",
    authorizer_id=this_authorizer.id,
    request_parameters={
        "method.request.path.proxy": True,
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var cognitoUserPoolName = config.RequireObject<dynamic>("cognitoUserPoolName");
    var thisUserPools = Aws.Cognito.GetUserPools.Invoke(new()
    {
        Name = cognitoUserPoolName,
    });

    var thisRestApi = new Aws.ApiGateway.RestApi("thisRestApi");

    var thisResource = new Aws.ApiGateway.Resource("thisResource", new()
    {
        RestApi = thisRestApi.Id,
        ParentId = thisRestApi.RootResourceId,
        PathPart = "{proxy+}",
    });

    var thisAuthorizer = new Aws.ApiGateway.Authorizer("thisAuthorizer", new()
    {
        Type = "COGNITO_USER_POOLS",
        RestApi = thisRestApi.Id,
        ProviderArns = thisUserPools.Apply(getUserPoolsResult => getUserPoolsResult.Arns),
    });

    var any = new Aws.ApiGateway.Method("any", new()
    {
        RestApi = thisRestApi.Id,
        ResourceId = thisResource.Id,
        HttpMethod = "ANY",
        Authorization = "COGNITO_USER_POOLS",
        AuthorizerId = thisAuthorizer.Id,
        RequestParameters = 
        {
            { "method.request.path.proxy", true },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		cognitoUserPoolName := cfg.RequireObject("cognitoUserPoolName")
		thisUserPools, err := cognito.GetUserPools(ctx, &cognito.GetUserPoolsArgs{
			Name: cognitoUserPoolName,
		}, nil)
		if err != nil {
			return err
		}
		thisRestApi, err := apigateway.NewRestApi(ctx, "thisRestApi", nil)
		if err != nil {
			return err
		}
		thisResource, err := apigateway.NewResource(ctx, "thisResource", &apigateway.ResourceArgs{
			RestApi:  thisRestApi.ID(),
			ParentId: thisRestApi.RootResourceId,
			PathPart: pulumi.String("{proxy+}"),
		})
		if err != nil {
			return err
		}
		thisAuthorizer, err := apigateway.NewAuthorizer(ctx, "thisAuthorizer", &apigateway.AuthorizerArgs{
			Type:         pulumi.String("COGNITO_USER_POOLS"),
			RestApi:      thisRestApi.ID(),
			ProviderArns: interface{}(thisUserPools.Arns),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewMethod(ctx, "any", &apigateway.MethodArgs{
			RestApi:       thisRestApi.ID(),
			ResourceId:    thisResource.ID(),
			HttpMethod:    pulumi.String("ANY"),
			Authorization: pulumi.String("COGNITO_USER_POOLS"),
			AuthorizerId:  thisAuthorizer.ID(),
			RequestParameters: pulumi.BoolMap{
				"method.request.path.proxy": pulumi.Bool(true),
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
import com.pulumi.aws.cognito.CognitoFunctions;
import com.pulumi.aws.cognito.inputs.GetUserPoolsArgs;
import com.pulumi.aws.apigateway.RestApi;
import com.pulumi.aws.apigateway.Resource;
import com.pulumi.aws.apigateway.ResourceArgs;
import com.pulumi.aws.apigateway.Authorizer;
import com.pulumi.aws.apigateway.AuthorizerArgs;
import com.pulumi.aws.apigateway.Method;
import com.pulumi.aws.apigateway.MethodArgs;
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
        final var config = ctx.config();
        final var cognitoUserPoolName = config.get("cognitoUserPoolName");
        final var thisUserPools = CognitoFunctions.getUserPools(GetUserPoolsArgs.builder()
            .name(cognitoUserPoolName)
            .build());

        var thisRestApi = new RestApi("thisRestApi");

        var thisResource = new Resource("thisResource", ResourceArgs.builder()        
            .restApi(thisRestApi.id())
            .parentId(thisRestApi.rootResourceId())
            .pathPart("{proxy+}")
            .build());

        var thisAuthorizer = new Authorizer("thisAuthorizer", AuthorizerArgs.builder()        
            .type("COGNITO_USER_POOLS")
            .restApi(thisRestApi.id())
            .providerArns(thisUserPools.applyValue(getUserPoolsResult -> getUserPoolsResult.arns()))
            .build());

        var any = new Method("any", MethodArgs.builder()        
            .restApi(thisRestApi.id())
            .resourceId(thisResource.id())
            .httpMethod("ANY")
            .authorization("COGNITO_USER_POOLS")
            .authorizerId(thisAuthorizer.id())
            .requestParameters(Map.of("method.request.path.proxy", true))
            .build());

    }
}
```
```yaml
configuration:
  cognitoUserPoolName:
    type: dynamic
resources:
  thisRestApi:
    type: aws:apigateway:RestApi
  thisResource:
    type: aws:apigateway:Resource
    properties:
      restApi: ${thisRestApi.id}
      parentId: ${thisRestApi.rootResourceId}
      pathPart: '{proxy+}'
  thisAuthorizer:
    type: aws:apigateway:Authorizer
    properties:
      type: COGNITO_USER_POOLS
      restApi: ${thisRestApi.id}
      providerArns: ${thisUserPools.arns}
  any:
    type: aws:apigateway:Method
    properties:
      restApi: ${thisRestApi.id}
      resourceId: ${thisResource.id}
      httpMethod: ANY
      authorization: COGNITO_USER_POOLS
      authorizerId: ${thisAuthorizer.id}
      requestParameters:
        method.request.path.proxy: true
variables:
  thisUserPools:
    fn::invoke:
      Function: aws:cognito:getUserPools
      Arguments:
        name: ${cognitoUserPoolName}
```


## Import

Using `pulumi import`, import `aws_api_gateway_method` using `REST-API-ID/RESOURCE-ID/HTTP-METHOD`. For example:

```sh
 $ pulumi import aws:apigateway/method:Method example 12345abcde/67890fghij/GET
```
 