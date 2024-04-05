Provides a AWS Transfer User SSH Key resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleServer = new aws.transfer.Server("exampleServer", {
    identityProviderType: "SERVICE_MANAGED",
    tags: {
        NAME: "tf-acc-test-transfer-server",
    },
});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["transfer.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const exampleUser = new aws.transfer.User("exampleUser", {
    serverId: exampleServer.id,
    userName: "tftestuser",
    role: exampleRole.arn,
    tags: {
        NAME: "tftestuser",
    },
});
const exampleSshKey = new aws.transfer.SshKey("exampleSshKey", {
    serverId: exampleServer.id,
    userName: exampleUser.userName,
    body: "... SSH key ...",
});
const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        sid: "AllowFullAccesstoS3",
        effect: "Allow",
        actions: ["s3:*"],
        resources: ["*"],
    }],
});
const exampleRolePolicy = new aws.iam.RolePolicy("exampleRolePolicy", {
    role: exampleRole.id,
    policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example_server = aws.transfer.Server("exampleServer",
    identity_provider_type="SERVICE_MANAGED",
    tags={
        "NAME": "tf-acc-test-transfer-server",
    })
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["transfer.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
example_user = aws.transfer.User("exampleUser",
    server_id=example_server.id,
    user_name="tftestuser",
    role=example_role.arn,
    tags={
        "NAME": "tftestuser",
    })
example_ssh_key = aws.transfer.SshKey("exampleSshKey",
    server_id=example_server.id,
    user_name=example_user.user_name,
    body="... SSH key ...")
example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="AllowFullAccesstoS3",
    effect="Allow",
    actions=["s3:*"],
    resources=["*"],
)])
example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
    role=example_role.id,
    policy=example_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleServer = new Aws.Transfer.Server("exampleServer", new()
    {
        IdentityProviderType = "SERVICE_MANAGED",
        Tags = 
        {
            { "NAME", "tf-acc-test-transfer-server" },
        },
    });

    var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "transfer.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleUser = new Aws.Transfer.User("exampleUser", new()
    {
        ServerId = exampleServer.Id,
        UserName = "tftestuser",
        Role = exampleRole.Arn,
        Tags = 
        {
            { "NAME", "tftestuser" },
        },
    });

    var exampleSshKey = new Aws.Transfer.SshKey("exampleSshKey", new()
    {
        ServerId = exampleServer.Id,
        UserName = exampleUser.UserName,
        Body = "... SSH key ...",
    });

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "AllowFullAccesstoS3",
                Effect = "Allow",
                Actions = new[]
                {
                    "s3:*",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var exampleRolePolicy = new Aws.Iam.RolePolicy("exampleRolePolicy", new()
    {
        Role = exampleRole.Id,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleServer, err := transfer.NewServer(ctx, "exampleServer", &transfer.ServerArgs{
			IdentityProviderType: pulumi.String("SERVICE_MANAGED"),
			Tags: pulumi.StringMap{
				"NAME": pulumi.String("tf-acc-test-transfer-server"),
			},
		})
		if err != nil {
			return err
		}
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"transfer.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		exampleUser, err := transfer.NewUser(ctx, "exampleUser", &transfer.UserArgs{
			ServerId: exampleServer.ID(),
			UserName: pulumi.String("tftestuser"),
			Role:     exampleRole.Arn,
			Tags: pulumi.StringMap{
				"NAME": pulumi.String("tftestuser"),
			},
		})
		if err != nil {
			return err
		}
		_, err = transfer.NewSshKey(ctx, "exampleSshKey", &transfer.SshKeyArgs{
			ServerId: exampleServer.ID(),
			UserName: exampleUser.UserName,
			Body:     pulumi.String("... SSH key ..."),
		})
		if err != nil {
			return err
		}
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("AllowFullAccesstoS3"),
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"s3:*",
					},
					Resources: []string{
						"*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "exampleRolePolicy", &iam.RolePolicyArgs{
			Role:   exampleRole.ID(),
			Policy: *pulumi.String(examplePolicyDocument.Json),
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
import com.pulumi.aws.transfer.Server;
import com.pulumi.aws.transfer.ServerArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.transfer.User;
import com.pulumi.aws.transfer.UserArgs;
import com.pulumi.aws.transfer.SshKey;
import com.pulumi.aws.transfer.SshKeyArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
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
        var exampleServer = new Server("exampleServer", ServerArgs.builder()        
            .identityProviderType("SERVICE_MANAGED")
            .tags(Map.of("NAME", "tf-acc-test-transfer-server"))
            .build());

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("transfer.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleUser = new User("exampleUser", UserArgs.builder()        
            .serverId(exampleServer.id())
            .userName("tftestuser")
            .role(exampleRole.arn())
            .tags(Map.of("NAME", "tftestuser"))
            .build());

        var exampleSshKey = new SshKey("exampleSshKey", SshKeyArgs.builder()        
            .serverId(exampleServer.id())
            .userName(exampleUser.userName())
            .body("... SSH key ...")
            .build());

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("AllowFullAccesstoS3")
                .effect("Allow")
                .actions("s3:*")
                .resources("*")
                .build())
            .build());

        var exampleRolePolicy = new RolePolicy("exampleRolePolicy", RolePolicyArgs.builder()        
            .role(exampleRole.id())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  exampleSshKey:
    type: aws:transfer:SshKey
    properties:
      serverId: ${exampleServer.id}
      userName: ${exampleUser.userName}
      body: '... SSH key ...'
  exampleServer:
    type: aws:transfer:Server
    properties:
      identityProviderType: SERVICE_MANAGED
      tags:
        NAME: tf-acc-test-transfer-server
  exampleUser:
    type: aws:transfer:User
    properties:
      serverId: ${exampleServer.id}
      userName: tftestuser
      role: ${exampleRole.arn}
      tags:
        NAME: tftestuser
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  exampleRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${exampleRole.id}
      policy: ${examplePolicyDocument.json}
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - transfer.amazonaws.com
            actions:
              - sts:AssumeRole
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: AllowFullAccesstoS3
            effect: Allow
            actions:
              - s3:*
            resources:
              - '*'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transfer SSH Public Key using the `server_id` and `user_name` and `ssh_public_key_id` separated by `/`. For example:

```sh
 $ pulumi import aws:transfer/sshKey:SshKey bar s-12345678/test-username/key-12345
```
 