Retrieve metadata information about a Secrets Manager secret. To retrieve a secret value, see the `aws.secretsmanager.SecretVersion` data source.

{{% examples %}}
## Example Usage
{{% example %}}
### ARN

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const by-arn = aws.secretsmanager.getSecret({
    arn: "arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456",
});
```
```python
import pulumi
import pulumi_aws as aws

by_arn = aws.secretsmanager.get_secret(arn="arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var by_arn = Aws.SecretsManager.GetSecret.Invoke(new()
    {
        Arn = "arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456",
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
		_, err := secretsmanager.LookupSecret(ctx, &secretsmanager.LookupSecretArgs{
			Arn: pulumi.StringRef("arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456"),
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
import com.pulumi.aws.secretsmanager.inputs.GetSecretArgs;
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
        final var by-arn = SecretsmanagerFunctions.getSecret(GetSecretArgs.builder()
            .arn("arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456")
            .build());

    }
}
```
```yaml
variables:
  by-arn:
    fn::invoke:
      Function: aws:secretsmanager:getSecret
      Arguments:
        arn: arn:aws:secretsmanager:us-east-1:123456789012:secret:example-123456
```
{{% /example %}}
{{% example %}}
### Name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const by-name = aws.secretsmanager.getSecret({
    name: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

by_name = aws.secretsmanager.get_secret(name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var by_name = Aws.SecretsManager.GetSecret.Invoke(new()
    {
        Name = "example",
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
		_, err := secretsmanager.LookupSecret(ctx, &secretsmanager.LookupSecretArgs{
			Name: pulumi.StringRef("example"),
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
import com.pulumi.aws.secretsmanager.inputs.GetSecretArgs;
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
        final var by-name = SecretsmanagerFunctions.getSecret(GetSecretArgs.builder()
            .name("example")
            .build());

    }
}
```
```yaml
variables:
  by-name:
    fn::invoke:
      Function: aws:secretsmanager:getSecret
      Arguments:
        name: example
```
{{% /example %}}
{{% /examples %}}