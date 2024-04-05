Manages an App Runner Custom Domain association.

> **NOTE:** After creation, you must use the information in the `certification_validation_records` attribute to add CNAME records to your Domain Name System (DNS). For each mapped domain name, add a mapping to the target App Runner subdomain (found in the `dns_target` attribute) and one or more certificate validation records. App Runner then performs DNS validation to verify that you own or control the domain name you associated. App Runner tracks domain validity in a certificate stored in AWS Certificate Manager (ACM).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.apprunner.CustomDomainAssociation("example", {
    domainName: "example.com",
    serviceArn: aws_apprunner_service.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.apprunner.CustomDomainAssociation("example",
    domain_name="example.com",
    service_arn=aws_apprunner_service["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppRunner.CustomDomainAssociation("example", new()
    {
        DomainName = "example.com",
        ServiceArn = aws_apprunner_service.Example.Arn,
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
		_, err := apprunner.NewCustomDomainAssociation(ctx, "example", &apprunner.CustomDomainAssociationArgs{
			DomainName: pulumi.String("example.com"),
			ServiceArn: pulumi.Any(aws_apprunner_service.Example.Arn),
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
import com.pulumi.aws.apprunner.CustomDomainAssociation;
import com.pulumi.aws.apprunner.CustomDomainAssociationArgs;
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
        var example = new CustomDomainAssociation("example", CustomDomainAssociationArgs.builder()        
            .domainName("example.com")
            .serviceArn(aws_apprunner_service.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:apprunner:CustomDomainAssociation
    properties:
      domainName: example.com
      serviceArn: ${aws_apprunner_service.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import App Runner Custom Domain Associations using the `domain_name` and `service_arn` separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:apprunner/customDomainAssociation:CustomDomainAssociation example example.com,arn:aws:apprunner:us-east-1:123456789012:service/example-app/8fe1e10304f84fd2b0df550fe98a71fa
```
 