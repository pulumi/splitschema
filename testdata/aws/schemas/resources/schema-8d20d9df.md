Provides an EventBridge Schema resource.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testRegistry = new aws.schemas.Registry("testRegistry", {});
const testSchema = new aws.schemas.Schema("testSchema", {
    registryName: testRegistry.name,
    type: "OpenApi3",
    description: "The schema definition for my event",
    content: JSON.stringify({
        openapi: "3.0.0",
        info: {
            version: "1.0.0",
            title: "Event",
        },
        paths: {},
        components: {
            schemas: {
                Event: {
                    type: "object",
                    properties: {
                        name: {
                            type: "string",
                        },
                    },
                },
            },
        },
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

test_registry = aws.schemas.Registry("testRegistry")
test_schema = aws.schemas.Schema("testSchema",
    registry_name=test_registry.name,
    type="OpenApi3",
    description="The schema definition for my event",
    content=json.dumps({
        "openapi": "3.0.0",
        "info": {
            "version": "1.0.0",
            "title": "Event",
        },
        "paths": {},
        "components": {
            "schemas": {
                "Event": {
                    "type": "object",
                    "properties": {
                        "name": {
                            "type": "string",
                        },
                    },
                },
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
    var testRegistry = new Aws.Schemas.Registry("testRegistry");

    var testSchema = new Aws.Schemas.Schema("testSchema", new()
    {
        RegistryName = testRegistry.Name,
        Type = "OpenApi3",
        Description = "The schema definition for my event",
        Content = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["openapi"] = "3.0.0",
            ["info"] = new Dictionary<string, object?>
            {
                ["version"] = "1.0.0",
                ["title"] = "Event",
            },
            ["paths"] = new Dictionary<string, object?>
            {
            },
            ["components"] = new Dictionary<string, object?>
            {
                ["schemas"] = new Dictionary<string, object?>
                {
                    ["Event"] = new Dictionary<string, object?>
                    {
                        ["type"] = "object",
                        ["properties"] = new Dictionary<string, object?>
                        {
                            ["name"] = new Dictionary<string, object?>
                            {
                                ["type"] = "string",
                            },
                        },
                    },
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/schemas"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testRegistry, err := schemas.NewRegistry(ctx, "testRegistry", nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"openapi": "3.0.0",
			"info": map[string]interface{}{
				"version": "1.0.0",
				"title":   "Event",
			},
			"paths": nil,
			"components": map[string]interface{}{
				"schemas": map[string]interface{}{
					"Event": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"name": map[string]interface{}{
								"type": "string",
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = schemas.NewSchema(ctx, "testSchema", &schemas.SchemaArgs{
			RegistryName: testRegistry.Name,
			Type:         pulumi.String("OpenApi3"),
			Description:  pulumi.String("The schema definition for my event"),
			Content:      pulumi.String(json0),
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
import com.pulumi.aws.schemas.Registry;
import com.pulumi.aws.schemas.Schema;
import com.pulumi.aws.schemas.SchemaArgs;
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
        var testRegistry = new Registry("testRegistry");

        var testSchema = new Schema("testSchema", SchemaArgs.builder()        
            .registryName(testRegistry.name())
            .type("OpenApi3")
            .description("The schema definition for my event")
            .content(serializeJson(
                jsonObject(
                    jsonProperty("openapi", "3.0.0"),
                    jsonProperty("info", jsonObject(
                        jsonProperty("version", "1.0.0"),
                        jsonProperty("title", "Event")
                    )),
                    jsonProperty("paths", jsonObject(

                    )),
                    jsonProperty("components", jsonObject(
                        jsonProperty("schemas", jsonObject(
                            jsonProperty("Event", jsonObject(
                                jsonProperty("type", "object"),
                                jsonProperty("properties", jsonObject(
                                    jsonProperty("name", jsonObject(
                                        jsonProperty("type", "string")
                                    ))
                                ))
                            ))
                        ))
                    ))
                )))
            .build());

    }
}
```
```yaml
resources:
  testRegistry:
    type: aws:schemas:Registry
  testSchema:
    type: aws:schemas:Schema
    properties:
      registryName: ${testRegistry.name}
      type: OpenApi3
      description: The schema definition for my event
      content:
        fn::toJSON:
          openapi: 3.0.0
          info:
            version: 1.0.0
            title: Event
          paths: {}
          components:
            schemas:
              Event:
                type: object
                properties:
                  name:
                    type: string
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EventBridge schema using the `name` and `registry_name`. For example:

```sh
 $ pulumi import aws:schemas/schema:Schema test name/registry
```
 