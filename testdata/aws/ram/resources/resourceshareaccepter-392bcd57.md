Manage accepting a Resource Access Manager (RAM) Resource Share invitation. From a _receiver_ AWS account, accept an invitation to share resources that were shared by a _sender_ AWS account. To create a resource share in the _sender_, see the `aws.ram.ResourceShare` resource.

> **Note:** If both AWS accounts are in the same Organization and [RAM Sharing with AWS Organizations is enabled](https://docs.aws.amazon.com/ram/latest/userguide/getting-started-sharing.html#getting-started-sharing-orgs), this resource is not necessary as RAM Resource Share invitations are not used.

{{% examples %}}
## Example Usage
{{% example %}}

This configuration provides an example of using multiple AWS providers to configure two different AWS accounts. In the _sender_ account, the configuration creates a `aws.ram.ResourceShare` and uses a data source in the _receiver_ account to create a `aws.ram.PrincipalAssociation` resource with the _receiver's_ account ID. In the _receiver_ account, the configuration accepts the invitation to share resources with the `aws.ram.ResourceShareAccepter`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const alternate = new aws.Provider("alternate", {profile: "profile1"});
const senderShare = new aws.ram.ResourceShare("senderShare", {
    allowExternalPrincipals: true,
    tags: {
        Name: "tf-test-resource-share",
    },
}, {
    provider: aws.alternate,
});
const receiver = aws.getCallerIdentity({});
const senderInvite = new aws.ram.PrincipalAssociation("senderInvite", {
    principal: receiver.then(receiver => receiver.accountId),
    resourceShareArn: senderShare.arn,
}, {
    provider: aws.alternate,
});
const receiverAccept = new aws.ram.ResourceShareAccepter("receiverAccept", {shareArn: senderInvite.resourceShareArn});
```
```python
import pulumi
import pulumi_aws as aws

alternate = aws.Provider("alternate", profile="profile1")
sender_share = aws.ram.ResourceShare("senderShare",
    allow_external_principals=True,
    tags={
        "Name": "tf-test-resource-share",
    },
    opts=pulumi.ResourceOptions(provider=aws["alternate"]))
receiver = aws.get_caller_identity()
sender_invite = aws.ram.PrincipalAssociation("senderInvite",
    principal=receiver.account_id,
    resource_share_arn=sender_share.arn,
    opts=pulumi.ResourceOptions(provider=aws["alternate"]))
receiver_accept = aws.ram.ResourceShareAccepter("receiverAccept", share_arn=sender_invite.resource_share_arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var alternate = new Aws.Provider("alternate", new()
    {
        Profile = "profile1",
    });

    var senderShare = new Aws.Ram.ResourceShare("senderShare", new()
    {
        AllowExternalPrincipals = true,
        Tags = 
        {
            { "Name", "tf-test-resource-share" },
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Alternate,
    });

    var receiver = Aws.GetCallerIdentity.Invoke();

    var senderInvite = new Aws.Ram.PrincipalAssociation("senderInvite", new()
    {
        Principal = receiver.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
        ResourceShareArn = senderShare.Arn,
    }, new CustomResourceOptions
    {
        Provider = aws.Alternate,
    });

    var receiverAccept = new Aws.Ram.ResourceShareAccepter("receiverAccept", new()
    {
        ShareArn = senderInvite.ResourceShareArn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ram"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "alternate", &aws.ProviderArgs{
			Profile: pulumi.String("profile1"),
		})
		if err != nil {
			return err
		}
		senderShare, err := ram.NewResourceShare(ctx, "senderShare", &ram.ResourceShareArgs{
			AllowExternalPrincipals: pulumi.Bool(true),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-test-resource-share"),
			},
		}, pulumi.Provider(aws.Alternate))
		if err != nil {
			return err
		}
		receiver, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		senderInvite, err := ram.NewPrincipalAssociation(ctx, "senderInvite", &ram.PrincipalAssociationArgs{
			Principal:        *pulumi.String(receiver.AccountId),
			ResourceShareArn: senderShare.Arn,
		}, pulumi.Provider(aws.Alternate))
		if err != nil {
			return err
		}
		_, err = ram.NewResourceShareAccepter(ctx, "receiverAccept", &ram.ResourceShareAccepterArgs{
			ShareArn: senderInvite.ResourceShareArn,
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
import com.pulumi.aws.Provider;
import com.pulumi.aws.ProviderArgs;
import com.pulumi.aws.ram.ResourceShare;
import com.pulumi.aws.ram.ResourceShareArgs;
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.ram.PrincipalAssociation;
import com.pulumi.aws.ram.PrincipalAssociationArgs;
import com.pulumi.aws.ram.ResourceShareAccepter;
import com.pulumi.aws.ram.ResourceShareAccepterArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var alternate = new Provider("alternate", ProviderArgs.builder()        
            .profile("profile1")
            .build());

        var senderShare = new ResourceShare("senderShare", ResourceShareArgs.builder()        
            .allowExternalPrincipals(true)
            .tags(Map.of("Name", "tf-test-resource-share"))
            .build(), CustomResourceOptions.builder()
                .provider(aws.alternate())
                .build());

        final var receiver = AwsFunctions.getCallerIdentity();

        var senderInvite = new PrincipalAssociation("senderInvite", PrincipalAssociationArgs.builder()        
            .principal(receiver.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
            .resourceShareArn(senderShare.arn())
            .build(), CustomResourceOptions.builder()
                .provider(aws.alternate())
                .build());

        var receiverAccept = new ResourceShareAccepter("receiverAccept", ResourceShareAccepterArgs.builder()        
            .shareArn(senderInvite.resourceShareArn())
            .build());

    }
}
```
```yaml
resources:
  alternate:
    type: pulumi:providers:aws
    properties:
      profile: profile1
  senderShare:
    type: aws:ram:ResourceShare
    properties:
      allowExternalPrincipals: true
      tags:
        Name: tf-test-resource-share
    options:
      provider: ${aws.alternate}
  senderInvite:
    type: aws:ram:PrincipalAssociation
    properties:
      principal: ${receiver.accountId}
      resourceShareArn: ${senderShare.arn}
    options:
      provider: ${aws.alternate}
  receiverAccept:
    type: aws:ram:ResourceShareAccepter
    properties:
      shareArn: ${senderInvite.resourceShareArn}
variables:
  receiver:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import resource share accepters using the resource share ARN. For example:

```sh
 $ pulumi import aws:ram/resourceShareAccepter:ResourceShareAccepter example arn:aws:ram:us-east-1:123456789012:resource-share/c4b56393-e8d9-89d9-6dc9-883752de4767
```
 