Provides a DMS (Data Migration Service) certificate resource. DMS certificates can be created, deleted, and imported.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new certificate
const test = new aws.dms.Certificate("test", {
    certificateId: "test-dms-certificate-tf",
    certificatePem: "...",
    tags: {
        Name: "test",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

# Create a new certificate
test = aws.dms.Certificate("test",
    certificate_id="test-dms-certificate-tf",
    certificate_pem="...",
    tags={
        "Name": "test",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Create a new certificate
    var test = new Aws.Dms.Certificate("test", new()
    {
        CertificateId = "test-dms-certificate-tf",
        CertificatePem = "...",
        Tags = 
        {
            { "Name", "test" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.NewCertificate(ctx, "test", &dms.CertificateArgs{
			CertificateId:  pulumi.String("test-dms-certificate-tf"),
			CertificatePem: pulumi.String("..."),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("test"),
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
import com.pulumi.aws.dms.Certificate;
import com.pulumi.aws.dms.CertificateArgs;
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
        var test = new Certificate("test", CertificateArgs.builder()        
            .certificateId("test-dms-certificate-tf")
            .certificatePem("...")
            .tags(Map.of("Name", "test"))
            .build());

    }
}
```
```yaml
resources:
  # Create a new certificate
  test:
    type: aws:dms:Certificate
    properties:
      certificateId: test-dms-certificate-tf
      certificatePem: '...'
      tags:
        Name: test
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import certificates using the `certificate_id`. For example:

```sh
 $ pulumi import aws:dms/certificate:Certificate test test-dms-certificate-tf
```
 