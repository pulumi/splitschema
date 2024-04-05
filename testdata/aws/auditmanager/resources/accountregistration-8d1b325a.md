Resource for managing AWS Audit Manager Account Registration.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.auditmanager.AccountRegistration("example", {});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.auditmanager.AccountRegistration("example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Auditmanager.AccountRegistration("example");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/auditmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := auditmanager.NewAccountRegistration(ctx, "example", nil)
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
import com.pulumi.aws.auditmanager.AccountRegistration;
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
        var example = new AccountRegistration("example");

    }
}
```
```yaml
resources:
  example:
    type: aws:auditmanager:AccountRegistration
```
{{% /example %}}
{{% example %}}
### Deregister On Destroy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.auditmanager.AccountRegistration("example", {deregisterOnDestroy: true});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.auditmanager.AccountRegistration("example", deregister_on_destroy=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Auditmanager.AccountRegistration("example", new()
    {
        DeregisterOnDestroy = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/auditmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := auditmanager.NewAccountRegistration(ctx, "example", &auditmanager.AccountRegistrationArgs{
			DeregisterOnDestroy: pulumi.Bool(true),
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
import com.pulumi.aws.auditmanager.AccountRegistration;
import com.pulumi.aws.auditmanager.AccountRegistrationArgs;
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
        var example = new AccountRegistration("example", AccountRegistrationArgs.builder()        
            .deregisterOnDestroy(true)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:auditmanager:AccountRegistration
    properties:
      deregisterOnDestroy: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Audit Manager Account Registration resources using the `id`. For example:

```sh
 $ pulumi import aws:auditmanager/accountRegistration:AccountRegistration example us-east-1
```
 