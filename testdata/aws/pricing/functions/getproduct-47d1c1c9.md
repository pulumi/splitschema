Use this data source to get the pricing information of all products in AWS.
This data source is only available in a us-east-1 or ap-south-1 provider.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.pricing.getProduct({
    filters: [
        {
            field: "instanceType",
            value: "c5.xlarge",
        },
        {
            field: "operatingSystem",
            value: "Linux",
        },
        {
            field: "location",
            value: "US East (N. Virginia)",
        },
        {
            field: "preInstalledSw",
            value: "NA",
        },
        {
            field: "licenseModel",
            value: "No License required",
        },
        {
            field: "tenancy",
            value: "Shared",
        },
        {
            field: "capacitystatus",
            value: "Used",
        },
    ],
    serviceCode: "AmazonEC2",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.pricing.get_product(filters=[
        aws.pricing.GetProductFilterArgs(
            field="instanceType",
            value="c5.xlarge",
        ),
        aws.pricing.GetProductFilterArgs(
            field="operatingSystem",
            value="Linux",
        ),
        aws.pricing.GetProductFilterArgs(
            field="location",
            value="US East (N. Virginia)",
        ),
        aws.pricing.GetProductFilterArgs(
            field="preInstalledSw",
            value="NA",
        ),
        aws.pricing.GetProductFilterArgs(
            field="licenseModel",
            value="No License required",
        ),
        aws.pricing.GetProductFilterArgs(
            field="tenancy",
            value="Shared",
        ),
        aws.pricing.GetProductFilterArgs(
            field="capacitystatus",
            value="Used",
        ),
    ],
    service_code="AmazonEC2")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Pricing.GetProduct.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Pricing.Inputs.GetProductFilterInputArgs
            {
                Field = "instanceType",
                Value = "c5.xlarge",
            },
            new Aws.Pricing.Inputs.GetProductFilterInputArgs
            {
                Field = "operatingSystem",
                Value = "Linux",
            },
            new Aws.Pricing.Inputs.GetProductFilterInputArgs
            {
                Field = "location",
                Value = "US East (N. Virginia)",
            },
            new Aws.Pricing.Inputs.GetProductFilterInputArgs
            {
                Field = "preInstalledSw",
                Value = "NA",
            },
            new Aws.Pricing.Inputs.GetProductFilterInputArgs
            {
                Field = "licenseModel",
                Value = "No License required",
            },
            new Aws.Pricing.Inputs.GetProductFilterInputArgs
            {
                Field = "tenancy",
                Value = "Shared",
            },
            new Aws.Pricing.Inputs.GetProductFilterInputArgs
            {
                Field = "capacitystatus",
                Value = "Used",
            },
        },
        ServiceCode = "AmazonEC2",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/pricing"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := pricing.GetProduct(ctx, &pricing.GetProductArgs{
			Filters: []pricing.GetProductFilter{
				{
					Field: "instanceType",
					Value: "c5.xlarge",
				},
				{
					Field: "operatingSystem",
					Value: "Linux",
				},
				{
					Field: "location",
					Value: "US East (N. Virginia)",
				},
				{
					Field: "preInstalledSw",
					Value: "NA",
				},
				{
					Field: "licenseModel",
					Value: "No License required",
				},
				{
					Field: "tenancy",
					Value: "Shared",
				},
				{
					Field: "capacitystatus",
					Value: "Used",
				},
			},
			ServiceCode: "AmazonEC2",
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
import com.pulumi.aws.pricing.PricingFunctions;
import com.pulumi.aws.pricing.inputs.GetProductArgs;
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
        final var example = PricingFunctions.getProduct(GetProductArgs.builder()
            .filters(            
                GetProductFilterArgs.builder()
                    .field("instanceType")
                    .value("c5.xlarge")
                    .build(),
                GetProductFilterArgs.builder()
                    .field("operatingSystem")
                    .value("Linux")
                    .build(),
                GetProductFilterArgs.builder()
                    .field("location")
                    .value("US East (N. Virginia)")
                    .build(),
                GetProductFilterArgs.builder()
                    .field("preInstalledSw")
                    .value("NA")
                    .build(),
                GetProductFilterArgs.builder()
                    .field("licenseModel")
                    .value("No License required")
                    .build(),
                GetProductFilterArgs.builder()
                    .field("tenancy")
                    .value("Shared")
                    .build(),
                GetProductFilterArgs.builder()
                    .field("capacitystatus")
                    .value("Used")
                    .build())
            .serviceCode("AmazonEC2")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:pricing:getProduct
      Arguments:
        filters:
          - field: instanceType
            value: c5.xlarge
          - field: operatingSystem
            value: Linux
          - field: location
            value: US East (N. Virginia)
          - field: preInstalledSw
            value: NA
          - field: licenseModel
            value: No License required
          - field: tenancy
            value: Shared
          - field: capacitystatus
            value: Used
        serviceCode: AmazonEC2
```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.pricing.getProduct({
    filters: [
        {
            field: "instanceType",
            value: "ds1.xlarge",
        },
        {
            field: "location",
            value: "US East (N. Virginia)",
        },
    ],
    serviceCode: "AmazonRedshift",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.pricing.get_product(filters=[
        aws.pricing.GetProductFilterArgs(
            field="instanceType",
            value="ds1.xlarge",
        ),
        aws.pricing.GetProductFilterArgs(
            field="location",
            value="US East (N. Virginia)",
        ),
    ],
    service_code="AmazonRedshift")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Pricing.GetProduct.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Pricing.Inputs.GetProductFilterInputArgs
            {
                Field = "instanceType",
                Value = "ds1.xlarge",
            },
            new Aws.Pricing.Inputs.GetProductFilterInputArgs
            {
                Field = "location",
                Value = "US East (N. Virginia)",
            },
        },
        ServiceCode = "AmazonRedshift",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/pricing"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := pricing.GetProduct(ctx, &pricing.GetProductArgs{
			Filters: []pricing.GetProductFilter{
				{
					Field: "instanceType",
					Value: "ds1.xlarge",
				},
				{
					Field: "location",
					Value: "US East (N. Virginia)",
				},
			},
			ServiceCode: "AmazonRedshift",
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
import com.pulumi.aws.pricing.PricingFunctions;
import com.pulumi.aws.pricing.inputs.GetProductArgs;
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
        final var example = PricingFunctions.getProduct(GetProductArgs.builder()
            .filters(            
                GetProductFilterArgs.builder()
                    .field("instanceType")
                    .value("ds1.xlarge")
                    .build(),
                GetProductFilterArgs.builder()
                    .field("location")
                    .value("US East (N. Virginia)")
                    .build())
            .serviceCode("AmazonRedshift")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:pricing:getProduct
      Arguments:
        filters:
          - field: instanceType
            value: ds1.xlarge
          - field: location
            value: US East (N. Virginia)
        serviceCode: AmazonRedshift
```
{{% /example %}}
{{% /examples %}}