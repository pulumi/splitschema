Provides a Resource Group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.resourcegroups.Group("test", {resourceQuery: {
    query: `{
  "ResourceTypeFilters": [
    "AWS::EC2::Instance"
  ],
  "TagFilters": [
    {
      "Key": "Stage",
      "Values": ["Test"]
    }
  ]
}

`,
}});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.resourcegroups.Group("test", resource_query=aws.resourcegroups.GroupResourceQueryArgs(
    query="""{
  "ResourceTypeFilters": [
    "AWS::EC2::Instance"
  ],
  "TagFilters": [
    {
      "Key": "Stage",
      "Values": ["Test"]
    }
  ]
}

""",
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.ResourceGroups.Group("test", new()
    {
        ResourceQuery = new Aws.ResourceGroups.Inputs.GroupResourceQueryArgs
        {
            Query = @"{
  ""ResourceTypeFilters"": [
    ""AWS::EC2::Instance""
  ],
  ""TagFilters"": [
    {
      ""Key"": ""Stage"",
      ""Values"": [""Test""]
    }
  ]
}

",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/resourcegroups"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resourcegroups.NewGroup(ctx, "test", &resourcegroups.GroupArgs{
			ResourceQuery: &resourcegroups.GroupResourceQueryArgs{
				Query: pulumi.String(`{
  "ResourceTypeFilters": [
    "AWS::EC2::Instance"
  ],
  "TagFilters": [
    {
      "Key": "Stage",
      "Values": ["Test"]
    }
  ]
}

`),
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
import com.pulumi.aws.resourcegroups.Group;
import com.pulumi.aws.resourcegroups.GroupArgs;
import com.pulumi.aws.resourcegroups.inputs.GroupResourceQueryArgs;
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
        var test = new Group("test", GroupArgs.builder()        
            .resourceQuery(GroupResourceQueryArgs.builder()
                .query("""
{
  "ResourceTypeFilters": [
    "AWS::EC2::Instance"
  ],
  "TagFilters": [
    {
      "Key": "Stage",
      "Values": ["Test"]
    }
  ]
}

                """)
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:resourcegroups:Group
    properties:
      resourceQuery:
        query: |+
          {
            "ResourceTypeFilters": [
              "AWS::EC2::Instance"
            ],
            "TagFilters": [
              {
                "Key": "Stage",
                "Values": ["Test"]
              }
            ]
          }
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import resource groups using the `name`. For example:

```sh
 $ pulumi import aws:resourcegroups/group:Group foo resource-group-name
```
 