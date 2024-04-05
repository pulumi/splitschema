Provides a Global Accelerator custom routing listener.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCustomRoutingAccelerator = new aws.globalaccelerator.CustomRoutingAccelerator("exampleCustomRoutingAccelerator", {
    ipAddressType: "IPV4",
    enabled: true,
    attributes: {
        flowLogsEnabled: true,
        flowLogsS3Bucket: "example-bucket",
        flowLogsS3Prefix: "flow-logs/",
    },
});
const exampleCustomRoutingListener = new aws.globalaccelerator.CustomRoutingListener("exampleCustomRoutingListener", {
    acceleratorArn: exampleCustomRoutingAccelerator.id,
    portRanges: [{
        fromPort: 80,
        toPort: 80,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_custom_routing_accelerator = aws.globalaccelerator.CustomRoutingAccelerator("exampleCustomRoutingAccelerator",
    ip_address_type="IPV4",
    enabled=True,
    attributes=aws.globalaccelerator.CustomRoutingAcceleratorAttributesArgs(
        flow_logs_enabled=True,
        flow_logs_s3_bucket="example-bucket",
        flow_logs_s3_prefix="flow-logs/",
    ))
example_custom_routing_listener = aws.globalaccelerator.CustomRoutingListener("exampleCustomRoutingListener",
    accelerator_arn=example_custom_routing_accelerator.id,
    port_ranges=[aws.globalaccelerator.CustomRoutingListenerPortRangeArgs(
        from_port=80,
        to_port=80,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleCustomRoutingAccelerator = new Aws.GlobalAccelerator.CustomRoutingAccelerator("exampleCustomRoutingAccelerator", new()
    {
        IpAddressType = "IPV4",
        Enabled = true,
        Attributes = new Aws.GlobalAccelerator.Inputs.CustomRoutingAcceleratorAttributesArgs
        {
            FlowLogsEnabled = true,
            FlowLogsS3Bucket = "example-bucket",
            FlowLogsS3Prefix = "flow-logs/",
        },
    });

    var exampleCustomRoutingListener = new Aws.GlobalAccelerator.CustomRoutingListener("exampleCustomRoutingListener", new()
    {
        AcceleratorArn = exampleCustomRoutingAccelerator.Id,
        PortRanges = new[]
        {
            new Aws.GlobalAccelerator.Inputs.CustomRoutingListenerPortRangeArgs
            {
                FromPort = 80,
                ToPort = 80,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/globalaccelerator"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleCustomRoutingAccelerator, err := globalaccelerator.NewCustomRoutingAccelerator(ctx, "exampleCustomRoutingAccelerator", &globalaccelerator.CustomRoutingAcceleratorArgs{
			IpAddressType: pulumi.String("IPV4"),
			Enabled:       pulumi.Bool(true),
			Attributes: &globalaccelerator.CustomRoutingAcceleratorAttributesArgs{
				FlowLogsEnabled:  pulumi.Bool(true),
				FlowLogsS3Bucket: pulumi.String("example-bucket"),
				FlowLogsS3Prefix: pulumi.String("flow-logs/"),
			},
		})
		if err != nil {
			return err
		}
		_, err = globalaccelerator.NewCustomRoutingListener(ctx, "exampleCustomRoutingListener", &globalaccelerator.CustomRoutingListenerArgs{
			AcceleratorArn: exampleCustomRoutingAccelerator.ID(),
			PortRanges: globalaccelerator.CustomRoutingListenerPortRangeArray{
				&globalaccelerator.CustomRoutingListenerPortRangeArgs{
					FromPort: pulumi.Int(80),
					ToPort:   pulumi.Int(80),
				},
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
import com.pulumi.aws.globalaccelerator.CustomRoutingAccelerator;
import com.pulumi.aws.globalaccelerator.CustomRoutingAcceleratorArgs;
import com.pulumi.aws.globalaccelerator.inputs.CustomRoutingAcceleratorAttributesArgs;
import com.pulumi.aws.globalaccelerator.CustomRoutingListener;
import com.pulumi.aws.globalaccelerator.CustomRoutingListenerArgs;
import com.pulumi.aws.globalaccelerator.inputs.CustomRoutingListenerPortRangeArgs;
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
        var exampleCustomRoutingAccelerator = new CustomRoutingAccelerator("exampleCustomRoutingAccelerator", CustomRoutingAcceleratorArgs.builder()        
            .ipAddressType("IPV4")
            .enabled(true)
            .attributes(CustomRoutingAcceleratorAttributesArgs.builder()
                .flowLogsEnabled(true)
                .flowLogsS3Bucket("example-bucket")
                .flowLogsS3Prefix("flow-logs/")
                .build())
            .build());

        var exampleCustomRoutingListener = new CustomRoutingListener("exampleCustomRoutingListener", CustomRoutingListenerArgs.builder()        
            .acceleratorArn(exampleCustomRoutingAccelerator.id())
            .portRanges(CustomRoutingListenerPortRangeArgs.builder()
                .fromPort(80)
                .toPort(80)
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleCustomRoutingAccelerator:
    type: aws:globalaccelerator:CustomRoutingAccelerator
    properties:
      ipAddressType: IPV4
      enabled: true
      attributes:
        flowLogsEnabled: true
        flowLogsS3Bucket: example-bucket
        flowLogsS3Prefix: flow-logs/
  exampleCustomRoutingListener:
    type: aws:globalaccelerator:CustomRoutingListener
    properties:
      acceleratorArn: ${exampleCustomRoutingAccelerator.id}
      portRanges:
        - fromPort: 80
          toPort: 80
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Global Accelerator custom routing listeners using the `id`. For example:

```sh
 $ pulumi import aws:globalaccelerator/customRoutingListener:CustomRoutingListener example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxxx
```
 