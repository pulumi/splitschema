Configures Session Stickiness for a Lightsail Load Balancer.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testLb = new aws.lightsail.Lb("testLb", {
    healthCheckPath: "/",
    instancePort: 80,
    tags: {
        foo: "bar",
    },
});
const testLbStickinessPolicy = new aws.lightsail.LbStickinessPolicy("testLbStickinessPolicy", {
    lbName: testLb.name,
    cookieDuration: 900,
    enabled: true,
});
```
```python
import pulumi
import pulumi_aws as aws

test_lb = aws.lightsail.Lb("testLb",
    health_check_path="/",
    instance_port=80,
    tags={
        "foo": "bar",
    })
test_lb_stickiness_policy = aws.lightsail.LbStickinessPolicy("testLbStickinessPolicy",
    lb_name=test_lb.name,
    cookie_duration=900,
    enabled=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testLb = new Aws.LightSail.Lb("testLb", new()
    {
        HealthCheckPath = "/",
        InstancePort = 80,
        Tags = 
        {
            { "foo", "bar" },
        },
    });

    var testLbStickinessPolicy = new Aws.LightSail.LbStickinessPolicy("testLbStickinessPolicy", new()
    {
        LbName = testLb.Name,
        CookieDuration = 900,
        Enabled = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testLb, err := lightsail.NewLb(ctx, "testLb", &lightsail.LbArgs{
			HealthCheckPath: pulumi.String("/"),
			InstancePort:    pulumi.Int(80),
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
			},
		})
		if err != nil {
			return err
		}
		_, err = lightsail.NewLbStickinessPolicy(ctx, "testLbStickinessPolicy", &lightsail.LbStickinessPolicyArgs{
			LbName:         testLb.Name,
			CookieDuration: pulumi.Int(900),
			Enabled:        pulumi.Bool(true),
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
import com.pulumi.aws.lightsail.Lb;
import com.pulumi.aws.lightsail.LbArgs;
import com.pulumi.aws.lightsail.LbStickinessPolicy;
import com.pulumi.aws.lightsail.LbStickinessPolicyArgs;
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
        var testLb = new Lb("testLb", LbArgs.builder()        
            .healthCheckPath("/")
            .instancePort("80")
            .tags(Map.of("foo", "bar"))
            .build());

        var testLbStickinessPolicy = new LbStickinessPolicy("testLbStickinessPolicy", LbStickinessPolicyArgs.builder()        
            .lbName(testLb.name())
            .cookieDuration(900)
            .enabled(true)
            .build());

    }
}
```
```yaml
resources:
  testLb:
    type: aws:lightsail:Lb
    properties:
      healthCheckPath: /
      instancePort: '80'
      tags:
        foo: bar
  testLbStickinessPolicy:
    type: aws:lightsail:LbStickinessPolicy
    properties:
      lbName: ${testLb.name}
      cookieDuration: 900
      enabled: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_lightsail_lb_stickiness_policy` using the `lb_name` attribute. For example:

```sh
 $ pulumi import aws:lightsail/lbStickinessPolicy:LbStickinessPolicy test example-load-balancer
```
 