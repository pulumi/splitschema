The ECR Repository data source allows the ARN, Repository URI and Registry ID to be retrieved for an ECR repository.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const service = aws.ecr.getRepository({
    name: "ecr-repository",
});
```
```python
import pulumi
import pulumi_aws as aws

service = aws.ecr.get_repository(name="ecr-repository")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var service = Aws.Ecr.GetRepository.Invoke(new()
    {
        Name = "ecr-repository",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecr.LookupRepository(ctx, &ecr.LookupRepositoryArgs{
			Name: "ecr-repository",
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
import com.pulumi.aws.ecr.EcrFunctions;
import com.pulumi.aws.ecr.inputs.GetRepositoryArgs;
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
        final var service = EcrFunctions.getRepository(GetRepositoryArgs.builder()
            .name("ecr-repository")
            .build());

    }
}
```
```yaml
variables:
  service:
    fn::invoke:
      Function: aws:ecr:getRepository
      Arguments:
        name: ecr-repository
```
{{% /example %}}
{{% /examples %}}