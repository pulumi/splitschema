Allows you to set a redrive policy of an SQS Queue
while referencing ARN of the dead letter queue inside the redrive policy.

This is useful when you want to set a dedicated
dead letter queue for a standard or FIFO queue, but need
the dead letter queue to exist before setting the redrive policy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const queue = new aws.sqs.Queue("queue", {});
const ddl = new aws.sqs.Queue("ddl", {redriveAllowPolicy: queue.arn.apply(arn => JSON.stringify({
    redrivePermission: "byQueue",
    sourceQueueArns: [arn],
}))});
const redrivePolicy = new aws.sqs.RedrivePolicy("redrivePolicy", {
    queueUrl: queue.id,
    redrivePolicy: ddl.arn.apply(arn => JSON.stringify({
        deadLetterTargetArn: arn,
        maxReceiveCount: 4,
    })),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

queue = aws.sqs.Queue("queue")
ddl = aws.sqs.Queue("ddl", redrive_allow_policy=queue.arn.apply(lambda arn: json.dumps({
    "redrivePermission": "byQueue",
    "sourceQueueArns": [arn],
})))
redrive_policy = aws.sqs.RedrivePolicy("redrivePolicy",
    queue_url=queue.id,
    redrive_policy=ddl.arn.apply(lambda arn: json.dumps({
        "deadLetterTargetArn": arn,
        "maxReceiveCount": 4,
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
    var queue = new Aws.Sqs.Queue("queue");

    var ddl = new Aws.Sqs.Queue("ddl", new()
    {
        RedriveAllowPolicy = queue.Arn.Apply(arn => JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["redrivePermission"] = "byQueue",
            ["sourceQueueArns"] = new[]
            {
                arn,
            },
        })),
    });

    var redrivePolicy = new Aws.Sqs.RedrivePolicy("redrivePolicy", new()
    {
        QueueUrl = queue.Id,
        RedrivePolicyName = ddl.Arn.Apply(arn => JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["deadLetterTargetArn"] = arn,
            ["maxReceiveCount"] = 4,
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
		queue, err := sqs.NewQueue(ctx, "queue", nil)
		if err != nil {
			return err
		}
		ddl, err := sqs.NewQueue(ctx, "ddl", &sqs.QueueArgs{
			RedriveAllowPolicy: queue.Arn.ApplyT(func(arn string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON0, err := json.Marshal(map[string]interface{}{
					"redrivePermission": "byQueue",
					"sourceQueueArns": []string{
						arn,
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
		_, err = sqs.NewRedrivePolicy(ctx, "redrivePolicy", &sqs.RedrivePolicyArgs{
			QueueUrl: queue.ID(),
			RedrivePolicy: ddl.Arn.ApplyT(func(arn string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON1, err := json.Marshal(map[string]interface{}{
					"deadLetterTargetArn": arn,
					"maxReceiveCount":     4,
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
import com.pulumi.aws.sqs.RedrivePolicy;
import com.pulumi.aws.sqs.RedrivePolicyArgs;
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
        var queue = new Queue("queue");

        var ddl = new Queue("ddl", QueueArgs.builder()        
            .redriveAllowPolicy(queue.arn().applyValue(arn -> serializeJson(
                jsonObject(
                    jsonProperty("redrivePermission", "byQueue"),
                    jsonProperty("sourceQueueArns", jsonArray(arn))
                ))))
            .build());

        var redrivePolicy = new RedrivePolicy("redrivePolicy", RedrivePolicyArgs.builder()        
            .queueUrl(queue.id())
            .redrivePolicy(ddl.arn().applyValue(arn -> serializeJson(
                jsonObject(
                    jsonProperty("deadLetterTargetArn", arn),
                    jsonProperty("maxReceiveCount", 4)
                ))))
            .build());

    }
}
```
```yaml
resources:
  queue:
    type: aws:sqs:Queue
  ddl:
    type: aws:sqs:Queue
    properties:
      redriveAllowPolicy:
        fn::toJSON:
          redrivePermission: byQueue
          sourceQueueArns:
            - ${queue.arn}
  redrivePolicy:
    type: aws:sqs:RedrivePolicy
    properties:
      queueUrl: ${queue.id}
      redrivePolicy:
        fn::toJSON:
          deadLetterTargetArn: ${ddl.arn}
          maxReceiveCount: 4
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SQS Queue Redrive Policies using the queue URL. For example:

```sh
 $ pulumi import aws:sqs/redrivePolicy:RedrivePolicy test https://queue.amazonaws.com/0123456789012/myqueue
```
 