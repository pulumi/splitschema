Provides a resource to manage an [Amazon Macie Findings Filter](https://docs.aws.amazon.com/macie/latest/APIReference/findingsfilters-id.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.macie2.Account("example", {});
const test = new aws.macie.FindingsFilter("test", {
    description: "DESCRIPTION",
    position: 1,
    action: "ARCHIVE",
    findingCriteria: {
        criterions: [{
            field: "region",
            eqs: [data.aws_region.current.name],
        }],
    },
}, {
    dependsOn: [aws_macie2_account.test],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.macie2.Account("example")
test = aws.macie.FindingsFilter("test",
    description="DESCRIPTION",
    position=1,
    action="ARCHIVE",
    finding_criteria=aws.macie.FindingsFilterFindingCriteriaArgs(
        criterions=[aws.macie.FindingsFilterFindingCriteriaCriterionArgs(
            field="region",
            eqs=[data["aws_region"]["current"]["name"]],
        )],
    ),
    opts=pulumi.ResourceOptions(depends_on=[aws_macie2_account["test"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Macie2.Account("example");

    var test = new Aws.Macie.FindingsFilter("test", new()
    {
        Description = "DESCRIPTION",
        Position = 1,
        Action = "ARCHIVE",
        FindingCriteria = new Aws.Macie.Inputs.FindingsFilterFindingCriteriaArgs
        {
            Criterions = new[]
            {
                new Aws.Macie.Inputs.FindingsFilterFindingCriteriaCriterionArgs
                {
                    Field = "region",
                    Eqs = new[]
                    {
                        data.Aws_region.Current.Name,
                    },
                },
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_macie2_account.Test,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/macie"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/macie2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := macie2.NewAccount(ctx, "example", nil)
		if err != nil {
			return err
		}
		_, err = macie.NewFindingsFilter(ctx, "test", &macie.FindingsFilterArgs{
			Description: pulumi.String("DESCRIPTION"),
			Position:    pulumi.Int(1),
			Action:      pulumi.String("ARCHIVE"),
			FindingCriteria: &macie.FindingsFilterFindingCriteriaArgs{
				Criterions: macie.FindingsFilterFindingCriteriaCriterionArray{
					&macie.FindingsFilterFindingCriteriaCriterionArgs{
						Field: pulumi.String("region"),
						Eqs: pulumi.StringArray{
							data.Aws_region.Current.Name,
						},
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_macie2_account.Test,
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
import com.pulumi.aws.macie2.Account;
import com.pulumi.aws.macie.FindingsFilter;
import com.pulumi.aws.macie.FindingsFilterArgs;
import com.pulumi.aws.macie.inputs.FindingsFilterFindingCriteriaArgs;
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
        var example = new Account("example");

        var test = new FindingsFilter("test", FindingsFilterArgs.builder()        
            .description("DESCRIPTION")
            .position(1)
            .action("ARCHIVE")
            .findingCriteria(FindingsFilterFindingCriteriaArgs.builder()
                .criterions(FindingsFilterFindingCriteriaCriterionArgs.builder()
                    .field("region")
                    .eqs(data.aws_region().current().name())
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_macie2_account.test())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:macie2:Account
  test:
    type: aws:macie:FindingsFilter
    properties:
      description: DESCRIPTION
      position: 1
      action: ARCHIVE
      findingCriteria:
        criterions:
          - field: region
            eqs:
              - ${data.aws_region.current.name}
    options:
      dependson:
        - ${aws_macie2_account.test}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_macie2_findings_filter` using the id. For example:

```sh
 $ pulumi import aws:macie/findingsFilter:FindingsFilter example abcd1
```
 