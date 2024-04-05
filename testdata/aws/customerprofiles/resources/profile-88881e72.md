Resource for managing an Amazon Customer Profiles Profile.
See the [Create Profile](https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_CreateProfile.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDomain = new aws.customerprofiles.Domain("exampleDomain", {domainName: "example"});
const exampleProfile = new aws.customerprofiles.Profile("exampleProfile", {domainName: exampleDomain.domainName});
```
```python
import pulumi
import pulumi_aws as aws

example_domain = aws.customerprofiles.Domain("exampleDomain", domain_name="example")
example_profile = aws.customerprofiles.Profile("exampleProfile", domain_name=example_domain.domain_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleDomain = new Aws.CustomerProfiles.Domain("exampleDomain", new()
    {
        DomainName = "example",
    });

    var exampleProfile = new Aws.CustomerProfiles.Profile("exampleProfile", new()
    {
        DomainName = exampleDomain.DomainName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/customerprofiles"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDomain, err := customerprofiles.NewDomain(ctx, "exampleDomain", &customerprofiles.DomainArgs{
			DomainName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = customerprofiles.NewProfile(ctx, "exampleProfile", &customerprofiles.ProfileArgs{
			DomainName: exampleDomain.DomainName,
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
import com.pulumi.aws.customerprofiles.Domain;
import com.pulumi.aws.customerprofiles.DomainArgs;
import com.pulumi.aws.customerprofiles.Profile;
import com.pulumi.aws.customerprofiles.ProfileArgs;
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
        var exampleDomain = new Domain("exampleDomain", DomainArgs.builder()        
            .domainName("example")
            .build());

        var exampleProfile = new Profile("exampleProfile", ProfileArgs.builder()        
            .domainName(exampleDomain.domainName())
            .build());

    }
}
```
```yaml
resources:
  exampleDomain:
    type: aws:customerprofiles:Domain
    properties:
      domainName: example
  exampleProfile:
    type: aws:customerprofiles:Profile
    properties:
      domainName: ${exampleDomain.domainName}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Customer Profiles Profile using the resource `id`. For example:

```sh
 $ pulumi import aws:customerprofiles/profile:Profile example domain-name/5f2f473dfbe841eb8d05cfc2a4c926df
```
 