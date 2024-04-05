Provides an AppSync Data Source.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleTable = new aws.dynamodb.Table("exampleTable", {
    readCapacity: 1,
    writeCapacity: 1,
    hashKey: "UserId",
    attributes: [{
        name: "UserId",
        type: "S",
    }],
});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["appsync.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const exampleRole = new aws.iam.Role("exampleRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const examplePolicyDocument = aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        actions: ["dynamodb:*"],
        resources: [exampleTable.arn],
    }],
});
const exampleRolePolicy = new aws.iam.RolePolicy("exampleRolePolicy", {
    role: exampleRole.id,
    policy: examplePolicyDocument.apply(examplePolicyDocument => examplePolicyDocument.json),
});
const exampleGraphQLApi = new aws.appsync.GraphQLApi("exampleGraphQLApi", {authenticationType: "API_KEY"});
const exampleDataSource = new aws.appsync.DataSource("exampleDataSource", {
    apiId: exampleGraphQLApi.id,
    name: "my_appsync_example",
    serviceRoleArn: exampleRole.arn,
    type: "AMAZON_DYNAMODB",
    dynamodbConfig: {
        tableName: exampleTable.name,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_table = aws.dynamodb.Table("exampleTable",
    read_capacity=1,
    write_capacity=1,
    hash_key="UserId",
    attributes=[aws.dynamodb.TableAttributeArgs(
        name="UserId",
        type="S",
    )])
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["appsync.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
example_role = aws.iam.Role("exampleRole", assume_role_policy=assume_role.json)
example_policy_document = aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    actions=["dynamodb:*"],
    resources=[example_table.arn],
)])
example_role_policy = aws.iam.RolePolicy("exampleRolePolicy",
    role=example_role.id,
    policy=example_policy_document.json)
example_graph_ql_api = aws.appsync.GraphQLApi("exampleGraphQLApi", authentication_type="API_KEY")
example_data_source = aws.appsync.DataSource("exampleDataSource",
    api_id=example_graph_ql_api.id,
    name="my_appsync_example",
    service_role_arn=example_role.arn,
    type="AMAZON_DYNAMODB",
    dynamodb_config=aws.appsync.DataSourceDynamodbConfigArgs(
        table_name=example_table.name,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleTable = new Aws.DynamoDB.Table("exampleTable", new()
    {
        ReadCapacity = 1,
        WriteCapacity = 1,
        HashKey = "UserId",
        Attributes = new[]
        {
            new Aws.DynamoDB.Inputs.TableAttributeArgs
            {
                Name = "UserId",
                Type = "S",
            },
        },
    });

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
                            "appsync.amazonaws.com",
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

    var exampleRole = new Aws.Iam.Role("exampleRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var examplePolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Actions = new[]
                {
                    "dynamodb:*",
                },
                Resources = new[]
                {
                    exampleTable.Arn,
                },
            },
        },
    });

    var exampleRolePolicy = new Aws.Iam.RolePolicy("exampleRolePolicy", new()
    {
        Role = exampleRole.Id,
        Policy = examplePolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var exampleGraphQLApi = new Aws.AppSync.GraphQLApi("exampleGraphQLApi", new()
    {
        AuthenticationType = "API_KEY",
    });

    var exampleDataSource = new Aws.AppSync.DataSource("exampleDataSource", new()
    {
        ApiId = exampleGraphQLApi.Id,
        Name = "my_appsync_example",
        ServiceRoleArn = exampleRole.Arn,
        Type = "AMAZON_DYNAMODB",
        DynamodbConfig = new Aws.AppSync.Inputs.DataSourceDynamodbConfigArgs
        {
            TableName = exampleTable.Name,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appsync"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleTable, err := dynamodb.NewTable(ctx, "exampleTable", &dynamodb.TableArgs{
			ReadCapacity:  pulumi.Int(1),
			WriteCapacity: pulumi.Int(1),
			HashKey:       pulumi.String("UserId"),
			Attributes: dynamodb.TableAttributeArray{
				&dynamodb.TableAttributeArgs{
					Name: pulumi.String("UserId"),
					Type: pulumi.String("S"),
				},
			},
		})
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
								"appsync.amazonaws.com",
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
		exampleRole, err := iam.NewRole(ctx, "exampleRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		examplePolicyDocument := iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
			Statements: iam.GetPolicyDocumentStatementArray{
				&iam.GetPolicyDocumentStatementArgs{
					Effect: pulumi.String("Allow"),
					Actions: pulumi.StringArray{
						pulumi.String("dynamodb:*"),
					},
					Resources: pulumi.StringArray{
						exampleTable.Arn,
					},
				},
			},
		}, nil)
		_, err = iam.NewRolePolicy(ctx, "exampleRolePolicy", &iam.RolePolicyArgs{
			Role: exampleRole.ID(),
			Policy: examplePolicyDocument.ApplyT(func(examplePolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
				return &examplePolicyDocument.Json, nil
			}).(pulumi.StringPtrOutput),
		})
		if err != nil {
			return err
		}
		exampleGraphQLApi, err := appsync.NewGraphQLApi(ctx, "exampleGraphQLApi", &appsync.GraphQLApiArgs{
			AuthenticationType: pulumi.String("API_KEY"),
		})
		if err != nil {
			return err
		}
		_, err = appsync.NewDataSource(ctx, "exampleDataSource", &appsync.DataSourceArgs{
			ApiId:          exampleGraphQLApi.ID(),
			Name:           pulumi.String("my_appsync_example"),
			ServiceRoleArn: exampleRole.Arn,
			Type:           pulumi.String("AMAZON_DYNAMODB"),
			DynamodbConfig: &appsync.DataSourceDynamodbConfigArgs{
				TableName: exampleTable.Name,
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
import com.pulumi.aws.dynamodb.Table;
import com.pulumi.aws.dynamodb.TableArgs;
import com.pulumi.aws.dynamodb.inputs.TableAttributeArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicy;
import com.pulumi.aws.iam.RolePolicyArgs;
import com.pulumi.aws.appsync.GraphQLApi;
import com.pulumi.aws.appsync.GraphQLApiArgs;
import com.pulumi.aws.appsync.DataSource;
import com.pulumi.aws.appsync.DataSourceArgs;
import com.pulumi.aws.appsync.inputs.DataSourceDynamodbConfigArgs;
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
        var exampleTable = new Table("exampleTable", TableArgs.builder()        
            .readCapacity(1)
            .writeCapacity(1)
            .hashKey("UserId")
            .attributes(TableAttributeArgs.builder()
                .name("UserId")
                .type("S")
                .build())
            .build());

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("appsync.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var exampleRole = new Role("exampleRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        final var examplePolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .actions("dynamodb:*")
                .resources(exampleTable.arn())
                .build())
            .build());

        var exampleRolePolicy = new RolePolicy("exampleRolePolicy", RolePolicyArgs.builder()        
            .role(exampleRole.id())
            .policy(examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(examplePolicyDocument -> examplePolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

        var exampleGraphQLApi = new GraphQLApi("exampleGraphQLApi", GraphQLApiArgs.builder()        
            .authenticationType("API_KEY")
            .build());

        var exampleDataSource = new DataSource("exampleDataSource", DataSourceArgs.builder()        
            .apiId(exampleGraphQLApi.id())
            .name("my_appsync_example")
            .serviceRoleArn(exampleRole.arn())
            .type("AMAZON_DYNAMODB")
            .dynamodbConfig(DataSourceDynamodbConfigArgs.builder()
                .tableName(exampleTable.name())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleTable:
    type: aws:dynamodb:Table
    properties:
      readCapacity: 1
      writeCapacity: 1
      hashKey: UserId
      attributes:
        - name: UserId
          type: S
  exampleRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  exampleRolePolicy:
    type: aws:iam:RolePolicy
    properties:
      role: ${exampleRole.id}
      policy: ${examplePolicyDocument.json}
  exampleGraphQLApi:
    type: aws:appsync:GraphQLApi
    properties:
      authenticationType: API_KEY
  exampleDataSource:
    type: aws:appsync:DataSource
    properties:
      apiId: ${exampleGraphQLApi.id}
      name: my_appsync_example
      serviceRoleArn: ${exampleRole.arn}
      type: AMAZON_DYNAMODB
      dynamodbConfig:
        tableName: ${exampleTable.name}
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
                  - appsync.amazonaws.com
            actions:
              - sts:AssumeRole
  examplePolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            actions:
              - dynamodb:*
            resources:
              - ${exampleTable.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_appsync_datasource` using the `api_id`, a hyphen, and `name`. For example:

```sh
 $ pulumi import aws:appsync/dataSource:DataSource example abcdef123456-example
```
 