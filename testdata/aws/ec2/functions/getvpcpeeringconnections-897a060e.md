Use this data source to get IDs of Amazon VPC peering connections
To get more details on each connection, use the data resource aws.ec2.VpcPeeringConnection

Note: To use this data source in a count, the resources should exist before trying to access
the data source.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const pcs = aws.ec2.getVpcPeeringConnections({
    filters: [{
        name: "requester-vpc-info.vpc-id",
        values: [aws_vpc.foo.id],
    }],
});
const pc = .map(__index => (aws.ec2.getVpcPeeringConnection({
    id: _arg0_.ids[__index],
})));
```
```python
import pulumi
import pulumi_aws as aws

pcs = aws.ec2.get_vpc_peering_connections(filters=[aws.ec2.GetVpcPeeringConnectionsFilterArgs(
    name="requester-vpc-info.vpc-id",
    values=[aws_vpc["foo"]["id"]],
)])
pc = [aws.ec2.get_vpc_peering_connection(id=pcs.ids[__index]) for __index in range(len(pcs.ids))]
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var pcs = Aws.Ec2.GetVpcPeeringConnections.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetVpcPeeringConnectionsFilterInputArgs
            {
                Name = "requester-vpc-info.vpc-id",
                Values = new[]
                {
                    aws_vpc.Foo.Id,
                },
            },
        },
    });

    var pc = ;

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
pcs, err := ec2.GetVpcPeeringConnections(ctx, &ec2.GetVpcPeeringConnectionsArgs{
Filters: []ec2.GetVpcPeeringConnectionsFilter{
{
Name: "requester-vpc-info.vpc-id",
Values: interface{}{
aws_vpc.Foo.Id,
},
},
},
}, nil);
if err != nil {
return err
}
_ := "TODO: For expression";
return nil
})
}
```
{{% /example %}}
{{% /examples %}}