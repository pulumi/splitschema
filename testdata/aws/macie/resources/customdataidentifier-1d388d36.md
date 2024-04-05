Provides a resource to manage an [AWS Macie Custom Data Identifier](https://docs.aws.amazon.com/macie/latest/APIReference/custom-data-identifiers-id.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.macie2.Account("exampleAccount", {});
const exampleCustomDataIdentifier = new aws.macie.CustomDataIdentifier("exampleCustomDataIdentifier", {
    regex: "[0-9]{3}-[0-9]{2}-[0-9]{4}",
    description: "DESCRIPTION",
    maximumMatchDistance: 10,
    keywords: ["keyword"],
    ignoreWords: ["ignore"],
}, {
    dependsOn: [aws_macie2_account.test],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.macie2.Account("exampleAccount")
example_custom_data_identifier = aws.macie.CustomDataIdentifier("exampleCustomDataIdentifier",
    regex="[0-9]{3}-[0-9]{2}-[0-9]{4}",
    description="DESCRIPTION",
    maximum_match_distance=10,
    keywords=["keyword"],
    ignore_words=["ignore"],
    opts=pulumi.ResourceOptions(depends_on=[aws_macie2_account["test"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.Macie2.Account("exampleAccount");

    var exampleCustomDataIdentifier = new Aws.Macie.CustomDataIdentifier("exampleCustomDataIdentifier", new()
    {
        Regex = "[0-9]{3}-[0-9]{2}-[0-9]{4}",
        Description = "DESCRIPTION",
        MaximumMatchDistance = 10,
        Keywords = new[]
        {
            "keyword",
        },
        IgnoreWords = new[]
        {
            "ignore",
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
		_, err := macie2.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = macie.NewCustomDataIdentifier(ctx, "exampleCustomDataIdentifier", &macie.CustomDataIdentifierArgs{
			Regex:                pulumi.String("[0-9]{3}-[0-9]{2}-[0-9]{4}"),
			Description:          pulumi.String("DESCRIPTION"),
			MaximumMatchDistance: pulumi.Int(10),
			Keywords: pulumi.StringArray{
				pulumi.String("keyword"),
			},
			IgnoreWords: pulumi.StringArray{
				pulumi.String("ignore"),
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
import com.pulumi.aws.macie.CustomDataIdentifier;
import com.pulumi.aws.macie.CustomDataIdentifierArgs;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleCustomDataIdentifier = new CustomDataIdentifier("exampleCustomDataIdentifier", CustomDataIdentifierArgs.builder()        
            .regex("[0-9]{3}-[0-9]{2}-[0-9]{4}")
            .description("DESCRIPTION")
            .maximumMatchDistance(10)
            .keywords("keyword")
            .ignoreWords("ignore")
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_macie2_account.test())
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:macie2:Account
  exampleCustomDataIdentifier:
    type: aws:macie:CustomDataIdentifier
    properties:
      regex: '[0-9]{3}-[0-9]{2}-[0-9]{4}'
      description: DESCRIPTION
      maximumMatchDistance: 10
      keywords:
        - keyword
      ignoreWords:
        - ignore
    options:
      dependson:
        - ${aws_macie2_account.test}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_macie2_custom_data_identifier` using the id. For example:

```sh
 $ pulumi import aws:macie/customDataIdentifier:CustomDataIdentifier example abcd1
```
 