Resource for managing an AWS Network Manager Attachment Accepter.

{{% examples %}}
## Example Usage
{{% example %}}
### Example with VPC attachment

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.networkmanager.AttachmentAccepter("test", {
    attachmentId: aws_networkmanager_vpc_attachment.vpc.id,
    attachmentType: aws_networkmanager_vpc_attachment.vpc.attachment_type,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.networkmanager.AttachmentAccepter("test",
    attachment_id=aws_networkmanager_vpc_attachment["vpc"]["id"],
    attachment_type=aws_networkmanager_vpc_attachment["vpc"]["attachment_type"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.NetworkManager.AttachmentAccepter("test", new()
    {
        AttachmentId = aws_networkmanager_vpc_attachment.Vpc.Id,
        AttachmentType = aws_networkmanager_vpc_attachment.Vpc.Attachment_type,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkmanager.NewAttachmentAccepter(ctx, "test", &networkmanager.AttachmentAccepterArgs{
			AttachmentId:   pulumi.Any(aws_networkmanager_vpc_attachment.Vpc.Id),
			AttachmentType: pulumi.Any(aws_networkmanager_vpc_attachment.Vpc.Attachment_type),
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
import com.pulumi.aws.networkmanager.AttachmentAccepter;
import com.pulumi.aws.networkmanager.AttachmentAccepterArgs;
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
        var test = new AttachmentAccepter("test", AttachmentAccepterArgs.builder()        
            .attachmentId(aws_networkmanager_vpc_attachment.vpc().id())
            .attachmentType(aws_networkmanager_vpc_attachment.vpc().attachment_type())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:networkmanager:AttachmentAccepter
    properties:
      attachmentId: ${aws_networkmanager_vpc_attachment.vpc.id}
      attachmentType: ${aws_networkmanager_vpc_attachment.vpc.attachment_type}
```
{{% /example %}}
{{% example %}}
### Example with site-to-site VPN attachment

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.networkmanager.AttachmentAccepter("test", {
    attachmentId: aws_networkmanager_site_to_site_vpn_attachment.vpn.id,
    attachmentType: aws_networkmanager_site_to_site_vpn_attachment.vpn.attachment_type,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.networkmanager.AttachmentAccepter("test",
    attachment_id=aws_networkmanager_site_to_site_vpn_attachment["vpn"]["id"],
    attachment_type=aws_networkmanager_site_to_site_vpn_attachment["vpn"]["attachment_type"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.NetworkManager.AttachmentAccepter("test", new()
    {
        AttachmentId = aws_networkmanager_site_to_site_vpn_attachment.Vpn.Id,
        AttachmentType = aws_networkmanager_site_to_site_vpn_attachment.Vpn.Attachment_type,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/networkmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkmanager.NewAttachmentAccepter(ctx, "test", &networkmanager.AttachmentAccepterArgs{
			AttachmentId:   pulumi.Any(aws_networkmanager_site_to_site_vpn_attachment.Vpn.Id),
			AttachmentType: pulumi.Any(aws_networkmanager_site_to_site_vpn_attachment.Vpn.Attachment_type),
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
import com.pulumi.aws.networkmanager.AttachmentAccepter;
import com.pulumi.aws.networkmanager.AttachmentAccepterArgs;
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
        var test = new AttachmentAccepter("test", AttachmentAccepterArgs.builder()        
            .attachmentId(aws_networkmanager_site_to_site_vpn_attachment.vpn().id())
            .attachmentType(aws_networkmanager_site_to_site_vpn_attachment.vpn().attachment_type())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:networkmanager:AttachmentAccepter
    properties:
      attachmentId: ${aws_networkmanager_site_to_site_vpn_attachment.vpn.id}
      attachmentType: ${aws_networkmanager_site_to_site_vpn_attachment.vpn.attachment_type}
```
{{% /example %}}
{{% /examples %}}