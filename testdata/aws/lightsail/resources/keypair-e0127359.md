Provides a Lightsail Key Pair, for use with Lightsail Instances. These key pairs
are separate from EC2 Key Pairs, and must be created or imported for use with
Lightsail.

> **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details

{{% examples %}}
## Example Usage
{{% example %}}
### Create New Key Pair

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new Lightsail Key Pair
const lgKeyPair = new aws.lightsail.KeyPair("lgKeyPair", {});
```
```python
import pulumi
import pulumi_aws as aws

# Create a new Lightsail Key Pair
lg_key_pair = aws.lightsail.KeyPair("lgKeyPair")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Create a new Lightsail Key Pair
    var lgKeyPair = new Aws.LightSail.KeyPair("lgKeyPair");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewKeyPair(ctx, "lgKeyPair", nil)
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
import com.pulumi.aws.lightsail.KeyPair;
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
        var lgKeyPair = new KeyPair("lgKeyPair");

    }
}
```
```yaml
resources:
  # Create a new Lightsail Key Pair
  lgKeyPair:
    type: aws:lightsail:KeyPair
```
{{% /example %}}
{{% example %}}
### Create New Key Pair with PGP Encrypted Private Key

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lgKeyPair = new aws.lightsail.KeyPair("lgKeyPair", {pgpKey: "keybase:keybaseusername"});
```
```python
import pulumi
import pulumi_aws as aws

lg_key_pair = aws.lightsail.KeyPair("lgKeyPair", pgp_key="keybase:keybaseusername")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var lgKeyPair = new Aws.LightSail.KeyPair("lgKeyPair", new()
    {
        PgpKey = "keybase:keybaseusername",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewKeyPair(ctx, "lgKeyPair", &lightsail.KeyPairArgs{
			PgpKey: pulumi.String("keybase:keybaseusername"),
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
import com.pulumi.aws.lightsail.KeyPair;
import com.pulumi.aws.lightsail.KeyPairArgs;
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
        var lgKeyPair = new KeyPair("lgKeyPair", KeyPairArgs.builder()        
            .pgpKey("keybase:keybaseusername")
            .build());

    }
}
```
```yaml
resources:
  lgKeyPair:
    type: aws:lightsail:KeyPair
    properties:
      pgpKey: keybase:keybaseusername
```
{{% /example %}}
{{% example %}}
### Existing Public Key Import

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const lgKeyPair = new aws.lightsail.KeyPair("lgKeyPair", {publicKey: fs.readFileSync("~/.ssh/id_rsa.pub", "utf8")});
```
```python
import pulumi
import pulumi_aws as aws

lg_key_pair = aws.lightsail.KeyPair("lgKeyPair", public_key=(lambda path: open(path).read())("~/.ssh/id_rsa.pub"))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var lgKeyPair = new Aws.LightSail.KeyPair("lgKeyPair", new()
    {
        PublicKey = File.ReadAllText("~/.ssh/id_rsa.pub"),
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
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
		_, err := lightsail.NewKeyPair(ctx, "lgKeyPair", &lightsail.KeyPairArgs{
			PublicKey: readFileOrPanic("~/.ssh/id_rsa.pub"),
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
import com.pulumi.aws.lightsail.KeyPair;
import com.pulumi.aws.lightsail.KeyPairArgs;
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
        var lgKeyPair = new KeyPair("lgKeyPair", KeyPairArgs.builder()        
            .publicKey(Files.readString(Paths.get("~/.ssh/id_rsa.pub")))
            .build());

    }
}
```
```yaml
resources:
  lgKeyPair:
    type: aws:lightsail:KeyPair
    properties:
      publicKey:
        fn::readFile: ~/.ssh/id_rsa.pub
```
{{% /example %}}
{{% /examples %}}

## Import

You cannot import Lightsail Key Pairs because the private and public key are only available on initial creation. 