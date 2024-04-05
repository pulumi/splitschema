Provides a [Route53 Delegation Set](https://docs.aws.amazon.com/Route53/latest/APIReference/API-actions-by-function.html#actions-by-function-reusable-delegation-sets) resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.route53.DelegationSet("main", {referenceName: "DynDNS"});
const primary = new aws.route53.Zone("primary", {delegationSetId: main.id});
const secondary = new aws.route53.Zone("secondary", {delegationSetId: main.id});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.route53.DelegationSet("main", reference_name="DynDNS")
primary = aws.route53.Zone("primary", delegation_set_id=main.id)
secondary = aws.route53.Zone("secondary", delegation_set_id=main.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Route53.DelegationSet("main", new()
    {
        ReferenceName = "DynDNS",
    });

    var primary = new Aws.Route53.Zone("primary", new()
    {
        DelegationSetId = main.Id,
    });

    var secondary = new Aws.Route53.Zone("secondary", new()
    {
        DelegationSetId = main.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := route53.NewDelegationSet(ctx, "main", &route53.DelegationSetArgs{
			ReferenceName: pulumi.String("DynDNS"),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewZone(ctx, "primary", &route53.ZoneArgs{
			DelegationSetId: main.ID(),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewZone(ctx, "secondary", &route53.ZoneArgs{
			DelegationSetId: main.ID(),
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
import com.pulumi.aws.route53.DelegationSet;
import com.pulumi.aws.route53.DelegationSetArgs;
import com.pulumi.aws.route53.Zone;
import com.pulumi.aws.route53.ZoneArgs;
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
        var main = new DelegationSet("main", DelegationSetArgs.builder()        
            .referenceName("DynDNS")
            .build());

        var primary = new Zone("primary", ZoneArgs.builder()        
            .delegationSetId(main.id())
            .build());

        var secondary = new Zone("secondary", ZoneArgs.builder()        
            .delegationSetId(main.id())
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:route53:DelegationSet
    properties:
      referenceName: DynDNS
  primary:
    type: aws:route53:Zone
    properties:
      delegationSetId: ${main.id}
  secondary:
    type: aws:route53:Zone
    properties:
      delegationSetId: ${main.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route53 Delegation Sets using the delegation set `id`. For example:

```sh
 $ pulumi import aws:route53/delegationSet:DelegationSet set1 N1PA6795SAMPLE
```
 