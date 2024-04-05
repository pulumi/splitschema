Resource for managing a Roles Anywhere Profile.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testRole = new aws.iam.Role("testRole", {
    path: "/",
    assumeRolePolicy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Action: [
                "sts:AssumeRole",
                "sts:TagSession",
                "sts:SetSourceIdentity",
            ],
            Principal: {
                Service: "rolesanywhere.amazonaws.com",
            },
            Effect: "Allow",
            Sid: "",
        }],
    }),
});
const testProfile = new aws.rolesanywhere.Profile("testProfile", {roleArns: [testRole.arn]});
```
```python
import pulumi
import json
import pulumi_aws as aws

test_role = aws.iam.Role("testRole",
    path="/",
    assume_role_policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Action": [
                "sts:AssumeRole",
                "sts:TagSession",
                "sts:SetSourceIdentity",
            ],
            "Principal": {
                "Service": "rolesanywhere.amazonaws.com",
            },
            "Effect": "Allow",
            "Sid": "",
        }],
    }))
test_profile = aws.rolesanywhere.Profile("testProfile", role_arns=[test_role.arn])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testRole = new Aws.Iam.Role("testRole", new()
    {
        Path = "/",
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = new[]
                    {
                        "sts:AssumeRole",
                        "sts:TagSession",
                        "sts:SetSourceIdentity",
                    },
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "rolesanywhere.amazonaws.com",
                    },
                    ["Effect"] = "Allow",
                    ["Sid"] = "",
                },
            },
        }),
    });

    var testProfile = new Aws.RolesAnywhere.Profile("testProfile", new()
    {
        RoleArns = new[]
        {
            testRole.Arn,
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rolesanywhere"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": []string{
						"sts:AssumeRole",
						"sts:TagSession",
						"sts:SetSourceIdentity",
					},
					"Principal": map[string]interface{}{
						"Service": "rolesanywhere.amazonaws.com",
					},
					"Effect": "Allow",
					"Sid":    "",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
			Path:             pulumi.String("/"),
			AssumeRolePolicy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = rolesanywhere.NewProfile(ctx, "testProfile", &rolesanywhere.ProfileArgs{
			RoleArns: pulumi.StringArray{
				testRole.Arn,
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.rolesanywhere.Profile;
import com.pulumi.aws.rolesanywhere.ProfileArgs;
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
        var testRole = new Role("testRole", RoleArgs.builder()        
            .path("/")
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", jsonArray(
                            "sts:AssumeRole", 
                            "sts:TagSession", 
                            "sts:SetSourceIdentity"
                        )),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "rolesanywhere.amazonaws.com")
                        )),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Sid", "")
                    )))
                )))
            .build());

        var testProfile = new Profile("testProfile", ProfileArgs.builder()        
            .roleArns(testRole.arn())
            .build());

    }
}
```
```yaml
resources:
  testRole:
    type: aws:iam:Role
    properties:
      path: /
      assumeRolePolicy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action:
                - sts:AssumeRole
                - sts:TagSession
                - sts:SetSourceIdentity
              Principal:
                Service: rolesanywhere.amazonaws.com
              Effect: Allow
              Sid:
  testProfile:
    type: aws:rolesanywhere:Profile
    properties:
      roleArns:
        - ${testRole.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_rolesanywhere_profile` using its `id`. For example:

```sh
 $ pulumi import aws:rolesanywhere/profile:Profile example db138a85-8925-4f9f-a409-08231233cacf
```
 