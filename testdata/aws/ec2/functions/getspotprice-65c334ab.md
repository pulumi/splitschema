Information about most recent Spot Price for a given EC2 instance.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getSpotPrice({
    availabilityZone: "us-west-2a",
    filters: [{
        name: "product-description",
        values: ["Linux/UNIX"],
    }],
    instanceType: "t3.medium",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_spot_price(availability_zone="us-west-2a",
    filters=[aws.ec2.GetSpotPriceFilterArgs(
        name="product-description",
        values=["Linux/UNIX"],
    )],
    instance_type="t3.medium")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetSpotPrice.Invoke(new()
    {
        AvailabilityZone = "us-west-2a",
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetSpotPriceFilterInputArgs
            {
                Name = "product-description",
                Values = new[]
                {
                    "Linux/UNIX",
                },
            },
        },
        InstanceType = "t3.medium",
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
		_, err := ec2.GetSpotPrice(ctx, &ec2.GetSpotPriceArgs{
			AvailabilityZone: pulumi.StringRef("us-west-2a"),
			Filters: []ec2.GetSpotPriceFilter{
				{
					Name: "product-description",
					Values: []string{
						"Linux/UNIX",
					},
				},
			},
			InstanceType: pulumi.StringRef("t3.medium"),
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
import com.pulumi.aws.ec2.inputs.GetSpotPriceArgs;
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
        final var example = Ec2Functions.getSpotPrice(GetSpotPriceArgs.builder()
            .availabilityZone("us-west-2a")
            .filters(GetSpotPriceFilterArgs.builder()
                .name("product-description")
                .values("Linux/UNIX")
                .build())
            .instanceType("t3.medium")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getSpotPrice
      Arguments:
        availabilityZone: us-west-2a
        filters:
          - name: product-description
            values:
              - Linux/UNIX
        instanceType: t3.medium
```
{{% /example %}}
{{% /examples %}}