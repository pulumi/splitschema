Data source for managing an AWS DMS (Database Migration) Certificate.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.dms.getCertificate({
    certificateId: aws_dms_certificate.test.certificate_id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.dms.get_certificate(certificate_id=aws_dms_certificate["test"]["certificate_id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Dms.GetCertificate.Invoke(new()
    {
        CertificateId = aws_dms_certificate.Test.Certificate_id,
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
		_, err := dms.LookupCertificate(ctx, &dms.LookupCertificateArgs{
			CertificateId: aws_dms_certificate.Test.Certificate_id,
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
import com.pulumi.aws.dms.DmsFunctions;
import com.pulumi.aws.dms.inputs.GetCertificateArgs;
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
        final var example = DmsFunctions.getCertificate(GetCertificateArgs.builder()
            .certificateId(aws_dms_certificate.test().certificate_id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:dms:getCertificate
      Arguments:
        certificateId: ${aws_dms_certificate.test.certificate_id}
```
{{% /example %}}
{{% /examples %}}