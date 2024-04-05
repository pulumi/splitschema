Provides a Glue Connection resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Non-VPC Connection

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Connection("example", {connectionProperties: {
    JDBC_CONNECTION_URL: "jdbc:mysql://example.com/exampledatabase",
    PASSWORD: "examplepassword",
    USERNAME: "exampleusername",
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Connection("example", connection_properties={
    "JDBC_CONNECTION_URL": "jdbc:mysql://example.com/exampledatabase",
    "PASSWORD": "examplepassword",
    "USERNAME": "exampleusername",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Connection("example", new()
    {
        ConnectionProperties = 
        {
            { "JDBC_CONNECTION_URL", "jdbc:mysql://example.com/exampledatabase" },
            { "PASSWORD", "examplepassword" },
            { "USERNAME", "exampleusername" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewConnection(ctx, "example", &glue.ConnectionArgs{
			ConnectionProperties: pulumi.StringMap{
				"JDBC_CONNECTION_URL": pulumi.String("jdbc:mysql://example.com/exampledatabase"),
				"PASSWORD":            pulumi.String("examplepassword"),
				"USERNAME":            pulumi.String("exampleusername"),
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
import com.pulumi.aws.glue.Connection;
import com.pulumi.aws.glue.ConnectionArgs;
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
        var example = new Connection("example", ConnectionArgs.builder()        
            .connectionProperties(Map.ofEntries(
                Map.entry("JDBC_CONNECTION_URL", "jdbc:mysql://example.com/exampledatabase"),
                Map.entry("PASSWORD", "examplepassword"),
                Map.entry("USERNAME", "exampleusername")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Connection
    properties:
      connectionProperties:
        JDBC_CONNECTION_URL: jdbc:mysql://example.com/exampledatabase
        PASSWORD: examplepassword
        USERNAME: exampleusername
```
{{% /example %}}
{{% example %}}
### VPC Connection

For more information, see the [AWS Documentation](https://docs.aws.amazon.com/glue/latest/dg/populate-add-connection.html#connection-JDBC-VPC).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Connection("example", {
    connectionProperties: {
        JDBC_CONNECTION_URL: `jdbc:mysql://${aws_rds_cluster.example.endpoint}/exampledatabase`,
        PASSWORD: "examplepassword",
        USERNAME: "exampleusername",
    },
    physicalConnectionRequirements: {
        availabilityZone: aws_subnet.example.availability_zone,
        securityGroupIdLists: [aws_security_group.example.id],
        subnetId: aws_subnet.example.id,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Connection("example",
    connection_properties={
        "JDBC_CONNECTION_URL": f"jdbc:mysql://{aws_rds_cluster['example']['endpoint']}/exampledatabase",
        "PASSWORD": "examplepassword",
        "USERNAME": "exampleusername",
    },
    physical_connection_requirements=aws.glue.ConnectionPhysicalConnectionRequirementsArgs(
        availability_zone=aws_subnet["example"]["availability_zone"],
        security_group_id_lists=[aws_security_group["example"]["id"]],
        subnet_id=aws_subnet["example"]["id"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Connection("example", new()
    {
        ConnectionProperties = 
        {
            { "JDBC_CONNECTION_URL", $"jdbc:mysql://{aws_rds_cluster.Example.Endpoint}/exampledatabase" },
            { "PASSWORD", "examplepassword" },
            { "USERNAME", "exampleusername" },
        },
        PhysicalConnectionRequirements = new Aws.Glue.Inputs.ConnectionPhysicalConnectionRequirementsArgs
        {
            AvailabilityZone = aws_subnet.Example.Availability_zone,
            SecurityGroupIdLists = new[]
            {
                aws_security_group.Example.Id,
            },
            SubnetId = aws_subnet.Example.Id,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/glue"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := glue.NewConnection(ctx, "example", &glue.ConnectionArgs{
			ConnectionProperties: pulumi.StringMap{
				"JDBC_CONNECTION_URL": pulumi.String(fmt.Sprintf("jdbc:mysql://%v/exampledatabase", aws_rds_cluster.Example.Endpoint)),
				"PASSWORD":            pulumi.String("examplepassword"),
				"USERNAME":            pulumi.String("exampleusername"),
			},
			PhysicalConnectionRequirements: &glue.ConnectionPhysicalConnectionRequirementsArgs{
				AvailabilityZone: pulumi.Any(aws_subnet.Example.Availability_zone),
				SecurityGroupIdLists: pulumi.StringArray{
					aws_security_group.Example.Id,
				},
				SubnetId: pulumi.Any(aws_subnet.Example.Id),
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
import com.pulumi.aws.glue.Connection;
import com.pulumi.aws.glue.ConnectionArgs;
import com.pulumi.aws.glue.inputs.ConnectionPhysicalConnectionRequirementsArgs;
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
        var example = new Connection("example", ConnectionArgs.builder()        
            .connectionProperties(Map.ofEntries(
                Map.entry("JDBC_CONNECTION_URL", String.format("jdbc:mysql://%s/exampledatabase", aws_rds_cluster.example().endpoint())),
                Map.entry("PASSWORD", "examplepassword"),
                Map.entry("USERNAME", "exampleusername")
            ))
            .physicalConnectionRequirements(ConnectionPhysicalConnectionRequirementsArgs.builder()
                .availabilityZone(aws_subnet.example().availability_zone())
                .securityGroupIdLists(aws_security_group.example().id())
                .subnetId(aws_subnet.example().id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Connection
    properties:
      connectionProperties:
        JDBC_CONNECTION_URL: jdbc:mysql://${aws_rds_cluster.example.endpoint}/exampledatabase
        PASSWORD: examplepassword
        USERNAME: exampleusername
      physicalConnectionRequirements:
        availabilityZone: ${aws_subnet.example.availability_zone}
        securityGroupIdLists:
          - ${aws_security_group.example.id}
        subnetId: ${aws_subnet.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Connections using the `CATALOG-ID` (AWS account ID if not custom) and `NAME`. For example:

```sh
 $ pulumi import aws:glue/connection:Connection MyConnection 123456789012:MyConnection
```
 