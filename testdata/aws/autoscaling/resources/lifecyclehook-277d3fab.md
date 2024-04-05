Provides an AutoScaling Lifecycle Hook resource.

> **NOTE:** This provider has two types of ways you can add lifecycle hooks - via
the `initial_lifecycle_hook` attribute from the
`aws.autoscaling.Group`
resource, or via this one. Hooks added via this resource will not be added
until the autoscaling group has been created, and depending on your
capacity
settings, after the initial instances have been launched, creating unintended
behavior. If you need hooks to run on all instances, add them with
`initial_lifecycle_hook` in
`aws.autoscaling.Group`,
but take care to not duplicate those hooks with this resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foobarGroup = new aws.autoscaling.Group("foobarGroup", {
    availabilityZones: ["us-west-2a"],
    healthCheckType: "EC2",
    terminationPolicies: ["OldestInstance"],
    tags: [{
        key: "Foo",
        value: "foo-bar",
        propagateAtLaunch: true,
    }],
});
const foobarLifecycleHook = new aws.autoscaling.LifecycleHook("foobarLifecycleHook", {
    autoscalingGroupName: foobarGroup.name,
    defaultResult: "CONTINUE",
    heartbeatTimeout: 2000,
    lifecycleTransition: "autoscaling:EC2_INSTANCE_LAUNCHING",
    notificationMetadata: JSON.stringify({
        foo: "bar",
    }),
    notificationTargetArn: "arn:aws:sqs:us-east-1:444455556666:queue1*",
    roleArn: "arn:aws:iam::123456789012:role/S3Access",
});
```
```python
import pulumi
import json
import pulumi_aws as aws

foobar_group = aws.autoscaling.Group("foobarGroup",
    availability_zones=["us-west-2a"],
    health_check_type="EC2",
    termination_policies=["OldestInstance"],
    tags=[aws.autoscaling.GroupTagArgs(
        key="Foo",
        value="foo-bar",
        propagate_at_launch=True,
    )])
foobar_lifecycle_hook = aws.autoscaling.LifecycleHook("foobarLifecycleHook",
    autoscaling_group_name=foobar_group.name,
    default_result="CONTINUE",
    heartbeat_timeout=2000,
    lifecycle_transition="autoscaling:EC2_INSTANCE_LAUNCHING",
    notification_metadata=json.dumps({
        "foo": "bar",
    }),
    notification_target_arn="arn:aws:sqs:us-east-1:444455556666:queue1*",
    role_arn="arn:aws:iam::123456789012:role/S3Access")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foobarGroup = new Aws.AutoScaling.Group("foobarGroup", new()
    {
        AvailabilityZones = new[]
        {
            "us-west-2a",
        },
        HealthCheckType = "EC2",
        TerminationPolicies = new[]
        {
            "OldestInstance",
        },
        Tags = new[]
        {
            new Aws.AutoScaling.Inputs.GroupTagArgs
            {
                Key = "Foo",
                Value = "foo-bar",
                PropagateAtLaunch = true,
            },
        },
    });

    var foobarLifecycleHook = new Aws.AutoScaling.LifecycleHook("foobarLifecycleHook", new()
    {
        AutoscalingGroupName = foobarGroup.Name,
        DefaultResult = "CONTINUE",
        HeartbeatTimeout = 2000,
        LifecycleTransition = "autoscaling:EC2_INSTANCE_LAUNCHING",
        NotificationMetadata = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["foo"] = "bar",
        }),
        NotificationTargetArn = "arn:aws:sqs:us-east-1:444455556666:queue1*",
        RoleArn = "arn:aws:iam::123456789012:role/S3Access",
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foobarGroup, err := autoscaling.NewGroup(ctx, "foobarGroup", &autoscaling.GroupArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
			},
			HealthCheckType: pulumi.String("EC2"),
			TerminationPolicies: pulumi.StringArray{
				pulumi.String("OldestInstance"),
			},
			Tags: autoscaling.GroupTagArray{
				&autoscaling.GroupTagArgs{
					Key:               pulumi.String("Foo"),
					Value:             pulumi.String("foo-bar"),
					PropagateAtLaunch: pulumi.Bool(true),
				},
			},
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"foo": "bar",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = autoscaling.NewLifecycleHook(ctx, "foobarLifecycleHook", &autoscaling.LifecycleHookArgs{
			AutoscalingGroupName:  foobarGroup.Name,
			DefaultResult:         pulumi.String("CONTINUE"),
			HeartbeatTimeout:      pulumi.Int(2000),
			LifecycleTransition:   pulumi.String("autoscaling:EC2_INSTANCE_LAUNCHING"),
			NotificationMetadata:  pulumi.String(json0),
			NotificationTargetArn: pulumi.String("arn:aws:sqs:us-east-1:444455556666:queue1*"),
			RoleArn:               pulumi.String("arn:aws:iam::123456789012:role/S3Access"),
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
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
import com.pulumi.aws.autoscaling.inputs.GroupTagArgs;
import com.pulumi.aws.autoscaling.LifecycleHook;
import com.pulumi.aws.autoscaling.LifecycleHookArgs;
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
        var foobarGroup = new Group("foobarGroup", GroupArgs.builder()        
            .availabilityZones("us-west-2a")
            .healthCheckType("EC2")
            .terminationPolicies("OldestInstance")
            .tags(GroupTagArgs.builder()
                .key("Foo")
                .value("foo-bar")
                .propagateAtLaunch(true)
                .build())
            .build());

        var foobarLifecycleHook = new LifecycleHook("foobarLifecycleHook", LifecycleHookArgs.builder()        
            .autoscalingGroupName(foobarGroup.name())
            .defaultResult("CONTINUE")
            .heartbeatTimeout(2000)
            .lifecycleTransition("autoscaling:EC2_INSTANCE_LAUNCHING")
            .notificationMetadata(serializeJson(
                jsonObject(
                    jsonProperty("foo", "bar")
                )))
            .notificationTargetArn("arn:aws:sqs:us-east-1:444455556666:queue1*")
            .roleArn("arn:aws:iam::123456789012:role/S3Access")
            .build());

    }
}
```
```yaml
resources:
  foobarGroup:
    type: aws:autoscaling:Group
    properties:
      availabilityZones:
        - us-west-2a
      healthCheckType: EC2
      terminationPolicies:
        - OldestInstance
      tags:
        - key: Foo
          value: foo-bar
          propagateAtLaunch: true
  foobarLifecycleHook:
    type: aws:autoscaling:LifecycleHook
    properties:
      autoscalingGroupName: ${foobarGroup.name}
      defaultResult: CONTINUE
      heartbeatTimeout: 2000
      lifecycleTransition: autoscaling:EC2_INSTANCE_LAUNCHING
      notificationMetadata:
        fn::toJSON:
          foo: bar
      notificationTargetArn: arn:aws:sqs:us-east-1:444455556666:queue1*
      roleArn: arn:aws:iam::123456789012:role/S3Access
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AutoScaling Lifecycle Hooks using the role autoscaling_group_name and name separated by `/`. For example:

```sh
 $ pulumi import aws:autoscaling/lifecycleHook:LifecycleHook test-lifecycle-hook asg-name/lifecycle-hook-name
```
 