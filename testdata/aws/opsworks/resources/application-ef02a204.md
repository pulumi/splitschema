Provides an OpsWorks application resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const foo_app = new aws.opsworks.Application("foo-app", {
    shortName: "foobar",
    stackId: aws_opsworks_stack.main.id,
    type: "rails",
    description: "This is a Rails application",
    domains: [
        "example.com",
        "sub.example.com",
    ],
    environments: [{
        key: "key",
        value: "value",
        secure: false,
    }],
    appSources: [{
        type: "git",
        revision: "master",
        url: "https://github.com/example.git",
    }],
    enableSsl: true,
    sslConfigurations: [{
        privateKey: fs.readFileSync("./foobar.key", "utf8"),
        certificate: fs.readFileSync("./foobar.crt", "utf8"),
    }],
    documentRoot: "public",
    autoBundleOnDeploy: "true",
    railsEnv: "staging",
});
```
```python
import pulumi
import pulumi_aws as aws

foo_app = aws.opsworks.Application("foo-app",
    short_name="foobar",
    stack_id=aws_opsworks_stack["main"]["id"],
    type="rails",
    description="This is a Rails application",
    domains=[
        "example.com",
        "sub.example.com",
    ],
    environments=[aws.opsworks.ApplicationEnvironmentArgs(
        key="key",
        value="value",
        secure=False,
    )],
    app_sources=[aws.opsworks.ApplicationAppSourceArgs(
        type="git",
        revision="master",
        url="https://github.com/example.git",
    )],
    enable_ssl=True,
    ssl_configurations=[aws.opsworks.ApplicationSslConfigurationArgs(
        private_key=(lambda path: open(path).read())("./foobar.key"),
        certificate=(lambda path: open(path).read())("./foobar.crt"),
    )],
    document_root="public",
    auto_bundle_on_deploy="true",
    rails_env="staging")
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo_app = new Aws.OpsWorks.Application("foo-app", new()
    {
        ShortName = "foobar",
        StackId = aws_opsworks_stack.Main.Id,
        Type = "rails",
        Description = "This is a Rails application",
        Domains = new[]
        {
            "example.com",
            "sub.example.com",
        },
        Environments = new[]
        {
            new Aws.OpsWorks.Inputs.ApplicationEnvironmentArgs
            {
                Key = "key",
                Value = "value",
                Secure = false,
            },
        },
        AppSources = new[]
        {
            new Aws.OpsWorks.Inputs.ApplicationAppSourceArgs
            {
                Type = "git",
                Revision = "master",
                Url = "https://github.com/example.git",
            },
        },
        EnableSsl = true,
        SslConfigurations = new[]
        {
            new Aws.OpsWorks.Inputs.ApplicationSslConfigurationArgs
            {
                PrivateKey = File.ReadAllText("./foobar.key"),
                Certificate = File.ReadAllText("./foobar.crt"),
            },
        },
        DocumentRoot = "public",
        AutoBundleOnDeploy = "true",
        RailsEnv = "staging",
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewApplication(ctx, "foo-app", &opsworks.ApplicationArgs{
			ShortName:   pulumi.String("foobar"),
			StackId:     pulumi.Any(aws_opsworks_stack.Main.Id),
			Type:        pulumi.String("rails"),
			Description: pulumi.String("This is a Rails application"),
			Domains: pulumi.StringArray{
				pulumi.String("example.com"),
				pulumi.String("sub.example.com"),
			},
			Environments: opsworks.ApplicationEnvironmentArray{
				&opsworks.ApplicationEnvironmentArgs{
					Key:    pulumi.String("key"),
					Value:  pulumi.String("value"),
					Secure: pulumi.Bool(false),
				},
			},
			AppSources: opsworks.ApplicationAppSourceArray{
				&opsworks.ApplicationAppSourceArgs{
					Type:     pulumi.String("git"),
					Revision: pulumi.String("master"),
					Url:      pulumi.String("https://github.com/example.git"),
				},
			},
			EnableSsl: pulumi.Bool(true),
			SslConfigurations: opsworks.ApplicationSslConfigurationArray{
				&opsworks.ApplicationSslConfigurationArgs{
					PrivateKey:  readFileOrPanic("./foobar.key"),
					Certificate: readFileOrPanic("./foobar.crt"),
				},
			},
			DocumentRoot:       pulumi.String("public"),
			AutoBundleOnDeploy: pulumi.String("true"),
			RailsEnv:           pulumi.String("staging"),
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
import com.pulumi.aws.opsworks.Application;
import com.pulumi.aws.opsworks.ApplicationArgs;
import com.pulumi.aws.opsworks.inputs.ApplicationEnvironmentArgs;
import com.pulumi.aws.opsworks.inputs.ApplicationAppSourceArgs;
import com.pulumi.aws.opsworks.inputs.ApplicationSslConfigurationArgs;
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
        var foo_app = new Application("foo-app", ApplicationArgs.builder()        
            .shortName("foobar")
            .stackId(aws_opsworks_stack.main().id())
            .type("rails")
            .description("This is a Rails application")
            .domains(            
                "example.com",
                "sub.example.com")
            .environments(ApplicationEnvironmentArgs.builder()
                .key("key")
                .value("value")
                .secure(false)
                .build())
            .appSources(ApplicationAppSourceArgs.builder()
                .type("git")
                .revision("master")
                .url("https://github.com/example.git")
                .build())
            .enableSsl(true)
            .sslConfigurations(ApplicationSslConfigurationArgs.builder()
                .privateKey(Files.readString(Paths.get("./foobar.key")))
                .certificate(Files.readString(Paths.get("./foobar.crt")))
                .build())
            .documentRoot("public")
            .autoBundleOnDeploy(true)
            .railsEnv("staging")
            .build());

    }
}
```
```yaml
resources:
  foo-app:
    type: aws:opsworks:Application
    properties:
      shortName: foobar
      stackId: ${aws_opsworks_stack.main.id}
      type: rails
      description: This is a Rails application
      domains:
        - example.com
        - sub.example.com
      environments:
        - key: key
          value: value
          secure: false
      appSources:
        - type: git
          revision: master
          url: https://github.com/example.git
      enableSsl: true
      sslConfigurations:
        - privateKey:
            fn::readFile: ./foobar.key
          certificate:
            fn::readFile: ./foobar.crt
      documentRoot: public
      autoBundleOnDeploy: true
      railsEnv: staging
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Opsworks Application using the `id`. For example:

```sh
 $ pulumi import aws:opsworks/application:Application test <id>
```
 