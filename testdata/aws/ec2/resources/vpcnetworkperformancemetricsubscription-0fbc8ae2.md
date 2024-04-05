Provides a resource to manage an Infrastructure Performance subscription.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.VpcNetworkPerformanceMetricSubscription("example", {
    destination: "us-west-1",
    source: "us-east-1",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.VpcNetworkPerformanceMetricSubscription("example",
    destination="us-west-1",
    source="us-east-1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.VpcNetworkPerformanceMetricSubscription("example", new()
    {
        Destination = "us-west-1",
        Source = "us-east-1",
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
		_, err := ec2.NewVpcNetworkPerformanceMetricSubscription(ctx, "example", &ec2.VpcNetworkPerformanceMetricSubscriptionArgs{
			Destination: pulumi.String("us-west-1"),
			Source:      pulumi.String("us-east-1"),
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
import com.pulumi.aws.ec2.VpcNetworkPerformanceMetricSubscription;
import com.pulumi.aws.ec2.VpcNetworkPerformanceMetricSubscriptionArgs;
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
        var example = new VpcNetworkPerformanceMetricSubscription("example", VpcNetworkPerformanceMetricSubscriptionArgs.builder()        
            .destination("us-west-1")
            .source("us-east-1")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:VpcNetworkPerformanceMetricSubscription
    properties:
      destination: us-west-1
      source: us-east-1
```
{{% /example %}}
{{% /examples %}}