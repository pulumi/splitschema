Provides a resource to manage the [default AWS DHCP Options Set](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_DHCP_Options.html#AmazonDNS)
in the current region.

Each AWS region comes with a default set of DHCP options.
**This is an advanced resource**, and has special caveats to be aware of when
using it. Please read this document in its entirety before using this resource.

The `aws.ec2.DefaultVpcDhcpOptions` behaves differently from normal resources, in that
this provider does not _create_ this resource, but instead "adopts" it
into management.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage with tags:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.ec2.DefaultVpcDhcpOptions("default", {tags: {
    Name: "Default DHCP Option Set",
}});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.ec2.DefaultVpcDhcpOptions("default", tags={
    "Name": "Default DHCP Option Set",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Ec2.DefaultVpcDhcpOptions("default", new()
    {
        Tags = 
        {
            { "Name", "Default DHCP Option Set" },
        },
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
		_, err := ec2.NewDefaultVpcDhcpOptions(ctx, "default", &ec2.DefaultVpcDhcpOptionsArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Default DHCP Option Set"),
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
import com.pulumi.aws.ec2.DefaultVpcDhcpOptions;
import com.pulumi.aws.ec2.DefaultVpcDhcpOptionsArgs;
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
        var default_ = new DefaultVpcDhcpOptions("default", DefaultVpcDhcpOptionsArgs.builder()        
            .tags(Map.of("Name", "Default DHCP Option Set"))
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:ec2:DefaultVpcDhcpOptions
    properties:
      tags:
        Name: Default DHCP Option Set
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC DHCP Options using the DHCP Options `id`. For example:

```sh
 $ pulumi import aws:ec2/defaultVpcDhcpOptions:DefaultVpcDhcpOptions default_options dopt-d9070ebb
```
 