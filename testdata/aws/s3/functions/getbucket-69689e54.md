Provides details about a specific S3 bucket.

This resource may prove useful when setting up a Route53 record, or an origin for a CloudFront
Distribution.

{{% examples %}}
## Example Usage
{{% example %}}
### Route53 Record

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const selected = aws.s3.getBucket({
    bucket: "bucket.test.com",
});
const testZone = aws.route53.getZone({
    name: "test.com.",
});
const example = new aws.route53.Record("example", {
    zoneId: testZone.then(testZone => testZone.id),
    name: "bucket",
    type: "A",
    aliases: [{
        name: selected.then(selected => selected.websiteDomain),
        zoneId: selected.then(selected => selected.hostedZoneId),
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

selected = aws.s3.get_bucket(bucket="bucket.test.com")
test_zone = aws.route53.get_zone(name="test.com.")
example = aws.route53.Record("example",
    zone_id=test_zone.id,
    name="bucket",
    type="A",
    aliases=[aws.route53.RecordAliasArgs(
        name=selected.website_domain,
        zone_id=selected.hosted_zone_id,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var selected = Aws.S3.GetBucket.Invoke(new()
    {
        Bucket = "bucket.test.com",
    });

    var testZone = Aws.Route53.GetZone.Invoke(new()
    {
        Name = "test.com.",
    });

    var example = new Aws.Route53.Record("example", new()
    {
        ZoneId = testZone.Apply(getZoneResult => getZoneResult.Id),
        Name = "bucket",
        Type = "A",
        Aliases = new[]
        {
            new Aws.Route53.Inputs.RecordAliasArgs
            {
                Name = selected.Apply(getBucketResult => getBucketResult.WebsiteDomain),
                ZoneId = selected.Apply(getBucketResult => getBucketResult.HostedZoneId),
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		selected, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
			Bucket: "bucket.test.com",
		}, nil)
		if err != nil {
			return err
		}
		testZone, err := route53.LookupZone(ctx, &route53.LookupZoneArgs{
			Name: pulumi.StringRef("test.com."),
		}, nil)
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "example", &route53.RecordArgs{
			ZoneId: *pulumi.String(testZone.Id),
			Name:   pulumi.String("bucket"),
			Type:   pulumi.String("A"),
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					Name:   *pulumi.String(selected.WebsiteDomain),
					ZoneId: *pulumi.String(selected.HostedZoneId),
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
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.inputs.GetBucketArgs;
import com.pulumi.aws.route53.Route53Functions;
import com.pulumi.aws.route53.inputs.GetZoneArgs;
import com.pulumi.aws.route53.Record;
import com.pulumi.aws.route53.RecordArgs;
import com.pulumi.aws.route53.inputs.RecordAliasArgs;
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
        final var selected = S3Functions.getBucket(GetBucketArgs.builder()
            .bucket("bucket.test.com")
            .build());

        final var testZone = Route53Functions.getZone(GetZoneArgs.builder()
            .name("test.com.")
            .build());

        var example = new Record("example", RecordArgs.builder()        
            .zoneId(testZone.applyValue(getZoneResult -> getZoneResult.id()))
            .name("bucket")
            .type("A")
            .aliases(RecordAliasArgs.builder()
                .name(selected.applyValue(getBucketResult -> getBucketResult.websiteDomain()))
                .zoneId(selected.applyValue(getBucketResult -> getBucketResult.hostedZoneId()))
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53:Record
    properties:
      zoneId: ${testZone.id}
      name: bucket
      type: A
      aliases:
        - name: ${selected.websiteDomain}
          zoneId: ${selected.hostedZoneId}
variables:
  selected:
    fn::invoke:
      Function: aws:s3:getBucket
      Arguments:
        bucket: bucket.test.com
  testZone:
    fn::invoke:
      Function: aws:route53:getZone
      Arguments:
        name: test.com.
```
{{% /example %}}
{{% example %}}
### CloudFront Origin

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const selected = aws.s3.getBucket({
    bucket: "a-test-bucket",
});
const test = new aws.cloudfront.Distribution("test", {origins: [{
    domainName: selected.then(selected => selected.bucketDomainName),
    originId: "s3-selected-bucket",
}]});
```
```python
import pulumi
import pulumi_aws as aws

selected = aws.s3.get_bucket(bucket="a-test-bucket")
test = aws.cloudfront.Distribution("test", origins=[aws.cloudfront.DistributionOriginArgs(
    domain_name=selected.bucket_domain_name,
    origin_id="s3-selected-bucket",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var selected = Aws.S3.GetBucket.Invoke(new()
    {
        Bucket = "a-test-bucket",
    });

    var test = new Aws.CloudFront.Distribution("test", new()
    {
        Origins = new[]
        {
            new Aws.CloudFront.Inputs.DistributionOriginArgs
            {
                DomainName = selected.Apply(getBucketResult => getBucketResult.BucketDomainName),
                OriginId = "s3-selected-bucket",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		selected, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
			Bucket: "a-test-bucket",
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudfront.NewDistribution(ctx, "test", &cloudfront.DistributionArgs{
			Origins: cloudfront.DistributionOriginArray{
				&cloudfront.DistributionOriginArgs{
					DomainName: *pulumi.String(selected.BucketDomainName),
					OriginId:   pulumi.String("s3-selected-bucket"),
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
import com.pulumi.aws.s3.S3Functions;
import com.pulumi.aws.s3.inputs.GetBucketArgs;
import com.pulumi.aws.cloudfront.Distribution;
import com.pulumi.aws.cloudfront.DistributionArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOriginArgs;
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
        final var selected = S3Functions.getBucket(GetBucketArgs.builder()
            .bucket("a-test-bucket")
            .build());

        var test = new Distribution("test", DistributionArgs.builder()        
            .origins(DistributionOriginArgs.builder()
                .domainName(selected.applyValue(getBucketResult -> getBucketResult.bucketDomainName()))
                .originId("s3-selected-bucket")
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:cloudfront:Distribution
    properties:
      origins:
        - domainName: ${selected.bucketDomainName}
          originId: s3-selected-bucket
variables:
  selected:
    fn::invoke:
      Function: aws:s3:getBucket
      Arguments:
        bucket: a-test-bucket
```
{{% /example %}}
{{% /examples %}}