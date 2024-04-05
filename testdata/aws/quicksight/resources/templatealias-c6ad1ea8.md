Resource for managing an AWS QuickSight Template Alias.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.TemplateAlias("example", {
    aliasName: "example-alias",
    templateId: aws_quicksight_template.test.template_id,
    templateVersionNumber: aws_quicksight_template.test.version_number,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.TemplateAlias("example",
    alias_name="example-alias",
    template_id=aws_quicksight_template["test"]["template_id"],
    template_version_number=aws_quicksight_template["test"]["version_number"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.TemplateAlias("example", new()
    {
        AliasName = "example-alias",
        TemplateId = aws_quicksight_template.Test.Template_id,
        TemplateVersionNumber = aws_quicksight_template.Test.Version_number,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewTemplateAlias(ctx, "example", &quicksight.TemplateAliasArgs{
			AliasName:             pulumi.String("example-alias"),
			TemplateId:            pulumi.Any(aws_quicksight_template.Test.Template_id),
			TemplateVersionNumber: pulumi.Any(aws_quicksight_template.Test.Version_number),
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
import com.pulumi.aws.quicksight.TemplateAlias;
import com.pulumi.aws.quicksight.TemplateAliasArgs;
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
        var example = new TemplateAlias("example", TemplateAliasArgs.builder()        
            .aliasName("example-alias")
            .templateId(aws_quicksight_template.test().template_id())
            .templateVersionNumber(aws_quicksight_template.test().version_number())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:TemplateAlias
    properties:
      aliasName: example-alias
      templateId: ${aws_quicksight_template.test.template_id}
      templateVersionNumber: ${aws_quicksight_template.test.version_number}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import QuickSight Template Alias using the AWS account ID, template ID, and alias name separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:quicksight/templateAlias:TemplateAlias example 123456789012,example-id,example-alias
```
 