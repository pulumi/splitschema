Provides a GameLift Fleet resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.gamelift.Fleet("example", {
    buildId: aws_gamelift_build.example.id,
    ec2InstanceType: "t2.micro",
    fleetType: "ON_DEMAND",
    runtimeConfiguration: {
        serverProcesses: [{
            concurrentExecutions: 1,
            launchPath: "C:\\game\\GomokuServer.exe",
        }],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.gamelift.Fleet("example",
    build_id=aws_gamelift_build["example"]["id"],
    ec2_instance_type="t2.micro",
    fleet_type="ON_DEMAND",
    runtime_configuration=aws.gamelift.FleetRuntimeConfigurationArgs(
        server_processes=[aws.gamelift.FleetRuntimeConfigurationServerProcessArgs(
            concurrent_executions=1,
            launch_path="C:\\game\\GomokuServer.exe",
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.GameLift.Fleet("example", new()
    {
        BuildId = aws_gamelift_build.Example.Id,
        Ec2InstanceType = "t2.micro",
        FleetType = "ON_DEMAND",
        RuntimeConfiguration = new Aws.GameLift.Inputs.FleetRuntimeConfigurationArgs
        {
            ServerProcesses = new[]
            {
                new Aws.GameLift.Inputs.FleetRuntimeConfigurationServerProcessArgs
                {
                    ConcurrentExecutions = 1,
                    LaunchPath = "C:\\game\\GomokuServer.exe",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/gamelift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := gamelift.NewFleet(ctx, "example", &gamelift.FleetArgs{
			BuildId:         pulumi.Any(aws_gamelift_build.Example.Id),
			Ec2InstanceType: pulumi.String("t2.micro"),
			FleetType:       pulumi.String("ON_DEMAND"),
			RuntimeConfiguration: &gamelift.FleetRuntimeConfigurationArgs{
				ServerProcesses: gamelift.FleetRuntimeConfigurationServerProcessArray{
					&gamelift.FleetRuntimeConfigurationServerProcessArgs{
						ConcurrentExecutions: pulumi.Int(1),
						LaunchPath:           pulumi.String("C:\\game\\GomokuServer.exe"),
					},
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
import com.pulumi.aws.gamelift.Fleet;
import com.pulumi.aws.gamelift.FleetArgs;
import com.pulumi.aws.gamelift.inputs.FleetRuntimeConfigurationArgs;
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
        var example = new Fleet("example", FleetArgs.builder()        
            .buildId(aws_gamelift_build.example().id())
            .ec2InstanceType("t2.micro")
            .fleetType("ON_DEMAND")
            .runtimeConfiguration(FleetRuntimeConfigurationArgs.builder()
                .serverProcesses(FleetRuntimeConfigurationServerProcessArgs.builder()
                    .concurrentExecutions(1)
                    .launchPath("C:\\game\\GomokuServer.exe")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:gamelift:Fleet
    properties:
      buildId: ${aws_gamelift_build.example.id}
      ec2InstanceType: t2.micro
      fleetType: ON_DEMAND
      runtimeConfiguration:
        serverProcesses:
          - concurrentExecutions: 1
            launchPath: C:\game\GomokuServer.exe
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import GameLift Fleets using the ID. For example:

```sh
 $ pulumi import aws:gamelift/fleet:Fleet example <fleet-id>
```
 