Provides a SageMaker App Image Config resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.sagemaker.AppImageConfig("test", {
    appImageConfigName: "example",
    kernelGatewayImageConfig: {
        kernelSpec: {
            name: "example",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.sagemaker.AppImageConfig("test",
    app_image_config_name="example",
    kernel_gateway_image_config=aws.sagemaker.AppImageConfigKernelGatewayImageConfigArgs(
        kernel_spec=aws.sagemaker.AppImageConfigKernelGatewayImageConfigKernelSpecArgs(
            name="example",
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Sagemaker.AppImageConfig("test", new()
    {
        AppImageConfigName = "example",
        KernelGatewayImageConfig = new Aws.Sagemaker.Inputs.AppImageConfigKernelGatewayImageConfigArgs
        {
            KernelSpec = new Aws.Sagemaker.Inputs.AppImageConfigKernelGatewayImageConfigKernelSpecArgs
            {
                Name = "example",
            },
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
		_, err := sagemaker.NewAppImageConfig(ctx, "test", &sagemaker.AppImageConfigArgs{
			AppImageConfigName: pulumi.String("example"),
			KernelGatewayImageConfig: &sagemaker.AppImageConfigKernelGatewayImageConfigArgs{
				KernelSpec: &sagemaker.AppImageConfigKernelGatewayImageConfigKernelSpecArgs{
					Name: pulumi.String("example"),
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
import com.pulumi.aws.sagemaker.AppImageConfig;
import com.pulumi.aws.sagemaker.AppImageConfigArgs;
import com.pulumi.aws.sagemaker.inputs.AppImageConfigKernelGatewayImageConfigArgs;
import com.pulumi.aws.sagemaker.inputs.AppImageConfigKernelGatewayImageConfigKernelSpecArgs;
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
        var test = new AppImageConfig("test", AppImageConfigArgs.builder()        
            .appImageConfigName("example")
            .kernelGatewayImageConfig(AppImageConfigKernelGatewayImageConfigArgs.builder()
                .kernelSpec(AppImageConfigKernelGatewayImageConfigKernelSpecArgs.builder()
                    .name("example")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:sagemaker:AppImageConfig
    properties:
      appImageConfigName: example
      kernelGatewayImageConfig:
        kernelSpec:
          name: example
```
{{% /example %}}
{{% example %}}
### Default File System Config

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.sagemaker.AppImageConfig("test", {
    appImageConfigName: "example",
    kernelGatewayImageConfig: {
        fileSystemConfig: {},
        kernelSpec: {
            name: "example",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.sagemaker.AppImageConfig("test",
    app_image_config_name="example",
    kernel_gateway_image_config=aws.sagemaker.AppImageConfigKernelGatewayImageConfigArgs(
        file_system_config=aws.sagemaker.AppImageConfigKernelGatewayImageConfigFileSystemConfigArgs(),
        kernel_spec=aws.sagemaker.AppImageConfigKernelGatewayImageConfigKernelSpecArgs(
            name="example",
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Sagemaker.AppImageConfig("test", new()
    {
        AppImageConfigName = "example",
        KernelGatewayImageConfig = new Aws.Sagemaker.Inputs.AppImageConfigKernelGatewayImageConfigArgs
        {
            FileSystemConfig = null,
            KernelSpec = new Aws.Sagemaker.Inputs.AppImageConfigKernelGatewayImageConfigKernelSpecArgs
            {
                Name = "example",
            },
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
		_, err := sagemaker.NewAppImageConfig(ctx, "test", &sagemaker.AppImageConfigArgs{
			AppImageConfigName: pulumi.String("example"),
			KernelGatewayImageConfig: &sagemaker.AppImageConfigKernelGatewayImageConfigArgs{
				FileSystemConfig: nil,
				KernelSpec: &sagemaker.AppImageConfigKernelGatewayImageConfigKernelSpecArgs{
					Name: pulumi.String("example"),
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
import com.pulumi.aws.sagemaker.AppImageConfig;
import com.pulumi.aws.sagemaker.AppImageConfigArgs;
import com.pulumi.aws.sagemaker.inputs.AppImageConfigKernelGatewayImageConfigArgs;
import com.pulumi.aws.sagemaker.inputs.AppImageConfigKernelGatewayImageConfigFileSystemConfigArgs;
import com.pulumi.aws.sagemaker.inputs.AppImageConfigKernelGatewayImageConfigKernelSpecArgs;
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
        var test = new AppImageConfig("test", AppImageConfigArgs.builder()        
            .appImageConfigName("example")
            .kernelGatewayImageConfig(AppImageConfigKernelGatewayImageConfigArgs.builder()
                .fileSystemConfig()
                .kernelSpec(AppImageConfigKernelGatewayImageConfigKernelSpecArgs.builder()
                    .name("example")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:sagemaker:AppImageConfig
    properties:
      appImageConfigName: example
      kernelGatewayImageConfig:
        fileSystemConfig: {}
        kernelSpec:
          name: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker App Image Configs using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/appImageConfig:AppImageConfig example example
```
 