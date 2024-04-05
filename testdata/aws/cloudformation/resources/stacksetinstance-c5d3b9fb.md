Manages a CloudFormation StackSet Instance. Instances are managed in the account and region of the StackSet after the target account permissions have been configured. Additional information about StackSets can be found in the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/what-is-cfnstacksets.html).

> **NOTE:** All target accounts must have an IAM Role created that matches the name of the execution role configured in the StackSet (the `execution_role_name` argument in the `aws.cloudformation.StackSet` resource) in a trust relationship with the administrative account or administration IAM Role. The execution role must have appropriate permissions to manage resources defined in the template along with those required for StackSets to operate. See the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/stacksets-prereqs.html) for more details.

> **NOTE:** To retain the Stack during resource destroy, ensure `retain_stack` has been set to `true` in the state first. This must be completed _before_ a deployment that would destroy the resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudformation.StackSetInstance("example", {
    accountId: "123456789012",
    region: "us-east-1",
    stackSetName: aws_cloudformation_stack_set.example.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudformation.StackSetInstance("example",
    account_id="123456789012",
    region="us-east-1",
    stack_set_name=aws_cloudformation_stack_set["example"]["name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudFormation.StackSetInstance("example", new()
    {
        AccountId = "123456789012",
        Region = "us-east-1",
        StackSetName = aws_cloudformation_stack_set.Example.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudformation.NewStackSetInstance(ctx, "example", &cloudformation.StackSetInstanceArgs{
			AccountId:    pulumi.String("123456789012"),
			Region:       pulumi.String("us-east-1"),
			StackSetName: pulumi.Any(aws_cloudformation_stack_set.Example.Name),
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
import com.pulumi.aws.cloudformation.StackSetInstance;
import com.pulumi.aws.cloudformation.StackSetInstanceArgs;
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
        var example = new StackSetInstance("example", StackSetInstanceArgs.builder()        
            .accountId("123456789012")
            .region("us-east-1")
            .stackSetName(aws_cloudformation_stack_set.example().name())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudformation:StackSetInstance
    properties:
      accountId: '123456789012'
      region: us-east-1
      stackSetName: ${aws_cloudformation_stack_set.example.name}
```
{{% /example %}}
{{% example %}}
### Example IAM Setup in Target Account

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        effect: "Allow",
        principals: [{
            identifiers: [aws_iam_role.AWSCloudFormationStackSetAdministrationRole.arn],
            type: "AWS",
        }],
    }],
});
const aWSCloudFormationStackSetExecutionRole = new aws.iam.Role("aWSCloudFormationStackSetExecutionRole", {assumeRolePolicy: aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy.then(aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy => aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy.json)});
const aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        actions: [
            "cloudformation:*",
            "s3:*",
            "sns:*",
        ],
        effect: "Allow",
        resources: ["*"],
    }],
});
const aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy = new aws.iam.RolePolicy("aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy", {
    policy: aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument.then(aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument => aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument.json),
    role: aWSCloudFormationStackSetExecutionRole.name,
});
```
```python
import pulumi
import pulumi_aws as aws

a_ws_cloud_formation_stack_set_execution_role_assume_role_policy = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        identifiers=[aws_iam_role["AWSCloudFormationStackSetAdministrationRole"]["arn"]],
        type="AWS",
    )],
)])
a_ws_cloud_formation_stack_set_execution_role = aws.iam.Role("aWSCloudFormationStackSetExecutionRole", assume_role_policy=a_ws_cloud_formation_stack_set_execution_role_assume_role_policy.json)
a_ws_cloud_formation_stack_set_execution_role_minimum_execution_policy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=[
        "cloudformation:*",
        "s3:*",
        "sns:*",
    ],
    effect="Allow",
    resources=["*"],
)])
a_ws_cloud_formation_stack_set_execution_role_minimum_execution_policy_role_policy = aws.iam.RolePolicy("aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy",
    policy=a_ws_cloud_formation_stack_set_execution_role_minimum_execution_policy_policy_document.json,
    role=a_ws_cloud_formation_stack_set_execution_role.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Identifiers = new[]
                        {
                            aws_iam_role.AWSCloudFormationStackSetAdministrationRole.Arn,
                        },
                        Type = "AWS",
                    },
                },
            },
        },
    });

    var aWSCloudFormationStackSetExecutionRole = new Aws.Iam.Role("aWSCloudFormationStackSetExecutionRole", new()
    {
        AssumeRolePolicy = aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "cloudformation:*",
                    "s3:*",
                    "sns:*",
                },
                Effect = "Allow",
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy = new Aws.Iam.RolePolicy("aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy", new()
    {
        Policy = aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
        Role = aWSCloudFormationStackSetExecutionRole.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Actions: []string{
"sts:AssumeRole",
},
Effect: pulumi.StringRef("Allow"),
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Identifiers: interface{}{
aws_iam_role.AWSCloudFormationStackSetAdministrationRole.Arn,
},
Type: "AWS",
},
},
},
},
}, nil);
if err != nil {
return err
}
aWSCloudFormationStackSetExecutionRole, err := iam.NewRole(ctx, "aWSCloudFormationStackSetExecutionRole", &iam.RoleArgs{
AssumeRolePolicy: *pulumi.String(aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy.Json),
})
if err != nil {
return err
}
aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Actions: []string{
"cloudformation:*",
"s3:*",
"sns:*",
},
Effect: pulumi.StringRef("Allow"),
Resources: []string{
"*",
},
},
},
}, nil);
if err != nil {
return err
}
_, err = iam.NewRolePolicy(ctx, "aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy", &iam.RolePolicyArgs{
Policy: *pulumi.String(aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument.Json),
Role: aWSCloudFormationStackSetExecutionRole.Name,
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
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
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
        final var aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .identifiers(aws_iam_role.AWSCloudFormationStackSetAdministrationRole().arn())
                    .type("AWS")
                    .build())
                .build())
            .build());

        var aWSCloudFormationStackSetExecutionRole = new Role("aWSCloudFormationStackSetExecutionRole", RoleArgs.builder()        
            .assumeRolePolicy(aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        final var aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions(                
                    "cloudformation:*",
                    "s3:*",
                    "sns:*")
                .effect("Allow")
                .resources("*")
                .build())
            .build());

        var aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy = new RolePolicy("aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy", RolePolicyArgs.builder()        
            .policy(aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .role(aWSCloudFormationStackSetExecutionRole.name())
            .build());

    }
}
```
```yaml
resources:
  aWSCloudFormationStackSetExecutionRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy.json}
  aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      policy: ${aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument.json}
      role: ${aWSCloudFormationStackSetExecutionRole.name}
variables:
  aWSCloudFormationStackSetExecutionRoleAssumeRolePolicy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            effect: Allow
            principals:
              - identifiers:
                  - ${aws_iam_role.AWSCloudFormationStackSetAdministrationRole.arn}
                type: AWS
  aWSCloudFormationStackSetExecutionRoleMinimumExecutionPolicyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - cloudformation:*
              - s3:*
              - sns:*
            effect: Allow
            resources:
              - '*'
```
{{% /example %}}
{{% example %}}
### Example Deployment across Organizations account

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudformation.StackSetInstance("example", {
    deploymentTargets: {
        organizationalUnitIds: [aws_organizations_organization.example.roots[0].id],
    },
    region: "us-east-1",
    stackSetName: aws_cloudformation_stack_set.example.name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudformation.StackSetInstance("example",
    deployment_targets=aws.cloudformation.StackSetInstanceDeploymentTargetsArgs(
        organizational_unit_ids=[aws_organizations_organization["example"]["roots"][0]["id"]],
    ),
    region="us-east-1",
    stack_set_name=aws_cloudformation_stack_set["example"]["name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudFormation.StackSetInstance("example", new()
    {
        DeploymentTargets = new Aws.CloudFormation.Inputs.StackSetInstanceDeploymentTargetsArgs
        {
            OrganizationalUnitIds = new[]
            {
                aws_organizations_organization.Example.Roots[0].Id,
            },
        },
        Region = "us-east-1",
        StackSetName = aws_cloudformation_stack_set.Example.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudformation.NewStackSetInstance(ctx, "example", &cloudformation.StackSetInstanceArgs{
			DeploymentTargets: &cloudformation.StackSetInstanceDeploymentTargetsArgs{
				OrganizationalUnitIds: pulumi.StringArray{
					aws_organizations_organization.Example.Roots[0].Id,
				},
			},
			Region:       pulumi.String("us-east-1"),
			StackSetName: pulumi.Any(aws_cloudformation_stack_set.Example.Name),
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
import com.pulumi.aws.cloudformation.StackSetInstance;
import com.pulumi.aws.cloudformation.StackSetInstanceArgs;
import com.pulumi.aws.cloudformation.inputs.StackSetInstanceDeploymentTargetsArgs;
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
        var example = new StackSetInstance("example", StackSetInstanceArgs.builder()        
            .deploymentTargets(StackSetInstanceDeploymentTargetsArgs.builder()
                .organizationalUnitIds(aws_organizations_organization.example().roots()[0].id())
                .build())
            .region("us-east-1")
            .stackSetName(aws_cloudformation_stack_set.example().name())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudformation:StackSetInstance
    properties:
      deploymentTargets:
        organizationalUnitIds:
          - ${aws_organizations_organization.example.roots[0].id}
      region: us-east-1
      stackSetName: ${aws_cloudformation_stack_set.example.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Import CloudFormation StackSet Instances that target AWS Organizational Units using the StackSet name, a slash (`/`) separated list of organizational unit IDs, and target AWS Region separated by commas (`,`). For example:

Import CloudFormation StackSet Instances when acting a delegated administrator in a member account using the StackSet name, target AWS account ID or slash (`/`) separated list of organizational unit IDs, target AWS Region and `call_as` value separated by commas (`,`). For example:

Using `pulumi import`, import CloudFormation StackSet Instances that target an AWS Account ID using the StackSet name, target AWS account ID, and target AWS Region separated by commas (`,`). For example:

```sh
 $ pulumi import aws:cloudformation/stackSetInstance:StackSetInstance example example,123456789012,us-east-1
```
 Using `pulumi import`, import CloudFormation StackSet Instances that target AWS Organizational Units using the StackSet name, a slash (`/`) separated list of organizational unit IDs, and target AWS Region separated by commas (`,`). For example:

```sh
 $ pulumi import aws:cloudformation/stackSetInstance:StackSetInstance example example,ou-sdas-123123123/ou-sdas-789789789,us-east-1
```
 Using `pulumi import`, import CloudFormation StackSet Instances when acting a delegated administrator in a member account using the StackSet name, target AWS account ID or slash (`/`) separated list of organizational unit IDs, target AWS Region and `call_as` value separated by commas (`,`). For example:

```sh
 $ pulumi import aws:cloudformation/stackSetInstance:StackSetInstance example example,ou-sdas-123123123/ou-sdas-789789789,us-east-1,DELEGATED_ADMIN
```
 