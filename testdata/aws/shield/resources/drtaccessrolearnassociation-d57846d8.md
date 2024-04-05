Authorizes the Shield Response Team (SRT) using the specified role, to access your AWS account to assist with DDoS attack mitigation during potential attacks. For more information see [Configure AWS SRT Support](https://docs.aws.amazon.com/waf/latest/developerguide/authorize-srt.html)

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testRole = new aws.iam.Role("testRole", {assumeRolePolicy: JSON.stringify({
    Version: "2012-10-17",
    Statement: [{
        Sid: "",
        Effect: "Allow",
        Principal: {
            Service: "drt.shield.amazonaws.com",
        },
        Action: "sts:AssumeRole",
    }],
})});
const testRolePolicyAttachment = new aws.iam.RolePolicyAttachment("testRolePolicyAttachment", {
    role: testRole.name,
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy",
});
const testDrtAccessRoleArnAssociation = new aws.shield.DrtAccessRoleArnAssociation("testDrtAccessRoleArnAssociation", {roleArn: testRole.arn});
```
```python
import pulumi
import json
import pulumi_aws as aws

test_role = aws.iam.Role("testRole", assume_role_policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Sid": "",
        "Effect": "Allow",
        "Principal": {
            "Service": "drt.shield.amazonaws.com",
        },
        "Action": "sts:AssumeRole",
    }],
}))
test_role_policy_attachment = aws.iam.RolePolicyAttachment("testRolePolicyAttachment",
    role=test_role.name,
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy")
test_drt_access_role_arn_association = aws.shield.DrtAccessRoleArnAssociation("testDrtAccessRoleArnAssociation", role_arn=test_role.arn)
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
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Sid"] = "",
                    ["Effect"] = "Allow",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "drt.shield.amazonaws.com",
                    },
                    ["Action"] = "sts:AssumeRole",
                },
            },
        }),
    });

    var testRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("testRolePolicyAttachment", new()
    {
        Role = testRole.Name,
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy",
    });

    var testDrtAccessRoleArnAssociation = new Aws.Shield.DrtAccessRoleArnAssociation("testDrtAccessRoleArnAssociation", new()
    {
        RoleArn = testRole.Arn,
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/shield"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Sid":    "",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "drt.shield.amazonaws.com",
					},
					"Action": "sts:AssumeRole",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "testRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      testRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy"),
		})
		if err != nil {
			return err
		}
		_, err = shield.NewDrtAccessRoleArnAssociation(ctx, "testDrtAccessRoleArnAssociation", &shield.DrtAccessRoleArnAssociationArgs{
			RoleArn: testRole.Arn,
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
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.shield.DrtAccessRoleArnAssociation;
import com.pulumi.aws.shield.DrtAccessRoleArnAssociationArgs;
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
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Sid", ""),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "drt.shield.amazonaws.com")
                        )),
                        jsonProperty("Action", "sts:AssumeRole")
                    )))
                )))
            .build());

        var testRolePolicyAttachment = new RolePolicyAttachment("testRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .role(testRole.name())
            .policyArn("arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy")
            .build());

        var testDrtAccessRoleArnAssociation = new DrtAccessRoleArnAssociation("testDrtAccessRoleArnAssociation", DrtAccessRoleArnAssociationArgs.builder()        
            .roleArn(testRole.arn())
            .build());

    }
}
```
```yaml
resources:
  testRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Sid:
              Effect: Allow
              Principal:
                Service: drt.shield.amazonaws.com
              Action: sts:AssumeRole
  testRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${testRole.name}
      policyArn: arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy
  testDrtAccessRoleArnAssociation:
    type: aws:shield:DrtAccessRoleArnAssociation
    properties:
      roleArn: ${testRole.arn}
```
{{% /example %}}
{{% /examples %}}