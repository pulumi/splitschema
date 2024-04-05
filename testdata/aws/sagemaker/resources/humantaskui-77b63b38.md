Provides a SageMaker Human Task UI resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const example = new aws.sagemaker.HumanTaskUI("example", {
    humanTaskUiName: "example",
    uiTemplate: {
        content: fs.readFileSync("sagemaker-human-task-ui-template.html", "utf8"),
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.HumanTaskUI("example",
    human_task_ui_name="example",
    ui_template=aws.sagemaker.HumanTaskUIUiTemplateArgs(
        content=(lambda path: open(path).read())("sagemaker-human-task-ui-template.html"),
    ))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.HumanTaskUI("example", new()
    {
        HumanTaskUiName = "example",
        UiTemplate = new Aws.Sagemaker.Inputs.HumanTaskUIUiTemplateArgs
        {
            Content = File.ReadAllText("sagemaker-human-task-ui-template.html"),
        },
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewHumanTaskUI(ctx, "example", &sagemaker.HumanTaskUIArgs{
			HumanTaskUiName: pulumi.String("example"),
			UiTemplate: &sagemaker.HumanTaskUIUiTemplateArgs{
				Content: readFileOrPanic("sagemaker-human-task-ui-template.html"),
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
import com.pulumi.aws.sagemaker.HumanTaskUI;
import com.pulumi.aws.sagemaker.HumanTaskUIArgs;
import com.pulumi.aws.sagemaker.inputs.HumanTaskUIUiTemplateArgs;
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
        var example = new HumanTaskUI("example", HumanTaskUIArgs.builder()        
            .humanTaskUiName("example")
            .uiTemplate(HumanTaskUIUiTemplateArgs.builder()
                .content(Files.readString(Paths.get("sagemaker-human-task-ui-template.html")))
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:HumanTaskUI
    properties:
      humanTaskUiName: example
      uiTemplate:
        content:
          fn::readFile: sagemaker-human-task-ui-template.html
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Human Task UIs using the `human_task_ui_name`. For example:

```sh
 $ pulumi import aws:sagemaker/humanTaskUI:HumanTaskUI example example
```
 