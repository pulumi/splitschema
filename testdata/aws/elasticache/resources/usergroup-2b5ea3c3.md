Provides an ElastiCache user group resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testUser = new aws.elasticache.User("testUser", {
    userId: "testUserId",
    userName: "default",
    accessString: "on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
    engine: "REDIS",
    passwords: ["password123456789"],
});
const testUserGroup = new aws.elasticache.UserGroup("testUserGroup", {
    engine: "REDIS",
    userGroupId: "userGroupId",
    userIds: [testUser.userId],
});
```
```python
import pulumi
import pulumi_aws as aws

test_user = aws.elasticache.User("testUser",
    user_id="testUserId",
    user_name="default",
    access_string="on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
    engine="REDIS",
    passwords=["password123456789"])
test_user_group = aws.elasticache.UserGroup("testUserGroup",
    engine="REDIS",
    user_group_id="userGroupId",
    user_ids=[test_user.user_id])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testUser = new Aws.ElastiCache.User("testUser", new()
    {
        UserId = "testUserId",
        UserName = "default",
        AccessString = "on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember",
        Engine = "REDIS",
        Passwords = new[]
        {
            "password123456789",
        },
    });

    var testUserGroup = new Aws.ElastiCache.UserGroup("testUserGroup", new()
    {
        Engine = "REDIS",
        UserGroupId = "userGroupId",
        UserIds = new[]
        {
            testUser.UserId,
        },
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
		testUser, err := elasticache.NewUser(ctx, "testUser", &elasticache.UserArgs{
			UserId:       pulumi.String("testUserId"),
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
		_, err = elasticache.NewUserGroup(ctx, "testUserGroup", &elasticache.UserGroupArgs{
			Engine:      pulumi.String("REDIS"),
			UserGroupId: pulumi.String("userGroupId"),
			UserIds: pulumi.StringArray{
				testUser.UserId,
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
import com.pulumi.aws.elasticache.User;
import com.pulumi.aws.elasticache.UserArgs;
import com.pulumi.aws.elasticache.UserGroup;
import com.pulumi.aws.elasticache.UserGroupArgs;
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
        var testUser = new User("testUser", UserArgs.builder()        
            .userId("testUserId")
            .userName("default")
            .accessString("on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember")
            .engine("REDIS")
            .passwords("password123456789")
            .build());

        var testUserGroup = new UserGroup("testUserGroup", UserGroupArgs.builder()        
            .engine("REDIS")
            .userGroupId("userGroupId")
            .userIds(testUser.userId())
            .build());

    }
}
```
```yaml
resources:
  testUser:
    type: aws:elasticache:User
    properties:
      userId: testUserId
      userName: default
      accessString: on ~app::* -@all +@read +@hash +@bitmap +@geo -setbit -bitfield -hset -hsetnx -hmset -hincrby -hincrbyfloat -hdel -bitop -geoadd -georadius -georadiusbymember
      engine: REDIS
      passwords:
        - password123456789
  testUserGroup:
    type: aws:elasticache:UserGroup
    properties:
      engine: REDIS
      userGroupId: userGroupId
      userIds:
        - ${testUser.userId}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ElastiCache user groups using the `user_group_id`. For example:

```sh
 $ pulumi import aws:elasticache/userGroup:UserGroup my_user_group userGoupId1
```
 