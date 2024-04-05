Adds the specified user to the specified group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleUserPool = new aws.cognito.UserPool("exampleUserPool", {passwordPolicy: {
    temporaryPasswordValidityDays: 7,
    minimumLength: 6,
    requireUppercase: false,
    requireSymbols: false,
    requireNumbers: false,
}});
const exampleUser = new aws.cognito.User("exampleUser", {
    userPoolId: exampleUserPool.id,
    username: "example",
});
const exampleUserGroup = new aws.cognito.UserGroup("exampleUserGroup", {userPoolId: exampleUserPool.id});
const exampleUserInGroup = new aws.cognito.UserInGroup("exampleUserInGroup", {
    userPoolId: exampleUserPool.id,
    groupName: exampleUserGroup.name,
    username: exampleUser.username,
});
```
```python
import pulumi
import pulumi_aws as aws

example_user_pool = aws.cognito.UserPool("exampleUserPool", password_policy=aws.cognito.UserPoolPasswordPolicyArgs(
    temporary_password_validity_days=7,
    minimum_length=6,
    require_uppercase=False,
    require_symbols=False,
    require_numbers=False,
))
example_user = aws.cognito.User("exampleUser",
    user_pool_id=example_user_pool.id,
    username="example")
example_user_group = aws.cognito.UserGroup("exampleUserGroup", user_pool_id=example_user_pool.id)
example_user_in_group = aws.cognito.UserInGroup("exampleUserInGroup",
    user_pool_id=example_user_pool.id,
    group_name=example_user_group.name,
    username=example_user.username)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleUserPool = new Aws.Cognito.UserPool("exampleUserPool", new()
    {
        PasswordPolicy = new Aws.Cognito.Inputs.UserPoolPasswordPolicyArgs
        {
            TemporaryPasswordValidityDays = 7,
            MinimumLength = 6,
            RequireUppercase = false,
            RequireSymbols = false,
            RequireNumbers = false,
        },
    });

    var exampleUser = new Aws.Cognito.User("exampleUser", new()
    {
        UserPoolId = exampleUserPool.Id,
        Username = "example",
    });

    var exampleUserGroup = new Aws.Cognito.UserGroup("exampleUserGroup", new()
    {
        UserPoolId = exampleUserPool.Id,
    });

    var exampleUserInGroup = new Aws.Cognito.UserInGroup("exampleUserInGroup", new()
    {
        UserPoolId = exampleUserPool.Id,
        GroupName = exampleUserGroup.Name,
        Username = exampleUser.Username,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleUserPool, err := cognito.NewUserPool(ctx, "exampleUserPool", &cognito.UserPoolArgs{
			PasswordPolicy: &cognito.UserPoolPasswordPolicyArgs{
				TemporaryPasswordValidityDays: pulumi.Int(7),
				MinimumLength:                 pulumi.Int(6),
				RequireUppercase:              pulumi.Bool(false),
				RequireSymbols:                pulumi.Bool(false),
				RequireNumbers:                pulumi.Bool(false),
			},
		})
		if err != nil {
			return err
		}
		exampleUser, err := cognito.NewUser(ctx, "exampleUser", &cognito.UserArgs{
			UserPoolId: exampleUserPool.ID(),
			Username:   pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		exampleUserGroup, err := cognito.NewUserGroup(ctx, "exampleUserGroup", &cognito.UserGroupArgs{
			UserPoolId: exampleUserPool.ID(),
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewUserInGroup(ctx, "exampleUserInGroup", &cognito.UserInGroupArgs{
			UserPoolId: exampleUserPool.ID(),
			GroupName:  exampleUserGroup.Name,
			Username:   exampleUser.Username,
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.UserPoolArgs;
import com.pulumi.aws.cognito.inputs.UserPoolPasswordPolicyArgs;
import com.pulumi.aws.cognito.User;
import com.pulumi.aws.cognito.UserArgs;
import com.pulumi.aws.cognito.UserGroup;
import com.pulumi.aws.cognito.UserGroupArgs;
import com.pulumi.aws.cognito.UserInGroup;
import com.pulumi.aws.cognito.UserInGroupArgs;
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
        var exampleUserPool = new UserPool("exampleUserPool", UserPoolArgs.builder()        
            .passwordPolicy(UserPoolPasswordPolicyArgs.builder()
                .temporaryPasswordValidityDays(7)
                .minimumLength(6)
                .requireUppercase(false)
                .requireSymbols(false)
                .requireNumbers(false)
                .build())
            .build());

        var exampleUser = new User("exampleUser", UserArgs.builder()        
            .userPoolId(exampleUserPool.id())
            .username("example")
            .build());

        var exampleUserGroup = new UserGroup("exampleUserGroup", UserGroupArgs.builder()        
            .userPoolId(exampleUserPool.id())
            .build());

        var exampleUserInGroup = new UserInGroup("exampleUserInGroup", UserInGroupArgs.builder()        
            .userPoolId(exampleUserPool.id())
            .groupName(exampleUserGroup.name())
            .username(exampleUser.username())
            .build());

    }
}
```
```yaml
resources:
  exampleUserPool:
    type: aws:cognito:UserPool
    properties:
      passwordPolicy:
        temporaryPasswordValidityDays: 7
        minimumLength: 6
        requireUppercase: false
        requireSymbols: false
        requireNumbers: false
  exampleUser:
    type: aws:cognito:User
    properties:
      userPoolId: ${exampleUserPool.id}
      username: example
  exampleUserGroup:
    type: aws:cognito:UserGroup
    properties:
      userPoolId: ${exampleUserPool.id}
  exampleUserInGroup:
    type: aws:cognito:UserInGroup
    properties:
      userPoolId: ${exampleUserPool.id}
      groupName: ${exampleUserGroup.name}
      username: ${exampleUser.username}
```
{{% /example %}}
{{% /examples %}}