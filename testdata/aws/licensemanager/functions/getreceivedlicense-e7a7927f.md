This resource can be used to get data on a received license using an ARN. This can be helpful for pulling in data on a license from the AWS marketplace and sharing that license with another account.

{{% examples %}}
## Example Usage
{{% example %}}

The following shows getting the received license data using and ARN.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.licensemanager.getReceivedLicense({
    licenseArn: "arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.licensemanager.get_received_license(license_arn="arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.LicenseManager.GetReceivedLicense.Invoke(new()
    {
        LicenseArn = "arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0",
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
		_, err := licensemanager.GetReceivedLicense(ctx, &licensemanager.GetReceivedLicenseArgs{
			LicenseArn: "arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0",
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
import com.pulumi.aws.licensemanager.inputs.GetReceivedLicenseArgs;
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
        final var test = LicensemanagerFunctions.getReceivedLicense(GetReceivedLicenseArgs.builder()
            .licenseArn("arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:licensemanager:getReceivedLicense
      Arguments:
        licenseArn: arn:aws:license-manager::111111111111:license:l-ecbaa94eb71a4830b6d7e49268fecaa0
```
{{% /example %}}
{{% /examples %}}