Manages a Service Catalog Tag Option Resource Association.

> **Tip:** A "resource" is either a Service Catalog portfolio or product.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicecatalog.TagOptionResourceAssociation("example", {
    resourceId: "prod-dnigbtea24ste",
    tagOptionId: "tag-pjtvyakdlyo3m",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.TagOptionResourceAssociation("example",
    resource_id="prod-dnigbtea24ste",
    tag_option_id="tag-pjtvyakdlyo3m")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceCatalog.TagOptionResourceAssociation("example", new()
    {
        ResourceId = "prod-dnigbtea24ste",
        TagOptionId = "tag-pjtvyakdlyo3m",
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
		_, err := servicecatalog.NewTagOptionResourceAssociation(ctx, "example", &servicecatalog.TagOptionResourceAssociationArgs{
			ResourceId:  pulumi.String("prod-dnigbtea24ste"),
			TagOptionId: pulumi.String("tag-pjtvyakdlyo3m"),
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
import com.pulumi.aws.servicecatalog.TagOptionResourceAssociation;
import com.pulumi.aws.servicecatalog.TagOptionResourceAssociationArgs;
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
        var example = new TagOptionResourceAssociation("example", TagOptionResourceAssociationArgs.builder()        
            .resourceId("prod-dnigbtea24ste")
            .tagOptionId("tag-pjtvyakdlyo3m")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicecatalog:TagOptionResourceAssociation
    properties:
      resourceId: prod-dnigbtea24ste
      tagOptionId: tag-pjtvyakdlyo3m
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_servicecatalog_tag_option_resource_association` using the tag option ID and resource ID. For example:

```sh
 $ pulumi import aws:servicecatalog/tagOptionResourceAssociation:TagOptionResourceAssociation example tag-pjtvyakdlyo3m:prod-dnigbtea24ste
```
 