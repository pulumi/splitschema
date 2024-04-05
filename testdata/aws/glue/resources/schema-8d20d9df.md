Provides a Glue Schema resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.glue.Schema("example", {
    schemaName: "example",
    registryArn: aws_glue_registry.test.arn,
    dataFormat: "AVRO",
    compatibility: "NONE",
    schemaDefinition: "{\"type\": \"record\", \"name\": \"r1\", \"fields\": [ {\"name\": \"f1\", \"type\": \"int\"}, {\"name\": \"f2\", \"type\": \"string\"} ]}",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.glue.Schema("example",
    schema_name="example",
    registry_arn=aws_glue_registry["test"]["arn"],
    data_format="AVRO",
    compatibility="NONE",
    schema_definition="{\"type\": \"record\", \"name\": \"r1\", \"fields\": [ {\"name\": \"f1\", \"type\": \"int\"}, {\"name\": \"f2\", \"type\": \"string\"} ]}")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Glue.Schema("example", new()
    {
        SchemaName = "example",
        RegistryArn = aws_glue_registry.Test.Arn,
        DataFormat = "AVRO",
        Compatibility = "NONE",
        SchemaDefinition = "{\"type\": \"record\", \"name\": \"r1\", \"fields\": [ {\"name\": \"f1\", \"type\": \"int\"}, {\"name\": \"f2\", \"type\": \"string\"} ]}",
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
		_, err := glue.NewSchema(ctx, "example", &glue.SchemaArgs{
			SchemaName:       pulumi.String("example"),
			RegistryArn:      pulumi.Any(aws_glue_registry.Test.Arn),
			DataFormat:       pulumi.String("AVRO"),
			Compatibility:    pulumi.String("NONE"),
			SchemaDefinition: pulumi.String("{\"type\": \"record\", \"name\": \"r1\", \"fields\": [ {\"name\": \"f1\", \"type\": \"int\"}, {\"name\": \"f2\", \"type\": \"string\"} ]}"),
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
import com.pulumi.aws.glue.Schema;
import com.pulumi.aws.glue.SchemaArgs;
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
        var example = new Schema("example", SchemaArgs.builder()        
            .schemaName("example")
            .registryArn(aws_glue_registry.test().arn())
            .dataFormat("AVRO")
            .compatibility("NONE")
            .schemaDefinition("{\"type\": \"record\", \"name\": \"r1\", \"fields\": [ {\"name\": \"f1\", \"type\": \"int\"}, {\"name\": \"f2\", \"type\": \"string\"} ]}")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:glue:Schema
    properties:
      schemaName: example
      registryArn: ${aws_glue_registry.test.arn}
      dataFormat: AVRO
      compatibility: NONE
      schemaDefinition: '{"type": "record", "name": "r1", "fields": [ {"name": "f1", "type": "int"}, {"name": "f2", "type": "string"} ]}'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Glue Registries using `arn`. For example:

```sh
 $ pulumi import aws:glue/schema:Schema example arn:aws:glue:us-west-2:123456789012:schema/example/example
```
 