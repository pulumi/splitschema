Creates a domain entry resource

> **NOTE on `id`:** In an effort to simplify imports, this resource `id` field has been updated to the standard resource id separator, a comma (`,`). For backward compatibility, the previous separator (underscore `_`) can still be used to read and import existing resources. When state is refreshed, the `id` will be updated to use the new standard separator. The previous separator will be deprecated in a future major release.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testDomain = new aws.lightsail.Domain("testDomain", {domainName: "mydomain.com"});
const testDomainEntry = new aws.lightsail.DomainEntry("testDomainEntry", {
    domainName: aws_lightsail_domain.domain_test.domain_name,
    type: "A",
    target: "127.0.0.1",
});
```
```python
import pulumi
import pulumi_aws as aws

test_domain = aws.lightsail.Domain("testDomain", domain_name="mydomain.com")
test_domain_entry = aws.lightsail.DomainEntry("testDomainEntry",
    domain_name=aws_lightsail_domain["domain_test"]["domain_name"],
    type="A",
    target="127.0.0.1")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testDomain = new Aws.LightSail.Domain("testDomain", new()
    {
        DomainName = "mydomain.com",
    });

    var testDomainEntry = new Aws.LightSail.DomainEntry("testDomainEntry", new()
    {
        DomainName = aws_lightsail_domain.Domain_test.Domain_name,
        Type = "A",
        Target = "127.0.0.1",
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
		_, err := lightsail.NewDomain(ctx, "testDomain", &lightsail.DomainArgs{
			DomainName: pulumi.String("mydomain.com"),
		})
		if err != nil {
			return err
		}
		_, err = lightsail.NewDomainEntry(ctx, "testDomainEntry", &lightsail.DomainEntryArgs{
			DomainName: pulumi.Any(aws_lightsail_domain.Domain_test.Domain_name),
			Type:       pulumi.String("A"),
			Target:     pulumi.String("127.0.0.1"),
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
import com.pulumi.aws.lightsail.DomainEntry;
import com.pulumi.aws.lightsail.DomainEntryArgs;
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
        var testDomain = new Domain("testDomain", DomainArgs.builder()        
            .domainName("mydomain.com")
            .build());

        var testDomainEntry = new DomainEntry("testDomainEntry", DomainEntryArgs.builder()        
            .domainName(aws_lightsail_domain.domain_test().domain_name())
            .type("A")
            .target("127.0.0.1")
            .build());

    }
}
```
```yaml
resources:
  testDomain:
    type: aws:lightsail:Domain
    properties:
      domainName: mydomain.com
  testDomainEntry:
    type: aws:lightsail:DomainEntry
    properties:
      domainName: ${aws_lightsail_domain.domain_test.domain_name}
      type: A
      target: 127.0.0.1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_lightsail_domain_entry` using the id attribute. For example:

```sh
 $ pulumi import aws:lightsail/domainEntry:DomainEntry example www,mydomain.com,A,127.0.0.1
```
 