Use this data source to get IDs or IPs of Amazon EC2 instances to be referenced elsewhere,
e.g., to allow easier migration from another management solution
or to make it easier for an operator to connect through bastion host(s).





> **Note:** It's strongly discouraged to use this data source for querying ephemeral
instances (e.g., managed via autoscaling group), as the output may change at any time
and you'd need to re-run `apply` every time an instance comes up or dies.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

export = async () => {
    const testInstances = await aws.ec2.getInstances({
        instanceTags: {
            Role: "HardWorker",
        },
        filters: [{
            name: "instance.group-id",
            values: ["sg-12345678"],
        }],
        instanceStateNames: [
            "running",
            "stopped",
        ],
    });
    const testEip: aws.ec2.Eip[] = [];
    for (const range = {value: 0}; range.value < testInstances.ids.length; range.value++) {
        testEip.push(new aws.ec2.Eip(`testEip-${range.value}`, {instance: testInstances.ids[range.value]}));
    }
}
```
```python
import pulumi
import pulumi_aws as aws

test_instances = aws.ec2.get_instances(instance_tags={
        "Role": "HardWorker",
    },
    filters=[aws.ec2.GetInstancesFilterArgs(
        name="instance.group-id",
        values=["sg-12345678"],
    )],
    instance_state_names=[
        "running",
        "stopped",
    ])
test_eip = []
for range in [{"value": i} for i in range(0, len(test_instances.ids))]:
    test_eip.append(aws.ec2.Eip(f"testEip-{range['value']}", instance=test_instances.ids[range["value"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(async() => 
{
    var testInstances = await Aws.Ec2.GetInstances.InvokeAsync(new()
    {
        InstanceTags = 
        {
            { "Role", "HardWorker" },
        },
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetInstancesFilterInputArgs
            {
                Name = "instance.group-id",
                Values = new[]
                {
                    "sg-12345678",
                },
            },
        },
        InstanceStateNames = new[]
        {
            "running",
            "stopped",
        },
    });

    var testEip = new List<Aws.Ec2.Eip>();
    for (var rangeIndex = 0; rangeIndex < testInstances.Ids.Length; rangeIndex++)
    {
        var range = new { Value = rangeIndex };
        testEip.Add(new Aws.Ec2.Eip($"testEip-{range.Value}", new()
        {
            Instance = testInstances.Ids[range.Value],
        }));
    }
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
		testInstances, err := ec2.GetInstances(ctx, &ec2.GetInstancesArgs{
			InstanceTags: map[string]interface{}{
				"Role": "HardWorker",
			},
			Filters: []ec2.GetInstancesFilter{
				{
					Name: "instance.group-id",
					Values: []string{
						"sg-12345678",
					},
				},
			},
			InstanceStateNames: []string{
				"running",
				"stopped",
			},
		}, nil)
		if err != nil {
			return err
		}
		var testEip []*ec2.Eip
		for index := 0; index < len(testInstances.Ids); index++ {
			key0 := index
			val0 := index
			__res, err := ec2.NewEip(ctx, fmt.Sprintf("testEip-%v", key0), &ec2.EipArgs{
				Instance: testInstances.Ids[val0],
			})
			if err != nil {
				return err
			}
			testEip = append(testEip, __res)
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetInstancesArgs;
import com.pulumi.aws.ec2.Eip;
import com.pulumi.aws.ec2.EipArgs;
import com.pulumi.codegen.internal.KeyedValue;
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
        final var testInstances = Ec2Functions.getInstances(GetInstancesArgs.builder()
            .instanceTags(Map.of("Role", "HardWorker"))
            .filters(GetInstancesFilterArgs.builder()
                .name("instance.group-id")
                .values("sg-12345678")
                .build())
            .instanceStateNames(            
                "running",
                "stopped")
            .build());

        for (var i = 0; i < testInstances.applyValue(getInstancesResult -> getInstancesResult.ids()).length(); i++) {
            new Eip("testEip-" + i, EipArgs.builder()            
                .instance(testInstances.applyValue(getInstancesResult -> getInstancesResult.ids())[range.value()])
                .build());

        
}
    }
}
```
{{% /example %}}
{{% /examples %}}