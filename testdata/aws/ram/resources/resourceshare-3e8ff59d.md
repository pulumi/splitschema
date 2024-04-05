Manages a Resource Access Manager (RAM) Resource Share. To associate principals with the share, see the `aws.ram.PrincipalAssociation` resource. To associate resources with the share, see the `aws.ram.ResourceAssociation` resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ram.ResourceShare("example", {
    allowExternalPrincipals: true,
    tags: {
        Environment: "Production",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ram.ResourceShare("example",
    allow_external_principals=True,
    tags={
        "Environment": "Production",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ram.ResourceShare("example", new()
    {
        AllowExternalPrincipals = true,
        Tags = 
        {
            { "Environment", "Production" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ram.NewResourceShare(ctx, "example", &ram.ResourceShareArgs{
			AllowExternalPrincipals: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("Production"),
			},
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
import com.pulumi.aws.ram.ResourceShare;
import com.pulumi.aws.ram.ResourceShareArgs;
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
        var example = new ResourceShare("example", ResourceShareArgs.builder()        
            .allowExternalPrincipals(true)
            .tags(Map.of("Environment", "Production"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ram:ResourceShare
    properties:
      allowExternalPrincipals: true
      tags:
        Environment: Production
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import resource shares using the `arn` of the resource share. For example:

```sh
 $ pulumi import aws:ram/resourceShare:ResourceShare example arn:aws:ram:eu-west-1:123456789012:resource-share/73da1ab9-b94a-4ba3-8eb4-45917f7f4b12
```
 