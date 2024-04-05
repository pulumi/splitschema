Resource for managing an AWS OpenSearch Serverless Access Policy. See AWS documentation for [data access policies](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html) and [supported data access policy permissions](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html#serverless-data-supported-permissions).

{{% examples %}}
## Example Usage
{{% example %}}
### Grant all collection and index permissions

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const example = new aws.opensearch.ServerlessAccessPolicy("example", {
    type: "data",
    description: "read and write permissions",
    policy: current.then(current => JSON.stringify([{
        Rules: [
            {
                ResourceType: "index",
                Resource: ["index/example-collection/*"],
                Permission: ["aoss:*"],
            },
            {
                ResourceType: "collection",
                Resource: ["collection/example-collection"],
                Permission: ["aoss:*"],
            },
        ],
        Principal: [current.arn],
    }])),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

current = aws.get_caller_identity()
example = aws.opensearch.ServerlessAccessPolicy("example",
    type="data",
    description="read and write permissions",
    policy=json.dumps([{
        "Rules": [
            {
                "ResourceType": "index",
                "Resource": ["index/example-collection/*"],
                "Permission": ["aoss:*"],
            },
            {
                "ResourceType": "collection",
                "Resource": ["collection/example-collection"],
                "Permission": ["aoss:*"],
            },
        ],
        "Principal": [current.arn],
    }]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var example = new Aws.OpenSearch.ServerlessAccessPolicy("example", new()
    {
        Type = "data",
        Description = "read and write permissions",
        Policy = JsonSerializer.Serialize(new[]
        {
            new Dictionary<string, object?>
            {
                ["Rules"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["ResourceType"] = "index",
                        ["Resource"] = new[]
                        {
                            "index/example-collection/*",
                        },
                        ["Permission"] = new[]
                        {
                            "aoss:*",
                        },
                    },
                    new Dictionary<string, object?>
                    {
                        ["ResourceType"] = "collection",
                        ["Resource"] = new[]
                        {
                            "collection/example-collection",
                        },
                        ["Permission"] = new[]
                        {
                            "aoss:*",
                        },
                    },
                },
                ["Principal"] = new[]
                {
                    current.Apply(getCallerIdentityResult => getCallerIdentityResult.Arn),
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal([]map[string]interface{}{
			map[string]interface{}{
				"Rules": []map[string]interface{}{
					map[string]interface{}{
						"ResourceType": "index",
						"Resource": []string{
							"index/example-collection/*",
						},
						"Permission": []string{
							"aoss:*",
						},
					},
					map[string]interface{}{
						"ResourceType": "collection",
						"Resource": []string{
							"collection/example-collection",
						},
						"Permission": []string{
							"aoss:*",
						},
					},
				},
				"Principal": []*string{
					current.Arn,
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = opensearch.NewServerlessAccessPolicy(ctx, "example", &opensearch.ServerlessAccessPolicyArgs{
			Type:        pulumi.String("data"),
			Description: pulumi.String("read and write permissions"),
			Policy:      pulumi.String(json0),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.opensearch.ServerlessAccessPolicy;
import com.pulumi.aws.opensearch.ServerlessAccessPolicyArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        var example = new ServerlessAccessPolicy("example", ServerlessAccessPolicyArgs.builder()        
            .type("data")
            .description("read and write permissions")
            .policy(serializeJson(
                jsonArray(jsonObject(
                    jsonProperty("Rules", jsonArray(
                        jsonObject(
                            jsonProperty("ResourceType", "index"),
                            jsonProperty("Resource", jsonArray("index/example-collection/*")),
                            jsonProperty("Permission", jsonArray("aoss:*"))
                        ), 
                        jsonObject(
                            jsonProperty("ResourceType", "collection"),
                            jsonProperty("Resource", jsonArray("collection/example-collection")),
                            jsonProperty("Permission", jsonArray("aoss:*"))
                        )
                    )),
                    jsonProperty("Principal", jsonArray(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.arn())))
                ))))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:opensearch:ServerlessAccessPolicy
    properties:
      type: data
      description: read and write permissions
      policy:
        fn::toJSON:
          - Rules:
              - ResourceType: index
                Resource:
                  - index/example-collection/*
                Permission:
                  - aoss:*
              - ResourceType: collection
                Resource:
                  - collection/example-collection
                Permission:
                  - aoss:*
            Principal:
              - ${current.arn}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### Grant read-only collection and index permissions

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const example = new aws.opensearch.ServerlessAccessPolicy("example", {
    type: "data",
    description: "read-only permissions",
    policy: current.then(current => JSON.stringify([{
        Rules: [
            {
                ResourceType: "index",
                Resource: ["index/example-collection/*"],
                Permission: [
                    "aoss:DescribeIndex",
                    "aoss:ReadDocument",
                ],
            },
            {
                ResourceType: "collection",
                Resource: ["collection/example-collection"],
                Permission: ["aoss:DescribeCollectionItems"],
            },
        ],
        Principal: [current.arn],
    }])),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

current = aws.get_caller_identity()
example = aws.opensearch.ServerlessAccessPolicy("example",
    type="data",
    description="read-only permissions",
    policy=json.dumps([{
        "Rules": [
            {
                "ResourceType": "index",
                "Resource": ["index/example-collection/*"],
                "Permission": [
                    "aoss:DescribeIndex",
                    "aoss:ReadDocument",
                ],
            },
            {
                "ResourceType": "collection",
                "Resource": ["collection/example-collection"],
                "Permission": ["aoss:DescribeCollectionItems"],
            },
        ],
        "Principal": [current.arn],
    }]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var example = new Aws.OpenSearch.ServerlessAccessPolicy("example", new()
    {
        Type = "data",
        Description = "read-only permissions",
        Policy = JsonSerializer.Serialize(new[]
        {
            new Dictionary<string, object?>
            {
                ["Rules"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["ResourceType"] = "index",
                        ["Resource"] = new[]
                        {
                            "index/example-collection/*",
                        },
                        ["Permission"] = new[]
                        {
                            "aoss:DescribeIndex",
                            "aoss:ReadDocument",
                        },
                    },
                    new Dictionary<string, object?>
                    {
                        ["ResourceType"] = "collection",
                        ["Resource"] = new[]
                        {
                            "collection/example-collection",
                        },
                        ["Permission"] = new[]
                        {
                            "aoss:DescribeCollectionItems",
                        },
                    },
                },
                ["Principal"] = new[]
                {
                    current.Apply(getCallerIdentityResult => getCallerIdentityResult.Arn),
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal([]map[string]interface{}{
			map[string]interface{}{
				"Rules": []interface{}{
					map[string]interface{}{
						"ResourceType": "index",
						"Resource": []string{
							"index/example-collection/*",
						},
						"Permission": []string{
							"aoss:DescribeIndex",
							"aoss:ReadDocument",
						},
					},
					map[string]interface{}{
						"ResourceType": "collection",
						"Resource": []string{
							"collection/example-collection",
						},
						"Permission": []string{
							"aoss:DescribeCollectionItems",
						},
					},
				},
				"Principal": []*string{
					current.Arn,
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = opensearch.NewServerlessAccessPolicy(ctx, "example", &opensearch.ServerlessAccessPolicyArgs{
			Type:        pulumi.String("data"),
			Description: pulumi.String("read-only permissions"),
			Policy:      pulumi.String(json0),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.opensearch.ServerlessAccessPolicy;
import com.pulumi.aws.opensearch.ServerlessAccessPolicyArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        var example = new ServerlessAccessPolicy("example", ServerlessAccessPolicyArgs.builder()        
            .type("data")
            .description("read-only permissions")
            .policy(serializeJson(
                jsonArray(jsonObject(
                    jsonProperty("Rules", jsonArray(
                        jsonObject(
                            jsonProperty("ResourceType", "index"),
                            jsonProperty("Resource", jsonArray("index/example-collection/*")),
                            jsonProperty("Permission", jsonArray(
                                "aoss:DescribeIndex", 
                                "aoss:ReadDocument"
                            ))
                        ), 
                        jsonObject(
                            jsonProperty("ResourceType", "collection"),
                            jsonProperty("Resource", jsonArray("collection/example-collection")),
                            jsonProperty("Permission", jsonArray("aoss:DescribeCollectionItems"))
                        )
                    )),
                    jsonProperty("Principal", jsonArray(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.arn())))
                ))))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:opensearch:ServerlessAccessPolicy
    properties:
      type: data
      description: read-only permissions
      policy:
        fn::toJSON:
          - Rules:
              - ResourceType: index
                Resource:
                  - index/example-collection/*
                Permission:
                  - aoss:DescribeIndex
                  - aoss:ReadDocument
              - ResourceType: collection
                Resource:
                  - collection/example-collection
                Permission:
                  - aoss:DescribeCollectionItems
            Principal:
              - ${current.arn}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### Grant SAML identity permissions

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.opensearch.ServerlessAccessPolicy("example", {
    type: "data",
    description: "saml permissions",
    policy: JSON.stringify([{
        Rules: [
            {
                ResourceType: "index",
                Resource: ["index/example-collection/*"],
                Permission: ["aoss:*"],
            },
            {
                ResourceType: "collection",
                Resource: ["collection/example-collection"],
                Permission: ["aoss:*"],
            },
        ],
        Principal: [
            "saml/123456789012/myprovider/user/Annie",
            "saml/123456789012/anotherprovider/group/Accounting",
        ],
    }]),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.opensearch.ServerlessAccessPolicy("example",
    type="data",
    description="saml permissions",
    policy=json.dumps([{
        "Rules": [
            {
                "ResourceType": "index",
                "Resource": ["index/example-collection/*"],
                "Permission": ["aoss:*"],
            },
            {
                "ResourceType": "collection",
                "Resource": ["collection/example-collection"],
                "Permission": ["aoss:*"],
            },
        ],
        "Principal": [
            "saml/123456789012/myprovider/user/Annie",
            "saml/123456789012/anotherprovider/group/Accounting",
        ],
    }]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.OpenSearch.ServerlessAccessPolicy("example", new()
    {
        Type = "data",
        Description = "saml permissions",
        Policy = JsonSerializer.Serialize(new[]
        {
            new Dictionary<string, object?>
            {
                ["Rules"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["ResourceType"] = "index",
                        ["Resource"] = new[]
                        {
                            "index/example-collection/*",
                        },
                        ["Permission"] = new[]
                        {
                            "aoss:*",
                        },
                    },
                    new Dictionary<string, object?>
                    {
                        ["ResourceType"] = "collection",
                        ["Resource"] = new[]
                        {
                            "collection/example-collection",
                        },
                        ["Permission"] = new[]
                        {
                            "aoss:*",
                        },
                    },
                },
                ["Principal"] = new[]
                {
                    "saml/123456789012/myprovider/user/Annie",
                    "saml/123456789012/anotherprovider/group/Accounting",
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
		tmpJSON0, err := json.Marshal([]map[string]interface{}{
			map[string]interface{}{
				"Rules": []map[string]interface{}{
					map[string]interface{}{
						"ResourceType": "index",
						"Resource": []string{
							"index/example-collection/*",
						},
						"Permission": []string{
							"aoss:*",
						},
					},
					map[string]interface{}{
						"ResourceType": "collection",
						"Resource": []string{
							"collection/example-collection",
						},
						"Permission": []string{
							"aoss:*",
						},
					},
				},
				"Principal": []string{
					"saml/123456789012/myprovider/user/Annie",
					"saml/123456789012/anotherprovider/group/Accounting",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = opensearch.NewServerlessAccessPolicy(ctx, "example", &opensearch.ServerlessAccessPolicyArgs{
			Type:        pulumi.String("data"),
			Description: pulumi.String("saml permissions"),
			Policy:      pulumi.String(json0),
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
import com.pulumi.aws.opensearch.ServerlessAccessPolicy;
import com.pulumi.aws.opensearch.ServerlessAccessPolicyArgs;
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
        var example = new ServerlessAccessPolicy("example", ServerlessAccessPolicyArgs.builder()        
            .type("data")
            .description("saml permissions")
            .policy(serializeJson(
                jsonArray(jsonObject(
                    jsonProperty("Rules", jsonArray(
                        jsonObject(
                            jsonProperty("ResourceType", "index"),
                            jsonProperty("Resource", jsonArray("index/example-collection/*")),
                            jsonProperty("Permission", jsonArray("aoss:*"))
                        ), 
                        jsonObject(
                            jsonProperty("ResourceType", "collection"),
                            jsonProperty("Resource", jsonArray("collection/example-collection")),
                            jsonProperty("Permission", jsonArray("aoss:*"))
                        )
                    )),
                    jsonProperty("Principal", jsonArray(
                        "saml/123456789012/myprovider/user/Annie", 
                        "saml/123456789012/anotherprovider/group/Accounting"
                    ))
                ))))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:opensearch:ServerlessAccessPolicy
    properties:
      type: data
      description: saml permissions
      policy:
        fn::toJSON:
          - Rules:
              - ResourceType: index
                Resource:
                  - index/example-collection/*
                Permission:
                  - aoss:*
              - ResourceType: collection
                Resource:
                  - collection/example-collection
                Permission:
                  - aoss:*
            Principal:
              - saml/123456789012/myprovider/user/Annie
              - saml/123456789012/anotherprovider/group/Accounting
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import OpenSearchServerless Access Policy using the `name` and `type` arguments separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:opensearch/serverlessAccessPolicy:ServerlessAccessPolicy example example/data
```
 