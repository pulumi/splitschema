Resource for managing an AWS Lex V2 Models Bot.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lex.V2modelsBot("example", {
    dataPrivacies: [{
        childDirected: "boolean",
    }],
    idleSessionTtlInSeconds: 10,
    roleArn: "bot_example_arn",
    tags: {
        foo: "bar",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lex.V2modelsBot("example",
    data_privacies=[aws.lex.V2modelsBotDataPrivacyArgs(
        child_directed="boolean",
    )],
    idle_session_ttl_in_seconds=10,
    role_arn="bot_example_arn",
    tags={
        "foo": "bar",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lex.V2modelsBot("example", new()
    {
        DataPrivacies = new[]
        {
            new Aws.Lex.Inputs.V2modelsBotDataPrivacyArgs
            {
                ChildDirected = "boolean",
            },
        },
        IdleSessionTtlInSeconds = 10,
        RoleArn = "bot_example_arn",
        Tags = 
        {
            { "foo", "bar" },
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
		_, err := lex.NewV2modelsBot(ctx, "example", &lex.V2modelsBotArgs{
			DataPrivacies: lex.V2modelsBotDataPrivacyArray{
				&lex.V2modelsBotDataPrivacyArgs{
					ChildDirected: pulumi.Bool("boolean"),
				},
			},
			IdleSessionTtlInSeconds: pulumi.Int(10),
			RoleArn:                 pulumi.String("bot_example_arn"),
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
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
import com.pulumi.aws.lex.V2modelsBot;
import com.pulumi.aws.lex.V2modelsBotArgs;
import com.pulumi.aws.lex.inputs.V2modelsBotDataPrivacyArgs;
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
        var example = new V2modelsBot("example", V2modelsBotArgs.builder()        
            .dataPrivacies(V2modelsBotDataPrivacyArgs.builder()
                .childDirected("boolean")
                .build())
            .idleSessionTtlInSeconds(10)
            .roleArn("bot_example_arn")
            .tags(Map.of("foo", "bar"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lex:V2modelsBot
    properties:
      dataPrivacies:
        - childDirected: boolean
      idleSessionTtlInSeconds: 10
      roleArn: bot_example_arn
      tags:
        foo: bar
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lex V2 Models Bot using the `id`. For example:

```sh
 $ pulumi import aws:lex/v2modelsBot:V2modelsBot example bot-id-12345678
```
 