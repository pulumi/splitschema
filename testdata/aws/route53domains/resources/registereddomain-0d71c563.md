Provides a resource to manage a domain that has been [registered](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/registrar-tld-list.html) and associated with the current AWS account.

**This is an advanced resource** and has special caveats to be aware of when using it. Please read this document in its entirety before using this resource.

The `aws.route53domains.RegisteredDomain` resource behaves differently from normal resources in that if a domain has been registered, the provider does not _register_ this domain, but instead "adopts" it into management. A destroy does not delete the domain but does remove the resource from state.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53domains.RegisteredDomain("example", {
    domainName: "example.com",
    nameServers: [
        {
            name: "ns-195.awsdns-24.com",
        },
        {
            name: "ns-874.awsdns-45.net",
        },
    ],
    tags: {
        Environment: "test",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53domains.RegisteredDomain("example",
    domain_name="example.com",
    name_servers=[
        aws.route53domains.RegisteredDomainNameServerArgs(
            name="ns-195.awsdns-24.com",
        ),
        aws.route53domains.RegisteredDomainNameServerArgs(
            name="ns-874.awsdns-45.net",
        ),
    ],
    tags={
        "Environment": "test",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53Domains.RegisteredDomain("example", new()
    {
        DomainName = "example.com",
        NameServers = new[]
        {
            new Aws.Route53Domains.Inputs.RegisteredDomainNameServerArgs
            {
                Name = "ns-195.awsdns-24.com",
            },
            new Aws.Route53Domains.Inputs.RegisteredDomainNameServerArgs
            {
                Name = "ns-874.awsdns-45.net",
            },
        },
        Tags = 
        {
            { "Environment", "test" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53domains"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53domains.NewRegisteredDomain(ctx, "example", &route53domains.RegisteredDomainArgs{
			DomainName: pulumi.String("example.com"),
			NameServers: route53domains.RegisteredDomainNameServerArray{
				&route53domains.RegisteredDomainNameServerArgs{
					Name: pulumi.String("ns-195.awsdns-24.com"),
				},
				&route53domains.RegisteredDomainNameServerArgs{
					Name: pulumi.String("ns-874.awsdns-45.net"),
				},
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("test"),
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
import com.pulumi.aws.route53domains.RegisteredDomain;
import com.pulumi.aws.route53domains.RegisteredDomainArgs;
import com.pulumi.aws.route53domains.inputs.RegisteredDomainNameServerArgs;
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
        var example = new RegisteredDomain("example", RegisteredDomainArgs.builder()        
            .domainName("example.com")
            .nameServers(            
                RegisteredDomainNameServerArgs.builder()
                    .name("ns-195.awsdns-24.com")
                    .build(),
                RegisteredDomainNameServerArgs.builder()
                    .name("ns-874.awsdns-45.net")
                    .build())
            .tags(Map.of("Environment", "test"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53domains:RegisteredDomain
    properties:
      domainName: example.com
      nameServers:
        - name: ns-195.awsdns-24.com
        - name: ns-874.awsdns-45.net
      tags:
        Environment: test
```
{{% /example %}}
{{% /examples %}}