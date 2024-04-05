Provides a resource to create a SES template.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myTemplate = new aws.ses.Template("myTemplate", {
    html: "<h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>",
    subject: "Greetings, {{name}}!",
    text: `Hello {{name}},
Your favorite animal is {{favoriteanimal}}.
`,
});
```
```python
import pulumi
import pulumi_aws as aws

my_template = aws.ses.Template("myTemplate",
    html="<h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>",
    subject="Greetings, {{name}}!",
    text="""Hello {{name}},
Your favorite animal is {{favoriteanimal}}.
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myTemplate = new Aws.Ses.Template("myTemplate", new()
    {
        Html = "<h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>",
        Subject = "Greetings, {{name}}!",
        Text = @"Hello {{name}},
Your favorite animal is {{favoriteanimal}}.
",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewTemplate(ctx, "myTemplate", &ses.TemplateArgs{
			Html:    pulumi.String("<h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>"),
			Subject: pulumi.String("Greetings, {{name}}!"),
			Text:    pulumi.String("Hello {{name}},\nYour favorite animal is {{favoriteanimal}}.\n"),
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
import com.pulumi.aws.ses.Template;
import com.pulumi.aws.ses.TemplateArgs;
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
        var myTemplate = new Template("myTemplate", TemplateArgs.builder()        
            .html("<h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>")
            .subject("Greetings, {{name}}!")
            .text("""
Hello {{name}},
Your favorite animal is {{favoriteanimal}}.
            """)
            .build());

    }
}
```
```yaml
resources:
  myTemplate:
    type: aws:ses:Template
    properties:
      html: <h1>Hello {{name}},</h1><p>Your favorite animal is {{favoriteanimal}}.</p>
      subject: Greetings, {{name}}!
      text: "Hello {{name}},\r\nYour favorite animal is {{favoriteanimal}}.\n"
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SES templates using the template name. For example:

```sh
 $ pulumi import aws:ses/template:Template MyTemplate MyTemplate
```
 