Connects a custom domain name registered via `aws.apigateway.DomainName`
with a deployed API so that its methods can be called via the
custom domain name.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const exampleStage = new aws.apigateway.Stage("exampleStage", {
    deployment: aws_api_gateway_deployment.example.id,
    restApi: aws_api_gateway_rest_api.example.id,
    stageName: "example",
});
const exampleDomainName = new aws.apigateway.DomainName("exampleDomainName", {
    domainName: "example.com",
    certificateName: "example-api",
    certificateBody: fs.readFileSync(`${path.module}/example.com/example.crt`, "utf8"),
    certificateChain: fs.readFileSync(`${path.module}/example.com/ca.crt`, "utf8"),
    certificatePrivateKey: fs.readFileSync(`${path.module}/example.com/example.key`, "utf8"),
});
const exampleBasePathMapping = new aws.apigateway.BasePathMapping("exampleBasePathMapping", {
    restApi: aws_api_gateway_rest_api.example.id,
    stageName: exampleStage.stageName,
    domainName: exampleDomainName.domainName,
});
```
```python
import pulumi
import pulumi_aws as aws

example_stage = aws.apigateway.Stage("exampleStage",
    deployment=aws_api_gateway_deployment["example"]["id"],
    rest_api=aws_api_gateway_rest_api["example"]["id"],
    stage_name="example")
example_domain_name = aws.apigateway.DomainName("exampleDomainName",
    domain_name="example.com",
    certificate_name="example-api",
    certificate_body=(lambda path: open(path).read())(f"{path['module']}/example.com/example.crt"),
    certificate_chain=(lambda path: open(path).read())(f"{path['module']}/example.com/ca.crt"),
    certificate_private_key=(lambda path: open(path).read())(f"{path['module']}/example.com/example.key"))
example_base_path_mapping = aws.apigateway.BasePathMapping("exampleBasePathMapping",
    rest_api=aws_api_gateway_rest_api["example"]["id"],
    stage_name=example_stage.stage_name,
    domain_name=example_domain_name.domain_name)
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleStage = new Aws.ApiGateway.Stage("exampleStage", new()
    {
        Deployment = aws_api_gateway_deployment.Example.Id,
        RestApi = aws_api_gateway_rest_api.Example.Id,
        StageName = "example",
    });

    var exampleDomainName = new Aws.ApiGateway.DomainName("exampleDomainName", new()
    {
        Domain = "example.com",
        CertificateName = "example-api",
        CertificateBody = File.ReadAllText($"{path.Module}/example.com/example.crt"),
        CertificateChain = File.ReadAllText($"{path.Module}/example.com/ca.crt"),
        CertificatePrivateKey = File.ReadAllText($"{path.Module}/example.com/example.key"),
    });

    var exampleBasePathMapping = new Aws.ApiGateway.BasePathMapping("exampleBasePathMapping", new()
    {
        RestApi = aws_api_gateway_rest_api.Example.Id,
        StageName = exampleStage.StageName,
        DomainName = exampleDomainName.Domain,
    });

});
```
```go
package main

import (
	"fmt"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
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
		exampleStage, err := apigateway.NewStage(ctx, "exampleStage", &apigateway.StageArgs{
			Deployment: pulumi.Any(aws_api_gateway_deployment.Example.Id),
			RestApi:    pulumi.Any(aws_api_gateway_rest_api.Example.Id),
			StageName:  pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		exampleDomainName, err := apigateway.NewDomainName(ctx, "exampleDomainName", &apigateway.DomainNameArgs{
			DomainName:            pulumi.String("example.com"),
			CertificateName:       pulumi.String("example-api"),
			CertificateBody:       readFileOrPanic(fmt.Sprintf("%v/example.com/example.crt", path.Module)),
			CertificateChain:      readFileOrPanic(fmt.Sprintf("%v/example.com/ca.crt", path.Module)),
			CertificatePrivateKey: readFileOrPanic(fmt.Sprintf("%v/example.com/example.key", path.Module)),
		})
		if err != nil {
			return err
		}
		_, err = apigateway.NewBasePathMapping(ctx, "exampleBasePathMapping", &apigateway.BasePathMappingArgs{
			RestApi:    pulumi.Any(aws_api_gateway_rest_api.Example.Id),
			StageName:  exampleStage.StageName,
			DomainName: exampleDomainName.DomainName,
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
import com.pulumi.aws.apigateway.Stage;
import com.pulumi.aws.apigateway.StageArgs;
import com.pulumi.aws.apigateway.DomainName;
import com.pulumi.aws.apigateway.DomainNameArgs;
import com.pulumi.aws.apigateway.BasePathMapping;
import com.pulumi.aws.apigateway.BasePathMappingArgs;
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
        var exampleStage = new Stage("exampleStage", StageArgs.builder()        
            .deployment(aws_api_gateway_deployment.example().id())
            .restApi(aws_api_gateway_rest_api.example().id())
            .stageName("example")
            .build());

        var exampleDomainName = new DomainName("exampleDomainName", DomainNameArgs.builder()        
            .domainName("example.com")
            .certificateName("example-api")
            .certificateBody(Files.readString(Paths.get(String.format("%s/example.com/example.crt", path.module()))))
            .certificateChain(Files.readString(Paths.get(String.format("%s/example.com/ca.crt", path.module()))))
            .certificatePrivateKey(Files.readString(Paths.get(String.format("%s/example.com/example.key", path.module()))))
            .build());

        var exampleBasePathMapping = new BasePathMapping("exampleBasePathMapping", BasePathMappingArgs.builder()        
            .restApi(aws_api_gateway_rest_api.example().id())
            .stageName(exampleStage.stageName())
            .domainName(exampleDomainName.domainName())
            .build());

    }
}
```
```yaml
resources:
  exampleStage:
    type: aws:apigateway:Stage
    properties:
      deployment: ${aws_api_gateway_deployment.example.id}
      restApi: ${aws_api_gateway_rest_api.example.id}
      stageName: example
  exampleDomainName:
    type: aws:apigateway:DomainName
    properties:
      domainName: example.com
      certificateName: example-api
      certificateBody:
        fn::readFile: ${path.module}/example.com/example.crt
      certificateChain:
        fn::readFile: ${path.module}/example.com/ca.crt
      certificatePrivateKey:
        fn::readFile: ${path.module}/example.com/example.key
  exampleBasePathMapping:
    type: aws:apigateway:BasePathMapping
    properties:
      restApi: ${aws_api_gateway_rest_api.example.id}
      stageName: ${exampleStage.stageName}
      domainName: ${exampleDomainName.domainName}
```
{{% /example %}}
{{% /examples %}}

## Import

For a non-root `base_path`:

Using `pulumi import`, import `aws_api_gateway_base_path_mapping` using the domain name and base path. For example:

For an empty `base_path` or, in other words, a root path (`/`):

```sh
 $ pulumi import aws:apigateway/basePathMapping:BasePathMapping example example.com/
```
 For a non-root `base_path`:

```sh
 $ pulumi import aws:apigateway/basePathMapping:BasePathMapping example example.com/base-path
```
 