Retrieve information about a Direct Connect Connection.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.directconnect.getConnection({
    name: "tf-dx-connection",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directconnect.get_connection(name="tf-dx-connection")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.DirectConnect.GetConnection.Invoke(new()
    {
        Name = "tf-dx-connection",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.LookupConnection(ctx, &directconnect.LookupConnectionArgs{
			Name: "tf-dx-connection",
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
import com.pulumi.aws.directconnect.DirectconnectFunctions;
import com.pulumi.aws.directconnect.inputs.GetConnectionArgs;
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
        final var example = DirectconnectFunctions.getConnection(GetConnectionArgs.builder()
            .name("tf-dx-connection")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:directconnect:getConnection
      Arguments:
        name: tf-dx-connection
```
{{% /example %}}
{{% /examples %}}