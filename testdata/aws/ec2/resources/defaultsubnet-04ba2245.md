Provides a resource to manage a [default subnet](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/default-vpc.html#default-vpc-basics) in the current region.

**This is an advanced resource** and has special caveats to be aware of when using it. Please read this document in its entirety before using this resource.

The `aws.ec2.DefaultSubnet` resource behaves differently from normal resources in that if a default subnet exists in the specified Availability Zone, this provider does not _create_ this resource, but instead "adopts" it into management.
If no default subnet exists, this provider creates a new default subnet.
By default, `pulumi destroy` does not delete the default subnet but does remove the resource from the state.
Set the `force_destroy` argument to `true` to delete the default subnet.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultAz1 = new aws.ec2.DefaultSubnet("defaultAz1", {
    availabilityZone: "us-west-2a",
    tags: {
        Name: "Default subnet for us-west-2a",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

default_az1 = aws.ec2.DefaultSubnet("defaultAz1",
    availability_zone="us-west-2a",
    tags={
        "Name": "Default subnet for us-west-2a",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultAz1 = new Aws.Ec2.DefaultSubnet("defaultAz1", new()
    {
        AvailabilityZone = "us-west-2a",
        Tags = 
        {
            { "Name", "Default subnet for us-west-2a" },
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
		_, err := ec2.NewDefaultSubnet(ctx, "defaultAz1", &ec2.DefaultSubnetArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Default subnet for us-west-2a"),
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
import com.pulumi.aws.ec2.DefaultSubnet;
import com.pulumi.aws.ec2.DefaultSubnetArgs;
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
        var defaultAz1 = new DefaultSubnet("defaultAz1", DefaultSubnetArgs.builder()        
            .availabilityZone("us-west-2a")
            .tags(Map.of("Name", "Default subnet for us-west-2a"))
            .build());

    }
}
```
```yaml
resources:
  defaultAz1:
    type: aws:ec2:DefaultSubnet
    properties:
      availabilityZone: us-west-2a
      tags:
        Name: Default subnet for us-west-2a
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import subnets using the subnet `id`. For example:

```sh
 $ pulumi import aws:ec2/defaultSubnet:DefaultSubnet public_subnet subnet-9d4a7b6c
```
 