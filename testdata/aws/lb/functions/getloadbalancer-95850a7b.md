> **Note:** `aws.alb.LoadBalancer` is known as `aws.lb.LoadBalancer`. The functionality is identical.

Provides information about a Load Balancer.

This data source can prove useful when a module accepts an LB as an input
variable and needs to, for example, determine the security groups associated
with it, etc.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const lbArn = config.get("lbArn") || "";
const lbName = config.get("lbName") || "";
const test = aws.lb.getLoadBalancer({
    arn: lbArn,
    name: lbName,
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
lb_arn = config.get("lbArn")
if lb_arn is None:
    lb_arn = ""
lb_name = config.get("lbName")
if lb_name is None:
    lb_name = ""
test = aws.lb.get_load_balancer(arn=lb_arn,
    name=lb_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var lbArn = config.Get("lbArn") ?? "";
    var lbName = config.Get("lbName") ?? "";
    var test = Aws.LB.GetLoadBalancer.Invoke(new()
    {
        Arn = lbArn,
        Name = lbName,
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
		lbArn := ""
		if param := cfg.Get("lbArn"); param != "" {
			lbArn = param
		}
		lbName := ""
		if param := cfg.Get("lbName"); param != "" {
			lbName = param
		}
		_, err := lb.LookupLoadBalancer(ctx, &lb.LookupLoadBalancerArgs{
			Arn:  pulumi.StringRef(lbArn),
			Name: pulumi.StringRef(lbName),
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
import com.pulumi.aws.lb.inputs.GetLoadBalancerArgs;
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
        final var lbArn = config.get("lbArn").orElse("");
        final var lbName = config.get("lbName").orElse("");
        final var test = LbFunctions.getLoadBalancer(GetLoadBalancerArgs.builder()
            .arn(lbArn)
            .name(lbName)
            .build());

    }
}
```
```yaml
configuration:
  lbArn:
    type: string
    default:
  lbName:
    type: string
    default:
variables:
  test:
    fn::invoke:
      Function: aws:lb:getLoadBalancer
      Arguments:
        arn: ${lbArn}
        name: ${lbName}
```
{{% /example %}}
{{% /examples %}}