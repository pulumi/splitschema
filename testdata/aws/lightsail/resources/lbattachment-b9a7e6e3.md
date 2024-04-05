Attaches a Lightsail Instance to a Lightsail Load Balancer.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const available = aws.getAvailabilityZones({
    state: "available",
    filters: [{
        name: "opt-in-status",
        values: ["opt-in-not-required"],
    }],
});
const testLb = new aws.lightsail.Lb("testLb", {
    healthCheckPath: "/",
    instancePort: 80,
    tags: {
        foo: "bar",
    },
});
const testInstance = new aws.lightsail.Instance("testInstance", {
    availabilityZone: available.then(available => available.names?.[0]),
    blueprintId: "amazon_linux_2",
    bundleId: "nano_1_0",
});
const testLbAttachment = new aws.lightsail.LbAttachment("testLbAttachment", {
    lbName: testLb.name,
    instanceName: testInstance.name,
});
```
```python
import pulumi
import pulumi_aws as aws

available = aws.get_availability_zones(state="available",
    filters=[aws.GetAvailabilityZonesFilterArgs(
        name="opt-in-status",
        values=["opt-in-not-required"],
    )])
test_lb = aws.lightsail.Lb("testLb",
    health_check_path="/",
    instance_port=80,
    tags={
        "foo": "bar",
    })
test_instance = aws.lightsail.Instance("testInstance",
    availability_zone=available.names[0],
    blueprint_id="amazon_linux_2",
    bundle_id="nano_1_0")
test_lb_attachment = aws.lightsail.LbAttachment("testLbAttachment",
    lb_name=test_lb.name,
    instance_name=test_instance.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var available = Aws.GetAvailabilityZones.Invoke(new()
    {
        State = "available",
        Filters = new[]
        {
            new Aws.Inputs.GetAvailabilityZonesFilterInputArgs
            {
                Name = "opt-in-status",
                Values = new[]
                {
                    "opt-in-not-required",
                },
            },
        },
    });

    var testLb = new Aws.LightSail.Lb("testLb", new()
    {
        HealthCheckPath = "/",
        InstancePort = 80,
        Tags = 
        {
            { "foo", "bar" },
        },
    });

    var testInstance = new Aws.LightSail.Instance("testInstance", new()
    {
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[0]),
        BlueprintId = "amazon_linux_2",
        BundleId = "nano_1_0",
    });

    var testLbAttachment = new Aws.LightSail.LbAttachment("testLbAttachment", new()
    {
        LbName = testLb.Name,
        InstanceName = testInstance.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		available, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			State: pulumi.StringRef("available"),
			Filters: []aws.GetAvailabilityZonesFilter{
				{
					Name: "opt-in-status",
					Values: []string{
						"opt-in-not-required",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
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
		testInstance, err := lightsail.NewInstance(ctx, "testInstance", &lightsail.InstanceArgs{
			AvailabilityZone: *pulumi.String(available.Names[0]),
			BlueprintId:      pulumi.String("amazon_linux_2"),
			BundleId:         pulumi.String("nano_1_0"),
		})
		if err != nil {
			return err
		}
		_, err = lightsail.NewLbAttachment(ctx, "testLbAttachment", &lightsail.LbAttachmentArgs{
			LbName:       testLb.Name,
			InstanceName: testInstance.Name,
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
import com.pulumi.aws.lightsail.Lb;
import com.pulumi.aws.lightsail.LbArgs;
import com.pulumi.aws.lightsail.Instance;
import com.pulumi.aws.lightsail.InstanceArgs;
import com.pulumi.aws.lightsail.LbAttachment;
import com.pulumi.aws.lightsail.LbAttachmentArgs;
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
        final var available = AwsFunctions.getAvailabilityZones(GetAvailabilityZonesArgs.builder()
            .state("available")
            .filters(GetAvailabilityZonesFilterArgs.builder()
                .name("opt-in-status")
                .values("opt-in-not-required")
                .build())
            .build());

        var testLb = new Lb("testLb", LbArgs.builder()        
            .healthCheckPath("/")
            .instancePort("80")
            .tags(Map.of("foo", "bar"))
            .build());

        var testInstance = new Instance("testInstance", InstanceArgs.builder()        
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[0]))
            .blueprintId("amazon_linux_2")
            .bundleId("nano_1_0")
            .build());

        var testLbAttachment = new LbAttachment("testLbAttachment", LbAttachmentArgs.builder()        
            .lbName(testLb.name())
            .instanceName(testInstance.name())
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
  testInstance:
    type: aws:lightsail:Instance
    properties:
      availabilityZone: ${available.names[0]}
      blueprintId: amazon_linux_2
      bundleId: nano_1_0
  testLbAttachment:
    type: aws:lightsail:LbAttachment
    properties:
      lbName: ${testLb.name}
      instanceName: ${testInstance.name}
variables:
  available:
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments:
        state: available
        filters:
          - name: opt-in-status
            values:
              - opt-in-not-required
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_lightsail_lb_attachment` using the name attribute. For example:

```sh
 $ pulumi import aws:lightsail/lbAttachment:LbAttachment test example-load-balancer,example-instance
```
 