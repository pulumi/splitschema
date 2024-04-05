Provides a resource to manage the [default AWS VPC](http://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/default-vpc.html)
in the current AWS Region.

If you created your AWS account after 2013-12-04 you have a default VPC in each AWS Region.

**This is an advanced resource** and has special caveats to be aware of when using it. Please read this document in its entirety before using this resource.

The `aws.ec2.DefaultVpc` resource behaves differently from normal resources in that if a default VPC exists, this provider does not _create_ this resource, but instead "adopts" it into management.
If no default VPC exists, the provider creates a new default VPC, which leads to the implicit creation of [other resources](https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html#default-vpc-components).
By default, `pulumi destroy` does not delete the default VPC but does remove the resource from the state.
Set the `force_destroy` argument to `true` to delete the default VPC.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage with tags:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.ec2.DefaultVpc("default", {tags: {
    Name: "Default VPC",
}});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.ec2.DefaultVpc("default", tags={
    "Name": "Default VPC",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Ec2.DefaultVpc("default", new()
    {
        Tags = 
        {
            { "Name", "Default VPC" },
        },
    });

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
		_, err := ec2.NewDefaultVpc(ctx, "default", &ec2.DefaultVpcArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Default VPC"),
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
import com.pulumi.aws.ec2.DefaultVpc;
import com.pulumi.aws.ec2.DefaultVpcArgs;
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
        var default_ = new DefaultVpc("default", DefaultVpcArgs.builder()        
            .tags(Map.of("Name", "Default VPC"))
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:ec2:DefaultVpc
    properties:
      tags:
        Name: Default VPC
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Default VPCs using the VPC `id`. For example:

```sh
 $ pulumi import aws:ec2/defaultVpc:DefaultVpc default vpc-a01106c2
```
 