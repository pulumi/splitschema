Provides a Timestream database resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.timestreamwrite.Database("example", {databaseName: "database-example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.timestreamwrite.Database("example", database_name="database-example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.TimestreamWrite.Database("example", new()
    {
        DatabaseName = "database-example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/timestreamwrite"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := timestreamwrite.NewDatabase(ctx, "example", &timestreamwrite.DatabaseArgs{
			DatabaseName: pulumi.String("database-example"),
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
import com.pulumi.aws.timestreamwrite.Database;
import com.pulumi.aws.timestreamwrite.DatabaseArgs;
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
        var example = new Database("example", DatabaseArgs.builder()        
            .databaseName("database-example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:timestreamwrite:Database
    properties:
      databaseName: database-example
```
{{% /example %}}
{{% example %}}
### Full usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.timestreamwrite.Database("example", {
    databaseName: "database-example",
    kmsKeyId: aws_kms_key.example.arn,
    tags: {
        Name: "value",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.timestreamwrite.Database("example",
    database_name="database-example",
    kms_key_id=aws_kms_key["example"]["arn"],
    tags={
        "Name": "value",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.TimestreamWrite.Database("example", new()
    {
        DatabaseName = "database-example",
        KmsKeyId = aws_kms_key.Example.Arn,
        Tags = 
        {
            { "Name", "value" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/timestreamwrite"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := timestreamwrite.NewDatabase(ctx, "example", &timestreamwrite.DatabaseArgs{
			DatabaseName: pulumi.String("database-example"),
			KmsKeyId:     pulumi.Any(aws_kms_key.Example.Arn),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("value"),
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
import com.pulumi.aws.timestreamwrite.Database;
import com.pulumi.aws.timestreamwrite.DatabaseArgs;
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
        var example = new Database("example", DatabaseArgs.builder()        
            .databaseName("database-example")
            .kmsKeyId(aws_kms_key.example().arn())
            .tags(Map.of("Name", "value"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:timestreamwrite:Database
    properties:
      databaseName: database-example
      kmsKeyId: ${aws_kms_key.example.arn}
      tags:
        Name: value
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Timestream databases using the `database_name`. For example:

```sh
 $ pulumi import aws:timestreamwrite/database:Database example example
```
 