Provides an Inspector Classic Assessment Template

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.inspector.AssessmentTemplate("example", {
    targetArn: aws_inspector_assessment_target.example.arn,
    duration: 3600,
    rulesPackageArns: [
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p",
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-H5hpSawc",
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-JJOtZiqQ",
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-vg5GGHSD",
    ],
    eventSubscriptions: [{
        event: "ASSESSMENT_RUN_COMPLETED",
        topicArn: aws_sns_topic.example.arn,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.inspector.AssessmentTemplate("example",
    target_arn=aws_inspector_assessment_target["example"]["arn"],
    duration=3600,
    rules_package_arns=[
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p",
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-H5hpSawc",
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-JJOtZiqQ",
        "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-vg5GGHSD",
    ],
    event_subscriptions=[aws.inspector.AssessmentTemplateEventSubscriptionArgs(
        event="ASSESSMENT_RUN_COMPLETED",
        topic_arn=aws_sns_topic["example"]["arn"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Inspector.AssessmentTemplate("example", new()
    {
        TargetArn = aws_inspector_assessment_target.Example.Arn,
        Duration = 3600,
        RulesPackageArns = new[]
        {
            "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p",
            "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-H5hpSawc",
            "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-JJOtZiqQ",
            "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-vg5GGHSD",
        },
        EventSubscriptions = new[]
        {
            new Aws.Inspector.Inputs.AssessmentTemplateEventSubscriptionArgs
            {
                Event = "ASSESSMENT_RUN_COMPLETED",
                TopicArn = aws_sns_topic.Example.Arn,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/inspector"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := inspector.NewAssessmentTemplate(ctx, "example", &inspector.AssessmentTemplateArgs{
			TargetArn: pulumi.Any(aws_inspector_assessment_target.Example.Arn),
			Duration:  pulumi.Int(3600),
			RulesPackageArns: pulumi.StringArray{
				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p"),
				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-H5hpSawc"),
				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-JJOtZiqQ"),
				pulumi.String("arn:aws:inspector:us-west-2:758058086616:rulespackage/0-vg5GGHSD"),
			},
			EventSubscriptions: inspector.AssessmentTemplateEventSubscriptionArray{
				&inspector.AssessmentTemplateEventSubscriptionArgs{
					Event:    pulumi.String("ASSESSMENT_RUN_COMPLETED"),
					TopicArn: pulumi.Any(aws_sns_topic.Example.Arn),
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
import com.pulumi.aws.inspector.AssessmentTemplate;
import com.pulumi.aws.inspector.AssessmentTemplateArgs;
import com.pulumi.aws.inspector.inputs.AssessmentTemplateEventSubscriptionArgs;
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
        var example = new AssessmentTemplate("example", AssessmentTemplateArgs.builder()        
            .targetArn(aws_inspector_assessment_target.example().arn())
            .duration(3600)
            .rulesPackageArns(            
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p",
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-H5hpSawc",
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-JJOtZiqQ",
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-vg5GGHSD")
            .eventSubscriptions(AssessmentTemplateEventSubscriptionArgs.builder()
                .event("ASSESSMENT_RUN_COMPLETED")
                .topicArn(aws_sns_topic.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:inspector:AssessmentTemplate
    properties:
      targetArn: ${aws_inspector_assessment_target.example.arn}
      duration: 3600
      rulesPackageArns:
        - arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p
        - arn:aws:inspector:us-west-2:758058086616:rulespackage/0-H5hpSawc
        - arn:aws:inspector:us-west-2:758058086616:rulespackage/0-JJOtZiqQ
        - arn:aws:inspector:us-west-2:758058086616:rulespackage/0-vg5GGHSD
      eventSubscriptions:
        - event: ASSESSMENT_RUN_COMPLETED
          topicArn: ${aws_sns_topic.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_inspector_assessment_template` using the template assessment ARN. For example:

```sh
 $ pulumi import aws:inspector/assessmentTemplate:AssessmentTemplate example arn:aws:inspector:us-west-2:123456789012:target/0-9IaAzhGR/template/0-WEcjR8CH
```
 