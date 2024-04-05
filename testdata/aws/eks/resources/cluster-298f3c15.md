Manages an EKS Cluster.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

export = async () => {
    const example = new aws.eks.Cluster("example", {
        roleArn: aws_iam_role.example.arn,
        vpcConfig: {
            subnetIds: [
                aws_subnet.example1.id,
                aws_subnet.example2.id,
            ],
        },
    }, {
        dependsOn: [
            aws_iam_role_policy_attachment["example-AmazonEKSClusterPolicy"],
            aws_iam_role_policy_attachment["example-AmazonEKSVPCResourceController"],
        ],
    });
    return {
        endpoint: example.endpoint,
        "kubeconfig-certificate-authority-data": example.certificateAuthority.apply(certificateAuthority => certificateAuthority.data),
    };
}
```
```python
import pulumi
import pulumi_aws as aws

example = aws.eks.Cluster("example",
    role_arn=aws_iam_role["example"]["arn"],
    vpc_config=aws.eks.ClusterVpcConfigArgs(
        subnet_ids=[
            aws_subnet["example1"]["id"],
            aws_subnet["example2"]["id"],
        ],
    ),
    opts=pulumi.ResourceOptions(depends_on=[
            aws_iam_role_policy_attachment["example-AmazonEKSClusterPolicy"],
            aws_iam_role_policy_attachment["example-AmazonEKSVPCResourceController"],
        ]))
pulumi.export("endpoint", example.endpoint)
pulumi.export("kubeconfig-certificate-authority-data", example.certificate_authority.data)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Eks.Cluster("example", new()
    {
        RoleArn = aws_iam_role.Example.Arn,
        VpcConfig = new Aws.Eks.Inputs.ClusterVpcConfigArgs
        {
            SubnetIds = new[]
            {
                aws_subnet.Example1.Id,
                aws_subnet.Example2.Id,
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_iam_role_policy_attachment.Example_AmazonEKSClusterPolicy,
            aws_iam_role_policy_attachment.Example_AmazonEKSVPCResourceController,
        },
    });

    return new Dictionary<string, object?>
    {
        ["endpoint"] = example.Endpoint,
        ["kubeconfig-certificate-authority-data"] = example.CertificateAuthority.Apply(certificateAuthority => certificateAuthority.Data),
    };
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := eks.NewCluster(ctx, "example", &eks.ClusterArgs{
			RoleArn: pulumi.Any(aws_iam_role.Example.Arn),
			VpcConfig: &eks.ClusterVpcConfigArgs{
				SubnetIds: pulumi.StringArray{
					aws_subnet.Example1.Id,
					aws_subnet.Example2.Id,
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_iam_role_policy_attachment.ExampleAmazonEKSClusterPolicy,
			aws_iam_role_policy_attachment.ExampleAmazonEKSVPCResourceController,
		}))
		if err != nil {
			return err
		}
		ctx.Export("endpoint", example.Endpoint)
		ctx.Export("kubeconfig-certificate-authority-data", example.CertificateAuthority.ApplyT(func(certificateAuthority eks.ClusterCertificateAuthority) (*string, error) {
			return &certificateAuthority.Data, nil
		}).(pulumi.StringPtrOutput))
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.eks.Cluster;
import com.pulumi.aws.eks.ClusterArgs;
import com.pulumi.aws.eks.inputs.ClusterVpcConfigArgs;
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
        var example = new Cluster("example", ClusterArgs.builder()        
            .roleArn(aws_iam_role.example().arn())
            .vpcConfig(ClusterVpcConfigArgs.builder()
                .subnetIds(                
                    aws_subnet.example1().id(),
                    aws_subnet.example2().id())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    aws_iam_role_policy_attachment.example-AmazonEKSClusterPolicy(),
                    aws_iam_role_policy_attachment.example-AmazonEKSVPCResourceController())
                .build());

        ctx.export("endpoint", example.endpoint());
        ctx.export("kubeconfig-certificate-authority-data", example.certificateAuthority().applyValue(certificateAuthority -> certificateAuthority.data()));
    }
}
```
```yaml
resources:
  example:
    type: aws:eks:Cluster
    properties:
      roleArn: ${aws_iam_role.example.arn}
      vpcConfig:
        subnetIds:
          - ${aws_subnet.example1.id}
          - ${aws_subnet.example2.id}
    options:
      dependson:
        - ${aws_iam_role_policy_attachment"example-AmazonEKSClusterPolicy"[%!s(MISSING)]}
        - ${aws_iam_role_policy_attachment"example-AmazonEKSVPCResourceController"[%!s(MISSING)]}
outputs:
  endpoint: ${example.endpoint}
  kubeconfig-certificate-authority-data: ${example.certificateAuthority.data}
```
{{% /example %}}
{{% example %}}
### Example IAM Role for EKS Cluster

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["eks.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const example = new aws.iam.Role("example", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const example_AmazonEKSClusterPolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSClusterPolicy", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
    role: example.name,
});
// Optionally, enable Security Groups for Pods
// Reference: https://docs.aws.amazon.com/eks/latest/userguide/security-groups-for-pods.html
const example_AmazonEKSVPCResourceController = new aws.iam.RolePolicyAttachment("example-AmazonEKSVPCResourceController", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEKSVPCResourceController",
    role: example.name,
});
```
```python
import pulumi
import pulumi_aws as aws

assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["eks.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
example = aws.iam.Role("example", assume_role_policy=assume_role.json)
example__amazon_eks_cluster_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSClusterPolicy",
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
    role=example.name)
# Optionally, enable Security Groups for Pods
# Reference: https://docs.aws.amazon.com/eks/latest/userguide/security-groups-for-pods.html
example__amazon_eksvpc_resource_controller = aws.iam.RolePolicyAttachment("example-AmazonEKSVPCResourceController",
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSVPCResourceController",
    role=example.name)
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
                            "eks.amazonaws.com",
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

    var example = new Aws.Iam.Role("example", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var example_AmazonEKSClusterPolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSClusterPolicy", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
        Role = example.Name,
    });

    // Optionally, enable Security Groups for Pods
    // Reference: https://docs.aws.amazon.com/eks/latest/userguide/security-groups-for-pods.html
    var example_AmazonEKSVPCResourceController = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSVPCResourceController", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSVPCResourceController",
        Role = example.Name,
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
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"eks.amazonaws.com",
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
		example, err := iam.NewRole(ctx, "example", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "example-AmazonEKSClusterPolicy", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"),
			Role:      example.Name,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "example-AmazonEKSVPCResourceController", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSVPCResourceController"),
			Role:      example.Name,
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
                    .identifiers("eks.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var example = new Role("example", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var example_AmazonEKSClusterPolicy = new RolePolicyAttachment("example-AmazonEKSClusterPolicy", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/AmazonEKSClusterPolicy")
            .role(example.name())
            .build());

        var example_AmazonEKSVPCResourceController = new RolePolicyAttachment("example-AmazonEKSVPCResourceController", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/AmazonEKSVPCResourceController")
            .role(example.name())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  example-AmazonEKSClusterPolicy:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/AmazonEKSClusterPolicy
      role: ${example.name}
  # Optionally, enable Security Groups for Pods
  # Reference: https://docs.aws.amazon.com/eks/latest/userguide/security-groups-for-pods.html
  example-AmazonEKSVPCResourceController:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/AmazonEKSVPCResourceController
      role: ${example.name}
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
                  - eks.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% example %}}
### Enabling Control Plane Logging

[EKS Control Plane Logging](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) can be enabled via the `enabled_cluster_log_types` argument. To manage the CloudWatch Log Group retention period, the `aws.cloudwatch.LogGroup` resource can be used.

> The below configuration uses [`dependsOn`](https://www.pulumi.com/docs/intro/concepts/programming-model/#dependson) to prevent ordering issues with EKS automatically creating the log group first and a variable for naming consistency. Other ordering and naming methodologies may be more appropriate for your environment.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const config = new pulumi.Config();
const clusterName = config.get("clusterName") || "example";
const exampleLogGroup = new aws.cloudwatch.LogGroup("exampleLogGroup", {retentionInDays: 7});
// ... potentially other configuration ...
const exampleCluster = new aws.eks.Cluster("exampleCluster", {enabledClusterLogTypes: [
    "api",
    "audit",
]}, {
    dependsOn: [exampleLogGroup],
});
// ... other configuration ...
```
```python
import pulumi
import pulumi_aws as aws

config = pulumi.Config()
cluster_name = config.get("clusterName")
if cluster_name is None:
    cluster_name = "example"
example_log_group = aws.cloudwatch.LogGroup("exampleLogGroup", retention_in_days=7)
# ... potentially other configuration ...
example_cluster = aws.eks.Cluster("exampleCluster", enabled_cluster_log_types=[
    "api",
    "audit",
],
opts=pulumi.ResourceOptions(depends_on=[example_log_group]))
# ... other configuration ...
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var config = new Config();
    var clusterName = config.Get("clusterName") ?? "example";
    var exampleLogGroup = new Aws.CloudWatch.LogGroup("exampleLogGroup", new()
    {
        RetentionInDays = 7,
    });

    // ... potentially other configuration ...
    var exampleCluster = new Aws.Eks.Cluster("exampleCluster", new()
    {
        EnabledClusterLogTypes = new[]
        {
            "api",
            "audit",
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleLogGroup,
        },
    });

    // ... other configuration ...
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		clusterName := "example"
		if param := cfg.Get("clusterName"); param != "" {
			clusterName = param
		}
		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", &cloudwatch.LogGroupArgs{
			RetentionInDays: pulumi.Int(7),
		})
		if err != nil {
			return err
		}
		_, err = eks.NewCluster(ctx, "exampleCluster", &eks.ClusterArgs{
			EnabledClusterLogTypes: pulumi.StringArray{
				pulumi.String("api"),
				pulumi.String("audit"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleLogGroup,
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.cloudwatch.LogGroupArgs;
import com.pulumi.aws.eks.Cluster;
import com.pulumi.aws.eks.ClusterArgs;
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
        final var config = ctx.config();
        final var clusterName = config.get("clusterName").orElse("example");
        var exampleLogGroup = new LogGroup("exampleLogGroup", LogGroupArgs.builder()        
            .retentionInDays(7)
            .build());

        var exampleCluster = new Cluster("exampleCluster", ClusterArgs.builder()        
            .enabledClusterLogTypes(            
                "api",
                "audit")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleLogGroup)
                .build());

    }
}
```
```yaml
configuration:
  clusterName:
    type: string
    default: example
resources:
  exampleCluster:
    type: aws:eks:Cluster
    properties:
      enabledClusterLogTypes:
        - api
        - audit
    options:
      dependson:
        - ${exampleLogGroup}
  exampleLogGroup:
    type: aws:cloudwatch:LogGroup
    properties:
      retentionInDays: 7
```
{{% /example %}}
{{% example %}}
### EKS Cluster on AWS Outpost

[Creating a local Amazon EKS cluster on an AWS Outpost](https://docs.aws.amazon.com/eks/latest/userguide/create-cluster-outpost.html)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: data.aws_iam_policy_document.example_assume_role_policy.json});
const exampleCluster = new aws.eks.Cluster("exampleCluster", {
    roleArn: exampleRole.arn,
    vpcConfig: {
        endpointPrivateAccess: true,
        endpointPublicAccess: false,
    },
    outpostConfig: {
        controlPlaneInstanceType: "m5d.large",
        outpostArns: [data.aws_outposts_outpost.example.arn],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_role = aws.iam.Role("exampleRole", assume_role_policy=data["aws_iam_policy_document"]["example_assume_role_policy"]["json"])
example_cluster = aws.eks.Cluster("exampleCluster",
    role_arn=example_role.arn,
    vpc_config=aws.eks.ClusterVpcConfigArgs(
        endpoint_private_access=True,
        endpoint_public_access=False,
    ),
    outpost_config=aws.eks.ClusterOutpostConfigArgs(
        control_plane_instance_type="m5d.large",
        outpost_arns=[data["aws_outposts_outpost"]["example"]["arn"]],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = data.Aws_iam_policy_document.Example_assume_role_policy.Json,
    });

    var exampleCluster = new Aws.Eks.Cluster("exampleCluster", new()
    {
        RoleArn = exampleRole.Arn,
        VpcConfig = new Aws.Eks.Inputs.ClusterVpcConfigArgs
        {
            EndpointPrivateAccess = true,
            EndpointPublicAccess = false,
        },
        OutpostConfig = new Aws.Eks.Inputs.ClusterOutpostConfigArgs
        {
            ControlPlaneInstanceType = "m5d.large",
            OutpostArns = new[]
            {
                data.Aws_outposts_outpost.Example.Arn,
            },
        },
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
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.Any(data.Aws_iam_policy_document.Example_assume_role_policy.Json),
		})
		if err != nil {
			return err
		}
		_, err = eks.NewCluster(ctx, "exampleCluster", &eks.ClusterArgs{
			RoleArn: exampleRole.Arn,
			VpcConfig: &eks.ClusterVpcConfigArgs{
				EndpointPrivateAccess: pulumi.Bool(true),
				EndpointPublicAccess:  pulumi.Bool(false),
			},
			OutpostConfig: &eks.ClusterOutpostConfigArgs{
				ControlPlaneInstanceType: pulumi.String("m5d.large"),
				OutpostArns: pulumi.StringArray{
					data.Aws_outposts_outpost.Example.Arn,
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.eks.Cluster;
import com.pulumi.aws.eks.ClusterArgs;
import com.pulumi.aws.eks.inputs.ClusterVpcConfigArgs;
import com.pulumi.aws.eks.inputs.ClusterOutpostConfigArgs;
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
        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(data.aws_iam_policy_document().example_assume_role_policy().json())
            .build());

        var exampleCluster = new Cluster("exampleCluster", ClusterArgs.builder()        
            .roleArn(exampleRole.arn())
            .vpcConfig(ClusterVpcConfigArgs.builder()
                .endpointPrivateAccess(true)
                .endpointPublicAccess(false)
                .build())
            .outpostConfig(ClusterOutpostConfigArgs.builder()
                .controlPlaneInstanceType("m5d.large")
                .outpostArns(data.aws_outposts_outpost().example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${data.aws_iam_policy_document.example_assume_role_policy.json}
  exampleCluster:
    type: aws:eks:Cluster
    properties:
      roleArn: ${exampleRole.arn}
      vpcConfig:
        endpointPrivateAccess: true
        endpointPublicAccess: false
      outpostConfig:
        controlPlaneInstanceType: m5d.large
        outpostArns:
          - ${data.aws_outposts_outpost.example.arn}
```

After adding inline IAM Policies (e.g., `aws.iam.RolePolicy` resource) or attaching IAM Policies (e.g., `aws.iam.Policy` resource and `aws.iam.RolePolicyAttachment` resource) with the desired permissions to the IAM Role, annotate the Kubernetes service account (e.g., `kubernetes_service_account` resource) and recreate any pods.
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EKS Clusters using the `name`. For example:

```sh
 $ pulumi import aws:eks/cluster:Cluster my_cluster my_cluster
```
 