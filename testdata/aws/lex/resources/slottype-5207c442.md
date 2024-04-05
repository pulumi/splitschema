Provides an Amazon Lex Slot Type resource. For more information see
[Amazon Lex: How It Works](https://docs.aws.amazon.com/lex/latest/dg/how-it-works.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const flowerTypes = new aws.lex.SlotType("flowerTypes", {
    createVersion: true,
    description: "Types of flowers to order",
    enumerationValues: [
        {
            synonyms: [
                "Lirium",
                "Martagon",
            ],
            value: "lilies",
        },
        {
            synonyms: [
                "Eduardoregelia",
                "Podonix",
            ],
            value: "tulips",
        },
    ],
    name: "FlowerTypes",
    valueSelectionStrategy: "ORIGINAL_VALUE",
});
```
```python
import pulumi
import pulumi_aws as aws

flower_types = aws.lex.SlotType("flowerTypes",
    create_version=True,
    description="Types of flowers to order",
    enumeration_values=[
        aws.lex.SlotTypeEnumerationValueArgs(
            synonyms=[
                "Lirium",
                "Martagon",
            ],
            value="lilies",
        ),
        aws.lex.SlotTypeEnumerationValueArgs(
            synonyms=[
                "Eduardoregelia",
                "Podonix",
            ],
            value="tulips",
        ),
    ],
    name="FlowerTypes",
    value_selection_strategy="ORIGINAL_VALUE")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var flowerTypes = new Aws.Lex.SlotType("flowerTypes", new()
    {
        CreateVersion = true,
        Description = "Types of flowers to order",
        EnumerationValues = new[]
        {
            new Aws.Lex.Inputs.SlotTypeEnumerationValueArgs
            {
                Synonyms = new[]
                {
                    "Lirium",
                    "Martagon",
                },
                Value = "lilies",
            },
            new Aws.Lex.Inputs.SlotTypeEnumerationValueArgs
            {
                Synonyms = new[]
                {
                    "Eduardoregelia",
                    "Podonix",
                },
                Value = "tulips",
            },
        },
        Name = "FlowerTypes",
        ValueSelectionStrategy = "ORIGINAL_VALUE",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lex"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lex.NewSlotType(ctx, "flowerTypes", &lex.SlotTypeArgs{
			CreateVersion: pulumi.Bool(true),
			Description:   pulumi.String("Types of flowers to order"),
			EnumerationValues: lex.SlotTypeEnumerationValueArray{
				&lex.SlotTypeEnumerationValueArgs{
					Synonyms: pulumi.StringArray{
						pulumi.String("Lirium"),
						pulumi.String("Martagon"),
					},
					Value: pulumi.String("lilies"),
				},
				&lex.SlotTypeEnumerationValueArgs{
					Synonyms: pulumi.StringArray{
						pulumi.String("Eduardoregelia"),
						pulumi.String("Podonix"),
					},
					Value: pulumi.String("tulips"),
				},
			},
			Name:                   pulumi.String("FlowerTypes"),
			ValueSelectionStrategy: pulumi.String("ORIGINAL_VALUE"),
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
import com.pulumi.aws.lex.SlotType;
import com.pulumi.aws.lex.SlotTypeArgs;
import com.pulumi.aws.lex.inputs.SlotTypeEnumerationValueArgs;
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
        var flowerTypes = new SlotType("flowerTypes", SlotTypeArgs.builder()        
            .createVersion(true)
            .description("Types of flowers to order")
            .enumerationValues(            
                SlotTypeEnumerationValueArgs.builder()
                    .synonyms(                    
                        "Lirium",
                        "Martagon")
                    .value("lilies")
                    .build(),
                SlotTypeEnumerationValueArgs.builder()
                    .synonyms(                    
                        "Eduardoregelia",
                        "Podonix")
                    .value("tulips")
                    .build())
            .name("FlowerTypes")
            .valueSelectionStrategy("ORIGINAL_VALUE")
            .build());

    }
}
```
```yaml
resources:
  flowerTypes:
    type: aws:lex:SlotType
    properties:
      createVersion: true
      description: Types of flowers to order
      enumerationValues:
        - synonyms:
            - Lirium
            - Martagon
          value: lilies
        - synonyms:
            - Eduardoregelia
            - Podonix
          value: tulips
      name: FlowerTypes
      valueSelectionStrategy: ORIGINAL_VALUE
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import slot types using their name. For example:

```sh
 $ pulumi import aws:lex/slotType:SlotType flower_types FlowerTypes
```
 