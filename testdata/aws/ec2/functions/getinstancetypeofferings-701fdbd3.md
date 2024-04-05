Information about EC2 Instance Type Offerings.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getInstanceTypeOfferings({
    filters: [
        {
            name: "instance-type",
            values: [
                "t2.micro",
                "t3.micro",
            ],
        },
        {
            name: "location",
            values: ["usw2-az4"],
        },
    ],
    locationType: "availability-zone-id",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_instance_type_offerings(filters=[
        aws.ec2.GetInstanceTypeOfferingsFilterArgs(
            name="instance-type",
            values=[
                "t2.micro",
                "t3.micro",
            ],
        ),
        aws.ec2.GetInstanceTypeOfferingsFilterArgs(
            name="location",
            values=["usw2-az4"],
        ),
    ],
    location_type="availability-zone-id")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetInstanceTypeOfferings.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetInstanceTypeOfferingsFilterInputArgs
            {
                Name = "instance-type",
                Values = new[]
                {
                    "t2.micro",
                    "t3.micro",
                },
            },
            new Aws.Ec2.Inputs.GetInstanceTypeOfferingsFilterInputArgs
            {
                Name = "location",
                Values = new[]
                {
                    "usw2-az4",
                },
            },
        },
        LocationType = "availability-zone-id",
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
		_, err := ec2.GetInstanceTypeOfferings(ctx, &ec2.GetInstanceTypeOfferingsArgs{
			Filters: []ec2.GetInstanceTypeOfferingsFilter{
				{
					Name: "instance-type",
					Values: []string{
						"t2.micro",
						"t3.micro",
					},
				},
				{
					Name: "location",
					Values: []string{
						"usw2-az4",
					},
				},
			},
			LocationType: pulumi.StringRef("availability-zone-id"),
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
import com.pulumi.aws.ec2.inputs.GetInstanceTypeOfferingsArgs;
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
        final var example = Ec2Functions.getInstanceTypeOfferings(GetInstanceTypeOfferingsArgs.builder()
            .filters(            
                GetInstanceTypeOfferingsFilterArgs.builder()
                    .name("instance-type")
                    .values(                    
                        "t2.micro",
                        "t3.micro")
                    .build(),
                GetInstanceTypeOfferingsFilterArgs.builder()
                    .name("location")
                    .values("usw2-az4")
                    .build())
            .locationType("availability-zone-id")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getInstanceTypeOfferings
      Arguments:
        filters:
          - name: instance-type
            values:
              - t2.micro
              - t3.micro
          - name: location
            values:
              - usw2-az4
        locationType: availability-zone-id
```
{{% /example %}}
{{% /examples %}}