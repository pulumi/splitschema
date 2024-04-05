This resource can be used to get a set of license ARNs matching a filter.

{{% examples %}}
## Example Usage
{{% example %}}

The following shows getting all license ARNs issued from the AWS marketplace. Providing no filter, would provide all license ARNs for the entire account.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.licensemanager.getReceivedLicenses({
    filters: [{
        name: "IssuerName",
        values: ["AWS/Marketplace"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.licensemanager.get_received_licenses(filters=[aws.licensemanager.GetReceivedLicensesFilterArgs(
    name="IssuerName",
    values=["AWS/Marketplace"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.LicenseManager.GetReceivedLicenses.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.LicenseManager.Inputs.GetReceivedLicensesFilterInputArgs
            {
                Name = "IssuerName",
                Values = new[]
                {
                    "AWS/Marketplace",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/licensemanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := licensemanager.GetReceivedLicenses(ctx, &licensemanager.GetReceivedLicensesArgs{
			Filters: []licensemanager.GetReceivedLicensesFilter{
				{
					Name: "IssuerName",
					Values: []string{
						"AWS/Marketplace",
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
import com.pulumi.aws.licensemanager.LicensemanagerFunctions;
import com.pulumi.aws.licensemanager.inputs.GetReceivedLicensesArgs;
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
        final var test = LicensemanagerFunctions.getReceivedLicenses(GetReceivedLicensesArgs.builder()
            .filters(GetReceivedLicensesFilterArgs.builder()
                .name("IssuerName")
                .values("AWS/Marketplace")
                .build())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:licensemanager:getReceivedLicenses
      Arguments:
        filters:
          - name: IssuerName
            values:
              - AWS/Marketplace
```
{{% /example %}}
{{% /examples %}}