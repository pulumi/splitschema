Use the `aws_prefix_list_entry` resource to manage a managed prefix list entry.

> **NOTE:** Pulumi currently provides two resources for managing Managed Prefix Lists and Managed Prefix List Entries. The standalone resource, Managed Prefix List Entry, is used to manage a single entry. The Managed Prefix List resource is used to manage multiple entries defined in-line. It is important to note that you cannot use a Managed Prefix List with in-line rules in conjunction with any Managed Prefix List Entry resources. This will result in a conflict of entries and will cause the entries to be overwritten.

> **NOTE:** To improve execution times on larger updates, it is recommended to use the inline `entry` block as part of the Managed Prefix List resource when creating a prefix list with more than 100 entries. You can find more information about the resource here.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.ManagedPrefixList("example", {
    addressFamily: "IPv4",
    maxEntries: 5,
    tags: {
        Env: "live",
    },
});
const entry1 = new aws.ec2.ManagedPrefixListEntry("entry1", {
    cidr: aws_vpc.example.cidr_block,
    description: "Primary",
    prefixListId: example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.ManagedPrefixList("example",
    address_family="IPv4",
    max_entries=5,
    tags={
        "Env": "live",
    })
entry1 = aws.ec2.ManagedPrefixListEntry("entry1",
    cidr=aws_vpc["example"]["cidr_block"],
    description="Primary",
    prefix_list_id=example.id)
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
        Tags = 
        {
            { "Env", "live" },
        },
    });

    var entry1 = new Aws.Ec2.ManagedPrefixListEntry("entry1", new()
    {
        Cidr = aws_vpc.Example.Cidr_block,
        Description = "Primary",
        PrefixListId = example.Id,
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
		example, err := ec2.NewManagedPrefixList(ctx, "example", &ec2.ManagedPrefixListArgs{
			AddressFamily: pulumi.String("IPv4"),
			MaxEntries:    pulumi.Int(5),
			Tags: pulumi.StringMap{
				"Env": pulumi.String("live"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewManagedPrefixListEntry(ctx, "entry1", &ec2.ManagedPrefixListEntryArgs{
			Cidr:         pulumi.Any(aws_vpc.Example.Cidr_block),
			Description:  pulumi.String("Primary"),
			PrefixListId: example.ID(),
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
import com.pulumi.aws.ec2.ManagedPrefixListEntry;
import com.pulumi.aws.ec2.ManagedPrefixListEntryArgs;
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
            .tags(Map.of("Env", "live"))
            .build());

        var entry1 = new ManagedPrefixListEntry("entry1", ManagedPrefixListEntryArgs.builder()        
            .cidr(aws_vpc.example().cidr_block())
            .description("Primary")
            .prefixListId(example.id())
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
      tags:
        Env: live
  entry1:
    type: aws:ec2:ManagedPrefixListEntry
    properties:
      cidr: ${aws_vpc.example.cidr_block}
      description: Primary
      prefixListId: ${example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import prefix list entries using `prefix_list_id` and `cidr` separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:ec2/managedPrefixListEntry:ManagedPrefixListEntry default pl-0570a1d2d725c16be,10.0.3.0/24
```
 