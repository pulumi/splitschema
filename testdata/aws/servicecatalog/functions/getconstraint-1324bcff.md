Provides information on a Service Catalog Constraint.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.servicecatalog.getConstraint({
    acceptLanguage: "en",
    id: "cons-hrvy0335",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.get_constraint(accept_language="en",
    id="cons-hrvy0335")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.ServiceCatalog.GetConstraint.Invoke(new()
    {
        AcceptLanguage = "en",
        Id = "cons-hrvy0335",
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
		_, err := servicecatalog.LookupConstraint(ctx, &servicecatalog.LookupConstraintArgs{
			AcceptLanguage: pulumi.StringRef("en"),
			Id:             "cons-hrvy0335",
		}, nil)
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
import com.pulumi.aws.servicecatalog.ServicecatalogFunctions;
import com.pulumi.aws.servicecatalog.inputs.GetConstraintArgs;
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
        final var example = ServicecatalogFunctions.getConstraint(GetConstraintArgs.builder()
            .acceptLanguage("en")
            .id("cons-hrvy0335")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:servicecatalog:getConstraint
      Arguments:
        acceptLanguage: en
        id: cons-hrvy0335
```
{{% /example %}}
{{% /examples %}}