Resource for managing an AWS FinSpace Kx User.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {
    description: "Example KMS Key",
    deletionWindowInDays: 7,
});
const exampleKxEnvironment = new aws.finspace.KxEnvironment("exampleKxEnvironment", {kmsKeyId: exampleKey.arn});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: JSON.stringify({
    Version: "2012-10-17",
    Statement: [{
        Action: "sts:AssumeRole",
        Effect: "Allow",
        Sid: "",
        Principal: {
            Service: "ec2.amazonaws.com",
        },
    }],
})});
const exampleKxUser = new aws.finspace.KxUser("exampleKxUser", {
    environmentId: exampleKxEnvironment.id,
    iamRole: exampleRole.arn,
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey",
    description="Example KMS Key",
    deletion_window_in_days=7)
example_kx_environment = aws.finspace.KxEnvironment("exampleKxEnvironment", kms_key_id=example_key.arn)
example_role = aws.iam.Role("exampleRole", assume_role_policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Sid": "",
        "Principal": {
            "Service": "ec2.amazonaws.com",
        },
    }],
}))
example_kx_user = aws.finspace.KxUser("exampleKxUser",
    environment_id=example_kx_environment.id,
    iam_role=example_role.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "Example KMS Key",
        DeletionWindowInDays = 7,
    });

    var exampleKxEnvironment = new Aws.FinSpace.KxEnvironment("exampleKxEnvironment", new()
    {
        KmsKeyId = exampleKey.Arn,
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = "sts:AssumeRole",
                    ["Effect"] = "Allow",
                    ["Sid"] = "",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "ec2.amazonaws.com",
                    },
                },
            },
        }),
    });

    var exampleKxUser = new Aws.FinSpace.KxUser("exampleKxUser", new()
    {
        EnvironmentId = exampleKxEnvironment.Id,
        IamRole = exampleRole.Arn,
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/finspace"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description:          pulumi.String("Example KMS Key"),
			DeletionWindowInDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		exampleKxEnvironment, err := finspace.NewKxEnvironment(ctx, "exampleKxEnvironment", &finspace.KxEnvironmentArgs{
			KmsKeyId: exampleKey.Arn,
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Sid":    "",
					"Principal": map[string]interface{}{
						"Service": "ec2.amazonaws.com",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = finspace.NewKxUser(ctx, "exampleKxUser", &finspace.KxUserArgs{
			EnvironmentId: exampleKxEnvironment.ID(),
			IamRole:       exampleRole.Arn,
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.finspace.KxEnvironment;
import com.pulumi.aws.finspace.KxEnvironmentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.finspace.KxUser;
import com.pulumi.aws.finspace.KxUserArgs;
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
        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("Example KMS Key")
            .deletionWindowInDays(7)
            .build());

        var exampleKxEnvironment = new KxEnvironment("exampleKxEnvironment", KxEnvironmentArgs.builder()        
            .kmsKeyId(exampleKey.arn())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", "sts:AssumeRole"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Sid", ""),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "ec2.amazonaws.com")
                        ))
                    )))
                )))
            .build());

        var exampleKxUser = new KxUser("exampleKxUser", KxUserArgs.builder()        
            .environmentId(exampleKxEnvironment.id())
            .iamRole(exampleRole.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: Example KMS Key
      deletionWindowInDays: 7
  exampleKxEnvironment:
    type: aws:finspace:KxEnvironment
    properties:
      kmsKeyId: ${exampleKey.arn}
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action: sts:AssumeRole
              Effect: Allow
              Sid:
              Principal:
                Service: ec2.amazonaws.com
  exampleKxUser:
    type: aws:finspace:KxUser
    properties:
      environmentId: ${exampleKxEnvironment.id}
      iamRole: ${exampleRole.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an AWS FinSpace Kx User using the `id` (environment ID and user name, comma-delimited). For example:

```sh
 $ pulumi import aws:finspace/kxUser:KxUser example n3ceo7wqxoxcti5tujqwzs,my-tf-kx-user
```
 