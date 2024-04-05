Provides an AppConfig Deployment resource for an `aws.appconfig.Application` resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appconfig.Deployment("example", {
    applicationId: aws_appconfig_application.example.id,
    configurationProfileId: aws_appconfig_configuration_profile.example.configuration_profile_id,
    configurationVersion: aws_appconfig_hosted_configuration_version.example.version_number,
    deploymentStrategyId: aws_appconfig_deployment_strategy.example.id,
    description: "My example deployment",
    environmentId: aws_appconfig_environment.example.environment_id,
    kmsKeyIdentifier: aws_kms_key.example.arn,
    tags: {
        Type: "AppConfig Deployment",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appconfig.Deployment("example",
    application_id=aws_appconfig_application["example"]["id"],
    configuration_profile_id=aws_appconfig_configuration_profile["example"]["configuration_profile_id"],
    configuration_version=aws_appconfig_hosted_configuration_version["example"]["version_number"],
    deployment_strategy_id=aws_appconfig_deployment_strategy["example"]["id"],
    description="My example deployment",
    environment_id=aws_appconfig_environment["example"]["environment_id"],
    kms_key_identifier=aws_kms_key["example"]["arn"],
    tags={
        "Type": "AppConfig Deployment",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppConfig.Deployment("example", new()
    {
        ApplicationId = aws_appconfig_application.Example.Id,
        ConfigurationProfileId = aws_appconfig_configuration_profile.Example.Configuration_profile_id,
        ConfigurationVersion = aws_appconfig_hosted_configuration_version.Example.Version_number,
        DeploymentStrategyId = aws_appconfig_deployment_strategy.Example.Id,
        Description = "My example deployment",
        EnvironmentId = aws_appconfig_environment.Example.Environment_id,
        KmsKeyIdentifier = aws_kms_key.Example.Arn,
        Tags = 
        {
            { "Type", "AppConfig Deployment" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appconfig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appconfig.NewDeployment(ctx, "example", &appconfig.DeploymentArgs{
			ApplicationId:          pulumi.Any(aws_appconfig_application.Example.Id),
			ConfigurationProfileId: pulumi.Any(aws_appconfig_configuration_profile.Example.Configuration_profile_id),
			ConfigurationVersion:   pulumi.Any(aws_appconfig_hosted_configuration_version.Example.Version_number),
			DeploymentStrategyId:   pulumi.Any(aws_appconfig_deployment_strategy.Example.Id),
			Description:            pulumi.String("My example deployment"),
			EnvironmentId:          pulumi.Any(aws_appconfig_environment.Example.Environment_id),
			KmsKeyIdentifier:       pulumi.Any(aws_kms_key.Example.Arn),
			Tags: pulumi.StringMap{
				"Type": pulumi.String("AppConfig Deployment"),
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
import com.pulumi.aws.appconfig.Deployment;
import com.pulumi.aws.appconfig.DeploymentArgs;
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
        var example = new Deployment("example", DeploymentArgs.builder()        
            .applicationId(aws_appconfig_application.example().id())
            .configurationProfileId(aws_appconfig_configuration_profile.example().configuration_profile_id())
            .configurationVersion(aws_appconfig_hosted_configuration_version.example().version_number())
            .deploymentStrategyId(aws_appconfig_deployment_strategy.example().id())
            .description("My example deployment")
            .environmentId(aws_appconfig_environment.example().environment_id())
            .kmsKeyIdentifier(aws_kms_key.example().arn())
            .tags(Map.of("Type", "AppConfig Deployment"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appconfig:Deployment
    properties:
      applicationId: ${aws_appconfig_application.example.id}
      configurationProfileId: ${aws_appconfig_configuration_profile.example.configuration_profile_id}
      configurationVersion: ${aws_appconfig_hosted_configuration_version.example.version_number}
      deploymentStrategyId: ${aws_appconfig_deployment_strategy.example.id}
      description: My example deployment
      environmentId: ${aws_appconfig_environment.example.environment_id}
      kmsKeyIdentifier: ${aws_kms_key.example.arn}
      tags:
        Type: AppConfig Deployment
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AppConfig Deployments using the application ID, environment ID, and deployment number separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:appconfig/deployment:Deployment example 71abcde/11xxxxx/1
```
 