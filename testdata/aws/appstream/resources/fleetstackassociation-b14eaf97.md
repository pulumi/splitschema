Manages an AppStream Fleet Stack association.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleFleet = new aws.appstream.Fleet("exampleFleet", {
    imageName: "Amazon-AppStream2-Sample-Image-03-11-2023",
    instanceType: "stream.standard.small",
    computeCapacity: {
        desiredInstances: 1,
    },
});
const exampleStack = new aws.appstream.Stack("exampleStack", {});
const exampleFleetStackAssociation = new aws.appstream.FleetStackAssociation("exampleFleetStackAssociation", {
    fleetName: exampleFleet.name,
    stackName: exampleStack.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example_fleet = aws.appstream.Fleet("exampleFleet",
    image_name="Amazon-AppStream2-Sample-Image-03-11-2023",
    instance_type="stream.standard.small",
    compute_capacity=aws.appstream.FleetComputeCapacityArgs(
        desired_instances=1,
    ))
example_stack = aws.appstream.Stack("exampleStack")
example_fleet_stack_association = aws.appstream.FleetStackAssociation("exampleFleetStackAssociation",
    fleet_name=example_fleet.name,
    stack_name=example_stack.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleFleet = new Aws.AppStream.Fleet("exampleFleet", new()
    {
        ImageName = "Amazon-AppStream2-Sample-Image-03-11-2023",
        InstanceType = "stream.standard.small",
        ComputeCapacity = new Aws.AppStream.Inputs.FleetComputeCapacityArgs
        {
            DesiredInstances = 1,
        },
    });

    var exampleStack = new Aws.AppStream.Stack("exampleStack");

    var exampleFleetStackAssociation = new Aws.AppStream.FleetStackAssociation("exampleFleetStackAssociation", new()
    {
        FleetName = exampleFleet.Name,
        StackName = exampleStack.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appstream"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleFleet, err := appstream.NewFleet(ctx, "exampleFleet", &appstream.FleetArgs{
			ImageName:    pulumi.String("Amazon-AppStream2-Sample-Image-03-11-2023"),
			InstanceType: pulumi.String("stream.standard.small"),
			ComputeCapacity: &appstream.FleetComputeCapacityArgs{
				DesiredInstances: pulumi.Int(1),
			},
		})
		if err != nil {
			return err
		}
		exampleStack, err := appstream.NewStack(ctx, "exampleStack", nil)
		if err != nil {
			return err
		}
		_, err = appstream.NewFleetStackAssociation(ctx, "exampleFleetStackAssociation", &appstream.FleetStackAssociationArgs{
			FleetName: exampleFleet.Name,
			StackName: exampleStack.Name,
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
import com.pulumi.aws.appstream.Fleet;
import com.pulumi.aws.appstream.FleetArgs;
import com.pulumi.aws.appstream.inputs.FleetComputeCapacityArgs;
import com.pulumi.aws.appstream.Stack;
import com.pulumi.aws.appstream.FleetStackAssociation;
import com.pulumi.aws.appstream.FleetStackAssociationArgs;
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
        var exampleFleet = new Fleet("exampleFleet", FleetArgs.builder()        
            .imageName("Amazon-AppStream2-Sample-Image-03-11-2023")
            .instanceType("stream.standard.small")
            .computeCapacity(FleetComputeCapacityArgs.builder()
                .desiredInstances(1)
                .build())
            .build());

        var exampleStack = new Stack("exampleStack");

        var exampleFleetStackAssociation = new FleetStackAssociation("exampleFleetStackAssociation", FleetStackAssociationArgs.builder()        
            .fleetName(exampleFleet.name())
            .stackName(exampleStack.name())
            .build());

    }
}
```
```yaml
resources:
  exampleFleet:
    type: aws:appstream:Fleet
    properties:
      imageName: Amazon-AppStream2-Sample-Image-03-11-2023
      instanceType: stream.standard.small
      computeCapacity:
        desiredInstances: 1
  exampleStack:
    type: aws:appstream:Stack
  exampleFleetStackAssociation:
    type: aws:appstream:FleetStackAssociation
    properties:
      fleetName: ${exampleFleet.name}
      stackName: ${exampleStack.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AppStream Stack Fleet Association using the `fleet_name` and `stack_name` separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:appstream/fleetStackAssociation:FleetStackAssociation example fleetName/stackName
```
 