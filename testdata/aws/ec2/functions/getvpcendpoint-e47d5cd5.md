The VPC Endpoint data source provides details about
a specific VPC endpoint.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const s3 = aws.ec2.getVpcEndpoint({
    vpcId: aws_vpc.foo.id,
    serviceName: "com.amazonaws.us-west-2.s3",
});
const privateS3 = new aws.ec2.VpcEndpointRouteTableAssociation("privateS3", {
    vpcEndpointId: s3.then(s3 => s3.id),
    routeTableId: aws_route_table["private"].id,
});
```
```python
import pulumi
import pulumi_aws as aws

s3 = aws.ec2.get_vpc_endpoint(vpc_id=aws_vpc["foo"]["id"],
    service_name="com.amazonaws.us-west-2.s3")
private_s3 = aws.ec2.VpcEndpointRouteTableAssociation("privateS3",
    vpc_endpoint_id=s3.id,
    route_table_id=aws_route_table["private"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var s3 = Aws.Ec2.GetVpcEndpoint.Invoke(new()
    {
        VpcId = aws_vpc.Foo.Id,
        ServiceName = "com.amazonaws.us-west-2.s3",
    });

    var privateS3 = new Aws.Ec2.VpcEndpointRouteTableAssociation("privateS3", new()
    {
        VpcEndpointId = s3.Apply(getVpcEndpointResult => getVpcEndpointResult.Id),
        RouteTableId = aws_route_table.Private.Id,
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
		s3, err := ec2.LookupVpcEndpoint(ctx, &ec2.LookupVpcEndpointArgs{
			VpcId:       pulumi.StringRef(aws_vpc.Foo.Id),
			ServiceName: pulumi.StringRef("com.amazonaws.us-west-2.s3"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcEndpointRouteTableAssociation(ctx, "privateS3", &ec2.VpcEndpointRouteTableAssociationArgs{
			VpcEndpointId: *pulumi.String(s3.Id),
			RouteTableId:  pulumi.Any(aws_route_table.Private.Id),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetVpcEndpointArgs;
import com.pulumi.aws.ec2.VpcEndpointRouteTableAssociation;
import com.pulumi.aws.ec2.VpcEndpointRouteTableAssociationArgs;
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
        final var s3 = Ec2Functions.getVpcEndpoint(GetVpcEndpointArgs.builder()
            .vpcId(aws_vpc.foo().id())
            .serviceName("com.amazonaws.us-west-2.s3")
            .build());

        var privateS3 = new VpcEndpointRouteTableAssociation("privateS3", VpcEndpointRouteTableAssociationArgs.builder()        
            .vpcEndpointId(s3.applyValue(getVpcEndpointResult -> getVpcEndpointResult.id()))
            .routeTableId(aws_route_table.private().id())
            .build());

    }
}
```
```yaml
resources:
  privateS3:
    type: aws:ec2:VpcEndpointRouteTableAssociation
    properties:
      vpcEndpointId: ${s3.id}
      routeTableId: ${aws_route_table.private.id}
variables:
  s3:
    fn::invoke:
      Function: aws:ec2:getVpcEndpoint
      Arguments:
        vpcId: ${aws_vpc.foo.id}
        serviceName: com.amazonaws.us-west-2.s3
```
{{% /example %}}
{{% /examples %}}