Information about a single RDS Reserved Instance Offering.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.rds.getReservedInstanceOffering({
    dbInstanceClass: "db.t2.micro",
    duration: 31536000,
    multiAz: false,
    offeringType: "All Upfront",
    productDescription: "mysql",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.rds.get_reserved_instance_offering(db_instance_class="db.t2.micro",
    duration=31536000,
    multi_az=False,
    offering_type="All Upfront",
    product_description="mysql")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Rds.GetReservedInstanceOffering.Invoke(new()
    {
        DbInstanceClass = "db.t2.micro",
        Duration = 31536000,
        MultiAz = false,
        OfferingType = "All Upfront",
        ProductDescription = "mysql",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.GetReservedInstanceOffering(ctx, &rds.GetReservedInstanceOfferingArgs{
			DbInstanceClass:    "db.t2.micro",
			Duration:           31536000,
			MultiAz:            false,
			OfferingType:       "All Upfront",
			ProductDescription: "mysql",
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
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetReservedInstanceOfferingArgs;
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
        final var test = RdsFunctions.getReservedInstanceOffering(GetReservedInstanceOfferingArgs.builder()
            .dbInstanceClass("db.t2.micro")
            .duration(31536000)
            .multiAz(false)
            .offeringType("All Upfront")
            .productDescription("mysql")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:rds:getReservedInstanceOffering
      Arguments:
        dbInstanceClass: db.t2.micro
        duration: 3.1536e+07
        multiAz: false
        offeringType: All Upfront
        productDescription: mysql
```
{{% /example %}}
{{% /examples %}}