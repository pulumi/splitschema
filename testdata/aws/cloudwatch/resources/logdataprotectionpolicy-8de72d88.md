Provides a CloudWatch Log Data Protection Policy resource.

Read more about protecting sensitive user data in the [User Guide](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/mask-sensitive-log-data.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {});
const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleLogDataProtectionPolicy = new aws.cloudwatch.LogDataProtectionPolicy("exampleLogDataProtectionPolicy", {
    logGroupName: exampleLogGroup.name,
    policyDocument: exampleBucketV2.bucket.apply(bucket => JSON.stringify({
        Name: "Example",
        Version: "2021-06-01",
        Statement: [
            {
                Sid: "Audit",
                DataIdentifier: ["arn:aws:dataprotection::aws:data-identifier/EmailAddress"],
                Operation: {
                    Audit: {
                        FindingsDestination: {
                            S3: {
                                Bucket: bucket,
                            },
                        },
                    },
                },
            },
            {
                Sid: "Redact",
                DataIdentifier: ["arn:aws:dataprotection::aws:data-identifier/EmailAddress"],
                Operation: {
                    Deidentify: {
                        MaskConfig: {},
                    },
                },
            },
        ],
    })),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup")
example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_log_data_protection_policy = aws.cloudwatch.LogDataProtectionPolicy("exampleLogDataProtectionPolicy",
    log_group_name=example_log_group.name,
    policy_document=example_bucket_v2.bucket.apply(lambda bucket: json.dumps({
        "Name": "Example",
        "Version": "2021-06-01",
        "Statement": [
            {
                "Sid": "Audit",
                "DataIdentifier": ["arn:aws:dataprotection::aws:data-identifier/EmailAddress"],
                "Operation": {
                    "Audit": {
                        "FindingsDestination": {
                            "S3": {
                                "Bucket": bucket,
                            },
                        },
                    },
                },
            },
            {
                "Sid": "Redact",
                "DataIdentifier": ["arn:aws:dataprotection::aws:data-identifier/EmailAddress"],
                "Operation": {
                    "Deidentify": {
                        "MaskConfig": {},
                    },
                },
            },
        ],
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
    var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup");

    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleLogDataProtectionPolicy = new Aws.CloudWatch.LogDataProtectionPolicy("exampleLogDataProtectionPolicy", new()
    {
        LogGroupName = exampleLogGroup.Name,
        PolicyDocument = exampleBucketV2.Bucket.Apply(bucket => JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Name"] = "Example",
            ["Version"] = "2021-06-01",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Sid"] = "Audit",
                    ["DataIdentifier"] = new[]
                    {
                        "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                    },
                    ["Operation"] = new Dictionary<string, object?>
                    {
                        ["Audit"] = new Dictionary<string, object?>
                        {
                            ["FindingsDestination"] = new Dictionary<string, object?>
                            {
                                ["S3"] = new Dictionary<string, object?>
                                {
                                    ["Bucket"] = bucket,
                                },
                            },
                        },
                    },
                },
                new Dictionary<string, object?>
                {
                    ["Sid"] = "Redact",
                    ["DataIdentifier"] = new[]
                    {
                        "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                    },
                    ["Operation"] = new Dictionary<string, object?>
                    {
                        ["Deidentify"] = new Dictionary<string, object?>
                        {
                            ["MaskConfig"] = new Dictionary<string, object?>
                            {
                            },
                        },
                    },
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
		if err != nil {
			return err
		}
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogDataProtectionPolicy(ctx, "exampleLogDataProtectionPolicy", &cloudwatch.LogDataProtectionPolicyArgs{
			LogGroupName: exampleLogGroup.Name,
			PolicyDocument: exampleBucketV2.Bucket.ApplyT(func(bucket string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON0, err := json.Marshal(map[string]interface{}{
					"Name":    "Example",
					"Version": "2021-06-01",
					"Statement": []interface{}{
						map[string]interface{}{
							"Sid": "Audit",
							"DataIdentifier": []string{
								"arn:aws:dataprotection::aws:data-identifier/EmailAddress",
							},
							"Operation": map[string]interface{}{
								"Audit": map[string]interface{}{
									"FindingsDestination": map[string]interface{}{
										"S3": map[string]interface{}{
											"Bucket": bucket,
										},
									},
								},
							},
						},
						map[string]interface{}{
							"Sid": "Redact",
							"DataIdentifier": []string{
								"arn:aws:dataprotection::aws:data-identifier/EmailAddress",
							},
							"Operation": map[string]interface{}{
								"Deidentify": map[string]interface{}{
									"MaskConfig": nil,
								},
							},
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.cloudwatch.LogDataProtectionPolicy;
import com.pulumi.aws.cloudwatch.LogDataProtectionPolicyArgs;
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
        var exampleLogGroup = new LogGroup("exampleLogGroup");

        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleLogDataProtectionPolicy = new LogDataProtectionPolicy("exampleLogDataProtectionPolicy", LogDataProtectionPolicyArgs.builder()        
            .logGroupName(exampleLogGroup.name())
            .policyDocument(exampleBucketV2.bucket().applyValue(bucket -> serializeJson(
                jsonObject(
                    jsonProperty("Name", "Example"),
                    jsonProperty("Version", "2021-06-01"),
                    jsonProperty("Statement", jsonArray(
                        jsonObject(
                            jsonProperty("Sid", "Audit"),
                            jsonProperty("DataIdentifier", jsonArray("arn:aws:dataprotection::aws:data-identifier/EmailAddress")),
                            jsonProperty("Operation", jsonObject(
                                jsonProperty("Audit", jsonObject(
                                    jsonProperty("FindingsDestination", jsonObject(
                                        jsonProperty("S3", jsonObject(
                                            jsonProperty("Bucket", bucket)
                                        ))
                                    ))
                                ))
                            ))
                        ), 
                        jsonObject(
                            jsonProperty("Sid", "Redact"),
                            jsonProperty("DataIdentifier", jsonArray("arn:aws:dataprotection::aws:data-identifier/EmailAddress")),
                            jsonProperty("Operation", jsonObject(
                                jsonProperty("Deidentify", jsonObject(
                                    jsonProperty("MaskConfig", jsonObject(

                                    ))
                                ))
                            ))
                        )
                    ))
                ))))
            .build());

    }
}
```
```yaml
resources:
  exampleLogGroup:
    type: aws:cloudwatch:LogGroup
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleLogDataProtectionPolicy:
    type: aws:cloudwatch:LogDataProtectionPolicy
    properties:
      logGroupName: ${exampleLogGroup.name}
      policyDocument:
        fn::toJSON:
          Name: Example
          Version: 2021-06-01
          Statement:
            - Sid: Audit
              DataIdentifier:
                - arn:aws:dataprotection::aws:data-identifier/EmailAddress
              Operation:
                Audit:
                  FindingsDestination:
                    S3:
                      Bucket: ${exampleBucketV2.bucket}
            - Sid: Redact
              DataIdentifier:
                - arn:aws:dataprotection::aws:data-identifier/EmailAddress
              Operation:
                Deidentify:
                  MaskConfig: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import this resource using the `log_group_name`. For example:

```sh
 $ pulumi import aws:cloudwatch/logDataProtectionPolicy:LogDataProtectionPolicy example my-log-group
```
 