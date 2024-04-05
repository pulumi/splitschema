Provides a Location Service Route Calculator.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.location.RouteCalculation("example", {
    calculatorName: "example",
    dataSource: "Here",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.location.RouteCalculation("example",
    calculator_name="example",
    data_source="Here")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Location.RouteCalculation("example", new()
    {
        CalculatorName = "example",
        DataSource = "Here",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/location"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := location.NewRouteCalculation(ctx, "example", &location.RouteCalculationArgs{
			CalculatorName: pulumi.String("example"),
			DataSource:     pulumi.String("Here"),
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
import com.pulumi.aws.location.RouteCalculation;
import com.pulumi.aws.location.RouteCalculationArgs;
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
        var example = new RouteCalculation("example", RouteCalculationArgs.builder()        
            .calculatorName("example")
            .dataSource("Here")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:location:RouteCalculation
    properties:
      calculatorName: example
      dataSource: Here
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_location_route_calculator` using the route calculator name. For example:

```sh
 $ pulumi import aws:location/routeCalculation:RouteCalculation example example
```
 