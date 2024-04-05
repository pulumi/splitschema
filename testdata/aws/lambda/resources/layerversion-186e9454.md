Provides a Lambda Layer Version resource. Lambda Layers allow you to reuse shared bits of code across multiple lambda functions.

For information about Lambda Layers and how to use them, see [AWS Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html).

> **NOTE:** Setting `skip_destroy` to `true` means that the AWS Provider will _not_ destroy any layer version, even when running destroy. Layer versions are thus intentional dangling resources that are _not_ managed by the provider and may incur extra expense in your AWS account.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lambdaLayer = new aws.lambda.LayerVersion("lambdaLayer", {
    compatibleRuntimes: ["nodejs16.x"],
    code: new pulumi.asset.FileArchive("lambda_layer_payload.zip"),
    layerName: "lambda_layer_name",
});
```
```python
import pulumi
import pulumi_aws as aws

lambda_layer = aws.lambda_.LayerVersion("lambdaLayer",
    compatible_runtimes=["nodejs16.x"],
    code=pulumi.FileArchive("lambda_layer_payload.zip"),
    layer_name="lambda_layer_name")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var lambdaLayer = new Aws.Lambda.LayerVersion("lambdaLayer", new()
    {
        CompatibleRuntimes = new[]
        {
            "nodejs16.x",
        },
        Code = new FileArchive("lambda_layer_payload.zip"),
        LayerName = "lambda_layer_name",
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
		_, err := lambda.NewLayerVersion(ctx, "lambdaLayer", &lambda.LayerVersionArgs{
			CompatibleRuntimes: pulumi.StringArray{
				pulumi.String("nodejs16.x"),
			},
			Code:      pulumi.NewFileArchive("lambda_layer_payload.zip"),
			LayerName: pulumi.String("lambda_layer_name"),
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
import com.pulumi.aws.lambda.LayerVersion;
import com.pulumi.aws.lambda.LayerVersionArgs;
import com.pulumi.asset.FileArchive;
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
        var lambdaLayer = new LayerVersion("lambdaLayer", LayerVersionArgs.builder()        
            .compatibleRuntimes("nodejs16.x")
            .code(new FileArchive("lambda_layer_payload.zip"))
            .layerName("lambda_layer_name")
            .build());

    }
}
```
```yaml
resources:
  lambdaLayer:
    type: aws:lambda:LayerVersion
    properties:
      compatibleRuntimes:
        - nodejs16.x
      code:
        fn::FileArchive: lambda_layer_payload.zip
      layerName: lambda_layer_name
```
{{% /example %}}
{{% /examples %}}
## Specifying the Deployment Package

AWS Lambda Layers expect source code to be provided as a deployment package whose structure varies depending on which `compatible_runtimes` this layer specifies.
See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_PublishLayerVersion.html#SSS-PublishLayerVersion-request-CompatibleRuntimes) for the valid values of `compatible_runtimes`.

Once you have created your deployment package you can specify it either directly as a local file (using the `filename` argument) or
indirectly via Amazon S3 (using the `s3_bucket`, `s3_key` and `s3_object_version` arguments). When providing the deployment
package via S3 it may be useful to use the `aws.s3.BucketObjectv2` resource to upload it.

For larger deployment packages it is recommended by Amazon to upload via S3, since the S3 API has better support for uploading large files efficiently.


## Import

Using `pulumi import`, import Lambda Layers using `arn`. For example:

```sh
 $ pulumi import aws:lambda/layerVersion:LayerVersion test_layer arn:aws:lambda:_REGION_:_ACCOUNT_ID_:layer:_LAYER_NAME_:_LAYER_VERSION_
```
 