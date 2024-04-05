Resource for managing an AWS Lex V2 Models Bot Version.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lex.V2modelsBotVersion("test", {
    botId: aws_lexv2models.test.id,
    localeSpecification: {
        en_US: {
            sourceBotVersion: "DRAFT",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lex.V2modelsBotVersion("test",
    bot_id=aws_lexv2models["test"]["id"],
    locale_specification={
        "en_US": aws.lex.V2modelsBotVersionLocaleSpecificationArgs(
            source_bot_version="DRAFT",
        ),
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Lex.V2modelsBotVersion("test", new()
    {
        BotId = aws_lexv2models.Test.Id,
        LocaleSpecification = 
        {
            { "en_US", new Aws.Lex.Inputs.V2modelsBotVersionLocaleSpecificationArgs
            {
                SourceBotVersion = "DRAFT",
            } },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lex"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lex.NewV2modelsBotVersion(ctx, "test", &lex.V2modelsBotVersionArgs{
			BotId: pulumi.Any(aws_lexv2models.Test.Id),
			LocaleSpecification: lex.V2modelsBotVersionLocaleSpecificationMap{
				"en_US": &lex.V2modelsBotVersionLocaleSpecificationArgs{
					SourceBotVersion: pulumi.String("DRAFT"),
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
import com.pulumi.aws.lex.V2modelsBotVersion;
import com.pulumi.aws.lex.V2modelsBotVersionArgs;
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
        var test = new V2modelsBotVersion("test", V2modelsBotVersionArgs.builder()        
            .botId(aws_lexv2models.test().id())
            .localeSpecification(Map.of("en_US", Map.of("sourceBotVersion", "DRAFT")))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lex:V2modelsBotVersion
    properties:
      botId: ${aws_lexv2models.test.id}
      localeSpecification:
        en_US:
          sourceBotVersion: DRAFT
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lex V2 Models Bot Version using the `id`. For example:

```sh
 $ pulumi import aws:lex/v2modelsBotVersion:V2modelsBotVersion example id-12345678,1
```
 