Resource for managing AWS Audit Manager Organization Admin Account Registration.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.auditmanager.OrganizationAdminAccountRegistration("example", {adminAccountId: "012345678901"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.auditmanager.OrganizationAdminAccountRegistration("example", admin_account_id="012345678901")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Auditmanager.OrganizationAdminAccountRegistration("example", new()
    {
        AdminAccountId = "012345678901",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/auditmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := auditmanager.NewOrganizationAdminAccountRegistration(ctx, "example", &auditmanager.OrganizationAdminAccountRegistrationArgs{
			AdminAccountId: pulumi.String("012345678901"),
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
import com.pulumi.aws.auditmanager.OrganizationAdminAccountRegistration;
import com.pulumi.aws.auditmanager.OrganizationAdminAccountRegistrationArgs;
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
        var example = new OrganizationAdminAccountRegistration("example", OrganizationAdminAccountRegistrationArgs.builder()        
            .adminAccountId("012345678901")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:auditmanager:OrganizationAdminAccountRegistration
    properties:
      adminAccountId: '012345678901'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Audit Manager Organization Admin Account Registration using the `id`. For example:

```sh
 $ pulumi import aws:auditmanager/organizationAdminAccountRegistration:OrganizationAdminAccountRegistration example 012345678901
```
 