Deploys an Application CloudFormation Stack from the Serverless Application Repository.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const currentPartition = aws.getPartition({});
const currentRegion = aws.getRegion({});
const postgres_rotator = new aws.serverlessrepository.CloudFormationStack("postgres-rotator", {
    applicationId: "arn:aws:serverlessrepo:us-east-1:297356227824:applications/SecretsManagerRDSPostgreSQLRotationSingleUser",
    capabilities: [
        "CAPABILITY_IAM",
        "CAPABILITY_RESOURCE_POLICY",
    ],
    parameters: {
        endpoint: Promise.all([currentRegion, currentPartition]).then(([currentRegion, currentPartition]) => `secretsmanager.${currentRegion.name}.${currentPartition.dnsSuffix}`),
        functionName: "func-postgres-rotator",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

current_partition = aws.get_partition()
current_region = aws.get_region()
postgres_rotator = aws.serverlessrepository.CloudFormationStack("postgres-rotator",
    application_id="arn:aws:serverlessrepo:us-east-1:297356227824:applications/SecretsManagerRDSPostgreSQLRotationSingleUser",
    capabilities=[
        "CAPABILITY_IAM",
        "CAPABILITY_RESOURCE_POLICY",
    ],
    parameters={
        "endpoint": f"secretsmanager.{current_region.name}.{current_partition.dns_suffix}",
        "functionName": "func-postgres-rotator",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var currentPartition = Aws.GetPartition.Invoke();

    var currentRegion = Aws.GetRegion.Invoke();

    var postgres_rotator = new Aws.ServerlessRepository.CloudFormationStack("postgres-rotator", new()
    {
        ApplicationId = "arn:aws:serverlessrepo:us-east-1:297356227824:applications/SecretsManagerRDSPostgreSQLRotationSingleUser",
        Capabilities = new[]
        {
            "CAPABILITY_IAM",
            "CAPABILITY_RESOURCE_POLICY",
        },
        Parameters = 
        {
            { "endpoint", Output.Tuple(currentRegion, currentPartition).Apply(values =>
            {
                var currentRegion = values.Item1;
                var currentPartition = values.Item2;
                return $"secretsmanager.{currentRegion.Apply(getRegionResult => getRegionResult.Name)}.{currentPartition.Apply(getPartitionResult => getPartitionResult.DnsSuffix)}";
            }) },
            { "functionName", "func-postgres-rotator" },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/serverlessrepository"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		currentPartition, err := aws.GetPartition(ctx, nil, nil)
		if err != nil {
			return err
		}
		currentRegion, err := aws.GetRegion(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = serverlessrepository.NewCloudFormationStack(ctx, "postgres-rotator", &serverlessrepository.CloudFormationStackArgs{
			ApplicationId: pulumi.String("arn:aws:serverlessrepo:us-east-1:297356227824:applications/SecretsManagerRDSPostgreSQLRotationSingleUser"),
			Capabilities: pulumi.StringArray{
				pulumi.String("CAPABILITY_IAM"),
				pulumi.String("CAPABILITY_RESOURCE_POLICY"),
			},
			Parameters: pulumi.StringMap{
				"endpoint":     pulumi.String(fmt.Sprintf("secretsmanager.%v.%v", currentRegion.Name, currentPartition.DnsSuffix)),
				"functionName": pulumi.String("func-postgres-rotator"),
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetPartitionArgs;
import com.pulumi.aws.inputs.GetRegionArgs;
import com.pulumi.aws.serverlessrepository.CloudFormationStack;
import com.pulumi.aws.serverlessrepository.CloudFormationStackArgs;
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
        final var currentPartition = AwsFunctions.getPartition();

        final var currentRegion = AwsFunctions.getRegion();

        var postgres_rotator = new CloudFormationStack("postgres-rotator", CloudFormationStackArgs.builder()        
            .applicationId("arn:aws:serverlessrepo:us-east-1:297356227824:applications/SecretsManagerRDSPostgreSQLRotationSingleUser")
            .capabilities(            
                "CAPABILITY_IAM",
                "CAPABILITY_RESOURCE_POLICY")
            .parameters(Map.ofEntries(
                Map.entry("endpoint", String.format("secretsmanager.%s.%s", currentRegion.applyValue(getRegionResult -> getRegionResult.name()),currentPartition.applyValue(getPartitionResult -> getPartitionResult.dnsSuffix()))),
                Map.entry("functionName", "func-postgres-rotator")
            ))
            .build());

    }
}
```
```yaml
resources:
  postgres-rotator:
    type: aws:serverlessrepository:CloudFormationStack
    properties:
      applicationId: arn:aws:serverlessrepo:us-east-1:297356227824:applications/SecretsManagerRDSPostgreSQLRotationSingleUser
      capabilities:
        - CAPABILITY_IAM
        - CAPABILITY_RESOURCE_POLICY
      parameters:
        endpoint: secretsmanager.${currentRegion.name}.${currentPartition.dnsSuffix}
        functionName: func-postgres-rotator
variables:
  currentPartition:
    fn::invoke:
      Function: aws:getPartition
      Arguments: {}
  currentRegion:
    fn::invoke:
      Function: aws:getRegion
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Serverless Application Repository Stack using the CloudFormation Stack name (with or without the `serverlessrepo-` prefix) or the CloudFormation Stack ID. For example:

```sh
 $ pulumi import aws:serverlessrepository/cloudFormationStack:CloudFormationStack example serverlessrepo-postgres-rotator
```
 