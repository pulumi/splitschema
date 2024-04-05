Manages an API Gateway REST Deployment. A deployment is a snapshot of the REST API configuration. The deployment can then be published to callable endpoints via the `aws.apigateway.Stage` resource and optionally managed further with the `aws.apigateway.BasePathMapping` resource, `aws.apigateway.DomainName` resource, and `aws_api_method_settings` resource. For more information, see the [API Gateway Developer Guide](https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-deploy-api.html).

To properly capture all REST API configuration in a deployment, this resource must have dependencies on all prior resources that manage resources/paths, methods, integrations, etc.

* For REST APIs that are configured via OpenAPI specification (`aws.apigateway.RestApi` resource `body` argument), no special dependency setup is needed beyond referencing the  `id` attribute of that resource unless additional resources have further customized the REST API.
* When the REST API configuration involves other resources (`aws.apigateway.Integration` resource), the dependency setup can be done with implicit resource references in the `triggers` argument or explicit resource references using the [resource `dependsOn` custom option](https://www.pulumi.com/docs/intro/concepts/resources/#dependson). The `triggers` argument should be preferred over `depends_on`, since `depends_on` can only capture dependency ordering and will not cause the resource to recreate (redeploy the REST API) with upstream configuration changes.

!> **WARNING:** It is recommended to use the `aws.apigateway.Stage` resource instead of managing an API Gateway Stage via the `stage_name` argument of this resource. When this resource is recreated (REST API redeployment) with the `stage_name` configured, the stage is deleted and recreated. This will cause a temporary service interruption, increase provide plan differences, and can require a second apply to recreate any downstream stage configuration such as associated `aws_api_method_settings` resources.


{{% examples %}}
## Example Usage
{{% example %}}
### OpenAPI Specification

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as crypto from "crypto";

const exampleRestApi = new aws.apigateway.RestApi("exampleRestApi", {body: JSON.stringify({
    openapi: "3.0.1",
    info: {
        title: "example",
        version: "1.0",
    },
    paths: {
        "/path1": {
            get: {
                "x-amazon-apigateway-integration": {
                    httpMethod: "GET",
                    payloadFormatVersion: "1.0",
                    type: "HTTP_PROXY",
                    uri: "https://ip-ranges.amazonaws.com/ip-ranges.json",
                },
            },
        },
    },
})});
const exampleDeployment = new aws.apigateway.Deployment("exampleDeployment", {
    restApi: exampleRestApi.id,
    triggers: {
        redeployment: exampleRestApi.body.apply(body => JSON.stringify(body)).apply(toJSON => crypto.createHash('sha1').update(toJSON).digest('hex')),
    },
});
const exampleStage = new aws.apigateway.Stage("exampleStage", {
    deployment: exampleDeployment.id,
    restApi: exampleRestApi.id,
    stageName: "example",
});
```
```python
import pulumi
import hashlib
import json
import pulumi_aws as aws

example_rest_api = aws.apigateway.RestApi("exampleRestApi", body=json.dumps({
    "openapi": "3.0.1",
    "info": {
        "title": "example",
        "version": "1.0",
    },
    "paths": {
        "/path1": {
            "get": {
                "x-amazon-apigateway-integration": {
                    "httpMethod": "GET",
                    "payloadFormatVersion": "1.0",
                    "type": "HTTP_PROXY",
                    "uri": "https://ip-ranges.amazonaws.com/ip-ranges.json",
                },
            },
        },
    },
}))
example_deployment = aws.apigateway.Deployment("exampleDeployment",
    rest_api=example_rest_api.id,
    triggers={
        "redeployment": example_rest_api.body.apply(lambda body: json.dumps(body)).apply(lambda to_json: hashlib.sha1(to_json.encode()).hexdigest()),
    })
example_stage = aws.apigateway.Stage("exampleStage",
    deployment=example_deployment.id,
    rest_api=example_rest_api.id,
    stage_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

	
string ComputeSHA1(string input) 
{
    var hash = SHA1.Create().ComputeHash(Encoding.UTF8.GetBytes(input));
    return BitConverter.ToString(hash).Replace("-","").ToLowerInvariant();
}

return await Deployment.RunAsync(() => 
{
    var exampleRestApi = new Aws.ApiGateway.RestApi("exampleRestApi", new()
    {
        Body = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["openapi"] = "3.0.1",
            ["info"] = new Dictionary<string, object?>
            {
                ["title"] = "example",
                ["version"] = "1.0",
            },
            ["paths"] = new Dictionary<string, object?>
            {
                ["/path1"] = new Dictionary<string, object?>
                {
                    ["get"] = new Dictionary<string, object?>
                    {
                        ["x-amazon-apigateway-integration"] = new Dictionary<string, object?>
                        {
                            ["httpMethod"] = "GET",
                            ["payloadFormatVersion"] = "1.0",
                            ["type"] = "HTTP_PROXY",
                            ["uri"] = "https://ip-ranges.amazonaws.com/ip-ranges.json",
                        },
                    },
                },
            },
        }),
    });

    var exampleDeployment = new Aws.ApiGateway.Deployment("exampleDeployment", new()
    {
        RestApi = exampleRestApi.Id,
        Triggers = 
        {
            { "redeployment", exampleRestApi.Body.Apply(body => JsonSerializer.Serialize(body)).Apply(toJSON => ComputeSHA1(toJSON)) },
        },
    });

    var exampleStage = new Aws.ApiGateway.Stage("exampleStage", new()
    {
        Deployment = exampleDeployment.Id,
        RestApi = exampleRestApi.Id,
        StageName = "example",
    });

});
```
```go
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func sha1Hash(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"openapi": "3.0.1",
			"info": map[string]interface{}{
				"title":   "example",
				"version": "1.0",
			},
			"paths": map[string]interface{}{
				"/path1": map[string]interface{}{
					"get": map[string]interface{}{
						"x-amazon-apigateway-integration": map[string]interface{}{
							"httpMethod":           "GET",
							"payloadFormatVersion": "1.0",
							"type":                 "HTTP_PROXY",
							"uri":                  "https://ip-ranges.amazonaws.com/ip-ranges.json",
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleRestApi, err := apigateway.NewRestApi(ctx, "exampleRestApi", &apigateway.RestApiArgs{
			Body: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		exampleDeployment, err := apigateway.NewDeployment(ctx, "exampleDeployment", &apigateway.DeploymentArgs{
			RestApi: exampleRestApi.ID(),
			Triggers: pulumi.StringMap{
				"redeployment": exampleRestApi.Body.ApplyT(func(body *string) (pulumi.String, error) {
					var _zero pulumi.String
					tmpJSON1, err := json.Marshal(body)
					if err != nil {
						return _zero, err
					}
					json1 := string(tmpJSON1)
					return pulumi.String(json1), nil
				}).(pulumi.StringOutput).ApplyT(func(toJSON string) (pulumi.String, error) {
					return pulumi.String(sha1Hash(toJSON)), nil
				}).(pulumi.StringOutput),
			},
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewStage(ctx, "exampleStage", &apigateway.StageArgs{
			Deployment: exampleDeployment.ID(),
			RestApi:    exampleRestApi.ID(),
			StageName:  pulumi.String("example"),
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
import com.pulumi.aws.apigateway.Deployment;
import com.pulumi.aws.apigateway.DeploymentArgs;
import com.pulumi.aws.apigateway.Stage;
import com.pulumi.aws.apigateway.StageArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var exampleRestApi = new RestApi("exampleRestApi", RestApiArgs.builder()        
            .body(serializeJson(
                jsonObject(
                    jsonProperty("openapi", "3.0.1"),
                    jsonProperty("info", jsonObject(
                        jsonProperty("title", "example"),
                        jsonProperty("version", "1.0")
                    )),
                    jsonProperty("paths", jsonObject(
                        jsonProperty("/path1", jsonObject(
                            jsonProperty("get", jsonObject(
                                jsonProperty("x-amazon-apigateway-integration", jsonObject(
                                    jsonProperty("httpMethod", "GET"),
                                    jsonProperty("payloadFormatVersion", "1.0"),
                                    jsonProperty("type", "HTTP_PROXY"),
                                    jsonProperty("uri", "https://ip-ranges.amazonaws.com/ip-ranges.json")
                                ))
                            ))
                        ))
                    ))
                )))
            .build());

        var exampleDeployment = new Deployment("exampleDeployment", DeploymentArgs.builder()        
            .restApi(exampleRestApi.id())
            .triggers(Map.of("redeployment", exampleRestApi.body().applyValue(body -> serializeJson(
                body)).applyValue(toJSON -> computeSHA1(toJSON))))
            .build());

        var exampleStage = new Stage("exampleStage", StageArgs.builder()        
            .deployment(exampleDeployment.id())
            .restApi(exampleRestApi.id())
            .stageName("example")
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Resources

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as crypto from "crypto";

const exampleRestApi = new aws.apigateway.RestApi("exampleRestApi", {});
const exampleResource = new aws.apigateway.Resource("exampleResource", {
    parentId: exampleRestApi.rootResourceId,
    pathPart: "example",
    restApi: exampleRestApi.id,
});
const exampleMethod = new aws.apigateway.Method("exampleMethod", {
    authorization: "NONE",
    httpMethod: "GET",
    resourceId: exampleResource.id,
    restApi: exampleRestApi.id,
});
const exampleIntegration = new aws.apigateway.Integration("exampleIntegration", {
    httpMethod: exampleMethod.httpMethod,
    resourceId: exampleResource.id,
    restApi: exampleRestApi.id,
    type: "MOCK",
});
const exampleDeployment = new aws.apigateway.Deployment("exampleDeployment", {
    restApi: exampleRestApi.id,
    triggers: {
        redeployment: pulumi.all([exampleResource.id, exampleMethod.id, exampleIntegration.id]).apply(([exampleResourceId, exampleMethodId, exampleIntegrationId]) => JSON.stringify([
            exampleResourceId,
            exampleMethodId,
            exampleIntegrationId,
        ])).apply(toJSON => crypto.createHash('sha1').update(toJSON).digest('hex')),
    },
});
const exampleStage = new aws.apigateway.Stage("exampleStage", {
    deployment: exampleDeployment.id,
    restApi: exampleRestApi.id,
    stageName: "example",
});
```
```python
import pulumi
import hashlib
import json
import pulumi_aws as aws

example_rest_api = aws.apigateway.RestApi("exampleRestApi")
example_resource = aws.apigateway.Resource("exampleResource",
    parent_id=example_rest_api.root_resource_id,
    path_part="example",
    rest_api=example_rest_api.id)
example_method = aws.apigateway.Method("exampleMethod",
    authorization="NONE",
    http_method="GET",
    resource_id=example_resource.id,
    rest_api=example_rest_api.id)
example_integration = aws.apigateway.Integration("exampleIntegration",
    http_method=example_method.http_method,
    resource_id=example_resource.id,
    rest_api=example_rest_api.id,
    type="MOCK")
example_deployment = aws.apigateway.Deployment("exampleDeployment",
    rest_api=example_rest_api.id,
    triggers={
        "redeployment": pulumi.Output.all(example_resource.id, example_method.id, example_integration.id).apply(lambda exampleResourceId, exampleMethodId, exampleIntegrationId: json.dumps([
            example_resource_id,
            example_method_id,
            example_integration_id,
        ])).apply(lambda to_json: hashlib.sha1(to_json.encode()).hexdigest()),
    })
example_stage = aws.apigateway.Stage("exampleStage",
    deployment=example_deployment.id,
    rest_api=example_rest_api.id,
    stage_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Security.Cryptography;
using System.Text;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

	
string ComputeSHA1(string input) 
{
    var hash = SHA1.Create().ComputeHash(Encoding.UTF8.GetBytes(input));
    return BitConverter.ToString(hash).Replace("-","").ToLowerInvariant();
}

return await Deployment.RunAsync(() => 
{
    var exampleRestApi = new Aws.ApiGateway.RestApi("exampleRestApi");

    var exampleResource = new Aws.ApiGateway.Resource("exampleResource", new()
    {
        ParentId = exampleRestApi.RootResourceId,
        PathPart = "example",
        RestApi = exampleRestApi.Id,
    });

    var exampleMethod = new Aws.ApiGateway.Method("exampleMethod", new()
    {
        Authorization = "NONE",
        HttpMethod = "GET",
        ResourceId = exampleResource.Id,
        RestApi = exampleRestApi.Id,
    });

    var exampleIntegration = new Aws.ApiGateway.Integration("exampleIntegration", new()
    {
        HttpMethod = exampleMethod.HttpMethod,
        ResourceId = exampleResource.Id,
        RestApi = exampleRestApi.Id,
        Type = "MOCK",
    });

    var exampleDeployment = new Aws.ApiGateway.Deployment("exampleDeployment", new()
    {
        RestApi = exampleRestApi.Id,
        Triggers = 
        {
            { "redeployment", Output.Tuple(exampleResource.Id, exampleMethod.Id, exampleIntegration.Id).Apply(values =>
            {
                var exampleResourceId = values.Item1;
                var exampleMethodId = values.Item2;
                var exampleIntegrationId = values.Item3;
                return JsonSerializer.Serialize(new[]
                {
                    exampleResourceId,
                    exampleMethodId,
                    exampleIntegrationId,
                });
            }).Apply(toJSON => ComputeSHA1(toJSON)) },
        },
    });

    var exampleStage = new Aws.ApiGateway.Stage("exampleStage", new()
    {
        Deployment = exampleDeployment.Id,
        RestApi = exampleRestApi.Id,
        StageName = "example",
    });

});
```
```go
package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func sha1Hash(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRestApi, err := apigateway.NewRestApi(ctx, "exampleRestApi", nil)
		if err != nil {
			return err
		}
		exampleResource, err := apigateway.NewResource(ctx, "exampleResource", &apigateway.ResourceArgs{
			ParentId: exampleRestApi.RootResourceId,
			PathPart: pulumi.String("example"),
			RestApi:  exampleRestApi.ID(),
		})
		if err != nil {
			return err
		}
		exampleMethod, err := apigateway.NewMethod(ctx, "exampleMethod", &apigateway.MethodArgs{
			Authorization: pulumi.String("NONE"),
			HttpMethod:    pulumi.String("GET"),
			ResourceId:    exampleResource.ID(),
			RestApi:       exampleRestApi.ID(),
		})
		if err != nil {
			return err
		}
		exampleIntegration, err := apigateway.NewIntegration(ctx, "exampleIntegration", &apigateway.IntegrationArgs{
			HttpMethod: exampleMethod.HttpMethod,
			ResourceId: exampleResource.ID(),
			RestApi:    exampleRestApi.ID(),
			Type:       pulumi.String("MOCK"),
		})
		if err != nil {
			return err
		}
		exampleDeployment, err := apigateway.NewDeployment(ctx, "exampleDeployment", &apigateway.DeploymentArgs{
			RestApi: exampleRestApi.ID(),
			Triggers: pulumi.StringMap{
				"redeployment": pulumi.All(exampleResource.ID(), exampleMethod.ID(), exampleIntegration.ID()).ApplyT(func(_args []interface{}) (string, error) {
					exampleResourceId := _args[0].(string)
					exampleMethodId := _args[1].(string)
					exampleIntegrationId := _args[2].(string)
					var _zero string
					tmpJSON0, err := json.Marshal([]string{
						exampleResourceId,
						exampleMethodId,
						exampleIntegrationId,
					})
					if err != nil {
						return _zero, err
					}
					json0 := string(tmpJSON0)
					return json0, nil
				}).(pulumi.StringOutput).ApplyT(func(toJSON string) (pulumi.String, error) {
					return pulumi.String(sha1Hash(toJSON)), nil
				}).(pulumi.StringOutput),
			},
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewStage(ctx, "exampleStage", &apigateway.StageArgs{
			Deployment: exampleDeployment.ID(),
			RestApi:    exampleRestApi.ID(),
			StageName:  pulumi.String("example"),
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
import com.pulumi.aws.apigateway.Resource;
import com.pulumi.aws.apigateway.ResourceArgs;
import com.pulumi.aws.apigateway.Method;
import com.pulumi.aws.apigateway.MethodArgs;
import com.pulumi.aws.apigateway.Integration;
import com.pulumi.aws.apigateway.IntegrationArgs;
import com.pulumi.aws.apigateway.Deployment;
import com.pulumi.aws.apigateway.DeploymentArgs;
import com.pulumi.aws.apigateway.Stage;
import com.pulumi.aws.apigateway.StageArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var exampleRestApi = new RestApi("exampleRestApi");

        var exampleResource = new Resource("exampleResource", ResourceArgs.builder()        
            .parentId(exampleRestApi.rootResourceId())
            .pathPart("example")
            .restApi(exampleRestApi.id())
            .build());

        var exampleMethod = new Method("exampleMethod", MethodArgs.builder()        
            .authorization("NONE")
            .httpMethod("GET")
            .resourceId(exampleResource.id())
            .restApi(exampleRestApi.id())
            .build());

        var exampleIntegration = new Integration("exampleIntegration", IntegrationArgs.builder()        
            .httpMethod(exampleMethod.httpMethod())
            .resourceId(exampleResource.id())
            .restApi(exampleRestApi.id())
            .type("MOCK")
            .build());

        var exampleDeployment = new Deployment("exampleDeployment", DeploymentArgs.builder()        
            .restApi(exampleRestApi.id())
            .triggers(Map.of("redeployment", Output.tuple(exampleResource.id(), exampleMethod.id(), exampleIntegration.id()).applyValue(values -> {
                var exampleResourceId = values.t1;
                var exampleMethodId = values.t2;
                var exampleIntegrationId = values.t3;
                return serializeJson(
                    jsonArray(
                        exampleResourceId, 
                        exampleMethodId, 
                        exampleIntegrationId
                    ));
            }).applyValue(toJSON -> computeSHA1(toJSON))))
            .build());

        var exampleStage = new Stage("exampleStage", StageArgs.builder()        
            .deployment(exampleDeployment.id())
            .restApi(exampleRestApi.id())
            .stageName("example")
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_api_gateway_deployment` using `REST-API-ID/DEPLOYMENT-ID`. For example:

```sh
 $ pulumi import aws:apigateway/deployment:Deployment example aabbccddee/1122334
```
 The `stage_name`, `stage_description`, and `variables` arguments cannot be imported. Use the `aws_api_gateway_stage` resource to import and manage stages.

The `triggers` argument cannot be imported.

