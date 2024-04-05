Provides an Elastic Container Registry Policy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentCallerIdentity = aws.getCallerIdentity({});
const currentRegion = aws.getRegion({});
const currentPartition = aws.getPartition({});
const example = new aws.ecr.RegistryPolicy("example", {policy: Promise.all([currentPartition, currentCallerIdentity, currentPartition, currentRegion, currentCallerIdentity]).then(([currentPartition, currentCallerIdentity, currentPartition1, currentRegion, currentCallerIdentity1]) => JSON.stringify({
    Version: "2012-10-17",
    Statement: [{
        Sid: "testpolicy",
        Effect: "Allow",
        Principal: {
            AWS: `arn:${currentPartition.partition}:iam::${currentCallerIdentity.accountId}:root`,
        },
        Action: ["ecr:ReplicateImage"],
        Resource: [`arn:${currentPartition1.partition}:ecr:${currentRegion.name}:${currentCallerIdentity1.accountId}:repository/*`],
    }],
}))});
```
```python
import pulumi
import json
import pulumi_aws as aws

current_caller_identity = aws.get_caller_identity()
current_region = aws.get_region()
current_partition = aws.get_partition()
example = aws.ecr.RegistryPolicy("example", policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Sid": "testpolicy",
        "Effect": "Allow",
        "Principal": {
            "AWS": f"arn:{current_partition.partition}:iam::{current_caller_identity.account_id}:root",
        },
        "Action": ["ecr:ReplicateImage"],
        "Resource": [f"arn:{current_partition.partition}:ecr:{current_region.name}:{current_caller_identity.account_id}:repository/*"],
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

    var currentRegion = Aws.GetRegion.Invoke();

    var currentPartition = Aws.GetPartition.Invoke();

    var example = new Aws.Ecr.RegistryPolicy("example", new()
    {
        Policy = Output.Tuple(currentPartition, currentCallerIdentity, currentPartition, currentRegion, currentCallerIdentity).Apply(values =>
        {
            var currentPartition = values.Item1;
            var currentCallerIdentity = values.Item2;
            var currentPartition1 = values.Item3;
            var currentRegion = values.Item4;
            var currentCallerIdentity1 = values.Item5;
            return JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                ["Version"] = "2012-10-17",
                ["Statement"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["Sid"] = "testpolicy",
                        ["Effect"] = "Allow",
                        ["Principal"] = new Dictionary<string, object?>
                        {
                            ["AWS"] = $"arn:{currentPartition.Apply(getPartitionResult => getPartitionResult.Partition)}:iam::{currentCallerIdentity.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId)}:root",
                        },
                        ["Action"] = new[]
                        {
                            "ecr:ReplicateImage",
                        },
                        ["Resource"] = new[]
                        {
                            $"arn:{currentPartition1.Partition}:ecr:{currentRegion.Apply(getRegionResult => getRegionResult.Name)}:{currentCallerIdentity1.AccountId}:repository/*",
                        },
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentCallerIdentity, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
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
					"Sid":    "testpolicy",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"AWS": fmt.Sprintf("arn:%v:iam::%v:root", currentPartition.Partition, currentCallerIdentity.AccountId),
					},
					"Action": []string{
						"ecr:ReplicateImage",
					},
					"Resource": []string{
						fmt.Sprintf("arn:%v:ecr:%v:%v:repository/*", currentPartition.Partition, currentRegion.Name, currentCallerIdentity.AccountId),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = ecr.NewRegistryPolicy(ctx, "example", &ecr.RegistryPolicyArgs{
			Policy: pulumi.String(json0),
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
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.ecr.RegistryPolicy;
import com.pulumi.aws.ecr.RegistryPolicyArgs;
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

        final var currentRegion = AwsFunctions.getRegion();

        final var currentPartition = AwsFunctions.getPartition();

        var example = new RegistryPolicy("example", RegistryPolicyArgs.builder()        
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Sid", "testpolicy"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", String.format("arn:%s:iam::%s:root", currentPartition.applyValue(getPartitionResult -> getPartitionResult.partition()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId())))
                        )),
                        jsonProperty("Action", jsonArray("ecr:ReplicateImage")),
                        jsonProperty("Resource", jsonArray(String.format("arn:%s:ecr:%s:%s:repository/*", currentPartition.applyValue(getPartitionResult -> getPartitionResult.partition()),currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentCallerIdentity.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))))
                    )))
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ecr:RegistryPolicy
    properties:
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Sid: testpolicy
              Effect: Allow
              Principal:
                AWS: arn:${currentPartition.partition}:iam::${currentCallerIdentity.accountId}:root
              Action:
                - ecr:ReplicateImage
              Resource:
                - arn:${currentPartition.partition}:ecr:${currentRegion.name}:${currentCallerIdentity.accountId}:repository/*
variables:
  currentCallerIdentity:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
  currentPartition:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECR Registry Policy using the registry id. For example:

```sh
 $ pulumi import aws:ecr/registryPolicy:RegistryPolicy example 123456789012
```
 