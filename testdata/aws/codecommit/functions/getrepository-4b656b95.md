The CodeCommit Repository data source allows the ARN, Repository ID, Repository URL for HTTP and Repository URL for SSH to be retrieved for an CodeCommit repository.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.codecommit.getRepository({
    repositoryName: "MyTestRepository",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.codecommit.get_repository(repository_name="MyTestRepository")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.CodeCommit.GetRepository.Invoke(new()
    {
        RepositoryName = "MyTestRepository",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codecommit"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codecommit.LookupRepository(ctx, &codecommit.LookupRepositoryArgs{
			RepositoryName: "MyTestRepository",
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
import com.pulumi.aws.codecommit.CodecommitFunctions;
import com.pulumi.aws.codecommit.inputs.GetRepositoryArgs;
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
        final var test = CodecommitFunctions.getRepository(GetRepositoryArgs.builder()
            .repositoryName("MyTestRepository")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:codecommit:getRepository
      Arguments:
        repositoryName: MyTestRepository
```
{{% /example %}}
{{% /examples %}}