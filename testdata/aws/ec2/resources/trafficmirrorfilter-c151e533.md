Provides an Traffic mirror filter.  
Read [limits and considerations](https://docs.aws.amazon.com/vpc/latest/mirroring/traffic-mirroring-considerations.html) for traffic mirroring

{{% examples %}}
## Example Usage
{{% example %}}

To create a basic traffic mirror filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ec2.TrafficMirrorFilter("foo", {
    description: "traffic mirror filter - example",
    networkServices: ["amazon-dns"],
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ec2.TrafficMirrorFilter("foo",
    description="traffic mirror filter - example",
    network_services=["amazon-dns"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ec2.TrafficMirrorFilter("foo", new()
    {
        Description = "traffic mirror filter - example",
        NetworkServices = new[]
        {
            "amazon-dns",
        },
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
		_, err := ec2.NewTrafficMirrorFilter(ctx, "foo", &ec2.TrafficMirrorFilterArgs{
			Description: pulumi.String("traffic mirror filter - example"),
			NetworkServices: pulumi.StringArray{
				pulumi.String("amazon-dns"),
			},
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
import com.pulumi.aws.ec2.TrafficMirrorFilter;
import com.pulumi.aws.ec2.TrafficMirrorFilterArgs;
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
        var foo = new TrafficMirrorFilter("foo", TrafficMirrorFilterArgs.builder()        
            .description("traffic mirror filter - example")
            .networkServices("amazon-dns")
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ec2:TrafficMirrorFilter
    properties:
      description: traffic mirror filter - example
      networkServices:
        - amazon-dns
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import traffic mirror filter using the `id`. For example:

```sh
 $ pulumi import aws:ec2/trafficMirrorFilter:TrafficMirrorFilter foo tmf-0fbb93ddf38198f64
```
 