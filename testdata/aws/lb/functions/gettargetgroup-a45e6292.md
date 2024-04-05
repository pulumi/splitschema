> **Note:** `aws.alb.TargetGroup` is known as `aws.lb.TargetGroup`. The functionality is identical.

Provides information about a Load Balancer Target Group.

This data source can prove useful when a module accepts an LB Target Group as an
input variable and needs to know its attributes. It can also be used to get the ARN of
an LB Target Group for use in other resources, given LB Target Group name.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const lbTgArn = config.get("lbTgArn") || "";
const lbTgName = config.get("lbTgName") || "";
const test = aws.lb.getTargetGroup({
    arn: lbTgArn,
    name: lbTgName,
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
lb_tg_arn = config.get("lbTgArn")
if lb_tg_arn is None:
    lb_tg_arn = ""
lb_tg_name = config.get("lbTgName")
if lb_tg_name is None:
    lb_tg_name = ""
test = aws.lb.get_target_group(arn=lb_tg_arn,
    name=lb_tg_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var lbTgArn = config.Get("lbTgArn") ?? "";
    var lbTgName = config.Get("lbTgName") ?? "";
    var test = Aws.LB.GetTargetGroup.Invoke(new()
    {
        Arn = lbTgArn,
        Name = lbTgName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		lbTgArn := ""
		if param := cfg.Get("lbTgArn"); param != "" {
			lbTgArn = param
		}
		lbTgName := ""
		if param := cfg.Get("lbTgName"); param != "" {
			lbTgName = param
		}
		_, err := lb.LookupTargetGroup(ctx, &lb.LookupTargetGroupArgs{
			Arn:  pulumi.StringRef(lbTgArn),
			Name: pulumi.StringRef(lbTgName),
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
import com.pulumi.aws.lb.LbFunctions;
import com.pulumi.aws.lb.inputs.GetTargetGroupArgs;
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
        final var config = ctx.config();
        final var lbTgArn = config.get("lbTgArn").orElse("");
        final var lbTgName = config.get("lbTgName").orElse("");
        final var test = LbFunctions.getTargetGroup(GetTargetGroupArgs.builder()
            .arn(lbTgArn)
            .name(lbTgName)
            .build());

    }
}
```
```yaml
configuration:
  lbTgArn:
    type: string
    default:
  lbTgName:
    type: string
    default:
variables:
  test:
    fn::invoke:
      Function: aws:lb:getTargetGroup
      Arguments:
        arn: ${lbTgArn}
        name: ${lbTgName}
```
{{% /example %}}
{{% /examples %}}