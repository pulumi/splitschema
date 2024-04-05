Manages an App Runner VPC Connector.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const connector = new aws.apprunner.VpcConnector("connector", {
    securityGroups: [
        "sg1",
        "sg2",
    ],
    subnets: [
        "subnet1",
        "subnet2",
    ],
    vpcConnectorName: "name",
});
```
```python
import pulumi
import pulumi_aws as aws

connector = aws.apprunner.VpcConnector("connector",
    security_groups=[
        "sg1",
        "sg2",
    ],
    subnets=[
        "subnet1",
        "subnet2",
    ],
    vpc_connector_name="name")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var connector = new Aws.AppRunner.VpcConnector("connector", new()
    {
        SecurityGroups = new[]
        {
            "sg1",
            "sg2",
        },
        Subnets = new[]
        {
            "subnet1",
            "subnet2",
        },
        VpcConnectorName = "name",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apprunner"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apprunner.NewVpcConnector(ctx, "connector", &apprunner.VpcConnectorArgs{
			SecurityGroups: pulumi.StringArray{
				pulumi.String("sg1"),
				pulumi.String("sg2"),
			},
			Subnets: pulumi.StringArray{
				pulumi.String("subnet1"),
				pulumi.String("subnet2"),
			},
			VpcConnectorName: pulumi.String("name"),
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
import com.pulumi.aws.apprunner.VpcConnector;
import com.pulumi.aws.apprunner.VpcConnectorArgs;
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
        var connector = new VpcConnector("connector", VpcConnectorArgs.builder()        
            .securityGroups(            
                "sg1",
                "sg2")
            .subnets(            
                "subnet1",
                "subnet2")
            .vpcConnectorName("name")
            .build());

    }
}
```
```yaml
resources:
  connector:
    type: aws:apprunner:VpcConnector
    properties:
      securityGroups:
        - sg1
        - sg2
      subnets:
        - subnet1
        - subnet2
      vpcConnectorName: name
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Runner vpc connector using the `arn`. For example:

```sh
 $ pulumi import aws:apprunner/vpcConnector:VpcConnector example arn:aws:apprunner:us-east-1:1234567890:vpcconnector/example/1/0a03292a89764e5882c41d8f991c82fe
```
 