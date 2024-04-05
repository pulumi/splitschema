Provides a GameLift Alias resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.gamelift.Alias("example", {
    description: "Example Description",
    routingStrategy: {
        message: "Example Message",
        type: "TERMINAL",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.gamelift.Alias("example",
    description="Example Description",
    routing_strategy=aws.gamelift.AliasRoutingStrategyArgs(
        message="Example Message",
        type="TERMINAL",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.GameLift.Alias("example", new()
    {
        Description = "Example Description",
        RoutingStrategy = new Aws.GameLift.Inputs.AliasRoutingStrategyArgs
        {
            Message = "Example Message",
            Type = "TERMINAL",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/gamelift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := gamelift.NewAlias(ctx, "example", &gamelift.AliasArgs{
			Description: pulumi.String("Example Description"),
			RoutingStrategy: &gamelift.AliasRoutingStrategyArgs{
				Message: pulumi.String("Example Message"),
				Type:    pulumi.String("TERMINAL"),
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
import com.pulumi.aws.gamelift.Alias;
import com.pulumi.aws.gamelift.AliasArgs;
import com.pulumi.aws.gamelift.inputs.AliasRoutingStrategyArgs;
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
        var example = new Alias("example", AliasArgs.builder()        
            .description("Example Description")
            .routingStrategy(AliasRoutingStrategyArgs.builder()
                .message("Example Message")
                .type("TERMINAL")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:gamelift:Alias
    properties:
      description: Example Description
      routingStrategy:
        message: Example Message
        type: TERMINAL
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import GameLift Aliases using the ID. For example:

```sh
 $ pulumi import aws:gamelift/alias:Alias example <alias-id>
```
 