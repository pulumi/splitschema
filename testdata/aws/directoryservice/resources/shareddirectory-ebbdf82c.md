Manages a directory in your account (directory owner) shared with another account (directory consumer).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleDirectory = new aws.directoryservice.Directory("exampleDirectory", {
    name: "tf-example",
    password: "SuperSecretPassw0rd",
    type: "MicrosoftAD",
    edition: "Standard",
    vpcSettings: {
        vpcId: aws_vpc.example.id,
        subnetIds: aws_subnet.example.map(__item => __item.id),
    },
});
const exampleSharedDirectory = new aws.directoryservice.SharedDirectory("exampleSharedDirectory", {
    directoryId: exampleDirectory.id,
    notes: "You wanna have a catch?",
    target: {
        id: data.aws_caller_identity.receiver.account_id,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_directory = aws.directoryservice.Directory("exampleDirectory",
    name="tf-example",
    password="SuperSecretPassw0rd",
    type="MicrosoftAD",
    edition="Standard",
    vpc_settings=aws.directoryservice.DirectoryVpcSettingsArgs(
        vpc_id=aws_vpc["example"]["id"],
        subnet_ids=[__item["id"] for __item in aws_subnet["example"]],
    ))
example_shared_directory = aws.directoryservice.SharedDirectory("exampleSharedDirectory",
    directory_id=example_directory.id,
    notes="You wanna have a catch?",
    target=aws.directoryservice.SharedDirectoryTargetArgs(
        id=data["aws_caller_identity"]["receiver"]["account_id"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleDirectory = new Aws.DirectoryService.Directory("exampleDirectory", new()
    {
        Name = "tf-example",
        Password = "SuperSecretPassw0rd",
        Type = "MicrosoftAD",
        Edition = "Standard",
        VpcSettings = new Aws.DirectoryService.Inputs.DirectoryVpcSettingsArgs
        {
            VpcId = aws_vpc.Example.Id,
            SubnetIds = aws_subnet.Example.Select(__item => __item.Id).ToList(),
        },
    });

    var exampleSharedDirectory = new Aws.DirectoryService.SharedDirectory("exampleSharedDirectory", new()
    {
        DirectoryId = exampleDirectory.Id,
        Notes = "You wanna have a catch?",
        Target = new Aws.DirectoryService.Inputs.SharedDirectoryTargetArgs
        {
            Id = data.Aws_caller_identity.Receiver.Account_id,
        },
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
exampleDirectory, err := directoryservice.NewDirectory(ctx, "exampleDirectory", &directoryservice.DirectoryArgs{
Name: pulumi.String("tf-example"),
Password: pulumi.String("SuperSecretPassw0rd"),
Type: pulumi.String("MicrosoftAD"),
Edition: pulumi.String("Standard"),
VpcSettings: &directoryservice.DirectoryVpcSettingsArgs{
VpcId: pulumi.Any(aws_vpc.Example.Id),
SubnetIds: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-resources-aws:directoryservice-sharedDirectory:SharedDirectory.pp:7,17-41),
},
})
if err != nil {
return err
}
_, err = directoryservice.NewSharedDirectory(ctx, "exampleSharedDirectory", &directoryservice.SharedDirectoryArgs{
DirectoryId: exampleDirectory.ID(),
Notes: pulumi.String("You wanna have a catch?"),
Target: &directoryservice.SharedDirectoryTargetArgs{
Id: pulumi.Any(data.Aws_caller_identity.Receiver.Account_id),
},
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
import com.pulumi.aws.directoryservice.Directory;
import com.pulumi.aws.directoryservice.DirectoryArgs;
import com.pulumi.aws.directoryservice.inputs.DirectoryVpcSettingsArgs;
import com.pulumi.aws.directoryservice.SharedDirectory;
import com.pulumi.aws.directoryservice.SharedDirectoryArgs;
import com.pulumi.aws.directoryservice.inputs.SharedDirectoryTargetArgs;
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
        var exampleDirectory = new Directory("exampleDirectory", DirectoryArgs.builder()        
            .name("tf-example")
            .password("SuperSecretPassw0rd")
            .type("MicrosoftAD")
            .edition("Standard")
            .vpcSettings(DirectoryVpcSettingsArgs.builder()
                .vpcId(aws_vpc.example().id())
                .subnetIds(aws_subnet.example().stream().map(element -> element.id()).collect(toList()))
                .build())
            .build());

        var exampleSharedDirectory = new SharedDirectory("exampleSharedDirectory", SharedDirectoryArgs.builder()        
            .directoryId(exampleDirectory.id())
            .notes("You wanna have a catch?")
            .target(SharedDirectoryTargetArgs.builder()
                .id(data.aws_caller_identity().receiver().account_id())
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Directory Service Shared Directories using the owner directory ID/shared directory ID. For example:

```sh
 $ pulumi import aws:directoryservice/sharedDirectory:SharedDirectory example d-1234567890/d-9267633ece
```
 