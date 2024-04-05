Provides a resource to create an EventBridge permission to support cross-account events in the current account default event bus.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

> **Note:** The EventBridge bus policy resource  (`aws.cloudwatch.EventBusPolicy`) is incompatible with the EventBridge permission resource (`aws.cloudwatch.EventPermission`) and will overwrite permissions.

{{% examples %}}
## Example Usage
{{% example %}}
### Account Access

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const devAccountAccess = new aws.cloudwatch.EventPermission("devAccountAccess", {
    principal: "123456789012",
    statementId: "DevAccountAccess",
});
```
```python
import pulumi
import pulumi_aws as aws

dev_account_access = aws.cloudwatch.EventPermission("devAccountAccess",
    principal="123456789012",
    statement_id="DevAccountAccess")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var devAccountAccess = new Aws.CloudWatch.EventPermission("devAccountAccess", new()
    {
        Principal = "123456789012",
        StatementId = "DevAccountAccess",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewEventPermission(ctx, "devAccountAccess", &cloudwatch.EventPermissionArgs{
			Principal:   pulumi.String("123456789012"),
			StatementId: pulumi.String("DevAccountAccess"),
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
import com.pulumi.aws.cloudwatch.EventPermission;
import com.pulumi.aws.cloudwatch.EventPermissionArgs;
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
        var devAccountAccess = new EventPermission("devAccountAccess", EventPermissionArgs.builder()        
            .principal("123456789012")
            .statementId("DevAccountAccess")
            .build());

    }
}
```
```yaml
resources:
  devAccountAccess:
    type: aws:cloudwatch:EventPermission
    properties:
      principal: '123456789012'
      statementId: DevAccountAccess
```
{{% /example %}}
{{% example %}}
### Organization Access

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const organizationAccess = new aws.cloudwatch.EventPermission("organizationAccess", {
    principal: "*",
    statementId: "OrganizationAccess",
    condition: {
        key: "aws:PrincipalOrgID",
        type: "StringEquals",
        value: aws_organizations_organization.example.id,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

organization_access = aws.cloudwatch.EventPermission("organizationAccess",
    principal="*",
    statement_id="OrganizationAccess",
    condition=aws.cloudwatch.EventPermissionConditionArgs(
        key="aws:PrincipalOrgID",
        type="StringEquals",
        value=aws_organizations_organization["example"]["id"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var organizationAccess = new Aws.CloudWatch.EventPermission("organizationAccess", new()
    {
        Principal = "*",
        StatementId = "OrganizationAccess",
        Condition = new Aws.CloudWatch.Inputs.EventPermissionConditionArgs
        {
            Key = "aws:PrincipalOrgID",
            Type = "StringEquals",
            Value = aws_organizations_organization.Example.Id,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewEventPermission(ctx, "organizationAccess", &cloudwatch.EventPermissionArgs{
			Principal:   pulumi.String("*"),
			StatementId: pulumi.String("OrganizationAccess"),
			Condition: &cloudwatch.EventPermissionConditionArgs{
				Key:   pulumi.String("aws:PrincipalOrgID"),
				Type:  pulumi.String("StringEquals"),
				Value: pulumi.Any(aws_organizations_organization.Example.Id),
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
import com.pulumi.aws.cloudwatch.EventPermission;
import com.pulumi.aws.cloudwatch.EventPermissionArgs;
import com.pulumi.aws.cloudwatch.inputs.EventPermissionConditionArgs;
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
        var organizationAccess = new EventPermission("organizationAccess", EventPermissionArgs.builder()        
            .principal("*")
            .statementId("OrganizationAccess")
            .condition(EventPermissionConditionArgs.builder()
                .key("aws:PrincipalOrgID")
                .type("StringEquals")
                .value(aws_organizations_organization.example().id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  organizationAccess:
    type: aws:cloudwatch:EventPermission
    properties:
      principal: '*'
      statementId: OrganizationAccess
      condition:
        key: aws:PrincipalOrgID
        type: StringEquals
        value: ${aws_organizations_organization.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EventBridge permissions using the `event_bus_name/statement_id` (if you omit `event_bus_name`, the `default` event bus will be used). For example:

```sh
 $ pulumi import aws:cloudwatch/eventPermission:EventPermission DevAccountAccess example-event-bus/DevAccountAccess
```
 