Provides a workspace in [AWS Workspaces](https://docs.aws.amazon.com/workspaces/latest/adminguide/amazon-workspaces.html) Service

> **NOTE:** AWS WorkSpaces service requires [`workspaces_DefaultRole`](https://docs.aws.amazon.com/workspaces/latest/adminguide/workspaces-access-control.html#create-default-role) IAM role to operate normally.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const valueWindows10 = aws.workspaces.getBundle({
    bundleId: "wsb-bh8rsxt14",
});
const example = new aws.workspaces.Workspace("example", {
    directoryId: aws_workspaces_directory.example.id,
    bundleId: valueWindows10.then(valueWindows10 => valueWindows10.id),
    userName: "john.doe",
    rootVolumeEncryptionEnabled: true,
    userVolumeEncryptionEnabled: true,
    volumeEncryptionKey: "alias/aws/workspaces",
    workspaceProperties: {
        computeTypeName: "VALUE",
        userVolumeSizeGib: 10,
        rootVolumeSizeGib: 80,
        runningMode: "AUTO_STOP",
        runningModeAutoStopTimeoutInMinutes: 60,
    },
    tags: {
        Department: "IT",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

value_windows10 = aws.workspaces.get_bundle(bundle_id="wsb-bh8rsxt14")
example = aws.workspaces.Workspace("example",
    directory_id=aws_workspaces_directory["example"]["id"],
    bundle_id=value_windows10.id,
    user_name="john.doe",
    root_volume_encryption_enabled=True,
    user_volume_encryption_enabled=True,
    volume_encryption_key="alias/aws/workspaces",
    workspace_properties=aws.workspaces.WorkspaceWorkspacePropertiesArgs(
        compute_type_name="VALUE",
        user_volume_size_gib=10,
        root_volume_size_gib=80,
        running_mode="AUTO_STOP",
        running_mode_auto_stop_timeout_in_minutes=60,
    ),
    tags={
        "Department": "IT",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var valueWindows10 = Aws.Workspaces.GetBundle.Invoke(new()
    {
        BundleId = "wsb-bh8rsxt14",
    });

    var example = new Aws.Workspaces.Workspace("example", new()
    {
        DirectoryId = aws_workspaces_directory.Example.Id,
        BundleId = valueWindows10.Apply(getBundleResult => getBundleResult.Id),
        UserName = "john.doe",
        RootVolumeEncryptionEnabled = true,
        UserVolumeEncryptionEnabled = true,
        VolumeEncryptionKey = "alias/aws/workspaces",
        WorkspaceProperties = new Aws.Workspaces.Inputs.WorkspaceWorkspacePropertiesArgs
        {
            ComputeTypeName = "VALUE",
            UserVolumeSizeGib = 10,
            RootVolumeSizeGib = 80,
            RunningMode = "AUTO_STOP",
            RunningModeAutoStopTimeoutInMinutes = 60,
        },
        Tags = 
        {
            { "Department", "IT" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/workspaces"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		valueWindows10, err := workspaces.GetBundle(ctx, &workspaces.GetBundleArgs{
			BundleId: pulumi.StringRef("wsb-bh8rsxt14"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = workspaces.NewWorkspace(ctx, "example", &workspaces.WorkspaceArgs{
			DirectoryId:                 pulumi.Any(aws_workspaces_directory.Example.Id),
			BundleId:                    *pulumi.String(valueWindows10.Id),
			UserName:                    pulumi.String("john.doe"),
			RootVolumeEncryptionEnabled: pulumi.Bool(true),
			UserVolumeEncryptionEnabled: pulumi.Bool(true),
			VolumeEncryptionKey:         pulumi.String("alias/aws/workspaces"),
			WorkspaceProperties: &workspaces.WorkspaceWorkspacePropertiesArgs{
				ComputeTypeName:                     pulumi.String("VALUE"),
				UserVolumeSizeGib:                   pulumi.Int(10),
				RootVolumeSizeGib:                   pulumi.Int(80),
				RunningMode:                         pulumi.String("AUTO_STOP"),
				RunningModeAutoStopTimeoutInMinutes: pulumi.Int(60),
			},
			Tags: pulumi.StringMap{
				"Department": pulumi.String("IT"),
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
import com.pulumi.aws.workspaces.WorkspacesFunctions;
import com.pulumi.aws.workspaces.inputs.GetBundleArgs;
import com.pulumi.aws.workspaces.Workspace;
import com.pulumi.aws.workspaces.WorkspaceArgs;
import com.pulumi.aws.workspaces.inputs.WorkspaceWorkspacePropertiesArgs;
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
        final var valueWindows10 = WorkspacesFunctions.getBundle(GetBundleArgs.builder()
            .bundleId("wsb-bh8rsxt14")
            .build());

        var example = new Workspace("example", WorkspaceArgs.builder()        
            .directoryId(aws_workspaces_directory.example().id())
            .bundleId(valueWindows10.applyValue(getBundleResult -> getBundleResult.id()))
            .userName("john.doe")
            .rootVolumeEncryptionEnabled(true)
            .userVolumeEncryptionEnabled(true)
            .volumeEncryptionKey("alias/aws/workspaces")
            .workspaceProperties(WorkspaceWorkspacePropertiesArgs.builder()
                .computeTypeName("VALUE")
                .userVolumeSizeGib(10)
                .rootVolumeSizeGib(80)
                .runningMode("AUTO_STOP")
                .runningModeAutoStopTimeoutInMinutes(60)
                .build())
            .tags(Map.of("Department", "IT"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:workspaces:Workspace
    properties:
      directoryId: ${aws_workspaces_directory.example.id}
      bundleId: ${valueWindows10.id}
      userName: john.doe
      rootVolumeEncryptionEnabled: true
      userVolumeEncryptionEnabled: true
      volumeEncryptionKey: alias/aws/workspaces
      workspaceProperties:
        computeTypeName: VALUE
        userVolumeSizeGib: 10
        rootVolumeSizeGib: 80
        runningMode: AUTO_STOP
        runningModeAutoStopTimeoutInMinutes: 60
      tags:
        Department: IT
variables:
  valueWindows10:
    fn::invoke:
      Function: aws:workspaces:getBundle
      Arguments:
        bundleId: wsb-bh8rsxt14
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Workspaces using their ID. For example:

```sh
 $ pulumi import aws:workspaces/workspace:Workspace example ws-9z9zmbkhv
```
 