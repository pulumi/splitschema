Provides a WAF Regional IPSet Resource for use with Application Load Balancer.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ipset = new aws.wafregional.IpSet("ipset", {ipSetDescriptors: [
    {
        type: "IPV4",
        value: "192.0.7.0/24",
    },
    {
        type: "IPV4",
        value: "10.16.16.0/16",
    },
]});
```
```python
import pulumi
import pulumi_aws as aws

ipset = aws.wafregional.IpSet("ipset", ip_set_descriptors=[
    aws.wafregional.IpSetIpSetDescriptorArgs(
        type="IPV4",
        value="192.0.7.0/24",
    ),
    aws.wafregional.IpSetIpSetDescriptorArgs(
        type="IPV4",
        value="10.16.16.0/16",
    ),
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ipset = new Aws.WafRegional.IpSet("ipset", new()
    {
        IpSetDescriptors = new[]
        {
            new Aws.WafRegional.Inputs.IpSetIpSetDescriptorArgs
            {
                Type = "IPV4",
                Value = "192.0.7.0/24",
            },
            new Aws.WafRegional.Inputs.IpSetIpSetDescriptorArgs
            {
                Type = "IPV4",
                Value = "10.16.16.0/16",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafregional.NewIpSet(ctx, "ipset", &wafregional.IpSetArgs{
			IpSetDescriptors: wafregional.IpSetIpSetDescriptorArray{
				&wafregional.IpSetIpSetDescriptorArgs{
					Type:  pulumi.String("IPV4"),
					Value: pulumi.String("192.0.7.0/24"),
				},
				&wafregional.IpSetIpSetDescriptorArgs{
					Type:  pulumi.String("IPV4"),
					Value: pulumi.String("10.16.16.0/16"),
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
import com.pulumi.aws.wafregional.IpSet;
import com.pulumi.aws.wafregional.IpSetArgs;
import com.pulumi.aws.wafregional.inputs.IpSetIpSetDescriptorArgs;
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
        var ipset = new IpSet("ipset", IpSetArgs.builder()        
            .ipSetDescriptors(            
                IpSetIpSetDescriptorArgs.builder()
                    .type("IPV4")
                    .value("192.0.7.0/24")
                    .build(),
                IpSetIpSetDescriptorArgs.builder()
                    .type("IPV4")
                    .value("10.16.16.0/16")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  ipset:
    type: aws:wafregional:IpSet
    properties:
      ipSetDescriptors:
        - type: IPV4
          value: 192.0.7.0/24
        - type: IPV4
          value: 10.16.16.0/16
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAF Regional IPSets using their ID. For example:

```sh
 $ pulumi import aws:wafregional/ipSet:IpSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 