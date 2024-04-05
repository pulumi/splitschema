> **Note:** `aws.alb.Listener` is known as `aws.lb.Listener`. The functionality is identical.

Provides information about a Load Balancer Listener.

This data source can prove useful when a module accepts an LB Listener as an input variable and needs to know the LB it is attached to, or other information specific to the listener in question.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const listenerArn = config.require("listenerArn");
const listener = aws.lb.getListener({
    arn: listenerArn,
});
const selected = aws.lb.getLoadBalancer({
    name: "default-public",
});
const selected443 = selected.then(selected => aws.lb.getListener({
    loadBalancerArn: selected.arn,
    port: 443,
}));
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
listener_arn = config.require("listenerArn")
listener = aws.lb.get_listener(arn=listener_arn)
selected = aws.lb.get_load_balancer(name="default-public")
selected443 = aws.lb.get_listener(load_balancer_arn=selected.arn,
    port=443)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var listenerArn = config.Require("listenerArn");
    var listener = Aws.LB.GetListener.Invoke(new()
    {
        Arn = listenerArn,
    });

    var selected = Aws.LB.GetLoadBalancer.Invoke(new()
    {
        Name = "default-public",
    });

    var selected443 = Aws.LB.GetListener.Invoke(new()
    {
        LoadBalancerArn = selected.Apply(getLoadBalancerResult => getLoadBalancerResult.Arn),
        Port = 443,
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
		listenerArn := cfg.Require("listenerArn")
		_, err := lb.LookupListener(ctx, &lb.LookupListenerArgs{
			Arn: pulumi.StringRef(listenerArn),
		}, nil)
		if err != nil {
			return err
		}
		selected, err := lb.LookupLoadBalancer(ctx, &lb.LookupLoadBalancerArgs{
			Name: pulumi.StringRef("default-public"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = lb.LookupListener(ctx, &lb.LookupListenerArgs{
			LoadBalancerArn: pulumi.StringRef(selected.Arn),
			Port:            pulumi.IntRef(443),
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
import com.pulumi.aws.lb.inputs.GetListenerArgs;
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
        final var listenerArn = config.get("listenerArn");
        final var listener = LbFunctions.getListener(GetListenerArgs.builder()
            .arn(listenerArn)
            .build());

        final var selected = LbFunctions.getLoadBalancer(GetLoadBalancerArgs.builder()
            .name("default-public")
            .build());

        final var selected443 = LbFunctions.getListener(GetListenerArgs.builder()
            .loadBalancerArn(selected.applyValue(getLoadBalancerResult -> getLoadBalancerResult.arn()))
            .port(443)
            .build());

    }
}
```
```yaml
configuration:
  # get listener from listener arn
  listenerArn:
    type: string
variables:
  listener: # get listener from load_balancer_arn and port
    fn::invoke:
      Function: aws:lb:getListener
      Arguments:
        arn: ${listenerArn}
  selected:
    fn::invoke:
      Function: aws:lb:getLoadBalancer
      Arguments:
        name: default-public
  selected443:
    fn::invoke:
      Function: aws:lb:getListener
      Arguments:
        loadBalancerArn: ${selected.arn}
        port: 443
```
{{% /example %}}
{{% /examples %}}