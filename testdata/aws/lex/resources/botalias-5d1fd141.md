Provides an Amazon Lex Bot Alias resource. For more information see
[Amazon Lex: How It Works](https://docs.aws.amazon.com/lex/latest/dg/how-it-works.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const orderFlowersProd = new aws.lex.BotAlias("orderFlowersProd", {
    botName: "OrderFlowers",
    botVersion: "1",
    description: "Production Version of the OrderFlowers Bot.",
    name: "OrderFlowersProd",
});
```
```python
import pulumi
import pulumi_aws as aws

order_flowers_prod = aws.lex.BotAlias("orderFlowersProd",
    bot_name="OrderFlowers",
    bot_version="1",
    description="Production Version of the OrderFlowers Bot.",
    name="OrderFlowersProd")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var orderFlowersProd = new Aws.Lex.BotAlias("orderFlowersProd", new()
    {
        BotName = "OrderFlowers",
        BotVersion = "1",
        Description = "Production Version of the OrderFlowers Bot.",
        Name = "OrderFlowersProd",
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
		_, err := lex.NewBotAlias(ctx, "orderFlowersProd", &lex.BotAliasArgs{
			BotName:     pulumi.String("OrderFlowers"),
			BotVersion:  pulumi.String("1"),
			Description: pulumi.String("Production Version of the OrderFlowers Bot."),
			Name:        pulumi.String("OrderFlowersProd"),
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
import com.pulumi.aws.lex.BotAlias;
import com.pulumi.aws.lex.BotAliasArgs;
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
        var orderFlowersProd = new BotAlias("orderFlowersProd", BotAliasArgs.builder()        
            .botName("OrderFlowers")
            .botVersion("1")
            .description("Production Version of the OrderFlowers Bot.")
            .name("OrderFlowersProd")
            .build());

    }
}
```
```yaml
resources:
  orderFlowersProd:
    type: aws:lex:BotAlias
    properties:
      botName: OrderFlowers
      botVersion: '1'
      description: Production Version of the OrderFlowers Bot.
      name: OrderFlowersProd
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import bot aliases using an ID with the format `bot_name:bot_alias_name`. For example:

```sh
 $ pulumi import aws:lex/botAlias:BotAlias order_flowers_prod OrderFlowers:OrderFlowersProd
```
 