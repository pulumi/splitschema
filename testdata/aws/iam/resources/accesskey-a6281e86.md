Provides an IAM access key. This is a set of credentials that allow API requests to be made as an IAM user.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lbUser = new aws.iam.User("lbUser", {path: "/system/"});
const lbAccessKey = new aws.iam.AccessKey("lbAccessKey", {
    user: lbUser.name,
    pgpKey: "keybase:some_person_that_exists",
});
const lbRoPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: ["ec2:Describe*"],
        resources: ["*"],
    }],
});
const lbRoUserPolicy = new aws.iam.UserPolicy("lbRoUserPolicy", {
    user: lbUser.name,
    policy: lbRoPolicyDocument.then(lbRoPolicyDocument => lbRoPolicyDocument.json),
});
export const secret = lbAccessKey.encryptedSecret;
```
```python
import pulumi
import pulumi_aws as aws

lb_user = aws.iam.User("lbUser", path="/system/")
lb_access_key = aws.iam.AccessKey("lbAccessKey",
    user=lb_user.name,
    pgp_key="keybase:some_person_that_exists")
lb_ro_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["ec2:Describe*"],
    resources=["*"],
)])
lb_ro_user_policy = aws.iam.UserPolicy("lbRoUserPolicy",
    user=lb_user.name,
    policy=lb_ro_policy_document.json)
pulumi.export("secret", lb_access_key.encrypted_secret)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var lbUser = new Aws.Iam.User("lbUser", new()
    {
        Path = "/system/",
    });

    var lbAccessKey = new Aws.Iam.AccessKey("lbAccessKey", new()
    {
        User = lbUser.Name,
        PgpKey = "keybase:some_person_that_exists",
    });

    var lbRoPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "ec2:Describe*",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var lbRoUserPolicy = new Aws.Iam.UserPolicy("lbRoUserPolicy", new()
    {
        User = lbUser.Name,
        Policy = lbRoPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    return new Dictionary<string, object?>
    {
        ["secret"] = lbAccessKey.EncryptedSecret,
    };
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
		lbUser, err := iam.NewUser(ctx, "lbUser", &iam.UserArgs{
			Path: pulumi.String("/system/"),
		})
		if err != nil {
			return err
		}
		lbAccessKey, err := iam.NewAccessKey(ctx, "lbAccessKey", &iam.AccessKeyArgs{
			User:   lbUser.Name,
			PgpKey: pulumi.String("keybase:some_person_that_exists"),
		})
		if err != nil {
			return err
		}
		lbRoPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"ec2:Describe*",
					},
					Resources: []string{
						"*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewUserPolicy(ctx, "lbRoUserPolicy", &iam.UserPolicyArgs{
			User:   lbUser.Name,
			Policy: *pulumi.String(lbRoPolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		ctx.Export("secret", lbAccessKey.EncryptedSecret)
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
import com.pulumi.aws.iam.AccessKey;
import com.pulumi.aws.iam.AccessKeyArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.UserPolicy;
import com.pulumi.aws.iam.UserPolicyArgs;
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

        var lbAccessKey = new AccessKey("lbAccessKey", AccessKeyArgs.builder()        
            .user(lbUser.name())
            .pgpKey("keybase:some_person_that_exists")
            .build());

        final var lbRoPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("ec2:Describe*")
                .resources("*")
                .build())
            .build());

        var lbRoUserPolicy = new UserPolicy("lbRoUserPolicy", UserPolicyArgs.builder()        
            .user(lbUser.name())
            .policy(lbRoPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        ctx.export("secret", lbAccessKey.encryptedSecret());
    }
}
```
```yaml
resources:
  lbAccessKey:
    type: aws:iam:AccessKey
    properties:
      user: ${lbUser.name}
      pgpKey: keybase:some_person_that_exists
  lbUser:
    type: aws:iam:User
    properties:
      path: /system/
  lbRoUserPolicy:
    type: aws:iam:UserPolicy
    properties:
      user: ${lbUser.name}
      policy: ${lbRoPolicyDocument.json}
variables:
  lbRoPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - ec2:Describe*
            resources:
              - '*'
outputs:
  secret: ${lbAccessKey.encryptedSecret}
```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testUser = new aws.iam.User("testUser", {path: "/test/"});
const testAccessKey = new aws.iam.AccessKey("testAccessKey", {user: testUser.name});
export const awsIamSmtpPasswordV4 = testAccessKey.sesSmtpPasswordV4;
```
```python
import pulumi
import pulumi_aws as aws

test_user = aws.iam.User("testUser", path="/test/")
test_access_key = aws.iam.AccessKey("testAccessKey", user=test_user.name)
pulumi.export("awsIamSmtpPasswordV4", test_access_key.ses_smtp_password_v4)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testUser = new Aws.Iam.User("testUser", new()
    {
        Path = "/test/",
    });

    var testAccessKey = new Aws.Iam.AccessKey("testAccessKey", new()
    {
        User = testUser.Name,
    });

    return new Dictionary<string, object?>
    {
        ["awsIamSmtpPasswordV4"] = testAccessKey.SesSmtpPasswordV4,
    };
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
		testUser, err := iam.NewUser(ctx, "testUser", &iam.UserArgs{
			Path: pulumi.String("/test/"),
		})
		if err != nil {
			return err
		}
		testAccessKey, err := iam.NewAccessKey(ctx, "testAccessKey", &iam.AccessKeyArgs{
			User: testUser.Name,
		})
		if err != nil {
			return err
		}
		ctx.Export("awsIamSmtpPasswordV4", testAccessKey.SesSmtpPasswordV4)
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
import com.pulumi.aws.iam.AccessKey;
import com.pulumi.aws.iam.AccessKeyArgs;
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
            .path("/test/")
            .build());

        var testAccessKey = new AccessKey("testAccessKey", AccessKeyArgs.builder()        
            .user(testUser.name())
            .build());

        ctx.export("awsIamSmtpPasswordV4", testAccessKey.sesSmtpPasswordV4());
    }
}
```
```yaml
resources:
  testUser:
    type: aws:iam:User
    properties:
      path: /test/
  testAccessKey:
    type: aws:iam:AccessKey
    properties:
      user: ${testUser.name}
outputs:
  awsIamSmtpPasswordV4: ${testAccessKey.sesSmtpPasswordV4}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM Access Keys using the identifier. For example:

```sh
 $ pulumi import aws:iam/accessKey:AccessKey example AKIA1234567890
```
 Resource attributes such as `encrypted_secret`, `key_fingerprint`, `pgp_key`, `secret`, `ses_smtp_password_v4`, and `encrypted_ses_smtp_password_v4` are not available for imported resources as this information cannot be read from the IAM API.

