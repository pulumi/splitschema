Provides the ability to register instances and containers with an Application Load Balancer (ALB) or Network Load Balancer (NLB) target group. For attaching resources with Elastic Load Balancer (ELB), see the `aws.elb.Attachment` resource.

> **Note:** `aws.alb.TargetGroupAttachment` is known as `aws.lb.TargetGroupAttachment`. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testTargetGroup = new aws.lb.TargetGroup("testTargetGroup", {});
// ... other configuration ...
const testInstance = new aws.ec2.Instance("testInstance", {});
// ... other configuration ...
const testTargetGroupAttachment = new aws.lb.TargetGroupAttachment("testTargetGroupAttachment", {
    targetGroupArn: testTargetGroup.arn,
    targetId: testInstance.id,
    port: 80,
});
```
```python
import pulumi
import pulumi_aws as aws

test_target_group = aws.lb.TargetGroup("testTargetGroup")
# ... other configuration ...
test_instance = aws.ec2.Instance("testInstance")
# ... other configuration ...
test_target_group_attachment = aws.lb.TargetGroupAttachment("testTargetGroupAttachment",
    target_group_arn=test_target_group.arn,
    target_id=test_instance.id,
    port=80)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testTargetGroup = new Aws.LB.TargetGroup("testTargetGroup");

    // ... other configuration ...
    var testInstance = new Aws.Ec2.Instance("testInstance");

    // ... other configuration ...
    var testTargetGroupAttachment = new Aws.LB.TargetGroupAttachment("testTargetGroupAttachment", new()
    {
        TargetGroupArn = testTargetGroup.Arn,
        TargetId = testInstance.Id,
        Port = 80,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testTargetGroup, err := lb.NewTargetGroup(ctx, "testTargetGroup", nil)
		if err != nil {
			return err
		}
		testInstance, err := ec2.NewInstance(ctx, "testInstance", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewTargetGroupAttachment(ctx, "testTargetGroupAttachment", &lb.TargetGroupAttachmentArgs{
			TargetGroupArn: testTargetGroup.Arn,
			TargetId:       testInstance.ID(),
			Port:           pulumi.Int(80),
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
import com.pulumi.aws.lb.TargetGroup;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.lb.TargetGroupAttachment;
import com.pulumi.aws.lb.TargetGroupAttachmentArgs;
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
        var testTargetGroup = new TargetGroup("testTargetGroup");

        var testInstance = new Instance("testInstance");

        var testTargetGroupAttachment = new TargetGroupAttachment("testTargetGroupAttachment", TargetGroupAttachmentArgs.builder()        
            .targetGroupArn(testTargetGroup.arn())
            .targetId(testInstance.id())
            .port(80)
            .build());

    }
}
```
```yaml
resources:
  testTargetGroupAttachment:
    type: aws:lb:TargetGroupAttachment
    properties:
      targetGroupArn: ${testTargetGroup.arn}
      targetId: ${testInstance.id}
      port: 80
  testTargetGroup:
    type: aws:lb:TargetGroup
  testInstance:
    type: aws:ec2:Instance
```
{{% /example %}}
{{% example %}}
### Lambda Target

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testTargetGroup = new aws.lb.TargetGroup("testTargetGroup", {targetType: "lambda"});
const testFunction = new aws.lambda.Function("testFunction", {});
// ... other configuration ...
const withLb = new aws.lambda.Permission("withLb", {
    action: "lambda:InvokeFunction",
    "function": testFunction.name,
    principal: "elasticloadbalancing.amazonaws.com",
    sourceArn: testTargetGroup.arn,
});
const testTargetGroupAttachment = new aws.lb.TargetGroupAttachment("testTargetGroupAttachment", {
    targetGroupArn: testTargetGroup.arn,
    targetId: testFunction.arn,
}, {
    dependsOn: [withLb],
});
```
```python
import pulumi
import pulumi_aws as aws

test_target_group = aws.lb.TargetGroup("testTargetGroup", target_type="lambda")
test_function = aws.lambda_.Function("testFunction")
# ... other configuration ...
with_lb = aws.lambda_.Permission("withLb",
    action="lambda:InvokeFunction",
    function=test_function.name,
    principal="elasticloadbalancing.amazonaws.com",
    source_arn=test_target_group.arn)
test_target_group_attachment = aws.lb.TargetGroupAttachment("testTargetGroupAttachment",
    target_group_arn=test_target_group.arn,
    target_id=test_function.arn,
    opts=pulumi.ResourceOptions(depends_on=[with_lb]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testTargetGroup = new Aws.LB.TargetGroup("testTargetGroup", new()
    {
        TargetType = "lambda",
    });

    var testFunction = new Aws.Lambda.Function("testFunction");

    // ... other configuration ...
    var withLb = new Aws.Lambda.Permission("withLb", new()
    {
        Action = "lambda:InvokeFunction",
        Function = testFunction.Name,
        Principal = "elasticloadbalancing.amazonaws.com",
        SourceArn = testTargetGroup.Arn,
    });

    var testTargetGroupAttachment = new Aws.LB.TargetGroupAttachment("testTargetGroupAttachment", new()
    {
        TargetGroupArn = testTargetGroup.Arn,
        TargetId = testFunction.Arn,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            withLb,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testTargetGroup, err := lb.NewTargetGroup(ctx, "testTargetGroup", &lb.TargetGroupArgs{
			TargetType: pulumi.String("lambda"),
		})
		if err != nil {
			return err
		}
		testFunction, err := lambda.NewFunction(ctx, "testFunction", nil)
		if err != nil {
			return err
		}
		withLb, err := lambda.NewPermission(ctx, "withLb", &lambda.PermissionArgs{
			Action:    pulumi.String("lambda:InvokeFunction"),
			Function:  testFunction.Name,
			Principal: pulumi.String("elasticloadbalancing.amazonaws.com"),
			SourceArn: testTargetGroup.Arn,
		})
		if err != nil {
			return err
		}
		_, err = lb.NewTargetGroupAttachment(ctx, "testTargetGroupAttachment", &lb.TargetGroupAttachmentArgs{
			TargetGroupArn: testTargetGroup.Arn,
			TargetId:       testFunction.Arn,
		}, pulumi.DependsOn([]pulumi.Resource{
			withLb,
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
import com.pulumi.aws.lb.TargetGroup;
import com.pulumi.aws.lb.TargetGroupArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.Permission;
import com.pulumi.aws.lambda.PermissionArgs;
import com.pulumi.aws.lb.TargetGroupAttachment;
import com.pulumi.aws.lb.TargetGroupAttachmentArgs;
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
        var testTargetGroup = new TargetGroup("testTargetGroup", TargetGroupArgs.builder()        
            .targetType("lambda")
            .build());

        var testFunction = new Function("testFunction");

        var withLb = new Permission("withLb", PermissionArgs.builder()        
            .action("lambda:InvokeFunction")
            .function(testFunction.name())
            .principal("elasticloadbalancing.amazonaws.com")
            .sourceArn(testTargetGroup.arn())
            .build());

        var testTargetGroupAttachment = new TargetGroupAttachment("testTargetGroupAttachment", TargetGroupAttachmentArgs.builder()        
            .targetGroupArn(testTargetGroup.arn())
            .targetId(testFunction.arn())
            .build(), CustomResourceOptions.builder()
                .dependsOn(withLb)
                .build());

    }
}
```
```yaml
resources:
  withLb:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunction
      function: ${testFunction.name}
      principal: elasticloadbalancing.amazonaws.com
      sourceArn: ${testTargetGroup.arn}
  testTargetGroup:
    type: aws:lb:TargetGroup
    properties:
      targetType: lambda
  testFunction:
    type: aws:lambda:Function
  testTargetGroupAttachment:
    type: aws:lb:TargetGroupAttachment
    properties:
      targetGroupArn: ${testTargetGroup.arn}
      targetId: ${testFunction.arn}
    options:
      dependson:
        - ${withLb}
```
{{% /example %}}
{{% example %}}
### Registering Multiple Targets

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleInstance: aws.ec2.Instance[] = [];
for (const range = {value: 0}; range.value < 3; range.value++) {
    exampleInstance.push(new aws.ec2.Instance(`exampleInstance-${range.value}`, {}));
}
// ... other configuration ...
const exampleTargetGroup = new aws.lb.TargetGroup("exampleTargetGroup", {});
// ... other configuration ...
const exampleTargetGroupAttachment: aws.lb.TargetGroupAttachment[] = [];
pulumi.all(exampleInstance.map((v, k) => [k, v]).reduce((__obj, [, ]) => ({ ...__obj, [v.id]: v }))).apply(rangeBody => {
    for (const range of Object.entries(rangeBody).map(([k, v]) => ({key: k, value: v}))) {
        exampleTargetGroupAttachment.push(new aws.lb.TargetGroupAttachment(`exampleTargetGroupAttachment-${range.key}`, {
            targetGroupArn: exampleTargetGroup.arn,
            targetId: range.value.id,
            port: 80,
        }));
    }
});
```
```python
import pulumi
import pulumi_aws as aws

example_instance = []
for range in [{"value": i} for i in range(0, 3)]:
    example_instance.append(aws.ec2.Instance(f"exampleInstance-{range['value']}"))
# ... other configuration ...
example_target_group = aws.lb.TargetGroup("exampleTargetGroup")
# ... other configuration ...
example_target_group_attachment = []
def create_example_target_group_attachment(range_body):
    for range in [{"key": k, "value": v} for [k, v] in enumerate(range_body)]:
        example_target_group_attachment.append(aws.lb.TargetGroupAttachment(f"exampleTargetGroupAttachment-{range['key']}",
            target_group_arn=example_target_group.arn,
            target_id=range["value"],
            port=80))

pulumi.Output.all({v.id: v for k, v in example_instance}).apply(lambda resolved_outputs: create_example_target_group_attachment(resolved_outputs[0]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleInstance = new List<Aws.Ec2.Instance>();
    for (var rangeIndex = 0; rangeIndex < 3; rangeIndex++)
    {
        var range = new { Value = rangeIndex };
        exampleInstance.Add(new Aws.Ec2.Instance($"exampleInstance-{range.Value}", new()
        {
        }));
    }
    // ... other configuration ...
    var exampleTargetGroup = new Aws.LB.TargetGroup("exampleTargetGroup");

    // ... other configuration ...
    var exampleTargetGroupAttachment = new List<Aws.LB.TargetGroupAttachment>();
    foreach (var range in exampleInstance.Select((value, i) => new { Key = i.ToString(), Value = pair.Value }).Select(pair => new { pair.Key, pair.Value }))
    {
        exampleTargetGroupAttachment.Add(new Aws.LB.TargetGroupAttachment($"exampleTargetGroupAttachment-{range.Key}", new()
        {
            TargetGroupArn = exampleTargetGroup.Arn,
            TargetId = range.Value.Id,
            Port = 80,
        }));
    }
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		var exampleInstance []*ec2.Instance
		for index := 0; index < 3; index++ {
			key0 := index
			_ := index
			__res, err := ec2.NewInstance(ctx, fmt.Sprintf("exampleInstance-%v", key0), nil)
			if err != nil {
				return err
			}
			exampleInstance = append(exampleInstance, __res)
		}
		exampleTargetGroup, err := lb.NewTargetGroup(ctx, "exampleTargetGroup", nil)
		if err != nil {
			return err
		}
		var exampleTargetGroupAttachment []*lb.TargetGroupAttachment
		for key0, val0 := range "TODO: For expression" {
			__res, err := lb.NewTargetGroupAttachment(ctx, fmt.Sprintf("exampleTargetGroupAttachment-%v", key0), &lb.TargetGroupAttachmentArgs{
				TargetGroupArn: exampleTargetGroup.Arn,
				TargetId:       pulumi.String(val0),
				Port:           pulumi.Int(80),
			})
			if err != nil {
				return err
			}
			exampleTargetGroupAttachment = append(exampleTargetGroupAttachment, __res)
		}
		return nil
	})
}
```
```yaml
resources:
  exampleInstance:
    type: aws:ec2:Instance
    options: {}
  exampleTargetGroup:
    type: aws:lb:TargetGroup
  exampleTargetGroupAttachment:
    type: aws:lb:TargetGroupAttachment
    properties:
      targetGroupArn: ${exampleTargetGroup.arn}
      targetId: ${range.value.id}
      port: 80
    options: {}
```
{{% /example %}}
{{% /examples %}}

## Import

You cannot import Target Group Attachments. 