Provides an AWS Route 53 Recovery Readiness Resource Set.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53recoveryreadiness.ResourceSet("example", {
    resourceSetName: my_cw_alarm_set,
    resourceSetType: "AWS::CloudWatch::Alarm",
    resources: [{
        resourceArn: aws_cloudwatch_metric_alarm.example.arn,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53recoveryreadiness.ResourceSet("example",
    resource_set_name=my_cw_alarm_set,
    resource_set_type="AWS::CloudWatch::Alarm",
    resources=[aws.route53recoveryreadiness.ResourceSetResourceArgs(
        resource_arn=aws_cloudwatch_metric_alarm["example"]["arn"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53RecoveryReadiness.ResourceSet("example", new()
    {
        ResourceSetName = my_cw_alarm_set,
        ResourceSetType = "AWS::CloudWatch::Alarm",
        Resources = new[]
        {
            new Aws.Route53RecoveryReadiness.Inputs.ResourceSetResourceArgs
            {
                ResourceArn = aws_cloudwatch_metric_alarm.Example.Arn,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53recoveryreadiness"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53recoveryreadiness.NewResourceSet(ctx, "example", &route53recoveryreadiness.ResourceSetArgs{
			ResourceSetName: pulumi.Any(my_cw_alarm_set),
			ResourceSetType: pulumi.String("AWS::CloudWatch::Alarm"),
			Resources: route53recoveryreadiness.ResourceSetResourceArray{
				&route53recoveryreadiness.ResourceSetResourceArgs{
					ResourceArn: pulumi.Any(aws_cloudwatch_metric_alarm.Example.Arn),
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
import com.pulumi.aws.route53recoveryreadiness.ResourceSet;
import com.pulumi.aws.route53recoveryreadiness.ResourceSetArgs;
import com.pulumi.aws.route53recoveryreadiness.inputs.ResourceSetResourceArgs;
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
        var example = new ResourceSet("example", ResourceSetArgs.builder()        
            .resourceSetName(my_cw_alarm_set)
            .resourceSetType("AWS::CloudWatch::Alarm")
            .resources(ResourceSetResourceArgs.builder()
                .resourceArn(aws_cloudwatch_metric_alarm.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53recoveryreadiness:ResourceSet
    properties:
      resourceSetName: ${["my-cw-alarm-set"]}
      resourceSetType: AWS::CloudWatch::Alarm
      resources:
        - resourceArn: ${aws_cloudwatch_metric_alarm.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route53 Recovery Readiness resource set name using the resource set name. For example:

```sh
 $ pulumi import aws:route53recoveryreadiness/resourceSet:ResourceSet my-cw-alarm-set example
```
 