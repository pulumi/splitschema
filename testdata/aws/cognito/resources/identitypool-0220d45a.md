Provides an AWS Cognito Identity Pool.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const _default = new aws.iam.SamlProvider("default", {samlMetadataDocument: fs.readFileSync("saml-metadata.xml", "utf8")});
const main = new aws.cognito.IdentityPool("main", {
    identityPoolName: "identity pool",
    allowUnauthenticatedIdentities: false,
    allowClassicFlow: false,
    cognitoIdentityProviders: [
        {
            clientId: "6lhlkkfbfb4q5kpp90urffae",
            providerName: "cognito-idp.us-east-1.amazonaws.com/us-east-1_Tv0493apJ",
            serverSideTokenCheck: false,
        },
        {
            clientId: "7kodkvfqfb4qfkp39eurffae",
            providerName: "cognito-idp.us-east-1.amazonaws.com/eu-west-1_Zr231apJu",
            serverSideTokenCheck: false,
        },
    ],
    supportedLoginProviders: {
        "graph.facebook.com": "7346241598935552",
        "accounts.google.com": "123456789012.apps.googleusercontent.com",
    },
    samlProviderArns: [_default.arn],
    openidConnectProviderArns: ["arn:aws:iam::123456789012:oidc-provider/id.example.com"],
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.iam.SamlProvider("default", saml_metadata_document=(lambda path: open(path).read())("saml-metadata.xml"))
main = aws.cognito.IdentityPool("main",
    identity_pool_name="identity pool",
    allow_unauthenticated_identities=False,
    allow_classic_flow=False,
    cognito_identity_providers=[
        aws.cognito.IdentityPoolCognitoIdentityProviderArgs(
            client_id="6lhlkkfbfb4q5kpp90urffae",
            provider_name="cognito-idp.us-east-1.amazonaws.com/us-east-1_Tv0493apJ",
            server_side_token_check=False,
        ),
        aws.cognito.IdentityPoolCognitoIdentityProviderArgs(
            client_id="7kodkvfqfb4qfkp39eurffae",
            provider_name="cognito-idp.us-east-1.amazonaws.com/eu-west-1_Zr231apJu",
            server_side_token_check=False,
        ),
    ],
    supported_login_providers={
        "graph.facebook.com": "7346241598935552",
        "accounts.google.com": "123456789012.apps.googleusercontent.com",
    },
    saml_provider_arns=[default.arn],
    openid_connect_provider_arns=["arn:aws:iam::123456789012:oidc-provider/id.example.com"])
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Iam.SamlProvider("default", new()
    {
        SamlMetadataDocument = File.ReadAllText("saml-metadata.xml"),
    });

    var main = new Aws.Cognito.IdentityPool("main", new()
    {
        IdentityPoolName = "identity pool",
        AllowUnauthenticatedIdentities = false,
        AllowClassicFlow = false,
        CognitoIdentityProviders = new[]
        {
            new Aws.Cognito.Inputs.IdentityPoolCognitoIdentityProviderArgs
            {
                ClientId = "6lhlkkfbfb4q5kpp90urffae",
                ProviderName = "cognito-idp.us-east-1.amazonaws.com/us-east-1_Tv0493apJ",
                ServerSideTokenCheck = false,
            },
            new Aws.Cognito.Inputs.IdentityPoolCognitoIdentityProviderArgs
            {
                ClientId = "7kodkvfqfb4qfkp39eurffae",
                ProviderName = "cognito-idp.us-east-1.amazonaws.com/eu-west-1_Zr231apJu",
                ServerSideTokenCheck = false,
            },
        },
        SupportedLoginProviders = 
        {
            { "graph.facebook.com", "7346241598935552" },
            { "accounts.google.com", "123456789012.apps.googleusercontent.com" },
        },
        SamlProviderArns = new[]
        {
            @default.Arn,
        },
        OpenidConnectProviderArns = new[]
        {
            "arn:aws:iam::123456789012:oidc-provider/id.example.com",
        },
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewSamlProvider(ctx, "default", &iam.SamlProviderArgs{
			SamlMetadataDocument: readFileOrPanic("saml-metadata.xml"),
		})
		if err != nil {
			return err
		}
		_, err = cognito.NewIdentityPool(ctx, "main", &cognito.IdentityPoolArgs{
			IdentityPoolName:               pulumi.String("identity pool"),
			AllowUnauthenticatedIdentities: pulumi.Bool(false),
			AllowClassicFlow:               pulumi.Bool(false),
			CognitoIdentityProviders: cognito.IdentityPoolCognitoIdentityProviderArray{
				&cognito.IdentityPoolCognitoIdentityProviderArgs{
					ClientId:             pulumi.String("6lhlkkfbfb4q5kpp90urffae"),
					ProviderName:         pulumi.String("cognito-idp.us-east-1.amazonaws.com/us-east-1_Tv0493apJ"),
					ServerSideTokenCheck: pulumi.Bool(false),
				},
				&cognito.IdentityPoolCognitoIdentityProviderArgs{
					ClientId:             pulumi.String("7kodkvfqfb4qfkp39eurffae"),
					ProviderName:         pulumi.String("cognito-idp.us-east-1.amazonaws.com/eu-west-1_Zr231apJu"),
					ServerSideTokenCheck: pulumi.Bool(false),
				},
			},
			SupportedLoginProviders: pulumi.StringMap{
				"graph.facebook.com":  pulumi.String("7346241598935552"),
				"accounts.google.com": pulumi.String("123456789012.apps.googleusercontent.com"),
			},
			SamlProviderArns: pulumi.StringArray{
				_default.Arn,
			},
			OpenidConnectProviderArns: pulumi.StringArray{
				pulumi.String("arn:aws:iam::123456789012:oidc-provider/id.example.com"),
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
import com.pulumi.aws.iam.SamlProvider;
import com.pulumi.aws.iam.SamlProviderArgs;
import com.pulumi.aws.cognito.IdentityPool;
import com.pulumi.aws.cognito.IdentityPoolArgs;
import com.pulumi.aws.cognito.inputs.IdentityPoolCognitoIdentityProviderArgs;
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
        var default_ = new SamlProvider("default", SamlProviderArgs.builder()        
            .samlMetadataDocument(Files.readString(Paths.get("saml-metadata.xml")))
            .build());

        var main = new IdentityPool("main", IdentityPoolArgs.builder()        
            .identityPoolName("identity pool")
            .allowUnauthenticatedIdentities(false)
            .allowClassicFlow(false)
            .cognitoIdentityProviders(            
                IdentityPoolCognitoIdentityProviderArgs.builder()
                    .clientId("6lhlkkfbfb4q5kpp90urffae")
                    .providerName("cognito-idp.us-east-1.amazonaws.com/us-east-1_Tv0493apJ")
                    .serverSideTokenCheck(false)
                    .build(),
                IdentityPoolCognitoIdentityProviderArgs.builder()
                    .clientId("7kodkvfqfb4qfkp39eurffae")
                    .providerName("cognito-idp.us-east-1.amazonaws.com/eu-west-1_Zr231apJu")
                    .serverSideTokenCheck(false)
                    .build())
            .supportedLoginProviders(Map.ofEntries(
                Map.entry("graph.facebook.com", "7346241598935552"),
                Map.entry("accounts.google.com", "123456789012.apps.googleusercontent.com")
            ))
            .samlProviderArns(default_.arn())
            .openidConnectProviderArns("arn:aws:iam::123456789012:oidc-provider/id.example.com")
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:iam:SamlProvider
    properties:
      samlMetadataDocument:
        fn::readFile: saml-metadata.xml
  main:
    type: aws:cognito:IdentityPool
    properties:
      identityPoolName: identity pool
      allowUnauthenticatedIdentities: false
      allowClassicFlow: false
      cognitoIdentityProviders:
        - clientId: 6lhlkkfbfb4q5kpp90urffae
          providerName: cognito-idp.us-east-1.amazonaws.com/us-east-1_Tv0493apJ
          serverSideTokenCheck: false
        - clientId: 7kodkvfqfb4qfkp39eurffae
          providerName: cognito-idp.us-east-1.amazonaws.com/eu-west-1_Zr231apJu
          serverSideTokenCheck: false
      supportedLoginProviders:
        graph.facebook.com: '7346241598935552'
        accounts.google.com: 123456789012.apps.googleusercontent.com
      samlProviderArns:
        - ${default.arn}
      openidConnectProviderArns:
        - arn:aws:iam::123456789012:oidc-provider/id.example.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cognito Identity Pool using its ID. For example:

```sh
 $ pulumi import aws:cognito/identityPool:IdentityPool mypool us-west-2:1a234567-8901-234b-5cde-f6789g01h2i3
```
 