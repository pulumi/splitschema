Manages a Service Catalog Product.

> **NOTE:** The user or role that uses this resources must have the `cloudformation:GetTemplate` IAM policy permission. This policy permission is required when using the `template_physical_id` argument.

> A "provisioning artifact" is also referred to as a "version." A "distributor" is also referred to as a "vendor."

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicecatalog.Product("example", {
    owner: "example-owner",
    provisioningArtifactParameters: {
        templateUrl: "https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/temp1.json",
    },
    tags: {
        foo: "bar",
    },
    type: "CLOUD_FORMATION_TEMPLATE",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.Product("example",
    owner="example-owner",
    provisioning_artifact_parameters=aws.servicecatalog.ProductProvisioningArtifactParametersArgs(
        template_url="https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/temp1.json",
    ),
    tags={
        "foo": "bar",
    },
    type="CLOUD_FORMATION_TEMPLATE")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceCatalog.Product("example", new()
    {
        Owner = "example-owner",
        ProvisioningArtifactParameters = new Aws.ServiceCatalog.Inputs.ProductProvisioningArtifactParametersArgs
        {
            TemplateUrl = "https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/temp1.json",
        },
        Tags = 
        {
            { "foo", "bar" },
        },
        Type = "CLOUD_FORMATION_TEMPLATE",
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
		_, err := servicecatalog.NewProduct(ctx, "example", &servicecatalog.ProductArgs{
			Owner: pulumi.String("example-owner"),
			ProvisioningArtifactParameters: &servicecatalog.ProductProvisioningArtifactParametersArgs{
				TemplateUrl: pulumi.String("https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/temp1.json"),
			},
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
			},
			Type: pulumi.String("CLOUD_FORMATION_TEMPLATE"),
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
import com.pulumi.aws.servicecatalog.Product;
import com.pulumi.aws.servicecatalog.ProductArgs;
import com.pulumi.aws.servicecatalog.inputs.ProductProvisioningArtifactParametersArgs;
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
        var example = new Product("example", ProductArgs.builder()        
            .owner("example-owner")
            .provisioningArtifactParameters(ProductProvisioningArtifactParametersArgs.builder()
                .templateUrl("https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/temp1.json")
                .build())
            .tags(Map.of("foo", "bar"))
            .type("CLOUD_FORMATION_TEMPLATE")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicecatalog:Product
    properties:
      owner: example-owner
      provisioningArtifactParameters:
        templateUrl: https://s3.amazonaws.com/cf-templates-ozkq9d3hgiq2-us-east-1/temp1.json
      tags:
        foo: bar
      type: CLOUD_FORMATION_TEMPLATE
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_servicecatalog_product` using the product ID. For example:

```sh
 $ pulumi import aws:servicecatalog/product:Product example prod-dnigbtea24ste
```
 