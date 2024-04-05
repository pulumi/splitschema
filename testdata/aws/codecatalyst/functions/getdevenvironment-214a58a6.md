Data source for managing an AWS CodeCatalyst Dev Environment.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.codecatalyst.getDevEnvironment({
    spaceName: "myspace",
    projectName: "myproject",
    envId: aws_codecatalyst_dev_environment.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codecatalyst.get_dev_environment(space_name="myspace",
    project_name="myproject",
    env_id=aws_codecatalyst_dev_environment["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.CodeCatalyst.GetDevEnvironment.Invoke(new()
    {
        SpaceName = "myspace",
        ProjectName = "myproject",
        EnvId = aws_codecatalyst_dev_environment.Example.Id,
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
		_, err := codecatalyst.LookupDevEnvironment(ctx, &codecatalyst.LookupDevEnvironmentArgs{
			SpaceName:   "myspace",
			ProjectName: "myproject",
			EnvId:       aws_codecatalyst_dev_environment.Example.Id,
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
import com.pulumi.aws.codecatalyst.CodecatalystFunctions;
import com.pulumi.aws.codecatalyst.inputs.GetDevEnvironmentArgs;
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
        final var example = CodecatalystFunctions.getDevEnvironment(GetDevEnvironmentArgs.builder()
            .spaceName("myspace")
            .projectName("myproject")
            .envId(aws_codecatalyst_dev_environment.example().id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:codecatalyst:getDevEnvironment
      Arguments:
        spaceName: myspace
        projectName: myproject
        envId: ${aws_codecatalyst_dev_environment.example.id}
```
{{% /example %}}
{{% /examples %}}