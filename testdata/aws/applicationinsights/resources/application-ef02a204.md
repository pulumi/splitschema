Provides a ApplicationInsights Application resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGroup = new aws.resourcegroups.Group("exampleGroup", {resourceQuery: {
    query: JSON.stringify({
        ResourceTypeFilters: ["AWS::EC2::Instance"],
        TagFilters: [{
            Key: "Stage",
            Values: ["Test"],
        }],
    }),
}});
const exampleApplication = new aws.applicationinsights.Application("exampleApplication", {resourceGroupName: exampleGroup.name});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_group = aws.resourcegroups.Group("exampleGroup", resource_query=aws.resourcegroups.GroupResourceQueryArgs(
    query=json.dumps({
        "ResourceTypeFilters": ["AWS::EC2::Instance"],
        "TagFilters": [{
            "Key": "Stage",
            "Values": ["Test"],
        }],
    }),
))
example_application = aws.applicationinsights.Application("exampleApplication", resource_group_name=example_group.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleGroup = new Aws.ResourceGroups.Group("exampleGroup", new()
    {
        ResourceQuery = new Aws.ResourceGroups.Inputs.GroupResourceQueryArgs
        {
            Query = JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                ["ResourceTypeFilters"] = new[]
                {
                    "AWS::EC2::Instance",
                },
                ["TagFilters"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["Key"] = "Stage",
                        ["Values"] = new[]
                        {
                            "Test",
                        },
                    },
                },
            }),
        },
    });

    var exampleApplication = new Aws.ApplicationInsights.Application("exampleApplication", new()
    {
        ResourceGroupName = exampleGroup.Name,
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/applicationinsights"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/resourcegroups"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"ResourceTypeFilters": []string{
				"AWS::EC2::Instance",
			},
			"TagFilters": []map[string]interface{}{
				map[string]interface{}{
					"Key": "Stage",
					"Values": []string{
						"Test",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleGroup, err := resourcegroups.NewGroup(ctx, "exampleGroup", &resourcegroups.GroupArgs{
			ResourceQuery: &resourcegroups.GroupResourceQueryArgs{
				Query: pulumi.String(json0),
			},
		})
		if err != nil {
			return err
		}
		_, err = applicationinsights.NewApplication(ctx, "exampleApplication", &applicationinsights.ApplicationArgs{
			ResourceGroupName: exampleGroup.Name,
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
import com.pulumi.aws.applicationinsights.Application;
import com.pulumi.aws.applicationinsights.ApplicationArgs;
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
        var exampleGroup = new Group("exampleGroup", GroupArgs.builder()        
            .resourceQuery(GroupResourceQueryArgs.builder()
                .query(serializeJson(
                    jsonObject(
                        jsonProperty("ResourceTypeFilters", jsonArray("AWS::EC2::Instance")),
                        jsonProperty("TagFilters", jsonArray(jsonObject(
                            jsonProperty("Key", "Stage"),
                            jsonProperty("Values", jsonArray("Test"))
                        )))
                    )))
                .build())
            .build());

        var exampleApplication = new Application("exampleApplication", ApplicationArgs.builder()        
            .resourceGroupName(exampleGroup.name())
            .build());

    }
}
```
```yaml
resources:
  exampleApplication:
    type: aws:applicationinsights:Application
    properties:
      resourceGroupName: ${exampleGroup.name}
  exampleGroup:
    type: aws:resourcegroups:Group
    properties:
      resourceQuery:
        query:
          fn::toJSON:
            ResourceTypeFilters:
              - AWS::EC2::Instance
            TagFilters:
              - Key: Stage
                Values:
                  - Test
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ApplicationInsights Applications using the `resource_group_name`. For example:

```sh
 $ pulumi import aws:applicationinsights/application:Application some some-application
```
 