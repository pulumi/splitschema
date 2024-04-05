Resource for managing an AWS CodeCatalyst Source Repository.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codecatalyst.SourceRepository("example", {
    projectName: "example-project",
    spaceName: "example-space",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codecatalyst.SourceRepository("example",
    project_name="example-project",
    space_name="example-space")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeCatalyst.SourceRepository("example", new()
    {
        ProjectName = "example-project",
        SpaceName = "example-space",
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
		_, err := codecatalyst.NewSourceRepository(ctx, "example", &codecatalyst.SourceRepositoryArgs{
			ProjectName: pulumi.String("example-project"),
			SpaceName:   pulumi.String("example-space"),
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
import com.pulumi.aws.codecatalyst.SourceRepository;
import com.pulumi.aws.codecatalyst.SourceRepositoryArgs;
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
        var example = new SourceRepository("example", SourceRepositoryArgs.builder()        
            .projectName("example-project")
            .spaceName("example-space")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codecatalyst:SourceRepository
    properties:
      projectName: example-project
      spaceName: example-space
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeCatalyst Source Repository using the `id`. For example:

```sh
 $ pulumi import aws:codecatalyst/sourceRepository:SourceRepository example example-repo
```
 