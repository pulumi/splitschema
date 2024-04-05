Provides an EC2 Spot Instance Request resource. This allows instances to be
requested on the spot market.

By default this provider creates Spot Instance Requests with a `persistent` type,
which means that for the duration of their lifetime, AWS will launch an
instance with the configured details if and when the spot market will accept
the requested price.

On destruction, this provider will make an attempt to terminate the associated Spot
Instance if there is one present.

Spot Instances requests with a `one-time` type will close the spot request
when the instance is terminated either by the request being below the current spot
price availability or by a user.

> **NOTE:** Because their behavior depends on the live status of the spot
market, Spot Instance Requests have a unique lifecycle that makes them behave
differently than other resources. Most importantly: there is __no
guarantee__ that a Spot Instance exists to fulfill the request at any given
point in time. See the [AWS Spot Instance
documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances.html)
for more information.

> **NOTE [AWS strongly discourages](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-best-practices.html#which-spot-request-method-to-use) the use of the legacy APIs called by this resource.
We recommend using the EC2 Instance resource with `instance_market_options` instead.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Request a spot instance at $0.03
const cheapWorker = new aws.ec2.SpotInstanceRequest("cheapWorker", {
    ami: "ami-1234",
    instanceType: "c4.xlarge",
    spotPrice: "0.03",
    tags: {
        Name: "CheapWorker",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

# Request a spot instance at $0.03
cheap_worker = aws.ec2.SpotInstanceRequest("cheapWorker",
    ami="ami-1234",
    instance_type="c4.xlarge",
    spot_price="0.03",
    tags={
        "Name": "CheapWorker",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Request a spot instance at $0.03
    var cheapWorker = new Aws.Ec2.SpotInstanceRequest("cheapWorker", new()
    {
        Ami = "ami-1234",
        InstanceType = "c4.xlarge",
        SpotPrice = "0.03",
        Tags = 
        {
            { "Name", "CheapWorker" },
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
		_, err := ec2.NewSpotInstanceRequest(ctx, "cheapWorker", &ec2.SpotInstanceRequestArgs{
			Ami:          pulumi.String("ami-1234"),
			InstanceType: pulumi.String("c4.xlarge"),
			SpotPrice:    pulumi.String("0.03"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("CheapWorker"),
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
import com.pulumi.aws.ec2.SpotInstanceRequest;
import com.pulumi.aws.ec2.SpotInstanceRequestArgs;
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
        var cheapWorker = new SpotInstanceRequest("cheapWorker", SpotInstanceRequestArgs.builder()        
            .ami("ami-1234")
            .instanceType("c4.xlarge")
            .spotPrice("0.03")
            .tags(Map.of("Name", "CheapWorker"))
            .build());

    }
}
```
```yaml
resources:
  # Request a spot instance at $0.03
  cheapWorker:
    type: aws:ec2:SpotInstanceRequest
    properties:
      ami: ami-1234
      instanceType: c4.xlarge
      spotPrice: '0.03'
      tags:
        Name: CheapWorker
```
{{% /example %}}
{{% /examples %}}