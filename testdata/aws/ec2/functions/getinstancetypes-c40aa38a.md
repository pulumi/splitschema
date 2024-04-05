Information about EC2 Instance Types.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.ec2.getInstanceTypes({
    filters: [
        {
            name: "auto-recovery-supported",
            values: ["true"],
        },
        {
            name: "network-info.encryption-in-transit-supported",
            values: ["true"],
        },
        {
            name: "instance-storage-supported",
            values: ["true"],
        },
        {
            name: "instance-type",
            values: [
                "g5.2xlarge",
                "g5.4xlarge",
            ],
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ec2.get_instance_types(filters=[
    aws.ec2.GetInstanceTypesFilterArgs(
        name="auto-recovery-supported",
        values=["true"],
    ),
    aws.ec2.GetInstanceTypesFilterArgs(
        name="network-info.encryption-in-transit-supported",
        values=["true"],
    ),
    aws.ec2.GetInstanceTypesFilterArgs(
        name="instance-storage-supported",
        values=["true"],
    ),
    aws.ec2.GetInstanceTypesFilterArgs(
        name="instance-type",
        values=[
            "g5.2xlarge",
            "g5.4xlarge",
        ],
    ),
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Ec2.GetInstanceTypes.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetInstanceTypesFilterInputArgs
            {
                Name = "auto-recovery-supported",
                Values = new[]
                {
                    "true",
                },
            },
            new Aws.Ec2.Inputs.GetInstanceTypesFilterInputArgs
            {
                Name = "network-info.encryption-in-transit-supported",
                Values = new[]
                {
                    "true",
                },
            },
            new Aws.Ec2.Inputs.GetInstanceTypesFilterInputArgs
            {
                Name = "instance-storage-supported",
                Values = new[]
                {
                    "true",
                },
            },
            new Aws.Ec2.Inputs.GetInstanceTypesFilterInputArgs
            {
                Name = "instance-type",
                Values = new[]
                {
                    "g5.2xlarge",
                    "g5.4xlarge",
                },
            },
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
		_, err := ec2.GetInstanceTypes(ctx, &ec2.GetInstanceTypesArgs{
			Filters: []ec2.GetInstanceTypesFilter{
				{
					Name: "auto-recovery-supported",
					Values: []string{
						"true",
					},
				},
				{
					Name: "network-info.encryption-in-transit-supported",
					Values: []string{
						"true",
					},
				},
				{
					Name: "instance-storage-supported",
					Values: []string{
						"true",
					},
				},
				{
					Name: "instance-type",
					Values: []string{
						"g5.2xlarge",
						"g5.4xlarge",
					},
				},
			},
		}, nil)
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetInstanceTypesArgs;
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
        final var test = Ec2Functions.getInstanceTypes(GetInstanceTypesArgs.builder()
            .filters(            
                GetInstanceTypesFilterArgs.builder()
                    .name("auto-recovery-supported")
                    .values("true")
                    .build(),
                GetInstanceTypesFilterArgs.builder()
                    .name("network-info.encryption-in-transit-supported")
                    .values("true")
                    .build(),
                GetInstanceTypesFilterArgs.builder()
                    .name("instance-storage-supported")
                    .values("true")
                    .build(),
                GetInstanceTypesFilterArgs.builder()
                    .name("instance-type")
                    .values(                    
                        "g5.2xlarge",
                        "g5.4xlarge")
                    .build())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:ec2:getInstanceTypes
      Arguments:
        filters:
          - name: auto-recovery-supported
            values:
              - 'true'
          - name: network-info.encryption-in-transit-supported
            values:
              - 'true'
          - name: instance-storage-supported
            values:
              - 'true'
          - name: instance-type
            values:
              - g5.2xlarge
              - g5.4xlarge
```
{{% /example %}}
{{% /examples %}}