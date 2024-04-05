Data source for getting an AWS Organizations Organizational Unit.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const org = aws.organizations.getOrganization({});
const ou = org.then(org => aws.organizations.getOrganizationalUnit({
    parentId: org.roots?.[0]?.id,
    name: "dev",
}));
```
```python
import pulumi
import pulumi_aws as aws

org = aws.organizations.get_organization()
ou = aws.organizations.get_organizational_unit(parent_id=org.roots[0].id,
    name="dev")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var org = Aws.Organizations.GetOrganization.Invoke();

    var ou = Aws.Organizations.GetOrganizationalUnit.Invoke(new()
    {
        ParentId = org.Apply(getOrganizationResult => getOrganizationResult.Roots[0]?.Id),
        Name = "dev",
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
		org, err := organizations.LookupOrganization(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = organizations.LookupOrganizationalUnit(ctx, &organizations.LookupOrganizationalUnitArgs{
			ParentId: org.Roots[0].Id,
			Name:     "dev",
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
import com.pulumi.aws.organizations.OrganizationsFunctions;
import com.pulumi.aws.organizations.inputs.GetOrganizationalUnitArgs;
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
        final var org = OrganizationsFunctions.getOrganization();

        final var ou = OrganizationsFunctions.getOrganizationalUnit(GetOrganizationalUnitArgs.builder()
            .parentId(org.applyValue(getOrganizationResult -> getOrganizationResult.roots()[0].id()))
            .name("dev")
            .build());

    }
}
```
```yaml
variables:
  org:
    fn::invoke:
      Function: aws:organizations:getOrganization
      Arguments: {}
  ou:
    fn::invoke:
      Function: aws:organizations:getOrganizationalUnit
      Arguments:
        parentId: ${org.roots[0].id}
        name: dev
```
{{% /example %}}
{{% /examples %}}