Accepts a shared directory in a consumer account.

> **NOTE:** Destroying this resource removes the shared directory from the consumer account only.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleSharedDirectory = new aws.directoryservice.SharedDirectory("exampleSharedDirectory", {
    directoryId: aws_directory_service_directory.example.id,
    notes: "example",
    target: {
        id: data.aws_caller_identity.receiver.account_id,
    },
});
const exampleSharedDirectoryAccepter = new aws.directoryservice.SharedDirectoryAccepter("exampleSharedDirectoryAccepter", {sharedDirectoryId: exampleSharedDirectory.sharedDirectoryId}, {
    provider: "awsalternate",
});
```
```python
import pulumi
import pulumi_aws as aws

example_shared_directory = aws.directoryservice.SharedDirectory("exampleSharedDirectory",
    directory_id=aws_directory_service_directory["example"]["id"],
    notes="example",
    target=aws.directoryservice.SharedDirectoryTargetArgs(
        id=data["aws_caller_identity"]["receiver"]["account_id"],
    ))
example_shared_directory_accepter = aws.directoryservice.SharedDirectoryAccepter("exampleSharedDirectoryAccepter", shared_directory_id=example_shared_directory.shared_directory_id,
opts=pulumi.ResourceOptions(provider="awsalternate"))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleSharedDirectory = new Aws.DirectoryService.SharedDirectory("exampleSharedDirectory", new()
    {
        DirectoryId = aws_directory_service_directory.Example.Id,
        Notes = "example",
        Target = new Aws.DirectoryService.Inputs.SharedDirectoryTargetArgs
        {
            Id = data.Aws_caller_identity.Receiver.Account_id,
        },
    });

    var exampleSharedDirectoryAccepter = new Aws.DirectoryService.SharedDirectoryAccepter("exampleSharedDirectoryAccepter", new()
    {
        SharedDirectoryId = exampleSharedDirectory.SharedDirectoryId,
    }, new CustomResourceOptions
    {
        Provider = "awsalternate",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directoryservice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleSharedDirectory, err := directoryservice.NewSharedDirectory(ctx, "exampleSharedDirectory", &directoryservice.SharedDirectoryArgs{
			DirectoryId: pulumi.Any(aws_directory_service_directory.Example.Id),
			Notes:       pulumi.String("example"),
			Target: &directoryservice.SharedDirectoryTargetArgs{
				Id: pulumi.Any(data.Aws_caller_identity.Receiver.Account_id),
			},
		})
		if err != nil {
			return err
		}
		_, err = directoryservice.NewSharedDirectoryAccepter(ctx, "exampleSharedDirectoryAccepter", &directoryservice.SharedDirectoryAccepterArgs{
			SharedDirectoryId: exampleSharedDirectory.SharedDirectoryId,
		}, pulumi.Provider("awsalternate"))
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
import com.pulumi.aws.directoryservice.SharedDirectory;
import com.pulumi.aws.directoryservice.SharedDirectoryArgs;
import com.pulumi.aws.directoryservice.inputs.SharedDirectoryTargetArgs;
import com.pulumi.aws.directoryservice.SharedDirectoryAccepter;
import com.pulumi.aws.directoryservice.SharedDirectoryAccepterArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleSharedDirectory = new SharedDirectory("exampleSharedDirectory", SharedDirectoryArgs.builder()        
            .directoryId(aws_directory_service_directory.example().id())
            .notes("example")
            .target(SharedDirectoryTargetArgs.builder()
                .id(data.aws_caller_identity().receiver().account_id())
                .build())
            .build());

        var exampleSharedDirectoryAccepter = new SharedDirectoryAccepter("exampleSharedDirectoryAccepter", SharedDirectoryAccepterArgs.builder()        
            .sharedDirectoryId(exampleSharedDirectory.sharedDirectoryId())
            .build(), CustomResourceOptions.builder()
                .provider("awsalternate")
                .build());

    }
}
```
```yaml
resources:
  exampleSharedDirectory:
    type: aws:directoryservice:SharedDirectory
    properties:
      directoryId: ${aws_directory_service_directory.example.id}
      notes: example
      target:
        id: ${data.aws_caller_identity.receiver.account_id}
  exampleSharedDirectoryAccepter:
    type: aws:directoryservice:SharedDirectoryAccepter
    properties:
      sharedDirectoryId: ${exampleSharedDirectory.sharedDirectoryId}
    options:
      provider: awsalternate
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Directory Service Shared Directories using the shared directory ID. For example:

```sh
 $ pulumi import aws:directoryservice/sharedDirectoryAccepter:SharedDirectoryAccepter example d-9267633ece
```
 