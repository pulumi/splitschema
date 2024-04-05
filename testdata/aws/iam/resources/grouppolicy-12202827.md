Provides an IAM policy attached to a group.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDevelopers = new aws.iam.Group("myDevelopers", {path: "/users/"});
const myDeveloperPolicy = new aws.iam.GroupPolicy("myDeveloperPolicy", {
    group: myDevelopers.name,
    policy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Action: ["ec2:Describe*"],
            Effect: "Allow",
            Resource: "*",
        }],
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

my_developers = aws.iam.Group("myDevelopers", path="/users/")
my_developer_policy = aws.iam.GroupPolicy("myDeveloperPolicy",
    group=my_developers.name,
    policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Action": ["ec2:Describe*"],
            "Effect": "Allow",
            "Resource": "*",
        }],
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myDevelopers = new Aws.Iam.Group("myDevelopers", new()
    {
        Path = "/users/",
    });

    var myDeveloperPolicy = new Aws.Iam.GroupPolicy("myDeveloperPolicy", new()
    {
        Group = myDevelopers.Name,
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
		myDevelopers, err := iam.NewGroup(ctx, "myDevelopers", &iam.GroupArgs{
			Path: pulumi.String("/users/"),
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
		_, err = iam.NewGroupPolicy(ctx, "myDeveloperPolicy", &iam.GroupPolicyArgs{
			Group:  myDevelopers.Name,
			Policy: pulumi.String(json0),
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
import com.pulumi.aws.iam.GroupPolicy;
import com.pulumi.aws.iam.GroupPolicyArgs;
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
        var myDevelopers = new Group("myDevelopers", GroupArgs.builder()        
            .path("/users/")
            .build());

        var myDeveloperPolicy = new GroupPolicy("myDeveloperPolicy", GroupPolicyArgs.builder()        
            .group(myDevelopers.name())
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

    }
}
```
```yaml
resources:
  myDeveloperPolicy:
    type: aws:iam:GroupPolicy
    properties:
      group: ${myDevelopers.name}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action:
                - ec2:Describe*
              Effect: Allow
              Resource: '*'
  myDevelopers:
    type: aws:iam:Group
    properties:
      path: /users/
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM Group Policies using the `group_name:group_policy_name`. For example:

```sh
 $ pulumi import aws:iam/groupPolicy:GroupPolicy mypolicy group_of_mypolicy_name:mypolicy_name
```
 