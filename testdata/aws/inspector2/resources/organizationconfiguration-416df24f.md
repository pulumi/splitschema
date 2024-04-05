Resource for managing an Amazon Inspector Organization Configuration.

> **NOTE:** In order for this resource to work, the account you use must be an Inspector Delegated Admin Account.

> **NOTE:** When this resource is deleted, EC2, ECR, Lambda, and Lambda code scans will no longer be automatically enabled for new members of your Amazon Inspector organization.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.inspector2.OrganizationConfiguration("example", {autoEnable: {
    ec2: true,
    ecr: false,
    lambda: true,
    lambdaCode: true,
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.inspector2.OrganizationConfiguration("example", auto_enable=aws.inspector2.OrganizationConfigurationAutoEnableArgs(
    ec2=True,
    ecr=False,
    lambda_=True,
    lambda_code=True,
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Inspector2.OrganizationConfiguration("example", new()
    {
        AutoEnable = new Aws.Inspector2.Inputs.OrganizationConfigurationAutoEnableArgs
        {
            Ec2 = true,
            Ecr = false,
            Lambda = true,
            LambdaCode = true,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/inspector2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := inspector2.NewOrganizationConfiguration(ctx, "example", &inspector2.OrganizationConfigurationArgs{
			AutoEnable: &inspector2.OrganizationConfigurationAutoEnableArgs{
				Ec2:        pulumi.Bool(true),
				Ecr:        pulumi.Bool(false),
				Lambda:     pulumi.Bool(true),
				LambdaCode: pulumi.Bool(true),
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
import com.pulumi.aws.inspector2.OrganizationConfiguration;
import com.pulumi.aws.inspector2.OrganizationConfigurationArgs;
import com.pulumi.aws.inspector2.inputs.OrganizationConfigurationAutoEnableArgs;
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
        var example = new OrganizationConfiguration("example", OrganizationConfigurationArgs.builder()        
            .autoEnable(OrganizationConfigurationAutoEnableArgs.builder()
                .ec2(true)
                .ecr(false)
                .lambda(true)
                .lambdaCode(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:inspector2:OrganizationConfiguration
    properties:
      autoEnable:
        ec2: true
        ecr: false
        lambda: true
        lambdaCode: true
```
{{% /example %}}
{{% /examples %}}