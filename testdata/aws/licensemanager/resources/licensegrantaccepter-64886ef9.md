Accepts a License Manager grant. This allows for sharing licenses with other aws accounts.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.licensemanager.LicenseGrantAccepter("test", {grantArn: "arn:aws:license-manager::123456789012:grant:g-1cf9fba4ba2f42dcab11c686c4b4d329"});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.licensemanager.LicenseGrantAccepter("test", grant_arn="arn:aws:license-manager::123456789012:grant:g-1cf9fba4ba2f42dcab11c686c4b4d329")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LicenseManager.LicenseGrantAccepter("test", new()
    {
        GrantArn = "arn:aws:license-manager::123456789012:grant:g-1cf9fba4ba2f42dcab11c686c4b4d329",
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
		_, err := licensemanager.NewLicenseGrantAccepter(ctx, "test", &licensemanager.LicenseGrantAccepterArgs{
			GrantArn: pulumi.String("arn:aws:license-manager::123456789012:grant:g-1cf9fba4ba2f42dcab11c686c4b4d329"),
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
import com.pulumi.aws.licensemanager.LicenseGrantAccepter;
import com.pulumi.aws.licensemanager.LicenseGrantAccepterArgs;
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
        var test = new LicenseGrantAccepter("test", LicenseGrantAccepterArgs.builder()        
            .grantArn("arn:aws:license-manager::123456789012:grant:g-1cf9fba4ba2f42dcab11c686c4b4d329")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:licensemanager:LicenseGrantAccepter
    properties:
      grantArn: arn:aws:license-manager::123456789012:grant:g-1cf9fba4ba2f42dcab11c686c4b4d329
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_licensemanager_grant_accepter` using the grant arn. For example:

```sh
 $ pulumi import aws:licensemanager/licenseGrantAccepter:LicenseGrantAccepter test arn:aws:license-manager::123456789012:grant:g-1cf9fba4ba2f42dcab11c686c4b4d329
```
 