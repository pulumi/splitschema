Information about single EC2 Instance Type Offering.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getInstanceTypeOffering({
    filters: [{
        name: "instance-type",
        values: [
            "t2.micro",
            "t3.micro",
        ],
    }],
    preferredInstanceTypes: [
        "t3.micro",
        "t2.micro",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_instance_type_offering(filters=[aws.ec2.GetInstanceTypeOfferingFilterArgs(
        name="instance-type",
        values=[
            "t2.micro",
            "t3.micro",
        ],
    )],
    preferred_instance_types=[
        "t3.micro",
        "t2.micro",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetInstanceTypeOffering.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetInstanceTypeOfferingFilterInputArgs
            {
                Name = "instance-type",
                Values = new[]
                {
                    "t2.micro",
                    "t3.micro",
                },
            },
        },
        PreferredInstanceTypes = new[]
        {
            "t3.micro",
            "t2.micro",
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
		_, err := ec2.GetInstanceTypeOffering(ctx, &ec2.GetInstanceTypeOfferingArgs{
			Filters: []ec2.GetInstanceTypeOfferingFilter{
				{
					Name: "instance-type",
					Values: []string{
						"t2.micro",
						"t3.micro",
					},
				},
			},
			PreferredInstanceTypes: []string{
				"t3.micro",
				"t2.micro",
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
import com.pulumi.aws.ec2.inputs.GetInstanceTypeOfferingArgs;
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
        final var example = Ec2Functions.getInstanceTypeOffering(GetInstanceTypeOfferingArgs.builder()
            .filters(GetInstanceTypeOfferingFilterArgs.builder()
                .name("instance-type")
                .values(                
                    "t2.micro",
                    "t3.micro")
                .build())
            .preferredInstanceTypes(            
                "t3.micro",
                "t2.micro")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getInstanceTypeOffering
      Arguments:
        filters:
          - name: instance-type
            values:
              - t2.micro
              - t3.micro
        preferredInstanceTypes:
          - t3.micro
          - t2.micro
```
{{% /example %}}
{{% /examples %}}