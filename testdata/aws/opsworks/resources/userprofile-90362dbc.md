Provides an OpsWorks User Profile resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myProfile = new aws.opsworks.UserProfile("myProfile", {
    userArn: aws_iam_user.user.arn,
    sshUsername: "my_user",
});
```
```python
import pulumi
import pulumi_aws as aws

my_profile = aws.opsworks.UserProfile("myProfile",
    user_arn=aws_iam_user["user"]["arn"],
    ssh_username="my_user")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myProfile = new Aws.OpsWorks.UserProfile("myProfile", new()
    {
        UserArn = aws_iam_user.User.Arn,
        SshUsername = "my_user",
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
		_, err := opsworks.NewUserProfile(ctx, "myProfile", &opsworks.UserProfileArgs{
			UserArn:     pulumi.Any(aws_iam_user.User.Arn),
			SshUsername: pulumi.String("my_user"),
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
import com.pulumi.aws.opsworks.UserProfile;
import com.pulumi.aws.opsworks.UserProfileArgs;
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
        var myProfile = new UserProfile("myProfile", UserProfileArgs.builder()        
            .userArn(aws_iam_user.user().arn())
            .sshUsername("my_user")
            .build());

    }
}
```
```yaml
resources:
  myProfile:
    type: aws:opsworks:UserProfile
    properties:
      userArn: ${aws_iam_user.user.arn}
      sshUsername: my_user
```
{{% /example %}}
{{% /examples %}}