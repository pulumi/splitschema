Creates a domain resource for the specified domain (e.g., example.com).
You cannot register a new domain name using Lightsail. You must register
a domain name using Amazon Route 53 or another domain name registrar.
If you have already registered your domain, you can enter its name in
this parameter to manage the DNS records for that domain.

> **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail) for more details

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const domainTest = new aws.lightsail.Domain("domainTest", {domainName: "mydomain.com"});
```
```python
import pulumi
import pulumi_aws as aws

domain_test = aws.lightsail.Domain("domainTest", domain_name="mydomain.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var domainTest = new Aws.LightSail.Domain("domainTest", new()
    {
        DomainName = "mydomain.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewDomain(ctx, "domainTest", &lightsail.DomainArgs{
			DomainName: pulumi.String("mydomain.com"),
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
import com.pulumi.aws.lightsail.Domain;
import com.pulumi.aws.lightsail.DomainArgs;
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
        var domainTest = new Domain("domainTest", DomainArgs.builder()        
            .domainName("mydomain.com")
            .build());

    }
}
```
```yaml
resources:
  domainTest:
    type: aws:lightsail:Domain
    properties:
      domainName: mydomain.com
```
{{% /example %}}
{{% /examples %}}