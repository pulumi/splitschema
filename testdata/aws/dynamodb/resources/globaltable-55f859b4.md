Manages [DynamoDB Global Tables V1 (version 2017.11.29)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V1.html). These are layered on top of existing DynamoDB Tables.

> **NOTE:** To instead manage [DynamoDB Global Tables V2 (version 2019.11.21)](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html), use the `aws.dynamodb.Table` resource `replica` configuration block.

> Note: There are many restrictions before you can properly create DynamoDB Global Tables in multiple regions. See the [AWS DynamoDB Global Table Requirements](http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables_reqs_bestpractices.html) for more information.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const us_east_1 = new aws.Provider("us-east-1", {region: "us-east-1"});
const us_west_2 = new aws.Provider("us-west-2", {region: "us-west-2"});
const us_east_1Table = new aws.dynamodb.Table("us-east-1Table", {
    hashKey: "myAttribute",
    streamEnabled: true,
    streamViewType: "NEW_AND_OLD_IMAGES",
    readCapacity: 1,
    writeCapacity: 1,
    attributes: [{
        name: "myAttribute",
        type: "S",
    }],
}, {
    provider: aws["us-east-1"],
});
const us_west_2Table = new aws.dynamodb.Table("us-west-2Table", {
    hashKey: "myAttribute",
    streamEnabled: true,
    streamViewType: "NEW_AND_OLD_IMAGES",
    readCapacity: 1,
    writeCapacity: 1,
    attributes: [{
        name: "myAttribute",
        type: "S",
    }],
}, {
    provider: aws["us-west-2"],
});
const myTable = new aws.dynamodb.GlobalTable("myTable", {replicas: [
    {
        regionName: "us-east-1",
    },
    {
        regionName: "us-west-2",
    },
]}, {
    provider: aws["us-east-1"],
    dependsOn: [
        us_east_1Table,
        us_west_2Table,
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

us_east_1 = aws.Provider("us-east-1", region="us-east-1")
us_west_2 = aws.Provider("us-west-2", region="us-west-2")
us_east_1_table = aws.dynamodb.Table("us-east-1Table",
    hash_key="myAttribute",
    stream_enabled=True,
    stream_view_type="NEW_AND_OLD_IMAGES",
    read_capacity=1,
    write_capacity=1,
    attributes=[aws.dynamodb.TableAttributeArgs(
        name="myAttribute",
        type="S",
    )],
    opts=pulumi.ResourceOptions(provider=aws["us-east-1"]))
us_west_2_table = aws.dynamodb.Table("us-west-2Table",
    hash_key="myAttribute",
    stream_enabled=True,
    stream_view_type="NEW_AND_OLD_IMAGES",
    read_capacity=1,
    write_capacity=1,
    attributes=[aws.dynamodb.TableAttributeArgs(
        name="myAttribute",
        type="S",
    )],
    opts=pulumi.ResourceOptions(provider=aws["us-west-2"]))
my_table = aws.dynamodb.GlobalTable("myTable", replicas=[
    aws.dynamodb.GlobalTableReplicaArgs(
        region_name="us-east-1",
    ),
    aws.dynamodb.GlobalTableReplicaArgs(
        region_name="us-west-2",
    ),
],
opts=pulumi.ResourceOptions(provider=aws["us-east-1"],
    depends_on=[
        us_east_1_table,
        us_west_2_table,
    ]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var us_east_1 = new Aws.Provider("us-east-1", new()
    {
        Region = "us-east-1",
    });

    var us_west_2 = new Aws.Provider("us-west-2", new()
    {
        Region = "us-west-2",
    });

    var us_east_1Table = new Aws.DynamoDB.Table("us-east-1Table", new()
    {
        HashKey = "myAttribute",
        StreamEnabled = true,
        StreamViewType = "NEW_AND_OLD_IMAGES",
        ReadCapacity = 1,
        WriteCapacity = 1,
        Attributes = new[]
        {
            new Aws.DynamoDB.Inputs.TableAttributeArgs
            {
                Name = "myAttribute",
                Type = "S",
            },
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Us_east_1,
    });

    var us_west_2Table = new Aws.DynamoDB.Table("us-west-2Table", new()
    {
        HashKey = "myAttribute",
        StreamEnabled = true,
        StreamViewType = "NEW_AND_OLD_IMAGES",
        ReadCapacity = 1,
        WriteCapacity = 1,
        Attributes = new[]
        {
            new Aws.DynamoDB.Inputs.TableAttributeArgs
            {
                Name = "myAttribute",
                Type = "S",
            },
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Us_west_2,
    });

    var myTable = new Aws.DynamoDB.GlobalTable("myTable", new()
    {
        Replicas = new[]
        {
            new Aws.DynamoDB.Inputs.GlobalTableReplicaArgs
            {
                RegionName = "us-east-1",
            },
            new Aws.DynamoDB.Inputs.GlobalTableReplicaArgs
            {
                RegionName = "us-west-2",
            },
        },
    }, new CustomResourceOptions
    {
        Provider = aws.Us_east_1,
        DependsOn = new[]
        {
            us_east_1Table,
            us_west_2Table,
        },
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
		_, err := aws.NewProvider(ctx, "us-east-1", &aws.ProviderArgs{
			Region: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}
		_, err = aws.NewProvider(ctx, "us-west-2", &aws.ProviderArgs{
			Region: pulumi.String("us-west-2"),
		})
		if err != nil {
			return err
		}
		_, err = dynamodb.NewTable(ctx, "us-east-1Table", &dynamodb.TableArgs{
			HashKey:        pulumi.String("myAttribute"),
			StreamEnabled:  pulumi.Bool(true),
			StreamViewType: pulumi.String("NEW_AND_OLD_IMAGES"),
			ReadCapacity:   pulumi.Int(1),
			WriteCapacity:  pulumi.Int(1),
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("myAttribute"),
					Type: pulumi.String("S"),
				},
			},
		}, pulumi.Provider(aws.UsEast1))
		if err != nil {
			return err
		}
		_, err = dynamodb.NewTable(ctx, "us-west-2Table", &dynamodb.TableArgs{
			HashKey:        pulumi.String("myAttribute"),
			StreamEnabled:  pulumi.Bool(true),
			StreamViewType: pulumi.String("NEW_AND_OLD_IMAGES"),
			ReadCapacity:   pulumi.Int(1),
			WriteCapacity:  pulumi.Int(1),
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("myAttribute"),
					Type: pulumi.String("S"),
				},
			},
		}, pulumi.Provider(aws.UsWest2))
		if err != nil {
			return err
		}
		_, err = dynamodb.NewGlobalTable(ctx, "myTable", &dynamodb.GlobalTableArgs{
			Replicas: dynamodb.GlobalTableReplicaArray{
				&dynamodb.GlobalTableReplicaArgs{
					RegionName: pulumi.String("us-east-1"),
				},
				&dynamodb.GlobalTableReplicaArgs{
					RegionName: pulumi.String("us-west-2"),
				},
			},
		}, pulumi.Provider(aws.UsEast1), pulumi.DependsOn([]pulumi.Resource{
			us_east_1Table,
			us_west_2Table,
		}))
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
import com.pulumi.aws.dynamodb.GlobalTable;
import com.pulumi.aws.dynamodb.GlobalTableArgs;
import com.pulumi.aws.dynamodb.inputs.GlobalTableReplicaArgs;
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
        var us_east_1 = new Provider("us-east-1", ProviderArgs.builder()        
            .region("us-east-1")
            .build());

        var us_west_2 = new Provider("us-west-2", ProviderArgs.builder()        
            .region("us-west-2")
            .build());

        var us_east_1Table = new Table("us-east-1Table", TableArgs.builder()        
            .hashKey("myAttribute")
            .streamEnabled(true)
            .streamViewType("NEW_AND_OLD_IMAGES")
            .readCapacity(1)
            .writeCapacity(1)
            .attributes(TableAttributeArgs.builder()
                .name("myAttribute")
                .type("S")
                .build())
            .build(), CustomResourceOptions.builder()
                .provider(aws.us-east-1())
                .build());

        var us_west_2Table = new Table("us-west-2Table", TableArgs.builder()        
            .hashKey("myAttribute")
            .streamEnabled(true)
            .streamViewType("NEW_AND_OLD_IMAGES")
            .readCapacity(1)
            .writeCapacity(1)
            .attributes(TableAttributeArgs.builder()
                .name("myAttribute")
                .type("S")
                .build())
            .build(), CustomResourceOptions.builder()
                .provider(aws.us-west-2())
                .build());

        var myTable = new GlobalTable("myTable", GlobalTableArgs.builder()        
            .replicas(            
                GlobalTableReplicaArgs.builder()
                    .regionName("us-east-1")
                    .build(),
                GlobalTableReplicaArgs.builder()
                    .regionName("us-west-2")
                    .build())
            .build(), CustomResourceOptions.builder()
                .provider(aws.us-east-1())
                .dependsOn(                
                    us_east_1Table,
                    us_west_2Table)
                .build());

    }
}
```
```yaml
resources:
  us-east-1:
    type: pulumi:providers:aws
    properties:
      region: us-east-1
  us-west-2:
    type: pulumi:providers:aws
    properties:
      region: us-west-2
  us-east-1Table:
    type: aws:dynamodb:Table
    properties:
      hashKey: myAttribute
      streamEnabled: true
      streamViewType: NEW_AND_OLD_IMAGES
      readCapacity: 1
      writeCapacity: 1
      attributes:
        - name: myAttribute
          type: S
    options:
      provider: ${aws"us-east-1"[%!s(MISSING)]}
  us-west-2Table:
    type: aws:dynamodb:Table
    properties:
      hashKey: myAttribute
      streamEnabled: true
      streamViewType: NEW_AND_OLD_IMAGES
      readCapacity: 1
      writeCapacity: 1
      attributes:
        - name: myAttribute
          type: S
    options:
      provider: ${aws"us-west-2"[%!s(MISSING)]}
  myTable:
    type: aws:dynamodb:GlobalTable
    properties:
      replicas:
        - regionName: us-east-1
        - regionName: us-west-2
    options:
      provider: ${aws"us-east-1"[%!s(MISSING)]}
      dependson:
        - ${["us-east-1Table"]}
        - ${["us-west-2Table"]}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DynamoDB Global Tables using the global table name. For example:

```sh
 $ pulumi import aws:dynamodb/globalTable:GlobalTable MyTable MyTable
```
 