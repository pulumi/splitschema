Attaches a permissions boundary policy to a Single Sign-On (SSO) Permission Set resource.

> **NOTE:** A permission set can have at most one permissions boundary attached; using more than one `aws.ssoadmin.PermissionsBoundaryAttachment` references the same permission set will show a permanent difference.

{{% examples %}}
## Example Usage
{{% example %}}
### Attaching an AWS-managed policy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ssoadmin.PermissionsBoundaryAttachment("example", {
    instanceArn: aws_ssoadmin_permission_set.example.instance_arn,
    permissionSetArn: aws_ssoadmin_permission_set.example.arn,
    permissionsBoundary: {
        managedPolicyArn: "arn:aws:iam::aws:policy/ReadOnlyAccess",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssoadmin.PermissionsBoundaryAttachment("example",
    instance_arn=aws_ssoadmin_permission_set["example"]["instance_arn"],
    permission_set_arn=aws_ssoadmin_permission_set["example"]["arn"],
    permissions_boundary=aws.ssoadmin.PermissionsBoundaryAttachmentPermissionsBoundaryArgs(
        managed_policy_arn="arn:aws:iam::aws:policy/ReadOnlyAccess",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SsoAdmin.PermissionsBoundaryAttachment("example", new()
    {
        InstanceArn = aws_ssoadmin_permission_set.Example.Instance_arn,
        PermissionSetArn = aws_ssoadmin_permission_set.Example.Arn,
        PermissionsBoundary = new Aws.SsoAdmin.Inputs.PermissionsBoundaryAttachmentPermissionsBoundaryArgs
        {
            ManagedPolicyArn = "arn:aws:iam::aws:policy/ReadOnlyAccess",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssoadmin"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssoadmin.NewPermissionsBoundaryAttachment(ctx, "example", &ssoadmin.PermissionsBoundaryAttachmentArgs{
			InstanceArn:      pulumi.Any(aws_ssoadmin_permission_set.Example.Instance_arn),
			PermissionSetArn: pulumi.Any(aws_ssoadmin_permission_set.Example.Arn),
			PermissionsBoundary: &ssoadmin.PermissionsBoundaryAttachmentPermissionsBoundaryArgs{
				ManagedPolicyArn: pulumi.String("arn:aws:iam::aws:policy/ReadOnlyAccess"),
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
import com.pulumi.aws.ssoadmin.PermissionsBoundaryAttachment;
import com.pulumi.aws.ssoadmin.PermissionsBoundaryAttachmentArgs;
import com.pulumi.aws.ssoadmin.inputs.PermissionsBoundaryAttachmentPermissionsBoundaryArgs;
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
        var example = new PermissionsBoundaryAttachment("example", PermissionsBoundaryAttachmentArgs.builder()        
            .instanceArn(aws_ssoadmin_permission_set.example().instance_arn())
            .permissionSetArn(aws_ssoadmin_permission_set.example().arn())
            .permissionsBoundary(PermissionsBoundaryAttachmentPermissionsBoundaryArgs.builder()
                .managedPolicyArn("arn:aws:iam::aws:policy/ReadOnlyAccess")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ssoadmin:PermissionsBoundaryAttachment
    properties:
      instanceArn: ${aws_ssoadmin_permission_set.example.instance_arn}
      permissionSetArn: ${aws_ssoadmin_permission_set.example.arn}
      permissionsBoundary:
        managedPolicyArn: arn:aws:iam::aws:policy/ReadOnlyAccess
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SSO Admin Permissions Boundary Attachments using the `permission_set_arn` and `instance_arn`, separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:ssoadmin/permissionsBoundaryAttachment:PermissionsBoundaryAttachment example arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
```
 