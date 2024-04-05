Provides a resource to manage an [Amazon Macie Organization Admin Account](https://docs.aws.amazon.com/macie/latest/APIReference/admin.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.macie2.Account("exampleAccount", {});
const exampleOrganizationAdminAccount = new aws.macie2.OrganizationAdminAccount("exampleOrganizationAdminAccount", {adminAccountId: "ID OF THE ADMIN ACCOUNT"}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.macie2.Account("exampleAccount")
example_organization_admin_account = aws.macie2.OrganizationAdminAccount("exampleOrganizationAdminAccount", admin_account_id="ID OF THE ADMIN ACCOUNT",
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

    var exampleOrganizationAdminAccount = new Aws.Macie2.OrganizationAdminAccount("exampleOrganizationAdminAccount", new()
    {
        AdminAccountId = "ID OF THE ADMIN ACCOUNT",
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
		_, err = macie2.NewOrganizationAdminAccount(ctx, "exampleOrganizationAdminAccount", &macie2.OrganizationAdminAccountArgs{
			AdminAccountId: pulumi.String("ID OF THE ADMIN ACCOUNT"),
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
import com.pulumi.aws.macie2.OrganizationAdminAccount;
import com.pulumi.aws.macie2.OrganizationAdminAccountArgs;
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

        var exampleOrganizationAdminAccount = new OrganizationAdminAccount("exampleOrganizationAdminAccount", OrganizationAdminAccountArgs.builder()        
            .adminAccountId("ID OF THE ADMIN ACCOUNT")
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
  exampleOrganizationAdminAccount:
    type: aws:macie2:OrganizationAdminAccount
    properties:
      adminAccountId: ID OF THE ADMIN ACCOUNT
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_macie2_organization_admin_account` using the id. For example:

```sh
 $ pulumi import aws:macie2/organizationAdminAccount:OrganizationAdminAccount example abcd1
```
 