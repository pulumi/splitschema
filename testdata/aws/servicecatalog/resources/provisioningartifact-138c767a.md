Manages a Service Catalog Provisioning Artifact for a specified product.

> A "provisioning artifact" is also referred to as a "version."

> **NOTE:** You cannot create a provisioning artifact for a product that was shared with you.

> **NOTE:** The user or role that use this resource must have the `cloudformation:GetTemplate` IAM policy permission. This policy permission is required when using the `template_physical_id` argument.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicecatalog.ProvisioningArtifact("example", {
    productId: aws_servicecatalog_product.example.id,
    type: "CLOUD_FORMATION_TEMPLATE",
    templateUrl: `https://${aws_s3_bucket.example.bucket_regional_domain_name}/${aws_s3_object.example.key}`,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.ProvisioningArtifact("example",
    product_id=aws_servicecatalog_product["example"]["id"],
    type="CLOUD_FORMATION_TEMPLATE",
    template_url=f"https://{aws_s3_bucket['example']['bucket_regional_domain_name']}/{aws_s3_object['example']['key']}")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceCatalog.ProvisioningArtifact("example", new()
    {
        ProductId = aws_servicecatalog_product.Example.Id,
        Type = "CLOUD_FORMATION_TEMPLATE",
        TemplateUrl = $"https://{aws_s3_bucket.Example.Bucket_regional_domain_name}/{aws_s3_object.Example.Key}",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicecatalog"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicecatalog.NewProvisioningArtifact(ctx, "example", &servicecatalog.ProvisioningArtifactArgs{
			ProductId:   pulumi.Any(aws_servicecatalog_product.Example.Id),
			Type:        pulumi.String("CLOUD_FORMATION_TEMPLATE"),
			TemplateUrl: pulumi.String(fmt.Sprintf("https://%v/%v", aws_s3_bucket.Example.Bucket_regional_domain_name, aws_s3_object.Example.Key)),
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
import com.pulumi.aws.servicecatalog.ProvisioningArtifact;
import com.pulumi.aws.servicecatalog.ProvisioningArtifactArgs;
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
        var example = new ProvisioningArtifact("example", ProvisioningArtifactArgs.builder()        
            .productId(aws_servicecatalog_product.example().id())
            .type("CLOUD_FORMATION_TEMPLATE")
            .templateUrl(String.format("https://%s/%s", aws_s3_bucket.example().bucket_regional_domain_name(),aws_s3_object.example().key()))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicecatalog:ProvisioningArtifact
    properties:
      productId: ${aws_servicecatalog_product.example.id}
      type: CLOUD_FORMATION_TEMPLATE
      templateUrl: https://${aws_s3_bucket.example.bucket_regional_domain_name}/${aws_s3_object.example.key}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_servicecatalog_provisioning_artifact` using the provisioning artifact ID and product ID separated by a colon. For example:

```sh
 $ pulumi import aws:servicecatalog/provisioningArtifact:ProvisioningArtifact example pa-ij2b6lusy6dec:prod-el3an0rma3
```
 