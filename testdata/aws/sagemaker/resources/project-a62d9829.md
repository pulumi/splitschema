Provides a SageMaker Project resource.

 > Note: If you are trying to use SageMaker projects with SageMaker studio you will need to add a tag with the key `sagemaker:studio-visibility` with value `true`. For more on requirements to use projects and permission needed see [AWS Docs](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-projects-templates-custom.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.Project("example", {
    projectName: "example",
    serviceCatalogProvisioningDetails: {
        productId: aws_servicecatalog_product.example.id,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.Project("example",
    project_name="example",
    service_catalog_provisioning_details=aws.sagemaker.ProjectServiceCatalogProvisioningDetailsArgs(
        product_id=aws_servicecatalog_product["example"]["id"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.Project("example", new()
    {
        ProjectName = "example",
        ServiceCatalogProvisioningDetails = new Aws.Sagemaker.Inputs.ProjectServiceCatalogProvisioningDetailsArgs
        {
            ProductId = aws_servicecatalog_product.Example.Id,
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
		_, err := sagemaker.NewProject(ctx, "example", &sagemaker.ProjectArgs{
			ProjectName: pulumi.String("example"),
			ServiceCatalogProvisioningDetails: &sagemaker.ProjectServiceCatalogProvisioningDetailsArgs{
				ProductId: pulumi.Any(aws_servicecatalog_product.Example.Id),
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
import com.pulumi.aws.sagemaker.Project;
import com.pulumi.aws.sagemaker.ProjectArgs;
import com.pulumi.aws.sagemaker.inputs.ProjectServiceCatalogProvisioningDetailsArgs;
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
        var example = new Project("example", ProjectArgs.builder()        
            .projectName("example")
            .serviceCatalogProvisioningDetails(ProjectServiceCatalogProvisioningDetailsArgs.builder()
                .productId(aws_servicecatalog_product.example().id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:Project
    properties:
      projectName: example
      serviceCatalogProvisioningDetails:
        productId: ${aws_servicecatalog_product.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Projects using the `project_name`. For example:

```sh
 $ pulumi import aws:sagemaker/project:Project example example
```
 