Provides a resource to create an organizational unit.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.organizations.OrganizationalUnit("example", {parentId: aws_organizations_organization.example.roots[0].id});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.organizations.OrganizationalUnit("example", parent_id=aws_organizations_organization["example"]["roots"][0]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Organizations.OrganizationalUnit("example", new()
    {
        ParentId = aws_organizations_organization.Example.Roots[0].Id,
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
		_, err := organizations.NewOrganizationalUnit(ctx, "example", &organizations.OrganizationalUnitArgs{
			ParentId: pulumi.Any(aws_organizations_organization.Example.Roots[0].Id),
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
import com.pulumi.aws.organizations.OrganizationalUnit;
import com.pulumi.aws.organizations.OrganizationalUnitArgs;
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
        var example = new OrganizationalUnit("example", OrganizationalUnitArgs.builder()        
            .parentId(aws_organizations_organization.example().roots()[0].id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:organizations:OrganizationalUnit
    properties:
      parentId: ${aws_organizations_organization.example.roots[0].id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AWS Organizations Organizational Units using the `id`. For example:

```sh
 $ pulumi import aws:organizations/organizationalUnit:OrganizationalUnit example ou-1234567
```
 