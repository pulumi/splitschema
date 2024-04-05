Provides a WAF Regional Geo Match Set Resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const geoMatchSet = new aws.wafregional.GeoMatchSet("geoMatchSet", {geoMatchConstraints: [
    {
        type: "Country",
        value: "US",
    },
    {
        type: "Country",
        value: "CA",
    },
]});
```
```python
import pulumi
import pulumi_aws as aws

geo_match_set = aws.wafregional.GeoMatchSet("geoMatchSet", geo_match_constraints=[
    aws.wafregional.GeoMatchSetGeoMatchConstraintArgs(
        type="Country",
        value="US",
    ),
    aws.wafregional.GeoMatchSetGeoMatchConstraintArgs(
        type="Country",
        value="CA",
    ),
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var geoMatchSet = new Aws.WafRegional.GeoMatchSet("geoMatchSet", new()
    {
        GeoMatchConstraints = new[]
        {
            new Aws.WafRegional.Inputs.GeoMatchSetGeoMatchConstraintArgs
            {
                Type = "Country",
                Value = "US",
            },
            new Aws.WafRegional.Inputs.GeoMatchSetGeoMatchConstraintArgs
            {
                Type = "Country",
                Value = "CA",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/wafregional"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := wafregional.NewGeoMatchSet(ctx, "geoMatchSet", &wafregional.GeoMatchSetArgs{
			GeoMatchConstraints: wafregional.GeoMatchSetGeoMatchConstraintArray{
				&wafregional.GeoMatchSetGeoMatchConstraintArgs{
					Type:  pulumi.String("Country"),
					Value: pulumi.String("US"),
				},
				&wafregional.GeoMatchSetGeoMatchConstraintArgs{
					Type:  pulumi.String("Country"),
					Value: pulumi.String("CA"),
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
import com.pulumi.aws.wafregional.GeoMatchSet;
import com.pulumi.aws.wafregional.GeoMatchSetArgs;
import com.pulumi.aws.wafregional.inputs.GeoMatchSetGeoMatchConstraintArgs;
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
        var geoMatchSet = new GeoMatchSet("geoMatchSet", GeoMatchSetArgs.builder()        
            .geoMatchConstraints(            
                GeoMatchSetGeoMatchConstraintArgs.builder()
                    .type("Country")
                    .value("US")
                    .build(),
                GeoMatchSetGeoMatchConstraintArgs.builder()
                    .type("Country")
                    .value("CA")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  geoMatchSet:
    type: aws:wafregional:GeoMatchSet
    properties:
      geoMatchConstraints:
        - type: Country
          value: US
        - type: Country
          value: CA
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WAF Regional Geo Match Set using the id. For example:

```sh
 $ pulumi import aws:wafregional/geoMatchSet:GeoMatchSet geo_match_set a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```
 