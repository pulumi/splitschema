Provides access to all Configuration Properties for an AppConfig Application. This will allow you to pass Configuration
Profile IDs to another resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleConfigurationProfiles = aws.appconfig.getConfigurationProfiles({
    applicationId: "a1d3rpe",
});
const exampleConfigurationProfile = exampleConfigurationProfiles.then(exampleConfigurationProfiles => .map(([, ]) => (aws.appconfig.getConfigurationProfile({
    configurationProfileId: __value,
    applicationId: aws_appconfig_application.example.id,
}))));
```
```python
import pulumi
import pulumi_aws as aws

example_configuration_profiles = aws.appconfig.get_configuration_profiles(application_id="a1d3rpe")
example_configuration_profile = [aws.appconfig.get_configuration_profile(configuration_profile_id=__value,
    application_id=aws_appconfig_application["example"]["id"]) for __key, __value in example_configuration_profiles.configuration_profile_ids]
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleConfigurationProfiles = Aws.AppConfig.GetConfigurationProfiles.Invoke(new()
    {
        ApplicationId = "a1d3rpe",
    });

    var exampleConfigurationProfile = .Select(__value => 
    {
        return Aws.AppConfig.GetConfigurationProfile.Invoke(new()
        {
            ConfigurationProfileId = __value,
            ApplicationId = aws_appconfig_application.Example.Id,
        });
    }).ToList();

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
		exampleConfigurationProfiles, err := appconfig.GetConfigurationProfiles(ctx, &appconfig.GetConfigurationProfilesArgs{
			ApplicationId: "a1d3rpe",
		}, nil)
		if err != nil {
			return err
		}
		_ := "TODO: For expression"
		return nil
	})
}
```
{{% /example %}}
{{% /examples %}}