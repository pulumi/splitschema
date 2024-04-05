Resource for managing an AWS VPC Lattice Resource Policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentCallerIdentity = aws.getCallerIdentity({});
const currentPartition = aws.getPartition({});
const exampleServiceNetwork = new aws.vpclattice.ServiceNetwork("exampleServiceNetwork", {});
const exampleResourcePolicy = new aws.vpclattice.ResourcePolicy("exampleResourcePolicy", {
    resourceArn: exampleServiceNetwork.arn,
    policy: pulumi.all([currentPartition, currentCallerIdentity, exampleServiceNetwork.arn]).apply(([currentPartition, currentCallerIdentity, arn]) => JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Sid: "test-pol-principals-6",
            Effect: "Allow",
            Principal: {
                AWS: `arn:${currentPartition.partition}:iam::${currentCallerIdentity.accountId}:root`,
            },
            Action: [
                "vpc-lattice:CreateServiceNetworkVpcAssociation",
                "vpc-lattice:CreateServiceNetworkServiceAssociation",
                "vpc-lattice:GetServiceNetwork",
            ],
            Resource: arn,
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
example_service_network = aws.vpclattice.ServiceNetwork("exampleServiceNetwork")
example_resource_policy = aws.vpclattice.ResourcePolicy("exampleResourcePolicy",
    resource_arn=example_service_network.arn,
    policy=example_service_network.arn.apply(lambda arn: json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Sid": "test-pol-principals-6",
            "Effect": "Allow",
            "Principal": {
                "AWS": f"arn:{current_partition.partition}:iam::{current_caller_identity.account_id}:root",
            },
            "Action": [
                "vpc-lattice:CreateServiceNetworkVpcAssociation",
                "vpc-lattice:CreateServiceNetworkServiceAssociation",
                "vpc-lattice:GetServiceNetwork",
            ],
            "Resource": arn,
        }],
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
    var currentCallerIdentity = Aws.GetCallerIdentity.Invoke();

    var currentPartition = Aws.GetPartition.Invoke();

    var exampleServiceNetwork = new Aws.VpcLattice.ServiceNetwork("exampleServiceNetwork");

    var exampleResourcePolicy = new Aws.VpcLattice.ResourcePolicy("exampleResourcePolicy", new()
    {
        ResourceArn = exampleServiceNetwork.Arn,
        Policy = Output.Tuple(currentPartition, currentCallerIdentity, exampleServiceNetwork.Arn).Apply(values =>
        {
            var currentPartition = values.Item1;
            var currentCallerIdentity = values.Item2;
            var arn = values.Item3;
            return JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                ["Version"] = "2012-10-17",
                ["Statement"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["Sid"] = "test-pol-principals-6",
                        ["Effect"] = "Allow",
                        ["Principal"] = new Dictionary<string, object?>
                        {
                            ["AWS"] = $"arn:{currentPartition.Apply(getPartitionResult => getPartitionResult.Partition)}:iam::{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:root",
                        },
                        ["Action"] = new[]
                        {
                            "vpc-lattice:CreateServiceNetworkVpcAssociation",
                            "vpc-lattice:CreateServiceNetworkServiceAssociation",
                            "vpc-lattice:GetServiceNetwork",
                        },
                        ["Resource"] = arn,
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
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
		exampleServiceNetwork, err := vpclattice.NewServiceNetwork(ctx, "exampleServiceNetwork", nil)
		if err != nil {
			return err
		}
		_, err = vpclattice.NewResourcePolicy(ctx, "exampleResourcePolicy", &vpclattice.ResourcePolicyArgs{
			ResourceArn: exampleServiceNetwork.Arn,
			Policy: exampleServiceNetwork.Arn.ApplyT(func(arn string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON0, err := json.Marshal(map[string]interface{}{
					"Version": "2012-10-17",
					"Statement": []map[string]interface{}{
						map[string]interface{}{
							"Sid":    "test-pol-principals-6",
							"Effect": "Allow",
							"Principal": map[string]interface{}{
								"AWS": fmt.Sprintf("arn:%v:iam::%v:root", currentPartition.Partition, currentCallerIdentity.AccountId),
							},
							"Action": []string{
								"vpc-lattice:CreateServiceNetworkVpcAssociation",
								"vpc-lattice:CreateServiceNetworkServiceAssociation",
								"vpc-lattice:GetServiceNetwork",
							},
							"Resource": arn,
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.vpclattice.ServiceNetwork;
import com.pulumi.aws.vpclattice.ResourcePolicy;
import com.pulumi.aws.vpclattice.ResourcePolicyArgs;
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

        var exampleServiceNetwork = new ServiceNetwork("exampleServiceNetwork");

        var exampleResourcePolicy = new ResourcePolicy("exampleResourcePolicy", ResourcePolicyArgs.builder()        
            .resourceArn(exampleServiceNetwork.arn())
            .policy(exampleServiceNetwork.arn().applyValue(arn -> serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Sid", "test-pol-principals-6"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", String.format("arn:%s:iam::%s:root", currentPartition.applyValue(getPartitionResult -> getPartitionResult.partition()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                        )),
                        jsonProperty("Action", jsonArray(
                            "vpc-lattice:CreateServiceNetworkVpcAssociation", 
                            "vpc-lattice:CreateServiceNetworkServiceAssociation", 
                            "vpc-lattice:GetServiceNetwork"
                        )),
                        jsonProperty("Resource", arn)
                    )))
                ))))
            .build());

    }
}
```
```yaml
resources:
  exampleServiceNetwork:
    type: aws:vpclattice:ServiceNetwork
  exampleResourcePolicy:
    type: aws:vpclattice:ResourcePolicy
    properties:
      resourceArn: ${exampleServiceNetwork.arn}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Sid: test-pol-principals-6
              Effect: Allow
              Principal:
                AWS: arn:${currentPartition.partition}:iam::${currentCallerIdentity.accountId}:root
              Action:
                - vpc-lattice:CreateServiceNetworkVpcAssociation
                - vpc-lattice:CreateServiceNetworkServiceAssociation
                - vpc-lattice:GetServiceNetwork
              Resource: ${exampleServiceNetwork.arn}
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

Using `pulumi import`, import VPC Lattice Resource Policy using the `resource_arn`. For example:

```sh
 $ pulumi import aws:vpclattice/resourcePolicy:ResourcePolicy example rft-8012925589
```
 