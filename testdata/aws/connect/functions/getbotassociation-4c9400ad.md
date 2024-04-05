Provides details about a specific Lex (V1) Bot associated with an Amazon Connect instance.

{{% examples %}}
## Example Usage
{{% example %}}
### By name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.connect.getBotAssociation({
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    lexBot: {
        name: "Test",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.get_bot_association(instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    lex_bot=aws.connect.GetBotAssociationLexBotArgs(
        name="Test",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Connect.GetBotAssociation.Invoke(new()
    {
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        LexBot = new Aws.Connect.Inputs.GetBotAssociationLexBotInputArgs
        {
            Name = "Test",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.LookupBotAssociation(ctx, &connect.LookupBotAssociationArgs{
			InstanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
			LexBot: connect.GetBotAssociationLexBot{
				Name: "Test",
			},
		}, nil)
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
import com.pulumi.aws.connect.ConnectFunctions;
import com.pulumi.aws.connect.inputs.GetBotAssociationArgs;
import com.pulumi.aws.connect.inputs.GetBotAssociationLexBotArgs;
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
        final var example = ConnectFunctions.getBotAssociation(GetBotAssociationArgs.builder()
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .lexBot(GetBotAssociationLexBotArgs.builder()
                .name("Test")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:connect:getBotAssociation
      Arguments:
        instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
        lexBot:
          name: Test
```
{{% /example %}}
{{% /examples %}}