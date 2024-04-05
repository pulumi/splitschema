Use this data source to get information about an OpenSearch Domain

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDomain = aws.opensearch.getDomain({
    domainName: "my-domain-name",
});
```
```python
import pulumi
import pulumi_aws as aws

my_domain = aws.opensearch.get_domain(domain_name="my-domain-name")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myDomain = Aws.OpenSearch.GetDomain.Invoke(new()
    {
        DomainName = "my-domain-name",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opensearch.LookupDomain(ctx, &opensearch.LookupDomainArgs{
			DomainName: "my-domain-name",
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
import com.pulumi.aws.opensearch.OpensearchFunctions;
import com.pulumi.aws.opensearch.inputs.GetDomainArgs;
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
        final var myDomain = OpensearchFunctions.getDomain(GetDomainArgs.builder()
            .domainName("my-domain-name")
            .build());

    }
}
```
```yaml
variables:
  myDomain:
    fn::invoke:
      Function: aws:opensearch:getDomain
      Arguments:
        domainName: my-domain-name
```
{{% /example %}}
{{% /examples %}}