Provides a Cognito User Identity Provider resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cognito.UserPool("example", {autoVerifiedAttributes: ["email"]});
const exampleProvider = new aws.cognito.IdentityProvider("exampleProvider", {
    userPoolId: example.id,
    providerName: "Google",
    providerType: "Google",
    providerDetails: {
        authorize_scopes: "email",
        client_id: "your client_id",
        client_secret: "your client_secret",
    },
    attributeMapping: {
        email: "email",
        username: "sub",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cognito.UserPool("example", auto_verified_attributes=["email"])
example_provider = aws.cognito.IdentityProvider("exampleProvider",
    user_pool_id=example.id,
    provider_name="Google",
    provider_type="Google",
    provider_details={
        "authorize_scopes": "email",
        "client_id": "your client_id",
        "client_secret": "your client_secret",
    },
    attribute_mapping={
        "email": "email",
        "username": "sub",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Cognito.UserPool("example", new()
    {
        AutoVerifiedAttributes = new[]
        {
            "email",
        },
    });

    var exampleProvider = new Aws.Cognito.IdentityProvider("exampleProvider", new()
    {
        UserPoolId = example.Id,
        ProviderName = "Google",
        ProviderType = "Google",
        ProviderDetails = 
        {
            { "authorize_scopes", "email" },
            { "client_id", "your client_id" },
            { "client_secret", "your client_secret" },
        },
        AttributeMapping = 
        {
            { "email", "email" },
            { "username", "sub" },
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
		example, err := cognito.NewUserPool(ctx, "example", &cognito.UserPoolArgs{
			AutoVerifiedAttributes: pulumi.StringArray{
				pulumi.String("email"),
			},
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewIdentityProvider(ctx, "exampleProvider", &cognito.IdentityProviderArgs{
			UserPoolId:   example.ID(),
			ProviderName: pulumi.String("Google"),
			ProviderType: pulumi.String("Google"),
			ProviderDetails: pulumi.StringMap{
				"authorize_scopes": pulumi.String("email"),
				"client_id":        pulumi.String("your client_id"),
				"client_secret":    pulumi.String("your client_secret"),
			},
			AttributeMapping: pulumi.StringMap{
				"email":    pulumi.String("email"),
				"username": pulumi.String("sub"),
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
import com.pulumi.aws.cognito.UserPoolArgs;
import com.pulumi.aws.cognito.IdentityProvider;
import com.pulumi.aws.cognito.IdentityProviderArgs;
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
        var example = new UserPool("example", UserPoolArgs.builder()        
            .autoVerifiedAttributes("email")
            .build());

        var exampleProvider = new IdentityProvider("exampleProvider", IdentityProviderArgs.builder()        
            .userPoolId(example.id())
            .providerName("Google")
            .providerType("Google")
            .providerDetails(Map.ofEntries(
                Map.entry("authorize_scopes", "email"),
                Map.entry("client_id", "your client_id"),
                Map.entry("client_secret", "your client_secret")
            ))
            .attributeMapping(Map.ofEntries(
                Map.entry("email", "email"),
                Map.entry("username", "sub")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cognito:UserPool
    properties:
      autoVerifiedAttributes:
        - email
  exampleProvider:
    type: aws:cognito:IdentityProvider
    properties:
      userPoolId: ${example.id}
      providerName: Google
      providerType: Google
      providerDetails:
        authorize_scopes: email
        client_id: your client_id
        client_secret: your client_secret
      attributeMapping:
        email: email
        username: sub
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_cognito_identity_provider` resources using their User Pool ID and Provider Name. For example:

```sh
 $ pulumi import aws:cognito/identityProvider:IdentityProvider example us-west-2_abc123:CorpAD
```
 