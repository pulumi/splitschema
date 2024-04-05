Provides a SageMaker Code Repository resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.CodeRepository("example", {
    codeRepositoryName: "example",
    gitConfig: {
        repositoryUrl: "https://github.com/github/docs.git",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.CodeRepository("example",
    code_repository_name="example",
    git_config=aws.sagemaker.CodeRepositoryGitConfigArgs(
        repository_url="https://github.com/github/docs.git",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.CodeRepository("example", new()
    {
        CodeRepositoryName = "example",
        GitConfig = new Aws.Sagemaker.Inputs.CodeRepositoryGitConfigArgs
        {
            RepositoryUrl = "https://github.com/github/docs.git",
        },
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
		_, err := sagemaker.NewCodeRepository(ctx, "example", &sagemaker.CodeRepositoryArgs{
			CodeRepositoryName: pulumi.String("example"),
			GitConfig: &sagemaker.CodeRepositoryGitConfigArgs{
				RepositoryUrl: pulumi.String("https://github.com/github/docs.git"),
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
import com.pulumi.aws.sagemaker.CodeRepository;
import com.pulumi.aws.sagemaker.CodeRepositoryArgs;
import com.pulumi.aws.sagemaker.inputs.CodeRepositoryGitConfigArgs;
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
        var example = new CodeRepository("example", CodeRepositoryArgs.builder()        
            .codeRepositoryName("example")
            .gitConfig(CodeRepositoryGitConfigArgs.builder()
                .repositoryUrl("https://github.com/github/docs.git")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:CodeRepository
    properties:
      codeRepositoryName: example
      gitConfig:
        repositoryUrl: https://github.com/github/docs.git
```
{{% /example %}}
{{% example %}}
### Example with Secret

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleSecret = new aws.secretsmanager.Secret("exampleSecret", {});
const exampleSecretVersion = new aws.secretsmanager.SecretVersion("exampleSecretVersion", {
    secretId: exampleSecret.id,
    secretString: JSON.stringify({
        username: "example",
        password: "example",
    }),
});
const exampleCodeRepository = new aws.sagemaker.CodeRepository("exampleCodeRepository", {
    codeRepositoryName: "example",
    gitConfig: {
        repositoryUrl: "https://github.com/github/docs.git",
        secretArn: exampleSecret.arn,
    },
}, {
    dependsOn: [exampleSecretVersion],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_secret = aws.secretsmanager.Secret("exampleSecret")
example_secret_version = aws.secretsmanager.SecretVersion("exampleSecretVersion",
    secret_id=example_secret.id,
    secret_string=json.dumps({
        "username": "example",
        "password": "example",
    }))
example_code_repository = aws.sagemaker.CodeRepository("exampleCodeRepository",
    code_repository_name="example",
    git_config=aws.sagemaker.CodeRepositoryGitConfigArgs(
        repository_url="https://github.com/github/docs.git",
        secret_arn=example_secret.arn,
    ),
    opts=pulumi.ResourceOptions(depends_on=[example_secret_version]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleSecret = new Aws.SecretsManager.Secret("exampleSecret");

    var exampleSecretVersion = new Aws.SecretsManager.SecretVersion("exampleSecretVersion", new()
    {
        SecretId = exampleSecret.Id,
        SecretString = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["username"] = "example",
            ["password"] = "example",
        }),
    });

    var exampleCodeRepository = new Aws.Sagemaker.CodeRepository("exampleCodeRepository", new()
    {
        CodeRepositoryName = "example",
        GitConfig = new Aws.Sagemaker.Inputs.CodeRepositoryGitConfigArgs
        {
            RepositoryUrl = "https://github.com/github/docs.git",
            SecretArn = exampleSecret.Arn,
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleSecretVersion,
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleSecret, err := secretsmanager.NewSecret(ctx, "exampleSecret", nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"username": "example",
			"password": "example",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleSecretVersion, err := secretsmanager.NewSecretVersion(ctx, "exampleSecretVersion", &secretsmanager.SecretVersionArgs{
			SecretId:     exampleSecret.ID(),
			SecretString: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = sagemaker.NewCodeRepository(ctx, "exampleCodeRepository", &sagemaker.CodeRepositoryArgs{
			CodeRepositoryName: pulumi.String("example"),
			GitConfig: &sagemaker.CodeRepositoryGitConfigArgs{
				RepositoryUrl: pulumi.String("https://github.com/github/docs.git"),
				SecretArn:     exampleSecret.Arn,
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleSecretVersion,
		}))
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
import com.pulumi.aws.secretsmanager.Secret;
import com.pulumi.aws.secretsmanager.SecretVersion;
import com.pulumi.aws.secretsmanager.SecretVersionArgs;
import com.pulumi.aws.sagemaker.CodeRepository;
import com.pulumi.aws.sagemaker.CodeRepositoryArgs;
import com.pulumi.aws.sagemaker.inputs.CodeRepositoryGitConfigArgs;
import static com.pulumi.codegen.internal.Serialization.*;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleSecret = new Secret("exampleSecret");

        var exampleSecretVersion = new SecretVersion("exampleSecretVersion", SecretVersionArgs.builder()        
            .secretId(exampleSecret.id())
            .secretString(serializeJson(
                jsonObject(
                    jsonProperty("username", "example"),
                    jsonProperty("password", "example")
                )))
            .build());

        var exampleCodeRepository = new CodeRepository("exampleCodeRepository", CodeRepositoryArgs.builder()        
            .codeRepositoryName("example")
            .gitConfig(CodeRepositoryGitConfigArgs.builder()
                .repositoryUrl("https://github.com/github/docs.git")
                .secretArn(exampleSecret.arn())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleSecretVersion)
                .build());

    }
}
```
```yaml
resources:
  exampleSecret:
    type: aws:secretsmanager:Secret
  exampleSecretVersion:
    type: aws:secretsmanager:SecretVersion
    properties:
      secretId: ${exampleSecret.id}
      secretString:
        fn::toJSON:
          username: example
          password: example
  exampleCodeRepository:
    type: aws:sagemaker:CodeRepository
    properties:
      codeRepositoryName: example
      gitConfig:
        repositoryUrl: https://github.com/github/docs.git
        secretArn: ${exampleSecret.arn}
    options:
      dependson:
        - ${exampleSecretVersion}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Code Repositories using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/codeRepository:CodeRepository test_code_repository my-code-repo
```
 