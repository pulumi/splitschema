Provides a SQS Queue Redrive Allow Policy resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleQueue = new aws.sqs.Queue("exampleQueue", {});
const src = new aws.sqs.Queue("src", {redrivePolicy: exampleQueue.arn.apply(arn => JSON.stringify({
    deadLetterTargetArn: arn,
    maxReceiveCount: 4,
}))});
const exampleRedriveAllowPolicy = new aws.sqs.RedriveAllowPolicy("exampleRedriveAllowPolicy", {
    queueUrl: exampleQueue.id,
    redriveAllowPolicy: src.arn.apply(arn => JSON.stringify({
        redrivePermission: "byQueue",
        sourceQueueArns: [arn],
    })),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_queue = aws.sqs.Queue("exampleQueue")
src = aws.sqs.Queue("src", redrive_policy=example_queue.arn.apply(lambda arn: json.dumps({
    "deadLetterTargetArn": arn,
    "maxReceiveCount": 4,
})))
example_redrive_allow_policy = aws.sqs.RedriveAllowPolicy("exampleRedriveAllowPolicy",
    queue_url=example_queue.id,
    redrive_allow_policy=src.arn.apply(lambda arn: json.dumps({
        "redrivePermission": "byQueue",
        "sourceQueueArns": [arn],
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
    var exampleQueue = new Aws.Sqs.Queue("exampleQueue");

    var src = new Aws.Sqs.Queue("src", new()
    {
        RedrivePolicy = exampleQueue.Arn.Apply(arn => JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["deadLetterTargetArn"] = arn,
            ["maxReceiveCount"] = 4,
        })),
    });

    var exampleRedriveAllowPolicy = new Aws.Sqs.RedriveAllowPolicy("exampleRedriveAllowPolicy", new()
    {
        QueueUrl = exampleQueue.Id,
        RedriveAllowPolicyName = src.Arn.Apply(arn => JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["redrivePermission"] = "byQueue",
            ["sourceQueueArns"] = new[]
            {
                arn,
            },
        })),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleQueue, err := sqs.NewQueue(ctx, "exampleQueue", nil)
		if err != nil {
			return err
		}
		src, err := sqs.NewQueue(ctx, "src", &sqs.QueueArgs{
			RedrivePolicy: exampleQueue.Arn.ApplyT(func(arn string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON0, err := json.Marshal(map[string]interface{}{
					"deadLetterTargetArn": arn,
					"maxReceiveCount":     4,
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
		_, err = sqs.NewRedriveAllowPolicy(ctx, "exampleRedriveAllowPolicy", &sqs.RedriveAllowPolicyArgs{
			QueueUrl: exampleQueue.ID(),
			RedriveAllowPolicy: src.Arn.ApplyT(func(arn string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON1, err := json.Marshal(map[string]interface{}{
					"redrivePermission": "byQueue",
					"sourceQueueArns": []string{
						arn,
					},
				})
				if err != nil {
					return _zero, err
				}
				json1 := string(tmpJSON1)
				return pulumi.String(json1), nil
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
import com.pulumi.aws.sqs.Queue;
import com.pulumi.aws.sqs.QueueArgs;
import com.pulumi.aws.sqs.RedriveAllowPolicy;
import com.pulumi.aws.sqs.RedriveAllowPolicyArgs;
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
        var exampleQueue = new Queue("exampleQueue");

        var src = new Queue("src", QueueArgs.builder()        
            .redrivePolicy(exampleQueue.arn().applyValue(arn -> serializeJson(
                jsonObject(
                    jsonProperty("deadLetterTargetArn", arn),
                    jsonProperty("maxReceiveCount", 4)
                ))))
            .build());

        var exampleRedriveAllowPolicy = new RedriveAllowPolicy("exampleRedriveAllowPolicy", RedriveAllowPolicyArgs.builder()        
            .queueUrl(exampleQueue.id())
            .redriveAllowPolicy(src.arn().applyValue(arn -> serializeJson(
                jsonObject(
                    jsonProperty("redrivePermission", "byQueue"),
                    jsonProperty("sourceQueueArns", jsonArray(arn))
                ))))
            .build());

    }
}
```
```yaml
resources:
  src:
    type: aws:sqs:Queue
    properties:
      redrivePolicy:
        fn::toJSON:
          deadLetterTargetArn: ${exampleQueue.arn}
          maxReceiveCount: 4
  exampleQueue:
    type: aws:sqs:Queue
  exampleRedriveAllowPolicy:
    type: aws:sqs:RedriveAllowPolicy
    properties:
      queueUrl: ${exampleQueue.id}
      redriveAllowPolicy:
        fn::toJSON:
          redrivePermission: byQueue
          sourceQueueArns:
            - ${src.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SQS Queue Redrive Allow Policies using the queue URL. For example:

```sh
 $ pulumi import aws:sqs/redriveAllowPolicy:RedriveAllowPolicy test https://queue.amazonaws.com/0123456789012/myqueue
```
 