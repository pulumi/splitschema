Provides a resource to create an association between a VPC endpoint and a security group.

> **NOTE on VPC Endpoints and VPC Endpoint Security Group Associations:** The provider provides
both a standalone VPC Endpoint Security Group Association (an association between a VPC endpoint
and a single `security_group_id`) and a VPC Endpoint resource with a `security_group_ids`
attribute. Do not use the same security group ID in both a VPC Endpoint resource and a VPC Endpoint Security
Group Association resource. Doing so will cause a conflict of associations and will overwrite the association.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sgEc2 = new aws.ec2.SecurityGroupAssociation("sgEc2", {
    vpcEndpointId: aws_vpc_endpoint.ec2.id,
    securityGroupId: aws_security_group.sg.id,
});
```
```python
import pulumi
import pulumi_aws as aws

sg_ec2 = aws.ec2.SecurityGroupAssociation("sgEc2",
    vpc_endpoint_id=aws_vpc_endpoint["ec2"]["id"],
    security_group_id=aws_security_group["sg"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var sgEc2 = new Aws.Ec2.SecurityGroupAssociation("sgEc2", new()
    {
        VpcEndpointId = aws_vpc_endpoint.Ec2.Id,
        SecurityGroupId = aws_security_group.Sg.Id,
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
		_, err := ec2.NewSecurityGroupAssociation(ctx, "sgEc2", &ec2.SecurityGroupAssociationArgs{
			VpcEndpointId:   pulumi.Any(aws_vpc_endpoint.Ec2.Id),
			SecurityGroupId: pulumi.Any(aws_security_group.Sg.Id),
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
import com.pulumi.aws.ec2.SecurityGroupAssociation;
import com.pulumi.aws.ec2.SecurityGroupAssociationArgs;
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
        var sgEc2 = new SecurityGroupAssociation("sgEc2", SecurityGroupAssociationArgs.builder()        
            .vpcEndpointId(aws_vpc_endpoint.ec2().id())
            .securityGroupId(aws_security_group.sg().id())
            .build());

    }
}
```
```yaml
resources:
  sgEc2:
    type: aws:ec2:SecurityGroupAssociation
    properties:
      vpcEndpointId: ${aws_vpc_endpoint.ec2.id}
      securityGroupId: ${aws_security_group.sg.id}
```
{{% /example %}}
{{% /examples %}}