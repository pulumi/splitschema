Resource for managing a Verified Access Instance Trust Provider Attachment.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleInstance = new aws.verifiedaccess.Instance("exampleInstance", {});
const exampleTrustProvider = new aws.verifiedaccess.TrustProvider("exampleTrustProvider", {
    deviceTrustProviderType: "jamf",
    policyReferenceName: "example",
    trustProviderType: "device",
    deviceOptions: {
        tenantId: "example",
    },
});
const exampleInstanceTrustProviderAttachment = new aws.verifiedaccess.InstanceTrustProviderAttachment("exampleInstanceTrustProviderAttachment", {
    verifiedaccessInstanceId: exampleInstance.id,
    verifiedaccessTrustProviderId: exampleTrustProvider.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example_instance = aws.verifiedaccess.Instance("exampleInstance")
example_trust_provider = aws.verifiedaccess.TrustProvider("exampleTrustProvider",
    device_trust_provider_type="jamf",
    policy_reference_name="example",
    trust_provider_type="device",
    device_options=aws.verifiedaccess.TrustProviderDeviceOptionsArgs(
        tenant_id="example",
    ))
example_instance_trust_provider_attachment = aws.verifiedaccess.InstanceTrustProviderAttachment("exampleInstanceTrustProviderAttachment",
    verifiedaccess_instance_id=example_instance.id,
    verifiedaccess_trust_provider_id=example_trust_provider.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleInstance = new Aws.VerifiedAccess.Instance("exampleInstance");

    var exampleTrustProvider = new Aws.VerifiedAccess.TrustProvider("exampleTrustProvider", new()
    {
        DeviceTrustProviderType = "jamf",
        PolicyReferenceName = "example",
        TrustProviderType = "device",
        DeviceOptions = new Aws.VerifiedAccess.Inputs.TrustProviderDeviceOptionsArgs
        {
            TenantId = "example",
        },
    });

    var exampleInstanceTrustProviderAttachment = new Aws.VerifiedAccess.InstanceTrustProviderAttachment("exampleInstanceTrustProviderAttachment", new()
    {
        VerifiedaccessInstanceId = exampleInstance.Id,
        VerifiedaccessTrustProviderId = exampleTrustProvider.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/verifiedaccess"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleInstance, err := verifiedaccess.NewInstance(ctx, "exampleInstance", nil)
		if err != nil {
			return err
		}
		exampleTrustProvider, err := verifiedaccess.NewTrustProvider(ctx, "exampleTrustProvider", &verifiedaccess.TrustProviderArgs{
			DeviceTrustProviderType: pulumi.String("jamf"),
			PolicyReferenceName:     pulumi.String("example"),
			TrustProviderType:       pulumi.String("device"),
			DeviceOptions: &verifiedaccess.TrustProviderDeviceOptionsArgs{
				TenantId: pulumi.String("example"),
			},
		})
		if err != nil {
			return err
		}
		_, err = verifiedaccess.NewInstanceTrustProviderAttachment(ctx, "exampleInstanceTrustProviderAttachment", &verifiedaccess.InstanceTrustProviderAttachmentArgs{
			VerifiedaccessInstanceId:      exampleInstance.ID(),
			VerifiedaccessTrustProviderId: exampleTrustProvider.ID(),
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
import com.pulumi.aws.verifiedaccess.Instance;
import com.pulumi.aws.verifiedaccess.TrustProvider;
import com.pulumi.aws.verifiedaccess.TrustProviderArgs;
import com.pulumi.aws.verifiedaccess.inputs.TrustProviderDeviceOptionsArgs;
import com.pulumi.aws.verifiedaccess.InstanceTrustProviderAttachment;
import com.pulumi.aws.verifiedaccess.InstanceTrustProviderAttachmentArgs;
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
        var exampleInstance = new Instance("exampleInstance");

        var exampleTrustProvider = new TrustProvider("exampleTrustProvider", TrustProviderArgs.builder()        
            .deviceTrustProviderType("jamf")
            .policyReferenceName("example")
            .trustProviderType("device")
            .deviceOptions(TrustProviderDeviceOptionsArgs.builder()
                .tenantId("example")
                .build())
            .build());

        var exampleInstanceTrustProviderAttachment = new InstanceTrustProviderAttachment("exampleInstanceTrustProviderAttachment", InstanceTrustProviderAttachmentArgs.builder()        
            .verifiedaccessInstanceId(exampleInstance.id())
            .verifiedaccessTrustProviderId(exampleTrustProvider.id())
            .build());

    }
}
```
```yaml
resources:
  exampleInstance:
    type: aws:verifiedaccess:Instance
  exampleTrustProvider:
    type: aws:verifiedaccess:TrustProvider
    properties:
      deviceTrustProviderType: jamf
      policyReferenceName: example
      trustProviderType: device
      deviceOptions:
        tenantId: example
  exampleInstanceTrustProviderAttachment:
    type: aws:verifiedaccess:InstanceTrustProviderAttachment
    properties:
      verifiedaccessInstanceId: ${exampleInstance.id}
      verifiedaccessTrustProviderId: ${exampleTrustProvider.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Verified Access Instance Trust Provider Attachments using the `verifiedaccess_instance_id` and `verifiedaccess_trust_provider_id` separated by a forward slash (`/`). For example:

```sh
 $ pulumi import aws:verifiedaccess/instanceTrustProviderAttachment:InstanceTrustProviderAttachment example vai-1234567890abcdef0/vatp-8012925589
```
 