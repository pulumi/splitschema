Provides an network ACL association resource which allows you to associate your network ACL with any subnet(s).

> **NOTE on Network ACLs and Network ACL Associations:** the provider provides both a standalone network ACL association resource
and a network ACL resource with a `subnet_ids` attribute. Do not use the same subnet ID in both a network ACL
resource and a network ACL association resource. Doing so will cause a conflict of associations and will overwrite the association.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.NetworkAclAssociation("main", {
    networkAclId: aws_network_acl.main.id,
    subnetId: aws_subnet.main.id,
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ec2.NetworkAclAssociation("main",
    network_acl_id=aws_network_acl["main"]["id"],
    subnet_id=aws_subnet["main"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ec2.NetworkAclAssociation("main", new()
    {
        NetworkAclId = aws_network_acl.Main.Id,
        SubnetId = aws_subnet.Main.Id,
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
		_, err := ec2.NewNetworkAclAssociation(ctx, "main", &ec2.NetworkAclAssociationArgs{
			NetworkAclId: pulumi.Any(aws_network_acl.Main.Id),
			SubnetId:     pulumi.Any(aws_subnet.Main.Id),
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
import com.pulumi.aws.ec2.NetworkAclAssociation;
import com.pulumi.aws.ec2.NetworkAclAssociationArgs;
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
        var main = new NetworkAclAssociation("main", NetworkAclAssociationArgs.builder()        
            .networkAclId(aws_network_acl.main().id())
            .subnetId(aws_subnet.main().id())
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:ec2:NetworkAclAssociation
    properties:
      networkAclId: ${aws_network_acl.main.id}
      subnetId: ${aws_subnet.main.id}
```
{{% /example %}}
{{% /examples %}}