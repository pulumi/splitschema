Provides a Resource Access Manager (RAM) principal association. Depending if [RAM Sharing with AWS Organizations is enabled](https://docs.aws.amazon.com/ram/latest/userguide/getting-started-sharing.html#getting-started-sharing-orgs), the RAM behavior with different principal types changes.

When RAM Sharing with AWS Organizations is enabled:

- For AWS Account ID, Organization, and Organizational Unit principals within the same AWS Organization, no resource share invitation is sent and resources become available automatically after creating the association.
- For AWS Account ID principals outside the AWS Organization, a resource share invitation is sent and must be accepted before resources become available. See the `aws.ram.ResourceShareAccepter` resource to accept these invitations.

When RAM Sharing with AWS Organizations is not enabled:

- Organization and Organizational Unit principals cannot be used.
- For AWS Account ID principals, a resource share invitation is sent and must be accepted before resources become available. See the `aws.ram.ResourceShareAccepter` resource to accept these invitations.

{{% examples %}}
## Example Usage
{{% example %}}
### AWS Account ID

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleResourceShare = new aws.ram.ResourceShare("exampleResourceShare", {allowExternalPrincipals: true});
const examplePrincipalAssociation = new aws.ram.PrincipalAssociation("examplePrincipalAssociation", {
    principal: "111111111111",
    resourceShareArn: exampleResourceShare.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example_resource_share = aws.ram.ResourceShare("exampleResourceShare", allow_external_principals=True)
example_principal_association = aws.ram.PrincipalAssociation("examplePrincipalAssociation",
    principal="111111111111",
    resource_share_arn=example_resource_share.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleResourceShare = new Aws.Ram.ResourceShare("exampleResourceShare", new()
    {
        AllowExternalPrincipals = true,
    });

    var examplePrincipalAssociation = new Aws.Ram.PrincipalAssociation("examplePrincipalAssociation", new()
    {
        Principal = "111111111111",
        ResourceShareArn = exampleResourceShare.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleResourceShare, err := ram.NewResourceShare(ctx, "exampleResourceShare", &ram.ResourceShareArgs{
			AllowExternalPrincipals: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = ram.NewPrincipalAssociation(ctx, "examplePrincipalAssociation", &ram.PrincipalAssociationArgs{
			Principal:        pulumi.String("111111111111"),
			ResourceShareArn: exampleResourceShare.Arn,
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
import com.pulumi.aws.ram.ResourceShare;
import com.pulumi.aws.ram.ResourceShareArgs;
import com.pulumi.aws.ram.PrincipalAssociation;
import com.pulumi.aws.ram.PrincipalAssociationArgs;
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
        var exampleResourceShare = new ResourceShare("exampleResourceShare", ResourceShareArgs.builder()        
            .allowExternalPrincipals(true)
            .build());

        var examplePrincipalAssociation = new PrincipalAssociation("examplePrincipalAssociation", PrincipalAssociationArgs.builder()        
            .principal("111111111111")
            .resourceShareArn(exampleResourceShare.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleResourceShare:
    type: aws:ram:ResourceShare
    properties:
      # ... other configuration ...
      allowExternalPrincipals: true
  examplePrincipalAssociation:
    type: aws:ram:PrincipalAssociation
    properties:
      principal: '111111111111'
      resourceShareArn: ${exampleResourceShare.arn}
```
{{% /example %}}
{{% example %}}
### AWS Organization

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ram.PrincipalAssociation("example", {
    principal: aws_organizations_organization.example.arn,
    resourceShareArn: aws_ram_resource_share.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ram.PrincipalAssociation("example",
    principal=aws_organizations_organization["example"]["arn"],
    resource_share_arn=aws_ram_resource_share["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ram.PrincipalAssociation("example", new()
    {
        Principal = aws_organizations_organization.Example.Arn,
        ResourceShareArn = aws_ram_resource_share.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ram.NewPrincipalAssociation(ctx, "example", &ram.PrincipalAssociationArgs{
			Principal:        pulumi.Any(aws_organizations_organization.Example.Arn),
			ResourceShareArn: pulumi.Any(aws_ram_resource_share.Example.Arn),
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
import com.pulumi.aws.ram.PrincipalAssociation;
import com.pulumi.aws.ram.PrincipalAssociationArgs;
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
        var example = new PrincipalAssociation("example", PrincipalAssociationArgs.builder()        
            .principal(aws_organizations_organization.example().arn())
            .resourceShareArn(aws_ram_resource_share.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ram:PrincipalAssociation
    properties:
      principal: ${aws_organizations_organization.example.arn}
      resourceShareArn: ${aws_ram_resource_share.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import RAM Principal Associations using their Resource Share ARN and the `principal` separated by a comma. For example:

```sh
 $ pulumi import aws:ram/principalAssociation:PrincipalAssociation example arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12,123456789012
```
 