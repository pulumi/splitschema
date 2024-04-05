Manages an IoT fleet provisioning template. For more info, see the AWS documentation on [fleet provisioning](https://docs.aws.amazon.com/iot/latest/developerguide/provision-wo-cert.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const iotAssumeRolePolicy = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["iot.amazonaws.com"],
        }],
    }],
});
const iotFleetProvisioning = new aws.iam.Role("iotFleetProvisioning", {
    path: "/service-role/",
    assumeRolePolicy: iotAssumeRolePolicy.then(iotAssumeRolePolicy => iotAssumeRolePolicy.json),
});
const iotFleetProvisioningRegistration = new aws.iam.RolePolicyAttachment("iotFleetProvisioningRegistration", {
    role: iotFleetProvisioning.name,
    policyArn: "arn:aws:iam::aws:policy/service-role/AWSIoTThingsRegistration",
});
const devicePolicyPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["iot:Subscribe"],
        resources: ["*"],
    }],
});
const devicePolicyPolicy = new aws.iot.Policy("devicePolicyPolicy", {policy: devicePolicyPolicyDocument.then(devicePolicyPolicyDocument => devicePolicyPolicyDocument.json)});
const fleet = new aws.iot.ProvisioningTemplate("fleet", {
    description: "My provisioning template",
    provisioningRoleArn: iotFleetProvisioning.arn,
    enabled: true,
    templateBody: devicePolicyPolicy.name.apply(name => JSON.stringify({
        Parameters: {
            SerialNumber: {
                Type: "String",
            },
        },
        Resources: {
            certificate: {
                Properties: {
                    CertificateId: {
                        Ref: "AWS::IoT::Certificate::Id",
                    },
                    Status: "Active",
                },
                Type: "AWS::IoT::Certificate",
            },
            policy: {
                Properties: {
                    PolicyName: name,
                },
                Type: "AWS::IoT::Policy",
            },
        },
    })),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

iot_assume_role_policy = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["iot.amazonaws.com"],
    )],
)])
iot_fleet_provisioning = aws.iam.Role("iotFleetProvisioning",
    path="/service-role/",
    assume_role_policy=iot_assume_role_policy.json)
iot_fleet_provisioning_registration = aws.iam.RolePolicyAttachment("iotFleetProvisioningRegistration",
    role=iot_fleet_provisioning.name,
    policy_arn="arn:aws:iam::aws:policy/service-role/AWSIoTThingsRegistration")
device_policy_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["iot:Subscribe"],
    resources=["*"],
)])
device_policy_policy = aws.iot.Policy("devicePolicyPolicy", policy=device_policy_policy_document.json)
fleet = aws.iot.ProvisioningTemplate("fleet",
    description="My provisioning template",
    provisioning_role_arn=iot_fleet_provisioning.arn,
    enabled=True,
    template_body=device_policy_policy.name.apply(lambda name: json.dumps({
        "Parameters": {
            "SerialNumber": {
                "Type": "String",
            },
        },
        "Resources": {
            "certificate": {
                "Properties": {
                    "CertificateId": {
                        "Ref": "AWS::IoT::Certificate::Id",
                    },
                    "Status": "Active",
                },
                "Type": "AWS::IoT::Certificate",
            },
            "policy": {
                "Properties": {
                    "PolicyName": name,
                },
                "Type": "AWS::IoT::Policy",
            },
        },
    })))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var iotAssumeRolePolicy = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "iot.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var iotFleetProvisioning = new Aws.Iam.Role("iotFleetProvisioning", new()
    {
        Path = "/service-role/",
        AssumeRolePolicy = iotAssumeRolePolicy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var iotFleetProvisioningRegistration = new Aws.Iam.RolePolicyAttachment("iotFleetProvisioningRegistration", new()
    {
        Role = iotFleetProvisioning.Name,
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSIoTThingsRegistration",
    });

    var devicePolicyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "iot:Subscribe",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var devicePolicyPolicy = new Aws.Iot.Policy("devicePolicyPolicy", new()
    {
        PolicyDocument = devicePolicyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var fleet = new Aws.Iot.ProvisioningTemplate("fleet", new()
    {
        Description = "My provisioning template",
        ProvisioningRoleArn = iotFleetProvisioning.Arn,
        Enabled = true,
        TemplateBody = devicePolicyPolicy.Name.Apply(name => JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Parameters"] = new Dictionary<string, object?>
            {
                ["SerialNumber"] = new Dictionary<string, object?>
                {
                    ["Type"] = "String",
                },
            },
            ["Resources"] = new Dictionary<string, object?>
            {
                ["certificate"] = new Dictionary<string, object?>
                {
                    ["Properties"] = new Dictionary<string, object?>
                    {
                        ["CertificateId"] = new Dictionary<string, object?>
                        {
                            ["Ref"] = "AWS::IoT::Certificate::Id",
                        },
                        ["Status"] = "Active",
                    },
                    ["Type"] = "AWS::IoT::Certificate",
                },
                ["policy"] = new Dictionary<string, object?>
                {
                    ["Properties"] = new Dictionary<string, object?>
                    {
                        ["PolicyName"] = name,
                    },
                    ["Type"] = "AWS::IoT::Policy",
                },
            },
        })),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		iotAssumeRolePolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"iot.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		iotFleetProvisioning, err := iam.NewRole(ctx, "iotFleetProvisioning", &iam.RoleArgs{
			Path:             pulumi.String("/service-role/"),
			AssumeRolePolicy: *pulumi.String(iotAssumeRolePolicy.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "iotFleetProvisioningRegistration", &iam.RolePolicyAttachmentArgs{
			Role:      iotFleetProvisioning.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AWSIoTThingsRegistration"),
		})
		if err != nil {
			return err
		}
		devicePolicyPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"iot:Subscribe",
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
		devicePolicyPolicy, err := iot.NewPolicy(ctx, "devicePolicyPolicy", &iot.PolicyArgs{
			Policy: *pulumi.String(devicePolicyPolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = iot.NewProvisioningTemplate(ctx, "fleet", &iot.ProvisioningTemplateArgs{
			Description:         pulumi.String("My provisioning template"),
			ProvisioningRoleArn: iotFleetProvisioning.Arn,
			Enabled:             pulumi.Bool(true),
			TemplateBody: devicePolicyPolicy.Name.ApplyT(func(name string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON0, err := json.Marshal(map[string]interface{}{
					"Parameters": map[string]interface{}{
						"SerialNumber": map[string]interface{}{
							"Type": "String",
						},
					},
					"Resources": map[string]interface{}{
						"certificate": map[string]interface{}{
							"Properties": map[string]interface{}{
								"CertificateId": map[string]interface{}{
									"Ref": "AWS::IoT::Certificate::Id",
								},
								"Status": "Active",
							},
							"Type": "AWS::IoT::Certificate",
						},
						"policy": map[string]interface{}{
							"Properties": map[string]interface{}{
								"PolicyName": name,
							},
							"Type": "AWS::IoT::Policy",
						},
					},
				})
				if err != nil {
					return _zero, err
				}
				json0 := string(tmpJSON0)
				return pulumi.String(json0), nil
			}).(pulumi.StringOutput),
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
import com.pulumi.aws.iot.Policy;
import com.pulumi.aws.iot.PolicyArgs;
import com.pulumi.aws.iot.ProvisioningTemplate;
import com.pulumi.aws.iot.ProvisioningTemplateArgs;
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
        final var iotAssumeRolePolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("iot.amazonaws.com")
                    .build())
                .build())
            .build());

        var iotFleetProvisioning = new Role("iotFleetProvisioning", RoleArgs.builder()        
            .path("/service-role/")
            .assumeRolePolicy(iotAssumeRolePolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var iotFleetProvisioningRegistration = new RolePolicyAttachment("iotFleetProvisioningRegistration", RolePolicyAttachmentArgs.builder()        
            .role(iotFleetProvisioning.name())
            .policyArn("arn:aws:iam::aws:policy/service-role/AWSIoTThingsRegistration")
            .build());

        final var devicePolicyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("iot:Subscribe")
                .resources("*")
                .build())
            .build());

        var devicePolicyPolicy = new Policy("devicePolicyPolicy", PolicyArgs.builder()        
            .policy(devicePolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var fleet = new ProvisioningTemplate("fleet", ProvisioningTemplateArgs.builder()        
            .description("My provisioning template")
            .provisioningRoleArn(iotFleetProvisioning.arn())
            .enabled(true)
            .templateBody(devicePolicyPolicy.name().applyValue(name -> serializeJson(
                jsonObject(
                    jsonProperty("Parameters", jsonObject(
                        jsonProperty("SerialNumber", jsonObject(
                            jsonProperty("Type", "String")
                        ))
                    )),
                    jsonProperty("Resources", jsonObject(
                        jsonProperty("certificate", jsonObject(
                            jsonProperty("Properties", jsonObject(
                                jsonProperty("CertificateId", jsonObject(
                                    jsonProperty("Ref", "AWS::IoT::Certificate::Id")
                                )),
                                jsonProperty("Status", "Active")
                            )),
                            jsonProperty("Type", "AWS::IoT::Certificate")
                        )),
                        jsonProperty("policy", jsonObject(
                            jsonProperty("Properties", jsonObject(
                                jsonProperty("PolicyName", name)
                            )),
                            jsonProperty("Type", "AWS::IoT::Policy")
                        ))
                    ))
                ))))
            .build());

    }
}
```
```yaml
resources:
  iotFleetProvisioning:
    type: aws:iam:Role
    properties:
      path: /service-role/
      assumeRolePolicy: ${iotAssumeRolePolicy.json}
  iotFleetProvisioningRegistration:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${iotFleetProvisioning.name}
      policyArn: arn:aws:iam::aws:policy/service-role/AWSIoTThingsRegistration
  devicePolicyPolicy:
    type: aws:iot:Policy
    properties:
      policy: ${devicePolicyPolicyDocument.json}
  fleet:
    type: aws:iot:ProvisioningTemplate
    properties:
      description: My provisioning template
      provisioningRoleArn: ${iotFleetProvisioning.arn}
      enabled: true
      templateBody:
        fn::toJSON:
          Parameters:
            SerialNumber:
              Type: String
          Resources:
            certificate:
              Properties:
                CertificateId:
                  Ref: AWS::IoT::Certificate::Id
                Status: Active
              Type: AWS::IoT::Certificate
            policy:
              Properties:
                PolicyName: ${devicePolicyPolicy.name}
              Type: AWS::IoT::Policy
variables:
  iotAssumeRolePolicy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            principals:
              - type: Service
                identifiers:
                  - iot.amazonaws.com
  devicePolicyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - iot:Subscribe
            resources:
              - '*'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IoT fleet provisioning templates using the `name`. For example:

```sh
 $ pulumi import aws:iot/provisioningTemplate:ProvisioningTemplate fleet FleetProvisioningTemplate
```
 