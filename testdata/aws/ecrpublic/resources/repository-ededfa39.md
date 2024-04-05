Provides a Public Elastic Container Registry Repository.

> **NOTE:** This resource can only be used in the `us-east-1` region.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const usEast1 = new aws.Provider("usEast1", {region: "us-east-1"});
const foo = new aws.ecrpublic.Repository("foo", {
    repositoryName: "bar",
    catalogData: {
        aboutText: "About Text",
        architectures: ["ARM"],
        description: "Description",
        logoImageBlob: Buffer.from(fs.readFileSync(image.png, 'binary')).toString('base64'),
        operatingSystems: ["Linux"],
        usageText: "Usage Text",
    },
    tags: {
        env: "production",
    },
}, {
    provider: aws.us_east_1,
});
```
```python
import pulumi
import base64
import pulumi_aws as aws

us_east1 = aws.Provider("usEast1", region="us-east-1")
foo = aws.ecrpublic.Repository("foo",
    repository_name="bar",
    catalog_data=aws.ecrpublic.RepositoryCatalogDataArgs(
        about_text="About Text",
        architectures=["ARM"],
        description="Description",
        logo_image_blob=(lambda path: base64.b64encode(open(path).read().encode()).decode())(image["png"]),
        operating_systems=["Linux"],
        usage_text="Usage Text",
    ),
    tags={
        "env": "production",
    },
    opts=pulumi.ResourceOptions(provider=aws["us_east_1"]))
```
```csharp
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

	
string ReadFileBase64(string path) 
{
    return Convert.ToBase64String(Encoding.UTF8.GetBytes(File.ReadAllText(path)));
}

return await Deployment.RunAsync(() => 
{
    var usEast1 = new Aws.Provider("usEast1", new()
    {
        Region = "us-east-1",
    });

    var foo = new Aws.EcrPublic.Repository("foo", new()
    {
        RepositoryName = "bar",
        CatalogData = new Aws.EcrPublic.Inputs.RepositoryCatalogDataArgs
        {
            AboutText = "About Text",
            Architectures = new[]
            {
                "ARM",
            },
            Description = "Description",
            LogoImageBlob = ReadFileBase64(image.Png),
            OperatingSystems = new[]
            {
                "Linux",
            },
            UsageText = "Usage Text",
        },
        Tags = 
        {
            { "env", "production" },
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Us_east_1,
    });

});
```
```go
package main

import (
	"encoding/base64"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecrpublic"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func filebase64OrPanic(path string) string {
	if fileData, err := os.ReadFile(path); err == nil {
		return base64.StdEncoding.EncodeToString(fileData[:])
	} else {
		panic(err.Error())
	}
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "usEast1", &aws.ProviderArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		_, err = ecrpublic.NewRepository(ctx, "foo", &ecrpublic.RepositoryArgs{
			RepositoryName: pulumi.String("bar"),
			CatalogData: &ecrpublic.RepositoryCatalogDataArgs{
				AboutText: pulumi.String("About Text"),
				Architectures: pulumi.StringArray{
					pulumi.String("ARM"),
				},
				Description:   pulumi.String("Description"),
				LogoImageBlob: filebase64OrPanic(image.Png),
				OperatingSystems: pulumi.StringArray{
					pulumi.String("Linux"),
				},
				UsageText: pulumi.String("Usage Text"),
			},
			Tags: pulumi.StringMap{
				"env": pulumi.String("production"),
			},
		}, pulumi.Provider(aws.Us_east_1))
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
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.ecrpublic.Repository;
import com.pulumi.aws.ecrpublic.RepositoryArgs;
import com.pulumi.aws.ecrpublic.inputs.RepositoryCatalogDataArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var usEast1 = new Provider("usEast1", ProviderArgs.builder()        
            .region("us-east-1")
            .build());

        var foo = new Repository("foo", RepositoryArgs.builder()        
            .repositoryName("bar")
            .catalogData(RepositoryCatalogDataArgs.builder()
                .aboutText("About Text")
                .architectures("ARM")
                .description("Description")
                .logoImageBlob(Base64.getEncoder().encodeToString(Files.readAllBytes(Paths.get(image.png()))))
                .operatingSystems("Linux")
                .usageText("Usage Text")
                .build())
            .tags(Map.of("env", "production"))
            .build(), CustomResourceOptions.builder()
                .provider(aws.us_east_1())
                .build());

    }
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECR Public Repositories using the `repository_name`. For example:

```sh
 $ pulumi import aws:ecrpublic/repository:Repository example example
```
 