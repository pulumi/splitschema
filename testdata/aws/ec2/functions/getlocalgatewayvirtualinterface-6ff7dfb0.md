Provides details about an EC2 Local Gateway Virtual Interface. More information can be found in the [Outposts User Guide](https://docs.aws.amazon.com/outposts/latest/userguide/outposts-networking-components.html#routing).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = .map(([, ]) => (aws.ec2.getLocalGatewayVirtualInterface({
    id: __value,
})));
```
```python
import pulumi
import pulumi_aws as aws

example = [aws.ec2.get_local_gateway_virtual_interface(id=__value) for __key, __value in data["aws_ec2_local_gateway_virtual_interface_group"]["example"]["local_gateway_virtual_interface_ids"]]
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = .Select(__value => 
    {
        return Aws.Ec2.GetLocalGatewayVirtualInterface.Invoke(new()
        {
            Id = __value,
        });
    }).ToList();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_ := "TODO: For expression"
		return nil
	})
}
```
{{% /example %}}
{{% /examples %}}