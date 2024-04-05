Resource for managing an AWS SESv2 (Simple Email V2) Email Identity.

{{% examples %}}
## Example Usage

### Basic Usage
{{% example %}}
### Email Address Identity

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sesv2.EmailIdentity("example", {emailIdentity: "testing@example.com"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sesv2.EmailIdentity("example", email_identity="testing@example.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SesV2.EmailIdentity("example", new()
    {
        EmailIdentityDetails = "testing@example.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sesv2.NewEmailIdentity(ctx, "example", &sesv2.EmailIdentityArgs{
			EmailIdentity: pulumi.String("testing@example.com"),
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
import com.pulumi.aws.sesv2.EmailIdentity;
import com.pulumi.aws.sesv2.EmailIdentityArgs;
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
        var example = new EmailIdentity("example", EmailIdentityArgs.builder()        
            .emailIdentity("testing@example.com")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sesv2:EmailIdentity
    properties:
      emailIdentity: testing@example.com
```
{{% /example %}}
{{% example %}}
### Domain Identity

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sesv2.EmailIdentity("example", {emailIdentity: "example.com"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sesv2.EmailIdentity("example", email_identity="example.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SesV2.EmailIdentity("example", new()
    {
        EmailIdentityDetails = "example.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sesv2.NewEmailIdentity(ctx, "example", &sesv2.EmailIdentityArgs{
			EmailIdentity: pulumi.String("example.com"),
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
import com.pulumi.aws.sesv2.EmailIdentity;
import com.pulumi.aws.sesv2.EmailIdentityArgs;
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
        var example = new EmailIdentity("example", EmailIdentityArgs.builder()        
            .emailIdentity("example.com")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sesv2:EmailIdentity
    properties:
      emailIdentity: example.com
```
{{% /example %}}
{{% example %}}
### Configuration Set

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleConfigurationSet = new aws.sesv2.ConfigurationSet("exampleConfigurationSet", {configurationSetName: "example"});
const exampleEmailIdentity = new aws.sesv2.EmailIdentity("exampleEmailIdentity", {
    emailIdentity: "example.com",
    configurationSetName: exampleConfigurationSet.configurationSetName,
});
```
```python
import pulumi
import pulumi_aws as aws

example_configuration_set = aws.sesv2.ConfigurationSet("exampleConfigurationSet", configuration_set_name="example")
example_email_identity = aws.sesv2.EmailIdentity("exampleEmailIdentity",
    email_identity="example.com",
    configuration_set_name=example_configuration_set.configuration_set_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleConfigurationSet = new Aws.SesV2.ConfigurationSet("exampleConfigurationSet", new()
    {
        ConfigurationSetName = "example",
    });

    var exampleEmailIdentity = new Aws.SesV2.EmailIdentity("exampleEmailIdentity", new()
    {
        EmailIdentityDetails = "example.com",
        ConfigurationSetName = exampleConfigurationSet.ConfigurationSetName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleConfigurationSet, err := sesv2.NewConfigurationSet(ctx, "exampleConfigurationSet", &sesv2.ConfigurationSetArgs{
			ConfigurationSetName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = sesv2.NewEmailIdentity(ctx, "exampleEmailIdentity", &sesv2.EmailIdentityArgs{
			EmailIdentity:        pulumi.String("example.com"),
			ConfigurationSetName: exampleConfigurationSet.ConfigurationSetName,
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
import com.pulumi.aws.sesv2.ConfigurationSet;
import com.pulumi.aws.sesv2.ConfigurationSetArgs;
import com.pulumi.aws.sesv2.EmailIdentity;
import com.pulumi.aws.sesv2.EmailIdentityArgs;
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
        var exampleConfigurationSet = new ConfigurationSet("exampleConfigurationSet", ConfigurationSetArgs.builder()        
            .configurationSetName("example")
            .build());

        var exampleEmailIdentity = new EmailIdentity("exampleEmailIdentity", EmailIdentityArgs.builder()        
            .emailIdentity("example.com")
            .configurationSetName(exampleConfigurationSet.configurationSetName())
            .build());

    }
}
```
```yaml
resources:
  exampleConfigurationSet:
    type: aws:sesv2:ConfigurationSet
    properties:
      configurationSetName: example
  exampleEmailIdentity:
    type: aws:sesv2:EmailIdentity
    properties:
      emailIdentity: example.com
      configurationSetName: ${exampleConfigurationSet.configurationSetName}
```
{{% /example %}}
{{% example %}}
### DKIM Signing Attributes (BYODKIM)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sesv2.EmailIdentity("example", {
    dkimSigningAttributes: {
        domainSigningPrivateKey: "MIIJKAIBAAKCAgEA2Se7p8zvnI4yh+Gh9j2rG5e2aRXjg03Y8saiupLnadPH9xvM...",
        domainSigningSelector: "example",
    },
    emailIdentity: "example.com",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sesv2.EmailIdentity("example",
    dkim_signing_attributes=aws.sesv2.EmailIdentityDkimSigningAttributesArgs(
        domain_signing_private_key="MIIJKAIBAAKCAgEA2Se7p8zvnI4yh+Gh9j2rG5e2aRXjg03Y8saiupLnadPH9xvM...",
        domain_signing_selector="example",
    ),
    email_identity="example.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.SesV2.EmailIdentity("example", new()
    {
        DkimSigningAttributes = new Aws.SesV2.Inputs.EmailIdentityDkimSigningAttributesArgs
        {
            DomainSigningPrivateKey = "MIIJKAIBAAKCAgEA2Se7p8zvnI4yh+Gh9j2rG5e2aRXjg03Y8saiupLnadPH9xvM...",
            DomainSigningSelector = "example",
        },
        EmailIdentityDetails = "example.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sesv2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sesv2.NewEmailIdentity(ctx, "example", &sesv2.EmailIdentityArgs{
			DkimSigningAttributes: &sesv2.EmailIdentityDkimSigningAttributesArgs{
				DomainSigningPrivateKey: pulumi.String("MIIJKAIBAAKCAgEA2Se7p8zvnI4yh+Gh9j2rG5e2aRXjg03Y8saiupLnadPH9xvM..."),
				DomainSigningSelector:   pulumi.String("example"),
			},
			EmailIdentity: pulumi.String("example.com"),
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
import com.pulumi.aws.sesv2.EmailIdentity;
import com.pulumi.aws.sesv2.EmailIdentityArgs;
import com.pulumi.aws.sesv2.inputs.EmailIdentityDkimSigningAttributesArgs;
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
        var example = new EmailIdentity("example", EmailIdentityArgs.builder()        
            .dkimSigningAttributes(EmailIdentityDkimSigningAttributesArgs.builder()
                .domainSigningPrivateKey("MIIJKAIBAAKCAgEA2Se7p8zvnI4yh+Gh9j2rG5e2aRXjg03Y8saiupLnadPH9xvM...")
                .domainSigningSelector("example")
                .build())
            .emailIdentity("example.com")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sesv2:EmailIdentity
    properties:
      dkimSigningAttributes:
        domainSigningPrivateKey: MIIJKAIBAAKCAgEA2Se7p8zvnI4yh+Gh9j2rG5e2aRXjg03Y8saiupLnadPH9xvM...
        domainSigningSelector: example
      emailIdentity: example.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SESv2 (Simple Email V2) Email Identity using the `email_identity`. For example:

```sh
 $ pulumi import aws:sesv2/emailIdentity:EmailIdentity example example.com
```
 