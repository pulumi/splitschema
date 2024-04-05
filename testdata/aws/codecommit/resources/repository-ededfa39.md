Provides a CodeCommit Repository Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.codecommit.Repository("test", {
    description: "This is the Sample App Repository",
    repositoryName: "MyTestRepository",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.codecommit.Repository("test",
    description="This is the Sample App Repository",
    repository_name="MyTestRepository")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.CodeCommit.Repository("test", new()
    {
        Description = "This is the Sample App Repository",
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
		_, err := codecommit.NewRepository(ctx, "test", &codecommit.RepositoryArgs{
			Description:    pulumi.String("This is the Sample App Repository"),
			RepositoryName: pulumi.String("MyTestRepository"),
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
import com.pulumi.aws.codecommit.Repository;
import com.pulumi.aws.codecommit.RepositoryArgs;
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
        var test = new Repository("test", RepositoryArgs.builder()        
            .description("This is the Sample App Repository")
            .repositoryName("MyTestRepository")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:codecommit:Repository
    properties:
      description: This is the Sample App Repository
      repositoryName: MyTestRepository
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeCommit repository using repository name. For example:

```sh
 $ pulumi import aws:codecommit/repository:Repository imported ExistingRepo
```
 