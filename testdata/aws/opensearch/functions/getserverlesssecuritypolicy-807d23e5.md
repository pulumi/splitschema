Use this data source to get information about an AWS OpenSearch Serverless Security Policy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.opensearch.getServerlessSecurityPolicy({
    name: "example-security-policy",
    type: "encryption",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.opensearch.get_serverless_security_policy(name="example-security-policy",
    type="encryption")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.OpenSearch.GetServerlessSecurityPolicy.Invoke(new()
    {
        Name = "example-security-policy",
        Type = "encryption",
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
		_, err := opensearch.LookupServerlessSecurityPolicy(ctx, &opensearch.LookupServerlessSecurityPolicyArgs{
			Name: "example-security-policy",
			Type: "encryption",
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
import com.pulumi.aws.opensearch.inputs.GetServerlessSecurityPolicyArgs;
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
        final var example = OpensearchFunctions.getServerlessSecurityPolicy(GetServerlessSecurityPolicyArgs.builder()
            .name("example-security-policy")
            .type("encryption")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:opensearch:getServerlessSecurityPolicy
      Arguments:
        name: example-security-policy
        type: encryption
```
{{% /example %}}
{{% /examples %}}