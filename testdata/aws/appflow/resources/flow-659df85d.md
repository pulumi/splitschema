Provides an AppFlow flow resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleSourceBucketV2 = new aws.s3.BucketV2("exampleSourceBucketV2", {});
const exampleSourcePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        sid: "AllowAppFlowSourceActions",
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["appflow.amazonaws.com"],
        }],
        actions: [
            "s3:ListBucket",
            "s3:GetObject",
        ],
        resources: [
            "arn:aws:s3:::example-source",
            "arn:aws:s3:::example-source/*",
        ],
    }],
});
const exampleSourceBucketPolicy = new aws.s3.BucketPolicy("exampleSourceBucketPolicy", {
    bucket: exampleSourceBucketV2.id,
    policy: exampleSourcePolicyDocument.then(exampleSourcePolicyDocument => exampleSourcePolicyDocument.json),
});
const exampleBucketObjectv2 = new aws.s3.BucketObjectv2("exampleBucketObjectv2", {
    bucket: exampleSourceBucketV2.id,
    key: "example_source.csv",
    source: new pulumi.asset.FileAsset("example_source.csv"),
});
const exampleDestinationBucketV2 = new aws.s3.BucketV2("exampleDestinationBucketV2", {});
const exampleDestinationPolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        sid: "AllowAppFlowDestinationActions",
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["appflow.amazonaws.com"],
        }],
        actions: [
            "s3:PutObject",
            "s3:AbortMultipartUpload",
            "s3:ListMultipartUploadParts",
            "s3:ListBucketMultipartUploads",
            "s3:GetBucketAcl",
            "s3:PutObjectAcl",
        ],
        resources: [
            "arn:aws:s3:::example-destination",
            "arn:aws:s3:::example-destination/*",
        ],
    }],
});
const exampleDestinationBucketPolicy = new aws.s3.BucketPolicy("exampleDestinationBucketPolicy", {
    bucket: exampleDestinationBucketV2.id,
    policy: exampleDestinationPolicyDocument.then(exampleDestinationPolicyDocument => exampleDestinationPolicyDocument.json),
});
const exampleFlow = new aws.appflow.Flow("exampleFlow", {
    sourceFlowConfig: {
        connectorType: "S3",
        sourceConnectorProperties: {
            s3: {
                bucketName: exampleSourceBucketPolicy.bucket,
                bucketPrefix: "example",
            },
        },
    },
    destinationFlowConfigs: [{
        connectorType: "S3",
        destinationConnectorProperties: {
            s3: {
                bucketName: exampleDestinationBucketPolicy.bucket,
                s3OutputFormatConfig: {
                    prefixConfig: {
                        prefixType: "PATH",
                    },
                },
            },
        },
    }],
    tasks: [{
        sourceFields: ["exampleField"],
        destinationField: "exampleField",
        taskType: "Map",
        connectorOperators: [{
            s3: "NO_OP",
        }],
    }],
    triggerConfig: {
        triggerType: "OnDemand",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_source_bucket_v2 = aws.s3.BucketV2("exampleSourceBucketV2")
example_source_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="AllowAppFlowSourceActions",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["appflow.amazonaws.com"],
    )],
    actions=[
        "s3:ListBucket",
        "s3:GetObject",
    ],
    resources=[
        "arn:aws:s3:::example-source",
        "arn:aws:s3:::example-source/*",
    ],
)])
example_source_bucket_policy = aws.s3.BucketPolicy("exampleSourceBucketPolicy",
    bucket=example_source_bucket_v2.id,
    policy=example_source_policy_document.json)
example_bucket_objectv2 = aws.s3.BucketObjectv2("exampleBucketObjectv2",
    bucket=example_source_bucket_v2.id,
    key="example_source.csv",
    source=pulumi.FileAsset("example_source.csv"))
example_destination_bucket_v2 = aws.s3.BucketV2("exampleDestinationBucketV2")
example_destination_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    sid="AllowAppFlowDestinationActions",
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["appflow.amazonaws.com"],
    )],
    actions=[
        "s3:PutObject",
        "s3:AbortMultipartUpload",
        "s3:ListMultipartUploadParts",
        "s3:ListBucketMultipartUploads",
        "s3:GetBucketAcl",
        "s3:PutObjectAcl",
    ],
    resources=[
        "arn:aws:s3:::example-destination",
        "arn:aws:s3:::example-destination/*",
    ],
)])
example_destination_bucket_policy = aws.s3.BucketPolicy("exampleDestinationBucketPolicy",
    bucket=example_destination_bucket_v2.id,
    policy=example_destination_policy_document.json)
example_flow = aws.appflow.Flow("exampleFlow",
    source_flow_config=aws.appflow.FlowSourceFlowConfigArgs(
        connector_type="S3",
        source_connector_properties=aws.appflow.FlowSourceFlowConfigSourceConnectorPropertiesArgs(
            s3=aws.appflow.FlowSourceFlowConfigSourceConnectorPropertiesS3Args(
                bucket_name=example_source_bucket_policy.bucket,
                bucket_prefix="example",
            ),
        ),
    ),
    destination_flow_configs=[aws.appflow.FlowDestinationFlowConfigArgs(
        connector_type="S3",
        destination_connector_properties=aws.appflow.FlowDestinationFlowConfigDestinationConnectorPropertiesArgs(
            s3=aws.appflow.FlowDestinationFlowConfigDestinationConnectorPropertiesS3Args(
                bucket_name=example_destination_bucket_policy.bucket,
                s3_output_format_config=aws.appflow.FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigArgs(
                    prefix_config=aws.appflow.FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigArgs(
                        prefix_type="PATH",
                    ),
                ),
            ),
        ),
    )],
    tasks=[aws.appflow.FlowTaskArgs(
        source_fields=["exampleField"],
        destination_field="exampleField",
        task_type="Map",
        connector_operators=[aws.appflow.FlowTaskConnectorOperatorArgs(
            s3="NO_OP",
        )],
    )],
    trigger_config=aws.appflow.FlowTriggerConfigArgs(
        trigger_type="OnDemand",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleSourceBucketV2 = new Aws.S3.BucketV2("exampleSourceBucketV2");

    var exampleSourcePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "AllowAppFlowSourceActions",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "appflow.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "s3:ListBucket",
                    "s3:GetObject",
                },
                Resources = new[]
                {
                    "arn:aws:s3:::example-source",
                    "arn:aws:s3:::example-source/*",
                },
            },
        },
    });

    var exampleSourceBucketPolicy = new Aws.S3.BucketPolicy("exampleSourceBucketPolicy", new()
    {
        Bucket = exampleSourceBucketV2.Id,
        Policy = exampleSourcePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleBucketObjectv2 = new Aws.S3.BucketObjectv2("exampleBucketObjectv2", new()
    {
        Bucket = exampleSourceBucketV2.Id,
        Key = "example_source.csv",
        Source = new FileAsset("example_source.csv"),
    });

    var exampleDestinationBucketV2 = new Aws.S3.BucketV2("exampleDestinationBucketV2");

    var exampleDestinationPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Sid = "AllowAppFlowDestinationActions",
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "appflow.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "s3:PutObject",
                    "s3:AbortMultipartUpload",
                    "s3:ListMultipartUploadParts",
                    "s3:ListBucketMultipartUploads",
                    "s3:GetBucketAcl",
                    "s3:PutObjectAcl",
                },
                Resources = new[]
                {
                    "arn:aws:s3:::example-destination",
                    "arn:aws:s3:::example-destination/*",
                },
            },
        },
    });

    var exampleDestinationBucketPolicy = new Aws.S3.BucketPolicy("exampleDestinationBucketPolicy", new()
    {
        Bucket = exampleDestinationBucketV2.Id,
        Policy = exampleDestinationPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleFlow = new Aws.AppFlow.Flow("exampleFlow", new()
    {
        SourceFlowConfig = new Aws.AppFlow.Inputs.FlowSourceFlowConfigArgs
        {
            ConnectorType = "S3",
            SourceConnectorProperties = new Aws.AppFlow.Inputs.FlowSourceFlowConfigSourceConnectorPropertiesArgs
            {
                S3 = new Aws.AppFlow.Inputs.FlowSourceFlowConfigSourceConnectorPropertiesS3Args
                {
                    BucketName = exampleSourceBucketPolicy.Bucket,
                    BucketPrefix = "example",
                },
            },
        },
        DestinationFlowConfigs = new[]
        {
            new Aws.AppFlow.Inputs.FlowDestinationFlowConfigArgs
            {
                ConnectorType = "S3",
                DestinationConnectorProperties = new Aws.AppFlow.Inputs.FlowDestinationFlowConfigDestinationConnectorPropertiesArgs
                {
                    S3 = new Aws.AppFlow.Inputs.FlowDestinationFlowConfigDestinationConnectorPropertiesS3Args
                    {
                        BucketName = exampleDestinationBucketPolicy.Bucket,
                        S3OutputFormatConfig = new Aws.AppFlow.Inputs.FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigArgs
                        {
                            PrefixConfig = new Aws.AppFlow.Inputs.FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigArgs
                            {
                                PrefixType = "PATH",
                            },
                        },
                    },
                },
            },
        },
        Tasks = new[]
        {
            new Aws.AppFlow.Inputs.FlowTaskArgs
            {
                SourceFields = new[]
                {
                    "exampleField",
                },
                DestinationField = "exampleField",
                TaskType = "Map",
                ConnectorOperators = new[]
                {
                    new Aws.AppFlow.Inputs.FlowTaskConnectorOperatorArgs
                    {
                        S3 = "NO_OP",
                    },
                },
            },
        },
        TriggerConfig = new Aws.AppFlow.Inputs.FlowTriggerConfigArgs
        {
            TriggerType = "OnDemand",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appflow"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleSourceBucketV2, err := s3.NewBucketV2(ctx, "exampleSourceBucketV2", nil)
		if err != nil {
			return err
		}
		exampleSourcePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("AllowAppFlowSourceActions"),
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"appflow.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"s3:ListBucket",
						"s3:GetObject",
					},
					Resources: []string{
						"arn:aws:s3:::example-source",
						"arn:aws:s3:::example-source/*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleSourceBucketPolicy, err := s3.NewBucketPolicy(ctx, "exampleSourceBucketPolicy", &s3.BucketPolicyArgs{
			Bucket: exampleSourceBucketV2.ID(),
			Policy: *pulumi.String(exampleSourcePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObjectv2(ctx, "exampleBucketObjectv2", &s3.BucketObjectv2Args{
			Bucket: exampleSourceBucketV2.ID(),
			Key:    pulumi.String("example_source.csv"),
			Source: pulumi.NewFileAsset("example_source.csv"),
		})
		if err != nil {
			return err
		}
		exampleDestinationBucketV2, err := s3.NewBucketV2(ctx, "exampleDestinationBucketV2", nil)
		if err != nil {
			return err
		}
		exampleDestinationPolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Sid:    pulumi.StringRef("AllowAppFlowDestinationActions"),
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"appflow.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"s3:PutObject",
						"s3:AbortMultipartUpload",
						"s3:ListMultipartUploadParts",
						"s3:ListBucketMultipartUploads",
						"s3:GetBucketAcl",
						"s3:PutObjectAcl",
					},
					Resources: []string{
						"arn:aws:s3:::example-destination",
						"arn:aws:s3:::example-destination/*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleDestinationBucketPolicy, err := s3.NewBucketPolicy(ctx, "exampleDestinationBucketPolicy", &s3.BucketPolicyArgs{
			Bucket: exampleDestinationBucketV2.ID(),
			Policy: *pulumi.String(exampleDestinationPolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		_, err = appflow.NewFlow(ctx, "exampleFlow", &appflow.FlowArgs{
			SourceFlowConfig: &appflow.FlowSourceFlowConfigArgs{
				ConnectorType: pulumi.String("S3"),
				SourceConnectorProperties: &appflow.FlowSourceFlowConfigSourceConnectorPropertiesArgs{
					S3: &appflow.FlowSourceFlowConfigSourceConnectorPropertiesS3Args{
						BucketName:   exampleSourceBucketPolicy.Bucket,
						BucketPrefix: pulumi.String("example"),
					},
				},
			},
			DestinationFlowConfigs: appflow.FlowDestinationFlowConfigArray{
				&appflow.FlowDestinationFlowConfigArgs{
					ConnectorType: pulumi.String("S3"),
					DestinationConnectorProperties: &appflow.FlowDestinationFlowConfigDestinationConnectorPropertiesArgs{
						S3: &appflow.FlowDestinationFlowConfigDestinationConnectorPropertiesS3Args{
							BucketName: exampleDestinationBucketPolicy.Bucket,
							S3OutputFormatConfig: &appflow.FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigArgs{
								PrefixConfig: &appflow.FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigArgs{
									PrefixType: pulumi.String("PATH"),
								},
							},
						},
					},
				},
			},
			Tasks: appflow.FlowTaskArray{
				&appflow.FlowTaskArgs{
					SourceFields: pulumi.StringArray{
						pulumi.String("exampleField"),
					},
					DestinationField: pulumi.String("exampleField"),
					TaskType:         pulumi.String("Map"),
					ConnectorOperators: appflow.FlowTaskConnectorOperatorArray{
						&appflow.FlowTaskConnectorOperatorArgs{
							S3: pulumi.String("NO_OP"),
						},
					},
				},
			},
			TriggerConfig: &appflow.FlowTriggerConfigArgs{
				TriggerType: pulumi.String("OnDemand"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.s3.BucketPolicy;
import com.pulumi.aws.s3.BucketPolicyArgs;
import com.pulumi.aws.s3.BucketObjectv2;
import com.pulumi.aws.s3.BucketObjectv2Args;
import com.pulumi.aws.appflow.Flow;
import com.pulumi.aws.appflow.FlowArgs;
import com.pulumi.aws.appflow.inputs.FlowSourceFlowConfigArgs;
import com.pulumi.aws.appflow.inputs.FlowSourceFlowConfigSourceConnectorPropertiesArgs;
import com.pulumi.aws.appflow.inputs.FlowSourceFlowConfigSourceConnectorPropertiesS3Args;
import com.pulumi.aws.appflow.inputs.FlowDestinationFlowConfigArgs;
import com.pulumi.aws.appflow.inputs.FlowDestinationFlowConfigDestinationConnectorPropertiesArgs;
import com.pulumi.aws.appflow.inputs.FlowDestinationFlowConfigDestinationConnectorPropertiesS3Args;
import com.pulumi.aws.appflow.inputs.FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigArgs;
import com.pulumi.aws.appflow.inputs.FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigArgs;
import com.pulumi.aws.appflow.inputs.FlowTaskArgs;
import com.pulumi.aws.appflow.inputs.FlowTriggerConfigArgs;
import com.pulumi.asset.FileAsset;
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
        var exampleSourceBucketV2 = new BucketV2("exampleSourceBucketV2");

        final var exampleSourcePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("AllowAppFlowSourceActions")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("appflow.amazonaws.com")
                    .build())
                .actions(                
                    "s3:ListBucket",
                    "s3:GetObject")
                .resources(                
                    "arn:aws:s3:::example-source",
                    "arn:aws:s3:::example-source/*")
                .build())
            .build());

        var exampleSourceBucketPolicy = new BucketPolicy("exampleSourceBucketPolicy", BucketPolicyArgs.builder()        
            .bucket(exampleSourceBucketV2.id())
            .policy(exampleSourcePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleBucketObjectv2 = new BucketObjectv2("exampleBucketObjectv2", BucketObjectv2Args.builder()        
            .bucket(exampleSourceBucketV2.id())
            .key("example_source.csv")
            .source(new FileAsset("example_source.csv"))
            .build());

        var exampleDestinationBucketV2 = new BucketV2("exampleDestinationBucketV2");

        final var exampleDestinationPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .sid("AllowAppFlowDestinationActions")
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("appflow.amazonaws.com")
                    .build())
                .actions(                
                    "s3:PutObject",
                    "s3:AbortMultipartUpload",
                    "s3:ListMultipartUploadParts",
                    "s3:ListBucketMultipartUploads",
                    "s3:GetBucketAcl",
                    "s3:PutObjectAcl")
                .resources(                
                    "arn:aws:s3:::example-destination",
                    "arn:aws:s3:::example-destination/*")
                .build())
            .build());

        var exampleDestinationBucketPolicy = new BucketPolicy("exampleDestinationBucketPolicy", BucketPolicyArgs.builder()        
            .bucket(exampleDestinationBucketV2.id())
            .policy(exampleDestinationPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var exampleFlow = new Flow("exampleFlow", FlowArgs.builder()        
            .sourceFlowConfig(FlowSourceFlowConfigArgs.builder()
                .connectorType("S3")
                .sourceConnectorProperties(FlowSourceFlowConfigSourceConnectorPropertiesArgs.builder()
                    .s3(FlowSourceFlowConfigSourceConnectorPropertiesS3Args.builder()
                        .bucketName(exampleSourceBucketPolicy.bucket())
                        .bucketPrefix("example")
                        .build())
                    .build())
                .build())
            .destinationFlowConfigs(FlowDestinationFlowConfigArgs.builder()
                .connectorType("S3")
                .destinationConnectorProperties(FlowDestinationFlowConfigDestinationConnectorPropertiesArgs.builder()
                    .s3(FlowDestinationFlowConfigDestinationConnectorPropertiesS3Args.builder()
                        .bucketName(exampleDestinationBucketPolicy.bucket())
                        .s3OutputFormatConfig(FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigArgs.builder()
                            .prefixConfig(FlowDestinationFlowConfigDestinationConnectorPropertiesS3S3OutputFormatConfigPrefixConfigArgs.builder()
                                .prefixType("PATH")
                                .build())
                            .build())
                        .build())
                    .build())
                .build())
            .tasks(FlowTaskArgs.builder()
                .sourceFields("exampleField")
                .destinationField("exampleField")
                .taskType("Map")
                .connectorOperators(FlowTaskConnectorOperatorArgs.builder()
                    .s3("NO_OP")
                    .build())
                .build())
            .triggerConfig(FlowTriggerConfigArgs.builder()
                .triggerType("OnDemand")
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleSourceBucketV2:
    type: aws:s3:BucketV2
  exampleSourceBucketPolicy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${exampleSourceBucketV2.id}
      policy: ${exampleSourcePolicyDocument.json}
  exampleBucketObjectv2:
    type: aws:s3:BucketObjectv2
    properties:
      bucket: ${exampleSourceBucketV2.id}
      key: example_source.csv
      source:
        fn::FileAsset: example_source.csv
  exampleDestinationBucketV2:
    type: aws:s3:BucketV2
  exampleDestinationBucketPolicy:
    type: aws:s3:BucketPolicy
    properties:
      bucket: ${exampleDestinationBucketV2.id}
      policy: ${exampleDestinationPolicyDocument.json}
  exampleFlow:
    type: aws:appflow:Flow
    properties:
      sourceFlowConfig:
        connectorType: S3
        sourceConnectorProperties:
          s3:
            bucketName: ${exampleSourceBucketPolicy.bucket}
            bucketPrefix: example
      destinationFlowConfigs:
        - connectorType: S3
          destinationConnectorProperties:
            s3:
              bucketName: ${exampleDestinationBucketPolicy.bucket}
              s3OutputFormatConfig:
                prefixConfig:
                  prefixType: PATH
      tasks:
        - sourceFields:
            - exampleField
          destinationField: exampleField
          taskType: Map
          connectorOperators:
            - s3: NO_OP
      triggerConfig:
        triggerType: OnDemand
variables:
  exampleSourcePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: AllowAppFlowSourceActions
            effect: Allow
            principals:
              - type: Service
                identifiers:
                  - appflow.amazonaws.com
            actions:
              - s3:ListBucket
              - s3:GetObject
            resources:
              - arn:aws:s3:::example-source
              - arn:aws:s3:::example-source/*
  exampleDestinationPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - sid: AllowAppFlowDestinationActions
            effect: Allow
            principals:
              - type: Service
                identifiers:
                  - appflow.amazonaws.com
            actions:
              - s3:PutObject
              - s3:AbortMultipartUpload
              - s3:ListMultipartUploadParts
              - s3:ListBucketMultipartUploads
              - s3:GetBucketAcl
              - s3:PutObjectAcl
            resources:
              - arn:aws:s3:::example-destination
              - arn:aws:s3:::example-destination/*
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AppFlow flows using the `arn`. For example:

```sh
 $ pulumi import aws:appflow/flow:Flow example arn:aws:appflow:us-west-2:123456789012:flow/example-flow
```
 