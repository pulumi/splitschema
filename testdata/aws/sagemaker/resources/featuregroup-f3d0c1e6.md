Provides a SageMaker Feature Group resource.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.sagemaker.FeatureGroup("example", {
    featureGroupName: "example",
    recordIdentifierFeatureName: "example",
    eventTimeFeatureName: "example",
    roleArn: aws_iam_role.test.arn,
    featureDefinitions: [{
        featureName: "example",
        featureType: "String",
    }],
    onlineStoreConfig: {
        enableOnlineStore: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sagemaker.FeatureGroup("example",
    feature_group_name="example",
    record_identifier_feature_name="example",
    event_time_feature_name="example",
    role_arn=aws_iam_role["test"]["arn"],
    feature_definitions=[aws.sagemaker.FeatureGroupFeatureDefinitionArgs(
        feature_name="example",
        feature_type="String",
    )],
    online_store_config=aws.sagemaker.FeatureGroupOnlineStoreConfigArgs(
        enable_online_store=True,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Sagemaker.FeatureGroup("example", new()
    {
        FeatureGroupName = "example",
        RecordIdentifierFeatureName = "example",
        EventTimeFeatureName = "example",
        RoleArn = aws_iam_role.Test.Arn,
        FeatureDefinitions = new[]
        {
            new Aws.Sagemaker.Inputs.FeatureGroupFeatureDefinitionArgs
            {
                FeatureName = "example",
                FeatureType = "String",
            },
        },
        OnlineStoreConfig = new Aws.Sagemaker.Inputs.FeatureGroupOnlineStoreConfigArgs
        {
            EnableOnlineStore = true,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewFeatureGroup(ctx, "example", &sagemaker.FeatureGroupArgs{
			FeatureGroupName:            pulumi.String("example"),
			RecordIdentifierFeatureName: pulumi.String("example"),
			EventTimeFeatureName:        pulumi.String("example"),
			RoleArn:                     pulumi.Any(aws_iam_role.Test.Arn),
			FeatureDefinitions: sagemaker.FeatureGroupFeatureDefinitionArray{
				&sagemaker.FeatureGroupFeatureDefinitionArgs{
					FeatureName: pulumi.String("example"),
					FeatureType: pulumi.String("String"),
				},
			},
			OnlineStoreConfig: &sagemaker.FeatureGroupOnlineStoreConfigArgs{
				EnableOnlineStore: pulumi.Bool(true),
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
import com.pulumi.aws.sagemaker.FeatureGroup;
import com.pulumi.aws.sagemaker.FeatureGroupArgs;
import com.pulumi.aws.sagemaker.inputs.FeatureGroupFeatureDefinitionArgs;
import com.pulumi.aws.sagemaker.inputs.FeatureGroupOnlineStoreConfigArgs;
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
        var example = new FeatureGroup("example", FeatureGroupArgs.builder()        
            .featureGroupName("example")
            .recordIdentifierFeatureName("example")
            .eventTimeFeatureName("example")
            .roleArn(aws_iam_role.test().arn())
            .featureDefinitions(FeatureGroupFeatureDefinitionArgs.builder()
                .featureName("example")
                .featureType("String")
                .build())
            .onlineStoreConfig(FeatureGroupOnlineStoreConfigArgs.builder()
                .enableOnlineStore(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:sagemaker:FeatureGroup
    properties:
      featureGroupName: example
      recordIdentifierFeatureName: example
      eventTimeFeatureName: example
      roleArn: ${aws_iam_role.test.arn}
      featureDefinitions:
        - featureName: example
          featureType: String
      onlineStoreConfig:
        enableOnlineStore: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Feature Groups using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/featureGroup:FeatureGroup test_feature_group feature_group-foo
```
 