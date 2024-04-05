Provides a Connection of Direct Connect.

{{% examples %}}
## Example Usage
{{% example %}}
### Create a connection

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const hoge = new aws.directconnect.Connection("hoge", {
    bandwidth: "1Gbps",
    location: "EqDC2",
});
```
```python
import pulumi
import pulumi_aws as aws

hoge = aws.directconnect.Connection("hoge",
    bandwidth="1Gbps",
    location="EqDC2")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var hoge = new Aws.DirectConnect.Connection("hoge", new()
    {
        Bandwidth = "1Gbps",
        Location = "EqDC2",
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
		_, err := directconnect.NewConnection(ctx, "hoge", &directconnect.ConnectionArgs{
			Bandwidth: pulumi.String("1Gbps"),
			Location:  pulumi.String("EqDC2"),
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
import com.pulumi.aws.directconnect.Connection;
import com.pulumi.aws.directconnect.ConnectionArgs;
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
        var hoge = new Connection("hoge", ConnectionArgs.builder()        
            .bandwidth("1Gbps")
            .location("EqDC2")
            .build());

    }
}
```
```yaml
resources:
  hoge:
    type: aws:directconnect:Connection
    properties:
      bandwidth: 1Gbps
      location: EqDC2
```
{{% /example %}}
{{% example %}}
### Request a MACsec-capable connection

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.directconnect.Connection("example", {
    bandwidth: "10Gbps",
    location: "EqDA2",
    requestMacsec: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directconnect.Connection("example",
    bandwidth="10Gbps",
    location="EqDA2",
    request_macsec=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DirectConnect.Connection("example", new()
    {
        Bandwidth = "10Gbps",
        Location = "EqDA2",
        RequestMacsec = true,
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
		_, err := directconnect.NewConnection(ctx, "example", &directconnect.ConnectionArgs{
			Bandwidth:     pulumi.String("10Gbps"),
			Location:      pulumi.String("EqDA2"),
			RequestMacsec: pulumi.Bool(true),
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
import com.pulumi.aws.directconnect.Connection;
import com.pulumi.aws.directconnect.ConnectionArgs;
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
        var example = new Connection("example", ConnectionArgs.builder()        
            .bandwidth("10Gbps")
            .location("EqDA2")
            .requestMacsec(true)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:directconnect:Connection
    properties:
      bandwidth: 10Gbps
      location: EqDA2
      requestMacsec: true
```
{{% /example %}}
{{% example %}}
### Configure encryption mode for MACsec-capable connections

> **NOTE:** You can only specify the `encryption_mode` argument once the connection is in an `Available` state.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.directconnect.Connection("example", {
    bandwidth: "10Gbps",
    encryptionMode: "must_encrypt",
    location: "EqDC2",
    requestMacsec: true,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directconnect.Connection("example",
    bandwidth="10Gbps",
    encryption_mode="must_encrypt",
    location="EqDC2",
    request_macsec=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DirectConnect.Connection("example", new()
    {
        Bandwidth = "10Gbps",
        EncryptionMode = "must_encrypt",
        Location = "EqDC2",
        RequestMacsec = true,
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
		_, err := directconnect.NewConnection(ctx, "example", &directconnect.ConnectionArgs{
			Bandwidth:      pulumi.String("10Gbps"),
			EncryptionMode: pulumi.String("must_encrypt"),
			Location:       pulumi.String("EqDC2"),
			RequestMacsec:  pulumi.Bool(true),
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
import com.pulumi.aws.directconnect.Connection;
import com.pulumi.aws.directconnect.ConnectionArgs;
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
        var example = new Connection("example", ConnectionArgs.builder()        
            .bandwidth("10Gbps")
            .encryptionMode("must_encrypt")
            .location("EqDC2")
            .requestMacsec(true)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:directconnect:Connection
    properties:
      bandwidth: 10Gbps
      encryptionMode: must_encrypt
      location: EqDC2
      requestMacsec: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Direct Connect connections using the connection `id`. For example:

```sh
 $ pulumi import aws:directconnect/connection:Connection test_connection dxcon-ffre0ec3
```
 