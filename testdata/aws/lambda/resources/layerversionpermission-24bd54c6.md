Provides a Lambda Layer Version Permission resource. It allows you to share you own Lambda Layers to another account by account ID, to all accounts in AWS organization or even to all AWS accounts.

For information about Lambda Layer Permissions and how to use them, see [Using Resource-based Policies for AWS Lambda][1]

> **NOTE:** Setting `skip_destroy` to `true` means that the AWS Provider will _not_ destroy any layer version permission, even when running `pulumi destroy`. Layer version permissions are thus intentional dangling resources that are _not_ managed by Pulumi and may incur extra expense in your AWS account.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lambdaLayerPermission = new aws.lambda.LayerVersionPermission("lambdaLayerPermission", {
    action: "lambda:GetLayerVersion",
    layerName: "arn:aws:lambda:us-west-2:123456654321:layer:test_layer1",
    principal: "111111111111",
    statementId: "dev-account",
    versionNumber: 1,
});
```
```python
import pulumi
import pulumi_aws as aws

lambda_layer_permission = aws.lambda_.LayerVersionPermission("lambdaLayerPermission",
    action="lambda:GetLayerVersion",
    layer_name="arn:aws:lambda:us-west-2:123456654321:layer:test_layer1",
    principal="111111111111",
    statement_id="dev-account",
    version_number=1)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var lambdaLayerPermission = new Aws.Lambda.LayerVersionPermission("lambdaLayerPermission", new()
    {
        Action = "lambda:GetLayerVersion",
        LayerName = "arn:aws:lambda:us-west-2:123456654321:layer:test_layer1",
        Principal = "111111111111",
        StatementId = "dev-account",
        VersionNumber = 1,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewLayerVersionPermission(ctx, "lambdaLayerPermission", &lambda.LayerVersionPermissionArgs{
			Action:        pulumi.String("lambda:GetLayerVersion"),
			LayerName:     pulumi.String("arn:aws:lambda:us-west-2:123456654321:layer:test_layer1"),
			Principal:     pulumi.String("111111111111"),
			StatementId:   pulumi.String("dev-account"),
			VersionNumber: pulumi.Int(1),
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
import com.pulumi.aws.lambda.LayerVersionPermission;
import com.pulumi.aws.lambda.LayerVersionPermissionArgs;
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
        var lambdaLayerPermission = new LayerVersionPermission("lambdaLayerPermission", LayerVersionPermissionArgs.builder()        
            .action("lambda:GetLayerVersion")
            .layerName("arn:aws:lambda:us-west-2:123456654321:layer:test_layer1")
            .principal("111111111111")
            .statementId("dev-account")
            .versionNumber(1)
            .build());

    }
}
```
```yaml
resources:
  lambdaLayerPermission:
    type: aws:lambda:LayerVersionPermission
    properties:
      action: lambda:GetLayerVersion
      layerName: arn:aws:lambda:us-west-2:123456654321:layer:test_layer1
      principal: '111111111111'
      statementId: dev-account
      versionNumber: 1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lambda Layer Permissions using `layer_name` and `version_number`, separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:lambda/layerVersionPermission:LayerVersionPermission example arn:aws:lambda:us-west-2:123456654321:layer:test_layer1,1
```
 