Provides an Amazon Managed Grafana workspace SAML configuration resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assume = new aws.iam.Role("assume", {assumeRolePolicy: JSON.stringify({
    Version: "2012-10-17",
    Statement: [{
        Action: "sts:AssumeRole",
        Effect: "Allow",
        Sid: "",
        Principal: {
            Service: "grafana.amazonaws.com",
        },
    }],
})});
const exampleWorkspace = new aws.grafana.Workspace("exampleWorkspace", {
    accountAccessType: "CURRENT_ACCOUNT",
    authenticationProviders: ["SAML"],
    permissionType: "SERVICE_MANAGED",
    roleArn: assume.arn,
});
const exampleWorkspaceSamlConfiguration = new aws.grafana.WorkspaceSamlConfiguration("exampleWorkspaceSamlConfiguration", {
    editorRoleValues: ["editor"],
    idpMetadataUrl: "https://my_idp_metadata.url",
    workspaceId: exampleWorkspace.id,
});
```
```python
import pulumi
import json
import pulumi_aws as aws

assume = aws.iam.Role("assume", assume_role_policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Sid": "",
        "Principal": {
            "Service": "grafana.amazonaws.com",
        },
    }],
}))
example_workspace = aws.grafana.Workspace("exampleWorkspace",
    account_access_type="CURRENT_ACCOUNT",
    authentication_providers=["SAML"],
    permission_type="SERVICE_MANAGED",
    role_arn=assume.arn)
example_workspace_saml_configuration = aws.grafana.WorkspaceSamlConfiguration("exampleWorkspaceSamlConfiguration",
    editor_role_values=["editor"],
    idp_metadata_url="https://my_idp_metadata.url",
    workspace_id=example_workspace.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var assume = new Aws.Iam.Role("assume", new()
    {
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = "sts:AssumeRole",
                    ["Effect"] = "Allow",
                    ["Sid"] = "",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "grafana.amazonaws.com",
                    },
                },
            },
        }),
    });

    var exampleWorkspace = new Aws.Grafana.Workspace("exampleWorkspace", new()
    {
        AccountAccessType = "CURRENT_ACCOUNT",
        AuthenticationProviders = new[]
        {
            "SAML",
        },
        PermissionType = "SERVICE_MANAGED",
        RoleArn = assume.Arn,
    });

    var exampleWorkspaceSamlConfiguration = new Aws.Grafana.WorkspaceSamlConfiguration("exampleWorkspaceSamlConfiguration", new()
    {
        EditorRoleValues = new[]
        {
            "editor",
        },
        IdpMetadataUrl = "https://my_idp_metadata.url",
        WorkspaceId = exampleWorkspace.Id,
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/grafana"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Sid":    "",
					"Principal": map[string]interface{}{
						"Service": "grafana.amazonaws.com",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		assume, err := iam.NewRole(ctx, "assume", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		exampleWorkspace, err := grafana.NewWorkspace(ctx, "exampleWorkspace", &grafana.WorkspaceArgs{
			AccountAccessType: pulumi.String("CURRENT_ACCOUNT"),
			AuthenticationProviders: pulumi.StringArray{
				pulumi.String("SAML"),
			},
			PermissionType: pulumi.String("SERVICE_MANAGED"),
			RoleArn:        assume.Arn,
		})
		if err != nil {
			return err
		}
		_, err = grafana.NewWorkspaceSamlConfiguration(ctx, "exampleWorkspaceSamlConfiguration", &grafana.WorkspaceSamlConfigurationArgs{
			EditorRoleValues: pulumi.StringArray{
				pulumi.String("editor"),
			},
			IdpMetadataUrl: pulumi.String("https://my_idp_metadata.url"),
			WorkspaceId:    exampleWorkspace.ID(),
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.grafana.Workspace;
import com.pulumi.aws.grafana.WorkspaceArgs;
import com.pulumi.aws.grafana.WorkspaceSamlConfiguration;
import com.pulumi.aws.grafana.WorkspaceSamlConfigurationArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var assume = new Role("assume", RoleArgs.builder()        
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", "sts:AssumeRole"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Sid", ""),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "grafana.amazonaws.com")
                        ))
                    )))
                )))
            .build());

        var exampleWorkspace = new Workspace("exampleWorkspace", WorkspaceArgs.builder()        
            .accountAccessType("CURRENT_ACCOUNT")
            .authenticationProviders("SAML")
            .permissionType("SERVICE_MANAGED")
            .roleArn(assume.arn())
            .build());

        var exampleWorkspaceSamlConfiguration = new WorkspaceSamlConfiguration("exampleWorkspaceSamlConfiguration", WorkspaceSamlConfigurationArgs.builder()        
            .editorRoleValues("editor")
            .idpMetadataUrl("https://my_idp_metadata.url")
            .workspaceId(exampleWorkspace.id())
            .build());

    }
}
```
```yaml
resources:
  exampleWorkspaceSamlConfiguration:
    type: aws:grafana:WorkspaceSamlConfiguration
    properties:
      editorRoleValues:
        - editor
      idpMetadataUrl: https://my_idp_metadata.url
      workspaceId: ${exampleWorkspace.id}
  exampleWorkspace:
    type: aws:grafana:Workspace
    properties:
      accountAccessType: CURRENT_ACCOUNT
      authenticationProviders:
        - SAML
      permissionType: SERVICE_MANAGED
      roleArn: ${assume.arn}
  assume:
    type: aws:iam:Role
    properties:
      assumeRolePolicy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action: sts:AssumeRole
              Effect: Allow
              Sid:
              Principal:
                Service: grafana.amazonaws.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Grafana Workspace SAML configuration using the workspace's `id`. For example:

```sh
 $ pulumi import aws:grafana/workspaceSamlConfiguration:WorkspaceSamlConfiguration example g-2054c75a02
```
 