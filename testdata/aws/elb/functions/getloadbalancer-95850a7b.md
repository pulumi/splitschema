Provides information about a "classic" Elastic Load Balancer (ELB).
See LB Data Source if you are looking for "v2"
Application Load Balancer (ALB) or Network Load Balancer (NLB).

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
const lbName = config.get("lbName") || "";
const test = aws.elb.getLoadBalancer({
    name: lbName,
});
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
lb_name = config.get("lbName")
if lb_name is None:
    lb_name = ""
test = aws.elb.get_load_balancer(name=lb_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var lbName = config.Get("lbName") ?? "";
    var test = Aws.Elb.GetLoadBalancer.Invoke(new()
    {
        Name = lbName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		lbName := ""
		if param := cfg.Get("lbName"); param != "" {
			lbName = param
		}
		_, err := elb.LookupLoadBalancer(ctx, &elb.LookupLoadBalancerArgs{
			Name: lbName,
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
import com.pulumi.aws.elb.ElbFunctions;
import com.pulumi.aws.elb.inputs.GetLoadBalancerArgs;
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
        final var lbName = config.get("lbName").orElse("");
        final var test = ElbFunctions.getLoadBalancer(GetLoadBalancerArgs.builder()
            .name(lbName)
            .build());

    }
}
```
```yaml
configuration:
  lbName:
    type: string
    default:
variables:
  test:
    fn::invoke:
      Function: aws:elb:getLoadBalancer
      Arguments:
        name: ${lbName}
```
{{% /example %}}
{{% /examples %}}