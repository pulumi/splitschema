Provides a resource to manage an S3 Access Point resource policy.

> **NOTE on Access Points and Access Point Policies:** The provider provides both a standalone Access Point Policy resource and an Access Point resource with a resource policy defined in-line. You cannot use an Access Point with in-line resource policy in conjunction with an Access Point Policy resource. Doing so will cause a conflict of policies and will overwrite the access point's resource policy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleAccessPoint = new aws.s3.AccessPoint("exampleAccessPoint", {
    bucket: exampleBucketV2.id,
    publicAccessBlockConfiguration: {
        blockPublicAcls: true,
        blockPublicPolicy: false,
        ignorePublicAcls: true,
        restrictPublicBuckets: false,
    },
});
const exampleAccessPointPolicy = new aws.s3control.AccessPointPolicy("exampleAccessPointPolicy", {
    accessPointArn: exampleAccessPoint.arn,
    policy: exampleAccessPoint.arn.apply(arn => JSON.stringify({
        Version: "2008-10-17",
        Statement: [{
            Effect: "Allow",
            Action: "s3:GetObjectTagging",
            Principal: {
                AWS: "*",
            },
            Resource: `${arn}/object/*`,
        }],
    })),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_access_point = aws.s3.AccessPoint("exampleAccessPoint",
    bucket=example_bucket_v2.id,
    public_access_block_configuration=aws.s3.AccessPointPublicAccessBlockConfigurationArgs(
        block_public_acls=True,
        block_public_policy=False,
        ignore_public_acls=True,
        restrict_public_buckets=False,
    ))
example_access_point_policy = aws.s3control.AccessPointPolicy("exampleAccessPointPolicy",
    access_point_arn=example_access_point.arn,
    policy=example_access_point.arn.apply(lambda arn: json.dumps({
        "Version": "2008-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Action": "s3:GetObjectTagging",
            "Principal": {
                "AWS": "*",
            },
            "Resource": f"{arn}/object/*",
        }],
    })))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleAccessPoint = new Aws.S3.AccessPoint("exampleAccessPoint", new()
    {
        Bucket = exampleBucketV2.Id,
        PublicAccessBlockConfiguration = new Aws.S3.Inputs.AccessPointPublicAccessBlockConfigurationArgs
        {
            BlockPublicAcls = true,
            BlockPublicPolicy = false,
            IgnorePublicAcls = true,
            RestrictPublicBuckets = false,
        },
    });

    var exampleAccessPointPolicy = new Aws.S3Control.AccessPointPolicy("exampleAccessPointPolicy", new()
    {
        AccessPointArn = exampleAccessPoint.Arn,
        Policy = exampleAccessPoint.Arn.Apply(arn => JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2008-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Effect"] = "Allow",
                    ["Action"] = "s3:GetObjectTagging",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["AWS"] = "*",
                    },
                    ["Resource"] = $"{arn}/object/*",
                },
            },
        })),
    });

});
```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		exampleAccessPoint, err := s3.NewAccessPoint(ctx, "exampleAccessPoint", &s3.AccessPointArgs{
			Bucket: exampleBucketV2.ID(),
			PublicAccessBlockConfiguration: &s3.AccessPointPublicAccessBlockConfigurationArgs{
				BlockPublicAcls:       pulumi.Bool(true),
				BlockPublicPolicy:     pulumi.Bool(false),
				IgnorePublicAcls:      pulumi.Bool(true),
				RestrictPublicBuckets: pulumi.Bool(false),
			},
		})
		if err != nil {
			return err
		}
		_, err = s3control.NewAccessPointPolicy(ctx, "exampleAccessPointPolicy", &s3control.AccessPointPolicyArgs{
			AccessPointArn: exampleAccessPoint.Arn,
			Policy: exampleAccessPoint.Arn.ApplyT(func(arn string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON0, err := json.Marshal(map[string]interface{}{
					"Version": "2008-10-17",
					"Statement": []map[string]interface{}{
						map[string]interface{}{
							"Effect": "Allow",
							"Action": "s3:GetObjectTagging",
							"Principal": map[string]interface{}{
								"AWS": "*",
							},
							"Resource": fmt.Sprintf("%v/object/*", arn),
						},
					},
				})
				if err != nil {
					return _zero, err
				}
				json0 := string(tmpJSON0)
				return pulumi.String(json0), nil
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.AccessPoint;
import com.pulumi.aws.s3.AccessPointArgs;
import com.pulumi.aws.s3.inputs.AccessPointPublicAccessBlockConfigurationArgs;
import com.pulumi.aws.s3control.AccessPointPolicy;
import com.pulumi.aws.s3control.AccessPointPolicyArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
            .publicAccessBlockConfiguration(AccessPointPublicAccessBlockConfigurationArgs.builder()
                .blockPublicAcls(true)
                .blockPublicPolicy(false)
                .ignorePublicAcls(true)
                .restrictPublicBuckets(false)
                .build())
            .build());

        var exampleAccessPointPolicy = new AccessPointPolicy("exampleAccessPointPolicy", AccessPointPolicyArgs.builder()        
            .accessPointArn(exampleAccessPoint.arn())
            .policy(exampleAccessPoint.arn().applyValue(arn -> serializeJson(
                jsonObject(
                    jsonProperty("Version", "2008-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Action", "s3:GetObjectTagging"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", "*")
                        )),
                        jsonProperty("Resource", String.format("%s/object/*", arn))
                    )))
                ))))
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
      publicAccessBlockConfiguration:
        blockPublicAcls: true
        blockPublicPolicy: false
        ignorePublicAcls: true
        restrictPublicBuckets: false
  exampleAccessPointPolicy:
    type: aws:s3control:AccessPointPolicy
    properties:
      accessPointArn: ${exampleAccessPoint.arn}
      policy:
        fn::toJSON:
          Version: 2008-10-17
          Statement:
            - Effect: Allow
              Action: s3:GetObjectTagging
              Principal:
                AWS: '*'
              Resource: ${exampleAccessPoint.arn}/object/*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Access Point policies using the `access_point_arn`. For example:

```sh
 $ pulumi import aws:s3control/accessPointPolicy:AccessPointPolicy example arn:aws:s3:us-west-2:123456789012:accesspoint/example
```
 