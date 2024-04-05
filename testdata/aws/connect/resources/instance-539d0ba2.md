Provides an Amazon Connect instance resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

!> **WARN:** Amazon Connect enforces a limit of [100 combined instance creation and deletions every 30 days](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-service-limits.html#feature-limits).  For example, if you create 80 instances and delete 20 of them, you must wait 30 days to create or delete another instance.  Use care when creating or deleting instances.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.connect.Instance("test", {
    identityManagementType: "CONNECT_MANAGED",
    inboundCallsEnabled: true,
    instanceAlias: "friendly-name-connect",
    outboundCallsEnabled: true,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.connect.Instance("test",
    identity_management_type="CONNECT_MANAGED",
    inbound_calls_enabled=True,
    instance_alias="friendly-name-connect",
    outbound_calls_enabled=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Connect.Instance("test", new()
    {
        IdentityManagementType = "CONNECT_MANAGED",
        InboundCallsEnabled = true,
        InstanceAlias = "friendly-name-connect",
        OutboundCallsEnabled = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewInstance(ctx, "test", &connect.InstanceArgs{
			IdentityManagementType: pulumi.String("CONNECT_MANAGED"),
			InboundCallsEnabled:    pulumi.Bool(true),
			InstanceAlias:          pulumi.String("friendly-name-connect"),
			OutboundCallsEnabled:   pulumi.Bool(true),
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
import com.pulumi.aws.connect.Instance;
import com.pulumi.aws.connect.InstanceArgs;
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
        var test = new Instance("test", InstanceArgs.builder()        
            .identityManagementType("CONNECT_MANAGED")
            .inboundCallsEnabled(true)
            .instanceAlias("friendly-name-connect")
            .outboundCallsEnabled(true)
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:connect:Instance
    properties:
      identityManagementType: CONNECT_MANAGED
      inboundCallsEnabled: true
      instanceAlias: friendly-name-connect
      outboundCallsEnabled: true
```

{{% /example %}}
{{% example %}}
### With Existing Active Directory

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.connect.Instance("test", {
    directoryId: aws_directory_service_directory.test.id,
    identityManagementType: "EXISTING_DIRECTORY",
    inboundCallsEnabled: true,
    instanceAlias: "friendly-name-connect",
    outboundCallsEnabled: true,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.connect.Instance("test",
    directory_id=aws_directory_service_directory["test"]["id"],
    identity_management_type="EXISTING_DIRECTORY",
    inbound_calls_enabled=True,
    instance_alias="friendly-name-connect",
    outbound_calls_enabled=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Connect.Instance("test", new()
    {
        DirectoryId = aws_directory_service_directory.Test.Id,
        IdentityManagementType = "EXISTING_DIRECTORY",
        InboundCallsEnabled = true,
        InstanceAlias = "friendly-name-connect",
        OutboundCallsEnabled = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewInstance(ctx, "test", &connect.InstanceArgs{
			DirectoryId:            pulumi.Any(aws_directory_service_directory.Test.Id),
			IdentityManagementType: pulumi.String("EXISTING_DIRECTORY"),
			InboundCallsEnabled:    pulumi.Bool(true),
			InstanceAlias:          pulumi.String("friendly-name-connect"),
			OutboundCallsEnabled:   pulumi.Bool(true),
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
import com.pulumi.aws.connect.Instance;
import com.pulumi.aws.connect.InstanceArgs;
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
        var test = new Instance("test", InstanceArgs.builder()        
            .directoryId(aws_directory_service_directory.test().id())
            .identityManagementType("EXISTING_DIRECTORY")
            .inboundCallsEnabled(true)
            .instanceAlias("friendly-name-connect")
            .outboundCallsEnabled(true)
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:connect:Instance
    properties:
      directoryId: ${aws_directory_service_directory.test.id}
      identityManagementType: EXISTING_DIRECTORY
      inboundCallsEnabled: true
      instanceAlias: friendly-name-connect
      outboundCallsEnabled: true
```

{{% /example %}}
{{% example %}}
### With SAML

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.connect.Instance("test", {
    identityManagementType: "SAML",
    inboundCallsEnabled: true,
    instanceAlias: "friendly-name-connect",
    outboundCallsEnabled: true,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.connect.Instance("test",
    identity_management_type="SAML",
    inbound_calls_enabled=True,
    instance_alias="friendly-name-connect",
    outbound_calls_enabled=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Connect.Instance("test", new()
    {
        IdentityManagementType = "SAML",
        InboundCallsEnabled = true,
        InstanceAlias = "friendly-name-connect",
        OutboundCallsEnabled = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewInstance(ctx, "test", &connect.InstanceArgs{
			IdentityManagementType: pulumi.String("SAML"),
			InboundCallsEnabled:    pulumi.Bool(true),
			InstanceAlias:          pulumi.String("friendly-name-connect"),
			OutboundCallsEnabled:   pulumi.Bool(true),
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
import com.pulumi.aws.connect.Instance;
import com.pulumi.aws.connect.InstanceArgs;
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
        var test = new Instance("test", InstanceArgs.builder()        
            .identityManagementType("SAML")
            .inboundCallsEnabled(true)
            .instanceAlias("friendly-name-connect")
            .outboundCallsEnabled(true)
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:connect:Instance
    properties:
      identityManagementType: SAML
      inboundCallsEnabled: true
      instanceAlias: friendly-name-connect
      outboundCallsEnabled: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Connect instances using the `id`. For example:

```sh
 $ pulumi import aws:connect/instance:Instance example f1288a1f-6193-445a-b47e-af739b2
```
 