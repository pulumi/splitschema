Provides a Security Hub member resource.

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
}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_member = aws.securityhub.Member("exampleMember",
    account_id="123456789012",
    email="example@example.com",
    invite=True,
    opts=pulumi.ResourceOptions(depends_on=[example_account]))
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewMember(ctx, "exampleMember", &securityhub.MemberArgs{
			AccountId: pulumi.String("123456789012"),
			Email:     pulumi.String("example@example.com"),
			Invite:    pulumi.Bool(true),
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
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.Member;
import com.pulumi.aws.securityhub.MemberArgs;
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
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
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
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Security Hub members using their account ID. For example:

```sh
 $ pulumi import aws:securityhub/member:Member example 123456789012
```
 