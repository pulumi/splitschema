Resource for managing an AWS Kendra Experience.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.kendra.Experience("example", {
    indexId: aws_kendra_index.example.id,
    description: "My Kendra Experience",
    roleArn: aws_iam_role.example.arn,
    configuration: {
        contentSourceConfiguration: {
            directPutContent: true,
            faqIds: [aws_kendra_faq.example.faq_id],
        },
        userIdentityConfiguration: {
            identityAttributeName: "12345ec453-1546651e-79c4-4554-91fa-00b43ccfa245",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.kendra.Experience("example",
    index_id=aws_kendra_index["example"]["id"],
    description="My Kendra Experience",
    role_arn=aws_iam_role["example"]["arn"],
    configuration=aws.kendra.ExperienceConfigurationArgs(
        content_source_configuration=aws.kendra.ExperienceConfigurationContentSourceConfigurationArgs(
            direct_put_content=True,
            faq_ids=[aws_kendra_faq["example"]["faq_id"]],
        ),
        user_identity_configuration=aws.kendra.ExperienceConfigurationUserIdentityConfigurationArgs(
            identity_attribute_name="12345ec453-1546651e-79c4-4554-91fa-00b43ccfa245",
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Kendra.Experience("example", new()
    {
        IndexId = aws_kendra_index.Example.Id,
        Description = "My Kendra Experience",
        RoleArn = aws_iam_role.Example.Arn,
        Configuration = new Aws.Kendra.Inputs.ExperienceConfigurationArgs
        {
            ContentSourceConfiguration = new Aws.Kendra.Inputs.ExperienceConfigurationContentSourceConfigurationArgs
            {
                DirectPutContent = true,
                FaqIds = new[]
                {
                    aws_kendra_faq.Example.Faq_id,
                },
            },
            UserIdentityConfiguration = new Aws.Kendra.Inputs.ExperienceConfigurationUserIdentityConfigurationArgs
            {
                IdentityAttributeName = "12345ec453-1546651e-79c4-4554-91fa-00b43ccfa245",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kendra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kendra.NewExperience(ctx, "example", &kendra.ExperienceArgs{
			IndexId:     pulumi.Any(aws_kendra_index.Example.Id),
			Description: pulumi.String("My Kendra Experience"),
			RoleArn:     pulumi.Any(aws_iam_role.Example.Arn),
			Configuration: &kendra.ExperienceConfigurationArgs{
				ContentSourceConfiguration: &kendra.ExperienceConfigurationContentSourceConfigurationArgs{
					DirectPutContent: pulumi.Bool(true),
					FaqIds: pulumi.StringArray{
						aws_kendra_faq.Example.Faq_id,
					},
				},
				UserIdentityConfiguration: &kendra.ExperienceConfigurationUserIdentityConfigurationArgs{
					IdentityAttributeName: pulumi.String("12345ec453-1546651e-79c4-4554-91fa-00b43ccfa245"),
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
import com.pulumi.aws.kendra.Experience;
import com.pulumi.aws.kendra.ExperienceArgs;
import com.pulumi.aws.kendra.inputs.ExperienceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.ExperienceConfigurationContentSourceConfigurationArgs;
import com.pulumi.aws.kendra.inputs.ExperienceConfigurationUserIdentityConfigurationArgs;
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
        var example = new Experience("example", ExperienceArgs.builder()        
            .indexId(aws_kendra_index.example().id())
            .description("My Kendra Experience")
            .roleArn(aws_iam_role.example().arn())
            .configuration(ExperienceConfigurationArgs.builder()
                .contentSourceConfiguration(ExperienceConfigurationContentSourceConfigurationArgs.builder()
                    .directPutContent(true)
                    .faqIds(aws_kendra_faq.example().faq_id())
                    .build())
                .userIdentityConfiguration(ExperienceConfigurationUserIdentityConfigurationArgs.builder()
                    .identityAttributeName("12345ec453-1546651e-79c4-4554-91fa-00b43ccfa245")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:kendra:Experience
    properties:
      indexId: ${aws_kendra_index.example.id}
      description: My Kendra Experience
      roleArn: ${aws_iam_role.example.arn}
      configuration:
        contentSourceConfiguration:
          directPutContent: true
          faqIds:
            - ${aws_kendra_faq.example.faq_id}
        userIdentityConfiguration:
          identityAttributeName: 12345ec453-1546651e-79c4-4554-91fa-00b43ccfa245
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Kendra Experience using the unique identifiers of the experience and index separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:kendra/experience:Experience example 1045d08d-66ef-4882-b3ed-dfb7df183e90/b34dfdf7-1f2b-4704-9581-79e00296845f
```
 