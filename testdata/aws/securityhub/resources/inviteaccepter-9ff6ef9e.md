> **Note:** AWS accounts can only be associated with a single Security Hub master account. Destroying this resource will disassociate the member account from the master account.

Accepts a Security Hub invitation.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const exampleMember = new aws.securityhub.Member("exampleMember", {
    accountId: "123456789012",
    email: "example@example.com",
    invite: true,
});
const inviteeAccount = new aws.securityhub.Account("inviteeAccount", {}, {
    provider: "aws.invitee",
});
const inviteeInviteAccepter = new aws.securityhub.InviteAccepter("inviteeInviteAccepter", {masterId: exampleMember.masterId}, {
    provider: "aws.invitee",
    dependsOn: [inviteeAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_member = aws.securityhub.Member("exampleMember",
    account_id="123456789012",
    email="example@example.com",
    invite=True)
invitee_account = aws.securityhub.Account("inviteeAccount", opts=pulumi.ResourceOptions(provider="aws.invitee"))
invitee_invite_accepter = aws.securityhub.InviteAccepter("inviteeInviteAccepter", master_id=example_member.master_id,
opts=pulumi.ResourceOptions(provider="aws.invitee",
    depends_on=[invitee_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.SecurityHub.Account("exampleAccount");

    var exampleMember = new Aws.SecurityHub.Member("exampleMember", new()
    {
        AccountId = "123456789012",
        Email = "example@example.com",
        Invite = true,
    });

    var inviteeAccount = new Aws.SecurityHub.Account("inviteeAccount", new()
    {
    }, new CustomResourceOptions
    {
        Provider = "aws.invitee",
    });

    var inviteeInviteAccepter = new Aws.SecurityHub.InviteAccepter("inviteeInviteAccepter", new()
    {
        MasterId = exampleMember.MasterId,
    }, new CustomResourceOptions
    {
        Provider = "aws.invitee",
        DependsOn = new[]
        {
            inviteeAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		exampleMember, err := securityhub.NewMember(ctx, "exampleMember", &securityhub.MemberArgs{
			AccountId: pulumi.String("123456789012"),
			Email:     pulumi.String("example@example.com"),
			Invite:    pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		inviteeAccount, err := securityhub.NewAccount(ctx, "inviteeAccount", nil, pulumi.Provider("aws.invitee"))
		if err != nil {
			return err
		}
		_, err = securityhub.NewInviteAccepter(ctx, "inviteeInviteAccepter", &securityhub.InviteAccepterArgs{
			MasterId: exampleMember.MasterId,
		}, pulumi.Provider("aws.invitee"), pulumi.DependsOn([]pulumi.Resource{
			inviteeAccount,
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
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.Member;
import com.pulumi.aws.securityhub.MemberArgs;
import com.pulumi.aws.securityhub.AccountArgs;
import com.pulumi.aws.securityhub.InviteAccepter;
import com.pulumi.aws.securityhub.InviteAccepterArgs;
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
            .accountId("123456789012")
            .email("example@example.com")
            .invite(true)
            .build());

        var inviteeAccount = new Account("inviteeAccount", AccountArgs.Empty, CustomResourceOptions.builder()
            .provider("aws.invitee")
            .build());

        var inviteeInviteAccepter = new InviteAccepter("inviteeInviteAccepter", InviteAccepterArgs.builder()        
            .masterId(exampleMember.masterId())
            .build(), CustomResourceOptions.builder()
                .provider("aws.invitee")
                .dependsOn(inviteeAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:securityhub:Account
  exampleMember:
    type: aws:securityhub:Member
    properties:
      accountId: '123456789012'
      email: example@example.com
      invite: true
  inviteeAccount:
    type: aws:securityhub:Account
    options:
      provider: aws.invitee
  inviteeInviteAccepter:
    type: aws:securityhub:InviteAccepter
    properties:
      masterId: ${exampleMember.masterId}
    options:
      provider: aws.invitee
      dependson:
        - ${inviteeAccount}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Security Hub invite acceptance using the account ID. For example:

```sh
 $ pulumi import aws:securityhub/inviteAccepter:InviteAccepter example 123456789012
```
 