This resource provisions and manages a Service Catalog provisioned product.

A provisioned product is a resourced instance of a product. For example, provisioning a product based on a CloudFormation template launches a CloudFormation stack and its underlying resources.

Like this resource, the `aws_servicecatalog_record` data source also provides information about a provisioned product. Although a Service Catalog record provides some overlapping information with this resource, a record is tied to a provisioned product event, such as provisioning, termination, and updating.

> **Tip:** If you include conflicted keys as tags, AWS will report an error, "Parameter validation failed: Missing required parameter in Tags[N]:Value".

> **Tip:** A "provisioning artifact" is also referred to as a "version." A "distributor" is also referred to as a "vendor."

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicecatalog.ProvisionedProduct("example", {
    productName: "Example product",
    provisioningArtifactName: "Example version",
    provisioningParameters: [{
        key: "foo",
        value: "bar",
    }],
    tags: {
        foo: "bar",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.ProvisionedProduct("example",
    product_name="Example product",
    provisioning_artifact_name="Example version",
    provisioning_parameters=[aws.servicecatalog.ProvisionedProductProvisioningParameterArgs(
        key="foo",
        value="bar",
    )],
    tags={
        "foo": "bar",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceCatalog.ProvisionedProduct("example", new()
    {
        ProductName = "Example product",
        ProvisioningArtifactName = "Example version",
        ProvisioningParameters = new[]
        {
            new Aws.ServiceCatalog.Inputs.ProvisionedProductProvisioningParameterArgs
            {
                Key = "foo",
                Value = "bar",
            },
        },
        Tags = 
        {
            { "foo", "bar" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicecatalog"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicecatalog.NewProvisionedProduct(ctx, "example", &servicecatalog.ProvisionedProductArgs{
			ProductName:              pulumi.String("Example product"),
			ProvisioningArtifactName: pulumi.String("Example version"),
			ProvisioningParameters: servicecatalog.ProvisionedProductProvisioningParameterArray{
				&servicecatalog.ProvisionedProductProvisioningParameterArgs{
					Key:   pulumi.String("foo"),
					Value: pulumi.String("bar"),
				},
			},
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
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
import com.pulumi.aws.servicecatalog.ProvisionedProduct;
import com.pulumi.aws.servicecatalog.ProvisionedProductArgs;
import com.pulumi.aws.servicecatalog.inputs.ProvisionedProductProvisioningParameterArgs;
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
        var example = new ProvisionedProduct("example", ProvisionedProductArgs.builder()        
            .productName("Example product")
            .provisioningArtifactName("Example version")
            .provisioningParameters(ProvisionedProductProvisioningParameterArgs.builder()
                .key("foo")
                .value("bar")
                .build())
            .tags(Map.of("foo", "bar"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicecatalog:ProvisionedProduct
    properties:
      productName: Example product
      provisioningArtifactName: Example version
      provisioningParameters:
        - key: foo
          value: bar
      tags:
        foo: bar
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_servicecatalog_provisioned_product` using the provisioned product ID. For example:

```sh
 $ pulumi import aws:servicecatalog/provisionedProduct:ProvisionedProduct example pp-dnigbtea24ste
```
 