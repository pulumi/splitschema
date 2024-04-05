Provides a Network Insights Analysis resource. Part of the "Reachability Analyzer" service in the AWS VPC console.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const path = new aws.ec2.NetworkInsightsPath("path", {
    source: aws_network_interface.source.id,
    destination: aws_network_interface.destination.id,
    protocol: "tcp",
});
const analysis = new aws.ec2.NetworkInsightsAnalysis("analysis", {networkInsightsPathId: path.id});
```
```python
import pulumi
import pulumi_aws as aws

path = aws.ec2.NetworkInsightsPath("path",
    source=aws_network_interface["source"]["id"],
    destination=aws_network_interface["destination"]["id"],
    protocol="tcp")
analysis = aws.ec2.NetworkInsightsAnalysis("analysis", network_insights_path_id=path.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var path = new Aws.Ec2.NetworkInsightsPath("path", new()
    {
        Source = aws_network_interface.Source.Id,
        Destination = aws_network_interface.Destination.Id,
        Protocol = "tcp",
    });

    var analysis = new Aws.Ec2.NetworkInsightsAnalysis("analysis", new()
    {
        NetworkInsightsPathId = path.Id,
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
		path, err := ec2.NewNetworkInsightsPath(ctx, "path", &ec2.NetworkInsightsPathArgs{
			Source:      pulumi.Any(aws_network_interface.Source.Id),
			Destination: pulumi.Any(aws_network_interface.Destination.Id),
			Protocol:    pulumi.String("tcp"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewNetworkInsightsAnalysis(ctx, "analysis", &ec2.NetworkInsightsAnalysisArgs{
			NetworkInsightsPathId: path.ID(),
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
import com.pulumi.aws.ec2.NetworkInsightsPath;
import com.pulumi.aws.ec2.NetworkInsightsPathArgs;
import com.pulumi.aws.ec2.NetworkInsightsAnalysis;
import com.pulumi.aws.ec2.NetworkInsightsAnalysisArgs;
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
        var path = new NetworkInsightsPath("path", NetworkInsightsPathArgs.builder()        
            .source(aws_network_interface.source().id())
            .destination(aws_network_interface.destination().id())
            .protocol("tcp")
            .build());

        var analysis = new NetworkInsightsAnalysis("analysis", NetworkInsightsAnalysisArgs.builder()        
            .networkInsightsPathId(path.id())
            .build());

    }
}
```
```yaml
resources:
  path:
    type: aws:ec2:NetworkInsightsPath
    properties:
      source: ${aws_network_interface.source.id}
      destination: ${aws_network_interface.destination.id}
      protocol: tcp
  analysis:
    type: aws:ec2:NetworkInsightsAnalysis
    properties:
      networkInsightsPathId: ${path.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Network Insights Analyses using the `id`. For example:

```sh
 $ pulumi import aws:ec2/networkInsightsAnalysis:NetworkInsightsAnalysis test nia-0462085c957f11a55
```
 