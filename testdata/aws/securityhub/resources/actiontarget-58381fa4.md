Creates Security Hub custom action.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const exampleActionTarget = new aws.securityhub.ActionTarget("exampleActionTarget", {
    identifier: "SendToChat",
    description: "This is custom action sends selected findings to chat",
}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_action_target = aws.securityhub.ActionTarget("exampleActionTarget",
    identifier="SendToChat",
    description="This is custom action sends selected findings to chat",
    opts=pulumi.ResourceOptions(depends_on=[example_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.SecurityHub.Account("exampleAccount");

    var exampleActionTarget = new Aws.SecurityHub.ActionTarget("exampleActionTarget", new()
    {
        Identifier = "SendToChat",
        Description = "This is custom action sends selected findings to chat",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewActionTarget(ctx, "exampleActionTarget", &securityhub.ActionTargetArgs{
			Identifier:  pulumi.String("SendToChat"),
			Description: pulumi.String("This is custom action sends selected findings to chat"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccount,
		}))
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
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.ActionTarget;
import com.pulumi.aws.securityhub.ActionTargetArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleActionTarget = new ActionTarget("exampleActionTarget", ActionTargetArgs.builder()        
            .identifier("SendToChat")
            .description("This is custom action sends selected findings to chat")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:securityhub:Account
  exampleActionTarget:
    type: aws:securityhub:ActionTarget
    properties:
      identifier: SendToChat
      description: This is custom action sends selected findings to chat
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Security Hub custom action using the action target ARN. For example:

```sh
 $ pulumi import aws:securityhub/actionTarget:ActionTarget example arn:aws:securityhub:eu-west-1:312940875350:action/custom/a
```
 