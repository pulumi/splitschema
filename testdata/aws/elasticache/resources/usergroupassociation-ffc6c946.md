Associate an existing ElastiCache user and an existing user group.

> **NOTE:** The provider will detect changes in the `aws.elasticache.UserGroup` since `aws.elasticache.UserGroupAssociation` changes the user IDs associated with the user group. You can ignore these changes with the `ignore_changes` option as shown in the example.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.elasticache.User("default", {
    userId: "defaultUserID",
    userName: "default",
    accessString: "on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
    engine: "REDIS",
    passwords: ["password123456789"],
});
const exampleUserGroup = new aws.elasticache.UserGroup("exampleUserGroup", {
    engine: "REDIS",
    userGroupId: "userGroupId",
    userIds: [_default.userId],
});
const exampleUser = new aws.elasticache.User("exampleUser", {
    userId: "exampleUserID",
    userName: "exampleuser",
    accessString: "on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
    engine: "REDIS",
    passwords: ["password123456789"],
});
const exampleUserGroupAssociation = new aws.elasticache.UserGroupAssociation("exampleUserGroupAssociation", {
    userGroupId: exampleUserGroup.userGroupId,
    userId: exampleUser.userId,
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.elasticache.User("default",
    user_id="defaultUserID",
    user_name="default",
    access_string="on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
    engine="REDIS",
    passwords=["password123456789"])
example_user_group = aws.elasticache.UserGroup("exampleUserGroup",
    engine="REDIS",
    user_group_id="userGroupId",
    user_ids=[default.user_id])
example_user = aws.elasticache.User("exampleUser",
    user_id="exampleUserID",
    user_name="exampleuser",
    access_string="on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
    engine="REDIS",
    passwords=["password123456789"])
example_user_group_association = aws.elasticache.UserGroupAssociation("exampleUserGroupAssociation",
    user_group_id=example_user_group.user_group_id,
    user_id=example_user.user_id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.ElastiCache.User("default", new()
    {
        UserId = "defaultUserID",
        UserName = "default",
        AccessString = "on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
        Engine = "REDIS",
        Passwords = new[]
        {
            "password123456789",
        },
    });

    var exampleUserGroup = new Aws.ElastiCache.UserGroup("exampleUserGroup", new()
    {
        Engine = "REDIS",
        UserGroupId = "userGroupId",
        UserIds = new[]
        {
            @default.UserId,
        },
    });

    var exampleUser = new Aws.ElastiCache.User("exampleUser", new()
    {
        UserId = "exampleUserID",
        UserName = "exampleuser",
        AccessString = "on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
        Engine = "REDIS",
        Passwords = new[]
        {
            "password123456789",
        },
    });

    var exampleUserGroupAssociation = new Aws.ElastiCache.UserGroupAssociation("exampleUserGroupAssociation", new()
    {
        UserGroupId = exampleUserGroup.UserGroupId,
        UserId = exampleUser.UserId,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticache"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticache.NewUser(ctx, "default", &elasticache.UserArgs{
			UserId:       pulumi.String("defaultUserID"),
			UserName:     pulumi.String("default"),
			AccessString: pulumi.String("on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember"),
			Engine:       pulumi.String("REDIS"),
			Passwords: pulumi.StringArray{
				pulumi.String("password123456789"),
			},
		})
		if err != nil {
			return err
		}
		exampleUserGroup, err := elasticache.NewUserGroup(ctx, "exampleUserGroup", &elasticache.UserGroupArgs{
			Engine:      pulumi.String("REDIS"),
			UserGroupId: pulumi.String("userGroupId"),
			UserIds: pulumi.StringArray{
				_default.UserId,
			},
		})
		if err != nil {
			return err
		}
		exampleUser, err := elasticache.NewUser(ctx, "exampleUser", &elasticache.UserArgs{
			UserId:       pulumi.String("exampleUserID"),
			UserName:     pulumi.String("exampleuser"),
			AccessString: pulumi.String("on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember"),
			Engine:       pulumi.String("REDIS"),
			Passwords: pulumi.StringArray{
				pulumi.String("password123456789"),
			},
		})
		if err != nil {
			return err
		}
		_, err = elasticache.NewUserGroupAssociation(ctx, "exampleUserGroupAssociation", &elasticache.UserGroupAssociationArgs{
			UserGroupId: exampleUserGroup.UserGroupId,
			UserId:      exampleUser.UserId,
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
import com.pulumi.aws.elasticache.User;
import com.pulumi.aws.elasticache.UserArgs;
import com.pulumi.aws.elasticache.UserGroup;
import com.pulumi.aws.elasticache.UserGroupArgs;
import com.pulumi.aws.elasticache.UserGroupAssociation;
import com.pulumi.aws.elasticache.UserGroupAssociationArgs;
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
        var default_ = new User("default", UserArgs.builder()        
            .userId("defaultUserID")
            .userName("default")
            .accessString("on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember")
            .engine("REDIS")
            .passwords("password123456789")
            .build());

        var exampleUserGroup = new UserGroup("exampleUserGroup", UserGroupArgs.builder()        
            .engine("REDIS")
            .userGroupId("userGroupId")
            .userIds(default_.userId())
            .build());

        var exampleUser = new User("exampleUser", UserArgs.builder()        
            .userId("exampleUserID")
            .userName("exampleuser")
            .accessString("on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember")
            .engine("REDIS")
            .passwords("password123456789")
            .build());

        var exampleUserGroupAssociation = new UserGroupAssociation("exampleUserGroupAssociation", UserGroupAssociationArgs.builder()        
            .userGroupId(exampleUserGroup.userGroupId())
            .userId(exampleUser.userId())
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:elasticache:User
    properties:
      userId: defaultUserID
      userName: default
      accessString: on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember
      engine: REDIS
      passwords:
        - password123456789
  exampleUserGroup:
    type: aws:elasticache:UserGroup
    properties:
      engine: REDIS
      userGroupId: userGroupId
      userIds:
        - ${default.userId}
  exampleUser:
    type: aws:elasticache:User
    properties:
      userId: exampleUserID
      userName: exampleuser
      accessString: on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember
      engine: REDIS
      passwords:
        - password123456789
  exampleUserGroupAssociation:
    type: aws:elasticache:UserGroupAssociation
    properties:
      userGroupId: ${exampleUserGroup.userGroupId}
      userId: ${exampleUser.userId}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ElastiCache user group associations using the `user_group_id` and `user_id`. For example:

```sh
 $ pulumi import aws:elasticache/userGroupAssociation:UserGroupAssociation example userGoupId1,userId
```
 