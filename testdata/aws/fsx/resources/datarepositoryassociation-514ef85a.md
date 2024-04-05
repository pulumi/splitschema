Manages a FSx for Lustre Data Repository Association. See [Linking your file system to an S3 bucket](https://docs.aws.amazon.com/fsx/latest/LustreGuide/create-dra-linked-data-repo.html) for more information.

> **NOTE:** Data Repository Associations are only compatible with AWS FSx for Lustre File Systems and `PERSISTENT_2` deployment type.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {});
const exampleBucketAclV2 = new aws.s3.BucketAclV2("exampleBucketAclV2", {
    bucket: exampleBucketV2.id,
    acl: "private",
});
const exampleLustreFileSystem = new aws.fsx.LustreFileSystem("exampleLustreFileSystem", {
    storageCapacity: 1200,
    subnetIds: [aws_subnet.example.id],
    deploymentType: "PERSISTENT_2",
    perUnitStorageThroughput: 125,
});
const exampleDataRepositoryAssociation = new aws.fsx.DataRepositoryAssociation("exampleDataRepositoryAssociation", {
    fileSystemId: exampleLustreFileSystem.id,
    dataRepositoryPath: pulumi.interpolate`s3://${exampleBucketV2.id}`,
    fileSystemPath: "/my-bucket",
    s3: {
        autoExportPolicy: {
            events: [
                "NEW",
                "CHANGED",
                "DELETED",
            ],
        },
        autoImportPolicy: {
            events: [
                "NEW",
                "CHANGED",
                "DELETED",
            ],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2")
example_bucket_acl_v2 = aws.s3.BucketAclV2("exampleBucketAclV2",
    bucket=example_bucket_v2.id,
    acl="private")
example_lustre_file_system = aws.fsx.LustreFileSystem("exampleLustreFileSystem",
    storage_capacity=1200,
    subnet_ids=[aws_subnet["example"]["id"]],
    deployment_type="PERSISTENT_2",
    per_unit_storage_throughput=125)
example_data_repository_association = aws.fsx.DataRepositoryAssociation("exampleDataRepositoryAssociation",
    file_system_id=example_lustre_file_system.id,
    data_repository_path=example_bucket_v2.id.apply(lambda id: f"s3://{id}"),
    file_system_path="/my-bucket",
    s3=aws.fsx.DataRepositoryAssociationS3Args(
        auto_export_policy=aws.fsx.DataRepositoryAssociationS3AutoExportPolicyArgs(
            events=[
                "NEW",
                "CHANGED",
                "DELETED",
            ],
        ),
        auto_import_policy=aws.fsx.DataRepositoryAssociationS3AutoImportPolicyArgs(
            events=[
                "NEW",
                "CHANGED",
                "DELETED",
            ],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2");

    var exampleBucketAclV2 = new Aws.S3.BucketAclV2("exampleBucketAclV2", new()
    {
        Bucket = exampleBucketV2.Id,
        Acl = "private",
    });

    var exampleLustreFileSystem = new Aws.Fsx.LustreFileSystem("exampleLustreFileSystem", new()
    {
        StorageCapacity = 1200,
        SubnetIds = new[]
        {
            aws_subnet.Example.Id,
        },
        DeploymentType = "PERSISTENT_2",
        PerUnitStorageThroughput = 125,
    });

    var exampleDataRepositoryAssociation = new Aws.Fsx.DataRepositoryAssociation("exampleDataRepositoryAssociation", new()
    {
        FileSystemId = exampleLustreFileSystem.Id,
        DataRepositoryPath = exampleBucketV2.Id.Apply(id => $"s3://{id}"),
        FileSystemPath = "/my-bucket",
        S3 = new Aws.Fsx.Inputs.DataRepositoryAssociationS3Args
        {
            AutoExportPolicy = new Aws.Fsx.Inputs.DataRepositoryAssociationS3AutoExportPolicyArgs
            {
                Events = new[]
                {
                    "NEW",
                    "CHANGED",
                    "DELETED",
                },
            },
            AutoImportPolicy = new Aws.Fsx.Inputs.DataRepositoryAssociationS3AutoImportPolicyArgs
            {
                Events = new[]
                {
                    "NEW",
                    "CHANGED",
                    "DELETED",
                },
            },
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fsx"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", nil)
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "exampleBucketAclV2", &s3.BucketAclV2Args{
			Bucket: exampleBucketV2.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		exampleLustreFileSystem, err := fsx.NewLustreFileSystem(ctx, "exampleLustreFileSystem", &fsx.LustreFileSystemArgs{
			StorageCapacity: pulumi.Int(1200),
			SubnetIds: pulumi.String{
				aws_subnet.Example.Id,
			},
			DeploymentType:           pulumi.String("PERSISTENT_2"),
			PerUnitStorageThroughput: pulumi.Int(125),
		})
		if err != nil {
			return err
		}
		_, err = fsx.NewDataRepositoryAssociation(ctx, "exampleDataRepositoryAssociation", &fsx.DataRepositoryAssociationArgs{
			FileSystemId: exampleLustreFileSystem.ID(),
			DataRepositoryPath: exampleBucketV2.ID().ApplyT(func(id string) (string, error) {
				return fmt.Sprintf("s3://%v", id), nil
			}).(pulumi.StringOutput),
			FileSystemPath: pulumi.String("/my-bucket"),
			S3: &fsx.DataRepositoryAssociationS3Args{
				AutoExportPolicy: &fsx.DataRepositoryAssociationS3AutoExportPolicyArgs{
					Events: pulumi.StringArray{
						pulumi.String("NEW"),
						pulumi.String("CHANGED"),
						pulumi.String("DELETED"),
					},
				},
				AutoImportPolicy: &fsx.DataRepositoryAssociationS3AutoImportPolicyArgs{
					Events: pulumi.StringArray{
						pulumi.String("NEW"),
						pulumi.String("CHANGED"),
						pulumi.String("DELETED"),
					},
				},
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
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.fsx.LustreFileSystem;
import com.pulumi.aws.fsx.LustreFileSystemArgs;
import com.pulumi.aws.fsx.DataRepositoryAssociation;
import com.pulumi.aws.fsx.DataRepositoryAssociationArgs;
import com.pulumi.aws.fsx.inputs.DataRepositoryAssociationS3Args;
import com.pulumi.aws.fsx.inputs.DataRepositoryAssociationS3AutoExportPolicyArgs;
import com.pulumi.aws.fsx.inputs.DataRepositoryAssociationS3AutoImportPolicyArgs;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2");

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .acl("private")
            .build());

        var exampleLustreFileSystem = new LustreFileSystem("exampleLustreFileSystem", LustreFileSystemArgs.builder()        
            .storageCapacity(1200)
            .subnetIds(aws_subnet.example().id())
            .deploymentType("PERSISTENT_2")
            .perUnitStorageThroughput(125)
            .build());

        var exampleDataRepositoryAssociation = new DataRepositoryAssociation("exampleDataRepositoryAssociation", DataRepositoryAssociationArgs.builder()        
            .fileSystemId(exampleLustreFileSystem.id())
            .dataRepositoryPath(exampleBucketV2.id().applyValue(id -> String.format("s3://%s", id)))
            .fileSystemPath("/my-bucket")
            .s3(DataRepositoryAssociationS3Args.builder()
                .autoExportPolicy(DataRepositoryAssociationS3AutoExportPolicyArgs.builder()
                    .events(                    
                        "NEW",
                        "CHANGED",
                        "DELETED")
                    .build())
                .autoImportPolicy(DataRepositoryAssociationS3AutoImportPolicyArgs.builder()
                    .events(                    
                        "NEW",
                        "CHANGED",
                        "DELETED")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${exampleBucketV2.id}
      acl: private
  exampleLustreFileSystem:
    type: aws:fsx:LustreFileSystem
    properties:
      storageCapacity: 1200
      subnetIds:
        - ${aws_subnet.example.id}
      deploymentType: PERSISTENT_2
      perUnitStorageThroughput: 125
  exampleDataRepositoryAssociation:
    type: aws:fsx:DataRepositoryAssociation
    properties:
      fileSystemId: ${exampleLustreFileSystem.id}
      dataRepositoryPath: s3://${exampleBucketV2.id}
      fileSystemPath: /my-bucket
      s3:
        autoExportPolicy:
          events:
            - NEW
            - CHANGED
            - DELETED
        autoImportPolicy:
          events:
            - NEW
            - CHANGED
            - DELETED
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import FSx Data Repository Associations using the `id`. For example:

```sh
 $ pulumi import aws:fsx/dataRepositoryAssociation:DataRepositoryAssociation example dra-0b1cfaeca11088b10
```
 