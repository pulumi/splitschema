Provides an AppStream fleet.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testFleet = new aws.appstream.Fleet("testFleet", {
    computeCapacity: {
        desiredInstances: 1,
    },
    description: "test fleet",
    displayName: "test-fleet",
    enableDefaultInternetAccess: false,
    fleetType: "ON_DEMAND",
    idleDisconnectTimeoutInSeconds: 60,
    imageName: "Amazon-AppStream2-Sample-Image-03-11-2023",
    instanceType: "stream.standard.large",
    maxUserDurationInSeconds: 600,
    tags: {
        TagName: "tag-value",
    },
    vpcConfig: {
        subnetIds: ["subnet-06e9b13400c225127"],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_fleet = aws.appstream.Fleet("testFleet",
    compute_capacity=aws.appstream.FleetComputeCapacityArgs(
        desired_instances=1,
    ),
    description="test fleet",
    display_name="test-fleet",
    enable_default_internet_access=False,
    fleet_type="ON_DEMAND",
    idle_disconnect_timeout_in_seconds=60,
    image_name="Amazon-AppStream2-Sample-Image-03-11-2023",
    instance_type="stream.standard.large",
    max_user_duration_in_seconds=600,
    tags={
        "TagName": "tag-value",
    },
    vpc_config=aws.appstream.FleetVpcConfigArgs(
        subnet_ids=["subnet-06e9b13400c225127"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testFleet = new Aws.AppStream.Fleet("testFleet", new()
    {
        ComputeCapacity = new Aws.AppStream.Inputs.FleetComputeCapacityArgs
        {
            DesiredInstances = 1,
        },
        Description = "test fleet",
        DisplayName = "test-fleet",
        EnableDefaultInternetAccess = false,
        FleetType = "ON_DEMAND",
        IdleDisconnectTimeoutInSeconds = 60,
        ImageName = "Amazon-AppStream2-Sample-Image-03-11-2023",
        InstanceType = "stream.standard.large",
        MaxUserDurationInSeconds = 600,
        Tags = 
        {
            { "TagName", "tag-value" },
        },
        VpcConfig = new Aws.AppStream.Inputs.FleetVpcConfigArgs
        {
            SubnetIds = new[]
            {
                "subnet-06e9b13400c225127",
            },
        },
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
		_, err := appstream.NewFleet(ctx, "testFleet", &appstream.FleetArgs{
			ComputeCapacity: &appstream.FleetComputeCapacityArgs{
				DesiredInstances: pulumi.Int(1),
			},
			Description:                    pulumi.String("test fleet"),
			DisplayName:                    pulumi.String("test-fleet"),
			EnableDefaultInternetAccess:    pulumi.Bool(false),
			FleetType:                      pulumi.String("ON_DEMAND"),
			IdleDisconnectTimeoutInSeconds: pulumi.Int(60),
			ImageName:                      pulumi.String("Amazon-AppStream2-Sample-Image-03-11-2023"),
			InstanceType:                   pulumi.String("stream.standard.large"),
			MaxUserDurationInSeconds:       pulumi.Int(600),
			Tags: pulumi.StringMap{
				"TagName": pulumi.String("tag-value"),
			},
			VpcConfig: &appstream.FleetVpcConfigArgs{
				SubnetIds: pulumi.StringArray{
					pulumi.String("subnet-06e9b13400c225127"),
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
import com.pulumi.aws.appstream.Fleet;
import com.pulumi.aws.appstream.FleetArgs;
import com.pulumi.aws.appstream.inputs.FleetComputeCapacityArgs;
import com.pulumi.aws.appstream.inputs.FleetVpcConfigArgs;
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
        var testFleet = new Fleet("testFleet", FleetArgs.builder()        
            .computeCapacity(FleetComputeCapacityArgs.builder()
                .desiredInstances(1)
                .build())
            .description("test fleet")
            .displayName("test-fleet")
            .enableDefaultInternetAccess(false)
            .fleetType("ON_DEMAND")
            .idleDisconnectTimeoutInSeconds(60)
            .imageName("Amazon-AppStream2-Sample-Image-03-11-2023")
            .instanceType("stream.standard.large")
            .maxUserDurationInSeconds(600)
            .tags(Map.of("TagName", "tag-value"))
            .vpcConfig(FleetVpcConfigArgs.builder()
                .subnetIds("subnet-06e9b13400c225127")
                .build())
            .build());

    }
}
```
```yaml
resources:
  testFleet:
    type: aws:appstream:Fleet
    properties:
      computeCapacity:
        desiredInstances: 1
      description: test fleet
      displayName: test-fleet
      enableDefaultInternetAccess: false
      fleetType: ON_DEMAND
      idleDisconnectTimeoutInSeconds: 60
      imageName: Amazon-AppStream2-Sample-Image-03-11-2023
      instanceType: stream.standard.large
      maxUserDurationInSeconds: 600
      tags:
        TagName: tag-value
      vpcConfig:
        subnetIds:
          - subnet-06e9b13400c225127
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_appstream_fleet` using the id. For example:

```sh
 $ pulumi import aws:appstream/fleet:Fleet example fleetNameExample
```
 