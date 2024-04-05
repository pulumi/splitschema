Manages an AWS IoT Thing Group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const parent = new aws.iot.ThingGroup("parent", {});
const example = new aws.iot.ThingGroup("example", {
    parentGroupName: parent.name,
    properties: {
        attributePayload: {
            attributes: {
                One: "11111",
                Two: "TwoTwo",
            },
        },
        description: "This is my thing group",
    },
    tags: {
        managed: "true",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

parent = aws.iot.ThingGroup("parent")
example = aws.iot.ThingGroup("example",
    parent_group_name=parent.name,
    properties=aws.iot.ThingGroupPropertiesArgs(
        attribute_payload=aws.iot.ThingGroupPropertiesAttributePayloadArgs(
            attributes={
                "One": "11111",
                "Two": "TwoTwo",
            },
        ),
        description="This is my thing group",
    ),
    tags={
        "managed": "true",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var parent = new Aws.Iot.ThingGroup("parent");

    var example = new Aws.Iot.ThingGroup("example", new()
    {
        ParentGroupName = parent.Name,
        Properties = new Aws.Iot.Inputs.ThingGroupPropertiesArgs
        {
            AttributePayload = new Aws.Iot.Inputs.ThingGroupPropertiesAttributePayloadArgs
            {
                Attributes = 
                {
                    { "One", "11111" },
                    { "Two", "TwoTwo" },
                },
            },
            Description = "This is my thing group",
        },
        Tags = 
        {
            { "managed", "true" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		parent, err := iot.NewThingGroup(ctx, "parent", nil)
		if err != nil {
			return err
		}
		_, err = iot.NewThingGroup(ctx, "example", &iot.ThingGroupArgs{
			ParentGroupName: parent.Name,
			Properties: &iot.ThingGroupPropertiesArgs{
				AttributePayload: &iot.ThingGroupPropertiesAttributePayloadArgs{
					Attributes: pulumi.StringMap{
						"One": pulumi.String("11111"),
						"Two": pulumi.String("TwoTwo"),
					},
				},
				Description: pulumi.String("This is my thing group"),
			},
			Tags: pulumi.StringMap{
				"managed": pulumi.String("true"),
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
import com.pulumi.aws.iot.ThingGroup;
import com.pulumi.aws.iot.ThingGroupArgs;
import com.pulumi.aws.iot.inputs.ThingGroupPropertiesArgs;
import com.pulumi.aws.iot.inputs.ThingGroupPropertiesAttributePayloadArgs;
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
        var parent = new ThingGroup("parent");

        var example = new ThingGroup("example", ThingGroupArgs.builder()        
            .parentGroupName(parent.name())
            .properties(ThingGroupPropertiesArgs.builder()
                .attributePayload(ThingGroupPropertiesAttributePayloadArgs.builder()
                    .attributes(Map.ofEntries(
                        Map.entry("One", "11111"),
                        Map.entry("Two", "TwoTwo")
                    ))
                    .build())
                .description("This is my thing group")
                .build())
            .tags(Map.of("managed", "true"))
            .build());

    }
}
```
```yaml
resources:
  parent:
    type: aws:iot:ThingGroup
  example:
    type: aws:iot:ThingGroup
    properties:
      parentGroupName: ${parent.name}
      properties:
        attributePayload:
          attributes:
            One: '11111'
            Two: TwoTwo
        description: This is my thing group
      tags:
        managed: 'true'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IoT Things Groups using the name. For example:

```sh
 $ pulumi import aws:iot/thingGroup:ThingGroup example example
```
 