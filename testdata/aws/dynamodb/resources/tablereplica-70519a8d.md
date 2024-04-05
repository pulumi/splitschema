Provides a DynamoDB table replica resource for [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html).

> **Note:** Use `lifecycle` `ignore_changes` for `replica` in the associated aws.dynamodb.Table configuration.

> **Note:** Do not use the `replica` configuration block of aws.dynamodb.Table together with this resource as the two configuration options are mutually exclusive.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.Provider("main", {region: "us-west-2"});
const alt = new aws.Provider("alt", {region: "us-east-2"});
const exampleTable = new aws.dynamodb.Table("exampleTable", {
    hashKey: "BrodoBaggins",
    billingMode: "PAY_PER_REQUEST",
    streamEnabled: true,
    streamViewType: "NEW_AND_OLD_IMAGES",
    attributes: [{
        name: "BrodoBaggins",
        type: "S",
    }],
}, {
    provider: aws.main,
});
const exampleTableReplica = new aws.dynamodb.TableReplica("exampleTableReplica", {
    globalTableArn: exampleTable.arn,
    tags: {
        Name: "IZPAWS",
        Pozo: "Amargo",
    },
}, {
    provider: aws.alt,
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.Provider("main", region="us-west-2")
alt = aws.Provider("alt", region="us-east-2")
example_table = aws.dynamodb.Table("exampleTable",
    hash_key="BrodoBaggins",
    billing_mode="PAY_PER_REQUEST",
    stream_enabled=True,
    stream_view_type="NEW_AND_OLD_IMAGES",
    attributes=[aws.dynamodb.TableAttributeArgs(
        name="BrodoBaggins",
        type="S",
    )],
    opts=pulumi.ResourceOptions(provider=aws["main"]))
example_table_replica = aws.dynamodb.TableReplica("exampleTableReplica",
    global_table_arn=example_table.arn,
    tags={
        "Name": "IZPAWS",
        "Pozo": "Amargo",
    },
    opts=pulumi.ResourceOptions(provider=aws["alt"]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Provider("main", new()
    {
        Region = "us-west-2",
    });

    var alt = new Aws.Provider("alt", new()
    {
        Region = "us-east-2",
    });

    var exampleTable = new Aws.DynamoDB.Table("exampleTable", new()
    {
        HashKey = "BrodoBaggins",
        BillingMode = "PAY_PER_REQUEST",
        StreamEnabled = true,
        StreamViewType = "NEW_AND_OLD_IMAGES",
        Attributes = new[]
        {
            new Aws.DynamoDB.Inputs.TableAttributeArgs
            {
                Name = "BrodoBaggins",
                Type = "S",
            },
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Main,
    });

    var exampleTableReplica = new Aws.DynamoDB.TableReplica("exampleTableReplica", new()
    {
        GlobalTableArn = exampleTable.Arn,
        Tags = 
        {
            { "Name", "IZPAWS" },
            { "Pozo", "Amargo" },
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Alt,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aws.NewProvider(ctx, "main", &aws.ProviderArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		_, err = aws.NewProvider(ctx, "alt", &aws.ProviderArgs{
			Region: pulumi.String("us-east-2"),
		})
		if err != nil {
			return err
		}
		exampleTable, err := dynamodb.NewTable(ctx, "exampleTable", &dynamodb.TableArgs{
			HashKey:        pulumi.String("BrodoBaggins"),
			BillingMode:    pulumi.String("PAY_PER_REQUEST"),
			StreamEnabled:  pulumi.Bool(true),
			StreamViewType: pulumi.String("NEW_AND_OLD_IMAGES"),
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("BrodoBaggins"),
					Type: pulumi.String("S"),
				},
			},
		}, pulumi.Provider(aws.Main))
		if err != nil {
			return err
		}
		_, err = dynamodb.NewTableReplica(ctx, "exampleTableReplica", &dynamodb.TableReplicaArgs{
			GlobalTableArn: exampleTable.Arn,
			Tags: pulumi.StringMap{
				"Name": pulumi.String("IZPAWS"),
				"Pozo": pulumi.String("Amargo"),
			},
		}, pulumi.Provider(aws.Alt))
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
import com.pulumi.aws.dynamodb.Table;
import com.pulumi.aws.dynamodb.TableArgs;
import com.pulumi.aws.dynamodb.inputs.TableAttributeArgs;
import com.pulumi.aws.dynamodb.TableReplica;
import com.pulumi.aws.dynamodb.TableReplicaArgs;
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
        var main = new Provider("main", ProviderArgs.builder()        
            .region("us-west-2")
            .build());

        var alt = new Provider("alt", ProviderArgs.builder()        
            .region("us-east-2")
            .build());

        var exampleTable = new Table("exampleTable", TableArgs.builder()        
            .hashKey("BrodoBaggins")
            .billingMode("PAY_PER_REQUEST")
            .streamEnabled(true)
            .streamViewType("NEW_AND_OLD_IMAGES")
            .attributes(TableAttributeArgs.builder()
                .name("BrodoBaggins")
                .type("S")
                .build())
            .build(), CustomResourceOptions.builder()
                .provider(aws.main())
                .build());

        var exampleTableReplica = new TableReplica("exampleTableReplica", TableReplicaArgs.builder()        
            .globalTableArn(exampleTable.arn())
            .tags(Map.ofEntries(
                Map.entry("Name", "IZPAWS"),
                Map.entry("Pozo", "Amargo")
            ))
            .build(), CustomResourceOptions.builder()
                .provider(aws.alt())
                .build());

    }
}
```
```yaml
resources:
  main:
    type: pulumi:providers:aws
    properties:
      region: us-west-2
  alt:
    type: pulumi:providers:aws
    properties:
      region: us-east-2
  exampleTable:
    type: aws:dynamodb:Table
    properties:
      hashKey: BrodoBaggins
      billingMode: PAY_PER_REQUEST
      streamEnabled: true
      streamViewType: NEW_AND_OLD_IMAGES
      attributes:
        - name: BrodoBaggins
          type: S
    options:
      provider: ${aws.main}
  exampleTableReplica:
    type: aws:dynamodb:TableReplica
    properties:
      globalTableArn: ${exampleTable.arn}
      tags:
        Name: IZPAWS
        Pozo: Amargo
    options:
      provider: ${aws.alt}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DynamoDB table replicas using the `table-name:main-region`. For example:

~> __Note:__ When importing, use the region where the initial or _main_ global table resides, _not_ the region of the replica.

```sh
 $ pulumi import aws:dynamodb/tableReplica:TableReplica example TestTable:us-west-2
```
 