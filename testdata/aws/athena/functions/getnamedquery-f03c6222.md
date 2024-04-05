Provides an Athena Named Query data source.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.athena.getNamedQuery({
    name: "athenaQueryName",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.athena.get_named_query(name="athenaQueryName")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Athena.GetNamedQuery.Invoke(new()
    {
        Name = "athenaQueryName",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/athena"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := athena.LookupNamedQuery(ctx, &athena.LookupNamedQueryArgs{
			Name: "athenaQueryName",
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
import com.pulumi.aws.athena.AthenaFunctions;
import com.pulumi.aws.athena.inputs.GetNamedQueryArgs;
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
        final var example = AthenaFunctions.getNamedQuery(GetNamedQueryArgs.builder()
            .name("athenaQueryName")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:athena:getNamedQuery
      Arguments:
        name: athenaQueryName
```
{{% /example %}}
{{% /examples %}}