Resource for managing an AWS OpenSearch Serverless Collection.

> **NOTE:** An `aws.opensearch.ServerlessCollection` cannot be created without having an applicable encryption security policy. Use the `depends_on` meta-argument to define this dependency.

> **NOTE:** An `aws.opensearch.ServerlessCollection` is not accessible without configuring an applicable network security policy. Data cannot be accessed without configuring an applicable data access policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleServerlessSecurityPolicy = new aws.opensearch.ServerlessSecurityPolicy("exampleServerlessSecurityPolicy", {
    type: "encryption",
    policy: JSON.stringify({
        Rules: [{
            Resource: ["collection/example"],
            ResourceType: "collection",
        }],
        AWSOwnedKey: true,
    }),
});
const exampleServerlessCollection = new aws.opensearch.ServerlessCollection("exampleServerlessCollection", {}, {
    dependsOn: [exampleServerlessSecurityPolicy],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_serverless_security_policy = aws.opensearch.ServerlessSecurityPolicy("exampleServerlessSecurityPolicy",
    type="encryption",
    policy=json.dumps({
        "Rules": [{
            "Resource": ["collection/example"],
            "ResourceType": "collection",
        }],
        "AWSOwnedKey": True,
    }))
example_serverless_collection = aws.opensearch.ServerlessCollection("exampleServerlessCollection", opts=pulumi.ResourceOptions(depends_on=[example_serverless_security_policy]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleServerlessSecurityPolicy = new Aws.OpenSearch.ServerlessSecurityPolicy("exampleServerlessSecurityPolicy", new()
    {
        Type = "encryption",
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Rules"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Resource"] = new[]
                    {
                        "collection/example",
                    },
                    ["ResourceType"] = "collection",
                },
            },
            ["AWSOwnedKey"] = true,
        }),
    });

    var exampleServerlessCollection = new Aws.OpenSearch.ServerlessCollection("exampleServerlessCollection", new()
    {
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleServerlessSecurityPolicy,
        },
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
			"Rules": []map[string]interface{}{
				map[string]interface{}{
					"Resource": []string{
						"collection/example",
					},
					"ResourceType": "collection",
				},
			},
			"AWSOwnedKey": true,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleServerlessSecurityPolicy, err := opensearch.NewServerlessSecurityPolicy(ctx, "exampleServerlessSecurityPolicy", &opensearch.ServerlessSecurityPolicyArgs{
			Type:   pulumi.String("encryption"),
			Policy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = opensearch.NewServerlessCollection(ctx, "exampleServerlessCollection", nil, pulumi.DependsOn([]pulumi.Resource{
			exampleServerlessSecurityPolicy,
		}))
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
import com.pulumi.aws.opensearch.ServerlessSecurityPolicy;
import com.pulumi.aws.opensearch.ServerlessSecurityPolicyArgs;
import com.pulumi.aws.opensearch.ServerlessCollection;
import com.pulumi.aws.opensearch.ServerlessCollectionArgs;
import static com.pulumi.codegen.internal.Serialization.*;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleServerlessSecurityPolicy = new ServerlessSecurityPolicy("exampleServerlessSecurityPolicy", ServerlessSecurityPolicyArgs.builder()        
            .type("encryption")
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Rules", jsonArray(jsonObject(
                        jsonProperty("Resource", jsonArray("collection/example")),
                        jsonProperty("ResourceType", "collection")
                    ))),
                    jsonProperty("AWSOwnedKey", true)
                )))
            .build());

        var exampleServerlessCollection = new ServerlessCollection("exampleServerlessCollection", ServerlessCollectionArgs.Empty, CustomResourceOptions.builder()
            .dependsOn(exampleServerlessSecurityPolicy)
            .build());

    }
}
```
```yaml
resources:
  exampleServerlessSecurityPolicy:
    type: aws:opensearch:ServerlessSecurityPolicy
    properties:
      type: encryption
      policy:
        fn::toJSON:
          Rules:
            - Resource:
                - collection/example
              ResourceType: collection
          AWSOwnedKey: true
  exampleServerlessCollection:
    type: aws:opensearch:ServerlessCollection
    options:
      dependson:
        - ${exampleServerlessSecurityPolicy}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import OpenSearchServerless Collection using the `id`. For example:

```sh
 $ pulumi import aws:opensearch/serverlessCollection:ServerlessCollection example example
```
 