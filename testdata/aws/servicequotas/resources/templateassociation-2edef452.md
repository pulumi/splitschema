Resource for managing an AWS Service Quotas Template Association.

> Only the management account of an organization can associate Service Quota templates, and this must be done from the `us-east-1` region.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicequotas.TemplateAssociation("example", {});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicequotas.TemplateAssociation("example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceQuotas.TemplateAssociation("example");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicequotas"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicequotas.NewTemplateAssociation(ctx, "example", nil)
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
import com.pulumi.aws.servicequotas.TemplateAssociation;
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
        var example = new TemplateAssociation("example");

    }
}
```
```yaml
resources:
  example:
    type: aws:servicequotas:TemplateAssociation
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Service Quotas Template Association using the `id`. For example:

```sh
 $ pulumi import aws:servicequotas/templateAssociation:TemplateAssociation example 012345678901
```
 