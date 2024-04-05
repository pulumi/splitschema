`aws.route53.getResolverRules` provides details about a set of Route53 Resolver rules.

{{% examples %}}
## Example Usage
{{% example %}}
### Retrieving the default resolver rule

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.route53.getResolverRules({
    ownerId: "Route 53 Resolver",
    ruleType: "RECURSIVE",
    shareStatus: "NOT_SHARED",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.get_resolver_rules(owner_id="Route 53 Resolver",
    rule_type="RECURSIVE",
    share_status="NOT_SHARED")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Route53.GetResolverRules.Invoke(new()
    {
        OwnerId = "Route 53 Resolver",
        RuleType = "RECURSIVE",
        ShareStatus = "NOT_SHARED",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.GetResolverRules(ctx, &route53.GetResolverRulesArgs{
			OwnerId:     pulumi.StringRef("Route 53 Resolver"),
			RuleType:    pulumi.StringRef("RECURSIVE"),
			ShareStatus: pulumi.StringRef("NOT_SHARED"),
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
import com.pulumi.aws.route53.Route53Functions;
import com.pulumi.aws.route53.inputs.GetResolverRulesArgs;
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
        final var example = Route53Functions.getResolverRules(GetResolverRulesArgs.builder()
            .ownerId("Route 53 Resolver")
            .ruleType("RECURSIVE")
            .shareStatus("NOT_SHARED")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:route53:getResolverRules
      Arguments:
        ownerId: Route 53 Resolver
        ruleType: RECURSIVE
        shareStatus: NOT_SHARED
```
{{% /example %}}
{{% example %}}
### Retrieving forward rules shared with me

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.route53.getResolverRules({
    ruleType: "FORWARD",
    shareStatus: "SHARED_WITH_ME",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.get_resolver_rules(rule_type="FORWARD",
    share_status="SHARED_WITH_ME")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Route53.GetResolverRules.Invoke(new()
    {
        RuleType = "FORWARD",
        ShareStatus = "SHARED_WITH_ME",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.GetResolverRules(ctx, &route53.GetResolverRulesArgs{
			RuleType:    pulumi.StringRef("FORWARD"),
			ShareStatus: pulumi.StringRef("SHARED_WITH_ME"),
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
import com.pulumi.aws.route53.Route53Functions;
import com.pulumi.aws.route53.inputs.GetResolverRulesArgs;
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
        final var example = Route53Functions.getResolverRules(GetResolverRulesArgs.builder()
            .ruleType("FORWARD")
            .shareStatus("SHARED_WITH_ME")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:route53:getResolverRules
      Arguments:
        ruleType: FORWARD
        shareStatus: SHARED_WITH_ME
```
{{% /example %}}
{{% example %}}
### Retrieving rules by name regex

Resolver rules whose name contains `abc`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.route53.getResolverRules({
    nameRegex: ".*abc.*",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.get_resolver_rules(name_regex=".*abc.*")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Route53.GetResolverRules.Invoke(new()
    {
        NameRegex = ".*abc.*",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.GetResolverRules(ctx, &route53.GetResolverRulesArgs{
			NameRegex: pulumi.StringRef(".*abc.*"),
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
import com.pulumi.aws.route53.Route53Functions;
import com.pulumi.aws.route53.inputs.GetResolverRulesArgs;
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
        final var example = Route53Functions.getResolverRules(GetResolverRulesArgs.builder()
            .nameRegex(".*abc.*")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:route53:getResolverRules
      Arguments:
        nameRegex: .*abc.*
```
{{% /example %}}
{{% /examples %}}