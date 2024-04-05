Gets the contents of the specified Systems Manager document.

{{% examples %}}
## Example Usage
{{% example %}}

To get the contents of the document owned by AWS.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = aws.ssm.getDocument({
    name: "AWS-GatherSoftwareInventory",
    documentFormat: "YAML",
});
export const content = foo.then(foo => foo.content);
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ssm.get_document(name="AWS-GatherSoftwareInventory",
    document_format="YAML")
pulumi.export("content", foo.content)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = Aws.Ssm.GetDocument.Invoke(new()
    {
        Name = "AWS-GatherSoftwareInventory",
        DocumentFormat = "YAML",
    });

    return new Dictionary<string, object?>
    {
        ["content"] = foo.Apply(getDocumentResult => getDocumentResult.Content),
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foo, err := ssm.LookupDocument(ctx, &ssm.LookupDocumentArgs{
			Name:           "AWS-GatherSoftwareInventory",
			DocumentFormat: pulumi.StringRef("YAML"),
		}, nil)
		if err != nil {
			return err
		}
		ctx.Export("content", foo.Content)
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ssm.SsmFunctions;
import com.pulumi.aws.ssm.inputs.GetDocumentArgs;
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
        final var foo = SsmFunctions.getDocument(GetDocumentArgs.builder()
            .name("AWS-GatherSoftwareInventory")
            .documentFormat("YAML")
            .build());

        ctx.export("content", foo.applyValue(getDocumentResult -> getDocumentResult.content()));
    }
}
```
```yaml
variables:
  foo:
    fn::invoke:
      Function: aws:ssm:getDocument
      Arguments:
        name: AWS-GatherSoftwareInventory
        documentFormat: YAML
outputs:
  content: ${foo.content}
```

To get the contents of the custom document.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.ssm.getDocument({
    name: aws_ssm_document.test.name,
    documentFormat: "JSON",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ssm.get_document(name=aws_ssm_document["test"]["name"],
    document_format="JSON")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Ssm.GetDocument.Invoke(new()
    {
        Name = aws_ssm_document.Test.Name,
        DocumentFormat = "JSON",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssm.LookupDocument(ctx, &ssm.LookupDocumentArgs{
			Name:           aws_ssm_document.Test.Name,
			DocumentFormat: pulumi.StringRef("JSON"),
		}, nil)
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
import com.pulumi.aws.ssm.SsmFunctions;
import com.pulumi.aws.ssm.inputs.GetDocumentArgs;
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
        final var test = SsmFunctions.getDocument(GetDocumentArgs.builder()
            .name(aws_ssm_document.test().name())
            .documentFormat("JSON")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:ssm:getDocument
      Arguments:
        name: ${aws_ssm_document.test.name}
        documentFormat: JSON
```
{{% /example %}}
{{% /examples %}}