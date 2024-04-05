Manages a CloudFormation StackSet. StackSets allow CloudFormation templates to be easily deployed across multiple accounts and regions via StackSet Instances (`aws.cloudformation.StackSetInstance` resource). Additional information about StackSets can be found in the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/what-is-cfnstacksets.html).

> **NOTE:** All template parameters, including those with a `Default`, must be configured or ignored with the `lifecycle` configuration block `ignore_changes` argument.

> **NOTE:** All `NoEcho` template parameters must be ignored with the `lifecycle` configuration block `ignore_changes` argument.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        effect: "Allow",
        principals: [{
            identifiers: ["cloudformation.amazonaws.com"],
            type: "Service",
        }],
    }],
});
const aWSCloudFormationStackSetAdministrationRole = new aws.iam.Role("aWSCloudFormationStackSetAdministrationRole", {assumeRolePolicy: aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.then(aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy => aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.json)});
const example = new aws.cloudformation.StackSet("example", {
    administrationRoleArn: aWSCloudFormationStackSetAdministrationRole.arn,
    parameters: {
        VPCCidr: "10.0.0.0/16",
    },
    templateBody: JSON.stringify({
        Parameters: {
            VPCCidr: {
                Type: "String",
                Default: "10.0.0.0/16",
                Description: "Enter the CIDR block for the VPC. Default is 10.0.0.0/16.",
            },
        },
        Resources: {
            myVpc: {
                Type: "AWS::EC2::VPC",
                Properties: {
                    CidrBlock: {
                        Ref: "VPCCidr",
                    },
                    Tags: [{
                        Key: "Name",
                        Value: "Primary_CF_VPC",
                    }],
                },
            },
        },
    }),
});
const aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        actions: ["sts:AssumeRole"],
        effect: "Allow",
        resources: [pulumi.interpolate`arn:aws:iam::*:role/${example.executionRoleName}`],
    }],
});
const aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy = new aws.iam.RolePolicy("aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy", {
    policy: aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.apply(aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument => aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.json),
    role: aWSCloudFormationStackSetAdministrationRole.name,
});
```
```python
import pulumi
import json
import pulumi_aws as aws

a_ws_cloud_formation_stack_set_administration_role_assume_role_policy = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        identifiers=["cloudformation.amazonaws.com"],
        type="Service",
    )],
)])
a_ws_cloud_formation_stack_set_administration_role = aws.iam.Role("aWSCloudFormationStackSetAdministrationRole", assume_role_policy=a_ws_cloud_formation_stack_set_administration_role_assume_role_policy.json)
example = aws.cloudformation.StackSet("example",
    administration_role_arn=a_ws_cloud_formation_stack_set_administration_role.arn,
    parameters={
        "VPCCidr": "10.0.0.0/16",
    },
    template_body=json.dumps({
        "Parameters": {
            "VPCCidr": {
                "Type": "String",
                "Default": "10.0.0.0/16",
                "Description": "Enter the CIDR block for the VPC. Default is 10.0.0.0/16.",
            },
        },
        "Resources": {
            "myVpc": {
                "Type": "AWS::EC2::VPC",
                "Properties": {
                    "CidrBlock": {
                        "Ref": "VPCCidr",
                    },
                    "Tags": [{
                        "Key": "Name",
                        "Value": "Primary_CF_VPC",
                    }],
                },
            },
        },
    }))
a_ws_cloud_formation_stack_set_administration_role_execution_policy_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    effect="Allow",
    resources=[example.execution_role_name.apply(lambda execution_role_name: f"arn:aws:iam::*:role/{execution_role_name}")],
)])
a_ws_cloud_formation_stack_set_administration_role_execution_policy_role_policy = aws.iam.RolePolicy("aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy",
    policy=a_ws_cloud_formation_stack_set_administration_role_execution_policy_policy_document.json,
    role=a_ws_cloud_formation_stack_set_administration_role.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                            "cloudformation.amazonaws.com",
                        },
                        Type = "Service",
                    },
                },
            },
        },
    });

    var aWSCloudFormationStackSetAdministrationRole = new Aws.Iam.Role("aWSCloudFormationStackSetAdministrationRole", new()
    {
        AssumeRolePolicy = aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var example = new Aws.CloudFormation.StackSet("example", new()
    {
        AdministrationRoleArn = aWSCloudFormationStackSetAdministrationRole.Arn,
        Parameters = 
        {
            { "VPCCidr", "10.0.0.0/16" },
        },
        TemplateBody = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Parameters"] = new Dictionary<string, object?>
            {
                ["VPCCidr"] = new Dictionary<string, object?>
                {
                    ["Type"] = "String",
                    ["Default"] = "10.0.0.0/16",
                    ["Description"] = "Enter the CIDR block for the VPC. Default is 10.0.0.0/16.",
                },
            },
            ["Resources"] = new Dictionary<string, object?>
            {
                ["myVpc"] = new Dictionary<string, object?>
                {
                    ["Type"] = "AWS::EC2::VPC",
                    ["Properties"] = new Dictionary<string, object?>
                    {
                        ["CidrBlock"] = new Dictionary<string, object?>
                        {
                            ["Ref"] = "VPCCidr",
                        },
                        ["Tags"] = new[]
                        {
                            new Dictionary<string, object?>
                            {
                                ["Key"] = "Name",
                                ["Value"] = "Primary_CF_VPC",
                            },
                        },
                    },
                },
            },
        }),
    });

    var aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                Resources = new[]
                {
                    $"arn:aws:iam::*:role/{example.ExecutionRoleName}",
                },
            },
        },
    });

    var aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy = new Aws.Iam.RolePolicy("aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy", new()
    {
        Policy = aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
        Role = aWSCloudFormationStackSetAdministrationRole.Name,
    });

});
```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudformation"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Identifiers: []string{
								"cloudformation.amazonaws.com",
							},
							Type: "Service",
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		aWSCloudFormationStackSetAdministrationRole, err := iam.NewRole(ctx, "aWSCloudFormationStackSetAdministrationRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.Json),
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Parameters": map[string]interface{}{
				"VPCCidr": map[string]interface{}{
					"Type":        "String",
					"Default":     "10.0.0.0/16",
					"Description": "Enter the CIDR block for the VPC. Default is 10.0.0.0/16.",
				},
			},
			"Resources": map[string]interface{}{
				"myVpc": map[string]interface{}{
					"Type": "AWS::EC2::VPC",
					"Properties": map[string]interface{}{
						"CidrBlock": map[string]interface{}{
							"Ref": "VPCCidr",
						},
						"Tags": []map[string]interface{}{
							map[string]interface{}{
								"Key":   "Name",
								"Value": "Primary_CF_VPC",
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		example, err := cloudformation.NewStackSet(ctx, "example", &cloudformation.StackSetArgs{
			AdministrationRoleArn: aWSCloudFormationStackSetAdministrationRole.Arn,
			Parameters: pulumi.StringMap{
				"VPCCidr": pulumi.String("10.0.0.0/16"),
			},
			TemplateBody: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Actions: pulumi.StringArray{
						pulumi.String("sts:AssumeRole"),
					},
					Effect: pulumi.String("Allow"),
					Resources: pulumi.StringArray{
						example.ExecutionRoleName.ApplyT(func(executionRoleName string) (string, error) {
							return fmt.Sprintf("arn:aws:iam::*:role/%v", executionRoleName), nil
						}).(pulumi.StringOutput),
					},
				},
			},
		}, nil)
		_, err = iam.NewRolePolicy(ctx, "aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy", &iam.RolePolicyArgs{
			Policy: aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.ApplyT(func(aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
			Role: aWSCloudFormationStackSetAdministrationRole.Name,
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
import com.pulumi.aws.cloudformation.StackSet;
import com.pulumi.aws.cloudformation.StackSetArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        final var aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .identifiers("cloudformation.amazonaws.com")
                    .type("Service")
                    .build())
                .build())
            .build());

        var aWSCloudFormationStackSetAdministrationRole = new Role("aWSCloudFormationStackSetAdministrationRole", RoleArgs.builder()        
            .assumeRolePolicy(aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var example = new StackSet("example", StackSetArgs.builder()        
            .administrationRoleArn(aWSCloudFormationStackSetAdministrationRole.arn())
            .parameters(Map.of("VPCCidr", "10.0.0.0/16"))
            .templateBody(serializeJson(
                jsonObject(
                    jsonProperty("Parameters", jsonObject(
                        jsonProperty("VPCCidr", jsonObject(
                            jsonProperty("Type", "String"),
                            jsonProperty("Default", "10.0.0.0/16"),
                            jsonProperty("Description", "Enter the CIDR block for the VPC. Default is 10.0.0.0/16.")
                        ))
                    )),
                    jsonProperty("Resources", jsonObject(
                        jsonProperty("myVpc", jsonObject(
                            jsonProperty("Type", "AWS::EC2::VPC"),
                            jsonProperty("Properties", jsonObject(
                                jsonProperty("CidrBlock", jsonObject(
                                    jsonProperty("Ref", "VPCCidr")
                                )),
                                jsonProperty("Tags", jsonArray(jsonObject(
                                    jsonProperty("Key", "Name"),
                                    jsonProperty("Value", "Primary_CF_VPC")
                                )))
                            ))
                        ))
                    ))
                )))
            .build());

        final var aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .effect("Allow")
                .resources(example.executionRoleName().applyValue(executionRoleName -> String.format("arn:aws:iam::*:role/%s", executionRoleName)))
                .build())
            .build());

        var aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy = new RolePolicy("aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy", RolePolicyArgs.builder()        
            .policy(aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument -> aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .role(aWSCloudFormationStackSetAdministrationRole.name())
            .build());

    }
}
```
```yaml
resources:
  aWSCloudFormationStackSetAdministrationRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy.json}
  example:
    type: aws:cloudformation:StackSet
    properties:
      administrationRoleArn: ${aWSCloudFormationStackSetAdministrationRole.arn}
      parameters:
        VPCCidr: 10.0.0.0/16
      templateBody:
        fn::toJSON:
          Parameters:
            VPCCidr:
              Type: String
              Default: 10.0.0.0/16
              Description: Enter the CIDR block for the VPC. Default is 10.0.0.0/16.
          Resources:
            myVpc:
              Type: AWS::EC2::VPC
              Properties:
                CidrBlock:
                  Ref: VPCCidr
                Tags:
                  - Key: Name
                    Value: Primary_CF_VPC
  aWSCloudFormationStackSetAdministrationRoleExecutionPolicyRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      policy: ${aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument.json}
      role: ${aWSCloudFormationStackSetAdministrationRole.name}
variables:
  aWSCloudFormationStackSetAdministrationRoleAssumeRolePolicy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            effect: Allow
            principals:
              - identifiers:
                  - cloudformation.amazonaws.com
                type: Service
  aWSCloudFormationStackSetAdministrationRoleExecutionPolicyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            effect: Allow
            resources:
              - arn:aws:iam::*:role/${example.executionRoleName}
```
{{% /example %}}
{{% /examples %}}

## Import

Import CloudFormation StackSets when acting a delegated administrator in a member account using the `name` and `call_as` values separated by a comma (`,`). For example:

Using `pulumi import`, import CloudFormation StackSets using the `name`. For example:

```sh
 $ pulumi import aws:cloudformation/stackSet:StackSet example example
```
 Using `pulumi import`, import CloudFormation StackSets when acting a delegated administrator in a member account using the `name` and `call_as` values separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:cloudformation/stackSet:StackSet example example,DELEGATED_ADMIN
```
 