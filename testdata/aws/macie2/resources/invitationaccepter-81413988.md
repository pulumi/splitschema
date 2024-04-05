Provides a resource to manage an [Amazon Macie Invitation Accepter](https://docs.aws.amazon.com/macie/latest/APIReference/invitations-accept.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primaryAccount = new aws.macie2.Account("primaryAccount", {}, {
    provider: "awsalternate",
});
const memberAccount = new aws.macie2.Account("memberAccount", {});
const primaryMember = new aws.macie2.Member("primaryMember", {
    accountId: "ACCOUNT ID",
    email: "EMAIL",
    invite: true,
    invitationMessage: "Message of the invite",
}, {
    provider: "awsalternate",
    dependsOn: [primaryAccount],
});
const memberInvitationAccepter = new aws.macie2.InvitationAccepter("memberInvitationAccepter", {administratorAccountId: "ADMINISTRATOR ACCOUNT ID"}, {
    dependsOn: [primaryMember],
});
```
```python
import pulumi
import pulumi_aws as aws

primary_account = aws.macie2.Account("primaryAccount", opts=pulumi.ResourceOptions(provider="awsalternate"))
member_account = aws.macie2.Account("memberAccount")
primary_member = aws.macie2.Member("primaryMember",
    account_id="ACCOUNT ID",
    email="EMAIL",
    invite=True,
    invitation_message="Message of the invite",
    opts=pulumi.ResourceOptions(provider="awsalternate",
        depends_on=[primary_account]))
member_invitation_accepter = aws.macie2.InvitationAccepter("memberInvitationAccepter", administrator_account_id="ADMINISTRATOR ACCOUNT ID",
opts=pulumi.ResourceOptions(depends_on=[primary_member]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var primaryAccount = new Aws.Macie2.Account("primaryAccount", new()
    {
    }, new CustomResourceOptions
    {
        Provider = "awsalternate",
    });

    var memberAccount = new Aws.Macie2.Account("memberAccount");

    var primaryMember = new Aws.Macie2.Member("primaryMember", new()
    {
        AccountId = "ACCOUNT ID",
        Email = "EMAIL",
        Invite = true,
        InvitationMessage = "Message of the invite",
    }, new CustomResourceOptions
    {
        Provider = "awsalternate",
        DependsOn = new[]
        {
            primaryAccount,
        },
    });

    var memberInvitationAccepter = new Aws.Macie2.InvitationAccepter("memberInvitationAccepter", new()
    {
        AdministratorAccountId = "ADMINISTRATOR ACCOUNT ID",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            primaryMember,
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
		primaryAccount, err := macie2.NewAccount(ctx, "primaryAccount", nil, pulumi.Provider("awsalternate"))
		if err != nil {
			return err
		}
		_, err = macie2.NewAccount(ctx, "memberAccount", nil)
		if err != nil {
			return err
		}
		primaryMember, err := macie2.NewMember(ctx, "primaryMember", &macie2.MemberArgs{
			AccountId:         pulumi.String("ACCOUNT ID"),
			Email:             pulumi.String("EMAIL"),
			Invite:            pulumi.Bool(true),
			InvitationMessage: pulumi.String("Message of the invite"),
		}, pulumi.Provider("awsalternate"), pulumi.DependsOn([]pulumi.Resource{
			primaryAccount,
		}))
		if err != nil {
			return err
		}
		_, err = macie2.NewInvitationAccepter(ctx, "memberInvitationAccepter", &macie2.InvitationAccepterArgs{
			AdministratorAccountId: pulumi.String("ADMINISTRATOR ACCOUNT ID"),
		}, pulumi.DependsOn([]pulumi.Resource{
			primaryMember,
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
import com.pulumi.aws.macie2.AccountArgs;
import com.pulumi.aws.macie2.Member;
import com.pulumi.aws.macie2.MemberArgs;
import com.pulumi.aws.macie2.InvitationAccepter;
import com.pulumi.aws.macie2.InvitationAccepterArgs;
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
        var primaryAccount = new Account("primaryAccount", AccountArgs.Empty, CustomResourceOptions.builder()
            .provider("awsalternate")
            .build());

        var memberAccount = new Account("memberAccount");

        var primaryMember = new Member("primaryMember", MemberArgs.builder()        
            .accountId("ACCOUNT ID")
            .email("EMAIL")
            .invite(true)
            .invitationMessage("Message of the invite")
            .build(), CustomResourceOptions.builder()
                .provider("awsalternate")
                .dependsOn(primaryAccount)
                .build());

        var memberInvitationAccepter = new InvitationAccepter("memberInvitationAccepter", InvitationAccepterArgs.builder()        
            .administratorAccountId("ADMINISTRATOR ACCOUNT ID")
            .build(), CustomResourceOptions.builder()
                .dependsOn(primaryMember)
                .build());

    }
}
```
```yaml
resources:
  primaryAccount:
    type: aws:macie2:Account
    options:
      provider: awsalternate
  memberAccount:
    type: aws:macie2:Account
  primaryMember:
    type: aws:macie2:Member
    properties:
      accountId: ACCOUNT ID
      email: EMAIL
      invite: true
      invitationMessage: Message of the invite
    options:
      provider: awsalternate
      dependson:
        - ${primaryAccount}
  memberInvitationAccepter:
    type: aws:macie2:InvitationAccepter
    properties:
      administratorAccountId: ADMINISTRATOR ACCOUNT ID
    options:
      dependson:
        - ${primaryMember}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_macie2_invitation_accepter` using the admin account ID. For example:

```sh
 $ pulumi import aws:macie2/invitationAccepter:InvitationAccepter example 123456789012
```
 