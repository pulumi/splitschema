This data source can be used to fetch information about a specific Glue Connection.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.glue.getConnection({
    id: "123456789123:connection",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.get_connection(id="123456789123:connection")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Glue.GetConnection.Invoke(new()
    {
        Id = "123456789123:connection",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.LookupConnection(ctx, &glue.LookupConnectionArgs{
			Id: "123456789123:connection",
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
import com.pulumi.aws.glue.GlueFunctions;
import com.pulumi.aws.glue.inputs.GetConnectionArgs;
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
        final var example = GlueFunctions.getConnection(GetConnectionArgs.builder()
            .id("123456789123:connection")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:glue:getConnection
      Arguments:
        id: 123456789123:connection
```
{{% /example %}}
{{% /examples %}}