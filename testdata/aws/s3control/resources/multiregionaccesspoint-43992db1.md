Provides a resource to manage an S3 Multi-Region Access Point associated with specified buckets.

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}
### Multiple AWS Buckets in Different Regions

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primaryRegion = new aws.Provider("primaryRegion", {region: "us-east-1"});
const secondaryRegion = new aws.Provider("secondaryRegion", {region: "us-west-2"});
const fooBucket = new aws.s3.BucketV2("fooBucket", {}, {
    provider: aws.primary_region,
});
const barBucket = new aws.s3.BucketV2("barBucket", {}, {
    provider: aws.secondary_region,
});
const example = new aws.s3control.MultiRegionAccessPoint("example", {details: {
    name: "example",
    regions: [
        {
            bucket: fooBucket.id,
        },
        {
            bucket: barBucket.id,
        },
    ],
}});
```
```python
import pulumi
import pulumi_aws as aws

primary_region = aws.Provider("primaryRegion", region="us-east-1")
secondary_region = aws.Provider("secondaryRegion", region="us-west-2")
foo_bucket = aws.s3.BucketV2("fooBucket", opts=pulumi.ResourceOptions(provider=aws["primary_region"]))
bar_bucket = aws.s3.BucketV2("barBucket", opts=pulumi.ResourceOptions(provider=aws["secondary_region"]))
example = aws.s3control.MultiRegionAccessPoint("example", details=aws.s3control.MultiRegionAccessPointDetailsArgs(
    name="example",
    regions=[
        aws.s3control.MultiRegionAccessPointDetailsRegionArgs(
            bucket=foo_bucket.id,
        ),
        aws.s3control.MultiRegionAccessPointDetailsRegionArgs(
            bucket=bar_bucket.id,
        ),
    ],
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var primaryRegion = new Aws.Provider("primaryRegion", new()
    {
        Region = "us-east-1",
    });

    var secondaryRegion = new Aws.Provider("secondaryRegion", new()
    {
        Region = "us-west-2",
    });

    var fooBucket = new Aws.S3.BucketV2("fooBucket", new()
    {
    }, new CustomResourceOptions
    {
        Provider = aws.Primary_region,
    });

    var barBucket = new Aws.S3.BucketV2("barBucket", new()
    {
    }, new CustomResourceOptions
    {
        Provider = aws.Secondary_region,
    });

    var example = new Aws.S3Control.MultiRegionAccessPoint("example", new()
    {
        Details = new Aws.S3Control.Inputs.MultiRegionAccessPointDetailsArgs
        {
            Name = "example",
            Regions = new[]
            {
                new Aws.S3Control.Inputs.MultiRegionAccessPointDetailsRegionArgs
                {
                    Bucket = fooBucket.Id,
                },
                new Aws.S3Control.Inputs.MultiRegionAccessPointDetailsRegionArgs
                {
                    Bucket = barBucket.Id,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "primaryRegion", &aws.ProviderArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		_, err = aws.NewProvider(ctx, "secondaryRegion", &aws.ProviderArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		fooBucket, err := s3.NewBucketV2(ctx, "fooBucket", nil, pulumi.Provider(aws.Primary_region))
		if err != nil {
			return err
		}
		barBucket, err := s3.NewBucketV2(ctx, "barBucket", nil, pulumi.Provider(aws.Secondary_region))
		if err != nil {
			return err
		}
		_, err = s3control.NewMultiRegionAccessPoint(ctx, "example", &s3control.MultiRegionAccessPointArgs{
			Details: &s3control.MultiRegionAccessPointDetailsArgs{
				Name: pulumi.String("example"),
				Regions: s3control.MultiRegionAccessPointDetailsRegionArray{
					&s3control.MultiRegionAccessPointDetailsRegionArgs{
						Bucket: fooBucket.ID(),
					},
					&s3control.MultiRegionAccessPointDetailsRegionArgs{
						Bucket: barBucket.ID(),
					},
				},
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
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.s3control.MultiRegionAccessPoint;
import com.pulumi.aws.s3control.MultiRegionAccessPointArgs;
import com.pulumi.aws.s3control.inputs.MultiRegionAccessPointDetailsArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var primaryRegion = new Provider("primaryRegion", ProviderArgs.builder()        
            .region("us-east-1")
            .build());

        var secondaryRegion = new Provider("secondaryRegion", ProviderArgs.builder()        
            .region("us-west-2")
            .build());

        var fooBucket = new BucketV2("fooBucket", BucketV2Args.Empty, CustomResourceOptions.builder()
            .provider(aws.primary_region())
            .build());

        var barBucket = new BucketV2("barBucket", BucketV2Args.Empty, CustomResourceOptions.builder()
            .provider(aws.secondary_region())
            .build());

        var example = new MultiRegionAccessPoint("example", MultiRegionAccessPointArgs.builder()        
            .details(MultiRegionAccessPointDetailsArgs.builder()
                .name("example")
                .regions(                
                    MultiRegionAccessPointDetailsRegionArgs.builder()
                        .bucket(fooBucket.id())
                        .build(),
                    MultiRegionAccessPointDetailsRegionArgs.builder()
                        .bucket(barBucket.id())
                        .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  primaryRegion:
    type: pulumi:providers:aws
    properties:
      region: us-east-1
  secondaryRegion:
    type: pulumi:providers:aws
    properties:
      region: us-west-2
  fooBucket:
    type: aws:s3:BucketV2
    options:
      provider: ${aws.primary_region}
  barBucket:
    type: aws:s3:BucketV2
    options:
      provider: ${aws.secondary_region}
  example:
    type: aws:s3control:MultiRegionAccessPoint
    properties:
      details:
        name: example
        regions:
          - bucket: ${fooBucket.id}
          - bucket: ${barBucket.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Multi-Region Access Points using the `account_id` and `name` of the Multi-Region Access Point separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:s3control/multiRegionAccessPoint:MultiRegionAccessPoint example 123456789012:example
```
 