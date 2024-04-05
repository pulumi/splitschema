Manages a Direct Connect Gateway Association Proposal, typically for enabling cross-account associations. For single account associations, see the `aws.directconnect.GatewayAssociation` resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.directconnect.GatewayAssociationProposal("example", {
    dxGatewayId: aws_dx_gateway.example.id,
    dxGatewayOwnerAccountId: aws_dx_gateway.example.owner_account_id,
    associatedGatewayId: aws_vpn_gateway.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.directconnect.GatewayAssociationProposal("example",
    dx_gateway_id=aws_dx_gateway["example"]["id"],
    dx_gateway_owner_account_id=aws_dx_gateway["example"]["owner_account_id"],
    associated_gateway_id=aws_vpn_gateway["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DirectConnect.GatewayAssociationProposal("example", new()
    {
        DxGatewayId = aws_dx_gateway.Example.Id,
        DxGatewayOwnerAccountId = aws_dx_gateway.Example.Owner_account_id,
        AssociatedGatewayId = aws_vpn_gateway.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/directconnect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := directconnect.NewGatewayAssociationProposal(ctx, "example", &directconnect.GatewayAssociationProposalArgs{
			DxGatewayId:             pulumi.Any(aws_dx_gateway.Example.Id),
			DxGatewayOwnerAccountId: pulumi.Any(aws_dx_gateway.Example.Owner_account_id),
			AssociatedGatewayId:     pulumi.Any(aws_vpn_gateway.Example.Id),
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
import com.pulumi.aws.directconnect.GatewayAssociationProposal;
import com.pulumi.aws.directconnect.GatewayAssociationProposalArgs;
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
        var example = new GatewayAssociationProposal("example", GatewayAssociationProposalArgs.builder()        
            .dxGatewayId(aws_dx_gateway.example().id())
            .dxGatewayOwnerAccountId(aws_dx_gateway.example().owner_account_id())
            .associatedGatewayId(aws_vpn_gateway.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:directconnect:GatewayAssociationProposal
    properties:
      dxGatewayId: ${aws_dx_gateway.example.id}
      dxGatewayOwnerAccountId: ${aws_dx_gateway.example.owner_account_id}
      associatedGatewayId: ${aws_vpn_gateway.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using a proposal ID, Direct Connect Gateway ID and associated gateway ID separated by `/`:

__With `pulumi import`__, import Direct Connect Gateway Association Proposals using either a proposal ID or proposal ID, Direct Connect Gateway ID and associated gateway ID separated by `/`. For example:

Using a proposal ID:

```sh
 $ pulumi import aws:directconnect/gatewayAssociationProposal:GatewayAssociationProposal example ac90e981-b718-4364-872d-65478c84fafe
```
 Using a proposal ID, Direct Connect Gateway ID and associated gateway ID separated by `/`:

```sh
 $ pulumi import aws:directconnect/gatewayAssociationProposal:GatewayAssociationProposal example ac90e981-b718-4364-872d-65478c84fafe/abcd1234-dcba-5678-be23-cdef9876ab45/vgw-12345678
```
 The latter case is useful when a previous proposal has been accepted and deleted by AWS. The `aws_dx_gateway_association_proposal` resource will then represent a pseudo-proposal for the same Direct Connect Gateway and associated gateway. If no previous proposal is available, use a tool like [`uuidgen`](http://manpages.ubuntu.com/manpages/bionic/man1/uuidgen.1.html) to generate a new random pseudo-proposal ID.

