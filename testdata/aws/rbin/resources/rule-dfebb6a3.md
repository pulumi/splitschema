Resource for managing an AWS RBin Rule.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rbin.Rule("example", {
    description: "example_rule",
    resourceTags: [{
        resourceTagKey: "tag_key",
        resourceTagValue: "tag_value",
    }],
    resourceType: "EBS_SNAPSHOT",
    retentionPeriod: {
        retentionPeriodUnit: "DAYS",
        retentionPeriodValue: 10,
    },
    tags: {
        test_tag_key: "test_tag_value",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rbin.Rule("example",
    description="example_rule",
    resource_tags=[aws.rbin.RuleResourceTagArgs(
        resource_tag_key="tag_key",
        resource_tag_value="tag_value",
    )],
    resource_type="EBS_SNAPSHOT",
    retention_period=aws.rbin.RuleRetentionPeriodArgs(
        retention_period_unit="DAYS",
        retention_period_value=10,
    ),
    tags={
        "test_tag_key": "test_tag_value",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Rbin.Rule("example", new()
    {
        Description = "example_rule",
        ResourceTags = new[]
        {
            new Aws.Rbin.Inputs.RuleResourceTagArgs
            {
                ResourceTagKey = "tag_key",
                ResourceTagValue = "tag_value",
            },
        },
        ResourceType = "EBS_SNAPSHOT",
        RetentionPeriod = new Aws.Rbin.Inputs.RuleRetentionPeriodArgs
        {
            RetentionPeriodUnit = "DAYS",
            RetentionPeriodValue = 10,
        },
        Tags = 
        {
            { "test_tag_key", "test_tag_value" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rbin"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rbin.NewRule(ctx, "example", &rbin.RuleArgs{
			Description: pulumi.String("example_rule"),
			ResourceTags: rbin.RuleResourceTagArray{
				&rbin.RuleResourceTagArgs{
					ResourceTagKey:   pulumi.String("tag_key"),
					ResourceTagValue: pulumi.String("tag_value"),
				},
			},
			ResourceType: pulumi.String("EBS_SNAPSHOT"),
			RetentionPeriod: &rbin.RuleRetentionPeriodArgs{
				RetentionPeriodUnit:  pulumi.String("DAYS"),
				RetentionPeriodValue: pulumi.Int(10),
			},
			Tags: pulumi.StringMap{
				"test_tag_key": pulumi.String("test_tag_value"),
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
import com.pulumi.aws.rbin.Rule;
import com.pulumi.aws.rbin.RuleArgs;
import com.pulumi.aws.rbin.inputs.RuleResourceTagArgs;
import com.pulumi.aws.rbin.inputs.RuleRetentionPeriodArgs;
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
        var example = new Rule("example", RuleArgs.builder()        
            .description("example_rule")
            .resourceTags(RuleResourceTagArgs.builder()
                .resourceTagKey("tag_key")
                .resourceTagValue("tag_value")
                .build())
            .resourceType("EBS_SNAPSHOT")
            .retentionPeriod(RuleRetentionPeriodArgs.builder()
                .retentionPeriodUnit("DAYS")
                .retentionPeriodValue(10)
                .build())
            .tags(Map.of("test_tag_key", "test_tag_value"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rbin:Rule
    properties:
      description: example_rule
      resourceTags:
        - resourceTagKey: tag_key
          resourceTagValue: tag_value
      resourceType: EBS_SNAPSHOT
      retentionPeriod:
        retentionPeriodUnit: DAYS
        retentionPeriodValue: 10
      tags:
        test_tag_key: test_tag_value
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import RBin Rule using the `id`. For example:

```sh
 $ pulumi import aws:rbin/rule:Rule example examplerule
```
 