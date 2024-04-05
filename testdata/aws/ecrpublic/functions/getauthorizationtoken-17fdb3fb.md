The Public ECR Authorization Token data source allows the authorization token, token expiration date, user name, and password to be retrieved for a Public ECR repository.

> **NOTE:** This data source can only be used in the `us-east-1` region.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const token = aws.ecrpublic.getAuthorizationToken({});
```
```python
import pulumi
import pulumi_aws as aws

token = aws.ecrpublic.get_authorization_token()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var token = Aws.EcrPublic.GetAuthorizationToken.Invoke();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecrpublic"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecrpublic.GetAuthorizationToken(ctx, nil, nil)
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
import com.pulumi.aws.ecrpublic.EcrpublicFunctions;
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
        final var token = EcrpublicFunctions.getAuthorizationToken();

    }
}
```
```yaml
variables:
  token:
    fn::invoke:
      Function: aws:ecrpublic:getAuthorizationToken
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}