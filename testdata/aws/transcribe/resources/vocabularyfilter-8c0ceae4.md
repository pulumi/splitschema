Resource for managing an AWS Transcribe VocabularyFilter.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transcribe.VocabularyFilter("example", {
    languageCode: "en-US",
    tags: {
        tag1: "value1",
        tag2: "value3",
    },
    vocabularyFilterName: "example",
    words: [
        "cars",
        "bucket",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transcribe.VocabularyFilter("example",
    language_code="en-US",
    tags={
        "tag1": "value1",
        "tag2": "value3",
    },
    vocabulary_filter_name="example",
    words=[
        "cars",
        "bucket",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transcribe.VocabularyFilter("example", new()
    {
        LanguageCode = "en-US",
        Tags = 
        {
            { "tag1", "value1" },
            { "tag2", "value3" },
        },
        VocabularyFilterName = "example",
        Words = new[]
        {
            "cars",
            "bucket",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transcribe"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transcribe.NewVocabularyFilter(ctx, "example", &transcribe.VocabularyFilterArgs{
			LanguageCode: pulumi.String("en-US"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value3"),
			},
			VocabularyFilterName: pulumi.String("example"),
			Words: pulumi.StringArray{
				pulumi.String("cars"),
				pulumi.String("bucket"),
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
import com.pulumi.aws.transcribe.VocabularyFilter;
import com.pulumi.aws.transcribe.VocabularyFilterArgs;
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
        var example = new VocabularyFilter("example", VocabularyFilterArgs.builder()        
            .languageCode("en-US")
            .tags(Map.ofEntries(
                Map.entry("tag1", "value1"),
                Map.entry("tag2", "value3")
            ))
            .vocabularyFilterName("example")
            .words(            
                "cars",
                "bucket")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transcribe:VocabularyFilter
    properties:
      languageCode: en-US
      tags:
        tag1: value1
        tag2: value3
      vocabularyFilterName: example
      words:
        - cars
        - bucket
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transcribe VocabularyFilter using the `vocabulary_filter_name`. For example:

```sh
 $ pulumi import aws:transcribe/vocabularyFilter:VocabularyFilter example example-name
```
 