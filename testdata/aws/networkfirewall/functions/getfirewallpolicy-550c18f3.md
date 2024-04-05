Retrieve information about a firewall policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Find firewall policy by name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.networkfirewall.getFirewallPolicy({
    name: _var.firewall_policy_name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.get_firewall_policy(name=var["firewall_policy_name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.NetworkFirewall.GetFirewallPolicy.Invoke(new()
    {
        Name = @var.Firewall_policy_name,
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
		_, err := networkfirewall.LookupFirewallPolicy(ctx, &networkfirewall.LookupFirewallPolicyArgs{
			Name: pulumi.StringRef(_var.Firewall_policy_name),
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
import com.pulumi.aws.networkfirewall.inputs.GetFirewallPolicyArgs;
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
        final var example = NetworkfirewallFunctions.getFirewallPolicy(GetFirewallPolicyArgs.builder()
            .name(var_.firewall_policy_name())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:networkfirewall:getFirewallPolicy
      Arguments:
        name: ${var.firewall_policy_name}
```
{{% /example %}}
{{% example %}}
### Find firewall policy by ARN

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.networkfirewall.getFirewallPolicy({
    arn: _var.firewall_policy_arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.get_firewall_policy(arn=var["firewall_policy_arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.NetworkFirewall.GetFirewallPolicy.Invoke(new()
    {
        Arn = @var.Firewall_policy_arn,
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
		_, err := networkfirewall.LookupFirewallPolicy(ctx, &networkfirewall.LookupFirewallPolicyArgs{
			Arn: pulumi.StringRef(_var.Firewall_policy_arn),
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
import com.pulumi.aws.networkfirewall.inputs.GetFirewallPolicyArgs;
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
        final var example = NetworkfirewallFunctions.getFirewallPolicy(GetFirewallPolicyArgs.builder()
            .arn(var_.firewall_policy_arn())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:networkfirewall:getFirewallPolicy
      Arguments:
        arn: ${var.firewall_policy_arn}
```
{{% /example %}}
{{% example %}}
### Find firewall policy by name and ARN

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.networkfirewall.getFirewallPolicy({
    arn: _var.firewall_policy_arn,
    name: _var.firewall_policy_name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.networkfirewall.get_firewall_policy(arn=var["firewall_policy_arn"],
    name=var["firewall_policy_name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.NetworkFirewall.GetFirewallPolicy.Invoke(new()
    {
        Arn = @var.Firewall_policy_arn,
        Name = @var.Firewall_policy_name,
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
		_, err := networkfirewall.LookupFirewallPolicy(ctx, &networkfirewall.LookupFirewallPolicyArgs{
			Arn:  pulumi.StringRef(_var.Firewall_policy_arn),
			Name: pulumi.StringRef(_var.Firewall_policy_name),
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
import com.pulumi.aws.networkfirewall.inputs.GetFirewallPolicyArgs;
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
        final var example = NetworkfirewallFunctions.getFirewallPolicy(GetFirewallPolicyArgs.builder()
            .arn(var_.firewall_policy_arn())
            .name(var_.firewall_policy_name())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:networkfirewall:getFirewallPolicy
      Arguments:
        arn: ${var.firewall_policy_arn}
        name: ${var.firewall_policy_name}
```

AWS Network Firewall does not allow multiple firewall policies with the same name to be created in an account. It is possible, however, to have multiple firewall policies available in a single account with identical `name` values but distinct `arn` values, e.g. firewall policies shared via a [Resource Access Manager (RAM) share][1]. In that case specifying `arn`, or `name` and `arn`, is recommended.

> **Note:** If there are multiple firewall policies in an account with the same `name`, and `arn` is not specified, the default behavior will return the firewall policy with `name` that was created in the account.
{{% /example %}}
{{% /examples %}}