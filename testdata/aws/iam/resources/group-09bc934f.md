Provides an IAM group.

> **NOTE on user management:** Using `aws.iam.GroupMembership` or `aws.iam.UserGroupMembership` resources in addition to manually managing user/group membership using the console may lead to configuration drift or conflicts. For this reason, it's recommended to either manage membership entirely with the provider or entirely within the AWS console.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const developers = new aws.iam.Group("developers", {path: "/users/"});
```
```python
import pulumi
import pulumi_aws as aws

developers = aws.iam.Group("developers", path="/users/")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var developers = new Aws.Iam.Group("developers", new()
    {
        Path = "/users/",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewGroup(ctx, "developers", &iam.GroupArgs{
			Path: pulumi.String("/users/"),
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
import com.pulumi.aws.iam.Group;
import com.pulumi.aws.iam.GroupArgs;
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
        var developers = new Group("developers", GroupArgs.builder()        
            .path("/users/")
            .build());

    }
}
```
```yaml
resources:
  developers:
    type: aws:iam:Group
    properties:
      path: /users/
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM Groups using the `name`. For example:

```sh
 $ pulumi import aws:iam/group:Group developers developers
```
 