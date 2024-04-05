Provides a Global Accelerator listener.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccelerator = new aws.globalaccelerator.Accelerator("exampleAccelerator", {
    ipAddressType: "IPV4",
    enabled: true,
    attributes: {
        flowLogsEnabled: true,
        flowLogsS3Bucket: "example-bucket",
        flowLogsS3Prefix: "flow-logs/",
    },
});
const exampleListener = new aws.globalaccelerator.Listener("exampleListener", {
    acceleratorArn: exampleAccelerator.id,
    clientAffinity: "SOURCE_IP",
    protocol: "TCP",
    portRanges: [{
        fromPort: 80,
        toPort: 80,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example_accelerator = aws.globalaccelerator.Accelerator("exampleAccelerator",
    ip_address_type="IPV4",
    enabled=True,
    attributes=aws.globalaccelerator.AcceleratorAttributesArgs(
        flow_logs_enabled=True,
        flow_logs_s3_bucket="example-bucket",
        flow_logs_s3_prefix="flow-logs/",
    ))
example_listener = aws.globalaccelerator.Listener("exampleListener",
    accelerator_arn=example_accelerator.id,
    client_affinity="SOURCE_IP",
    protocol="TCP",
    port_ranges=[aws.globalaccelerator.ListenerPortRangeArgs(
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
    var exampleAccelerator = new Aws.GlobalAccelerator.Accelerator("exampleAccelerator", new()
    {
        IpAddressType = "IPV4",
        Enabled = true,
        Attributes = new Aws.GlobalAccelerator.Inputs.AcceleratorAttributesArgs
        {
            FlowLogsEnabled = true,
            FlowLogsS3Bucket = "example-bucket",
            FlowLogsS3Prefix = "flow-logs/",
        },
    });

    var exampleListener = new Aws.GlobalAccelerator.Listener("exampleListener", new()
    {
        AcceleratorArn = exampleAccelerator.Id,
        ClientAffinity = "SOURCE_IP",
        Protocol = "TCP",
        PortRanges = new[]
        {
            new Aws.GlobalAccelerator.Inputs.ListenerPortRangeArgs
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
		exampleAccelerator, err := globalaccelerator.NewAccelerator(ctx, "exampleAccelerator", &globalaccelerator.AcceleratorArgs{
			IpAddressType: pulumi.String("IPV4"),
			Enabled:       pulumi.Bool(true),
			Attributes: &globalaccelerator.AcceleratorAttributesArgs{
				FlowLogsEnabled:  pulumi.Bool(true),
				FlowLogsS3Bucket: pulumi.String("example-bucket"),
				FlowLogsS3Prefix: pulumi.String("flow-logs/"),
			},
		})
		if err != nil {
			return err
		}
		_, err = globalaccelerator.NewListener(ctx, "exampleListener", &globalaccelerator.ListenerArgs{
			AcceleratorArn: exampleAccelerator.ID(),
			ClientAffinity: pulumi.String("SOURCE_IP"),
			Protocol:       pulumi.String("TCP"),
			PortRanges: globalaccelerator.ListenerPortRangeArray{
				&globalaccelerator.ListenerPortRangeArgs{
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
import com.pulumi.aws.globalaccelerator.Accelerator;
import com.pulumi.aws.globalaccelerator.AcceleratorArgs;
import com.pulumi.aws.globalaccelerator.inputs.AcceleratorAttributesArgs;
import com.pulumi.aws.globalaccelerator.Listener;
import com.pulumi.aws.globalaccelerator.ListenerArgs;
import com.pulumi.aws.globalaccelerator.inputs.ListenerPortRangeArgs;
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
        var exampleAccelerator = new Accelerator("exampleAccelerator", AcceleratorArgs.builder()        
            .ipAddressType("IPV4")
            .enabled(true)
            .attributes(AcceleratorAttributesArgs.builder()
                .flowLogsEnabled(true)
                .flowLogsS3Bucket("example-bucket")
                .flowLogsS3Prefix("flow-logs/")
                .build())
            .build());

        var exampleListener = new Listener("exampleListener", ListenerArgs.builder()        
            .acceleratorArn(exampleAccelerator.id())
            .clientAffinity("SOURCE_IP")
            .protocol("TCP")
            .portRanges(ListenerPortRangeArgs.builder()
                .fromPort(80)
                .toPort(80)
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleAccelerator:
    type: aws:globalaccelerator:Accelerator
    properties:
      ipAddressType: IPV4
      enabled: true
      attributes:
        flowLogsEnabled: true
        flowLogsS3Bucket: example-bucket
        flowLogsS3Prefix: flow-logs/
  exampleListener:
    type: aws:globalaccelerator:Listener
    properties:
      acceleratorArn: ${exampleAccelerator.id}
      clientAffinity: SOURCE_IP
      protocol: TCP
      portRanges:
        - fromPort: 80
          toPort: 80
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Global Accelerator listeners using the `id`. For example:

```sh
 $ pulumi import aws:globalaccelerator/listener:Listener example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/listener/xxxxxxxx
```
 