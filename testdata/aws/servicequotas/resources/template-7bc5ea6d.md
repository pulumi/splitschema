Resource for managing an AWS Service Quotas Template.

> Only the management account of an organization can alter Service Quota templates, and this must be done from the `us-east-1` region.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicequotas.Template("example", {
    quotaCode: "L-2ACBD22F",
    region: "us-east-1",
    serviceCode: "lambda",
    value: 80,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicequotas.Template("example",
    quota_code="L-2ACBD22F",
    region="us-east-1",
    service_code="lambda",
    value=80)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceQuotas.Template("example", new()
    {
        QuotaCode = "L-2ACBD22F",
        Region = "us-east-1",
        ServiceCode = "lambda",
        Value = 80,
    });

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
		_, err := servicequotas.NewTemplate(ctx, "example", &servicequotas.TemplateArgs{
			QuotaCode:   pulumi.String("L-2ACBD22F"),
			Region:      pulumi.String("us-east-1"),
			ServiceCode: pulumi.String("lambda"),
			Value:       pulumi.Float64(80),
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
import com.pulumi.aws.servicequotas.Template;
import com.pulumi.aws.servicequotas.TemplateArgs;
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
        var example = new Template("example", TemplateArgs.builder()        
            .quotaCode("L-2ACBD22F")
            .region("us-east-1")
            .serviceCode("lambda")
            .value("80")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicequotas:Template
    properties:
      quotaCode: L-2ACBD22F
      # function and layer storage (default: 75 GB)
      region: us-east-1
      serviceCode: lambda
      value: '80'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Service Quotas Template using the `id`. For example:

```sh
 $ pulumi import aws:servicequotas/template:Template example us-east-1,L-2ACBD22F,lambda
```
 