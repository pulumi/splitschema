Provides an IAM policy attached to a user.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lbUser = new aws.iam.User("lbUser", {path: "/system/"});
const lbRo = new aws.iam.UserPolicy("lbRo", {
    user: lbUser.name,
    policy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Action: ["ec2:Describe*"],
            Effect: "Allow",
            Resource: "*",
        }],
    }),
});
const lbAccessKey = new aws.iam.AccessKey("lbAccessKey", {user: lbUser.name});
```
```python
import pulumi
import json
import pulumi_aws as aws

lb_user = aws.iam.User("lbUser", path="/system/")
lb_ro = aws.iam.UserPolicy("lbRo",
    user=lb_user.name,
    policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Action": ["ec2:Describe*"],
            "Effect": "Allow",
            "Resource": "*",
        }],
    }))
lb_access_key = aws.iam.AccessKey("lbAccessKey", user=lb_user.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var lbUser = new Aws.Iam.User("lbUser", new()
    {
        Path = "/system/",
    });

    var lbRo = new Aws.Iam.UserPolicy("lbRo", new()
    {
        User = lbUser.Name,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = new[]
                    {
                        "ec2:Describe*",
                    },
                    ["Effect"] = "Allow",
                    ["Resource"] = "*",
                },
            },
        }),
    });

    var lbAccessKey = new Aws.Iam.AccessKey("lbAccessKey", new()
    {
        User = lbUser.Name,
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		lbUser, err := iam.NewUser(ctx, "lbUser", &iam.UserArgs{
			Path: pulumi.String("/system/"),
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": []string{
						"ec2:Describe*",
					},
					"Effect":   "Allow",
					"Resource": "*",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = iam.NewUserPolicy(ctx, "lbRo", &iam.UserPolicyArgs{
			User:   lbUser.Name,
			Policy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewAccessKey(ctx, "lbAccessKey", &iam.AccessKeyArgs{
			User: lbUser.Name,
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
import com.pulumi.aws.iam.User;
import com.pulumi.aws.iam.UserArgs;
import com.pulumi.aws.iam.UserPolicy;
import com.pulumi.aws.iam.UserPolicyArgs;
import com.pulumi.aws.iam.AccessKey;
import com.pulumi.aws.iam.AccessKeyArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var lbUser = new User("lbUser", UserArgs.builder()        
            .path("/system/")
            .build());

        var lbRo = new UserPolicy("lbRo", UserPolicyArgs.builder()        
            .user(lbUser.name())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", jsonArray("ec2:Describe*")),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Resource", "*")
                    )))
                )))
            .build());

        var lbAccessKey = new AccessKey("lbAccessKey", AccessKeyArgs.builder()        
            .user(lbUser.name())
            .build());

    }
}
```
```yaml
resources:
  lbRo:
    type: aws:iam:UserPolicy
    properties:
      user: ${lbUser.name}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action:
                - ec2:Describe*
              Effect: Allow
              Resource: '*'
  lbUser:
    type: aws:iam:User
    properties:
      path: /system/
  lbAccessKey:
    type: aws:iam:AccessKey
    properties:
      user: ${lbUser.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM User Policies using the `user_name:user_policy_name`. For example:

```sh
 $ pulumi import aws:iam/userPolicy:UserPolicy mypolicy user_of_mypolicy_name:mypolicy_name
```
 