Data source for managing an AWS SSO Admin Application.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ssoadmin.getApplication({
    applicationArn: "arn:aws:sso::012345678901:application/ssoins-1234/apl-5678",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssoadmin.get_application(application_arn="arn:aws:sso::012345678901:application/ssoins-1234/apl-5678")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.SsoAdmin.GetApplication.Invoke(new()
    {
        ApplicationArn = "arn:aws:sso::012345678901:application/ssoins-1234/apl-5678",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssoadmin"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssoadmin.LookupApplication(ctx, &ssoadmin.LookupApplicationArgs{
			ApplicationArn: "arn:aws:sso::012345678901:application/ssoins-1234/apl-5678",
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
import com.pulumi.aws.ssoadmin.SsoadminFunctions;
import com.pulumi.aws.ssoadmin.inputs.GetApplicationArgs;
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
        final var example = SsoadminFunctions.getApplication(GetApplicationArgs.builder()
            .applicationArn("arn:aws:sso::012345678901:application/ssoins-1234/apl-5678")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ssoadmin:getApplication
      Arguments:
        applicationArn: arn:aws:sso::012345678901:application/ssoins-1234/apl-5678
```
{{% /example %}}
{{% /examples %}}