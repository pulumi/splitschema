Use this data source to get the ARN of an AWS Transfer Server for use in other
resources.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.transfer.getServer({
    serverId: "s-1234567",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.get_server(server_id="s-1234567")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Transfer.GetServer.Invoke(new()
    {
        ServerId = "s-1234567",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.LookupServer(ctx, &transfer.LookupServerArgs{
			ServerId: "s-1234567",
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
import com.pulumi.aws.transfer.TransferFunctions;
import com.pulumi.aws.transfer.inputs.GetServerArgs;
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
        final var example = TransferFunctions.getServer(GetServerArgs.builder()
            .serverId("s-1234567")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:transfer:getServer
      Arguments:
        serverId: s-1234567
```
{{% /example %}}
{{% /examples %}}