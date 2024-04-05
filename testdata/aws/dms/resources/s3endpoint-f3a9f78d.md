Provides a DMS (Data Migration Service) S3 endpoint resource. DMS S3 endpoints can be created, updated, deleted, and imported.

> **Note:** AWS is deprecating `extra_connection_attributes`, such as used with `aws.dms.Endpoint`. This resource is an alternative to `aws.dms.Endpoint` and does not use `extra_connection_attributes`. (AWS currently includes `extra_connection_attributes` in the raw responses to the AWS Provider requests and so they may be visible in the logs.)

> **Note:** Some of this resource's arguments have default values that come from the AWS Provider. Other default values are provided by AWS and subject to change without notice. When relying on AWS defaults, the provider state will often have a zero value. For example, the AWS Provider does not provide a default for `cdc_max_batch_interval` but the AWS default is `60` (seconds). However, the provider state will show `0` since this is the value return by AWS when no value is present. Below, we aim to flag the defaults that come from AWS (_e.g._, "AWS default...").

{{% examples %}}
## Example Usage
{{% example %}}
### Minimal Configuration

This is the minimal configuration for an `aws.dms.S3Endpoint`. This endpoint will rely on the AWS Provider and AWS defaults.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dms.S3Endpoint("example", {
    endpointId: "donnedtipi",
    endpointType: "target",
    bucketName: "beckut_name",
    serviceAccessRoleArn: aws_iam_role.example.arn,
}, {
    dependsOn: [aws_iam_role_policy.example],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.dms.S3Endpoint("example",
    endpoint_id="donnedtipi",
    endpoint_type="target",
    bucket_name="beckut_name",
    service_access_role_arn=aws_iam_role["example"]["arn"],
    opts=pulumi.ResourceOptions(depends_on=[aws_iam_role_policy["example"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Dms.S3Endpoint("example", new()
    {
        EndpointId = "donnedtipi",
        EndpointType = "target",
        BucketName = "beckut_name",
        ServiceAccessRoleArn = aws_iam_role.Example.Arn,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_iam_role_policy.Example,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.NewS3Endpoint(ctx, "example", &dms.S3EndpointArgs{
			EndpointId:           pulumi.String("donnedtipi"),
			EndpointType:         pulumi.String("target"),
			BucketName:           pulumi.String("beckut_name"),
			ServiceAccessRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_iam_role_policy.Example,
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
import com.pulumi.aws.dms.S3Endpoint;
import com.pulumi.aws.dms.S3EndpointArgs;
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
        var example = new S3Endpoint("example", S3EndpointArgs.builder()        
            .endpointId("donnedtipi")
            .endpointType("target")
            .bucketName("beckut_name")
            .serviceAccessRoleArn(aws_iam_role.example().arn())
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_iam_role_policy.example())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:dms:S3Endpoint
    properties:
      endpointId: donnedtipi
      endpointType: target
      bucketName: beckut_name
      serviceAccessRoleArn: ${aws_iam_role.example.arn}
    options:
      dependson:
        - ${aws_iam_role_policy.example}
```
{{% /example %}}
{{% example %}}
### Complete Configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dms.S3Endpoint("example", {
    endpointId: "donnedtipi",
    endpointType: "target",
    sslMode: "none",
    tags: {
        Name: "donnedtipi",
        Update: "to-update",
        Remove: "to-remove",
    },
    addColumnName: true,
    addTrailingPaddingCharacter: false,
    bucketFolder: "folder",
    bucketName: "bucket_name",
    cannedAclForObjects: "private",
    cdcInsertsAndUpdates: true,
    cdcInsertsOnly: false,
    cdcMaxBatchInterval: 100,
    cdcMinFileSize: 16,
    cdcPath: "cdc/path",
    compressionType: "GZIP",
    csvDelimiter: ";",
    csvNoSupValue: "x",
    csvNullValue: "?",
    csvRowDelimiter: "\\r\\n",
    dataFormat: "parquet",
    dataPageSize: 1100000,
    datePartitionDelimiter: "UNDERSCORE",
    datePartitionEnabled: true,
    datePartitionSequence: "yyyymmddhh",
    datePartitionTimezone: "Asia/Seoul",
    dictPageSizeLimit: 1000000,
    enableStatistics: false,
    encodingType: "plain",
    encryptionMode: "SSE_S3",
    expectedBucketOwner: data.aws_caller_identity.current.account_id,
    externalTableDefinition: "etd",
    ignoreHeaderRows: 1,
    includeOpForFullLoad: true,
    maxFileSize: 1000000,
    parquetTimestampInMillisecond: true,
    parquetVersion: "parquet-2-0",
    preserveTransactions: false,
    rfc4180: false,
    rowGroupLength: 11000,
    serverSideEncryptionKmsKeyId: aws_kms_key.example.arn,
    serviceAccessRoleArn: aws_iam_role.example.arn,
    timestampColumnName: "tx_commit_time",
    useCsvNoSupValue: false,
    useTaskStartTimeForFullLoadTimestamp: true,
    glueCatalogGeneration: true,
}, {
    dependsOn: [aws_iam_role_policy.example],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.dms.S3Endpoint("example",
    endpoint_id="donnedtipi",
    endpoint_type="target",
    ssl_mode="none",
    tags={
        "Name": "donnedtipi",
        "Update": "to-update",
        "Remove": "to-remove",
    },
    add_column_name=True,
    add_trailing_padding_character=False,
    bucket_folder="folder",
    bucket_name="bucket_name",
    canned_acl_for_objects="private",
    cdc_inserts_and_updates=True,
    cdc_inserts_only=False,
    cdc_max_batch_interval=100,
    cdc_min_file_size=16,
    cdc_path="cdc/path",
    compression_type="GZIP",
    csv_delimiter=";",
    csv_no_sup_value="x",
    csv_null_value="?",
    csv_row_delimiter="\\r\\n",
    data_format="parquet",
    data_page_size=1100000,
    date_partition_delimiter="UNDERSCORE",
    date_partition_enabled=True,
    date_partition_sequence="yyyymmddhh",
    date_partition_timezone="Asia/Seoul",
    dict_page_size_limit=1000000,
    enable_statistics=False,
    encoding_type="plain",
    encryption_mode="SSE_S3",
    expected_bucket_owner=data["aws_caller_identity"]["current"]["account_id"],
    external_table_definition="etd",
    ignore_header_rows=1,
    include_op_for_full_load=True,
    max_file_size=1000000,
    parquet_timestamp_in_millisecond=True,
    parquet_version="parquet-2-0",
    preserve_transactions=False,
    rfc4180=False,
    row_group_length=11000,
    server_side_encryption_kms_key_id=aws_kms_key["example"]["arn"],
    service_access_role_arn=aws_iam_role["example"]["arn"],
    timestamp_column_name="tx_commit_time",
    use_csv_no_sup_value=False,
    use_task_start_time_for_full_load_timestamp=True,
    glue_catalog_generation=True,
    opts=pulumi.ResourceOptions(depends_on=[aws_iam_role_policy["example"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Dms.S3Endpoint("example", new()
    {
        EndpointId = "donnedtipi",
        EndpointType = "target",
        SslMode = "none",
        Tags = 
        {
            { "Name", "donnedtipi" },
            { "Update", "to-update" },
            { "Remove", "to-remove" },
        },
        AddColumnName = true,
        AddTrailingPaddingCharacter = false,
        BucketFolder = "folder",
        BucketName = "bucket_name",
        CannedAclForObjects = "private",
        CdcInsertsAndUpdates = true,
        CdcInsertsOnly = false,
        CdcMaxBatchInterval = 100,
        CdcMinFileSize = 16,
        CdcPath = "cdc/path",
        CompressionType = "GZIP",
        CsvDelimiter = ";",
        CsvNoSupValue = "x",
        CsvNullValue = "?",
        CsvRowDelimiter = "\\r\\n",
        DataFormat = "parquet",
        DataPageSize = 1100000,
        DatePartitionDelimiter = "UNDERSCORE",
        DatePartitionEnabled = true,
        DatePartitionSequence = "yyyymmddhh",
        DatePartitionTimezone = "Asia/Seoul",
        DictPageSizeLimit = 1000000,
        EnableStatistics = false,
        EncodingType = "plain",
        EncryptionMode = "SSE_S3",
        ExpectedBucketOwner = data.Aws_caller_identity.Current.Account_id,
        ExternalTableDefinition = "etd",
        IgnoreHeaderRows = 1,
        IncludeOpForFullLoad = true,
        MaxFileSize = 1000000,
        ParquetTimestampInMillisecond = true,
        ParquetVersion = "parquet-2-0",
        PreserveTransactions = false,
        Rfc4180 = false,
        RowGroupLength = 11000,
        ServerSideEncryptionKmsKeyId = aws_kms_key.Example.Arn,
        ServiceAccessRoleArn = aws_iam_role.Example.Arn,
        TimestampColumnName = "tx_commit_time",
        UseCsvNoSupValue = false,
        UseTaskStartTimeForFullLoadTimestamp = true,
        GlueCatalogGeneration = true,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_iam_role_policy.Example,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dms.NewS3Endpoint(ctx, "example", &dms.S3EndpointArgs{
			EndpointId:   pulumi.String("donnedtipi"),
			EndpointType: pulumi.String("target"),
			SslMode:      pulumi.String("none"),
			Tags: pulumi.StringMap{
				"Name":   pulumi.String("donnedtipi"),
				"Update": pulumi.String("to-update"),
				"Remove": pulumi.String("to-remove"),
			},
			AddColumnName:                        pulumi.Bool(true),
			AddTrailingPaddingCharacter:          pulumi.Bool(false),
			BucketFolder:                         pulumi.String("folder"),
			BucketName:                           pulumi.String("bucket_name"),
			CannedAclForObjects:                  pulumi.String("private"),
			CdcInsertsAndUpdates:                 pulumi.Bool(true),
			CdcInsertsOnly:                       pulumi.Bool(false),
			CdcMaxBatchInterval:                  pulumi.Int(100),
			CdcMinFileSize:                       pulumi.Int(16),
			CdcPath:                              pulumi.String("cdc/path"),
			CompressionType:                      pulumi.String("GZIP"),
			CsvDelimiter:                         pulumi.String(";"),
			CsvNoSupValue:                        pulumi.String("x"),
			CsvNullValue:                         pulumi.String("?"),
			CsvRowDelimiter:                      pulumi.String("\\r\\n"),
			DataFormat:                           pulumi.String("parquet"),
			DataPageSize:                         pulumi.Int(1100000),
			DatePartitionDelimiter:               pulumi.String("UNDERSCORE"),
			DatePartitionEnabled:                 pulumi.Bool(true),
			DatePartitionSequence:                pulumi.String("yyyymmddhh"),
			DatePartitionTimezone:                pulumi.String("Asia/Seoul"),
			DictPageSizeLimit:                    pulumi.Int(1000000),
			EnableStatistics:                     pulumi.Bool(false),
			EncodingType:                         pulumi.String("plain"),
			EncryptionMode:                       pulumi.String("SSE_S3"),
			ExpectedBucketOwner:                  pulumi.Any(data.Aws_caller_identity.Current.Account_id),
			ExternalTableDefinition:              pulumi.String("etd"),
			IgnoreHeaderRows:                     pulumi.Int(1),
			IncludeOpForFullLoad:                 pulumi.Bool(true),
			MaxFileSize:                          pulumi.Int(1000000),
			ParquetTimestampInMillisecond:        pulumi.Bool(true),
			ParquetVersion:                       pulumi.String("parquet-2-0"),
			PreserveTransactions:                 pulumi.Bool(false),
			Rfc4180:                              pulumi.Bool(false),
			RowGroupLength:                       pulumi.Int(11000),
			ServerSideEncryptionKmsKeyId:         pulumi.Any(aws_kms_key.Example.Arn),
			ServiceAccessRoleArn:                 pulumi.Any(aws_iam_role.Example.Arn),
			TimestampColumnName:                  pulumi.String("tx_commit_time"),
			UseCsvNoSupValue:                     pulumi.Bool(false),
			UseTaskStartTimeForFullLoadTimestamp: pulumi.Bool(true),
			GlueCatalogGeneration:                pulumi.Bool(true),
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_iam_role_policy.Example,
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
import com.pulumi.aws.dms.S3Endpoint;
import com.pulumi.aws.dms.S3EndpointArgs;
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
        var example = new S3Endpoint("example", S3EndpointArgs.builder()        
            .endpointId("donnedtipi")
            .endpointType("target")
            .sslMode("none")
            .tags(Map.ofEntries(
                Map.entry("Name", "donnedtipi"),
                Map.entry("Update", "to-update"),
                Map.entry("Remove", "to-remove")
            ))
            .addColumnName(true)
            .addTrailingPaddingCharacter(false)
            .bucketFolder("folder")
            .bucketName("bucket_name")
            .cannedAclForObjects("private")
            .cdcInsertsAndUpdates(true)
            .cdcInsertsOnly(false)
            .cdcMaxBatchInterval(100)
            .cdcMinFileSize(16)
            .cdcPath("cdc/path")
            .compressionType("GZIP")
            .csvDelimiter(";")
            .csvNoSupValue("x")
            .csvNullValue("?")
            .csvRowDelimiter("\\r\\n")
            .dataFormat("parquet")
            .dataPageSize(1100000)
            .datePartitionDelimiter("UNDERSCORE")
            .datePartitionEnabled(true)
            .datePartitionSequence("yyyymmddhh")
            .datePartitionTimezone("Asia/Seoul")
            .dictPageSizeLimit(1000000)
            .enableStatistics(false)
            .encodingType("plain")
            .encryptionMode("SSE_S3")
            .expectedBucketOwner(data.aws_caller_identity().current().account_id())
            .externalTableDefinition("etd")
            .ignoreHeaderRows(1)
            .includeOpForFullLoad(true)
            .maxFileSize(1000000)
            .parquetTimestampInMillisecond(true)
            .parquetVersion("parquet-2-0")
            .preserveTransactions(false)
            .rfc4180(false)
            .rowGroupLength(11000)
            .serverSideEncryptionKmsKeyId(aws_kms_key.example().arn())
            .serviceAccessRoleArn(aws_iam_role.example().arn())
            .timestampColumnName("tx_commit_time")
            .useCsvNoSupValue(false)
            .useTaskStartTimeForFullLoadTimestamp(true)
            .glueCatalogGeneration(true)
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_iam_role_policy.example())
                .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:dms:S3Endpoint
    properties:
      endpointId: donnedtipi
      endpointType: target
      sslMode: none
      tags:
        Name: donnedtipi
        Update: to-update
        Remove: to-remove
      addColumnName: true
      addTrailingPaddingCharacter: false
      bucketFolder: folder
      bucketName: bucket_name
      cannedAclForObjects: private
      cdcInsertsAndUpdates: true
      cdcInsertsOnly: false
      cdcMaxBatchInterval: 100
      cdcMinFileSize: 16
      cdcPath: cdc/path
      compressionType: GZIP
      csvDelimiter: ;
      csvNoSupValue: x
      csvNullValue: '?'
      csvRowDelimiter: \r\n
      dataFormat: parquet
      dataPageSize: 1.1e+06
      datePartitionDelimiter: UNDERSCORE
      datePartitionEnabled: true
      datePartitionSequence: yyyymmddhh
      datePartitionTimezone: Asia/Seoul
      dictPageSizeLimit: 1e+06
      enableStatistics: false
      encodingType: plain
      encryptionMode: SSE_S3
      expectedBucketOwner: ${data.aws_caller_identity.current.account_id}
      externalTableDefinition: etd
      ignoreHeaderRows: 1
      includeOpForFullLoad: true
      maxFileSize: 1e+06
      parquetTimestampInMillisecond: true
      parquetVersion: parquet-2-0
      preserveTransactions: false
      rfc4180: false
      rowGroupLength: 11000
      serverSideEncryptionKmsKeyId: ${aws_kms_key.example.arn}
      serviceAccessRoleArn: ${aws_iam_role.example.arn}
      timestampColumnName: tx_commit_time
      useCsvNoSupValue: false
      useTaskStartTimeForFullLoadTimestamp: true
      glueCatalogGeneration: true
    options:
      dependson:
        - ${aws_iam_role_policy.example}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import endpoints using the `endpoint_id`. For example:

```sh
 $ pulumi import aws:dms/s3Endpoint:S3Endpoint example example-dms-endpoint-tf
```
 