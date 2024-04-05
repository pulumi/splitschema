Provides an API Gateway Usage Plan.

{{% examples %}}
## Example Usage
{{% example %}}

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
const development = new aws.apigateway.Stage("development", {
    deployment: exampleDeployment.id,
    restApi: exampleRestApi.id,
    stageName: "development",
});
const production = new aws.apigateway.Stage("production", {
    deployment: exampleDeployment.id,
    restApi: exampleRestApi.id,
    stageName: "production",
});
const exampleUsagePlan = new aws.apigateway.UsagePlan("exampleUsagePlan", {
    description: "my description",
    productCode: "MYCODE",
    apiStages: [
        {
            apiId: exampleRestApi.id,
            stage: development.stageName,
        },
        {
            apiId: exampleRestApi.id,
            stage: production.stageName,
        },
    ],
    quotaSettings: {
        limit: 20,
        offset: 2,
        period: "WEEK",
    },
    throttleSettings: {
        burstLimit: 5,
        rateLimit: 10,
    },
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
development = aws.apigateway.Stage("development",
    deployment=example_deployment.id,
    rest_api=example_rest_api.id,
    stage_name="development")
production = aws.apigateway.Stage("production",
    deployment=example_deployment.id,
    rest_api=example_rest_api.id,
    stage_name="production")
example_usage_plan = aws.apigateway.UsagePlan("exampleUsagePlan",
    description="my description",
    product_code="MYCODE",
    api_stages=[
        aws.apigateway.UsagePlanApiStageArgs(
            api_id=example_rest_api.id,
            stage=development.stage_name,
        ),
        aws.apigateway.UsagePlanApiStageArgs(
            api_id=example_rest_api.id,
            stage=production.stage_name,
        ),
    ],
    quota_settings=aws.apigateway.UsagePlanQuotaSettingsArgs(
        limit=20,
        offset=2,
        period="WEEK",
    ),
    throttle_settings=aws.apigateway.UsagePlanThrottleSettingsArgs(
        burst_limit=5,
        rate_limit=10,
    ))
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

    var development = new Aws.ApiGateway.Stage("development", new()
    {
        Deployment = exampleDeployment.Id,
        RestApi = exampleRestApi.Id,
        StageName = "development",
    });

    var production = new Aws.ApiGateway.Stage("production", new()
    {
        Deployment = exampleDeployment.Id,
        RestApi = exampleRestApi.Id,
        StageName = "production",
    });

    var exampleUsagePlan = new Aws.ApiGateway.UsagePlan("exampleUsagePlan", new()
    {
        Description = "my description",
        ProductCode = "MYCODE",
        ApiStages = new[]
        {
            new Aws.ApiGateway.Inputs.UsagePlanApiStageArgs
            {
                ApiId = exampleRestApi.Id,
                Stage = development.StageName,
            },
            new Aws.ApiGateway.Inputs.UsagePlanApiStageArgs
            {
                ApiId = exampleRestApi.Id,
                Stage = production.StageName,
            },
        },
        QuotaSettings = new Aws.ApiGateway.Inputs.UsagePlanQuotaSettingsArgs
        {
            Limit = 20,
            Offset = 2,
            Period = "WEEK",
        },
        ThrottleSettings = new Aws.ApiGateway.Inputs.UsagePlanThrottleSettingsArgs
        {
            BurstLimit = 5,
            RateLimit = 10,
        },
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
		development, err := apigateway.NewStage(ctx, "development", &apigateway.StageArgs{
			Deployment: exampleDeployment.ID(),
			RestApi:    exampleRestApi.ID(),
			StageName:  pulumi.String("development"),
		})
		if err != nil {
			return err
		}
		production, err := apigateway.NewStage(ctx, "production", &apigateway.StageArgs{
			Deployment: exampleDeployment.ID(),
			RestApi:    exampleRestApi.ID(),
			StageName:  pulumi.String("production"),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewUsagePlan(ctx, "exampleUsagePlan", &apigateway.UsagePlanArgs{
			Description: pulumi.String("my description"),
			ProductCode: pulumi.String("MYCODE"),
			ApiStages: apigateway.UsagePlanApiStageArray{
				&apigateway.UsagePlanApiStageArgs{
					ApiId: exampleRestApi.ID(),
					Stage: development.StageName,
				},
				&apigateway.UsagePlanApiStageArgs{
					ApiId: exampleRestApi.ID(),
					Stage: production.StageName,
				},
			},
			QuotaSettings: &apigateway.UsagePlanQuotaSettingsArgs{
				Limit:  pulumi.Int(20),
				Offset: pulumi.Int(2),
				Period: pulumi.String("WEEK"),
			},
			ThrottleSettings: &apigateway.UsagePlanThrottleSettingsArgs{
				BurstLimit: pulumi.Int(5),
				RateLimit:  pulumi.Float64(10),
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
import com.pulumi.aws.apigateway.RestApiArgs;
import com.pulumi.aws.apigateway.Deployment;
import com.pulumi.aws.apigateway.DeploymentArgs;
import com.pulumi.aws.apigateway.Stage;
import com.pulumi.aws.apigateway.StageArgs;
import com.pulumi.aws.apigateway.UsagePlan;
import com.pulumi.aws.apigateway.UsagePlanArgs;
import com.pulumi.aws.apigateway.inputs.UsagePlanApiStageArgs;
import com.pulumi.aws.apigateway.inputs.UsagePlanQuotaSettingsArgs;
import com.pulumi.aws.apigateway.inputs.UsagePlanThrottleSettingsArgs;
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

        var development = new Stage("development", StageArgs.builder()        
            .deployment(exampleDeployment.id())
            .restApi(exampleRestApi.id())
            .stageName("development")
            .build());

        var production = new Stage("production", StageArgs.builder()        
            .deployment(exampleDeployment.id())
            .restApi(exampleRestApi.id())
            .stageName("production")
            .build());

        var exampleUsagePlan = new UsagePlan("exampleUsagePlan", UsagePlanArgs.builder()        
            .description("my description")
            .productCode("MYCODE")
            .apiStages(            
                UsagePlanApiStageArgs.builder()
                    .apiId(exampleRestApi.id())
                    .stage(development.stageName())
                    .build(),
                UsagePlanApiStageArgs.builder()
                    .apiId(exampleRestApi.id())
                    .stage(production.stageName())
                    .build())
            .quotaSettings(UsagePlanQuotaSettingsArgs.builder()
                .limit(20)
                .offset(2)
                .period("WEEK")
                .build())
            .throttleSettings(UsagePlanThrottleSettingsArgs.builder()
                .burstLimit(5)
                .rateLimit(10)
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS API Gateway Usage Plan using the `id`. For example:

```sh
 $ pulumi import aws:apigateway/usagePlan:UsagePlan myusageplan <usage_plan_id>
```
 