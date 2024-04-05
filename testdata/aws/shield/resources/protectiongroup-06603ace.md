Creates a grouping of protected resources so they can be handled as a collective.
This resource grouping improves the accuracy of detection and reduces false positives. For more information see
[Managing AWS Shield Advanced protection groups](https://docs.aws.amazon.com/waf/latest/developerguide/manage-protection-group.html)

{{% examples %}}
## Example Usage
{{% example %}}
### Create protection group for all resources

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.shield.ProtectionGroup("example", {
    aggregation: "MAX",
    pattern: "ALL",
    protectionGroupId: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.shield.ProtectionGroup("example",
    aggregation="MAX",
    pattern="ALL",
    protection_group_id="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Shield.ProtectionGroup("example", new()
    {
        Aggregation = "MAX",
        Pattern = "ALL",
        ProtectionGroupId = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/shield"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := shield.NewProtectionGroup(ctx, "example", &shield.ProtectionGroupArgs{
			Aggregation:       pulumi.String("MAX"),
			Pattern:           pulumi.String("ALL"),
			ProtectionGroupId: pulumi.String("example"),
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
import com.pulumi.aws.shield.ProtectionGroup;
import com.pulumi.aws.shield.ProtectionGroupArgs;
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
        var example = new ProtectionGroup("example", ProtectionGroupArgs.builder()        
            .aggregation("MAX")
            .pattern("ALL")
            .protectionGroupId("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:shield:ProtectionGroup
    properties:
      aggregation: MAX
      pattern: ALL
      protectionGroupId: example
```
{{% /example %}}
{{% example %}}
### Create protection group for arbitrary number of resources

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentRegion = aws.getRegion({});
const currentCallerIdentity = aws.getCallerIdentity({});
const exampleEip = new aws.ec2.Eip("exampleEip", {domain: "vpc"});
const exampleProtection = new aws.shield.Protection("exampleProtection", {resourceArn: pulumi.all([currentRegion, currentCallerIdentity, exampleEip.id]).apply(([currentRegion, currentCallerIdentity, id]) => `arn:aws:ec2:${currentRegion.name}:${currentCallerIdentity.accountId}:eip-allocation/${id}`)});
const exampleProtectionGroup = new aws.shield.ProtectionGroup("exampleProtectionGroup", {
    protectionGroupId: "example",
    aggregation: "MEAN",
    pattern: "ARBITRARY",
    members: [pulumi.all([currentRegion, currentCallerIdentity, exampleEip.id]).apply(([currentRegion, currentCallerIdentity, id]) => `arn:aws:ec2:${currentRegion.name}:${currentCallerIdentity.accountId}:eip-allocation/${id}`)],
}, {
    dependsOn: [exampleProtection],
});
```
```python
import pulumi
import pulumi_aws as aws

current_region = aws.get_region()
current_caller_identity = aws.get_caller_identity()
example_eip = aws.ec2.Eip("exampleEip", domain="vpc")
example_protection = aws.shield.Protection("exampleProtection", resource_arn=example_eip.id.apply(lambda id: f"arn:aws:ec2:{current_region.name}:{current_caller_identity.account_id}:eip-allocation/{id}"))
example_protection_group = aws.shield.ProtectionGroup("exampleProtectionGroup",
    protection_group_id="example",
    aggregation="MEAN",
    pattern="ARBITRARY",
    members=[example_eip.id.apply(lambda id: f"arn:aws:ec2:{current_region.name}:{current_caller_identity.account_id}:eip-allocation/{id}")],
    opts=pulumi.ResourceOptions(depends_on=[example_protection]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var currentRegion = Aws.GetRegion.Invoke();

    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var exampleEip = new Aws.Ec2.Eip("exampleEip", new()
    {
        Domain = "vpc",
    });

    var exampleProtection = new Aws.Shield.Protection("exampleProtection", new()
    {
        ResourceArn = Output.Tuple(currentRegion, currentCallerIdentity, exampleEip.Id).Apply(values =>
        {
            var currentRegion = values.Item1;
            var currentCallerIdentity = values.Item2;
            var id = values.Item3;
            return $"arn:aws:ec2:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:eip-allocation/{id}";
        }),
    });

    var exampleProtectionGroup = new Aws.Shield.ProtectionGroup("exampleProtectionGroup", new()
    {
        ProtectionGroupId = "example",
        Aggregation = "MEAN",
        Pattern = "ARBITRARY",
        Members = new[]
        {
            Output.Tuple(currentRegion, currentCallerIdentity, exampleEip.Id).Apply(values =>
            {
                var currentRegion = values.Item1;
                var currentCallerIdentity = values.Item2;
                var id = values.Item3;
                return $"arn:aws:ec2:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:eip-allocation/{id}";
            }),
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleProtection,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/shield"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleEip, err := ec2.NewEip(ctx, "exampleEip", &ec2.EipArgs{
			Domain: pulumi.String("vpc"),
		})
		if err != nil {
			return err
		}
		exampleProtection, err := shield.NewProtection(ctx, "exampleProtection", &shield.ProtectionArgs{
			ResourceArn: exampleEip.ID().ApplyT(func(id string) (string, error) {
				return fmt.Sprintf("arn:aws:ec2:%v:%v:eip-allocation/%v", currentRegion.Name, currentCallerIdentity.AccountId, id), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = shield.NewProtectionGroup(ctx, "exampleProtectionGroup", &shield.ProtectionGroupArgs{
			ProtectionGroupId: pulumi.String("example"),
			Aggregation:       pulumi.String("MEAN"),
			Pattern:           pulumi.String("ARBITRARY"),
			Members: pulumi.StringArray{
				exampleEip.ID().ApplyT(func(id string) (string, error) {
					return fmt.Sprintf("arn:aws:ec2:%v:%v:eip-allocation/%v", currentRegion.Name, currentCallerIdentity.AccountId, id), nil
				}).(pulumi.StringOutput),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleProtection,
		}))
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.ec2.Eip;
import com.pulumi.aws.ec2.EipArgs;
import com.pulumi.aws.shield.Protection;
import com.pulumi.aws.shield.ProtectionArgs;
import com.pulumi.aws.shield.ProtectionGroup;
import com.pulumi.aws.shield.ProtectionGroupArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        final var currentRegion = AwsFunctions.getRegion();

        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        var exampleEip = new Eip("exampleEip", EipArgs.builder()        
            .domain("vpc")
            .build());

        var exampleProtection = new Protection("exampleProtection", ProtectionArgs.builder()        
            .resourceArn(exampleEip.id().applyValue(id -> String.format("arn:aws:ec2:%s:%s:eip-allocation/%s", currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()),id)))
            .build());

        var exampleProtectionGroup = new ProtectionGroup("exampleProtectionGroup", ProtectionGroupArgs.builder()        
            .protectionGroupId("example")
            .aggregation("MEAN")
            .pattern("ARBITRARY")
            .members(exampleEip.id().applyValue(id -> String.format("arn:aws:ec2:%s:%s:eip-allocation/%s", currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()),id)))
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleProtection)
                .build());

    }
}
```
```yaml
resources:
  exampleEip:
    type: aws:ec2:Eip
    properties:
      domain: vpc
  exampleProtection:
    type: aws:shield:Protection
    properties:
      resourceArn: arn:aws:ec2:${currentRegion.name}:${currentCallerIdentity.accountId}:eip-allocation/${exampleEip.id}
  exampleProtectionGroup:
    type: aws:shield:ProtectionGroup
    properties:
      protectionGroupId: example
      aggregation: MEAN
      pattern: ARBITRARY
      members:
        - arn:aws:ec2:${currentRegion.name}:${currentCallerIdentity.accountId}:eip-allocation/${exampleEip.id}
    options:
      dependson:
        - ${exampleProtection}
variables:
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### Create protection group for a type of resource

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.shield.ProtectionGroup("example", {
    aggregation: "SUM",
    pattern: "BY_RESOURCE_TYPE",
    protectionGroupId: "example",
    resourceType: "ELASTIC_IP_ALLOCATION",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.shield.ProtectionGroup("example",
    aggregation="SUM",
    pattern="BY_RESOURCE_TYPE",
    protection_group_id="example",
    resource_type="ELASTIC_IP_ALLOCATION")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Shield.ProtectionGroup("example", new()
    {
        Aggregation = "SUM",
        Pattern = "BY_RESOURCE_TYPE",
        ProtectionGroupId = "example",
        ResourceType = "ELASTIC_IP_ALLOCATION",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/shield"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := shield.NewProtectionGroup(ctx, "example", &shield.ProtectionGroupArgs{
			Aggregation:       pulumi.String("SUM"),
			Pattern:           pulumi.String("BY_RESOURCE_TYPE"),
			ProtectionGroupId: pulumi.String("example"),
			ResourceType:      pulumi.String("ELASTIC_IP_ALLOCATION"),
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
import com.pulumi.aws.shield.ProtectionGroup;
import com.pulumi.aws.shield.ProtectionGroupArgs;
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
        var example = new ProtectionGroup("example", ProtectionGroupArgs.builder()        
            .aggregation("SUM")
            .pattern("BY_RESOURCE_TYPE")
            .protectionGroupId("example")
            .resourceType("ELASTIC_IP_ALLOCATION")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:shield:ProtectionGroup
    properties:
      aggregation: SUM
      pattern: BY_RESOURCE_TYPE
      protectionGroupId: example
      resourceType: ELASTIC_IP_ALLOCATION
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Shield protection group resources using their protection group id. For example:

```sh
 $ pulumi import aws:shield/protectionGroup:ProtectionGroup example example
```
 