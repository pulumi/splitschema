Provides a CodePipeline.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codestarconnections.Connection("example", {providerType: "GitHub"});
const codepipelineBucket = new aws.s3.BucketV2("codepipelineBucket", {});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["codepipeline.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const codepipelineRole = new aws.iam.Role("codepipelineRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const s3kmskey = aws.kms.getAlias({
    name: "alias/myKmsKey",
});
const codepipeline = new aws.codepipeline.Pipeline("codepipeline", {
    roleArn: codepipelineRole.arn,
    artifactStores: [{
        location: codepipelineBucket.bucket,
        type: "S3",
        encryptionKey: {
            id: s3kmskey.then(s3kmskey => s3kmskey.arn),
            type: "KMS",
        },
    }],
    stages: [
        {
            name: "Source",
            actions: [{
                name: "Source",
                category: "Source",
                owner: "AWS",
                provider: "CodeStarSourceConnection",
                version: "1",
                outputArtifacts: ["source_output"],
                configuration: {
                    ConnectionArn: example.arn,
                    FullRepositoryId: "my-organization/example",
                    BranchName: "main",
                },
            }],
        },
        {
            name: "Build",
            actions: [{
                name: "Build",
                category: "Build",
                owner: "AWS",
                provider: "CodeBuild",
                inputArtifacts: ["source_output"],
                outputArtifacts: ["build_output"],
                version: "1",
                configuration: {
                    ProjectName: "test",
                },
            }],
        },
        {
            name: "Deploy",
            actions: [{
                name: "Deploy",
                category: "Deploy",
                owner: "AWS",
                provider: "CloudFormation",
                inputArtifacts: ["build_output"],
                version: "1",
                configuration: {
                    ActionMode: "REPLACE_ON_FAILURE",
                    Capabilities: "CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM",
                    OutputFileName: "CreateStackOutput.json",
                    StackName: "MyStack",
                    TemplatePath: "build_output::sam-templated.yaml",
                },
            }],
        },
    ],
});
const codepipelineBucketAcl = new aws.s3.BucketAclV2("codepipelineBucketAcl", {
    bucket: codepipelineBucket.id,
    acl: "private",
});
const codepipelinePolicyPolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [
        {
            effect: "Allow",
            actions: [
                "s3:GetObject",
                "s3:GetObjectVersion",
                "s3:GetBucketVersioning",
                "s3:PutObjectAcl",
                "s3:PutObject",
            ],
            resources: [
                codepipelineBucket.arn,
                pulumi.interpolate`${codepipelineBucket.arn}/*`,
            ],
        },
        {
            effect: "Allow",
            actions: ["codestar-connections:UseConnection"],
            resources: [example.arn],
        },
        {
            effect: "Allow",
            actions: [
                "codebuild:BatchGetBuilds",
                "codebuild:StartBuild",
            ],
            resources: ["*"],
        },
    ],
});
const codepipelinePolicyRolePolicy = new aws.iam.RolePolicy("codepipelinePolicyRolePolicy", {
    role: codepipelineRole.id,
    policy: codepipelinePolicyPolicyDocument.apply(codepipelinePolicyPolicyDocument => codepipelinePolicyPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codestarconnections.Connection("example", provider_type="GitHub")
codepipeline_bucket = aws.s3.BucketV2("codepipelineBucket")
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["codepipeline.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
codepipeline_role = aws.iam.Role("codepipelineRole", assume_role_policy=assume_role.json)
s3kmskey = aws.kms.get_alias(name="alias/myKmsKey")
codepipeline = aws.codepipeline.Pipeline("codepipeline",
    role_arn=codepipeline_role.arn,
    artifact_stores=[aws.codepipeline.PipelineArtifactStoreArgs(
        location=codepipeline_bucket.bucket,
        type="S3",
        encryption_key=aws.codepipeline.PipelineArtifactStoreEncryptionKeyArgs(
            id=s3kmskey.arn,
            type="KMS",
        ),
    )],
    stages=[
        aws.codepipeline.PipelineStageArgs(
            name="Source",
            actions=[aws.codepipeline.PipelineStageActionArgs(
                name="Source",
                category="Source",
                owner="AWS",
                provider="CodeStarSourceConnection",
                version="1",
                output_artifacts=["source_output"],
                configuration={
                    "ConnectionArn": example.arn,
                    "FullRepositoryId": "my-organization/example",
                    "BranchName": "main",
                },
            )],
        ),
        aws.codepipeline.PipelineStageArgs(
            name="Build",
            actions=[aws.codepipeline.PipelineStageActionArgs(
                name="Build",
                category="Build",
                owner="AWS",
                provider="CodeBuild",
                input_artifacts=["source_output"],
                output_artifacts=["build_output"],
                version="1",
                configuration={
                    "ProjectName": "test",
                },
            )],
        ),
        aws.codepipeline.PipelineStageArgs(
            name="Deploy",
            actions=[aws.codepipeline.PipelineStageActionArgs(
                name="Deploy",
                category="Deploy",
                owner="AWS",
                provider="CloudFormation",
                input_artifacts=["build_output"],
                version="1",
                configuration={
                    "ActionMode": "REPLACE_ON_FAILURE",
                    "Capabilities": "CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM",
                    "OutputFileName": "CreateStackOutput.json",
                    "StackName": "MyStack",
                    "TemplatePath": "build_output::sam-templated.yaml",
                },
            )],
        ),
    ])
codepipeline_bucket_acl = aws.s3.BucketAclV2("codepipelineBucketAcl",
    bucket=codepipeline_bucket.id,
    acl="private")
codepipeline_policy_policy_document = aws.iam.get_policy_document_output(statements=[
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=[
            "s3:GetObject",
            "s3:GetObjectVersion",
            "s3:GetBucketVersioning",
            "s3:PutObjectAcl",
            "s3:PutObject",
        ],
        resources=[
            codepipeline_bucket.arn,
            codepipeline_bucket.arn.apply(lambda arn: f"{arn}/*"),
        ],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=["codestar-connections:UseConnection"],
        resources=[example.arn],
    ),
    aws.iam.GetPolicyDocumentStatementArgs(
        effect="Allow",
        actions=[
            "codebuild:BatchGetBuilds",
            "codebuild:StartBuild",
        ],
        resources=["*"],
    ),
])
codepipeline_policy_role_policy = aws.iam.RolePolicy("codepipelinePolicyRolePolicy",
    role=codepipeline_role.id,
    policy=codepipeline_policy_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeStarConnections.Connection("example", new()
    {
        ProviderType = "GitHub",
    });

    var codepipelineBucket = new Aws.S3.BucketV2("codepipelineBucket");

    var assumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "codepipeline.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var codepipelineRole = new Aws.Iam.Role("codepipelineRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var s3kmskey = Aws.Kms.GetAlias.Invoke(new()
    {
        Name = "alias/myKmsKey",
    });

    var codepipeline = new Aws.CodePipeline.Pipeline("codepipeline", new()
    {
        RoleArn = codepipelineRole.Arn,
        ArtifactStores = new[]
        {
            new Aws.CodePipeline.Inputs.PipelineArtifactStoreArgs
            {
                Location = codepipelineBucket.Bucket,
                Type = "S3",
                EncryptionKey = new Aws.CodePipeline.Inputs.PipelineArtifactStoreEncryptionKeyArgs
                {
                    Id = s3kmskey.Apply(getAliasResult => getAliasResult.Arn),
                    Type = "KMS",
                },
            },
        },
        Stages = new[]
        {
            new Aws.CodePipeline.Inputs.PipelineStageArgs
            {
                Name = "Source",
                Actions = new[]
                {
                    new Aws.CodePipeline.Inputs.PipelineStageActionArgs
                    {
                        Name = "Source",
                        Category = "Source",
                        Owner = "AWS",
                        Provider = "CodeStarSourceConnection",
                        Version = "1",
                        OutputArtifacts = new[]
                        {
                            "source_output",
                        },
                        Configuration = 
                        {
                            { "ConnectionArn", example.Arn },
                            { "FullRepositoryId", "my-organization/example" },
                            { "BranchName", "main" },
                        },
                    },
                },
            },
            new Aws.CodePipeline.Inputs.PipelineStageArgs
            {
                Name = "Build",
                Actions = new[]
                {
                    new Aws.CodePipeline.Inputs.PipelineStageActionArgs
                    {
                        Name = "Build",
                        Category = "Build",
                        Owner = "AWS",
                        Provider = "CodeBuild",
                        InputArtifacts = new[]
                        {
                            "source_output",
                        },
                        OutputArtifacts = new[]
                        {
                            "build_output",
                        },
                        Version = "1",
                        Configuration = 
                        {
                            { "ProjectName", "test" },
                        },
                    },
                },
            },
            new Aws.CodePipeline.Inputs.PipelineStageArgs
            {
                Name = "Deploy",
                Actions = new[]
                {
                    new Aws.CodePipeline.Inputs.PipelineStageActionArgs
                    {
                        Name = "Deploy",
                        Category = "Deploy",
                        Owner = "AWS",
                        Provider = "CloudFormation",
                        InputArtifacts = new[]
                        {
                            "build_output",
                        },
                        Version = "1",
                        Configuration = 
                        {
                            { "ActionMode", "REPLACE_ON_FAILURE" },
                            { "Capabilities", "CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM" },
                            { "OutputFileName", "CreateStackOutput.json" },
                            { "StackName", "MyStack" },
                            { "TemplatePath", "build_output::sam-templated.yaml" },
                        },
                    },
                },
            },
        },
    });

    var codepipelineBucketAcl = new Aws.S3.BucketAclV2("codepipelineBucketAcl", new()
    {
        Bucket = codepipelineBucket.Id,
        Acl = "private",
    });

    var codepipelinePolicyPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "s3:GetObject",
                    "s3:GetObjectVersion",
                    "s3:GetBucketVersioning",
                    "s3:PutObjectAcl",
                    "s3:PutObject",
                },
                Resources = new[]
                {
                    codepipelineBucket.Arn,
                    $"{codepipelineBucket.Arn}/*",
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "codestar-connections:UseConnection",
                },
                Resources = new[]
                {
                    example.Arn,
                },
            },
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "codebuild:BatchGetBuilds",
                    "codebuild:StartBuild",
                },
                Resources = new[]
                {
                    "*",
                },
            },
        },
    });

    var codepipelinePolicyRolePolicy = new Aws.Iam.RolePolicy("codepipelinePolicyRolePolicy", new()
    {
        Role = codepipelineRole.Id,
        Policy = codepipelinePolicyPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codepipeline"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codestarconnections"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := codestarconnections.NewConnection(ctx, "example", &codestarconnections.ConnectionArgs{
			ProviderType: pulumi.String("GitHub"),
		})
		if err != nil {
			return err
		}
		codepipelineBucket, err := s3.NewBucketV2(ctx, "codepipelineBucket", nil)
		if err != nil {
			return err
		}
		assumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"codepipeline.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		codepipelineRole, err := iam.NewRole(ctx, "codepipelineRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		s3kmskey, err := kms.LookupAlias(ctx, &kms.LookupAliasArgs{
			Name: "alias/myKmsKey",
		}, nil)
		if err != nil {
			return err
		}
		_, err = codepipeline.NewPipeline(ctx, "codepipeline", &codepipeline.PipelineArgs{
			RoleArn: codepipelineRole.Arn,
			ArtifactStores: codepipeline.PipelineArtifactStoreArray{
				&codepipeline.PipelineArtifactStoreArgs{
					Location: codepipelineBucket.Bucket,
					Type:     pulumi.String("S3"),
					EncryptionKey: &codepipeline.PipelineArtifactStoreEncryptionKeyArgs{
						Id:   *pulumi.String(s3kmskey.Arn),
						Type: pulumi.String("KMS"),
					},
				},
			},
			Stages: codepipeline.PipelineStageArray{
				&codepipeline.PipelineStageArgs{
					Name: pulumi.String("Source"),
					Actions: codepipeline.PipelineStageActionArray{
						&codepipeline.PipelineStageActionArgs{
							Name:     pulumi.String("Source"),
							Category: pulumi.String("Source"),
							Owner:    pulumi.String("AWS"),
							Provider: pulumi.String("CodeStarSourceConnection"),
							Version:  pulumi.String("1"),
							OutputArtifacts: pulumi.StringArray{
								pulumi.String("source_output"),
							},
							Configuration: pulumi.StringMap{
								"ConnectionArn":    example.Arn,
								"FullRepositoryId": pulumi.String("my-organization/example"),
								"BranchName":       pulumi.String("main"),
							},
						},
					},
				},
				&codepipeline.PipelineStageArgs{
					Name: pulumi.String("Build"),
					Actions: codepipeline.PipelineStageActionArray{
						&codepipeline.PipelineStageActionArgs{
							Name:     pulumi.String("Build"),
							Category: pulumi.String("Build"),
							Owner:    pulumi.String("AWS"),
							Provider: pulumi.String("CodeBuild"),
							InputArtifacts: pulumi.StringArray{
								pulumi.String("source_output"),
							},
							OutputArtifacts: pulumi.StringArray{
								pulumi.String("build_output"),
							},
							Version: pulumi.String("1"),
							Configuration: pulumi.StringMap{
								"ProjectName": pulumi.String("test"),
							},
						},
					},
				},
				&codepipeline.PipelineStageArgs{
					Name: pulumi.String("Deploy"),
					Actions: codepipeline.PipelineStageActionArray{
						&codepipeline.PipelineStageActionArgs{
							Name:     pulumi.String("Deploy"),
							Category: pulumi.String("Deploy"),
							Owner:    pulumi.String("AWS"),
							Provider: pulumi.String("CloudFormation"),
							InputArtifacts: pulumi.StringArray{
								pulumi.String("build_output"),
							},
							Version: pulumi.String("1"),
							Configuration: pulumi.StringMap{
								"ActionMode":     pulumi.String("REPLACE_ON_FAILURE"),
								"Capabilities":   pulumi.String("CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM"),
								"OutputFileName": pulumi.String("CreateStackOutput.json"),
								"StackName":      pulumi.String("MyStack"),
								"TemplatePath":   pulumi.String("build_output::sam-templated.yaml"),
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "codepipelineBucketAcl", &s3.BucketAclV2Args{
			Bucket: codepipelineBucket.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		codepipelinePolicyPolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("s3:GetObject"),
						pulumi.String("s3:GetObjectVersion"),
						pulumi.String("s3:GetBucketVersioning"),
						pulumi.String("s3:PutObjectAcl"),
						pulumi.String("s3:PutObject"),
					},
					Resources: pulumi.StringArray{
						codepipelineBucket.Arn,
						codepipelineBucket.Arn.ApplyT(func(arn string) (string, error) {
							return fmt.Sprintf("%v/*", arn), nil
						}).(pulumi.StringOutput),
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("codestar-connections:UseConnection"),
					},
					Resources: pulumi.StringArray{
						example.Arn,
					},
				},
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("codebuild:BatchGetBuilds"),
						pulumi.String("codebuild:StartBuild"),
					},
					Resources: pulumi.StringArray{
						pulumi.String("*"),
					},
				},
			},
		}, nil)
		_, err = iam.NewRolePolicy(ctx, "codepipelinePolicyRolePolicy", &iam.RolePolicyArgs{
			Role: codepipelineRole.ID(),
			Policy: codepipelinePolicyPolicyDocument.ApplyT(func(codepipelinePolicyPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &codepipelinePolicyPolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
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
import com.pulumi.aws.codestarconnections.Connection;
import com.pulumi.aws.codestarconnections.ConnectionArgs;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.kms.KmsFunctions;
import com.pulumi.aws.kms.inputs.GetAliasArgs;
import com.pulumi.aws.codepipeline.Pipeline;
import com.pulumi.aws.codepipeline.PipelineArgs;
import com.pulumi.aws.codepipeline.inputs.PipelineArtifactStoreArgs;
import com.pulumi.aws.codepipeline.inputs.PipelineArtifactStoreEncryptionKeyArgs;
import com.pulumi.aws.codepipeline.inputs.PipelineStageArgs;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
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
            .providerType("GitHub")
            .build());

        var codepipelineBucket = new BucketV2("codepipelineBucket");

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("codepipeline.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var codepipelineRole = new Role("codepipelineRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        final var s3kmskey = KmsFunctions.getAlias(GetAliasArgs.builder()
            .name("alias/myKmsKey")
            .build());

        var codepipeline = new Pipeline("codepipeline", PipelineArgs.builder()        
            .roleArn(codepipelineRole.arn())
            .artifactStores(PipelineArtifactStoreArgs.builder()
                .location(codepipelineBucket.bucket())
                .type("S3")
                .encryptionKey(PipelineArtifactStoreEncryptionKeyArgs.builder()
                    .id(s3kmskey.applyValue(getAliasResult -> getAliasResult.arn()))
                    .type("KMS")
                    .build())
                .build())
            .stages(            
                PipelineStageArgs.builder()
                    .name("Source")
                    .actions(PipelineStageActionArgs.builder()
                        .name("Source")
                        .category("Source")
                        .owner("AWS")
                        .provider("CodeStarSourceConnection")
                        .version("1")
                        .outputArtifacts("source_output")
                        .configuration(Map.ofEntries(
                            Map.entry("ConnectionArn", example.arn()),
                            Map.entry("FullRepositoryId", "my-organization/example"),
                            Map.entry("BranchName", "main")
                        ))
                        .build())
                    .build(),
                PipelineStageArgs.builder()
                    .name("Build")
                    .actions(PipelineStageActionArgs.builder()
                        .name("Build")
                        .category("Build")
                        .owner("AWS")
                        .provider("CodeBuild")
                        .inputArtifacts("source_output")
                        .outputArtifacts("build_output")
                        .version("1")
                        .configuration(Map.of("ProjectName", "test"))
                        .build())
                    .build(),
                PipelineStageArgs.builder()
                    .name("Deploy")
                    .actions(PipelineStageActionArgs.builder()
                        .name("Deploy")
                        .category("Deploy")
                        .owner("AWS")
                        .provider("CloudFormation")
                        .inputArtifacts("build_output")
                        .version("1")
                        .configuration(Map.ofEntries(
                            Map.entry("ActionMode", "REPLACE_ON_FAILURE"),
                            Map.entry("Capabilities", "CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM"),
                            Map.entry("OutputFileName", "CreateStackOutput.json"),
                            Map.entry("StackName", "MyStack"),
                            Map.entry("TemplatePath", "build_output::sam-templated.yaml")
                        ))
                        .build())
                    .build())
            .build());

        var codepipelineBucketAcl = new BucketAclV2("codepipelineBucketAcl", BucketAclV2Args.builder()        
            .bucket(codepipelineBucket.id())
            .acl("private")
            .build());

        final var codepipelinePolicyPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(            
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions(                    
                        "s3:GetObject",
                        "s3:GetObjectVersion",
                        "s3:GetBucketVersioning",
                        "s3:PutObjectAcl",
                        "s3:PutObject")
                    .resources(                    
                        codepipelineBucket.arn(),
                        codepipelineBucket.arn().applyValue(arn -> String.format("%s/*", arn)))
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions("codestar-connections:UseConnection")
                    .resources(example.arn())
                    .build(),
                GetPolicyDocumentStatementArgs.builder()
                    .effect("Allow")
                    .actions(                    
                        "codebuild:BatchGetBuilds",
                        "codebuild:StartBuild")
                    .resources("*")
                    .build())
            .build());

        var codepipelinePolicyRolePolicy = new RolePolicy("codepipelinePolicyRolePolicy", RolePolicyArgs.builder()        
            .role(codepipelineRole.id())
            .policy(codepipelinePolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(codepipelinePolicyPolicyDocument -> codepipelinePolicyPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  codepipeline:
    type: aws:codepipeline:Pipeline
    properties:
      roleArn: ${codepipelineRole.arn}
      artifactStores:
        - location: ${codepipelineBucket.bucket}
          type: S3
          encryptionKey:
            id: ${s3kmskey.arn}
            type: KMS
      stages:
        - name: Source
          actions:
            - name: Source
              category: Source
              owner: AWS
              provider: CodeStarSourceConnection
              version: '1'
              outputArtifacts:
                - source_output
              configuration:
                ConnectionArn: ${example.arn}
                FullRepositoryId: my-organization/example
                BranchName: main
        - name: Build
          actions:
            - name: Build
              category: Build
              owner: AWS
              provider: CodeBuild
              inputArtifacts:
                - source_output
              outputArtifacts:
                - build_output
              version: '1'
              configuration:
                ProjectName: test
        - name: Deploy
          actions:
            - name: Deploy
              category: Deploy
              owner: AWS
              provider: CloudFormation
              inputArtifacts:
                - build_output
              version: '1'
              configuration:
                ActionMode: REPLACE_ON_FAILURE
                Capabilities: CAPABILITY_AUTO_EXPAND,CAPABILITY_IAM
                OutputFileName: CreateStackOutput.json
                StackName: MyStack
                TemplatePath: build_output::sam-templated.yaml
  example:
    type: aws:codestarconnections:Connection
    properties:
      providerType: GitHub
  codepipelineBucket:
    type: aws:s3:BucketV2
  codepipelineBucketAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${codepipelineBucket.id}
      acl: private
  codepipelineRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  codepipelinePolicyRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${codepipelineRole.id}
      policy: ${codepipelinePolicyPolicyDocument.json}
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - codepipeline.amazonaws.com
            actions:
              - sts:AssumeRole
  codepipelinePolicyPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - s3:GetObject
              - s3:GetObjectVersion
              - s3:GetBucketVersioning
              - s3:PutObjectAcl
              - s3:PutObject
            resources:
              - ${codepipelineBucket.arn}
              - ${codepipelineBucket.arn}/*
          - effect: Allow
            actions:
              - codestar-connections:UseConnection
            resources:
              - ${example.arn}
          - effect: Allow
            actions:
              - codebuild:BatchGetBuilds
              - codebuild:StartBuild
            resources:
              - '*'
  s3kmskey:
    fn::invoke:
      Function: aws:kms:getAlias
      Arguments:
        name: alias/myKmsKey
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodePipelines using the name. For example:

```sh
 $ pulumi import aws:codepipeline/pipeline:Pipeline foo example
```
 