Provides a managed prefix list resource.

> **NOTE on Managed Prefix Lists and Managed Prefix List Entries:** The provider
currently provides both a standalone Managed Prefix List Entry resource (a single entry),
and a Managed Prefix List resource with entries defined in-line. At this time you
cannot use a Managed Prefix List with in-line rules in conjunction with any Managed
Prefix List Entry resources. Doing so will cause a conflict of entries and will overwrite entries.

> **NOTE on `max_entries`:** When you reference a Prefix List in a resource,
the maximum number of entries for the prefix lists counts as the same number of rules
or entries for the resource. For example, if you create a prefix list with a maximum
of 20 entries and you reference that prefix list in a security group rule, this counts
as 20 rules for the security group.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.ManagedPrefixList("example", {
    addressFamily: "IPv4",
    maxEntries: 5,
    entries: [
        {
            cidr: aws_vpc.example.cidr_block,
            description: "Primary",
        },
        {
            cidr: aws_vpc_ipv4_cidr_block_association.example.cidr_block,
            description: "Secondary",
        },
    ],
    tags: {
        Env: "live",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.ManagedPrefixList("example",
    address_family="IPv4",
    max_entries=5,
    entries=[
        aws.ec2.ManagedPrefixListEntryArgs(
            cidr=aws_vpc["example"]["cidr_block"],
            description="Primary",
        ),
        aws.ec2.ManagedPrefixListEntryArgs(
            cidr=aws_vpc_ipv4_cidr_block_association["example"]["cidr_block"],
            description="Secondary",
        ),
    ],
    tags={
        "Env": "live",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.ManagedPrefixList("example", new()
    {
        AddressFamily = "IPv4",
        MaxEntries = 5,
        Entries = new[]
        {
            new Aws.Ec2.Inputs.ManagedPrefixListEntryArgs
            {
                Cidr = aws_vpc.Example.Cidr_block,
                Description = "Primary",
            },
            new Aws.Ec2.Inputs.ManagedPrefixListEntryArgs
            {
                Cidr = aws_vpc_ipv4_cidr_block_association.Example.Cidr_block,
                Description = "Secondary",
            },
        },
        Tags = 
        {
            { "Env", "live" },
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
		_, err := ec2.NewManagedPrefixList(ctx, "example", &ec2.ManagedPrefixListArgs{
			AddressFamily: pulumi.String("IPv4"),
			MaxEntries:    pulumi.Int(5),
			Entries: ec2.ManagedPrefixListEntryTypeArray{
				&ec2.ManagedPrefixListEntryTypeArgs{
					Cidr:        pulumi.Any(aws_vpc.Example.Cidr_block),
					Description: pulumi.String("Primary"),
				},
				&ec2.ManagedPrefixListEntryTypeArgs{
					Cidr:        pulumi.Any(aws_vpc_ipv4_cidr_block_association.Example.Cidr_block),
					Description: pulumi.String("Secondary"),
				},
			},
			Tags: pulumi.StringMap{
				"Env": pulumi.String("live"),
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
import com.pulumi.aws.ec2.ManagedPrefixList;
import com.pulumi.aws.ec2.ManagedPrefixListArgs;
import com.pulumi.aws.ec2.inputs.ManagedPrefixListEntryArgs;
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
        var example = new ManagedPrefixList("example", ManagedPrefixListArgs.builder()        
            .addressFamily("IPv4")
            .maxEntries(5)
            .entries(            
                ManagedPrefixListEntryArgs.builder()
                    .cidr(aws_vpc.example().cidr_block())
                    .description("Primary")
                    .build(),
                ManagedPrefixListEntryArgs.builder()
                    .cidr(aws_vpc_ipv4_cidr_block_association.example().cidr_block())
                    .description("Secondary")
                    .build())
            .tags(Map.of("Env", "live"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:ManagedPrefixList
    properties:
      addressFamily: IPv4
      maxEntries: 5
      entries:
        - cidr: ${aws_vpc.example.cidr_block}
          description: Primary
        - cidr: ${aws_vpc_ipv4_cidr_block_association.example.cidr_block}
          description: Secondary
      tags:
        Env: live
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Prefix Lists using the `id`. For example:

```sh
 $ pulumi import aws:ec2/managedPrefixList:ManagedPrefixList default pl-0570a1d2d725c16be
```
 