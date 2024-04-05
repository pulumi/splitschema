Retrieve information about a firewall.

{{% examples %}}
## Example Usage
{{% example %}}
### Find firewall policy by ARN

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.networkfirewall.getFirewall({
    arn: aws_networkfirewall_firewall.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.get_firewall(arn=aws_networkfirewall_firewall["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.NetworkFirewall.GetFirewall.Invoke(new()
    {
        Arn = aws_networkfirewall_firewall.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfirewall.LookupFirewall(ctx, &networkfirewall.LookupFirewallArgs{
			Arn: pulumi.StringRef(aws_networkfirewall_firewall.Arn),
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
import com.pulumi.aws.networkfirewall.NetworkfirewallFunctions;
import com.pulumi.aws.networkfirewall.inputs.GetFirewallArgs;
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
        final var example = NetworkfirewallFunctions.getFirewall(GetFirewallArgs.builder()
            .arn(aws_networkfirewall_firewall.arn())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:networkfirewall:getFirewall
      Arguments:
        arn: ${aws_networkfirewall_firewall.arn}
```
{{% /example %}}
{{% example %}}
### Find firewall policy by Name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.networkfirewall.getFirewall({
    name: "Test",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.get_firewall(name="Test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.NetworkFirewall.GetFirewall.Invoke(new()
    {
        Name = "Test",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfirewall.LookupFirewall(ctx, &networkfirewall.LookupFirewallArgs{
			Name: pulumi.StringRef("Test"),
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
import com.pulumi.aws.networkfirewall.NetworkfirewallFunctions;
import com.pulumi.aws.networkfirewall.inputs.GetFirewallArgs;
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
        final var example = NetworkfirewallFunctions.getFirewall(GetFirewallArgs.builder()
            .name("Test")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:networkfirewall:getFirewall
      Arguments:
        name: Test
```
{{% /example %}}
{{% example %}}
### Find firewall policy by ARN and Name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.networkfirewall.getFirewall({
    arn: aws_networkfirewall_firewall.arn,
    name: "Test",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.get_firewall(arn=aws_networkfirewall_firewall["arn"],
    name="Test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.NetworkFirewall.GetFirewall.Invoke(new()
    {
        Arn = aws_networkfirewall_firewall.Arn,
        Name = "Test",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkfirewall"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfirewall.LookupFirewall(ctx, &networkfirewall.LookupFirewallArgs{
			Arn:  pulumi.StringRef(aws_networkfirewall_firewall.Arn),
			Name: pulumi.StringRef("Test"),
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
import com.pulumi.aws.networkfirewall.NetworkfirewallFunctions;
import com.pulumi.aws.networkfirewall.inputs.GetFirewallArgs;
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
        final var example = NetworkfirewallFunctions.getFirewall(GetFirewallArgs.builder()
            .arn(aws_networkfirewall_firewall.arn())
            .name("Test")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:networkfirewall:getFirewall
      Arguments:
        arn: ${aws_networkfirewall_firewall.arn}
        name: Test
```
{{% /example %}}
{{% /examples %}}