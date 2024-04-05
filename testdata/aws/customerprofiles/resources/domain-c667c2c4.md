Resource for managing an Amazon Customer Profiles Domain.
See the [Create Domain](https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_CreateDomain.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.customerprofiles.Domain("example", {domainName: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.customerprofiles.Domain("example", domain_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CustomerProfiles.Domain("example", new()
    {
        DomainName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/customerprofiles"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := customerprofiles.NewDomain(ctx, "example", &customerprofiles.DomainArgs{
			DomainName: pulumi.String("example"),
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
import com.pulumi.aws.customerprofiles.Domain;
import com.pulumi.aws.customerprofiles.DomainArgs;
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
        var example = new Domain("example", DomainArgs.builder()        
            .domainName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:customerprofiles:Domain
    properties:
      domainName: example
```
{{% /example %}}
{{% example %}}
### With SQS DLQ and KMS set

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleQueue = new aws.sqs.Queue("exampleQueue", {policy: JSON.stringify({
    Version: "2012-10-17",
    Statement: [{
        Sid: "Customer Profiles SQS policy",
        Effect: "Allow",
        Action: ["sqs:SendMessage"],
        Resource: "*",
        Principal: {
            Service: "profile.amazonaws.com",
        },
    }],
})});
const exampleKey = new aws.kms.Key("exampleKey", {
    description: "example",
    deletionWindowInDays: 10,
});
const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {forceDestroy: true});
const exampleBucketPolicy = new aws.s3.BucketPolicy("exampleBucketPolicy", {
    bucket: exampleBucketV2.id,
    policy: pulumi.all([exampleBucketV2.arn, exampleBucketV2.arn]).apply(([exampleBucketV2Arn, exampleBucketV2Arn1]) => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Sid: "Customer Profiles S3 policy",
            Effect: "Allow",
            Action: [
                "s3:GetObject",
                "s3:PutObject",
                "s3:ListBucket",
            ],
            Resource: [
                exampleBucketV2Arn,
                `${exampleBucketV2Arn1}/*`,
            ],
            Principal: {
                Service: "profile.amazonaws.com",
            },
        }],
    })),
});
const test = new aws.customerprofiles.Domain("test", {
    domainName: example,
    deadLetterQueueUrl: exampleQueue.id,
    defaultEncryptionKey: exampleKey.arn,
    defaultExpirationDays: 365,
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_queue = aws.sqs.Queue("exampleQueue", policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Sid": "Customer Profiles SQS policy",
        "Effect": "Allow",
        "Action": ["sqs:SendMessage"],
        "Resource": "*",
        "Principal": {
            "Service": "profile.amazonaws.com",
        },
    }],
}))
example_key = aws.kms.Key("exampleKey",
    description="example",
    deletion_window_in_days=10)
example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2", force_destroy=True)
example_bucket_policy = aws.s3.BucketPolicy("exampleBucketPolicy",
    bucket=example_bucket_v2.id,
    policy=pulumi.Output.all(example_bucket_v2.arn, example_bucket_v2.arn).apply(lambda exampleBucketV2Arn, exampleBucketV2Arn1: json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Sid": "Customer Profiles S3 policy",
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:PutObject",
                "s3:ListBucket",
            ],
            "Resource": [
                example_bucket_v2_arn,
                f"{example_bucket_v2_arn1}/*",
            ],
            "Principal": {
                "Service": "profile.amazonaws.com",
            },
        }],
    })))
test = aws.customerprofiles.Domain("test",
    domain_name=example,
    dead_letter_queue_url=example_queue.id,
    default_encryption_key=example_key.arn,
    default_expiration_days=365)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleQueue = new Aws.Sqs.Queue("exampleQueue", new()
    {
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Sid"] = "Customer Profiles SQS policy",
                    ["Effect"] = "Allow",
                    ["Action"] = new[]
                    {
                        "sqs:SendMessage",
                    },
                    ["Resource"] = "*",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "profile.amazonaws.com",
                    },
                },
            },
        }),
    });

    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "example",
        DeletionWindowInDays = 10,
    });

    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2", new()
    {
        ForceDestroy = true,
    });

    var exampleBucketPolicy = new Aws.S3.BucketPolicy("exampleBucketPolicy", new()
    {
        Bucket = exampleBucketV2.Id,
        Policy = Output.Tuple(exampleBucketV2.Arn, exampleBucketV2.Arn).Apply(values =>
        {
            var exampleBucketV2Arn = values.Item1;
            var exampleBucketV2Arn1 = values.Item2;
            return JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                ["Version"] = "2012-10-17",
                ["Statement"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["Sid"] = "Customer Profiles S3 policy",
                        ["Effect"] = "Allow",
                        ["Action"] = new[]
                        {
                            "s3:GetObject",
                            "s3:PutObject",
                            "s3:ListBucket",
                        },
                        ["Resource"] = new[]
                        {
                            exampleBucketV2Arn,
                            $"{exampleBucketV2Arn1}/*",
                        },
                        ["Principal"] = new Dictionary<string, object?>
                        {
                            ["Service"] = "profile.amazonaws.com",
                        },
                    },
                },
            });
        }),
    });

    var test = new Aws.CustomerProfiles.Domain("test", new()
    {
        DomainName = example,
        DeadLetterQueueUrl = exampleQueue.Id,
        DefaultEncryptionKey = exampleKey.Arn,
        DefaultExpirationDays = 365,
    });

});
```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/customerprofiles"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Sid":    "Customer Profiles SQS policy",
					"Effect": "Allow",
					"Action": []string{
						"sqs:SendMessage",
					},
					"Resource": "*",
					"Principal": map[string]interface{}{
						"Service": "profile.amazonaws.com",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleQueue, err := sqs.NewQueue(ctx, "exampleQueue", &sqs.QueueArgs{
			Policy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description:          pulumi.String("example"),
			DeletionWindowInDays: pulumi.Int(10),
		})
		if err != nil {
			return err
		}
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", &s3.BucketV2Args{
			ForceDestroy: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketPolicy(ctx, "exampleBucketPolicy", &s3.BucketPolicyArgs{
			Bucket: exampleBucketV2.ID(),
			Policy: pulumi.All(exampleBucketV2.Arn, exampleBucketV2.Arn).ApplyT(func(_args []interface{}) (string, error) {
				exampleBucketV2Arn := _args[0].(string)
				exampleBucketV2Arn1 := _args[1].(string)
				var _zero string
				tmpJSON1, err := json.Marshal(map[string]interface{}{
					"Version": "2012-10-17",
					"Statement": []map[string]interface{}{
						map[string]interface{}{
							"Sid":    "Customer Profiles S3 policy",
							"Effect": "Allow",
							"Action": []string{
								"s3:GetObject",
								"s3:PutObject",
								"s3:ListBucket",
							},
							"Resource": []string{
								exampleBucketV2Arn,
								fmt.Sprintf("%v/*", exampleBucketV2Arn1),
							},
							"Principal": map[string]interface{}{
								"Service": "profile.amazonaws.com",
							},
						},
					},
				})
				if err != nil {
					return _zero, err
				}
				json1 := string(tmpJSON1)
				return json1, nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = customerprofiles.NewDomain(ctx, "test", &customerprofiles.DomainArgs{
			DomainName:            pulumi.Any(example),
			DeadLetterQueueUrl:    exampleQueue.ID(),
			DefaultEncryptionKey:  exampleKey.Arn,
			DefaultExpirationDays: pulumi.Int(365),
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
import com.pulumi.aws.sqs.Queue;
import com.pulumi.aws.sqs.QueueArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
import com.pulumi.aws.customerprofiles.Domain;
import com.pulumi.aws.customerprofiles.DomainArgs;
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
        var exampleQueue = new Queue("exampleQueue", QueueArgs.builder()        
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Sid", "Customer Profiles SQS policy"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Action", jsonArray("sqs:SendMessage")),
                        jsonProperty("Resource", "*"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "profile.amazonaws.com")
                        ))
                    )))
                )))
            .build());

        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("example")
            .deletionWindowInDays(10)
            .build());

        var exampleBucketV2 = new BucketV2("exampleBucketV2", BucketV2Args.builder()        
            .forceDestroy(true)
            .build());

        var exampleBucketPolicy = new BucketPolicy("exampleBucketPolicy", BucketPolicyArgs.builder()        
            .bucket(exampleBucketV2.id())
            .policy(Output.tuple(exampleBucketV2.arn(), exampleBucketV2.arn()).applyValue(values -> {
                var exampleBucketV2Arn = values.t1;
                var exampleBucketV2Arn1 = values.t2;
                return serializeJson(
                    jsonObject(
                        jsonProperty("Version", "2012-10-17"),
                        jsonProperty("Statement", jsonArray(jsonObject(
                            jsonProperty("Sid", "Customer Profiles S3 policy"),
                            jsonProperty("Effect", "Allow"),
                            jsonProperty("Action", jsonArray(
                                "s3:GetObject", 
                                "s3:PutObject", 
                                "s3:ListBucket"
                            )),
                            jsonProperty("Resource", jsonArray(
                                exampleBucketV2Arn, 
                                String.format("%s/*", exampleBucketV2Arn1)
                            )),
                            jsonProperty("Principal", jsonObject(
                                jsonProperty("Service", "profile.amazonaws.com")
                            ))
                        )))
                    ));
            }))
            .build());

        var test = new Domain("test", DomainArgs.builder()        
            .domainName(example)
            .deadLetterQueueUrl(exampleQueue.id())
            .defaultEncryptionKey(exampleKey.arn())
            .defaultExpirationDays(365)
            .build());

    }
}
```
```yaml
resources:
  exampleQueue:
    type: aws:sqs:Queue
    properties:
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Sid: Customer Profiles SQS policy
              Effect: Allow
              Action:
                - sqs:SendMessage
              Resource: '*'
              Principal:
                Service: profile.amazonaws.com
  exampleKey:
    type: aws:kms:Key
    properties:
      description: example
      deletionWindowInDays: 10
  exampleBucketV2:
    type: aws:s3:BucketV2
    properties:
      forceDestroy: true
  exampleBucketPolicy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${exampleBucketV2.id}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Sid: Customer Profiles S3 policy
              Effect: Allow
              Action:
                - s3:GetObject
                - s3:PutObject
                - s3:ListBucket
              Resource:
                - ${exampleBucketV2.arn}
                - ${exampleBucketV2.arn}/*
              Principal:
                Service: profile.amazonaws.com
  test:
    type: aws:customerprofiles:Domain
    properties:
      domainName: ${example}
      deadLetterQueueUrl: ${exampleQueue.id}
      defaultEncryptionKey: ${exampleKey.arn}
      defaultExpirationDays: 365
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Customer Profiles Domain using the resource `id`. For example:

```sh
 $ pulumi import aws:customerprofiles/domain:Domain example e6f777be-22d0-4b40-b307-5d2720ef16b2
```
 