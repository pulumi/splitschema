Provides a SageMaker endpoint configuration resource.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ec = new aws.sagemaker.EndpointConfiguration("ec", {
    productionVariants: [{
        variantName: "variant-1",
        modelName: aws_sagemaker_model.m.name,
        initialInstanceCount: 1,
        instanceType: "ml.t2.medium",
    }],
    tags: {
        Name: "foo",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

ec = aws.sagemaker.EndpointConfiguration("ec",
    production_variants=[aws.sagemaker.EndpointConfigurationProductionVariantArgs(
        variant_name="variant-1",
        model_name=aws_sagemaker_model["m"]["name"],
        initial_instance_count=1,
        instance_type="ml.t2.medium",
    )],
    tags={
        "Name": "foo",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ec = new Aws.Sagemaker.EndpointConfiguration("ec", new()
    {
        ProductionVariants = new[]
        {
            new Aws.Sagemaker.Inputs.EndpointConfigurationProductionVariantArgs
            {
                VariantName = "variant-1",
                ModelName = aws_sagemaker_model.M.Name,
                InitialInstanceCount = 1,
                InstanceType = "ml.t2.medium",
            },
        },
        Tags = 
        {
            { "Name", "foo" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewEndpointConfiguration(ctx, "ec", &sagemaker.EndpointConfigurationArgs{
			ProductionVariants: sagemaker.EndpointConfigurationProductionVariantArray{
				&sagemaker.EndpointConfigurationProductionVariantArgs{
					VariantName:          pulumi.String("variant-1"),
					ModelName:            pulumi.Any(aws_sagemaker_model.M.Name),
					InitialInstanceCount: pulumi.Int(1),
					InstanceType:         pulumi.String("ml.t2.medium"),
				},
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("foo"),
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
import com.pulumi.aws.sagemaker.EndpointConfiguration;
import com.pulumi.aws.sagemaker.EndpointConfigurationArgs;
import com.pulumi.aws.sagemaker.inputs.EndpointConfigurationProductionVariantArgs;
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
        var ec = new EndpointConfiguration("ec", EndpointConfigurationArgs.builder()        
            .productionVariants(EndpointConfigurationProductionVariantArgs.builder()
                .variantName("variant-1")
                .modelName(aws_sagemaker_model.m().name())
                .initialInstanceCount(1)
                .instanceType("ml.t2.medium")
                .build())
            .tags(Map.of("Name", "foo"))
            .build());

    }
}
```
```yaml
resources:
  ec:
    type: aws:sagemaker:EndpointConfiguration
    properties:
      productionVariants:
        - variantName: variant-1
          modelName: ${aws_sagemaker_model.m.name}
          initialInstanceCount: 1
          instanceType: ml.t2.medium
      tags:
        Name: foo
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import endpoint configurations using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/endpointConfiguration:EndpointConfiguration test_endpoint_config endpoint-config-foo
```
 