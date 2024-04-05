Managing [IoT Thing indexing](https://docs.aws.amazon.com/iot/latest/developerguide/managing-index.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iot.IndexingConfiguration("example", {thingIndexingConfiguration: {
    customFields: [
        {
            name: "shadow.desired.power",
            type: "Boolean",
        },
        {
            name: "attributes.version",
            type: "Number",
        },
        {
            name: "shadow.name.thing1shadow.desired.DefaultDesired",
            type: "String",
        },
        {
            name: "deviceDefender.securityProfile1.NUMBER_VALUE_BEHAVIOR.lastViolationValue.number",
            type: "Number",
        },
    ],
    deviceDefenderIndexingMode: "VIOLATIONS",
    filter: {
        namedShadowNames: ["thing1shadow"],
    },
    namedShadowIndexingMode: "ON",
    thingConnectivityIndexingMode: "STATUS",
    thingIndexingMode: "REGISTRY_AND_SHADOW",
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iot.IndexingConfiguration("example", thing_indexing_configuration=aws.iot.IndexingConfigurationThingIndexingConfigurationArgs(
    custom_fields=[
        aws.iot.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs(
            name="shadow.desired.power",
            type="Boolean",
        ),
        aws.iot.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs(
            name="attributes.version",
            type="Number",
        ),
        aws.iot.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs(
            name="shadow.name.thing1shadow.desired.DefaultDesired",
            type="String",
        ),
        aws.iot.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs(
            name="deviceDefender.securityProfile1.NUMBER_VALUE_BEHAVIOR.lastViolationValue.number",
            type="Number",
        ),
    ],
    device_defender_indexing_mode="VIOLATIONS",
    filter=aws.iot.IndexingConfigurationThingIndexingConfigurationFilterArgs(
        named_shadow_names=["thing1shadow"],
    ),
    named_shadow_indexing_mode="ON",
    thing_connectivity_indexing_mode="STATUS",
    thing_indexing_mode="REGISTRY_AND_SHADOW",
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Iot.IndexingConfiguration("example", new()
    {
        ThingIndexingConfiguration = new Aws.Iot.Inputs.IndexingConfigurationThingIndexingConfigurationArgs
        {
            CustomFields = new[]
            {
                new Aws.Iot.Inputs.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs
                {
                    Name = "shadow.desired.power",
                    Type = "Boolean",
                },
                new Aws.Iot.Inputs.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs
                {
                    Name = "attributes.version",
                    Type = "Number",
                },
                new Aws.Iot.Inputs.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs
                {
                    Name = "shadow.name.thing1shadow.desired.DefaultDesired",
                    Type = "String",
                },
                new Aws.Iot.Inputs.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs
                {
                    Name = "deviceDefender.securityProfile1.NUMBER_VALUE_BEHAVIOR.lastViolationValue.number",
                    Type = "Number",
                },
            },
            DeviceDefenderIndexingMode = "VIOLATIONS",
            Filter = new Aws.Iot.Inputs.IndexingConfigurationThingIndexingConfigurationFilterArgs
            {
                NamedShadowNames = new[]
                {
                    "thing1shadow",
                },
            },
            NamedShadowIndexingMode = "ON",
            ThingConnectivityIndexingMode = "STATUS",
            ThingIndexingMode = "REGISTRY_AND_SHADOW",
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
		_, err := iot.NewIndexingConfiguration(ctx, "example", &iot.IndexingConfigurationArgs{
			ThingIndexingConfiguration: &iot.IndexingConfigurationThingIndexingConfigurationArgs{
				CustomFields: iot.IndexingConfigurationThingIndexingConfigurationCustomFieldArray{
					&iot.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs{
						Name: pulumi.String("shadow.desired.power"),
						Type: pulumi.String("Boolean"),
					},
					&iot.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs{
						Name: pulumi.String("attributes.version"),
						Type: pulumi.String("Number"),
					},
					&iot.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs{
						Name: pulumi.String("shadow.name.thing1shadow.desired.DefaultDesired"),
						Type: pulumi.String("String"),
					},
					&iot.IndexingConfigurationThingIndexingConfigurationCustomFieldArgs{
						Name: pulumi.String("deviceDefender.securityProfile1.NUMBER_VALUE_BEHAVIOR.lastViolationValue.number"),
						Type: pulumi.String("Number"),
					},
				},
				DeviceDefenderIndexingMode: pulumi.String("VIOLATIONS"),
				Filter: &iot.IndexingConfigurationThingIndexingConfigurationFilterArgs{
					NamedShadowNames: pulumi.StringArray{
						pulumi.String("thing1shadow"),
					},
				},
				NamedShadowIndexingMode:       pulumi.String("ON"),
				ThingConnectivityIndexingMode: pulumi.String("STATUS"),
				ThingIndexingMode:             pulumi.String("REGISTRY_AND_SHADOW"),
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
import com.pulumi.aws.iot.IndexingConfiguration;
import com.pulumi.aws.iot.IndexingConfigurationArgs;
import com.pulumi.aws.iot.inputs.IndexingConfigurationThingIndexingConfigurationArgs;
import com.pulumi.aws.iot.inputs.IndexingConfigurationThingIndexingConfigurationFilterArgs;
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
        var example = new IndexingConfiguration("example", IndexingConfigurationArgs.builder()        
            .thingIndexingConfiguration(IndexingConfigurationThingIndexingConfigurationArgs.builder()
                .customFields(                
                    IndexingConfigurationThingIndexingConfigurationCustomFieldArgs.builder()
                        .name("shadow.desired.power")
                        .type("Boolean")
                        .build(),
                    IndexingConfigurationThingIndexingConfigurationCustomFieldArgs.builder()
                        .name("attributes.version")
                        .type("Number")
                        .build(),
                    IndexingConfigurationThingIndexingConfigurationCustomFieldArgs.builder()
                        .name("shadow.name.thing1shadow.desired.DefaultDesired")
                        .type("String")
                        .build(),
                    IndexingConfigurationThingIndexingConfigurationCustomFieldArgs.builder()
                        .name("deviceDefender.securityProfile1.NUMBER_VALUE_BEHAVIOR.lastViolationValue.number")
                        .type("Number")
                        .build())
                .deviceDefenderIndexingMode("VIOLATIONS")
                .filter(IndexingConfigurationThingIndexingConfigurationFilterArgs.builder()
                    .namedShadowNames("thing1shadow")
                    .build())
                .namedShadowIndexingMode("ON")
                .thingConnectivityIndexingMode("STATUS")
                .thingIndexingMode("REGISTRY_AND_SHADOW")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:iot:IndexingConfiguration
    properties:
      thingIndexingConfiguration:
        customFields:
          - name: shadow.desired.power
            type: Boolean
          - name: attributes.version
            type: Number
          - name: shadow.name.thing1shadow.desired.DefaultDesired
            type: String
          - name: deviceDefender.securityProfile1.NUMBER_VALUE_BEHAVIOR.lastViolationValue.number
            type: Number
        deviceDefenderIndexingMode: VIOLATIONS
        filter:
          namedShadowNames:
            - thing1shadow
        namedShadowIndexingMode: ON
        thingConnectivityIndexingMode: STATUS
        thingIndexingMode: REGISTRY_AND_SHADOW
```
{{% /example %}}
{{% /examples %}}