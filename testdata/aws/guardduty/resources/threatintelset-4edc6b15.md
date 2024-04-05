Provides a resource to manage a GuardDuty ThreatIntelSet.

> **Note:** Currently in GuardDuty, users from member accounts cannot upload and further manage ThreatIntelSets. ThreatIntelSets that are uploaded by the primary account are imposed on GuardDuty functionality in its member accounts. See the [GuardDuty API Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/create-threat-intel-set.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.guardduty.Detector("primary", {enable: true});
const bucket = new aws.s3.BucketV2("bucket", {});
// ... other configuration ...
const bucketAcl = new aws.s3.BucketAclV2("bucketAcl", {
    bucket: bucket.id,
    acl: "private",
});
const myThreatIntelSetBucketObjectv2 = new aws.s3.BucketObjectv2("myThreatIntelSetBucketObjectv2", {
    acl: "public-read",
    content: "10.0.0.0/8\n",
    bucket: bucket.id,
    key: "MyThreatIntelSet",
});
const myThreatIntelSetThreatIntelSet = new aws.guardduty.ThreatIntelSet("myThreatIntelSetThreatIntelSet", {
    activate: true,
    detectorId: primary.id,
    format: "TXT",
    location: pulumi.interpolate`https://s3.amazonaws.com/${myThreatIntelSetBucketObjectv2.bucket}/${myThreatIntelSetBucketObjectv2.key}`,
});
```
```python
import pulumi
import pulumi_aws as aws

primary = aws.guardduty.Detector("primary", enable=True)
bucket = aws.s3.BucketV2("bucket")
# ... other configuration ...
bucket_acl = aws.s3.BucketAclV2("bucketAcl",
    bucket=bucket.id,
    acl="private")
my_threat_intel_set_bucket_objectv2 = aws.s3.BucketObjectv2("myThreatIntelSetBucketObjectv2",
    acl="public-read",
    content="10.0.0.0/8\n",
    bucket=bucket.id,
    key="MyThreatIntelSet")
my_threat_intel_set_threat_intel_set = aws.guardduty.ThreatIntelSet("myThreatIntelSetThreatIntelSet",
    activate=True,
    detector_id=primary.id,
    format="TXT",
    location=pulumi.Output.all(my_threat_intel_set_bucket_objectv2.bucket, my_threat_intel_set_bucket_objectv2.key).apply(lambda bucket, key: f"https://s3.amazonaws.com/{bucket}/{key}"))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var primary = new Aws.GuardDuty.Detector("primary", new()
    {
        Enable = true,
    });

    var bucket = new Aws.S3.BucketV2("bucket");

    // ... other configuration ...
    var bucketAcl = new Aws.S3.BucketAclV2("bucketAcl", new()
    {
        Bucket = bucket.Id,
        Acl = "private",
    });

    var myThreatIntelSetBucketObjectv2 = new Aws.S3.BucketObjectv2("myThreatIntelSetBucketObjectv2", new()
    {
        Acl = "public-read",
        Content = @"10.0.0.0/8
",
        Bucket = bucket.Id,
        Key = "MyThreatIntelSet",
    });

    var myThreatIntelSetThreatIntelSet = new Aws.GuardDuty.ThreatIntelSet("myThreatIntelSetThreatIntelSet", new()
    {
        Activate = true,
        DetectorId = primary.Id,
        Format = "TXT",
        Location = Output.Tuple(myThreatIntelSetBucketObjectv2.Bucket, myThreatIntelSetBucketObjectv2.Key).Apply(values =>
        {
            var bucket = values.Item1;
            var key = values.Item2;
            return $"https://s3.amazonaws.com/{bucket}/{key}";
        }),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/guardduty"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		primary, err := guardduty.NewDetector(ctx, "primary", &guardduty.DetectorArgs{
			Enable: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		bucket, err := s3.NewBucketV2(ctx, "bucket", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "bucketAcl", &s3.BucketAclV2Args{
			Bucket: bucket.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		myThreatIntelSetBucketObjectv2, err := s3.NewBucketObjectv2(ctx, "myThreatIntelSetBucketObjectv2", &s3.BucketObjectv2Args{
			Acl:     pulumi.String("public-read"),
			Content: pulumi.String("10.0.0.0/8\n"),
			Bucket:  bucket.ID(),
			Key:     pulumi.String("MyThreatIntelSet"),
		})
		if err != nil {
			return err
		}
		_, err = guardduty.NewThreatIntelSet(ctx, "myThreatIntelSetThreatIntelSet", &guardduty.ThreatIntelSetArgs{
			Activate:   pulumi.Bool(true),
			DetectorId: primary.ID(),
			Format:     pulumi.String("TXT"),
			Location: pulumi.All(myThreatIntelSetBucketObjectv2.Bucket, myThreatIntelSetBucketObjectv2.Key).ApplyT(func(_args []interface{}) (string, error) {
				bucket := _args[0].(string)
				key := _args[1].(string)
				return fmt.Sprintf("https://s3.amazonaws.com/%v/%v", bucket, key), nil
			}).(pulumi.StringOutput),
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
import com.pulumi.aws.guardduty.Detector;
import com.pulumi.aws.guardduty.DetectorArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.s3.BucketObjectv2;
import com.pulumi.aws.s3.BucketObjectv2Args;
import com.pulumi.aws.guardduty.ThreatIntelSet;
import com.pulumi.aws.guardduty.ThreatIntelSetArgs;
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
        var primary = new Detector("primary", DetectorArgs.builder()        
            .enable(true)
            .build());

        var bucket = new BucketV2("bucket");

        var bucketAcl = new BucketAclV2("bucketAcl", BucketAclV2Args.builder()        
            .bucket(bucket.id())
            .acl("private")
            .build());

        var myThreatIntelSetBucketObjectv2 = new BucketObjectv2("myThreatIntelSetBucketObjectv2", BucketObjectv2Args.builder()        
            .acl("public-read")
            .content("""
10.0.0.0/8
            """)
            .bucket(bucket.id())
            .key("MyThreatIntelSet")
            .build());

        var myThreatIntelSetThreatIntelSet = new ThreatIntelSet("myThreatIntelSetThreatIntelSet", ThreatIntelSetArgs.builder()        
            .activate(true)
            .detectorId(primary.id())
            .format("TXT")
            .location(Output.tuple(myThreatIntelSetBucketObjectv2.bucket(), myThreatIntelSetBucketObjectv2.key()).applyValue(values -> {
                var bucket = values.t1;
                var key = values.t2;
                return String.format("https://s3.amazonaws.com/%s/%s", bucket,key);
            }))
            .build());

    }
}
```
```yaml
resources:
  primary:
    type: aws:guardduty:Detector
    properties:
      enable: true
  bucket:
    type: aws:s3:BucketV2
  bucketAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${bucket.id}
      acl: private
  myThreatIntelSetBucketObjectv2:
    type: aws:s3:BucketObjectv2
    properties:
      acl: public-read
      content: |
        10.0.0.0/8
      bucket: ${bucket.id}
      key: MyThreatIntelSet
  myThreatIntelSetThreatIntelSet:
    type: aws:guardduty:ThreatIntelSet
    properties:
      activate: true
      detectorId: ${primary.id}
      format: TXT
      location: https://s3.amazonaws.com/${myThreatIntelSetBucketObjectv2.bucket}/${myThreatIntelSetBucketObjectv2.key}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import GuardDuty ThreatIntelSet using the primary GuardDuty detector ID and ThreatIntelSetID. For example:

```sh
 $ pulumi import aws:guardduty/threatIntelSet:ThreatIntelSet MyThreatIntelSet 00b00fd5aecc0ab60a708659477e9617:123456789012
```
 