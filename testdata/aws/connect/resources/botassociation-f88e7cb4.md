Allows the specified Amazon Connect instance to access the specified Amazon Lex (V1) bot. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html) and [Add an Amazon Lex bot](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-lex.html).

> **NOTE:** This resource only currently supports Amazon Lex (V1) Associations.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.BotAssociation("example", {
    instanceId: aws_connect_instance.example.id,
    lexBot: {
        lexRegion: "us-west-2",
        name: "Test",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.BotAssociation("example",
    instance_id=aws_connect_instance["example"]["id"],
    lex_bot=aws.connect.BotAssociationLexBotArgs(
        lex_region="us-west-2",
        name="Test",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.BotAssociation("example", new()
    {
        InstanceId = aws_connect_instance.Example.Id,
        LexBot = new Aws.Connect.Inputs.BotAssociationLexBotArgs
        {
            LexRegion = "us-west-2",
            Name = "Test",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewBotAssociation(ctx, "example", &connect.BotAssociationArgs{
			InstanceId: pulumi.Any(aws_connect_instance.Example.Id),
			LexBot: &connect.BotAssociationLexBotArgs{
				LexRegion: pulumi.String("us-west-2"),
				Name:      pulumi.String("Test"),
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
import com.pulumi.aws.connect.BotAssociation;
import com.pulumi.aws.connect.BotAssociationArgs;
import com.pulumi.aws.connect.inputs.BotAssociationLexBotArgs;
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
        var example = new BotAssociation("example", BotAssociationArgs.builder()        
            .instanceId(aws_connect_instance.example().id())
            .lexBot(BotAssociationLexBotArgs.builder()
                .lexRegion("us-west-2")
                .name("Test")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:BotAssociation
    properties:
      instanceId: ${aws_connect_instance.example.id}
      lexBot:
        lexRegion: us-west-2
        name: Test
```
{{% /example %}}
{{% example %}}
### Including a sample Lex bot

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getRegion({});
const exampleIntent = new aws.lex.Intent("exampleIntent", {
    createVersion: true,
    name: "connect_lex_intent",
    fulfillmentActivity: {
        type: "ReturnIntent",
    },
    sampleUtterances: ["I would like to pick up flowers."],
});
const exampleBot = new aws.lex.Bot("exampleBot", {
    abortStatement: {
        messages: [{
            content: "Sorry, I am not able to assist at this time.",
            contentType: "PlainText",
        }],
    },
    clarificationPrompt: {
        maxAttempts: 2,
        messages: [{
            content: "I didn't understand you, what would you like to do?",
            contentType: "PlainText",
        }],
    },
    intents: [{
        intentName: exampleIntent.name,
        intentVersion: "1",
    }],
    childDirected: false,
    name: "connect_lex_bot",
    processBehavior: "BUILD",
});
const exampleBotAssociation = new aws.connect.BotAssociation("exampleBotAssociation", {
    instanceId: aws_connect_instance.example.id,
    lexBot: {
        lexRegion: current.then(current => current.name),
        name: exampleBot.name,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_region()
example_intent = aws.lex.Intent("exampleIntent",
    create_version=True,
    name="connect_lex_intent",
    fulfillment_activity=aws.lex.IntentFulfillmentActivityArgs(
        type="ReturnIntent",
    ),
    sample_utterances=["I would like to pick up flowers."])
example_bot = aws.lex.Bot("exampleBot",
    abort_statement=aws.lex.BotAbortStatementArgs(
        messages=[aws.lex.BotAbortStatementMessageArgs(
            content="Sorry, I am not able to assist at this time.",
            content_type="PlainText",
        )],
    ),
    clarification_prompt=aws.lex.BotClarificationPromptArgs(
        max_attempts=2,
        messages=[aws.lex.BotClarificationPromptMessageArgs(
            content="I didn't understand you, what would you like to do?",
            content_type="PlainText",
        )],
    ),
    intents=[aws.lex.BotIntentArgs(
        intent_name=example_intent.name,
        intent_version="1",
    )],
    child_directed=False,
    name="connect_lex_bot",
    process_behavior="BUILD")
example_bot_association = aws.connect.BotAssociation("exampleBotAssociation",
    instance_id=aws_connect_instance["example"]["id"],
    lex_bot=aws.connect.BotAssociationLexBotArgs(
        lex_region=current.name,
        name=example_bot.name,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetRegion.Invoke();

    var exampleIntent = new Aws.Lex.Intent("exampleIntent", new()
    {
        CreateVersion = true,
        Name = "connect_lex_intent",
        FulfillmentActivity = new Aws.Lex.Inputs.IntentFulfillmentActivityArgs
        {
            Type = "ReturnIntent",
        },
        SampleUtterances = new[]
        {
            "I would like to pick up flowers.",
        },
    });

    var exampleBot = new Aws.Lex.Bot("exampleBot", new()
    {
        AbortStatement = new Aws.Lex.Inputs.BotAbortStatementArgs
        {
            Messages = new[]
            {
                new Aws.Lex.Inputs.BotAbortStatementMessageArgs
                {
                    Content = "Sorry, I am not able to assist at this time.",
                    ContentType = "PlainText",
                },
            },
        },
        ClarificationPrompt = new Aws.Lex.Inputs.BotClarificationPromptArgs
        {
            MaxAttempts = 2,
            Messages = new[]
            {
                new Aws.Lex.Inputs.BotClarificationPromptMessageArgs
                {
                    Content = "I didn't understand you, what would you like to do?",
                    ContentType = "PlainText",
                },
            },
        },
        Intents = new[]
        {
            new Aws.Lex.Inputs.BotIntentArgs
            {
                IntentName = exampleIntent.Name,
                IntentVersion = "1",
            },
        },
        ChildDirected = false,
        Name = "connect_lex_bot",
        ProcessBehavior = "BUILD",
    });

    var exampleBotAssociation = new Aws.Connect.BotAssociation("exampleBotAssociation", new()
    {
        InstanceId = aws_connect_instance.Example.Id,
        LexBot = new Aws.Connect.Inputs.BotAssociationLexBotArgs
        {
            LexRegion = current.Apply(getRegionResult => getRegionResult.Name),
            Name = exampleBot.Name,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lex"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleIntent, err := lex.NewIntent(ctx, "exampleIntent", &lex.IntentArgs{
			CreateVersion: pulumi.Bool(true),
			Name:          pulumi.String("connect_lex_intent"),
			FulfillmentActivity: &lex.IntentFulfillmentActivityArgs{
				Type: pulumi.String("ReturnIntent"),
			},
			SampleUtterances: pulumi.StringArray{
				pulumi.String("I would like to pick up flowers."),
			},
		})
		if err != nil {
			return err
		}
		exampleBot, err := lex.NewBot(ctx, "exampleBot", &lex.BotArgs{
			AbortStatement: &lex.BotAbortStatementArgs{
				Messages: lex.BotAbortStatementMessageArray{
					&lex.BotAbortStatementMessageArgs{
						Content:     pulumi.String("Sorry, I am not able to assist at this time."),
						ContentType: pulumi.String("PlainText"),
					},
				},
			},
			ClarificationPrompt: &lex.BotClarificationPromptArgs{
				MaxAttempts: pulumi.Int(2),
				Messages: lex.BotClarificationPromptMessageArray{
					&lex.BotClarificationPromptMessageArgs{
						Content:     pulumi.String("I didn't understand you, what would you like to do?"),
						ContentType: pulumi.String("PlainText"),
					},
				},
			},
			Intents: lex.BotIntentArray{
				&lex.BotIntentArgs{
					IntentName:    exampleIntent.Name,
					IntentVersion: pulumi.String("1"),
				},
			},
			ChildDirected:   pulumi.Bool(false),
			Name:            pulumi.String("connect_lex_bot"),
			ProcessBehavior: pulumi.String("BUILD"),
		})
		if err != nil {
			return err
		}
		_, err = connect.NewBotAssociation(ctx, "exampleBotAssociation", &connect.BotAssociationArgs{
			InstanceId: pulumi.Any(aws_connect_instance.Example.Id),
			LexBot: &connect.BotAssociationLexBotArgs{
				LexRegion: *pulumi.String(current.Name),
				Name:      exampleBot.Name,
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.lex.Intent;
import com.pulumi.aws.lex.IntentArgs;
import com.pulumi.aws.lex.inputs.IntentFulfillmentActivityArgs;
import com.pulumi.aws.lex.Bot;
import com.pulumi.aws.lex.BotArgs;
import com.pulumi.aws.lex.inputs.BotAbortStatementArgs;
import com.pulumi.aws.lex.inputs.BotClarificationPromptArgs;
import com.pulumi.aws.lex.inputs.BotIntentArgs;
import com.pulumi.aws.connect.BotAssociation;
import com.pulumi.aws.connect.BotAssociationArgs;
import com.pulumi.aws.connect.inputs.BotAssociationLexBotArgs;
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
        final var current = AwsFunctions.getRegion();

        var exampleIntent = new Intent("exampleIntent", IntentArgs.builder()        
            .createVersion(true)
            .name("connect_lex_intent")
            .fulfillmentActivity(IntentFulfillmentActivityArgs.builder()
                .type("ReturnIntent")
                .build())
            .sampleUtterances("I would like to pick up flowers.")
            .build());

        var exampleBot = new Bot("exampleBot", BotArgs.builder()        
            .abortStatement(BotAbortStatementArgs.builder()
                .messages(BotAbortStatementMessageArgs.builder()
                    .content("Sorry, I am not able to assist at this time.")
                    .contentType("PlainText")
                    .build())
                .build())
            .clarificationPrompt(BotClarificationPromptArgs.builder()
                .maxAttempts(2)
                .messages(BotClarificationPromptMessageArgs.builder()
                    .content("I didn't understand you, what would you like to do?")
                    .contentType("PlainText")
                    .build())
                .build())
            .intents(BotIntentArgs.builder()
                .intentName(exampleIntent.name())
                .intentVersion("1")
                .build())
            .childDirected(false)
            .name("connect_lex_bot")
            .processBehavior("BUILD")
            .build());

        var exampleBotAssociation = new BotAssociation("exampleBotAssociation", BotAssociationArgs.builder()        
            .instanceId(aws_connect_instance.example().id())
            .lexBot(BotAssociationLexBotArgs.builder()
                .lexRegion(current.applyValue(getRegionResult -> getRegionResult.name()))
                .name(exampleBot.name())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleIntent:
    type: aws:lex:Intent
    properties:
      createVersion: true
      name: connect_lex_intent
      fulfillmentActivity:
        type: ReturnIntent
      sampleUtterances:
        - I would like to pick up flowers.
  exampleBot:
    type: aws:lex:Bot
    properties:
      abortStatement:
        messages:
          - content: Sorry, I am not able to assist at this time.
            contentType: PlainText
      clarificationPrompt:
        maxAttempts: 2
        messages:
          - content: I didn't understand you, what would you like to do?
            contentType: PlainText
      intents:
        - intentName: ${exampleIntent.name}
          intentVersion: '1'
      childDirected: false
      name: connect_lex_bot
      processBehavior: BUILD
  exampleBotAssociation:
    type: aws:connect:BotAssociation
    properties:
      instanceId: ${aws_connect_instance.example.id}
      lexBot:
        lexRegion: ${current.name}
        name: ${exampleBot.name}
variables:
  current:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_connect_bot_association` using the Amazon Connect instance ID, Lex (V1) bot name, and Lex (V1) bot region separated by colons (`:`). For example:

```sh
 $ pulumi import aws:connect/botAssociation:BotAssociation example aaaaaaaa-bbbb-cccc-dddd-111111111111:Example:us-west-2
```
 