This resource can be used to get a set of license grant ARNs matching a filter.

{{% examples %}}
## Example Usage
{{% example %}}

The following shows getting all license grant ARNs granted to your account.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const test = current.then(current => aws.licensemanager.getLicenseGrants({
    filters: [{
        name: "GranteePrincipalARN",
        values: [`arn:aws:iam::${current.accountId}:root`],
    }],
}));
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
test = aws.licensemanager.get_license_grants(filters=[aws.licensemanager.GetLicenseGrantsFilterArgs(
    name="GranteePrincipalARN",
    values=[f"arn:aws:iam::{current.account_id}:root"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var test = Aws.LicenseManager.GetLicenseGrants.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.LicenseManager.Inputs.GetLicenseGrantsFilterInputArgs
            {
                Name = "GranteePrincipalARN",
                Values = new[]
                {
                    $"arn:aws:iam::{current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:root",
                },
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/licensemanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = licensemanager.GetLicenseGrants(ctx, &licensemanager.GetLicenseGrantsArgs{
			Filters: []licensemanager.GetLicenseGrantsFilter{
				{
					Name: "GranteePrincipalARN",
					Values: []string{
						fmt.Sprintf("arn:aws:iam::%v:root", current.AccountId),
					},
				},
			},
		}, nil)
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
import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
import com.pulumi.aws.licensemanager.inputs.GetLicenseGrantsArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        final var test = LicensemanagerFunctions.getLicenseGrants(GetLicenseGrantsArgs.builder()
            .filters(GetLicenseGrantsFilterArgs.builder()
                .name("GranteePrincipalARN")
                .values(String.format("arn:aws:iam::%s:root", current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                .build())
            .build());

    }
}
```
```yaml
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  test:
    fn::invoke:
      Function: aws:licensemanager:getLicenseGrants
      Arguments:
        filters:
          - name: GranteePrincipalARN
            values:
              - arn:aws:iam::${current.accountId}:root
```
{{% /example %}}
{{% /examples %}}