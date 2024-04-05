Provides a Direct Connect LAG. Connections can be added to the LAG via the `aws.directconnect.Connection` and `aws.directconnect.ConnectionAssociation` resources.

> *NOTE:* When creating a LAG, if no existing connection is specified, Direct Connect will create a connection and this provider will remove this unmanaged connection during resource creation.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const hoge = new aws.directconnect.LinkAggregationGroup("hoge", {
    connectionsBandwidth: "1Gbps",
    forceDestroy: true,
    location: "EqDC2",
});
```
```python
import pulumi
import pulumi_aws as aws

hoge = aws.directconnect.LinkAggregationGroup("hoge",
    connections_bandwidth="1Gbps",
    force_destroy=True,
    location="EqDC2")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var hoge = new Aws.DirectConnect.LinkAggregationGroup("hoge", new()
    {
        ConnectionsBandwidth = "1Gbps",
        ForceDestroy = true,
        Location = "EqDC2",
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
		_, err := directconnect.NewLinkAggregationGroup(ctx, "hoge", &directconnect.LinkAggregationGroupArgs{
			ConnectionsBandwidth: pulumi.String("1Gbps"),
			ForceDestroy:         pulumi.Bool(true),
			Location:             pulumi.String("EqDC2"),
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
import com.pulumi.aws.directconnect.LinkAggregationGroup;
import com.pulumi.aws.directconnect.LinkAggregationGroupArgs;
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
        var hoge = new LinkAggregationGroup("hoge", LinkAggregationGroupArgs.builder()        
            .connectionsBandwidth("1Gbps")
            .forceDestroy(true)
            .location("EqDC2")
            .build());

    }
}
```
```yaml
resources:
  hoge:
    type: aws:directconnect:LinkAggregationGroup
    properties:
      connectionsBandwidth: 1Gbps
      forceDestroy: true
      location: EqDC2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Direct Connect LAGs using the LAG `id`. For example:

```sh
 $ pulumi import aws:directconnect/linkAggregationGroup:LinkAggregationGroup test_lag dxlag-fgnsp5rq
```
 