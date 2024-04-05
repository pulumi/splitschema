Attaches a Managed IAM Policy to an IAM user

> **NOTE:** The usage of this resource conflicts with the `aws.iam.PolicyAttachment` resource and will permanently show a difference if both are defined.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const user = new aws.iam.User("user", {});
const policy = new aws.iam.Policy("policy", {
    description: "A test policy",
    policy: "{ ... policy JSON ... }",
});
const test_attach = new aws.iam.UserPolicyAttachment("test-attach", {
    user: user.name,
    policyArn: policy.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

user = aws.iam.User("user")
policy = aws.iam.Policy("policy",
    description="A test policy",
    policy="{ ... policy JSON ... }")
test_attach = aws.iam.UserPolicyAttachment("test-attach",
    user=user.name,
    policy_arn=policy.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var user = new Aws.Iam.User("user");

    var policy = new Aws.Iam.Policy("policy", new()
    {
        Description = "A test policy",
        PolicyDocument = "{ ... policy JSON ... }",
    });

    var test_attach = new Aws.Iam.UserPolicyAttachment("test-attach", new()
    {
        User = user.Name,
        PolicyArn = policy.Arn,
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
		user, err := iam.NewUser(ctx, "user", nil)
		if err != nil {
			return err
		}
		policy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
			Description: pulumi.String("A test policy"),
			Policy:      pulumi.Any("{ ... policy JSON ... }"),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewUserPolicyAttachment(ctx, "test-attach", &iam.UserPolicyAttachmentArgs{
			User:      user.Name,
			PolicyArn: policy.Arn,
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
import com.pulumi.aws.iam.Policy;
import com.pulumi.aws.iam.PolicyArgs;
import com.pulumi.aws.iam.UserPolicyAttachment;
import com.pulumi.aws.iam.UserPolicyAttachmentArgs;
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
        var user = new User("user");

        var policy = new Policy("policy", PolicyArgs.builder()        
            .description("A test policy")
            .policy("{ ... policy JSON ... }")
            .build());

        var test_attach = new UserPolicyAttachment("test-attach", UserPolicyAttachmentArgs.builder()        
            .user(user.name())
            .policyArn(policy.arn())
            .build());

    }
}
```
```yaml
resources:
  user:
    type: aws:iam:User
  policy:
    type: aws:iam:Policy
    properties:
      description: A test policy
      policy: '{ ... policy JSON ... }'
  test-attach:
    type: aws:iam:UserPolicyAttachment
    properties:
      user: ${user.name}
      policyArn: ${policy.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM user policy attachments using the user name and policy arn separated by `/`. For example:

```sh
 $ pulumi import aws:iam/userPolicyAttachment:UserPolicyAttachment test-attach test-user/arn:aws:iam::xxxxxxxxxxxx:policy/test-policy
```
 