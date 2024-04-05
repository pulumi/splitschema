Provides an OpsWorks permission resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myStackPermission = new aws.opsworks.Permission("myStackPermission", {
    allowSsh: true,
    allowSudo: true,
    level: "iam_only",
    userArn: aws_iam_user.user.arn,
    stackId: aws_opsworks_stack.stack.id,
});
```
```python
import pulumi
import pulumi_aws as aws

my_stack_permission = aws.opsworks.Permission("myStackPermission",
    allow_ssh=True,
    allow_sudo=True,
    level="iam_only",
    user_arn=aws_iam_user["user"]["arn"],
    stack_id=aws_opsworks_stack["stack"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myStackPermission = new Aws.OpsWorks.Permission("myStackPermission", new()
    {
        AllowSsh = true,
        AllowSudo = true,
        Level = "iam_only",
        UserArn = aws_iam_user.User.Arn,
        StackId = aws_opsworks_stack.Stack.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewPermission(ctx, "myStackPermission", &opsworks.PermissionArgs{
			AllowSsh:  pulumi.Bool(true),
			AllowSudo: pulumi.Bool(true),
			Level:     pulumi.String("iam_only"),
			UserArn:   pulumi.Any(aws_iam_user.User.Arn),
			StackId:   pulumi.Any(aws_opsworks_stack.Stack.Id),
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
import com.pulumi.aws.opsworks.Permission;
import com.pulumi.aws.opsworks.PermissionArgs;
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
        var myStackPermission = new Permission("myStackPermission", PermissionArgs.builder()        
            .allowSsh(true)
            .allowSudo(true)
            .level("iam_only")
            .userArn(aws_iam_user.user().arn())
            .stackId(aws_opsworks_stack.stack().id())
            .build());

    }
}
```
```yaml
resources:
  myStackPermission:
    type: aws:opsworks:Permission
    properties:
      allowSsh: true
      allowSudo: true
      level: iam_only
      userArn: ${aws_iam_user.user.arn}
      stackId: ${aws_opsworks_stack.stack.id}
```
{{% /example %}}
{{% /examples %}}