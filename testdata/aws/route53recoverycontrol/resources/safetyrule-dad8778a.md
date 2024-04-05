Provides an AWS Route 53 Recovery Control Config Safety Rule

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53recoverycontrol.SafetyRule("example", {
    assertedControls: [aws_route53recoverycontrolconfig_routing_control.example.arn],
    controlPanelArn: "arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8",
    waitPeriodMs: 5000,
    ruleConfig: {
        inverted: false,
        threshold: 1,
        type: "ATLEAST",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53recoverycontrol.SafetyRule("example",
    asserted_controls=[aws_route53recoverycontrolconfig_routing_control["example"]["arn"]],
    control_panel_arn="arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8",
    wait_period_ms=5000,
    rule_config=aws.route53recoverycontrol.SafetyRuleRuleConfigArgs(
        inverted=False,
        threshold=1,
        type="ATLEAST",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53RecoveryControl.SafetyRule("example", new()
    {
        AssertedControls = new[]
        {
            aws_route53recoverycontrolconfig_routing_control.Example.Arn,
        },
        ControlPanelArn = "arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8",
        WaitPeriodMs = 5000,
        RuleConfig = new Aws.Route53RecoveryControl.Inputs.SafetyRuleRuleConfigArgs
        {
            Inverted = false,
            Threshold = 1,
            Type = "ATLEAST",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53recoverycontrol"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53recoverycontrol.NewSafetyRule(ctx, "example", &route53recoverycontrol.SafetyRuleArgs{
			AssertedControls: pulumi.StringArray{
				aws_route53recoverycontrolconfig_routing_control.Example.Arn,
			},
			ControlPanelArn: pulumi.String("arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8"),
			WaitPeriodMs:    pulumi.Int(5000),
			RuleConfig: &route53recoverycontrol.SafetyRuleRuleConfigArgs{
				Inverted:  pulumi.Bool(false),
				Threshold: pulumi.Int(1),
				Type:      pulumi.String("ATLEAST"),
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
import com.pulumi.aws.route53recoverycontrol.SafetyRule;
import com.pulumi.aws.route53recoverycontrol.SafetyRuleArgs;
import com.pulumi.aws.route53recoverycontrol.inputs.SafetyRuleRuleConfigArgs;
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
        var example = new SafetyRule("example", SafetyRuleArgs.builder()        
            .assertedControls(aws_route53recoverycontrolconfig_routing_control.example().arn())
            .controlPanelArn("arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8")
            .waitPeriodMs(5000)
            .ruleConfig(SafetyRuleRuleConfigArgs.builder()
                .inverted(false)
                .threshold(1)
                .type("ATLEAST")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53recoverycontrol:SafetyRule
    properties:
      assertedControls:
        - ${aws_route53recoverycontrolconfig_routing_control.example.arn}
      controlPanelArn: arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8
      waitPeriodMs: 5000
      ruleConfig:
        inverted: false
        threshold: 1
        type: ATLEAST
```

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53recoverycontrol.SafetyRule("example", {
    controlPanelArn: "arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8",
    waitPeriodMs: 5000,
    gatingControls: [aws_route53recoverycontrolconfig_routing_control.example.arn],
    targetControls: [aws_route53recoverycontrolconfig_routing_control.example.arn],
    ruleConfig: {
        inverted: false,
        threshold: 1,
        type: "ATLEAST",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53recoverycontrol.SafetyRule("example",
    control_panel_arn="arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8",
    wait_period_ms=5000,
    gating_controls=[aws_route53recoverycontrolconfig_routing_control["example"]["arn"]],
    target_controls=[aws_route53recoverycontrolconfig_routing_control["example"]["arn"]],
    rule_config=aws.route53recoverycontrol.SafetyRuleRuleConfigArgs(
        inverted=False,
        threshold=1,
        type="ATLEAST",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53RecoveryControl.SafetyRule("example", new()
    {
        ControlPanelArn = "arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8",
        WaitPeriodMs = 5000,
        GatingControls = new[]
        {
            aws_route53recoverycontrolconfig_routing_control.Example.Arn,
        },
        TargetControls = new[]
        {
            aws_route53recoverycontrolconfig_routing_control.Example.Arn,
        },
        RuleConfig = new Aws.Route53RecoveryControl.Inputs.SafetyRuleRuleConfigArgs
        {
            Inverted = false,
            Threshold = 1,
            Type = "ATLEAST",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53recoverycontrol"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53recoverycontrol.NewSafetyRule(ctx, "example", &route53recoverycontrol.SafetyRuleArgs{
			ControlPanelArn: pulumi.String("arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8"),
			WaitPeriodMs:    pulumi.Int(5000),
			GatingControls: pulumi.StringArray{
				aws_route53recoverycontrolconfig_routing_control.Example.Arn,
			},
			TargetControls: pulumi.StringArray{
				aws_route53recoverycontrolconfig_routing_control.Example.Arn,
			},
			RuleConfig: &route53recoverycontrol.SafetyRuleRuleConfigArgs{
				Inverted:  pulumi.Bool(false),
				Threshold: pulumi.Int(1),
				Type:      pulumi.String("ATLEAST"),
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
import com.pulumi.aws.route53recoverycontrol.SafetyRule;
import com.pulumi.aws.route53recoverycontrol.SafetyRuleArgs;
import com.pulumi.aws.route53recoverycontrol.inputs.SafetyRuleRuleConfigArgs;
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
        var example = new SafetyRule("example", SafetyRuleArgs.builder()        
            .controlPanelArn("arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8")
            .waitPeriodMs(5000)
            .gatingControls(aws_route53recoverycontrolconfig_routing_control.example().arn())
            .targetControls(aws_route53recoverycontrolconfig_routing_control.example().arn())
            .ruleConfig(SafetyRuleRuleConfigArgs.builder()
                .inverted(false)
                .threshold(1)
                .type("ATLEAST")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53recoverycontrol:SafetyRule
    properties:
      controlPanelArn: arn:aws:route53-recovery-control::313517334327:controlpanel/abd5fbfc052d4844a082dbf400f61da8
      waitPeriodMs: 5000
      gatingControls:
        - ${aws_route53recoverycontrolconfig_routing_control.example.arn}
      targetControls:
        - ${aws_route53recoverycontrolconfig_routing_control.example.arn}
      ruleConfig:
        inverted: false
        threshold: 1
        type: ATLEAST
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route53 Recovery Control Config Safety Rule using the safety rule ARN. For example:

```sh
 $ pulumi import aws:route53recoverycontrol/safetyRule:SafetyRule myrule arn:aws:route53-recovery-control::313517334327:controlpanel/1bfba17df8684f5dab0467b71424f7e8/safetyrule/3bacc77003364c0f
```
 