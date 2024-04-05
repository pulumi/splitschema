Resource for managing an AWS Lex V2 Models Bot Locale.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lex.V2modelsBotLocale("example", {
    botId: aws_lexv2models_bot.example.id,
    botVersion: "DRAFT",
    localeId: "en_US",
    nLuIntentConfidenceThreshold: 0.7,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lex.V2modelsBotLocale("example",
    bot_id=aws_lexv2models_bot["example"]["id"],
    bot_version="DRAFT",
    locale_id="en_US",
    n_lu_intent_confidence_threshold=0.7)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lex.V2modelsBotLocale("example", new()
    {
        BotId = aws_lexv2models_bot.Example.Id,
        BotVersion = "DRAFT",
        LocaleId = "en_US",
        NLuIntentConfidenceThreshold = 0.7,
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
		_, err := lex.NewV2modelsBotLocale(ctx, "example", &lex.V2modelsBotLocaleArgs{
			BotId:                        pulumi.Any(aws_lexv2models_bot.Example.Id),
			BotVersion:                   pulumi.String("DRAFT"),
			LocaleId:                     pulumi.String("en_US"),
			NLuIntentConfidenceThreshold: pulumi.Float64(0.7),
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
import com.pulumi.aws.lex.V2modelsBotLocale;
import com.pulumi.aws.lex.V2modelsBotLocaleArgs;
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
        var example = new V2modelsBotLocale("example", V2modelsBotLocaleArgs.builder()        
            .botId(aws_lexv2models_bot.example().id())
            .botVersion("DRAFT")
            .localeId("en_US")
            .nLuIntentConfidenceThreshold(0.7)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lex:V2modelsBotLocale
    properties:
      botId: ${aws_lexv2models_bot.example.id}
      botVersion: DRAFT
      localeId: en_US
      nLuIntentConfidenceThreshold: 0.7
```
{{% /example %}}
{{% example %}}
### Voice Settings

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lex.V2modelsBotLocale("example", {
    botId: aws_lexv2models_bot.example.id,
    botVersion: "DRAFT",
    localeId: "en_US",
    nLuIntentConfidenceThreshold: 0.7,
    voiceSettings: {
        voiceId: "Kendra",
        engine: "standard",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lex.V2modelsBotLocale("example",
    bot_id=aws_lexv2models_bot["example"]["id"],
    bot_version="DRAFT",
    locale_id="en_US",
    n_lu_intent_confidence_threshold=0.7,
    voice_settings=aws.lex.V2modelsBotLocaleVoiceSettingsArgs(
        voice_id="Kendra",
        engine="standard",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lex.V2modelsBotLocale("example", new()
    {
        BotId = aws_lexv2models_bot.Example.Id,
        BotVersion = "DRAFT",
        LocaleId = "en_US",
        NLuIntentConfidenceThreshold = 0.7,
        VoiceSettings = new Aws.Lex.Inputs.V2modelsBotLocaleVoiceSettingsArgs
        {
            VoiceId = "Kendra",
            Engine = "standard",
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
		_, err := lex.NewV2modelsBotLocale(ctx, "example", &lex.V2modelsBotLocaleArgs{
			BotId:                        pulumi.Any(aws_lexv2models_bot.Example.Id),
			BotVersion:                   pulumi.String("DRAFT"),
			LocaleId:                     pulumi.String("en_US"),
			NLuIntentConfidenceThreshold: pulumi.Float64(0.7),
			VoiceSettings: &lex.V2modelsBotLocaleVoiceSettingsArgs{
				VoiceId: pulumi.String("Kendra"),
				Engine:  pulumi.String("standard"),
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
import com.pulumi.aws.lex.V2modelsBotLocale;
import com.pulumi.aws.lex.V2modelsBotLocaleArgs;
import com.pulumi.aws.lex.inputs.V2modelsBotLocaleVoiceSettingsArgs;
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
        var example = new V2modelsBotLocale("example", V2modelsBotLocaleArgs.builder()        
            .botId(aws_lexv2models_bot.example().id())
            .botVersion("DRAFT")
            .localeId("en_US")
            .nLuIntentConfidenceThreshold(0.7)
            .voiceSettings(V2modelsBotLocaleVoiceSettingsArgs.builder()
                .voiceId("Kendra")
                .engine("standard")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lex:V2modelsBotLocale
    properties:
      botId: ${aws_lexv2models_bot.example.id}
      botVersion: DRAFT
      localeId: en_US
      nLuIntentConfidenceThreshold: 0.7
      voiceSettings:
        voiceId: Kendra
        engine: standard
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lex V2 Models Bot Locale using the `id`. For example:

```sh
 $ pulumi import aws:lex/v2modelsBotLocale:V2modelsBotLocale example en_US,abcd-12345678,1
```
 