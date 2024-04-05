Provides an EC2 Host resource. This allows Dedicated Hosts to be allocated, modified, and released.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new host with instance type of c5.18xlarge with Auto Placement
// and Host Recovery enabled.
const test = new aws.ec2.DedicatedHost("test", {
    autoPlacement: "on",
    availabilityZone: "us-west-2a",
    hostRecovery: "on",
    instanceType: "c5.18xlarge",
});
```
```python
import pulumi
import pulumi_aws as aws

# Create a new host with instance type of c5.18xlarge with Auto Placement
# and Host Recovery enabled.
test = aws.ec2.DedicatedHost("test",
    auto_placement="on",
    availability_zone="us-west-2a",
    host_recovery="on",
    instance_type="c5.18xlarge")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Create a new host with instance type of c5.18xlarge with Auto Placement
    // and Host Recovery enabled.
    var test = new Aws.Ec2.DedicatedHost("test", new()
    {
        AutoPlacement = "on",
        AvailabilityZone = "us-west-2a",
        HostRecovery = "on",
        InstanceType = "c5.18xlarge",
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
		_, err := ec2.NewDedicatedHost(ctx, "test", &ec2.DedicatedHostArgs{
			AutoPlacement:    pulumi.String("on"),
			AvailabilityZone: pulumi.String("us-west-2a"),
			HostRecovery:     pulumi.String("on"),
			InstanceType:     pulumi.String("c5.18xlarge"),
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
import com.pulumi.aws.ec2.DedicatedHost;
import com.pulumi.aws.ec2.DedicatedHostArgs;
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
        var test = new DedicatedHost("test", DedicatedHostArgs.builder()        
            .autoPlacement("on")
            .availabilityZone("us-west-2a")
            .hostRecovery("on")
            .instanceType("c5.18xlarge")
            .build());

    }
}
```
```yaml
resources:
  # Create a new host with instance type of c5.18xlarge with Auto Placement
  # // and Host Recovery enabled.
  test:
    type: aws:ec2:DedicatedHost
    properties:
      autoPlacement: on
      availabilityZone: us-west-2a
      hostRecovery: on
      instanceType: c5.18xlarge
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import hosts using the host `id`. For example:

```sh
 $ pulumi import aws:ec2/dedicatedHost:DedicatedHost example h-0385a99d0e4b20cbb
```
 