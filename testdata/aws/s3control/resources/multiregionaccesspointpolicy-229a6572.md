Provides a resource to manage an S3 Multi-Region Access Point access control policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentCallerIdentity = aws.getCallerIdentity({});
const currentPartition = aws.getPartition({});
const fooBucket = new aws.s3.BucketV2("fooBucket", {});
const exampleMultiRegionAccessPoint = new aws.s3control.MultiRegionAccessPoint("exampleMultiRegionAccessPoint", {details: {
    name: "example",
    regions: [{
        bucket: fooBucket.id,
    }],
}});
const exampleMultiRegionAccessPointPolicy = new aws.s3control.MultiRegionAccessPointPolicy("exampleMultiRegionAccessPointPolicy", {details: {
    name: exampleMultiRegionAccessPoint.id.apply(id => id.split(":"))[1],
    policy: pulumi.all([currentCallerIdentity, currentPartition, currentCallerIdentity, exampleMultiRegionAccessPoint.alias]).apply(([currentCallerIdentity, currentPartition, currentCallerIdentity1, alias]) => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Sid: "Example",
            Effect: "Allow",
            Principal: {
                AWS: currentCallerIdentity.accountId,
            },
            Action: [
                "s3:GetObject",
                "s3:PutObject",
            ],
            Resource: `arn:${currentPartition.partition}:s3::${currentCallerIdentity1.accountId}:accesspoint/${alias}/object/*`,
        }],
    })),
}});
```
```python
import pulumi
import json
import pulumi_aws as aws

current_caller_identity = aws.get_caller_identity()
current_partition = aws.get_partition()
foo_bucket = aws.s3.BucketV2("fooBucket")
example_multi_region_access_point = aws.s3control.MultiRegionAccessPoint("exampleMultiRegionAccessPoint", details=aws.s3control.MultiRegionAccessPointDetailsArgs(
    name="example",
    regions=[aws.s3control.MultiRegionAccessPointDetailsRegionArgs(
        bucket=foo_bucket.id,
    )],
))
example_multi_region_access_point_policy = aws.s3control.MultiRegionAccessPointPolicy("exampleMultiRegionAccessPointPolicy", details=aws.s3control.MultiRegionAccessPointPolicyDetailsArgs(
    name=example_multi_region_access_point.id.apply(lambda id: id.split(":"))[1],
    policy=example_multi_region_access_point.alias.apply(lambda alias: json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Sid": "Example",
            "Effect": "Allow",
            "Principal": {
                "AWS": current_caller_identity.account_id,
            },
            "Action": [
                "s3:GetObject",
                "s3:PutObject",
            ],
            "Resource": f"arn:{current_partition.partition}:s3::{current_caller_identity.account_id}:accesspoint/{alias}/object/*",
        }],
    })),
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var currentPartition = Aws.GetPartition.Invoke();

    var fooBucket = new Aws.S3.BucketV2("fooBucket");

    var exampleMultiRegionAccessPoint = new Aws.S3Control.MultiRegionAccessPoint("exampleMultiRegionAccessPoint", new()
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
            },
        },
    });

    var exampleMultiRegionAccessPointPolicy = new Aws.S3Control.MultiRegionAccessPointPolicy("exampleMultiRegionAccessPointPolicy", new()
    {
        Details = new Aws.S3Control.Inputs.MultiRegionAccessPointPolicyDetailsArgs
        {
            Name = exampleMultiRegionAccessPoint.Id.Apply(id => id.Split(":"))[1],
            Policy = Output.Tuple(currentCallerIdentity, currentPartition, currentCallerIdentity, exampleMultiRegionAccessPoint.Alias).Apply(values =>
            {
                var currentCallerIdentity = values.Item1;
                var currentPartition = values.Item2;
                var currentCallerIdentity1 = values.Item3;
                var @alias = values.Item4;
                return JsonSerializer.Serialize(new Dictionary<string, object?>
                {
                    ["Version"] = "2012-10-17",
                    ["Statement"] = new[]
                    {
                        new Dictionary<string, object?>
                        {
                            ["Sid"] = "Example",
                            ["Effect"] = "Allow",
                            ["Principal"] = new Dictionary<string, object?>
                            {
                                ["AWS"] = currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
                            },
                            ["Action"] = new[]
                            {
                                "s3:GetObject",
                                "s3:PutObject",
                            },
                            ["Resource"] = $"arn:{currentPartition.Apply(getPartitionResult => getPartitionResult.Partition)}:s3::{currentCallerIdentity1.AccountId}:accesspoint/{@alias}/object/*",
                        },
                    },
                });
            }),
        },
    });

});
```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentPartition, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		fooBucket, err := s3.NewBucketV2(ctx, "fooBucket", nil)
		if err != nil {
			return err
		}
		exampleMultiRegionAccessPoint, err := s3control.NewMultiRegionAccessPoint(ctx, "exampleMultiRegionAccessPoint", &s3control.MultiRegionAccessPointArgs{
			Details: &s3control.MultiRegionAccessPointDetailsArgs{
				Name: pulumi.String("example"),
				Regions: s3control.MultiRegionAccessPointDetailsRegionArray{
					&s3control.MultiRegionAccessPointDetailsRegionArgs{
						Bucket: fooBucket.ID(),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = s3control.NewMultiRegionAccessPointPolicy(ctx, "exampleMultiRegionAccessPointPolicy", &s3control.MultiRegionAccessPointPolicyArgs{
			Details: &s3control.MultiRegionAccessPointPolicyDetailsArgs{
				Name: "TODO: call element",
				Policy: exampleMultiRegionAccessPoint.Alias.ApplyT(func(alias string) (pulumi.String, error) {
					var _zero pulumi.String
					tmpJSON0, err := json.Marshal(map[string]interface{}{
						"Version": "2012-10-17",
						"Statement": []map[string]interface{}{
							map[string]interface{}{
								"Sid":    "Example",
								"Effect": "Allow",
								"Principal": map[string]interface{}{
									"AWS": currentCallerIdentity.AccountId,
								},
								"Action": []string{
									"s3:GetObject",
									"s3:PutObject",
								},
								"Resource": fmt.Sprintf("arn:%v:s3::%v:accesspoint/%v/object/*", currentPartition.Partition, currentCallerIdentity.AccountId, alias),
							},
						},
					})
					if err != nil {
						return _zero, err
					}
					json0 := string(tmpJSON0)
					return pulumi.String(json0), nil
				}).(pulumi.StringOutput),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3control.MultiRegionAccessPoint;
import com.pulumi.aws.s3control.MultiRegionAccessPointArgs;
import com.pulumi.aws.s3control.inputs.MultiRegionAccessPointDetailsArgs;
import com.pulumi.aws.s3control.MultiRegionAccessPointPolicy;
import com.pulumi.aws.s3control.MultiRegionAccessPointPolicyArgs;
import com.pulumi.aws.s3control.inputs.MultiRegionAccessPointPolicyDetailsArgs;
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
        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        final var currentPartition = AwsFunctions.getPartition();

        var fooBucket = new BucketV2("fooBucket");

        var exampleMultiRegionAccessPoint = new MultiRegionAccessPoint("exampleMultiRegionAccessPoint", MultiRegionAccessPointArgs.builder()        
            .details(MultiRegionAccessPointDetailsArgs.builder()
                .name("example")
                .regions(MultiRegionAccessPointDetailsRegionArgs.builder()
                    .bucket(fooBucket.id())
                    .build())
                .build())
            .build());

        var exampleMultiRegionAccessPointPolicy = new MultiRegionAccessPointPolicy("exampleMultiRegionAccessPointPolicy", MultiRegionAccessPointPolicyArgs.builder()        
            .details(MultiRegionAccessPointPolicyDetailsArgs.builder()
                .name(exampleMultiRegionAccessPoint.id().applyValue(id -> id.split(":"))[1])
                .policy(exampleMultiRegionAccessPoint.alias().applyValue(alias -> serializeJson(
                    jsonObject(
                        jsonProperty("Version", "2012-10-17"),
                        jsonProperty("Statement", jsonArray(jsonObject(
                            jsonProperty("Sid", "Example"),
                            jsonProperty("Effect", "Allow"),
                            jsonProperty("Principal", jsonObject(
                                jsonProperty("AWS", currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
                            )),
                            jsonProperty("Action", jsonArray(
                                "s3:GetObject", 
                                "s3:PutObject"
                            )),
                            jsonProperty("Resource", String.format("arn:%s:s3::%s:accesspoint/%s/object/*", currentPartition.applyValue(getPartitionResult -> getPartitionResult.partition()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()),alias))
                        )))
                    ))))
                .build())
            .build());

    }
}
```
```yaml
resources:
  fooBucket:
    type: aws:s3:BucketV2
  exampleMultiRegionAccessPoint:
    type: aws:s3control:MultiRegionAccessPoint
    properties:
      details:
        name: example
        regions:
          - bucket: ${fooBucket.id}
  exampleMultiRegionAccessPointPolicy:
    type: aws:s3control:MultiRegionAccessPointPolicy
    properties:
      details:
        name:
          fn::select:
            - 1
            - fn::split:
                - ${exampleMultiRegionAccessPoint.id}
                - ':'
        policy:
          fn::toJSON:
            Version: 2012-10-17
            Statement:
              - Sid: Example
                Effect: Allow
                Principal:
                  AWS: ${currentCallerIdentity.accountId}
                Action:
                  - s3:GetObject
                  - s3:PutObject
                Resource: arn:${currentPartition.partition}:s3::${currentCallerIdentity.accountId}:accesspoint/${exampleMultiRegionAccessPoint.alias}/object/*
variables:
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  currentPartition:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Multi-Region Access Point Policies using the `account_id` and `name` of the Multi-Region Access Point separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:s3control/multiRegionAccessPointPolicy:MultiRegionAccessPointPolicy example 123456789012:example
```
 