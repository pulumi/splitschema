{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.worklink.Fleet("example", {});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.worklink.Fleet("example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.WorkLink.Fleet("example");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/worklink"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := worklink.NewFleet(ctx, "example", nil)
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
import com.pulumi.aws.worklink.Fleet;
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
        var example = new Fleet("example");

    }
}
```
```yaml
resources:
  example:
    type: aws:worklink:Fleet
```

Network Configuration Usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.worklink.Fleet("example", {network: {
    vpcId: aws_vpc.test.id,
    subnetIds: [aws_subnet.test.map(__item => __item.id)],
    securityGroupIds: [aws_security_group.test.id],
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.worklink.Fleet("example", network=aws.worklink.FleetNetworkArgs(
    vpc_id=aws_vpc["test"]["id"],
    subnet_ids=[[__item["id"] for __item in aws_subnet["test"]]],
    security_group_ids=[aws_security_group["test"]["id"]],
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.WorkLink.Fleet("example", new()
    {
        Network = new Aws.WorkLink.Inputs.FleetNetworkArgs
        {
            VpcId = aws_vpc.Test.Id,
            SubnetIds = new[]
            {
                aws_subnet.Test.Select(__item => __item.Id).ToList(),
            },
            SecurityGroupIds = new[]
            {
                aws_security_group.Test.Id,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/worklink"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := worklink.NewFleet(ctx, "example", &worklink.FleetArgs{
Network: &worklink.FleetNetworkArgs{
VpcId: pulumi.Any(aws_vpc.Test.Id),
SubnetIds: pulumi.StringArray{
%!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-resources-aws:worklink-fleet:Fleet.pp:3,26-47),
},
SecurityGroupIds: pulumi.StringArray{
aws_security_group.Test.Id,
},
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
import com.pulumi.aws.worklink.Fleet;
import com.pulumi.aws.worklink.FleetArgs;
import com.pulumi.aws.worklink.inputs.FleetNetworkArgs;
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
        var example = new Fleet("example", FleetArgs.builder()        
            .network(FleetNetworkArgs.builder()
                .vpcId(aws_vpc.test().id())
                .subnetIds(aws_subnet.test().stream().map(element -> element.id()).collect(toList()))
                .securityGroupIds(aws_security_group.test().id())
                .build())
            .build());

    }
}
```

Identity Provider Configuration Usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const test = new aws.worklink.Fleet("test", {identityProvider: {
    type: "SAML",
    samlMetadata: fs.readFileSync("saml-metadata.xml", "utf8"),
}});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.worklink.Fleet("test", identity_provider=aws.worklink.FleetIdentityProviderArgs(
    type="SAML",
    saml_metadata=(lambda path: open(path).read())("saml-metadata.xml"),
))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.WorkLink.Fleet("test", new()
    {
        IdentityProvider = new Aws.WorkLink.Inputs.FleetIdentityProviderArgs
        {
            Type = "SAML",
            SamlMetadata = File.ReadAllText("saml-metadata.xml"),
        },
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/worklink"
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
		_, err := worklink.NewFleet(ctx, "test", &worklink.FleetArgs{
			IdentityProvider: &worklink.FleetIdentityProviderArgs{
				Type:         pulumi.String("SAML"),
				SamlMetadata: readFileOrPanic("saml-metadata.xml"),
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
import com.pulumi.aws.worklink.Fleet;
import com.pulumi.aws.worklink.FleetArgs;
import com.pulumi.aws.worklink.inputs.FleetIdentityProviderArgs;
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
        var test = new Fleet("test", FleetArgs.builder()        
            .identityProvider(FleetIdentityProviderArgs.builder()
                .type("SAML")
                .samlMetadata(Files.readString(Paths.get("saml-metadata.xml")))
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:worklink:Fleet
    properties:
      identityProvider:
        type: SAML
        samlMetadata:
          fn::readFile: saml-metadata.xml
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WorkLink using the ARN. For example:

```sh
 $ pulumi import aws:worklink/fleet:Fleet test arn:aws:worklink::123456789012:fleet/example
```
 