Provides a Cognito User Pool Client resource.

To manage a User Pool Client created by another service, such as when [configuring an OpenSearch Domain to use Cognito authentication](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html),
use the `aws.cognito.ManagedUserPoolClient` resource instead.

{{% examples %}}
## Example Usage
{{% example %}}
### Create a basic user pool client

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const pool = new aws.cognito.UserPool("pool", {});
const client = new aws.cognito.UserPoolClient("client", {userPoolId: pool.id});
```
```python
import pulumi
import pulumi_aws as aws

pool = aws.cognito.UserPool("pool")
client = aws.cognito.UserPoolClient("client", user_pool_id=pool.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var pool = new Aws.Cognito.UserPool("pool");

    var client = new Aws.Cognito.UserPoolClient("client", new()
    {
        UserPoolId = pool.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		pool, err := cognito.NewUserPool(ctx, "pool", nil)
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPoolClient(ctx, "client", &cognito.UserPoolClientArgs{
			UserPoolId: pool.ID(),
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.UserPoolClient;
import com.pulumi.aws.cognito.UserPoolClientArgs;
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
        var pool = new UserPool("pool");

        var client = new UserPoolClient("client", UserPoolClientArgs.builder()        
            .userPoolId(pool.id())
            .build());

    }
}
```
```yaml
resources:
  client:
    type: aws:cognito:UserPoolClient
    properties:
      userPoolId: ${pool.id}
  pool:
    type: aws:cognito:UserPool
```
{{% /example %}}
{{% example %}}
### Create a user pool client with no SRP authentication

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const pool = new aws.cognito.UserPool("pool", {});
const client = new aws.cognito.UserPoolClient("client", {
    userPoolId: pool.id,
    generateSecret: true,
    explicitAuthFlows: ["ADMIN_NO_SRP_AUTH"],
});
```
```python
import pulumi
import pulumi_aws as aws

pool = aws.cognito.UserPool("pool")
client = aws.cognito.UserPoolClient("client",
    user_pool_id=pool.id,
    generate_secret=True,
    explicit_auth_flows=["ADMIN_NO_SRP_AUTH"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var pool = new Aws.Cognito.UserPool("pool");

    var client = new Aws.Cognito.UserPoolClient("client", new()
    {
        UserPoolId = pool.Id,
        GenerateSecret = true,
        ExplicitAuthFlows = new[]
        {
            "ADMIN_NO_SRP_AUTH",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		pool, err := cognito.NewUserPool(ctx, "pool", nil)
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPoolClient(ctx, "client", &cognito.UserPoolClientArgs{
			UserPoolId:     pool.ID(),
			GenerateSecret: pulumi.Bool(true),
			ExplicitAuthFlows: pulumi.StringArray{
				pulumi.String("ADMIN_NO_SRP_AUTH"),
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.UserPoolClient;
import com.pulumi.aws.cognito.UserPoolClientArgs;
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
        var pool = new UserPool("pool");

        var client = new UserPoolClient("client", UserPoolClientArgs.builder()        
            .userPoolId(pool.id())
            .generateSecret(true)
            .explicitAuthFlows("ADMIN_NO_SRP_AUTH")
            .build());

    }
}
```
```yaml
resources:
  client:
    type: aws:cognito:UserPoolClient
    properties:
      userPoolId: ${pool.id}
      generateSecret: true
      explicitAuthFlows:
        - ADMIN_NO_SRP_AUTH
  pool:
    type: aws:cognito:UserPool
```
{{% /example %}}
{{% example %}}
### Create a user pool client with pinpoint analytics

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testUserPool = new aws.cognito.UserPool("testUserPool", {});
const testApp = new aws.pinpoint.App("testApp", {});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["cognito-idp.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const testRole = new aws.iam.Role("testRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const testUserPoolClient = new aws.cognito.UserPoolClient("testUserPoolClient", {
    userPoolId: testUserPool.id,
    analyticsConfiguration: {
        applicationId: testApp.applicationId,
        externalId: "some_id",
        roleArn: testRole.arn,
        userDataShared: true,
    },
});
const current = aws.getCallerIdentity({});
const testPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        actions: [
            "mobiletargeting:UpdateEndpoint",
            "mobiletargeting:PutEvents",
        ],
        resources: [pulumi.all([current, testApp.applicationId]).apply(([current, applicationId]) => `arn:aws:mobiletargeting:*:${current.accountId}:apps/${applicationId}*`)],
    }],
});
const testRolePolicy = new aws.iam.RolePolicy("testRolePolicy", {
    role: testRole.id,
    policy: testPolicyDocument.apply(testPolicyDocument => testPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

test_user_pool = aws.cognito.UserPool("testUserPool")
test_app = aws.pinpoint.App("testApp")
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["cognito-idp.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
test_role = aws.iam.Role("testRole", assume_role_policy=assume_role.json)
test_user_pool_client = aws.cognito.UserPoolClient("testUserPoolClient",
    user_pool_id=test_user_pool.id,
    analytics_configuration=aws.cognito.UserPoolClientAnalyticsConfigurationArgs(
        application_id=test_app.application_id,
        external_id="some_id",
        role_arn=test_role.arn,
        user_data_shared=True,
    ))
current = aws.get_caller_identity()
test_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=[
        "mobiletargeting:UpdateEndpoint",
        "mobiletargeting:PutEvents",
    ],
    resources=[test_app.application_id.apply(lambda application_id: f"arn:aws:mobiletargeting:*:{current.account_id}:apps/{application_id}*")],
)])
test_role_policy = aws.iam.RolePolicy("testRolePolicy",
    role=test_role.id,
    policy=test_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testUserPool = new Aws.Cognito.UserPool("testUserPool");

    var testApp = new Aws.Pinpoint.App("testApp");

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
                            "cognito-idp.amazonaws.com",
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

    var testRole = new Aws.Iam.Role("testRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var testUserPoolClient = new Aws.Cognito.UserPoolClient("testUserPoolClient", new()
    {
        UserPoolId = testUserPool.Id,
        AnalyticsConfiguration = new Aws.Cognito.Inputs.UserPoolClientAnalyticsConfigurationArgs
        {
            ApplicationId = testApp.ApplicationId,
            ExternalId = "some_id",
            RoleArn = testRole.Arn,
            UserDataShared = true,
        },
    });

    var current = Aws.GetCallerIdentity.Invoke();

    var testPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "mobiletargeting:UpdateEndpoint",
                    "mobiletargeting:PutEvents",
                },
                Resources = new[]
                {
                    $"arn:aws:mobiletargeting:*:{current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:apps/{testApp.ApplicationId}*",
                },
            },
        },
    });

    var testRolePolicy = new Aws.Iam.RolePolicy("testRolePolicy", new()
    {
        Role = testRole.Id,
        Policy = testPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/pinpoint"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testUserPool, err := cognito.NewUserPool(ctx, "testUserPool", nil)
		if err != nil {
			return err
		}
		testApp, err := pinpoint.NewApp(ctx, "testApp", nil)
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
								"cognito-idp.amazonaws.com",
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
		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPoolClient(ctx, "testUserPoolClient", &cognito.UserPoolClientArgs{
			UserPoolId: testUserPool.ID(),
			AnalyticsConfiguration: &cognito.UserPoolClientAnalyticsConfigurationArgs{
				ApplicationId:  testApp.ApplicationId,
				ExternalId:     pulumi.String("some_id"),
				RoleArn:        testRole.Arn,
				UserDataShared: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		testPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("mobiletargeting:UpdateEndpoint"),
						pulumi.String("mobiletargeting:PutEvents"),
					},
					Resources: pulumi.StringArray{
						testApp.ApplicationId.ApplyT(func(applicationId string) (string, error) {
							return fmt.Sprintf("arn:aws:mobiletargeting:*:%v:apps/%v*", current.AccountId, applicationId), nil
						}).(pulumi.StringOutput),
					},
				},
			},
		}, nil)
		_, err = iam.NewRolePolicy(ctx, "testRolePolicy", &iam.RolePolicyArgs{
			Role: testRole.ID(),
			Policy: testPolicyDocument.ApplyT(func(testPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &testPolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.pinpoint.App;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.cognito.UserPoolClient;
import com.pulumi.aws.cognito.UserPoolClientArgs;
import com.pulumi.aws.cognito.inputs.UserPoolClientAnalyticsConfigurationArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
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
        var testUserPool = new UserPool("testUserPool");

        var testApp = new App("testApp");

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("cognito-idp.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var testRole = new Role("testRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var testUserPoolClient = new UserPoolClient("testUserPoolClient", UserPoolClientArgs.builder()        
            .userPoolId(testUserPool.id())
            .analyticsConfiguration(UserPoolClientAnalyticsConfigurationArgs.builder()
                .applicationId(testApp.applicationId())
                .externalId("some_id")
                .roleArn(testRole.arn())
                .userDataShared(true)
                .build())
            .build());

        final var current = AwsFunctions.getCallerIdentity();

        final var testPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions(                
                    "mobiletargeting:UpdateEndpoint",
                    "mobiletargeting:PutEvents")
                .resources(testApp.applicationId().applyValue(applicationId -> String.format("arn:aws:mobiletargeting:*:%s:apps/%s*", current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()),applicationId)))
                .build())
            .build());

        var testRolePolicy = new RolePolicy("testRolePolicy", RolePolicyArgs.builder()        
            .role(testRole.id())
            .policy(testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(testPolicyDocument -> testPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  testUserPoolClient:
    type: aws:cognito:UserPoolClient
    properties:
      userPoolId: ${testUserPool.id}
      analyticsConfiguration:
        applicationId: ${testApp.applicationId}
        externalId: some_id
        roleArn: ${testRole.arn}
        userDataShared: true
  testUserPool:
    type: aws:cognito:UserPool
  testApp:
    type: aws:pinpoint:App
  testRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  testRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${testRole.id}
      policy: ${testPolicyDocument.json}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - cognito-idp.amazonaws.com
            actions:
              - sts:AssumeRole
  testPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - mobiletargeting:UpdateEndpoint
              - mobiletargeting:PutEvents
            resources:
              - arn:aws:mobiletargeting:*:${current.accountId}:apps/${testApp.applicationId}*
```
{{% /example %}}
{{% example %}}
### Create a user pool client with Cognito as the identity provider

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const pool = new aws.cognito.UserPool("pool", {});
const userpoolClient = new aws.cognito.UserPoolClient("userpoolClient", {
    userPoolId: pool.id,
    callbackUrls: ["https://example.com"],
    allowedOauthFlowsUserPoolClient: true,
    allowedOauthFlows: [
        "code",
        "implicit",
    ],
    allowedOauthScopes: [
        "email",
        "openid",
    ],
    supportedIdentityProviders: ["COGNITO"],
});
```
```python
import pulumi
import pulumi_aws as aws

pool = aws.cognito.UserPool("pool")
userpool_client = aws.cognito.UserPoolClient("userpoolClient",
    user_pool_id=pool.id,
    callback_urls=["https://example.com"],
    allowed_oauth_flows_user_pool_client=True,
    allowed_oauth_flows=[
        "code",
        "implicit",
    ],
    allowed_oauth_scopes=[
        "email",
        "openid",
    ],
    supported_identity_providers=["COGNITO"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var pool = new Aws.Cognito.UserPool("pool");

    var userpoolClient = new Aws.Cognito.UserPoolClient("userpoolClient", new()
    {
        UserPoolId = pool.Id,
        CallbackUrls = new[]
        {
            "https://example.com",
        },
        AllowedOauthFlowsUserPoolClient = true,
        AllowedOauthFlows = new[]
        {
            "code",
            "implicit",
        },
        AllowedOauthScopes = new[]
        {
            "email",
            "openid",
        },
        SupportedIdentityProviders = new[]
        {
            "COGNITO",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		pool, err := cognito.NewUserPool(ctx, "pool", nil)
		if err != nil {
			return err
		}
		_, err = cognito.NewUserPoolClient(ctx, "userpoolClient", &cognito.UserPoolClientArgs{
			UserPoolId: pool.ID(),
			CallbackUrls: pulumi.StringArray{
				pulumi.String("https://example.com"),
			},
			AllowedOauthFlowsUserPoolClient: pulumi.Bool(true),
			AllowedOauthFlows: pulumi.StringArray{
				pulumi.String("code"),
				pulumi.String("implicit"),
			},
			AllowedOauthScopes: pulumi.StringArray{
				pulumi.String("email"),
				pulumi.String("openid"),
			},
			SupportedIdentityProviders: pulumi.StringArray{
				pulumi.String("COGNITO"),
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.UserPoolClient;
import com.pulumi.aws.cognito.UserPoolClientArgs;
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
        var pool = new UserPool("pool");

        var userpoolClient = new UserPoolClient("userpoolClient", UserPoolClientArgs.builder()        
            .userPoolId(pool.id())
            .callbackUrls("https://example.com")
            .allowedOauthFlowsUserPoolClient(true)
            .allowedOauthFlows(            
                "code",
                "implicit")
            .allowedOauthScopes(            
                "email",
                "openid")
            .supportedIdentityProviders("COGNITO")
            .build());

    }
}
```
```yaml
resources:
  userpoolClient:
    type: aws:cognito:UserPoolClient
    properties:
      userPoolId: ${pool.id}
      callbackUrls:
        - https://example.com
      allowedOauthFlowsUserPoolClient: true
      allowedOauthFlows:
        - code
        - implicit
      allowedOauthScopes:
        - email
        - openid
      supportedIdentityProviders:
        - COGNITO
  pool:
    type: aws:cognito:UserPool
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cognito User Pool Clients using the `id` of the Cognito User Pool, and the `id` of the Cognito User Pool Client. For example:

```sh
 $ pulumi import aws:cognito/userPoolClient:UserPoolClient client us-west-2_abc123/3ho4ek12345678909nh3fmhpko
```
 