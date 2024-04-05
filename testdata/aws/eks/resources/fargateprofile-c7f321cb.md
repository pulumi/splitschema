Manages an EKS Fargate Profile.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.eks.FargateProfile("example", {
    clusterName: aws_eks_cluster.example.name,
    podExecutionRoleArn: aws_iam_role.example.arn,
    subnetIds: aws_subnet.example.map(__item => __item.id),
    selectors: [{
        namespace: "example",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.eks.FargateProfile("example",
    cluster_name=aws_eks_cluster["example"]["name"],
    pod_execution_role_arn=aws_iam_role["example"]["arn"],
    subnet_ids=[__item["id"] for __item in aws_subnet["example"]],
    selectors=[aws.eks.FargateProfileSelectorArgs(
        namespace="example",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Eks.FargateProfile("example", new()
    {
        ClusterName = aws_eks_cluster.Example.Name,
        PodExecutionRoleArn = aws_iam_role.Example.Arn,
        SubnetIds = aws_subnet.Example.Select(__item => __item.Id).ToList(),
        Selectors = new[]
        {
            new Aws.Eks.Inputs.FargateProfileSelectorArgs
            {
                Namespace = "example",
            },
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
_, err := eks.NewFargateProfile(ctx, "example", &eks.FargateProfileArgs{
ClusterName: pulumi.Any(aws_eks_cluster.Example.Name),
PodExecutionRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
SubnetIds: toPulumiArray(splat0),
Selectors: eks.FargateProfileSelectorArray{
&eks.FargateProfileSelectorArgs{
Namespace: pulumi.String("example"),
},
},
})
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
import com.pulumi.aws.eks.FargateProfile;
import com.pulumi.aws.eks.FargateProfileArgs;
import com.pulumi.aws.eks.inputs.FargateProfileSelectorArgs;
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
        var example = new FargateProfile("example", FargateProfileArgs.builder()        
            .clusterName(aws_eks_cluster.example().name())
            .podExecutionRoleArn(aws_iam_role.example().arn())
            .subnetIds(aws_subnet.example().stream().map(element -> element.id()).collect(toList()))
            .selectors(FargateProfileSelectorArgs.builder()
                .namespace("example")
                .build())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Example IAM Role for EKS Fargate Profile

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.iam.Role("example", {assumeRolePolicy: JSON.stringify({
    Statement: [{
        Action: "sts:AssumeRole",
        Effect: "Allow",
        Principal: {
            Service: "eks-fargate-pods.amazonaws.com",
        },
    }],
    Version: "2012-10-17",
})});
const example_AmazonEKSFargatePodExecutionRolePolicy = new aws.iam.RolePolicyAttachment("example-AmazonEKSFargatePodExecutionRolePolicy", {
    policyArn: "arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy",
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
            "Service": "eks-fargate-pods.amazonaws.com",
        },
    }],
    "Version": "2012-10-17",
}))
example__amazon_eks_fargate_pod_execution_role_policy = aws.iam.RolePolicyAttachment("example-AmazonEKSFargatePodExecutionRolePolicy",
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy",
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
                        ["Service"] = "eks-fargate-pods.amazonaws.com",
                    },
                },
            },
            ["Version"] = "2012-10-17",
        }),
    });

    var example_AmazonEKSFargatePodExecutionRolePolicy = new Aws.Iam.RolePolicyAttachment("example-AmazonEKSFargatePodExecutionRolePolicy", new()
    {
        PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy",
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
						"Service": "eks-fargate-pods.amazonaws.com",
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
		_, err = iam.NewRolePolicyAttachment(ctx, "example-AmazonEKSFargatePodExecutionRolePolicy", &iam.RolePolicyAttachmentArgs{
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy"),
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
                            jsonProperty("Service", "eks-fargate-pods.amazonaws.com")
                        ))
                    ))),
                    jsonProperty("Version", "2012-10-17")
                )))
            .build());

        var example_AmazonEKSFargatePodExecutionRolePolicy = new RolePolicyAttachment("example-AmazonEKSFargatePodExecutionRolePolicy", RolePolicyAttachmentArgs.builder()        
            .policyArn("arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy")
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
                Service: eks-fargate-pods.amazonaws.com
          Version: 2012-10-17
  example-AmazonEKSFargatePodExecutionRolePolicy:
    type: aws:iam:RolePolicyAttachment
    properties:
      policyArn: arn:aws:iam::aws:policy/AmazonEKSFargatePodExecutionRolePolicy
      role: ${example.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EKS Fargate Profiles using the `cluster_name` and `fargate_profile_name` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:eks/fargateProfile:FargateProfile my_fargate_profile my_cluster:my_fargate_profile
```
 