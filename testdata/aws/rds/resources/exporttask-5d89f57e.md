Resource for managing an AWS RDS (Relational Database) Export Task.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.ExportTask("example", {
    exportTaskIdentifier: "example",
    sourceArn: aws_db_snapshot.example.db_snapshot_arn,
    s3BucketName: aws_s3_bucket.example.id,
    iamRoleArn: aws_iam_role.example.arn,
    kmsKeyId: aws_kms_key.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.ExportTask("example",
    export_task_identifier="example",
    source_arn=aws_db_snapshot["example"]["db_snapshot_arn"],
    s3_bucket_name=aws_s3_bucket["example"]["id"],
    iam_role_arn=aws_iam_role["example"]["arn"],
    kms_key_id=aws_kms_key["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Rds.ExportTask("example", new()
    {
        ExportTaskIdentifier = "example",
        SourceArn = aws_db_snapshot.Example.Db_snapshot_arn,
        S3BucketName = aws_s3_bucket.Example.Id,
        IamRoleArn = aws_iam_role.Example.Arn,
        KmsKeyId = aws_kms_key.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.NewExportTask(ctx, "example", &rds.ExportTaskArgs{
			ExportTaskIdentifier: pulumi.String("example"),
			SourceArn:            pulumi.Any(aws_db_snapshot.Example.Db_snapshot_arn),
			S3BucketName:         pulumi.Any(aws_s3_bucket.Example.Id),
			IamRoleArn:           pulumi.Any(aws_iam_role.Example.Arn),
			KmsKeyId:             pulumi.Any(aws_kms_key.Example.Arn),
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
import com.pulumi.aws.rds.ExportTask;
import com.pulumi.aws.rds.ExportTaskArgs;
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
        var example = new ExportTask("example", ExportTaskArgs.builder()        
            .exportTaskIdentifier("example")
            .sourceArn(aws_db_snapshot.example().db_snapshot_arn())
            .s3BucketName(aws_s3_bucket.example().id())
            .iamRoleArn(aws_iam_role.example().arn())
            .kmsKeyId(aws_kms_key.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:rds:ExportTask
    properties:
      exportTaskIdentifier: example
      sourceArn: ${aws_db_snapshot.example.db_snapshot_arn}
      s3BucketName: ${aws_s3_bucket.example.id}
      iamRoleArn: ${aws_iam_role.example.arn}
      kmsKeyId: ${aws_kms_key.example.arn}
```
{{% /example %}}
{{% example %}}
### Complete Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {forceDestroy: true});
const exampleBucketAclV2 = new aws.s3.BucketAclV2("exampleBucketAclV2", {
    bucket: exampleBucketV2.id,
    acl: "private",
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: JSON.stringify({
    Version: "2012-10-17",
    Statement: [{
        Action: "sts:AssumeRole",
        Effect: "Allow",
        Sid: "",
        Principal: {
            Service: "export.rds.amazonaws.com",
        },
    }],
})});
const examplePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            actions: ["s3:ListAllMyBuckets"],
            resources: ["*"],
        },
        {
            actions: [
                "s3:GetBucketLocation",
                "s3:ListBucket",
            ],
            resources: [exampleBucketV2.arn],
        },
        {
            actions: [
                "s3:GetObject",
                "s3:PutObject",
                "s3:DeleteObject",
            ],
            resources: [pulumi.interpolate`${exampleBucketV2.arn}/*`],
        },
    ],
});
const examplePolicy = new aws.iam.Policy("examplePolicy", {policy: examplePolicyDocument.apply(examplePolicyDocument => examplePolicyDocument.json)});
const exampleRolePolicyAttachment = new aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment", {
    role: exampleRole.name,
    policyArn: examplePolicy.arn,
});
const exampleKey = new aws.kms.Key("exampleKey", {deletionWindowInDays: 10});
const exampleInstance = new aws.rds.Instance("exampleInstance", {
    identifier: "example",
    allocatedStorage: 10,
    dbName: "test",
    engine: "mysql",
    engineVersion: "5.7",
    instanceClass: "db.t3.micro",
    username: "foo",
    password: "foobarbaz",
    parameterGroupName: "default.mysql5.7",
    skipFinalSnapshot: true,
});
const exampleSnapshot = new aws.rds.Snapshot("exampleSnapshot", {
    dbInstanceIdentifier: exampleInstance.identifier,
    dbSnapshotIdentifier: "example",
});
const exampleExportTask = new aws.rds.ExportTask("exampleExportTask", {
    exportTaskIdentifier: "example",
    sourceArn: exampleSnapshot.dbSnapshotArn,
    s3BucketName: exampleBucketV2.id,
    iamRoleArn: exampleRole.arn,
    kmsKeyId: exampleKey.arn,
    exportOnlies: ["database"],
    s3Prefix: "my_prefix/example",
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2", force_destroy=True)
example_bucket_acl_v2 = aws.s3.BucketAclV2("exampleBucketAclV2",
    bucket=example_bucket_v2.id,
    acl="private")
example_role = aws.iam.Role("exampleRole", assume_role_policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Sid": "",
        "Principal": {
            "Service": "export.rds.amazonaws.com",
        },
    }],
}))
example_policy_document = aws.iam.get_policy_document_output(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        actions=["s3:ListAllMyBuckets"],
        resources=["*"],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        actions=[
            "s3:GetBucketLocation",
            "s3:ListBucket",
        ],
        resources=[example_bucket_v2.arn],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        actions=[
            "s3:GetObject",
            "s3:PutObject",
            "s3:DeleteObject",
        ],
        resources=[example_bucket_v2.arn.apply(lambda arn: f"{arn}/*")],
    ),
])
example_policy = aws.iam.Policy("examplePolicy", policy=example_policy_document.json)
example_role_policy_attachment = aws.iam.RolePolicyAttachment("exampleRolePolicyAttachment",
    role=example_role.name,
    policy_arn=example_policy.arn)
example_key = aws.kms.Key("exampleKey", deletion_window_in_days=10)
example_instance = aws.rds.Instance("exampleInstance",
    identifier="example",
    allocated_storage=10,
    db_name="test",
    engine="mysql",
    engine_version="5.7",
    instance_class="db.t3.micro",
    username="foo",
    password="foobarbaz",
    parameter_group_name="default.mysql5.7",
    skip_final_snapshot=True)
example_snapshot = aws.rds.Snapshot("exampleSnapshot",
    db_instance_identifier=example_instance.identifier,
    db_snapshot_identifier="example")
example_export_task = aws.rds.ExportTask("exampleExportTask",
    export_task_identifier="example",
    source_arn=example_snapshot.db_snapshot_arn,
    s3_bucket_name=example_bucket_v2.id,
    iam_role_arn=example_role.arn,
    kms_key_id=example_key.arn,
    export_onlies=["database"],
    s3_prefix="my_prefix/example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2", new()
    {
        ForceDestroy = true,
    });

    var exampleBucketAclV2 = new Aws.S3.BucketAclV2("exampleBucketAclV2", new()
    {
        Bucket = exampleBucketV2.Id,
        Acl = "private",
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = "sts:AssumeRole",
                    ["Effect"] = "Allow",
                    ["Sid"] = "",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["Service"] = "export.rds.amazonaws.com",
                    },
                },
            },
        }),
    });

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "s3:ListAllMyBuckets",
                },
                Resources = new[]
                {
                    "*",
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "s3:GetBucketLocation",
                    "s3:ListBucket",
                },
                Resources = new[]
                {
                    exampleBucketV2.Arn,
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "s3:GetObject",
                    "s3:PutObject",
                    "s3:DeleteObject",
                },
                Resources = new[]
                {
                    $"{exampleBucketV2.Arn}/*",
                },
            },
        },
    });

    var examplePolicy = new Aws.Iam.Policy("examplePolicy", new()
    {
        PolicyDocument = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("exampleRolePolicyAttachment", new()
    {
        Role = exampleRole.Name,
        PolicyArn = examplePolicy.Arn,
    });

    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        DeletionWindowInDays = 10,
    });

    var exampleInstance = new Aws.Rds.Instance("exampleInstance", new()
    {
        Identifier = "example",
        AllocatedStorage = 10,
        DbName = "test",
        Engine = "mysql",
        EngineVersion = "5.7",
        InstanceClass = "db.t3.micro",
        Username = "foo",
        Password = "foobarbaz",
        ParameterGroupName = "default.mysql5.7",
        SkipFinalSnapshot = true,
    });

    var exampleSnapshot = new Aws.Rds.Snapshot("exampleSnapshot", new()
    {
        DbInstanceIdentifier = exampleInstance.Identifier,
        DbSnapshotIdentifier = "example",
    });

    var exampleExportTask = new Aws.Rds.ExportTask("exampleExportTask", new()
    {
        ExportTaskIdentifier = "example",
        SourceArn = exampleSnapshot.DbSnapshotArn,
        S3BucketName = exampleBucketV2.Id,
        IamRoleArn = exampleRole.Arn,
        KmsKeyId = exampleKey.Arn,
        ExportOnlies = new[]
        {
            "database",
        },
        S3Prefix = "my_prefix/example",
    });

});
```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", &s3.BucketV2Args{
			ForceDestroy: pulumi.Bool(true),
		})
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
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Sid":    "",
					"Principal": map[string]interface{}{
						"Service": "export.rds.amazonaws.com",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		examplePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Actions: pulumi.StringArray{
						pulumi.String("s3:ListAllMyBuckets"),
					},
					Resources: pulumi.StringArray{
						pulumi.String("*"),
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Actions: pulumi.StringArray{
						pulumi.String("s3:GetBucketLocation"),
						pulumi.String("s3:ListBucket"),
					},
					Resources: pulumi.StringArray{
						exampleBucketV2.Arn,
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Actions: pulumi.StringArray{
						pulumi.String("s3:GetObject"),
						pulumi.String("s3:PutObject"),
						pulumi.String("s3:DeleteObject"),
					},
					Resources: pulumi.StringArray{
						exampleBucketV2.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
				},
			},
		}, nil)
		examplePolicy, err := iam.NewPolicy(ctx, "examplePolicy", &iam.PolicyArgs{
			Policy: examplePolicyDocument.ApplyT(func(examplePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &examplePolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "exampleRolePolicyAttachment", &iam.RolePolicyAttachmentArgs{
			Role:      exampleRole.Name,
			PolicyArn: examplePolicy.Arn,
		})
		if err != nil {
			return err
		}
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			DeletionWindowInDays: pulumi.Int(10),
		})
		if err != nil {
			return err
		}
		exampleInstance, err := rds.NewInstance(ctx, "exampleInstance", &rds.InstanceArgs{
			Identifier:         pulumi.String("example"),
			AllocatedStorage:   pulumi.Int(10),
			DbName:             pulumi.String("test"),
			Engine:             pulumi.String("mysql"),
			EngineVersion:      pulumi.String("5.7"),
			InstanceClass:      pulumi.String("db.t3.micro"),
			Username:           pulumi.String("foo"),
			Password:           pulumi.String("foobarbaz"),
			ParameterGroupName: pulumi.String("default.mysql5.7"),
			SkipFinalSnapshot:  pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		exampleSnapshot, err := rds.NewSnapshot(ctx, "exampleSnapshot", &rds.SnapshotArgs{
			DbInstanceIdentifier: exampleInstance.Identifier,
			DbSnapshotIdentifier: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		_, err = rds.NewExportTask(ctx, "exampleExportTask", &rds.ExportTaskArgs{
			ExportTaskIdentifier: pulumi.String("example"),
			SourceArn:            exampleSnapshot.DbSnapshotArn,
			S3BucketName:         exampleBucketV2.ID(),
			IamRoleArn:           exampleRole.Arn,
			KmsKeyId:             exampleKey.Arn,
			ExportOnlies: pulumi.StringArray{
				pulumi.String("database"),
			},
			S3Prefix: pulumi.String("my_prefix/example"),
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
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Policy;
import com.pulumi.aws.iam.PolicyArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.rds.Instance;
import com.pulumi.aws.rds.InstanceArgs;
import com.pulumi.aws.rds.Snapshot;
import com.pulumi.aws.rds.SnapshotArgs;
import com.pulumi.aws.rds.ExportTask;
import com.pulumi.aws.rds.ExportTaskArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var exampleBucketV2 = new BucketV2("exampleBucketV2", BucketV2Args.builder()        
            .forceDestroy(true)
            .build());

        var exampleBucketAclV2 = new BucketAclV2("exampleBucketAclV2", BucketAclV2Args.builder()        
            .bucket(exampleBucketV2.id())
            .acl("private")
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", "sts:AssumeRole"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Sid", ""),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "export.rds.amazonaws.com")
                        ))
                    )))
                )))
            .build());

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .actions("s3:ListAllMyBuckets")
                    .resources("*")
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .actions(                    
                        "s3:GetBucketLocation",
                        "s3:ListBucket")
                    .resources(exampleBucketV2.arn())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .actions(                    
                        "s3:GetObject",
                        "s3:PutObject",
                        "s3:DeleteObject")
                    .resources(exampleBucketV2.arn().applyValue(arn -> String.format("%s/*", arn)))
                    .build())
            .build());

        var examplePolicy = new Policy("examplePolicy", PolicyArgs.builder()        
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(examplePolicyDocument -> examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var exampleRolePolicyAttachment = new RolePolicyAttachment("exampleRolePolicyAttachment", RolePolicyAttachmentArgs.builder()        
            .role(exampleRole.name())
            .policyArn(examplePolicy.arn())
            .build());

        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .deletionWindowInDays(10)
            .build());

        var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()        
            .identifier("example")
            .allocatedStorage(10)
            .dbName("test")
            .engine("mysql")
            .engineVersion("5.7")
            .instanceClass("db.t3.micro")
            .username("foo")
            .password("foobarbaz")
            .parameterGroupName("default.mysql5.7")
            .skipFinalSnapshot(true)
            .build());

        var exampleSnapshot = new Snapshot("exampleSnapshot", SnapshotArgs.builder()        
            .dbInstanceIdentifier(exampleInstance.identifier())
            .dbSnapshotIdentifier("example")
            .build());

        var exampleExportTask = new ExportTask("exampleExportTask", ExportTaskArgs.builder()        
            .exportTaskIdentifier("example")
            .sourceArn(exampleSnapshot.dbSnapshotArn())
            .s3BucketName(exampleBucketV2.id())
            .iamRoleArn(exampleRole.arn())
            .kmsKeyId(exampleKey.arn())
            .exportOnlies("database")
            .s3Prefix("my_prefix/example")
            .build());

    }
}
```
```yaml
resources:
  exampleBucketV2:
    type: aws:s3:BucketV2
    properties:
      forceDestroy: true
  exampleBucketAclV2:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${exampleBucketV2.id}
      acl: private
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action: sts:AssumeRole
              Effect: Allow
              Sid:
              Principal:
                Service: export.rds.amazonaws.com
  examplePolicy:
    type: aws:iam:Policy
    properties:
      policy: ${examplePolicyDocument.json}
  exampleRolePolicyAttachment:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${exampleRole.name}
      policyArn: ${examplePolicy.arn}
  exampleKey:
    type: aws:kms:Key
    properties:
      deletionWindowInDays: 10
  exampleInstance:
    type: aws:rds:Instance
    properties:
      identifier: example
      allocatedStorage: 10
      dbName: test
      engine: mysql
      engineVersion: '5.7'
      instanceClass: db.t3.micro
      username: foo
      password: foobarbaz
      parameterGroupName: default.mysql5.7
      skipFinalSnapshot: true
  exampleSnapshot:
    type: aws:rds:Snapshot
    properties:
      dbInstanceIdentifier: ${exampleInstance.identifier}
      dbSnapshotIdentifier: example
  exampleExportTask:
    type: aws:rds:ExportTask
    properties:
      exportTaskIdentifier: example
      sourceArn: ${exampleSnapshot.dbSnapshotArn}
      s3BucketName: ${exampleBucketV2.id}
      iamRoleArn: ${exampleRole.arn}
      kmsKeyId: ${exampleKey.arn}
      exportOnlies:
        - database
      s3Prefix: my_prefix/example
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - s3:ListAllMyBuckets
            resources:
              - '*'
          - actions:
              - s3:GetBucketLocation
              - s3:ListBucket
            resources:
              - ${exampleBucketV2.arn}
          - actions:
              - s3:GetObject
              - s3:PutObject
              - s3:DeleteObject
            resources:
              - ${exampleBucketV2.arn}/*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a RDS (Relational Database) Export Task using the `export_task_identifier`. For example:

```sh
 $ pulumi import aws:rds/exportTask:ExportTask example example
```
 