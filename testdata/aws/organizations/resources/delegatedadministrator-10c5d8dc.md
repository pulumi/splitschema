Provides a resource to manage an [AWS Organizations Delegated Administrator](https://docs.aws.amazon.com/organizations/latest/APIReference/API_RegisterDelegatedAdministrator.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.organizations.DelegatedAdministrator("example", {
    accountId: "123456789012",
    servicePrincipal: "principal",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.organizations.DelegatedAdministrator("example",
    account_id="123456789012",
    service_principal="principal")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Organizations.DelegatedAdministrator("example", new()
    {
        AccountId = "123456789012",
        ServicePrincipal = "principal",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewDelegatedAdministrator(ctx, "example", &organizations.DelegatedAdministratorArgs{
			AccountId:        pulumi.String("123456789012"),
			ServicePrincipal: pulumi.String("principal"),
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
import com.pulumi.aws.organizations.DelegatedAdministrator;
import com.pulumi.aws.organizations.DelegatedAdministratorArgs;
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
        var example = new DelegatedAdministrator("example", DelegatedAdministratorArgs.builder()        
            .accountId("123456789012")
            .servicePrincipal("principal")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:organizations:DelegatedAdministrator
    properties:
      accountId: '123456789012'
      servicePrincipal: principal
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_organizations_delegated_administrator` using the account ID and its service principal. For example:

```sh
 $ pulumi import aws:organizations/delegatedAdministrator:DelegatedAdministrator example 123456789012/config.amazonaws.com
```
 