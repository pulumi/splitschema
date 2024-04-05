Resource for managing an AWS AccessAnalyzer Archive Rule.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.accessanalyzer.ArchiveRule("example", {
    analyzerName: "example-analyzer",
    filters: [
        {
            criteria: "condition.aws:UserId",
            eqs: ["userid"],
        },
        {
            criteria: "error",
            exists: "true",
        },
        {
            criteria: "isPublic",
            eqs: ["false"],
        },
    ],
    ruleName: "example-rule",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.accessanalyzer.ArchiveRule("example",
    analyzer_name="example-analyzer",
    filters=[
        aws.accessanalyzer.ArchiveRuleFilterArgs(
            criteria="condition.aws:UserId",
            eqs=["userid"],
        ),
        aws.accessanalyzer.ArchiveRuleFilterArgs(
            criteria="error",
            exists="true",
        ),
        aws.accessanalyzer.ArchiveRuleFilterArgs(
            criteria="isPublic",
            eqs=["false"],
        ),
    ],
    rule_name="example-rule")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AccessAnalyzer.ArchiveRule("example", new()
    {
        AnalyzerName = "example-analyzer",
        Filters = new[]
        {
            new Aws.AccessAnalyzer.Inputs.ArchiveRuleFilterArgs
            {
                Criteria = "condition.aws:UserId",
                Eqs = new[]
                {
                    "userid",
                },
            },
            new Aws.AccessAnalyzer.Inputs.ArchiveRuleFilterArgs
            {
                Criteria = "error",
                Exists = "true",
            },
            new Aws.AccessAnalyzer.Inputs.ArchiveRuleFilterArgs
            {
                Criteria = "isPublic",
                Eqs = new[]
                {
                    "false",
                },
            },
        },
        RuleName = "example-rule",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/accessanalyzer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := accessanalyzer.NewArchiveRule(ctx, "example", &accessanalyzer.ArchiveRuleArgs{
			AnalyzerName: pulumi.String("example-analyzer"),
			Filters: accessanalyzer.ArchiveRuleFilterArray{
				&accessanalyzer.ArchiveRuleFilterArgs{
					Criteria: pulumi.String("condition.aws:UserId"),
					Eqs: pulumi.StringArray{
						pulumi.String("userid"),
					},
				},
				&accessanalyzer.ArchiveRuleFilterArgs{
					Criteria: pulumi.String("error"),
					Exists:   pulumi.String("true"),
				},
				&accessanalyzer.ArchiveRuleFilterArgs{
					Criteria: pulumi.String("isPublic"),
					Eqs: pulumi.StringArray{
						pulumi.String("false"),
					},
				},
			},
			RuleName: pulumi.String("example-rule"),
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
import com.pulumi.aws.accessanalyzer.ArchiveRule;
import com.pulumi.aws.accessanalyzer.ArchiveRuleArgs;
import com.pulumi.aws.accessanalyzer.inputs.ArchiveRuleFilterArgs;
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
        var example = new ArchiveRule("example", ArchiveRuleArgs.builder()        
            .analyzerName("example-analyzer")
            .filters(            
                ArchiveRuleFilterArgs.builder()
                    .criteria("condition.aws:UserId")
                    .eqs("userid")
                    .build(),
                ArchiveRuleFilterArgs.builder()
                    .criteria("error")
                    .exists(true)
                    .build(),
                ArchiveRuleFilterArgs.builder()
                    .criteria("isPublic")
                    .eqs("false")
                    .build())
            .ruleName("example-rule")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:accessanalyzer:ArchiveRule
    properties:
      analyzerName: example-analyzer
      filters:
        - criteria: condition.aws:UserId
          eqs:
            - userid
        - criteria: error
          exists: true
        - criteria: isPublic
          eqs:
            - 'false'
      ruleName: example-rule
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AccessAnalyzer ArchiveRule using the `analyzer_name/rule_name`. For example:

```sh
 $ pulumi import aws:accessanalyzer/archiveRule:ArchiveRule example example-analyzer/example-rule
```
 