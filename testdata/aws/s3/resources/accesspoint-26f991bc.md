Provides a resource to manage an S3 Access Point.

> **NOTE on Access Points and Access Point Policies:** This provider provides both a standalone Access Point Policy resource and an Access Point resource with a resource policy defined in-line. You cannot use an Access Point with in-line resource policy in conjunction with an Access Point Policy resource. Doing so will cause a conflict of policies and will overwrite the access point's resource policy.

> Advanced usage: To use a custom API endpoint for this resource, use the `s3control` endpoint provider configuration), not the `s3` endpoint provider configuration.

> This resource cannot be used with S3 directory buckets.

{{% examples %}}
## Example Usage
{{% example %}}
### AWS Partition General Purpose Bucket

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleAccessPoint = new aws.s3.AccessPoint("exampleAccessPoint", {bucket: exampleBucketV2.id});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_access_point = aws.s3.AccessPoint("exampleAccessPoint", bucket=example_bucket_v2.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleAccessPoint = new Aws.S3.AccessPoint("exampleAccessPoint", new()
    {
        Bucket = exampleBucketV2.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewAccessPoint(ctx, "exampleAccessPoint", &s3.AccessPointArgs{
			Bucket: exampleBucketV2.ID(),
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
import com.pulumi.aws.s3.AccessPoint;
import com.pulumi.aws.s3.AccessPointArgs;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleAccessPoint = new AccessPoint("exampleAccessPoint", AccessPointArgs.builder()        
            .bucket(exampleBucketV2.id())
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleAccessPoint:
    type: aws:s3:AccessPoint
    properties:
      bucket: ${exampleBucketV2.id}
```
{{% /example %}}
{{% example %}}
### S3 on Outposts Bucket

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucket = new aws.s3control.Bucket("exampleBucket", {bucket: "example"});
const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.0.0.0/16"});
const exampleAccessPoint = new aws.s3.AccessPoint("exampleAccessPoint", {
    bucket: exampleBucket.arn,
    vpcConfiguration: {
        vpcId: exampleVpc.id,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket = aws.s3control.Bucket("exampleBucket", bucket="example")
example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
example_access_point = aws.s3.AccessPoint("exampleAccessPoint",
    bucket=example_bucket.arn,
    vpc_configuration=aws.s3.AccessPointVpcConfigurationArgs(
        vpc_id=example_vpc.id,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucket = new Aws.S3Control.Bucket("exampleBucket", new()
    {
        BucketName = "example",
    });

    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var exampleAccessPoint = new Aws.S3.AccessPoint("exampleAccessPoint", new()
    {
        Bucket = exampleBucket.Arn,
        VpcConfiguration = new Aws.S3.Inputs.AccessPointVpcConfigurationArgs
        {
            VpcId = exampleVpc.Id,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucket, err := s3control.NewBucket(ctx, "exampleBucket", &s3control.BucketArgs{
			Bucket: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewAccessPoint(ctx, "exampleAccessPoint", &s3.AccessPointArgs{
			Bucket: exampleBucket.Arn,
			VpcConfiguration: &s3.AccessPointVpcConfigurationArgs{
				VpcId: exampleVpc.ID(),
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
import com.pulumi.aws.s3control.Bucket;
import com.pulumi.aws.s3control.BucketArgs;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.s3.AccessPoint;
import com.pulumi.aws.s3.AccessPointArgs;
import com.pulumi.aws.s3.inputs.AccessPointVpcConfigurationArgs;
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
        var exampleBucket = new Bucket("exampleBucket", BucketArgs.builder()        
            .bucket("example")
            .build());

        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var exampleAccessPoint = new AccessPoint("exampleAccessPoint", AccessPointArgs.builder()        
            .bucket(exampleBucket.arn())
            .vpcConfiguration(AccessPointVpcConfigurationArgs.builder()
                .vpcId(exampleVpc.id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleBucket:
    type: aws:s3control:Bucket
    properties:
      bucket: example
  exampleAccessPoint:
    type: aws:s3:AccessPoint
    properties:
      bucket: ${exampleBucket.arn}
      vpcConfiguration:
        vpcId: ${exampleVpc.id}
  exampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
```
{{% /example %}}
{{% /examples %}}

## Import

Import using the ARN for Access Points associated with an S3 on Outposts Bucket:

__Using `pulumi import` to import.__ For example:

Import using the `account_id` and `name` separated by a colon (`:`) for Access Points associated with an AWS Partition S3 Bucket:

```sh
 $ pulumi import aws:s3/accessPoint:AccessPoint example 123456789012:example
```
 Import using the ARN for Access Points associated with an S3 on Outposts Bucket:

```sh
 $ pulumi import aws:s3/accessPoint:AccessPoint example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-1234567890123456/accesspoint/example
```
 