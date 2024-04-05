Provides a resource to manage an [Amazon Macie Member](https://docs.aws.amazon.com/macie/latest/APIReference/members-id.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.macie2.Account("exampleAccount", {});
const exampleMember = new aws.macie2.Member("exampleMember", {
    accountId: "AWS ACCOUNT ID",
    email: "EMAIL",
    invite: true,
    invitationMessage: "Message of the invitation",
    invitationDisableEmailNotification: true,
}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.macie2.Account("exampleAccount")
example_member = aws.macie2.Member("exampleMember",
    account_id="AWS ACCOUNT ID",
    email="EMAIL",
    invite=True,
    invitation_message="Message of the invitation",
    invitation_disable_email_notification=True,
    opts=pulumi.ResourceOptions(depends_on=[example_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.Macie2.Account("exampleAccount");

    var exampleMember = new Aws.Macie2.Member("exampleMember", new()
    {
        AccountId = "AWS ACCOUNT ID",
        Email = "EMAIL",
        Invite = true,
        InvitationMessage = "Message of the invitation",
        InvitationDisableEmailNotification = true,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/macie2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := macie2.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = macie2.NewMember(ctx, "exampleMember", &macie2.MemberArgs{
			AccountId:                          pulumi.String("AWS ACCOUNT ID"),
			Email:                              pulumi.String("EMAIL"),
			Invite:                             pulumi.Bool(true),
			InvitationMessage:                  pulumi.String("Message of the invitation"),
			InvitationDisableEmailNotification: pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccount,
		}))
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
import com.pulumi.aws.macie2.Account;
import com.pulumi.aws.macie2.Member;
import com.pulumi.aws.macie2.MemberArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleMember = new Member("exampleMember", MemberArgs.builder()        
            .accountId("AWS ACCOUNT ID")
            .email("EMAIL")
            .invite(true)
            .invitationMessage("Message of the invitation")
            .invitationDisableEmailNotification(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:macie2:Account
  exampleMember:
    type: aws:macie2:Member
    properties:
      accountId: AWS ACCOUNT ID
      email: EMAIL
      invite: true
      invitationMessage: Message of the invitation
      invitationDisableEmailNotification: true
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_macie2_member` using the account ID of the member account. For example:

```sh
 $ pulumi import aws:macie2/member:Member example 123456789012
```
 