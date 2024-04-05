Provides an EC2 placement group. Read more about placement groups
in [AWS Docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/placement-groups.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const web = new aws.ec2.PlacementGroup("web", {strategy: "cluster"});
```
```python
import pulumi
import pulumi_aws as aws

web = aws.ec2.PlacementGroup("web", strategy="cluster")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var web = new Aws.Ec2.PlacementGroup("web", new()
    {
        Strategy = "cluster",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewPlacementGroup(ctx, "web", &ec2.PlacementGroupArgs{
			Strategy: pulumi.String("cluster"),
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
import com.pulumi.aws.ec2.PlacementGroup;
import com.pulumi.aws.ec2.PlacementGroupArgs;
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
        var web = new PlacementGroup("web", PlacementGroupArgs.builder()        
            .strategy("cluster")
            .build());

    }
}
```
```yaml
resources:
  web:
    type: aws:ec2:PlacementGroup
    properties:
      strategy: cluster
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import placement groups using the `name`. For example:

```sh
 $ pulumi import aws:ec2/placementGroup:PlacementGroup prod_pg production-placement-group
```
 