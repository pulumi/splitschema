Provides a static IP address attachment - relationship between a Lightsail static IP & Lightsail instance.

> **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testStaticIp = new aws.lightsail.StaticIp("testStaticIp", {});
const testInstance = new aws.lightsail.Instance("testInstance", {
    availabilityZone: "us-east-1b",
    blueprintId: "string",
    bundleId: "string",
    keyPairName: "some_key_name",
});
const testStaticIpAttachment = new aws.lightsail.StaticIpAttachment("testStaticIpAttachment", {
    staticIpName: testStaticIp.id,
    instanceName: testInstance.id,
});
```
```python
import pulumi
import pulumi_aws as aws

test_static_ip = aws.lightsail.StaticIp("testStaticIp")
test_instance = aws.lightsail.Instance("testInstance",
    availability_zone="us-east-1b",
    blueprint_id="string",
    bundle_id="string",
    key_pair_name="some_key_name")
test_static_ip_attachment = aws.lightsail.StaticIpAttachment("testStaticIpAttachment",
    static_ip_name=test_static_ip.id,
    instance_name=test_instance.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testStaticIp = new Aws.LightSail.StaticIp("testStaticIp");

    var testInstance = new Aws.LightSail.Instance("testInstance", new()
    {
        AvailabilityZone = "us-east-1b",
        BlueprintId = "string",
        BundleId = "string",
        KeyPairName = "some_key_name",
    });

    var testStaticIpAttachment = new Aws.LightSail.StaticIpAttachment("testStaticIpAttachment", new()
    {
        StaticIpName = testStaticIp.Id,
        InstanceName = testInstance.Id,
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
		testStaticIp, err := lightsail.NewStaticIp(ctx, "testStaticIp", nil)
		if err != nil {
			return err
		}
		testInstance, err := lightsail.NewInstance(ctx, "testInstance", &lightsail.InstanceArgs{
			AvailabilityZone: pulumi.String("us-east-1b"),
			BlueprintId:      pulumi.String("string"),
			BundleId:         pulumi.String("string"),
			KeyPairName:      pulumi.String("some_key_name"),
		})
		if err != nil {
			return err
		}
		_, err = lightsail.NewStaticIpAttachment(ctx, "testStaticIpAttachment", &lightsail.StaticIpAttachmentArgs{
			StaticIpName: testStaticIp.ID(),
			InstanceName: testInstance.ID(),
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
import com.pulumi.aws.lightsail.StaticIp;
import com.pulumi.aws.lightsail.Instance;
import com.pulumi.aws.lightsail.InstanceArgs;
import com.pulumi.aws.lightsail.StaticIpAttachment;
import com.pulumi.aws.lightsail.StaticIpAttachmentArgs;
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
        var testStaticIp = new StaticIp("testStaticIp");

        var testInstance = new Instance("testInstance", InstanceArgs.builder()        
            .availabilityZone("us-east-1b")
            .blueprintId("string")
            .bundleId("string")
            .keyPairName("some_key_name")
            .build());

        var testStaticIpAttachment = new StaticIpAttachment("testStaticIpAttachment", StaticIpAttachmentArgs.builder()        
            .staticIpName(testStaticIp.id())
            .instanceName(testInstance.id())
            .build());

    }
}
```
```yaml
resources:
  testStaticIpAttachment:
    type: aws:lightsail:StaticIpAttachment
    properties:
      staticIpName: ${testStaticIp.id}
      instanceName: ${testInstance.id}
  testStaticIp:
    type: aws:lightsail:StaticIp
  testInstance:
    type: aws:lightsail:Instance
    properties:
      availabilityZone: us-east-1b
      blueprintId: string
      bundleId: string
      keyPairName: some_key_name
```
{{% /example %}}
{{% /examples %}}