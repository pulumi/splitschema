Resource for managing an Athena Prepared Statement.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testBucketV2 = new aws.s3.BucketV2("testBucketV2", {forceDestroy: true});
const testWorkgroup = new aws.athena.Workgroup("testWorkgroup", {});
const testDatabase = new aws.athena.Database("testDatabase", {
    name: "example",
    bucket: testBucketV2.bucket,
});
const testPreparedStatement = new aws.athena.PreparedStatement("testPreparedStatement", {
    queryStatement: pulumi.interpolate`SELECT * FROM ${testDatabase.name} WHERE x = ?`,
    workgroup: testWorkgroup.name,
});
```
```python
import pulumi
import pulumi_aws as aws

test_bucket_v2 = aws.s3.BucketV2("testBucketV2", force_destroy=True)
test_workgroup = aws.athena.Workgroup("testWorkgroup")
test_database = aws.athena.Database("testDatabase",
    name="example",
    bucket=test_bucket_v2.bucket)
test_prepared_statement = aws.athena.PreparedStatement("testPreparedStatement",
    query_statement=test_database.name.apply(lambda name: f"SELECT * FROM {name} WHERE x = ?"),
    workgroup=test_workgroup.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testBucketV2 = new Aws.S3.BucketV2("testBucketV2", new()
    {
        ForceDestroy = true,
    });

    var testWorkgroup = new Aws.Athena.Workgroup("testWorkgroup");

    var testDatabase = new Aws.Athena.Database("testDatabase", new()
    {
        Name = "example",
        Bucket = testBucketV2.Bucket,
    });

    var testPreparedStatement = new Aws.Athena.PreparedStatement("testPreparedStatement", new()
    {
        QueryStatement = testDatabase.Name.Apply(name => $"SELECT * FROM {name} WHERE x = ?"),
        Workgroup = testWorkgroup.Name,
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/athena"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testBucketV2, err := s3.NewBucketV2(ctx, "testBucketV2", &s3.BucketV2Args{
			ForceDestroy: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		testWorkgroup, err := athena.NewWorkgroup(ctx, "testWorkgroup", nil)
		if err != nil {
			return err
		}
		testDatabase, err := athena.NewDatabase(ctx, "testDatabase", &athena.DatabaseArgs{
			Name:   pulumi.String("example"),
			Bucket: testBucketV2.Bucket,
		})
		if err != nil {
			return err
		}
		_, err = athena.NewPreparedStatement(ctx, "testPreparedStatement", &athena.PreparedStatementArgs{
			QueryStatement: testDatabase.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("SELECT * FROM %v WHERE x = ?", name), nil
			}).(pulumi.StringOutput),
			Workgroup: testWorkgroup.Name,
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.athena.Workgroup;
import com.pulumi.aws.athena.Database;
import com.pulumi.aws.athena.DatabaseArgs;
import com.pulumi.aws.athena.PreparedStatement;
import com.pulumi.aws.athena.PreparedStatementArgs;
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
        var testBucketV2 = new BucketV2("testBucketV2", BucketV2Args.builder()        
            .forceDestroy(true)
            .build());

        var testWorkgroup = new Workgroup("testWorkgroup");

        var testDatabase = new Database("testDatabase", DatabaseArgs.builder()        
            .name("example")
            .bucket(testBucketV2.bucket())
            .build());

        var testPreparedStatement = new PreparedStatement("testPreparedStatement", PreparedStatementArgs.builder()        
            .queryStatement(testDatabase.name().applyValue(name -> String.format("SELECT * FROM %s WHERE x = ?", name)))
            .workgroup(testWorkgroup.name())
            .build());

    }
}
```
```yaml
resources:
  testBucketV2:
    type: aws:s3:BucketV2
    properties:
      forceDestroy: true
  testWorkgroup:
    type: aws:athena:Workgroup
  testDatabase:
    type: aws:athena:Database
    properties:
      name: example
      bucket: ${testBucketV2.bucket}
  testPreparedStatement:
    type: aws:athena:PreparedStatement
    properties:
      queryStatement: SELECT * FROM ${testDatabase.name} WHERE x = ?
      workgroup: ${testWorkgroup.name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Athena Prepared Statement using the `WORKGROUP-NAME/STATEMENT-NAME`. For example:

```sh
 $ pulumi import aws:athena/preparedStatement:PreparedStatement example 12345abcde/example
```
 