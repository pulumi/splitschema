Manages an individual Transfer Family resource tag. This resource should only be used in cases where Transfer Family resources are created outside the provider (e.g., Servers without AWS Management Console) or the tag key has the `aws:` prefix.

> **NOTE:** This tagging resource should not be combined with the resource for managing the parent resource. For example, using `aws.transfer.Server` and `aws.transfer.Tag` to manage tags of the same server will cause a perpetual difference where the `aws.transfer.Server` resource will try to remove the tag being added by the `aws.transfer.Tag` resource.

> **NOTE:** This tagging resource does not use the provider `ignore_tags` configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Server("example", {identityProviderType: "SERVICE_MANAGED"});
const zoneId = new aws.transfer.Tag("zoneId", {
    resourceArn: example.arn,
    key: "aws:transfer:route53HostedZoneId",
    value: "/hostedzone/MyHostedZoneId",
});
const hostname = new aws.transfer.Tag("hostname", {
    resourceArn: example.arn,
    key: "aws:transfer:customHostname",
    value: "example.com",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Server("example", identity_provider_type="SERVICE_MANAGED")
zone_id = aws.transfer.Tag("zoneId",
    resource_arn=example.arn,
    key="aws:transfer:route53HostedZoneId",
    value="/hostedzone/MyHostedZoneId")
hostname = aws.transfer.Tag("hostname",
    resource_arn=example.arn,
    key="aws:transfer:customHostname",
    value="example.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Server("example", new()
    {
        IdentityProviderType = "SERVICE_MANAGED",
    });

    var zoneId = new Aws.Transfer.Tag("zoneId", new()
    {
        ResourceArn = example.Arn,
        Key = "aws:transfer:route53HostedZoneId",
        Value = "/hostedzone/MyHostedZoneId",
    });

    var hostname = new Aws.Transfer.Tag("hostname", new()
    {
        ResourceArn = example.Arn,
        Key = "aws:transfer:customHostname",
        Value = "example.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := transfer.NewServer(ctx, "example", &transfer.ServerArgs{
			IdentityProviderType: pulumi.String("SERVICE_MANAGED"),
		})
		if err != nil {
			return err
		}
		_, err = transfer.NewTag(ctx, "zoneId", &transfer.TagArgs{
			ResourceArn: example.Arn,
			Key:         pulumi.String("aws:transfer:route53HostedZoneId"),
			Value:       pulumi.String("/hostedzone/MyHostedZoneId"),
		})
		if err != nil {
			return err
		}
		_, err = transfer.NewTag(ctx, "hostname", &transfer.TagArgs{
			ResourceArn: example.Arn,
			Key:         pulumi.String("aws:transfer:customHostname"),
			Value:       pulumi.String("example.com"),
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
import com.pulumi.aws.transfer.Server;
import com.pulumi.aws.transfer.ServerArgs;
import com.pulumi.aws.transfer.Tag;
import com.pulumi.aws.transfer.TagArgs;
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
        var example = new Server("example", ServerArgs.builder()        
            .identityProviderType("SERVICE_MANAGED")
            .build());

        var zoneId = new Tag("zoneId", TagArgs.builder()        
            .resourceArn(example.arn())
            .key("aws:transfer:route53HostedZoneId")
            .value("/hostedzone/MyHostedZoneId")
            .build());

        var hostname = new Tag("hostname", TagArgs.builder()        
            .resourceArn(example.arn())
            .key("aws:transfer:customHostname")
            .value("example.com")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Server
    properties:
      identityProviderType: SERVICE_MANAGED
  zoneId:
    type: aws:transfer:Tag
    properties:
      resourceArn: ${example.arn}
      key: aws:transfer:route53HostedZoneId
      value: /hostedzone/MyHostedZoneId
  hostname:
    type: aws:transfer:Tag
    properties:
      resourceArn: ${example.arn}
      key: aws:transfer:customHostname
      value: example.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_transfer_tag` using the Transfer Family resource identifier and key, separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:transfer/tag:Tag example arn:aws:transfer:us-east-1:123456789012:server/s-1234567890abcdef0,Name
```
 