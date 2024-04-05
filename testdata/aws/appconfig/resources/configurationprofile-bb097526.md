Provides an AppConfig Configuration Profile resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appconfig.ConfigurationProfile("example", {
    applicationId: aws_appconfig_application.example.id,
    description: "Example Configuration Profile",
    locationUri: "hosted",
    validators: [{
        content: aws_lambda_function.example.arn,
        type: "LAMBDA",
    }],
    tags: {
        Type: "AppConfig Configuration Profile",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appconfig.ConfigurationProfile("example",
    application_id=aws_appconfig_application["example"]["id"],
    description="Example Configuration Profile",
    location_uri="hosted",
    validators=[aws.appconfig.ConfigurationProfileValidatorArgs(
        content=aws_lambda_function["example"]["arn"],
        type="LAMBDA",
    )],
    tags={
        "Type": "AppConfig Configuration Profile",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppConfig.ConfigurationProfile("example", new()
    {
        ApplicationId = aws_appconfig_application.Example.Id,
        Description = "Example Configuration Profile",
        LocationUri = "hosted",
        Validators = new[]
        {
            new Aws.AppConfig.Inputs.ConfigurationProfileValidatorArgs
            {
                Content = aws_lambda_function.Example.Arn,
                Type = "LAMBDA",
            },
        },
        Tags = 
        {
            { "Type", "AppConfig Configuration Profile" },
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
		_, err := appconfig.NewConfigurationProfile(ctx, "example", &appconfig.ConfigurationProfileArgs{
			ApplicationId: pulumi.Any(aws_appconfig_application.Example.Id),
			Description:   pulumi.String("Example Configuration Profile"),
			LocationUri:   pulumi.String("hosted"),
			Validators: appconfig.ConfigurationProfileValidatorArray{
				&appconfig.ConfigurationProfileValidatorArgs{
					Content: pulumi.Any(aws_lambda_function.Example.Arn),
					Type:    pulumi.String("LAMBDA"),
				},
			},
			Tags: pulumi.StringMap{
				"Type": pulumi.String("AppConfig Configuration Profile"),
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
import com.pulumi.aws.appconfig.ConfigurationProfile;
import com.pulumi.aws.appconfig.ConfigurationProfileArgs;
import com.pulumi.aws.appconfig.inputs.ConfigurationProfileValidatorArgs;
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
        var example = new ConfigurationProfile("example", ConfigurationProfileArgs.builder()        
            .applicationId(aws_appconfig_application.example().id())
            .description("Example Configuration Profile")
            .locationUri("hosted")
            .validators(ConfigurationProfileValidatorArgs.builder()
                .content(aws_lambda_function.example().arn())
                .type("LAMBDA")
                .build())
            .tags(Map.of("Type", "AppConfig Configuration Profile"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appconfig:ConfigurationProfile
    properties:
      applicationId: ${aws_appconfig_application.example.id}
      description: Example Configuration Profile
      locationUri: hosted
      validators:
        - content: ${aws_lambda_function.example.arn}
          type: LAMBDA
      tags:
        Type: AppConfig Configuration Profile
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AppConfig Configuration Profiles using the configuration profile ID and application ID separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:appconfig/configurationProfile:ConfigurationProfile example 71abcde:11xxxxx
```
 