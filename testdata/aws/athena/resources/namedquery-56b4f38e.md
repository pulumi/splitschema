Provides an Athena Named Query resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const hogeBucketV2 = new aws.s3.BucketV2("hogeBucketV2", {});
const testKey = new aws.kms.Key("testKey", {
    deletionWindowInDays: 7,
    description: "Athena KMS Key",
});
const testWorkgroup = new aws.athena.Workgroup("testWorkgroup", {configuration: {
    resultConfiguration: {
        encryptionConfiguration: {
            encryptionOption: "SSE_KMS",
            kmsKeyArn: testKey.arn,
        },
    },
}});
const hogeDatabase = new aws.athena.Database("hogeDatabase", {
    name: "users",
    bucket: hogeBucketV2.id,
});
const foo = new aws.athena.NamedQuery("foo", {
    workgroup: testWorkgroup.id,
    database: hogeDatabase.name,
    query: pulumi.interpolate`SELECT * FROM ${hogeDatabase.name} limit 10;`,
});
```
```python
import pulumi
import pulumi_aws as aws

hoge_bucket_v2 = aws.s3.BucketV2("hogeBucketV2")
test_key = aws.kms.Key("testKey",
    deletion_window_in_days=7,
    description="Athena KMS Key")
test_workgroup = aws.athena.Workgroup("testWorkgroup", configuration=aws.athena.WorkgroupConfigurationArgs(
    result_configuration=aws.athena.WorkgroupConfigurationResultConfigurationArgs(
        encryption_configuration=aws.athena.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs(
            encryption_option="SSE_KMS",
            kms_key_arn=test_key.arn,
        ),
    ),
))
hoge_database = aws.athena.Database("hogeDatabase",
    name="users",
    bucket=hoge_bucket_v2.id)
foo = aws.athena.NamedQuery("foo",
    workgroup=test_workgroup.id,
    database=hoge_database.name,
    query=hoge_database.name.apply(lambda name: f"SELECT * FROM {name} limit 10;"))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var hogeBucketV2 = new Aws.S3.BucketV2("hogeBucketV2");

    var testKey = new Aws.Kms.Key("testKey", new()
    {
        DeletionWindowInDays = 7,
        Description = "Athena KMS Key",
    });

    var testWorkgroup = new Aws.Athena.Workgroup("testWorkgroup", new()
    {
        Configuration = new Aws.Athena.Inputs.WorkgroupConfigurationArgs
        {
            ResultConfiguration = new Aws.Athena.Inputs.WorkgroupConfigurationResultConfigurationArgs
            {
                EncryptionConfiguration = new Aws.Athena.Inputs.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs
                {
                    EncryptionOption = "SSE_KMS",
                    KmsKeyArn = testKey.Arn,
                },
            },
        },
    });

    var hogeDatabase = new Aws.Athena.Database("hogeDatabase", new()
    {
        Name = "users",
        Bucket = hogeBucketV2.Id,
    });

    var foo = new Aws.Athena.NamedQuery("foo", new()
    {
        Workgroup = testWorkgroup.Id,
        Database = hogeDatabase.Name,
        Query = hogeDatabase.Name.Apply(name => $"SELECT * FROM {name} limit 10;"),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/athena"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		hogeBucketV2, err := s3.NewBucketV2(ctx, "hogeBucketV2", nil)
		if err != nil {
			return err
		}
		testKey, err := kms.NewKey(ctx, "testKey", &kms.KeyArgs{
			DeletionWindowInDays: pulumi.Int(7),
			Description:          pulumi.String("Athena KMS Key"),
		})
		if err != nil {
			return err
		}
		testWorkgroup, err := athena.NewWorkgroup(ctx, "testWorkgroup", &athena.WorkgroupArgs{
			Configuration: &athena.WorkgroupConfigurationArgs{
				ResultConfiguration: &athena.WorkgroupConfigurationResultConfigurationArgs{
					EncryptionConfiguration: &athena.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs{
						EncryptionOption: pulumi.String("SSE_KMS"),
						KmsKeyArn:        testKey.Arn,
					},
				},
			},
		})
		if err != nil {
			return err
		}
		hogeDatabase, err := athena.NewDatabase(ctx, "hogeDatabase", &athena.DatabaseArgs{
			Name:   pulumi.String("users"),
			Bucket: hogeBucketV2.ID(),
		})
		if err != nil {
			return err
		}
		_, err = athena.NewNamedQuery(ctx, "foo", &athena.NamedQueryArgs{
			Workgroup: testWorkgroup.ID(),
			Database:  hogeDatabase.Name,
			Query: hogeDatabase.Name.ApplyT(func(name string) (string, error) {
				return fmt.Sprintf("SELECT * FROM %v limit 10;", name), nil
			}).(pulumi.StringOutput),
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.athena.Workgroup;
import com.pulumi.aws.athena.WorkgroupArgs;
import com.pulumi.aws.athena.inputs.WorkgroupConfigurationArgs;
import com.pulumi.aws.athena.inputs.WorkgroupConfigurationResultConfigurationArgs;
import com.pulumi.aws.athena.inputs.WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs;
import com.pulumi.aws.athena.Database;
import com.pulumi.aws.athena.DatabaseArgs;
import com.pulumi.aws.athena.NamedQuery;
import com.pulumi.aws.athena.NamedQueryArgs;
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
        var hogeBucketV2 = new BucketV2("hogeBucketV2");

        var testKey = new Key("testKey", KeyArgs.builder()        
            .deletionWindowInDays(7)
            .description("Athena KMS Key")
            .build());

        var testWorkgroup = new Workgroup("testWorkgroup", WorkgroupArgs.builder()        
            .configuration(WorkgroupConfigurationArgs.builder()
                .resultConfiguration(WorkgroupConfigurationResultConfigurationArgs.builder()
                    .encryptionConfiguration(WorkgroupConfigurationResultConfigurationEncryptionConfigurationArgs.builder()
                        .encryptionOption("SSE_KMS")
                        .kmsKeyArn(testKey.arn())
                        .build())
                    .build())
                .build())
            .build());

        var hogeDatabase = new Database("hogeDatabase", DatabaseArgs.builder()        
            .name("users")
            .bucket(hogeBucketV2.id())
            .build());

        var foo = new NamedQuery("foo", NamedQueryArgs.builder()        
            .workgroup(testWorkgroup.id())
            .database(hogeDatabase.name())
            .query(hogeDatabase.name().applyValue(name -> String.format("SELECT * FROM %s limit 10;", name)))
            .build());

    }
}
```
```yaml
resources:
  hogeBucketV2:
    type: aws:s3:BucketV2
  testKey:
    type: aws:kms:Key
    properties:
      deletionWindowInDays: 7
      description: Athena KMS Key
  testWorkgroup:
    type: aws:athena:Workgroup
    properties:
      configuration:
        resultConfiguration:
          encryptionConfiguration:
            encryptionOption: SSE_KMS
            kmsKeyArn: ${testKey.arn}
  hogeDatabase:
    type: aws:athena:Database
    properties:
      name: users
      bucket: ${hogeBucketV2.id}
  foo:
    type: aws:athena:NamedQuery
    properties:
      workgroup: ${testWorkgroup.id}
      database: ${hogeDatabase.name}
      query: SELECT * FROM ${hogeDatabase.name} limit 10;
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Athena Named Query using the query ID. For example:

```sh
 $ pulumi import aws:athena/namedQuery:NamedQuery example 0123456789
```
 