Manages an RDS DB Reserved Instance.

> **NOTE:** Once created, a reservation is valid for the `duration` of the provided `offering_id` and cannot be deleted. Performing a `destroy` will only remove the resource from state. For more information see [RDS Reserved Instances Documentation](https://aws.amazon.com/rds/reserved-instances/) and [PurchaseReservedDBInstancesOffering](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_PurchaseReservedDBInstancesOffering.html).

> **NOTE:** Due to the expense of testing this resource, we provide it as best effort. If you find it useful, and have the ability to help test or notice issues, consider reaching out to us on GitHub.

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
const example = new aws.rds.ReservedInstance("example", {
    offeringId: test.then(test => test.offeringId),
    reservationId: "optionalCustomReservationID",
    instanceCount: 3,
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
example = aws.rds.ReservedInstance("example",
    offering_id=test.offering_id,
    reservation_id="optionalCustomReservationID",
    instance_count=3)
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

    var example = new Aws.Rds.ReservedInstance("example", new()
    {
        OfferingId = test.Apply(getReservedInstanceOfferingResult => getReservedInstanceOfferingResult.OfferingId),
        ReservationId = "optionalCustomReservationID",
        InstanceCount = 3,
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
		test, err := rds.GetReservedInstanceOffering(ctx, &rds.GetReservedInstanceOfferingArgs{
			DbInstanceClass:    "db.t2.micro",
			Duration:           31536000,
			MultiAz:            false,
			OfferingType:       "All Upfront",
			ProductDescription: "mysql",
		}, nil)
		if err != nil {
			return err
		}
		_, err = rds.NewReservedInstance(ctx, "example", &rds.ReservedInstanceArgs{
			OfferingId:    *pulumi.String(test.OfferingId),
			ReservationId: pulumi.String("optionalCustomReservationID"),
			InstanceCount: pulumi.Int(3),
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
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetReservedInstanceOfferingArgs;
import com.pulumi.aws.rds.ReservedInstance;
import com.pulumi.aws.rds.ReservedInstanceArgs;
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

        var example = new ReservedInstance("example", ReservedInstanceArgs.builder()        
            .offeringId(test.applyValue(getReservedInstanceOfferingResult -> getReservedInstanceOfferingResult.offeringId()))
            .reservationId("optionalCustomReservationID")
            .instanceCount(3)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rds:ReservedInstance
    properties:
      offeringId: ${test.offeringId}
      reservationId: optionalCustomReservationID
      instanceCount: 3
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

## Import

Using `pulumi import`, import RDS DB Instance Reservations using the `instance_id`. For example:

```sh
 $ pulumi import aws:rds/reservedInstance:ReservedInstance reservation_instance CustomReservationID
```
 