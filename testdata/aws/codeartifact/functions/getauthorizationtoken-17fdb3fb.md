The CodeArtifact Authorization Token data source generates a temporary authentication token for accessing repositories in a CodeArtifact domain.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.codeartifact.getAuthorizationToken({
    domain: aws_codeartifact_domain.test.domain,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.codeartifact.get_authorization_token(domain=aws_codeartifact_domain["test"]["domain"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.CodeArtifact.GetAuthorizationToken.Invoke(new()
    {
        Domain = aws_codeartifact_domain.Test.Domain,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codeartifact"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codeartifact.GetAuthorizationToken(ctx, &codeartifact.GetAuthorizationTokenArgs{
			Domain: aws_codeartifact_domain.Test.Domain,
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
import com.pulumi.aws.codeartifact.CodeartifactFunctions;
import com.pulumi.aws.codeartifact.inputs.GetAuthorizationTokenArgs;
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
        final var test = CodeartifactFunctions.getAuthorizationToken(GetAuthorizationTokenArgs.builder()
            .domain(aws_codeartifact_domain.test().domain())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:codeartifact:getAuthorizationToken
      Arguments:
        domain: ${aws_codeartifact_domain.test.domain}
```
{{% /example %}}
{{% /examples %}}