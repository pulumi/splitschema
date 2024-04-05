Provides an AppStream stack.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appstream.Stack("example", {
    applicationSettings: {
        enabled: true,
        settingsGroup: "SettingsGroup",
    },
    description: "stack description",
    displayName: "stack display name",
    feedbackUrl: "http://your-domain/feedback",
    redirectUrl: "http://your-domain/redirect",
    storageConnectors: [{
        connectorType: "HOMEFOLDERS",
    }],
    tags: {
        TagName: "TagValue",
    },
    userSettings: [
        {
            action: "CLIPBOARD_COPY_FROM_LOCAL_DEVICE",
            permission: "ENABLED",
        },
        {
            action: "CLIPBOARD_COPY_TO_LOCAL_DEVICE",
            permission: "ENABLED",
        },
        {
            action: "DOMAIN_PASSWORD_SIGNIN",
            permission: "ENABLED",
        },
        {
            action: "DOMAIN_SMART_CARD_SIGNIN",
            permission: "DISABLED",
        },
        {
            action: "FILE_DOWNLOAD",
            permission: "ENABLED",
        },
        {
            action: "FILE_UPLOAD",
            permission: "ENABLED",
        },
        {
            action: "PRINTING_TO_LOCAL_DEVICE",
            permission: "ENABLED",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appstream.Stack("example",
    application_settings=aws.appstream.StackApplicationSettingsArgs(
        enabled=True,
        settings_group="SettingsGroup",
    ),
    description="stack description",
    display_name="stack display name",
    feedback_url="http://your-domain/feedback",
    redirect_url="http://your-domain/redirect",
    storage_connectors=[aws.appstream.StackStorageConnectorArgs(
        connector_type="HOMEFOLDERS",
    )],
    tags={
        "TagName": "TagValue",
    },
    user_settings=[
        aws.appstream.StackUserSettingArgs(
            action="CLIPBOARD_COPY_FROM_LOCAL_DEVICE",
            permission="ENABLED",
        ),
        aws.appstream.StackUserSettingArgs(
            action="CLIPBOARD_COPY_TO_LOCAL_DEVICE",
            permission="ENABLED",
        ),
        aws.appstream.StackUserSettingArgs(
            action="DOMAIN_PASSWORD_SIGNIN",
            permission="ENABLED",
        ),
        aws.appstream.StackUserSettingArgs(
            action="DOMAIN_SMART_CARD_SIGNIN",
            permission="DISABLED",
        ),
        aws.appstream.StackUserSettingArgs(
            action="FILE_DOWNLOAD",
            permission="ENABLED",
        ),
        aws.appstream.StackUserSettingArgs(
            action="FILE_UPLOAD",
            permission="ENABLED",
        ),
        aws.appstream.StackUserSettingArgs(
            action="PRINTING_TO_LOCAL_DEVICE",
            permission="ENABLED",
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppStream.Stack("example", new()
    {
        ApplicationSettings = new Aws.AppStream.Inputs.StackApplicationSettingsArgs
        {
            Enabled = true,
            SettingsGroup = "SettingsGroup",
        },
        Description = "stack description",
        DisplayName = "stack display name",
        FeedbackUrl = "http://your-domain/feedback",
        RedirectUrl = "http://your-domain/redirect",
        StorageConnectors = new[]
        {
            new Aws.AppStream.Inputs.StackStorageConnectorArgs
            {
                ConnectorType = "HOMEFOLDERS",
            },
        },
        Tags = 
        {
            { "TagName", "TagValue" },
        },
        UserSettings = new[]
        {
            new Aws.AppStream.Inputs.StackUserSettingArgs
            {
                Action = "CLIPBOARD_COPY_FROM_LOCAL_DEVICE",
                Permission = "ENABLED",
            },
            new Aws.AppStream.Inputs.StackUserSettingArgs
            {
                Action = "CLIPBOARD_COPY_TO_LOCAL_DEVICE",
                Permission = "ENABLED",
            },
            new Aws.AppStream.Inputs.StackUserSettingArgs
            {
                Action = "DOMAIN_PASSWORD_SIGNIN",
                Permission = "ENABLED",
            },
            new Aws.AppStream.Inputs.StackUserSettingArgs
            {
                Action = "DOMAIN_SMART_CARD_SIGNIN",
                Permission = "DISABLED",
            },
            new Aws.AppStream.Inputs.StackUserSettingArgs
            {
                Action = "FILE_DOWNLOAD",
                Permission = "ENABLED",
            },
            new Aws.AppStream.Inputs.StackUserSettingArgs
            {
                Action = "FILE_UPLOAD",
                Permission = "ENABLED",
            },
            new Aws.AppStream.Inputs.StackUserSettingArgs
            {
                Action = "PRINTING_TO_LOCAL_DEVICE",
                Permission = "ENABLED",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appstream"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appstream.NewStack(ctx, "example", &appstream.StackArgs{
			ApplicationSettings: &appstream.StackApplicationSettingsArgs{
				Enabled:       pulumi.Bool(true),
				SettingsGroup: pulumi.String("SettingsGroup"),
			},
			Description: pulumi.String("stack description"),
			DisplayName: pulumi.String("stack display name"),
			FeedbackUrl: pulumi.String("http://your-domain/feedback"),
			RedirectUrl: pulumi.String("http://your-domain/redirect"),
			StorageConnectors: appstream.StackStorageConnectorArray{
				&appstream.StackStorageConnectorArgs{
					ConnectorType: pulumi.String("HOMEFOLDERS"),
				},
			},
			Tags: pulumi.StringMap{
				"TagName": pulumi.String("TagValue"),
			},
			UserSettings: appstream.StackUserSettingArray{
				&appstream.StackUserSettingArgs{
					Action:     pulumi.String("CLIPBOARD_COPY_FROM_LOCAL_DEVICE"),
					Permission: pulumi.String("ENABLED"),
				},
				&appstream.StackUserSettingArgs{
					Action:     pulumi.String("CLIPBOARD_COPY_TO_LOCAL_DEVICE"),
					Permission: pulumi.String("ENABLED"),
				},
				&appstream.StackUserSettingArgs{
					Action:     pulumi.String("DOMAIN_PASSWORD_SIGNIN"),
					Permission: pulumi.String("ENABLED"),
				},
				&appstream.StackUserSettingArgs{
					Action:     pulumi.String("DOMAIN_SMART_CARD_SIGNIN"),
					Permission: pulumi.String("DISABLED"),
				},
				&appstream.StackUserSettingArgs{
					Action:     pulumi.String("FILE_DOWNLOAD"),
					Permission: pulumi.String("ENABLED"),
				},
				&appstream.StackUserSettingArgs{
					Action:     pulumi.String("FILE_UPLOAD"),
					Permission: pulumi.String("ENABLED"),
				},
				&appstream.StackUserSettingArgs{
					Action:     pulumi.String("PRINTING_TO_LOCAL_DEVICE"),
					Permission: pulumi.String("ENABLED"),
				},
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
import com.pulumi.aws.appstream.Stack;
import com.pulumi.aws.appstream.StackArgs;
import com.pulumi.aws.appstream.inputs.StackApplicationSettingsArgs;
import com.pulumi.aws.appstream.inputs.StackStorageConnectorArgs;
import com.pulumi.aws.appstream.inputs.StackUserSettingArgs;
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
        var example = new Stack("example", StackArgs.builder()        
            .applicationSettings(StackApplicationSettingsArgs.builder()
                .enabled(true)
                .settingsGroup("SettingsGroup")
                .build())
            .description("stack description")
            .displayName("stack display name")
            .feedbackUrl("http://your-domain/feedback")
            .redirectUrl("http://your-domain/redirect")
            .storageConnectors(StackStorageConnectorArgs.builder()
                .connectorType("HOMEFOLDERS")
                .build())
            .tags(Map.of("TagName", "TagValue"))
            .userSettings(            
                StackUserSettingArgs.builder()
                    .action("CLIPBOARD_COPY_FROM_LOCAL_DEVICE")
                    .permission("ENABLED")
                    .build(),
                StackUserSettingArgs.builder()
                    .action("CLIPBOARD_COPY_TO_LOCAL_DEVICE")
                    .permission("ENABLED")
                    .build(),
                StackUserSettingArgs.builder()
                    .action("DOMAIN_PASSWORD_SIGNIN")
                    .permission("ENABLED")
                    .build(),
                StackUserSettingArgs.builder()
                    .action("DOMAIN_SMART_CARD_SIGNIN")
                    .permission("DISABLED")
                    .build(),
                StackUserSettingArgs.builder()
                    .action("FILE_DOWNLOAD")
                    .permission("ENABLED")
                    .build(),
                StackUserSettingArgs.builder()
                    .action("FILE_UPLOAD")
                    .permission("ENABLED")
                    .build(),
                StackUserSettingArgs.builder()
                    .action("PRINTING_TO_LOCAL_DEVICE")
                    .permission("ENABLED")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appstream:Stack
    properties:
      applicationSettings:
        enabled: true
        settingsGroup: SettingsGroup
      description: stack description
      displayName: stack display name
      feedbackUrl: http://your-domain/feedback
      redirectUrl: http://your-domain/redirect
      storageConnectors:
        - connectorType: HOMEFOLDERS
      tags:
        TagName: TagValue
      userSettings:
        - action: CLIPBOARD_COPY_FROM_LOCAL_DEVICE
          permission: ENABLED
        - action: CLIPBOARD_COPY_TO_LOCAL_DEVICE
          permission: ENABLED
        - action: DOMAIN_PASSWORD_SIGNIN
          permission: ENABLED
        - action: DOMAIN_SMART_CARD_SIGNIN
          permission: DISABLED
        - action: FILE_DOWNLOAD
          permission: ENABLED
        - action: FILE_UPLOAD
          permission: ENABLED
        - action: PRINTING_TO_LOCAL_DEVICE
          permission: ENABLED
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_appstream_stack` using the id. For example:

```sh
 $ pulumi import aws:appstream/stack:Stack example stackID
```
 