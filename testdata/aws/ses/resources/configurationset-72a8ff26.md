Provides an SES configuration set resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ses.ConfigurationSet("test", {});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ses.ConfigurationSet("test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Ses.ConfigurationSet("test");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewConfigurationSet(ctx, "test", nil)
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
import com.pulumi.aws.ses.ConfigurationSet;
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
        var test = new ConfigurationSet("test");

    }
}
```
```yaml
resources:
  test:
    type: aws:ses:ConfigurationSet
```
{{% /example %}}
{{% example %}}
### Require TLS Connections

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ses.ConfigurationSet("test", {deliveryOptions: {
    tlsPolicy: "Require",
}});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ses.ConfigurationSet("test", delivery_options=aws.ses.ConfigurationSetDeliveryOptionsArgs(
    tls_policy="Require",
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Ses.ConfigurationSet("test", new()
    {
        DeliveryOptions = new Aws.Ses.Inputs.ConfigurationSetDeliveryOptionsArgs
        {
            TlsPolicy = "Require",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewConfigurationSet(ctx, "test", &ses.ConfigurationSetArgs{
			DeliveryOptions: &ses.ConfigurationSetDeliveryOptionsArgs{
				TlsPolicy: pulumi.String("Require"),
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
import com.pulumi.aws.ses.ConfigurationSet;
import com.pulumi.aws.ses.ConfigurationSetArgs;
import com.pulumi.aws.ses.inputs.ConfigurationSetDeliveryOptionsArgs;
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
        var test = new ConfigurationSet("test", ConfigurationSetArgs.builder()        
            .deliveryOptions(ConfigurationSetDeliveryOptionsArgs.builder()
                .tlsPolicy("Require")
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:ses:ConfigurationSet
    properties:
      deliveryOptions:
        tlsPolicy: Require
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SES Configuration Sets using their `name`. For example:

```sh
 $ pulumi import aws:ses/configurationSet:ConfigurationSet test some-configuration-set-test
```
 