Manages an Amazon API Gateway Version 2 [model](https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html#models-mappings-models).

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apigatewayv2.Model("example", {
    apiId: aws_apigatewayv2_api.example.id,
    contentType: "application/json",
    schema: JSON.stringify({
        $schema: "http://json-schema.org/draft-04/schema#",
        title: "ExampleModel",
        type: "object",
        properties: {
            id: {
                type: "string",
            },
        },
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.apigatewayv2.Model("example",
    api_id=aws_apigatewayv2_api["example"]["id"],
    content_type="application/json",
    schema=json.dumps({
        "$schema": "http://json-schema.org/draft-04/schema#",
        "title": "ExampleModel",
        "type": "object",
        "properties": {
            "id": {
                "type": "string",
            },
        },
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ApiGatewayV2.Model("example", new()
    {
        ApiId = aws_apigatewayv2_api.Example.Id,
        ContentType = "application/json",
        Schema = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["$schema"] = "http://json-schema.org/draft-04/schema#",
            ["title"] = "ExampleModel",
            ["type"] = "object",
            ["properties"] = new Dictionary<string, object?>
            {
                ["id"] = new Dictionary<string, object?>
                {
                    ["type"] = "string",
                },
            },
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigatewayv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"$schema": "http://json-schema.org/draft-04/schema#",
			"title":   "ExampleModel",
			"type":    "object",
			"properties": map[string]interface{}{
				"id": map[string]interface{}{
					"type": "string",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = apigatewayv2.NewModel(ctx, "example", &apigatewayv2.ModelArgs{
			ApiId:       pulumi.Any(aws_apigatewayv2_api.Example.Id),
			ContentType: pulumi.String("application/json"),
			Schema:      pulumi.String(json0),
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
import com.pulumi.aws.apigatewayv2.Model;
import com.pulumi.aws.apigatewayv2.ModelArgs;
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
        var example = new Model("example", ModelArgs.builder()        
            .apiId(aws_apigatewayv2_api.example().id())
            .contentType("application/json")
            .schema(serializeJson(
                jsonObject(
                    jsonProperty("$schema", "http://json-schema.org/draft-04/schema#"),
                    jsonProperty("title", "ExampleModel"),
                    jsonProperty("type", "object"),
                    jsonProperty("properties", jsonObject(
                        jsonProperty("id", jsonObject(
                            jsonProperty("type", "string")
                        ))
                    ))
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apigatewayv2:Model
    properties:
      apiId: ${aws_apigatewayv2_api.example.id}
      contentType: application/json
      schema:
        fn::toJSON:
          $schema: http://json-schema.org/draft-04/schema#
          title: ExampleModel
          type: object
          properties:
            id:
              type: string
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_apigatewayv2_model` using the API identifier and model identifier. For example:

```sh
 $ pulumi import aws:apigatewayv2/model:Model example aabbccddee/1122334
```
 