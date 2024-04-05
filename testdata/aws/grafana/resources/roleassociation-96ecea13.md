Provides an Amazon Managed Grafana workspace role association resource.

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
const exampleRoleAssociation = new aws.grafana.RoleAssociation("exampleRoleAssociation", {
    role: "ADMIN",
    userIds: [
        "USER_ID_1",
        "USER_ID_2",
    ],
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
example_role_association = aws.grafana.RoleAssociation("exampleRoleAssociation",
    role="ADMIN",
    user_ids=[
        "USER_ID_1",
        "USER_ID_2",
    ],
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

    var exampleRoleAssociation = new Aws.Grafana.RoleAssociation("exampleRoleAssociation", new()
    {
        Role = "ADMIN",
        UserIds = new[]
        {
            "USER_ID_1",
            "USER_ID_2",
        },
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
		_, err = grafana.NewRoleAssociation(ctx, "exampleRoleAssociation", &grafana.RoleAssociationArgs{
			Role: pulumi.String("ADMIN"),
			UserIds: pulumi.StringArray{
				pulumi.String("USER_ID_1"),
				pulumi.String("USER_ID_2"),
			},
			WorkspaceId: exampleWorkspace.ID(),
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
import com.pulumi.aws.grafana.RoleAssociation;
import com.pulumi.aws.grafana.RoleAssociationArgs;
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

        var exampleRoleAssociation = new RoleAssociation("exampleRoleAssociation", RoleAssociationArgs.builder()        
            .role("ADMIN")
            .userIds(            
                "USER_ID_1",
                "USER_ID_2")
            .workspaceId(exampleWorkspace.id())
            .build());

    }
}
```
```yaml
resources:
  exampleRoleAssociation:
    type: aws:grafana:RoleAssociation
    properties:
      role: ADMIN
      userIds:
        - USER_ID_1
        - USER_ID_2
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