Associates a Direct Connect Connection with a LAG.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleConnection = new aws.directconnect.Connection("exampleConnection", {
    bandwidth: "1Gbps",
    location: "EqSe2-EQ",
});
const exampleLinkAggregationGroup = new aws.directconnect.LinkAggregationGroup("exampleLinkAggregationGroup", {
    connectionsBandwidth: "1Gbps",
    location: "EqSe2-EQ",
});
const exampleConnectionAssociation = new aws.directconnect.ConnectionAssociation("exampleConnectionAssociation", {
    connectionId: exampleConnection.id,
    lagId: exampleLinkAggregationGroup.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_connection = aws.directconnect.Connection("exampleConnection",
    bandwidth="1Gbps",
    location="EqSe2-EQ")
example_link_aggregation_group = aws.directconnect.LinkAggregationGroup("exampleLinkAggregationGroup",
    connections_bandwidth="1Gbps",
    location="EqSe2-EQ")
example_connection_association = aws.directconnect.ConnectionAssociation("exampleConnectionAssociation",
    connection_id=example_connection.id,
    lag_id=example_link_aggregation_group.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleConnection = new Aws.DirectConnect.Connection("exampleConnection", new()
    {
        Bandwidth = "1Gbps",
        Location = "EqSe2-EQ",
    });

    var exampleLinkAggregationGroup = new Aws.DirectConnect.LinkAggregationGroup("exampleLinkAggregationGroup", new()
    {
        ConnectionsBandwidth = "1Gbps",
        Location = "EqSe2-EQ",
    });

    var exampleConnectionAssociation = new Aws.DirectConnect.ConnectionAssociation("exampleConnectionAssociation", new()
    {
        ConnectionId = exampleConnection.Id,
        LagId = exampleLinkAggregationGroup.Id,
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
		exampleConnection, err := directconnect.NewConnection(ctx, "exampleConnection", &directconnect.ConnectionArgs{
			Bandwidth: pulumi.String("1Gbps"),
			Location:  pulumi.String("EqSe2-EQ"),
		})
		if err != nil {
			return err
		}
		exampleLinkAggregationGroup, err := directconnect.NewLinkAggregationGroup(ctx, "exampleLinkAggregationGroup", &directconnect.LinkAggregationGroupArgs{
			ConnectionsBandwidth: pulumi.String("1Gbps"),
			Location:             pulumi.String("EqSe2-EQ"),
		})
		if err != nil {
			return err
		}
		_, err = directconnect.NewConnectionAssociation(ctx, "exampleConnectionAssociation", &directconnect.ConnectionAssociationArgs{
			ConnectionId: exampleConnection.ID(),
			LagId:        exampleLinkAggregationGroup.ID(),
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
import com.pulumi.aws.directconnect.Connection;
import com.pulumi.aws.directconnect.ConnectionArgs;
import com.pulumi.aws.directconnect.LinkAggregationGroup;
import com.pulumi.aws.directconnect.LinkAggregationGroupArgs;
import com.pulumi.aws.directconnect.ConnectionAssociation;
import com.pulumi.aws.directconnect.ConnectionAssociationArgs;
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
        var exampleConnection = new Connection("exampleConnection", ConnectionArgs.builder()        
            .bandwidth("1Gbps")
            .location("EqSe2-EQ")
            .build());

        var exampleLinkAggregationGroup = new LinkAggregationGroup("exampleLinkAggregationGroup", LinkAggregationGroupArgs.builder()        
            .connectionsBandwidth("1Gbps")
            .location("EqSe2-EQ")
            .build());

        var exampleConnectionAssociation = new ConnectionAssociation("exampleConnectionAssociation", ConnectionAssociationArgs.builder()        
            .connectionId(exampleConnection.id())
            .lagId(exampleLinkAggregationGroup.id())
            .build());

    }
}
```
```yaml
resources:
  exampleConnection:
    type: aws:directconnect:Connection
    properties:
      bandwidth: 1Gbps
      location: EqSe2-EQ
  exampleLinkAggregationGroup:
    type: aws:directconnect:LinkAggregationGroup
    properties:
      connectionsBandwidth: 1Gbps
      location: EqSe2-EQ
  exampleConnectionAssociation:
    type: aws:directconnect:ConnectionAssociation
    properties:
      connectionId: ${exampleConnection.id}
      lagId: ${exampleLinkAggregationGroup.id}
```
{{% /example %}}
{{% /examples %}}