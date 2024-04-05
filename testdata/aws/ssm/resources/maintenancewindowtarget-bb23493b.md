Provides an SSM Maintenance Window Target resource

{{% examples %}}
## Example Usage
{{% example %}}
### Instance Target

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const window = new aws.ssm.MaintenanceWindow("window", {
    schedule: "cron(0 16 ? * TUE *)",
    duration: 3,
    cutoff: 1,
});
const target1 = new aws.ssm.MaintenanceWindowTarget("target1", {
    windowId: window.id,
    description: "This is a maintenance window target",
    resourceType: "INSTANCE",
    targets: [{
        key: "tag:Name",
        values: ["acceptance_test"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

window = aws.ssm.MaintenanceWindow("window",
    schedule="cron(0 16 ? * TUE *)",
    duration=3,
    cutoff=1)
target1 = aws.ssm.MaintenanceWindowTarget("target1",
    window_id=window.id,
    description="This is a maintenance window target",
    resource_type="INSTANCE",
    targets=[aws.ssm.MaintenanceWindowTargetTargetArgs(
        key="tag:Name",
        values=["acceptance_test"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var window = new Aws.Ssm.MaintenanceWindow("window", new()
    {
        Schedule = "cron(0 16 ? * TUE *)",
        Duration = 3,
        Cutoff = 1,
    });

    var target1 = new Aws.Ssm.MaintenanceWindowTarget("target1", new()
    {
        WindowId = window.Id,
        Description = "This is a maintenance window target",
        ResourceType = "INSTANCE",
        Targets = new[]
        {
            new Aws.Ssm.Inputs.MaintenanceWindowTargetTargetArgs
            {
                Key = "tag:Name",
                Values = new[]
                {
                    "acceptance_test",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		window, err := ssm.NewMaintenanceWindow(ctx, "window", &ssm.MaintenanceWindowArgs{
			Schedule: pulumi.String("cron(0 16 ? * TUE *)"),
			Duration: pulumi.Int(3),
			Cutoff:   pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		_, err = ssm.NewMaintenanceWindowTarget(ctx, "target1", &ssm.MaintenanceWindowTargetArgs{
			WindowId:     window.ID(),
			Description:  pulumi.String("This is a maintenance window target"),
			ResourceType: pulumi.String("INSTANCE"),
			Targets: ssm.MaintenanceWindowTargetTargetArray{
				&ssm.MaintenanceWindowTargetTargetArgs{
					Key: pulumi.String("tag:Name"),
					Values: pulumi.StringArray{
						pulumi.String("acceptance_test"),
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
import com.pulumi.aws.ssm.MaintenanceWindow;
import com.pulumi.aws.ssm.MaintenanceWindowArgs;
import com.pulumi.aws.ssm.MaintenanceWindowTarget;
import com.pulumi.aws.ssm.MaintenanceWindowTargetArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTargetTargetArgs;
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
        var window = new MaintenanceWindow("window", MaintenanceWindowArgs.builder()        
            .schedule("cron(0 16 ? * TUE *)")
            .duration(3)
            .cutoff(1)
            .build());

        var target1 = new MaintenanceWindowTarget("target1", MaintenanceWindowTargetArgs.builder()        
            .windowId(window.id())
            .description("This is a maintenance window target")
            .resourceType("INSTANCE")
            .targets(MaintenanceWindowTargetTargetArgs.builder()
                .key("tag:Name")
                .values("acceptance_test")
                .build())
            .build());

    }
}
```
```yaml
resources:
  window:
    type: aws:ssm:MaintenanceWindow
    properties:
      schedule: cron(0 16 ? * TUE *)
      duration: 3
      cutoff: 1
  target1:
    type: aws:ssm:MaintenanceWindowTarget
    properties:
      windowId: ${window.id}
      description: This is a maintenance window target
      resourceType: INSTANCE
      targets:
        - key: tag:Name
          values:
            - acceptance_test
```
{{% /example %}}
{{% example %}}
### Resource Group Target

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const window = new aws.ssm.MaintenanceWindow("window", {
    schedule: "cron(0 16 ? * TUE *)",
    duration: 3,
    cutoff: 1,
});
const target1 = new aws.ssm.MaintenanceWindowTarget("target1", {
    windowId: window.id,
    description: "This is a maintenance window target",
    resourceType: "RESOURCE_GROUP",
    targets: [{
        key: "resource-groups:ResourceTypeFilters",
        values: ["AWS::EC2::Instance"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

window = aws.ssm.MaintenanceWindow("window",
    schedule="cron(0 16 ? * TUE *)",
    duration=3,
    cutoff=1)
target1 = aws.ssm.MaintenanceWindowTarget("target1",
    window_id=window.id,
    description="This is a maintenance window target",
    resource_type="RESOURCE_GROUP",
    targets=[aws.ssm.MaintenanceWindowTargetTargetArgs(
        key="resource-groups:ResourceTypeFilters",
        values=["AWS::EC2::Instance"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var window = new Aws.Ssm.MaintenanceWindow("window", new()
    {
        Schedule = "cron(0 16 ? * TUE *)",
        Duration = 3,
        Cutoff = 1,
    });

    var target1 = new Aws.Ssm.MaintenanceWindowTarget("target1", new()
    {
        WindowId = window.Id,
        Description = "This is a maintenance window target",
        ResourceType = "RESOURCE_GROUP",
        Targets = new[]
        {
            new Aws.Ssm.Inputs.MaintenanceWindowTargetTargetArgs
            {
                Key = "resource-groups:ResourceTypeFilters",
                Values = new[]
                {
                    "AWS::EC2::Instance",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		window, err := ssm.NewMaintenanceWindow(ctx, "window", &ssm.MaintenanceWindowArgs{
			Schedule: pulumi.String("cron(0 16 ? * TUE *)"),
			Duration: pulumi.Int(3),
			Cutoff:   pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		_, err = ssm.NewMaintenanceWindowTarget(ctx, "target1", &ssm.MaintenanceWindowTargetArgs{
			WindowId:     window.ID(),
			Description:  pulumi.String("This is a maintenance window target"),
			ResourceType: pulumi.String("RESOURCE_GROUP"),
			Targets: ssm.MaintenanceWindowTargetTargetArray{
				&ssm.MaintenanceWindowTargetTargetArgs{
					Key: pulumi.String("resource-groups:ResourceTypeFilters"),
					Values: pulumi.StringArray{
						pulumi.String("AWS::EC2::Instance"),
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
import com.pulumi.aws.ssm.MaintenanceWindow;
import com.pulumi.aws.ssm.MaintenanceWindowArgs;
import com.pulumi.aws.ssm.MaintenanceWindowTarget;
import com.pulumi.aws.ssm.MaintenanceWindowTargetArgs;
import com.pulumi.aws.ssm.inputs.MaintenanceWindowTargetTargetArgs;
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
        var window = new MaintenanceWindow("window", MaintenanceWindowArgs.builder()        
            .schedule("cron(0 16 ? * TUE *)")
            .duration(3)
            .cutoff(1)
            .build());

        var target1 = new MaintenanceWindowTarget("target1", MaintenanceWindowTargetArgs.builder()        
            .windowId(window.id())
            .description("This is a maintenance window target")
            .resourceType("RESOURCE_GROUP")
            .targets(MaintenanceWindowTargetTargetArgs.builder()
                .key("resource-groups:ResourceTypeFilters")
                .values("AWS::EC2::Instance")
                .build())
            .build());

    }
}
```
```yaml
resources:
  window:
    type: aws:ssm:MaintenanceWindow
    properties:
      schedule: cron(0 16 ? * TUE *)
      duration: 3
      cutoff: 1
  target1:
    type: aws:ssm:MaintenanceWindowTarget
    properties:
      windowId: ${window.id}
      description: This is a maintenance window target
      resourceType: RESOURCE_GROUP
      targets:
        - key: resource-groups:ResourceTypeFilters
          values:
            - AWS::EC2::Instance
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SSM Maintenance Window targets using `WINDOW_ID/WINDOW_TARGET_ID`. For example:

```sh
 $ pulumi import aws:ssm/maintenanceWindowTarget:MaintenanceWindowTarget example mw-0c50858d01EXAMPLE/23639a0b-ddbc-4bca-9e72-78d96EXAMPLE
```
 