Provides a resource to create an association between a VPC endpoint and a subnet.

> **NOTE on VPC Endpoints and VPC Endpoint Subnet Associations:** This provider provides
both a standalone VPC Endpoint Subnet Association (an association between a VPC endpoint
and a single `subnet_id`) and a VPC Endpoint resource with a `subnet_ids`
attribute. Do not use the same subnet ID in both a VPC Endpoint resource and a VPC Endpoint Subnet
Association resource. Doing so will cause a conflict of associations and will overwrite the association.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const snEc2 = new aws.ec2.VpcEndpointSubnetAssociation("snEc2", {
    vpcEndpointId: aws_vpc_endpoint.ec2.id,
    subnetId: aws_subnet.sn.id,
});
```
```python
import pulumi
import pulumi_aws as aws

sn_ec2 = aws.ec2.VpcEndpointSubnetAssociation("snEc2",
    vpc_endpoint_id=aws_vpc_endpoint["ec2"]["id"],
    subnet_id=aws_subnet["sn"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var snEc2 = new Aws.Ec2.VpcEndpointSubnetAssociation("snEc2", new()
    {
        VpcEndpointId = aws_vpc_endpoint.Ec2.Id,
        SubnetId = aws_subnet.Sn.Id,
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
		_, err := ec2.NewVpcEndpointSubnetAssociation(ctx, "snEc2", &ec2.VpcEndpointSubnetAssociationArgs{
			VpcEndpointId: pulumi.Any(aws_vpc_endpoint.Ec2.Id),
			SubnetId:      pulumi.Any(aws_subnet.Sn.Id),
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
import com.pulumi.aws.ec2.VpcEndpointSubnetAssociation;
import com.pulumi.aws.ec2.VpcEndpointSubnetAssociationArgs;
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
        var snEc2 = new VpcEndpointSubnetAssociation("snEc2", VpcEndpointSubnetAssociationArgs.builder()        
            .vpcEndpointId(aws_vpc_endpoint.ec2().id())
            .subnetId(aws_subnet.sn().id())
            .build());

    }
}
```
```yaml
resources:
  snEc2:
    type: aws:ec2:VpcEndpointSubnetAssociation
    properties:
      vpcEndpointId: ${aws_vpc_endpoint.ec2.id}
      subnetId: ${aws_subnet.sn.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Endpoint Subnet Associations using `vpc_endpoint_id` together with `subnet_id`. For example:

```sh
 $ pulumi import aws:ec2/vpcEndpointSubnetAssociation:VpcEndpointSubnetAssociation example vpce-aaaaaaaa/subnet-bbbbbbbbbbbbbbbbb
```
 