Resource for managing an AWS VPC Lattice Service Network VPC Association.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.vpclattice.ServiceNetworkVpcAssociation("example", {
    vpcIdentifier: aws_vpc.example.id,
    serviceNetworkIdentifier: aws_vpclattice_service_network.example.id,
    securityGroupIds: [aws_security_group.example.id],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.ServiceNetworkVpcAssociation("example",
    vpc_identifier=aws_vpc["example"]["id"],
    service_network_identifier=aws_vpclattice_service_network["example"]["id"],
    security_group_ids=[aws_security_group["example"]["id"]])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VpcLattice.ServiceNetworkVpcAssociation("example", new()
    {
        VpcIdentifier = aws_vpc.Example.Id,
        ServiceNetworkIdentifier = aws_vpclattice_service_network.Example.Id,
        SecurityGroupIds = new[]
        {
            aws_security_group.Example.Id,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpclattice.NewServiceNetworkVpcAssociation(ctx, "example", &vpclattice.ServiceNetworkVpcAssociationArgs{
			VpcIdentifier:            pulumi.Any(aws_vpc.Example.Id),
			ServiceNetworkIdentifier: pulumi.Any(aws_vpclattice_service_network.Example.Id),
			SecurityGroupIds: pulumi.StringArray{
				aws_security_group.Example.Id,
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
import com.pulumi.aws.vpclattice.ServiceNetworkVpcAssociation;
import com.pulumi.aws.vpclattice.ServiceNetworkVpcAssociationArgs;
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
        var example = new ServiceNetworkVpcAssociation("example", ServiceNetworkVpcAssociationArgs.builder()        
            .vpcIdentifier(aws_vpc.example().id())
            .serviceNetworkIdentifier(aws_vpclattice_service_network.example().id())
            .securityGroupIds(aws_security_group.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:vpclattice:ServiceNetworkVpcAssociation
    properties:
      vpcIdentifier: ${aws_vpc.example.id}
      serviceNetworkIdentifier: ${aws_vpclattice_service_network.example.id}
      securityGroupIds:
        - ${aws_security_group.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Lattice Service Network VPC Association using the `id`. For example:

```sh
 $ pulumi import aws:vpclattice/serviceNetworkVpcAssociation:ServiceNetworkVpcAssociation example snsa-05e2474658a88f6ba
```
 