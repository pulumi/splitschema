Provides a confirmation of the creation of the specified hosted connection on an interconnect.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const confirmation = new aws.directconnect.ConnectionConfirmation("confirmation", {connectionId: "dxcon-ffabc123"});
```
```python
import pulumi
import pulumi_aws as aws

confirmation = aws.directconnect.ConnectionConfirmation("confirmation", connection_id="dxcon-ffabc123")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var confirmation = new Aws.DirectConnect.ConnectionConfirmation("confirmation", new()
    {
        ConnectionId = "dxcon-ffabc123",
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
		_, err := directconnect.NewConnectionConfirmation(ctx, "confirmation", &directconnect.ConnectionConfirmationArgs{
			ConnectionId: pulumi.String("dxcon-ffabc123"),
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
import com.pulumi.aws.directconnect.ConnectionConfirmation;
import com.pulumi.aws.directconnect.ConnectionConfirmationArgs;
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
        var confirmation = new ConnectionConfirmation("confirmation", ConnectionConfirmationArgs.builder()        
            .connectionId("dxcon-ffabc123")
            .build());

    }
}
```
```yaml
resources:
  confirmation:
    type: aws:directconnect:ConnectionConfirmation
    properties:
      connectionId: dxcon-ffabc123
```
{{% /example %}}
{{% /examples %}}