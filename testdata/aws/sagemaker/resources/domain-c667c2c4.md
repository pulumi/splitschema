Provides a SageMaker Domain resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["sagemaker.amazonaws.com"],
        }],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {
    path: "/",
    assumeRolePolicy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json),
});
const exampleDomain = new aws.sagemaker.Domain("exampleDomain", {
    domainName: "example",
    authMode: "IAM",
    vpcId: aws_vpc.example.id,
    subnetIds: [aws_subnet.example.id],
    defaultUserSettings: {
        executionRole: exampleRole.arn,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["sagemaker.amazonaws.com"],
    )],
)])
example_role = aws.iam.Role("exampleRole",
    path="/",
    assume_role_policy=example_policy_document.json)
example_domain = aws.sagemaker.Domain("exampleDomain",
    domain_name="example",
    auth_mode="IAM",
    vpc_id=aws_vpc["example"]["id"],
    subnet_ids=[aws_subnet["example"]["id"]],
    default_user_settings=aws.sagemaker.DomainDefaultUserSettingsArgs(
        execution_role=example_role.arn,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "sagemaker.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        Path = "/",
        AssumeRolePolicy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleDomain = new Aws.Sagemaker.Domain("exampleDomain", new()
    {
        DomainName = "example",
        AuthMode = "IAM",
        VpcId = aws_vpc.Example.Id,
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
        DefaultUserSettings = new Aws.Sagemaker.Inputs.DomainDefaultUserSettingsArgs
        {
            ExecutionRole = exampleRole.Arn,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"sagemaker.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			Path:             pulumi.String("/"),
			AssumeRolePolicy: *pulumi.String(examplePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = sagemaker.NewDomain(ctx, "exampleDomain", &sagemaker.DomainArgs{
			DomainName: pulumi.String("example"),
			AuthMode:   pulumi.String("IAM"),
			VpcId:      pulumi.Any(aws_vpc.Example.Id),
			SubnetIds: pulumi.StringArray{
				aws_subnet.Example.Id,
			},
			DefaultUserSettings: &sagemaker.DomainDefaultUserSettingsArgs{
				ExecutionRole: exampleRole.Arn,
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.sagemaker.Domain;
import com.pulumi.aws.sagemaker.DomainArgs;
import com.pulumi.aws.sagemaker.inputs.DomainDefaultUserSettingsArgs;
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
        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("sagemaker.amazonaws.com")
                    .build())
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .path("/")
            .assumeRolePolicy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleDomain = new Domain("exampleDomain", DomainArgs.builder()        
            .domainName("example")
            .authMode("IAM")
            .vpcId(aws_vpc.example().id())
            .subnetIds(aws_subnet.example().id())
            .defaultUserSettings(DomainDefaultUserSettingsArgs.builder()
                .executionRole(exampleRole.arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleDomain:
    type: aws:sagemaker:Domain
    properties:
      domainName: example
      authMode: IAM
      vpcId: ${aws_vpc.example.id}
      subnetIds:
        - ${aws_subnet.example.id}
      defaultUserSettings:
        executionRole: ${exampleRole.arn}
  exampleRole:
    type: aws:iam:Role
    properties:
      path: /
      assumeRolePolicy: ${examplePolicyDocument.json}
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            principals:
              - type: Service
                identifiers:
                  - sagemaker.amazonaws.com
```
{{% /example %}}
{{% example %}}
### Using Custom Images

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleImage = new aws.sagemaker.Image("exampleImage", {
    imageName: "example",
    roleArn: aws_iam_role.example.arn,
});
const exampleAppImageConfig = new aws.sagemaker.AppImageConfig("exampleAppImageConfig", {
    appImageConfigName: "example",
    kernelGatewayImageConfig: {
        kernelSpec: {
            name: "example",
        },
    },
});
const exampleImageVersion = new aws.sagemaker.ImageVersion("exampleImageVersion", {
    imageName: exampleImage.id,
    baseImage: "base-image",
});
const exampleDomain = new aws.sagemaker.Domain("exampleDomain", {
    domainName: "example",
    authMode: "IAM",
    vpcId: aws_vpc.example.id,
    subnetIds: [aws_subnet.example.id],
    defaultUserSettings: {
        executionRole: aws_iam_role.example.arn,
        kernelGatewayAppSettings: {
            customImages: [{
                appImageConfigName: exampleAppImageConfig.appImageConfigName,
                imageName: exampleImageVersion.imageName,
            }],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_image = aws.sagemaker.Image("exampleImage",
    image_name="example",
    role_arn=aws_iam_role["example"]["arn"])
example_app_image_config = aws.sagemaker.AppImageConfig("exampleAppImageConfig",
    app_image_config_name="example",
    kernel_gateway_image_config=aws.sagemaker.AppImageConfigKernelGatewayImageConfigArgs(
        kernel_spec=aws.sagemaker.AppImageConfigKernelGatewayImageConfigKernelSpecArgs(
            name="example",
        ),
    ))
example_image_version = aws.sagemaker.ImageVersion("exampleImageVersion",
    image_name=example_image.id,
    base_image="base-image")
example_domain = aws.sagemaker.Domain("exampleDomain",
    domain_name="example",
    auth_mode="IAM",
    vpc_id=aws_vpc["example"]["id"],
    subnet_ids=[aws_subnet["example"]["id"]],
    default_user_settings=aws.sagemaker.DomainDefaultUserSettingsArgs(
        execution_role=aws_iam_role["example"]["arn"],
        kernel_gateway_app_settings=aws.sagemaker.DomainDefaultUserSettingsKernelGatewayAppSettingsArgs(
            custom_images=[aws.sagemaker.DomainDefaultUserSettingsKernelGatewayAppSettingsCustomImageArgs(
                app_image_config_name=example_app_image_config.app_image_config_name,
                image_name=example_image_version.image_name,
            )],
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
    var exampleImage = new Aws.Sagemaker.Image("exampleImage", new()
    {
        ImageName = "example",
        RoleArn = aws_iam_role.Example.Arn,
    });

    var exampleAppImageConfig = new Aws.Sagemaker.AppImageConfig("exampleAppImageConfig", new()
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

    var exampleImageVersion = new Aws.Sagemaker.ImageVersion("exampleImageVersion", new()
    {
        ImageName = exampleImage.Id,
        BaseImage = "base-image",
    });

    var exampleDomain = new Aws.Sagemaker.Domain("exampleDomain", new()
    {
        DomainName = "example",
        AuthMode = "IAM",
        VpcId = aws_vpc.Example.Id,
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
        DefaultUserSettings = new Aws.Sagemaker.Inputs.DomainDefaultUserSettingsArgs
        {
            ExecutionRole = aws_iam_role.Example.Arn,
            KernelGatewayAppSettings = new Aws.Sagemaker.Inputs.DomainDefaultUserSettingsKernelGatewayAppSettingsArgs
            {
                CustomImages = new[]
                {
                    new Aws.Sagemaker.Inputs.DomainDefaultUserSettingsKernelGatewayAppSettingsCustomImageArgs
                    {
                        AppImageConfigName = exampleAppImageConfig.AppImageConfigName,
                        ImageName = exampleImageVersion.ImageName,
                    },
                },
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
		exampleImage, err := sagemaker.NewImage(ctx, "exampleImage", &sagemaker.ImageArgs{
			ImageName: pulumi.String("example"),
			RoleArn:   pulumi.Any(aws_iam_role.Example.Arn),
		})
		if err != nil {
			return err
		}
		exampleAppImageConfig, err := sagemaker.NewAppImageConfig(ctx, "exampleAppImageConfig", &sagemaker.AppImageConfigArgs{
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
		exampleImageVersion, err := sagemaker.NewImageVersion(ctx, "exampleImageVersion", &sagemaker.ImageVersionArgs{
			ImageName: exampleImage.ID(),
			BaseImage: pulumi.String("base-image"),
		})
		if err != nil {
			return err
		}
		_, err = sagemaker.NewDomain(ctx, "exampleDomain", &sagemaker.DomainArgs{
			DomainName: pulumi.String("example"),
			AuthMode:   pulumi.String("IAM"),
			VpcId:      pulumi.Any(aws_vpc.Example.Id),
			SubnetIds: pulumi.StringArray{
				aws_subnet.Example.Id,
			},
			DefaultUserSettings: &sagemaker.DomainDefaultUserSettingsArgs{
				ExecutionRole: pulumi.Any(aws_iam_role.Example.Arn),
				KernelGatewayAppSettings: &sagemaker.DomainDefaultUserSettingsKernelGatewayAppSettingsArgs{
					CustomImages: sagemaker.DomainDefaultUserSettingsKernelGatewayAppSettingsCustomImageArray{
						&sagemaker.DomainDefaultUserSettingsKernelGatewayAppSettingsCustomImageArgs{
							AppImageConfigName: exampleAppImageConfig.AppImageConfigName,
							ImageName:          exampleImageVersion.ImageName,
						},
					},
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
import com.pulumi.aws.sagemaker.Image;
import com.pulumi.aws.sagemaker.ImageArgs;
import com.pulumi.aws.sagemaker.AppImageConfig;
import com.pulumi.aws.sagemaker.AppImageConfigArgs;
import com.pulumi.aws.sagemaker.inputs.AppImageConfigKernelGatewayImageConfigArgs;
import com.pulumi.aws.sagemaker.inputs.AppImageConfigKernelGatewayImageConfigKernelSpecArgs;
import com.pulumi.aws.sagemaker.ImageVersion;
import com.pulumi.aws.sagemaker.ImageVersionArgs;
import com.pulumi.aws.sagemaker.Domain;
import com.pulumi.aws.sagemaker.DomainArgs;
import com.pulumi.aws.sagemaker.inputs.DomainDefaultUserSettingsArgs;
import com.pulumi.aws.sagemaker.inputs.DomainDefaultUserSettingsKernelGatewayAppSettingsArgs;
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
        var exampleImage = new Image("exampleImage", ImageArgs.builder()        
            .imageName("example")
            .roleArn(aws_iam_role.example().arn())
            .build());

        var exampleAppImageConfig = new AppImageConfig("exampleAppImageConfig", AppImageConfigArgs.builder()        
            .appImageConfigName("example")
            .kernelGatewayImageConfig(AppImageConfigKernelGatewayImageConfigArgs.builder()
                .kernelSpec(AppImageConfigKernelGatewayImageConfigKernelSpecArgs.builder()
                    .name("example")
                    .build())
                .build())
            .build());

        var exampleImageVersion = new ImageVersion("exampleImageVersion", ImageVersionArgs.builder()        
            .imageName(exampleImage.id())
            .baseImage("base-image")
            .build());

        var exampleDomain = new Domain("exampleDomain", DomainArgs.builder()        
            .domainName("example")
            .authMode("IAM")
            .vpcId(aws_vpc.example().id())
            .subnetIds(aws_subnet.example().id())
            .defaultUserSettings(DomainDefaultUserSettingsArgs.builder()
                .executionRole(aws_iam_role.example().arn())
                .kernelGatewayAppSettings(DomainDefaultUserSettingsKernelGatewayAppSettingsArgs.builder()
                    .customImages(DomainDefaultUserSettingsKernelGatewayAppSettingsCustomImageArgs.builder()
                        .appImageConfigName(exampleAppImageConfig.appImageConfigName())
                        .imageName(exampleImageVersion.imageName())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleImage:
    type: aws:sagemaker:Image
    properties:
      imageName: example
      roleArn: ${aws_iam_role.example.arn}
  exampleAppImageConfig:
    type: aws:sagemaker:AppImageConfig
    properties:
      appImageConfigName: example
      kernelGatewayImageConfig:
        kernelSpec:
          name: example
  exampleImageVersion:
    type: aws:sagemaker:ImageVersion
    properties:
      imageName: ${exampleImage.id}
      baseImage: base-image
  exampleDomain:
    type: aws:sagemaker:Domain
    properties:
      domainName: example
      authMode: IAM
      vpcId: ${aws_vpc.example.id}
      subnetIds:
        - ${aws_subnet.example.id}
      defaultUserSettings:
        executionRole: ${aws_iam_role.example.arn}
        kernelGatewayAppSettings:
          customImages:
            - appImageConfigName: ${exampleAppImageConfig.appImageConfigName}
              imageName: ${exampleImageVersion.imageName}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SageMaker Domains using the `id`. For example:

```sh
 $ pulumi import aws:sagemaker/domain:Domain test_domain d-8jgsjtilstu8
```
 