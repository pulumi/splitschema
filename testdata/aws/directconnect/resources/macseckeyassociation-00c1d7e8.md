Provides a MAC Security (MACSec) secret key resource for use with Direct Connect. See [MACsec prerequisites](https://docs.aws.amazon.com/directconnect/latest/UserGuide/direct-connect-mac-sec-getting-started.html#mac-sec-prerequisites) for information about MAC Security (MACsec) prerequisites.

Creating this resource will also create a resource of type `aws.secretsmanager.Secret` which is managed by Direct Connect. While you can import this resource into your state, because this secret is managed by Direct Connect, you will not be able to make any modifications to it. See [How AWS Direct Connect uses AWS Secrets Manager](https://docs.aws.amazon.com/secretsmanager/latest/userguide/integrating_how-services-use-secrets_directconnect.html) for details.

> **Note:** All arguments including `ckn` and `cak` will be stored in the raw state as plain-text.
> **Note:** The `secret_arn` argument can only be used to reference a previously created MACSec key. You cannot associate a Secrets Manager secret created outside of the `aws.directconnect.MacsecKeyAssociation` resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Create MACSec key with CKN and CAK

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.directconnect.getConnection({
    name: "tf-dx-connection",
});
const test = new aws.directconnect.MacsecKeyAssociation("test", {
    connectionId: example.then(example => example.id),
    ckn: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
    cak: "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directconnect.get_connection(name="tf-dx-connection")
test = aws.directconnect.MacsecKeyAssociation("test",
    connection_id=example.id,
    ckn="0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
    cak="abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.DirectConnect.GetConnection.Invoke(new()
    {
        Name = "tf-dx-connection",
    });

    var test = new Aws.DirectConnect.MacsecKeyAssociation("test", new()
    {
        ConnectionId = example.Apply(getConnectionResult => getConnectionResult.Id),
        Ckn = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
        Cak = "abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := directconnect.LookupConnection(ctx, &directconnect.LookupConnectionArgs{
			Name: "tf-dx-connection",
		}, nil)
		if err != nil {
			return err
		}
		_, err = directconnect.NewMacsecKeyAssociation(ctx, "test", &directconnect.MacsecKeyAssociationArgs{
			ConnectionId: *pulumi.String(example.Id),
			Ckn:          pulumi.String("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"),
			Cak:          pulumi.String("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789"),
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
import com.pulumi.aws.directconnect.DirectconnectFunctions;
import com.pulumi.aws.directconnect.inputs.GetConnectionArgs;
import com.pulumi.aws.directconnect.MacsecKeyAssociation;
import com.pulumi.aws.directconnect.MacsecKeyAssociationArgs;
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
        final var example = DirectconnectFunctions.getConnection(GetConnectionArgs.builder()
            .name("tf-dx-connection")
            .build());

        var test = new MacsecKeyAssociation("test", MacsecKeyAssociationArgs.builder()        
            .connectionId(example.applyValue(getConnectionResult -> getConnectionResult.id()))
            .ckn("0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")
            .cak("abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:directconnect:MacsecKeyAssociation
    properties:
      connectionId: ${example.id}
      ckn: 0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
      cak: abcdef0123456789abcdef0123456789abcdef0123456789abcdef0123456789
variables:
  example:
    fn::invoke:
      Function: aws:directconnect:getConnection
      Arguments:
        name: tf-dx-connection
```
{{% /example %}}
{{% example %}}
### Create MACSec key with existing Secrets Manager secret

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleConnection = aws.directconnect.getConnection({
    name: "tf-dx-connection",
});
const exampleSecret = aws.secretsmanager.getSecret({
    name: "directconnect!prod/us-east-1/directconnect/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
});
const test = new aws.directconnect.MacsecKeyAssociation("test", {
    connectionId: exampleConnection.then(exampleConnection => exampleConnection.id),
    secretArn: exampleSecret.then(exampleSecret => exampleSecret.arn),
});
```
```python
import pulumi
import pulumi_aws as aws

example_connection = aws.directconnect.get_connection(name="tf-dx-connection")
example_secret = aws.secretsmanager.get_secret(name="directconnect!prod/us-east-1/directconnect/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")
test = aws.directconnect.MacsecKeyAssociation("test",
    connection_id=example_connection.id,
    secret_arn=example_secret.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleConnection = Aws.DirectConnect.GetConnection.Invoke(new()
    {
        Name = "tf-dx-connection",
    });

    var exampleSecret = Aws.SecretsManager.GetSecret.Invoke(new()
    {
        Name = "directconnect!prod/us-east-1/directconnect/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
    });

    var test = new Aws.DirectConnect.MacsecKeyAssociation("test", new()
    {
        ConnectionId = exampleConnection.Apply(getConnectionResult => getConnectionResult.Id),
        SecretArn = exampleSecret.Apply(getSecretResult => getSecretResult.Arn),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleConnection, err := directconnect.LookupConnection(ctx, &directconnect.LookupConnectionArgs{
			Name: "tf-dx-connection",
		}, nil)
		if err != nil {
			return err
		}
		exampleSecret, err := secretsmanager.LookupSecret(ctx, &secretsmanager.LookupSecretArgs{
			Name: pulumi.StringRef("directconnect!prod/us-east-1/directconnect/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = directconnect.NewMacsecKeyAssociation(ctx, "test", &directconnect.MacsecKeyAssociationArgs{
			ConnectionId: *pulumi.String(exampleConnection.Id),
			SecretArn:    *pulumi.String(exampleSecret.Arn),
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
import com.pulumi.aws.directconnect.DirectconnectFunctions;
import com.pulumi.aws.directconnect.inputs.GetConnectionArgs;
import com.pulumi.aws.secretsmanager.SecretsmanagerFunctions;
import com.pulumi.aws.secretsmanager.inputs.GetSecretArgs;
import com.pulumi.aws.directconnect.MacsecKeyAssociation;
import com.pulumi.aws.directconnect.MacsecKeyAssociationArgs;
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
        final var exampleConnection = DirectconnectFunctions.getConnection(GetConnectionArgs.builder()
            .name("tf-dx-connection")
            .build());

        final var exampleSecret = SecretsmanagerFunctions.getSecret(GetSecretArgs.builder()
            .name("directconnect!prod/us-east-1/directconnect/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef")
            .build());

        var test = new MacsecKeyAssociation("test", MacsecKeyAssociationArgs.builder()        
            .connectionId(exampleConnection.applyValue(getConnectionResult -> getConnectionResult.id()))
            .secretArn(exampleSecret.applyValue(getSecretResult -> getSecretResult.arn()))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:directconnect:MacsecKeyAssociation
    properties:
      connectionId: ${exampleConnection.id}
      secretArn: ${exampleSecret.arn}
variables:
  exampleConnection:
    fn::invoke:
      Function: aws:directconnect:getConnection
      Arguments:
        name: tf-dx-connection
  exampleSecret:
    fn::invoke:
      Function: aws:secretsmanager:getSecret
      Arguments:
        name: directconnect!prod/us-east-1/directconnect/0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
```
{{% /example %}}
{{% /examples %}}