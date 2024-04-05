Provides a Lambda Code Signing Config resource. A code signing configuration defines a list of allowed signing profiles and defines the code-signing validation policy (action to be taken if deployment validation checks fail).

For information about Lambda code signing configurations and how to use them, see [configuring code signing for Lambda functions](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html)

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const newCsc = new aws.lambda.CodeSigningConfig("newCsc", {
    allowedPublishers: {
        signingProfileVersionArns: [
            aws_signer_signing_profile.example1.arn,
            aws_signer_signing_profile.example2.arn,
        ],
    },
    policies: {
        untrustedArtifactOnDeployment: "Warn",
    },
    description: "My awesome code signing config.",
});
```
```python
import pulumi
import pulumi_aws as aws

new_csc = aws.lambda_.CodeSigningConfig("newCsc",
    allowed_publishers=aws.lambda_.CodeSigningConfigAllowedPublishersArgs(
        signing_profile_version_arns=[
            aws_signer_signing_profile["example1"]["arn"],
            aws_signer_signing_profile["example2"]["arn"],
        ],
    ),
    policies=aws.lambda_.CodeSigningConfigPoliciesArgs(
        untrusted_artifact_on_deployment="Warn",
    ),
    description="My awesome code signing config.")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var newCsc = new Aws.Lambda.CodeSigningConfig("newCsc", new()
    {
        AllowedPublishers = new Aws.Lambda.Inputs.CodeSigningConfigAllowedPublishersArgs
        {
            SigningProfileVersionArns = new[]
            {
                aws_signer_signing_profile.Example1.Arn,
                aws_signer_signing_profile.Example2.Arn,
            },
        },
        Policies = new Aws.Lambda.Inputs.CodeSigningConfigPoliciesArgs
        {
            UntrustedArtifactOnDeployment = "Warn",
        },
        Description = "My awesome code signing config.",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewCodeSigningConfig(ctx, "newCsc", &lambda.CodeSigningConfigArgs{
			AllowedPublishers: &lambda.CodeSigningConfigAllowedPublishersArgs{
				SigningProfileVersionArns: pulumi.StringArray{
					aws_signer_signing_profile.Example1.Arn,
					aws_signer_signing_profile.Example2.Arn,
				},
			},
			Policies: &lambda.CodeSigningConfigPoliciesArgs{
				UntrustedArtifactOnDeployment: pulumi.String("Warn"),
			},
			Description: pulumi.String("My awesome code signing config."),
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
import com.pulumi.aws.lambda.CodeSigningConfig;
import com.pulumi.aws.lambda.CodeSigningConfigArgs;
import com.pulumi.aws.lambda.inputs.CodeSigningConfigAllowedPublishersArgs;
import com.pulumi.aws.lambda.inputs.CodeSigningConfigPoliciesArgs;
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
        var newCsc = new CodeSigningConfig("newCsc", CodeSigningConfigArgs.builder()        
            .allowedPublishers(CodeSigningConfigAllowedPublishersArgs.builder()
                .signingProfileVersionArns(                
                    aws_signer_signing_profile.example1().arn(),
                    aws_signer_signing_profile.example2().arn())
                .build())
            .policies(CodeSigningConfigPoliciesArgs.builder()
                .untrustedArtifactOnDeployment("Warn")
                .build())
            .description("My awesome code signing config.")
            .build());

    }
}
```
```yaml
resources:
  newCsc:
    type: aws:lambda:CodeSigningConfig
    properties:
      allowedPublishers:
        signingProfileVersionArns:
          - ${aws_signer_signing_profile.example1.arn}
          - ${aws_signer_signing_profile.example2.arn}
      policies:
        untrustedArtifactOnDeployment: Warn
      description: My awesome code signing config.
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Code Signing Configs using their ARN. For example:

```sh
 $ pulumi import aws:lambda/codeSigningConfig:CodeSigningConfig imported_csc arn:aws:lambda:us-west-2:123456789012:code-signing-config:csc-0f6c334abcdea4d8b
```
 