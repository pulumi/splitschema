Provides a Cognito Risk Configuration resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cognito.RiskConfiguration("example", {
    userPoolId: aws_cognito_user_pool.example.id,
    riskExceptionConfiguration: {
        blockedIpRangeLists: ["10.10.10.10/32"],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cognito.RiskConfiguration("example",
    user_pool_id=aws_cognito_user_pool["example"]["id"],
    risk_exception_configuration=aws.cognito.RiskConfigurationRiskExceptionConfigurationArgs(
        blocked_ip_range_lists=["10.10.10.10/32"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Cognito.RiskConfiguration("example", new()
    {
        UserPoolId = aws_cognito_user_pool.Example.Id,
        RiskExceptionConfiguration = new Aws.Cognito.Inputs.RiskConfigurationRiskExceptionConfigurationArgs
        {
            BlockedIpRangeLists = new[]
            {
                "10.10.10.10/32",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognito.NewRiskConfiguration(ctx, "example", &cognito.RiskConfigurationArgs{
			UserPoolId: pulumi.Any(aws_cognito_user_pool.Example.Id),
			RiskExceptionConfiguration: &cognito.RiskConfigurationRiskExceptionConfigurationArgs{
				BlockedIpRangeLists: pulumi.StringArray{
					pulumi.String("10.10.10.10/32"),
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
import com.pulumi.aws.cognito.RiskConfiguration;
import com.pulumi.aws.cognito.RiskConfigurationArgs;
import com.pulumi.aws.cognito.inputs.RiskConfigurationRiskExceptionConfigurationArgs;
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
        var example = new RiskConfiguration("example", RiskConfigurationArgs.builder()        
            .userPoolId(aws_cognito_user_pool.example().id())
            .riskExceptionConfiguration(RiskConfigurationRiskExceptionConfigurationArgs.builder()
                .blockedIpRangeLists("10.10.10.10/32")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cognito:RiskConfiguration
    properties:
      userPoolId: ${aws_cognito_user_pool.example.id}
      riskExceptionConfiguration:
        blockedIpRangeLists:
          - 10.10.10.10/32
```
{{% /example %}}
{{% /examples %}}

## Import

Import using the user pool ID and Client ID separated by a `:`:

__Using `pulumi import` to import__ Cognito Risk Configurations using the user pool ID or the user pool ID and Client Id separated by a `:`. For example:

Import using the user pool ID:

```sh
 $ pulumi import aws:cognito/riskConfiguration:RiskConfiguration main example
```
 Import using the user pool ID and Client ID separated by a `:`:

```sh
 $ pulumi import aws:cognito/riskConfiguration:RiskConfiguration main example:example
```
 