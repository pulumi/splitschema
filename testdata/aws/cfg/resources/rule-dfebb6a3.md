Provides an AWS Config Rule.

> **Note:** Config Rule requires an existing Configuration Recorder to be present. Use of `depends_on` is recommended (as shown below) to avoid race conditions.

{{% examples %}}
## Example Usage
{{% example %}}
### AWS Managed Rules

AWS managed rules can be used by setting the source owner to `AWS` and the source identifier to the name of the managed rule. More information about AWS managed rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["config.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const role = new aws.iam.Role("role", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const foo = new aws.cfg.Recorder("foo", {roleArn: role.arn});
const rule = new aws.cfg.Rule("rule", {source: {
    owner: "AWS",
    sourceIdentifier: "S3_BUCKET_VERSIONING_ENABLED",
}}, {
    dependsOn: [foo],
});
const policyDocument = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        actions: ["config:Put*"],
        resources: ["*"],
    }],
});
const rolePolicy = new aws.iam.RolePolicy("rolePolicy", {
    role: role.id,
    policy: policyDocument.then(policyDocument => policyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["config.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
role = aws.iam.Role("role", assume_role_policy=assume_role.json)
foo = aws.cfg.Recorder("foo", role_arn=role.arn)
rule = aws.cfg.Rule("rule", source=aws.cfg.RuleSourceArgs(
    owner="AWS",
    source_identifier="S3_BUCKET_VERSIONING_ENABLED",
),
opts=pulumi.ResourceOptions(depends_on=[foo]))
policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["config:Put*"],
    resources=["*"],
)])
role_policy = aws.iam.RolePolicy("rolePolicy",
    role=role.id,
    policy=policy_document.json)
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
                            "config.amazonaws.com",
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

    var role = new Aws.Iam.Role("role", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var foo = new Aws.Cfg.Recorder("foo", new()
    {
        RoleArn = role.Arn,
    });

    var rule = new Aws.Cfg.Rule("rule", new()
    {
        Source = new Aws.Cfg.Inputs.RuleSourceArgs
        {
            Owner = "AWS",
            SourceIdentifier = "S3_BUCKET_VERSIONING_ENABLED",
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            foo,
        },
    });

    var policyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "config:Put*",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var rolePolicy = new Aws.Iam.RolePolicy("rolePolicy", new()
    {
        Role = role.Id,
        Policy = policyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
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
								"config.amazonaws.com",
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
		role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		foo, err := cfg.NewRecorder(ctx, "foo", &cfg.RecorderArgs{
			RoleArn: role.Arn,
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewRule(ctx, "rule", &cfg.RuleArgs{
			Source: &cfg.RuleSourceArgs{
				Owner:            pulumi.String("AWS"),
				SourceIdentifier: pulumi.String("S3_BUCKET_VERSIONING_ENABLED"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			foo,
		}))
		if err != nil {
			return err
		}
		policyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Actions: []string{
						"config:Put*",
					},
					Resources: []string{
						"*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicy(ctx, "rolePolicy", &iam.RolePolicyArgs{
			Role:   role.ID(),
			Policy: *pulumi.String(policyDocument.Json),
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
import com.pulumi.aws.cfg.Recorder;
import com.pulumi.aws.cfg.RecorderArgs;
import com.pulumi.aws.cfg.Rule;
import com.pulumi.aws.cfg.RuleArgs;
import com.pulumi.aws.cfg.inputs.RuleSourceArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
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
        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("config.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var role = new Role("role", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var foo = new Recorder("foo", RecorderArgs.builder()        
            .roleArn(role.arn())
            .build());

        var rule = new Rule("rule", RuleArgs.builder()        
            .source(RuleSourceArgs.builder()
                .owner("AWS")
                .sourceIdentifier("S3_BUCKET_VERSIONING_ENABLED")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(foo)
                .build());

        final var policyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("config:Put*")
                .resources("*")
                .build())
            .build());

        var rolePolicy = new RolePolicy("rolePolicy", RolePolicyArgs.builder()        
            .role(role.id())
            .policy(policyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

    }
}
```
```yaml
resources:
  rule:
    type: aws:cfg:Rule
    properties:
      source:
        owner: AWS
        sourceIdentifier: S3_BUCKET_VERSIONING_ENABLED
    options:
      dependson:
        - ${foo}
  foo:
    type: aws:cfg:Recorder
    properties:
      roleArn: ${role.arn}
  role:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  rolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${role.id}
      policy: ${policyDocument.json}
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
                  - config.amazonaws.com
            actions:
              - sts:AssumeRole
  policyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - config:Put*
            resources:
              - '*'
```
{{% /example %}}
{{% example %}}
### Custom Rules

Custom rules can be used by setting the source owner to `CUSTOM_LAMBDA` and the source identifier to the Amazon Resource Name (ARN) of the Lambda Function. The AWS Config service must have permissions to invoke the Lambda Function, e.g., via the `aws.lambda.Permission` resource. More information about custom rules can be found in the [AWS Config Developer Guide](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules.html).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRecorder = new aws.cfg.Recorder("exampleRecorder", {});
// ... other configuration ...
const exampleFunction = new aws.lambda.Function("exampleFunction", {});
// ... other configuration ...
const examplePermission = new aws.lambda.Permission("examplePermission", {
    action: "lambda:InvokeFunction",
    "function": exampleFunction.arn,
    principal: "config.amazonaws.com",
});
// ... other configuration ...
const exampleRule = new aws.cfg.Rule("exampleRule", {source: {
    owner: "CUSTOM_LAMBDA",
    sourceIdentifier: exampleFunction.arn,
}}, {
    dependsOn: [
        exampleRecorder,
        examplePermission,
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_recorder = aws.cfg.Recorder("exampleRecorder")
# ... other configuration ...
example_function = aws.lambda_.Function("exampleFunction")
# ... other configuration ...
example_permission = aws.lambda_.Permission("examplePermission",
    action="lambda:InvokeFunction",
    function=example_function.arn,
    principal="config.amazonaws.com")
# ... other configuration ...
example_rule = aws.cfg.Rule("exampleRule", source=aws.cfg.RuleSourceArgs(
    owner="CUSTOM_LAMBDA",
    source_identifier=example_function.arn,
),
opts=pulumi.ResourceOptions(depends_on=[
        example_recorder,
        example_permission,
    ]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleRecorder = new Aws.Cfg.Recorder("exampleRecorder");

    // ... other configuration ...
    var exampleFunction = new Aws.Lambda.Function("exampleFunction");

    // ... other configuration ...
    var examplePermission = new Aws.Lambda.Permission("examplePermission", new()
    {
        Action = "lambda:InvokeFunction",
        Function = exampleFunction.Arn,
        Principal = "config.amazonaws.com",
    });

    // ... other configuration ...
    var exampleRule = new Aws.Cfg.Rule("exampleRule", new()
    {
        Source = new Aws.Cfg.Inputs.RuleSourceArgs
        {
            Owner = "CUSTOM_LAMBDA",
            SourceIdentifier = exampleFunction.Arn,
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleRecorder,
            examplePermission,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleRecorder, err := cfg.NewRecorder(ctx, "exampleRecorder", nil)
		if err != nil {
			return err
		}
		exampleFunction, err := lambda.NewFunction(ctx, "exampleFunction", nil)
		if err != nil {
			return err
		}
		examplePermission, err := lambda.NewPermission(ctx, "examplePermission", &lambda.PermissionArgs{
			Action:    pulumi.String("lambda:InvokeFunction"),
			Function:  exampleFunction.Arn,
			Principal: pulumi.String("config.amazonaws.com"),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewRule(ctx, "exampleRule", &cfg.RuleArgs{
			Source: &cfg.RuleSourceArgs{
				Owner:            pulumi.String("CUSTOM_LAMBDA"),
				SourceIdentifier: exampleFunction.Arn,
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleRecorder,
			examplePermission,
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
import com.pulumi.aws.cfg.Recorder;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.Permission;
import com.pulumi.aws.lambda.PermissionArgs;
import com.pulumi.aws.cfg.Rule;
import com.pulumi.aws.cfg.RuleArgs;
import com.pulumi.aws.cfg.inputs.RuleSourceArgs;
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
        var exampleRecorder = new Recorder("exampleRecorder");

        var exampleFunction = new Function("exampleFunction");

        var examplePermission = new Permission("examplePermission", PermissionArgs.builder()        
            .action("lambda:InvokeFunction")
            .function(exampleFunction.arn())
            .principal("config.amazonaws.com")
            .build());

        var exampleRule = new Rule("exampleRule", RuleArgs.builder()        
            .source(RuleSourceArgs.builder()
                .owner("CUSTOM_LAMBDA")
                .sourceIdentifier(exampleFunction.arn())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    exampleRecorder,
                    examplePermission)
                .build());

    }
}
```
```yaml
resources:
  exampleRecorder:
    type: aws:cfg:Recorder
  exampleFunction:
    type: aws:lambda:Function
  examplePermission:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunction
      function: ${exampleFunction.arn}
      principal: config.amazonaws.com
  exampleRule:
    type: aws:cfg:Rule
    properties:
      source:
        owner: CUSTOM_LAMBDA
        sourceIdentifier: ${exampleFunction.arn}
    options:
      dependson:
        - ${exampleRecorder}
        - ${examplePermission}
```
{{% /example %}}
{{% example %}}
### Custom Policies

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cfg.Rule("example", {source: {
    owner: "CUSTOM_POLICY",
    sourceDetails: [{
        messageType: "ConfigurationItemChangeNotification",
    }],
    customPolicyDetails: {
        policyRuntime: "guard-2.x.x",
        policyText: `	  rule tableisactive when
		  resourceType == "AWS::DynamoDB::Table" {
		  configuration.tableStatus == ['ACTIVE']
	  }
	  
	  rule checkcompliance when
		  resourceType == "AWS::DynamoDB::Table"
		  tableisactive {
			  supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus == "ENABLED"
	  }
`,
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cfg.Rule("example", source=aws.cfg.RuleSourceArgs(
    owner="CUSTOM_POLICY",
    source_details=[aws.cfg.RuleSourceSourceDetailArgs(
        message_type="ConfigurationItemChangeNotification",
    )],
    custom_policy_details=aws.cfg.RuleSourceCustomPolicyDetailsArgs(
        policy_runtime="guard-2.x.x",
        policy_text="""	  rule tableisactive when
		  resourceType == "AWS::DynamoDB::Table" {
		  configuration.tableStatus == ['ACTIVE']
	  }
	  
	  rule checkcompliance when
		  resourceType == "AWS::DynamoDB::Table"
		  tableisactive {
			  supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus == "ENABLED"
	  }
""",
    ),
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Cfg.Rule("example", new()
    {
        Source = new Aws.Cfg.Inputs.RuleSourceArgs
        {
            Owner = "CUSTOM_POLICY",
            SourceDetails = new[]
            {
                new Aws.Cfg.Inputs.RuleSourceSourceDetailArgs
                {
                    MessageType = "ConfigurationItemChangeNotification",
                },
            },
            CustomPolicyDetails = new Aws.Cfg.Inputs.RuleSourceCustomPolicyDetailsArgs
            {
                PolicyRuntime = "guard-2.x.x",
                PolicyText = @"	  rule tableisactive when
		  resourceType == ""AWS::DynamoDB::Table"" {
		  configuration.tableStatus == ['ACTIVE']
	  }
	  
	  rule checkcompliance when
		  resourceType == ""AWS::DynamoDB::Table""
		  tableisactive {
			  supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus == ""ENABLED""
	  }
",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cfg"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cfg.NewRule(ctx, "example", &cfg.RuleArgs{
			Source: &cfg.RuleSourceArgs{
				Owner: pulumi.String("CUSTOM_POLICY"),
				SourceDetails: cfg.RuleSourceSourceDetailArray{
					&cfg.RuleSourceSourceDetailArgs{
						MessageType: pulumi.String("ConfigurationItemChangeNotification"),
					},
				},
				CustomPolicyDetails: &cfg.RuleSourceCustomPolicyDetailsArgs{
					PolicyRuntime: pulumi.String("guard-2.x.x"),
					PolicyText: pulumi.String(`	  rule tableisactive when
		  resourceType == "AWS::DynamoDB::Table" {
		  configuration.tableStatus == ['ACTIVE']
	  }
	  
	  rule checkcompliance when
		  resourceType == "AWS::DynamoDB::Table"
		  tableisactive {
			  supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus == "ENABLED"
	  }
`),
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
import com.pulumi.aws.cfg.Rule;
import com.pulumi.aws.cfg.RuleArgs;
import com.pulumi.aws.cfg.inputs.RuleSourceArgs;
import com.pulumi.aws.cfg.inputs.RuleSourceCustomPolicyDetailsArgs;
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
        var example = new Rule("example", RuleArgs.builder()        
            .source(RuleSourceArgs.builder()
                .owner("CUSTOM_POLICY")
                .sourceDetails(RuleSourceSourceDetailArgs.builder()
                    .messageType("ConfigurationItemChangeNotification")
                    .build())
                .customPolicyDetails(RuleSourceCustomPolicyDetailsArgs.builder()
                    .policyRuntime("guard-2.x.x")
                    .policyText("""
	  rule tableisactive when
		  resourceType == "AWS::DynamoDB::Table" {
		  configuration.tableStatus == ['ACTIVE']
	  }
	  
	  rule checkcompliance when
		  resourceType == "AWS::DynamoDB::Table"
		  tableisactive {
			  supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus == "ENABLED"
	  }
                    """)
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cfg:Rule
    properties:
      source:
        owner: CUSTOM_POLICY
        sourceDetails:
          - messageType: ConfigurationItemChangeNotification
        customPolicyDetails:
          policyRuntime: guard-2.x.x
          policyText: "\t  rule tableisactive when\n\t\t  resourceType == \"AWS::DynamoDB::Table\" {\n\t\t  configuration.tableStatus == ['ACTIVE']\n\t  }\n\t  \n\t  rule checkcompliance when\n\t\t  resourceType == \"AWS::DynamoDB::Table\"\n\t\t  tableisactive {\n\t\t\t  supplementaryConfiguration.ContinuousBackupsDescription.pointInTimeRecoveryDescription.pointInTimeRecoveryStatus == \"ENABLED\"\n\t  }\n"
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Config Rule using the name. For example:

```sh
 $ pulumi import aws:cfg/rule:Rule foo example
```
 