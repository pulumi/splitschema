> **WARNING:** Multiple aws.iam.GroupMembership resources with the same group name will produce inconsistent behavior!

Provides a top level resource to manage IAM Group membership for IAM Users. For
more information on managing IAM Groups or IAM Users, see IAM Groups or
IAM Users

> **Note:** `aws.iam.GroupMembership` will conflict with itself if used more than once with the same group. To non-exclusively manage the users in a group, see the
`aws.iam.UserGroupMembership` resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const group = new aws.iam.Group("group", {});
const userOne = new aws.iam.User("userOne", {});
const userTwo = new aws.iam.User("userTwo", {});
const team = new aws.iam.GroupMembership("team", {
    users: [
        userOne.name,
        userTwo.name,
    ],
    group: group.name,
});
```
```python
import pulumi
import pulumi_aws as aws

group = aws.iam.Group("group")
user_one = aws.iam.User("userOne")
user_two = aws.iam.User("userTwo")
team = aws.iam.GroupMembership("team",
    users=[
        user_one.name,
        user_two.name,
    ],
    group=group.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @group = new Aws.Iam.Group("group");

    var userOne = new Aws.Iam.User("userOne");

    var userTwo = new Aws.Iam.User("userTwo");

    var team = new Aws.Iam.GroupMembership("team", new()
    {
        Users = new[]
        {
            userOne.Name,
            userTwo.Name,
        },
        Group = @group.Name,
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
		group, err := iam.NewGroup(ctx, "group", nil)
		if err != nil {
			return err
		}
		userOne, err := iam.NewUser(ctx, "userOne", nil)
		if err != nil {
			return err
		}
		userTwo, err := iam.NewUser(ctx, "userTwo", nil)
		if err != nil {
			return err
		}
		_, err = iam.NewGroupMembership(ctx, "team", &iam.GroupMembershipArgs{
			Users: pulumi.StringArray{
				userOne.Name,
				userTwo.Name,
			},
			Group: group.Name,
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
import com.pulumi.aws.iam.User;
import com.pulumi.aws.iam.GroupMembership;
import com.pulumi.aws.iam.GroupMembershipArgs;
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
        var group = new Group("group");

        var userOne = new User("userOne");

        var userTwo = new User("userTwo");

        var team = new GroupMembership("team", GroupMembershipArgs.builder()        
            .users(            
                userOne.name(),
                userTwo.name())
            .group(group.name())
            .build());

    }
}
```
```yaml
resources:
  team:
    type: aws:iam:GroupMembership
    properties:
      users:
        - ${userOne.name}
        - ${userTwo.name}
      group: ${group.name}
  group:
    type: aws:iam:Group
  userOne:
    type: aws:iam:User
  userTwo:
    type: aws:iam:User
```
{{% /example %}}
{{% /examples %}}