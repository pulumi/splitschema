> **Note:** There is only a single subscription allowed per account.

To help you understand the charges for your Spot instances, Amazon EC2 provides a data feed that describes your Spot instance usage and pricing.
This data feed is sent to an Amazon S3 bucket that you specify when you subscribe to the data feed.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultBucketV2 = new aws.s3.BucketV2("defaultBucketV2", {});
const defaultSpotDatafeedSubscription = new aws.ec2.SpotDatafeedSubscription("defaultSpotDatafeedSubscription", {
    bucket: defaultBucketV2.id,
    prefix: "my_subdirectory",
});
```
```python
import pulumi
import pulumi_aws as aws

default_bucket_v2 = aws.s3.BucketV2("defaultBucketV2")
default_spot_datafeed_subscription = aws.ec2.SpotDatafeedSubscription("defaultSpotDatafeedSubscription",
    bucket=default_bucket_v2.id,
    prefix="my_subdirectory")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultBucketV2 = new Aws.S3.BucketV2("defaultBucketV2");

    var defaultSpotDatafeedSubscription = new Aws.Ec2.SpotDatafeedSubscription("defaultSpotDatafeedSubscription", new()
    {
        Bucket = defaultBucketV2.Id,
        Prefix = "my_subdirectory",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultBucketV2, err := s3.NewBucketV2(ctx, "defaultBucketV2", nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewSpotDatafeedSubscription(ctx, "defaultSpotDatafeedSubscription", &ec2.SpotDatafeedSubscriptionArgs{
			Bucket: defaultBucketV2.ID(),
			Prefix: pulumi.String("my_subdirectory"),
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.ec2.SpotDatafeedSubscription;
import com.pulumi.aws.ec2.SpotDatafeedSubscriptionArgs;
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
        var defaultBucketV2 = new BucketV2("defaultBucketV2");

        var defaultSpotDatafeedSubscription = new SpotDatafeedSubscription("defaultSpotDatafeedSubscription", SpotDatafeedSubscriptionArgs.builder()        
            .bucket(defaultBucketV2.id())
            .prefix("my_subdirectory")
            .build());

    }
}
```
```yaml
resources:
  defaultBucketV2:
    type: aws:s3:BucketV2
  defaultSpotDatafeedSubscription:
    type: aws:ec2:SpotDatafeedSubscription
    properties:
      bucket: ${defaultBucketV2.id}
      prefix: my_subdirectory
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a Spot Datafeed Subscription using the word `spot-datafeed-subscription`. For example:

```sh
 $ pulumi import aws:ec2/spotDatafeedSubscription:SpotDatafeedSubscription mysubscription spot-datafeed-subscription
```
 