Provides an Amazon AppIntegrations Data Integration resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appintegrations.DataIntegration("example", {
    description: "example",
    kmsKey: aws_kms_key.test.arn,
    sourceUri: "Salesforce://AppFlow/example",
    scheduleConfig: {
        firstExecutionFrom: "1439788442681",
        object: "Account",
        scheduleExpression: "rate(1 hour)",
    },
    tags: {
        Key1: "Value1",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appintegrations.DataIntegration("example",
    description="example",
    kms_key=aws_kms_key["test"]["arn"],
    source_uri="Salesforce://AppFlow/example",
    schedule_config=aws.appintegrations.DataIntegrationScheduleConfigArgs(
        first_execution_from="1439788442681",
        object="Account",
        schedule_expression="rate(1 hour)",
    ),
    tags={
        "Key1": "Value1",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppIntegrations.DataIntegration("example", new()
    {
        Description = "example",
        KmsKey = aws_kms_key.Test.Arn,
        SourceUri = "Salesforce://AppFlow/example",
        ScheduleConfig = new Aws.AppIntegrations.Inputs.DataIntegrationScheduleConfigArgs
        {
            FirstExecutionFrom = "1439788442681",
            Object = "Account",
            ScheduleExpression = "rate(1 hour)",
        },
        Tags = 
        {
            { "Key1", "Value1" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appintegrations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appintegrations.NewDataIntegration(ctx, "example", &appintegrations.DataIntegrationArgs{
			Description: pulumi.String("example"),
			KmsKey:      pulumi.Any(aws_kms_key.Test.Arn),
			SourceUri:   pulumi.String("Salesforce://AppFlow/example"),
			ScheduleConfig: &appintegrations.DataIntegrationScheduleConfigArgs{
				FirstExecutionFrom: pulumi.String("1439788442681"),
				Object:             pulumi.String("Account"),
				ScheduleExpression: pulumi.String("rate(1 hour)"),
			},
			Tags: pulumi.StringMap{
				"Key1": pulumi.String("Value1"),
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
import com.pulumi.aws.appintegrations.DataIntegration;
import com.pulumi.aws.appintegrations.DataIntegrationArgs;
import com.pulumi.aws.appintegrations.inputs.DataIntegrationScheduleConfigArgs;
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
        var example = new DataIntegration("example", DataIntegrationArgs.builder()        
            .description("example")
            .kmsKey(aws_kms_key.test().arn())
            .sourceUri("Salesforce://AppFlow/example")
            .scheduleConfig(DataIntegrationScheduleConfigArgs.builder()
                .firstExecutionFrom("1439788442681")
                .object("Account")
                .scheduleExpression("rate(1 hour)")
                .build())
            .tags(Map.of("Key1", "Value1"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appintegrations:DataIntegration
    properties:
      description: example
      kmsKey: ${aws_kms_key.test.arn}
      sourceUri: Salesforce://AppFlow/example
      scheduleConfig:
        firstExecutionFrom: '1439788442681'
        object: Account
        scheduleExpression: rate(1 hour)
      tags:
        Key1: Value1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon AppIntegrations Data Integrations using the `id`. For example:

```sh
 $ pulumi import aws:appintegrations/dataIntegration:DataIntegration example 12345678-1234-1234-1234-123456789123
```
 