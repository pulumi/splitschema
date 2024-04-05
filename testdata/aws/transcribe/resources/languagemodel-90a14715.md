Resource for managing an AWS Transcribe LanguageModel.

> This resource can take a significant amount of time to provision. See Language Model [FAQ](https://aws.amazon.com/transcribe/faqs/) for more details.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const examplePolicyDocument = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["transcribe.amazonaws.com"],
        }],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: examplePolicyDocument.then(examplePolicyDocument => examplePolicyDocument.json)});
const testPolicy = new aws.iam.RolePolicy("testPolicy", {
    role: exampleRole.id,
    policy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Action: [
                "s3:GetObject",
                "s3:ListBucket",
            ],
            Effect: "Allow",
            Resource: ["*"],
        }],
    }),
});
const exampleBucketV2 = new aws.s3.BucketV2("exampleBucketV2", {forceDestroy: true});
const object = new aws.s3.BucketObjectv2("object", {
    bucket: exampleBucketV2.id,
    key: "transcribe/test1.txt",
    source: new pulumi.asset.FileAsset("test1.txt"),
});
const exampleLanguageModel = new aws.transcribe.LanguageModel("exampleLanguageModel", {
    modelName: "example",
    baseModelName: "NarrowBand",
    inputDataConfig: {
        dataAccessRoleArn: exampleRole.arn,
        s3Uri: pulumi.interpolate`s3://${exampleBucketV2.id}/transcribe/`,
    },
    languageCode: "en-US",
    tags: {
        ENVIRONMENT: "development",
    },
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_policy_document = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["transcribe.amazonaws.com"],
    )],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=example_policy_document.json)
test_policy = aws.iam.RolePolicy("testPolicy",
    role=example_role.id,
    policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Action": [
                "s3:GetObject",
                "s3:ListBucket",
            ],
            "Effect": "Allow",
            "Resource": ["*"],
        }],
    }))
example_bucket_v2 = aws.s3.BucketV2("exampleBucketV2", force_destroy=True)
object = aws.s3.BucketObjectv2("object",
    bucket=example_bucket_v2.id,
    key="transcribe/test1.txt",
    source=pulumi.FileAsset("test1.txt"))
example_language_model = aws.transcribe.LanguageModel("exampleLanguageModel",
    model_name="example",
    base_model_name="NarrowBand",
    input_data_config=aws.transcribe.LanguageModelInputDataConfigArgs(
        data_access_role_arn=example_role.arn,
        s3_uri=example_bucket_v2.id.apply(lambda id: f"s3://{id}/transcribe/"),
    ),
    language_code="en-US",
    tags={
        "ENVIRONMENT": "development",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "transcribe.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var testPolicy = new Aws.Iam.RolePolicy("testPolicy", new()
    {
        Role = exampleRole.Id,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = new[]
                    {
                        "s3:GetObject",
                        "s3:ListBucket",
                    },
                    ["Effect"] = "Allow",
                    ["Resource"] = new[]
                    {
                        "*",
                    },
                },
            },
        }),
    });

    var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2", new()
    {
        ForceDestroy = true,
    });

    var @object = new Aws.S3.BucketObjectv2("object", new()
    {
        Bucket = exampleBucketV2.Id,
        Key = "transcribe/test1.txt",
        Source = new FileAsset("test1.txt"),
    });

    var exampleLanguageModel = new Aws.Transcribe.LanguageModel("exampleLanguageModel", new()
    {
        ModelName = "example",
        BaseModelName = "NarrowBand",
        InputDataConfig = new Aws.Transcribe.Inputs.LanguageModelInputDataConfigArgs
        {
            DataAccessRoleArn = exampleRole.Arn,
            S3Uri = exampleBucketV2.Id.Apply(id => $"s3://{id}/transcribe/"),
        },
        LanguageCode = "en-US",
        Tags = 
        {
            { "ENVIRONMENT", "development" },
        },
    });

});
```
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transcribe"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		examplePolicyDocument, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"transcribe.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(examplePolicyDocument.Json),
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": []string{
						"s3:GetObject",
						"s3:ListBucket",
					},
					"Effect": "Allow",
					"Resource": []string{
						"*",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = iam.NewRolePolicy(ctx, "testPolicy", &iam.RolePolicyArgs{
			Role:   exampleRole.ID(),
			Policy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		exampleBucketV2, err := s3.NewBucketV2(ctx, "exampleBucketV2", &s3.BucketV2Args{
			ForceDestroy: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketObjectv2(ctx, "object", &s3.BucketObjectv2Args{
			Bucket: exampleBucketV2.ID(),
			Key:    pulumi.String("transcribe/test1.txt"),
			Source: pulumi.NewFileAsset("test1.txt"),
		})
		if err != nil {
			return err
		}
		_, err = transcribe.NewLanguageModel(ctx, "exampleLanguageModel", &transcribe.LanguageModelArgs{
			ModelName:     pulumi.String("example"),
			BaseModelName: pulumi.String("NarrowBand"),
			InputDataConfig: &transcribe.LanguageModelInputDataConfigArgs{
				DataAccessRoleArn: exampleRole.Arn,
				S3Uri: exampleBucketV2.ID().ApplyT(func(id string) (string, error) {
					return fmt.Sprintf("s3://%v/transcribe/", id), nil
				}).(pulumi.StringOutput),
			},
			LanguageCode: pulumi.String("en-US"),
			Tags: pulumi.StringMap{
				"ENVIRONMENT": pulumi.String("development"),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.s3.BucketObjectv2;
import com.pulumi.aws.s3.BucketObjectv2Args;
import com.pulumi.aws.transcribe.LanguageModel;
import com.pulumi.aws.transcribe.LanguageModelArgs;
import com.pulumi.aws.transcribe.inputs.LanguageModelInputDataConfigArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("transcribe.amazonaws.com")
                    .build())
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var testPolicy = new RolePolicy("testPolicy", RolePolicyArgs.builder()        
            .role(exampleRole.id())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", jsonArray(
                            "s3:GetObject", 
                            "s3:ListBucket"
                        )),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Resource", jsonArray("*"))
                    )))
                )))
            .build());

        var exampleBucketV2 = new BucketV2("exampleBucketV2", BucketV2Args.builder()        
            .forceDestroy(true)
            .build());

        var object = new BucketObjectv2("object", BucketObjectv2Args.builder()        
            .bucket(exampleBucketV2.id())
            .key("transcribe/test1.txt")
            .source(new FileAsset("test1.txt"))
            .build());

        var exampleLanguageModel = new LanguageModel("exampleLanguageModel", LanguageModelArgs.builder()        
            .modelName("example")
            .baseModelName("NarrowBand")
            .inputDataConfig(LanguageModelInputDataConfigArgs.builder()
                .dataAccessRoleArn(exampleRole.arn())
                .s3Uri(exampleBucketV2.id().applyValue(id -> String.format("s3://%s/transcribe/", id)))
                .build())
            .languageCode("en-US")
            .tags(Map.of("ENVIRONMENT", "development"))
            .build());

    }
}
```
```yaml
resources:
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${examplePolicyDocument.json}
  testPolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${exampleRole.id}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action:
                - s3:GetObject
                - s3:ListBucket
              Effect: Allow
              Resource:
                - '*'
  exampleBucketV2:
    type: aws:s3:BucketV2
    properties:
      forceDestroy: true
  object:
    type: aws:s3:BucketObjectv2
    properties:
      bucket: ${exampleBucketV2.id}
      key: transcribe/test1.txt
      source:
        fn::FileAsset: test1.txt
  exampleLanguageModel:
    type: aws:transcribe:LanguageModel
    properties:
      modelName: example
      baseModelName: NarrowBand
      inputDataConfig:
        dataAccessRoleArn: ${exampleRole.arn}
        s3Uri: s3://${exampleBucketV2.id}/transcribe/
      languageCode: en-US
      tags:
        ENVIRONMENT: development
variables:
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            principals:
              - type: Service
                identifiers:
                  - transcribe.amazonaws.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transcribe LanguageModel using the `model_name`. For example:

```sh
 $ pulumi import aws:transcribe/languageModel:LanguageModel example example-name
```
 