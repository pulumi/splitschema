Provides a Route53 CIDR location resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleCidrCollection = new aws.route53.CidrCollection("exampleCidrCollection", {});
const exampleCidrLocation = new aws.route53.CidrLocation("exampleCidrLocation", {
    cidrCollectionId: exampleCidrCollection.id,
    cidrBlocks: [
        "200.5.3.0/24",
        "200.6.3.0/24",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_cidr_collection = aws.route53.CidrCollection("exampleCidrCollection")
example_cidr_location = aws.route53.CidrLocation("exampleCidrLocation",
    cidr_collection_id=example_cidr_collection.id,
    cidr_blocks=[
        "200.5.3.0/24",
        "200.6.3.0/24",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleCidrCollection = new Aws.Route53.CidrCollection("exampleCidrCollection");

    var exampleCidrLocation = new Aws.Route53.CidrLocation("exampleCidrLocation", new()
    {
        CidrCollectionId = exampleCidrCollection.Id,
        CidrBlocks = new[]
        {
            "200.5.3.0/24",
            "200.6.3.0/24",
        },
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
		exampleCidrCollection, err := route53.NewCidrCollection(ctx, "exampleCidrCollection", nil)
		if err != nil {
			return err
		}
		_, err = route53.NewCidrLocation(ctx, "exampleCidrLocation", &route53.CidrLocationArgs{
			CidrCollectionId: exampleCidrCollection.ID(),
			CidrBlocks: pulumi.StringArray{
				pulumi.String("200.5.3.0/24"),
				pulumi.String("200.6.3.0/24"),
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
import com.pulumi.aws.route53.CidrCollection;
import com.pulumi.aws.route53.CidrLocation;
import com.pulumi.aws.route53.CidrLocationArgs;
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
        var exampleCidrCollection = new CidrCollection("exampleCidrCollection");

        var exampleCidrLocation = new CidrLocation("exampleCidrLocation", CidrLocationArgs.builder()        
            .cidrCollectionId(exampleCidrCollection.id())
            .cidrBlocks(            
                "200.5.3.0/24",
                "200.6.3.0/24")
            .build());

    }
}
```
```yaml
resources:
  exampleCidrCollection:
    type: aws:route53:CidrCollection
  exampleCidrLocation:
    type: aws:route53:CidrLocation
    properties:
      cidrCollectionId: ${exampleCidrCollection.id}
      cidrBlocks:
        - 200.5.3.0/24
        - 200.6.3.0/24
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CIDR locations using their the CIDR collection ID and location name. For example:

```sh
 $ pulumi import aws:route53/cidrLocation:CidrLocation example 9ac32814-3e67-0932-6048-8d779cc6f511,office
```
 