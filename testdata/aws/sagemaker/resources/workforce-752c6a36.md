Provides a SageMaker Workforce resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Cognito Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleUserPool = new aws.cognito.UserPool("exampleUserPool", {});
const exampleUserPoolClient = new aws.cognito.UserPoolClient("exampleUserPoolClient", {
    generateSecret: true,
    userPoolId: exampleUserPool.id,
});
const exampleUserPoolDomain = new aws.cognito.UserPoolDomain("exampleUserPoolDomain", {
    domain: "example",
    userPoolId: exampleUserPool.id,
});
const exampleWorkforce = new aws.sagemaker.Workforce("exampleWorkforce", {
    workforceName: "example",
    cognitoConfig: {
        clientId: exampleUserPoolClient.id,
        userPool: exampleUserPoolDomain.userPoolId,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_user_pool = aws.cognito.UserPool("exampleUserPool")
example_user_pool_client = aws.cognito.UserPoolClient("exampleUserPoolClient",
    generate_secret=True,
    user_pool_id=example_user_pool.id)
example_user_pool_domain = aws.cognito.UserPoolDomain("exampleUserPoolDomain",
    domain="example",
    user_pool_id=example_user_pool.id)
example_workforce = aws.sagemaker.Workforce("exampleWorkforce",
    workforce_name="example",
    cognito_config=aws.sagemaker.WorkforceCognitoConfigArgs(
        client_id=example_user_pool_client.id,
        user_pool=example_user_pool_domain.user_pool_id,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleUserPool = new Aws.Cognito.UserPool("exampleUserPool");

    var exampleUserPoolClient = new Aws.Cognito.UserPoolClient("exampleUserPoolClient", new()
    {
        GenerateSecret = true,
        UserPoolId = exampleUserPool.Id,
    });

    var exampleUserPoolDomain = new Aws.Cognito.UserPoolDomain("exampleUserPoolDomain", new()
    {
        Domain = "example",
        UserPoolId = exampleUserPool.Id,
    });

    var exampleWorkforce = new Aws.Sagemaker.Workforce("exampleWorkforce", new()
    {
        WorkforceName = "example",
        CognitoConfig = new Aws.Sagemaker.Inputs.WorkforceCognitoConfigArgs
        {
            ClientId = exampleUserPoolClient.Id,
            UserPool = exampleUserPoolDomain.UserPoolId,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleUserPool, err := cognito.NewUserPool(ctx, "exampleUserPool", nil)
		if err != nil {
			return err
		}
		exampleUserPoolClient, err := cognito.NewUserPoolClient(ctx, "exampleUserPoolClient", &cognito.UserPoolClientArgs{
			GenerateSecret: pulumi.Bool(true),
			UserPoolId:     exampleUserPool.ID(),
		})
		if err != nil {
			return err
		}
		exampleUserPoolDomain, err := cognito.NewUserPoolDomain(ctx, "exampleUserPoolDomain", &cognito.UserPoolDomainArgs{
			Domain:     pulumi.String("example"),
			UserPoolId: exampleUserPool.ID(),
		})
		if err != nil {
			return err
		}
		_, err = sagemaker.NewWorkforce(ctx, "exampleWorkforce", &sagemaker.WorkforceArgs{
			WorkforceName: pulumi.String("example"),
			CognitoConfig: &sagemaker.WorkforceCognitoConfigArgs{
				ClientId: exampleUserPoolClient.ID(),
				UserPool: exampleUserPoolDomain.UserPoolId,
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
import com.pulumi.aws.cognito.UserPoolDomain;
import com.pulumi.aws.cognito.UserPoolDomainArgs;
import com.pulumi.aws.sagemaker.Workforce;
import com.pulumi.aws.sagemaker.WorkforceArgs;
import com.pulumi.aws.sagemaker.inputs.WorkforceCognitoConfigArgs;
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
        var exampleUserPool = new UserPool("exampleUserPool");

        var exampleUserPoolClient = new UserPoolClient("exampleUserPoolClient", UserPoolClientArgs.builder()        
            .generateSecret(true)
            .userPoolId(exampleUserPool.id())
            .build());

        var exampleUserPoolDomain = new UserPoolDomain("exampleUserPoolDomain", UserPoolDomainArgs.builder()        
            .domain("example")
            .userPoolId(exampleUserPool.id())
            .build());

        var exampleWorkforce = new Workforce("exampleWorkforce", WorkforceArgs.builder()        
            .workforceName("example")
            .cognitoConfig(WorkforceCognitoConfigArgs.builder()
                .clientId(exampleUserPoolClient.id())
                .userPool(exampleUserPoolDomain.userPoolId())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleWorkforce:
    type: aws:sagemaker:Workforce
    properties:
      workforceName: example
      cognitoConfig:
        clientId: ${exampleUserPoolClient.id}
        userPool: ${exampleUserPoolDomain.userPoolId}
  exampleUserPool:
    type: aws:cognito:UserPool
  exampleUserPoolClient:
    type: aws:cognito:UserPoolClient
    properties:
      generateSecret: true
      userPoolId: ${exampleUserPool.id}
  exampleUserPoolDomain:
    type: aws:cognito:UserPoolDomain
    properties:
      domain: example
      userPoolId: ${exampleUserPool.id}
```
{{% /example %}}
{{% example %}}
### Oidc Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.Workforce("example", {
    oidcConfig: {
        authorizationEndpoint: "https://example.com",
        clientId: "example",
        clientSecret: "example",
        issuer: "https://example.com",
        jwksUri: "https://example.com",
        logoutEndpoint: "https://example.com",
        tokenEndpoint: "https://example.com",
        userInfoEndpoint: "https://example.com",
    },
    workforceName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.Workforce("example",
    oidc_config=aws.sagemaker.WorkforceOidcConfigArgs(
        authorization_endpoint="https://example.com",
        client_id="example",
        client_secret="example",
        issuer="https://example.com",
        jwks_uri="https://example.com",
        logout_endpoint="https://example.com",
        token_endpoint="https://example.com",
        user_info_endpoint="https://example.com",
    ),
    workforce_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.Workforce("example", new()
    {
        OidcConfig = new Aws.Sagemaker.Inputs.WorkforceOidcConfigArgs
        {
            AuthorizationEndpoint = "https://example.com",
            ClientId = "example",
            ClientSecret = "example",
            Issuer = "https://example.com",
            JwksUri = "https://example.com",
            LogoutEndpoint = "https://example.com",
            TokenEndpoint = "https://example.com",
            UserInfoEndpoint = "https://example.com",
        },
        WorkforceName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewWorkforce(ctx, "example", &sagemaker.WorkforceArgs{
			OidcConfig: &sagemaker.WorkforceOidcConfigArgs{
				AuthorizationEndpoint: pulumi.String("https://example.com"),
				ClientId:              pulumi.String("example"),
				ClientSecret:          pulumi.String("example"),
				Issuer:                pulumi.String("https://example.com"),
				JwksUri:               pulumi.String("https://example.com"),
				LogoutEndpoint:        pulumi.String("https://example.com"),
				TokenEndpoint:         pulumi.String("https://example.com"),
				UserInfoEndpoint:      pulumi.String("https://example.com"),
			},
			WorkforceName: pulumi.String("example"),
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
import com.pulumi.aws.sagemaker.Workforce;
import com.pulumi.aws.sagemaker.WorkforceArgs;
import com.pulumi.aws.sagemaker.inputs.WorkforceOidcConfigArgs;
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
        var example = new Workforce("example", WorkforceArgs.builder()        
            .oidcConfig(WorkforceOidcConfigArgs.builder()
                .authorizationEndpoint("https://example.com")
                .clientId("example")
                .clientSecret("example")
                .issuer("https://example.com")
                .jwksUri("https://example.com")
                .logoutEndpoint("https://example.com")
                .tokenEndpoint("https://example.com")
                .userInfoEndpoint("https://example.com")
                .build())
            .workforceName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:Workforce
    properties:
      oidcConfig:
        authorizationEndpoint: https://example.com
        clientId: example
        clientSecret: example
        issuer: https://example.com
        jwksUri: https://example.com
        logoutEndpoint: https://example.com
        tokenEndpoint: https://example.com
        userInfoEndpoint: https://example.com
      workforceName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Workforces using the `workforce_name`. For example:

```sh
 $ pulumi import aws:sagemaker/workforce:Workforce example example
```
 