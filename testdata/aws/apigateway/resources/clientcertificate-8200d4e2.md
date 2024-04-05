Provides an API Gateway Client Certificate.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const demo = new aws.apigateway.ClientCertificate("demo", {description: "My client certificate"});
```
```python
import pulumi
import pulumi_aws as aws

demo = aws.apigateway.ClientCertificate("demo", description="My client certificate")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var demo = new Aws.ApiGateway.ClientCertificate("demo", new()
    {
        Description = "My client certificate",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apigateway.NewClientCertificate(ctx, "demo", &apigateway.ClientCertificateArgs{
			Description: pulumi.String("My client certificate"),
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
import com.pulumi.aws.apigateway.ClientCertificate;
import com.pulumi.aws.apigateway.ClientCertificateArgs;
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
        var demo = new ClientCertificate("demo", ClientCertificateArgs.builder()        
            .description("My client certificate")
            .build());

    }
}
```
```yaml
resources:
  demo:
    type: aws:apigateway:ClientCertificate
    properties:
      description: My client certificate
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import API Gateway Client Certificates using the id. For example:

```sh
 $ pulumi import aws:apigateway/clientCertificate:ClientCertificate demo ab1cqe
```
 