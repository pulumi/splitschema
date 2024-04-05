Provides a resource to manage a GuardDuty IPSet.

> **Note:** Currently in GuardDuty, users from member accounts cannot upload and further manage IPSets. IPSets that are uploaded by the primary account are imposed on GuardDuty functionality in its member accounts. See the [GuardDuty API Documentation](https://docs.aws.amazon.com/guardduty/latest/ug/create-ip-set.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const primary = new aws.guardduty.Detector("primary", {enable: true});
const bucket = new aws.s3.BucketV2("bucket", {});
// ... other configuration
const myIPSet = new aws.s3.BucketObjectv2("myIPSet", {
    content: "10.0.0.0/8\n",
    bucket: bucket.id,
    key: "MyIPSet",
});
const example = new aws.guardduty.IPSet("example", {
    activate: true,
    detectorId: primary.id,
    format: "TXT",
    location: pulumi.interpolate`https://s3.amazonaws.com/${myIPSet.bucket}/${myIPSet.key}`,
});
const bucketAcl = new aws.s3.BucketAclV2("bucketAcl", {
    bucket: bucket.id,
    acl: "private",
});
```
```python
import pulumi
import pulumi_aws as aws

primary = aws.guardduty.Detector("primary", enable=True)
bucket = aws.s3.BucketV2("bucket")
# ... other configuration
my_ip_set = aws.s3.BucketObjectv2("myIPSet",
    content="10.0.0.0/8\n",
    bucket=bucket.id,
    key="MyIPSet")
example = aws.guardduty.IPSet("example",
    activate=True,
    detector_id=primary.id,
    format="TXT",
    location=pulumi.Output.all(my_ip_set.bucket, my_ip_set.key).apply(lambda bucket, key: f"https://s3.amazonaws.com/{bucket}/{key}"))
bucket_acl = aws.s3.BucketAclV2("bucketAcl",
    bucket=bucket.id,
    acl="private")
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

    // ... other configuration
    var myIPSet = new Aws.S3.BucketObjectv2("myIPSet", new()
    {
        Content = @"10.0.0.0/8
",
        Bucket = bucket.Id,
        Key = "MyIPSet",
    });

    var example = new Aws.GuardDuty.IPSet("example", new()
    {
        Activate = true,
        DetectorId = primary.Id,
        Format = "TXT",
        Location = Output.Tuple(myIPSet.Bucket, myIPSet.Key).Apply(values =>
        {
            var bucket = values.Item1;
            var key = values.Item2;
            return $"https://s3.amazonaws.com/{bucket}/{key}";
        }),
    });

    var bucketAcl = new Aws.S3.BucketAclV2("bucketAcl", new()
    {
        Bucket = bucket.Id,
        Acl = "private",
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
		myIPSet, err := s3.NewBucketObjectv2(ctx, "myIPSet", &s3.BucketObjectv2Args{
			Content: pulumi.String("10.0.0.0/8\n"),
			Bucket:  bucket.ID(),
			Key:     pulumi.String("MyIPSet"),
		})
		if err != nil {
			return err
		}
		_, err = guardduty.NewIPSet(ctx, "example", &guardduty.IPSetArgs{
			Activate:   pulumi.Bool(true),
			DetectorId: primary.ID(),
			Format:     pulumi.String("TXT"),
			Location: pulumi.All(myIPSet.Bucket, myIPSet.Key).ApplyT(func(_args []interface{}) (string, error) {
				bucket := _args[0].(string)
				key := _args[1].(string)
				return fmt.Sprintf("https://s3.amazonaws.com/%v/%v", bucket, key), nil
			}).(pulumi.StringOutput),
		})
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
import com.pulumi.aws.s3.BucketObjectv2;
import com.pulumi.aws.s3.BucketObjectv2Args;
import com.pulumi.aws.guardduty.IPSet;
import com.pulumi.aws.guardduty.IPSetArgs;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
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

        var myIPSet = new BucketObjectv2("myIPSet", BucketObjectv2Args.builder()        
            .content("""
10.0.0.0/8
            """)
            .bucket(bucket.id())
            .key("MyIPSet")
            .build());

        var example = new IPSet("example", IPSetArgs.builder()        
            .activate(true)
            .detectorId(primary.id())
            .format("TXT")
            .location(Output.tuple(myIPSet.bucket(), myIPSet.key()).applyValue(values -> {
                var bucket = values.t1;
                var key = values.t2;
                return String.format("https://s3.amazonaws.com/%s/%s", bucket,key);
            }))
            .build());

        var bucketAcl = new BucketAclV2("bucketAcl", BucketAclV2Args.builder()        
            .bucket(bucket.id())
            .acl("private")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:guardduty:IPSet
    properties:
      activate: true
      detectorId: ${primary.id}
      format: TXT
      location: https://s3.amazonaws.com/${myIPSet.bucket}/${myIPSet.key}
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
  myIPSet:
    type: aws:s3:BucketObjectv2
    properties:
      content: |
        10.0.0.0/8
      bucket: ${bucket.id}
      key: MyIPSet
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import GuardDuty IPSet using the primary GuardDuty detector ID and IPSet ID. For example:

```sh
 $ pulumi import aws:guardduty/iPSet:IPSet MyIPSet 00b00fd5aecc0ab60a708659477e9617:123456789012
```
 