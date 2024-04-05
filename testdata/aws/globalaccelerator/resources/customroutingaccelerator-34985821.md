Creates a Global Accelerator custom routing accelerator.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.globalaccelerator.CustomRoutingAccelerator("example", {
    attributes: {
        flowLogsEnabled: true,
        flowLogsS3Bucket: "example-bucket",
        flowLogsS3Prefix: "flow-logs/",
    },
    enabled: true,
    ipAddressType: "IPV4",
    ipAddresses: ["1.2.3.4"],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.globalaccelerator.CustomRoutingAccelerator("example",
    attributes=aws.globalaccelerator.CustomRoutingAcceleratorAttributesArgs(
        flow_logs_enabled=True,
        flow_logs_s3_bucket="example-bucket",
        flow_logs_s3_prefix="flow-logs/",
    ),
    enabled=True,
    ip_address_type="IPV4",
    ip_addresses=["1.2.3.4"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.GlobalAccelerator.CustomRoutingAccelerator("example", new()
    {
        Attributes = new Aws.GlobalAccelerator.Inputs.CustomRoutingAcceleratorAttributesArgs
        {
            FlowLogsEnabled = true,
            FlowLogsS3Bucket = "example-bucket",
            FlowLogsS3Prefix = "flow-logs/",
        },
        Enabled = true,
        IpAddressType = "IPV4",
        IpAddresses = new[]
        {
            "1.2.3.4",
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
		_, err := globalaccelerator.NewCustomRoutingAccelerator(ctx, "example", &globalaccelerator.CustomRoutingAcceleratorArgs{
			Attributes: &globalaccelerator.CustomRoutingAcceleratorAttributesArgs{
				FlowLogsEnabled:  pulumi.Bool(true),
				FlowLogsS3Bucket: pulumi.String("example-bucket"),
				FlowLogsS3Prefix: pulumi.String("flow-logs/"),
			},
			Enabled:       pulumi.Bool(true),
			IpAddressType: pulumi.String("IPV4"),
			IpAddresses: pulumi.StringArray{
				pulumi.String("1.2.3.4"),
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
        var example = new CustomRoutingAccelerator("example", CustomRoutingAcceleratorArgs.builder()        
            .attributes(CustomRoutingAcceleratorAttributesArgs.builder()
                .flowLogsEnabled(true)
                .flowLogsS3Bucket("example-bucket")
                .flowLogsS3Prefix("flow-logs/")
                .build())
            .enabled(true)
            .ipAddressType("IPV4")
            .ipAddresses("1.2.3.4")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:globalaccelerator:CustomRoutingAccelerator
    properties:
      attributes:
        flowLogsEnabled: true
        flowLogsS3Bucket: example-bucket
        flowLogsS3Prefix: flow-logs/
      enabled: true
      ipAddressType: IPV4
      ipAddresses:
        - 1.2.3.4
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Global Accelerator custom routing accelerators using the `arn`. For example:

```sh
 $ pulumi import aws:globalaccelerator/customRoutingAccelerator:CustomRoutingAccelerator example arn:aws:globalaccelerator::111111111111:accelerator/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
```
 