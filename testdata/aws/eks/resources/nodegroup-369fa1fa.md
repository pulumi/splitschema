Manages an EKS Node Group, which can provision and optionally update an Auto Scaling Group of Kubernetes worker nodes compatible with EKS. Additional documentation about this functionality can be found in the [EKS User Guide](https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.eks.NodeGroup("example", {
    clusterName: aws_eks_cluster.example.name,
    nodeRoleArn: aws_iam_role.example.arn,
    subnetIds: aws_subnet.example.map(__item => __item.id),
    scalingConfig: {
        desiredSize: 1,
        maxSize: 2,
        minSize: 1,
    },
    updateConfig: {
        maxUnavailable: 1,
    },
}, {
    dependsOn: [
        aws_iam_role_policy_attachment["example-AmazonEKSWorkerNodePolicy"],
        aws_iam_role_policy_attachment["example-AmazonEKS_CNI_Policy"],
        aws_iam_role_policy_attachment["example-AmazonEC2ContainerRegistryReadOnly"],
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.eks.NodeGroup("example",
    cluster_name=aws_eks_cluster["example"]["name"],
    node_role_arn=aws_iam_role["example"]["arn"],
    subnet_ids=[__item["id"] for __item in aws_subnet["example"]],
    scaling_config=aws.eks.NodeGroupScalingConfigArgs(
        desired_size=1,
        max_size=2,
        min_size=1,
    ),
    update_config=aws.eks.NodeGroupUpdateConfigArgs(
        max_unavailable=1,
    ),
    opts=pulumi.ResourceOptions(depends_on=[
            aws_iam_role_policy_attachment["example-AmazonEKSWorkerNodePolicy"],
            aws_iam_role_policy_attachment["example-AmazonEKS_CNI_Policy"],
            aws_iam_role_policy_attachment["example-AmazonEC2ContainerRegistryReadOnly"],
        ]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Eks.NodeGroup("example", new()
    {
        ClusterName = aws_eks_cluster.Example.Name,
        NodeRoleArn = aws_iam_role.Example.Arn,
        SubnetIds = aws_subnet.Example.Select(__item => __item.Id).ToList(),
        ScalingConfig = new Aws.Eks.Inputs.NodeGroupScalingConfigArgs
        {
            DesiredSize = 1,
            MaxSize = 2,
            MinSize = 1,
        },
        UpdateConfig = new Aws.Eks.Inputs.NodeGroupUpdateConfigArgs
        {
            MaxUnavailable = 1,
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_iam_role_policy_attachment.Example_AmazonEKSWorkerNodePolicy,
            aws_iam_role_policy_attachment.Example_AmazonEKS_CNI_Policy,
            aws_iam_role_policy_attachment.Example_AmazonEC2ContainerRegistryReadOnly,
        },
    });

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
var splat0 []interface{}
for _, val0 := range aws_subnet.Example {
splat0 = append(splat0, val0.Id)
}
_, err := eks.NewNodeGroup(ctx, "example", &eks.NodeGroupArgs{
ClusterName: pulumi.Any(aws_eks_cluster.Example.Name),
NodeRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
SubnetIds: toPulumiArray(splat0),
ScalingConfig: &eks.NodeGroupScalingConfigArgs{
DesiredSize: pulumi.Int(1),
MaxSize: pulumi.Int(2),
MinSize: pulumi.Int(1),
},
UpdateConfig: &eks.NodeGroupUpdateConfigArgs{
MaxUnavailable: pulumi.Int(1),
},
}, pulumi.DependsOn([]pulumi.Resource{
aws_iam_role_policy_attachment.ExampleAmazonEKSWorkerNodePolicy,
aws_iam_role_policy_attachment.ExampleAmazonEKS_CNI_Policy,
aws_iam_role_policy_attachment.ExampleAmazonEC2ContainerRegistryReadOnly,
}))
if err != nil {
return err
}
return nil
})
}
func toPulumiArray(arr []) pulumi.Array {
var pulumiArr pulumi.Array
for _, v := range arr {
pulumiArr = append(pulumiArr, pulumi.(v))
}
return pulumiArr
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.eks.NodeGroup;
import com.pulumi.aws.eks.NodeGroupArgs;
import com.pulumi.aws.eks.inputs.NodeGroupScalingConfigArgs;
import com.pulumi.aws.eks.inputs.NodeGroupUpdateConfigArgs;
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
        var example = new NodeGroup("example", NodeGroupArgs.builder()        
            .clusterName(aws_eks_cluster.example().name())
            .nodeRoleArn(aws_iam_role.example().arn())
            .subnetIds(aws_subnet.example().stream().map(element -> element.id()).collect(toList()))
            .scalingConfig(NodeGroupScalingConfigArgs.builder()
                .desiredSize(1)
                .maxSize(2)
                .minSize(1)
                .build())
            .updateConfig(NodeGroupUpdateConfigArgs.builder()
                .maxUnavailable(1)
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    aws_iam_role_policy_attachment.example-AmazonEKSWorkerNodePolicy(),
                    aws_iam_role_policy_attachment.example-AmazonEKS_CNI_Policy(),
                    aws_iam_role_policy_attachment.example-AmazonEC2ContainerRegistryReadOnly())
                .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Ignoring Changes to Desired Size

You can utilize [ignoreChanges](https://www.pulumi.com/docs/intro/concepts/programming-model/#ignorechanges) create an EKS Node Group with an initial size of running instances, then ignore any changes to that count caused externally (e.g. Application Autoscaling).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ... other configurations ...
const example = new aws.eks.NodeGroup("example", {scalingConfig: {
    desiredSize: 2,
}});
```
```python
import pulumi
import pulumi_aws as aws

# ... other configurations ...
example = aws.eks.NodeGroup("example", scaling_config=aws.eks.NodeGroupScalingConfigArgs(
    desired_size=2,
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ... other configurations ...
    var example = new Aws.Eks.NodeGroup("example", new()
    {
        ScalingConfig = new Aws.Eks.Inputs.NodeGroupScalingConfigArgs
        {
            DesiredSize = 2,
        },
    });

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
		_, err := eks.NewNodeGroup(ctx, "example", &eks.NodeGroupArgs{
			ScalingConfig: &eks.NodeGroupScalingConfigArgs{
				DesiredSize: pulumi.Int(2),
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
import com.pulumi.aws.eks.NodeGroup;
import com.pulumi.aws.eks.NodeGroupArgs;
import com.pulumi.aws.eks.inputs.NodeGroupScalingConfigArgs;
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
        var example = new NodeGroup("example", NodeGroupArgs.builder()        
            .scalingConfig(NodeGroupScalingConfigArgs.builder()
                .desiredSize(2)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:eks:NodeGroup
    properties:
      scalingConfig:
        desiredSize: 2
```
{{% /example %}}
{{% example %}}
### Example IAM Role for EKS Node Group

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iam.Role("example", {assumeRolePolicy: JSON.stringify({
    Statement: [{
        Action: "sts:AssumeRole",
        Effect: "Allow",
        Principal: {
            Service: "ec2.amazonaws.com",
        },
    }],
    Version: "2012-10-17",
})});
const example_AmazonEKSWorkerNodePolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSWorkerNodePolicy", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
    role: example.name,
});
const example_AmazonEKSCNIPolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSCNIPolicy", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
    role: example.name,
});
const example_AmazonEC2ContainerRegistryReadOnly = new aws.iam.RolePolicyAttachment("example-AmazonEC2ContainerRegistryReadOnly", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
    role: example.name,
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.iam.Role("example", assume_role_policy=json.dumps({
    "Statement": [{
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Principal": {
            "Service": "ec2.amazonaws.com",
        },
    }],
    "Version": "2012-10-17",
}))
example__amazon_eks_worker_node_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSWorkerNodePolicy",
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
    role=example.name)
example__amazon_ekscni_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSCNIPolicy",
    policy_arn="arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
    role=example.name)
example__amazon_ec2_container_registry_read_only = aws.iam.RolePolicyAttachment("example-AmazonEC2ContainerRegistryReadOnly",
    policy_arn="arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
    role=example.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Iam.Role("example", new()
    {
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = "sts:AssumeRole",
                    ["Effect"] = "Allow",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "ec2.amazonaws.com",
                    },
                },
            },
            ["Version"] = "2012-10-17",
        }),
    });

    var example_AmazonEKSWorkerNodePolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSWorkerNodePolicy", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
        Role = example.Name,
    });

    var example_AmazonEKSCNIPolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSCNIPolicy", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
        Role = example.Name,
    });

    var example_AmazonEC2ContainerRegistryReadOnly = new Aws.Iam.RolePolicyAttachment("example-AmazonEC2ContainerRegistryReadOnly", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
        Role = example.Name,
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Service": "ec2.amazonaws.com",
					},
				},
			},
			"Version": "2012-10-17",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		example, err := iam.NewRole(ctx, "example", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "example-AmazonEKSWorkerNodePolicy", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"),
			Role:      example.Name,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "example-AmazonEKSCNIPolicy", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"),
			Role:      example.Name,
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "example-AmazonEC2ContainerRegistryReadOnly", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"),
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
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
        var example = new Role("example", RoleArgs.builder()        
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", "sts:AssumeRole"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "ec2.amazonaws.com")
                        ))
                    ))),
                    jsonProperty("Version", "2012-10-17")
                )))
            .build());

        var example_AmazonEKSWorkerNodePolicy = new RolePolicyAttachment("example-AmazonEKSWorkerNodePolicy", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy")
            .role(example.name())
            .build());

        var example_AmazonEKSCNIPolicy = new RolePolicyAttachment("example-AmazonEKSCNIPolicy", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy")
            .role(example.name())
            .build());

        var example_AmazonEC2ContainerRegistryReadOnly = new RolePolicyAttachment("example-AmazonEC2ContainerRegistryReadOnly", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly")
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
      assumeRolePolicy:
        fn::toJSON:
          Statement:
            - Action: sts:AssumeRole
              Effect: Allow
              Principal:
                Service: ec2.amazonaws.com
          Version: 2012-10-17
  example-AmazonEKSWorkerNodePolicy:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy
      role: ${example.name}
  example-AmazonEKSCNIPolicy:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy
      role: ${example.name}
  example-AmazonEC2ContainerRegistryReadOnly:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly
      role: ${example.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EKS Node Groups using the `cluster_name` and `node_group_name` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:eks/nodeGroup:NodeGroup my_node_group my_cluster:my_node_group
```
 