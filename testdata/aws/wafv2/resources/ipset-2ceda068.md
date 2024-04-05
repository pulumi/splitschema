Provides a WAFv2 IP Set Resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.wafv2.IpSet("example", {
    addresses: [
        "1.2.3.4/32",
        "5.6.7.8/32",
    ],
    description: "Example IP set",
    ipAddressVersion: "IPV4",
    scope: "REGIONAL",
    tags: {
        Tag1: "Value1",
        Tag2: "Value2",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.wafv2.IpSet("example",
    addresses=[
        "1.2.3.4/32",
        "5.6.7.8/32",
    ],
    description="Example IP set",
    ip_address_version="IPV4",
    scope="REGIONAL",
    tags={
        "Tag1": "Value1",
        "Tag2": "Value2",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.WafV2.IpSet("example", new()
    {
        Addresses = new[]
        {
            "1.2.3.4/32",
            "5.6.7.8/32",
        },
        Description = "Example IP set",
        IpAddressVersion = "IPV4",
        Scope = "REGIONAL",
        Tags = 
        {
            { "Tag1", "Value1" },
            { "Tag2", "Value2" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafv2.NewIpSet(ctx, "example", &wafv2.IpSetArgs{
			Addresses: pulumi.StringArray{
				pulumi.String("1.2.3.4/32"),
				pulumi.String("5.6.7.8/32"),
			},
			Description:      pulumi.String("Example IP set"),
			IpAddressVersion: pulumi.String("IPV4"),
			Scope:            pulumi.String("REGIONAL"),
			Tags: pulumi.StringMap{
				"Tag1": pulumi.String("Value1"),
				"Tag2": pulumi.String("Value2"),
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
import com.pulumi.aws.wafv2.IpSet;
import com.pulumi.aws.wafv2.IpSetArgs;
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
        var example = new IpSet("example", IpSetArgs.builder()        
            .addresses(            
                "1.2.3.4/32",
                "5.6.7.8/32")
            .description("Example IP set")
            .ipAddressVersion("IPV4")
            .scope("REGIONAL")
            .tags(Map.ofEntries(
                Map.entry("Tag1", "Value1"),
                Map.entry("Tag2", "Value2")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:wafv2:IpSet
    properties:
      addresses:
        - 1.2.3.4/32
        - 5.6.7.8/32
      description: Example IP set
      ipAddressVersion: IPV4
      scope: REGIONAL
      tags:
        Tag1: Value1
        Tag2: Value2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAFv2 IP Sets using `ID/name/scope`. For example:

```sh
 $ pulumi import aws:wafv2/ipSet:IpSet example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc/example/REGIONAL
```
 