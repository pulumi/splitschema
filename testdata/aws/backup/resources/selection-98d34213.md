Manages selection conditions for AWS Backup plan resources.

{{% examples %}}
## Example Usage
{{% example %}}
### IAM Role

> For more information about creating and managing IAM Roles for backups and restores, see the [AWS Backup Developer Guide](https://docs.aws.amazon.com/aws-backup/latest/devguide/iam-service-roles.html).

The below example creates an IAM role with the default managed IAM Policy for allowing AWS Backup to create backups.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["backup.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const exampleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment", {
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup",
    role: exampleRole.name,
});
// ... other configuration ...
const exampleSelection = new aws.backup.Selection("exampleSelection", {iamRoleArn: exampleRole.arn});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["backup.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
example_role_policy_attachment = aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment",
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup",
    role=example_role.name)
# ... other configuration ...
example_selection = aws.backup.Selection("exampleSelection", iam_role_arn=example_role.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "backup.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("exampleRolePolicyAttachment", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup",
        Role = exampleRole.Name,
    });

    // ... other configuration ...
    var exampleSelection = new Aws.Backup.Selection("exampleSelection", new()
    {
        IamRoleArn = exampleRole.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"backup.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup"),
			Role:      exampleRole.Name,
		})
		if err != nil {
			return err
		}
		_, err = backup.NewSelection(ctx, "exampleSelection", &backup.SelectionArgs{
			IamRoleArn: exampleRole.Arn,
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.backup.Selection;
import com.pulumi.aws.backup.SelectionArgs;
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
        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("backup.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleRolePolicyAttachment = new RolePolicyAttachment("exampleRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup")
            .role(exampleRole.name())
            .build());

        var exampleSelection = new Selection("exampleSelection", SelectionArgs.builder()        
            .iamRoleArn(exampleRole.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  exampleRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/service-role/AWSBackupServiceRolePolicyForBackup
      role: ${exampleRole.name}
  exampleSelection:
    type: aws:backup:Selection
    properties:
      iamRoleArn: ${exampleRole.arn}
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - backup.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% example %}}
### Selecting Backups By Tag

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.backup.Selection("example", {
    iamRoleArn: aws_iam_role.example.arn,
    planId: aws_backup_plan.example.id,
    selectionTags: [{
        type: "STRINGEQUALS",
        key: "foo",
        value: "bar",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.backup.Selection("example",
    iam_role_arn=aws_iam_role["example"]["arn"],
    plan_id=aws_backup_plan["example"]["id"],
    selection_tags=[aws.backup.SelectionSelectionTagArgs(
        type="STRINGEQUALS",
        key="foo",
        value="bar",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Backup.Selection("example", new()
    {
        IamRoleArn = aws_iam_role.Example.Arn,
        PlanId = aws_backup_plan.Example.Id,
        SelectionTags = new[]
        {
            new Aws.Backup.Inputs.SelectionSelectionTagArgs
            {
                Type = "STRINGEQUALS",
                Key = "foo",
                Value = "bar",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.NewSelection(ctx, "example", &backup.SelectionArgs{
			IamRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			PlanId:     pulumi.Any(aws_backup_plan.Example.Id),
			SelectionTags: backup.SelectionSelectionTagArray{
				&backup.SelectionSelectionTagArgs{
					Type:  pulumi.String("STRINGEQUALS"),
					Key:   pulumi.String("foo"),
					Value: pulumi.String("bar"),
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
import com.pulumi.aws.backup.Selection;
import com.pulumi.aws.backup.SelectionArgs;
import com.pulumi.aws.backup.inputs.SelectionSelectionTagArgs;
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
        var example = new Selection("example", SelectionArgs.builder()        
            .iamRoleArn(aws_iam_role.example().arn())
            .planId(aws_backup_plan.example().id())
            .selectionTags(SelectionSelectionTagArgs.builder()
                .type("STRINGEQUALS")
                .key("foo")
                .value("bar")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:backup:Selection
    properties:
      iamRoleArn: ${aws_iam_role.example.arn}
      planId: ${aws_backup_plan.example.id}
      selectionTags:
        - type: STRINGEQUALS
          key: foo
          value: bar
```
{{% /example %}}
{{% example %}}
### Selecting Backups By Conditions

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.backup.Selection("example", {
    iamRoleArn: aws_iam_role.example.arn,
    planId: aws_backup_plan.example.id,
    resources: ["*"],
    conditions: [{
        stringEquals: [{
            key: "aws:ResourceTag/Component",
            value: "rds",
        }],
        stringLikes: [{
            key: "aws:ResourceTag/Application",
            value: "app*",
        }],
        stringNotEquals: [{
            key: "aws:ResourceTag/Backup",
            value: "false",
        }],
        stringNotLikes: [{
            key: "aws:ResourceTag/Environment",
            value: "test*",
        }],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.backup.Selection("example",
    iam_role_arn=aws_iam_role["example"]["arn"],
    plan_id=aws_backup_plan["example"]["id"],
    resources=["*"],
    conditions=[aws.backup.SelectionConditionArgs(
        string_equals=[aws.backup.SelectionConditionStringEqualArgs(
            key="aws:ResourceTag/Component",
            value="rds",
        )],
        string_likes=[aws.backup.SelectionConditionStringLikeArgs(
            key="aws:ResourceTag/Application",
            value="app*",
        )],
        string_not_equals=[aws.backup.SelectionConditionStringNotEqualArgs(
            key="aws:ResourceTag/Backup",
            value="false",
        )],
        string_not_likes=[aws.backup.SelectionConditionStringNotLikeArgs(
            key="aws:ResourceTag/Environment",
            value="test*",
        )],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Backup.Selection("example", new()
    {
        IamRoleArn = aws_iam_role.Example.Arn,
        PlanId = aws_backup_plan.Example.Id,
        Resources = new[]
        {
            "*",
        },
        Conditions = new[]
        {
            new Aws.Backup.Inputs.SelectionConditionArgs
            {
                StringEquals = new[]
                {
                    new Aws.Backup.Inputs.SelectionConditionStringEqualArgs
                    {
                        Key = "aws:ResourceTag/Component",
                        Value = "rds",
                    },
                },
                StringLikes = new[]
                {
                    new Aws.Backup.Inputs.SelectionConditionStringLikeArgs
                    {
                        Key = "aws:ResourceTag/Application",
                        Value = "app*",
                    },
                },
                StringNotEquals = new[]
                {
                    new Aws.Backup.Inputs.SelectionConditionStringNotEqualArgs
                    {
                        Key = "aws:ResourceTag/Backup",
                        Value = "false",
                    },
                },
                StringNotLikes = new[]
                {
                    new Aws.Backup.Inputs.SelectionConditionStringNotLikeArgs
                    {
                        Key = "aws:ResourceTag/Environment",
                        Value = "test*",
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.NewSelection(ctx, "example", &backup.SelectionArgs{
			IamRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			PlanId:     pulumi.Any(aws_backup_plan.Example.Id),
			Resources: pulumi.StringArray{
				pulumi.String("*"),
			},
			Conditions: backup.SelectionConditionArray{
				&backup.SelectionConditionArgs{
					StringEquals: backup.SelectionConditionStringEqualArray{
						&backup.SelectionConditionStringEqualArgs{
							Key:   pulumi.String("aws:ResourceTag/Component"),
							Value: pulumi.String("rds"),
						},
					},
					StringLikes: backup.SelectionConditionStringLikeArray{
						&backup.SelectionConditionStringLikeArgs{
							Key:   pulumi.String("aws:ResourceTag/Application"),
							Value: pulumi.String("app*"),
						},
					},
					StringNotEquals: backup.SelectionConditionStringNotEqualArray{
						&backup.SelectionConditionStringNotEqualArgs{
							Key:   pulumi.String("aws:ResourceTag/Backup"),
							Value: pulumi.String("false"),
						},
					},
					StringNotLikes: backup.SelectionConditionStringNotLikeArray{
						&backup.SelectionConditionStringNotLikeArgs{
							Key:   pulumi.String("aws:ResourceTag/Environment"),
							Value: pulumi.String("test*"),
						},
					},
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
import com.pulumi.aws.backup.Selection;
import com.pulumi.aws.backup.SelectionArgs;
import com.pulumi.aws.backup.inputs.SelectionConditionArgs;
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
        var example = new Selection("example", SelectionArgs.builder()        
            .iamRoleArn(aws_iam_role.example().arn())
            .planId(aws_backup_plan.example().id())
            .resources("*")
            .conditions(SelectionConditionArgs.builder()
                .stringEquals(SelectionConditionStringEqualArgs.builder()
                    .key("aws:ResourceTag/Component")
                    .value("rds")
                    .build())
                .stringLikes(SelectionConditionStringLikeArgs.builder()
                    .key("aws:ResourceTag/Application")
                    .value("app*")
                    .build())
                .stringNotEquals(SelectionConditionStringNotEqualArgs.builder()
                    .key("aws:ResourceTag/Backup")
                    .value("false")
                    .build())
                .stringNotLikes(SelectionConditionStringNotLikeArgs.builder()
                    .key("aws:ResourceTag/Environment")
                    .value("test*")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:backup:Selection
    properties:
      iamRoleArn: ${aws_iam_role.example.arn}
      planId: ${aws_backup_plan.example.id}
      resources:
        - '*'
      conditions:
        - stringEquals:
            - key: aws:ResourceTag/Component
              value: rds
          stringLikes:
            - key: aws:ResourceTag/Application
              value: app*
          stringNotEquals:
            - key: aws:ResourceTag/Backup
              value: 'false'
          stringNotLikes:
            - key: aws:ResourceTag/Environment
              value: test*
```
{{% /example %}}
{{% example %}}
### Selecting Backups By Resource

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.backup.Selection("example", {
    iamRoleArn: aws_iam_role.example.arn,
    planId: aws_backup_plan.example.id,
    resources: [
        aws_db_instance.example.arn,
        aws_ebs_volume.example.arn,
        aws_efs_file_system.example.arn,
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.backup.Selection("example",
    iam_role_arn=aws_iam_role["example"]["arn"],
    plan_id=aws_backup_plan["example"]["id"],
    resources=[
        aws_db_instance["example"]["arn"],
        aws_ebs_volume["example"]["arn"],
        aws_efs_file_system["example"]["arn"],
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Backup.Selection("example", new()
    {
        IamRoleArn = aws_iam_role.Example.Arn,
        PlanId = aws_backup_plan.Example.Id,
        Resources = new[]
        {
            aws_db_instance.Example.Arn,
            aws_ebs_volume.Example.Arn,
            aws_efs_file_system.Example.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.NewSelection(ctx, "example", &backup.SelectionArgs{
			IamRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			PlanId:     pulumi.Any(aws_backup_plan.Example.Id),
			Resources: pulumi.StringArray{
				aws_db_instance.Example.Arn,
				aws_ebs_volume.Example.Arn,
				aws_efs_file_system.Example.Arn,
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
import com.pulumi.aws.backup.Selection;
import com.pulumi.aws.backup.SelectionArgs;
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
        var example = new Selection("example", SelectionArgs.builder()        
            .iamRoleArn(aws_iam_role.example().arn())
            .planId(aws_backup_plan.example().id())
            .resources(            
                aws_db_instance.example().arn(),
                aws_ebs_volume.example().arn(),
                aws_efs_file_system.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:backup:Selection
    properties:
      iamRoleArn: ${aws_iam_role.example.arn}
      planId: ${aws_backup_plan.example.id}
      resources:
        - ${aws_db_instance.example.arn}
        - ${aws_ebs_volume.example.arn}
        - ${aws_efs_file_system.example.arn}
```
{{% /example %}}
{{% example %}}
### Selecting Backups By Not Resource

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.backup.Selection("example", {
    iamRoleArn: aws_iam_role.example.arn,
    planId: aws_backup_plan.example.id,
    notResources: [
        aws_db_instance.example.arn,
        aws_ebs_volume.example.arn,
        aws_efs_file_system.example.arn,
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.backup.Selection("example",
    iam_role_arn=aws_iam_role["example"]["arn"],
    plan_id=aws_backup_plan["example"]["id"],
    not_resources=[
        aws_db_instance["example"]["arn"],
        aws_ebs_volume["example"]["arn"],
        aws_efs_file_system["example"]["arn"],
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Backup.Selection("example", new()
    {
        IamRoleArn = aws_iam_role.Example.Arn,
        PlanId = aws_backup_plan.Example.Id,
        NotResources = new[]
        {
            aws_db_instance.Example.Arn,
            aws_ebs_volume.Example.Arn,
            aws_efs_file_system.Example.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.NewSelection(ctx, "example", &backup.SelectionArgs{
			IamRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			PlanId:     pulumi.Any(aws_backup_plan.Example.Id),
			NotResources: pulumi.StringArray{
				aws_db_instance.Example.Arn,
				aws_ebs_volume.Example.Arn,
				aws_efs_file_system.Example.Arn,
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
import com.pulumi.aws.backup.Selection;
import com.pulumi.aws.backup.SelectionArgs;
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
        var example = new Selection("example", SelectionArgs.builder()        
            .iamRoleArn(aws_iam_role.example().arn())
            .planId(aws_backup_plan.example().id())
            .notResources(            
                aws_db_instance.example().arn(),
                aws_ebs_volume.example().arn(),
                aws_efs_file_system.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:backup:Selection
    properties:
      iamRoleArn: ${aws_iam_role.example.arn}
      planId: ${aws_backup_plan.example.id}
      notResources:
        - ${aws_db_instance.example.arn}
        - ${aws_ebs_volume.example.arn}
        - ${aws_efs_file_system.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Backup selection using the role plan_id and id separated by `|`. For example:

```sh
 $ pulumi import aws:backup/selection:Selection example plan-id|selection-id
```
 