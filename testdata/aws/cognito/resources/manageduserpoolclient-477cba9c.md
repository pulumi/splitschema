Use the `aws.cognito.UserPoolClient` resource to manage a Cognito User Pool Client.

**This resource is advanced** and has special caveats to consider before use. Please read this document completely before using the resource.

Use the `aws.cognito.ManagedUserPoolClient` resource to manage a Cognito User Pool Client that is automatically created by an AWS service. For instance, when [configuring an OpenSearch Domain to use Cognito authentication](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html), the OpenSearch service creates the User Pool Client during setup and removes it when it is no longer required. As a result, the `aws.cognito.ManagedUserPoolClient` resource does not create or delete this resource, but instead assumes management of it.

Use the `aws.cognito.UserPoolClient` resource to manage Cognito User Pool Clients for normal use cases.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleUserPool = new aws.cognito.UserPool("exampleUserPool", {});
const exampleIdentityPool = new aws.cognito.IdentityPool("exampleIdentityPool", {identityPoolName: "example"});
const current = aws.getPartition({});
const examplePolicyDocument = current.then(current => aws.iam.getPolicyDocument({
    statements: [{
        sid: "",
        actions: ["sts:AssumeRole"],
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: [`es.${current.dnsSuffix}`],
        }],
    }],
}));
const exampleRole = new aws.iam.Role("exampleRole", {
    path: "/service-role/",
    assumeRolePolicy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
const exampleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment", {
    role: exampleRole.name,
    policyArn: current.then(current => `arn:${current.partition}:iam::aws:policy/AmazonESCognitoAccess`),
});
const exampleDomain = new aws.opensearch.Domain("exampleDomain", {
    cognitoOptions: {
        enabled: true,
        userPoolId: exampleUserPool.id,
        identityPoolId: exampleIdentityPool.id,
        roleArn: exampleRole.arn,
    },
    ebsOptions: {
        ebsEnabled: true,
        volumeSize: 10,
    },
}, {
    dependsOn: [
        aws_cognito_user_pool_domain.example,
        exampleRolePolicyAttachment,
    ],
});
const exampleManagedUserPoolClient = new aws.cognito.ManagedUserPoolClient("exampleManagedUserPoolClient", {
    namePrefix: "AmazonOpenSearchService-example",
    userPoolId: exampleUserPool.id,
}, {
    dependsOn: [exampleDomain],
});
```
```python
import pulumi
import pulumi_aws as aws

example_user_pool = aws.cognito.UserPool("exampleUserPool")
example_identity_pool = aws.cognito.IdentityPool("exampleIdentityPool", identity_pool_name="example")
current = aws.get_partition()
example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="",
    actions=["sts:AssumeRole"],
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=[f"es.{current.dns_suffix}"],
    )],
)])
example_role = aws.iam.Role("exampleRole",
    path="/service-role/",
    assume_role_policy=example_policy_document.json)
example_role_policy_attachment = aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment",
    role=example_role.name,
    policy_arn=f"arn:{current.partition}:iam::aws:policy/AmazonESCognitoAccess")
example_domain = aws.opensearch.Domain("exampleDomain",
    cognito_options=aws.opensearch.DomainCognitoOptionsArgs(
        enabled=True,
        user_pool_id=example_user_pool.id,
        identity_pool_id=example_identity_pool.id,
        role_arn=example_role.arn,
    ),
    ebs_options=aws.opensearch.DomainEbsOptionsArgs(
        ebs_enabled=True,
        volume_size=10,
    ),
    opts=pulumi.ResourceOptions(depends_on=[
            aws_cognito_user_pool_domain["example"],
            example_role_policy_attachment,
        ]))
example_managed_user_pool_client = aws.cognito.ManagedUserPoolClient("exampleManagedUserPoolClient",
    name_prefix="AmazonOpenSearchService-example",
    user_pool_id=example_user_pool.id,
    opts=pulumi.ResourceOptions(depends_on=[example_domain]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleUserPool = new Aws.Cognito.UserPool("exampleUserPool");

    var exampleIdentityPool = new Aws.Cognito.IdentityPool("exampleIdentityPool", new()
    {
        IdentityPoolName = "example",
    });

    var current = Aws.GetPartition.Invoke();

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "",
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            $"es.{current.Apply(getPartitionResult => getPartitionResult.DnsSuffix)}",
                        },
                    },
                },
            },
        },
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        Path = "/service-role/",
        AssumeRolePolicy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("exampleRolePolicyAttachment", new()
    {
        Role = exampleRole.Name,
        PolicyArn = $"arn:{current.Apply(getPartitionResult => getPartitionResult.Partition)}:iam::aws:policy/AmazonESCognitoAccess",
    });

    var exampleDomain = new Aws.OpenSearch.Domain("exampleDomain", new()
    {
        CognitoOptions = new Aws.OpenSearch.Inputs.DomainCognitoOptionsArgs
        {
            Enabled = true,
            UserPoolId = exampleUserPool.Id,
            IdentityPoolId = exampleIdentityPool.Id,
            RoleArn = exampleRole.Arn,
        },
        EbsOptions = new Aws.OpenSearch.Inputs.DomainEbsOptionsArgs
        {
            EbsEnabled = true,
            VolumeSize = 10,
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_cognito_user_pool_domain.Example,
            exampleRolePolicyAttachment,
        },
    });

    var exampleManagedUserPoolClient = new Aws.Cognito.ManagedUserPoolClient("exampleManagedUserPoolClient", new()
    {
        NamePrefix = "AmazonOpenSearchService-example",
        UserPoolId = exampleUserPool.Id,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleDomain,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleUserPool, err := cognito.NewUserPool(ctx, "exampleUserPool", nil)
		if err != nil {
			return err
		}
		exampleIdentityPool, err := cognito.NewIdentityPool(ctx, "exampleIdentityPool", &cognito.IdentityPoolArgs{
			IdentityPoolName: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		current, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid: pulumi.StringRef(""),
					Actions: []string{
						"sts:AssumeRole",
					},
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								fmt.Sprintf("es.%v", current.DnsSuffix),
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			Path:             pulumi.String("/service-role/"),
			AssumeRolePolicy: *pulumi.String(examplePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		exampleRolePolicyAttachment, err := iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      exampleRole.Name,
			PolicyArn: pulumi.String(fmt.Sprintf("arn:%v:iam::aws:policy/AmazonESCognitoAccess", current.Partition)),
		})
		if err != nil {
			return err
		}
		exampleDomain, err := opensearch.NewDomain(ctx, "exampleDomain", &opensearch.DomainArgs{
			CognitoOptions: &opensearch.DomainCognitoOptionsArgs{
				Enabled:        pulumi.Bool(true),
				UserPoolId:     exampleUserPool.ID(),
				IdentityPoolId: exampleIdentityPool.ID(),
				RoleArn:        exampleRole.Arn,
			},
			EbsOptions: &opensearch.DomainEbsOptionsArgs{
				EbsEnabled: pulumi.Bool(true),
				VolumeSize: pulumi.Int(10),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_cognito_user_pool_domain.Example,
			exampleRolePolicyAttachment,
		}))
		if err != nil {
			return err
		}
		_, err = cognito.NewManagedUserPoolClient(ctx, "exampleManagedUserPoolClient", &cognito.ManagedUserPoolClientArgs{
			NamePrefix: pulumi.String("AmazonOpenSearchService-example"),
			UserPoolId: exampleUserPool.ID(),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleDomain,
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.IdentityPool;
import com.pulumi.aws.cognito.IdentityPoolArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.opensearch.Domain;
import com.pulumi.aws.opensearch.DomainArgs;
import com.pulumi.aws.opensearch.inputs.DomainCognitoOptionsArgs;
import com.pulumi.aws.opensearch.inputs.DomainEbsOptionsArgs;
import com.pulumi.aws.cognito.ManagedUserPoolClient;
import com.pulumi.aws.cognito.ManagedUserPoolClientArgs;
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
        var exampleUserPool = new UserPool("exampleUserPool");

        var exampleIdentityPool = new IdentityPool("exampleIdentityPool", IdentityPoolArgs.builder()        
            .identityPoolName("example")
            .build());

        final var current = AwsFunctions.getPartition();

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("")
                .actions("sts:AssumeRole")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers(String.format("es.%s", current.applyValue(getPartitionResult -> getPartitionResult.dnsSuffix())))
                    .build())
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .path("/service-role/")
            .assumeRolePolicy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleRolePolicyAttachment = new RolePolicyAttachment("exampleRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .role(exampleRole.name())
            .policyArn(String.format("arn:%s:iam::aws:policy/AmazonESCognitoAccess", current.applyValue(getPartitionResult -> getPartitionResult.partition())))
            .build());

        var exampleDomain = new Domain("exampleDomain", DomainArgs.builder()        
            .cognitoOptions(DomainCognitoOptionsArgs.builder()
                .enabled(true)
                .userPoolId(exampleUserPool.id())
                .identityPoolId(exampleIdentityPool.id())
                .roleArn(exampleRole.arn())
                .build())
            .ebsOptions(DomainEbsOptionsArgs.builder()
                .ebsEnabled(true)
                .volumeSize(10)
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    aws_cognito_user_pool_domain.example(),
                    exampleRolePolicyAttachment)
                .build());

        var exampleManagedUserPoolClient = new ManagedUserPoolClient("exampleManagedUserPoolClient", ManagedUserPoolClientArgs.builder()        
            .namePrefix("AmazonOpenSearchService-example")
            .userPoolId(exampleUserPool.id())
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleDomain)
                .build());

    }
}
```
```yaml
resources:
  exampleManagedUserPoolClient:
    type: aws:cognito:ManagedUserPoolClient
    properties:
      namePrefix: AmazonOpenSearchService-example
      userPoolId: ${exampleUserPool.id}
    options:
      dependson:
        - ${exampleDomain}
  exampleUserPool:
    type: aws:cognito:UserPool
  exampleIdentityPool:
    type: aws:cognito:IdentityPool
    properties:
      identityPoolName: example
  exampleDomain:
    type: aws:opensearch:Domain
    properties:
      cognitoOptions:
        enabled: true
        userPoolId: ${exampleUserPool.id}
        identityPoolId: ${exampleIdentityPool.id}
        roleArn: ${exampleRole.arn}
      ebsOptions:
        ebsEnabled: true
        volumeSize: 10
    options:
      dependson:
        - ${aws_cognito_user_pool_domain.example}
        - ${exampleRolePolicyAttachment}
  exampleRole:
    type: aws:iam:Role
    properties:
      path: /service-role/
      assumeRolePolicy: ${examplePolicyDocument.json}
  exampleRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${exampleRole.name}
      policyArn: arn:${current.partition}:iam::aws:policy/AmazonESCognitoAccess
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid:
            actions:
              - sts:AssumeRole
            effect: Allow
            principals:
              - type: Service
                identifiers:
                  - es.${current.dnsSuffix}
  current:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cognito User Pool Clients using the `id` of the Cognito User Pool and the `id` of the Cognito User Pool Client. For example:

```sh
 $ pulumi import aws:cognito/managedUserPoolClient:ManagedUserPoolClient client us-west-2_abc123/3ho4ek12345678909nh3fmhpko
```
 