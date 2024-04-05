Use this data source to get the name of a elastic beanstalk solution stack.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const multiDocker = aws.elasticbeanstalk.getSolutionStack({
    mostRecent: true,
    nameRegex: "^64bit Amazon Linux (.*) Multi-container Docker (.*)$",
});
```
```python
import pulumi
import pulumi_aws as aws

multi_docker = aws.elasticbeanstalk.get_solution_stack(most_recent=True,
    name_regex="^64bit Amazon Linux (.*) Multi-container Docker (.*)$")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var multiDocker = Aws.ElasticBeanstalk.GetSolutionStack.Invoke(new()
    {
        MostRecent = true,
        NameRegex = "^64bit Amazon Linux (.*) Multi-container Docker (.*)$",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticbeanstalk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticbeanstalk.GetSolutionStack(ctx, &elasticbeanstalk.GetSolutionStackArgs{
			MostRecent: pulumi.BoolRef(true),
			NameRegex:  "^64bit Amazon Linux (.*) Multi-container Docker (.*)$",
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
import com.pulumi.aws.elasticbeanstalk.ElasticbeanstalkFunctions;
import com.pulumi.aws.elasticbeanstalk.inputs.GetSolutionStackArgs;
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
        final var multiDocker = ElasticbeanstalkFunctions.getSolutionStack(GetSolutionStackArgs.builder()
            .mostRecent(true)
            .nameRegex("^64bit Amazon Linux (.*) Multi-container Docker (.*)$")
            .build());

    }
}
```
```yaml
variables:
  multiDocker:
    fn::invoke:
      Function: aws:elasticbeanstalk:getSolutionStack
      Arguments:
        mostRecent: true
        nameRegex: ^64bit Amazon Linux (.*) Multi-container Docker (.*)$
```
{{% /example %}}
{{% /examples %}}