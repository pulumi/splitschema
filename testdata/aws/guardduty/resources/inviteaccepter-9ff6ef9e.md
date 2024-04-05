Provides a resource to accept a pending GuardDuty invite on creation, ensure the detector has the correct primary account on read, and disassociate with the primary account upon removal.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.Provider("primary", {});
const member = new aws.Provider("member", {});
const primaryDetector = new aws.guardduty.Detector("primaryDetector", {}, {
    provider: aws.primary,
});
const memberDetector = new aws.guardduty.Detector("memberDetector", {}, {
    provider: aws.member,
});
const memberMember = new aws.guardduty.Member("memberMember", {
    accountId: memberDetector.accountId,
    detectorId: primaryDetector.id,
    email: "required@example.com",
    invite: true,
}, {
    provider: aws.primary,
});
const memberInviteAccepter = new aws.guardduty.InviteAccepter("memberInviteAccepter", {
    detectorId: memberDetector.id,
    masterAccountId: primaryDetector.accountId,
}, {
    provider: aws.member,
    dependsOn: [memberMember],
});
```
```python
import pulumi
import pulumi_aws as aws

primary = aws.Provider("primary")
member = aws.Provider("member")
primary_detector = aws.guardduty.Detector("primaryDetector", opts=pulumi.ResourceOptions(provider=aws["primary"]))
member_detector = aws.guardduty.Detector("memberDetector", opts=pulumi.ResourceOptions(provider=aws["member"]))
member_member = aws.guardduty.Member("memberMember",
    account_id=member_detector.account_id,
    detector_id=primary_detector.id,
    email="required@example.com",
    invite=True,
    opts=pulumi.ResourceOptions(provider=aws["primary"]))
member_invite_accepter = aws.guardduty.InviteAccepter("memberInviteAccepter",
    detector_id=member_detector.id,
    master_account_id=primary_detector.account_id,
    opts=pulumi.ResourceOptions(provider=aws["member"],
        depends_on=[member_member]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var primary = new Aws.Provider("primary");

    var member = new Aws.Provider("member");

    var primaryDetector = new Aws.GuardDuty.Detector("primaryDetector", new()
    {
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var memberDetector = new Aws.GuardDuty.Detector("memberDetector", new()
    {
    }, new CustomResourceOptions
    {
        Provider = aws.Member,
    });

    var memberMember = new Aws.GuardDuty.Member("memberMember", new()
    {
        AccountId = memberDetector.AccountId,
        DetectorId = primaryDetector.Id,
        Email = "required@example.com",
        Invite = true,
    }, new CustomResourceOptions
    {
        Provider = aws.Primary,
    });

    var memberInviteAccepter = new Aws.GuardDuty.InviteAccepter("memberInviteAccepter", new()
    {
        DetectorId = memberDetector.Id,
        MasterAccountId = primaryDetector.AccountId,
    }, new CustomResourceOptions
    {
        Provider = aws.Member,
        DependsOn = new[]
        {
            memberMember,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/guardduty"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "primary", nil)
		if err != nil {
			return err
		}
		_, err = aws.NewProvider(ctx, "member", nil)
		if err != nil {
			return err
		}
		primaryDetector, err := guardduty.NewDetector(ctx, "primaryDetector", nil, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		memberDetector, err := guardduty.NewDetector(ctx, "memberDetector", nil, pulumi.Provider(aws.Member))
		if err != nil {
			return err
		}
		memberMember, err := guardduty.NewMember(ctx, "memberMember", &guardduty.MemberArgs{
			AccountId:  memberDetector.AccountId,
			DetectorId: primaryDetector.ID(),
			Email:      pulumi.String("required@example.com"),
			Invite:     pulumi.Bool(true),
		}, pulumi.Provider(aws.Primary))
		if err != nil {
			return err
		}
		_, err = guardduty.NewInviteAccepter(ctx, "memberInviteAccepter", &guardduty.InviteAccepterArgs{
			DetectorId:      memberDetector.ID(),
			MasterAccountId: primaryDetector.AccountId,
		}, pulumi.Provider(aws.Member), pulumi.DependsOn([]pulumi.Resource{
			memberMember,
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
import com.pulumi.aws.Provider;
import com.pulumi.aws.guardduty.Detector;
import com.pulumi.aws.guardduty.DetectorArgs;
import com.pulumi.aws.guardduty.Member;
import com.pulumi.aws.guardduty.MemberArgs;
import com.pulumi.aws.guardduty.InviteAccepter;
import com.pulumi.aws.guardduty.InviteAccepterArgs;
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
        var primary = new Provider("primary");

        var member = new Provider("member");

        var primaryDetector = new Detector("primaryDetector", DetectorArgs.Empty, CustomResourceOptions.builder()
            .provider(aws.primary())
            .build());

        var memberDetector = new Detector("memberDetector", DetectorArgs.Empty, CustomResourceOptions.builder()
            .provider(aws.member())
            .build());

        var memberMember = new Member("memberMember", MemberArgs.builder()        
            .accountId(memberDetector.accountId())
            .detectorId(primaryDetector.id())
            .email("required@example.com")
            .invite(true)
            .build(), CustomResourceOptions.builder()
                .provider(aws.primary())
                .build());

        var memberInviteAccepter = new InviteAccepter("memberInviteAccepter", InviteAccepterArgs.builder()        
            .detectorId(memberDetector.id())
            .masterAccountId(primaryDetector.accountId())
            .build(), CustomResourceOptions.builder()
                .provider(aws.member())
                .dependsOn(memberMember)
                .build());

    }
}
```
```yaml
resources:
  primary:
    type: pulumi:providers:aws
  member:
    type: pulumi:providers:aws
  memberInviteAccepter:
    type: aws:guardduty:InviteAccepter
    properties:
      detectorId: ${memberDetector.id}
      masterAccountId: ${primaryDetector.accountId}
    options:
      provider: ${aws.member}
      dependson:
        - ${memberMember}
  memberMember:
    type: aws:guardduty:Member
    properties:
      accountId: ${memberDetector.accountId}
      detectorId: ${primaryDetector.id}
      email: required@example.com
      invite: true
    options:
      provider: ${aws.primary}
  primaryDetector:
    type: aws:guardduty:Detector
    options:
      provider: ${aws.primary}
  memberDetector:
    type: aws:guardduty:Detector
    options:
      provider: ${aws.member}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_guardduty_invite_accepter` using the member GuardDuty detector ID. For example:

```sh
 $ pulumi import aws:guardduty/inviteAccepter:InviteAccepter member 00b00fd5aecc0ab60a708659477e9617
```
 