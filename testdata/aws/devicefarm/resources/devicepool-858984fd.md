Provides a resource to manage AWS Device Farm Device Pools.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.devicefarm.DevicePool("example", {
    projectArn: aws_devicefarm_project.example.arn,
    rules: [{
        attribute: "OS_VERSION",
        operator: "EQUALS",
        value: "\"AVAILABLE\"",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.devicefarm.DevicePool("example",
    project_arn=aws_devicefarm_project["example"]["arn"],
    rules=[aws.devicefarm.DevicePoolRuleArgs(
        attribute="OS_VERSION",
        operator="EQUALS",
        value="\"AVAILABLE\"",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DeviceFarm.DevicePool("example", new()
    {
        ProjectArn = aws_devicefarm_project.Example.Arn,
        Rules = new[]
        {
            new Aws.DeviceFarm.Inputs.DevicePoolRuleArgs
            {
                Attribute = "OS_VERSION",
                Operator = "EQUALS",
                Value = "\"AVAILABLE\"",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/devicefarm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devicefarm.NewDevicePool(ctx, "example", &devicefarm.DevicePoolArgs{
			ProjectArn: pulumi.Any(aws_devicefarm_project.Example.Arn),
			Rules: devicefarm.DevicePoolRuleArray{
				&devicefarm.DevicePoolRuleArgs{
					Attribute: pulumi.String("OS_VERSION"),
					Operator:  pulumi.String("EQUALS"),
					Value:     pulumi.String("\"AVAILABLE\""),
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
import com.pulumi.aws.devicefarm.DevicePool;
import com.pulumi.aws.devicefarm.DevicePoolArgs;
import com.pulumi.aws.devicefarm.inputs.DevicePoolRuleArgs;
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
        var example = new DevicePool("example", DevicePoolArgs.builder()        
            .projectArn(aws_devicefarm_project.example().arn())
            .rules(DevicePoolRuleArgs.builder()
                .attribute("OS_VERSION")
                .operator("EQUALS")
                .value("\"AVAILABLE\"")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:devicefarm:DevicePool
    properties:
      projectArn: ${aws_devicefarm_project.example.arn}
      rules:
        - attribute: OS_VERSION
          operator: EQUALS
          value: '"AVAILABLE"'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DeviceFarm Device Pools using their ARN. For example:

```sh
 $ pulumi import aws:devicefarm/devicePool:DevicePool example arn:aws:devicefarm:us-west-2:123456789012:devicepool:4fa784c7-ccb4-4dbf-ba4f-02198320daa1/4fa784c7-ccb4-4dbf-ba4f-02198320daa1
```
 