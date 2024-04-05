Retrieve the SES domain identity

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ses.getDomainIdentity({
    domain: "example.com",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ses.get_domain_identity(domain="example.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ses.GetDomainIdentity.Invoke(new()
    {
        Domain = "example.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.LookupDomainIdentity(ctx, &ses.LookupDomainIdentityArgs{
			Domain: "example.com",
		}, nil)
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
import com.pulumi.aws.ses.SesFunctions;
import com.pulumi.aws.ses.inputs.GetDomainIdentityArgs;
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
        final var example = SesFunctions.getDomainIdentity(GetDomainIdentityArgs.builder()
            .domain("example.com")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ses:getDomainIdentity
      Arguments:
        domain: example.com
```
{{% /example %}}
{{% /examples %}}