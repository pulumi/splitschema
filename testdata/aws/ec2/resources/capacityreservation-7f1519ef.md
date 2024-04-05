Provides an EC2 Capacity Reservation. This allows you to reserve capacity for your Amazon EC2 instances in a specific Availability Zone for any duration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.ec2.CapacityReservation("default", {
    availabilityZone: "eu-west-1a",
    instanceCount: 1,
    instancePlatform: "Linux/UNIX",
    instanceType: "t2.micro",
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.ec2.CapacityReservation("default",
    availability_zone="eu-west-1a",
    instance_count=1,
    instance_platform="Linux/UNIX",
    instance_type="t2.micro")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Ec2.CapacityReservation("default", new()
    {
        AvailabilityZone = "eu-west-1a",
        InstanceCount = 1,
        InstancePlatform = "Linux/UNIX",
        InstanceType = "t2.micro",
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
		_, err := ec2.NewCapacityReservation(ctx, "default", &ec2.CapacityReservationArgs{
			AvailabilityZone: pulumi.String("eu-west-1a"),
			InstanceCount:    pulumi.Int(1),
			InstancePlatform: pulumi.String("Linux/UNIX"),
			InstanceType:     pulumi.String("t2.micro"),
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
import com.pulumi.aws.ec2.CapacityReservation;
import com.pulumi.aws.ec2.CapacityReservationArgs;
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
        var default_ = new CapacityReservation("default", CapacityReservationArgs.builder()        
            .availabilityZone("eu-west-1a")
            .instanceCount(1)
            .instancePlatform("Linux/UNIX")
            .instanceType("t2.micro")
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:ec2:CapacityReservation
    properties:
      availabilityZone: eu-west-1a
      instanceCount: 1
      instancePlatform: Linux/UNIX
      instanceType: t2.micro
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Capacity Reservations using the `id`. For example:

```sh
 $ pulumi import aws:ec2/capacityReservation:CapacityReservation web cr-0123456789abcdef0
```
 