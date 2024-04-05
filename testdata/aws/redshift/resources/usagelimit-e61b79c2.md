Creates a new Amazon Redshift Usage Limit.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshift.UsageLimit("example", {
    clusterIdentifier: aws_redshift_cluster.example.id,
    featureType: "concurrency-scaling",
    limitType: "time",
    amount: 60,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.UsageLimit("example",
    cluster_identifier=aws_redshift_cluster["example"]["id"],
    feature_type="concurrency-scaling",
    limit_type="time",
    amount=60)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedShift.UsageLimit("example", new()
    {
        ClusterIdentifier = aws_redshift_cluster.Example.Id,
        FeatureType = "concurrency-scaling",
        LimitType = "time",
        Amount = 60,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewUsageLimit(ctx, "example", &redshift.UsageLimitArgs{
			ClusterIdentifier: pulumi.Any(aws_redshift_cluster.Example.Id),
			FeatureType:       pulumi.String("concurrency-scaling"),
			LimitType:         pulumi.String("time"),
			Amount:            pulumi.Int(60),
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
import com.pulumi.aws.redshift.UsageLimit;
import com.pulumi.aws.redshift.UsageLimitArgs;
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
        var example = new UsageLimit("example", UsageLimitArgs.builder()        
            .clusterIdentifier(aws_redshift_cluster.example().id())
            .featureType("concurrency-scaling")
            .limitType("time")
            .amount(60)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:UsageLimit
    properties:
      clusterIdentifier: ${aws_redshift_cluster.example.id}
      featureType: concurrency-scaling
      limitType: time
      amount: 60
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift usage limits using the `id`. For example:

```sh
 $ pulumi import aws:redshift/usageLimit:UsageLimit example example-id
```
 