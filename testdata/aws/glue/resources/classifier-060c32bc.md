Provides a Glue Classifier resource.

> **NOTE:** It is only valid to create one type of classifier (csv, grok, JSON, or XML). Changing classifier types will recreate the classifier.

{{% examples %}}
## Example Usage
{{% example %}}
### Csv Classifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Classifier("example", {csvClassifier: {
    allowSingleColumn: false,
    containsHeader: "PRESENT",
    delimiter: ",",
    disableValueTrimming: false,
    headers: [
        "example1",
        "example2",
    ],
    quoteSymbol: "'",
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Classifier("example", csv_classifier=aws.glue.ClassifierCsvClassifierArgs(
    allow_single_column=False,
    contains_header="PRESENT",
    delimiter=",",
    disable_value_trimming=False,
    headers=[
        "example1",
        "example2",
    ],
    quote_symbol="'",
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Classifier("example", new()
    {
        CsvClassifier = new Aws.Glue.Inputs.ClassifierCsvClassifierArgs
        {
            AllowSingleColumn = false,
            ContainsHeader = "PRESENT",
            Delimiter = ",",
            DisableValueTrimming = false,
            Headers = new[]
            {
                "example1",
                "example2",
            },
            QuoteSymbol = "'",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewClassifier(ctx, "example", &glue.ClassifierArgs{
			CsvClassifier: &glue.ClassifierCsvClassifierArgs{
				AllowSingleColumn:    pulumi.Bool(false),
				ContainsHeader:       pulumi.String("PRESENT"),
				Delimiter:            pulumi.String(","),
				DisableValueTrimming: pulumi.Bool(false),
				Headers: pulumi.StringArray{
					pulumi.String("example1"),
					pulumi.String("example2"),
				},
				QuoteSymbol: pulumi.String("'"),
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
import com.pulumi.aws.glue.Classifier;
import com.pulumi.aws.glue.ClassifierArgs;
import com.pulumi.aws.glue.inputs.ClassifierCsvClassifierArgs;
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
        var example = new Classifier("example", ClassifierArgs.builder()        
            .csvClassifier(ClassifierCsvClassifierArgs.builder()
                .allowSingleColumn(false)
                .containsHeader("PRESENT")
                .delimiter(",")
                .disableValueTrimming(false)
                .headers(                
                    "example1",
                    "example2")
                .quoteSymbol("'")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Classifier
    properties:
      csvClassifier:
        allowSingleColumn: false
        containsHeader: PRESENT
        delimiter: ','
        disableValueTrimming: false
        headers:
          - example1
          - example2
        quoteSymbol: ''''
```
{{% /example %}}
{{% example %}}
### Grok Classifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Classifier("example", {grokClassifier: {
    classification: "example",
    grokPattern: "example",
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Classifier("example", grok_classifier=aws.glue.ClassifierGrokClassifierArgs(
    classification="example",
    grok_pattern="example",
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Classifier("example", new()
    {
        GrokClassifier = new Aws.Glue.Inputs.ClassifierGrokClassifierArgs
        {
            Classification = "example",
            GrokPattern = "example",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewClassifier(ctx, "example", &glue.ClassifierArgs{
			GrokClassifier: &glue.ClassifierGrokClassifierArgs{
				Classification: pulumi.String("example"),
				GrokPattern:    pulumi.String("example"),
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
import com.pulumi.aws.glue.Classifier;
import com.pulumi.aws.glue.ClassifierArgs;
import com.pulumi.aws.glue.inputs.ClassifierGrokClassifierArgs;
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
        var example = new Classifier("example", ClassifierArgs.builder()        
            .grokClassifier(ClassifierGrokClassifierArgs.builder()
                .classification("example")
                .grokPattern("example")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Classifier
    properties:
      grokClassifier:
        classification: example
        grokPattern: example
```
{{% /example %}}
{{% example %}}
### JSON Classifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Classifier("example", {jsonClassifier: {
    jsonPath: "example",
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Classifier("example", json_classifier=aws.glue.ClassifierJsonClassifierArgs(
    json_path="example",
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Classifier("example", new()
    {
        JsonClassifier = new Aws.Glue.Inputs.ClassifierJsonClassifierArgs
        {
            JsonPath = "example",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewClassifier(ctx, "example", &glue.ClassifierArgs{
			JsonClassifier: &glue.ClassifierJsonClassifierArgs{
				JsonPath: pulumi.String("example"),
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
import com.pulumi.aws.glue.Classifier;
import com.pulumi.aws.glue.ClassifierArgs;
import com.pulumi.aws.glue.inputs.ClassifierJsonClassifierArgs;
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
        var example = new Classifier("example", ClassifierArgs.builder()        
            .jsonClassifier(ClassifierJsonClassifierArgs.builder()
                .jsonPath("example")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Classifier
    properties:
      jsonClassifier:
        jsonPath: example
```
{{% /example %}}
{{% example %}}
### XML Classifier

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Classifier("example", {xmlClassifier: {
    classification: "example",
    rowTag: "example",
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Classifier("example", xml_classifier=aws.glue.ClassifierXmlClassifierArgs(
    classification="example",
    row_tag="example",
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Classifier("example", new()
    {
        XmlClassifier = new Aws.Glue.Inputs.ClassifierXmlClassifierArgs
        {
            Classification = "example",
            RowTag = "example",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewClassifier(ctx, "example", &glue.ClassifierArgs{
			XmlClassifier: &glue.ClassifierXmlClassifierArgs{
				Classification: pulumi.String("example"),
				RowTag:         pulumi.String("example"),
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
import com.pulumi.aws.glue.Classifier;
import com.pulumi.aws.glue.ClassifierArgs;
import com.pulumi.aws.glue.inputs.ClassifierXmlClassifierArgs;
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
        var example = new Classifier("example", ClassifierArgs.builder()        
            .xmlClassifier(ClassifierXmlClassifierArgs.builder()
                .classification("example")
                .rowTag("example")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Classifier
    properties:
      xmlClassifier:
        classification: example
        rowTag: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Classifiers using their name. For example:

```sh
 $ pulumi import aws:glue/classifier:Classifier MyClassifier MyClassifier
```
 