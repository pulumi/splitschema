Resource for managing an AWS CodeCatalyst Project.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.codecatalyst.Project("test", {
    description: "My CodeCatalyst Project created using Pulumi",
    displayName: "MyProject",
    spaceName: "myproject",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.codecatalyst.Project("test",
    description="My CodeCatalyst Project created using Pulumi",
    display_name="MyProject",
    space_name="myproject")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.CodeCatalyst.Project("test", new()
    {
        Description = "My CodeCatalyst Project created using Pulumi",
        DisplayName = "MyProject",
        SpaceName = "myproject",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codecatalyst"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codecatalyst.NewProject(ctx, "test", &codecatalyst.ProjectArgs{
			Description: pulumi.String("My CodeCatalyst Project created using Pulumi"),
			DisplayName: pulumi.String("MyProject"),
			SpaceName:   pulumi.String("myproject"),
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
import com.pulumi.aws.codecatalyst.Project;
import com.pulumi.aws.codecatalyst.ProjectArgs;
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
        var test = new Project("test", ProjectArgs.builder()        
            .description("My CodeCatalyst Project created using Pulumi")
            .displayName("MyProject")
            .spaceName("myproject")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:codecatalyst:Project
    properties:
      description: My CodeCatalyst Project created using Pulumi
      displayName: MyProject
      spaceName: myproject
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeCatalyst Project using the `id`. For example:

```sh
 $ pulumi import aws:codecatalyst/project:Project example project-id-12345678
```
 