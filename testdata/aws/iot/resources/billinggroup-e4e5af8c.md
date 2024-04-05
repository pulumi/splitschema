Manages an AWS IoT Billing Group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iot.BillingGroup("example", {
    properties: {
        description: "This is my billing group",
    },
    tags: {
        terraform: "true",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iot.BillingGroup("example",
    properties=aws.iot.BillingGroupPropertiesArgs(
        description="This is my billing group",
    ),
    tags={
        "terraform": "true",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Iot.BillingGroup("example", new()
    {
        Properties = new Aws.Iot.Inputs.BillingGroupPropertiesArgs
        {
            Description = "This is my billing group",
        },
        Tags = 
        {
            { "terraform", "true" },
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
		_, err := iot.NewBillingGroup(ctx, "example", &iot.BillingGroupArgs{
			Properties: &iot.BillingGroupPropertiesArgs{
				Description: pulumi.String("This is my billing group"),
			},
			Tags: pulumi.StringMap{
				"terraform": pulumi.String("true"),
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
import com.pulumi.aws.iot.BillingGroup;
import com.pulumi.aws.iot.BillingGroupArgs;
import com.pulumi.aws.iot.inputs.BillingGroupPropertiesArgs;
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
        var example = new BillingGroup("example", BillingGroupArgs.builder()        
            .properties(BillingGroupPropertiesArgs.builder()
                .description("This is my billing group")
                .build())
            .tags(Map.of("terraform", "true"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:iot:BillingGroup
    properties:
      properties:
        description: This is my billing group
      tags:
        terraform: 'true'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IoT Billing Groups using the name. For example:

```sh
 $ pulumi import aws:iot/billingGroup:BillingGroup example example
```
 