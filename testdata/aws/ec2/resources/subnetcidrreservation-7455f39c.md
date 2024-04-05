Provides a subnet CIDR reservation resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.SubnetCidrReservation("example", {
    cidrBlock: "10.0.0.16/28",
    reservationType: "prefix",
    subnetId: aws_subnet.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.SubnetCidrReservation("example",
    cidr_block="10.0.0.16/28",
    reservation_type="prefix",
    subnet_id=aws_subnet["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.SubnetCidrReservation("example", new()
    {
        CidrBlock = "10.0.0.16/28",
        ReservationType = "prefix",
        SubnetId = aws_subnet.Example.Id,
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
		_, err := ec2.NewSubnetCidrReservation(ctx, "example", &ec2.SubnetCidrReservationArgs{
			CidrBlock:       pulumi.String("10.0.0.16/28"),
			ReservationType: pulumi.String("prefix"),
			SubnetId:        pulumi.Any(aws_subnet.Example.Id),
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
import com.pulumi.aws.ec2.SubnetCidrReservation;
import com.pulumi.aws.ec2.SubnetCidrReservationArgs;
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
        var example = new SubnetCidrReservation("example", SubnetCidrReservationArgs.builder()        
            .cidrBlock("10.0.0.16/28")
            .reservationType("prefix")
            .subnetId(aws_subnet.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:SubnetCidrReservation
    properties:
      cidrBlock: 10.0.0.16/28
      reservationType: prefix
      subnetId: ${aws_subnet.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Existing CIDR reservations using `SUBNET_ID:RESERVATION_ID`. For example:

```sh
 $ pulumi import aws:ec2/subnetCidrReservation:SubnetCidrReservation example subnet-01llsxvsxabqiymcz:scr-4mnvz6wb7otksjcs9
```
 