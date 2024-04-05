Attach an Elastic network interface (ENI) resource with EC2 instance.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ec2.NetworkInterfaceAttachment("test", {
    instanceId: aws_instance.test.id,
    networkInterfaceId: aws_network_interface.test.id,
    deviceIndex: 0,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ec2.NetworkInterfaceAttachment("test",
    instance_id=aws_instance["test"]["id"],
    network_interface_id=aws_network_interface["test"]["id"],
    device_index=0)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Ec2.NetworkInterfaceAttachment("test", new()
    {
        InstanceId = aws_instance.Test.Id,
        NetworkInterfaceId = aws_network_interface.Test.Id,
        DeviceIndex = 0,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewNetworkInterfaceAttachment(ctx, "test", &ec2.NetworkInterfaceAttachmentArgs{
			InstanceId:         pulumi.Any(aws_instance.Test.Id),
			NetworkInterfaceId: pulumi.Any(aws_network_interface.Test.Id),
			DeviceIndex:        pulumi.Int(0),
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
import com.pulumi.aws.ec2.NetworkInterfaceAttachment;
import com.pulumi.aws.ec2.NetworkInterfaceAttachmentArgs;
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
        var test = new NetworkInterfaceAttachment("test", NetworkInterfaceAttachmentArgs.builder()        
            .instanceId(aws_instance.test().id())
            .networkInterfaceId(aws_network_interface.test().id())
            .deviceIndex(0)
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:ec2:NetworkInterfaceAttachment
    properties:
      instanceId: ${aws_instance.test.id}
      networkInterfaceId: ${aws_network_interface.test.id}
      deviceIndex: 0
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Elastic network interface (ENI) Attachments using its Attachment ID. For example:

```sh
 $ pulumi import aws:ec2/networkInterfaceAttachment:NetworkInterfaceAttachment secondary_nic eni-attach-0a33842b4ec347c4c
```
 