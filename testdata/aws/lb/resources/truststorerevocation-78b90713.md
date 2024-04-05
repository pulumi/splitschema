Provides a ELBv2 Trust Store Revocation for use with Application Load Balancer Listener resources.

{{% examples %}}
## Example Usage
{{% example %}}
### Trust Store With Revocations

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testTrustStore = new aws.lb.TrustStore("testTrustStore", {
    caCertificatesBundleS3Bucket: "...",
    caCertificatesBundleS3Key: "...",
});
const testTrustStoreRevocation = new aws.lb.TrustStoreRevocation("testTrustStoreRevocation", {
    trustStoreArn: testTrustStore.arn,
    revocationsS3Bucket: "...",
    revocationsS3Key: "...",
});
```
```python
import pulumi
import pulumi_aws as aws

test_trust_store = aws.lb.TrustStore("testTrustStore",
    ca_certificates_bundle_s3_bucket="...",
    ca_certificates_bundle_s3_key="...")
test_trust_store_revocation = aws.lb.TrustStoreRevocation("testTrustStoreRevocation",
    trust_store_arn=test_trust_store.arn,
    revocations_s3_bucket="...",
    revocations_s3_key="...")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testTrustStore = new Aws.LB.TrustStore("testTrustStore", new()
    {
        CaCertificatesBundleS3Bucket = "...",
        CaCertificatesBundleS3Key = "...",
    });

    var testTrustStoreRevocation = new Aws.LB.TrustStoreRevocation("testTrustStoreRevocation", new()
    {
        TrustStoreArn = testTrustStore.Arn,
        RevocationsS3Bucket = "...",
        RevocationsS3Key = "...",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testTrustStore, err := lb.NewTrustStore(ctx, "testTrustStore", &lb.TrustStoreArgs{
			CaCertificatesBundleS3Bucket: pulumi.String("..."),
			CaCertificatesBundleS3Key:    pulumi.String("..."),
		})
		if err != nil {
			return err
		}
		_, err = lb.NewTrustStoreRevocation(ctx, "testTrustStoreRevocation", &lb.TrustStoreRevocationArgs{
			TrustStoreArn:       testTrustStore.Arn,
			RevocationsS3Bucket: pulumi.String("..."),
			RevocationsS3Key:    pulumi.String("..."),
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
import com.pulumi.aws.lb.TrustStore;
import com.pulumi.aws.lb.TrustStoreArgs;
import com.pulumi.aws.lb.TrustStoreRevocation;
import com.pulumi.aws.lb.TrustStoreRevocationArgs;
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
        var testTrustStore = new TrustStore("testTrustStore", TrustStoreArgs.builder()        
            .caCertificatesBundleS3Bucket("...")
            .caCertificatesBundleS3Key("...")
            .build());

        var testTrustStoreRevocation = new TrustStoreRevocation("testTrustStoreRevocation", TrustStoreRevocationArgs.builder()        
            .trustStoreArn(testTrustStore.arn())
            .revocationsS3Bucket("...")
            .revocationsS3Key("...")
            .build());

    }
}
```
```yaml
resources:
  testTrustStore:
    type: aws:lb:TrustStore
    properties:
      caCertificatesBundleS3Bucket: '...'
      caCertificatesBundleS3Key: '...'
  testTrustStoreRevocation:
    type: aws:lb:TrustStoreRevocation
    properties:
      trustStoreArn: ${testTrustStore.arn}
      revocationsS3Bucket: '...'
      revocationsS3Key: '...'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Trust Store Revocations using their ARN. For example:

```sh
 $ pulumi import aws:lb/trustStoreRevocation:TrustStoreRevocation example arn:aws:elasticloadbalancing:us-west-2:187416307283:truststore/my-trust-store/20cfe21448b66314,6
```
 