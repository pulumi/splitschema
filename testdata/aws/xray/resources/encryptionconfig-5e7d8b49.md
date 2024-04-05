Creates and manages an AWS XRay Encryption Config.

> **NOTE:** Removing this resource from the provider has no effect to the encryption configuration within X-Ray.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.xray.EncryptionConfig("example", {type: "NONE"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.xray.EncryptionConfig("example", type="NONE")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Xray.EncryptionConfig("example", new()
    {
        Type = "NONE",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/xray"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := xray.NewEncryptionConfig(ctx, "example", &xray.EncryptionConfigArgs{
			Type: pulumi.String("NONE"),
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
import com.pulumi.aws.xray.EncryptionConfig;
import com.pulumi.aws.xray.EncryptionConfigArgs;
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
        var example = new EncryptionConfig("example", EncryptionConfigArgs.builder()        
            .type("NONE")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:xray:EncryptionConfig
    properties:
      type: NONE
```

{{% /example %}}
{{% example %}}
### With KMS Key

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const examplePolicyDocument = current.then(current => aws.iam.getPolicyDocument({
    statements: [{
        sid: "Enable IAM User Permissions",
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: [`arn:aws:iam::${current.accountId}:root`],
        }],
        actions: ["kms:*"],
        resources: ["*"],
    }],
}));
const exampleKey = new aws.kms.Key("exampleKey", {
    description: "Some Key",
    deletionWindowInDays: 7,
    policy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
const exampleEncryptionConfig = new aws.xray.EncryptionConfig("exampleEncryptionConfig", {
    type: "KMS",
    keyId: exampleKey.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="Enable IAM User Permissions",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=[f"arn:aws:iam::{current.account_id}:root"],
    )],
    actions=["kms:*"],
    resources=["*"],
)])
example_key = aws.kms.Key("exampleKey",
    description="Some Key",
    deletion_window_in_days=7,
    policy=example_policy_document.json)
example_encryption_config = aws.xray.EncryptionConfig("exampleEncryptionConfig",
    type="KMS",
    key_id=example_key.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "Enable IAM User Permissions",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            $"arn:aws:iam::{current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:root",
                        },
                    },
                },
                Actions = new[]
                {
                    "kms:*",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "Some Key",
        DeletionWindowInDays = 7,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleEncryptionConfig = new Aws.Xray.EncryptionConfig("exampleEncryptionConfig", new()
    {
        Type = "KMS",
        KeyId = exampleKey.Arn,
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/xray"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("Enable IAM User Permissions"),
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "AWS",
							Identifiers: []string{
								fmt.Sprintf("arn:aws:iam::%v:root", current.AccountId),
							},
						},
					},
					Actions: []string{
						"kms:*",
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
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description:          pulumi.String("Some Key"),
			DeletionWindowInDays: pulumi.Int(7),
			Policy:               *pulumi.String(examplePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = xray.NewEncryptionConfig(ctx, "exampleEncryptionConfig", &xray.EncryptionConfigArgs{
			Type:  pulumi.String("KMS"),
			KeyId: exampleKey.Arn,
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.xray.EncryptionConfig;
import com.pulumi.aws.xray.EncryptionConfigArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("Enable IAM User Permissions")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers(String.format("arn:aws:iam::%s:root", current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                    .build())
                .actions("kms:*")
                .resources("*")
                .build())
            .build());

        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("Some Key")
            .deletionWindowInDays(7)
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleEncryptionConfig = new EncryptionConfig("exampleEncryptionConfig", EncryptionConfigArgs.builder()        
            .type("KMS")
            .keyId(exampleKey.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: Some Key
      deletionWindowInDays: 7
      policy: ${examplePolicyDocument.json}
  exampleEncryptionConfig:
    type: aws:xray:EncryptionConfig
    properties:
      type: KMS
      keyId: ${exampleKey.arn}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: Enable IAM User Permissions
            effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - arn:aws:iam::${current.accountId}:root
            actions:
              - kms:*
            resources:
              - '*'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import XRay Encryption Config using the region name. For example:

```sh
 $ pulumi import aws:xray/encryptionConfig:EncryptionConfig example us-west-2
```
 