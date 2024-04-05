Resource for managing an AWS CodeCatalyst Dev Environment.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.codecatalyst.DevEnvironment("test", {
    alias: "devenv",
    ides: {
        name: "PyCharm",
        runtime: "public.ecr.aws/jetbrains/py",
    },
    inactivityTimeoutMinutes: 40,
    instanceType: "dev.standard1.small",
    persistentStorage: {
        size: 16,
    },
    projectName: "myproject",
    repositories: [{
        branchName: "main",
        repositoryName: "pulumi-provider-aws",
    }],
    spaceName: "myspace",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.codecatalyst.DevEnvironment("test",
    alias="devenv",
    ides=aws.codecatalyst.DevEnvironmentIdesArgs(
        name="PyCharm",
        runtime="public.ecr.aws/jetbrains/py",
    ),
    inactivity_timeout_minutes=40,
    instance_type="dev.standard1.small",
    persistent_storage=aws.codecatalyst.DevEnvironmentPersistentStorageArgs(
        size=16,
    ),
    project_name="myproject",
    repositories=[aws.codecatalyst.DevEnvironmentRepositoryArgs(
        branch_name="main",
        repository_name="pulumi-provider-aws",
    )],
    space_name="myspace")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.CodeCatalyst.DevEnvironment("test", new()
    {
        Alias = "devenv",
        Ides = new Aws.CodeCatalyst.Inputs.DevEnvironmentIdesArgs
        {
            Name = "PyCharm",
            Runtime = "public.ecr.aws/jetbrains/py",
        },
        InactivityTimeoutMinutes = 40,
        InstanceType = "dev.standard1.small",
        PersistentStorage = new Aws.CodeCatalyst.Inputs.DevEnvironmentPersistentStorageArgs
        {
            Size = 16,
        },
        ProjectName = "myproject",
        Repositories = new[]
        {
            new Aws.CodeCatalyst.Inputs.DevEnvironmentRepositoryArgs
            {
                BranchName = "main",
                RepositoryName = "pulumi-provider-aws",
            },
        },
        SpaceName = "myspace",
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
		_, err := codecatalyst.NewDevEnvironment(ctx, "test", &codecatalyst.DevEnvironmentArgs{
			Alias: pulumi.String("devenv"),
			Ides: &codecatalyst.DevEnvironmentIdesArgs{
				Name:    pulumi.String("PyCharm"),
				Runtime: pulumi.String("public.ecr.aws/jetbrains/py"),
			},
			InactivityTimeoutMinutes: pulumi.Int(40),
			InstanceType:             pulumi.String("dev.standard1.small"),
			PersistentStorage: &codecatalyst.DevEnvironmentPersistentStorageArgs{
				Size: pulumi.Int(16),
			},
			ProjectName: pulumi.String("myproject"),
			Repositories: codecatalyst.DevEnvironmentRepositoryArray{
				&codecatalyst.DevEnvironmentRepositoryArgs{
					BranchName:     pulumi.String("main"),
					RepositoryName: pulumi.String("pulumi-provider-aws"),
				},
			},
			SpaceName: pulumi.String("myspace"),
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
import com.pulumi.aws.codecatalyst.DevEnvironment;
import com.pulumi.aws.codecatalyst.DevEnvironmentArgs;
import com.pulumi.aws.codecatalyst.inputs.DevEnvironmentIdesArgs;
import com.pulumi.aws.codecatalyst.inputs.DevEnvironmentPersistentStorageArgs;
import com.pulumi.aws.codecatalyst.inputs.DevEnvironmentRepositoryArgs;
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
        var test = new DevEnvironment("test", DevEnvironmentArgs.builder()        
            .alias("devenv")
            .ides(DevEnvironmentIdesArgs.builder()
                .name("PyCharm")
                .runtime("public.ecr.aws/jetbrains/py")
                .build())
            .inactivityTimeoutMinutes(40)
            .instanceType("dev.standard1.small")
            .persistentStorage(DevEnvironmentPersistentStorageArgs.builder()
                .size(16)
                .build())
            .projectName("myproject")
            .repositories(DevEnvironmentRepositoryArgs.builder()
                .branchName("main")
                .repositoryName("pulumi-provider-aws")
                .build())
            .spaceName("myspace")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:codecatalyst:DevEnvironment
    properties:
      alias: devenv
      ides:
        name: PyCharm
        runtime: public.ecr.aws/jetbrains/py
      inactivityTimeoutMinutes: 40
      instanceType: dev.standard1.small
      persistentStorage:
        size: 16
      projectName: myproject
      repositories:
        - branchName: main
          repositoryName: pulumi-provider-aws
      spaceName: myspace
```
{{% /example %}}
{{% /examples %}}