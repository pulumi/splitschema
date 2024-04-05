Provides a resource to manage a GuardDuty member. To accept invitations in member accounts, see the `aws.guardduty.InviteAccepter` resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.guardduty.Detector("primary", {enable: true});
const memberDetector = new aws.guardduty.Detector("memberDetector", {enable: true}, {
    provider: aws.dev,
});
const memberMember = new aws.guardduty.Member("memberMember", {
    accountId: memberDetector.accountId,
    detectorId: primary.id,
    email: "required@example.com",
    invite: true,
    invitationMessage: "please accept guardduty invitation",
});
```
```python
import pulumi
import pulumi_aws as aws

primary = aws.guardduty.Detector("primary", enable=True)
member_detector = aws.guardduty.Detector("memberDetector", enable=True,
opts=pulumi.ResourceOptions(provider=aws["dev"]))
member_member = aws.guardduty.Member("memberMember",
    account_id=member_detector.account_id,
    detector_id=primary.id,
    email="required@example.com",
    invite=True,
    invitation_message="please accept guardduty invitation")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var primary = new Aws.GuardDuty.Detector("primary", new()
    {
        Enable = true,
    });

    var memberDetector = new Aws.GuardDuty.Detector("memberDetector", new()
    {
        Enable = true,
    }, new CustomResourceOptions
    {
        Provider = aws.Dev,
    });

    var memberMember = new Aws.GuardDuty.Member("memberMember", new()
    {
        AccountId = memberDetector.AccountId,
        DetectorId = primary.Id,
        Email = "required@example.com",
        Invite = true,
        InvitationMessage = "please accept guardduty invitation",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/guardduty"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		primary, err := guardduty.NewDetector(ctx, "primary", &guardduty.DetectorArgs{
			Enable: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		memberDetector, err := guardduty.NewDetector(ctx, "memberDetector", &guardduty.DetectorArgs{
			Enable: pulumi.Bool(true),
		}, pulumi.Provider(aws.Dev))
		if err != nil {
			return err
		}
		_, err = guardduty.NewMember(ctx, "memberMember", &guardduty.MemberArgs{
			AccountId:         memberDetector.AccountId,
			DetectorId:        primary.ID(),
			Email:             pulumi.String("required@example.com"),
			Invite:            pulumi.Bool(true),
			InvitationMessage: pulumi.String("please accept guardduty invitation"),
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
import com.pulumi.aws.guardduty.Detector;
import com.pulumi.aws.guardduty.DetectorArgs;
import com.pulumi.aws.guardduty.Member;
import com.pulumi.aws.guardduty.MemberArgs;
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
        var primary = new Detector("primary", DetectorArgs.builder()        
            .enable(true)
            .build());

        var memberDetector = new Detector("memberDetector", DetectorArgs.builder()        
            .enable(true)
            .build(), CustomResourceOptions.builder()
                .provider(aws.dev())
                .build());

        var memberMember = new Member("memberMember", MemberArgs.builder()        
            .accountId(memberDetector.accountId())
            .detectorId(primary.id())
            .email("required@example.com")
            .invite(true)
            .invitationMessage("please accept guardduty invitation")
            .build());

    }
}
```
```yaml
resources:
  primary:
    type: aws:guardduty:Detector
    properties:
      enable: true
  memberDetector:
    type: aws:guardduty:Detector
    properties:
      enable: true
    options:
      provider: ${aws.dev}
  memberMember:
    type: aws:guardduty:Member
    properties:
      accountId: ${memberDetector.accountId}
      detectorId: ${primary.id}
      email: required@example.com
      invite: true
      invitationMessage: please accept guardduty invitation
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import GuardDuty members using the primary GuardDuty detector ID and member AWS account ID. For example:

```sh
 $ pulumi import aws:guardduty/member:Member MyMember 00b00fd5aecc0ab60a708659477e9617:123456789012
```
 