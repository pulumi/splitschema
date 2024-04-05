Manages an Image Builder Distribution Configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.imagebuilder.DistributionConfiguration("example", {distributions: [{
    amiDistributionConfiguration: {
        amiTags: {
            CostCenter: "IT",
        },
        launchPermission: {
            userIds: ["123456789012"],
        },
        name: "example-{{ imagebuilder:buildDate }}",
    },
    launchTemplateConfigurations: [{
        launchTemplateId: "lt-0aaa1bcde2ff3456",
    }],
    region: "us-east-1",
}]});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.imagebuilder.DistributionConfiguration("example", distributions=[aws.imagebuilder.DistributionConfigurationDistributionArgs(
    ami_distribution_configuration=aws.imagebuilder.DistributionConfigurationDistributionAmiDistributionConfigurationArgs(
        ami_tags={
            "CostCenter": "IT",
        },
        launch_permission=aws.imagebuilder.DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs(
            user_ids=["123456789012"],
        ),
        name="example-{{ imagebuilder:buildDate }}",
    ),
    launch_template_configurations=[aws.imagebuilder.DistributionConfigurationDistributionLaunchTemplateConfigurationArgs(
        launch_template_id="lt-0aaa1bcde2ff3456",
    )],
    region="us-east-1",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ImageBuilder.DistributionConfiguration("example", new()
    {
        Distributions = new[]
        {
            new Aws.ImageBuilder.Inputs.DistributionConfigurationDistributionArgs
            {
                AmiDistributionConfiguration = new Aws.ImageBuilder.Inputs.DistributionConfigurationDistributionAmiDistributionConfigurationArgs
                {
                    AmiTags = 
                    {
                        { "CostCenter", "IT" },
                    },
                    LaunchPermission = new Aws.ImageBuilder.Inputs.DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs
                    {
                        UserIds = new[]
                        {
                            "123456789012",
                        },
                    },
                    Name = "example-{{ imagebuilder:buildDate }}",
                },
                LaunchTemplateConfigurations = new[]
                {
                    new Aws.ImageBuilder.Inputs.DistributionConfigurationDistributionLaunchTemplateConfigurationArgs
                    {
                        LaunchTemplateId = "lt-0aaa1bcde2ff3456",
                    },
                },
                Region = "us-east-1",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/imagebuilder"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := imagebuilder.NewDistributionConfiguration(ctx, "example", &imagebuilder.DistributionConfigurationArgs{
			Distributions: imagebuilder.DistributionConfigurationDistributionArray{
				&imagebuilder.DistributionConfigurationDistributionArgs{
					AmiDistributionConfiguration: &imagebuilder.DistributionConfigurationDistributionAmiDistributionConfigurationArgs{
						AmiTags: pulumi.StringMap{
							"CostCenter": pulumi.String("IT"),
						},
						LaunchPermission: &imagebuilder.DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs{
							UserIds: pulumi.StringArray{
								pulumi.String("123456789012"),
							},
						},
						Name: pulumi.String("example-{{ imagebuilder:buildDate }}"),
					},
					LaunchTemplateConfigurations: imagebuilder.DistributionConfigurationDistributionLaunchTemplateConfigurationArray{
						&imagebuilder.DistributionConfigurationDistributionLaunchTemplateConfigurationArgs{
							LaunchTemplateId: pulumi.String("lt-0aaa1bcde2ff3456"),
						},
					},
					Region: pulumi.String("us-east-1"),
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
import com.pulumi.aws.imagebuilder.DistributionConfiguration;
import com.pulumi.aws.imagebuilder.DistributionConfigurationArgs;
import com.pulumi.aws.imagebuilder.inputs.DistributionConfigurationDistributionArgs;
import com.pulumi.aws.imagebuilder.inputs.DistributionConfigurationDistributionAmiDistributionConfigurationArgs;
import com.pulumi.aws.imagebuilder.inputs.DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs;
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
        var example = new DistributionConfiguration("example", DistributionConfigurationArgs.builder()        
            .distributions(DistributionConfigurationDistributionArgs.builder()
                .amiDistributionConfiguration(DistributionConfigurationDistributionAmiDistributionConfigurationArgs.builder()
                    .amiTags(Map.of("CostCenter", "IT"))
                    .launchPermission(DistributionConfigurationDistributionAmiDistributionConfigurationLaunchPermissionArgs.builder()
                        .userIds("123456789012")
                        .build())
                    .name("example-{{ imagebuilder:buildDate }}")
                    .build())
                .launchTemplateConfigurations(DistributionConfigurationDistributionLaunchTemplateConfigurationArgs.builder()
                    .launchTemplateId("lt-0aaa1bcde2ff3456")
                    .build())
                .region("us-east-1")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:imagebuilder:DistributionConfiguration
    properties:
      distributions:
        - amiDistributionConfiguration:
            amiTags:
              CostCenter: IT
            launchPermission:
              userIds:
                - '123456789012'
            name: example-{{ imagebuilder:buildDate }}
          launchTemplateConfigurations:
            - launchTemplateId: lt-0aaa1bcde2ff3456
          region: us-east-1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_imagebuilder_distribution_configurations` resources using the Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:imagebuilder/distributionConfiguration:DistributionConfiguration example arn:aws:imagebuilder:us-east-1:123456789012:distribution-configuration/example
```
 