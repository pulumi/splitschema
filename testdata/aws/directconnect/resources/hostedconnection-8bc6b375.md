Provides a hosted connection on the specified interconnect or a link aggregation group (LAG) of interconnects. Intended for use by AWS Direct Connect Partners only.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const hosted = new aws.directconnect.HostedConnection("hosted", {
    bandwidth: "100Mbps",
    connectionId: "dxcon-ffabc123",
    ownerAccountId: "123456789012",
    vlan: 1,
});
```
```python
import pulumi
import pulumi_aws as aws

hosted = aws.directconnect.HostedConnection("hosted",
    bandwidth="100Mbps",
    connection_id="dxcon-ffabc123",
    owner_account_id="123456789012",
    vlan=1)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var hosted = new Aws.DirectConnect.HostedConnection("hosted", new()
    {
        Bandwidth = "100Mbps",
        ConnectionId = "dxcon-ffabc123",
        OwnerAccountId = "123456789012",
        Vlan = 1,
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
		_, err := directconnect.NewHostedConnection(ctx, "hosted", &directconnect.HostedConnectionArgs{
			Bandwidth:      pulumi.String("100Mbps"),
			ConnectionId:   pulumi.String("dxcon-ffabc123"),
			OwnerAccountId: pulumi.String("123456789012"),
			Vlan:           pulumi.Int(1),
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
import com.pulumi.aws.directconnect.HostedConnection;
import com.pulumi.aws.directconnect.HostedConnectionArgs;
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
        var hosted = new HostedConnection("hosted", HostedConnectionArgs.builder()        
            .bandwidth("100Mbps")
            .connectionId("dxcon-ffabc123")
            .ownerAccountId("123456789012")
            .vlan(1)
            .build());

    }
}
```
```yaml
resources:
  hosted:
    type: aws:directconnect:HostedConnection
    properties:
      bandwidth: 100Mbps
      connectionId: dxcon-ffabc123
      ownerAccountId: '123456789012'
      vlan: 1
```
{{% /example %}}
{{% /examples %}}