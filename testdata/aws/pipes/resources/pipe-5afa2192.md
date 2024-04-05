Resource for managing an AWS EventBridge Pipes Pipe.

You can find out more about EventBridge Pipes in the [User Guide](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html).

EventBridge Pipes are very configurable, and may require IAM permissions to work correctly. More information on the configuration options and IAM permissions can be found in the [User Guide](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes.html).

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = aws.getCallerIdentity({});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: main.then(main => JSON.stringify({
    Version: "2012-10-17",
    Statement: {
        Effect: "Allow",
        Action: "sts:AssumeRole",
        Principal: {
            Service: "pipes.amazonaws.com",
        },
        Condition: {
            StringEquals: {
                "aws:SourceAccount": main.accountId,
            },
        },
    },
}))});
const sourceQueue = new aws.sqs.Queue("sourceQueue", {});
const sourceRolePolicy = new aws.iam.RolePolicy("sourceRolePolicy", {
    role: exampleRole.id,
    policy: sourceQueue.arn.apply(arn => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Action: [
                "sqs:DeleteMessage",
                "sqs:GetQueueAttributes",
                "sqs:ReceiveMessage",
            ],
            Resource: [arn],
        }],
    })),
});
const targetQueue = new aws.sqs.Queue("targetQueue", {});
const targetRolePolicy = new aws.iam.RolePolicy("targetRolePolicy", {
    role: exampleRole.id,
    policy: targetQueue.arn.apply(arn => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Action: ["sqs:SendMessage"],
            Resource: [arn],
        }],
    })),
});
const examplePipe = new aws.pipes.Pipe("examplePipe", {
    roleArn: exampleRole.arn,
    source: sourceQueue.arn,
    target: targetQueue.arn,
}, {
    dependsOn: [
        sourceRolePolicy,
        targetRolePolicy,
    ],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

main = aws.get_caller_identity()
example_role = aws.iam.Role("exampleRole", assume_role_policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": {
        "Effect": "Allow",
        "Action": "sts:AssumeRole",
        "Principal": {
            "Service": "pipes.amazonaws.com",
        },
        "Condition": {
            "StringEquals": {
                "aws:SourceAccount": main.account_id,
            },
        },
    },
}))
source_queue = aws.sqs.Queue("sourceQueue")
source_role_policy = aws.iam.RolePolicy("sourceRolePolicy",
    role=example_role.id,
    policy=source_queue.arn.apply(lambda arn: json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Action": [
                "sqs:DeleteMessage",
                "sqs:GetQueueAttributes",
                "sqs:ReceiveMessage",
            ],
            "Resource": [arn],
        }],
    })))
target_queue = aws.sqs.Queue("targetQueue")
target_role_policy = aws.iam.RolePolicy("targetRolePolicy",
    role=example_role.id,
    policy=target_queue.arn.apply(lambda arn: json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Action": ["sqs:SendMessage"],
            "Resource": [arn],
        }],
    })))
example_pipe = aws.pipes.Pipe("examplePipe",
    role_arn=example_role.arn,
    source=source_queue.arn,
    target=target_queue.arn,
    opts=pulumi.ResourceOptions(depends_on=[
            source_role_policy,
            target_role_policy,
        ]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = Aws.GetCallerIdentity.Invoke();

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new Dictionary<string, object?>
            {
                ["Effect"] = "Allow",
                ["Action"] = "sts:AssumeRole",
                ["Principal"] = new Dictionary<string, object?>
                {
                    ["Service"] = "pipes.amazonaws.com",
                },
                ["Condition"] = new Dictionary<string, object?>
                {
                    ["StringEquals"] = new Dictionary<string, object?>
                    {
                        ["aws:SourceAccount"] = main.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
                    },
                },
            },
        }),
    });

    var sourceQueue = new Aws.Sqs.Queue("sourceQueue");

    var sourceRolePolicy = new Aws.Iam.RolePolicy("sourceRolePolicy", new()
    {
        Role = exampleRole.Id,
        Policy = sourceQueue.Arn.Apply(arn => JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Effect"] = "Allow",
                    ["Action"] = new[]
                    {
                        "sqs:DeleteMessage",
                        "sqs:GetQueueAttributes",
                        "sqs:ReceiveMessage",
                    },
                    ["Resource"] = new[]
                    {
                        arn,
                    },
                },
            },
        })),
    });

    var targetQueue = new Aws.Sqs.Queue("targetQueue");

    var targetRolePolicy = new Aws.Iam.RolePolicy("targetRolePolicy", new()
    {
        Role = exampleRole.Id,
        Policy = targetQueue.Arn.Apply(arn => JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Effect"] = "Allow",
                    ["Action"] = new[]
                    {
                        "sqs:SendMessage",
                    },
                    ["Resource"] = new[]
                    {
                        arn,
                    },
                },
            },
        })),
    });

    var examplePipe = new Aws.Pipes.Pipe("examplePipe", new()
    {
        RoleArn = exampleRole.Arn,
        Source = sourceQueue.Arn,
        Target = targetQueue.Arn,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            sourceRolePolicy,
            targetRolePolicy,
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/pipes"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sqs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": map[string]interface{}{
				"Effect": "Allow",
				"Action": "sts:AssumeRole",
				"Principal": map[string]interface{}{
					"Service": "pipes.amazonaws.com",
				},
				"Condition": map[string]interface{}{
					"StringEquals": map[string]interface{}{
						"aws:SourceAccount": main.AccountId,
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		sourceQueue, err := sqs.NewQueue(ctx, "sourceQueue", nil)
		if err != nil {
			return err
		}
		sourceRolePolicy, err := iam.NewRolePolicy(ctx, "sourceRolePolicy", &iam.RolePolicyArgs{
			Role: exampleRole.ID(),
			Policy: sourceQueue.Arn.ApplyT(func(arn string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON1, err := json.Marshal(map[string]interface{}{
					"Version": "2012-10-17",
					"Statement": []map[string]interface{}{
						map[string]interface{}{
							"Effect": "Allow",
							"Action": []string{
								"sqs:DeleteMessage",
								"sqs:GetQueueAttributes",
								"sqs:ReceiveMessage",
							},
							"Resource": []string{
								arn,
							},
						},
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
		targetQueue, err := sqs.NewQueue(ctx, "targetQueue", nil)
		if err != nil {
			return err
		}
		targetRolePolicy, err := iam.NewRolePolicy(ctx, "targetRolePolicy", &iam.RolePolicyArgs{
			Role: exampleRole.ID(),
			Policy: targetQueue.Arn.ApplyT(func(arn string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON2, err := json.Marshal(map[string]interface{}{
					"Version": "2012-10-17",
					"Statement": []map[string]interface{}{
						map[string]interface{}{
							"Effect": "Allow",
							"Action": []string{
								"sqs:SendMessage",
							},
							"Resource": []string{
								arn,
							},
						},
					},
				})
				if err != nil {
					return _zero, err
				}
				json2 := string(tmpJSON2)
				return pulumi.String(json2), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = pipes.NewPipe(ctx, "examplePipe", &pipes.PipeArgs{
			RoleArn: exampleRole.Arn,
			Source:  sourceQueue.Arn,
			Target:  targetQueue.Arn,
		}, pulumi.DependsOn([]pulumi.Resource{
			sourceRolePolicy,
			targetRolePolicy,
		}))
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.sqs.Queue;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
import com.pulumi.aws.pipes.Pipe;
import com.pulumi.aws.pipes.PipeArgs;
import static com.pulumi.codegen.internal.Serialization.*;
import com.pulumi.resources.CustomResourceOptions;
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
        final var main = AwsFunctions.getCallerIdentity();

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonObject(
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Action", "sts:AssumeRole"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "pipes.amazonaws.com")
                        )),
                        jsonProperty("Condition", jsonObject(
                            jsonProperty("StringEquals", jsonObject(
                                jsonProperty("aws:SourceAccount", main.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
                            ))
                        ))
                    ))
                )))
            .build());

        var sourceQueue = new Queue("sourceQueue");

        var sourceRolePolicy = new RolePolicy("sourceRolePolicy", RolePolicyArgs.builder()        
            .role(exampleRole.id())
            .policy(sourceQueue.arn().applyValue(arn -> serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Action", jsonArray(
                            "sqs:DeleteMessage", 
                            "sqs:GetQueueAttributes", 
                            "sqs:ReceiveMessage"
                        )),
                        jsonProperty("Resource", jsonArray(arn))
                    )))
                ))))
            .build());

        var targetQueue = new Queue("targetQueue");

        var targetRolePolicy = new RolePolicy("targetRolePolicy", RolePolicyArgs.builder()        
            .role(exampleRole.id())
            .policy(targetQueue.arn().applyValue(arn -> serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Action", jsonArray("sqs:SendMessage")),
                        jsonProperty("Resource", jsonArray(arn))
                    )))
                ))))
            .build());

        var examplePipe = new Pipe("examplePipe", PipeArgs.builder()        
            .roleArn(exampleRole.arn())
            .source(sourceQueue.arn())
            .target(targetQueue.arn())
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    sourceRolePolicy,
                    targetRolePolicy)
                .build());

    }
}
```
```yaml
resources:
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            Effect: Allow
            Action: sts:AssumeRole
            Principal:
              Service: pipes.amazonaws.com
            Condition:
              StringEquals:
                aws:SourceAccount: ${main.accountId}
  sourceRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${exampleRole.id}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Effect: Allow
              Action:
                - sqs:DeleteMessage
                - sqs:GetQueueAttributes
                - sqs:ReceiveMessage
              Resource:
                - ${sourceQueue.arn}
  sourceQueue:
    type: aws:sqs:Queue
  targetRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${exampleRole.id}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Effect: Allow
              Action:
                - sqs:SendMessage
              Resource:
                - ${targetQueue.arn}
  targetQueue:
    type: aws:sqs:Queue
  examplePipe:
    type: aws:pipes:Pipe
    properties:
      roleArn: ${exampleRole.arn}
      source: ${sourceQueue.arn}
      target: ${targetQueue.arn}
    options:
      dependson:
        - ${sourceRolePolicy}
        - ${targetRolePolicy}
variables:
  main:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### Enrichment Usage

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.pipes.Pipe;
import com.pulumi.aws.pipes.PipeArgs;
import com.pulumi.aws.pipes.inputs.PipeEnrichmentParametersArgs;
import com.pulumi.aws.pipes.inputs.PipeEnrichmentParametersHttpParametersArgs;
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
        var example = new Pipe("example", PipeArgs.builder()        
            .roleArn(aws_iam_role.example().arn())
            .source(aws_sqs_queue.source().arn())
            .target(aws_sqs_queue.target().arn())
            .enrichment(aws_cloudwatch_event_api_destination.example().arn())
            .enrichmentParameters(PipeEnrichmentParametersArgs.builder()
                .httpParameters(PipeEnrichmentParametersHttpParametersArgs.builder()
                    .pathParameterValues("example-path-param")
                    .headerParameters(Map.ofEntries(
                        Map.entry("example-header", "example-value"),
                        Map.entry("second-example-header", "second-example-value")
                    ))
                    .queryStringParameters(Map.ofEntries(
                        Map.entry("example-query-string", "example-value"),
                        Map.entry("second-example-query-string", "second-example-value")
                    ))
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:pipes:Pipe
    properties:
      roleArn: ${aws_iam_role.example.arn}
      source: ${aws_sqs_queue.source.arn}
      target: ${aws_sqs_queue.target.arn}
      enrichment: ${aws_cloudwatch_event_api_destination.example.arn}
      enrichmentParameters:
        httpParameters:
          pathParameterValues:
            - example-path-param
          headerParameters:
            example-header: example-value
            second-example-header: second-example-value
          queryStringParameters:
            example-query-string: example-value
            second-example-query-string: second-example-value
```
{{% /example %}}
{{% example %}}
### Filter Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.pipes.Pipe("example", {
    roleArn: aws_iam_role.example.arn,
    source: aws_sqs_queue.source.arn,
    target: aws_sqs_queue.target.arn,
    sourceParameters: {
        filterCriteria: {
            filters: [{
                pattern: JSON.stringify({
                    source: ["event-source"],
                }),
            }],
        },
    },
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.pipes.Pipe("example",
    role_arn=aws_iam_role["example"]["arn"],
    source=aws_sqs_queue["source"]["arn"],
    target=aws_sqs_queue["target"]["arn"],
    source_parameters=aws.pipes.PipeSourceParametersArgs(
        filter_criteria=aws.pipes.PipeSourceParametersFilterCriteriaArgs(
            filters=[aws.pipes.PipeSourceParametersFilterCriteriaFilterArgs(
                pattern=json.dumps({
                    "source": ["event-source"],
                }),
            )],
        ),
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
    var example = new Aws.Pipes.Pipe("example", new()
    {
        RoleArn = aws_iam_role.Example.Arn,
        Source = aws_sqs_queue.Source.Arn,
        Target = aws_sqs_queue.Target.Arn,
        SourceParameters = new Aws.Pipes.Inputs.PipeSourceParametersArgs
        {
            FilterCriteria = new Aws.Pipes.Inputs.PipeSourceParametersFilterCriteriaArgs
            {
                Filters = new[]
                {
                    new Aws.Pipes.Inputs.PipeSourceParametersFilterCriteriaFilterArgs
                    {
                        Pattern = JsonSerializer.Serialize(new Dictionary<string, object?>
                        {
                            ["source"] = new[]
                            {
                                "event-source",
                            },
                        }),
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/pipes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"source": []string{
				"event-source",
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = pipes.NewPipe(ctx, "example", &pipes.PipeArgs{
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			Source:  pulumi.Any(aws_sqs_queue.Source.Arn),
			Target:  pulumi.Any(aws_sqs_queue.Target.Arn),
			SourceParameters: &pipes.PipeSourceParametersArgs{
				FilterCriteria: &pipes.PipeSourceParametersFilterCriteriaArgs{
					Filters: pipes.PipeSourceParametersFilterCriteriaFilterArray{
						&pipes.PipeSourceParametersFilterCriteriaFilterArgs{
							Pattern: pulumi.String(json0),
						},
					},
				},
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
import com.pulumi.aws.pipes.Pipe;
import com.pulumi.aws.pipes.PipeArgs;
import com.pulumi.aws.pipes.inputs.PipeSourceParametersArgs;
import com.pulumi.aws.pipes.inputs.PipeSourceParametersFilterCriteriaArgs;
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
        var example = new Pipe("example", PipeArgs.builder()        
            .roleArn(aws_iam_role.example().arn())
            .source(aws_sqs_queue.source().arn())
            .target(aws_sqs_queue.target().arn())
            .sourceParameters(PipeSourceParametersArgs.builder()
                .filterCriteria(PipeSourceParametersFilterCriteriaArgs.builder()
                    .filters(PipeSourceParametersFilterCriteriaFilterArgs.builder()
                        .pattern(serializeJson(
                            jsonObject(
                                jsonProperty("source", jsonArray("event-source"))
                            )))
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:pipes:Pipe
    properties:
      roleArn: ${aws_iam_role.example.arn}
      source: ${aws_sqs_queue.source.arn}
      target: ${aws_sqs_queue.target.arn}
      sourceParameters:
        filterCriteria:
          filters:
            - pattern:
                fn::toJSON:
                  source:
                    - event-source
```
{{% /example %}}
{{% example %}}
### SQS Source and Target Configuration Usage

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.pipes.Pipe;
import com.pulumi.aws.pipes.PipeArgs;
import com.pulumi.aws.pipes.inputs.PipeSourceParametersArgs;
import com.pulumi.aws.pipes.inputs.PipeSourceParametersSqsQueueParametersArgs;
import com.pulumi.aws.pipes.inputs.PipeTargetParametersArgs;
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
        var example = new Pipe("example", PipeArgs.builder()        
            .roleArn(aws_iam_role.example().arn())
            .source(aws_sqs_queue.source().arn())
            .target(aws_sqs_queue.target().arn())
            .sourceParameters(PipeSourceParametersArgs.builder()
                .sqsQueueParameters(PipeSourceParametersSqsQueueParametersArgs.builder()
                    .batchSize(1)
                    .maximumBatchingWindowInSeconds(2)
                    .build())
                .build())
            .targetParameters(PipeTargetParametersArgs.builder()
                .sqsQueue(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:pipes:Pipe
    properties:
      roleArn: ${aws_iam_role.example.arn}
      source: ${aws_sqs_queue.source.arn}
      target: ${aws_sqs_queue.target.arn}
      sourceParameters:
        sqsQueueParameters:
          batchSize: 1
          maximumBatchingWindowInSeconds: 2
      targetParameters:
        sqsQueue:
          - messageDeduplicationId: example-dedupe
            messageGroupId: example-group
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import pipes using the `name`. For example:

```sh
 $ pulumi import aws:pipes/pipe:Pipe example my-pipe
```
 