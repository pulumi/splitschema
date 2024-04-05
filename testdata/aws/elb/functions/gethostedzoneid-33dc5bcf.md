Use this data source to get the HostedZoneId of the AWS Elastic Load Balancing HostedZoneId
in a given region for the purpose of using in an AWS Route53 Alias.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = aws.elb.getHostedZoneId({});
const www = new aws.route53.Record("www", {
    zoneId: aws_route53_zone.primary.zone_id,
    name: "example.com",
    type: "A",
    aliases: [{
        name: aws_elb.main.dns_name,
        zoneId: main.then(main => main.id),
        evaluateTargetHealth: true,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.elb.get_hosted_zone_id()
www = aws.route53.Record("www",
    zone_id=aws_route53_zone["primary"]["zone_id"],
    name="example.com",
    type="A",
    aliases=[aws.route53.RecordAliasArgs(
        name=aws_elb["main"]["dns_name"],
        zone_id=main.id,
        evaluate_target_health=True,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = Aws.Elb.GetHostedZoneId.Invoke();

    var www = new Aws.Route53.Record("www", new()
    {
        ZoneId = aws_route53_zone.Primary.Zone_id,
        Name = "example.com",
        Type = "A",
        Aliases = new[]
        {
            new Aws.Route53.Inputs.RecordAliasArgs
            {
                Name = aws_elb.Main.Dns_name,
                ZoneId = main.Apply(getHostedZoneIdResult => getHostedZoneIdResult.Id),
                EvaluateTargetHealth = true,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := elb.GetHostedZoneId(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "www", &route53.RecordArgs{
			ZoneId: pulumi.Any(aws_route53_zone.Primary.Zone_id),
			Name:   pulumi.String("example.com"),
			Type:   pulumi.String("A"),
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					Name:                 pulumi.Any(aws_elb.Main.Dns_name),
					ZoneId:               *pulumi.String(main.Id),
					EvaluateTargetHealth: pulumi.Bool(true),
				},
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
import com.pulumi.aws.elb.ElbFunctions;
import com.pulumi.aws.elb.inputs.GetHostedZoneIdArgs;
import com.pulumi.aws.route53.Record;
import com.pulumi.aws.route53.RecordArgs;
import com.pulumi.aws.route53.inputs.RecordAliasArgs;
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
        final var main = ElbFunctions.getHostedZoneId();

        var www = new Record("www", RecordArgs.builder()        
            .zoneId(aws_route53_zone.primary().zone_id())
            .name("example.com")
            .type("A")
            .aliases(RecordAliasArgs.builder()
                .name(aws_elb.main().dns_name())
                .zoneId(main.applyValue(getHostedZoneIdResult -> getHostedZoneIdResult.id()))
                .evaluateTargetHealth(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  www:
    type: aws:route53:Record
    properties:
      zoneId: ${aws_route53_zone.primary.zone_id}
      name: example.com
      type: A
      aliases:
        - name: ${aws_elb.main.dns_name}
          zoneId: ${main.id}
          evaluateTargetHealth: true
variables:
  main:
    fn::invoke:
      Function: aws:elb:getHostedZoneId
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}