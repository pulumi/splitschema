Provides an Amazon Connect Vocabulary resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.Vocabulary("example", {
    content: `Phrase	IPA	SoundsLike	DisplayAs
Los-Angeles			Los Angeles
F.B.I.	ɛ f b i aɪ		FBI
Etienne		eh-tee-en	
`,
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    languageCode: "en-US",
    tags: {
        Key1: "Value1",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.Vocabulary("example",
    content="""Phrase	IPA	SoundsLike	DisplayAs
Los-Angeles			Los Angeles
F.B.I.	ɛ f b i aɪ		FBI
Etienne		eh-tee-en	
""",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    language_code="en-US",
    tags={
        "Key1": "Value1",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.Vocabulary("example", new()
    {
        Content = @"Phrase	IPA	SoundsLike	DisplayAs
Los-Angeles			Los Angeles
F.B.I.	ɛ f b i aɪ		FBI
Etienne		eh-tee-en	
",
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        LanguageCode = "en-US",
        Tags = 
        {
            { "Key1", "Value1" },
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
		_, err := connect.NewVocabulary(ctx, "example", &connect.VocabularyArgs{
			Content:      pulumi.String("Phrase	IPA	SoundsLike	DisplayAs\nLos-Angeles			Los Angeles\nF.B.I.	ɛ f b i aɪ		FBI\nEtienne		eh-tee-en	\n"),
			InstanceId:   pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
			LanguageCode: pulumi.String("en-US"),
			Tags: pulumi.StringMap{
				"Key1": pulumi.String("Value1"),
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
import com.pulumi.aws.connect.Vocabulary;
import com.pulumi.aws.connect.VocabularyArgs;
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
        var example = new Vocabulary("example", VocabularyArgs.builder()        
            .content("""
Phrase	IPA	SoundsLike	DisplayAs
Los-Angeles			Los Angeles
F.B.I.	ɛ f b i aɪ		FBI
Etienne		eh-tee-en	
            """)
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .languageCode("en-US")
            .tags(Map.of("Key1", "Value1"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:Vocabulary
    properties:
      content: |
        Phrase	IPA	SoundsLike	DisplayAs
        Los-Angeles			Los Angeles
        F.B.I.	ɛ f b i aɪ		FBI
        Etienne		eh-tee-en	
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
      languageCode: en-US
      tags:
        Key1: Value1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Connect Vocabularies using the `instance_id` and `vocabulary_id` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:connect/vocabulary:Vocabulary example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
```
 