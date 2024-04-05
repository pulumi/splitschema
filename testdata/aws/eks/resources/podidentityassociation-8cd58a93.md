Resource for managing an AWS EKS (Elastic Kubernetes) Pod Identity Association.

Creates an EKS Pod Identity association between a service account in an Amazon EKS cluster and an IAM role with EKS Pod Identity. Use EKS Pod Identity to give temporary IAM credentials to pods and the credentials are rotated automatically.

Amazon EKS Pod Identity associations provide the ability to manage credentials for your applications, similar to the way that EC2 instance profiles provide credentials to Amazon EC2 instances.

If a pod uses a service account that has an association, Amazon EKS sets environment variables in the containers of the pod. The environment variables configure the Amazon Web Services SDKs, including the Command Line Interface, to use the EKS Pod Identity credentials.

Pod Identity is a simpler method than IAM roles for service accounts, as this method doesn’t use OIDC identity providers. Additionally, you can configure a role for Pod Identity once, and reuse it across clusters.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["pods.eks.amazonaws.com"],
        }],
        actions: [
            "sts:AssumeRole",
            "sts:TagSession",
        ],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const exampleS3 = new aws.iam.RolePolicyAttachment("exampleS3", {
    policyArn: "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess",
    role: exampleRole.name,
});
const examplePodIdentityAssociation = new aws.eks.PodIdentityAssociation("examplePodIdentityAssociation", {
    clusterName: aws_eks_cluster.example.name,
    namespace: "example",
    serviceAccount: "example-sa",
    roleArn: exampleRole.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["pods.eks.amazonaws.com"],
    )],
    actions=[
        "sts:AssumeRole",
        "sts:TagSession",
    ],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
example_s3 = aws.iam.RolePolicyAttachment("exampleS3",
    policy_arn="arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess",
    role=example_role.name)
example_pod_identity_association = aws.eks.PodIdentityAssociation("examplePodIdentityAssociation",
    cluster_name=aws_eks_cluster["example"]["name"],
    namespace="example",
    service_account="example-sa",
    role_arn=example_role.arn)
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
                            "pods.eks.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                    "sts:TagSession",
                },
            },
        },
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleS3 = new Aws.Iam.RolePolicyAttachment("exampleS3", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess",
        Role = exampleRole.Name,
    });

    var examplePodIdentityAssociation = new Aws.Eks.PodIdentityAssociation("examplePodIdentityAssociation", new()
    {
        ClusterName = aws_eks_cluster.Example.Name,
        Namespace = "example",
        ServiceAccount = "example-sa",
        RoleArn = exampleRole.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
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
								"pods.eks.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
						"sts:TagSession",
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
		_, err = iam.NewRolePolicyAttachment(ctx, "exampleS3", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"),
			Role:      exampleRole.Name,
		})
		if err != nil {
			return err
		}
		_, err = eks.NewPodIdentityAssociation(ctx, "examplePodIdentityAssociation", &eks.PodIdentityAssociationArgs{
			ClusterName:    pulumi.Any(aws_eks_cluster.Example.Name),
			Namespace:      pulumi.String("example"),
			ServiceAccount: pulumi.String("example-sa"),
			RoleArn:        exampleRole.Arn,
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
import com.pulumi.aws.eks.PodIdentityAssociation;
import com.pulumi.aws.eks.PodIdentityAssociationArgs;
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
                    .identifiers("pods.eks.amazonaws.com")
                    .build())
                .actions(                
                    "sts:AssumeRole",
                    "sts:TagSession")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleS3 = new RolePolicyAttachment("exampleS3", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess")
            .role(exampleRole.name())
            .build());

        var examplePodIdentityAssociation = new PodIdentityAssociation("examplePodIdentityAssociation", PodIdentityAssociationArgs.builder()        
            .clusterName(aws_eks_cluster.example().name())
            .namespace("example")
            .serviceAccount("example-sa")
            .roleArn(exampleRole.arn())
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
  exampleS3:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess
      role: ${exampleRole.name}
  examplePodIdentityAssociation:
    type: aws:eks:PodIdentityAssociation
    properties:
      clusterName: ${aws_eks_cluster.example.name}
      namespace: example
      serviceAccount: example-sa
      roleArn: ${exampleRole.arn}
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
                  - pods.eks.amazonaws.com
            actions:
              - sts:AssumeRole
              - sts:TagSession
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EKS (Elastic Kubernetes) Pod Identity Association using the `cluster_name` and `association_id` separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:eks/podIdentityAssociation:PodIdentityAssociation example example,a-12345678
```
 