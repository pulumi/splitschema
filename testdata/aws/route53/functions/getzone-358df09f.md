`aws.route53.Zone` provides details about a specific Route 53 Hosted Zone.

This data source allows to find a Hosted Zone ID given Hosted Zone name and certain search criteria.

{{% examples %}}
## Example Usage
{{% example %}}

The following example shows how to get a Hosted Zone from its name and from this data how to create a Record Set.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const selected = aws.route53.getZone({
    name: "test.com.",
    privateZone: true,
});
const www = new aws.route53.Record("www", {
    zoneId: selected.then(selected => selected.zoneId),
    name: selected.then(selected => `www.${selected.name}`),
    type: "A",
    ttl: 300,
    records: ["10.0.0.1"],
});
```
```python
import pulumi
import pulumi_aws as aws

selected = aws.route53.get_zone(name="test.com.",
    private_zone=True)
www = aws.route53.Record("www",
    zone_id=selected.zone_id,
    name=f"www.{selected.name}",
    type="A",
    ttl=300,
    records=["10.0.0.1"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var selected = Aws.Route53.GetZone.Invoke(new()
    {
        Name = "test.com.",
        PrivateZone = true,
    });

    var www = new Aws.Route53.Record("www", new()
    {
        ZoneId = selected.Apply(getZoneResult => getZoneResult.ZoneId),
        Name = $"www.{selected.Apply(getZoneResult => getZoneResult.Name)}",
        Type = "A",
        Ttl = 300,
        Records = new[]
        {
            "10.0.0.1",
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		selected, err := route53.LookupZone(ctx, &route53.LookupZoneArgs{
			Name:        pulumi.StringRef("test.com."),
			PrivateZone: pulumi.BoolRef(true),
		}, nil)
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "www", &route53.RecordArgs{
			ZoneId: *pulumi.String(selected.ZoneId),
			Name:   pulumi.String(fmt.Sprintf("www.%v", selected.Name)),
			Type:   pulumi.String("A"),
			Ttl:    pulumi.Int(300),
			Records: pulumi.StringArray{
				pulumi.String("10.0.0.1"),
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
import com.pulumi.aws.route53.Route53Functions;
import com.pulumi.aws.route53.inputs.GetZoneArgs;
import com.pulumi.aws.route53.Record;
import com.pulumi.aws.route53.RecordArgs;
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
        final var selected = Route53Functions.getZone(GetZoneArgs.builder()
            .name("test.com.")
            .privateZone(true)
            .build());

        var www = new Record("www", RecordArgs.builder()        
            .zoneId(selected.applyValue(getZoneResult -> getZoneResult.zoneId()))
            .name(String.format("www.%s", selected.applyValue(getZoneResult -> getZoneResult.name())))
            .type("A")
            .ttl("300")
            .records("10.0.0.1")
            .build());

    }
}
```
```yaml
resources:
  www:
    type: aws:route53:Record
    properties:
      zoneId: ${selected.zoneId}
      name: www.${selected.name}
      type: A
      ttl: '300'
      records:
        - 10.0.0.1
variables:
  selected:
    fn::invoke:
      Function: aws:route53:getZone
      Arguments:
        name: test.com.
        privateZone: true
```
{{% /example %}}
{{% /examples %}}