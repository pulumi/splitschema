Manages an AWS Config Configuration Aggregator

{{% examples %}}
## Example Usage
{{% example %}}
### Account Based Aggregation

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const account = new aws.cfg.ConfigurationAggregator("account", {accountAggregationSource: {
    accountIds: ["123456789012"],
    regions: ["us-west-2"],
}});
```
```python
import pulumi
import pulumi_aws as aws

account = aws.cfg.ConfigurationAggregator("account", account_aggregation_source=aws.cfg.ConfigurationAggregatorAccountAggregationSourceArgs(
    account_ids=["123456789012"],
    regions=["us-west-2"],
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var account = new Aws.Cfg.ConfigurationAggregator("account", new()
    {
        AccountAggregationSource = new Aws.Cfg.Inputs.ConfigurationAggregatorAccountAggregationSourceArgs
        {
            AccountIds = new[]
            {
                "123456789012",
            },
            Regions = new[]
            {
                "us-west-2",
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
		_, err := cfg.NewConfigurationAggregator(ctx, "account", &cfg.ConfigurationAggregatorArgs{
			AccountAggregationSource: &cfg.ConfigurationAggregatorAccountAggregationSourceArgs{
				AccountIds: pulumi.StringArray{
					pulumi.String("123456789012"),
				},
				Regions: pulumi.StringArray{
					pulumi.String("us-west-2"),
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
import com.pulumi.aws.cfg.ConfigurationAggregator;
import com.pulumi.aws.cfg.ConfigurationAggregatorArgs;
import com.pulumi.aws.cfg.inputs.ConfigurationAggregatorAccountAggregationSourceArgs;
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
        var account = new ConfigurationAggregator("account", ConfigurationAggregatorArgs.builder()        
            .accountAggregationSource(ConfigurationAggregatorAccountAggregationSourceArgs.builder()
                .accountIds("123456789012")
                .regions("us-west-2")
                .build())
            .build());

    }
}
```
```yaml
resources:
  account:
    type: aws:cfg:ConfigurationAggregator
    properties:
      accountAggregationSource:
        accountIds:
          - '123456789012'
        regions:
          - us-west-2
```
{{% /example %}}
{{% example %}}
### Organization Based Aggregation

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
const organizationRole = new aws.iam.Role("organizationRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const organizationRolePolicyAttachment = new aws.iam.RolePolicyAttachment("organizationRolePolicyAttachment", {
    role: organizationRole.name,
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations",
});
const organizationConfigurationAggregator = new aws.cfg.ConfigurationAggregator("organizationConfigurationAggregator", {organizationAggregationSource: {
    allRegions: true,
    roleArn: organizationRole.arn,
}}, {
    dependsOn: [organizationRolePolicyAttachment],
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
organization_role = aws.iam.Role("organizationRole", assume_role_policy=assume_role.json)
organization_role_policy_attachment = aws.iam.RolePolicyAttachment("organizationRolePolicyAttachment",
    role=organization_role.name,
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations")
organization_configuration_aggregator = aws.cfg.ConfigurationAggregator("organizationConfigurationAggregator", organization_aggregation_source=aws.cfg.ConfigurationAggregatorOrganizationAggregationSourceArgs(
    all_regions=True,
    role_arn=organization_role.arn,
),
opts=pulumi.ResourceOptions(depends_on=[organization_role_policy_attachment]))
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

    var organizationRole = new Aws.Iam.Role("organizationRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var organizationRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("organizationRolePolicyAttachment", new()
    {
        Role = organizationRole.Name,
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations",
    });

    var organizationConfigurationAggregator = new Aws.Cfg.ConfigurationAggregator("organizationConfigurationAggregator", new()
    {
        OrganizationAggregationSource = new Aws.Cfg.Inputs.ConfigurationAggregatorOrganizationAggregationSourceArgs
        {
            AllRegions = true,
            RoleArn = organizationRole.Arn,
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            organizationRolePolicyAttachment,
        },
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
		organizationRole, err := iam.NewRole(ctx, "organizationRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		organizationRolePolicyAttachment, err := iam.NewRolePolicyAttachment(ctx, "organizationRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      organizationRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations"),
		})
		if err != nil {
			return err
		}
		_, err = cfg.NewConfigurationAggregator(ctx, "organizationConfigurationAggregator", &cfg.ConfigurationAggregatorArgs{
			OrganizationAggregationSource: &cfg.ConfigurationAggregatorOrganizationAggregationSourceArgs{
				AllRegions: pulumi.Bool(true),
				RoleArn:    organizationRole.Arn,
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			organizationRolePolicyAttachment,
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.cfg.ConfigurationAggregator;
import com.pulumi.aws.cfg.ConfigurationAggregatorArgs;
import com.pulumi.aws.cfg.inputs.ConfigurationAggregatorOrganizationAggregationSourceArgs;
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

        var organizationRole = new Role("organizationRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var organizationRolePolicyAttachment = new RolePolicyAttachment("organizationRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .role(organizationRole.name())
            .policyArn("arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations")
            .build());

        var organizationConfigurationAggregator = new ConfigurationAggregator("organizationConfigurationAggregator", ConfigurationAggregatorArgs.builder()        
            .organizationAggregationSource(ConfigurationAggregatorOrganizationAggregationSourceArgs.builder()
                .allRegions(true)
                .roleArn(organizationRole.arn())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(organizationRolePolicyAttachment)
                .build());

    }
}
```
```yaml
resources:
  organizationConfigurationAggregator:
    type: aws:cfg:ConfigurationAggregator
    properties:
      organizationAggregationSource:
        allRegions: true
        roleArn: ${organizationRole.arn}
    options:
      dependson:
        - ${organizationRolePolicyAttachment}
  organizationRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  organizationRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${organizationRole.name}
      policyArn: arn:aws:iam::aws:policy/service-role/AWSConfigRoleForOrganizations
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
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Configuration Aggregators using the name. For example:

```sh
 $ pulumi import aws:cfg/configurationAggregator:ConfigurationAggregator example foo
```
 