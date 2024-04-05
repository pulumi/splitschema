Retrieve information about a Secrets Manager secret version, including its secret value. To retrieve secret metadata, see the `aws.secretsmanager.Secret` data source.

{{% examples %}}
## Example Usage
{{% example %}}
### Retrieve Current Secret Version

By default, this data sources retrieves information based on the `AWSCURRENT` staging label.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const secret-version = aws.secretsmanager.getSecretVersion({
    secretId: data.aws_secretsmanager_secret.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

secret_version = aws.secretsmanager.get_secret_version(secret_id=data["aws_secretsmanager_secret"]["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var secret_version = Aws.SecretsManager.GetSecretVersion.Invoke(new()
    {
        SecretId = data.Aws_secretsmanager_secret.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := secretsmanager.LookupSecretVersion(ctx, &secretsmanager.LookupSecretVersionArgs{
			SecretId: data.Aws_secretsmanager_secret.Example.Id,
		}, nil)
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
import com.pulumi.aws.secretsmanager.SecretsmanagerFunctions;
import com.pulumi.aws.secretsmanager.inputs.GetSecretVersionArgs;
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
        final var secret-version = SecretsmanagerFunctions.getSecretVersion(GetSecretVersionArgs.builder()
            .secretId(data.aws_secretsmanager_secret().example().id())
            .build());

    }
}
```
```yaml
variables:
  secret-version:
    fn::invoke:
      Function: aws:secretsmanager:getSecretVersion
      Arguments:
        secretId: ${data.aws_secretsmanager_secret.example.id}
```
{{% /example %}}
{{% example %}}
### Retrieve Specific Secret Version

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const by-version-stage = aws.secretsmanager.getSecretVersion({
    secretId: data.aws_secretsmanager_secret.example.id,
    versionStage: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

by_version_stage = aws.secretsmanager.get_secret_version(secret_id=data["aws_secretsmanager_secret"]["example"]["id"],
    version_stage="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var by_version_stage = Aws.SecretsManager.GetSecretVersion.Invoke(new()
    {
        SecretId = data.Aws_secretsmanager_secret.Example.Id,
        VersionStage = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := secretsmanager.LookupSecretVersion(ctx, &secretsmanager.LookupSecretVersionArgs{
			SecretId:     data.Aws_secretsmanager_secret.Example.Id,
			VersionStage: pulumi.StringRef("example"),
		}, nil)
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
import com.pulumi.aws.secretsmanager.SecretsmanagerFunctions;
import com.pulumi.aws.secretsmanager.inputs.GetSecretVersionArgs;
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
        final var by-version-stage = SecretsmanagerFunctions.getSecretVersion(GetSecretVersionArgs.builder()
            .secretId(data.aws_secretsmanager_secret().example().id())
            .versionStage("example")
            .build());

    }
}
```
```yaml
variables:
  by-version-stage:
    fn::invoke:
      Function: aws:secretsmanager:getSecretVersion
      Arguments:
        secretId: ${data.aws_secretsmanager_secret.example.id}
        versionStage: example
```
{{% /example %}}
{{% /examples %}}