Provides a resource to attach an AWS Organizations policy to an organization account, root, or unit.

{{% examples %}}
## Example Usage
{{% example %}}
### Organization Account

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const account = new aws.organizations.PolicyAttachment("account", {
    policyId: aws_organizations_policy.example.id,
    targetId: "123456789012",
});
```
```python
import pulumi
import pulumi_aws as aws

account = aws.organizations.PolicyAttachment("account",
    policy_id=aws_organizations_policy["example"]["id"],
    target_id="123456789012")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var account = new Aws.Organizations.PolicyAttachment("account", new()
    {
        PolicyId = aws_organizations_policy.Example.Id,
        TargetId = "123456789012",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewPolicyAttachment(ctx, "account", &organizations.PolicyAttachmentArgs{
			PolicyId: pulumi.Any(aws_organizations_policy.Example.Id),
			TargetId: pulumi.String("123456789012"),
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
import com.pulumi.aws.organizations.PolicyAttachment;
import com.pulumi.aws.organizations.PolicyAttachmentArgs;
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
        var account = new PolicyAttachment("account", PolicyAttachmentArgs.builder()        
            .policyId(aws_organizations_policy.example().id())
            .targetId("123456789012")
            .build());

    }
}
```
```yaml
resources:
  account:
    type: aws:organizations:PolicyAttachment
    properties:
      policyId: ${aws_organizations_policy.example.id}
      targetId: '123456789012'
```
{{% /example %}}
{{% example %}}
### Organization Root

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const root = new aws.organizations.PolicyAttachment("root", {
    policyId: aws_organizations_policy.example.id,
    targetId: aws_organizations_organization.example.roots[0].id,
});
```
```python
import pulumi
import pulumi_aws as aws

root = aws.organizations.PolicyAttachment("root",
    policy_id=aws_organizations_policy["example"]["id"],
    target_id=aws_organizations_organization["example"]["roots"][0]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var root = new Aws.Organizations.PolicyAttachment("root", new()
    {
        PolicyId = aws_organizations_policy.Example.Id,
        TargetId = aws_organizations_organization.Example.Roots[0].Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewPolicyAttachment(ctx, "root", &organizations.PolicyAttachmentArgs{
			PolicyId: pulumi.Any(aws_organizations_policy.Example.Id),
			TargetId: pulumi.Any(aws_organizations_organization.Example.Roots[0].Id),
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
import com.pulumi.aws.organizations.PolicyAttachment;
import com.pulumi.aws.organizations.PolicyAttachmentArgs;
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
        var root = new PolicyAttachment("root", PolicyAttachmentArgs.builder()        
            .policyId(aws_organizations_policy.example().id())
            .targetId(aws_organizations_organization.example().roots()[0].id())
            .build());

    }
}
```
```yaml
resources:
  root:
    type: aws:organizations:PolicyAttachment
    properties:
      policyId: ${aws_organizations_policy.example.id}
      targetId: ${aws_organizations_organization.example.roots[0].id}
```
{{% /example %}}
{{% example %}}
### Organization Unit

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const unit = new aws.organizations.PolicyAttachment("unit", {
    policyId: aws_organizations_policy.example.id,
    targetId: aws_organizations_organizational_unit.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

unit = aws.organizations.PolicyAttachment("unit",
    policy_id=aws_organizations_policy["example"]["id"],
    target_id=aws_organizations_organizational_unit["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var unit = new Aws.Organizations.PolicyAttachment("unit", new()
    {
        PolicyId = aws_organizations_policy.Example.Id,
        TargetId = aws_organizations_organizational_unit.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := organizations.NewPolicyAttachment(ctx, "unit", &organizations.PolicyAttachmentArgs{
			PolicyId: pulumi.Any(aws_organizations_policy.Example.Id),
			TargetId: pulumi.Any(aws_organizations_organizational_unit.Example.Id),
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
import com.pulumi.aws.organizations.PolicyAttachment;
import com.pulumi.aws.organizations.PolicyAttachmentArgs;
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
        var unit = new PolicyAttachment("unit", PolicyAttachmentArgs.builder()        
            .policyId(aws_organizations_policy.example().id())
            .targetId(aws_organizations_organizational_unit.example().id())
            .build());

    }
}
```
```yaml
resources:
  unit:
    type: aws:organizations:PolicyAttachment
    properties:
      policyId: ${aws_organizations_policy.example.id}
      targetId: ${aws_organizations_organizational_unit.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_organizations_policy_attachment` using the target ID and policy ID. For example:

With an account target:

```sh
 $ pulumi import aws:organizations/policyAttachment:PolicyAttachment account 123456789012:p-12345678
```
 