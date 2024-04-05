Resource for managing an AWS OpenSearch Serverless Lifecycle Policy. See AWS documentation for [lifecycle policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-lifecycle.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.opensearch.ServerlessLifecyclePolicy("example", {
    type: "retention",
    policy: JSON.stringify({
        Rules: [
            {
                ResourceType: "index",
                Resource: ["index/autoparts-inventory/*"],
                MinIndexRetention: "81d",
            },
            {
                ResourceType: "index",
                Resource: ["index/sales/orders*"],
                NoMinIndexRetention: true,
            },
        ],
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.opensearch.ServerlessLifecyclePolicy("example",
    type="retention",
    policy=json.dumps({
        "Rules": [
            {
                "ResourceType": "index",
                "Resource": ["index/autoparts-inventory/*"],
                "MinIndexRetention": "81d",
            },
            {
                "ResourceType": "index",
                "Resource": ["index/sales/orders*"],
                "NoMinIndexRetention": True,
            },
        ],
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
    var example = new Aws.OpenSearch.ServerlessLifecyclePolicy("example", new()
    {
        Type = "retention",
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Rules"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["ResourceType"] = "index",
                    ["Resource"] = new[]
                    {
                        "index/autoparts-inventory/*",
                    },
                    ["MinIndexRetention"] = "81d",
                },
                new Dictionary<string, object?>
                {
                    ["ResourceType"] = "index",
                    ["Resource"] = new[]
                    {
                        "index/sales/orders*",
                    },
                    ["NoMinIndexRetention"] = true,
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Rules": []interface{}{
				map[string]interface{}{
					"ResourceType": "index",
					"Resource": []string{
						"index/autoparts-inventory/*",
					},
					"MinIndexRetention": "81d",
				},
				map[string]interface{}{
					"ResourceType": "index",
					"Resource": []string{
						"index/sales/orders*",
					},
					"NoMinIndexRetention": true,
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = opensearch.NewServerlessLifecyclePolicy(ctx, "example", &opensearch.ServerlessLifecyclePolicyArgs{
			Type:   pulumi.String("retention"),
			Policy: pulumi.String(json0),
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
import com.pulumi.aws.opensearch.ServerlessLifecyclePolicy;
import com.pulumi.aws.opensearch.ServerlessLifecyclePolicyArgs;
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
        var example = new ServerlessLifecyclePolicy("example", ServerlessLifecyclePolicyArgs.builder()        
            .type("retention")
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Rules", jsonArray(
                        jsonObject(
                            jsonProperty("ResourceType", "index"),
                            jsonProperty("Resource", jsonArray("index/autoparts-inventory/*")),
                            jsonProperty("MinIndexRetention", "81d")
                        ), 
                        jsonObject(
                            jsonProperty("ResourceType", "index"),
                            jsonProperty("Resource", jsonArray("index/sales/orders*")),
                            jsonProperty("NoMinIndexRetention", true)
                        )
                    ))
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:opensearch:ServerlessLifecyclePolicy
    properties:
      type: retention
      policy:
        fn::toJSON:
          Rules:
            - ResourceType: index
              Resource:
                - index/autoparts-inventory/*
              MinIndexRetention: 81d
            - ResourceType: index
              Resource:
                - index/sales/orders*
              NoMinIndexRetention: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import OpenSearch Serverless Lifecycle Policy using the `name` and `type` arguments separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:opensearch/serverlessLifecyclePolicy:ServerlessLifecyclePolicy example example/retention
```
 