Gives an external source (like an EventBridge Rule, SNS, or S3) permission to access the Lambda function.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const iamForLambda = new aws.iam.Role("iamForLambda", {assumeRolePolicy: JSON.stringify({
    Version: "2012-10-17",
    Statement: [{
        Action: "sts:AssumeRole",
        Effect: "Allow",
        Sid: "",
        Principal: {
            Service: "lambda.amazonaws.com",
        },
    }],
})});
const testLambda = new aws.lambda.Function("testLambda", {
    code: new pulumi.asset.FileArchive("lambdatest.zip"),
    role: iamForLambda.arn,
    handler: "exports.handler",
    runtime: "nodejs16.x",
});
const testAlias = new aws.lambda.Alias("testAlias", {
    description: "a sample description",
    functionName: testLambda.name,
    functionVersion: "$LATEST",
});
const allowCloudwatch = new aws.lambda.Permission("allowCloudwatch", {
    action: "lambda:InvokeFunction",
    "function": testLambda.name,
    principal: "events.amazonaws.com",
    sourceArn: "arn:aws:events:eu-west-1:111122223333:rule/RunDaily",
    qualifier: testAlias.name,
});
```
```python
import pulumi
import json
import pulumi_aws as aws

iam_for_lambda = aws.iam.Role("iamForLambda", assume_role_policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Sid": "",
        "Principal": {
            "Service": "lambda.amazonaws.com",
        },
    }],
}))
test_lambda = aws.lambda_.Function("testLambda",
    code=pulumi.FileArchive("lambdatest.zip"),
    role=iam_for_lambda.arn,
    handler="exports.handler",
    runtime="nodejs16.x")
test_alias = aws.lambda_.Alias("testAlias",
    description="a sample description",
    function_name=test_lambda.name,
    function_version="$LATEST")
allow_cloudwatch = aws.lambda_.Permission("allowCloudwatch",
    action="lambda:InvokeFunction",
    function=test_lambda.name,
    principal="events.amazonaws.com",
    source_arn="arn:aws:events:eu-west-1:111122223333:rule/RunDaily",
    qualifier=test_alias.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var iamForLambda = new Aws.Iam.Role("iamForLambda", new()
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
                        ["Service"] = "lambda.amazonaws.com",
                    },
                },
            },
        }),
    });

    var testLambda = new Aws.Lambda.Function("testLambda", new()
    {
        Code = new FileArchive("lambdatest.zip"),
        Role = iamForLambda.Arn,
        Handler = "exports.handler",
        Runtime = "nodejs16.x",
    });

    var testAlias = new Aws.Lambda.Alias("testAlias", new()
    {
        Description = "a sample description",
        FunctionName = testLambda.Name,
        FunctionVersion = "$LATEST",
    });

    var allowCloudwatch = new Aws.Lambda.Permission("allowCloudwatch", new()
    {
        Action = "lambda:InvokeFunction",
        Function = testLambda.Name,
        Principal = "events.amazonaws.com",
        SourceArn = "arn:aws:events:eu-west-1:111122223333:rule/RunDaily",
        Qualifier = testAlias.Name,
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "sts:AssumeRole",
					"Effect": "Allow",
					"Sid":    "",
					"Principal": map[string]interface{}{
						"Service": "lambda.amazonaws.com",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		iamForLambda, err := iam.NewRole(ctx, "iamForLambda", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		testLambda, err := lambda.NewFunction(ctx, "testLambda", &lambda.FunctionArgs{
			Code:    pulumi.NewFileArchive("lambdatest.zip"),
			Role:    iamForLambda.Arn,
			Handler: pulumi.String("exports.handler"),
			Runtime: pulumi.String("nodejs16.x"),
		})
		if err != nil {
			return err
		}
		testAlias, err := lambda.NewAlias(ctx, "testAlias", &lambda.AliasArgs{
			Description:     pulumi.String("a sample description"),
			FunctionName:    testLambda.Name,
			FunctionVersion: pulumi.String("$LATEST"),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewPermission(ctx, "allowCloudwatch", &lambda.PermissionArgs{
			Action:    pulumi.String("lambda:InvokeFunction"),
			Function:  testLambda.Name,
			Principal: pulumi.String("events.amazonaws.com"),
			SourceArn: pulumi.String("arn:aws:events:eu-west-1:111122223333:rule/RunDaily"),
			Qualifier: testAlias.Name,
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
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
import com.pulumi.aws.lambda.Alias;
import com.pulumi.aws.lambda.AliasArgs;
import com.pulumi.aws.lambda.Permission;
import com.pulumi.aws.lambda.PermissionArgs;
import static com.pulumi.codegen.internal.Serialization.*;
import com.pulumi.asset.FileArchive;
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
        var iamForLambda = new Role("iamForLambda", RoleArgs.builder()        
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", "sts:AssumeRole"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Sid", ""),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "lambda.amazonaws.com")
                        ))
                    )))
                )))
            .build());

        var testLambda = new Function("testLambda", FunctionArgs.builder()        
            .code(new FileArchive("lambdatest.zip"))
            .role(iamForLambda.arn())
            .handler("exports.handler")
            .runtime("nodejs16.x")
            .build());

        var testAlias = new Alias("testAlias", AliasArgs.builder()        
            .description("a sample description")
            .functionName(testLambda.name())
            .functionVersion("$LATEST")
            .build());

        var allowCloudwatch = new Permission("allowCloudwatch", PermissionArgs.builder()        
            .action("lambda:InvokeFunction")
            .function(testLambda.name())
            .principal("events.amazonaws.com")
            .sourceArn("arn:aws:events:eu-west-1:111122223333:rule/RunDaily")
            .qualifier(testAlias.name())
            .build());

    }
}
```
```yaml
resources:
  allowCloudwatch:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunction
      function: ${testLambda.name}
      principal: events.amazonaws.com
      sourceArn: arn:aws:events:eu-west-1:111122223333:rule/RunDaily
      qualifier: ${testAlias.name}
  testAlias:
    type: aws:lambda:Alias
    properties:
      description: a sample description
      functionName: ${testLambda.name}
      functionVersion: $LATEST
  testLambda:
    type: aws:lambda:Function
    properties:
      code:
        fn::FileArchive: lambdatest.zip
      role: ${iamForLambda.arn}
      handler: exports.handler
      runtime: nodejs16.x
  iamForLambda:
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
                Service: lambda.amazonaws.com
```
{{% /example %}}
{{% example %}}
### With SNS

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultTopic = new aws.sns.Topic("defaultTopic", {});
const defaultRole = new aws.iam.Role("defaultRole", {assumeRolePolicy: JSON.stringify({
    Version: "2012-10-17",
    Statement: [{
        Action: "sts:AssumeRole",
        Effect: "Allow",
        Sid: "",
        Principal: {
            Service: "lambda.amazonaws.com",
        },
    }],
})});
const func = new aws.lambda.Function("func", {
    code: new pulumi.asset.FileArchive("lambdatest.zip"),
    role: defaultRole.arn,
    handler: "exports.handler",
    runtime: "python3.7",
});
const withSns = new aws.lambda.Permission("withSns", {
    action: "lambda:InvokeFunction",
    "function": func.name,
    principal: "sns.amazonaws.com",
    sourceArn: defaultTopic.arn,
});
const lambda = new aws.sns.TopicSubscription("lambda", {
    topic: defaultTopic.arn,
    protocol: "lambda",
    endpoint: func.arn,
});
```
```python
import pulumi
import json
import pulumi_aws as aws

default_topic = aws.sns.Topic("defaultTopic")
default_role = aws.iam.Role("defaultRole", assume_role_policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Action": "sts:AssumeRole",
        "Effect": "Allow",
        "Sid": "",
        "Principal": {
            "Service": "lambda.amazonaws.com",
        },
    }],
}))
func = aws.lambda_.Function("func",
    code=pulumi.FileArchive("lambdatest.zip"),
    role=default_role.arn,
    handler="exports.handler",
    runtime="python3.7")
with_sns = aws.lambda_.Permission("withSns",
    action="lambda:InvokeFunction",
    function=func.name,
    principal="sns.amazonaws.com",
    source_arn=default_topic.arn)
lambda_ = aws.sns.TopicSubscription("lambda",
    topic=default_topic.arn,
    protocol="lambda",
    endpoint=func.arn)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultTopic = new Aws.Sns.Topic("defaultTopic");

    var defaultRole = new Aws.Iam.Role("defaultRole", new()
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
                        ["Service"] = "lambda.amazonaws.com",
                    },
                },
            },
        }),
    });

    var func = new Aws.Lambda.Function("func", new()
    {
        Code = new FileArchive("lambdatest.zip"),
        Role = defaultRole.Arn,
        Handler = "exports.handler",
        Runtime = "python3.7",
    });

    var withSns = new Aws.Lambda.Permission("withSns", new()
    {
        Action = "lambda:InvokeFunction",
        Function = func.Name,
        Principal = "sns.amazonaws.com",
        SourceArn = defaultTopic.Arn,
    });

    var lambda = new Aws.Sns.TopicSubscription("lambda", new()
    {
        Topic = defaultTopic.Arn,
        Protocol = "lambda",
        Endpoint = func.Arn,
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultTopic, err := sns.NewTopic(ctx, "defaultTopic", nil)
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
						"Service": "lambda.amazonaws.com",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		defaultRole, err := iam.NewRole(ctx, "defaultRole", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewFunction(ctx, "func", &lambda.FunctionArgs{
			Code:    pulumi.NewFileArchive("lambdatest.zip"),
			Role:    defaultRole.Arn,
			Handler: pulumi.String("exports.handler"),
			Runtime: pulumi.String("python3.7"),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewPermission(ctx, "withSns", &lambda.PermissionArgs{
			Action:    pulumi.String("lambda:InvokeFunction"),
			Function:  _func.Name,
			Principal: pulumi.String("sns.amazonaws.com"),
			SourceArn: defaultTopic.Arn,
		})
		if err != nil {
			return err
		}
		_, err = sns.NewTopicSubscription(ctx, "lambda", &sns.TopicSubscriptionArgs{
			Topic:    defaultTopic.Arn,
			Protocol: pulumi.String("lambda"),
			Endpoint: _func.Arn,
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
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
import com.pulumi.aws.lambda.Permission;
import com.pulumi.aws.lambda.PermissionArgs;
import com.pulumi.aws.sns.TopicSubscription;
import com.pulumi.aws.sns.TopicSubscriptionArgs;
import static com.pulumi.codegen.internal.Serialization.*;
import com.pulumi.asset.FileArchive;
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
        var defaultTopic = new Topic("defaultTopic");

        var defaultRole = new Role("defaultRole", RoleArgs.builder()        
            .assumeRolePolicy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", "sts:AssumeRole"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Sid", ""),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("Service", "lambda.amazonaws.com")
                        ))
                    )))
                )))
            .build());

        var func = new Function("func", FunctionArgs.builder()        
            .code(new FileArchive("lambdatest.zip"))
            .role(defaultRole.arn())
            .handler("exports.handler")
            .runtime("python3.7")
            .build());

        var withSns = new Permission("withSns", PermissionArgs.builder()        
            .action("lambda:InvokeFunction")
            .function(func.name())
            .principal("sns.amazonaws.com")
            .sourceArn(defaultTopic.arn())
            .build());

        var lambda = new TopicSubscription("lambda", TopicSubscriptionArgs.builder()        
            .topic(defaultTopic.arn())
            .protocol("lambda")
            .endpoint(func.arn())
            .build());

    }
}
```
```yaml
resources:
  withSns:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunction
      function: ${func.name}
      principal: sns.amazonaws.com
      sourceArn: ${defaultTopic.arn}
  defaultTopic:
    type: aws:sns:Topic
  lambda:
    type: aws:sns:TopicSubscription
    properties:
      topic: ${defaultTopic.arn}
      protocol: lambda
      endpoint: ${func.arn}
  func:
    type: aws:lambda:Function
    properties:
      code:
        fn::FileArchive: lambdatest.zip
      role: ${defaultRole.arn}
      handler: exports.handler
      runtime: python3.7
  defaultRole:
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
                Service: lambda.amazonaws.com
```
{{% /example %}}
{{% example %}}
### With API Gateway REST API

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDemoAPI = new aws.apigateway.RestApi("myDemoAPI", {description: "This is my API for demonstration purposes"});
const lambdaPermission = new aws.lambda.Permission("lambdaPermission", {
    action: "lambda:InvokeFunction",
    "function": "MyDemoFunction",
    principal: "apigateway.amazonaws.com",
    sourceArn: pulumi.interpolate`${myDemoAPI.executionArn}/*`,
});
```
```python
import pulumi
import pulumi_aws as aws

my_demo_api = aws.apigateway.RestApi("myDemoAPI", description="This is my API for demonstration purposes")
lambda_permission = aws.lambda_.Permission("lambdaPermission",
    action="lambda:InvokeFunction",
    function="MyDemoFunction",
    principal="apigateway.amazonaws.com",
    source_arn=my_demo_api.execution_arn.apply(lambda execution_arn: f"{execution_arn}/*"))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myDemoAPI = new Aws.ApiGateway.RestApi("myDemoAPI", new()
    {
        Description = "This is my API for demonstration purposes",
    });

    var lambdaPermission = new Aws.Lambda.Permission("lambdaPermission", new()
    {
        Action = "lambda:InvokeFunction",
        Function = "MyDemoFunction",
        Principal = "apigateway.amazonaws.com",
        SourceArn = myDemoAPI.ExecutionArn.Apply(executionArn => $"{executionArn}/*"),
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/apigateway"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myDemoAPI, err := apigateway.NewRestApi(ctx, "myDemoAPI", &apigateway.RestApiArgs{
			Description: pulumi.String("This is my API for demonstration purposes"),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewPermission(ctx, "lambdaPermission", &lambda.PermissionArgs{
			Action:    pulumi.String("lambda:InvokeFunction"),
			Function:  pulumi.Any("MyDemoFunction"),
			Principal: pulumi.String("apigateway.amazonaws.com"),
			SourceArn: myDemoAPI.ExecutionArn.ApplyT(func(executionArn string) (string, error) {
				return fmt.Sprintf("%v/*", executionArn), nil
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
import com.pulumi.aws.apigateway.RestApi;
import com.pulumi.aws.apigateway.RestApiArgs;
import com.pulumi.aws.lambda.Permission;
import com.pulumi.aws.lambda.PermissionArgs;
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
        var myDemoAPI = new RestApi("myDemoAPI", RestApiArgs.builder()        
            .description("This is my API for demonstration purposes")
            .build());

        var lambdaPermission = new Permission("lambdaPermission", PermissionArgs.builder()        
            .action("lambda:InvokeFunction")
            .function("MyDemoFunction")
            .principal("apigateway.amazonaws.com")
            .sourceArn(myDemoAPI.executionArn().applyValue(executionArn -> String.format("%s/*", executionArn)))
            .build());

    }
}
```
```yaml
resources:
  myDemoAPI:
    type: aws:apigateway:RestApi
    properties:
      description: This is my API for demonstration purposes
  lambdaPermission:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunction
      function: MyDemoFunction
      principal: apigateway.amazonaws.com
      # The /* part allows invocation from any stage, method and resource path
      #     // within API Gateway.
      sourceArn: ${myDemoAPI.executionArn}/*
```
{{% /example %}}
{{% example %}}
### With CloudWatch Log Group

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const defaultLogGroup = new aws.cloudwatch.LogGroup("defaultLogGroup", {});
const assumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["lambda.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const defaultRole = new aws.iam.Role("defaultRole", {assumeRolePolicy: assumeRole.then(assumeRole => assumeRole.json)});
const loggingFunction = new aws.lambda.Function("loggingFunction", {
    code: new pulumi.asset.FileArchive("lamba_logging.zip"),
    handler: "exports.handler",
    role: defaultRole.arn,
    runtime: "python3.7",
});
const loggingPermission = new aws.lambda.Permission("loggingPermission", {
    action: "lambda:InvokeFunction",
    "function": loggingFunction.name,
    principal: "logs.eu-west-1.amazonaws.com",
    sourceArn: pulumi.interpolate`${defaultLogGroup.arn}:*`,
});
const loggingLogSubscriptionFilter = new aws.cloudwatch.LogSubscriptionFilter("loggingLogSubscriptionFilter", {
    destinationArn: loggingFunction.arn,
    filterPattern: "",
    logGroup: defaultLogGroup.name,
}, {
    dependsOn: [loggingPermission],
});
```
```python
import pulumi
import pulumi_aws as aws

default_log_group = aws.cloudwatch.LogGroup("defaultLogGroup")
assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["lambda.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
default_role = aws.iam.Role("defaultRole", assume_role_policy=assume_role.json)
logging_function = aws.lambda_.Function("loggingFunction",
    code=pulumi.FileArchive("lamba_logging.zip"),
    handler="exports.handler",
    role=default_role.arn,
    runtime="python3.7")
logging_permission = aws.lambda_.Permission("loggingPermission",
    action="lambda:InvokeFunction",
    function=logging_function.name,
    principal="logs.eu-west-1.amazonaws.com",
    source_arn=default_log_group.arn.apply(lambda arn: f"{arn}:*"))
logging_log_subscription_filter = aws.cloudwatch.LogSubscriptionFilter("loggingLogSubscriptionFilter",
    destination_arn=logging_function.arn,
    filter_pattern="",
    log_group=default_log_group.name,
    opts=pulumi.ResourceOptions(depends_on=[logging_permission]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var defaultLogGroup = new Aws.CloudWatch.LogGroup("defaultLogGroup");

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
                            "lambda.amazonaws.com",
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

    var defaultRole = new Aws.Iam.Role("defaultRole", new()
    {
        AssumeRolePolicy = assumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var loggingFunction = new Aws.Lambda.Function("loggingFunction", new()
    {
        Code = new FileArchive("lamba_logging.zip"),
        Handler = "exports.handler",
        Role = defaultRole.Arn,
        Runtime = "python3.7",
    });

    var loggingPermission = new Aws.Lambda.Permission("loggingPermission", new()
    {
        Action = "lambda:InvokeFunction",
        Function = loggingFunction.Name,
        Principal = "logs.eu-west-1.amazonaws.com",
        SourceArn = defaultLogGroup.Arn.Apply(arn => $"{arn}:*"),
    });

    var loggingLogSubscriptionFilter = new Aws.CloudWatch.LogSubscriptionFilter("loggingLogSubscriptionFilter", new()
    {
        DestinationArn = loggingFunction.Arn,
        FilterPattern = "",
        LogGroup = defaultLogGroup.Name,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            loggingPermission,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		defaultLogGroup, err := cloudwatch.NewLogGroup(ctx, "defaultLogGroup", nil)
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
								"lambda.amazonaws.com",
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
		defaultRole, err := iam.NewRole(ctx, "defaultRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRole.Json),
		})
		if err != nil {
			return err
		}
		loggingFunction, err := lambda.NewFunction(ctx, "loggingFunction", &lambda.FunctionArgs{
			Code:    pulumi.NewFileArchive("lamba_logging.zip"),
			Handler: pulumi.String("exports.handler"),
			Role:    defaultRole.Arn,
			Runtime: pulumi.String("python3.7"),
		})
		if err != nil {
			return err
		}
		loggingPermission, err := lambda.NewPermission(ctx, "loggingPermission", &lambda.PermissionArgs{
			Action:    pulumi.String("lambda:InvokeFunction"),
			Function:  loggingFunction.Name,
			Principal: pulumi.String("logs.eu-west-1.amazonaws.com"),
			SourceArn: defaultLogGroup.Arn.ApplyT(func(arn string) (string, error) {
				return fmt.Sprintf("%v:*", arn), nil
			}).(pulumi.StringOutput),
		})
		if err != nil {
			return err
		}
		_, err = cloudwatch.NewLogSubscriptionFilter(ctx, "loggingLogSubscriptionFilter", &cloudwatch.LogSubscriptionFilterArgs{
			DestinationArn: loggingFunction.Arn,
			FilterPattern:  pulumi.String(""),
			LogGroup:       defaultLogGroup.Name,
		}, pulumi.DependsOn([]pulumi.Resource{
			loggingPermission,
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.lambda.Function;
import com.pulumi.aws.lambda.FunctionArgs;
import com.pulumi.aws.lambda.Permission;
import com.pulumi.aws.lambda.PermissionArgs;
import com.pulumi.aws.cloudwatch.LogSubscriptionFilter;
import com.pulumi.aws.cloudwatch.LogSubscriptionFilterArgs;
import com.pulumi.resources.CustomResourceOptions;
import com.pulumi.asset.FileArchive;
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
        var defaultLogGroup = new LogGroup("defaultLogGroup");

        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("lambda.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var defaultRole = new Role("defaultRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var loggingFunction = new Function("loggingFunction", FunctionArgs.builder()        
            .code(new FileArchive("lamba_logging.zip"))
            .handler("exports.handler")
            .role(defaultRole.arn())
            .runtime("python3.7")
            .build());

        var loggingPermission = new Permission("loggingPermission", PermissionArgs.builder()        
            .action("lambda:InvokeFunction")
            .function(loggingFunction.name())
            .principal("logs.eu-west-1.amazonaws.com")
            .sourceArn(defaultLogGroup.arn().applyValue(arn -> String.format("%s:*", arn)))
            .build());

        var loggingLogSubscriptionFilter = new LogSubscriptionFilter("loggingLogSubscriptionFilter", LogSubscriptionFilterArgs.builder()        
            .destinationArn(loggingFunction.arn())
            .filterPattern("")
            .logGroup(defaultLogGroup.name())
            .build(), CustomResourceOptions.builder()
                .dependsOn(loggingPermission)
                .build());

    }
}
```
```yaml
resources:
  loggingPermission:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunction
      function: ${loggingFunction.name}
      principal: logs.eu-west-1.amazonaws.com
      sourceArn: ${defaultLogGroup.arn}:*
  defaultLogGroup:
    type: aws:cloudwatch:LogGroup
  loggingLogSubscriptionFilter:
    type: aws:cloudwatch:LogSubscriptionFilter
    properties:
      destinationArn: ${loggingFunction.arn}
      filterPattern:
      logGroup: ${defaultLogGroup.name}
    options:
      dependson:
        - ${loggingPermission}
  loggingFunction:
    type: aws:lambda:Function
    properties:
      code:
        fn::FileArchive: lamba_logging.zip
      handler: exports.handler
      role: ${defaultRole.arn}
      runtime: python3.7
  defaultRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
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
                  - lambda.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% example %}}
### With Cross-Account Invocation Policy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const urlFunctionUrl = new aws.lambda.FunctionUrl("urlFunctionUrl", {
    functionName: aws_lambda_function.example.function_name,
    authorizationType: "AWS_IAM",
});
const urlPermission = new aws.lambda.Permission("urlPermission", {
    action: "lambda:InvokeFunctionUrl",
    "function": aws_lambda_function.example.function_name,
    principal: "arn:aws:iam::444455556666:role/example",
    sourceAccount: "444455556666",
    functionUrlAuthType: "AWS_IAM",
});
```
```python
import pulumi
import pulumi_aws as aws

url_function_url = aws.lambda_.FunctionUrl("urlFunctionUrl",
    function_name=aws_lambda_function["example"]["function_name"],
    authorization_type="AWS_IAM")
url_permission = aws.lambda_.Permission("urlPermission",
    action="lambda:InvokeFunctionUrl",
    function=aws_lambda_function["example"]["function_name"],
    principal="arn:aws:iam::444455556666:role/example",
    source_account="444455556666",
    function_url_auth_type="AWS_IAM")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var urlFunctionUrl = new Aws.Lambda.FunctionUrl("urlFunctionUrl", new()
    {
        FunctionName = aws_lambda_function.Example.Function_name,
        AuthorizationType = "AWS_IAM",
    });

    var urlPermission = new Aws.Lambda.Permission("urlPermission", new()
    {
        Action = "lambda:InvokeFunctionUrl",
        Function = aws_lambda_function.Example.Function_name,
        Principal = "arn:aws:iam::444455556666:role/example",
        SourceAccount = "444455556666",
        FunctionUrlAuthType = "AWS_IAM",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewFunctionUrl(ctx, "urlFunctionUrl", &lambda.FunctionUrlArgs{
			FunctionName:      pulumi.Any(aws_lambda_function.Example.Function_name),
			AuthorizationType: pulumi.String("AWS_IAM"),
		})
		if err != nil {
			return err
		}
		_, err = lambda.NewPermission(ctx, "urlPermission", &lambda.PermissionArgs{
			Action:              pulumi.String("lambda:InvokeFunctionUrl"),
			Function:            pulumi.Any(aws_lambda_function.Example.Function_name),
			Principal:           pulumi.String("arn:aws:iam::444455556666:role/example"),
			SourceAccount:       pulumi.String("444455556666"),
			FunctionUrlAuthType: pulumi.String("AWS_IAM"),
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
import com.pulumi.aws.lambda.FunctionUrl;
import com.pulumi.aws.lambda.FunctionUrlArgs;
import com.pulumi.aws.lambda.Permission;
import com.pulumi.aws.lambda.PermissionArgs;
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
        var urlFunctionUrl = new FunctionUrl("urlFunctionUrl", FunctionUrlArgs.builder()        
            .functionName(aws_lambda_function.example().function_name())
            .authorizationType("AWS_IAM")
            .build());

        var urlPermission = new Permission("urlPermission", PermissionArgs.builder()        
            .action("lambda:InvokeFunctionUrl")
            .function(aws_lambda_function.example().function_name())
            .principal("arn:aws:iam::444455556666:role/example")
            .sourceAccount("444455556666")
            .functionUrlAuthType("AWS_IAM")
            .build());

    }
}
```
```yaml
resources:
  urlFunctionUrl:
    type: aws:lambda:FunctionUrl
    properties:
      functionName: ${aws_lambda_function.example.function_name}
      authorizationType: AWS_IAM
  urlPermission:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunctionUrl
      function: ${aws_lambda_function.example.function_name}
      principal: arn:aws:iam::444455556666:role/example
      sourceAccount: '444455556666'
      functionUrlAuthType: AWS_IAM # Adds the following condition keys
      #   # "Condition": {
      #   #      "StringEquals": {
      #   #        "AWS:SourceAccount": "444455556666",
      #   #        "lambda:FunctionUrlAuthType": "AWS_IAM"
      #   #      }
      #   #    }
```
{{% /example %}}
{{% example %}}
### With `replace_triggered_by` Lifecycle Configuration

If omitting the `qualifier` argument (which forces re-creation each time a function version is published), a `lifecycle` block can be used to ensure permissions are re-applied on any change to the underlying function.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const logging = new aws.lambda.Permission("logging", {
    action: "lambda:InvokeFunction",
    "function": aws_lambda_function.example.function_name,
    principal: "events.amazonaws.com",
    sourceArn: "arn:aws:events:eu-west-1:111122223333:rule/RunDaily",
});
```
```python
import pulumi
import pulumi_aws as aws

logging = aws.lambda_.Permission("logging",
    action="lambda:InvokeFunction",
    function=aws_lambda_function["example"]["function_name"],
    principal="events.amazonaws.com",
    source_arn="arn:aws:events:eu-west-1:111122223333:rule/RunDaily")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var logging = new Aws.Lambda.Permission("logging", new()
    {
        Action = "lambda:InvokeFunction",
        Function = aws_lambda_function.Example.Function_name,
        Principal = "events.amazonaws.com",
        SourceArn = "arn:aws:events:eu-west-1:111122223333:rule/RunDaily",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewPermission(ctx, "logging", &lambda.PermissionArgs{
			Action:    pulumi.String("lambda:InvokeFunction"),
			Function:  pulumi.Any(aws_lambda_function.Example.Function_name),
			Principal: pulumi.String("events.amazonaws.com"),
			SourceArn: pulumi.String("arn:aws:events:eu-west-1:111122223333:rule/RunDaily"),
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
import com.pulumi.aws.lambda.Permission;
import com.pulumi.aws.lambda.PermissionArgs;
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
        var logging = new Permission("logging", PermissionArgs.builder()        
            .action("lambda:InvokeFunction")
            .function(aws_lambda_function.example().function_name())
            .principal("events.amazonaws.com")
            .sourceArn("arn:aws:events:eu-west-1:111122223333:rule/RunDaily")
            .build());

    }
}
```
```yaml
resources:
  logging:
    type: aws:lambda:Permission
    properties:
      action: lambda:InvokeFunction
      function: ${aws_lambda_function.example.function_name}
      principal: events.amazonaws.com
      sourceArn: arn:aws:events:eu-west-1:111122223333:rule/RunDaily
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lambda permission statements using function_name/statement_id with an optional qualifier. For example:

```sh
 $ pulumi import aws:lambda/permission:Permission test_lambda_permission my_test_lambda_function/AllowExecutionFromCloudWatch
```
 ```sh
 $ pulumi import aws:lambda/permission:Permission test_lambda_permission my_test_lambda_function:qualifier_name/AllowExecutionFromCloudWatch
```
 