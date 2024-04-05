Resource for managing an AWS VPC Lattice Service Network or Service Access log subscription.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.vpclattice.AccessLogSubscription("example", {
    resourceIdentifier: aws_vpclattice_service_network.example.id,
    destinationArn: aws_s3.bucket.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.AccessLogSubscription("example",
    resource_identifier=aws_vpclattice_service_network["example"]["id"],
    destination_arn=aws_s3["bucket"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VpcLattice.AccessLogSubscription("example", new()
    {
        ResourceIdentifier = aws_vpclattice_service_network.Example.Id,
        DestinationArn = aws_s3.Bucket.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpclattice.NewAccessLogSubscription(ctx, "example", &vpclattice.AccessLogSubscriptionArgs{
			ResourceIdentifier: pulumi.Any(aws_vpclattice_service_network.Example.Id),
			DestinationArn:     pulumi.Any(aws_s3.Bucket.Arn),
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
import com.pulumi.aws.vpclattice.AccessLogSubscription;
import com.pulumi.aws.vpclattice.AccessLogSubscriptionArgs;
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
        var example = new AccessLogSubscription("example", AccessLogSubscriptionArgs.builder()        
            .resourceIdentifier(aws_vpclattice_service_network.example().id())
            .destinationArn(aws_s3.bucket().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:vpclattice:AccessLogSubscription
    properties:
      resourceIdentifier: ${aws_vpclattice_service_network.example.id}
      destinationArn: ${aws_s3.bucket.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Lattice Access Log Subscription using the access log subscription ID. For example:

```sh
 $ pulumi import aws:vpclattice/accessLogSubscription:AccessLogSubscription example rft-8012925589
```
 