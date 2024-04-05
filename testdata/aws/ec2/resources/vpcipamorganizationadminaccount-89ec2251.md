Enables the IPAM Service and promotes a delegated administrator.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const delegated = aws.getCallerIdentity({});
const example = new aws.ec2.VpcIpamOrganizationAdminAccount("example", {delegatedAdminAccountId: delegated.then(delegated => delegated.accountId)});
const ipamDelegateAccount = new aws.Provider("ipamDelegateAccount", {});
// authentication arguments omitted
```
```python
import pulumi
import pulumi_aws as aws

delegated = aws.get_caller_identity()
example = aws.ec2.VpcIpamOrganizationAdminAccount("example", delegated_admin_account_id=delegated.account_id)
ipam_delegate_account = aws.Provider("ipamDelegateAccount")
# authentication arguments omitted
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var delegated = Aws.GetCallerIdentity.Invoke();

    var example = new Aws.Ec2.VpcIpamOrganizationAdminAccount("example", new()
    {
        DelegatedAdminAccountId = delegated.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
    });

    var ipamDelegateAccount = new Aws.Provider("ipamDelegateAccount");

    // authentication arguments omitted
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		delegated, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewVpcIpamOrganizationAdminAccount(ctx, "example", &ec2.VpcIpamOrganizationAdminAccountArgs{
			DelegatedAdminAccountId: *pulumi.String(delegated.AccountId),
		})
		if err != nil {
			return err
		}
		_, err = aws.NewProvider(ctx, "ipamDelegateAccount", nil)
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.ec2.VpcIpamOrganizationAdminAccount;
import com.pulumi.aws.ec2.VpcIpamOrganizationAdminAccountArgs;
import com.pulumi.aws.Provider;
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
        final var delegated = AwsFunctions.getCallerIdentity();

        var example = new VpcIpamOrganizationAdminAccount("example", VpcIpamOrganizationAdminAccountArgs.builder()        
            .delegatedAdminAccountId(delegated.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
            .build());

        var ipamDelegateAccount = new Provider("ipamDelegateAccount");

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:VpcIpamOrganizationAdminAccount
    properties:
      delegatedAdminAccountId: ${delegated.accountId}
  ipamDelegateAccount:
    type: pulumi:providers:aws
variables:
  delegated:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IPAMs using the delegate account `id`. For example:

```sh
 $ pulumi import aws:ec2/vpcIpamOrganizationAdminAccount:VpcIpamOrganizationAdminAccount example 12345678901
```
 