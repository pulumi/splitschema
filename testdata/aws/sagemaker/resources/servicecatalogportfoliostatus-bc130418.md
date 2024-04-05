Manages status of Service Catalog in SageMaker. Service Catalog is used to create SageMaker projects.

{{% examples %}}
## Example Usage
{{% example %}}

Usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.ServicecatalogPortfolioStatus("example", {status: "Enabled"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.ServicecatalogPortfolioStatus("example", status="Enabled")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.ServicecatalogPortfolioStatus("example", new()
    {
        Status = "Enabled",
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
		_, err := sagemaker.NewServicecatalogPortfolioStatus(ctx, "example", &sagemaker.ServicecatalogPortfolioStatusArgs{
			Status: pulumi.String("Enabled"),
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
import com.pulumi.aws.sagemaker.ServicecatalogPortfolioStatus;
import com.pulumi.aws.sagemaker.ServicecatalogPortfolioStatusArgs;
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
        var example = new ServicecatalogPortfolioStatus("example", ServicecatalogPortfolioStatusArgs.builder()        
            .status("Enabled")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:ServicecatalogPortfolioStatus
    properties:
      status: Enabled
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import models using the `id`. For example:

```sh
 $ pulumi import aws:sagemaker/servicecatalogPortfolioStatus:ServicecatalogPortfolioStatus example us-east-1
```
 