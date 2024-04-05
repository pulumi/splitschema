Provides a Network Insights Path resource. Part of the "Reachability Analyzer" service in the AWS VPC console.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ec2.NetworkInsightsPath("test", {
    source: aws_network_interface.source.id,
    destination: aws_network_interface.destination.id,
    protocol: "tcp",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ec2.NetworkInsightsPath("test",
    source=aws_network_interface["source"]["id"],
    destination=aws_network_interface["destination"]["id"],
    protocol="tcp")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Ec2.NetworkInsightsPath("test", new()
    {
        Source = aws_network_interface.Source.Id,
        Destination = aws_network_interface.Destination.Id,
        Protocol = "tcp",
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
		_, err := ec2.NewNetworkInsightsPath(ctx, "test", &ec2.NetworkInsightsPathArgs{
			Source:      pulumi.Any(aws_network_interface.Source.Id),
			Destination: pulumi.Any(aws_network_interface.Destination.Id),
			Protocol:    pulumi.String("tcp"),
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
        var test = new NetworkInsightsPath("test", NetworkInsightsPathArgs.builder()        
            .source(aws_network_interface.source().id())
            .destination(aws_network_interface.destination().id())
            .protocol("tcp")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:ec2:NetworkInsightsPath
    properties:
      source: ${aws_network_interface.source.id}
      destination: ${aws_network_interface.destination.id}
      protocol: tcp
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Network Insights Paths using the `id`. For example:

```sh
 $ pulumi import aws:ec2/networkInsightsPath:NetworkInsightsPath test nip-00edfba169923aefd
```
 