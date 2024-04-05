Resource for managing an AWS Managed Streaming for Kafka Cluster Policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentCallerIdentity = aws.getCallerIdentity({});
const currentPartition = aws.getPartition({});
const example = new aws.msk.ClusterPolicy("example", {
    clusterArn: aws_msk_cluster.example.arn,
    policy: Promise.all([currentPartition, currentCallerIdentity]).then(([currentPartition, currentCallerIdentity]) => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Sid: "ExampleMskClusterPolicy",
            Effect: "Allow",
            Principal: {
                AWS: `arn:${currentPartition.partition}:iam::${currentCallerIdentity.accountId}:root`,
            },
            Action: [
                "kafka:Describe*",
                "kafka:Get*",
                "kafka:CreateVpcConnection",
                "kafka:GetBootstrapBrokers",
            ],
            Resource: aws_msk_cluster.example.arn,
        }],
    })),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

current_caller_identity = aws.get_caller_identity()
current_partition = aws.get_partition()
example = aws.msk.ClusterPolicy("example",
    cluster_arn=aws_msk_cluster["example"]["arn"],
    policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Sid": "ExampleMskClusterPolicy",
            "Effect": "Allow",
            "Principal": {
                "AWS": f"arn:{current_partition.partition}:iam::{current_caller_identity.account_id}:root",
            },
            "Action": [
                "kafka:Describe*",
                "kafka:Get*",
                "kafka:CreateVpcConnection",
                "kafka:GetBootstrapBrokers",
            ],
            "Resource": aws_msk_cluster["example"]["arn"],
        }],
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var currentPartition = Aws.GetPartition.Invoke();

    var example = new Aws.Msk.ClusterPolicy("example", new()
    {
        ClusterArn = aws_msk_cluster.Example.Arn,
        Policy = Output.Tuple(currentPartition, currentCallerIdentity).Apply(values =>
        {
            var currentPartition = values.Item1;
            var currentCallerIdentity = values.Item2;
            return JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                ["Version"] = "2012-10-17",
                ["Statement"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["Sid"] = "ExampleMskClusterPolicy",
                        ["Effect"] = "Allow",
                        ["Principal"] = new Dictionary<string, object?>
                        {
                            ["AWS"] = $"arn:{currentPartition.Apply(getPartitionResult => getPartitionResult.Partition)}:iam::{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:root",
                        },
                        ["Action"] = new[]
                        {
                            "kafka:Describe*",
                            "kafka:Get*",
                            "kafka:CreateVpcConnection",
                            "kafka:GetBootstrapBrokers",
                        },
                        ["Resource"] = aws_msk_cluster.Example.Arn,
                    },
                },
            });
        }),
    });

});
```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/msk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentPartition, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Sid":    "ExampleMskClusterPolicy",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"AWS": fmt.Sprintf("arn:%v:iam::%v:root", currentPartition.Partition, currentCallerIdentity.AccountId),
					},
					"Action": []string{
						"kafka:Describe*",
						"kafka:Get*",
						"kafka:CreateVpcConnection",
						"kafka:GetBootstrapBrokers",
					},
					"Resource": aws_msk_cluster.Example.Arn,
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = msk.NewClusterPolicy(ctx, "example", &msk.ClusterPolicyArgs{
			ClusterArn: pulumi.Any(aws_msk_cluster.Example.Arn),
			Policy:     pulumi.String(json0),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.msk.ClusterPolicy;
import com.pulumi.aws.msk.ClusterPolicyArgs;
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
        final var currentCallerIdentity = AwsFunctions.getCallerIdentity();

        final var currentPartition = AwsFunctions.getPartition();

        var example = new ClusterPolicy("example", ClusterPolicyArgs.builder()        
            .clusterArn(aws_msk_cluster.example().arn())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Sid", "ExampleMskClusterPolicy"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", String.format("arn:%s:iam::%s:root", currentPartition.applyValue(getPartitionResult -> getPartitionResult.partition()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                        )),
                        jsonProperty("Action", jsonArray(
                            "kafka:Describe*", 
                            "kafka:Get*", 
                            "kafka:CreateVpcConnection", 
                            "kafka:GetBootstrapBrokers"
                        )),
                        jsonProperty("Resource", aws_msk_cluster.example().arn())
                    )))
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:msk:ClusterPolicy
    properties:
      clusterArn: ${aws_msk_cluster.example.arn}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Sid: ExampleMskClusterPolicy
              Effect: Allow
              Principal:
                AWS: arn:${currentPartition.partition}:iam::${currentCallerIdentity.accountId}:root
              Action:
                - kafka:Describe*
                - kafka:Get*
                - kafka:CreateVpcConnection
                - kafka:GetBootstrapBrokers
              Resource: ${aws_msk_cluster.example.arn}
variables:
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  currentPartition:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Managed Streaming for Kafka Cluster Policy using the `cluster_arn`. For example:

```sh
 $ pulumi import aws:msk/clusterPolicy:ClusterPolicy example arn:aws:kafka:us-west-2:123456789012:cluster/example/279c0212-d057-4dba-9aa9-1c4e5a25bfc7-3
```
 