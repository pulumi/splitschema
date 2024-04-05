Provides a resource to manage an S3 Control Bucket Policy.

> This functionality is for managing [S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html). To manage S3 Bucket Policies in an AWS Partition, see the `aws.s3.BucketPolicy` resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3control.BucketPolicy("example", {
    bucket: aws_s3control_bucket.example.arn,
    policy: JSON.stringify({
        Id: "testBucketPolicy",
        Statement: [{
            Action: "s3-outposts:PutBucketLifecycleConfiguration",
            Effect: "Deny",
            Principal: {
                AWS: "*",
            },
            Resource: aws_s3control_bucket.example.arn,
            Sid: "statement1",
        }],
        Version: "2012-10-17",
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.s3control.BucketPolicy("example",
    bucket=aws_s3control_bucket["example"]["arn"],
    policy=json.dumps({
        "Id": "testBucketPolicy",
        "Statement": [{
            "Action": "s3-outposts:PutBucketLifecycleConfiguration",
            "Effect": "Deny",
            "Principal": {
                "AWS": "*",
            },
            "Resource": aws_s3control_bucket["example"]["arn"],
            "Sid": "statement1",
        }],
        "Version": "2012-10-17",
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3Control.BucketPolicy("example", new()
    {
        Bucket = aws_s3control_bucket.Example.Arn,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Id"] = "testBucketPolicy",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = "s3-outposts:PutBucketLifecycleConfiguration",
                    ["Effect"] = "Deny",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["AWS"] = "*",
                    },
                    ["Resource"] = aws_s3control_bucket.Example.Arn,
                    ["Sid"] = "statement1",
                },
            },
            ["Version"] = "2012-10-17",
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Id": "testBucketPolicy",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "s3-outposts:PutBucketLifecycleConfiguration",
					"Effect": "Deny",
					"Principal": map[string]interface{}{
						"AWS": "*",
					},
					"Resource": aws_s3control_bucket.Example.Arn,
					"Sid":      "statement1",
				},
			},
			"Version": "2012-10-17",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = s3control.NewBucketPolicy(ctx, "example", &s3control.BucketPolicyArgs{
			Bucket: pulumi.Any(aws_s3control_bucket.Example.Arn),
			Policy: pulumi.String(json0),
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
import com.pulumi.aws.s3control.BucketPolicy;
import com.pulumi.aws.s3control.BucketPolicyArgs;
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
        var example = new BucketPolicy("example", BucketPolicyArgs.builder()        
            .bucket(aws_s3control_bucket.example().arn())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Id", "testBucketPolicy"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", "s3-outposts:PutBucketLifecycleConfiguration"),
                        jsonProperty("Effect", "Deny"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", "*")
                        )),
                        jsonProperty("Resource", aws_s3control_bucket.example().arn()),
                        jsonProperty("Sid", "statement1")
                    ))),
                    jsonProperty("Version", "2012-10-17")
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3control:BucketPolicy
    properties:
      bucket: ${aws_s3control_bucket.example.arn}
      policy:
        fn::toJSON:
          Id: testBucketPolicy
          Statement:
            - Action: s3-outposts:PutBucketLifecycleConfiguration
              Effect: Deny
              Principal:
                AWS: '*'
              Resource: ${aws_s3control_bucket.example.arn}
              Sid: statement1
          Version: 2012-10-17
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 Control Bucket Policies using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:s3control/bucketPolicy:BucketPolicy example arn:aws:s3-outposts:us-east-1:123456789012:outpost/op-12345678/bucket/example
```
 